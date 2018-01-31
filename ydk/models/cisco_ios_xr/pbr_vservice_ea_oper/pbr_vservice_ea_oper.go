// This module contains a collection of YANG definitions
// for Cisco IOS-XR pbr-vservice-ea package operational data.
// 
// This module contains definitions
// for the following management objects:
//   service-function-chaining: NSH Service Function Chaining
//     operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pbr_vservice_ea_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pbr_vservice_ea_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pbr-vservice-ea-oper service-function-chaining}", reflect.TypeOf(ServiceFunctionChaining{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pbr-vservice-ea-oper:service-function-chaining", reflect.TypeOf(ServiceFunctionChaining{}))
}

// VsNshStats represents Vs nsh stats
type VsNshStats string

const (
    // vs nsh stats spi si
    VsNshStats_vs_nsh_stats_spi_si VsNshStats = "vs-nsh-stats-spi-si"

    // vs nsh stats ter min ate
    VsNshStats_vs_nsh_stats_ter_min_ate VsNshStats = "vs-nsh-stats-ter-min-ate"

    // vs nsh stats sf
    VsNshStats_vs_nsh_stats_sf VsNshStats = "vs-nsh-stats-sf"

    // vs nsh stats sff
    VsNshStats_vs_nsh_stats_sff VsNshStats = "vs-nsh-stats-sff"

    // vs nsh stats sff local
    VsNshStats_vs_nsh_stats_sff_local VsNshStats = "vs-nsh-stats-sff-local"

    // vs nsh stats sfp
    VsNshStats_vs_nsh_stats_sfp VsNshStats = "vs-nsh-stats-sfp"

    // vs nsh stats sfp detail
    VsNshStats_vs_nsh_stats_sfp_detail VsNshStats = "vs-nsh-stats-sfp-detail"

    // vs nsh stats max
    VsNshStats_vs_nsh_stats_max VsNshStats = "vs-nsh-stats-max"
)

// ServiceFunctionChaining
// NSH Service Function Chaining operational data
type ServiceFunctionChaining struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific NSH Service Function Chaining operational data.
    Nodes ServiceFunctionChaining_Nodes
}

func (serviceFunctionChaining *ServiceFunctionChaining) GetFilter() yfilter.YFilter { return serviceFunctionChaining.YFilter }

func (serviceFunctionChaining *ServiceFunctionChaining) SetFilter(yf yfilter.YFilter) { serviceFunctionChaining.YFilter = yf }

func (serviceFunctionChaining *ServiceFunctionChaining) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (serviceFunctionChaining *ServiceFunctionChaining) GetSegmentPath() string {
    return "Cisco-IOS-XR-pbr-vservice-ea-oper:service-function-chaining"
}

func (serviceFunctionChaining *ServiceFunctionChaining) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &serviceFunctionChaining.Nodes
    }
    return nil
}

func (serviceFunctionChaining *ServiceFunctionChaining) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &serviceFunctionChaining.Nodes
    return children
}

func (serviceFunctionChaining *ServiceFunctionChaining) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunctionChaining *ServiceFunctionChaining) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionChaining *ServiceFunctionChaining) GetYangName() string { return "service-function-chaining" }

func (serviceFunctionChaining *ServiceFunctionChaining) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionChaining *ServiceFunctionChaining) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionChaining *ServiceFunctionChaining) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionChaining *ServiceFunctionChaining) SetParent(parent types.Entity) { serviceFunctionChaining.parent = parent }

func (serviceFunctionChaining *ServiceFunctionChaining) GetParent() types.Entity { return serviceFunctionChaining.parent }

func (serviceFunctionChaining *ServiceFunctionChaining) GetParentYangName() string { return "Cisco-IOS-XR-pbr-vservice-ea-oper" }

// ServiceFunctionChaining_Nodes
// Node-specific NSH Service Function Chaining
// operational data
type ServiceFunctionChaining_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NSH operational data for a particular node. The type is slice of
    // ServiceFunctionChaining_Nodes_Node.
    Node []ServiceFunctionChaining_Nodes_Node
}

func (nodes *ServiceFunctionChaining_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *ServiceFunctionChaining_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *ServiceFunctionChaining_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *ServiceFunctionChaining_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *ServiceFunctionChaining_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *ServiceFunctionChaining_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *ServiceFunctionChaining_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *ServiceFunctionChaining_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *ServiceFunctionChaining_Nodes) GetYangName() string { return "nodes" }

func (nodes *ServiceFunctionChaining_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *ServiceFunctionChaining_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *ServiceFunctionChaining_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *ServiceFunctionChaining_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *ServiceFunctionChaining_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *ServiceFunctionChaining_Nodes) GetParentYangName() string { return "service-function-chaining" }

// ServiceFunctionChaining_Nodes_Node
// NSH operational data for a particular node
type ServiceFunctionChaining_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node to collect statistics from. The type is
    // string with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Client Process.
    Process ServiceFunctionChaining_Nodes_Node_Process
}

func (node *ServiceFunctionChaining_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *ServiceFunctionChaining_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *ServiceFunctionChaining_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "process" { return "Process" }
    return ""
}

func (node *ServiceFunctionChaining_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *ServiceFunctionChaining_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process" {
        return &node.Process
    }
    return nil
}

func (node *ServiceFunctionChaining_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process"] = &node.Process
    return children
}

func (node *ServiceFunctionChaining_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *ServiceFunctionChaining_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *ServiceFunctionChaining_Nodes_Node) GetYangName() string { return "node" }

func (node *ServiceFunctionChaining_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *ServiceFunctionChaining_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *ServiceFunctionChaining_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *ServiceFunctionChaining_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *ServiceFunctionChaining_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *ServiceFunctionChaining_Nodes_Node) GetParentYangName() string { return "nodes" }

// ServiceFunctionChaining_Nodes_Node_Process
// Client Process
type ServiceFunctionChaining_Nodes_Node_Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service Function Path operational data.
    ServiceFunctionPath ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath

    // Service Function operational data.
    ServiceFunction ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction

    // Service Function Forwarder operational data.
    ServiceFunctionForwarder ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder
}

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetFilter() yfilter.YFilter { return process.YFilter }

func (process *ServiceFunctionChaining_Nodes_Node_Process) SetFilter(yf yfilter.YFilter) { process.YFilter = yf }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetGoName(yname string) string {
    if yname == "service-function-path" { return "ServiceFunctionPath" }
    if yname == "service-function" { return "ServiceFunction" }
    if yname == "service-function-forwarder" { return "ServiceFunctionForwarder" }
    return ""
}

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetSegmentPath() string {
    return "process"
}

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-function-path" {
        return &process.ServiceFunctionPath
    }
    if childYangName == "service-function" {
        return &process.ServiceFunction
    }
    if childYangName == "service-function-forwarder" {
        return &process.ServiceFunctionForwarder
    }
    return nil
}

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-function-path"] = &process.ServiceFunctionPath
    children["service-function"] = &process.ServiceFunction
    children["service-function-forwarder"] = &process.ServiceFunctionForwarder
    return children
}

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetBundleName() string { return "cisco_ios_xr" }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetYangName() string { return "process" }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (process *ServiceFunctionChaining_Nodes_Node_Process) SetParent(parent types.Entity) { process.parent = parent }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetParent() types.Entity { return process.parent }

func (process *ServiceFunctionChaining_Nodes_Node_Process) GetParentYangName() string { return "node" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath
// Service Function Path operational data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service Function Path Id .
    PathIds ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds
}

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetFilter() yfilter.YFilter { return serviceFunctionPath.YFilter }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) SetFilter(yf yfilter.YFilter) { serviceFunctionPath.YFilter = yf }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetGoName(yname string) string {
    if yname == "path-ids" { return "PathIds" }
    return ""
}

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetSegmentPath() string {
    return "service-function-path"
}

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-ids" {
        return &serviceFunctionPath.PathIds
    }
    return nil
}

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-ids"] = &serviceFunctionPath.PathIds
    return children
}

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetYangName() string { return "service-function-path" }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) SetParent(parent types.Entity) { serviceFunctionPath.parent = parent }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetParent() types.Entity { return serviceFunctionPath.parent }

func (serviceFunctionPath *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath) GetParentYangName() string { return "process" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds
// Service Function Path Id 
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specific Service-Function-Path identifier . The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId.
    PathId []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId
}

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetFilter() yfilter.YFilter { return pathIds.YFilter }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) SetFilter(yf yfilter.YFilter) { pathIds.YFilter = yf }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    return ""
}

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetSegmentPath() string {
    return "path-ids"
}

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-id" {
        for _, c := range pathIds.PathId {
            if pathIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId{}
        pathIds.PathId = append(pathIds.PathId, child)
        return &pathIds.PathId[len(pathIds.PathId)-1]
    }
    return nil
}

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pathIds.PathId {
        children[pathIds.PathId[i].GetSegmentPath()] = &pathIds.PathId[i]
    }
    return children
}

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetBundleName() string { return "cisco_ios_xr" }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetYangName() string { return "path-ids" }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) SetParent(parent types.Entity) { pathIds.parent = parent }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetParent() types.Entity { return pathIds.parent }

func (pathIds *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds) GetParentYangName() string { return "service-function-path" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId
// Specific Service-Function-Path identifier 
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Service-Function-Path identifier. The
    // type is interface{} with range: 1..16777215.
    Id interface{}

    // Service Index Belonging to Path.
    ServiceIndexes ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes

    // SFP Statistics.
    Stats ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats
}

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetFilter() yfilter.YFilter { return pathId.YFilter }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) SetFilter(yf yfilter.YFilter) { pathId.YFilter = yf }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "service-indexes" { return "ServiceIndexes" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetSegmentPath() string {
    return "path-id" + "[id='" + fmt.Sprintf("%v", pathId.Id) + "']"
}

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-indexes" {
        return &pathId.ServiceIndexes
    }
    if childYangName == "stats" {
        return &pathId.Stats
    }
    return nil
}

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-indexes"] = &pathId.ServiceIndexes
    children["stats"] = &pathId.Stats
    return children
}

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = pathId.Id
    return leafs
}

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetBundleName() string { return "cisco_ios_xr" }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetYangName() string { return "path-id" }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) SetParent(parent types.Entity) { pathId.parent = parent }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetParent() types.Entity { return pathId.parent }

func (pathId *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId) GetParentYangName() string { return "path-ids" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes
// Service Index Belonging to Path
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index operational data belonging to this path. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex.
    ServiceIndex []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex
}

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetFilter() yfilter.YFilter { return serviceIndexes.YFilter }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) SetFilter(yf yfilter.YFilter) { serviceIndexes.YFilter = yf }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetGoName(yname string) string {
    if yname == "service-index" { return "ServiceIndex" }
    return ""
}

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetSegmentPath() string {
    return "service-indexes"
}

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-index" {
        for _, c := range serviceIndexes.ServiceIndex {
            if serviceIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex{}
        serviceIndexes.ServiceIndex = append(serviceIndexes.ServiceIndex, child)
        return &serviceIndexes.ServiceIndex[len(serviceIndexes.ServiceIndex)-1]
    }
    return nil
}

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serviceIndexes.ServiceIndex {
        children[serviceIndexes.ServiceIndex[i].GetSegmentPath()] = &serviceIndexes.ServiceIndex[i]
    }
    return children
}

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetYangName() string { return "service-indexes" }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) SetParent(parent types.Entity) { serviceIndexes.parent = parent }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetParent() types.Entity { return serviceIndexes.parent }

func (serviceIndexes *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes) GetParentYangName() string { return "path-id" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex
// Service index operational data belonging
// to this path
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Service Index. The type is interface{} with range:
    // 1..255.
    Index interface{}

    // Statistics data.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data

    // SI array in case of detail stats. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr.
    SiArr []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr
}

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetFilter() yfilter.YFilter { return serviceIndex.YFilter }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) SetFilter(yf yfilter.YFilter) { serviceIndex.YFilter = yf }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "data" { return "Data" }
    if yname == "si-arr" { return "SiArr" }
    return ""
}

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetSegmentPath() string {
    return "service-index" + "[index='" + fmt.Sprintf("%v", serviceIndex.Index) + "']"
}

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &serviceIndex.Data
    }
    if childYangName == "si-arr" {
        for _, c := range serviceIndex.SiArr {
            if serviceIndex.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr{}
        serviceIndex.SiArr = append(serviceIndex.SiArr, child)
        return &serviceIndex.SiArr[len(serviceIndex.SiArr)-1]
    }
    return nil
}

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &serviceIndex.Data
    for i := range serviceIndex.SiArr {
        children[serviceIndex.SiArr[i].GetSegmentPath()] = &serviceIndex.SiArr[i]
    }
    return children
}

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = serviceIndex.Index
    return leafs
}

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetBundleName() string { return "cisco_ios_xr" }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetYangName() string { return "service-index" }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) SetParent(parent types.Entity) { serviceIndex.parent = parent }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetParent() types.Entity { return serviceIndex.parent }

func (serviceIndex *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex) GetParentYangName() string { return "service-indexes" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data
// Statistics data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SFP stats.
    Sfp ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp

    // SPI SI stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term

    // Service function stats.
    Sf ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf

    // Service function forwarder stats.
    Sff ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff

    // Local service function forwarder stats.
    SffLocal ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sfp" { return "Sfp" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    if yname == "sf" { return "Sf" }
    if yname == "sff" { return "Sff" }
    if yname == "sff-local" { return "SffLocal" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sfp" {
        return &data.Sfp
    }
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    if childYangName == "sf" {
        return &data.Sf
    }
    if childYangName == "sff" {
        return &data.Sff
    }
    if childYangName == "sff-local" {
        return &data.SffLocal
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sfp"] = &data.Sfp
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    children["sf"] = &data.Sf
    children["sff"] = &data.Sff
    children["sff-local"] = &data.SffLocal
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data) GetParentYangName() string { return "service-index" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp
// SFP stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi

    // Terminate counters.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetFilter() yfilter.YFilter { return sfp.YFilter }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) SetFilter(yf yfilter.YFilter) { sfp.YFilter = yf }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetGoName(yname string) string {
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetSegmentPath() string {
    return "sfp"
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &sfp.SpiSi
    }
    if childYangName == "term" {
        return &sfp.Term
    }
    return nil
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &sfp.SpiSi
    children["term"] = &sfp.Term
    return children
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetBundleName() string { return "cisco_ios_xr" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetYangName() string { return "sfp" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) SetParent(parent types.Entity) { sfp.parent = parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetParent() types.Entity { return sfp.parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi
// Service index counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_SpiSi) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term
// Terminate counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sfp_Term) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi
// SPI SI stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf
// Service function stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetFilter() yfilter.YFilter { return sf.YFilter }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) SetFilter(yf yfilter.YFilter) { sf.YFilter = yf }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetSegmentPath() string {
    return "sf"
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sf.ProcessedPkts
    leafs["processed-bytes"] = sf.ProcessedBytes
    return leafs
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetBundleName() string { return "cisco_ios_xr" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetYangName() string { return "sf" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) SetParent(parent types.Entity) { sf.parent = parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetParent() types.Entity { return sf.parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sf) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff
// Service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetFilter() yfilter.YFilter { return sff.YFilter }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) SetFilter(yf yfilter.YFilter) { sff.YFilter = yf }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetSegmentPath() string {
    return "sff"
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sff.ProcessedPkts
    leafs["processed-bytes"] = sff.ProcessedBytes
    return leafs
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetBundleName() string { return "cisco_ios_xr" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetYangName() string { return "sff" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) SetParent(parent types.Entity) { sff.parent = parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetParent() types.Entity { return sff.parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_Sff) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal
// Local service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetFilter() yfilter.YFilter { return sffLocal.YFilter }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) SetFilter(yf yfilter.YFilter) { sffLocal.YFilter = yf }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetGoName(yname string) string {
    if yname == "malformed-err-pkts" { return "MalformedErrPkts" }
    if yname == "lookup-err-pkts" { return "LookupErrPkts" }
    if yname == "malformed-err-bytes" { return "MalformedErrBytes" }
    if yname == "lookup-err-bytes" { return "LookupErrBytes" }
    return ""
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetSegmentPath() string {
    return "sff-local"
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-err-pkts"] = sffLocal.MalformedErrPkts
    leafs["lookup-err-pkts"] = sffLocal.LookupErrPkts
    leafs["malformed-err-bytes"] = sffLocal.MalformedErrBytes
    leafs["lookup-err-bytes"] = sffLocal.LookupErrBytes
    return leafs
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetBundleName() string { return "cisco_ios_xr" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetYangName() string { return "sff-local" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) SetParent(parent types.Entity) { sffLocal.parent = parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetParent() types.Entity { return sffLocal.parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_Data_SffLocal) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr
// SI array in case of detail stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetFilter() yfilter.YFilter { return siArr.YFilter }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) SetFilter(yf yfilter.YFilter) { siArr.YFilter = yf }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetGoName(yname string) string {
    if yname == "si" { return "Si" }
    if yname == "data" { return "Data" }
    return ""
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetSegmentPath() string {
    return "si-arr"
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &siArr.Data
    }
    return nil
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &siArr.Data
    return children
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["si"] = siArr.Si
    return leafs
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetBundleName() string { return "cisco_ios_xr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetYangName() string { return "si-arr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) SetParent(parent types.Entity) { siArr.parent = parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetParent() types.Entity { return siArr.parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr) GetParentYangName() string { return "service-index" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data
// Stats counter for this index
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SF/SFF stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data) GetParentYangName() string { return "si-arr" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi
// SF/SFF stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_ServiceIndexes_ServiceIndex_SiArr_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats
// SFP Statistics
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail statistics per service index .
    Detail ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail

    // Combined statistics of all service index in service functionpath.
    Summarized ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized
}

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    if yname == "summarized" { return "Summarized" }
    return ""
}

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &stats.Detail
    }
    if childYangName == "summarized" {
        return &stats.Summarized
    }
    return nil
}

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &stats.Detail
    children["summarized"] = &stats.Summarized
    return children
}

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetYangName() string { return "stats" }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats) GetParentYangName() string { return "path-id" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail
// Detail statistics per service index 
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics data.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data

    // SI array in case of detail stats. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr.
    SiArr []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr
}

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetGoName(yname string) string {
    if yname == "data" { return "Data" }
    if yname == "si-arr" { return "SiArr" }
    return ""
}

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &detail.Data
    }
    if childYangName == "si-arr" {
        for _, c := range detail.SiArr {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr{}
        detail.SiArr = append(detail.SiArr, child)
        return &detail.SiArr[len(detail.SiArr)-1]
    }
    return nil
}

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &detail.Data
    for i := range detail.SiArr {
        children[detail.SiArr[i].GetSegmentPath()] = &detail.SiArr[i]
    }
    return children
}

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetYangName() string { return "detail" }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetParent() types.Entity { return detail.parent }

func (detail *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail) GetParentYangName() string { return "stats" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data
// Statistics data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SFP stats.
    Sfp ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp

    // SPI SI stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term

    // Service function stats.
    Sf ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf

    // Service function forwarder stats.
    Sff ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff

    // Local service function forwarder stats.
    SffLocal ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sfp" { return "Sfp" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    if yname == "sf" { return "Sf" }
    if yname == "sff" { return "Sff" }
    if yname == "sff-local" { return "SffLocal" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sfp" {
        return &data.Sfp
    }
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    if childYangName == "sf" {
        return &data.Sf
    }
    if childYangName == "sff" {
        return &data.Sff
    }
    if childYangName == "sff-local" {
        return &data.SffLocal
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sfp"] = &data.Sfp
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    children["sf"] = &data.Sf
    children["sff"] = &data.Sff
    children["sff-local"] = &data.SffLocal
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data) GetParentYangName() string { return "detail" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp
// SFP stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi

    // Terminate counters.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetFilter() yfilter.YFilter { return sfp.YFilter }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) SetFilter(yf yfilter.YFilter) { sfp.YFilter = yf }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetGoName(yname string) string {
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetSegmentPath() string {
    return "sfp"
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &sfp.SpiSi
    }
    if childYangName == "term" {
        return &sfp.Term
    }
    return nil
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &sfp.SpiSi
    children["term"] = &sfp.Term
    return children
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetBundleName() string { return "cisco_ios_xr" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetYangName() string { return "sfp" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) SetParent(parent types.Entity) { sfp.parent = parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetParent() types.Entity { return sfp.parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi
// Service index counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_SpiSi) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term
// Terminate counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sfp_Term) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi
// SPI SI stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf
// Service function stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetFilter() yfilter.YFilter { return sf.YFilter }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) SetFilter(yf yfilter.YFilter) { sf.YFilter = yf }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetSegmentPath() string {
    return "sf"
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sf.ProcessedPkts
    leafs["processed-bytes"] = sf.ProcessedBytes
    return leafs
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetBundleName() string { return "cisco_ios_xr" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetYangName() string { return "sf" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) SetParent(parent types.Entity) { sf.parent = parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetParent() types.Entity { return sf.parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sf) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff
// Service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetFilter() yfilter.YFilter { return sff.YFilter }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) SetFilter(yf yfilter.YFilter) { sff.YFilter = yf }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetSegmentPath() string {
    return "sff"
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sff.ProcessedPkts
    leafs["processed-bytes"] = sff.ProcessedBytes
    return leafs
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetBundleName() string { return "cisco_ios_xr" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetYangName() string { return "sff" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) SetParent(parent types.Entity) { sff.parent = parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetParent() types.Entity { return sff.parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_Sff) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal
// Local service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetFilter() yfilter.YFilter { return sffLocal.YFilter }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) SetFilter(yf yfilter.YFilter) { sffLocal.YFilter = yf }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetGoName(yname string) string {
    if yname == "malformed-err-pkts" { return "MalformedErrPkts" }
    if yname == "lookup-err-pkts" { return "LookupErrPkts" }
    if yname == "malformed-err-bytes" { return "MalformedErrBytes" }
    if yname == "lookup-err-bytes" { return "LookupErrBytes" }
    return ""
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetSegmentPath() string {
    return "sff-local"
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-err-pkts"] = sffLocal.MalformedErrPkts
    leafs["lookup-err-pkts"] = sffLocal.LookupErrPkts
    leafs["malformed-err-bytes"] = sffLocal.MalformedErrBytes
    leafs["lookup-err-bytes"] = sffLocal.LookupErrBytes
    return leafs
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetBundleName() string { return "cisco_ios_xr" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetYangName() string { return "sff-local" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) SetParent(parent types.Entity) { sffLocal.parent = parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetParent() types.Entity { return sffLocal.parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_Data_SffLocal) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr
// SI array in case of detail stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetFilter() yfilter.YFilter { return siArr.YFilter }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) SetFilter(yf yfilter.YFilter) { siArr.YFilter = yf }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetGoName(yname string) string {
    if yname == "si" { return "Si" }
    if yname == "data" { return "Data" }
    return ""
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetSegmentPath() string {
    return "si-arr"
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &siArr.Data
    }
    return nil
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &siArr.Data
    return children
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["si"] = siArr.Si
    return leafs
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetBundleName() string { return "cisco_ios_xr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetYangName() string { return "si-arr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) SetParent(parent types.Entity) { siArr.parent = parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetParent() types.Entity { return siArr.parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr) GetParentYangName() string { return "detail" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data
// Stats counter for this index
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SF/SFF stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data) GetParentYangName() string { return "si-arr" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi
// SF/SFF stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Detail_SiArr_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized
// Combined statistics of all service index
// in service functionpath
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics data.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data

    // SI array in case of detail stats. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr.
    SiArr []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr
}

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetFilter() yfilter.YFilter { return summarized.YFilter }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) SetFilter(yf yfilter.YFilter) { summarized.YFilter = yf }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetGoName(yname string) string {
    if yname == "data" { return "Data" }
    if yname == "si-arr" { return "SiArr" }
    return ""
}

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetSegmentPath() string {
    return "summarized"
}

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &summarized.Data
    }
    if childYangName == "si-arr" {
        for _, c := range summarized.SiArr {
            if summarized.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr{}
        summarized.SiArr = append(summarized.SiArr, child)
        return &summarized.SiArr[len(summarized.SiArr)-1]
    }
    return nil
}

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &summarized.Data
    for i := range summarized.SiArr {
        children[summarized.SiArr[i].GetSegmentPath()] = &summarized.SiArr[i]
    }
    return children
}

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetBundleName() string { return "cisco_ios_xr" }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetYangName() string { return "summarized" }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) SetParent(parent types.Entity) { summarized.parent = parent }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetParent() types.Entity { return summarized.parent }

func (summarized *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized) GetParentYangName() string { return "stats" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data
// Statistics data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SFP stats.
    Sfp ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp

    // SPI SI stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term

    // Service function stats.
    Sf ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf

    // Service function forwarder stats.
    Sff ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff

    // Local service function forwarder stats.
    SffLocal ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sfp" { return "Sfp" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    if yname == "sf" { return "Sf" }
    if yname == "sff" { return "Sff" }
    if yname == "sff-local" { return "SffLocal" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sfp" {
        return &data.Sfp
    }
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    if childYangName == "sf" {
        return &data.Sf
    }
    if childYangName == "sff" {
        return &data.Sff
    }
    if childYangName == "sff-local" {
        return &data.SffLocal
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sfp"] = &data.Sfp
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    children["sf"] = &data.Sf
    children["sff"] = &data.Sff
    children["sff-local"] = &data.SffLocal
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data) GetParentYangName() string { return "summarized" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp
// SFP stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi

    // Terminate counters.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetFilter() yfilter.YFilter { return sfp.YFilter }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) SetFilter(yf yfilter.YFilter) { sfp.YFilter = yf }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetGoName(yname string) string {
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetSegmentPath() string {
    return "sfp"
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &sfp.SpiSi
    }
    if childYangName == "term" {
        return &sfp.Term
    }
    return nil
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &sfp.SpiSi
    children["term"] = &sfp.Term
    return children
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetBundleName() string { return "cisco_ios_xr" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetYangName() string { return "sfp" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) SetParent(parent types.Entity) { sfp.parent = parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetParent() types.Entity { return sfp.parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi
// Service index counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_SpiSi) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term
// Terminate counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sfp_Term) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi
// SPI SI stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf
// Service function stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetFilter() yfilter.YFilter { return sf.YFilter }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) SetFilter(yf yfilter.YFilter) { sf.YFilter = yf }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetSegmentPath() string {
    return "sf"
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sf.ProcessedPkts
    leafs["processed-bytes"] = sf.ProcessedBytes
    return leafs
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetBundleName() string { return "cisco_ios_xr" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetYangName() string { return "sf" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) SetParent(parent types.Entity) { sf.parent = parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetParent() types.Entity { return sf.parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sf) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff
// Service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetFilter() yfilter.YFilter { return sff.YFilter }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) SetFilter(yf yfilter.YFilter) { sff.YFilter = yf }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetSegmentPath() string {
    return "sff"
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sff.ProcessedPkts
    leafs["processed-bytes"] = sff.ProcessedBytes
    return leafs
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetBundleName() string { return "cisco_ios_xr" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetYangName() string { return "sff" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) SetParent(parent types.Entity) { sff.parent = parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetParent() types.Entity { return sff.parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_Sff) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal
// Local service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetFilter() yfilter.YFilter { return sffLocal.YFilter }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) SetFilter(yf yfilter.YFilter) { sffLocal.YFilter = yf }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetGoName(yname string) string {
    if yname == "malformed-err-pkts" { return "MalformedErrPkts" }
    if yname == "lookup-err-pkts" { return "LookupErrPkts" }
    if yname == "malformed-err-bytes" { return "MalformedErrBytes" }
    if yname == "lookup-err-bytes" { return "LookupErrBytes" }
    return ""
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetSegmentPath() string {
    return "sff-local"
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-err-pkts"] = sffLocal.MalformedErrPkts
    leafs["lookup-err-pkts"] = sffLocal.LookupErrPkts
    leafs["malformed-err-bytes"] = sffLocal.MalformedErrBytes
    leafs["lookup-err-bytes"] = sffLocal.LookupErrBytes
    return leafs
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetBundleName() string { return "cisco_ios_xr" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetYangName() string { return "sff-local" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) SetParent(parent types.Entity) { sffLocal.parent = parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetParent() types.Entity { return sffLocal.parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_Data_SffLocal) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr
// SI array in case of detail stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetFilter() yfilter.YFilter { return siArr.YFilter }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) SetFilter(yf yfilter.YFilter) { siArr.YFilter = yf }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetGoName(yname string) string {
    if yname == "si" { return "Si" }
    if yname == "data" { return "Data" }
    return ""
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetSegmentPath() string {
    return "si-arr"
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &siArr.Data
    }
    return nil
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &siArr.Data
    return children
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["si"] = siArr.Si
    return leafs
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetBundleName() string { return "cisco_ios_xr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetYangName() string { return "si-arr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) SetParent(parent types.Entity) { siArr.parent = parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetParent() types.Entity { return siArr.parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr) GetParentYangName() string { return "summarized" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data
// Stats counter for this index
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SF/SFF stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data) GetParentYangName() string { return "si-arr" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi
// SF/SFF stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionPath_PathIds_PathId_Stats_Summarized_SiArr_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction
// Service Function operational data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of Service Function Names.
    SfNames ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames
}

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetFilter() yfilter.YFilter { return serviceFunction.YFilter }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) SetFilter(yf yfilter.YFilter) { serviceFunction.YFilter = yf }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetGoName(yname string) string {
    if yname == "sf-names" { return "SfNames" }
    return ""
}

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetSegmentPath() string {
    return "service-function"
}

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sf-names" {
        return &serviceFunction.SfNames
    }
    return nil
}

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sf-names"] = &serviceFunction.SfNames
    return children
}

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetYangName() string { return "service-function" }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) SetParent(parent types.Entity) { serviceFunction.parent = parent }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetParent() types.Entity { return serviceFunction.parent }

func (serviceFunction *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction) GetParentYangName() string { return "process" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames
// List of Service Function Names
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of Service Function. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName.
    SfName []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName
}

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetFilter() yfilter.YFilter { return sfNames.YFilter }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) SetFilter(yf yfilter.YFilter) { sfNames.YFilter = yf }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetGoName(yname string) string {
    if yname == "sf-name" { return "SfName" }
    return ""
}

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetSegmentPath() string {
    return "sf-names"
}

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sf-name" {
        for _, c := range sfNames.SfName {
            if sfNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName{}
        sfNames.SfName = append(sfNames.SfName, child)
        return &sfNames.SfName[len(sfNames.SfName)-1]
    }
    return nil
}

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sfNames.SfName {
        children[sfNames.SfName[i].GetSegmentPath()] = &sfNames.SfName[i]
    }
    return children
}

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetBundleName() string { return "cisco_ios_xr" }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetYangName() string { return "sf-names" }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) SetParent(parent types.Entity) { sfNames.parent = parent }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetParent() types.Entity { return sfNames.parent }

func (sfNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames) GetParentYangName() string { return "service-function" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName
// Name of Service Function
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name. The type is string with length: 1..32.
    Name interface{}

    // Statistics data.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data

    // SI array in case of detail stats. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr.
    SiArr []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr
}

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetFilter() yfilter.YFilter { return sfName.YFilter }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) SetFilter(yf yfilter.YFilter) { sfName.YFilter = yf }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "data" { return "Data" }
    if yname == "si-arr" { return "SiArr" }
    return ""
}

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetSegmentPath() string {
    return "sf-name" + "[name='" + fmt.Sprintf("%v", sfName.Name) + "']"
}

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &sfName.Data
    }
    if childYangName == "si-arr" {
        for _, c := range sfName.SiArr {
            if sfName.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr{}
        sfName.SiArr = append(sfName.SiArr, child)
        return &sfName.SiArr[len(sfName.SiArr)-1]
    }
    return nil
}

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &sfName.Data
    for i := range sfName.SiArr {
        children[sfName.SiArr[i].GetSegmentPath()] = &sfName.SiArr[i]
    }
    return children
}

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sfName.Name
    return leafs
}

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetBundleName() string { return "cisco_ios_xr" }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetYangName() string { return "sf-name" }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) SetParent(parent types.Entity) { sfName.parent = parent }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetParent() types.Entity { return sfName.parent }

func (sfName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName) GetParentYangName() string { return "sf-names" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data
// Statistics data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SFP stats.
    Sfp ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp

    // SPI SI stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term

    // Service function stats.
    Sf ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf

    // Service function forwarder stats.
    Sff ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff

    // Local service function forwarder stats.
    SffLocal ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sfp" { return "Sfp" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    if yname == "sf" { return "Sf" }
    if yname == "sff" { return "Sff" }
    if yname == "sff-local" { return "SffLocal" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sfp" {
        return &data.Sfp
    }
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    if childYangName == "sf" {
        return &data.Sf
    }
    if childYangName == "sff" {
        return &data.Sff
    }
    if childYangName == "sff-local" {
        return &data.SffLocal
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sfp"] = &data.Sfp
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    children["sf"] = &data.Sf
    children["sff"] = &data.Sff
    children["sff-local"] = &data.SffLocal
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data) GetParentYangName() string { return "sf-name" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp
// SFP stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi

    // Terminate counters.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetFilter() yfilter.YFilter { return sfp.YFilter }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) SetFilter(yf yfilter.YFilter) { sfp.YFilter = yf }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetGoName(yname string) string {
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetSegmentPath() string {
    return "sfp"
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &sfp.SpiSi
    }
    if childYangName == "term" {
        return &sfp.Term
    }
    return nil
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &sfp.SpiSi
    children["term"] = &sfp.Term
    return children
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetBundleName() string { return "cisco_ios_xr" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetYangName() string { return "sfp" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) SetParent(parent types.Entity) { sfp.parent = parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetParent() types.Entity { return sfp.parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi
// Service index counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_SpiSi) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term
// Terminate counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sfp_Term) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi
// SPI SI stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf
// Service function stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetFilter() yfilter.YFilter { return sf.YFilter }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) SetFilter(yf yfilter.YFilter) { sf.YFilter = yf }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetSegmentPath() string {
    return "sf"
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sf.ProcessedPkts
    leafs["processed-bytes"] = sf.ProcessedBytes
    return leafs
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetBundleName() string { return "cisco_ios_xr" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetYangName() string { return "sf" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) SetParent(parent types.Entity) { sf.parent = parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetParent() types.Entity { return sf.parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sf) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff
// Service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetFilter() yfilter.YFilter { return sff.YFilter }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) SetFilter(yf yfilter.YFilter) { sff.YFilter = yf }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetSegmentPath() string {
    return "sff"
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sff.ProcessedPkts
    leafs["processed-bytes"] = sff.ProcessedBytes
    return leafs
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetBundleName() string { return "cisco_ios_xr" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetYangName() string { return "sff" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) SetParent(parent types.Entity) { sff.parent = parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetParent() types.Entity { return sff.parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_Sff) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal
// Local service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetFilter() yfilter.YFilter { return sffLocal.YFilter }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) SetFilter(yf yfilter.YFilter) { sffLocal.YFilter = yf }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetGoName(yname string) string {
    if yname == "malformed-err-pkts" { return "MalformedErrPkts" }
    if yname == "lookup-err-pkts" { return "LookupErrPkts" }
    if yname == "malformed-err-bytes" { return "MalformedErrBytes" }
    if yname == "lookup-err-bytes" { return "LookupErrBytes" }
    return ""
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetSegmentPath() string {
    return "sff-local"
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-err-pkts"] = sffLocal.MalformedErrPkts
    leafs["lookup-err-pkts"] = sffLocal.LookupErrPkts
    leafs["malformed-err-bytes"] = sffLocal.MalformedErrBytes
    leafs["lookup-err-bytes"] = sffLocal.LookupErrBytes
    return leafs
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetBundleName() string { return "cisco_ios_xr" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetYangName() string { return "sff-local" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) SetParent(parent types.Entity) { sffLocal.parent = parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetParent() types.Entity { return sffLocal.parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_Data_SffLocal) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr
// SI array in case of detail stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetFilter() yfilter.YFilter { return siArr.YFilter }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) SetFilter(yf yfilter.YFilter) { siArr.YFilter = yf }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetGoName(yname string) string {
    if yname == "si" { return "Si" }
    if yname == "data" { return "Data" }
    return ""
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetSegmentPath() string {
    return "si-arr"
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &siArr.Data
    }
    return nil
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &siArr.Data
    return children
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["si"] = siArr.Si
    return leafs
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetBundleName() string { return "cisco_ios_xr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetYangName() string { return "si-arr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) SetParent(parent types.Entity) { siArr.parent = parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetParent() types.Entity { return siArr.parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr) GetParentYangName() string { return "sf-name" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data
// Stats counter for this index
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SF/SFF stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data) GetParentYangName() string { return "si-arr" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi
// SF/SFF stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunction_SfNames_SfName_SiArr_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder
// Service Function Forwarder operational data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local Service Function Forwarder operational data.
    Local ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local

    // List of Service Function Forwarder Names.
    SffNames ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames
}

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetFilter() yfilter.YFilter { return serviceFunctionForwarder.YFilter }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) SetFilter(yf yfilter.YFilter) { serviceFunctionForwarder.YFilter = yf }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetGoName(yname string) string {
    if yname == "local" { return "Local" }
    if yname == "sff-names" { return "SffNames" }
    return ""
}

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetSegmentPath() string {
    return "service-function-forwarder"
}

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local" {
        return &serviceFunctionForwarder.Local
    }
    if childYangName == "sff-names" {
        return &serviceFunctionForwarder.SffNames
    }
    return nil
}

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local"] = &serviceFunctionForwarder.Local
    children["sff-names"] = &serviceFunctionForwarder.SffNames
    return children
}

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetYangName() string { return "service-function-forwarder" }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) SetParent(parent types.Entity) { serviceFunctionForwarder.parent = parent }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetParent() types.Entity { return serviceFunctionForwarder.parent }

func (serviceFunctionForwarder *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder) GetParentYangName() string { return "process" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local
// Local Service Function Forwarder operational
// data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error Statistics for local service function forwarder.
    Error ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error
}

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetFilter() yfilter.YFilter { return local.YFilter }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) SetFilter(yf yfilter.YFilter) { local.YFilter = yf }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    return ""
}

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetSegmentPath() string {
    return "local"
}

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        return &local.Error
    }
    return nil
}

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["error"] = &local.Error
    return children
}

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetBundleName() string { return "cisco_ios_xr" }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetYangName() string { return "local" }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) SetParent(parent types.Entity) { local.parent = parent }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetParent() types.Entity { return local.parent }

func (local *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local) GetParentYangName() string { return "service-function-forwarder" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error
// Error Statistics for local service function
// forwarder
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics data.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data

    // SI array in case of detail stats. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr.
    SiArr []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr
}

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetGoName(yname string) string {
    if yname == "data" { return "Data" }
    if yname == "si-arr" { return "SiArr" }
    return ""
}

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetSegmentPath() string {
    return "error"
}

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &error.Data
    }
    if childYangName == "si-arr" {
        for _, c := range error.SiArr {
            if error.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr{}
        error.SiArr = append(error.SiArr, child)
        return &error.SiArr[len(error.SiArr)-1]
    }
    return nil
}

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &error.Data
    for i := range error.SiArr {
        children[error.SiArr[i].GetSegmentPath()] = &error.SiArr[i]
    }
    return children
}

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetYangName() string { return "error" }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetParent() types.Entity { return error.parent }

func (error *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error) GetParentYangName() string { return "local" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data
// Statistics data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SFP stats.
    Sfp ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp

    // SPI SI stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term

    // Service function stats.
    Sf ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf

    // Service function forwarder stats.
    Sff ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff

    // Local service function forwarder stats.
    SffLocal ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sfp" { return "Sfp" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    if yname == "sf" { return "Sf" }
    if yname == "sff" { return "Sff" }
    if yname == "sff-local" { return "SffLocal" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sfp" {
        return &data.Sfp
    }
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    if childYangName == "sf" {
        return &data.Sf
    }
    if childYangName == "sff" {
        return &data.Sff
    }
    if childYangName == "sff-local" {
        return &data.SffLocal
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sfp"] = &data.Sfp
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    children["sf"] = &data.Sf
    children["sff"] = &data.Sff
    children["sff-local"] = &data.SffLocal
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data) GetParentYangName() string { return "error" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp
// SFP stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi

    // Terminate counters.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetFilter() yfilter.YFilter { return sfp.YFilter }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) SetFilter(yf yfilter.YFilter) { sfp.YFilter = yf }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetGoName(yname string) string {
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetSegmentPath() string {
    return "sfp"
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &sfp.SpiSi
    }
    if childYangName == "term" {
        return &sfp.Term
    }
    return nil
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &sfp.SpiSi
    children["term"] = &sfp.Term
    return children
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetBundleName() string { return "cisco_ios_xr" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetYangName() string { return "sfp" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) SetParent(parent types.Entity) { sfp.parent = parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetParent() types.Entity { return sfp.parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi
// Service index counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_SpiSi) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term
// Terminate counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sfp_Term) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi
// SPI SI stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf
// Service function stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetFilter() yfilter.YFilter { return sf.YFilter }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) SetFilter(yf yfilter.YFilter) { sf.YFilter = yf }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetSegmentPath() string {
    return "sf"
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sf.ProcessedPkts
    leafs["processed-bytes"] = sf.ProcessedBytes
    return leafs
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetBundleName() string { return "cisco_ios_xr" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetYangName() string { return "sf" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) SetParent(parent types.Entity) { sf.parent = parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetParent() types.Entity { return sf.parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sf) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff
// Service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetFilter() yfilter.YFilter { return sff.YFilter }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) SetFilter(yf yfilter.YFilter) { sff.YFilter = yf }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetSegmentPath() string {
    return "sff"
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sff.ProcessedPkts
    leafs["processed-bytes"] = sff.ProcessedBytes
    return leafs
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetBundleName() string { return "cisco_ios_xr" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetYangName() string { return "sff" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) SetParent(parent types.Entity) { sff.parent = parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetParent() types.Entity { return sff.parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_Sff) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal
// Local service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetFilter() yfilter.YFilter { return sffLocal.YFilter }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) SetFilter(yf yfilter.YFilter) { sffLocal.YFilter = yf }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetGoName(yname string) string {
    if yname == "malformed-err-pkts" { return "MalformedErrPkts" }
    if yname == "lookup-err-pkts" { return "LookupErrPkts" }
    if yname == "malformed-err-bytes" { return "MalformedErrBytes" }
    if yname == "lookup-err-bytes" { return "LookupErrBytes" }
    return ""
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetSegmentPath() string {
    return "sff-local"
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-err-pkts"] = sffLocal.MalformedErrPkts
    leafs["lookup-err-pkts"] = sffLocal.LookupErrPkts
    leafs["malformed-err-bytes"] = sffLocal.MalformedErrBytes
    leafs["lookup-err-bytes"] = sffLocal.LookupErrBytes
    return leafs
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetBundleName() string { return "cisco_ios_xr" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetYangName() string { return "sff-local" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) SetParent(parent types.Entity) { sffLocal.parent = parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetParent() types.Entity { return sffLocal.parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_Data_SffLocal) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr
// SI array in case of detail stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetFilter() yfilter.YFilter { return siArr.YFilter }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) SetFilter(yf yfilter.YFilter) { siArr.YFilter = yf }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetGoName(yname string) string {
    if yname == "si" { return "Si" }
    if yname == "data" { return "Data" }
    return ""
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetSegmentPath() string {
    return "si-arr"
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &siArr.Data
    }
    return nil
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &siArr.Data
    return children
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["si"] = siArr.Si
    return leafs
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetBundleName() string { return "cisco_ios_xr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetYangName() string { return "si-arr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) SetParent(parent types.Entity) { siArr.parent = parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetParent() types.Entity { return siArr.parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr) GetParentYangName() string { return "error" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data
// Stats counter for this index
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SF/SFF stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data) GetParentYangName() string { return "si-arr" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi
// SF/SFF stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_Local_Error_SiArr_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames
// List of Service Function Forwarder Names
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of Service Function Forwarder. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName.
    SffName []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName
}

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetFilter() yfilter.YFilter { return sffNames.YFilter }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) SetFilter(yf yfilter.YFilter) { sffNames.YFilter = yf }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetGoName(yname string) string {
    if yname == "sff-name" { return "SffName" }
    return ""
}

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetSegmentPath() string {
    return "sff-names"
}

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sff-name" {
        for _, c := range sffNames.SffName {
            if sffNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName{}
        sffNames.SffName = append(sffNames.SffName, child)
        return &sffNames.SffName[len(sffNames.SffName)-1]
    }
    return nil
}

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sffNames.SffName {
        children[sffNames.SffName[i].GetSegmentPath()] = &sffNames.SffName[i]
    }
    return children
}

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetBundleName() string { return "cisco_ios_xr" }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetYangName() string { return "sff-names" }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) SetParent(parent types.Entity) { sffNames.parent = parent }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetParent() types.Entity { return sffNames.parent }

func (sffNames *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames) GetParentYangName() string { return "service-function-forwarder" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName
// Name of Service Function Forwarder
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name. The type is string with length: 1..32.
    Name interface{}

    // Statistics data.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data

    // SI array in case of detail stats. The type is slice of
    // ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr.
    SiArr []ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr
}

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetFilter() yfilter.YFilter { return sffName.YFilter }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) SetFilter(yf yfilter.YFilter) { sffName.YFilter = yf }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "data" { return "Data" }
    if yname == "si-arr" { return "SiArr" }
    return ""
}

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetSegmentPath() string {
    return "sff-name" + "[name='" + fmt.Sprintf("%v", sffName.Name) + "']"
}

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &sffName.Data
    }
    if childYangName == "si-arr" {
        for _, c := range sffName.SiArr {
            if sffName.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr{}
        sffName.SiArr = append(sffName.SiArr, child)
        return &sffName.SiArr[len(sffName.SiArr)-1]
    }
    return nil
}

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &sffName.Data
    for i := range sffName.SiArr {
        children[sffName.SiArr[i].GetSegmentPath()] = &sffName.SiArr[i]
    }
    return children
}

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = sffName.Name
    return leafs
}

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetBundleName() string { return "cisco_ios_xr" }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetYangName() string { return "sff-name" }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) SetParent(parent types.Entity) { sffName.parent = parent }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetParent() types.Entity { return sffName.parent }

func (sffName *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName) GetParentYangName() string { return "sff-names" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data
// Statistics data
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SFP stats.
    Sfp ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp

    // SPI SI stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term

    // Service function stats.
    Sf ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf

    // Service function forwarder stats.
    Sff ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff

    // Local service function forwarder stats.
    SffLocal ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sfp" { return "Sfp" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    if yname == "sf" { return "Sf" }
    if yname == "sff" { return "Sff" }
    if yname == "sff-local" { return "SffLocal" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sfp" {
        return &data.Sfp
    }
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    if childYangName == "sf" {
        return &data.Sf
    }
    if childYangName == "sff" {
        return &data.Sff
    }
    if childYangName == "sff-local" {
        return &data.SffLocal
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sfp"] = &data.Sfp
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    children["sf"] = &data.Sf
    children["sff"] = &data.Sff
    children["sff-local"] = &data.SffLocal
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data) GetParentYangName() string { return "sff-name" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp
// SFP stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index counters.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi

    // Terminate counters.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetFilter() yfilter.YFilter { return sfp.YFilter }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) SetFilter(yf yfilter.YFilter) { sfp.YFilter = yf }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetGoName(yname string) string {
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetSegmentPath() string {
    return "sfp"
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &sfp.SpiSi
    }
    if childYangName == "term" {
        return &sfp.Term
    }
    return nil
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &sfp.SpiSi
    children["term"] = &sfp.Term
    return children
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetBundleName() string { return "cisco_ios_xr" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetYangName() string { return "sfp" }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) SetParent(parent types.Entity) { sfp.parent = parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetParent() types.Entity { return sfp.parent }

func (sfp *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi
// Service index counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_SpiSi) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term
// Terminate counters
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sfp_Term) GetParentYangName() string { return "sfp" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi
// SPI SI stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Term) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf
// Service function stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetFilter() yfilter.YFilter { return sf.YFilter }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) SetFilter(yf yfilter.YFilter) { sf.YFilter = yf }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetSegmentPath() string {
    return "sf"
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sf.ProcessedPkts
    leafs["processed-bytes"] = sf.ProcessedBytes
    return leafs
}

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetBundleName() string { return "cisco_ios_xr" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetYangName() string { return "sf" }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) SetParent(parent types.Entity) { sf.parent = parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetParent() types.Entity { return sf.parent }

func (sf *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sf) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff
// Service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetFilter() yfilter.YFilter { return sff.YFilter }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) SetFilter(yf yfilter.YFilter) { sff.YFilter = yf }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetSegmentPath() string {
    return "sff"
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = sff.ProcessedPkts
    leafs["processed-bytes"] = sff.ProcessedBytes
    return leafs
}

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetBundleName() string { return "cisco_ios_xr" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetYangName() string { return "sff" }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) SetParent(parent types.Entity) { sff.parent = parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetParent() types.Entity { return sff.parent }

func (sff *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_Sff) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal
// Local service function forwarder stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets with invalid NSH header. The type is interface{} with
    // range: 0..18446744073709551615.
    MalformedErrPkts interface{}

    // Number of packets with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615.
    LookupErrPkts interface{}

    // Total bytes with invalid NSH header. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    MalformedErrBytes interface{}

    // Total bytes with unknown spi-si. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    LookupErrBytes interface{}
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetFilter() yfilter.YFilter { return sffLocal.YFilter }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) SetFilter(yf yfilter.YFilter) { sffLocal.YFilter = yf }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetGoName(yname string) string {
    if yname == "malformed-err-pkts" { return "MalformedErrPkts" }
    if yname == "lookup-err-pkts" { return "LookupErrPkts" }
    if yname == "malformed-err-bytes" { return "MalformedErrBytes" }
    if yname == "lookup-err-bytes" { return "LookupErrBytes" }
    return ""
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetSegmentPath() string {
    return "sff-local"
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["malformed-err-pkts"] = sffLocal.MalformedErrPkts
    leafs["lookup-err-pkts"] = sffLocal.LookupErrPkts
    leafs["malformed-err-bytes"] = sffLocal.MalformedErrBytes
    leafs["lookup-err-bytes"] = sffLocal.LookupErrBytes
    return leafs
}

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetBundleName() string { return "cisco_ios_xr" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetYangName() string { return "sff-local" }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) SetParent(parent types.Entity) { sffLocal.parent = parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetParent() types.Entity { return sffLocal.parent }

func (sffLocal *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_Data_SffLocal) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr
// SI array in case of detail stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service index. The type is interface{} with range: 0..255.
    Si interface{}

    // Stats counter for this index.
    Data ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetFilter() yfilter.YFilter { return siArr.YFilter }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) SetFilter(yf yfilter.YFilter) { siArr.YFilter = yf }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetGoName(yname string) string {
    if yname == "si" { return "Si" }
    if yname == "data" { return "Data" }
    return ""
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetSegmentPath() string {
    return "si-arr"
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data" {
        return &siArr.Data
    }
    return nil
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data"] = &siArr.Data
    return children
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["si"] = siArr.Si
    return leafs
}

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetBundleName() string { return "cisco_ios_xr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetYangName() string { return "si-arr" }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) SetParent(parent types.Entity) { siArr.parent = parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetParent() types.Entity { return siArr.parent }

func (siArr *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr) GetParentYangName() string { return "sff-name" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data
// Stats counter for this index
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // type. The type is VsNshStats.
    Type interface{}

    // SF/SFF stats.
    SpiSi ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi

    // Terminate stats.
    Term ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetFilter() yfilter.YFilter { return data.YFilter }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) SetFilter(yf yfilter.YFilter) { data.YFilter = yf }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "spi-si" { return "SpiSi" }
    if yname == "term" { return "Term" }
    return ""
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetSegmentPath() string {
    return "data"
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spi-si" {
        return &data.SpiSi
    }
    if childYangName == "term" {
        return &data.Term
    }
    return nil
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["spi-si"] = &data.SpiSi
    children["term"] = &data.Term
    return children
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = data.Type
    return leafs
}

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetBundleName() string { return "cisco_ios_xr" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetYangName() string { return "data" }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) SetParent(parent types.Entity) { data.parent = parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetParent() types.Entity { return data.parent }

func (data *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data) GetParentYangName() string { return "si-arr" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi
// SF/SFF stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets processed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProcessedPkts interface{}

    // Total bytes processed. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ProcessedBytes interface{}
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetFilter() yfilter.YFilter { return spiSi.YFilter }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) SetFilter(yf yfilter.YFilter) { spiSi.YFilter = yf }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetGoName(yname string) string {
    if yname == "processed-pkts" { return "ProcessedPkts" }
    if yname == "processed-bytes" { return "ProcessedBytes" }
    return ""
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetSegmentPath() string {
    return "spi-si"
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processed-pkts"] = spiSi.ProcessedPkts
    leafs["processed-bytes"] = spiSi.ProcessedBytes
    return leafs
}

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetBundleName() string { return "cisco_ios_xr" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetYangName() string { return "spi-si" }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) SetParent(parent types.Entity) { spiSi.parent = parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetParent() types.Entity { return spiSi.parent }

func (spiSi *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_SpiSi) GetParentYangName() string { return "data" }

// ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term
// Terminate stats
type ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of terminated packets. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedPkts interface{}

    // Total bytes terminated. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TerminatedBytes interface{}
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetFilter() yfilter.YFilter { return term.YFilter }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) SetFilter(yf yfilter.YFilter) { term.YFilter = yf }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetGoName(yname string) string {
    if yname == "terminated-pkts" { return "TerminatedPkts" }
    if yname == "terminated-bytes" { return "TerminatedBytes" }
    return ""
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetSegmentPath() string {
    return "term"
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["terminated-pkts"] = term.TerminatedPkts
    leafs["terminated-bytes"] = term.TerminatedBytes
    return leafs
}

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetBundleName() string { return "cisco_ios_xr" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetYangName() string { return "term" }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) SetParent(parent types.Entity) { term.parent = parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetParent() types.Entity { return term.parent }

func (term *ServiceFunctionChaining_Nodes_Node_Process_ServiceFunctionForwarder_SffNames_SffName_SiArr_Data_Term) GetParentYangName() string { return "data" }

