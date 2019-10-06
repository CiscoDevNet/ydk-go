// This module contains a collection of YANG definitions
// for Cisco IOS-XR fia-internal-tcam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   controller: Controller Resources
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller DPA operational data.
    Dpa Controller_Dpa
}

func (controller *Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "Cisco-IOS-XR-fia-internal-tcam-oper"
    controller.EntityData.SegmentPath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller"
    controller.EntityData.AbsolutePath = controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("dpa", types.YChild{"Dpa", &controller.Dpa})
    controller.EntityData.Leafs = types.NewOrderedMap()

    controller.EntityData.YListKeys = []string {}

    return &(controller.EntityData)
}

// Controller_Dpa
// Controller DPA operational data
type Controller_Dpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA data for available nodes.
    Nodes Controller_Dpa_Nodes
}

func (dpa *Controller_Dpa) GetEntityData() *types.CommonEntityData {
    dpa.EntityData.YFilter = dpa.YFilter
    dpa.EntityData.YangName = "dpa"
    dpa.EntityData.BundleName = "cisco_ios_xr"
    dpa.EntityData.ParentYangName = "controller"
    dpa.EntityData.SegmentPath = "dpa"
    dpa.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/" + dpa.EntityData.SegmentPath
    dpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dpa.EntityData.Children = types.NewOrderedMap()
    dpa.EntityData.Children.Append("nodes", types.YChild{"Nodes", &dpa.Nodes})
    dpa.EntityData.Leafs = types.NewOrderedMap()

    dpa.EntityData.YListKeys = []string {}

    return &(dpa.EntityData)
}

// Controller_Dpa_Nodes
// DPA data for available nodes
type Controller_Dpa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Controller_Dpa_Nodes_Node.
    Node []*Controller_Dpa_Nodes_Node
}

func (nodes *Controller_Dpa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dpa"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/" + nodes.EntityData.SegmentPath
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

// Controller_Dpa_Nodes_Node
// DPA operational data for a particular node
type Controller_Dpa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // External TCAM Resource Information.
    ExternalTcamResources Controller_Dpa_Nodes_Node_ExternalTcamResources

    // Internal TCAM Resource Information.
    InternalTcamResources Controller_Dpa_Nodes_Node_InternalTcamResources
}

func (node *Controller_Dpa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("external-tcam-resources", types.YChild{"ExternalTcamResources", &node.ExternalTcamResources})
    node.EntityData.Children.Append("internal-tcam-resources", types.YChild{"InternalTcamResources", &node.InternalTcamResources})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Controller_Dpa_Nodes_Node_ExternalTcamResources
// External TCAM Resource Information
type Controller_Dpa_Nodes_Node_ExternalTcamResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // npu tcam. The type is slice of
    // Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam.
    NpuTcam []*Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam
}

func (externalTcamResources *Controller_Dpa_Nodes_Node_ExternalTcamResources) GetEntityData() *types.CommonEntityData {
    externalTcamResources.EntityData.YFilter = externalTcamResources.YFilter
    externalTcamResources.EntityData.YangName = "external-tcam-resources"
    externalTcamResources.EntityData.BundleName = "cisco_ios_xr"
    externalTcamResources.EntityData.ParentYangName = "node"
    externalTcamResources.EntityData.SegmentPath = "external-tcam-resources"
    externalTcamResources.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/" + externalTcamResources.EntityData.SegmentPath
    externalTcamResources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalTcamResources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalTcamResources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalTcamResources.EntityData.Children = types.NewOrderedMap()
    externalTcamResources.EntityData.Children.Append("npu-tcam", types.YChild{"NpuTcam", nil})
    for i := range externalTcamResources.NpuTcam {
        types.SetYListKey(externalTcamResources.NpuTcam[i], i)
        externalTcamResources.EntityData.Children.Append(types.GetSegmentPath(externalTcamResources.NpuTcam[i]), types.YChild{"NpuTcam", externalTcamResources.NpuTcam[i]})
    }
    externalTcamResources.EntityData.Leafs = types.NewOrderedMap()

    externalTcamResources.EntityData.YListKeys = []string {}

    return &(externalTcamResources.EntityData)
}

// Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam
// npu tcam
type Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // npu id. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // tcam bank. The type is slice of
    // Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank.
    TcamBank []*Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank
}

func (npuTcam *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam) GetEntityData() *types.CommonEntityData {
    npuTcam.EntityData.YFilter = npuTcam.YFilter
    npuTcam.EntityData.YangName = "npu-tcam"
    npuTcam.EntityData.BundleName = "cisco_ios_xr"
    npuTcam.EntityData.ParentYangName = "external-tcam-resources"
    npuTcam.EntityData.SegmentPath = "npu-tcam" + types.AddNoKeyToken(npuTcam)
    npuTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/external-tcam-resources/" + npuTcam.EntityData.SegmentPath
    npuTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuTcam.EntityData.Children = types.NewOrderedMap()
    npuTcam.EntityData.Children.Append("tcam-bank", types.YChild{"TcamBank", nil})
    for i := range npuTcam.TcamBank {
        types.SetYListKey(npuTcam.TcamBank[i], i)
        npuTcam.EntityData.Children.Append(types.GetSegmentPath(npuTcam.TcamBank[i]), types.YChild{"TcamBank", npuTcam.TcamBank[i]})
    }
    npuTcam.EntityData.Leafs = types.NewOrderedMap()
    npuTcam.EntityData.Leafs.Append("npu-id", types.YLeaf{"NpuId", npuTcam.NpuId})

    npuTcam.EntityData.YListKeys = []string {}

    return &(npuTcam.EntityData)
}

// Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank
// tcam bank
type Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    BankDb []*Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb
}

func (tcamBank *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank) GetEntityData() *types.CommonEntityData {
    tcamBank.EntityData.YFilter = tcamBank.YFilter
    tcamBank.EntityData.YangName = "tcam-bank"
    tcamBank.EntityData.BundleName = "cisco_ios_xr"
    tcamBank.EntityData.ParentYangName = "npu-tcam"
    tcamBank.EntityData.SegmentPath = "tcam-bank" + types.AddNoKeyToken(tcamBank)
    tcamBank.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/external-tcam-resources/npu-tcam/" + tcamBank.EntityData.SegmentPath
    tcamBank.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamBank.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamBank.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamBank.EntityData.Children = types.NewOrderedMap()
    tcamBank.EntityData.Children.Append("bank-db", types.YChild{"BankDb", nil})
    for i := range tcamBank.BankDb {
        types.SetYListKey(tcamBank.BankDb[i], i)
        tcamBank.EntityData.Children.Append(types.GetSegmentPath(tcamBank.BankDb[i]), types.YChild{"BankDb", tcamBank.BankDb[i]})
    }
    tcamBank.EntityData.Leafs = types.NewOrderedMap()
    tcamBank.EntityData.Leafs.Append("bank-id", types.YLeaf{"BankId", tcamBank.BankId})
    tcamBank.EntityData.Leafs.Append("bank-key-size", types.YLeaf{"BankKeySize", tcamBank.BankKeySize})
    tcamBank.EntityData.Leafs.Append("bank-free-entries", types.YLeaf{"BankFreeEntries", tcamBank.BankFreeEntries})
    tcamBank.EntityData.Leafs.Append("bank-inuse-entries", types.YLeaf{"BankInuseEntries", tcamBank.BankInuseEntries})
    tcamBank.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", tcamBank.Owner})
    tcamBank.EntityData.Leafs.Append("nof-dbs", types.YLeaf{"NofDbs", tcamBank.NofDbs})

    tcamBank.EntityData.YListKeys = []string {}

    return &(tcamBank.EntityData)
}

// Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb
// bank db
type Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // db id. The type is interface{} with range: 0..4294967295.
    DbId interface{}

    // db inuse entries. The type is interface{} with range: 0..4294967295.
    DbInuseEntries interface{}

    // db prefix. The type is string.
    DbPrefix interface{}
}

func (bankDb *Controller_Dpa_Nodes_Node_ExternalTcamResources_NpuTcam_TcamBank_BankDb) GetEntityData() *types.CommonEntityData {
    bankDb.EntityData.YFilter = bankDb.YFilter
    bankDb.EntityData.YangName = "bank-db"
    bankDb.EntityData.BundleName = "cisco_ios_xr"
    bankDb.EntityData.ParentYangName = "tcam-bank"
    bankDb.EntityData.SegmentPath = "bank-db" + types.AddNoKeyToken(bankDb)
    bankDb.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/external-tcam-resources/npu-tcam/tcam-bank/" + bankDb.EntityData.SegmentPath
    bankDb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bankDb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bankDb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bankDb.EntityData.Children = types.NewOrderedMap()
    bankDb.EntityData.Leafs = types.NewOrderedMap()
    bankDb.EntityData.Leafs.Append("db-id", types.YLeaf{"DbId", bankDb.DbId})
    bankDb.EntityData.Leafs.Append("db-inuse-entries", types.YLeaf{"DbInuseEntries", bankDb.DbInuseEntries})
    bankDb.EntityData.Leafs.Append("db-prefix", types.YLeaf{"DbPrefix", bankDb.DbPrefix})

    bankDb.EntityData.YListKeys = []string {}

    return &(bankDb.EntityData)
}

// Controller_Dpa_Nodes_Node_InternalTcamResources
// Internal TCAM Resource Information
type Controller_Dpa_Nodes_Node_InternalTcamResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // npu tcam. The type is slice of
    // Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam.
    NpuTcam []*Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam
}

func (internalTcamResources *Controller_Dpa_Nodes_Node_InternalTcamResources) GetEntityData() *types.CommonEntityData {
    internalTcamResources.EntityData.YFilter = internalTcamResources.YFilter
    internalTcamResources.EntityData.YangName = "internal-tcam-resources"
    internalTcamResources.EntityData.BundleName = "cisco_ios_xr"
    internalTcamResources.EntityData.ParentYangName = "node"
    internalTcamResources.EntityData.SegmentPath = "internal-tcam-resources"
    internalTcamResources.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/" + internalTcamResources.EntityData.SegmentPath
    internalTcamResources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalTcamResources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalTcamResources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalTcamResources.EntityData.Children = types.NewOrderedMap()
    internalTcamResources.EntityData.Children.Append("npu-tcam", types.YChild{"NpuTcam", nil})
    for i := range internalTcamResources.NpuTcam {
        types.SetYListKey(internalTcamResources.NpuTcam[i], i)
        internalTcamResources.EntityData.Children.Append(types.GetSegmentPath(internalTcamResources.NpuTcam[i]), types.YChild{"NpuTcam", internalTcamResources.NpuTcam[i]})
    }
    internalTcamResources.EntityData.Leafs = types.NewOrderedMap()

    internalTcamResources.EntityData.YListKeys = []string {}

    return &(internalTcamResources.EntityData)
}

// Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam
// npu tcam
type Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // npu id. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // tcam bank. The type is slice of
    // Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank.
    TcamBank []*Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank
}

func (npuTcam *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam) GetEntityData() *types.CommonEntityData {
    npuTcam.EntityData.YFilter = npuTcam.YFilter
    npuTcam.EntityData.YangName = "npu-tcam"
    npuTcam.EntityData.BundleName = "cisco_ios_xr"
    npuTcam.EntityData.ParentYangName = "internal-tcam-resources"
    npuTcam.EntityData.SegmentPath = "npu-tcam" + types.AddNoKeyToken(npuTcam)
    npuTcam.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/internal-tcam-resources/" + npuTcam.EntityData.SegmentPath
    npuTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuTcam.EntityData.Children = types.NewOrderedMap()
    npuTcam.EntityData.Children.Append("tcam-bank", types.YChild{"TcamBank", nil})
    for i := range npuTcam.TcamBank {
        types.SetYListKey(npuTcam.TcamBank[i], i)
        npuTcam.EntityData.Children.Append(types.GetSegmentPath(npuTcam.TcamBank[i]), types.YChild{"TcamBank", npuTcam.TcamBank[i]})
    }
    npuTcam.EntityData.Leafs = types.NewOrderedMap()
    npuTcam.EntityData.Leafs.Append("npu-id", types.YLeaf{"NpuId", npuTcam.NpuId})

    npuTcam.EntityData.YListKeys = []string {}

    return &(npuTcam.EntityData)
}

// Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank
// tcam bank
type Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    BankDb []*Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb
}

func (tcamBank *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank) GetEntityData() *types.CommonEntityData {
    tcamBank.EntityData.YFilter = tcamBank.YFilter
    tcamBank.EntityData.YangName = "tcam-bank"
    tcamBank.EntityData.BundleName = "cisco_ios_xr"
    tcamBank.EntityData.ParentYangName = "npu-tcam"
    tcamBank.EntityData.SegmentPath = "tcam-bank" + types.AddNoKeyToken(tcamBank)
    tcamBank.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/internal-tcam-resources/npu-tcam/" + tcamBank.EntityData.SegmentPath
    tcamBank.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcamBank.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcamBank.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcamBank.EntityData.Children = types.NewOrderedMap()
    tcamBank.EntityData.Children.Append("bank-db", types.YChild{"BankDb", nil})
    for i := range tcamBank.BankDb {
        types.SetYListKey(tcamBank.BankDb[i], i)
        tcamBank.EntityData.Children.Append(types.GetSegmentPath(tcamBank.BankDb[i]), types.YChild{"BankDb", tcamBank.BankDb[i]})
    }
    tcamBank.EntityData.Leafs = types.NewOrderedMap()
    tcamBank.EntityData.Leafs.Append("bank-id", types.YLeaf{"BankId", tcamBank.BankId})
    tcamBank.EntityData.Leafs.Append("bank-key-size", types.YLeaf{"BankKeySize", tcamBank.BankKeySize})
    tcamBank.EntityData.Leafs.Append("bank-free-entries", types.YLeaf{"BankFreeEntries", tcamBank.BankFreeEntries})
    tcamBank.EntityData.Leafs.Append("bank-inuse-entries", types.YLeaf{"BankInuseEntries", tcamBank.BankInuseEntries})
    tcamBank.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", tcamBank.Owner})
    tcamBank.EntityData.Leafs.Append("nof-dbs", types.YLeaf{"NofDbs", tcamBank.NofDbs})

    tcamBank.EntityData.YListKeys = []string {}

    return &(tcamBank.EntityData)
}

// Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb
// bank db
type Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // db id. The type is interface{} with range: 0..4294967295.
    DbId interface{}

    // db inuse entries. The type is interface{} with range: 0..4294967295.
    DbInuseEntries interface{}

    // db prefix. The type is string.
    DbPrefix interface{}
}

func (bankDb *Controller_Dpa_Nodes_Node_InternalTcamResources_NpuTcam_TcamBank_BankDb) GetEntityData() *types.CommonEntityData {
    bankDb.EntityData.YFilter = bankDb.YFilter
    bankDb.EntityData.YangName = "bank-db"
    bankDb.EntityData.BundleName = "cisco_ios_xr"
    bankDb.EntityData.ParentYangName = "tcam-bank"
    bankDb.EntityData.SegmentPath = "bank-db" + types.AddNoKeyToken(bankDb)
    bankDb.EntityData.AbsolutePath = "Cisco-IOS-XR-fia-internal-tcam-oper:controller/dpa/nodes/node/internal-tcam-resources/npu-tcam/tcam-bank/" + bankDb.EntityData.SegmentPath
    bankDb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bankDb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bankDb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bankDb.EntityData.Children = types.NewOrderedMap()
    bankDb.EntityData.Leafs = types.NewOrderedMap()
    bankDb.EntityData.Leafs.Append("db-id", types.YLeaf{"DbId", bankDb.DbId})
    bankDb.EntityData.Leafs.Append("db-inuse-entries", types.YLeaf{"DbInuseEntries", bankDb.DbInuseEntries})
    bankDb.EntityData.Leafs.Append("db-prefix", types.YLeaf{"DbPrefix", bankDb.DbPrefix})

    bankDb.EntityData.YListKeys = []string {}

    return &(bankDb.EntityData)
}

