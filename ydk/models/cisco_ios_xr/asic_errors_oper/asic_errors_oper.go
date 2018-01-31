// This module contains a collection of YANG definitions
// for Cisco IOS-XR asic-errors package operational data.
// 
// This module contains definitions
// for the following management objects:
//   asic-errors: Error summary of all asics
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asic_errors_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asic_errors_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asic-errors-oper asic-errors}", reflect.TypeOf(AsicErrors{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asic-errors-oper:asic-errors", reflect.TypeOf(AsicErrors{}))
}

// AsicErrors
// Error summary of all asics
type AsicErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Asic errors for each available nodes.
    Nodes AsicErrors_Nodes
}

func (asicErrors *AsicErrors) GetFilter() yfilter.YFilter { return asicErrors.YFilter }

func (asicErrors *AsicErrors) SetFilter(yf yfilter.YFilter) { asicErrors.YFilter = yf }

func (asicErrors *AsicErrors) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (asicErrors *AsicErrors) GetSegmentPath() string {
    return "Cisco-IOS-XR-asic-errors-oper:asic-errors"
}

func (asicErrors *AsicErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &asicErrors.Nodes
    }
    return nil
}

func (asicErrors *AsicErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &asicErrors.Nodes
    return children
}

func (asicErrors *AsicErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrors *AsicErrors) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrors *AsicErrors) GetYangName() string { return "asic-errors" }

func (asicErrors *AsicErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrors *AsicErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrors *AsicErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrors *AsicErrors) SetParent(parent types.Entity) { asicErrors.parent = parent }

func (asicErrors *AsicErrors) GetParent() types.Entity { return asicErrors.parent }

func (asicErrors *AsicErrors) GetParentYangName() string { return "Cisco-IOS-XR-asic-errors-oper" }

// AsicErrors_Nodes
// Asic errors for each available nodes
type AsicErrors_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Asic error for a particular node. The type is slice of
    // AsicErrors_Nodes_Node.
    Node []AsicErrors_Nodes_Node
}

func (nodes *AsicErrors_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *AsicErrors_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *AsicErrors_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *AsicErrors_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *AsicErrors_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *AsicErrors_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *AsicErrors_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *AsicErrors_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *AsicErrors_Nodes) GetYangName() string { return "nodes" }

func (nodes *AsicErrors_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *AsicErrors_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *AsicErrors_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *AsicErrors_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *AsicErrors_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *AsicErrors_Nodes) GetParentYangName() string { return "asic-errors" }

// AsicErrors_Nodes_Node
// Asic error for a particular node
type AsicErrors_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Asic on the node. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation.
    AsicInformation []AsicErrors_Nodes_Node_AsicInformation
}

func (node *AsicErrors_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *AsicErrors_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *AsicErrors_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "asic-information" { return "AsicInformation" }
    return ""
}

func (node *AsicErrors_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *AsicErrors_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-information" {
        for _, c := range node.AsicInformation {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation{}
        node.AsicInformation = append(node.AsicInformation, child)
        return &node.AsicInformation[len(node.AsicInformation)-1]
    }
    return nil
}

func (node *AsicErrors_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.AsicInformation {
        children[node.AsicInformation[i].GetSegmentPath()] = &node.AsicInformation[i]
    }
    return children
}

func (node *AsicErrors_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *AsicErrors_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *AsicErrors_Nodes_Node) GetYangName() string { return "node" }

func (node *AsicErrors_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *AsicErrors_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *AsicErrors_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *AsicErrors_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *AsicErrors_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *AsicErrors_Nodes_Node) GetParentYangName() string { return "nodes" }

// AsicErrors_Nodes_Node_AsicInformation
// Asic on the node
type AsicErrors_Nodes_Node_AsicInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Asic string. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Asic interface{}

    // All asic instance on the node.
    AllInstances AsicErrors_Nodes_Node_AsicInformation_AllInstances

    // All asic errors  on the node.
    Instances AsicErrors_Nodes_Node_AsicInformation_Instances
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetFilter() yfilter.YFilter { return asicInformation.YFilter }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) SetFilter(yf yfilter.YFilter) { asicInformation.YFilter = yf }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetGoName(yname string) string {
    if yname == "asic" { return "Asic" }
    if yname == "all-instances" { return "AllInstances" }
    if yname == "instances" { return "Instances" }
    return ""
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetSegmentPath() string {
    return "asic-information" + "[asic='" + fmt.Sprintf("%v", asicInformation.Asic) + "']"
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-instances" {
        return &asicInformation.AllInstances
    }
    if childYangName == "instances" {
        return &asicInformation.Instances
    }
    return nil
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["all-instances"] = &asicInformation.AllInstances
    children["instances"] = &asicInformation.Instances
    return children
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asic"] = asicInformation.Asic
    return leafs
}

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetBundleName() string { return "cisco_ios_xr" }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetYangName() string { return "asic-information" }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) SetParent(parent types.Entity) { asicInformation.parent = parent }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetParent() types.Entity { return asicInformation.parent }

func (asicInformation *AsicErrors_Nodes_Node_AsicInformation) GetParentYangName() string { return "node" }

// AsicErrors_Nodes_Node_AsicInformation_AllInstances
// All asic instance on the node
type AsicErrors_Nodes_Node_AsicInformation_AllInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error path of all instances.
    AllErrorPath AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetFilter() yfilter.YFilter { return allInstances.YFilter }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) SetFilter(yf yfilter.YFilter) { allInstances.YFilter = yf }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetGoName(yname string) string {
    if yname == "all-error-path" { return "AllErrorPath" }
    return ""
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetSegmentPath() string {
    return "all-instances"
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-error-path" {
        return &allInstances.AllErrorPath
    }
    return nil
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["all-error-path"] = &allInstances.AllErrorPath
    return children
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetBundleName() string { return "cisco_ios_xr" }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetYangName() string { return "all-instances" }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) SetParent(parent types.Entity) { allInstances.parent = parent }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetParent() types.Entity { return allInstances.parent }

func (allInstances *AsicErrors_Nodes_Node_AsicInformation_AllInstances) GetParentYangName() string { return "asic-information" }

// AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath
// Error path of all instances
type AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary of all instances errors.
    Summary AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetFilter() yfilter.YFilter { return allErrorPath.YFilter }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) SetFilter(yf yfilter.YFilter) { allErrorPath.YFilter = yf }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    return ""
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetSegmentPath() string {
    return "all-error-path"
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &allErrorPath.Summary
    }
    return nil
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &allErrorPath.Summary
    return children
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetBundleName() string { return "cisco_ios_xr" }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetYangName() string { return "all-error-path" }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) SetParent(parent types.Entity) { allErrorPath.parent = parent }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetParent() types.Entity { return allErrorPath.parent }

func (allErrorPath *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath) GetParentYangName() string { return "all-instances" }

// AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary
// Summary of all instances errors
type AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // legacy client. The type is bool.
    LegacyClient interface{}

    // cih client. The type is bool.
    CihClient interface{}

    // sum data. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData.
    SumData []AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetGoName(yname string) string {
    if yname == "legacy-client" { return "LegacyClient" }
    if yname == "cih-client" { return "CihClient" }
    if yname == "sum-data" { return "SumData" }
    return ""
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sum-data" {
        for _, c := range summary.SumData {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData{}
        summary.SumData = append(summary.SumData, child)
        return &summary.SumData[len(summary.SumData)-1]
    }
    return nil
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.SumData {
        children[summary.SumData[i].GetSegmentPath()] = &summary.SumData[i]
    }
    return children
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["legacy-client"] = summary.LegacyClient
    leafs["cih-client"] = summary.CihClient
    return leafs
}

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetYangName() string { return "summary" }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetParent() types.Entity { return summary.parent }

func (summary *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary) GetParentYangName() string { return "all-error-path" }

// AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData
// sum data
type AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // num nodes. The type is interface{} with range: 0..4294967295.
    NumNodes interface{}

    // crc err count. The type is interface{} with range: 0..4294967295.
    CrcErrCount interface{}

    // sbe err count. The type is interface{} with range: 0..4294967295.
    SbeErrCount interface{}

    // mbe err count. The type is interface{} with range: 0..4294967295.
    MbeErrCount interface{}

    // par err count. The type is interface{} with range: 0..4294967295.
    ParErrCount interface{}

    // gen err count. The type is interface{} with range: 0..4294967295.
    GenErrCount interface{}

    // reset err count. The type is interface{} with range: 0..4294967295.
    ResetErrCount interface{}

    // err count. The type is slice of interface{} with range: 0..4294967295.
    ErrCount []interface{}

    // pcie err count. The type is slice of interface{} with range: 0..4294967295.
    PcieErrCount []interface{}

    // node key. The type is slice of interface{} with range: 0..4294967295.
    NodeKey []interface{}
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetFilter() yfilter.YFilter { return sumData.YFilter }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) SetFilter(yf yfilter.YFilter) { sumData.YFilter = yf }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetGoName(yname string) string {
    if yname == "num-nodes" { return "NumNodes" }
    if yname == "crc-err-count" { return "CrcErrCount" }
    if yname == "sbe-err-count" { return "SbeErrCount" }
    if yname == "mbe-err-count" { return "MbeErrCount" }
    if yname == "par-err-count" { return "ParErrCount" }
    if yname == "gen-err-count" { return "GenErrCount" }
    if yname == "reset-err-count" { return "ResetErrCount" }
    if yname == "err-count" { return "ErrCount" }
    if yname == "pcie-err-count" { return "PcieErrCount" }
    if yname == "node-key" { return "NodeKey" }
    return ""
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetSegmentPath() string {
    return "sum-data"
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-nodes"] = sumData.NumNodes
    leafs["crc-err-count"] = sumData.CrcErrCount
    leafs["sbe-err-count"] = sumData.SbeErrCount
    leafs["mbe-err-count"] = sumData.MbeErrCount
    leafs["par-err-count"] = sumData.ParErrCount
    leafs["gen-err-count"] = sumData.GenErrCount
    leafs["reset-err-count"] = sumData.ResetErrCount
    leafs["err-count"] = sumData.ErrCount
    leafs["pcie-err-count"] = sumData.PcieErrCount
    leafs["node-key"] = sumData.NodeKey
    return leafs
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetBundleName() string { return "cisco_ios_xr" }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetYangName() string { return "sum-data" }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) SetParent(parent types.Entity) { sumData.parent = parent }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetParent() types.Entity { return sumData.parent }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_AllInstances_AllErrorPath_Summary_SumData) GetParentYangName() string { return "summary" }

// AsicErrors_Nodes_Node_AsicInformation_Instances
// All asic errors  on the node
type AsicErrors_Nodes_Node_AsicInformation_Instances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Particular asic instance on the node. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance.
    Instance []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetFilter() yfilter.YFilter { return instances.YFilter }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) SetFilter(yf yfilter.YFilter) { instances.YFilter = yf }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetSegmentPath() string {
    return "instances"
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        for _, c := range instances.Instance {
            if instances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance{}
        instances.Instance = append(instances.Instance, child)
        return &instances.Instance[len(instances.Instance)-1]
    }
    return nil
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instances.Instance {
        children[instances.Instance[i].GetSegmentPath()] = &instances.Instance[i]
    }
    return children
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetBundleName() string { return "cisco_ios_xr" }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetYangName() string { return "instances" }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) SetParent(parent types.Entity) { instances.parent = parent }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetParent() types.Entity { return instances.parent }

func (instances *AsicErrors_Nodes_Node_AsicInformation_Instances) GetParentYangName() string { return "asic-information" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance
// Particular asic instance on the node
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. asic instance. The type is interface{} with range:
    // -2147483648..2147483647.
    AsicInstance interface{}

    // Error path of the instances.
    ErrorPath AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetGoName(yname string) string {
    if yname == "asic-instance" { return "AsicInstance" }
    if yname == "error-path" { return "ErrorPath" }
    return ""
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetSegmentPath() string {
    return "instance" + "[asic-instance='" + fmt.Sprintf("%v", instance.AsicInstance) + "']"
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error-path" {
        return &instance.ErrorPath
    }
    return nil
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["error-path"] = &instance.ErrorPath
    return children
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asic-instance"] = instance.AsicInstance
    return leafs
}

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetYangName() string { return "instance" }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetParent() types.Entity { return instance.parent }

func (instance *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance) GetParentYangName() string { return "instances" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath
// Error path of the instances
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Multiple bit soft error information.
    MultipleBitSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors

    // Indirect hard error information.
    AsicErrorGenericSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft

    // CRC hard error information.
    CrcHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors

    // Indirect hard error information.
    AsicErrorSbeSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft

    // Hardware soft error information.
    HardwareSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors

    // Indirect hard error information.
    AsicErrorCrcSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft

    // Indirect hard error information.
    AsicErrorParitySoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft

    // IO soft error information.
    IoSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors

    // Reset soft error information.
    ResetSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors

    // Barrier hard error information.
    BarrierHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors

    // Ucode soft error information.
    UcodeSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors

    // Indirect hard error information.
    AsicErrorResetHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard

    // Single bit hard error information.
    SingleBitHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors

    // Indirect hard error information.
    IndirectHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors

    // OOR thresh information.
    OutofResourceSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft

    // CRC soft error information.
    CrcSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors

    // Time out hard error information.
    TimeOutHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors

    // Barrier soft error information.
    BarrierSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors

    // Indirect hard error information.
    AsicErrorMbeSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft

    // BP hard error information.
    BackPressureHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors

    // Single bit soft error information.
    SingleBitSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors

    // Indirect soft error information.
    IndirectSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors

    // Generic hard error information.
    GenericHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors

    // Link hard error information.
    LinkHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors

    // Configuration hard error information.
    ConfigurationHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors

    // Summary for a specific instance.
    InstanceSummary AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary

    // Unexpected hard error information.
    UnexpectedHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors

    // Time out soft error information.
    TimeOutSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors

    // Indirect hard error information.
    AsicErrorGenericHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard

    // Parity hard error information.
    ParityHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors

    // Descriptor hard error information.
    DescriptorHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors

    // Interface hard error information.
    InterfaceHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors

    // Indirect hard error information.
    AsicErrorSbeHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard

    // Indirect hard error information.
    AsicErrorCrcHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard

    // Indirect hard error information.
    AsicErrorParityHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard

    // Indirect hard error information.
    AsicErrorResetSoft AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft

    // BP soft error information.
    BackPressureSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors

    // Generic soft error information.
    GenericSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors

    // Link soft error information.
    LinkSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors

    // Configuration soft error information.
    ConfigurationSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors

    // Multiple bit hard error information.
    MultipleBitHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors

    // Unexpected soft error information.
    UnexpectedSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors

    // OOR thresh information.
    OutofResourceHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard

    // Hardware hard error information.
    HardwareHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors

    // Parity soft error information.
    ParitySoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors

    // Descriptor soft error information.
    DescriptorSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors

    // Interface soft error information.
    InterfaceSoftErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors

    // IO hard error information.
    IoHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors

    // Reset hard error information.
    ResetHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors

    // UCode hard error information.
    UcodeHardErrors AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors

    // Indirect hard error information.
    AsicErrorMbeHard AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetFilter() yfilter.YFilter { return errorPath.YFilter }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) SetFilter(yf yfilter.YFilter) { errorPath.YFilter = yf }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetGoName(yname string) string {
    if yname == "multiple-bit-soft-errors" { return "MultipleBitSoftErrors" }
    if yname == "asic-error-generic-soft" { return "AsicErrorGenericSoft" }
    if yname == "crc-hard-errors" { return "CrcHardErrors" }
    if yname == "asic-error-sbe-soft" { return "AsicErrorSbeSoft" }
    if yname == "hardware-soft-errors" { return "HardwareSoftErrors" }
    if yname == "asic-error-crc-soft" { return "AsicErrorCrcSoft" }
    if yname == "asic-error-parity-soft" { return "AsicErrorParitySoft" }
    if yname == "io-soft-errors" { return "IoSoftErrors" }
    if yname == "reset-soft-errors" { return "ResetSoftErrors" }
    if yname == "barrier-hard-errors" { return "BarrierHardErrors" }
    if yname == "ucode-soft-errors" { return "UcodeSoftErrors" }
    if yname == "asic-error-reset-hard" { return "AsicErrorResetHard" }
    if yname == "single-bit-hard-errors" { return "SingleBitHardErrors" }
    if yname == "indirect-hard-errors" { return "IndirectHardErrors" }
    if yname == "outof-resource-soft" { return "OutofResourceSoft" }
    if yname == "crc-soft-errors" { return "CrcSoftErrors" }
    if yname == "time-out-hard-errors" { return "TimeOutHardErrors" }
    if yname == "barrier-soft-errors" { return "BarrierSoftErrors" }
    if yname == "asic-error-mbe-soft" { return "AsicErrorMbeSoft" }
    if yname == "back-pressure-hard-errors" { return "BackPressureHardErrors" }
    if yname == "single-bit-soft-errors" { return "SingleBitSoftErrors" }
    if yname == "indirect-soft-errors" { return "IndirectSoftErrors" }
    if yname == "generic-hard-errors" { return "GenericHardErrors" }
    if yname == "link-hard-errors" { return "LinkHardErrors" }
    if yname == "configuration-hard-errors" { return "ConfigurationHardErrors" }
    if yname == "instance-summary" { return "InstanceSummary" }
    if yname == "unexpected-hard-errors" { return "UnexpectedHardErrors" }
    if yname == "time-out-soft-errors" { return "TimeOutSoftErrors" }
    if yname == "asic-error-generic-hard" { return "AsicErrorGenericHard" }
    if yname == "parity-hard-errors" { return "ParityHardErrors" }
    if yname == "descriptor-hard-errors" { return "DescriptorHardErrors" }
    if yname == "interface-hard-errors" { return "InterfaceHardErrors" }
    if yname == "asic-error-sbe-hard" { return "AsicErrorSbeHard" }
    if yname == "asic-error-crc-hard" { return "AsicErrorCrcHard" }
    if yname == "asic-error-parity-hard" { return "AsicErrorParityHard" }
    if yname == "asic-error-reset-soft" { return "AsicErrorResetSoft" }
    if yname == "back-pressure-soft-errors" { return "BackPressureSoftErrors" }
    if yname == "generic-soft-errors" { return "GenericSoftErrors" }
    if yname == "link-soft-errors" { return "LinkSoftErrors" }
    if yname == "configuration-soft-errors" { return "ConfigurationSoftErrors" }
    if yname == "multiple-bit-hard-errors" { return "MultipleBitHardErrors" }
    if yname == "unexpected-soft-errors" { return "UnexpectedSoftErrors" }
    if yname == "outof-resource-hard" { return "OutofResourceHard" }
    if yname == "hardware-hard-errors" { return "HardwareHardErrors" }
    if yname == "parity-soft-errors" { return "ParitySoftErrors" }
    if yname == "descriptor-soft-errors" { return "DescriptorSoftErrors" }
    if yname == "interface-soft-errors" { return "InterfaceSoftErrors" }
    if yname == "io-hard-errors" { return "IoHardErrors" }
    if yname == "reset-hard-errors" { return "ResetHardErrors" }
    if yname == "ucode-hard-errors" { return "UcodeHardErrors" }
    if yname == "asic-error-mbe-hard" { return "AsicErrorMbeHard" }
    return ""
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetSegmentPath() string {
    return "error-path"
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multiple-bit-soft-errors" {
        return &errorPath.MultipleBitSoftErrors
    }
    if childYangName == "asic-error-generic-soft" {
        return &errorPath.AsicErrorGenericSoft
    }
    if childYangName == "crc-hard-errors" {
        return &errorPath.CrcHardErrors
    }
    if childYangName == "asic-error-sbe-soft" {
        return &errorPath.AsicErrorSbeSoft
    }
    if childYangName == "hardware-soft-errors" {
        return &errorPath.HardwareSoftErrors
    }
    if childYangName == "asic-error-crc-soft" {
        return &errorPath.AsicErrorCrcSoft
    }
    if childYangName == "asic-error-parity-soft" {
        return &errorPath.AsicErrorParitySoft
    }
    if childYangName == "io-soft-errors" {
        return &errorPath.IoSoftErrors
    }
    if childYangName == "reset-soft-errors" {
        return &errorPath.ResetSoftErrors
    }
    if childYangName == "barrier-hard-errors" {
        return &errorPath.BarrierHardErrors
    }
    if childYangName == "ucode-soft-errors" {
        return &errorPath.UcodeSoftErrors
    }
    if childYangName == "asic-error-reset-hard" {
        return &errorPath.AsicErrorResetHard
    }
    if childYangName == "single-bit-hard-errors" {
        return &errorPath.SingleBitHardErrors
    }
    if childYangName == "indirect-hard-errors" {
        return &errorPath.IndirectHardErrors
    }
    if childYangName == "outof-resource-soft" {
        return &errorPath.OutofResourceSoft
    }
    if childYangName == "crc-soft-errors" {
        return &errorPath.CrcSoftErrors
    }
    if childYangName == "time-out-hard-errors" {
        return &errorPath.TimeOutHardErrors
    }
    if childYangName == "barrier-soft-errors" {
        return &errorPath.BarrierSoftErrors
    }
    if childYangName == "asic-error-mbe-soft" {
        return &errorPath.AsicErrorMbeSoft
    }
    if childYangName == "back-pressure-hard-errors" {
        return &errorPath.BackPressureHardErrors
    }
    if childYangName == "single-bit-soft-errors" {
        return &errorPath.SingleBitSoftErrors
    }
    if childYangName == "indirect-soft-errors" {
        return &errorPath.IndirectSoftErrors
    }
    if childYangName == "generic-hard-errors" {
        return &errorPath.GenericHardErrors
    }
    if childYangName == "link-hard-errors" {
        return &errorPath.LinkHardErrors
    }
    if childYangName == "configuration-hard-errors" {
        return &errorPath.ConfigurationHardErrors
    }
    if childYangName == "instance-summary" {
        return &errorPath.InstanceSummary
    }
    if childYangName == "unexpected-hard-errors" {
        return &errorPath.UnexpectedHardErrors
    }
    if childYangName == "time-out-soft-errors" {
        return &errorPath.TimeOutSoftErrors
    }
    if childYangName == "asic-error-generic-hard" {
        return &errorPath.AsicErrorGenericHard
    }
    if childYangName == "parity-hard-errors" {
        return &errorPath.ParityHardErrors
    }
    if childYangName == "descriptor-hard-errors" {
        return &errorPath.DescriptorHardErrors
    }
    if childYangName == "interface-hard-errors" {
        return &errorPath.InterfaceHardErrors
    }
    if childYangName == "asic-error-sbe-hard" {
        return &errorPath.AsicErrorSbeHard
    }
    if childYangName == "asic-error-crc-hard" {
        return &errorPath.AsicErrorCrcHard
    }
    if childYangName == "asic-error-parity-hard" {
        return &errorPath.AsicErrorParityHard
    }
    if childYangName == "asic-error-reset-soft" {
        return &errorPath.AsicErrorResetSoft
    }
    if childYangName == "back-pressure-soft-errors" {
        return &errorPath.BackPressureSoftErrors
    }
    if childYangName == "generic-soft-errors" {
        return &errorPath.GenericSoftErrors
    }
    if childYangName == "link-soft-errors" {
        return &errorPath.LinkSoftErrors
    }
    if childYangName == "configuration-soft-errors" {
        return &errorPath.ConfigurationSoftErrors
    }
    if childYangName == "multiple-bit-hard-errors" {
        return &errorPath.MultipleBitHardErrors
    }
    if childYangName == "unexpected-soft-errors" {
        return &errorPath.UnexpectedSoftErrors
    }
    if childYangName == "outof-resource-hard" {
        return &errorPath.OutofResourceHard
    }
    if childYangName == "hardware-hard-errors" {
        return &errorPath.HardwareHardErrors
    }
    if childYangName == "parity-soft-errors" {
        return &errorPath.ParitySoftErrors
    }
    if childYangName == "descriptor-soft-errors" {
        return &errorPath.DescriptorSoftErrors
    }
    if childYangName == "interface-soft-errors" {
        return &errorPath.InterfaceSoftErrors
    }
    if childYangName == "io-hard-errors" {
        return &errorPath.IoHardErrors
    }
    if childYangName == "reset-hard-errors" {
        return &errorPath.ResetHardErrors
    }
    if childYangName == "ucode-hard-errors" {
        return &errorPath.UcodeHardErrors
    }
    if childYangName == "asic-error-mbe-hard" {
        return &errorPath.AsicErrorMbeHard
    }
    return nil
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["multiple-bit-soft-errors"] = &errorPath.MultipleBitSoftErrors
    children["asic-error-generic-soft"] = &errorPath.AsicErrorGenericSoft
    children["crc-hard-errors"] = &errorPath.CrcHardErrors
    children["asic-error-sbe-soft"] = &errorPath.AsicErrorSbeSoft
    children["hardware-soft-errors"] = &errorPath.HardwareSoftErrors
    children["asic-error-crc-soft"] = &errorPath.AsicErrorCrcSoft
    children["asic-error-parity-soft"] = &errorPath.AsicErrorParitySoft
    children["io-soft-errors"] = &errorPath.IoSoftErrors
    children["reset-soft-errors"] = &errorPath.ResetSoftErrors
    children["barrier-hard-errors"] = &errorPath.BarrierHardErrors
    children["ucode-soft-errors"] = &errorPath.UcodeSoftErrors
    children["asic-error-reset-hard"] = &errorPath.AsicErrorResetHard
    children["single-bit-hard-errors"] = &errorPath.SingleBitHardErrors
    children["indirect-hard-errors"] = &errorPath.IndirectHardErrors
    children["outof-resource-soft"] = &errorPath.OutofResourceSoft
    children["crc-soft-errors"] = &errorPath.CrcSoftErrors
    children["time-out-hard-errors"] = &errorPath.TimeOutHardErrors
    children["barrier-soft-errors"] = &errorPath.BarrierSoftErrors
    children["asic-error-mbe-soft"] = &errorPath.AsicErrorMbeSoft
    children["back-pressure-hard-errors"] = &errorPath.BackPressureHardErrors
    children["single-bit-soft-errors"] = &errorPath.SingleBitSoftErrors
    children["indirect-soft-errors"] = &errorPath.IndirectSoftErrors
    children["generic-hard-errors"] = &errorPath.GenericHardErrors
    children["link-hard-errors"] = &errorPath.LinkHardErrors
    children["configuration-hard-errors"] = &errorPath.ConfigurationHardErrors
    children["instance-summary"] = &errorPath.InstanceSummary
    children["unexpected-hard-errors"] = &errorPath.UnexpectedHardErrors
    children["time-out-soft-errors"] = &errorPath.TimeOutSoftErrors
    children["asic-error-generic-hard"] = &errorPath.AsicErrorGenericHard
    children["parity-hard-errors"] = &errorPath.ParityHardErrors
    children["descriptor-hard-errors"] = &errorPath.DescriptorHardErrors
    children["interface-hard-errors"] = &errorPath.InterfaceHardErrors
    children["asic-error-sbe-hard"] = &errorPath.AsicErrorSbeHard
    children["asic-error-crc-hard"] = &errorPath.AsicErrorCrcHard
    children["asic-error-parity-hard"] = &errorPath.AsicErrorParityHard
    children["asic-error-reset-soft"] = &errorPath.AsicErrorResetSoft
    children["back-pressure-soft-errors"] = &errorPath.BackPressureSoftErrors
    children["generic-soft-errors"] = &errorPath.GenericSoftErrors
    children["link-soft-errors"] = &errorPath.LinkSoftErrors
    children["configuration-soft-errors"] = &errorPath.ConfigurationSoftErrors
    children["multiple-bit-hard-errors"] = &errorPath.MultipleBitHardErrors
    children["unexpected-soft-errors"] = &errorPath.UnexpectedSoftErrors
    children["outof-resource-hard"] = &errorPath.OutofResourceHard
    children["hardware-hard-errors"] = &errorPath.HardwareHardErrors
    children["parity-soft-errors"] = &errorPath.ParitySoftErrors
    children["descriptor-soft-errors"] = &errorPath.DescriptorSoftErrors
    children["interface-soft-errors"] = &errorPath.InterfaceSoftErrors
    children["io-hard-errors"] = &errorPath.IoHardErrors
    children["reset-hard-errors"] = &errorPath.ResetHardErrors
    children["ucode-hard-errors"] = &errorPath.UcodeHardErrors
    children["asic-error-mbe-hard"] = &errorPath.AsicErrorMbeHard
    return children
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetBundleName() string { return "cisco_ios_xr" }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetYangName() string { return "error-path" }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) SetParent(parent types.Entity) { errorPath.parent = parent }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetParent() types.Entity { return errorPath.parent }

func (errorPath *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath) GetParentYangName() string { return "instance" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors
// Multiple bit soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetFilter() yfilter.YFilter { return multipleBitSoftErrors.YFilter }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) SetFilter(yf yfilter.YFilter) { multipleBitSoftErrors.YFilter = yf }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetSegmentPath() string {
    return "multiple-bit-soft-errors"
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range multipleBitSoftErrors.Error {
            if multipleBitSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error{}
        multipleBitSoftErrors.Error = append(multipleBitSoftErrors.Error, child)
        return &multipleBitSoftErrors.Error[len(multipleBitSoftErrors.Error)-1]
    }
    return nil
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multipleBitSoftErrors.Error {
        children[multipleBitSoftErrors.Error[i].GetSegmentPath()] = &multipleBitSoftErrors.Error[i]
    }
    return children
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetYangName() string { return "multiple-bit-soft-errors" }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) SetParent(parent types.Entity) { multipleBitSoftErrors.parent = parent }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetParent() types.Entity { return multipleBitSoftErrors.parent }

func (multipleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error) GetParentYangName() string { return "multiple-bit-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetFilter() yfilter.YFilter { return asicErrorGenericSoft.YFilter }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) SetFilter(yf yfilter.YFilter) { asicErrorGenericSoft.YFilter = yf }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetSegmentPath() string {
    return "asic-error-generic-soft"
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorGenericSoft.Error {
            if asicErrorGenericSoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error{}
        asicErrorGenericSoft.Error = append(asicErrorGenericSoft.Error, child)
        return &asicErrorGenericSoft.Error[len(asicErrorGenericSoft.Error)-1]
    }
    return nil
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorGenericSoft.Error {
        children[asicErrorGenericSoft.Error[i].GetSegmentPath()] = &asicErrorGenericSoft.Error[i]
    }
    return children
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetYangName() string { return "asic-error-generic-soft" }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) SetParent(parent types.Entity) { asicErrorGenericSoft.parent = parent }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetParent() types.Entity { return asicErrorGenericSoft.parent }

func (asicErrorGenericSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error) GetParentYangName() string { return "asic-error-generic-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericSoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors
// CRC hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetFilter() yfilter.YFilter { return crcHardErrors.YFilter }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) SetFilter(yf yfilter.YFilter) { crcHardErrors.YFilter = yf }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetSegmentPath() string {
    return "crc-hard-errors"
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range crcHardErrors.Error {
            if crcHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error{}
        crcHardErrors.Error = append(crcHardErrors.Error, child)
        return &crcHardErrors.Error[len(crcHardErrors.Error)-1]
    }
    return nil
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crcHardErrors.Error {
        children[crcHardErrors.Error[i].GetSegmentPath()] = &crcHardErrors.Error[i]
    }
    return children
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetYangName() string { return "crc-hard-errors" }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) SetParent(parent types.Entity) { crcHardErrors.parent = parent }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetParent() types.Entity { return crcHardErrors.parent }

func (crcHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error) GetParentYangName() string { return "crc-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetFilter() yfilter.YFilter { return asicErrorSbeSoft.YFilter }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) SetFilter(yf yfilter.YFilter) { asicErrorSbeSoft.YFilter = yf }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetSegmentPath() string {
    return "asic-error-sbe-soft"
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorSbeSoft.Error {
            if asicErrorSbeSoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error{}
        asicErrorSbeSoft.Error = append(asicErrorSbeSoft.Error, child)
        return &asicErrorSbeSoft.Error[len(asicErrorSbeSoft.Error)-1]
    }
    return nil
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorSbeSoft.Error {
        children[asicErrorSbeSoft.Error[i].GetSegmentPath()] = &asicErrorSbeSoft.Error[i]
    }
    return children
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetYangName() string { return "asic-error-sbe-soft" }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) SetParent(parent types.Entity) { asicErrorSbeSoft.parent = parent }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetParent() types.Entity { return asicErrorSbeSoft.parent }

func (asicErrorSbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error) GetParentYangName() string { return "asic-error-sbe-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeSoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors
// Hardware soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetFilter() yfilter.YFilter { return hardwareSoftErrors.YFilter }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) SetFilter(yf yfilter.YFilter) { hardwareSoftErrors.YFilter = yf }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetSegmentPath() string {
    return "hardware-soft-errors"
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range hardwareSoftErrors.Error {
            if hardwareSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error{}
        hardwareSoftErrors.Error = append(hardwareSoftErrors.Error, child)
        return &hardwareSoftErrors.Error[len(hardwareSoftErrors.Error)-1]
    }
    return nil
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareSoftErrors.Error {
        children[hardwareSoftErrors.Error[i].GetSegmentPath()] = &hardwareSoftErrors.Error[i]
    }
    return children
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetYangName() string { return "hardware-soft-errors" }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) SetParent(parent types.Entity) { hardwareSoftErrors.parent = parent }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetParent() types.Entity { return hardwareSoftErrors.parent }

func (hardwareSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error) GetParentYangName() string { return "hardware-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetFilter() yfilter.YFilter { return asicErrorCrcSoft.YFilter }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) SetFilter(yf yfilter.YFilter) { asicErrorCrcSoft.YFilter = yf }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetSegmentPath() string {
    return "asic-error-crc-soft"
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorCrcSoft.Error {
            if asicErrorCrcSoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error{}
        asicErrorCrcSoft.Error = append(asicErrorCrcSoft.Error, child)
        return &asicErrorCrcSoft.Error[len(asicErrorCrcSoft.Error)-1]
    }
    return nil
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorCrcSoft.Error {
        children[asicErrorCrcSoft.Error[i].GetSegmentPath()] = &asicErrorCrcSoft.Error[i]
    }
    return children
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetYangName() string { return "asic-error-crc-soft" }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) SetParent(parent types.Entity) { asicErrorCrcSoft.parent = parent }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetParent() types.Entity { return asicErrorCrcSoft.parent }

func (asicErrorCrcSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error) GetParentYangName() string { return "asic-error-crc-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcSoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetFilter() yfilter.YFilter { return asicErrorParitySoft.YFilter }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) SetFilter(yf yfilter.YFilter) { asicErrorParitySoft.YFilter = yf }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetSegmentPath() string {
    return "asic-error-parity-soft"
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorParitySoft.Error {
            if asicErrorParitySoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error{}
        asicErrorParitySoft.Error = append(asicErrorParitySoft.Error, child)
        return &asicErrorParitySoft.Error[len(asicErrorParitySoft.Error)-1]
    }
    return nil
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorParitySoft.Error {
        children[asicErrorParitySoft.Error[i].GetSegmentPath()] = &asicErrorParitySoft.Error[i]
    }
    return children
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetYangName() string { return "asic-error-parity-soft" }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) SetParent(parent types.Entity) { asicErrorParitySoft.parent = parent }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetParent() types.Entity { return asicErrorParitySoft.parent }

func (asicErrorParitySoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error) GetParentYangName() string { return "asic-error-parity-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParitySoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors
// IO soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetFilter() yfilter.YFilter { return ioSoftErrors.YFilter }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) SetFilter(yf yfilter.YFilter) { ioSoftErrors.YFilter = yf }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetSegmentPath() string {
    return "io-soft-errors"
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range ioSoftErrors.Error {
            if ioSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error{}
        ioSoftErrors.Error = append(ioSoftErrors.Error, child)
        return &ioSoftErrors.Error[len(ioSoftErrors.Error)-1]
    }
    return nil
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ioSoftErrors.Error {
        children[ioSoftErrors.Error[i].GetSegmentPath()] = &ioSoftErrors.Error[i]
    }
    return children
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetYangName() string { return "io-soft-errors" }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) SetParent(parent types.Entity) { ioSoftErrors.parent = parent }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetParent() types.Entity { return ioSoftErrors.parent }

func (ioSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error) GetParentYangName() string { return "io-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors
// Reset soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetFilter() yfilter.YFilter { return resetSoftErrors.YFilter }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) SetFilter(yf yfilter.YFilter) { resetSoftErrors.YFilter = yf }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetSegmentPath() string {
    return "reset-soft-errors"
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range resetSoftErrors.Error {
            if resetSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error{}
        resetSoftErrors.Error = append(resetSoftErrors.Error, child)
        return &resetSoftErrors.Error[len(resetSoftErrors.Error)-1]
    }
    return nil
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range resetSoftErrors.Error {
        children[resetSoftErrors.Error[i].GetSegmentPath()] = &resetSoftErrors.Error[i]
    }
    return children
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetYangName() string { return "reset-soft-errors" }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) SetParent(parent types.Entity) { resetSoftErrors.parent = parent }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetParent() types.Entity { return resetSoftErrors.parent }

func (resetSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error) GetParentYangName() string { return "reset-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors
// Barrier hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetFilter() yfilter.YFilter { return barrierHardErrors.YFilter }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) SetFilter(yf yfilter.YFilter) { barrierHardErrors.YFilter = yf }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetSegmentPath() string {
    return "barrier-hard-errors"
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range barrierHardErrors.Error {
            if barrierHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error{}
        barrierHardErrors.Error = append(barrierHardErrors.Error, child)
        return &barrierHardErrors.Error[len(barrierHardErrors.Error)-1]
    }
    return nil
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range barrierHardErrors.Error {
        children[barrierHardErrors.Error[i].GetSegmentPath()] = &barrierHardErrors.Error[i]
    }
    return children
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetYangName() string { return "barrier-hard-errors" }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) SetParent(parent types.Entity) { barrierHardErrors.parent = parent }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetParent() types.Entity { return barrierHardErrors.parent }

func (barrierHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error) GetParentYangName() string { return "barrier-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors
// Ucode soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetFilter() yfilter.YFilter { return ucodeSoftErrors.YFilter }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) SetFilter(yf yfilter.YFilter) { ucodeSoftErrors.YFilter = yf }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetSegmentPath() string {
    return "ucode-soft-errors"
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range ucodeSoftErrors.Error {
            if ucodeSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error{}
        ucodeSoftErrors.Error = append(ucodeSoftErrors.Error, child)
        return &ucodeSoftErrors.Error[len(ucodeSoftErrors.Error)-1]
    }
    return nil
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ucodeSoftErrors.Error {
        children[ucodeSoftErrors.Error[i].GetSegmentPath()] = &ucodeSoftErrors.Error[i]
    }
    return children
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetYangName() string { return "ucode-soft-errors" }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) SetParent(parent types.Entity) { ucodeSoftErrors.parent = parent }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetParent() types.Entity { return ucodeSoftErrors.parent }

func (ucodeSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error) GetParentYangName() string { return "ucode-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetFilter() yfilter.YFilter { return asicErrorResetHard.YFilter }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) SetFilter(yf yfilter.YFilter) { asicErrorResetHard.YFilter = yf }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetSegmentPath() string {
    return "asic-error-reset-hard"
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorResetHard.Error {
            if asicErrorResetHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error{}
        asicErrorResetHard.Error = append(asicErrorResetHard.Error, child)
        return &asicErrorResetHard.Error[len(asicErrorResetHard.Error)-1]
    }
    return nil
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorResetHard.Error {
        children[asicErrorResetHard.Error[i].GetSegmentPath()] = &asicErrorResetHard.Error[i]
    }
    return children
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetYangName() string { return "asic-error-reset-hard" }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) SetParent(parent types.Entity) { asicErrorResetHard.parent = parent }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetParent() types.Entity { return asicErrorResetHard.parent }

func (asicErrorResetHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error) GetParentYangName() string { return "asic-error-reset-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetHard_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors
// Single bit hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetFilter() yfilter.YFilter { return singleBitHardErrors.YFilter }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) SetFilter(yf yfilter.YFilter) { singleBitHardErrors.YFilter = yf }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetSegmentPath() string {
    return "single-bit-hard-errors"
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range singleBitHardErrors.Error {
            if singleBitHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error{}
        singleBitHardErrors.Error = append(singleBitHardErrors.Error, child)
        return &singleBitHardErrors.Error[len(singleBitHardErrors.Error)-1]
    }
    return nil
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range singleBitHardErrors.Error {
        children[singleBitHardErrors.Error[i].GetSegmentPath()] = &singleBitHardErrors.Error[i]
    }
    return children
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetYangName() string { return "single-bit-hard-errors" }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) SetParent(parent types.Entity) { singleBitHardErrors.parent = parent }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetParent() types.Entity { return singleBitHardErrors.parent }

func (singleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error) GetParentYangName() string { return "single-bit-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetFilter() yfilter.YFilter { return indirectHardErrors.YFilter }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) SetFilter(yf yfilter.YFilter) { indirectHardErrors.YFilter = yf }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetSegmentPath() string {
    return "indirect-hard-errors"
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range indirectHardErrors.Error {
            if indirectHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error{}
        indirectHardErrors.Error = append(indirectHardErrors.Error, child)
        return &indirectHardErrors.Error[len(indirectHardErrors.Error)-1]
    }
    return nil
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indirectHardErrors.Error {
        children[indirectHardErrors.Error[i].GetSegmentPath()] = &indirectHardErrors.Error[i]
    }
    return children
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetYangName() string { return "indirect-hard-errors" }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) SetParent(parent types.Entity) { indirectHardErrors.parent = parent }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetParent() types.Entity { return indirectHardErrors.parent }

func (indirectHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error) GetParentYangName() string { return "indirect-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft
// OOR thresh information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetFilter() yfilter.YFilter { return outofResourceSoft.YFilter }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) SetFilter(yf yfilter.YFilter) { outofResourceSoft.YFilter = yf }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetSegmentPath() string {
    return "outof-resource-soft"
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range outofResourceSoft.Error {
            if outofResourceSoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error{}
        outofResourceSoft.Error = append(outofResourceSoft.Error, child)
        return &outofResourceSoft.Error[len(outofResourceSoft.Error)-1]
    }
    return nil
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range outofResourceSoft.Error {
        children[outofResourceSoft.Error[i].GetSegmentPath()] = &outofResourceSoft.Error[i]
    }
    return children
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetBundleName() string { return "cisco_ios_xr" }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetYangName() string { return "outof-resource-soft" }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) SetParent(parent types.Entity) { outofResourceSoft.parent = parent }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetParent() types.Entity { return outofResourceSoft.parent }

func (outofResourceSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error) GetParentYangName() string { return "outof-resource-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceSoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors
// CRC soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetFilter() yfilter.YFilter { return crcSoftErrors.YFilter }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) SetFilter(yf yfilter.YFilter) { crcSoftErrors.YFilter = yf }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetSegmentPath() string {
    return "crc-soft-errors"
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range crcSoftErrors.Error {
            if crcSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error{}
        crcSoftErrors.Error = append(crcSoftErrors.Error, child)
        return &crcSoftErrors.Error[len(crcSoftErrors.Error)-1]
    }
    return nil
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crcSoftErrors.Error {
        children[crcSoftErrors.Error[i].GetSegmentPath()] = &crcSoftErrors.Error[i]
    }
    return children
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetYangName() string { return "crc-soft-errors" }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) SetParent(parent types.Entity) { crcSoftErrors.parent = parent }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetParent() types.Entity { return crcSoftErrors.parent }

func (crcSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error) GetParentYangName() string { return "crc-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_CrcSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors
// Time out hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetFilter() yfilter.YFilter { return timeOutHardErrors.YFilter }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) SetFilter(yf yfilter.YFilter) { timeOutHardErrors.YFilter = yf }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetSegmentPath() string {
    return "time-out-hard-errors"
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range timeOutHardErrors.Error {
            if timeOutHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error{}
        timeOutHardErrors.Error = append(timeOutHardErrors.Error, child)
        return &timeOutHardErrors.Error[len(timeOutHardErrors.Error)-1]
    }
    return nil
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range timeOutHardErrors.Error {
        children[timeOutHardErrors.Error[i].GetSegmentPath()] = &timeOutHardErrors.Error[i]
    }
    return children
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetYangName() string { return "time-out-hard-errors" }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) SetParent(parent types.Entity) { timeOutHardErrors.parent = parent }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetParent() types.Entity { return timeOutHardErrors.parent }

func (timeOutHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error) GetParentYangName() string { return "time-out-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors
// Barrier soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetFilter() yfilter.YFilter { return barrierSoftErrors.YFilter }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) SetFilter(yf yfilter.YFilter) { barrierSoftErrors.YFilter = yf }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetSegmentPath() string {
    return "barrier-soft-errors"
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range barrierSoftErrors.Error {
            if barrierSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error{}
        barrierSoftErrors.Error = append(barrierSoftErrors.Error, child)
        return &barrierSoftErrors.Error[len(barrierSoftErrors.Error)-1]
    }
    return nil
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range barrierSoftErrors.Error {
        children[barrierSoftErrors.Error[i].GetSegmentPath()] = &barrierSoftErrors.Error[i]
    }
    return children
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetYangName() string { return "barrier-soft-errors" }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) SetParent(parent types.Entity) { barrierSoftErrors.parent = parent }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetParent() types.Entity { return barrierSoftErrors.parent }

func (barrierSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error) GetParentYangName() string { return "barrier-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BarrierSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetFilter() yfilter.YFilter { return asicErrorMbeSoft.YFilter }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) SetFilter(yf yfilter.YFilter) { asicErrorMbeSoft.YFilter = yf }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetSegmentPath() string {
    return "asic-error-mbe-soft"
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorMbeSoft.Error {
            if asicErrorMbeSoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error{}
        asicErrorMbeSoft.Error = append(asicErrorMbeSoft.Error, child)
        return &asicErrorMbeSoft.Error[len(asicErrorMbeSoft.Error)-1]
    }
    return nil
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorMbeSoft.Error {
        children[asicErrorMbeSoft.Error[i].GetSegmentPath()] = &asicErrorMbeSoft.Error[i]
    }
    return children
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetYangName() string { return "asic-error-mbe-soft" }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) SetParent(parent types.Entity) { asicErrorMbeSoft.parent = parent }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetParent() types.Entity { return asicErrorMbeSoft.parent }

func (asicErrorMbeSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error) GetParentYangName() string { return "asic-error-mbe-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeSoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors
// BP hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetFilter() yfilter.YFilter { return backPressureHardErrors.YFilter }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) SetFilter(yf yfilter.YFilter) { backPressureHardErrors.YFilter = yf }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetSegmentPath() string {
    return "back-pressure-hard-errors"
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range backPressureHardErrors.Error {
            if backPressureHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error{}
        backPressureHardErrors.Error = append(backPressureHardErrors.Error, child)
        return &backPressureHardErrors.Error[len(backPressureHardErrors.Error)-1]
    }
    return nil
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range backPressureHardErrors.Error {
        children[backPressureHardErrors.Error[i].GetSegmentPath()] = &backPressureHardErrors.Error[i]
    }
    return children
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetYangName() string { return "back-pressure-hard-errors" }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) SetParent(parent types.Entity) { backPressureHardErrors.parent = parent }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetParent() types.Entity { return backPressureHardErrors.parent }

func (backPressureHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error) GetParentYangName() string { return "back-pressure-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors
// Single bit soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetFilter() yfilter.YFilter { return singleBitSoftErrors.YFilter }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) SetFilter(yf yfilter.YFilter) { singleBitSoftErrors.YFilter = yf }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetSegmentPath() string {
    return "single-bit-soft-errors"
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range singleBitSoftErrors.Error {
            if singleBitSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error{}
        singleBitSoftErrors.Error = append(singleBitSoftErrors.Error, child)
        return &singleBitSoftErrors.Error[len(singleBitSoftErrors.Error)-1]
    }
    return nil
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range singleBitSoftErrors.Error {
        children[singleBitSoftErrors.Error[i].GetSegmentPath()] = &singleBitSoftErrors.Error[i]
    }
    return children
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetYangName() string { return "single-bit-soft-errors" }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) SetParent(parent types.Entity) { singleBitSoftErrors.parent = parent }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetParent() types.Entity { return singleBitSoftErrors.parent }

func (singleBitSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error) GetParentYangName() string { return "single-bit-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_SingleBitSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors
// Indirect soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetFilter() yfilter.YFilter { return indirectSoftErrors.YFilter }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) SetFilter(yf yfilter.YFilter) { indirectSoftErrors.YFilter = yf }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetSegmentPath() string {
    return "indirect-soft-errors"
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range indirectSoftErrors.Error {
            if indirectSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error{}
        indirectSoftErrors.Error = append(indirectSoftErrors.Error, child)
        return &indirectSoftErrors.Error[len(indirectSoftErrors.Error)-1]
    }
    return nil
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indirectSoftErrors.Error {
        children[indirectSoftErrors.Error[i].GetSegmentPath()] = &indirectSoftErrors.Error[i]
    }
    return children
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetYangName() string { return "indirect-soft-errors" }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) SetParent(parent types.Entity) { indirectSoftErrors.parent = parent }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetParent() types.Entity { return indirectSoftErrors.parent }

func (indirectSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error) GetParentYangName() string { return "indirect-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IndirectSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors
// Generic hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetFilter() yfilter.YFilter { return genericHardErrors.YFilter }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) SetFilter(yf yfilter.YFilter) { genericHardErrors.YFilter = yf }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetSegmentPath() string {
    return "generic-hard-errors"
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range genericHardErrors.Error {
            if genericHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error{}
        genericHardErrors.Error = append(genericHardErrors.Error, child)
        return &genericHardErrors.Error[len(genericHardErrors.Error)-1]
    }
    return nil
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range genericHardErrors.Error {
        children[genericHardErrors.Error[i].GetSegmentPath()] = &genericHardErrors.Error[i]
    }
    return children
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetYangName() string { return "generic-hard-errors" }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) SetParent(parent types.Entity) { genericHardErrors.parent = parent }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetParent() types.Entity { return genericHardErrors.parent }

func (genericHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error) GetParentYangName() string { return "generic-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors
// Link hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetFilter() yfilter.YFilter { return linkHardErrors.YFilter }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) SetFilter(yf yfilter.YFilter) { linkHardErrors.YFilter = yf }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetSegmentPath() string {
    return "link-hard-errors"
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range linkHardErrors.Error {
            if linkHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error{}
        linkHardErrors.Error = append(linkHardErrors.Error, child)
        return &linkHardErrors.Error[len(linkHardErrors.Error)-1]
    }
    return nil
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkHardErrors.Error {
        children[linkHardErrors.Error[i].GetSegmentPath()] = &linkHardErrors.Error[i]
    }
    return children
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetYangName() string { return "link-hard-errors" }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) SetParent(parent types.Entity) { linkHardErrors.parent = parent }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetParent() types.Entity { return linkHardErrors.parent }

func (linkHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error) GetParentYangName() string { return "link-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors
// Configuration hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetFilter() yfilter.YFilter { return configurationHardErrors.YFilter }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) SetFilter(yf yfilter.YFilter) { configurationHardErrors.YFilter = yf }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetSegmentPath() string {
    return "configuration-hard-errors"
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range configurationHardErrors.Error {
            if configurationHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error{}
        configurationHardErrors.Error = append(configurationHardErrors.Error, child)
        return &configurationHardErrors.Error[len(configurationHardErrors.Error)-1]
    }
    return nil
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configurationHardErrors.Error {
        children[configurationHardErrors.Error[i].GetSegmentPath()] = &configurationHardErrors.Error[i]
    }
    return children
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetYangName() string { return "configuration-hard-errors" }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) SetParent(parent types.Entity) { configurationHardErrors.parent = parent }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetParent() types.Entity { return configurationHardErrors.parent }

func (configurationHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error) GetParentYangName() string { return "configuration-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary
// Summary for a specific instance
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // legacy client. The type is bool.
    LegacyClient interface{}

    // cih client. The type is bool.
    CihClient interface{}

    // sum data. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData.
    SumData []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetFilter() yfilter.YFilter { return instanceSummary.YFilter }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) SetFilter(yf yfilter.YFilter) { instanceSummary.YFilter = yf }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetGoName(yname string) string {
    if yname == "legacy-client" { return "LegacyClient" }
    if yname == "cih-client" { return "CihClient" }
    if yname == "sum-data" { return "SumData" }
    return ""
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetSegmentPath() string {
    return "instance-summary"
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sum-data" {
        for _, c := range instanceSummary.SumData {
            if instanceSummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData{}
        instanceSummary.SumData = append(instanceSummary.SumData, child)
        return &instanceSummary.SumData[len(instanceSummary.SumData)-1]
    }
    return nil
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instanceSummary.SumData {
        children[instanceSummary.SumData[i].GetSegmentPath()] = &instanceSummary.SumData[i]
    }
    return children
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["legacy-client"] = instanceSummary.LegacyClient
    leafs["cih-client"] = instanceSummary.CihClient
    return leafs
}

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetYangName() string { return "instance-summary" }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) SetParent(parent types.Entity) { instanceSummary.parent = parent }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetParent() types.Entity { return instanceSummary.parent }

func (instanceSummary *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData
// sum data
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // num nodes. The type is interface{} with range: 0..4294967295.
    NumNodes interface{}

    // crc err count. The type is interface{} with range: 0..4294967295.
    CrcErrCount interface{}

    // sbe err count. The type is interface{} with range: 0..4294967295.
    SbeErrCount interface{}

    // mbe err count. The type is interface{} with range: 0..4294967295.
    MbeErrCount interface{}

    // par err count. The type is interface{} with range: 0..4294967295.
    ParErrCount interface{}

    // gen err count. The type is interface{} with range: 0..4294967295.
    GenErrCount interface{}

    // reset err count. The type is interface{} with range: 0..4294967295.
    ResetErrCount interface{}

    // err count. The type is slice of interface{} with range: 0..4294967295.
    ErrCount []interface{}

    // pcie err count. The type is slice of interface{} with range: 0..4294967295.
    PcieErrCount []interface{}

    // node key. The type is slice of interface{} with range: 0..4294967295.
    NodeKey []interface{}
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetFilter() yfilter.YFilter { return sumData.YFilter }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) SetFilter(yf yfilter.YFilter) { sumData.YFilter = yf }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetGoName(yname string) string {
    if yname == "num-nodes" { return "NumNodes" }
    if yname == "crc-err-count" { return "CrcErrCount" }
    if yname == "sbe-err-count" { return "SbeErrCount" }
    if yname == "mbe-err-count" { return "MbeErrCount" }
    if yname == "par-err-count" { return "ParErrCount" }
    if yname == "gen-err-count" { return "GenErrCount" }
    if yname == "reset-err-count" { return "ResetErrCount" }
    if yname == "err-count" { return "ErrCount" }
    if yname == "pcie-err-count" { return "PcieErrCount" }
    if yname == "node-key" { return "NodeKey" }
    return ""
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetSegmentPath() string {
    return "sum-data"
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-nodes"] = sumData.NumNodes
    leafs["crc-err-count"] = sumData.CrcErrCount
    leafs["sbe-err-count"] = sumData.SbeErrCount
    leafs["mbe-err-count"] = sumData.MbeErrCount
    leafs["par-err-count"] = sumData.ParErrCount
    leafs["gen-err-count"] = sumData.GenErrCount
    leafs["reset-err-count"] = sumData.ResetErrCount
    leafs["err-count"] = sumData.ErrCount
    leafs["pcie-err-count"] = sumData.PcieErrCount
    leafs["node-key"] = sumData.NodeKey
    return leafs
}

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetBundleName() string { return "cisco_ios_xr" }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetYangName() string { return "sum-data" }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) SetParent(parent types.Entity) { sumData.parent = parent }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetParent() types.Entity { return sumData.parent }

func (sumData *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InstanceSummary_SumData) GetParentYangName() string { return "instance-summary" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors
// Unexpected hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetFilter() yfilter.YFilter { return unexpectedHardErrors.YFilter }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) SetFilter(yf yfilter.YFilter) { unexpectedHardErrors.YFilter = yf }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetSegmentPath() string {
    return "unexpected-hard-errors"
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range unexpectedHardErrors.Error {
            if unexpectedHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error{}
        unexpectedHardErrors.Error = append(unexpectedHardErrors.Error, child)
        return &unexpectedHardErrors.Error[len(unexpectedHardErrors.Error)-1]
    }
    return nil
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unexpectedHardErrors.Error {
        children[unexpectedHardErrors.Error[i].GetSegmentPath()] = &unexpectedHardErrors.Error[i]
    }
    return children
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetYangName() string { return "unexpected-hard-errors" }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) SetParent(parent types.Entity) { unexpectedHardErrors.parent = parent }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetParent() types.Entity { return unexpectedHardErrors.parent }

func (unexpectedHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error) GetParentYangName() string { return "unexpected-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors
// Time out soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetFilter() yfilter.YFilter { return timeOutSoftErrors.YFilter }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) SetFilter(yf yfilter.YFilter) { timeOutSoftErrors.YFilter = yf }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetSegmentPath() string {
    return "time-out-soft-errors"
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range timeOutSoftErrors.Error {
            if timeOutSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error{}
        timeOutSoftErrors.Error = append(timeOutSoftErrors.Error, child)
        return &timeOutSoftErrors.Error[len(timeOutSoftErrors.Error)-1]
    }
    return nil
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range timeOutSoftErrors.Error {
        children[timeOutSoftErrors.Error[i].GetSegmentPath()] = &timeOutSoftErrors.Error[i]
    }
    return children
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetYangName() string { return "time-out-soft-errors" }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) SetParent(parent types.Entity) { timeOutSoftErrors.parent = parent }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetParent() types.Entity { return timeOutSoftErrors.parent }

func (timeOutSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error) GetParentYangName() string { return "time-out-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_TimeOutSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetFilter() yfilter.YFilter { return asicErrorGenericHard.YFilter }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) SetFilter(yf yfilter.YFilter) { asicErrorGenericHard.YFilter = yf }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetSegmentPath() string {
    return "asic-error-generic-hard"
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorGenericHard.Error {
            if asicErrorGenericHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error{}
        asicErrorGenericHard.Error = append(asicErrorGenericHard.Error, child)
        return &asicErrorGenericHard.Error[len(asicErrorGenericHard.Error)-1]
    }
    return nil
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorGenericHard.Error {
        children[asicErrorGenericHard.Error[i].GetSegmentPath()] = &asicErrorGenericHard.Error[i]
    }
    return children
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetYangName() string { return "asic-error-generic-hard" }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) SetParent(parent types.Entity) { asicErrorGenericHard.parent = parent }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetParent() types.Entity { return asicErrorGenericHard.parent }

func (asicErrorGenericHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error) GetParentYangName() string { return "asic-error-generic-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorGenericHard_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors
// Parity hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetFilter() yfilter.YFilter { return parityHardErrors.YFilter }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) SetFilter(yf yfilter.YFilter) { parityHardErrors.YFilter = yf }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetSegmentPath() string {
    return "parity-hard-errors"
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range parityHardErrors.Error {
            if parityHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error{}
        parityHardErrors.Error = append(parityHardErrors.Error, child)
        return &parityHardErrors.Error[len(parityHardErrors.Error)-1]
    }
    return nil
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range parityHardErrors.Error {
        children[parityHardErrors.Error[i].GetSegmentPath()] = &parityHardErrors.Error[i]
    }
    return children
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetYangName() string { return "parity-hard-errors" }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) SetParent(parent types.Entity) { parityHardErrors.parent = parent }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetParent() types.Entity { return parityHardErrors.parent }

func (parityHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error) GetParentYangName() string { return "parity-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParityHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors
// Descriptor hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetFilter() yfilter.YFilter { return descriptorHardErrors.YFilter }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) SetFilter(yf yfilter.YFilter) { descriptorHardErrors.YFilter = yf }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetSegmentPath() string {
    return "descriptor-hard-errors"
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range descriptorHardErrors.Error {
            if descriptorHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error{}
        descriptorHardErrors.Error = append(descriptorHardErrors.Error, child)
        return &descriptorHardErrors.Error[len(descriptorHardErrors.Error)-1]
    }
    return nil
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range descriptorHardErrors.Error {
        children[descriptorHardErrors.Error[i].GetSegmentPath()] = &descriptorHardErrors.Error[i]
    }
    return children
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetYangName() string { return "descriptor-hard-errors" }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) SetParent(parent types.Entity) { descriptorHardErrors.parent = parent }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetParent() types.Entity { return descriptorHardErrors.parent }

func (descriptorHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error) GetParentYangName() string { return "descriptor-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors
// Interface hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetFilter() yfilter.YFilter { return interfaceHardErrors.YFilter }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) SetFilter(yf yfilter.YFilter) { interfaceHardErrors.YFilter = yf }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetSegmentPath() string {
    return "interface-hard-errors"
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range interfaceHardErrors.Error {
            if interfaceHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error{}
        interfaceHardErrors.Error = append(interfaceHardErrors.Error, child)
        return &interfaceHardErrors.Error[len(interfaceHardErrors.Error)-1]
    }
    return nil
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceHardErrors.Error {
        children[interfaceHardErrors.Error[i].GetSegmentPath()] = &interfaceHardErrors.Error[i]
    }
    return children
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetYangName() string { return "interface-hard-errors" }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) SetParent(parent types.Entity) { interfaceHardErrors.parent = parent }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetParent() types.Entity { return interfaceHardErrors.parent }

func (interfaceHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error) GetParentYangName() string { return "interface-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetFilter() yfilter.YFilter { return asicErrorSbeHard.YFilter }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) SetFilter(yf yfilter.YFilter) { asicErrorSbeHard.YFilter = yf }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetSegmentPath() string {
    return "asic-error-sbe-hard"
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorSbeHard.Error {
            if asicErrorSbeHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error{}
        asicErrorSbeHard.Error = append(asicErrorSbeHard.Error, child)
        return &asicErrorSbeHard.Error[len(asicErrorSbeHard.Error)-1]
    }
    return nil
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorSbeHard.Error {
        children[asicErrorSbeHard.Error[i].GetSegmentPath()] = &asicErrorSbeHard.Error[i]
    }
    return children
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetYangName() string { return "asic-error-sbe-hard" }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) SetParent(parent types.Entity) { asicErrorSbeHard.parent = parent }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetParent() types.Entity { return asicErrorSbeHard.parent }

func (asicErrorSbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error) GetParentYangName() string { return "asic-error-sbe-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorSbeHard_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetFilter() yfilter.YFilter { return asicErrorCrcHard.YFilter }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) SetFilter(yf yfilter.YFilter) { asicErrorCrcHard.YFilter = yf }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetSegmentPath() string {
    return "asic-error-crc-hard"
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorCrcHard.Error {
            if asicErrorCrcHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error{}
        asicErrorCrcHard.Error = append(asicErrorCrcHard.Error, child)
        return &asicErrorCrcHard.Error[len(asicErrorCrcHard.Error)-1]
    }
    return nil
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorCrcHard.Error {
        children[asicErrorCrcHard.Error[i].GetSegmentPath()] = &asicErrorCrcHard.Error[i]
    }
    return children
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetYangName() string { return "asic-error-crc-hard" }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) SetParent(parent types.Entity) { asicErrorCrcHard.parent = parent }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetParent() types.Entity { return asicErrorCrcHard.parent }

func (asicErrorCrcHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error) GetParentYangName() string { return "asic-error-crc-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorCrcHard_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetFilter() yfilter.YFilter { return asicErrorParityHard.YFilter }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) SetFilter(yf yfilter.YFilter) { asicErrorParityHard.YFilter = yf }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetSegmentPath() string {
    return "asic-error-parity-hard"
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorParityHard.Error {
            if asicErrorParityHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error{}
        asicErrorParityHard.Error = append(asicErrorParityHard.Error, child)
        return &asicErrorParityHard.Error[len(asicErrorParityHard.Error)-1]
    }
    return nil
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorParityHard.Error {
        children[asicErrorParityHard.Error[i].GetSegmentPath()] = &asicErrorParityHard.Error[i]
    }
    return children
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetYangName() string { return "asic-error-parity-hard" }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) SetParent(parent types.Entity) { asicErrorParityHard.parent = parent }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetParent() types.Entity { return asicErrorParityHard.parent }

func (asicErrorParityHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error) GetParentYangName() string { return "asic-error-parity-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorParityHard_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetFilter() yfilter.YFilter { return asicErrorResetSoft.YFilter }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) SetFilter(yf yfilter.YFilter) { asicErrorResetSoft.YFilter = yf }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetSegmentPath() string {
    return "asic-error-reset-soft"
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorResetSoft.Error {
            if asicErrorResetSoft.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error{}
        asicErrorResetSoft.Error = append(asicErrorResetSoft.Error, child)
        return &asicErrorResetSoft.Error[len(asicErrorResetSoft.Error)-1]
    }
    return nil
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorResetSoft.Error {
        children[asicErrorResetSoft.Error[i].GetSegmentPath()] = &asicErrorResetSoft.Error[i]
    }
    return children
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetYangName() string { return "asic-error-reset-soft" }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) SetParent(parent types.Entity) { asicErrorResetSoft.parent = parent }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetParent() types.Entity { return asicErrorResetSoft.parent }

func (asicErrorResetSoft *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error) GetParentYangName() string { return "asic-error-reset-soft" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorResetSoft_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors
// BP soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetFilter() yfilter.YFilter { return backPressureSoftErrors.YFilter }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) SetFilter(yf yfilter.YFilter) { backPressureSoftErrors.YFilter = yf }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetSegmentPath() string {
    return "back-pressure-soft-errors"
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range backPressureSoftErrors.Error {
            if backPressureSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error{}
        backPressureSoftErrors.Error = append(backPressureSoftErrors.Error, child)
        return &backPressureSoftErrors.Error[len(backPressureSoftErrors.Error)-1]
    }
    return nil
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range backPressureSoftErrors.Error {
        children[backPressureSoftErrors.Error[i].GetSegmentPath()] = &backPressureSoftErrors.Error[i]
    }
    return children
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetYangName() string { return "back-pressure-soft-errors" }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) SetParent(parent types.Entity) { backPressureSoftErrors.parent = parent }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetParent() types.Entity { return backPressureSoftErrors.parent }

func (backPressureSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error) GetParentYangName() string { return "back-pressure-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_BackPressureSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors
// Generic soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetFilter() yfilter.YFilter { return genericSoftErrors.YFilter }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) SetFilter(yf yfilter.YFilter) { genericSoftErrors.YFilter = yf }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetSegmentPath() string {
    return "generic-soft-errors"
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range genericSoftErrors.Error {
            if genericSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error{}
        genericSoftErrors.Error = append(genericSoftErrors.Error, child)
        return &genericSoftErrors.Error[len(genericSoftErrors.Error)-1]
    }
    return nil
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range genericSoftErrors.Error {
        children[genericSoftErrors.Error[i].GetSegmentPath()] = &genericSoftErrors.Error[i]
    }
    return children
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetYangName() string { return "generic-soft-errors" }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) SetParent(parent types.Entity) { genericSoftErrors.parent = parent }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetParent() types.Entity { return genericSoftErrors.parent }

func (genericSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error) GetParentYangName() string { return "generic-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_GenericSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors
// Link soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetFilter() yfilter.YFilter { return linkSoftErrors.YFilter }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) SetFilter(yf yfilter.YFilter) { linkSoftErrors.YFilter = yf }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetSegmentPath() string {
    return "link-soft-errors"
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range linkSoftErrors.Error {
            if linkSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error{}
        linkSoftErrors.Error = append(linkSoftErrors.Error, child)
        return &linkSoftErrors.Error[len(linkSoftErrors.Error)-1]
    }
    return nil
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkSoftErrors.Error {
        children[linkSoftErrors.Error[i].GetSegmentPath()] = &linkSoftErrors.Error[i]
    }
    return children
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetYangName() string { return "link-soft-errors" }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) SetParent(parent types.Entity) { linkSoftErrors.parent = parent }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetParent() types.Entity { return linkSoftErrors.parent }

func (linkSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error) GetParentYangName() string { return "link-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_LinkSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors
// Configuration soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetFilter() yfilter.YFilter { return configurationSoftErrors.YFilter }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) SetFilter(yf yfilter.YFilter) { configurationSoftErrors.YFilter = yf }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetSegmentPath() string {
    return "configuration-soft-errors"
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range configurationSoftErrors.Error {
            if configurationSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error{}
        configurationSoftErrors.Error = append(configurationSoftErrors.Error, child)
        return &configurationSoftErrors.Error[len(configurationSoftErrors.Error)-1]
    }
    return nil
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configurationSoftErrors.Error {
        children[configurationSoftErrors.Error[i].GetSegmentPath()] = &configurationSoftErrors.Error[i]
    }
    return children
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetYangName() string { return "configuration-soft-errors" }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) SetParent(parent types.Entity) { configurationSoftErrors.parent = parent }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetParent() types.Entity { return configurationSoftErrors.parent }

func (configurationSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error) GetParentYangName() string { return "configuration-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ConfigurationSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors
// Multiple bit hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetFilter() yfilter.YFilter { return multipleBitHardErrors.YFilter }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) SetFilter(yf yfilter.YFilter) { multipleBitHardErrors.YFilter = yf }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetSegmentPath() string {
    return "multiple-bit-hard-errors"
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range multipleBitHardErrors.Error {
            if multipleBitHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error{}
        multipleBitHardErrors.Error = append(multipleBitHardErrors.Error, child)
        return &multipleBitHardErrors.Error[len(multipleBitHardErrors.Error)-1]
    }
    return nil
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multipleBitHardErrors.Error {
        children[multipleBitHardErrors.Error[i].GetSegmentPath()] = &multipleBitHardErrors.Error[i]
    }
    return children
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetYangName() string { return "multiple-bit-hard-errors" }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) SetParent(parent types.Entity) { multipleBitHardErrors.parent = parent }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetParent() types.Entity { return multipleBitHardErrors.parent }

func (multipleBitHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error) GetParentYangName() string { return "multiple-bit-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_MultipleBitHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors
// Unexpected soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetFilter() yfilter.YFilter { return unexpectedSoftErrors.YFilter }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) SetFilter(yf yfilter.YFilter) { unexpectedSoftErrors.YFilter = yf }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetSegmentPath() string {
    return "unexpected-soft-errors"
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range unexpectedSoftErrors.Error {
            if unexpectedSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error{}
        unexpectedSoftErrors.Error = append(unexpectedSoftErrors.Error, child)
        return &unexpectedSoftErrors.Error[len(unexpectedSoftErrors.Error)-1]
    }
    return nil
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unexpectedSoftErrors.Error {
        children[unexpectedSoftErrors.Error[i].GetSegmentPath()] = &unexpectedSoftErrors.Error[i]
    }
    return children
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetYangName() string { return "unexpected-soft-errors" }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) SetParent(parent types.Entity) { unexpectedSoftErrors.parent = parent }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetParent() types.Entity { return unexpectedSoftErrors.parent }

func (unexpectedSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error) GetParentYangName() string { return "unexpected-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UnexpectedSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard
// OOR thresh information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetFilter() yfilter.YFilter { return outofResourceHard.YFilter }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) SetFilter(yf yfilter.YFilter) { outofResourceHard.YFilter = yf }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetSegmentPath() string {
    return "outof-resource-hard"
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range outofResourceHard.Error {
            if outofResourceHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error{}
        outofResourceHard.Error = append(outofResourceHard.Error, child)
        return &outofResourceHard.Error[len(outofResourceHard.Error)-1]
    }
    return nil
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range outofResourceHard.Error {
        children[outofResourceHard.Error[i].GetSegmentPath()] = &outofResourceHard.Error[i]
    }
    return children
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetBundleName() string { return "cisco_ios_xr" }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetYangName() string { return "outof-resource-hard" }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) SetParent(parent types.Entity) { outofResourceHard.parent = parent }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetParent() types.Entity { return outofResourceHard.parent }

func (outofResourceHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error) GetParentYangName() string { return "outof-resource-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_OutofResourceHard_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors
// Hardware hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetFilter() yfilter.YFilter { return hardwareHardErrors.YFilter }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) SetFilter(yf yfilter.YFilter) { hardwareHardErrors.YFilter = yf }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetSegmentPath() string {
    return "hardware-hard-errors"
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range hardwareHardErrors.Error {
            if hardwareHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error{}
        hardwareHardErrors.Error = append(hardwareHardErrors.Error, child)
        return &hardwareHardErrors.Error[len(hardwareHardErrors.Error)-1]
    }
    return nil
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareHardErrors.Error {
        children[hardwareHardErrors.Error[i].GetSegmentPath()] = &hardwareHardErrors.Error[i]
    }
    return children
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetYangName() string { return "hardware-hard-errors" }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) SetParent(parent types.Entity) { hardwareHardErrors.parent = parent }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetParent() types.Entity { return hardwareHardErrors.parent }

func (hardwareHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error) GetParentYangName() string { return "hardware-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_HardwareHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors
// Parity soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetFilter() yfilter.YFilter { return paritySoftErrors.YFilter }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) SetFilter(yf yfilter.YFilter) { paritySoftErrors.YFilter = yf }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetSegmentPath() string {
    return "parity-soft-errors"
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range paritySoftErrors.Error {
            if paritySoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error{}
        paritySoftErrors.Error = append(paritySoftErrors.Error, child)
        return &paritySoftErrors.Error[len(paritySoftErrors.Error)-1]
    }
    return nil
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paritySoftErrors.Error {
        children[paritySoftErrors.Error[i].GetSegmentPath()] = &paritySoftErrors.Error[i]
    }
    return children
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetYangName() string { return "parity-soft-errors" }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) SetParent(parent types.Entity) { paritySoftErrors.parent = parent }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetParent() types.Entity { return paritySoftErrors.parent }

func (paritySoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error) GetParentYangName() string { return "parity-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ParitySoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors
// Descriptor soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetFilter() yfilter.YFilter { return descriptorSoftErrors.YFilter }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) SetFilter(yf yfilter.YFilter) { descriptorSoftErrors.YFilter = yf }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetSegmentPath() string {
    return "descriptor-soft-errors"
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range descriptorSoftErrors.Error {
            if descriptorSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error{}
        descriptorSoftErrors.Error = append(descriptorSoftErrors.Error, child)
        return &descriptorSoftErrors.Error[len(descriptorSoftErrors.Error)-1]
    }
    return nil
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range descriptorSoftErrors.Error {
        children[descriptorSoftErrors.Error[i].GetSegmentPath()] = &descriptorSoftErrors.Error[i]
    }
    return children
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetYangName() string { return "descriptor-soft-errors" }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) SetParent(parent types.Entity) { descriptorSoftErrors.parent = parent }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetParent() types.Entity { return descriptorSoftErrors.parent }

func (descriptorSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error) GetParentYangName() string { return "descriptor-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_DescriptorSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors
// Interface soft error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetFilter() yfilter.YFilter { return interfaceSoftErrors.YFilter }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) SetFilter(yf yfilter.YFilter) { interfaceSoftErrors.YFilter = yf }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetSegmentPath() string {
    return "interface-soft-errors"
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range interfaceSoftErrors.Error {
            if interfaceSoftErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error{}
        interfaceSoftErrors.Error = append(interfaceSoftErrors.Error, child)
        return &interfaceSoftErrors.Error[len(interfaceSoftErrors.Error)-1]
    }
    return nil
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceSoftErrors.Error {
        children[interfaceSoftErrors.Error[i].GetSegmentPath()] = &interfaceSoftErrors.Error[i]
    }
    return children
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetYangName() string { return "interface-soft-errors" }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) SetParent(parent types.Entity) { interfaceSoftErrors.parent = parent }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetParent() types.Entity { return interfaceSoftErrors.parent }

func (interfaceSoftErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error) GetParentYangName() string { return "interface-soft-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_InterfaceSoftErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors
// IO hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetFilter() yfilter.YFilter { return ioHardErrors.YFilter }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) SetFilter(yf yfilter.YFilter) { ioHardErrors.YFilter = yf }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetSegmentPath() string {
    return "io-hard-errors"
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range ioHardErrors.Error {
            if ioHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error{}
        ioHardErrors.Error = append(ioHardErrors.Error, child)
        return &ioHardErrors.Error[len(ioHardErrors.Error)-1]
    }
    return nil
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ioHardErrors.Error {
        children[ioHardErrors.Error[i].GetSegmentPath()] = &ioHardErrors.Error[i]
    }
    return children
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetYangName() string { return "io-hard-errors" }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) SetParent(parent types.Entity) { ioHardErrors.parent = parent }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetParent() types.Entity { return ioHardErrors.parent }

func (ioHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error) GetParentYangName() string { return "io-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_IoHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors
// Reset hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetFilter() yfilter.YFilter { return resetHardErrors.YFilter }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) SetFilter(yf yfilter.YFilter) { resetHardErrors.YFilter = yf }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetSegmentPath() string {
    return "reset-hard-errors"
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range resetHardErrors.Error {
            if resetHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error{}
        resetHardErrors.Error = append(resetHardErrors.Error, child)
        return &resetHardErrors.Error[len(resetHardErrors.Error)-1]
    }
    return nil
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range resetHardErrors.Error {
        children[resetHardErrors.Error[i].GetSegmentPath()] = &resetHardErrors.Error[i]
    }
    return children
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetYangName() string { return "reset-hard-errors" }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) SetParent(parent types.Entity) { resetHardErrors.parent = parent }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetParent() types.Entity { return resetHardErrors.parent }

func (resetHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error) GetParentYangName() string { return "reset-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_ResetHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors
// UCode hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetFilter() yfilter.YFilter { return ucodeHardErrors.YFilter }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) SetFilter(yf yfilter.YFilter) { ucodeHardErrors.YFilter = yf }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetSegmentPath() string {
    return "ucode-hard-errors"
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range ucodeHardErrors.Error {
            if ucodeHardErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error{}
        ucodeHardErrors.Error = append(ucodeHardErrors.Error, child)
        return &ucodeHardErrors.Error[len(ucodeHardErrors.Error)-1]
    }
    return nil
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ucodeHardErrors.Error {
        children[ucodeHardErrors.Error[i].GetSegmentPath()] = &ucodeHardErrors.Error[i]
    }
    return children
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetBundleName() string { return "cisco_ios_xr" }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetYangName() string { return "ucode-hard-errors" }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) SetParent(parent types.Entity) { ucodeHardErrors.parent = parent }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetParent() types.Entity { return ucodeHardErrors.parent }

func (ucodeHardErrors *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error) GetParentYangName() string { return "ucode-hard-errors" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_UcodeHardErrors_Error_LastErr) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard
// Indirect hard error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of errors. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error.
    Error []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetFilter() yfilter.YFilter { return asicErrorMbeHard.YFilter }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) SetFilter(yf yfilter.YFilter) { asicErrorMbeHard.YFilter = yf }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetSegmentPath() string {
    return "asic-error-mbe-hard"
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        for _, c := range asicErrorMbeHard.Error {
            if asicErrorMbeHard.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error{}
        asicErrorMbeHard.Error = append(asicErrorMbeHard.Error, child)
        return &asicErrorMbeHard.Error[len(asicErrorMbeHard.Error)-1]
    }
    return nil
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicErrorMbeHard.Error {
        children[asicErrorMbeHard.Error[i].GetSegmentPath()] = &asicErrorMbeHard.Error[i]
    }
    return children
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetYangName() string { return "asic-error-mbe-hard" }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) SetParent(parent types.Entity) { asicErrorMbeHard.parent = parent }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetParent() types.Entity { return asicErrorMbeHard.parent }

func (asicErrorMbeHard *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard) GetParentYangName() string { return "error-path" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error
// Collection of errors
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name assigned to mem. The type is string.
    Name interface{}

    // Name of rack/board/asic. The type is string.
    AsicInfo interface{}

    // 32 bit key. The type is interface{} with range: 0..4294967295.
    NodeKey interface{}

    // High threshold crossed. The type is bool.
    AlarmOn interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshHi interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodHi interface{}

    // High threshold value. The type is interface{} with range: 0..4294967295.
    ThreshLo interface{}

    // High period value. The type is interface{} with range: 0..4294967295.
    PeriodLo interface{}

    // Accumulated count. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Type of error. The type is interface{} with range: 0..4294967295.
    IntrType interface{}

    // Leaf ID defined in user data. The type is interface{} with range:
    // 0..4294967295.
    LeafId interface{}

    // Time  cleared. The type is interface{} with range: 0..18446744073709551615.
    LastCleared interface{}

    // List of csrs_info. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo.
    CsrsInfo []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo

    // Last Printable error information. The type is slice of
    // AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr.
    LastErr []AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "asic-info" { return "AsicInfo" }
    if yname == "node-key" { return "NodeKey" }
    if yname == "alarm-on" { return "AlarmOn" }
    if yname == "thresh-hi" { return "ThreshHi" }
    if yname == "period-hi" { return "PeriodHi" }
    if yname == "thresh-lo" { return "ThreshLo" }
    if yname == "period-lo" { return "PeriodLo" }
    if yname == "count" { return "Count" }
    if yname == "intr-type" { return "IntrType" }
    if yname == "leaf-id" { return "LeafId" }
    if yname == "last-cleared" { return "LastCleared" }
    if yname == "csrs-info" { return "CsrsInfo" }
    if yname == "last-err" { return "LastErr" }
    return ""
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetSegmentPath() string {
    return "error"
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csrs-info" {
        for _, c := range error.CsrsInfo {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo{}
        error.CsrsInfo = append(error.CsrsInfo, child)
        return &error.CsrsInfo[len(error.CsrsInfo)-1]
    }
    if childYangName == "last-err" {
        for _, c := range error.LastErr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr{}
        error.LastErr = append(error.LastErr, child)
        return &error.LastErr[len(error.LastErr)-1]
    }
    return nil
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range error.CsrsInfo {
        children[error.CsrsInfo[i].GetSegmentPath()] = &error.CsrsInfo[i]
    }
    for i := range error.LastErr {
        children[error.LastErr[i].GetSegmentPath()] = &error.LastErr[i]
    }
    return children
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = error.Name
    leafs["asic-info"] = error.AsicInfo
    leafs["node-key"] = error.NodeKey
    leafs["alarm-on"] = error.AlarmOn
    leafs["thresh-hi"] = error.ThreshHi
    leafs["period-hi"] = error.PeriodHi
    leafs["thresh-lo"] = error.ThreshLo
    leafs["period-lo"] = error.PeriodLo
    leafs["count"] = error.Count
    leafs["intr-type"] = error.IntrType
    leafs["leaf-id"] = error.LeafId
    leafs["last-cleared"] = error.LastCleared
    return leafs
}

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetYangName() string { return "error" }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetParent() types.Entity { return error.parent }

func (error *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error) GetParentYangName() string { return "asic-error-mbe-hard" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo
// List of csrs_info
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // name. The type is string.
    Name interface{}

    // address. The type is interface{} with range: 0..18446744073709551615.
    Address interface{}

    // width. The type is interface{} with range: 0..4294967295.
    Width interface{}
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetFilter() yfilter.YFilter { return csrsInfo.YFilter }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) SetFilter(yf yfilter.YFilter) { csrsInfo.YFilter = yf }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "address" { return "Address" }
    if yname == "width" { return "Width" }
    return ""
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetSegmentPath() string {
    return "csrs-info"
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = csrsInfo.Name
    leafs["address"] = csrsInfo.Address
    leafs["width"] = csrsInfo.Width
    return leafs
}

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetYangName() string { return "csrs-info" }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) SetParent(parent types.Entity) { csrsInfo.parent = parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetParent() types.Entity { return csrsInfo.parent }

func (csrsInfo *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_CsrsInfo) GetParentYangName() string { return "error" }

// AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr
// Last Printable error information
type AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // at time. The type is interface{} with range: 0..18446744073709551615.
    AtTime interface{}

    // at time nsec. The type is interface{} with range: 0..18446744073709551615.
    AtTimeNsec interface{}

    // counter val. The type is interface{} with range: 0..4294967295.
    CounterVal interface{}

    // error desc. The type is string.
    ErrorDesc interface{}

    // error regval. The type is slice of interface{} with range: 0..255.
    ErrorRegval []interface{}
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetFilter() yfilter.YFilter { return lastErr.YFilter }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) SetFilter(yf yfilter.YFilter) { lastErr.YFilter = yf }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetGoName(yname string) string {
    if yname == "at-time" { return "AtTime" }
    if yname == "at-time-nsec" { return "AtTimeNsec" }
    if yname == "counter-val" { return "CounterVal" }
    if yname == "error-desc" { return "ErrorDesc" }
    if yname == "error-regval" { return "ErrorRegval" }
    return ""
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetSegmentPath() string {
    return "last-err"
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["at-time"] = lastErr.AtTime
    leafs["at-time-nsec"] = lastErr.AtTimeNsec
    leafs["counter-val"] = lastErr.CounterVal
    leafs["error-desc"] = lastErr.ErrorDesc
    leafs["error-regval"] = lastErr.ErrorRegval
    return leafs
}

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetBundleName() string { return "cisco_ios_xr" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetYangName() string { return "last-err" }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) SetParent(parent types.Entity) { lastErr.parent = parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetParent() types.Entity { return lastErr.parent }

func (lastErr *AsicErrors_Nodes_Node_AsicInformation_Instances_Instance_ErrorPath_AsicErrorMbeHard_Error_LastErr) GetParentYangName() string { return "error" }

