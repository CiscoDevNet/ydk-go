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
    parent types.Entity
    YFilter yfilter.YFilter

    // Dot1x operational data.
    Statistics Dot1X_Statistics

    // Node-specific Dot1x operational data.
    Nodes Dot1X_Nodes

    // Dot1x operational data.
    Session Dot1X_Session
}

func (dot1X *Dot1X) GetFilter() yfilter.YFilter { return dot1X.YFilter }

func (dot1X *Dot1X) SetFilter(yf yfilter.YFilter) { dot1X.YFilter = yf }

func (dot1X *Dot1X) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "nodes" { return "Nodes" }
    if yname == "session" { return "Session" }
    return ""
}

func (dot1X *Dot1X) GetSegmentPath() string {
    return "Cisco-IOS-XR-dot1x-oper:dot1x"
}

func (dot1X *Dot1X) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &dot1X.Statistics
    }
    if childYangName == "nodes" {
        return &dot1X.Nodes
    }
    if childYangName == "session" {
        return &dot1X.Session
    }
    return nil
}

func (dot1X *Dot1X) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &dot1X.Statistics
    children["nodes"] = &dot1X.Nodes
    children["session"] = &dot1X.Session
    return children
}

func (dot1X *Dot1X) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dot1X *Dot1X) GetBundleName() string { return "cisco_ios_xr" }

func (dot1X *Dot1X) GetYangName() string { return "dot1x" }

func (dot1X *Dot1X) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1X *Dot1X) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1X *Dot1X) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1X *Dot1X) SetParent(parent types.Entity) { dot1X.parent = parent }

func (dot1X *Dot1X) GetParent() types.Entity { return dot1X.parent }

func (dot1X *Dot1X) GetParentYangName() string { return "Cisco-IOS-XR-dot1x-oper" }

// Dot1X_Statistics
// Dot1x operational data
type Dot1X_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interfaces with Dot1x.
    InterfaceStatistics Dot1X_Statistics_InterfaceStatistics
}

func (statistics *Dot1X_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dot1X_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dot1X_Statistics) GetGoName(yname string) string {
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (statistics *Dot1X_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dot1X_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-statistics" {
        return &statistics.InterfaceStatistics
    }
    return nil
}

func (statistics *Dot1X_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-statistics"] = &statistics.InterfaceStatistics
    return children
}

func (statistics *Dot1X_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dot1X_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dot1X_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dot1X_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dot1X_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dot1X_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dot1X_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dot1X_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dot1X_Statistics) GetParentYangName() string { return "dot1x" }

// Dot1X_Statistics_InterfaceStatistics
// Interfaces with Dot1x
type Dot1X_Statistics_InterfaceStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dot1x Data for that Interface. The type is slice of
    // Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic.
    InterfaceStatistic []Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetFilter() yfilter.YFilter { return interfaceStatistics.YFilter }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) SetFilter(yf yfilter.YFilter) { interfaceStatistics.YFilter = yf }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetGoName(yname string) string {
    if yname == "interface-statistic" { return "InterfaceStatistic" }
    return ""
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetSegmentPath() string {
    return "interface-statistics"
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-statistic" {
        for _, c := range interfaceStatistics.InterfaceStatistic {
            if interfaceStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic{}
        interfaceStatistics.InterfaceStatistic = append(interfaceStatistics.InterfaceStatistic, child)
        return &interfaceStatistics.InterfaceStatistic[len(interfaceStatistics.InterfaceStatistic)-1]
    }
    return nil
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceStatistics.InterfaceStatistic {
        children[interfaceStatistics.InterfaceStatistic[i].GetSegmentPath()] = &interfaceStatistics.InterfaceStatistic[i]
    }
    return children
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetYangName() string { return "interface-statistics" }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) SetParent(parent types.Entity) { interfaceStatistics.parent = parent }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetParent() types.Entity { return interfaceStatistics.parent }

func (interfaceStatistics *Dot1X_Statistics_InterfaceStatistics) GetParentYangName() string { return "statistics" }

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic
// Dot1x Data for that Interface
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetFilter() yfilter.YFilter { return interfaceStatistic.YFilter }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) SetFilter(yf yfilter.YFilter) { interfaceStatistic.YFilter = yf }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "pae" { return "Pae" }
    if yname == "idb" { return "Idb" }
    if yname == "auth" { return "Auth" }
    if yname == "supp" { return "Supp" }
    return ""
}

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetSegmentPath() string {
    return "interface-statistic" + "[name='" + fmt.Sprintf("%v", interfaceStatistic.Name) + "']"
}

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "idb" {
        return &interfaceStatistic.Idb
    }
    if childYangName == "auth" {
        return &interfaceStatistic.Auth
    }
    if childYangName == "supp" {
        return &interfaceStatistic.Supp
    }
    return nil
}

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["idb"] = &interfaceStatistic.Idb
    children["auth"] = &interfaceStatistic.Auth
    children["supp"] = &interfaceStatistic.Supp
    return children
}

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = interfaceStatistic.Name
    leafs["interface-name"] = interfaceStatistic.InterfaceName
    leafs["pae"] = interfaceStatistic.Pae
    return leafs
}

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetYangName() string { return "interface-statistic" }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) SetParent(parent types.Entity) { interfaceStatistic.parent = parent }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetParent() types.Entity { return interfaceStatistic.parent }

func (interfaceStatistic *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic) GetParentYangName() string { return "interface-statistics" }

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb
// Dot1x interface database Statistics
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // NoRxOnIntfDown. The type is interface{} with range: 0..4294967295.
    NoRxOnIntfDown interface{}
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetFilter() yfilter.YFilter { return idb.YFilter }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) SetFilter(yf yfilter.YFilter) { idb.YFilter = yf }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetGoName(yname string) string {
    if yname == "rx-total" { return "RxTotal" }
    if yname == "tx-total" { return "TxTotal" }
    if yname == "no-rx-on-intf-down" { return "NoRxOnIntfDown" }
    return ""
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetSegmentPath() string {
    return "idb"
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-total"] = idb.RxTotal
    leafs["tx-total"] = idb.TxTotal
    leafs["no-rx-on-intf-down"] = idb.NoRxOnIntfDown
    return leafs
}

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetBundleName() string { return "cisco_ios_xr" }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetYangName() string { return "idb" }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) SetParent(parent types.Entity) { idb.parent = parent }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetParent() types.Entity { return idb.parent }

func (idb *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Idb) GetParentYangName() string { return "interface-statistic" }

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth
// Dot1x Authenticator Port Statistics
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth struct {
    parent types.Entity
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

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetFilter() yfilter.YFilter { return auth.YFilter }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) SetFilter(yf yfilter.YFilter) { auth.YFilter = yf }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetGoName(yname string) string {
    if yname == "rx-start" { return "RxStart" }
    if yname == "rx-logoff" { return "RxLogoff" }
    if yname == "rx-resp" { return "RxResp" }
    if yname == "rx-resp-id" { return "RxRespId" }
    if yname == "rx-invalid" { return "RxInvalid" }
    if yname == "rx-len-err" { return "RxLenErr" }
    if yname == "rx-my-mac-err" { return "RxMyMacErr" }
    if yname == "rx-total" { return "RxTotal" }
    if yname == "tx-req" { return "TxReq" }
    if yname == "tx-reqid" { return "TxReqid" }
    if yname == "tx-total" { return "TxTotal" }
    return ""
}

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetSegmentPath() string {
    return "auth"
}

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-start"] = auth.RxStart
    leafs["rx-logoff"] = auth.RxLogoff
    leafs["rx-resp"] = auth.RxResp
    leafs["rx-resp-id"] = auth.RxRespId
    leafs["rx-invalid"] = auth.RxInvalid
    leafs["rx-len-err"] = auth.RxLenErr
    leafs["rx-my-mac-err"] = auth.RxMyMacErr
    leafs["rx-total"] = auth.RxTotal
    leafs["tx-req"] = auth.TxReq
    leafs["tx-reqid"] = auth.TxReqid
    leafs["tx-total"] = auth.TxTotal
    return leafs
}

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetBundleName() string { return "cisco_ios_xr" }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetYangName() string { return "auth" }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) SetParent(parent types.Entity) { auth.parent = parent }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetParent() types.Entity { return auth.parent }

func (auth *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Auth) GetParentYangName() string { return "interface-statistic" }

// Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp
// Dot1x Supplicant Port Statistics
type Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp struct {
    parent types.Entity
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

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetFilter() yfilter.YFilter { return supp.YFilter }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) SetFilter(yf yfilter.YFilter) { supp.YFilter = yf }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetGoName(yname string) string {
    if yname == "rx-req" { return "RxReq" }
    if yname == "rx-invalid" { return "RxInvalid" }
    if yname == "rx-len-err" { return "RxLenErr" }
    if yname == "rx-my-mac-err" { return "RxMyMacErr" }
    if yname == "rx-total" { return "RxTotal" }
    if yname == "tx-start" { return "TxStart" }
    if yname == "tx-logoff" { return "TxLogoff" }
    if yname == "tx-resp" { return "TxResp" }
    if yname == "tx-total" { return "TxTotal" }
    return ""
}

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetSegmentPath() string {
    return "supp"
}

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-req"] = supp.RxReq
    leafs["rx-invalid"] = supp.RxInvalid
    leafs["rx-len-err"] = supp.RxLenErr
    leafs["rx-my-mac-err"] = supp.RxMyMacErr
    leafs["rx-total"] = supp.RxTotal
    leafs["tx-start"] = supp.TxStart
    leafs["tx-logoff"] = supp.TxLogoff
    leafs["tx-resp"] = supp.TxResp
    leafs["tx-total"] = supp.TxTotal
    return leafs
}

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetBundleName() string { return "cisco_ios_xr" }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetYangName() string { return "supp" }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) SetParent(parent types.Entity) { supp.parent = parent }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetParent() types.Entity { return supp.parent }

func (supp *Dot1X_Statistics_InterfaceStatistics_InterfaceStatistic_Supp) GetParentYangName() string { return "interface-statistic" }

// Dot1X_Nodes
// Node-specific Dot1x operational data
type Dot1X_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dot1x operational data for a particular node. The type is slice of
    // Dot1X_Nodes_Node.
    Node []Dot1X_Nodes_Node
}

func (nodes *Dot1X_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Dot1X_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Dot1X_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Dot1X_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Dot1X_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Dot1X_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Dot1X_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Dot1X_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Dot1X_Nodes) GetYangName() string { return "nodes" }

func (nodes *Dot1X_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Dot1X_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Dot1X_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Dot1X_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Dot1X_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Dot1X_Nodes) GetParentYangName() string { return "dot1x" }

// Dot1X_Nodes_Node
// Dot1x operational data for a particular node
type Dot1X_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Dot1x Default Values.
    Dot1XDefaults Dot1X_Nodes_Node_Dot1XDefaults

    // Dot1x Default Values.
    Statistics Dot1X_Nodes_Node_Statistics
}

func (node *Dot1X_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Dot1X_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Dot1X_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "dot1x-defaults" { return "Dot1XDefaults" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *Dot1X_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Dot1X_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dot1x-defaults" {
        return &node.Dot1XDefaults
    }
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *Dot1X_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dot1x-defaults"] = &node.Dot1XDefaults
    children["statistics"] = &node.Statistics
    return children
}

func (node *Dot1X_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Dot1X_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Dot1X_Nodes_Node) GetYangName() string { return "node" }

func (node *Dot1X_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Dot1X_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Dot1X_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Dot1X_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Dot1X_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Dot1X_Nodes_Node) GetParentYangName() string { return "nodes" }

// Dot1X_Nodes_Node_Dot1XDefaults
// Dot1x Default Values
type Dot1X_Nodes_Node_Dot1XDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dot1x Protocol Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Dot1x Authenticator default Timer values.
    AuthTimers Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers

    // Dot1x Supllicant default Timer values.
    SuppTimers Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetFilter() yfilter.YFilter { return dot1XDefaults.YFilter }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) SetFilter(yf yfilter.YFilter) { dot1XDefaults.YFilter = yf }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "auth-timers" { return "AuthTimers" }
    if yname == "supp-timers" { return "SuppTimers" }
    return ""
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetSegmentPath() string {
    return "dot1x-defaults"
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auth-timers" {
        return &dot1XDefaults.AuthTimers
    }
    if childYangName == "supp-timers" {
        return &dot1XDefaults.SuppTimers
    }
    return nil
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auth-timers"] = &dot1XDefaults.AuthTimers
    children["supp-timers"] = &dot1XDefaults.SuppTimers
    return children
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = dot1XDefaults.Version
    return leafs
}

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetYangName() string { return "dot1x-defaults" }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) SetParent(parent types.Entity) { dot1XDefaults.parent = parent }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetParent() types.Entity { return dot1XDefaults.parent }

func (dot1XDefaults *Dot1X_Nodes_Node_Dot1XDefaults) GetParentYangName() string { return "node" }

// Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers
// Dot1x Authenticator default Timer values
type Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers struct {
    parent types.Entity
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

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetFilter() yfilter.YFilter { return authTimers.YFilter }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) SetFilter(yf yfilter.YFilter) { authTimers.YFilter = yf }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetGoName(yname string) string {
    if yname == "quiet-period" { return "QuietPeriod" }
    if yname == "tx-period" { return "TxPeriod" }
    if yname == "max-reauth-req" { return "MaxReauthReq" }
    if yname == "supp-timeout" { return "SuppTimeout" }
    if yname == "max-req" { return "MaxReq" }
    if yname == "reauth-period" { return "ReauthPeriod" }
    return ""
}

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetSegmentPath() string {
    return "auth-timers"
}

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quiet-period"] = authTimers.QuietPeriod
    leafs["tx-period"] = authTimers.TxPeriod
    leafs["max-reauth-req"] = authTimers.MaxReauthReq
    leafs["supp-timeout"] = authTimers.SuppTimeout
    leafs["max-req"] = authTimers.MaxReq
    leafs["reauth-period"] = authTimers.ReauthPeriod
    return leafs
}

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetBundleName() string { return "cisco_ios_xr" }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetYangName() string { return "auth-timers" }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) SetParent(parent types.Entity) { authTimers.parent = parent }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetParent() types.Entity { return authTimers.parent }

func (authTimers *Dot1X_Nodes_Node_Dot1XDefaults_AuthTimers) GetParentYangName() string { return "dot1x-defaults" }

// Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers
// Dot1x Supllicant default Timer values
type Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers struct {
    parent types.Entity
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

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetFilter() yfilter.YFilter { return suppTimers.YFilter }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) SetFilter(yf yfilter.YFilter) { suppTimers.YFilter = yf }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetGoName(yname string) string {
    if yname == "auth-period" { return "AuthPeriod" }
    if yname == "held-period" { return "HeldPeriod" }
    if yname == "start-period" { return "StartPeriod" }
    if yname == "max-start" { return "MaxStart" }
    return ""
}

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetSegmentPath() string {
    return "supp-timers"
}

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auth-period"] = suppTimers.AuthPeriod
    leafs["held-period"] = suppTimers.HeldPeriod
    leafs["start-period"] = suppTimers.StartPeriod
    leafs["max-start"] = suppTimers.MaxStart
    return leafs
}

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetBundleName() string { return "cisco_ios_xr" }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetYangName() string { return "supp-timers" }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) SetParent(parent types.Entity) { suppTimers.parent = parent }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetParent() types.Entity { return suppTimers.parent }

func (suppTimers *Dot1X_Nodes_Node_Dot1XDefaults_SuppTimers) GetParentYangName() string { return "dot1x-defaults" }

// Dot1X_Nodes_Node_Statistics
// Dot1x Default Values
type Dot1X_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global statistics.
    GlStats Dot1X_Nodes_Node_Statistics_GlStats

    // dot1x interface statistics list. The type is slice of
    // Dot1X_Nodes_Node_Statistics_IfStats.
    IfStats []Dot1X_Nodes_Node_Statistics_IfStats
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dot1X_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dot1X_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "gl-stats" { return "GlStats" }
    if yname == "if-stats" { return "IfStats" }
    return ""
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "gl-stats" {
        return &statistics.GlStats
    }
    if childYangName == "if-stats" {
        for _, c := range statistics.IfStats {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Nodes_Node_Statistics_IfStats{}
        statistics.IfStats = append(statistics.IfStats, child)
        return &statistics.IfStats[len(statistics.IfStats)-1]
    }
    return nil
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["gl-stats"] = &statistics.GlStats
    for i := range statistics.IfStats {
        children[statistics.IfStats[i].GetSegmentPath()] = &statistics.IfStats[i]
    }
    return children
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Dot1X_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dot1X_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dot1X_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dot1X_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dot1X_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dot1X_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dot1X_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dot1X_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Dot1X_Nodes_Node_Statistics_GlStats
// Global statistics
type Dot1X_Nodes_Node_Statistics_GlStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // RxNoIDB. The type is interface{} with range: 0..4294967295.
    RxNoIdb interface{}
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetFilter() yfilter.YFilter { return glStats.YFilter }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) SetFilter(yf yfilter.YFilter) { glStats.YFilter = yf }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetGoName(yname string) string {
    if yname == "tx-total" { return "TxTotal" }
    if yname == "rx-total" { return "RxTotal" }
    if yname == "rx-no-idb" { return "RxNoIdb" }
    return ""
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetSegmentPath() string {
    return "gl-stats"
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-total"] = glStats.TxTotal
    leafs["rx-total"] = glStats.RxTotal
    leafs["rx-no-idb"] = glStats.RxNoIdb
    return leafs
}

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetBundleName() string { return "cisco_ios_xr" }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetYangName() string { return "gl-stats" }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) SetParent(parent types.Entity) { glStats.parent = parent }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetParent() types.Entity { return glStats.parent }

func (glStats *Dot1X_Nodes_Node_Statistics_GlStats) GetParentYangName() string { return "statistics" }

// Dot1X_Nodes_Node_Statistics_IfStats
// dot1x interface statistics list
type Dot1X_Nodes_Node_Statistics_IfStats struct {
    parent types.Entity
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

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetFilter() yfilter.YFilter { return ifStats.YFilter }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) SetFilter(yf yfilter.YFilter) { ifStats.YFilter = yf }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "pae" { return "Pae" }
    if yname == "idb" { return "Idb" }
    if yname == "auth" { return "Auth" }
    if yname == "supp" { return "Supp" }
    return ""
}

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetSegmentPath() string {
    return "if-stats"
}

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "idb" {
        return &ifStats.Idb
    }
    if childYangName == "auth" {
        return &ifStats.Auth
    }
    if childYangName == "supp" {
        return &ifStats.Supp
    }
    return nil
}

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["idb"] = &ifStats.Idb
    children["auth"] = &ifStats.Auth
    children["supp"] = &ifStats.Supp
    return children
}

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = ifStats.InterfaceName
    leafs["pae"] = ifStats.Pae
    return leafs
}

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetBundleName() string { return "cisco_ios_xr" }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetYangName() string { return "if-stats" }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) SetParent(parent types.Entity) { ifStats.parent = parent }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetParent() types.Entity { return ifStats.parent }

func (ifStats *Dot1X_Nodes_Node_Statistics_IfStats) GetParentYangName() string { return "statistics" }

// Dot1X_Nodes_Node_Statistics_IfStats_Idb
// Dot1x interface database Statistics
type Dot1X_Nodes_Node_Statistics_IfStats_Idb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RxTotal. The type is interface{} with range: 0..4294967295.
    RxTotal interface{}

    // TxTotal. The type is interface{} with range: 0..4294967295.
    TxTotal interface{}

    // NoRxOnIntfDown. The type is interface{} with range: 0..4294967295.
    NoRxOnIntfDown interface{}
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetFilter() yfilter.YFilter { return idb.YFilter }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) SetFilter(yf yfilter.YFilter) { idb.YFilter = yf }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetGoName(yname string) string {
    if yname == "rx-total" { return "RxTotal" }
    if yname == "tx-total" { return "TxTotal" }
    if yname == "no-rx-on-intf-down" { return "NoRxOnIntfDown" }
    return ""
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetSegmentPath() string {
    return "idb"
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-total"] = idb.RxTotal
    leafs["tx-total"] = idb.TxTotal
    leafs["no-rx-on-intf-down"] = idb.NoRxOnIntfDown
    return leafs
}

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetBundleName() string { return "cisco_ios_xr" }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetYangName() string { return "idb" }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) SetParent(parent types.Entity) { idb.parent = parent }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetParent() types.Entity { return idb.parent }

func (idb *Dot1X_Nodes_Node_Statistics_IfStats_Idb) GetParentYangName() string { return "if-stats" }

// Dot1X_Nodes_Node_Statistics_IfStats_Auth
// Dot1x Authenticator Port Statistics
type Dot1X_Nodes_Node_Statistics_IfStats_Auth struct {
    parent types.Entity
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

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetFilter() yfilter.YFilter { return auth.YFilter }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) SetFilter(yf yfilter.YFilter) { auth.YFilter = yf }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetGoName(yname string) string {
    if yname == "rx-start" { return "RxStart" }
    if yname == "rx-logoff" { return "RxLogoff" }
    if yname == "rx-resp" { return "RxResp" }
    if yname == "rx-resp-id" { return "RxRespId" }
    if yname == "rx-invalid" { return "RxInvalid" }
    if yname == "rx-len-err" { return "RxLenErr" }
    if yname == "rx-my-mac-err" { return "RxMyMacErr" }
    if yname == "rx-total" { return "RxTotal" }
    if yname == "tx-req" { return "TxReq" }
    if yname == "tx-reqid" { return "TxReqid" }
    if yname == "tx-total" { return "TxTotal" }
    return ""
}

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetSegmentPath() string {
    return "auth"
}

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-start"] = auth.RxStart
    leafs["rx-logoff"] = auth.RxLogoff
    leafs["rx-resp"] = auth.RxResp
    leafs["rx-resp-id"] = auth.RxRespId
    leafs["rx-invalid"] = auth.RxInvalid
    leafs["rx-len-err"] = auth.RxLenErr
    leafs["rx-my-mac-err"] = auth.RxMyMacErr
    leafs["rx-total"] = auth.RxTotal
    leafs["tx-req"] = auth.TxReq
    leafs["tx-reqid"] = auth.TxReqid
    leafs["tx-total"] = auth.TxTotal
    return leafs
}

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetBundleName() string { return "cisco_ios_xr" }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetYangName() string { return "auth" }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) SetParent(parent types.Entity) { auth.parent = parent }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetParent() types.Entity { return auth.parent }

func (auth *Dot1X_Nodes_Node_Statistics_IfStats_Auth) GetParentYangName() string { return "if-stats" }

// Dot1X_Nodes_Node_Statistics_IfStats_Supp
// Dot1x Supplicant Port Statistics
type Dot1X_Nodes_Node_Statistics_IfStats_Supp struct {
    parent types.Entity
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

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetFilter() yfilter.YFilter { return supp.YFilter }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) SetFilter(yf yfilter.YFilter) { supp.YFilter = yf }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetGoName(yname string) string {
    if yname == "rx-req" { return "RxReq" }
    if yname == "rx-invalid" { return "RxInvalid" }
    if yname == "rx-len-err" { return "RxLenErr" }
    if yname == "rx-my-mac-err" { return "RxMyMacErr" }
    if yname == "rx-total" { return "RxTotal" }
    if yname == "tx-start" { return "TxStart" }
    if yname == "tx-logoff" { return "TxLogoff" }
    if yname == "tx-resp" { return "TxResp" }
    if yname == "tx-total" { return "TxTotal" }
    return ""
}

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetSegmentPath() string {
    return "supp"
}

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-req"] = supp.RxReq
    leafs["rx-invalid"] = supp.RxInvalid
    leafs["rx-len-err"] = supp.RxLenErr
    leafs["rx-my-mac-err"] = supp.RxMyMacErr
    leafs["rx-total"] = supp.RxTotal
    leafs["tx-start"] = supp.TxStart
    leafs["tx-logoff"] = supp.TxLogoff
    leafs["tx-resp"] = supp.TxResp
    leafs["tx-total"] = supp.TxTotal
    return leafs
}

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetBundleName() string { return "cisco_ios_xr" }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetYangName() string { return "supp" }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) SetParent(parent types.Entity) { supp.parent = parent }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetParent() types.Entity { return supp.parent }

func (supp *Dot1X_Nodes_Node_Statistics_IfStats_Supp) GetParentYangName() string { return "if-stats" }

// Dot1X_Session
// Dot1x operational data
type Dot1X_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interfaces with Dot1x.
    InterfaceSessions Dot1X_Session_InterfaceSessions
}

func (session *Dot1X_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Dot1X_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Dot1X_Session) GetGoName(yname string) string {
    if yname == "interface-sessions" { return "InterfaceSessions" }
    return ""
}

func (session *Dot1X_Session) GetSegmentPath() string {
    return "session"
}

func (session *Dot1X_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-sessions" {
        return &session.InterfaceSessions
    }
    return nil
}

func (session *Dot1X_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-sessions"] = &session.InterfaceSessions
    return children
}

func (session *Dot1X_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *Dot1X_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Dot1X_Session) GetYangName() string { return "session" }

func (session *Dot1X_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Dot1X_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Dot1X_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Dot1X_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Dot1X_Session) GetParent() types.Entity { return session.parent }

func (session *Dot1X_Session) GetParentYangName() string { return "dot1x" }

// Dot1X_Session_InterfaceSessions
// Interfaces with Dot1x
type Dot1X_Session_InterfaceSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dot1x Data for that Interface. The type is slice of
    // Dot1X_Session_InterfaceSessions_InterfaceSession.
    InterfaceSession []Dot1X_Session_InterfaceSessions_InterfaceSession
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetFilter() yfilter.YFilter { return interfaceSessions.YFilter }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) SetFilter(yf yfilter.YFilter) { interfaceSessions.YFilter = yf }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetGoName(yname string) string {
    if yname == "interface-session" { return "InterfaceSession" }
    return ""
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetSegmentPath() string {
    return "interface-sessions"
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-session" {
        for _, c := range interfaceSessions.InterfaceSession {
            if interfaceSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Session_InterfaceSessions_InterfaceSession{}
        interfaceSessions.InterfaceSession = append(interfaceSessions.InterfaceSession, child)
        return &interfaceSessions.InterfaceSession[len(interfaceSessions.InterfaceSession)-1]
    }
    return nil
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceSessions.InterfaceSession {
        children[interfaceSessions.InterfaceSession[i].GetSegmentPath()] = &interfaceSessions.InterfaceSession[i]
    }
    return children
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetYangName() string { return "interface-sessions" }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) SetParent(parent types.Entity) { interfaceSessions.parent = parent }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetParent() types.Entity { return interfaceSessions.parent }

func (interfaceSessions *Dot1X_Session_InterfaceSessions) GetParentYangName() string { return "session" }

// Dot1X_Session_InterfaceSessions_InterfaceSession
// Dot1x Data for that Interface
type Dot1X_Session_InterfaceSessions_InterfaceSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetFilter() yfilter.YFilter { return interfaceSession.YFilter }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) SetFilter(yf yfilter.YFilter) { interfaceSession.YFilter = yf }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-sname" { return "InterfaceSname" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "mac" { return "Mac" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "intf-info" { return "IntfInfo" }
    if yname == "mka-status-info" { return "MkaStatusInfo" }
    return ""
}

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetSegmentPath() string {
    return "interface-session" + "[name='" + fmt.Sprintf("%v", interfaceSession.Name) + "']"
}

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intf-info" {
        return &interfaceSession.IntfInfo
    }
    if childYangName == "mka-status-info" {
        return &interfaceSession.MkaStatusInfo
    }
    return nil
}

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intf-info"] = &interfaceSession.IntfInfo
    children["mka-status-info"] = &interfaceSession.MkaStatusInfo
    return children
}

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = interfaceSession.Name
    leafs["interface-name"] = interfaceSession.InterfaceName
    leafs["interface-sname"] = interfaceSession.InterfaceSname
    leafs["if-handle"] = interfaceSession.IfHandle
    leafs["mac"] = interfaceSession.Mac
    leafs["ethertype"] = interfaceSession.Ethertype
    return leafs
}

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetYangName() string { return "interface-session" }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) SetParent(parent types.Entity) { interfaceSession.parent = parent }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetParent() types.Entity { return interfaceSession.parent }

func (interfaceSession *Dot1X_Session_InterfaceSessions_InterfaceSession) GetParentYangName() string { return "interface-sessions" }

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo
// Dot1x interface Info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo struct {
    parent types.Entity
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

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetFilter() yfilter.YFilter { return intfInfo.YFilter }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) SetFilter(yf yfilter.YFilter) { intfInfo.YFilter = yf }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetGoName(yname string) string {
    if yname == "pae" { return "Pae" }
    if yname == "port-status" { return "PortStatus" }
    if yname == "dot1x-profile" { return "Dot1XProfile" }
    if yname == "auth-info" { return "AuthInfo" }
    if yname == "supp-info" { return "SuppInfo" }
    return ""
}

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetSegmentPath() string {
    return "intf-info"
}

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auth-info" {
        return &intfInfo.AuthInfo
    }
    if childYangName == "supp-info" {
        return &intfInfo.SuppInfo
    }
    return nil
}

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auth-info"] = &intfInfo.AuthInfo
    children["supp-info"] = &intfInfo.SuppInfo
    return children
}

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pae"] = intfInfo.Pae
    leafs["port-status"] = intfInfo.PortStatus
    leafs["dot1x-profile"] = intfInfo.Dot1XProfile
    return leafs
}

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetBundleName() string { return "cisco_ios_xr" }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetYangName() string { return "intf-info" }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) SetParent(parent types.Entity) { intfInfo.parent = parent }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetParent() types.Entity { return intfInfo.parent }

func (intfInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo) GetParentYangName() string { return "interface-session" }

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo
// Dot1x Authenticator info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Re-Authentication enabled status. The type is string.
    Reauth interface{}

    // Configuration Dependency . The type is string.
    ConfigDependency interface{}

    // Authenticator client list. The type is slice of
    // Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client.
    Client []Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetFilter() yfilter.YFilter { return authInfo.YFilter }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) SetFilter(yf yfilter.YFilter) { authInfo.YFilter = yf }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetGoName(yname string) string {
    if yname == "reauth" { return "Reauth" }
    if yname == "config-dependency" { return "ConfigDependency" }
    if yname == "client" { return "Client" }
    return ""
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetSegmentPath() string {
    return "auth-info"
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range authInfo.Client {
            if authInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client{}
        authInfo.Client = append(authInfo.Client, child)
        return &authInfo.Client[len(authInfo.Client)-1]
    }
    return nil
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authInfo.Client {
        children[authInfo.Client[i].GetSegmentPath()] = &authInfo.Client[i]
    }
    return children
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reauth"] = authInfo.Reauth
    leafs["config-dependency"] = authInfo.ConfigDependency
    return leafs
}

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetBundleName() string { return "cisco_ios_xr" }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetYangName() string { return "auth-info" }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) SetParent(parent types.Entity) { authInfo.parent = parent }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetParent() types.Entity { return authInfo.parent }

func (authInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo) GetParentYangName() string { return "intf-info" }

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client
// Authenticator client list
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client struct {
    parent types.Entity
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

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    if yname == "auth-sm-state" { return "AuthSmState" }
    if yname == "auth-bend-sm-state" { return "AuthBendSmState" }
    if yname == "time-to-next-reauth" { return "TimeToNextReauth" }
    if yname == "last-auth-time" { return "LastAuthTime" }
    return ""
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetSegmentPath() string {
    return "client"
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac"] = client.Mac
    leafs["auth-sm-state"] = client.AuthSmState
    leafs["auth-bend-sm-state"] = client.AuthBendSmState
    leafs["time-to-next-reauth"] = client.TimeToNextReauth
    leafs["last-auth-time"] = client.LastAuthTime
    return leafs
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetYangName() string { return "client" }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetParent() types.Entity { return client.parent }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_AuthInfo_Client) GetParentYangName() string { return "auth-info" }

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo
// Dot1x Supplicant info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EAP profile. The type is string.
    EapProfile interface{}

    // Configuration Dependency . The type is string.
    ConfigDependency interface{}

    // Supp Client info. The type is slice of
    // Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client.
    Client []Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetFilter() yfilter.YFilter { return suppInfo.YFilter }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) SetFilter(yf yfilter.YFilter) { suppInfo.YFilter = yf }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetGoName(yname string) string {
    if yname == "eap-profile" { return "EapProfile" }
    if yname == "config-dependency" { return "ConfigDependency" }
    if yname == "client" { return "Client" }
    return ""
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetSegmentPath() string {
    return "supp-info"
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range suppInfo.Client {
            if suppInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client{}
        suppInfo.Client = append(suppInfo.Client, child)
        return &suppInfo.Client[len(suppInfo.Client)-1]
    }
    return nil
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range suppInfo.Client {
        children[suppInfo.Client[i].GetSegmentPath()] = &suppInfo.Client[i]
    }
    return children
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["eap-profile"] = suppInfo.EapProfile
    leafs["config-dependency"] = suppInfo.ConfigDependency
    return leafs
}

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetBundleName() string { return "cisco_ios_xr" }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetYangName() string { return "supp-info" }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) SetParent(parent types.Entity) { suppInfo.parent = parent }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetParent() types.Entity { return suppInfo.parent }

func (suppInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo) GetParentYangName() string { return "intf-info" }

// Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client
// Supp Client info
type Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client struct {
    parent types.Entity
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

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    if yname == "eap-method" { return "EapMethod" }
    if yname == "last-auth-time" { return "LastAuthTime" }
    if yname == "auth-sm-state" { return "AuthSmState" }
    if yname == "auth-bend-sm-state" { return "AuthBendSmState" }
    return ""
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetSegmentPath() string {
    return "client"
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac"] = client.Mac
    leafs["eap-method"] = client.EapMethod
    leafs["last-auth-time"] = client.LastAuthTime
    leafs["auth-sm-state"] = client.AuthSmState
    leafs["auth-bend-sm-state"] = client.AuthBendSmState
    return leafs
}

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetYangName() string { return "client" }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetParent() types.Entity { return client.parent }

func (client *Dot1X_Session_InterfaceSessions_InterfaceSession_IntfInfo_SuppInfo_Client) GetParentYangName() string { return "supp-info" }

// Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo
// MKA session secure status
type Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo struct {
    parent types.Entity
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

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetFilter() yfilter.YFilter { return mkaStatusInfo.YFilter }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) SetFilter(yf yfilter.YFilter) { mkaStatusInfo.YFilter = yf }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetGoName(yname string) string {
    if yname == "tie-break-role" { return "TieBreakRole" }
    if yname == "eap-based-macsec" { return "EapBasedMacsec" }
    if yname == "mka-start-time" { return "MkaStartTime" }
    if yname == "mka-stop-time" { return "MkaStopTime" }
    if yname == "mka-response-time" { return "MkaResponseTime" }
    return ""
}

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetSegmentPath() string {
    return "mka-status-info"
}

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tie-break-role"] = mkaStatusInfo.TieBreakRole
    leafs["eap-based-macsec"] = mkaStatusInfo.EapBasedMacsec
    leafs["mka-start-time"] = mkaStatusInfo.MkaStartTime
    leafs["mka-stop-time"] = mkaStatusInfo.MkaStopTime
    leafs["mka-response-time"] = mkaStatusInfo.MkaResponseTime
    return leafs
}

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetBundleName() string { return "cisco_ios_xr" }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetYangName() string { return "mka-status-info" }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) SetParent(parent types.Entity) { mkaStatusInfo.parent = parent }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetParent() types.Entity { return mkaStatusInfo.parent }

func (mkaStatusInfo *Dot1X_Session_InterfaceSessions_InterfaceSession_MkaStatusInfo) GetParentYangName() string { return "interface-session" }

