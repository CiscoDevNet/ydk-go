// This module contains a collection of YANG definitions
// for Cisco IOS-XR dot1x package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dot1x: Dot1x operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dot1x-oper dot1x}", reflect.TypeOf(Dot1X{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dot1x-oper:dot1x", reflect.TypeOf(Dot1X{}))
}

// Dot1X
// Dot1x operational data
type Dot1X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x operational data.
    Statistics Dot1X_Statistics

    // Node-specific Dot1x operational data.
    Nodes Dot1X_Nodes

    // Dot1x operational data.
    Session Dot1X_Session
}

func (dot1X *Dot1X) GetEntityData() *types.CommonEntityData {
    dot1X.EntityData.YFilter = dot1X.YFilter
    dot1X.EntityData.YangName = "dot1x"
    dot1X.EntityData.BundleName = "cisco_ios_xr"
    dot1X.EntityData.ParentYangName = "Cisco-IOS-XR-dot1x-oper"
    dot1X.EntityData.SegmentPath = "Cisco-IOS-XR-dot1x-oper:dot1x"
    dot1X.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1X.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1X.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1X.EntityData.Children = make(map[string]types.YChild)
    dot1X.EntityData.Children["statistics"] = types.YChild{"Statistics", &dot1X.Statistics}
    dot1X.EntityData.Children["nodes"] = types.YChild{"Nodes", &dot1X.Nodes}
    dot1X.EntityData.Children["session"] = types.YChild{"Session", &dot1X.Session}
    dot1X.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dot1X.EntityData)
}

// Dot1X_Statistics
// Dot1x operational data
type Dot1X_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces with Dot1x.
    InterfaceStatistics Dot1X_Statistics_InterfaceStatistics
}

func (statistics *Dot1X_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "dot1x"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["interface-statistics"] = types.YChild{"InterfaceStatistics", &statistics.InterfaceStatistics}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Dot1X_Statistics_InterfaceStatistics
// Interfaces with Dot1x
type Dot1X_Statistics_InterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Data for that Interface. The type is slice of
    // Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic.
    InterfaceStatistic []Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetEntityData() *types.CommonEntityData {
    interfaceStatistics.EntityData.YFilter = interfaceStatistics.YFilter
    interfaceStatistics.EntityData.YangName = "interface-statistics"
    interfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistics.EntityData.ParentYangName = "statistics"
    interfaceStatistics.EntityData.SegmentPath = "interface-statistics"
    interfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistics.EntityData.Children = make(map[string]types.YChild)
    interfaceStatistics.EntityData.Children["interface-statistic"] = types.YChild{"InterfaceStatistic", nil}
    for i := range interfaceStatistics.InterfaceStatistic {
        interfaceStatistics.EntityData.Children[types.GetSegmentPath(&interfaceStatistics.InterfaceStatistic[i])] = types.YChild{"InterfaceStatistic", &interfaceStatistics.InterfaceStatistic[i]}
    }
    interfaceStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceStatistics.EntityData)
}

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic
// Dot1x Data for that Interface
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Interface Display name . The type is string.
    InterfaceName interface{}

    // PAE type on interface. The type is string.
    Pae interface{}

    // Dot1x interface database Statistics.
    Idb Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb

    // Dot1x Authenticator Port Statistics.
    Auth Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth

    // Dot1x Supplicant Port Statistics.
    Supp Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp
}

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetEntityData() *types.CommonEntityData {
    interfaceStatistic.EntityData.YFilter = interfaceStatistic.YFilter
    interfaceStatistic.EntityData.YangName = "interface-statistic"
    interfaceStatistic.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistic.EntityData.ParentYangName = "interface-statistics"
    interfaceStatistic.EntityData.SegmentPath = "interface-statistic" + "[name='" + fmt.Sprintf("%v", interfaceStatistic.Name) + "']"
    interfaceStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistic.EntityData.Children = make(map[string]types.YChild)
    interfaceStatistic.EntityData.Children["idb"] = types.YChild{"Idb", &interfaceStatistic.Idb}
    interfaceStatistic.EntityData.Children["auth"] = types.YChild{"Auth", &interfaceStatistic.Auth}
    interfaceStatistic.EntityData.Children["supp"] = types.YChild{"Supp", &interfaceStatistic.Supp}
    interfaceStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceStatistic.EntityData.Leafs["name"] = types.YLeaf{"Name", interfaceStatistic.Name}
    interfaceStatistic.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceStatistic.InterfaceName}
    interfaceStatistic.EntityData.Leafs["pae"] = types.YLeaf{"Pae", interfaceStatistic.Pae}
    return &(interfaceStatistic.EntityData)
}

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb
// Dot1x interface database Statistics
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // NoRxOnIntfDown. The type is interface{} with range: 0..4294967295.
    NoRxOnIntfDown interface{}
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetEntityData() *types.CommonEntityData {
    idb.EntityData.YFilter = idb.YFilter
    idb.EntityData.YangName = "idb"
    idb.EntityData.BundleName = "cisco_ios_xr"
    idb.EntityData.ParentYangName = "interface-statistic"
    idb.EntityData.SegmentPath = "idb"
    idb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idb.EntityData.Children = make(map[string]types.YChild)
    idb.EntityData.Leafs = make(map[string]types.YLeaf)
    idb.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", idb.RxTotal}
    idb.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", idb.TxTotal}
    idb.EntityData.Leafs["no-rx-on-intf-down"] = types.YLeaf{"NoRxOnIntfDown", idb.NoRxOnIntfDown}
    return &(idb.EntityData)
}

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth
// Dot1x Authenticator Port Statistics
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth struct {
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
}

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetEntityData() *types.CommonEntityData {
    auth.EntityData.YFilter = auth.YFilter
    auth.EntityData.YangName = "auth"
    auth.EntityData.BundleName = "cisco_ios_xr"
    auth.EntityData.ParentYangName = "interface-statistic"
    auth.EntityData.SegmentPath = "auth"
    auth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auth.EntityData.Children = make(map[string]types.YChild)
    auth.EntityData.Leafs = make(map[string]types.YLeaf)
    auth.EntityData.Leafs["rx-start"] = types.YLeaf{"RxStart", auth.RxStart}
    auth.EntityData.Leafs["rx-logoff"] = types.YLeaf{"RxLogoff", auth.RxLogoff}
    auth.EntityData.Leafs["rx-resp"] = types.YLeaf{"RxResp", auth.RxResp}
    auth.EntityData.Leafs["rx-resp-id"] = types.YLeaf{"RxRespId", auth.RxRespId}
    auth.EntityData.Leafs["rx-invalid"] = types.YLeaf{"RxInvalid", auth.RxInvalid}
    auth.EntityData.Leafs["rx-len-err"] = types.YLeaf{"RxLenErr", auth.RxLenErr}
    auth.EntityData.Leafs["rx-my-mac-err"] = types.YLeaf{"RxMyMacErr", auth.RxMyMacErr}
    auth.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", auth.RxTotal}
    auth.EntityData.Leafs["tx-req"] = types.YLeaf{"TxReq", auth.TxReq}
    auth.EntityData.Leafs["tx-reqid"] = types.YLeaf{"TxReqid", auth.TxReqid}
    auth.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", auth.TxTotal}
    return &(auth.EntityData)
}

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp
// Dot1x Supplicant Port Statistics
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp struct {
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

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetEntityData() *types.CommonEntityData {
    supp.EntityData.YFilter = supp.YFilter
    supp.EntityData.YangName = "supp"
    supp.EntityData.BundleName = "cisco_ios_xr"
    supp.EntityData.ParentYangName = "interface-statistic"
    supp.EntityData.SegmentPath = "supp"
    supp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supp.EntityData.Children = make(map[string]types.YChild)
    supp.EntityData.Leafs = make(map[string]types.YLeaf)
    supp.EntityData.Leafs["rx-req"] = types.YLeaf{"RxReq", supp.RxReq}
    supp.EntityData.Leafs["rx-invalid"] = types.YLeaf{"RxInvalid", supp.RxInvalid}
    supp.EntityData.Leafs["rx-len-err"] = types.YLeaf{"RxLenErr", supp.RxLenErr}
    supp.EntityData.Leafs["rx-my-mac-err"] = types.YLeaf{"RxMyMacErr", supp.RxMyMacErr}
    supp.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", supp.RxTotal}
    supp.EntityData.Leafs["tx-start"] = types.YLeaf{"TxStart", supp.TxStart}
    supp.EntityData.Leafs["tx-logoff"] = types.YLeaf{"TxLogoff", supp.TxLogoff}
    supp.EntityData.Leafs["tx-resp"] = types.YLeaf{"TxResp", supp.TxResp}
    supp.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", supp.TxTotal}
    return &(supp.EntityData)
}

// Dot1X_Nodes
// Node-specific Dot1x operational data
type Dot1X_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x operational data for a particular node. The type is slice of
    // Dot1X_Nodes_Node.
    Node []Dot1X_Nodes_Node
}

func (nodes *Dot1X_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dot1x"
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

// Dot1X_Nodes_Node
// Dot1x operational data for a particular node
type Dot1X_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Dot1x Default Values.
    Dot1XDefaults Dot1X_Nodes_Node_Dot1XDefaults

    // Dot1x Default Values.
    Statistics Dot1X_Nodes_Node_Statistics
}

func (node *Dot1X_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["dot1x-defaults"] = types.YChild{"Dot1XDefaults", &node.Dot1XDefaults}
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Dot1X_Nodes_Node_Dot1XDefaults
// Dot1x Default Values
type Dot1X_Nodes_Node_Dot1XDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Protocol Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Dot1x Authenticator default Timer values.
    AuthTimers Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers

    // Dot1x Supllicant default Timer values.
    SuppTimers Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetEntityData() *types.CommonEntityData {
    dot1XDefaults.EntityData.YFilter = dot1XDefaults.YFilter
    dot1XDefaults.EntityData.YangName = "dot1x-defaults"
    dot1XDefaults.EntityData.BundleName = "cisco_ios_xr"
    dot1XDefaults.EntityData.ParentYangName = "node"
    dot1XDefaults.EntityData.SegmentPath = "dot1x-defaults"
    dot1XDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dot1XDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dot1XDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dot1XDefaults.EntityData.Children = make(map[string]types.YChild)
    dot1XDefaults.EntityData.Children["auth-timers"] = types.YChild{"AuthTimers", &dot1XDefaults.AuthTimers}
    dot1XDefaults.EntityData.Children["supp-timers"] = types.YChild{"SuppTimers", &dot1XDefaults.SuppTimers}
    dot1XDefaults.EntityData.Leafs = make(map[string]types.YLeaf)
    dot1XDefaults.EntityData.Leafs["version"] = types.YLeaf{"Version", dot1XDefaults.Version}
    return &(dot1XDefaults.EntityData)
}

// Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers
// Dot1x Authenticator default Timer values
type Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers struct {
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

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetEntityData() *types.CommonEntityData {
    authTimers.EntityData.YFilter = authTimers.YFilter
    authTimers.EntityData.YangName = "auth-timers"
    authTimers.EntityData.BundleName = "cisco_ios_xr"
    authTimers.EntityData.ParentYangName = "dot1x-defaults"
    authTimers.EntityData.SegmentPath = "auth-timers"
    authTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authTimers.EntityData.Children = make(map[string]types.YChild)
    authTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    authTimers.EntityData.Leafs["quiet-period"] = types.YLeaf{"QuietPeriod", authTimers.QuietPeriod}
    authTimers.EntityData.Leafs["tx-period"] = types.YLeaf{"TxPeriod", authTimers.TxPeriod}
    authTimers.EntityData.Leafs["max-reauth-req"] = types.YLeaf{"MaxReauthReq", authTimers.MaxReauthReq}
    authTimers.EntityData.Leafs["supp-timeout"] = types.YLeaf{"SuppTimeout", authTimers.SuppTimeout}
    authTimers.EntityData.Leafs["max-req"] = types.YLeaf{"MaxReq", authTimers.MaxReq}
    authTimers.EntityData.Leafs["reauth-period"] = types.YLeaf{"ReauthPeriod", authTimers.ReauthPeriod}
    return &(authTimers.EntityData)
}

// Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers
// Dot1x Supllicant default Timer values
type Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers struct {
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

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetEntityData() *types.CommonEntityData {
    suppTimers.EntityData.YFilter = suppTimers.YFilter
    suppTimers.EntityData.YangName = "supp-timers"
    suppTimers.EntityData.BundleName = "cisco_ios_xr"
    suppTimers.EntityData.ParentYangName = "dot1x-defaults"
    suppTimers.EntityData.SegmentPath = "supp-timers"
    suppTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppTimers.EntityData.Children = make(map[string]types.YChild)
    suppTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    suppTimers.EntityData.Leafs["auth-period"] = types.YLeaf{"AuthPeriod", suppTimers.AuthPeriod}
    suppTimers.EntityData.Leafs["held-period"] = types.YLeaf{"HeldPeriod", suppTimers.HeldPeriod}
    suppTimers.EntityData.Leafs["start-period"] = types.YLeaf{"StartPeriod", suppTimers.StartPeriod}
    suppTimers.EntityData.Leafs["max-start"] = types.YLeaf{"MaxStart", suppTimers.MaxStart}
    return &(suppTimers.EntityData)
}

// Dot1X_Nodes_Node_Statistics
// Dot1x Default Values
type Dot1X_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global statistics.
    GlStats Dot1X_Nodes_Node_Statistics_GlStats

    // dot1x interface statistics list. The type is slice of
    // Dot1X_Nodes_Node_Statistics_IfStats.
    IfStats []Dot1X_Nodes_Node_Statistics_IfStats
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["gl-stats"] = types.YChild{"GlStats", &statistics.GlStats}
    statistics.EntityData.Children["if-stats"] = types.YChild{"IfStats", nil}
    for i := range statistics.IfStats {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.IfStats[i])] = types.YChild{"IfStats", &statistics.IfStats[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Dot1X_Nodes_Node_Statistics_GlStats
// Global statistics
type Dot1X_Nodes_Node_Statistics_GlStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // RxNoIDB. The type is interface{} with range: 0..4294967295.
    RxNoIdb interface{}
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetEntityData() *types.CommonEntityData {
    glStats.EntityData.YFilter = glStats.YFilter
    glStats.EntityData.YangName = "gl-stats"
    glStats.EntityData.BundleName = "cisco_ios_xr"
    glStats.EntityData.ParentYangName = "statistics"
    glStats.EntityData.SegmentPath = "gl-stats"
    glStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    glStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    glStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    glStats.EntityData.Children = make(map[string]types.YChild)
    glStats.EntityData.Leafs = make(map[string]types.YLeaf)
    glStats.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", glStats.TxTotal}
    glStats.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", glStats.RxTotal}
    glStats.EntityData.Leafs["rx-no-idb"] = types.YLeaf{"RxNoIdb", glStats.RxNoIdb}
    return &(glStats.EntityData)
}

// Dot1X_Nodes_Node_Statistics_IfStats
// dot1x interface statistics list
type Dot1X_Nodes_Node_Statistics_IfStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Display name . The type is string.
    InterfaceName interface{}

    // PAE type on interface. The type is string.
    Pae interface{}

    // Dot1x interface database Statistics.
    Idb Dot1X_Nodes_Node_Statistics_IfStats_Idb

    // Dot1x Authenticator Port Statistics.
    Auth Dot1X_Nodes_Node_Statistics_IfStats_Auth

    // Dot1x Supplicant Port Statistics.
    Supp Dot1X_Nodes_Node_Statistics_IfStats_Supp
}

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetEntityData() *types.CommonEntityData {
    ifStats.EntityData.YFilter = ifStats.YFilter
    ifStats.EntityData.YangName = "if-stats"
    ifStats.EntityData.BundleName = "cisco_ios_xr"
    ifStats.EntityData.ParentYangName = "statistics"
    ifStats.EntityData.SegmentPath = "if-stats"
    ifStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifStats.EntityData.Children = make(map[string]types.YChild)
    ifStats.EntityData.Children["idb"] = types.YChild{"Idb", &ifStats.Idb}
    ifStats.EntityData.Children["auth"] = types.YChild{"Auth", &ifStats.Auth}
    ifStats.EntityData.Children["supp"] = types.YChild{"Supp", &ifStats.Supp}
    ifStats.EntityData.Leafs = make(map[string]types.YLeaf)
    ifStats.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", ifStats.InterfaceName}
    ifStats.EntityData.Leafs["pae"] = types.YLeaf{"Pae", ifStats.Pae}
    return &(ifStats.EntityData)
}

// Dot1X_Nodes_Node_Statistics_IfStats_Idb
// Dot1x interface database Statistics
type Dot1X_Nodes_Node_Statistics_IfStats_Idb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // NoRxOnIntfDown. The type is interface{} with range: 0..4294967295.
    NoRxOnIntfDown interface{}
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetEntityData() *types.CommonEntityData {
    idb.EntityData.YFilter = idb.YFilter
    idb.EntityData.YangName = "idb"
    idb.EntityData.BundleName = "cisco_ios_xr"
    idb.EntityData.ParentYangName = "if-stats"
    idb.EntityData.SegmentPath = "idb"
    idb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idb.EntityData.Children = make(map[string]types.YChild)
    idb.EntityData.Leafs = make(map[string]types.YLeaf)
    idb.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", idb.RxTotal}
    idb.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", idb.TxTotal}
    idb.EntityData.Leafs["no-rx-on-intf-down"] = types.YLeaf{"NoRxOnIntfDown", idb.NoRxOnIntfDown}
    return &(idb.EntityData)
}

// Dot1X_Nodes_Node_Statistics_IfStats_Auth
// Dot1x Authenticator Port Statistics
type Dot1X_Nodes_Node_Statistics_IfStats_Auth struct {
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
}

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetEntityData() *types.CommonEntityData {
    auth.EntityData.YFilter = auth.YFilter
    auth.EntityData.YangName = "auth"
    auth.EntityData.BundleName = "cisco_ios_xr"
    auth.EntityData.ParentYangName = "if-stats"
    auth.EntityData.SegmentPath = "auth"
    auth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auth.EntityData.Children = make(map[string]types.YChild)
    auth.EntityData.Leafs = make(map[string]types.YLeaf)
    auth.EntityData.Leafs["rx-start"] = types.YLeaf{"RxStart", auth.RxStart}
    auth.EntityData.Leafs["rx-logoff"] = types.YLeaf{"RxLogoff", auth.RxLogoff}
    auth.EntityData.Leafs["rx-resp"] = types.YLeaf{"RxResp", auth.RxResp}
    auth.EntityData.Leafs["rx-resp-id"] = types.YLeaf{"RxRespId", auth.RxRespId}
    auth.EntityData.Leafs["rx-invalid"] = types.YLeaf{"RxInvalid", auth.RxInvalid}
    auth.EntityData.Leafs["rx-len-err"] = types.YLeaf{"RxLenErr", auth.RxLenErr}
    auth.EntityData.Leafs["rx-my-mac-err"] = types.YLeaf{"RxMyMacErr", auth.RxMyMacErr}
    auth.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", auth.RxTotal}
    auth.EntityData.Leafs["tx-req"] = types.YLeaf{"TxReq", auth.TxReq}
    auth.EntityData.Leafs["tx-reqid"] = types.YLeaf{"TxReqid", auth.TxReqid}
    auth.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", auth.TxTotal}
    return &(auth.EntityData)
}

// Dot1X_Nodes_Node_Statistics_IfStats_Supp
// Dot1x Supplicant Port Statistics
type Dot1X_Nodes_Node_Statistics_IfStats_Supp struct {
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

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetEntityData() *types.CommonEntityData {
    supp.EntityData.YFilter = supp.YFilter
    supp.EntityData.YangName = "supp"
    supp.EntityData.BundleName = "cisco_ios_xr"
    supp.EntityData.ParentYangName = "if-stats"
    supp.EntityData.SegmentPath = "supp"
    supp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supp.EntityData.Children = make(map[string]types.YChild)
    supp.EntityData.Leafs = make(map[string]types.YLeaf)
    supp.EntityData.Leafs["rx-req"] = types.YLeaf{"RxReq", supp.RxReq}
    supp.EntityData.Leafs["rx-invalid"] = types.YLeaf{"RxInvalid", supp.RxInvalid}
    supp.EntityData.Leafs["rx-len-err"] = types.YLeaf{"RxLenErr", supp.RxLenErr}
    supp.EntityData.Leafs["rx-my-mac-err"] = types.YLeaf{"RxMyMacErr", supp.RxMyMacErr}
    supp.EntityData.Leafs["rx-total"] = types.YLeaf{"RxTotal", supp.RxTotal}
    supp.EntityData.Leafs["tx-start"] = types.YLeaf{"TxStart", supp.TxStart}
    supp.EntityData.Leafs["tx-logoff"] = types.YLeaf{"TxLogoff", supp.TxLogoff}
    supp.EntityData.Leafs["tx-resp"] = types.YLeaf{"TxResp", supp.TxResp}
    supp.EntityData.Leafs["tx-total"] = types.YLeaf{"TxTotal", supp.TxTotal}
    return &(supp.EntityData)
}

// Dot1X_Session
// Dot1x operational data
type Dot1X_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces with Dot1x.
    InterfaceSessions Dot1X_Session_InterfaceSessions
}

func (session *Dot1X_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "dot1x"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["interface-sessions"] = types.YChild{"InterfaceSessions", &session.InterfaceSessions}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(session.EntityData)
}

// Dot1X_Session_InterfaceSessions
// Interfaces with Dot1x
type Dot1X_Session_InterfaceSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dot1x Data for that Interface. The type is slice of
    // Dot1X_Session_InterfaceSessions_InterfaceSession.
    InterfaceSession []Dot1X_Session_InterfaceSessions_InterfaceSession
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetEntityData() *types.CommonEntityData {
    interfaceSessions.EntityData.YFilter = interfaceSessions.YFilter
    interfaceSessions.EntityData.YangName = "interface-sessions"
    interfaceSessions.EntityData.BundleName = "cisco_ios_xr"
    interfaceSessions.EntityData.ParentYangName = "session"
    interfaceSessions.EntityData.SegmentPath = "interface-sessions"
    interfaceSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSessions.EntityData.Children = make(map[string]types.YChild)
    interfaceSessions.EntityData.Children["interface-session"] = types.YChild{"InterfaceSession", nil}
    for i := range interfaceSessions.InterfaceSession {
        interfaceSessions.EntityData.Children[types.GetSegmentPath(&interfaceSessions.InterfaceSession[i])] = types.YChild{"InterfaceSession", &interfaceSessions.InterfaceSession[i]}
    }
    interfaceSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSessions.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession
// Dot1x Data for that Interface
type Dot1X_Session_InterfaceSessions_InterfaceSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

    // Dot1x interface Info.
    IntfInfo Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo

    // MKA session secure status.
    MkaStatusInfo Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo
}

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetEntityData() *types.CommonEntityData {
    interfaceSession.EntityData.YFilter = interfaceSession.YFilter
    interfaceSession.EntityData.YangName = "interface-session"
    interfaceSession.EntityData.BundleName = "cisco_ios_xr"
    interfaceSession.EntityData.ParentYangName = "interface-sessions"
    interfaceSession.EntityData.SegmentPath = "interface-session" + "[name='" + fmt.Sprintf("%v", interfaceSession.Name) + "']"
    interfaceSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSession.EntityData.Children = make(map[string]types.YChild)
    interfaceSession.EntityData.Children["intf-info"] = types.YChild{"IntfInfo", &interfaceSession.IntfInfo}
    interfaceSession.EntityData.Children["mka-status-info"] = types.YChild{"MkaStatusInfo", &interfaceSession.MkaStatusInfo}
    interfaceSession.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSession.EntityData.Leafs["name"] = types.YLeaf{"Name", interfaceSession.Name}
    interfaceSession.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceSession.InterfaceName}
    interfaceSession.EntityData.Leafs["interface-sname"] = types.YLeaf{"InterfaceSname", interfaceSession.InterfaceSname}
    interfaceSession.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", interfaceSession.IfHandle}
    interfaceSession.EntityData.Leafs["mac"] = types.YLeaf{"Mac", interfaceSession.Mac}
    interfaceSession.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", interfaceSession.Ethertype}
    return &(interfaceSession.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo
// Dot1x interface Info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PAE type on interface. The type is string.
    Pae interface{}

    // Dot1x Port Status. The type is string.
    PortStatus interface{}

    // Dot1x Profile. The type is string.
    Dot1XProfile interface{}

    // Dot1x Authenticator info.
    AuthInfo Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo

    // Dot1x Supplicant info.
    SuppInfo Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo
}

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetEntityData() *types.CommonEntityData {
    intfInfo.EntityData.YFilter = intfInfo.YFilter
    intfInfo.EntityData.YangName = "intf-info"
    intfInfo.EntityData.BundleName = "cisco_ios_xr"
    intfInfo.EntityData.ParentYangName = "interface-session"
    intfInfo.EntityData.SegmentPath = "intf-info"
    intfInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfInfo.EntityData.Children = make(map[string]types.YChild)
    intfInfo.EntityData.Children["auth-info"] = types.YChild{"AuthInfo", &intfInfo.AuthInfo}
    intfInfo.EntityData.Children["supp-info"] = types.YChild{"SuppInfo", &intfInfo.SuppInfo}
    intfInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    intfInfo.EntityData.Leafs["pae"] = types.YLeaf{"Pae", intfInfo.Pae}
    intfInfo.EntityData.Leafs["port-status"] = types.YLeaf{"PortStatus", intfInfo.PortStatus}
    intfInfo.EntityData.Leafs["dot1x-profile"] = types.YLeaf{"Dot1XProfile", intfInfo.Dot1XProfile}
    return &(intfInfo.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo
// Dot1x Authenticator info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Re-Authentication enabled status. The type is string.
    Reauth interface{}

    // Configuration Dependency . The type is string.
    ConfigDependency interface{}

    // Authenticator client list. The type is slice of
    // Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client.
    Client []Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetEntityData() *types.CommonEntityData {
    authInfo.EntityData.YFilter = authInfo.YFilter
    authInfo.EntityData.YangName = "auth-info"
    authInfo.EntityData.BundleName = "cisco_ios_xr"
    authInfo.EntityData.ParentYangName = "intf-info"
    authInfo.EntityData.SegmentPath = "auth-info"
    authInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authInfo.EntityData.Children = make(map[string]types.YChild)
    authInfo.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range authInfo.Client {
        authInfo.EntityData.Children[types.GetSegmentPath(&authInfo.Client[i])] = types.YChild{"Client", &authInfo.Client[i]}
    }
    authInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    authInfo.EntityData.Leafs["reauth"] = types.YLeaf{"Reauth", authInfo.Reauth}
    authInfo.EntityData.Leafs["config-dependency"] = types.YLeaf{"ConfigDependency", authInfo.ConfigDependency}
    return &(authInfo.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client
// Authenticator client list
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "auth-info"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["mac"] = types.YLeaf{"Mac", client.Mac}
    client.EntityData.Leafs["auth-sm-state"] = types.YLeaf{"AuthSmState", client.AuthSmState}
    client.EntityData.Leafs["auth-bend-sm-state"] = types.YLeaf{"AuthBendSmState", client.AuthBendSmState}
    client.EntityData.Leafs["time-to-next-reauth"] = types.YLeaf{"TimeToNextReauth", client.TimeToNextReauth}
    client.EntityData.Leafs["last-auth-time"] = types.YLeaf{"LastAuthTime", client.LastAuthTime}
    return &(client.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo
// Dot1x Supplicant info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EAP profile. The type is string.
    EapProfile interface{}

    // Configuration Dependency . The type is string.
    ConfigDependency interface{}

    // Supp Client info. The type is slice of
    // Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client.
    Client []Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetEntityData() *types.CommonEntityData {
    suppInfo.EntityData.YFilter = suppInfo.YFilter
    suppInfo.EntityData.YangName = "supp-info"
    suppInfo.EntityData.BundleName = "cisco_ios_xr"
    suppInfo.EntityData.ParentYangName = "intf-info"
    suppInfo.EntityData.SegmentPath = "supp-info"
    suppInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppInfo.EntityData.Children = make(map[string]types.YChild)
    suppInfo.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range suppInfo.Client {
        suppInfo.EntityData.Children[types.GetSegmentPath(&suppInfo.Client[i])] = types.YChild{"Client", &suppInfo.Client[i]}
    }
    suppInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    suppInfo.EntityData.Leafs["eap-profile"] = types.YLeaf{"EapProfile", suppInfo.EapProfile}
    suppInfo.EntityData.Leafs["config-dependency"] = types.YLeaf{"ConfigDependency", suppInfo.ConfigDependency}
    return &(suppInfo.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client
// Supp Client info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "supp-info"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["mac"] = types.YLeaf{"Mac", client.Mac}
    client.EntityData.Leafs["eap-method"] = types.YLeaf{"EapMethod", client.EapMethod}
    client.EntityData.Leafs["last-auth-time"] = types.YLeaf{"LastAuthTime", client.LastAuthTime}
    client.EntityData.Leafs["auth-sm-state"] = types.YLeaf{"AuthSmState", client.AuthSmState}
    client.EntityData.Leafs["auth-bend-sm-state"] = types.YLeaf{"AuthBendSmState", client.AuthBendSmState}
    return &(client.EntityData)
}

// Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo
// MKA session secure status
type Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo struct {
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

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetEntityData() *types.CommonEntityData {
    mkaStatusInfo.EntityData.YFilter = mkaStatusInfo.YFilter
    mkaStatusInfo.EntityData.YangName = "mka-status-info"
    mkaStatusInfo.EntityData.BundleName = "cisco_ios_xr"
    mkaStatusInfo.EntityData.ParentYangName = "interface-session"
    mkaStatusInfo.EntityData.SegmentPath = "mka-status-info"
    mkaStatusInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mkaStatusInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mkaStatusInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mkaStatusInfo.EntityData.Children = make(map[string]types.YChild)
    mkaStatusInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    mkaStatusInfo.EntityData.Leafs["tie-break-role"] = types.YLeaf{"TieBreakRole", mkaStatusInfo.TieBreakRole}
    mkaStatusInfo.EntityData.Leafs["eap-based-macsec"] = types.YLeaf{"EapBasedMacsec", mkaStatusInfo.EapBasedMacsec}
    mkaStatusInfo.EntityData.Leafs["mka-start-time"] = types.YLeaf{"MkaStartTime", mkaStatusInfo.MkaStartTime}
    mkaStatusInfo.EntityData.Leafs["mka-stop-time"] = types.YLeaf{"MkaStopTime", mkaStatusInfo.MkaStopTime}
    mkaStatusInfo.EntityData.Leafs["mka-response-time"] = types.YLeaf{"MkaResponseTime", mkaStatusInfo.MkaResponseTime}
    return &(mkaStatusInfo.EntityData)
}

