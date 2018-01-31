// This module contains a collection of YANG definitions
// for Cisco IOS-XR fia-internal-tcam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   controller: Controller Resources
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fia_internal_tcam_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fia_internal_tcam_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fia-internal-tcam-oper controller}", reflect.TypeOf(Controller{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fia-internal-tcam-oper:controller", reflect.TypeOf(Controller{}))
}

// Controller
// Controller Resources
type Controller struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Controller DPA operational data.
    Dpa Controller_Dpa
}

func (controller *Controller) GetFilter() yfilter.YFilter { return controller.YFilter }

func (controller *Controller) SetFilter(yf yfilter.YFilter) { controller.YFilter = yf }

func (controller *Controller) GetGoName(yname string) string {
    if yname == "dpa" { return "Dpa" }
    return ""
}

func (controller *Controller) GetSegmentPath() string {
    return "Cisco-IOS-XR-fia-internal-tcam-oper:controller"
}

func (controller *Controller) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dpa" {
        return &controller.Dpa
    }
    return nil
}

func (controller *Controller) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dpa"] = &controller.Dpa
    return children
}

func (controller *Controller) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controller *Controller) GetBundleName() string { return "cisco_ios_xr" }

func (controller *Controller) GetYangName() string { return "controller" }

func (controller *Controller) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controller *Controller) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controller *Controller) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controller *Controller) SetParent(parent types.Entity) { controller.parent = parent }

func (controller *Controller) GetParent() types.Entity { return controller.parent }

func (controller *Controller) GetParentYangName() string { return "Cisco-IOS-XR-fia-internal-tcam-oper" }

// Controller_Dpa
// Controller DPA operational data
type Controller_Dpa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DPA data for available nodes.
    Nodes Controller_Dpa_Nodes
}

func (dpa *Controller_Dpa) GetFilter() yfilter.YFilter { return dpa.YFilter }

func (dpa *Controller_Dpa) SetFilter(yf yfilter.YFilter) { dpa.YFilter = yf }

func (dpa *Controller_Dpa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (dpa *Controller_Dpa) GetSegmentPath() string {
    return "dpa"
}

func (dpa *Controller_Dpa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &dpa.Nodes
    }
    return nil
}

func (dpa *Controller_Dpa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &dpa.Nodes
    return children
}

func (dpa *Controller_Dpa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dpa *Controller_Dpa) GetBundleName() string { return "cisco_ios_xr" }

func (dpa *Controller_Dpa) GetYangName() string { return "dpa" }

func (dpa *Controller_Dpa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dpa *Controller_Dpa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dpa *Controller_Dpa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dpa *Controller_Dpa) SetParent(parent types.Entity) { dpa.parent = parent }

func (dpa *Controller_Dpa) GetParent() types.Entity { return dpa.parent }

func (dpa *Controller_Dpa) GetParentYangName() string { return "controller" }

// Controller_Dpa_Nodes
// DPA data for available nodes
type Controller_Dpa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Controller_Dpa_Nodes_Node.
    Node []Controller_Dpa_Nodes_Node
}

func (nodes *Controller_Dpa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Controller_Dpa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Controller_Dpa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Controller_Dpa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Controller_Dpa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Controller_Dpa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Controller_Dpa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Controller_Dpa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Controller_Dpa_Nodes) GetYangName() string { return "nodes" }

func (nodes *Controller_Dpa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Controller_Dpa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Controller_Dpa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Controller_Dpa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Controller_Dpa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Controller_Dpa_Nodes) GetParentYangName() string { return "dpa" }

// Controller_Dpa_Nodes_Node
// DPA operational data for a particular node
type Controller_Dpa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // External TCAM Resource Information.
    ExternalTcamResources Controller_Dpa_Nodes_Node_ExternalTcamResources

    // Internal TCAM Resource Information.
    InternalTcamResources Controller_Dpa_Nodes_Node_InternalTcamResources
}

func (node *Controller_Dpa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Controller_Dpa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Controller_Dpa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "external-tcam-resources" { return "ExternalTcamResources" }
    if yname == "internal-tcam-resources" { return "InternalTcamResources" }
    return ""
}

func (node *Controller_Dpa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Controller_Dpa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "external-tcam-resources" {
        return &node.ExternalTcamResources
    }
    if childYangName == "internal-tcam-resources" {
        return &node.InternalTcamResources
    }
    return nil
}

func (node *Controller_Dpa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["external-tcam-resources"] = &node.ExternalTcamResources
    children["internal-tcam-resources"] = &node.InternalTcamResources
    return children
}

func (node *Controller_Dpa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Controller_Dpa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Controller_Dpa_Nodes_Node) GetYangName() string { return "node" }

func (node *Controller_Dpa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Controller_Dpa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Controller_Dpa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Controller_Dpa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Controller_Dpa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Controller_Dpa_Nodes_Node) GetParentYangName() string { return "nodes" }

// Controller_Dpa_Nodes_Node_ExternalTcamResources
// External TCAM Resource Information
type Controller_Dpa_Nodes_Node_ExternalTcamResources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // npu tcam. The type is slice of
    // Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam.
    NpuTcam []Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetFilter() yfilter.YFilter { return externalTcamResources.YFilter }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) SetFilter(yf yfilter.YFilter) { externalTcamResources.YFilter = yf }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetGoName(yname string) string {
    if yname == "npu-tcam" { return "NpuTcam" }
    return ""
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetSegmentPath() string {
    return "external-tcam-resources"
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-tcam" {
        for _, c := range externalTcamResources.NpuTcam {
            if externalTcamResources.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam{}
        externalTcamResources.NpuTcam = append(externalTcamResources.NpuTcam, child)
        return &externalTcamResources.NpuTcam[len(externalTcamResources.NpuTcam)-1]
    }
    return nil
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range externalTcamResources.NpuTcam {
        children[externalTcamResources.NpuTcam[i].GetSegmentPath()] = &externalTcamResources.NpuTcam[i]
    }
    return children
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetBundleName() string { return "cisco_ios_xr" }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetYangName() string { return "external-tcam-resources" }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) SetParent(parent types.Entity) { externalTcamResources.parent = parent }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetParent() types.Entity { return externalTcamResources.parent }

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetParentYangName() string { return "node" }

// Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam
// npu tcam
type Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // npu id. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // tcam bank. The type is slice of
    // Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank.
    TcamBank []Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetFilter() yfilter.YFilter { return npuTcam.YFilter }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) SetFilter(yf yfilter.YFilter) { npuTcam.YFilter = yf }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "tcam-bank" { return "TcamBank" }
    return ""
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetSegmentPath() string {
    return "npu-tcam"
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcam-bank" {
        for _, c := range npuTcam.TcamBank {
            if npuTcam.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank{}
        npuTcam.TcamBank = append(npuTcam.TcamBank, child)
        return &npuTcam.TcamBank[len(npuTcam.TcamBank)-1]
    }
    return nil
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuTcam.TcamBank {
        children[npuTcam.TcamBank[i].GetSegmentPath()] = &npuTcam.TcamBank[i]
    }
    return children
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = npuTcam.NpuId
    return leafs
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetBundleName() string { return "cisco_ios_xr" }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetYangName() string { return "npu-tcam" }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) SetParent(parent types.Entity) { npuTcam.parent = parent }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetParent() types.Entity { return npuTcam.parent }

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetParentYangName() string { return "external-tcam-resources" }

// Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank
// tcam bank
type Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bank id. The type is string.
    BankId interface{}

    // bank key size. The type is string.
    BankKeySize interface{}

    // bank free entries. The type is interface{} with range: 0..4294967295.
    BankFreeEntries interface{}

    // bank inuse entries. The type is interface{} with range: 0..4294967295.
    BankInuseEntries interface{}

    // owner. The type is string.
    Owner interface{}

    // nof dbs. The type is interface{} with range: 0..4294967295.
    NofDbs interface{}

    // bank db. The type is slice of
    // Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb.
    BankDb []Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetFilter() yfilter.YFilter { return tcamBank.YFilter }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) SetFilter(yf yfilter.YFilter) { tcamBank.YFilter = yf }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetGoName(yname string) string {
    if yname == "bank-id" { return "BankId" }
    if yname == "bank-key-size" { return "BankKeySize" }
    if yname == "bank-free-entries" { return "BankFreeEntries" }
    if yname == "bank-inuse-entries" { return "BankInuseEntries" }
    if yname == "owner" { return "Owner" }
    if yname == "nof-dbs" { return "NofDbs" }
    if yname == "bank-db" { return "BankDb" }
    return ""
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetSegmentPath() string {
    return "tcam-bank"
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bank-db" {
        for _, c := range tcamBank.BankDb {
            if tcamBank.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb{}
        tcamBank.BankDb = append(tcamBank.BankDb, child)
        return &tcamBank.BankDb[len(tcamBank.BankDb)-1]
    }
    return nil
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tcamBank.BankDb {
        children[tcamBank.BankDb[i].GetSegmentPath()] = &tcamBank.BankDb[i]
    }
    return children
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bank-id"] = tcamBank.BankId
    leafs["bank-key-size"] = tcamBank.BankKeySize
    leafs["bank-free-entries"] = tcamBank.BankFreeEntries
    leafs["bank-inuse-entries"] = tcamBank.BankInuseEntries
    leafs["owner"] = tcamBank.Owner
    leafs["nof-dbs"] = tcamBank.NofDbs
    return leafs
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetBundleName() string { return "cisco_ios_xr" }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetYangName() string { return "tcam-bank" }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) SetParent(parent types.Entity) { tcamBank.parent = parent }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetParent() types.Entity { return tcamBank.parent }

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetParentYangName() string { return "npu-tcam" }

// Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb
// bank db
type Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // db id. The type is interface{} with range: 0..4294967295.
    DbId interface{}

    // db inuse entries. The type is interface{} with range: 0..4294967295.
    DbInuseEntries interface{}

    // db prefix. The type is string.
    DbPrefix interface{}
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetFilter() yfilter.YFilter { return bankDb.YFilter }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) SetFilter(yf yfilter.YFilter) { bankDb.YFilter = yf }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetGoName(yname string) string {
    if yname == "db-id" { return "DbId" }
    if yname == "db-inuse-entries" { return "DbInuseEntries" }
    if yname == "db-prefix" { return "DbPrefix" }
    return ""
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetSegmentPath() string {
    return "bank-db"
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["db-id"] = bankDb.DbId
    leafs["db-inuse-entries"] = bankDb.DbInuseEntries
    leafs["db-prefix"] = bankDb.DbPrefix
    return leafs
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetBundleName() string { return "cisco_ios_xr" }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetYangName() string { return "bank-db" }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) SetParent(parent types.Entity) { bankDb.parent = parent }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetParent() types.Entity { return bankDb.parent }

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetParentYangName() string { return "tcam-bank" }

// Controller_Dpa_Nodes_Node_InternalTcamResources
// Internal TCAM Resource Information
type Controller_Dpa_Nodes_Node_InternalTcamResources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // npu tcam. The type is slice of
    // Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam.
    NpuTcam []Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetFilter() yfilter.YFilter { return internalTcamResources.YFilter }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) SetFilter(yf yfilter.YFilter) { internalTcamResources.YFilter = yf }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetGoName(yname string) string {
    if yname == "npu-tcam" { return "NpuTcam" }
    return ""
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetSegmentPath() string {
    return "internal-tcam-resources"
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-tcam" {
        for _, c := range internalTcamResources.NpuTcam {
            if internalTcamResources.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam{}
        internalTcamResources.NpuTcam = append(internalTcamResources.NpuTcam, child)
        return &internalTcamResources.NpuTcam[len(internalTcamResources.NpuTcam)-1]
    }
    return nil
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range internalTcamResources.NpuTcam {
        children[internalTcamResources.NpuTcam[i].GetSegmentPath()] = &internalTcamResources.NpuTcam[i]
    }
    return children
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetBundleName() string { return "cisco_ios_xr" }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetYangName() string { return "internal-tcam-resources" }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) SetParent(parent types.Entity) { internalTcamResources.parent = parent }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetParent() types.Entity { return internalTcamResources.parent }

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetParentYangName() string { return "node" }

// Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam
// npu tcam
type Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // npu id. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // tcam bank. The type is slice of
    // Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank.
    TcamBank []Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetFilter() yfilter.YFilter { return npuTcam.YFilter }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) SetFilter(yf yfilter.YFilter) { npuTcam.YFilter = yf }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "tcam-bank" { return "TcamBank" }
    return ""
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetSegmentPath() string {
    return "npu-tcam"
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcam-bank" {
        for _, c := range npuTcam.TcamBank {
            if npuTcam.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank{}
        npuTcam.TcamBank = append(npuTcam.TcamBank, child)
        return &npuTcam.TcamBank[len(npuTcam.TcamBank)-1]
    }
    return nil
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuTcam.TcamBank {
        children[npuTcam.TcamBank[i].GetSegmentPath()] = &npuTcam.TcamBank[i]
    }
    return children
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = npuTcam.NpuId
    return leafs
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetBundleName() string { return "cisco_ios_xr" }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetYangName() string { return "npu-tcam" }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) SetParent(parent types.Entity) { npuTcam.parent = parent }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetParent() types.Entity { return npuTcam.parent }

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetParentYangName() string { return "internal-tcam-resources" }

// Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank
// tcam bank
type Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bank id. The type is string.
    BankId interface{}

    // bank key size. The type is string.
    BankKeySize interface{}

    // bank free entries. The type is interface{} with range: 0..4294967295.
    BankFreeEntries interface{}

    // bank inuse entries. The type is interface{} with range: 0..4294967295.
    BankInuseEntries interface{}

    // owner. The type is string.
    Owner interface{}

    // nof dbs. The type is interface{} with range: 0..4294967295.
    NofDbs interface{}

    // bank db. The type is slice of
    // Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb.
    BankDb []Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetFilter() yfilter.YFilter { return tcamBank.YFilter }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) SetFilter(yf yfilter.YFilter) { tcamBank.YFilter = yf }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetGoName(yname string) string {
    if yname == "bank-id" { return "BankId" }
    if yname == "bank-key-size" { return "BankKeySize" }
    if yname == "bank-free-entries" { return "BankFreeEntries" }
    if yname == "bank-inuse-entries" { return "BankInuseEntries" }
    if yname == "owner" { return "Owner" }
    if yname == "nof-dbs" { return "NofDbs" }
    if yname == "bank-db" { return "BankDb" }
    return ""
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetSegmentPath() string {
    return "tcam-bank"
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bank-db" {
        for _, c := range tcamBank.BankDb {
            if tcamBank.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb{}
        tcamBank.BankDb = append(tcamBank.BankDb, child)
        return &tcamBank.BankDb[len(tcamBank.BankDb)-1]
    }
    return nil
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tcamBank.BankDb {
        children[tcamBank.BankDb[i].GetSegmentPath()] = &tcamBank.BankDb[i]
    }
    return children
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bank-id"] = tcamBank.BankId
    leafs["bank-key-size"] = tcamBank.BankKeySize
    leafs["bank-free-entries"] = tcamBank.BankFreeEntries
    leafs["bank-inuse-entries"] = tcamBank.BankInuseEntries
    leafs["owner"] = tcamBank.Owner
    leafs["nof-dbs"] = tcamBank.NofDbs
    return leafs
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetBundleName() string { return "cisco_ios_xr" }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetYangName() string { return "tcam-bank" }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) SetParent(parent types.Entity) { tcamBank.parent = parent }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetParent() types.Entity { return tcamBank.parent }

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetParentYangName() string { return "npu-tcam" }

// Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb
// bank db
type Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // db id. The type is interface{} with range: 0..4294967295.
    DbId interface{}

    // db inuse entries. The type is interface{} with range: 0..4294967295.
    DbInuseEntries interface{}

    // db prefix. The type is string.
    DbPrefix interface{}
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetFilter() yfilter.YFilter { return bankDb.YFilter }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) SetFilter(yf yfilter.YFilter) { bankDb.YFilter = yf }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetGoName(yname string) string {
    if yname == "db-id" { return "DbId" }
    if yname == "db-inuse-entries" { return "DbInuseEntries" }
    if yname == "db-prefix" { return "DbPrefix" }
    return ""
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetSegmentPath() string {
    return "bank-db"
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["db-id"] = bankDb.DbId
    leafs["db-inuse-entries"] = bankDb.DbInuseEntries
    leafs["db-prefix"] = bankDb.DbPrefix
    return leafs
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetBundleName() string { return "cisco_ios_xr" }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetYangName() string { return "bank-db" }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) SetParent(parent types.Entity) { bankDb.parent = parent }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetParent() types.Entity { return bankDb.parent }

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetParentYangName() string { return "tcam-bank" }

