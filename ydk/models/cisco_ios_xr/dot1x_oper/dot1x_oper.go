// This module contains a collection of YANG definitions
// for Cisco IOS-XR dot1x package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dot1x: Dot1x operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package dot1x_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dot1x_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dot1x-oper dot1x}", reflect.TypeOf(Dot1x{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dot1x-oper:dot1x", reflect.TypeOf(Dot1x{}))
}

// Dot1x
// Dot1x operational data
type Dot1x struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x operational data.
    Statistics Dot1x_Statistics

    // Node-specific Dot1x operational data.
    Nodes Dot1x_Nodes

    // Dot1x operational data.
    Session Dot1x_Session
}

func (dot1x *Dot1x) GetEntityData() *types.CommonEntityData {
    dot1x.EntityData.YFilter = dot1x.YFilter
    dot1x.EntityData.YangName = "dot1x"
    dot1x.EntityData.BundleName = "cisco_ios_xr"
    dot1x.EntityData.ParentYangName = "Cisco-IOS-XR-dot1x-oper"
    dot1x.EntityData.SegmentPath = "Cisco-IOS-XR-dot1x-oper:dot1x"
    dot1x.EntityData.AbsolutePath = dot1x.EntityData.SegmentPath
    dot1x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1x.EntityData.Children = types.NewOrderedMap()
    dot1x.EntityData.Children.Append("statistics", types.YChild{"Statistics", &dot1x.Statistics})
    dot1x.EntityData.Children.Append("nodes", types.YChild{"Nodes", &dot1x.Nodes})
    dot1x.EntityData.Children.Append("session", types.YChild{"Session", &dot1x.Session})
    dot1x.EntityData.Leafs = types.NewOrderedMap()

    dot1x.EntityData.YListKeys = []string {}

    return &(dot1x.EntityData)
}

// Dot1x_Statistics
// Dot1x operational data
type Dot1x_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces with Dot1x.
    InterfaceStatistics Dot1x_Statistics_InterfaceStatistics
}

func (statistics *Dot1x_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "dot1x"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("interface-statistics", types.YChild{"InterfaceStatistics", &statistics.InterfaceStatistics})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics
// Interfaces with Dot1x
type Dot1x_Statistics_InterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Data for that Interface. The type is slice of
    // Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic.
    InterfaceStatistic []*Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic
}

func (interfaceStatistics *Dot1x_Statistics_InterfaceStatistics) GetEntityData() *types.CommonEntityData {
    interfaceStatistics.EntityData.YFilter = interfaceStatistics.YFilter
    interfaceStatistics.EntityData.YangName = "interface-statistics"
    interfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistics.EntityData.ParentYangName = "statistics"
    interfaceStatistics.EntityData.SegmentPath = "interface-statistics"
    interfaceStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/" + interfaceStatistics.EntityData.SegmentPath
    interfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistics.EntityData.Children = types.NewOrderedMap()
    interfaceStatistics.EntityData.Children.Append("interface-statistic", types.YChild{"InterfaceStatistic", nil})
    for i := range interfaceStatistics.InterfaceStatistic {
        interfaceStatistics.EntityData.Children.Append(types.GetSegmentPath(interfaceStatistics.InterfaceStatistic[i]), types.YChild{"InterfaceStatistic", interfaceStatistics.InterfaceStatistic[i]})
    }
    interfaceStatistics.EntityData.Leafs = types.NewOrderedMap()

    interfaceStatistics.EntityData.YListKeys = []string {}

    return &(interfaceStatistics.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic
// Dot1x Data for that Interface
type Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    Name interface{}

    // Interface Display name . The type is string.
    InterfaceName interface{}

    // PAE type on interface. The type is string.
    Pae interface{}

    // Dot1x interface database Statistics.
    Idb Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Idb

    // Dot1x Authenticator Port Statistics.
    Auth Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth

    // Dot1x Supplicant Port Statistics.
    Supp Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Supp

    // Dot1x Local EAP Port Statistics.
    LocalEap Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_LocalEap
}

func (interfaceStatistic *Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic) GetEntityData() *types.CommonEntityData {
    interfaceStatistic.EntityData.YFilter = interfaceStatistic.YFilter
    interfaceStatistic.EntityData.YangName = "interface-statistic"
    interfaceStatistic.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistic.EntityData.ParentYangName = "interface-statistics"
    interfaceStatistic.EntityData.SegmentPath = "interface-statistic" + types.AddKeyToken(interfaceStatistic.Name, "name")
    interfaceStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/interface-statistics/" + interfaceStatistic.EntityData.SegmentPath
    interfaceStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistic.EntityData.Children = types.NewOrderedMap()
    interfaceStatistic.EntityData.Children.Append("idb", types.YChild{"Idb", &interfaceStatistic.Idb})
    interfaceStatistic.EntityData.Children.Append("auth", types.YChild{"Auth", &interfaceStatistic.Auth})
    interfaceStatistic.EntityData.Children.Append("supp", types.YChild{"Supp", &interfaceStatistic.Supp})
    interfaceStatistic.EntityData.Children.Append("local-eap", types.YChild{"LocalEap", &interfaceStatistic.LocalEap})
    interfaceStatistic.EntityData.Leafs = types.NewOrderedMap()
    interfaceStatistic.EntityData.Leafs.Append("name", types.YLeaf{"Name", interfaceStatistic.Name})
    interfaceStatistic.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceStatistic.InterfaceName})
    interfaceStatistic.EntityData.Leafs.Append("pae", types.YLeaf{"Pae", interfaceStatistic.Pae})

    interfaceStatistic.EntityData.YListKeys = []string {"Name"}

    return &(interfaceStatistic.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Idb
// Dot1x interface database Statistics
type Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Idb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // NoRxOnIntfDown. The type is interface{} with range: 0..4294967295.
    NoRxOnIntfDown interface{}
}

func (idb *Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetEntityData() *types.CommonEntityData {
    idb.EntityData.YFilter = idb.YFilter
    idb.EntityData.YangName = "idb"
    idb.EntityData.BundleName = "cisco_ios_xr"
    idb.EntityData.ParentYangName = "interface-statistic"
    idb.EntityData.SegmentPath = "idb"
    idb.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/interface-statistics/interface-statistic/" + idb.EntityData.SegmentPath
    idb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idb.EntityData.Children = types.NewOrderedMap()
    idb.EntityData.Leafs = types.NewOrderedMap()
    idb.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", idb.RxTotal})
    idb.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", idb.TxTotal})
    idb.EntityData.Leafs.Append("no-rx-on-intf-down", types.YLeaf{"NoRxOnIntfDown", idb.NoRxOnIntfDown})

    idb.EntityData.YListKeys = []string {}

    return &(idb.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth
// Dot1x Authenticator Port Statistics
type Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxStart. The type is interface{} with range: 0..4294967295.
    RxStart interface{}

    // RxLogoff. The type is interface{} with range: 0..4294967295.
    RxLogoff interface{}

    // RxResp. The type is interface{} with range: 0..4294967295.
    RxResp interface{}

    // RxRespID. The type is interface{} with range: 0..4294967295.
    RxRespId interface{}

    // RxInvalid. The type is interface{} with range: 0..4294967295.
    RxInvalid interface{}

    // RxLenErr. The type is interface{} with range: 0..4294967295.
    RxLenErr interface{}

    // RxMyMacErr. The type is interface{} with range: 0..4294967295.
    RxMyMacErr interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxReq. The type is interface{} with range: 0..4294967295.
    TxReq interface{}

    // TxReqID. The type is interface{} with range: 0..4294967295.
    TxReqid interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // PacketDrop. The type is interface{} with range: 0..4294967295.
    PacketDropNoConfigReceived interface{}

    // PortControl.
    PortControl Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth_PortControl
}

func (auth *Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetEntityData() *types.CommonEntityData {
    auth.EntityData.YFilter = auth.YFilter
    auth.EntityData.YangName = "auth"
    auth.EntityData.BundleName = "cisco_ios_xr"
    auth.EntityData.ParentYangName = "interface-statistic"
    auth.EntityData.SegmentPath = "auth"
    auth.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/interface-statistics/interface-statistic/" + auth.EntityData.SegmentPath
    auth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auth.EntityData.Children = types.NewOrderedMap()
    auth.EntityData.Children.Append("port-control", types.YChild{"PortControl", &auth.PortControl})
    auth.EntityData.Leafs = types.NewOrderedMap()
    auth.EntityData.Leafs.Append("rx-start", types.YLeaf{"RxStart", auth.RxStart})
    auth.EntityData.Leafs.Append("rx-logoff", types.YLeaf{"RxLogoff", auth.RxLogoff})
    auth.EntityData.Leafs.Append("rx-resp", types.YLeaf{"RxResp", auth.RxResp})
    auth.EntityData.Leafs.Append("rx-resp-id", types.YLeaf{"RxRespId", auth.RxRespId})
    auth.EntityData.Leafs.Append("rx-invalid", types.YLeaf{"RxInvalid", auth.RxInvalid})
    auth.EntityData.Leafs.Append("rx-len-err", types.YLeaf{"RxLenErr", auth.RxLenErr})
    auth.EntityData.Leafs.Append("rx-my-mac-err", types.YLeaf{"RxMyMacErr", auth.RxMyMacErr})
    auth.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", auth.RxTotal})
    auth.EntityData.Leafs.Append("tx-req", types.YLeaf{"TxReq", auth.TxReq})
    auth.EntityData.Leafs.Append("tx-reqid", types.YLeaf{"TxReqid", auth.TxReqid})
    auth.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", auth.TxTotal})
    auth.EntityData.Leafs.Append("packet-drop-no-config-received", types.YLeaf{"PacketDropNoConfigReceived", auth.PacketDropNoConfigReceived})

    auth.EntityData.YListKeys = []string {}

    return &(auth.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth_PortControl
// PortControl
type Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth_PortControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EnableSucc. The type is interface{} with range: 0..4294967295.
    EnableSucc interface{}

    // EnableFail. The type is interface{} with range: 0..4294967295.
    EnableFail interface{}

    // AddClientSucc. The type is interface{} with range: 0..4294967295.
    AddClientSucc interface{}

    // AddClientFail. The type is interface{} with range: 0..4294967295.
    AddClientFail interface{}

    // RemoveClientSucc. The type is interface{} with range: 0..4294967295.
    RemoveClientSucc interface{}

    // RemoveClientFail. The type is interface{} with range: 0..4294967295.
    RemoveClientFail interface{}
}

func (portControl *Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Auth_PortControl) GetEntityData() *types.CommonEntityData {
    portControl.EntityData.YFilter = portControl.YFilter
    portControl.EntityData.YangName = "port-control"
    portControl.EntityData.BundleName = "cisco_ios_xr"
    portControl.EntityData.ParentYangName = "auth"
    portControl.EntityData.SegmentPath = "port-control"
    portControl.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/interface-statistics/interface-statistic/auth/" + portControl.EntityData.SegmentPath
    portControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portControl.EntityData.Children = types.NewOrderedMap()
    portControl.EntityData.Leafs = types.NewOrderedMap()
    portControl.EntityData.Leafs.Append("enable-succ", types.YLeaf{"EnableSucc", portControl.EnableSucc})
    portControl.EntityData.Leafs.Append("enable-fail", types.YLeaf{"EnableFail", portControl.EnableFail})
    portControl.EntityData.Leafs.Append("add-client-succ", types.YLeaf{"AddClientSucc", portControl.AddClientSucc})
    portControl.EntityData.Leafs.Append("add-client-fail", types.YLeaf{"AddClientFail", portControl.AddClientFail})
    portControl.EntityData.Leafs.Append("remove-client-succ", types.YLeaf{"RemoveClientSucc", portControl.RemoveClientSucc})
    portControl.EntityData.Leafs.Append("remove-client-fail", types.YLeaf{"RemoveClientFail", portControl.RemoveClientFail})

    portControl.EntityData.YListKeys = []string {}

    return &(portControl.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Supp
// Dot1x Supplicant Port Statistics
type Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Supp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxReq. The type is interface{} with range: 0..4294967295.
    RxReq interface{}

    // RxInvalid. The type is interface{} with range: 0..4294967295.
    RxInvalid interface{}

    // RxLenErr. The type is interface{} with range: 0..4294967295.
    RxLenErr interface{}

    // RxMyMacErr. The type is interface{} with range: 0..4294967295.
    RxMyMacErr interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxStart. The type is interface{} with range: 0..4294967295.
    TxStart interface{}

    // TxLogoff. The type is interface{} with range: 0..4294967295.
    TxLogoff interface{}

    // TxResp. The type is interface{} with range: 0..4294967295.
    TxResp interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}
}

func (supp *Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetEntityData() *types.CommonEntityData {
    supp.EntityData.YFilter = supp.YFilter
    supp.EntityData.YangName = "supp"
    supp.EntityData.BundleName = "cisco_ios_xr"
    supp.EntityData.ParentYangName = "interface-statistic"
    supp.EntityData.SegmentPath = "supp"
    supp.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/interface-statistics/interface-statistic/" + supp.EntityData.SegmentPath
    supp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supp.EntityData.Children = types.NewOrderedMap()
    supp.EntityData.Leafs = types.NewOrderedMap()
    supp.EntityData.Leafs.Append("rx-req", types.YLeaf{"RxReq", supp.RxReq})
    supp.EntityData.Leafs.Append("rx-invalid", types.YLeaf{"RxInvalid", supp.RxInvalid})
    supp.EntityData.Leafs.Append("rx-len-err", types.YLeaf{"RxLenErr", supp.RxLenErr})
    supp.EntityData.Leafs.Append("rx-my-mac-err", types.YLeaf{"RxMyMacErr", supp.RxMyMacErr})
    supp.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", supp.RxTotal})
    supp.EntityData.Leafs.Append("tx-start", types.YLeaf{"TxStart", supp.TxStart})
    supp.EntityData.Leafs.Append("tx-logoff", types.YLeaf{"TxLogoff", supp.TxLogoff})
    supp.EntityData.Leafs.Append("tx-resp", types.YLeaf{"TxResp", supp.TxResp})
    supp.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", supp.TxTotal})

    supp.EntityData.YListKeys = []string {}

    return &(supp.EntityData)
}

// Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_LocalEap
// Dot1x Local EAP Port Statistics
type Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_LocalEap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests. The type is interface{} with range: 0..4294967295.
    Requests interface{}

    // Replies. The type is interface{} with range: 0..4294967295.
    Replies interface{}

    // Timeout. The type is interface{} with range: 0..4294967295.
    Timeout interface{}

    // DroppedNoEAP. The type is interface{} with range: 0..4294967295.
    DroppedNoEap interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}

    // Success. The type is interface{} with range: 0..4294967295.
    Success interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}
}

func (localEap *Dot1x_Statistics_InterfaceStatistics_InterfaceStatistic_LocalEap) GetEntityData() *types.CommonEntityData {
    localEap.EntityData.YFilter = localEap.YFilter
    localEap.EntityData.YangName = "local-eap"
    localEap.EntityData.BundleName = "cisco_ios_xr"
    localEap.EntityData.ParentYangName = "interface-statistic"
    localEap.EntityData.SegmentPath = "local-eap"
    localEap.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/statistics/interface-statistics/interface-statistic/" + localEap.EntityData.SegmentPath
    localEap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localEap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localEap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localEap.EntityData.Children = types.NewOrderedMap()
    localEap.EntityData.Leafs = types.NewOrderedMap()
    localEap.EntityData.Leafs.Append("requests", types.YLeaf{"Requests", localEap.Requests})
    localEap.EntityData.Leafs.Append("replies", types.YLeaf{"Replies", localEap.Replies})
    localEap.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", localEap.Timeout})
    localEap.EntityData.Leafs.Append("dropped-no-eap", types.YLeaf{"DroppedNoEap", localEap.DroppedNoEap})
    localEap.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", localEap.Dropped})
    localEap.EntityData.Leafs.Append("success", types.YLeaf{"Success", localEap.Success})
    localEap.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", localEap.Failed})

    localEap.EntityData.YListKeys = []string {}

    return &(localEap.EntityData)
}

// Dot1x_Nodes
// Node-specific Dot1x operational data
type Dot1x_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x operational data for a particular node. The type is slice of
    // Dot1x_Nodes_Node.
    Node []*Dot1x_Nodes_Node
}

func (nodes *Dot1x_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dot1x"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/" + nodes.EntityData.SegmentPath
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

// Dot1x_Nodes_Node
// Dot1x operational data for a particular node
type Dot1x_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Dot1x Default Values.
    Dot1xDefaults Dot1x_Nodes_Node_Dot1xDefaults

    // Dot1x Default Values.
    Statistics Dot1x_Nodes_Node_Statistics
}

func (node *Dot1x_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("dot1x-defaults", types.YChild{"Dot1xDefaults", &node.Dot1xDefaults})
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Dot1x_Nodes_Node_Dot1xDefaults
// Dot1x Default Values
type Dot1x_Nodes_Node_Dot1xDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Protocol Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Dot1x Authenticator default Timer values.
    AuthTimers Dot1x_Nodes_Node_Dot1xDefaults_AuthTimers

    // Dot1x Supllicant default Timer values.
    SuppTimers Dot1x_Nodes_Node_Dot1xDefaults_SuppTimers
}

func (dot1xDefaults *Dot1x_Nodes_Node_Dot1xDefaults) GetEntityData() *types.CommonEntityData {
    dot1xDefaults.EntityData.YFilter = dot1xDefaults.YFilter
    dot1xDefaults.EntityData.YangName = "dot1x-defaults"
    dot1xDefaults.EntityData.BundleName = "cisco_ios_xr"
    dot1xDefaults.EntityData.ParentYangName = "node"
    dot1xDefaults.EntityData.SegmentPath = "dot1x-defaults"
    dot1xDefaults.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/" + dot1xDefaults.EntityData.SegmentPath
    dot1xDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1xDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1xDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1xDefaults.EntityData.Children = types.NewOrderedMap()
    dot1xDefaults.EntityData.Children.Append("auth-timers", types.YChild{"AuthTimers", &dot1xDefaults.AuthTimers})
    dot1xDefaults.EntityData.Children.Append("supp-timers", types.YChild{"SuppTimers", &dot1xDefaults.SuppTimers})
    dot1xDefaults.EntityData.Leafs = types.NewOrderedMap()
    dot1xDefaults.EntityData.Leafs.Append("version", types.YLeaf{"Version", dot1xDefaults.Version})

    dot1xDefaults.EntityData.YListKeys = []string {}

    return &(dot1xDefaults.EntityData)
}

// Dot1x_Nodes_Node_Dot1xDefaults_AuthTimers
// Dot1x Authenticator default Timer values
type Dot1x_Nodes_Node_Dot1xDefaults_AuthTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // in Seconds, authenticator remains quiet (in the HELD state) following a
    // failed authentication exchange before trying to reauthenticate the client.
    // The type is interface{} with range: 0..4294967295. Units are second.
    QuietPeriod interface{}

    // in Seconds, Timeout for supplicant reply, authenticator-to-supplicant
    // retransmission time for EAP-request-ID packets (assuming that no response
    // is received) from the client. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TxPeriod interface{}

    // Max No. of Reauthentication Attempts (or) retransmits an EAP-request-ID
    // frame to the client before restarting the authentication process. The type
    // is interface{} with range: 0..4294967295.
    MaxReauthReq interface{}

    // in Seconds, Timeout for supplicant reply, authenticator-to-supplicant
    // retransmission time for all EAP messages except for EAP Request ID packets.
    // The type is interface{} with range: 0..4294967295. Units are second.
    SuppTimeout interface{}

    // Max No. of EAP-Req (except for EAP-Request-ID) retransmits
    // (authenticator-to-supplicant) before sending EAP-Failure. The type is
    // interface{} with range: 0..4294967295.
    MaxReq interface{}

    // in Seconds,  after which an automatic reauthentication should be initiated.
    // The type is interface{} with range: 0..4294967295. Units are second.
    ReauthPeriod interface{}
}

func (authTimers *Dot1x_Nodes_Node_Dot1xDefaults_AuthTimers) GetEntityData() *types.CommonEntityData {
    authTimers.EntityData.YFilter = authTimers.YFilter
    authTimers.EntityData.YangName = "auth-timers"
    authTimers.EntityData.BundleName = "cisco_ios_xr"
    authTimers.EntityData.ParentYangName = "dot1x-defaults"
    authTimers.EntityData.SegmentPath = "auth-timers"
    authTimers.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/dot1x-defaults/" + authTimers.EntityData.SegmentPath
    authTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authTimers.EntityData.Children = types.NewOrderedMap()
    authTimers.EntityData.Leafs = types.NewOrderedMap()
    authTimers.EntityData.Leafs.Append("quiet-period", types.YLeaf{"QuietPeriod", authTimers.QuietPeriod})
    authTimers.EntityData.Leafs.Append("tx-period", types.YLeaf{"TxPeriod", authTimers.TxPeriod})
    authTimers.EntityData.Leafs.Append("max-reauth-req", types.YLeaf{"MaxReauthReq", authTimers.MaxReauthReq})
    authTimers.EntityData.Leafs.Append("supp-timeout", types.YLeaf{"SuppTimeout", authTimers.SuppTimeout})
    authTimers.EntityData.Leafs.Append("max-req", types.YLeaf{"MaxReq", authTimers.MaxReq})
    authTimers.EntityData.Leafs.Append("reauth-period", types.YLeaf{"ReauthPeriod", authTimers.ReauthPeriod})

    authTimers.EntityData.YListKeys = []string {}

    return &(authTimers.EntityData)
}

// Dot1x_Nodes_Node_Dot1xDefaults_SuppTimers
// Dot1x Supllicant default Timer values
type Dot1x_Nodes_Node_Dot1xDefaults_SuppTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // in Seconds, supplicant waits for a response from an authenticator except
    // for EAPOL-START before timing out. The type is interface{} with range:
    // 0..4294967295. Units are second.
    AuthPeriod interface{}

    // in Seconds, supplicant will stay in the HELD state (that is, the length of
    // time it will wait before trying to send the credentials again after a
    // failed attempt). The type is interface{} with range: 0..4294967295. Units
    // are second.
    HeldPeriod interface{}

    // Configures the interval, in seconds, between two successive EAPOL-Start
    // frames when they are being retransmitted. The type is interface{} with
    // range: 0..4294967295. Units are second.
    StartPeriod interface{}

    // Max No. of EAPOL-Start frames supplicant can send to the authenticator. The
    // type is interface{} with range: 0..4294967295.
    MaxStart interface{}
}

func (suppTimers *Dot1x_Nodes_Node_Dot1xDefaults_SuppTimers) GetEntityData() *types.CommonEntityData {
    suppTimers.EntityData.YFilter = suppTimers.YFilter
    suppTimers.EntityData.YangName = "supp-timers"
    suppTimers.EntityData.BundleName = "cisco_ios_xr"
    suppTimers.EntityData.ParentYangName = "dot1x-defaults"
    suppTimers.EntityData.SegmentPath = "supp-timers"
    suppTimers.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/dot1x-defaults/" + suppTimers.EntityData.SegmentPath
    suppTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppTimers.EntityData.Children = types.NewOrderedMap()
    suppTimers.EntityData.Leafs = types.NewOrderedMap()
    suppTimers.EntityData.Leafs.Append("auth-period", types.YLeaf{"AuthPeriod", suppTimers.AuthPeriod})
    suppTimers.EntityData.Leafs.Append("held-period", types.YLeaf{"HeldPeriod", suppTimers.HeldPeriod})
    suppTimers.EntityData.Leafs.Append("start-period", types.YLeaf{"StartPeriod", suppTimers.StartPeriod})
    suppTimers.EntityData.Leafs.Append("max-start", types.YLeaf{"MaxStart", suppTimers.MaxStart})

    suppTimers.EntityData.YListKeys = []string {}

    return &(suppTimers.EntityData)
}

// Dot1x_Nodes_Node_Statistics
// Dot1x Default Values
type Dot1x_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global statistics.
    GlStats Dot1x_Nodes_Node_Statistics_GlStats

    // dot1x interface statistics list. The type is slice of
    // Dot1x_Nodes_Node_Statistics_IfStats.
    IfStats []*Dot1x_Nodes_Node_Statistics_IfStats
}

func (statistics *Dot1x_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("gl-stats", types.YChild{"GlStats", &statistics.GlStats})
    statistics.EntityData.Children.Append("if-stats", types.YChild{"IfStats", nil})
    for i := range statistics.IfStats {
        types.SetYListKey(statistics.IfStats[i], i)
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.IfStats[i]), types.YChild{"IfStats", statistics.IfStats[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Dot1x_Nodes_Node_Statistics_GlStats
// Global statistics
type Dot1x_Nodes_Node_Statistics_GlStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // RxNoIDB. The type is interface{} with range: 0..4294967295.
    RxNoIdb interface{}

    // PacketDrop. The type is interface{} with range: 0..4294967295.
    PacketDropNoConfigReceived interface{}

    // PortControl.
    PortControl Dot1x_Nodes_Node_Statistics_GlStats_PortControl
}

func (glStats *Dot1x_Nodes_Node_Statistics_GlStats) GetEntityData() *types.CommonEntityData {
    glStats.EntityData.YFilter = glStats.YFilter
    glStats.EntityData.YangName = "gl-stats"
    glStats.EntityData.BundleName = "cisco_ios_xr"
    glStats.EntityData.ParentYangName = "statistics"
    glStats.EntityData.SegmentPath = "gl-stats"
    glStats.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/" + glStats.EntityData.SegmentPath
    glStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    glStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    glStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    glStats.EntityData.Children = types.NewOrderedMap()
    glStats.EntityData.Children.Append("port-control", types.YChild{"PortControl", &glStats.PortControl})
    glStats.EntityData.Leafs = types.NewOrderedMap()
    glStats.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", glStats.TxTotal})
    glStats.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", glStats.RxTotal})
    glStats.EntityData.Leafs.Append("rx-no-idb", types.YLeaf{"RxNoIdb", glStats.RxNoIdb})
    glStats.EntityData.Leafs.Append("packet-drop-no-config-received", types.YLeaf{"PacketDropNoConfigReceived", glStats.PacketDropNoConfigReceived})

    glStats.EntityData.YListKeys = []string {}

    return &(glStats.EntityData)
}

// Dot1x_Nodes_Node_Statistics_GlStats_PortControl
// PortControl
type Dot1x_Nodes_Node_Statistics_GlStats_PortControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EnableSucc. The type is interface{} with range: 0..4294967295.
    EnableSucc interface{}

    // EnableFail. The type is interface{} with range: 0..4294967295.
    EnableFail interface{}

    // DisableSucc. The type is interface{} with range: 0..4294967295.
    DisableSucc interface{}

    // DisableFail. The type is interface{} with range: 0..4294967295.
    DisableFail interface{}

    // AddClientSucc. The type is interface{} with range: 0..4294967295.
    AddClientSucc interface{}

    // AddClientFail. The type is interface{} with range: 0..4294967295.
    AddClientFail interface{}

    // RemoveClientSucc. The type is interface{} with range: 0..4294967295.
    RemoveClientSucc interface{}

    // RemoveClientFail. The type is interface{} with range: 0..4294967295.
    RemoveClientFail interface{}
}

func (portControl *Dot1x_Nodes_Node_Statistics_GlStats_PortControl) GetEntityData() *types.CommonEntityData {
    portControl.EntityData.YFilter = portControl.YFilter
    portControl.EntityData.YangName = "port-control"
    portControl.EntityData.BundleName = "cisco_ios_xr"
    portControl.EntityData.ParentYangName = "gl-stats"
    portControl.EntityData.SegmentPath = "port-control"
    portControl.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/gl-stats/" + portControl.EntityData.SegmentPath
    portControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portControl.EntityData.Children = types.NewOrderedMap()
    portControl.EntityData.Leafs = types.NewOrderedMap()
    portControl.EntityData.Leafs.Append("enable-succ", types.YLeaf{"EnableSucc", portControl.EnableSucc})
    portControl.EntityData.Leafs.Append("enable-fail", types.YLeaf{"EnableFail", portControl.EnableFail})
    portControl.EntityData.Leafs.Append("disable-succ", types.YLeaf{"DisableSucc", portControl.DisableSucc})
    portControl.EntityData.Leafs.Append("disable-fail", types.YLeaf{"DisableFail", portControl.DisableFail})
    portControl.EntityData.Leafs.Append("add-client-succ", types.YLeaf{"AddClientSucc", portControl.AddClientSucc})
    portControl.EntityData.Leafs.Append("add-client-fail", types.YLeaf{"AddClientFail", portControl.AddClientFail})
    portControl.EntityData.Leafs.Append("remove-client-succ", types.YLeaf{"RemoveClientSucc", portControl.RemoveClientSucc})
    portControl.EntityData.Leafs.Append("remove-client-fail", types.YLeaf{"RemoveClientFail", portControl.RemoveClientFail})

    portControl.EntityData.YListKeys = []string {}

    return &(portControl.EntityData)
}

// Dot1x_Nodes_Node_Statistics_IfStats
// dot1x interface statistics list
type Dot1x_Nodes_Node_Statistics_IfStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface Display name . The type is string.
    InterfaceName interface{}

    // PAE type on interface. The type is string.
    Pae interface{}

    // Dot1x interface database Statistics.
    Idb Dot1x_Nodes_Node_Statistics_IfStats_Idb

    // Dot1x Authenticator Port Statistics.
    Auth Dot1x_Nodes_Node_Statistics_IfStats_Auth

    // Dot1x Supplicant Port Statistics.
    Supp Dot1x_Nodes_Node_Statistics_IfStats_Supp

    // Dot1x Local EAP Port Statistics.
    LocalEap Dot1x_Nodes_Node_Statistics_IfStats_LocalEap
}

func (ifStats *Dot1x_Nodes_Node_Statistics_IfStats) GetEntityData() *types.CommonEntityData {
    ifStats.EntityData.YFilter = ifStats.YFilter
    ifStats.EntityData.YangName = "if-stats"
    ifStats.EntityData.BundleName = "cisco_ios_xr"
    ifStats.EntityData.ParentYangName = "statistics"
    ifStats.EntityData.SegmentPath = "if-stats" + types.AddNoKeyToken(ifStats)
    ifStats.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/" + ifStats.EntityData.SegmentPath
    ifStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifStats.EntityData.Children = types.NewOrderedMap()
    ifStats.EntityData.Children.Append("idb", types.YChild{"Idb", &ifStats.Idb})
    ifStats.EntityData.Children.Append("auth", types.YChild{"Auth", &ifStats.Auth})
    ifStats.EntityData.Children.Append("supp", types.YChild{"Supp", &ifStats.Supp})
    ifStats.EntityData.Children.Append("local-eap", types.YChild{"LocalEap", &ifStats.LocalEap})
    ifStats.EntityData.Leafs = types.NewOrderedMap()
    ifStats.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ifStats.InterfaceName})
    ifStats.EntityData.Leafs.Append("pae", types.YLeaf{"Pae", ifStats.Pae})

    ifStats.EntityData.YListKeys = []string {}

    return &(ifStats.EntityData)
}

// Dot1x_Nodes_Node_Statistics_IfStats_Idb
// Dot1x interface database Statistics
type Dot1x_Nodes_Node_Statistics_IfStats_Idb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // NoRxOnIntfDown. The type is interface{} with range: 0..4294967295.
    NoRxOnIntfDown interface{}
}

func (idb *Dot1x_Nodes_Node_Statistics_IfStats_Idb) GetEntityData() *types.CommonEntityData {
    idb.EntityData.YFilter = idb.YFilter
    idb.EntityData.YangName = "idb"
    idb.EntityData.BundleName = "cisco_ios_xr"
    idb.EntityData.ParentYangName = "if-stats"
    idb.EntityData.SegmentPath = "idb"
    idb.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/if-stats/" + idb.EntityData.SegmentPath
    idb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idb.EntityData.Children = types.NewOrderedMap()
    idb.EntityData.Leafs = types.NewOrderedMap()
    idb.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", idb.RxTotal})
    idb.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", idb.TxTotal})
    idb.EntityData.Leafs.Append("no-rx-on-intf-down", types.YLeaf{"NoRxOnIntfDown", idb.NoRxOnIntfDown})

    idb.EntityData.YListKeys = []string {}

    return &(idb.EntityData)
}

// Dot1x_Nodes_Node_Statistics_IfStats_Auth
// Dot1x Authenticator Port Statistics
type Dot1x_Nodes_Node_Statistics_IfStats_Auth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxStart. The type is interface{} with range: 0..4294967295.
    RxStart interface{}

    // RxLogoff. The type is interface{} with range: 0..4294967295.
    RxLogoff interface{}

    // RxResp. The type is interface{} with range: 0..4294967295.
    RxResp interface{}

    // RxRespID. The type is interface{} with range: 0..4294967295.
    RxRespId interface{}

    // RxInvalid. The type is interface{} with range: 0..4294967295.
    RxInvalid interface{}

    // RxLenErr. The type is interface{} with range: 0..4294967295.
    RxLenErr interface{}

    // RxMyMacErr. The type is interface{} with range: 0..4294967295.
    RxMyMacErr interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxReq. The type is interface{} with range: 0..4294967295.
    TxReq interface{}

    // TxReqID. The type is interface{} with range: 0..4294967295.
    TxReqid interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // PacketDrop. The type is interface{} with range: 0..4294967295.
    PacketDropNoConfigReceived interface{}

    // PortControl.
    PortControl Dot1x_Nodes_Node_Statistics_IfStats_Auth_PortControl
}

func (auth *Dot1x_Nodes_Node_Statistics_IfStats_Auth) GetEntityData() *types.CommonEntityData {
    auth.EntityData.YFilter = auth.YFilter
    auth.EntityData.YangName = "auth"
    auth.EntityData.BundleName = "cisco_ios_xr"
    auth.EntityData.ParentYangName = "if-stats"
    auth.EntityData.SegmentPath = "auth"
    auth.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/if-stats/" + auth.EntityData.SegmentPath
    auth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auth.EntityData.Children = types.NewOrderedMap()
    auth.EntityData.Children.Append("port-control", types.YChild{"PortControl", &auth.PortControl})
    auth.EntityData.Leafs = types.NewOrderedMap()
    auth.EntityData.Leafs.Append("rx-start", types.YLeaf{"RxStart", auth.RxStart})
    auth.EntityData.Leafs.Append("rx-logoff", types.YLeaf{"RxLogoff", auth.RxLogoff})
    auth.EntityData.Leafs.Append("rx-resp", types.YLeaf{"RxResp", auth.RxResp})
    auth.EntityData.Leafs.Append("rx-resp-id", types.YLeaf{"RxRespId", auth.RxRespId})
    auth.EntityData.Leafs.Append("rx-invalid", types.YLeaf{"RxInvalid", auth.RxInvalid})
    auth.EntityData.Leafs.Append("rx-len-err", types.YLeaf{"RxLenErr", auth.RxLenErr})
    auth.EntityData.Leafs.Append("rx-my-mac-err", types.YLeaf{"RxMyMacErr", auth.RxMyMacErr})
    auth.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", auth.RxTotal})
    auth.EntityData.Leafs.Append("tx-req", types.YLeaf{"TxReq", auth.TxReq})
    auth.EntityData.Leafs.Append("tx-reqid", types.YLeaf{"TxReqid", auth.TxReqid})
    auth.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", auth.TxTotal})
    auth.EntityData.Leafs.Append("packet-drop-no-config-received", types.YLeaf{"PacketDropNoConfigReceived", auth.PacketDropNoConfigReceived})

    auth.EntityData.YListKeys = []string {}

    return &(auth.EntityData)
}

// Dot1x_Nodes_Node_Statistics_IfStats_Auth_PortControl
// PortControl
type Dot1x_Nodes_Node_Statistics_IfStats_Auth_PortControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EnableSucc. The type is interface{} with range: 0..4294967295.
    EnableSucc interface{}

    // EnableFail. The type is interface{} with range: 0..4294967295.
    EnableFail interface{}

    // AddClientSucc. The type is interface{} with range: 0..4294967295.
    AddClientSucc interface{}

    // AddClientFail. The type is interface{} with range: 0..4294967295.
    AddClientFail interface{}

    // RemoveClientSucc. The type is interface{} with range: 0..4294967295.
    RemoveClientSucc interface{}

    // RemoveClientFail. The type is interface{} with range: 0..4294967295.
    RemoveClientFail interface{}
}

func (portControl *Dot1x_Nodes_Node_Statistics_IfStats_Auth_PortControl) GetEntityData() *types.CommonEntityData {
    portControl.EntityData.YFilter = portControl.YFilter
    portControl.EntityData.YangName = "port-control"
    portControl.EntityData.BundleName = "cisco_ios_xr"
    portControl.EntityData.ParentYangName = "auth"
    portControl.EntityData.SegmentPath = "port-control"
    portControl.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/if-stats/auth/" + portControl.EntityData.SegmentPath
    portControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portControl.EntityData.Children = types.NewOrderedMap()
    portControl.EntityData.Leafs = types.NewOrderedMap()
    portControl.EntityData.Leafs.Append("enable-succ", types.YLeaf{"EnableSucc", portControl.EnableSucc})
    portControl.EntityData.Leafs.Append("enable-fail", types.YLeaf{"EnableFail", portControl.EnableFail})
    portControl.EntityData.Leafs.Append("add-client-succ", types.YLeaf{"AddClientSucc", portControl.AddClientSucc})
    portControl.EntityData.Leafs.Append("add-client-fail", types.YLeaf{"AddClientFail", portControl.AddClientFail})
    portControl.EntityData.Leafs.Append("remove-client-succ", types.YLeaf{"RemoveClientSucc", portControl.RemoveClientSucc})
    portControl.EntityData.Leafs.Append("remove-client-fail", types.YLeaf{"RemoveClientFail", portControl.RemoveClientFail})

    portControl.EntityData.YListKeys = []string {}

    return &(portControl.EntityData)
}

// Dot1x_Nodes_Node_Statistics_IfStats_Supp
// Dot1x Supplicant Port Statistics
type Dot1x_Nodes_Node_Statistics_IfStats_Supp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxReq. The type is interface{} with range: 0..4294967295.
    RxReq interface{}

    // RxInvalid. The type is interface{} with range: 0..4294967295.
    RxInvalid interface{}

    // RxLenErr. The type is interface{} with range: 0..4294967295.
    RxLenErr interface{}

    // RxMyMacErr. The type is interface{} with range: 0..4294967295.
    RxMyMacErr interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxStart. The type is interface{} with range: 0..4294967295.
    TxStart interface{}

    // TxLogoff. The type is interface{} with range: 0..4294967295.
    TxLogoff interface{}

    // TxResp. The type is interface{} with range: 0..4294967295.
    TxResp interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}
}

func (supp *Dot1x_Nodes_Node_Statistics_IfStats_Supp) GetEntityData() *types.CommonEntityData {
    supp.EntityData.YFilter = supp.YFilter
    supp.EntityData.YangName = "supp"
    supp.EntityData.BundleName = "cisco_ios_xr"
    supp.EntityData.ParentYangName = "if-stats"
    supp.EntityData.SegmentPath = "supp"
    supp.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/if-stats/" + supp.EntityData.SegmentPath
    supp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supp.EntityData.Children = types.NewOrderedMap()
    supp.EntityData.Leafs = types.NewOrderedMap()
    supp.EntityData.Leafs.Append("rx-req", types.YLeaf{"RxReq", supp.RxReq})
    supp.EntityData.Leafs.Append("rx-invalid", types.YLeaf{"RxInvalid", supp.RxInvalid})
    supp.EntityData.Leafs.Append("rx-len-err", types.YLeaf{"RxLenErr", supp.RxLenErr})
    supp.EntityData.Leafs.Append("rx-my-mac-err", types.YLeaf{"RxMyMacErr", supp.RxMyMacErr})
    supp.EntityData.Leafs.Append("rx-total", types.YLeaf{"RxTotal", supp.RxTotal})
    supp.EntityData.Leafs.Append("tx-start", types.YLeaf{"TxStart", supp.TxStart})
    supp.EntityData.Leafs.Append("tx-logoff", types.YLeaf{"TxLogoff", supp.TxLogoff})
    supp.EntityData.Leafs.Append("tx-resp", types.YLeaf{"TxResp", supp.TxResp})
    supp.EntityData.Leafs.Append("tx-total", types.YLeaf{"TxTotal", supp.TxTotal})

    supp.EntityData.YListKeys = []string {}

    return &(supp.EntityData)
}

// Dot1x_Nodes_Node_Statistics_IfStats_LocalEap
// Dot1x Local EAP Port Statistics
type Dot1x_Nodes_Node_Statistics_IfStats_LocalEap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests. The type is interface{} with range: 0..4294967295.
    Requests interface{}

    // Replies. The type is interface{} with range: 0..4294967295.
    Replies interface{}

    // Timeout. The type is interface{} with range: 0..4294967295.
    Timeout interface{}

    // DroppedNoEAP. The type is interface{} with range: 0..4294967295.
    DroppedNoEap interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}

    // Success. The type is interface{} with range: 0..4294967295.
    Success interface{}

    // Failed. The type is interface{} with range: 0..4294967295.
    Failed interface{}
}

func (localEap *Dot1x_Nodes_Node_Statistics_IfStats_LocalEap) GetEntityData() *types.CommonEntityData {
    localEap.EntityData.YFilter = localEap.YFilter
    localEap.EntityData.YangName = "local-eap"
    localEap.EntityData.BundleName = "cisco_ios_xr"
    localEap.EntityData.ParentYangName = "if-stats"
    localEap.EntityData.SegmentPath = "local-eap"
    localEap.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/nodes/node/statistics/if-stats/" + localEap.EntityData.SegmentPath
    localEap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localEap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localEap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localEap.EntityData.Children = types.NewOrderedMap()
    localEap.EntityData.Leafs = types.NewOrderedMap()
    localEap.EntityData.Leafs.Append("requests", types.YLeaf{"Requests", localEap.Requests})
    localEap.EntityData.Leafs.Append("replies", types.YLeaf{"Replies", localEap.Replies})
    localEap.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", localEap.Timeout})
    localEap.EntityData.Leafs.Append("dropped-no-eap", types.YLeaf{"DroppedNoEap", localEap.DroppedNoEap})
    localEap.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", localEap.Dropped})
    localEap.EntityData.Leafs.Append("success", types.YLeaf{"Success", localEap.Success})
    localEap.EntityData.Leafs.Append("failed", types.YLeaf{"Failed", localEap.Failed})

    localEap.EntityData.YListKeys = []string {}

    return &(localEap.EntityData)
}

// Dot1x_Session
// Dot1x operational data
type Dot1x_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces with Dot1x.
    InterfaceSessions Dot1x_Session_InterfaceSessions
}

func (session *Dot1x_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "dot1x"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("interface-sessions", types.YChild{"InterfaceSessions", &session.InterfaceSessions})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// Dot1x_Session_InterfaceSessions
// Interfaces with Dot1x
type Dot1x_Session_InterfaceSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Data for that Interface. The type is slice of
    // Dot1x_Session_InterfaceSessions_InterfaceSession.
    InterfaceSession []*Dot1x_Session_InterfaceSessions_InterfaceSession
}

func (interfaceSessions *Dot1x_Session_InterfaceSessions) GetEntityData() *types.CommonEntityData {
    interfaceSessions.EntityData.YFilter = interfaceSessions.YFilter
    interfaceSessions.EntityData.YangName = "interface-sessions"
    interfaceSessions.EntityData.BundleName = "cisco_ios_xr"
    interfaceSessions.EntityData.ParentYangName = "session"
    interfaceSessions.EntityData.SegmentPath = "interface-sessions"
    interfaceSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/" + interfaceSessions.EntityData.SegmentPath
    interfaceSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSessions.EntityData.Children = types.NewOrderedMap()
    interfaceSessions.EntityData.Children.Append("interface-session", types.YChild{"InterfaceSession", nil})
    for i := range interfaceSessions.InterfaceSession {
        interfaceSessions.EntityData.Children.Append(types.GetSegmentPath(interfaceSessions.InterfaceSession[i]), types.YChild{"InterfaceSession", interfaceSessions.InterfaceSession[i]})
    }
    interfaceSessions.EntityData.Leafs = types.NewOrderedMap()

    interfaceSessions.EntityData.YListKeys = []string {}

    return &(interfaceSessions.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession
// Dot1x Data for that Interface
type Dot1x_Session_InterfaceSessions_InterfaceSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    Name interface{}

    // Interface Display name . The type is string.
    InterfaceName interface{}

    // Interface Display short_name . The type is string.
    InterfaceSname interface{}

    // Interface handle. The type is string.
    IfHandle interface{}

    // formatted MAC Address. The type is string.
    Mac interface{}

    // EAPOL Ethertype. The type is string.
    Ethertype interface{}

    // EAPOL Destination Address. The type is string.
    EapolAddr interface{}

    // Dot1x interface Info.
    IntfInfo Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo

    // MKA session secure status.
    MkaStatusInfo Dot1x_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo
}

func (interfaceSession *Dot1x_Session_InterfaceSessions_InterfaceSession) GetEntityData() *types.CommonEntityData {
    interfaceSession.EntityData.YFilter = interfaceSession.YFilter
    interfaceSession.EntityData.YangName = "interface-session"
    interfaceSession.EntityData.BundleName = "cisco_ios_xr"
    interfaceSession.EntityData.ParentYangName = "interface-sessions"
    interfaceSession.EntityData.SegmentPath = "interface-session" + types.AddKeyToken(interfaceSession.Name, "name")
    interfaceSession.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/" + interfaceSession.EntityData.SegmentPath
    interfaceSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSession.EntityData.Children = types.NewOrderedMap()
    interfaceSession.EntityData.Children.Append("intf-info", types.YChild{"IntfInfo", &interfaceSession.IntfInfo})
    interfaceSession.EntityData.Children.Append("mka-status-info", types.YChild{"MkaStatusInfo", &interfaceSession.MkaStatusInfo})
    interfaceSession.EntityData.Leafs = types.NewOrderedMap()
    interfaceSession.EntityData.Leafs.Append("name", types.YLeaf{"Name", interfaceSession.Name})
    interfaceSession.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceSession.InterfaceName})
    interfaceSession.EntityData.Leafs.Append("interface-sname", types.YLeaf{"InterfaceSname", interfaceSession.InterfaceSname})
    interfaceSession.EntityData.Leafs.Append("if-handle", types.YLeaf{"IfHandle", interfaceSession.IfHandle})
    interfaceSession.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", interfaceSession.Mac})
    interfaceSession.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", interfaceSession.Ethertype})
    interfaceSession.EntityData.Leafs.Append("eapol-addr", types.YLeaf{"EapolAddr", interfaceSession.EapolAddr})

    interfaceSession.EntityData.YListKeys = []string {"Name"}

    return &(interfaceSession.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo
// Dot1x interface Info
type Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PAE type on interface. The type is string.
    Pae interface{}

    // Dot1x Port Status. The type is string.
    PortStatus interface{}

    // Dot1x Profile. The type is string.
    Dot1xProfile interface{}

    // L2 Transport Info. The type is bool.
    L2Transport interface{}

    // Dot1x Authenticator info.
    AuthInfo Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo

    // Dot1x Supplicant info.
    SuppInfo Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo
}

func (intfInfo *Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetEntityData() *types.CommonEntityData {
    intfInfo.EntityData.YFilter = intfInfo.YFilter
    intfInfo.EntityData.YangName = "intf-info"
    intfInfo.EntityData.BundleName = "cisco_ios_xr"
    intfInfo.EntityData.ParentYangName = "interface-session"
    intfInfo.EntityData.SegmentPath = "intf-info"
    intfInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/interface-session/" + intfInfo.EntityData.SegmentPath
    intfInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfInfo.EntityData.Children = types.NewOrderedMap()
    intfInfo.EntityData.Children.Append("auth-info", types.YChild{"AuthInfo", &intfInfo.AuthInfo})
    intfInfo.EntityData.Children.Append("supp-info", types.YChild{"SuppInfo", &intfInfo.SuppInfo})
    intfInfo.EntityData.Leafs = types.NewOrderedMap()
    intfInfo.EntityData.Leafs.Append("pae", types.YLeaf{"Pae", intfInfo.Pae})
    intfInfo.EntityData.Leafs.Append("port-status", types.YLeaf{"PortStatus", intfInfo.PortStatus})
    intfInfo.EntityData.Leafs.Append("dot1x-profile", types.YLeaf{"Dot1xProfile", intfInfo.Dot1xProfile})
    intfInfo.EntityData.Leafs.Append("l2-transport", types.YLeaf{"L2Transport", intfInfo.L2Transport})

    intfInfo.EntityData.YListKeys = []string {}

    return &(intfInfo.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo
// Dot1x Authenticator info
type Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port Control Feature. The type is string.
    PortControl interface{}

    // Re-Authentication enabled status. The type is string.
    Reauth interface{}

    // Configuration Dependency . The type is string.
    ConfigDependency interface{}

    // EAP profile. The type is string.
    EapProfile interface{}

    // Authenticator client list. The type is slice of
    // Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client.
    Client []*Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client
}

func (authInfo *Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetEntityData() *types.CommonEntityData {
    authInfo.EntityData.YFilter = authInfo.YFilter
    authInfo.EntityData.YangName = "auth-info"
    authInfo.EntityData.BundleName = "cisco_ios_xr"
    authInfo.EntityData.ParentYangName = "intf-info"
    authInfo.EntityData.SegmentPath = "auth-info"
    authInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/interface-session/intf-info/" + authInfo.EntityData.SegmentPath
    authInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authInfo.EntityData.Children = types.NewOrderedMap()
    authInfo.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range authInfo.Client {
        types.SetYListKey(authInfo.Client[i], i)
        authInfo.EntityData.Children.Append(types.GetSegmentPath(authInfo.Client[i]), types.YChild{"Client", authInfo.Client[i]})
    }
    authInfo.EntityData.Leafs = types.NewOrderedMap()
    authInfo.EntityData.Leafs.Append("port-control", types.YLeaf{"PortControl", authInfo.PortControl})
    authInfo.EntityData.Leafs.Append("reauth", types.YLeaf{"Reauth", authInfo.Reauth})
    authInfo.EntityData.Leafs.Append("config-dependency", types.YLeaf{"ConfigDependency", authInfo.ConfigDependency})
    authInfo.EntityData.Leafs.Append("eap-profile", types.YLeaf{"EapProfile", authInfo.EapProfile})

    authInfo.EntityData.YListKeys = []string {}

    return &(authInfo.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client
// Authenticator client list
type Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // formatted MAC Address. The type is string.
    Mac interface{}

    // Auth SM State. The type is string.
    AuthSmState interface{}

    // Auth back end SM State. The type is string.
    AuthBendSmState interface{}

    // remaining time for next reauthentication. The type is string.
    TimeToNextReauth interface{}

    // Last Authenticated Timestamp (formatted). The type is string.
    LastAuthTime interface{}

    // Last Authenticated Server. The type is string.
    LastAuthServer interface{}

    // Auth Client Port Control Status. The type is string.
    PortControl interface{}
}

func (client *Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "auth-info"
    client.EntityData.SegmentPath = "client" + types.AddNoKeyToken(client)
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/interface-session/intf-info/auth-info/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", client.Mac})
    client.EntityData.Leafs.Append("auth-sm-state", types.YLeaf{"AuthSmState", client.AuthSmState})
    client.EntityData.Leafs.Append("auth-bend-sm-state", types.YLeaf{"AuthBendSmState", client.AuthBendSmState})
    client.EntityData.Leafs.Append("time-to-next-reauth", types.YLeaf{"TimeToNextReauth", client.TimeToNextReauth})
    client.EntityData.Leafs.Append("last-auth-time", types.YLeaf{"LastAuthTime", client.LastAuthTime})
    client.EntityData.Leafs.Append("last-auth-server", types.YLeaf{"LastAuthServer", client.LastAuthServer})
    client.EntityData.Leafs.Append("port-control", types.YLeaf{"PortControl", client.PortControl})

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo
// Dot1x Supplicant info
type Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EAP profile. The type is string.
    EapProfile interface{}

    // Configuration Dependency . The type is string.
    ConfigDependency interface{}

    // Supp Client info. The type is slice of
    // Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client.
    Client []*Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client
}

func (suppInfo *Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetEntityData() *types.CommonEntityData {
    suppInfo.EntityData.YFilter = suppInfo.YFilter
    suppInfo.EntityData.YangName = "supp-info"
    suppInfo.EntityData.BundleName = "cisco_ios_xr"
    suppInfo.EntityData.ParentYangName = "intf-info"
    suppInfo.EntityData.SegmentPath = "supp-info"
    suppInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/interface-session/intf-info/" + suppInfo.EntityData.SegmentPath
    suppInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppInfo.EntityData.Children = types.NewOrderedMap()
    suppInfo.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range suppInfo.Client {
        types.SetYListKey(suppInfo.Client[i], i)
        suppInfo.EntityData.Children.Append(types.GetSegmentPath(suppInfo.Client[i]), types.YChild{"Client", suppInfo.Client[i]})
    }
    suppInfo.EntityData.Leafs = types.NewOrderedMap()
    suppInfo.EntityData.Leafs.Append("eap-profile", types.YLeaf{"EapProfile", suppInfo.EapProfile})
    suppInfo.EntityData.Leafs.Append("config-dependency", types.YLeaf{"ConfigDependency", suppInfo.ConfigDependency})

    suppInfo.EntityData.YListKeys = []string {}

    return &(suppInfo.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client
// Supp Client info
type Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // formatted MAC Address. The type is string.
    Mac interface{}

    // EAP Method. The type is string.
    EapMethod interface{}

    // Last Authenticated Timestamp (formatted). The type is string.
    LastAuthTime interface{}

    // supp SM State. The type is string.
    AuthSmState interface{}

    // supp back end SM State. The type is string.
    AuthBendSmState interface{}
}

func (client *Dot1x_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "supp-info"
    client.EntityData.SegmentPath = "client" + types.AddNoKeyToken(client)
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/interface-session/intf-info/supp-info/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", client.Mac})
    client.EntityData.Leafs.Append("eap-method", types.YLeaf{"EapMethod", client.EapMethod})
    client.EntityData.Leafs.Append("last-auth-time", types.YLeaf{"LastAuthTime", client.LastAuthTime})
    client.EntityData.Leafs.Append("auth-sm-state", types.YLeaf{"AuthSmState", client.AuthSmState})
    client.EntityData.Leafs.Append("auth-bend-sm-state", types.YLeaf{"AuthBendSmState", client.AuthBendSmState})

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// Dot1x_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo
// MKA session secure status
type Dot1x_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Tie breaker role chosen for mka when PAE type is BOTH. The type is
    // string.
    TieBreakRole interface{}

    // EAP Mode status for MKA. The type is string.
    EapBasedMacsec interface{}

    // Time stamp when Dot1x posting a message to  MKA to start session. The type
    // is string.
    MkaStartTime interface{}

    // Time stamp when Dot1x posting a message to  MKA to stop session. The type
    // is string.
    MkaStopTime interface{}

    // Time Stamp of MKA acknowledgement to Dot1x. The type is string.
    MkaResponseTime interface{}
}

func (mkaStatusInfo *Dot1x_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetEntityData() *types.CommonEntityData {
    mkaStatusInfo.EntityData.YFilter = mkaStatusInfo.YFilter
    mkaStatusInfo.EntityData.YangName = "mka-status-info"
    mkaStatusInfo.EntityData.BundleName = "cisco_ios_xr"
    mkaStatusInfo.EntityData.ParentYangName = "interface-session"
    mkaStatusInfo.EntityData.SegmentPath = "mka-status-info"
    mkaStatusInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dot1x-oper:dot1x/session/interface-sessions/interface-session/" + mkaStatusInfo.EntityData.SegmentPath
    mkaStatusInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mkaStatusInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mkaStatusInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mkaStatusInfo.EntityData.Children = types.NewOrderedMap()
    mkaStatusInfo.EntityData.Leafs = types.NewOrderedMap()
    mkaStatusInfo.EntityData.Leafs.Append("tie-break-role", types.YLeaf{"TieBreakRole", mkaStatusInfo.TieBreakRole})
    mkaStatusInfo.EntityData.Leafs.Append("eap-based-macsec", types.YLeaf{"EapBasedMacsec", mkaStatusInfo.EapBasedMacsec})
    mkaStatusInfo.EntityData.Leafs.Append("mka-start-time", types.YLeaf{"MkaStartTime", mkaStatusInfo.MkaStartTime})
    mkaStatusInfo.EntityData.Leafs.Append("mka-stop-time", types.YLeaf{"MkaStopTime", mkaStatusInfo.MkaStopTime})
    mkaStatusInfo.EntityData.Leafs.Append("mka-response-time", types.YLeaf{"MkaResponseTime", mkaStatusInfo.MkaResponseTime})

    mkaStatusInfo.EntityData.YListKeys = []string {}

    return &(mkaStatusInfo.EntityData)
}

