// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-np package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-np: Hardware NP Counters
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes HardwareModuleNp_Nodes
}

func (hardwareModuleNp *HardwareModuleNp) GetFilter() yfilter.YFilter { return hardwareModuleNp.YFilter }

func (hardwareModuleNp *HardwareModuleNp) SetFilter(yf yfilter.YFilter) { hardwareModuleNp.YFilter = yf }

func (hardwareModuleNp *HardwareModuleNp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModuleNp *HardwareModuleNp) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-np-oper:hardware-module-np"
}

func (hardwareModuleNp *HardwareModuleNp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &hardwareModuleNp.Nodes
    }
    return nil
}

func (hardwareModuleNp *HardwareModuleNp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &hardwareModuleNp.Nodes
    return children
}

func (hardwareModuleNp *HardwareModuleNp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModuleNp *HardwareModuleNp) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleNp *HardwareModuleNp) GetYangName() string { return "hardware-module-np" }

func (hardwareModuleNp *HardwareModuleNp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleNp *HardwareModuleNp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleNp *HardwareModuleNp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleNp *HardwareModuleNp) SetParent(parent types.Entity) { hardwareModuleNp.parent = parent }

func (hardwareModuleNp *HardwareModuleNp) GetParent() types.Entity { return hardwareModuleNp.parent }

func (hardwareModuleNp *HardwareModuleNp) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-np-oper" }

// HardwareModuleNp_Nodes
// Table of nodes
type HardwareModuleNp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of HardwareModuleNp_Nodes_Node.
    Node []HardwareModuleNp_Nodes_Node
}

func (nodes *HardwareModuleNp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModuleNp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModuleNp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModuleNp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModuleNp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModuleNp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModuleNp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModuleNp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModuleNp_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModuleNp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModuleNp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModuleNp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModuleNp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModuleNp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModuleNp_Nodes) GetParentYangName() string { return "hardware-module-np" }

// HardwareModuleNp_Nodes_Node
// Number
type HardwareModuleNp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of all NP.
    Nps HardwareModuleNp_Nodes_Node_Nps
}

func (node *HardwareModuleNp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModuleNp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModuleNp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "nps" { return "Nps" }
    return ""
}

func (node *HardwareModuleNp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModuleNp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nps" {
        return &node.Nps
    }
    return nil
}

func (node *HardwareModuleNp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nps"] = &node.Nps
    return children
}

func (node *HardwareModuleNp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *HardwareModuleNp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModuleNp_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModuleNp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModuleNp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModuleNp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModuleNp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModuleNp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModuleNp_Nodes_Node) GetParentYangName() string { return "nodes" }

// HardwareModuleNp_Nodes_Node_Nps
// List of all NP
type HardwareModuleNp_Nodes_Node_Nps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // np0 to np7. The type is slice of HardwareModuleNp_Nodes_Node_Nps_Np.
    Np []HardwareModuleNp_Nodes_Node_Nps_Np
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetFilter() yfilter.YFilter { return nps.YFilter }

func (nps *HardwareModuleNp_Nodes_Node_Nps) SetFilter(yf yfilter.YFilter) { nps.YFilter = yf }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetGoName(yname string) string {
    if yname == "np" { return "Np" }
    return ""
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetSegmentPath() string {
    return "nps"
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np" {
        for _, c := range nps.Np {
            if nps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node_Nps_Np{}
        nps.Np = append(nps.Np, child)
        return &nps.Np[len(nps.Np)-1]
    }
    return nil
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nps.Np {
        children[nps.Np[i].GetSegmentPath()] = &nps.Np[i]
    }
    return children
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetBundleName() string { return "cisco_ios_xr" }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetYangName() string { return "nps" }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nps *HardwareModuleNp_Nodes_Node_Nps) SetParent(parent types.Entity) { nps.parent = parent }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetParent() types.Entity { return nps.parent }

func (nps *HardwareModuleNp_Nodes_Node_Nps) GetParentYangName() string { return "node" }

// HardwareModuleNp_Nodes_Node_Nps_Np
// np0 to np7
type HardwareModuleNp_Nodes_Node_Nps_Np struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NP name. The type is string with pattern:
    // (np0)|(np1)|(np2)|(np3)|(np4)|(np5)|(np6)|(np7).
    NpName interface{}

    // prm channel load info.
    ChnLoad HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad

    // prm tcam summary info.
    TcamSummary HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary

    // prm counters info.
    Counters HardwareModuleNp_Nodes_Node_Nps_Np_Counters

    // prm fast drop counters info.
    FastDrop HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetFilter() yfilter.YFilter { return np.YFilter }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) SetFilter(yf yfilter.YFilter) { np.YFilter = yf }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetGoName(yname string) string {
    if yname == "np-name" { return "NpName" }
    if yname == "chn-load" { return "ChnLoad" }
    if yname == "tcam-summary" { return "TcamSummary" }
    if yname == "counters" { return "Counters" }
    if yname == "fast-drop" { return "FastDrop" }
    return ""
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetSegmentPath() string {
    return "np" + "[np-name='" + fmt.Sprintf("%v", np.NpName) + "']"
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "chn-load" {
        return &np.ChnLoad
    }
    if childYangName == "tcam-summary" {
        return &np.TcamSummary
    }
    if childYangName == "counters" {
        return &np.Counters
    }
    if childYangName == "fast-drop" {
        return &np.FastDrop
    }
    return nil
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["chn-load"] = &np.ChnLoad
    children["tcam-summary"] = &np.TcamSummary
    children["counters"] = &np.Counters
    children["fast-drop"] = &np.FastDrop
    return children
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["np-name"] = np.NpName
    return leafs
}

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetBundleName() string { return "cisco_ios_xr" }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetYangName() string { return "np" }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) SetParent(parent types.Entity) { np.parent = parent }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetParent() types.Entity { return np.parent }

func (np *HardwareModuleNp_Nodes_Node_Nps_Np) GetParentYangName() string { return "nps" }

// HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad
// prm channel load info
type HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array of NP Channel load counters. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad.
    NpChnLoad []HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetFilter() yfilter.YFilter { return chnLoad.YFilter }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) SetFilter(yf yfilter.YFilter) { chnLoad.YFilter = yf }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetGoName(yname string) string {
    if yname == "np-chn-load" { return "NpChnLoad" }
    return ""
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetSegmentPath() string {
    return "chn-load"
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np-chn-load" {
        for _, c := range chnLoad.NpChnLoad {
            if chnLoad.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad{}
        chnLoad.NpChnLoad = append(chnLoad.NpChnLoad, child)
        return &chnLoad.NpChnLoad[len(chnLoad.NpChnLoad)-1]
    }
    return nil
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range chnLoad.NpChnLoad {
        children[chnLoad.NpChnLoad[i].GetSegmentPath()] = &chnLoad.NpChnLoad[i]
    }
    return children
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetBundleName() string { return "cisco_ios_xr" }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetYangName() string { return "chn-load" }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) SetParent(parent types.Entity) { chnLoad.parent = parent }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetParent() types.Entity { return chnLoad.parent }

func (chnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad) GetParentYangName() string { return "np" }

// HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad
// Array of NP Channel load counters
type HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetFilter() yfilter.YFilter { return npChnLoad.YFilter }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) SetFilter(yf yfilter.YFilter) { npChnLoad.YFilter = yf }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetGoName(yname string) string {
    if yname == "flow-ctr-counter" { return "FlowCtrCounter" }
    if yname == "avg-rfd-usage" { return "AvgRfdUsage" }
    if yname == "peak-rfd-usage" { return "PeakRfdUsage" }
    if yname == "avg-guar-rfd-usage" { return "AvgGuarRfdUsage" }
    if yname == "peak-guar-rfd-usage" { return "PeakGuarRfdUsage" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetSegmentPath() string {
    return "np-chn-load"
}

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-ctr-counter"] = npChnLoad.FlowCtrCounter
    leafs["avg-rfd-usage"] = npChnLoad.AvgRfdUsage
    leafs["peak-rfd-usage"] = npChnLoad.PeakRfdUsage
    leafs["avg-guar-rfd-usage"] = npChnLoad.AvgGuarRfdUsage
    leafs["peak-guar-rfd-usage"] = npChnLoad.PeakGuarRfdUsage
    leafs["interface-name"] = npChnLoad.InterfaceName
    return leafs
}

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetBundleName() string { return "cisco_ios_xr" }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetYangName() string { return "np-chn-load" }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) SetParent(parent types.Entity) { npChnLoad.parent = parent }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetParent() types.Entity { return npChnLoad.parent }

func (npChnLoad *HardwareModuleNp_Nodes_Node_Nps_Np_ChnLoad_NpChnLoad) GetParentYangName() string { return "chn-load" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary
// prm tcam summary info
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Internal tcam summary info.
    InternalTcamInfo HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo

    // External tcam summary info.
    TcamInfo HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetFilter() yfilter.YFilter { return tcamSummary.YFilter }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) SetFilter(yf yfilter.YFilter) { tcamSummary.YFilter = yf }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetGoName(yname string) string {
    if yname == "internal-tcam-info" { return "InternalTcamInfo" }
    if yname == "tcam-info" { return "TcamInfo" }
    return ""
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetSegmentPath() string {
    return "tcam-summary"
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "internal-tcam-info" {
        return &tcamSummary.InternalTcamInfo
    }
    if childYangName == "tcam-info" {
        return &tcamSummary.TcamInfo
    }
    return nil
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["internal-tcam-info"] = &tcamSummary.InternalTcamInfo
    children["tcam-info"] = &tcamSummary.TcamInfo
    return children
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetBundleName() string { return "cisco_ios_xr" }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetYangName() string { return "tcam-summary" }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) SetParent(parent types.Entity) { tcamSummary.parent = parent }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetParent() types.Entity { return tcamSummary.parent }

func (tcamSummary *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary) GetParentYangName() string { return "np" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo
// Internal tcam summary info
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCAM LT ODS 2 summary.
    TcamLtOds2 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2

    // TCAM LT_ODS 8 summary.
    TcamLtOds8 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8

    // Array of TCAM LT L2 partition summaries. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2.
    TcamLtL2 []HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetFilter() yfilter.YFilter { return internalTcamInfo.YFilter }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) SetFilter(yf yfilter.YFilter) { internalTcamInfo.YFilter = yf }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetGoName(yname string) string {
    if yname == "tcam-lt-ods2" { return "TcamLtOds2" }
    if yname == "tcam-lt-ods8" { return "TcamLtOds8" }
    if yname == "tcam-lt-l2" { return "TcamLtL2" }
    return ""
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetSegmentPath() string {
    return "internal-tcam-info"
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcam-lt-ods2" {
        return &internalTcamInfo.TcamLtOds2
    }
    if childYangName == "tcam-lt-ods8" {
        return &internalTcamInfo.TcamLtOds8
    }
    if childYangName == "tcam-lt-l2" {
        for _, c := range internalTcamInfo.TcamLtL2 {
            if internalTcamInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2{}
        internalTcamInfo.TcamLtL2 = append(internalTcamInfo.TcamLtL2, child)
        return &internalTcamInfo.TcamLtL2[len(internalTcamInfo.TcamLtL2)-1]
    }
    return nil
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcam-lt-ods2"] = &internalTcamInfo.TcamLtOds2
    children["tcam-lt-ods8"] = &internalTcamInfo.TcamLtOds8
    for i := range internalTcamInfo.TcamLtL2 {
        children[internalTcamInfo.TcamLtL2[i].GetSegmentPath()] = &internalTcamInfo.TcamLtL2[i]
    }
    return children
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetBundleName() string { return "cisco_ios_xr" }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetYangName() string { return "internal-tcam-info" }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) SetParent(parent types.Entity) { internalTcamInfo.parent = parent }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetParent() types.Entity { return internalTcamInfo.parent }

func (internalTcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo) GetParentYangName() string { return "tcam-summary" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2
// TCAM LT ODS 2 summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2 struct {
    parent types.Entity
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

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetFilter() yfilter.YFilter { return tcamLtOds2.YFilter }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) SetFilter(yf yfilter.YFilter) { tcamLtOds2.YFilter = yf }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetGoName(yname string) string {
    if yname == "max-entries" { return "MaxEntries" }
    if yname == "free-entries" { return "FreeEntries" }
    if yname == "app-id-ifib" { return "AppIdIfib" }
    if yname == "app-id-qos" { return "AppIdQos" }
    if yname == "app-id-acl" { return "AppIdAcl" }
    if yname == "app-id-afmon" { return "AppIdAfmon" }
    if yname == "app-id-li" { return "AppIdLi" }
    if yname == "app-id-pbr" { return "AppIdPbr" }
    if yname == "application-edpl-entry" { return "ApplicationEdplEntry" }
    return ""
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetSegmentPath() string {
    return "tcam-lt-ods2"
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "app-id-ifib" {
        return &tcamLtOds2.AppIdIfib
    }
    if childYangName == "app-id-qos" {
        return &tcamLtOds2.AppIdQos
    }
    if childYangName == "app-id-acl" {
        return &tcamLtOds2.AppIdAcl
    }
    if childYangName == "app-id-afmon" {
        return &tcamLtOds2.AppIdAfmon
    }
    if childYangName == "app-id-li" {
        return &tcamLtOds2.AppIdLi
    }
    if childYangName == "app-id-pbr" {
        return &tcamLtOds2.AppIdPbr
    }
    if childYangName == "application-edpl-entry" {
        return &tcamLtOds2.ApplicationEdplEntry
    }
    return nil
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["app-id-ifib"] = &tcamLtOds2.AppIdIfib
    children["app-id-qos"] = &tcamLtOds2.AppIdQos
    children["app-id-acl"] = &tcamLtOds2.AppIdAcl
    children["app-id-afmon"] = &tcamLtOds2.AppIdAfmon
    children["app-id-li"] = &tcamLtOds2.AppIdLi
    children["app-id-pbr"] = &tcamLtOds2.AppIdPbr
    children["application-edpl-entry"] = &tcamLtOds2.ApplicationEdplEntry
    return children
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-entries"] = tcamLtOds2.MaxEntries
    leafs["free-entries"] = tcamLtOds2.FreeEntries
    return leafs
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetBundleName() string { return "cisco_ios_xr" }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetYangName() string { return "tcam-lt-ods2" }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) SetParent(parent types.Entity) { tcamLtOds2.parent = parent }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetParent() types.Entity { return tcamLtOds2.parent }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2) GetParentYangName() string { return "internal-tcam-info" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib struct {
    parent types.Entity
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

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetFilter() yfilter.YFilter { return appIdIfib.YFilter }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) SetFilter(yf yfilter.YFilter) { appIdIfib.YFilter = yf }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetSegmentPath() string {
    return "app-id-ifib"
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdIfib.NumVmrIds
    leafs["total-used-entries"] = appIdIfib.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdIfib.TotalAllocatedEntries
    return leafs
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetBundleName() string { return "cisco_ios_xr" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetYangName() string { return "app-id-ifib" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) SetParent(parent types.Entity) { appIdIfib.parent = parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetParent() types.Entity { return appIdIfib.parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdIfib) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos struct {
    parent types.Entity
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

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetFilter() yfilter.YFilter { return appIdQos.YFilter }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) SetFilter(yf yfilter.YFilter) { appIdQos.YFilter = yf }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetSegmentPath() string {
    return "app-id-qos"
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdQos.NumVmrIds
    leafs["total-used-entries"] = appIdQos.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdQos.TotalAllocatedEntries
    return leafs
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetBundleName() string { return "cisco_ios_xr" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetYangName() string { return "app-id-qos" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) SetParent(parent types.Entity) { appIdQos.parent = parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetParent() types.Entity { return appIdQos.parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdQos) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl struct {
    parent types.Entity
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

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetFilter() yfilter.YFilter { return appIdAcl.YFilter }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) SetFilter(yf yfilter.YFilter) { appIdAcl.YFilter = yf }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetSegmentPath() string {
    return "app-id-acl"
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAcl.NumVmrIds
    leafs["total-used-entries"] = appIdAcl.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdAcl.TotalAllocatedEntries
    return leafs
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetYangName() string { return "app-id-acl" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) SetParent(parent types.Entity) { appIdAcl.parent = parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetParent() types.Entity { return appIdAcl.parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAcl) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon struct {
    parent types.Entity
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

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetFilter() yfilter.YFilter { return appIdAfmon.YFilter }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) SetFilter(yf yfilter.YFilter) { appIdAfmon.YFilter = yf }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetSegmentPath() string {
    return "app-id-afmon"
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAfmon.NumVmrIds
    leafs["total-used-entries"] = appIdAfmon.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdAfmon.TotalAllocatedEntries
    return leafs
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetYangName() string { return "app-id-afmon" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) SetParent(parent types.Entity) { appIdAfmon.parent = parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetParent() types.Entity { return appIdAfmon.parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdAfmon) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi struct {
    parent types.Entity
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

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetFilter() yfilter.YFilter { return appIdLi.YFilter }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) SetFilter(yf yfilter.YFilter) { appIdLi.YFilter = yf }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetSegmentPath() string {
    return "app-id-li"
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdLi.NumVmrIds
    leafs["total-used-entries"] = appIdLi.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdLi.TotalAllocatedEntries
    return leafs
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetBundleName() string { return "cisco_ios_xr" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetYangName() string { return "app-id-li" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) SetParent(parent types.Entity) { appIdLi.parent = parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetParent() types.Entity { return appIdLi.parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdLi) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr struct {
    parent types.Entity
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

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetFilter() yfilter.YFilter { return appIdPbr.YFilter }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) SetFilter(yf yfilter.YFilter) { appIdPbr.YFilter = yf }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetSegmentPath() string {
    return "app-id-pbr"
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdPbr.NumVmrIds
    leafs["total-used-entries"] = appIdPbr.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdPbr.TotalAllocatedEntries
    return leafs
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetBundleName() string { return "cisco_ios_xr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetYangName() string { return "app-id-pbr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) SetParent(parent types.Entity) { appIdPbr.parent = parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetParent() types.Entity { return appIdPbr.parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_AppIdPbr) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry struct {
    parent types.Entity
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

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetFilter() yfilter.YFilter { return applicationEdplEntry.YFilter }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) SetFilter(yf yfilter.YFilter) { applicationEdplEntry.YFilter = yf }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetSegmentPath() string {
    return "application-edpl-entry"
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = applicationEdplEntry.NumVmrIds
    leafs["total-used-entries"] = applicationEdplEntry.TotalUsedEntries
    leafs["total-allocated-entries"] = applicationEdplEntry.TotalAllocatedEntries
    return leafs
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetBundleName() string { return "cisco_ios_xr" }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetYangName() string { return "application-edpl-entry" }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) SetParent(parent types.Entity) { applicationEdplEntry.parent = parent }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetParent() types.Entity { return applicationEdplEntry.parent }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds2_ApplicationEdplEntry) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8
// TCAM LT_ODS 8 summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8 struct {
    parent types.Entity
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

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetFilter() yfilter.YFilter { return tcamLtOds8.YFilter }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) SetFilter(yf yfilter.YFilter) { tcamLtOds8.YFilter = yf }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetGoName(yname string) string {
    if yname == "max-entries" { return "MaxEntries" }
    if yname == "free-entries" { return "FreeEntries" }
    if yname == "app-id-ifib" { return "AppIdIfib" }
    if yname == "app-id-qos" { return "AppIdQos" }
    if yname == "app-id-acl" { return "AppIdAcl" }
    if yname == "app-id-afmon" { return "AppIdAfmon" }
    if yname == "app-id-li" { return "AppIdLi" }
    if yname == "app-id-pbr" { return "AppIdPbr" }
    if yname == "application-edpl-entry" { return "ApplicationEdplEntry" }
    return ""
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetSegmentPath() string {
    return "tcam-lt-ods8"
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "app-id-ifib" {
        return &tcamLtOds8.AppIdIfib
    }
    if childYangName == "app-id-qos" {
        return &tcamLtOds8.AppIdQos
    }
    if childYangName == "app-id-acl" {
        return &tcamLtOds8.AppIdAcl
    }
    if childYangName == "app-id-afmon" {
        return &tcamLtOds8.AppIdAfmon
    }
    if childYangName == "app-id-li" {
        return &tcamLtOds8.AppIdLi
    }
    if childYangName == "app-id-pbr" {
        return &tcamLtOds8.AppIdPbr
    }
    if childYangName == "application-edpl-entry" {
        return &tcamLtOds8.ApplicationEdplEntry
    }
    return nil
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["app-id-ifib"] = &tcamLtOds8.AppIdIfib
    children["app-id-qos"] = &tcamLtOds8.AppIdQos
    children["app-id-acl"] = &tcamLtOds8.AppIdAcl
    children["app-id-afmon"] = &tcamLtOds8.AppIdAfmon
    children["app-id-li"] = &tcamLtOds8.AppIdLi
    children["app-id-pbr"] = &tcamLtOds8.AppIdPbr
    children["application-edpl-entry"] = &tcamLtOds8.ApplicationEdplEntry
    return children
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-entries"] = tcamLtOds8.MaxEntries
    leafs["free-entries"] = tcamLtOds8.FreeEntries
    return leafs
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetBundleName() string { return "cisco_ios_xr" }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetYangName() string { return "tcam-lt-ods8" }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) SetParent(parent types.Entity) { tcamLtOds8.parent = parent }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetParent() types.Entity { return tcamLtOds8.parent }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8) GetParentYangName() string { return "internal-tcam-info" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib struct {
    parent types.Entity
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

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetFilter() yfilter.YFilter { return appIdIfib.YFilter }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) SetFilter(yf yfilter.YFilter) { appIdIfib.YFilter = yf }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetSegmentPath() string {
    return "app-id-ifib"
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdIfib.NumVmrIds
    leafs["total-used-entries"] = appIdIfib.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdIfib.TotalAllocatedEntries
    return leafs
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetBundleName() string { return "cisco_ios_xr" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetYangName() string { return "app-id-ifib" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) SetParent(parent types.Entity) { appIdIfib.parent = parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetParent() types.Entity { return appIdIfib.parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdIfib) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos struct {
    parent types.Entity
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

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetFilter() yfilter.YFilter { return appIdQos.YFilter }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) SetFilter(yf yfilter.YFilter) { appIdQos.YFilter = yf }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetSegmentPath() string {
    return "app-id-qos"
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdQos.NumVmrIds
    leafs["total-used-entries"] = appIdQos.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdQos.TotalAllocatedEntries
    return leafs
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetBundleName() string { return "cisco_ios_xr" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetYangName() string { return "app-id-qos" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) SetParent(parent types.Entity) { appIdQos.parent = parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetParent() types.Entity { return appIdQos.parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdQos) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl struct {
    parent types.Entity
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

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetFilter() yfilter.YFilter { return appIdAcl.YFilter }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) SetFilter(yf yfilter.YFilter) { appIdAcl.YFilter = yf }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetSegmentPath() string {
    return "app-id-acl"
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAcl.NumVmrIds
    leafs["total-used-entries"] = appIdAcl.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdAcl.TotalAllocatedEntries
    return leafs
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetYangName() string { return "app-id-acl" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) SetParent(parent types.Entity) { appIdAcl.parent = parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetParent() types.Entity { return appIdAcl.parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAcl) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon struct {
    parent types.Entity
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

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetFilter() yfilter.YFilter { return appIdAfmon.YFilter }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) SetFilter(yf yfilter.YFilter) { appIdAfmon.YFilter = yf }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetSegmentPath() string {
    return "app-id-afmon"
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAfmon.NumVmrIds
    leafs["total-used-entries"] = appIdAfmon.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdAfmon.TotalAllocatedEntries
    return leafs
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetYangName() string { return "app-id-afmon" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) SetParent(parent types.Entity) { appIdAfmon.parent = parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetParent() types.Entity { return appIdAfmon.parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdAfmon) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi struct {
    parent types.Entity
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

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetFilter() yfilter.YFilter { return appIdLi.YFilter }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) SetFilter(yf yfilter.YFilter) { appIdLi.YFilter = yf }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetSegmentPath() string {
    return "app-id-li"
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdLi.NumVmrIds
    leafs["total-used-entries"] = appIdLi.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdLi.TotalAllocatedEntries
    return leafs
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetBundleName() string { return "cisco_ios_xr" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetYangName() string { return "app-id-li" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) SetParent(parent types.Entity) { appIdLi.parent = parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetParent() types.Entity { return appIdLi.parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdLi) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr struct {
    parent types.Entity
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

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetFilter() yfilter.YFilter { return appIdPbr.YFilter }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) SetFilter(yf yfilter.YFilter) { appIdPbr.YFilter = yf }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetSegmentPath() string {
    return "app-id-pbr"
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdPbr.NumVmrIds
    leafs["total-used-entries"] = appIdPbr.TotalUsedEntries
    leafs["total-allocated-entries"] = appIdPbr.TotalAllocatedEntries
    return leafs
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetBundleName() string { return "cisco_ios_xr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetYangName() string { return "app-id-pbr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) SetParent(parent types.Entity) { appIdPbr.parent = parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetParent() types.Entity { return appIdPbr.parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_AppIdPbr) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry struct {
    parent types.Entity
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

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetFilter() yfilter.YFilter { return applicationEdplEntry.YFilter }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) SetFilter(yf yfilter.YFilter) { applicationEdplEntry.YFilter = yf }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "total-used-entries" { return "TotalUsedEntries" }
    if yname == "total-allocated-entries" { return "TotalAllocatedEntries" }
    return ""
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetSegmentPath() string {
    return "application-edpl-entry"
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = applicationEdplEntry.NumVmrIds
    leafs["total-used-entries"] = applicationEdplEntry.TotalUsedEntries
    leafs["total-allocated-entries"] = applicationEdplEntry.TotalAllocatedEntries
    return leafs
}

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetBundleName() string { return "cisco_ios_xr" }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetYangName() string { return "application-edpl-entry" }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) SetParent(parent types.Entity) { applicationEdplEntry.parent = parent }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetParent() types.Entity { return applicationEdplEntry.parent }

func (applicationEdplEntry *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtOds8_ApplicationEdplEntry) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2
// Array of TCAM LT L2 partition summaries
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PartitionID. The type is interface{} with range: 0..4294967295.
    PartitionId interface{}

    // Valid Entries. The type is interface{} with range: 0..4294967295.
    ValidEntries interface{}

    // Free Entries. The type is interface{} with range: 0..4294967295.
    FreeEntries interface{}
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetFilter() yfilter.YFilter { return tcamLtL2.YFilter }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) SetFilter(yf yfilter.YFilter) { tcamLtL2.YFilter = yf }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetGoName(yname string) string {
    if yname == "partition-id" { return "PartitionId" }
    if yname == "valid-entries" { return "ValidEntries" }
    if yname == "free-entries" { return "FreeEntries" }
    return ""
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetSegmentPath() string {
    return "tcam-lt-l2"
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["partition-id"] = tcamLtL2.PartitionId
    leafs["valid-entries"] = tcamLtL2.ValidEntries
    leafs["free-entries"] = tcamLtL2.FreeEntries
    return leafs
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetBundleName() string { return "cisco_ios_xr" }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetYangName() string { return "tcam-lt-l2" }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) SetParent(parent types.Entity) { tcamLtL2.parent = parent }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetParent() types.Entity { return tcamLtL2.parent }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_InternalTcamInfo_TcamLtL2) GetParentYangName() string { return "internal-tcam-info" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo
// External tcam summary info
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCAM ODS2 partition summary.
    TcamLtOds2 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2

    // TCAM ODS8 partition summary.
    TcamLtOds8 HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8

    // Array of TCAM L2 partition summaries. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2.
    TcamLtL2 []HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetFilter() yfilter.YFilter { return tcamInfo.YFilter }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) SetFilter(yf yfilter.YFilter) { tcamInfo.YFilter = yf }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetGoName(yname string) string {
    if yname == "tcam-lt-ods2" { return "TcamLtOds2" }
    if yname == "tcam-lt-ods8" { return "TcamLtOds8" }
    if yname == "tcam-lt-l2" { return "TcamLtL2" }
    return ""
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetSegmentPath() string {
    return "tcam-info"
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcam-lt-ods2" {
        return &tcamInfo.TcamLtOds2
    }
    if childYangName == "tcam-lt-ods8" {
        return &tcamInfo.TcamLtOds8
    }
    if childYangName == "tcam-lt-l2" {
        for _, c := range tcamInfo.TcamLtL2 {
            if tcamInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2{}
        tcamInfo.TcamLtL2 = append(tcamInfo.TcamLtL2, child)
        return &tcamInfo.TcamLtL2[len(tcamInfo.TcamLtL2)-1]
    }
    return nil
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tcam-lt-ods2"] = &tcamInfo.TcamLtOds2
    children["tcam-lt-ods8"] = &tcamInfo.TcamLtOds8
    for i := range tcamInfo.TcamLtL2 {
        children[tcamInfo.TcamLtL2[i].GetSegmentPath()] = &tcamInfo.TcamLtL2[i]
    }
    return children
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetBundleName() string { return "cisco_ios_xr" }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetYangName() string { return "tcam-info" }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) SetParent(parent types.Entity) { tcamInfo.parent = parent }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetParent() types.Entity { return tcamInfo.parent }

func (tcamInfo *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo) GetParentYangName() string { return "tcam-summary" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2
// TCAM ODS2 partition summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2 struct {
    parent types.Entity
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

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetFilter() yfilter.YFilter { return tcamLtOds2.YFilter }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) SetFilter(yf yfilter.YFilter) { tcamLtOds2.YFilter = yf }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetGoName(yname string) string {
    if yname == "free-entries" { return "FreeEntries" }
    if yname == "reserved-entries" { return "ReservedEntries" }
    if yname == "acl-common" { return "AclCommon" }
    if yname == "app-id-ifib" { return "AppIdIfib" }
    if yname == "app-id-qos" { return "AppIdQos" }
    if yname == "app-id-acl" { return "AppIdAcl" }
    if yname == "app-id-afmon" { return "AppIdAfmon" }
    if yname == "app-id-li" { return "AppIdLi" }
    if yname == "app-id-pbr" { return "AppIdPbr" }
    if yname == "app-id-edpl" { return "AppIdEdpl" }
    return ""
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetSegmentPath() string {
    return "tcam-lt-ods2"
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-common" {
        return &tcamLtOds2.AclCommon
    }
    if childYangName == "app-id-ifib" {
        return &tcamLtOds2.AppIdIfib
    }
    if childYangName == "app-id-qos" {
        return &tcamLtOds2.AppIdQos
    }
    if childYangName == "app-id-acl" {
        return &tcamLtOds2.AppIdAcl
    }
    if childYangName == "app-id-afmon" {
        return &tcamLtOds2.AppIdAfmon
    }
    if childYangName == "app-id-li" {
        return &tcamLtOds2.AppIdLi
    }
    if childYangName == "app-id-pbr" {
        return &tcamLtOds2.AppIdPbr
    }
    if childYangName == "app-id-edpl" {
        return &tcamLtOds2.AppIdEdpl
    }
    return nil
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl-common"] = &tcamLtOds2.AclCommon
    children["app-id-ifib"] = &tcamLtOds2.AppIdIfib
    children["app-id-qos"] = &tcamLtOds2.AppIdQos
    children["app-id-acl"] = &tcamLtOds2.AppIdAcl
    children["app-id-afmon"] = &tcamLtOds2.AppIdAfmon
    children["app-id-li"] = &tcamLtOds2.AppIdLi
    children["app-id-pbr"] = &tcamLtOds2.AppIdPbr
    children["app-id-edpl"] = &tcamLtOds2.AppIdEdpl
    return children
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["free-entries"] = tcamLtOds2.FreeEntries
    leafs["reserved-entries"] = tcamLtOds2.ReservedEntries
    return leafs
}

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetBundleName() string { return "cisco_ios_xr" }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetYangName() string { return "tcam-lt-ods2" }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) SetParent(parent types.Entity) { tcamLtOds2.parent = parent }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetParent() types.Entity { return tcamLtOds2.parent }

func (tcamLtOds2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2) GetParentYangName() string { return "tcam-info" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon
// ACL common region
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Free entries in the table. The type is interface{} with range:
    // 0..4294967295.
    FreeEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    AllocatedEntries interface{}
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetFilter() yfilter.YFilter { return aclCommon.YFilter }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) SetFilter(yf yfilter.YFilter) { aclCommon.YFilter = yf }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetGoName(yname string) string {
    if yname == "free-entries" { return "FreeEntries" }
    if yname == "allocated-entries" { return "AllocatedEntries" }
    return ""
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetSegmentPath() string {
    return "acl-common"
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["free-entries"] = aclCommon.FreeEntries
    leafs["allocated-entries"] = aclCommon.AllocatedEntries
    return leafs
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetBundleName() string { return "cisco_ios_xr" }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetYangName() string { return "acl-common" }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) SetParent(parent types.Entity) { aclCommon.parent = parent }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetParent() types.Entity { return aclCommon.parent }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AclCommon) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib struct {
    parent types.Entity
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

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetFilter() yfilter.YFilter { return appIdIfib.YFilter }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) SetFilter(yf yfilter.YFilter) { appIdIfib.YFilter = yf }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetSegmentPath() string {
    return "app-id-ifib"
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdIfib.NumVmrIds
    leafs["num-active-entries"] = appIdIfib.NumActiveEntries
    leafs["num-allocated-entries"] = appIdIfib.NumAllocatedEntries
    return leafs
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetBundleName() string { return "cisco_ios_xr" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetYangName() string { return "app-id-ifib" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) SetParent(parent types.Entity) { appIdIfib.parent = parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetParent() types.Entity { return appIdIfib.parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdIfib) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos struct {
    parent types.Entity
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

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetFilter() yfilter.YFilter { return appIdQos.YFilter }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) SetFilter(yf yfilter.YFilter) { appIdQos.YFilter = yf }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetSegmentPath() string {
    return "app-id-qos"
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdQos.NumVmrIds
    leafs["num-active-entries"] = appIdQos.NumActiveEntries
    leafs["num-allocated-entries"] = appIdQos.NumAllocatedEntries
    return leafs
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetBundleName() string { return "cisco_ios_xr" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetYangName() string { return "app-id-qos" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) SetParent(parent types.Entity) { appIdQos.parent = parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetParent() types.Entity { return appIdQos.parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdQos) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl struct {
    parent types.Entity
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

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetFilter() yfilter.YFilter { return appIdAcl.YFilter }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) SetFilter(yf yfilter.YFilter) { appIdAcl.YFilter = yf }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetSegmentPath() string {
    return "app-id-acl"
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAcl.NumVmrIds
    leafs["num-active-entries"] = appIdAcl.NumActiveEntries
    leafs["num-allocated-entries"] = appIdAcl.NumAllocatedEntries
    return leafs
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetYangName() string { return "app-id-acl" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) SetParent(parent types.Entity) { appIdAcl.parent = parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetParent() types.Entity { return appIdAcl.parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAcl) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon struct {
    parent types.Entity
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

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetFilter() yfilter.YFilter { return appIdAfmon.YFilter }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) SetFilter(yf yfilter.YFilter) { appIdAfmon.YFilter = yf }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetSegmentPath() string {
    return "app-id-afmon"
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAfmon.NumVmrIds
    leafs["num-active-entries"] = appIdAfmon.NumActiveEntries
    leafs["num-allocated-entries"] = appIdAfmon.NumAllocatedEntries
    return leafs
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetYangName() string { return "app-id-afmon" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) SetParent(parent types.Entity) { appIdAfmon.parent = parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetParent() types.Entity { return appIdAfmon.parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdAfmon) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi struct {
    parent types.Entity
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

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetFilter() yfilter.YFilter { return appIdLi.YFilter }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) SetFilter(yf yfilter.YFilter) { appIdLi.YFilter = yf }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetSegmentPath() string {
    return "app-id-li"
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdLi.NumVmrIds
    leafs["num-active-entries"] = appIdLi.NumActiveEntries
    leafs["num-allocated-entries"] = appIdLi.NumAllocatedEntries
    return leafs
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetBundleName() string { return "cisco_ios_xr" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetYangName() string { return "app-id-li" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) SetParent(parent types.Entity) { appIdLi.parent = parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetParent() types.Entity { return appIdLi.parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdLi) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr struct {
    parent types.Entity
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

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetFilter() yfilter.YFilter { return appIdPbr.YFilter }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) SetFilter(yf yfilter.YFilter) { appIdPbr.YFilter = yf }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetSegmentPath() string {
    return "app-id-pbr"
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdPbr.NumVmrIds
    leafs["num-active-entries"] = appIdPbr.NumActiveEntries
    leafs["num-allocated-entries"] = appIdPbr.NumAllocatedEntries
    return leafs
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetBundleName() string { return "cisco_ios_xr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetYangName() string { return "app-id-pbr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) SetParent(parent types.Entity) { appIdPbr.parent = parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetParent() types.Entity { return appIdPbr.parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdPbr) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl struct {
    parent types.Entity
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

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetFilter() yfilter.YFilter { return appIdEdpl.YFilter }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) SetFilter(yf yfilter.YFilter) { appIdEdpl.YFilter = yf }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetSegmentPath() string {
    return "app-id-edpl"
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdEdpl.NumVmrIds
    leafs["num-active-entries"] = appIdEdpl.NumActiveEntries
    leafs["num-allocated-entries"] = appIdEdpl.NumAllocatedEntries
    return leafs
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetBundleName() string { return "cisco_ios_xr" }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetYangName() string { return "app-id-edpl" }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) SetParent(parent types.Entity) { appIdEdpl.parent = parent }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetParent() types.Entity { return appIdEdpl.parent }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds2_AppIdEdpl) GetParentYangName() string { return "tcam-lt-ods2" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8
// TCAM ODS8 partition summary
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8 struct {
    parent types.Entity
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

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetFilter() yfilter.YFilter { return tcamLtOds8.YFilter }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) SetFilter(yf yfilter.YFilter) { tcamLtOds8.YFilter = yf }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetGoName(yname string) string {
    if yname == "free-entries" { return "FreeEntries" }
    if yname == "reserved-entries" { return "ReservedEntries" }
    if yname == "acl-common" { return "AclCommon" }
    if yname == "app-id-ifib" { return "AppIdIfib" }
    if yname == "app-id-qos" { return "AppIdQos" }
    if yname == "app-id-acl" { return "AppIdAcl" }
    if yname == "app-id-afmon" { return "AppIdAfmon" }
    if yname == "app-id-li" { return "AppIdLi" }
    if yname == "app-id-pbr" { return "AppIdPbr" }
    if yname == "app-id-edpl" { return "AppIdEdpl" }
    return ""
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetSegmentPath() string {
    return "tcam-lt-ods8"
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-common" {
        return &tcamLtOds8.AclCommon
    }
    if childYangName == "app-id-ifib" {
        return &tcamLtOds8.AppIdIfib
    }
    if childYangName == "app-id-qos" {
        return &tcamLtOds8.AppIdQos
    }
    if childYangName == "app-id-acl" {
        return &tcamLtOds8.AppIdAcl
    }
    if childYangName == "app-id-afmon" {
        return &tcamLtOds8.AppIdAfmon
    }
    if childYangName == "app-id-li" {
        return &tcamLtOds8.AppIdLi
    }
    if childYangName == "app-id-pbr" {
        return &tcamLtOds8.AppIdPbr
    }
    if childYangName == "app-id-edpl" {
        return &tcamLtOds8.AppIdEdpl
    }
    return nil
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl-common"] = &tcamLtOds8.AclCommon
    children["app-id-ifib"] = &tcamLtOds8.AppIdIfib
    children["app-id-qos"] = &tcamLtOds8.AppIdQos
    children["app-id-acl"] = &tcamLtOds8.AppIdAcl
    children["app-id-afmon"] = &tcamLtOds8.AppIdAfmon
    children["app-id-li"] = &tcamLtOds8.AppIdLi
    children["app-id-pbr"] = &tcamLtOds8.AppIdPbr
    children["app-id-edpl"] = &tcamLtOds8.AppIdEdpl
    return children
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["free-entries"] = tcamLtOds8.FreeEntries
    leafs["reserved-entries"] = tcamLtOds8.ReservedEntries
    return leafs
}

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetBundleName() string { return "cisco_ios_xr" }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetYangName() string { return "tcam-lt-ods8" }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) SetParent(parent types.Entity) { tcamLtOds8.parent = parent }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetParent() types.Entity { return tcamLtOds8.parent }

func (tcamLtOds8 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8) GetParentYangName() string { return "tcam-info" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon
// ACL common region
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Free entries in the table. The type is interface{} with range:
    // 0..4294967295.
    FreeEntries interface{}

    // The number of active vmr entries. The type is interface{} with range:
    // 0..4294967295.
    AllocatedEntries interface{}
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetFilter() yfilter.YFilter { return aclCommon.YFilter }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) SetFilter(yf yfilter.YFilter) { aclCommon.YFilter = yf }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetGoName(yname string) string {
    if yname == "free-entries" { return "FreeEntries" }
    if yname == "allocated-entries" { return "AllocatedEntries" }
    return ""
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetSegmentPath() string {
    return "acl-common"
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["free-entries"] = aclCommon.FreeEntries
    leafs["allocated-entries"] = aclCommon.AllocatedEntries
    return leafs
}

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetBundleName() string { return "cisco_ios_xr" }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetYangName() string { return "acl-common" }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) SetParent(parent types.Entity) { aclCommon.parent = parent }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetParent() types.Entity { return aclCommon.parent }

func (aclCommon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AclCommon) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib
// app IFIB entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib struct {
    parent types.Entity
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

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetFilter() yfilter.YFilter { return appIdIfib.YFilter }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) SetFilter(yf yfilter.YFilter) { appIdIfib.YFilter = yf }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetSegmentPath() string {
    return "app-id-ifib"
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdIfib.NumVmrIds
    leafs["num-active-entries"] = appIdIfib.NumActiveEntries
    leafs["num-allocated-entries"] = appIdIfib.NumAllocatedEntries
    return leafs
}

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetBundleName() string { return "cisco_ios_xr" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetYangName() string { return "app-id-ifib" }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) SetParent(parent types.Entity) { appIdIfib.parent = parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetParent() types.Entity { return appIdIfib.parent }

func (appIdIfib *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdIfib) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos
// app qos entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos struct {
    parent types.Entity
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

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetFilter() yfilter.YFilter { return appIdQos.YFilter }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) SetFilter(yf yfilter.YFilter) { appIdQos.YFilter = yf }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetSegmentPath() string {
    return "app-id-qos"
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdQos.NumVmrIds
    leafs["num-active-entries"] = appIdQos.NumActiveEntries
    leafs["num-allocated-entries"] = appIdQos.NumAllocatedEntries
    return leafs
}

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetBundleName() string { return "cisco_ios_xr" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetYangName() string { return "app-id-qos" }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) SetParent(parent types.Entity) { appIdQos.parent = parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetParent() types.Entity { return appIdQos.parent }

func (appIdQos *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdQos) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl
// app acl entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl struct {
    parent types.Entity
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

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetFilter() yfilter.YFilter { return appIdAcl.YFilter }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) SetFilter(yf yfilter.YFilter) { appIdAcl.YFilter = yf }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetSegmentPath() string {
    return "app-id-acl"
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAcl.NumVmrIds
    leafs["num-active-entries"] = appIdAcl.NumActiveEntries
    leafs["num-allocated-entries"] = appIdAcl.NumAllocatedEntries
    return leafs
}

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetYangName() string { return "app-id-acl" }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) SetParent(parent types.Entity) { appIdAcl.parent = parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetParent() types.Entity { return appIdAcl.parent }

func (appIdAcl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAcl) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon
// app afmon entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon struct {
    parent types.Entity
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

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetFilter() yfilter.YFilter { return appIdAfmon.YFilter }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) SetFilter(yf yfilter.YFilter) { appIdAfmon.YFilter = yf }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetSegmentPath() string {
    return "app-id-afmon"
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdAfmon.NumVmrIds
    leafs["num-active-entries"] = appIdAfmon.NumActiveEntries
    leafs["num-allocated-entries"] = appIdAfmon.NumAllocatedEntries
    return leafs
}

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetBundleName() string { return "cisco_ios_xr" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetYangName() string { return "app-id-afmon" }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) SetParent(parent types.Entity) { appIdAfmon.parent = parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetParent() types.Entity { return appIdAfmon.parent }

func (appIdAfmon *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdAfmon) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi
// app LI entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi struct {
    parent types.Entity
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

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetFilter() yfilter.YFilter { return appIdLi.YFilter }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) SetFilter(yf yfilter.YFilter) { appIdLi.YFilter = yf }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetSegmentPath() string {
    return "app-id-li"
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdLi.NumVmrIds
    leafs["num-active-entries"] = appIdLi.NumActiveEntries
    leafs["num-allocated-entries"] = appIdLi.NumAllocatedEntries
    return leafs
}

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetBundleName() string { return "cisco_ios_xr" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetYangName() string { return "app-id-li" }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) SetParent(parent types.Entity) { appIdLi.parent = parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetParent() types.Entity { return appIdLi.parent }

func (appIdLi *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdLi) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr
// app PBR entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr struct {
    parent types.Entity
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

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetFilter() yfilter.YFilter { return appIdPbr.YFilter }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) SetFilter(yf yfilter.YFilter) { appIdPbr.YFilter = yf }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetSegmentPath() string {
    return "app-id-pbr"
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdPbr.NumVmrIds
    leafs["num-active-entries"] = appIdPbr.NumActiveEntries
    leafs["num-allocated-entries"] = appIdPbr.NumAllocatedEntries
    return leafs
}

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetBundleName() string { return "cisco_ios_xr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetYangName() string { return "app-id-pbr" }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) SetParent(parent types.Entity) { appIdPbr.parent = parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetParent() types.Entity { return appIdPbr.parent }

func (appIdPbr *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdPbr) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl
// app EDPL entry
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl struct {
    parent types.Entity
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

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetFilter() yfilter.YFilter { return appIdEdpl.YFilter }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) SetFilter(yf yfilter.YFilter) { appIdEdpl.YFilter = yf }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetGoName(yname string) string {
    if yname == "num-vmr-ids" { return "NumVmrIds" }
    if yname == "num-active-entries" { return "NumActiveEntries" }
    if yname == "num-allocated-entries" { return "NumAllocatedEntries" }
    return ""
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetSegmentPath() string {
    return "app-id-edpl"
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-vmr-ids"] = appIdEdpl.NumVmrIds
    leafs["num-active-entries"] = appIdEdpl.NumActiveEntries
    leafs["num-allocated-entries"] = appIdEdpl.NumAllocatedEntries
    return leafs
}

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetBundleName() string { return "cisco_ios_xr" }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetYangName() string { return "app-id-edpl" }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) SetParent(parent types.Entity) { appIdEdpl.parent = parent }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetParent() types.Entity { return appIdEdpl.parent }

func (appIdEdpl *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtOds8_AppIdEdpl) GetParentYangName() string { return "tcam-lt-ods8" }

// HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2
// Array of TCAM L2 partition summaries
type HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PartitionID. The type is interface{} with range: 0..4294967295.
    PartitionId interface{}

    // Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Valid Entries. The type is interface{} with range: 0..4294967295.
    ValidEntries interface{}

    // Free Entries. The type is interface{} with range: 0..4294967295.
    FreeEntries interface{}
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetFilter() yfilter.YFilter { return tcamLtL2.YFilter }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) SetFilter(yf yfilter.YFilter) { tcamLtL2.YFilter = yf }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetGoName(yname string) string {
    if yname == "partition-id" { return "PartitionId" }
    if yname == "priority" { return "Priority" }
    if yname == "valid-entries" { return "ValidEntries" }
    if yname == "free-entries" { return "FreeEntries" }
    return ""
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetSegmentPath() string {
    return "tcam-lt-l2"
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["partition-id"] = tcamLtL2.PartitionId
    leafs["priority"] = tcamLtL2.Priority
    leafs["valid-entries"] = tcamLtL2.ValidEntries
    leafs["free-entries"] = tcamLtL2.FreeEntries
    return leafs
}

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetBundleName() string { return "cisco_ios_xr" }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetYangName() string { return "tcam-lt-l2" }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) SetParent(parent types.Entity) { tcamLtL2.parent = parent }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetParent() types.Entity { return tcamLtL2.parent }

func (tcamLtL2 *HardwareModuleNp_Nodes_Node_Nps_Np_TcamSummary_TcamInfo_TcamLtL2) GetParentYangName() string { return "tcam-info" }

// HardwareModuleNp_Nodes_Node_Nps_Np_Counters
// prm counters info
type HardwareModuleNp_Nodes_Node_Nps_Np_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array of NP Counters. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter.
    NpCounter []HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetGoName(yname string) string {
    if yname == "np-counter" { return "NpCounter" }
    return ""
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np-counter" {
        for _, c := range counters.NpCounter {
            if counters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter{}
        counters.NpCounter = append(counters.NpCounter, child)
        return &counters.NpCounter[len(counters.NpCounter)-1]
    }
    return nil
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range counters.NpCounter {
        children[counters.NpCounter[i].GetSegmentPath()] = &counters.NpCounter[i]
    }
    return children
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetYangName() string { return "counters" }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetParent() types.Entity { return counters.parent }

func (counters *HardwareModuleNp_Nodes_Node_Nps_Np_Counters) GetParentYangName() string { return "np" }

// HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter
// Array of NP Counters
type HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetFilter() yfilter.YFilter { return npCounter.YFilter }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) SetFilter(yf yfilter.YFilter) { npCounter.YFilter = yf }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetGoName(yname string) string {
    if yname == "counter-index" { return "CounterIndex" }
    if yname == "counter-value" { return "CounterValue" }
    if yname == "rate" { return "Rate" }
    if yname == "counter-type" { return "CounterType" }
    if yname == "counter-name" { return "CounterName" }
    return ""
}

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetSegmentPath() string {
    return "np-counter"
}

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter-index"] = npCounter.CounterIndex
    leafs["counter-value"] = npCounter.CounterValue
    leafs["rate"] = npCounter.Rate
    leafs["counter-type"] = npCounter.CounterType
    leafs["counter-name"] = npCounter.CounterName
    return leafs
}

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetBundleName() string { return "cisco_ios_xr" }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetYangName() string { return "np-counter" }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) SetParent(parent types.Entity) { npCounter.parent = parent }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetParent() types.Entity { return npCounter.parent }

func (npCounter *HardwareModuleNp_Nodes_Node_Nps_Np_Counters_NpCounter) GetParentYangName() string { return "counters" }

// HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop
// prm fast drop counters info
type HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array of NP Fast Drop Counters. The type is slice of
    // HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop.
    NpFastDrop []HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetFilter() yfilter.YFilter { return fastDrop.YFilter }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) SetFilter(yf yfilter.YFilter) { fastDrop.YFilter = yf }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetGoName(yname string) string {
    if yname == "np-fast-drop" { return "NpFastDrop" }
    return ""
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetSegmentPath() string {
    return "fast-drop"
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np-fast-drop" {
        for _, c := range fastDrop.NpFastDrop {
            if fastDrop.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop{}
        fastDrop.NpFastDrop = append(fastDrop.NpFastDrop, child)
        return &fastDrop.NpFastDrop[len(fastDrop.NpFastDrop)-1]
    }
    return nil
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fastDrop.NpFastDrop {
        children[fastDrop.NpFastDrop[i].GetSegmentPath()] = &fastDrop.NpFastDrop[i]
    }
    return children
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetBundleName() string { return "cisco_ios_xr" }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetYangName() string { return "fast-drop" }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) SetParent(parent types.Entity) { fastDrop.parent = parent }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetParent() types.Entity { return fastDrop.parent }

func (fastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop) GetParentYangName() string { return "np" }

// HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop
// Array of NP Fast Drop Counters
type HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // The Value of the counter. The type is interface{} with range:
    // 0..18446744073709551615.
    CounterValue interface{}
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetFilter() yfilter.YFilter { return npFastDrop.YFilter }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) SetFilter(yf yfilter.YFilter) { npFastDrop.YFilter = yf }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "counter-value" { return "CounterValue" }
    return ""
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetSegmentPath() string {
    return "np-fast-drop"
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = npFastDrop.InterfaceName
    leafs["counter-value"] = npFastDrop.CounterValue
    return leafs
}

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetBundleName() string { return "cisco_ios_xr" }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetYangName() string { return "np-fast-drop" }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) SetParent(parent types.Entity) { npFastDrop.parent = parent }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetParent() types.Entity { return npFastDrop.parent }

func (npFastDrop *HardwareModuleNp_Nodes_Node_Nps_Np_FastDrop_NpFastDrop) GetParentYangName() string { return "fast-drop" }

