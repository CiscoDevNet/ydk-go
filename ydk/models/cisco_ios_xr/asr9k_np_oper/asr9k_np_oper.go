// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-np package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-np: Hardware NP Counters
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_np_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_np_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-np-oper hardware-module-np}", reflect.TypeOf(HardwareModuleNp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-np-oper:hardware-module-np", reflect.TypeOf(HardwareModuleNp{}))
}

// HardwareModuleNp
// Hardware NP Counters
type HardwareModuleNp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes HardwareModuleNp_Nodes
}

func (hardwareModuleNp *HardwareModuleNp) GetEntityData() *types.CommonEntityData {
    hardwareModuleNp.EntityData.YFilter = hardwareModuleNp.YFilter
    hardwareModuleNp.EntityData.YangName = "hardware-module-np"
    hardwareModuleNp.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleNp.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-np-oper"
    hardwareModuleNp.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np"
    hardwareModuleNp.EntityData.AbsolutePath = hardwareModuleNp.EntityData.SegmentPath
    hardwareModuleNp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleNp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleNp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleNp.EntityData.Children = types.NewOrderedMap()
    hardwareModuleNp.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleNp.Nodes})
    hardwareModuleNp.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleNp.EntityData.YListKeys = []string {}

    return &(hardwareModuleNp.EntityData)
}

// HardwareModuleNp_Nodes
// Table of nodes
type HardwareModuleNp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of HardwareModuleNp_Nodes_Node.
    Node []*HardwareModuleNp_Nodes_Node
}

func (nodes *HardwareModuleNp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-np"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/" + nodes.EntityData.SegmentPath
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

// HardwareModuleNp_Nodes_Node
// Number
type HardwareModuleNp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of all NP.
    Nps HardwareModuleNp_Nodes_Node_Nps
}

func (node *HardwareModuleNp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("nps", types.YChild{"Nps", &node.Nps})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps
// List of all NP
type HardwareModuleNp_Nodes_Node_Nps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // np0 to np7. The type is slice of HardwareModuleNp_Nodes_Node_Nps_Np.
    Np []*HardwareModuleNp_Nodes_Node_Nps_Np
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetEntityData() *types.CommonEntityData {
    nps.EntityData.YFilter = nps.YFilter
    nps.EntityData.YangName = "nps"
    nps.EntityData.BundleName = "cisco_ios_xr"
    nps.EntityData.ParentYangName = "node"
    nps.EntityData.SegmentPath = "nps"
    nps.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/" + nps.EntityData.SegmentPath
    nps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nps.EntityData.Children = types.NewOrderedMap()
    nps.EntityData.Children.Append("np", types.YChild{"Np", nil})
    for i := range nps.Np {
        nps.EntityData.Children.Append(types.GetSegmentPath(nps.Np[i]), types.YChild{"Np", nps.Np[i]})
    }
    nps.EntityData.Leafs = types.NewOrderedMap()

    nps.EntityData.YListKeys = []string {}

    return &(nps.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np
// np0 to np7
type HardwareModuleNp_Nodes_Node_Nps_Np struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NP name. The type is string with pattern:
    // (np0)|(np1)|(np2)|(np3)|(np4)|(np5)|(np6)|(np7).
    NpName interface{}

    // prm channel load info.
    ChnLoad HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad

    // Hardware np load or utilization.
    LoadUtilization HardwareModuleNp_Nodes_Node_Nps_Np_LoadUtilization

    // prm tcam summary info.
    TcamSummary HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary

    // prm counters info.
    Counters HardwareModuleNp_Nodes_Node_Nps_Np_Counters

    // prm fast drop counters info.
    FastDrop HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop

    // EFD COS to Priority Mapping.
    Efd HardwareModuleNp_Nodes_Node_Nps_Np_Efd
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetEntityData() *types.CommonEntityData {
    np.EntityData.YFilter = np.YFilter
    np.EntityData.YangName = "np"
    np.EntityData.BundleName = "cisco_ios_xr"
    np.EntityData.ParentYangName = "nps"
    np.EntityData.SegmentPath = "np" + types.AddKeyToken(np.NpName, "np-name")
    np.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/" + np.EntityData.SegmentPath
    np.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    np.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    np.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    np.EntityData.Children = types.NewOrderedMap()
    np.EntityData.Children.Append("chn-load", types.YChild{"ChnLoad", &np.ChnLoad})
    np.EntityData.Children.Append("load-utilization", types.YChild{"LoadUtilization", &np.LoadUtilization})
    np.EntityData.Children.Append("tcam-summary", types.YChild{"TcamSummary", &np.TcamSummary})
    np.EntityData.Children.Append("counters", types.YChild{"Counters", &np.Counters})
    np.EntityData.Children.Append("fast-drop", types.YChild{"FastDrop", &np.FastDrop})
    np.EntityData.Children.Append("efd", types.YChild{"Efd", &np.Efd})
    np.EntityData.Leafs = types.NewOrderedMap()
    np.EntityData.Leafs.Append("np-name", types.YLeaf{"NpName", np.NpName})

    np.EntityData.YListKeys = []string {"NpName"}

    return &(np.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad
// prm channel load info
type HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of NP Channel load counters. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad.
    NpChnLoad []*HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetEntityData() *types.CommonEntityData {
    chnLoad.EntityData.YFilter = chnLoad.YFilter
    chnLoad.EntityData.YangName = "chn-load"
    chnLoad.EntityData.BundleName = "cisco_ios_xr"
    chnLoad.EntityData.ParentYangName = "np"
    chnLoad.EntityData.SegmentPath = "chn-load"
    chnLoad.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/" + chnLoad.EntityData.SegmentPath
    chnLoad.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chnLoad.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chnLoad.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chnLoad.EntityData.Children = types.NewOrderedMap()
    chnLoad.EntityData.Children.Append("np-chn-load", types.YChild{"NpChnLoad", nil})
    for i := range chnLoad.NpChnLoad {
        types.SetYListKey(chnLoad.NpChnLoad[i], i)
        chnLoad.EntityData.Children.Append(types.GetSegmentPath(chnLoad.NpChnLoad[i]), types.YChild{"NpChnLoad", chnLoad.NpChnLoad[i]})
    }
    chnLoad.EntityData.Leafs = types.NewOrderedMap()

    chnLoad.EntityData.YListKeys = []string {}

    return &(chnLoad.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad
// Array of NP Channel load counters
type HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Flow control counters. The type is interface{} with range: 0..4294967295.
    FlowCtrCounter interface{}

    // Average RFD Usage. The type is interface{} with range: 0..4294967295.
    AvgRfdUsage interface{}

    // Peak RFD Usage. The type is interface{} with range: 0..4294967295.
    PeakRfdUsage interface{}

    // Average of garanteed RFD usage. The type is interface{} with range:
    // 0..4294967295.
    AvgGuarRfdUsage interface{}

    // Peak of garanteed RFD usage. The type is interface{} with range:
    // 0..4294967295.
    PeakGuarRfdUsage interface{}

    // Inerface Name. The type is string.
    InterfaceName interface{}
}

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetEntityData() *types.CommonEntityData {
    npChnLoad.EntityData.YFilter = npChnLoad.YFilter
    npChnLoad.EntityData.YangName = "np-chn-load"
    npChnLoad.EntityData.BundleName = "cisco_ios_xr"
    npChnLoad.EntityData.ParentYangName = "chn-load"
    npChnLoad.EntityData.SegmentPath = "np-chn-load" + types.AddNoKeyToken(npChnLoad)
    npChnLoad.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/chn-load/" + npChnLoad.EntityData.SegmentPath
    npChnLoad.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npChnLoad.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npChnLoad.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npChnLoad.EntityData.Children = types.NewOrderedMap()
    npChnLoad.EntityData.Leafs = types.NewOrderedMap()
    npChnLoad.EntityData.Leafs.Append("flow-ctr-counter", types.YLeaf{"FlowCtrCounter", npChnLoad.FlowCtrCounter})
    npChnLoad.EntityData.Leafs.Append("avg-rfd-usage", types.YLeaf{"AvgRfdUsage", npChnLoad.AvgRfdUsage})
    npChnLoad.EntityData.Leafs.Append("peak-rfd-usage", types.YLeaf{"PeakRfdUsage", npChnLoad.PeakRfdUsage})
    npChnLoad.EntityData.Leafs.Append("avg-guar-rfd-usage", types.YLeaf{"AvgGuarRfdUsage", npChnLoad.AvgGuarRfdUsage})
    npChnLoad.EntityData.Leafs.Append("peak-guar-rfd-usage", types.YLeaf{"PeakGuarRfdUsage", npChnLoad.PeakGuarRfdUsage})
    npChnLoad.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", npChnLoad.InterfaceName})

    npChnLoad.EntityData.YListKeys = []string {}

    return &(npChnLoad.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_LoadUtilization
// Hardware np load or utilization
type HardwareModuleNp_Nodes_Node_Nps_Np_LoadUtilization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Percent load on the NP engines. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Utilization interface{}

    // Packet rate load on NP engines. The type is interface{} with range:
    // 0..4294967295.
    PacketPerSecondRate interface{}
}

func (loadUtilization *HardwareModuleNp_Nodes_Node_Nps_Np_LoadUtilization) GetEntityData() *types.CommonEntityData {
    loadUtilization.EntityData.YFilter = loadUtilization.YFilter
    loadUtilization.EntityData.YangName = "load-utilization"
    loadUtilization.EntityData.BundleName = "cisco_ios_xr"
    loadUtilization.EntityData.ParentYangName = "np"
    loadUtilization.EntityData.SegmentPath = "load-utilization"
    loadUtilization.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/" + loadUtilization.EntityData.SegmentPath
    loadUtilization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadUtilization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadUtilization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadUtilization.EntityData.Children = types.NewOrderedMap()
    loadUtilization.EntityData.Leafs = types.NewOrderedMap()
    loadUtilization.EntityData.Leafs.Append("utilization", types.YLeaf{"Utilization", loadUtilization.Utilization})
    loadUtilization.EntityData.Leafs.Append("packet-per-second-rate", types.YLeaf{"PacketPerSecondRate", loadUtilization.PacketPerSecondRate})

    loadUtilization.EntityData.YListKeys = []string {}

    return &(loadUtilization.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary
// prm tcam summary info
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Internal tcam summary info.
    InternalTcamInfo HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo

    // External tcam summary info.
    TcamInfo HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetEntityData() *types.CommonEntityData {
    tcamSummary.EntityData.YFilter = tcamSummary.YFilter
    tcamSummary.EntityData.YangName = "tcam-summary"
    tcamSummary.EntityData.BundleName = "cisco_ios_xr"
    tcamSummary.EntityData.ParentYangName = "np"
    tcamSummary.EntityData.SegmentPath = "tcam-summary"
    tcamSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/" + tcamSummary.EntityData.SegmentPath
    tcamSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamSummary.EntityData.Children = types.NewOrderedMap()
    tcamSummary.EntityData.Children.Append("internal-tcam-info", types.YChild{"InternalTcamInfo", &tcamSummary.InternalTcamInfo})
    tcamSummary.EntityData.Children.Append("tcam-info", types.YChild{"TcamInfo", &tcamSummary.TcamInfo})
    tcamSummary.EntityData.Leafs = types.NewOrderedMap()

    tcamSummary.EntityData.YListKeys = []string {}

    return &(tcamSummary.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo
// Internal tcam summary info
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCAM LT ODS 2 summary.
    TcamLtOds2 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2

    // TCAM LT_ODS 8 summary.
    TcamLtOds8 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8

    // Array of TCAM LT L2 partition summaries. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2.
    TcamLtL2 []*HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetEntityData() *types.CommonEntityData {
    internalTcamInfo.EntityData.YFilter = internalTcamInfo.YFilter
    internalTcamInfo.EntityData.YangName = "internal-tcam-info"
    internalTcamInfo.EntityData.BundleName = "cisco_ios_xr"
    internalTcamInfo.EntityData.ParentYangName = "tcam-summary"
    internalTcamInfo.EntityData.SegmentPath = "internal-tcam-info"
    internalTcamInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/" + internalTcamInfo.EntityData.SegmentPath
    internalTcamInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalTcamInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalTcamInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalTcamInfo.EntityData.Children = types.NewOrderedMap()
    internalTcamInfo.EntityData.Children.Append("tcam-lt-ods2", types.YChild{"TcamLtOds2", &internalTcamInfo.TcamLtOds2})
    internalTcamInfo.EntityData.Children.Append("tcam-lt-ods8", types.YChild{"TcamLtOds8", &internalTcamInfo.TcamLtOds8})
    internalTcamInfo.EntityData.Children.Append("tcam-lt-l2", types.YChild{"TcamLtL2", nil})
    for i := range internalTcamInfo.TcamLtL2 {
        types.SetYListKey(internalTcamInfo.TcamLtL2[i], i)
        internalTcamInfo.EntityData.Children.Append(types.GetSegmentPath(internalTcamInfo.TcamLtL2[i]), types.YChild{"TcamLtL2", internalTcamInfo.TcamLtL2[i]})
    }
    internalTcamInfo.EntityData.Leafs = types.NewOrderedMap()

    internalTcamInfo.EntityData.YListKeys = []string {}

    return &(internalTcamInfo.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2
// TCAM LT ODS 2 summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max entries. The type is interface{} with range: 0..4294967295.
    MaxEntries interface{}

    // free entries. The type is interface{} with range: 0..4294967295.
    FreeEntries interface{}

    // app IFIB entry.
    AppIdIfib HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib

    // app qos entry.
    AppIdQos HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos

    // app acl entry.
    AppIdAcl HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl

    // app afmon entry.
    AppIdAfmon HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon

    // app LI entry.
    AppIdLi HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi

    // app PBR entry.
    AppIdPbr HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr

    // app EDPL entry.
    ApplicationEdplEntry HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetEntityData() *types.CommonEntityData {
    tcamLtOds2.EntityData.YFilter = tcamLtOds2.YFilter
    tcamLtOds2.EntityData.YangName = "tcam-lt-ods2"
    tcamLtOds2.EntityData.BundleName = "cisco_ios_xr"
    tcamLtOds2.EntityData.ParentYangName = "internal-tcam-info"
    tcamLtOds2.EntityData.SegmentPath = "tcam-lt-ods2"
    tcamLtOds2.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/" + tcamLtOds2.EntityData.SegmentPath
    tcamLtOds2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamLtOds2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamLtOds2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamLtOds2.EntityData.Children = types.NewOrderedMap()
    tcamLtOds2.EntityData.Children.Append("app-id-ifib", types.YChild{"AppIdIfib", &tcamLtOds2.AppIdIfib})
    tcamLtOds2.EntityData.Children.Append("app-id-qos", types.YChild{"AppIdQos", &tcamLtOds2.AppIdQos})
    tcamLtOds2.EntityData.Children.Append("app-id-acl", types.YChild{"AppIdAcl", &tcamLtOds2.AppIdAcl})
    tcamLtOds2.EntityData.Children.Append("app-id-afmon", types.YChild{"AppIdAfmon", &tcamLtOds2.AppIdAfmon})
    tcamLtOds2.EntityData.Children.Append("app-id-li", types.YChild{"AppIdLi", &tcamLtOds2.AppIdLi})
    tcamLtOds2.EntityData.Children.Append("app-id-pbr", types.YChild{"AppIdPbr", &tcamLtOds2.AppIdPbr})
    tcamLtOds2.EntityData.Children.Append("application-edpl-entry", types.YChild{"ApplicationEdplEntry", &tcamLtOds2.ApplicationEdplEntry})
    tcamLtOds2.EntityData.Leafs = types.NewOrderedMap()
    tcamLtOds2.EntityData.Leafs.Append("max-entries", types.YLeaf{"MaxEntries", tcamLtOds2.MaxEntries})
    tcamLtOds2.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", tcamLtOds2.FreeEntries})

    tcamLtOds2.EntityData.YListKeys = []string {}

    return &(tcamLtOds2.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetEntityData() *types.CommonEntityData {
    appIdIfib.EntityData.YFilter = appIdIfib.YFilter
    appIdIfib.EntityData.YangName = "app-id-ifib"
    appIdIfib.EntityData.BundleName = "cisco_ios_xr"
    appIdIfib.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdIfib.EntityData.SegmentPath = "app-id-ifib"
    appIdIfib.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + appIdIfib.EntityData.SegmentPath
    appIdIfib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdIfib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdIfib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdIfib.EntityData.Children = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdIfib.NumVmrIds})
    appIdIfib.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdIfib.TotalUsedEntries})
    appIdIfib.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdIfib.TotalAllocatedEntries})

    appIdIfib.EntityData.YListKeys = []string {}

    return &(appIdIfib.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetEntityData() *types.CommonEntityData {
    appIdQos.EntityData.YFilter = appIdQos.YFilter
    appIdQos.EntityData.YangName = "app-id-qos"
    appIdQos.EntityData.BundleName = "cisco_ios_xr"
    appIdQos.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdQos.EntityData.SegmentPath = "app-id-qos"
    appIdQos.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + appIdQos.EntityData.SegmentPath
    appIdQos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdQos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdQos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdQos.EntityData.Children = types.NewOrderedMap()
    appIdQos.EntityData.Leafs = types.NewOrderedMap()
    appIdQos.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdQos.NumVmrIds})
    appIdQos.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdQos.TotalUsedEntries})
    appIdQos.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdQos.TotalAllocatedEntries})

    appIdQos.EntityData.YListKeys = []string {}

    return &(appIdQos.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetEntityData() *types.CommonEntityData {
    appIdAcl.EntityData.YFilter = appIdAcl.YFilter
    appIdAcl.EntityData.YangName = "app-id-acl"
    appIdAcl.EntityData.BundleName = "cisco_ios_xr"
    appIdAcl.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdAcl.EntityData.SegmentPath = "app-id-acl"
    appIdAcl.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + appIdAcl.EntityData.SegmentPath
    appIdAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAcl.EntityData.Children = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAcl.NumVmrIds})
    appIdAcl.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdAcl.TotalUsedEntries})
    appIdAcl.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdAcl.TotalAllocatedEntries})

    appIdAcl.EntityData.YListKeys = []string {}

    return &(appIdAcl.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetEntityData() *types.CommonEntityData {
    appIdAfmon.EntityData.YFilter = appIdAfmon.YFilter
    appIdAfmon.EntityData.YangName = "app-id-afmon"
    appIdAfmon.EntityData.BundleName = "cisco_ios_xr"
    appIdAfmon.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdAfmon.EntityData.SegmentPath = "app-id-afmon"
    appIdAfmon.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + appIdAfmon.EntityData.SegmentPath
    appIdAfmon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAfmon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAfmon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAfmon.EntityData.Children = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAfmon.NumVmrIds})
    appIdAfmon.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdAfmon.TotalUsedEntries})
    appIdAfmon.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdAfmon.TotalAllocatedEntries})

    appIdAfmon.EntityData.YListKeys = []string {}

    return &(appIdAfmon.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetEntityData() *types.CommonEntityData {
    appIdLi.EntityData.YFilter = appIdLi.YFilter
    appIdLi.EntityData.YangName = "app-id-li"
    appIdLi.EntityData.BundleName = "cisco_ios_xr"
    appIdLi.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdLi.EntityData.SegmentPath = "app-id-li"
    appIdLi.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + appIdLi.EntityData.SegmentPath
    appIdLi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdLi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdLi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdLi.EntityData.Children = types.NewOrderedMap()
    appIdLi.EntityData.Leafs = types.NewOrderedMap()
    appIdLi.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdLi.NumVmrIds})
    appIdLi.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdLi.TotalUsedEntries})
    appIdLi.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdLi.TotalAllocatedEntries})

    appIdLi.EntityData.YListKeys = []string {}

    return &(appIdLi.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetEntityData() *types.CommonEntityData {
    appIdPbr.EntityData.YFilter = appIdPbr.YFilter
    appIdPbr.EntityData.YangName = "app-id-pbr"
    appIdPbr.EntityData.BundleName = "cisco_ios_xr"
    appIdPbr.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdPbr.EntityData.SegmentPath = "app-id-pbr"
    appIdPbr.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + appIdPbr.EntityData.SegmentPath
    appIdPbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdPbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdPbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdPbr.EntityData.Children = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdPbr.NumVmrIds})
    appIdPbr.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdPbr.TotalUsedEntries})
    appIdPbr.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdPbr.TotalAllocatedEntries})

    appIdPbr.EntityData.YListKeys = []string {}

    return &(appIdPbr.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetEntityData() *types.CommonEntityData {
    applicationEdplEntry.EntityData.YFilter = applicationEdplEntry.YFilter
    applicationEdplEntry.EntityData.YangName = "application-edpl-entry"
    applicationEdplEntry.EntityData.BundleName = "cisco_ios_xr"
    applicationEdplEntry.EntityData.ParentYangName = "tcam-lt-ods2"
    applicationEdplEntry.EntityData.SegmentPath = "application-edpl-entry"
    applicationEdplEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods2/" + applicationEdplEntry.EntityData.SegmentPath
    applicationEdplEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applicationEdplEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applicationEdplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applicationEdplEntry.EntityData.Children = types.NewOrderedMap()
    applicationEdplEntry.EntityData.Leafs = types.NewOrderedMap()
    applicationEdplEntry.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", applicationEdplEntry.NumVmrIds})
    applicationEdplEntry.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", applicationEdplEntry.TotalUsedEntries})
    applicationEdplEntry.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", applicationEdplEntry.TotalAllocatedEntries})

    applicationEdplEntry.EntityData.YListKeys = []string {}

    return &(applicationEdplEntry.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8
// TCAM LT_ODS 8 summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max entries. The type is interface{} with range: 0..4294967295.
    MaxEntries interface{}

    // free entries. The type is interface{} with range: 0..4294967295.
    FreeEntries interface{}

    // app IFIB entry.
    AppIdIfib HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib

    // app qos entry.
    AppIdQos HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos

    // app acl entry.
    AppIdAcl HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl

    // app afmon entry.
    AppIdAfmon HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon

    // app LI entry.
    AppIdLi HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi

    // app PBR entry.
    AppIdPbr HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr

    // app EDPL entry.
    ApplicationEdplEntry HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetEntityData() *types.CommonEntityData {
    tcamLtOds8.EntityData.YFilter = tcamLtOds8.YFilter
    tcamLtOds8.EntityData.YangName = "tcam-lt-ods8"
    tcamLtOds8.EntityData.BundleName = "cisco_ios_xr"
    tcamLtOds8.EntityData.ParentYangName = "internal-tcam-info"
    tcamLtOds8.EntityData.SegmentPath = "tcam-lt-ods8"
    tcamLtOds8.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/" + tcamLtOds8.EntityData.SegmentPath
    tcamLtOds8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamLtOds8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamLtOds8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamLtOds8.EntityData.Children = types.NewOrderedMap()
    tcamLtOds8.EntityData.Children.Append("app-id-ifib", types.YChild{"AppIdIfib", &tcamLtOds8.AppIdIfib})
    tcamLtOds8.EntityData.Children.Append("app-id-qos", types.YChild{"AppIdQos", &tcamLtOds8.AppIdQos})
    tcamLtOds8.EntityData.Children.Append("app-id-acl", types.YChild{"AppIdAcl", &tcamLtOds8.AppIdAcl})
    tcamLtOds8.EntityData.Children.Append("app-id-afmon", types.YChild{"AppIdAfmon", &tcamLtOds8.AppIdAfmon})
    tcamLtOds8.EntityData.Children.Append("app-id-li", types.YChild{"AppIdLi", &tcamLtOds8.AppIdLi})
    tcamLtOds8.EntityData.Children.Append("app-id-pbr", types.YChild{"AppIdPbr", &tcamLtOds8.AppIdPbr})
    tcamLtOds8.EntityData.Children.Append("application-edpl-entry", types.YChild{"ApplicationEdplEntry", &tcamLtOds8.ApplicationEdplEntry})
    tcamLtOds8.EntityData.Leafs = types.NewOrderedMap()
    tcamLtOds8.EntityData.Leafs.Append("max-entries", types.YLeaf{"MaxEntries", tcamLtOds8.MaxEntries})
    tcamLtOds8.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", tcamLtOds8.FreeEntries})

    tcamLtOds8.EntityData.YListKeys = []string {}

    return &(tcamLtOds8.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetEntityData() *types.CommonEntityData {
    appIdIfib.EntityData.YFilter = appIdIfib.YFilter
    appIdIfib.EntityData.YangName = "app-id-ifib"
    appIdIfib.EntityData.BundleName = "cisco_ios_xr"
    appIdIfib.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdIfib.EntityData.SegmentPath = "app-id-ifib"
    appIdIfib.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + appIdIfib.EntityData.SegmentPath
    appIdIfib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdIfib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdIfib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdIfib.EntityData.Children = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdIfib.NumVmrIds})
    appIdIfib.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdIfib.TotalUsedEntries})
    appIdIfib.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdIfib.TotalAllocatedEntries})

    appIdIfib.EntityData.YListKeys = []string {}

    return &(appIdIfib.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetEntityData() *types.CommonEntityData {
    appIdQos.EntityData.YFilter = appIdQos.YFilter
    appIdQos.EntityData.YangName = "app-id-qos"
    appIdQos.EntityData.BundleName = "cisco_ios_xr"
    appIdQos.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdQos.EntityData.SegmentPath = "app-id-qos"
    appIdQos.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + appIdQos.EntityData.SegmentPath
    appIdQos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdQos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdQos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdQos.EntityData.Children = types.NewOrderedMap()
    appIdQos.EntityData.Leafs = types.NewOrderedMap()
    appIdQos.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdQos.NumVmrIds})
    appIdQos.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdQos.TotalUsedEntries})
    appIdQos.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdQos.TotalAllocatedEntries})

    appIdQos.EntityData.YListKeys = []string {}

    return &(appIdQos.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetEntityData() *types.CommonEntityData {
    appIdAcl.EntityData.YFilter = appIdAcl.YFilter
    appIdAcl.EntityData.YangName = "app-id-acl"
    appIdAcl.EntityData.BundleName = "cisco_ios_xr"
    appIdAcl.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdAcl.EntityData.SegmentPath = "app-id-acl"
    appIdAcl.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + appIdAcl.EntityData.SegmentPath
    appIdAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAcl.EntityData.Children = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAcl.NumVmrIds})
    appIdAcl.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdAcl.TotalUsedEntries})
    appIdAcl.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdAcl.TotalAllocatedEntries})

    appIdAcl.EntityData.YListKeys = []string {}

    return &(appIdAcl.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetEntityData() *types.CommonEntityData {
    appIdAfmon.EntityData.YFilter = appIdAfmon.YFilter
    appIdAfmon.EntityData.YangName = "app-id-afmon"
    appIdAfmon.EntityData.BundleName = "cisco_ios_xr"
    appIdAfmon.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdAfmon.EntityData.SegmentPath = "app-id-afmon"
    appIdAfmon.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + appIdAfmon.EntityData.SegmentPath
    appIdAfmon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAfmon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAfmon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAfmon.EntityData.Children = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAfmon.NumVmrIds})
    appIdAfmon.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdAfmon.TotalUsedEntries})
    appIdAfmon.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdAfmon.TotalAllocatedEntries})

    appIdAfmon.EntityData.YListKeys = []string {}

    return &(appIdAfmon.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetEntityData() *types.CommonEntityData {
    appIdLi.EntityData.YFilter = appIdLi.YFilter
    appIdLi.EntityData.YangName = "app-id-li"
    appIdLi.EntityData.BundleName = "cisco_ios_xr"
    appIdLi.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdLi.EntityData.SegmentPath = "app-id-li"
    appIdLi.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + appIdLi.EntityData.SegmentPath
    appIdLi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdLi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdLi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdLi.EntityData.Children = types.NewOrderedMap()
    appIdLi.EntityData.Leafs = types.NewOrderedMap()
    appIdLi.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdLi.NumVmrIds})
    appIdLi.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdLi.TotalUsedEntries})
    appIdLi.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdLi.TotalAllocatedEntries})

    appIdLi.EntityData.YListKeys = []string {}

    return &(appIdLi.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetEntityData() *types.CommonEntityData {
    appIdPbr.EntityData.YFilter = appIdPbr.YFilter
    appIdPbr.EntityData.YangName = "app-id-pbr"
    appIdPbr.EntityData.BundleName = "cisco_ios_xr"
    appIdPbr.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdPbr.EntityData.SegmentPath = "app-id-pbr"
    appIdPbr.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + appIdPbr.EntityData.SegmentPath
    appIdPbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdPbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdPbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdPbr.EntityData.Children = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdPbr.NumVmrIds})
    appIdPbr.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", appIdPbr.TotalUsedEntries})
    appIdPbr.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", appIdPbr.TotalAllocatedEntries})

    appIdPbr.EntityData.YListKeys = []string {}

    return &(appIdPbr.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // number of used vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalUsedEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    TotalAllocatedEntries interface{}
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetEntityData() *types.CommonEntityData {
    applicationEdplEntry.EntityData.YFilter = applicationEdplEntry.YFilter
    applicationEdplEntry.EntityData.YangName = "application-edpl-entry"
    applicationEdplEntry.EntityData.BundleName = "cisco_ios_xr"
    applicationEdplEntry.EntityData.ParentYangName = "tcam-lt-ods8"
    applicationEdplEntry.EntityData.SegmentPath = "application-edpl-entry"
    applicationEdplEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/tcam-lt-ods8/" + applicationEdplEntry.EntityData.SegmentPath
    applicationEdplEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applicationEdplEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applicationEdplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applicationEdplEntry.EntityData.Children = types.NewOrderedMap()
    applicationEdplEntry.EntityData.Leafs = types.NewOrderedMap()
    applicationEdplEntry.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", applicationEdplEntry.NumVmrIds})
    applicationEdplEntry.EntityData.Leafs.Append("total-used-entries", types.YLeaf{"TotalUsedEntries", applicationEdplEntry.TotalUsedEntries})
    applicationEdplEntry.EntityData.Leafs.Append("total-allocated-entries", types.YLeaf{"TotalAllocatedEntries", applicationEdplEntry.TotalAllocatedEntries})

    applicationEdplEntry.EntityData.YListKeys = []string {}

    return &(applicationEdplEntry.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2
// Array of TCAM LT L2 partition summaries
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // PartitionID. The type is interface{} with range: 0..4294967295.
    PartitionId interface{}

    // Valid Entries. The type is interface{} with range: 0..4294967295.
    ValidEntries interface{}

    // Free Entries. The type is interface{} with range: 0..4294967295.
    FreeEntries interface{}
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetEntityData() *types.CommonEntityData {
    tcamLtL2.EntityData.YFilter = tcamLtL2.YFilter
    tcamLtL2.EntityData.YangName = "tcam-lt-l2"
    tcamLtL2.EntityData.BundleName = "cisco_ios_xr"
    tcamLtL2.EntityData.ParentYangName = "internal-tcam-info"
    tcamLtL2.EntityData.SegmentPath = "tcam-lt-l2" + types.AddNoKeyToken(tcamLtL2)
    tcamLtL2.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/internal-tcam-info/" + tcamLtL2.EntityData.SegmentPath
    tcamLtL2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamLtL2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamLtL2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamLtL2.EntityData.Children = types.NewOrderedMap()
    tcamLtL2.EntityData.Leafs = types.NewOrderedMap()
    tcamLtL2.EntityData.Leafs.Append("partition-id", types.YLeaf{"PartitionId", tcamLtL2.PartitionId})
    tcamLtL2.EntityData.Leafs.Append("valid-entries", types.YLeaf{"ValidEntries", tcamLtL2.ValidEntries})
    tcamLtL2.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", tcamLtL2.FreeEntries})

    tcamLtL2.EntityData.YListKeys = []string {}

    return &(tcamLtL2.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo
// External tcam summary info
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCAM ODS2 partition summary.
    TcamLtOds2 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2

    // TCAM ODS8 partition summary.
    TcamLtOds8 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8

    // Array of TCAM L2 partition summaries. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2.
    TcamLtL2 []*HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetEntityData() *types.CommonEntityData {
    tcamInfo.EntityData.YFilter = tcamInfo.YFilter
    tcamInfo.EntityData.YangName = "tcam-info"
    tcamInfo.EntityData.BundleName = "cisco_ios_xr"
    tcamInfo.EntityData.ParentYangName = "tcam-summary"
    tcamInfo.EntityData.SegmentPath = "tcam-info"
    tcamInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/" + tcamInfo.EntityData.SegmentPath
    tcamInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamInfo.EntityData.Children = types.NewOrderedMap()
    tcamInfo.EntityData.Children.Append("tcam-lt-ods2", types.YChild{"TcamLtOds2", &tcamInfo.TcamLtOds2})
    tcamInfo.EntityData.Children.Append("tcam-lt-ods8", types.YChild{"TcamLtOds8", &tcamInfo.TcamLtOds8})
    tcamInfo.EntityData.Children.Append("tcam-lt-l2", types.YChild{"TcamLtL2", nil})
    for i := range tcamInfo.TcamLtL2 {
        types.SetYListKey(tcamInfo.TcamLtL2[i], i)
        tcamInfo.EntityData.Children.Append(types.GetSegmentPath(tcamInfo.TcamLtL2[i]), types.YChild{"TcamLtL2", tcamInfo.TcamLtL2[i]})
    }
    tcamInfo.EntityData.Leafs = types.NewOrderedMap()

    tcamInfo.EntityData.YListKeys = []string {}

    return &(tcamInfo.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2
// TCAM ODS2 partition summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Free entries in the table. The type is interface{} with range:
    // 0..4294967295.
    FreeEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    ReservedEntries interface{}

    // ACL common region.
    AclCommon HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon

    // app IFIB entry.
    AppIdIfib HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib

    // app qos entry.
    AppIdQos HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos

    // app acl entry.
    AppIdAcl HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl

    // app afmon entry.
    AppIdAfmon HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon

    // app LI entry.
    AppIdLi HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi

    // app PBR entry.
    AppIdPbr HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr

    // app EDPL entry.
    AppIdEdpl HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetEntityData() *types.CommonEntityData {
    tcamLtOds2.EntityData.YFilter = tcamLtOds2.YFilter
    tcamLtOds2.EntityData.YangName = "tcam-lt-ods2"
    tcamLtOds2.EntityData.BundleName = "cisco_ios_xr"
    tcamLtOds2.EntityData.ParentYangName = "tcam-info"
    tcamLtOds2.EntityData.SegmentPath = "tcam-lt-ods2"
    tcamLtOds2.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/" + tcamLtOds2.EntityData.SegmentPath
    tcamLtOds2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamLtOds2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamLtOds2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamLtOds2.EntityData.Children = types.NewOrderedMap()
    tcamLtOds2.EntityData.Children.Append("acl-common", types.YChild{"AclCommon", &tcamLtOds2.AclCommon})
    tcamLtOds2.EntityData.Children.Append("app-id-ifib", types.YChild{"AppIdIfib", &tcamLtOds2.AppIdIfib})
    tcamLtOds2.EntityData.Children.Append("app-id-qos", types.YChild{"AppIdQos", &tcamLtOds2.AppIdQos})
    tcamLtOds2.EntityData.Children.Append("app-id-acl", types.YChild{"AppIdAcl", &tcamLtOds2.AppIdAcl})
    tcamLtOds2.EntityData.Children.Append("app-id-afmon", types.YChild{"AppIdAfmon", &tcamLtOds2.AppIdAfmon})
    tcamLtOds2.EntityData.Children.Append("app-id-li", types.YChild{"AppIdLi", &tcamLtOds2.AppIdLi})
    tcamLtOds2.EntityData.Children.Append("app-id-pbr", types.YChild{"AppIdPbr", &tcamLtOds2.AppIdPbr})
    tcamLtOds2.EntityData.Children.Append("app-id-edpl", types.YChild{"AppIdEdpl", &tcamLtOds2.AppIdEdpl})
    tcamLtOds2.EntityData.Leafs = types.NewOrderedMap()
    tcamLtOds2.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", tcamLtOds2.FreeEntries})
    tcamLtOds2.EntityData.Leafs.Append("reserved-entries", types.YLeaf{"ReservedEntries", tcamLtOds2.ReservedEntries})

    tcamLtOds2.EntityData.YListKeys = []string {}

    return &(tcamLtOds2.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon
// ACL common region
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Free entries in the table. The type is interface{} with range:
    // 0..4294967295.
    FreeEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    AllocatedEntries interface{}
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetEntityData() *types.CommonEntityData {
    aclCommon.EntityData.YFilter = aclCommon.YFilter
    aclCommon.EntityData.YangName = "acl-common"
    aclCommon.EntityData.BundleName = "cisco_ios_xr"
    aclCommon.EntityData.ParentYangName = "tcam-lt-ods2"
    aclCommon.EntityData.SegmentPath = "acl-common"
    aclCommon.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + aclCommon.EntityData.SegmentPath
    aclCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclCommon.EntityData.Children = types.NewOrderedMap()
    aclCommon.EntityData.Leafs = types.NewOrderedMap()
    aclCommon.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", aclCommon.FreeEntries})
    aclCommon.EntityData.Leafs.Append("allocated-entries", types.YLeaf{"AllocatedEntries", aclCommon.AllocatedEntries})

    aclCommon.EntityData.YListKeys = []string {}

    return &(aclCommon.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetEntityData() *types.CommonEntityData {
    appIdIfib.EntityData.YFilter = appIdIfib.YFilter
    appIdIfib.EntityData.YangName = "app-id-ifib"
    appIdIfib.EntityData.BundleName = "cisco_ios_xr"
    appIdIfib.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdIfib.EntityData.SegmentPath = "app-id-ifib"
    appIdIfib.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdIfib.EntityData.SegmentPath
    appIdIfib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdIfib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdIfib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdIfib.EntityData.Children = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdIfib.NumVmrIds})
    appIdIfib.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdIfib.NumActiveEntries})
    appIdIfib.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdIfib.NumAllocatedEntries})

    appIdIfib.EntityData.YListKeys = []string {}

    return &(appIdIfib.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetEntityData() *types.CommonEntityData {
    appIdQos.EntityData.YFilter = appIdQos.YFilter
    appIdQos.EntityData.YangName = "app-id-qos"
    appIdQos.EntityData.BundleName = "cisco_ios_xr"
    appIdQos.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdQos.EntityData.SegmentPath = "app-id-qos"
    appIdQos.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdQos.EntityData.SegmentPath
    appIdQos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdQos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdQos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdQos.EntityData.Children = types.NewOrderedMap()
    appIdQos.EntityData.Leafs = types.NewOrderedMap()
    appIdQos.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdQos.NumVmrIds})
    appIdQos.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdQos.NumActiveEntries})
    appIdQos.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdQos.NumAllocatedEntries})

    appIdQos.EntityData.YListKeys = []string {}

    return &(appIdQos.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetEntityData() *types.CommonEntityData {
    appIdAcl.EntityData.YFilter = appIdAcl.YFilter
    appIdAcl.EntityData.YangName = "app-id-acl"
    appIdAcl.EntityData.BundleName = "cisco_ios_xr"
    appIdAcl.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdAcl.EntityData.SegmentPath = "app-id-acl"
    appIdAcl.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdAcl.EntityData.SegmentPath
    appIdAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAcl.EntityData.Children = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAcl.NumVmrIds})
    appIdAcl.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdAcl.NumActiveEntries})
    appIdAcl.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdAcl.NumAllocatedEntries})

    appIdAcl.EntityData.YListKeys = []string {}

    return &(appIdAcl.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetEntityData() *types.CommonEntityData {
    appIdAfmon.EntityData.YFilter = appIdAfmon.YFilter
    appIdAfmon.EntityData.YangName = "app-id-afmon"
    appIdAfmon.EntityData.BundleName = "cisco_ios_xr"
    appIdAfmon.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdAfmon.EntityData.SegmentPath = "app-id-afmon"
    appIdAfmon.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdAfmon.EntityData.SegmentPath
    appIdAfmon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAfmon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAfmon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAfmon.EntityData.Children = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAfmon.NumVmrIds})
    appIdAfmon.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdAfmon.NumActiveEntries})
    appIdAfmon.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdAfmon.NumAllocatedEntries})

    appIdAfmon.EntityData.YListKeys = []string {}

    return &(appIdAfmon.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetEntityData() *types.CommonEntityData {
    appIdLi.EntityData.YFilter = appIdLi.YFilter
    appIdLi.EntityData.YangName = "app-id-li"
    appIdLi.EntityData.BundleName = "cisco_ios_xr"
    appIdLi.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdLi.EntityData.SegmentPath = "app-id-li"
    appIdLi.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdLi.EntityData.SegmentPath
    appIdLi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdLi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdLi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdLi.EntityData.Children = types.NewOrderedMap()
    appIdLi.EntityData.Leafs = types.NewOrderedMap()
    appIdLi.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdLi.NumVmrIds})
    appIdLi.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdLi.NumActiveEntries})
    appIdLi.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdLi.NumAllocatedEntries})

    appIdLi.EntityData.YListKeys = []string {}

    return &(appIdLi.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetEntityData() *types.CommonEntityData {
    appIdPbr.EntityData.YFilter = appIdPbr.YFilter
    appIdPbr.EntityData.YangName = "app-id-pbr"
    appIdPbr.EntityData.BundleName = "cisco_ios_xr"
    appIdPbr.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdPbr.EntityData.SegmentPath = "app-id-pbr"
    appIdPbr.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdPbr.EntityData.SegmentPath
    appIdPbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdPbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdPbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdPbr.EntityData.Children = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdPbr.NumVmrIds})
    appIdPbr.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdPbr.NumActiveEntries})
    appIdPbr.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdPbr.NumAllocatedEntries})

    appIdPbr.EntityData.YListKeys = []string {}

    return &(appIdPbr.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetEntityData() *types.CommonEntityData {
    appIdEdpl.EntityData.YFilter = appIdEdpl.YFilter
    appIdEdpl.EntityData.YangName = "app-id-edpl"
    appIdEdpl.EntityData.BundleName = "cisco_ios_xr"
    appIdEdpl.EntityData.ParentYangName = "tcam-lt-ods2"
    appIdEdpl.EntityData.SegmentPath = "app-id-edpl"
    appIdEdpl.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods2/" + appIdEdpl.EntityData.SegmentPath
    appIdEdpl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdEdpl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdEdpl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdEdpl.EntityData.Children = types.NewOrderedMap()
    appIdEdpl.EntityData.Leafs = types.NewOrderedMap()
    appIdEdpl.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdEdpl.NumVmrIds})
    appIdEdpl.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdEdpl.NumActiveEntries})
    appIdEdpl.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdEdpl.NumAllocatedEntries})

    appIdEdpl.EntityData.YListKeys = []string {}

    return &(appIdEdpl.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8
// TCAM ODS8 partition summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Free entries in the table. The type is interface{} with range:
    // 0..4294967295.
    FreeEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    ReservedEntries interface{}

    // ACL common region.
    AclCommon HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon

    // app IFIB entry.
    AppIdIfib HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib

    // app qos entry.
    AppIdQos HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos

    // app acl entry.
    AppIdAcl HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl

    // app afmon entry.
    AppIdAfmon HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon

    // app LI entry.
    AppIdLi HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi

    // app PBR entry.
    AppIdPbr HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr

    // app EDPL entry.
    AppIdEdpl HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetEntityData() *types.CommonEntityData {
    tcamLtOds8.EntityData.YFilter = tcamLtOds8.YFilter
    tcamLtOds8.EntityData.YangName = "tcam-lt-ods8"
    tcamLtOds8.EntityData.BundleName = "cisco_ios_xr"
    tcamLtOds8.EntityData.ParentYangName = "tcam-info"
    tcamLtOds8.EntityData.SegmentPath = "tcam-lt-ods8"
    tcamLtOds8.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/" + tcamLtOds8.EntityData.SegmentPath
    tcamLtOds8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamLtOds8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamLtOds8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamLtOds8.EntityData.Children = types.NewOrderedMap()
    tcamLtOds8.EntityData.Children.Append("acl-common", types.YChild{"AclCommon", &tcamLtOds8.AclCommon})
    tcamLtOds8.EntityData.Children.Append("app-id-ifib", types.YChild{"AppIdIfib", &tcamLtOds8.AppIdIfib})
    tcamLtOds8.EntityData.Children.Append("app-id-qos", types.YChild{"AppIdQos", &tcamLtOds8.AppIdQos})
    tcamLtOds8.EntityData.Children.Append("app-id-acl", types.YChild{"AppIdAcl", &tcamLtOds8.AppIdAcl})
    tcamLtOds8.EntityData.Children.Append("app-id-afmon", types.YChild{"AppIdAfmon", &tcamLtOds8.AppIdAfmon})
    tcamLtOds8.EntityData.Children.Append("app-id-li", types.YChild{"AppIdLi", &tcamLtOds8.AppIdLi})
    tcamLtOds8.EntityData.Children.Append("app-id-pbr", types.YChild{"AppIdPbr", &tcamLtOds8.AppIdPbr})
    tcamLtOds8.EntityData.Children.Append("app-id-edpl", types.YChild{"AppIdEdpl", &tcamLtOds8.AppIdEdpl})
    tcamLtOds8.EntityData.Leafs = types.NewOrderedMap()
    tcamLtOds8.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", tcamLtOds8.FreeEntries})
    tcamLtOds8.EntityData.Leafs.Append("reserved-entries", types.YLeaf{"ReservedEntries", tcamLtOds8.ReservedEntries})

    tcamLtOds8.EntityData.YListKeys = []string {}

    return &(tcamLtOds8.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon
// ACL common region
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Free entries in the table. The type is interface{} with range:
    // 0..4294967295.
    FreeEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    AllocatedEntries interface{}
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetEntityData() *types.CommonEntityData {
    aclCommon.EntityData.YFilter = aclCommon.YFilter
    aclCommon.EntityData.YangName = "acl-common"
    aclCommon.EntityData.BundleName = "cisco_ios_xr"
    aclCommon.EntityData.ParentYangName = "tcam-lt-ods8"
    aclCommon.EntityData.SegmentPath = "acl-common"
    aclCommon.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + aclCommon.EntityData.SegmentPath
    aclCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclCommon.EntityData.Children = types.NewOrderedMap()
    aclCommon.EntityData.Leafs = types.NewOrderedMap()
    aclCommon.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", aclCommon.FreeEntries})
    aclCommon.EntityData.Leafs.Append("allocated-entries", types.YLeaf{"AllocatedEntries", aclCommon.AllocatedEntries})

    aclCommon.EntityData.YListKeys = []string {}

    return &(aclCommon.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetEntityData() *types.CommonEntityData {
    appIdIfib.EntityData.YFilter = appIdIfib.YFilter
    appIdIfib.EntityData.YangName = "app-id-ifib"
    appIdIfib.EntityData.BundleName = "cisco_ios_xr"
    appIdIfib.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdIfib.EntityData.SegmentPath = "app-id-ifib"
    appIdIfib.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdIfib.EntityData.SegmentPath
    appIdIfib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdIfib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdIfib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdIfib.EntityData.Children = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs = types.NewOrderedMap()
    appIdIfib.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdIfib.NumVmrIds})
    appIdIfib.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdIfib.NumActiveEntries})
    appIdIfib.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdIfib.NumAllocatedEntries})

    appIdIfib.EntityData.YListKeys = []string {}

    return &(appIdIfib.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetEntityData() *types.CommonEntityData {
    appIdQos.EntityData.YFilter = appIdQos.YFilter
    appIdQos.EntityData.YangName = "app-id-qos"
    appIdQos.EntityData.BundleName = "cisco_ios_xr"
    appIdQos.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdQos.EntityData.SegmentPath = "app-id-qos"
    appIdQos.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdQos.EntityData.SegmentPath
    appIdQos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdQos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdQos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdQos.EntityData.Children = types.NewOrderedMap()
    appIdQos.EntityData.Leafs = types.NewOrderedMap()
    appIdQos.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdQos.NumVmrIds})
    appIdQos.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdQos.NumActiveEntries})
    appIdQos.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdQos.NumAllocatedEntries})

    appIdQos.EntityData.YListKeys = []string {}

    return &(appIdQos.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetEntityData() *types.CommonEntityData {
    appIdAcl.EntityData.YFilter = appIdAcl.YFilter
    appIdAcl.EntityData.YangName = "app-id-acl"
    appIdAcl.EntityData.BundleName = "cisco_ios_xr"
    appIdAcl.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdAcl.EntityData.SegmentPath = "app-id-acl"
    appIdAcl.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdAcl.EntityData.SegmentPath
    appIdAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAcl.EntityData.Children = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs = types.NewOrderedMap()
    appIdAcl.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAcl.NumVmrIds})
    appIdAcl.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdAcl.NumActiveEntries})
    appIdAcl.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdAcl.NumAllocatedEntries})

    appIdAcl.EntityData.YListKeys = []string {}

    return &(appIdAcl.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetEntityData() *types.CommonEntityData {
    appIdAfmon.EntityData.YFilter = appIdAfmon.YFilter
    appIdAfmon.EntityData.YangName = "app-id-afmon"
    appIdAfmon.EntityData.BundleName = "cisco_ios_xr"
    appIdAfmon.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdAfmon.EntityData.SegmentPath = "app-id-afmon"
    appIdAfmon.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdAfmon.EntityData.SegmentPath
    appIdAfmon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdAfmon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdAfmon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdAfmon.EntityData.Children = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs = types.NewOrderedMap()
    appIdAfmon.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdAfmon.NumVmrIds})
    appIdAfmon.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdAfmon.NumActiveEntries})
    appIdAfmon.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdAfmon.NumAllocatedEntries})

    appIdAfmon.EntityData.YListKeys = []string {}

    return &(appIdAfmon.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetEntityData() *types.CommonEntityData {
    appIdLi.EntityData.YFilter = appIdLi.YFilter
    appIdLi.EntityData.YangName = "app-id-li"
    appIdLi.EntityData.BundleName = "cisco_ios_xr"
    appIdLi.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdLi.EntityData.SegmentPath = "app-id-li"
    appIdLi.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdLi.EntityData.SegmentPath
    appIdLi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdLi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdLi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdLi.EntityData.Children = types.NewOrderedMap()
    appIdLi.EntityData.Leafs = types.NewOrderedMap()
    appIdLi.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdLi.NumVmrIds})
    appIdLi.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdLi.NumActiveEntries})
    appIdLi.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdLi.NumAllocatedEntries})

    appIdLi.EntityData.YListKeys = []string {}

    return &(appIdLi.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetEntityData() *types.CommonEntityData {
    appIdPbr.EntityData.YFilter = appIdPbr.YFilter
    appIdPbr.EntityData.YangName = "app-id-pbr"
    appIdPbr.EntityData.BundleName = "cisco_ios_xr"
    appIdPbr.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdPbr.EntityData.SegmentPath = "app-id-pbr"
    appIdPbr.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdPbr.EntityData.SegmentPath
    appIdPbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdPbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdPbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdPbr.EntityData.Children = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs = types.NewOrderedMap()
    appIdPbr.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdPbr.NumVmrIds})
    appIdPbr.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdPbr.NumActiveEntries})
    appIdPbr.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdPbr.NumAllocatedEntries})

    appIdPbr.EntityData.YListKeys = []string {}

    return &(appIdPbr.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vmr IDs. The type is interface{} with range: 0..4294967295.
    NumVmrIds interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumActiveEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    NumAllocatedEntries interface{}
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetEntityData() *types.CommonEntityData {
    appIdEdpl.EntityData.YFilter = appIdEdpl.YFilter
    appIdEdpl.EntityData.YangName = "app-id-edpl"
    appIdEdpl.EntityData.BundleName = "cisco_ios_xr"
    appIdEdpl.EntityData.ParentYangName = "tcam-lt-ods8"
    appIdEdpl.EntityData.SegmentPath = "app-id-edpl"
    appIdEdpl.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/tcam-lt-ods8/" + appIdEdpl.EntityData.SegmentPath
    appIdEdpl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appIdEdpl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appIdEdpl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appIdEdpl.EntityData.Children = types.NewOrderedMap()
    appIdEdpl.EntityData.Leafs = types.NewOrderedMap()
    appIdEdpl.EntityData.Leafs.Append("num-vmr-ids", types.YLeaf{"NumVmrIds", appIdEdpl.NumVmrIds})
    appIdEdpl.EntityData.Leafs.Append("num-active-entries", types.YLeaf{"NumActiveEntries", appIdEdpl.NumActiveEntries})
    appIdEdpl.EntityData.Leafs.Append("num-allocated-entries", types.YLeaf{"NumAllocatedEntries", appIdEdpl.NumAllocatedEntries})

    appIdEdpl.EntityData.YListKeys = []string {}

    return &(appIdEdpl.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2
// Array of TCAM L2 partition summaries
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // PartitionID. The type is interface{} with range: 0..4294967295.
    PartitionId interface{}

    // Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Valid Entries. The type is interface{} with range: 0..4294967295.
    ValidEntries interface{}

    // Free Entries. The type is interface{} with range: 0..4294967295.
    FreeEntries interface{}
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetEntityData() *types.CommonEntityData {
    tcamLtL2.EntityData.YFilter = tcamLtL2.YFilter
    tcamLtL2.EntityData.YangName = "tcam-lt-l2"
    tcamLtL2.EntityData.BundleName = "cisco_ios_xr"
    tcamLtL2.EntityData.ParentYangName = "tcam-info"
    tcamLtL2.EntityData.SegmentPath = "tcam-lt-l2" + types.AddNoKeyToken(tcamLtL2)
    tcamLtL2.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/tcam-summary/tcam-info/" + tcamLtL2.EntityData.SegmentPath
    tcamLtL2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamLtL2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamLtL2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamLtL2.EntityData.Children = types.NewOrderedMap()
    tcamLtL2.EntityData.Leafs = types.NewOrderedMap()
    tcamLtL2.EntityData.Leafs.Append("partition-id", types.YLeaf{"PartitionId", tcamLtL2.PartitionId})
    tcamLtL2.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", tcamLtL2.Priority})
    tcamLtL2.EntityData.Leafs.Append("valid-entries", types.YLeaf{"ValidEntries", tcamLtL2.ValidEntries})
    tcamLtL2.EntityData.Leafs.Append("free-entries", types.YLeaf{"FreeEntries", tcamLtL2.FreeEntries})

    tcamLtL2.EntityData.YListKeys = []string {}

    return &(tcamLtL2.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_Counters
// prm counters info
type HardwareModuleNp_Nodes_Node_Nps_Np_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of NP Counters. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter.
    NpCounter []*HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "np"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/" + counters.EntityData.SegmentPath
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Children.Append("np-counter", types.YChild{"NpCounter", nil})
    for i := range counters.NpCounter {
        types.SetYListKey(counters.NpCounter[i], i)
        counters.EntityData.Children.Append(types.GetSegmentPath(counters.NpCounter[i]), types.YChild{"NpCounter", counters.NpCounter[i]})
    }
    counters.EntityData.Leafs = types.NewOrderedMap()

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter
// Array of NP Counters
type HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Counter Index. The type is interface{} with range: 0..4294967295.
    CounterIndex interface{}

    // The accurate value of the counter. The type is interface{} with range:
    // 0..18446744073709551615.
    CounterValue interface{}

    // Rate in Packets Per Second. The type is interface{} with range:
    // 0..4294967295. Units are packet/s.
    Rate interface{}

    // Counter TypeDROP: Drop counterPUNT: Punt counterFWD:  Forward or generic
    // counterUNKNOWN: Counter type unknown. The type is string.
    CounterType interface{}

    // Counter name. The type is string.
    CounterName interface{}
}

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetEntityData() *types.CommonEntityData {
    npCounter.EntityData.YFilter = npCounter.YFilter
    npCounter.EntityData.YangName = "np-counter"
    npCounter.EntityData.BundleName = "cisco_ios_xr"
    npCounter.EntityData.ParentYangName = "counters"
    npCounter.EntityData.SegmentPath = "np-counter" + types.AddNoKeyToken(npCounter)
    npCounter.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/counters/" + npCounter.EntityData.SegmentPath
    npCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npCounter.EntityData.Children = types.NewOrderedMap()
    npCounter.EntityData.Leafs = types.NewOrderedMap()
    npCounter.EntityData.Leafs.Append("counter-index", types.YLeaf{"CounterIndex", npCounter.CounterIndex})
    npCounter.EntityData.Leafs.Append("counter-value", types.YLeaf{"CounterValue", npCounter.CounterValue})
    npCounter.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", npCounter.Rate})
    npCounter.EntityData.Leafs.Append("counter-type", types.YLeaf{"CounterType", npCounter.CounterType})
    npCounter.EntityData.Leafs.Append("counter-name", types.YLeaf{"CounterName", npCounter.CounterName})

    npCounter.EntityData.YListKeys = []string {}

    return &(npCounter.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop
// prm fast drop counters info
type HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of NP Fast Drop Counters. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop.
    NpFastDrop []*HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetEntityData() *types.CommonEntityData {
    fastDrop.EntityData.YFilter = fastDrop.YFilter
    fastDrop.EntityData.YangName = "fast-drop"
    fastDrop.EntityData.BundleName = "cisco_ios_xr"
    fastDrop.EntityData.ParentYangName = "np"
    fastDrop.EntityData.SegmentPath = "fast-drop"
    fastDrop.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/" + fastDrop.EntityData.SegmentPath
    fastDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastDrop.EntityData.Children = types.NewOrderedMap()
    fastDrop.EntityData.Children.Append("np-fast-drop", types.YChild{"NpFastDrop", nil})
    for i := range fastDrop.NpFastDrop {
        types.SetYListKey(fastDrop.NpFastDrop[i], i)
        fastDrop.EntityData.Children.Append(types.GetSegmentPath(fastDrop.NpFastDrop[i]), types.YChild{"NpFastDrop", fastDrop.NpFastDrop[i]})
    }
    fastDrop.EntityData.Leafs = types.NewOrderedMap()

    fastDrop.EntityData.YListKeys = []string {}

    return &(fastDrop.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop
// Array of NP Fast Drop Counters
type HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface name. The type is string.
    InterfaceName interface{}

    // The Value of the counter. The type is interface{} with range:
    // 0..18446744073709551615.
    CounterValue interface{}
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetEntityData() *types.CommonEntityData {
    npFastDrop.EntityData.YFilter = npFastDrop.YFilter
    npFastDrop.EntityData.YangName = "np-fast-drop"
    npFastDrop.EntityData.BundleName = "cisco_ios_xr"
    npFastDrop.EntityData.ParentYangName = "fast-drop"
    npFastDrop.EntityData.SegmentPath = "np-fast-drop" + types.AddNoKeyToken(npFastDrop)
    npFastDrop.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/fast-drop/" + npFastDrop.EntityData.SegmentPath
    npFastDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npFastDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npFastDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npFastDrop.EntityData.Children = types.NewOrderedMap()
    npFastDrop.EntityData.Leafs = types.NewOrderedMap()
    npFastDrop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", npFastDrop.InterfaceName})
    npFastDrop.EntityData.Leafs.Append("counter-value", types.YLeaf{"CounterValue", npFastDrop.CounterValue})

    npFastDrop.EntityData.YListKeys = []string {}

    return &(npFastDrop.EntityData)
}

// HardwareModuleNp_Nodes_Node_Nps_Np_Efd
// EFD COS to Priority Mapping
type HardwareModuleNp_Nodes_Node_Nps_Np_Efd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // high-priority-list CLI is supported. The type is interface{} with range:
    // 0..4294967295.
    HpListIsSupported interface{}

    // Array of VLAN COS Mapping. The type is slice of interface{} with range:
    // 0..4294967295.
    VlanCosMapping []interface{}

    // Array of MPLS COS Mapping. The type is slice of interface{} with range:
    // 0..4294967295.
    MplscosMapping []interface{}

    // Array of IP COS Mapping. The type is slice of interface{} with range:
    // 0..4294967295.
    IpcosMapping []interface{}
}

func (efd *HardwareModuleNp_Nodes_Node_Nps_Np_Efd) GetEntityData() *types.CommonEntityData {
    efd.EntityData.YFilter = efd.YFilter
    efd.EntityData.YangName = "efd"
    efd.EntityData.BundleName = "cisco_ios_xr"
    efd.EntityData.ParentYangName = "np"
    efd.EntityData.SegmentPath = "efd"
    efd.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np/nodes/node/nps/np/" + efd.EntityData.SegmentPath
    efd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    efd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    efd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    efd.EntityData.Children = types.NewOrderedMap()
    efd.EntityData.Leafs = types.NewOrderedMap()
    efd.EntityData.Leafs.Append("hp-list-is-supported", types.YLeaf{"HpListIsSupported", efd.HpListIsSupported})
    efd.EntityData.Leafs.Append("vlan-cos-mapping", types.YLeaf{"VlanCosMapping", efd.VlanCosMapping})
    efd.EntityData.Leafs.Append("mplscos-mapping", types.YLeaf{"MplscosMapping", efd.MplscosMapping})
    efd.EntityData.Leafs.Append("ipcos-mapping", types.YLeaf{"IpcosMapping", efd.IpcosMapping})

    efd.EntityData.YListKeys = []string {}

    return &(efd.EntityData)
}

