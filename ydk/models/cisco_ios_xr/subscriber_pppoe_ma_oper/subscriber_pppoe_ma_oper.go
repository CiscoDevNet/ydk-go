// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-pppoe-ma package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pppoe: PPPoE operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_pppoe_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_pppoe_ma_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-pppoe-ma-oper pppoe}", reflect.TypeOf(Pppoe{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-pppoe-ma-oper:pppoe", reflect.TypeOf(Pppoe{}))
}

// PppoeMaThrottleState represents Pppoe ma throttle state
type PppoeMaThrottleState string

const (
    // Idle State
    PppoeMaThrottleState_idle PppoeMaThrottleState = "idle"

    // Monitor State
    PppoeMaThrottleState_monitor PppoeMaThrottleState = "monitor"

    // Block State
    PppoeMaThrottleState_block PppoeMaThrottleState = "block"
)

// PppoeMaLimitState represents Pppoe ma limit state
type PppoeMaLimitState string

const (
    // OK State
    PppoeMaLimitState_ok PppoeMaLimitState = "ok"

    // Warn State
    PppoeMaLimitState_warning PppoeMaLimitState = "warning"

    // Block State
    PppoeMaLimitState_block PppoeMaLimitState = "block"
)

// PppoeMaSessionIdbSrgState represents Pppoe ma session idb srg state
type PppoeMaSessionIdbSrgState string

const (
    // SRG-None state
    PppoeMaSessionIdbSrgState_none PppoeMaSessionIdbSrgState = "none"

    // SRG-Active state
    PppoeMaSessionIdbSrgState_active PppoeMaSessionIdbSrgState = "active"

    // SRG-Standby state
    PppoeMaSessionIdbSrgState_standby PppoeMaSessionIdbSrgState = "standby"
)

// Pppoe
// PPPoE operational data
type Pppoe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE access interface statistics information.
    AccessInterfaceStatistics Pppoe_AccessInterfaceStatistics

    // Per-node PPPoE operational data.
    Nodes Pppoe_Nodes
}

func (pppoe *Pppoe) GetFilter() yfilter.YFilter { return pppoe.YFilter }

func (pppoe *Pppoe) SetFilter(yf yfilter.YFilter) { pppoe.YFilter = yf }

func (pppoe *Pppoe) GetGoName(yname string) string {
    if yname == "access-interface-statistics" { return "AccessInterfaceStatistics" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (pppoe *Pppoe) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-pppoe-ma-oper:pppoe"
}

func (pppoe *Pppoe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface-statistics" {
        return &pppoe.AccessInterfaceStatistics
    }
    if childYangName == "nodes" {
        return &pppoe.Nodes
    }
    return nil
}

func (pppoe *Pppoe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-interface-statistics"] = &pppoe.AccessInterfaceStatistics
    children["nodes"] = &pppoe.Nodes
    return children
}

func (pppoe *Pppoe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pppoe *Pppoe) GetBundleName() string { return "cisco_ios_xr" }

func (pppoe *Pppoe) GetYangName() string { return "pppoe" }

func (pppoe *Pppoe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoe *Pppoe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoe *Pppoe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoe *Pppoe) SetParent(parent types.Entity) { pppoe.parent = parent }

func (pppoe *Pppoe) GetParent() types.Entity { return pppoe.parent }

func (pppoe *Pppoe) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-pppoe-ma-oper" }

// Pppoe_AccessInterfaceStatistics
// PPPoE access interface statistics information
type Pppoe_AccessInterfaceStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics information for a PPPoE-enabled access interface. The type is
    // slice of Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic.
    AccessInterfaceStatistic []Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetFilter() yfilter.YFilter { return accessInterfaceStatistics.YFilter }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) SetFilter(yf yfilter.YFilter) { accessInterfaceStatistics.YFilter = yf }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetGoName(yname string) string {
    if yname == "access-interface-statistic" { return "AccessInterfaceStatistic" }
    return ""
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetSegmentPath() string {
    return "access-interface-statistics"
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface-statistic" {
        for _, c := range accessInterfaceStatistics.AccessInterfaceStatistic {
            if accessInterfaceStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic{}
        accessInterfaceStatistics.AccessInterfaceStatistic = append(accessInterfaceStatistics.AccessInterfaceStatistic, child)
        return &accessInterfaceStatistics.AccessInterfaceStatistic[len(accessInterfaceStatistics.AccessInterfaceStatistic)-1]
    }
    return nil
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessInterfaceStatistics.AccessInterfaceStatistic {
        children[accessInterfaceStatistics.AccessInterfaceStatistic[i].GetSegmentPath()] = &accessInterfaceStatistics.AccessInterfaceStatistic[i]
    }
    return children
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetYangName() string { return "access-interface-statistics" }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) SetParent(parent types.Entity) { accessInterfaceStatistics.parent = parent }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetParent() types.Entity { return accessInterfaceStatistics.parent }

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetParentYangName() string { return "pppoe" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic
// Statistics information for a PPPoE-enabled
// access interface
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. PPPoE Access Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Packet Counts.
    PacketCounts Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetFilter() yfilter.YFilter { return accessInterfaceStatistic.YFilter }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) SetFilter(yf yfilter.YFilter) { accessInterfaceStatistic.YFilter = yf }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "packet-counts" { return "PacketCounts" }
    return ""
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetSegmentPath() string {
    return "access-interface-statistic" + "[interface-name='" + fmt.Sprintf("%v", accessInterfaceStatistic.InterfaceName) + "']"
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packet-counts" {
        return &accessInterfaceStatistic.PacketCounts
    }
    return nil
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-counts"] = &accessInterfaceStatistic.PacketCounts
    return children
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = accessInterfaceStatistic.InterfaceName
    return leafs
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetYangName() string { return "access-interface-statistic" }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) SetParent(parent types.Entity) { accessInterfaceStatistic.parent = parent }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetParent() types.Entity { return accessInterfaceStatistic.parent }

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetParentYangName() string { return "access-interface-statistics" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts
// Packet Counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PADI counts.
    Padi Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi

    // PADO counts.
    Pado Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado

    // PADR counts.
    Padr Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr

    // PADS Success counts.
    PadsSuccess Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess

    // PADS Error counts.
    PadsError Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError

    // PADT counts.
    Padt Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt

    // Session Stage counts.
    SessionState Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState

    // Other counts.
    Other Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other
}

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetFilter() yfilter.YFilter { return packetCounts.YFilter }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) SetFilter(yf yfilter.YFilter) { packetCounts.YFilter = yf }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetGoName(yname string) string {
    if yname == "padi" { return "Padi" }
    if yname == "pado" { return "Pado" }
    if yname == "padr" { return "Padr" }
    if yname == "pads-success" { return "PadsSuccess" }
    if yname == "pads-error" { return "PadsError" }
    if yname == "padt" { return "Padt" }
    if yname == "session-state" { return "SessionState" }
    if yname == "other" { return "Other" }
    return ""
}

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetSegmentPath() string {
    return "packet-counts"
}

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "padi" {
        return &packetCounts.Padi
    }
    if childYangName == "pado" {
        return &packetCounts.Pado
    }
    if childYangName == "padr" {
        return &packetCounts.Padr
    }
    if childYangName == "pads-success" {
        return &packetCounts.PadsSuccess
    }
    if childYangName == "pads-error" {
        return &packetCounts.PadsError
    }
    if childYangName == "padt" {
        return &packetCounts.Padt
    }
    if childYangName == "session-state" {
        return &packetCounts.SessionState
    }
    if childYangName == "other" {
        return &packetCounts.Other
    }
    return nil
}

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["padi"] = &packetCounts.Padi
    children["pado"] = &packetCounts.Pado
    children["padr"] = &packetCounts.Padr
    children["pads-success"] = &packetCounts.PadsSuccess
    children["pads-error"] = &packetCounts.PadsError
    children["padt"] = &packetCounts.Padt
    children["session-state"] = &packetCounts.SessionState
    children["other"] = &packetCounts.Other
    return children
}

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetBundleName() string { return "cisco_ios_xr" }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetYangName() string { return "packet-counts" }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) SetParent(parent types.Entity) { packetCounts.parent = parent }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetParent() types.Entity { return packetCounts.parent }

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetParentYangName() string { return "access-interface-statistic" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi
// PADI counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetFilter() yfilter.YFilter { return padi.YFilter }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) SetFilter(yf yfilter.YFilter) { padi.YFilter = yf }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetSegmentPath() string {
    return "padi"
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padi.Sent
    leafs["received"] = padi.Received
    leafs["dropped"] = padi.Dropped
    return leafs
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetBundleName() string { return "cisco_ios_xr" }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetYangName() string { return "padi" }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) SetParent(parent types.Entity) { padi.parent = parent }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetParent() types.Entity { return padi.parent }

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado
// PADO counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetFilter() yfilter.YFilter { return pado.YFilter }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) SetFilter(yf yfilter.YFilter) { pado.YFilter = yf }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetSegmentPath() string {
    return "pado"
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = pado.Sent
    leafs["received"] = pado.Received
    leafs["dropped"] = pado.Dropped
    return leafs
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetBundleName() string { return "cisco_ios_xr" }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetYangName() string { return "pado" }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) SetParent(parent types.Entity) { pado.parent = parent }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetParent() types.Entity { return pado.parent }

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr
// PADR counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetFilter() yfilter.YFilter { return padr.YFilter }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) SetFilter(yf yfilter.YFilter) { padr.YFilter = yf }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetSegmentPath() string {
    return "padr"
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padr.Sent
    leafs["received"] = padr.Received
    leafs["dropped"] = padr.Dropped
    return leafs
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetBundleName() string { return "cisco_ios_xr" }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetYangName() string { return "padr" }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) SetParent(parent types.Entity) { padr.parent = parent }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetParent() types.Entity { return padr.parent }

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess
// PADS Success counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetFilter() yfilter.YFilter { return padsSuccess.YFilter }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) SetFilter(yf yfilter.YFilter) { padsSuccess.YFilter = yf }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetSegmentPath() string {
    return "pads-success"
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padsSuccess.Sent
    leafs["received"] = padsSuccess.Received
    leafs["dropped"] = padsSuccess.Dropped
    return leafs
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetBundleName() string { return "cisco_ios_xr" }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetYangName() string { return "pads-success" }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) SetParent(parent types.Entity) { padsSuccess.parent = parent }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetParent() types.Entity { return padsSuccess.parent }

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError
// PADS Error counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetFilter() yfilter.YFilter { return padsError.YFilter }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) SetFilter(yf yfilter.YFilter) { padsError.YFilter = yf }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetSegmentPath() string {
    return "pads-error"
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padsError.Sent
    leafs["received"] = padsError.Received
    leafs["dropped"] = padsError.Dropped
    return leafs
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetBundleName() string { return "cisco_ios_xr" }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetYangName() string { return "pads-error" }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) SetParent(parent types.Entity) { padsError.parent = parent }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetParent() types.Entity { return padsError.parent }

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt
// PADT counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetFilter() yfilter.YFilter { return padt.YFilter }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) SetFilter(yf yfilter.YFilter) { padt.YFilter = yf }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetSegmentPath() string {
    return "padt"
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padt.Sent
    leafs["received"] = padt.Received
    leafs["dropped"] = padt.Dropped
    return leafs
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetBundleName() string { return "cisco_ios_xr" }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetYangName() string { return "padt" }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) SetParent(parent types.Entity) { padt.parent = parent }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetParent() types.Entity { return padt.parent }

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState
// Session Stage counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetFilter() yfilter.YFilter { return sessionState.YFilter }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) SetFilter(yf yfilter.YFilter) { sessionState.YFilter = yf }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetSegmentPath() string {
    return "session-state"
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = sessionState.Sent
    leafs["received"] = sessionState.Received
    leafs["dropped"] = sessionState.Dropped
    return leafs
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetBundleName() string { return "cisco_ios_xr" }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetYangName() string { return "session-state" }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) SetParent(parent types.Entity) { sessionState.parent = parent }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetParent() types.Entity { return sessionState.parent }

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetParentYangName() string { return "packet-counts" }

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other
// Other counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetFilter() yfilter.YFilter { return other.YFilter }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) SetFilter(yf yfilter.YFilter) { other.YFilter = yf }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetSegmentPath() string {
    return "other"
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = other.Sent
    leafs["received"] = other.Received
    leafs["dropped"] = other.Dropped
    return leafs
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetBundleName() string { return "cisco_ios_xr" }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetYangName() string { return "other" }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) SetParent(parent types.Entity) { other.parent = parent }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetParent() types.Entity { return other.parent }

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes
// Per-node PPPoE operational data
type Pppoe_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE operational data for a particular node. The type is slice of
    // Pppoe_Nodes_Node.
    Node []Pppoe_Nodes_Node
}

func (nodes *Pppoe_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Pppoe_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Pppoe_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Pppoe_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Pppoe_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Pppoe_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Pppoe_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Pppoe_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Pppoe_Nodes) GetYangName() string { return "nodes" }

func (nodes *Pppoe_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Pppoe_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Pppoe_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Pppoe_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Pppoe_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Pppoe_Nodes) GetParentYangName() string { return "pppoe" }

// Pppoe_Nodes_Node
// PPPoE operational data for a particular node
type Pppoe_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // PPPoE statistics for a given node.
    Statistics Pppoe_Nodes_Node_Statistics

    // PPPoE access interface information.
    AccessInterface Pppoe_Nodes_Node_AccessInterface

    // Per interface PPPoE operational data.
    Interfaces Pppoe_Nodes_Node_Interfaces

    // PPPoE BBA-Group information.
    BbaGroups Pppoe_Nodes_Node_BbaGroups

    // PPPoE statistics for a given node.
    SummaryTotal Pppoe_Nodes_Node_SummaryTotal
}

func (node *Pppoe_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Pppoe_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Pppoe_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statistics" { return "Statistics" }
    if yname == "access-interface" { return "AccessInterface" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "bba-groups" { return "BbaGroups" }
    if yname == "summary-total" { return "SummaryTotal" }
    return ""
}

func (node *Pppoe_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Pppoe_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    if childYangName == "access-interface" {
        return &node.AccessInterface
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "bba-groups" {
        return &node.BbaGroups
    }
    if childYangName == "summary-total" {
        return &node.SummaryTotal
    }
    return nil
}

func (node *Pppoe_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    children["access-interface"] = &node.AccessInterface
    children["interfaces"] = &node.Interfaces
    children["bba-groups"] = &node.BbaGroups
    children["summary-total"] = &node.SummaryTotal
    return children
}

func (node *Pppoe_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Pppoe_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Pppoe_Nodes_Node) GetYangName() string { return "node" }

func (node *Pppoe_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Pppoe_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Pppoe_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Pppoe_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Pppoe_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Pppoe_Nodes_Node) GetParentYangName() string { return "nodes" }

// Pppoe_Nodes_Node_Statistics
// PPPoE statistics for a given node
type Pppoe_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet Counts.
    PacketCounts Pppoe_Nodes_Node_Statistics_PacketCounts

    // Packet Error Counts.
    PacketErrorCounts Pppoe_Nodes_Node_Statistics_PacketErrorCounts
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Pppoe_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Pppoe_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "packet-counts" { return "PacketCounts" }
    if yname == "packet-error-counts" { return "PacketErrorCounts" }
    return ""
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packet-counts" {
        return &statistics.PacketCounts
    }
    if childYangName == "packet-error-counts" {
        return &statistics.PacketErrorCounts
    }
    return nil
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-counts"] = &statistics.PacketCounts
    children["packet-error-counts"] = &statistics.PacketErrorCounts
    return children
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Pppoe_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Pppoe_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Pppoe_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Pppoe_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Pppoe_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Pppoe_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Pppoe_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Pppoe_Nodes_Node_Statistics_PacketCounts
// Packet Counts
type Pppoe_Nodes_Node_Statistics_PacketCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PADI counts.
    Padi Pppoe_Nodes_Node_Statistics_PacketCounts_Padi

    // PADO counts.
    Pado Pppoe_Nodes_Node_Statistics_PacketCounts_Pado

    // PADR counts.
    Padr Pppoe_Nodes_Node_Statistics_PacketCounts_Padr

    // PADS Success counts.
    PadsSuccess Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess

    // PADS Error counts.
    PadsError Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError

    // PADT counts.
    Padt Pppoe_Nodes_Node_Statistics_PacketCounts_Padt

    // Session Stage counts.
    SessionState Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState

    // Other counts.
    Other Pppoe_Nodes_Node_Statistics_PacketCounts_Other
}

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetFilter() yfilter.YFilter { return packetCounts.YFilter }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) SetFilter(yf yfilter.YFilter) { packetCounts.YFilter = yf }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetGoName(yname string) string {
    if yname == "padi" { return "Padi" }
    if yname == "pado" { return "Pado" }
    if yname == "padr" { return "Padr" }
    if yname == "pads-success" { return "PadsSuccess" }
    if yname == "pads-error" { return "PadsError" }
    if yname == "padt" { return "Padt" }
    if yname == "session-state" { return "SessionState" }
    if yname == "other" { return "Other" }
    return ""
}

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetSegmentPath() string {
    return "packet-counts"
}

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "padi" {
        return &packetCounts.Padi
    }
    if childYangName == "pado" {
        return &packetCounts.Pado
    }
    if childYangName == "padr" {
        return &packetCounts.Padr
    }
    if childYangName == "pads-success" {
        return &packetCounts.PadsSuccess
    }
    if childYangName == "pads-error" {
        return &packetCounts.PadsError
    }
    if childYangName == "padt" {
        return &packetCounts.Padt
    }
    if childYangName == "session-state" {
        return &packetCounts.SessionState
    }
    if childYangName == "other" {
        return &packetCounts.Other
    }
    return nil
}

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["padi"] = &packetCounts.Padi
    children["pado"] = &packetCounts.Pado
    children["padr"] = &packetCounts.Padr
    children["pads-success"] = &packetCounts.PadsSuccess
    children["pads-error"] = &packetCounts.PadsError
    children["padt"] = &packetCounts.Padt
    children["session-state"] = &packetCounts.SessionState
    children["other"] = &packetCounts.Other
    return children
}

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetBundleName() string { return "cisco_ios_xr" }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetYangName() string { return "packet-counts" }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) SetParent(parent types.Entity) { packetCounts.parent = parent }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetParent() types.Entity { return packetCounts.parent }

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetParentYangName() string { return "statistics" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_Padi
// PADI counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Padi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetFilter() yfilter.YFilter { return padi.YFilter }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) SetFilter(yf yfilter.YFilter) { padi.YFilter = yf }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetSegmentPath() string {
    return "padi"
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padi.Sent
    leafs["received"] = padi.Received
    leafs["dropped"] = padi.Dropped
    return leafs
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetBundleName() string { return "cisco_ios_xr" }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetYangName() string { return "padi" }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) SetParent(parent types.Entity) { padi.parent = parent }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetParent() types.Entity { return padi.parent }

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_Pado
// PADO counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Pado struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetFilter() yfilter.YFilter { return pado.YFilter }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) SetFilter(yf yfilter.YFilter) { pado.YFilter = yf }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetSegmentPath() string {
    return "pado"
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = pado.Sent
    leafs["received"] = pado.Received
    leafs["dropped"] = pado.Dropped
    return leafs
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetBundleName() string { return "cisco_ios_xr" }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetYangName() string { return "pado" }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) SetParent(parent types.Entity) { pado.parent = parent }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetParent() types.Entity { return pado.parent }

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_Padr
// PADR counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Padr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetFilter() yfilter.YFilter { return padr.YFilter }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) SetFilter(yf yfilter.YFilter) { padr.YFilter = yf }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetSegmentPath() string {
    return "padr"
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padr.Sent
    leafs["received"] = padr.Received
    leafs["dropped"] = padr.Dropped
    return leafs
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetBundleName() string { return "cisco_ios_xr" }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetYangName() string { return "padr" }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) SetParent(parent types.Entity) { padr.parent = parent }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetParent() types.Entity { return padr.parent }

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess
// PADS Success counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetFilter() yfilter.YFilter { return padsSuccess.YFilter }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) SetFilter(yf yfilter.YFilter) { padsSuccess.YFilter = yf }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetSegmentPath() string {
    return "pads-success"
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padsSuccess.Sent
    leafs["received"] = padsSuccess.Received
    leafs["dropped"] = padsSuccess.Dropped
    return leafs
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetBundleName() string { return "cisco_ios_xr" }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetYangName() string { return "pads-success" }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) SetParent(parent types.Entity) { padsSuccess.parent = parent }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetParent() types.Entity { return padsSuccess.parent }

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError
// PADS Error counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetFilter() yfilter.YFilter { return padsError.YFilter }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) SetFilter(yf yfilter.YFilter) { padsError.YFilter = yf }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetSegmentPath() string {
    return "pads-error"
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padsError.Sent
    leafs["received"] = padsError.Received
    leafs["dropped"] = padsError.Dropped
    return leafs
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetBundleName() string { return "cisco_ios_xr" }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetYangName() string { return "pads-error" }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) SetParent(parent types.Entity) { padsError.parent = parent }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetParent() types.Entity { return padsError.parent }

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_Padt
// PADT counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Padt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetFilter() yfilter.YFilter { return padt.YFilter }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) SetFilter(yf yfilter.YFilter) { padt.YFilter = yf }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetSegmentPath() string {
    return "padt"
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = padt.Sent
    leafs["received"] = padt.Received
    leafs["dropped"] = padt.Dropped
    return leafs
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetBundleName() string { return "cisco_ios_xr" }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetYangName() string { return "padt" }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) SetParent(parent types.Entity) { padt.parent = parent }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetParent() types.Entity { return padt.parent }

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState
// Session Stage counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetFilter() yfilter.YFilter { return sessionState.YFilter }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) SetFilter(yf yfilter.YFilter) { sessionState.YFilter = yf }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetSegmentPath() string {
    return "session-state"
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = sessionState.Sent
    leafs["received"] = sessionState.Received
    leafs["dropped"] = sessionState.Dropped
    return leafs
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetBundleName() string { return "cisco_ios_xr" }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetYangName() string { return "session-state" }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) SetParent(parent types.Entity) { sessionState.parent = parent }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetParent() types.Entity { return sessionState.parent }

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketCounts_Other
// Other counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Other struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetFilter() yfilter.YFilter { return other.YFilter }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) SetFilter(yf yfilter.YFilter) { other.YFilter = yf }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetSegmentPath() string {
    return "other"
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent"] = other.Sent
    leafs["received"] = other.Received
    leafs["dropped"] = other.Dropped
    return leafs
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetBundleName() string { return "cisco_ios_xr" }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetYangName() string { return "other" }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) SetParent(parent types.Entity) { other.parent = parent }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetParent() types.Entity { return other.parent }

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetParentYangName() string { return "packet-counts" }

// Pppoe_Nodes_Node_Statistics_PacketErrorCounts
// Packet Error Counts
type Pppoe_Nodes_Node_Statistics_PacketErrorCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // No interface handle. The type is interface{} with range: 0..4294967295.
    NoInterfaceHandle interface{}

    // No packet payload. The type is interface{} with range: 0..4294967295.
    NoPacketPayload interface{}

    // No packet mac-address. The type is interface{} with range: 0..4294967295.
    NoPacketMacAddress interface{}

    // Invalid version-type value. The type is interface{} with range:
    // 0..4294967295.
    InvalidVersionTypeValue interface{}

    // Bad packet length. The type is interface{} with range: 0..4294967295.
    BadPacketLength interface{}

    // Unknown interface. The type is interface{} with range: 0..4294967295.
    UnknownInterface interface{}

    // PADO received. The type is interface{} with range: 0..4294967295.
    PadoReceived interface{}

    // PADS received. The type is interface{} with range: 0..4294967295.
    PadsReceived interface{}

    // Unknown packet type received. The type is interface{} with range:
    // 0..4294967295.
    UnknownPacketTypeReceived interface{}

    // Unexpected Session-ID in packet. The type is interface{} with range:
    // 0..4294967295.
    UnexpectedSessionIdInPacket interface{}

    // No Service-Name Tag. The type is interface{} with range: 0..4294967295.
    NoServiceNameTag interface{}

    // PADT for unknown session. The type is interface{} with range:
    // 0..4294967295.
    PadtForUnknownSession interface{}

    // PADT with wrong peer-mac. The type is interface{} with range:
    // 0..4294967295.
    PadtWithWrongPeerMac interface{}

    // PADT with wrong VLAN tags. The type is interface{} with range:
    // 0..4294967295.
    PadtWithWrongVlanTags interface{}

    // Zero-length Host-Uniq tag. The type is interface{} with range:
    // 0..4294967295.
    ZeroLengthHostUniq interface{}

    // PADT before PADS sent. The type is interface{} with range: 0..4294967295.
    PadtBeforePadsSent interface{}

    // Session-stage packet for unknown session. The type is interface{} with
    // range: 0..4294967295.
    SessionStagePacketForUnknownSession interface{}

    // Session-stage packet with wrong mac. The type is interface{} with range:
    // 0..4294967295.
    SessionStagePacketWithWrongMac interface{}

    // Session-stage packet with wrong VLAN tags. The type is interface{} with
    // range: 0..4294967295.
    SessionStagePacketWithWrongVlanTags interface{}

    // Session-stage packet with no error. The type is interface{} with range:
    // 0..4294967295.
    SessionStagePacketWithNoError interface{}

    // Tag too short. The type is interface{} with range: 0..4294967295.
    TagTooShort interface{}

    // Bad tag-length field. The type is interface{} with range: 0..4294967295.
    BadTagLengthField interface{}

    // Multiple Service-Name tags. The type is interface{} with range:
    // 0..4294967295.
    MultipleServiceNameTags interface{}

    // Multiple Max-Payload tags. The type is interface{} with range:
    // 0..4294967295.
    MultipleMaxPayloadTags interface{}

    // Invalid Max-Payload tag. The type is interface{} with range: 0..4294967295.
    InvalidMaxPayloadTag interface{}

    // Multiple Vendor-specific tags. The type is interface{} with range:
    // 0..4294967295.
    MultipleVendorSpecificTags interface{}

    // Unexpected AC-Name tag. The type is interface{} with range: 0..4294967295.
    UnexpectedAcNameTag interface{}

    // Unexpected error tags. The type is interface{} with range: 0..4294967295.
    UnexpectedErrorTags interface{}

    // Unknown tag received. The type is interface{} with range: 0..4294967295.
    UnknownTagReceived interface{}

    // No IANA code in vendor tag. The type is interface{} with range:
    // 0..4294967295.
    NoIanaCodeInvendorTag interface{}

    // Invalid IANA code in vendor tag. The type is interface{} with range:
    // 0..4294967295.
    InvalidIanaCodeInvendorTag interface{}

    // Vendor tag too short. The type is interface{} with range: 0..4294967295.
    VendorTagTooShort interface{}

    // Bad vendor tag length field. The type is interface{} with range:
    // 0..4294967295.
    BadVendorTagLengthField interface{}

    // Multiple Host-Uniq tags. The type is interface{} with range: 0..4294967295.
    MultipleHostUniqTags interface{}

    // Multiple relay-session-id tags. The type is interface{} with range:
    // 0..4294967295.
    MultipleRelaySessionIdTags interface{}

    // Multiple Circuit-ID tags. The type is interface{} with range:
    // 0..4294967295.
    MultipleCircuitIdTags interface{}

    // Multiple Remote-ID tags. The type is interface{} with range: 0..4294967295.
    MultipleRemoteIdTags interface{}

    // Invalid DSL tag. The type is interface{} with range: 0..4294967295.
    InvalidDslTag interface{}

    // Multiple of the same DSL tag. The type is interface{} with range:
    // 0..4294967295.
    MultipleOfTheSameDslTag interface{}

    // Invalid IWF tag. The type is interface{} with range: 0..4294967295.
    InvalidIwfTag interface{}

    // Multiple IWF tags. The type is interface{} with range: 0..4294967295.
    MultipleIwfTags interface{}

    // Unknown vendor-tag. The type is interface{} with range: 0..4294967295.
    UnknownvendorTag interface{}

    // No space left in packet. The type is interface{} with range: 0..4294967295.
    NoSpaceLeftInPacket interface{}

    // Duplicate Host-Uniq tag received. The type is interface{} with range:
    // 0..4294967295.
    DuplicateHostUniqTagReceived interface{}

    // Duplicate Relay Session ID tag received. The type is interface{} with
    // range: 0..4294967295.
    DuplicateRelaySessionIdTagReceived interface{}

    // Packet too long. The type is interface{} with range: 0..4294967295.
    PacketTooLong interface{}

    // Invalid ALE tag. The type is interface{} with range: 0..4294967295.
    InvalidAleTag interface{}

    // Multiple ALE tags. The type is interface{} with range: 0..4294967295.
    MultipleAleTags interface{}

    // Invalid Service Name. The type is interface{} with range: 0..4294967295.
    InvalidServiceName interface{}

    // Invalid Peer MAC. The type is interface{} with range: 0..4294967295.
    InvalidPeerMac interface{}

    // Invalid VLAN Tags. The type is interface{} with range: 0..4294967295.
    InvalidVlanTags interface{}

    // Packet Received on SRG Slave. The type is interface{} with range:
    // 0..4294967295.
    PacketOnSrgSlave interface{}
}

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetFilter() yfilter.YFilter { return packetErrorCounts.YFilter }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) SetFilter(yf yfilter.YFilter) { packetErrorCounts.YFilter = yf }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetGoName(yname string) string {
    if yname == "no-interface-handle" { return "NoInterfaceHandle" }
    if yname == "no-packet-payload" { return "NoPacketPayload" }
    if yname == "no-packet-mac-address" { return "NoPacketMacAddress" }
    if yname == "invalid-version-type-value" { return "InvalidVersionTypeValue" }
    if yname == "bad-packet-length" { return "BadPacketLength" }
    if yname == "unknown-interface" { return "UnknownInterface" }
    if yname == "pado-received" { return "PadoReceived" }
    if yname == "pads-received" { return "PadsReceived" }
    if yname == "unknown-packet-type-received" { return "UnknownPacketTypeReceived" }
    if yname == "unexpected-session-id-in-packet" { return "UnexpectedSessionIdInPacket" }
    if yname == "no-service-name-tag" { return "NoServiceNameTag" }
    if yname == "padt-for-unknown-session" { return "PadtForUnknownSession" }
    if yname == "padt-with-wrong-peer-mac" { return "PadtWithWrongPeerMac" }
    if yname == "padt-with-wrong-vlan-tags" { return "PadtWithWrongVlanTags" }
    if yname == "zero-length-host-uniq" { return "ZeroLengthHostUniq" }
    if yname == "padt-before-pads-sent" { return "PadtBeforePadsSent" }
    if yname == "session-stage-packet-for-unknown-session" { return "SessionStagePacketForUnknownSession" }
    if yname == "session-stage-packet-with-wrong-mac" { return "SessionStagePacketWithWrongMac" }
    if yname == "session-stage-packet-with-wrong-vlan-tags" { return "SessionStagePacketWithWrongVlanTags" }
    if yname == "session-stage-packet-with-no-error" { return "SessionStagePacketWithNoError" }
    if yname == "tag-too-short" { return "TagTooShort" }
    if yname == "bad-tag-length-field" { return "BadTagLengthField" }
    if yname == "multiple-service-name-tags" { return "MultipleServiceNameTags" }
    if yname == "multiple-max-payload-tags" { return "MultipleMaxPayloadTags" }
    if yname == "invalid-max-payload-tag" { return "InvalidMaxPayloadTag" }
    if yname == "multiple-vendor-specific-tags" { return "MultipleVendorSpecificTags" }
    if yname == "unexpected-ac-name-tag" { return "UnexpectedAcNameTag" }
    if yname == "unexpected-error-tags" { return "UnexpectedErrorTags" }
    if yname == "unknown-tag-received" { return "UnknownTagReceived" }
    if yname == "no-iana-code-invendor-tag" { return "NoIanaCodeInvendorTag" }
    if yname == "invalid-iana-code-invendor-tag" { return "InvalidIanaCodeInvendorTag" }
    if yname == "vendor-tag-too-short" { return "VendorTagTooShort" }
    if yname == "bad-vendor-tag-length-field" { return "BadVendorTagLengthField" }
    if yname == "multiple-host-uniq-tags" { return "MultipleHostUniqTags" }
    if yname == "multiple-relay-session-id-tags" { return "MultipleRelaySessionIdTags" }
    if yname == "multiple-circuit-id-tags" { return "MultipleCircuitIdTags" }
    if yname == "multiple-remote-id-tags" { return "MultipleRemoteIdTags" }
    if yname == "invalid-dsl-tag" { return "InvalidDslTag" }
    if yname == "multiple-of-the-same-dsl-tag" { return "MultipleOfTheSameDslTag" }
    if yname == "invalid-iwf-tag" { return "InvalidIwfTag" }
    if yname == "multiple-iwf-tags" { return "MultipleIwfTags" }
    if yname == "unknownvendor-tag" { return "UnknownvendorTag" }
    if yname == "no-space-left-in-packet" { return "NoSpaceLeftInPacket" }
    if yname == "duplicate-host-uniq-tag-received" { return "DuplicateHostUniqTagReceived" }
    if yname == "duplicate-relay-session-id-tag-received" { return "DuplicateRelaySessionIdTagReceived" }
    if yname == "packet-too-long" { return "PacketTooLong" }
    if yname == "invalid-ale-tag" { return "InvalidAleTag" }
    if yname == "multiple-ale-tags" { return "MultipleAleTags" }
    if yname == "invalid-service-name" { return "InvalidServiceName" }
    if yname == "invalid-peer-mac" { return "InvalidPeerMac" }
    if yname == "invalid-vlan-tags" { return "InvalidVlanTags" }
    if yname == "packet-on-srg-slave" { return "PacketOnSrgSlave" }
    return ""
}

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetSegmentPath() string {
    return "packet-error-counts"
}

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["no-interface-handle"] = packetErrorCounts.NoInterfaceHandle
    leafs["no-packet-payload"] = packetErrorCounts.NoPacketPayload
    leafs["no-packet-mac-address"] = packetErrorCounts.NoPacketMacAddress
    leafs["invalid-version-type-value"] = packetErrorCounts.InvalidVersionTypeValue
    leafs["bad-packet-length"] = packetErrorCounts.BadPacketLength
    leafs["unknown-interface"] = packetErrorCounts.UnknownInterface
    leafs["pado-received"] = packetErrorCounts.PadoReceived
    leafs["pads-received"] = packetErrorCounts.PadsReceived
    leafs["unknown-packet-type-received"] = packetErrorCounts.UnknownPacketTypeReceived
    leafs["unexpected-session-id-in-packet"] = packetErrorCounts.UnexpectedSessionIdInPacket
    leafs["no-service-name-tag"] = packetErrorCounts.NoServiceNameTag
    leafs["padt-for-unknown-session"] = packetErrorCounts.PadtForUnknownSession
    leafs["padt-with-wrong-peer-mac"] = packetErrorCounts.PadtWithWrongPeerMac
    leafs["padt-with-wrong-vlan-tags"] = packetErrorCounts.PadtWithWrongVlanTags
    leafs["zero-length-host-uniq"] = packetErrorCounts.ZeroLengthHostUniq
    leafs["padt-before-pads-sent"] = packetErrorCounts.PadtBeforePadsSent
    leafs["session-stage-packet-for-unknown-session"] = packetErrorCounts.SessionStagePacketForUnknownSession
    leafs["session-stage-packet-with-wrong-mac"] = packetErrorCounts.SessionStagePacketWithWrongMac
    leafs["session-stage-packet-with-wrong-vlan-tags"] = packetErrorCounts.SessionStagePacketWithWrongVlanTags
    leafs["session-stage-packet-with-no-error"] = packetErrorCounts.SessionStagePacketWithNoError
    leafs["tag-too-short"] = packetErrorCounts.TagTooShort
    leafs["bad-tag-length-field"] = packetErrorCounts.BadTagLengthField
    leafs["multiple-service-name-tags"] = packetErrorCounts.MultipleServiceNameTags
    leafs["multiple-max-payload-tags"] = packetErrorCounts.MultipleMaxPayloadTags
    leafs["invalid-max-payload-tag"] = packetErrorCounts.InvalidMaxPayloadTag
    leafs["multiple-vendor-specific-tags"] = packetErrorCounts.MultipleVendorSpecificTags
    leafs["unexpected-ac-name-tag"] = packetErrorCounts.UnexpectedAcNameTag
    leafs["unexpected-error-tags"] = packetErrorCounts.UnexpectedErrorTags
    leafs["unknown-tag-received"] = packetErrorCounts.UnknownTagReceived
    leafs["no-iana-code-invendor-tag"] = packetErrorCounts.NoIanaCodeInvendorTag
    leafs["invalid-iana-code-invendor-tag"] = packetErrorCounts.InvalidIanaCodeInvendorTag
    leafs["vendor-tag-too-short"] = packetErrorCounts.VendorTagTooShort
    leafs["bad-vendor-tag-length-field"] = packetErrorCounts.BadVendorTagLengthField
    leafs["multiple-host-uniq-tags"] = packetErrorCounts.MultipleHostUniqTags
    leafs["multiple-relay-session-id-tags"] = packetErrorCounts.MultipleRelaySessionIdTags
    leafs["multiple-circuit-id-tags"] = packetErrorCounts.MultipleCircuitIdTags
    leafs["multiple-remote-id-tags"] = packetErrorCounts.MultipleRemoteIdTags
    leafs["invalid-dsl-tag"] = packetErrorCounts.InvalidDslTag
    leafs["multiple-of-the-same-dsl-tag"] = packetErrorCounts.MultipleOfTheSameDslTag
    leafs["invalid-iwf-tag"] = packetErrorCounts.InvalidIwfTag
    leafs["multiple-iwf-tags"] = packetErrorCounts.MultipleIwfTags
    leafs["unknownvendor-tag"] = packetErrorCounts.UnknownvendorTag
    leafs["no-space-left-in-packet"] = packetErrorCounts.NoSpaceLeftInPacket
    leafs["duplicate-host-uniq-tag-received"] = packetErrorCounts.DuplicateHostUniqTagReceived
    leafs["duplicate-relay-session-id-tag-received"] = packetErrorCounts.DuplicateRelaySessionIdTagReceived
    leafs["packet-too-long"] = packetErrorCounts.PacketTooLong
    leafs["invalid-ale-tag"] = packetErrorCounts.InvalidAleTag
    leafs["multiple-ale-tags"] = packetErrorCounts.MultipleAleTags
    leafs["invalid-service-name"] = packetErrorCounts.InvalidServiceName
    leafs["invalid-peer-mac"] = packetErrorCounts.InvalidPeerMac
    leafs["invalid-vlan-tags"] = packetErrorCounts.InvalidVlanTags
    leafs["packet-on-srg-slave"] = packetErrorCounts.PacketOnSrgSlave
    return leafs
}

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetBundleName() string { return "cisco_ios_xr" }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetYangName() string { return "packet-error-counts" }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) SetParent(parent types.Entity) { packetErrorCounts.parent = parent }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetParent() types.Entity { return packetErrorCounts.parent }

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetParentYangName() string { return "statistics" }

// Pppoe_Nodes_Node_AccessInterface
// PPPoE access interface information
type Pppoe_Nodes_Node_AccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE access interface summary information.
    Summaries Pppoe_Nodes_Node_AccessInterface_Summaries
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetFilter() yfilter.YFilter { return accessInterface.YFilter }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) SetFilter(yf yfilter.YFilter) { accessInterface.YFilter = yf }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetGoName(yname string) string {
    if yname == "summaries" { return "Summaries" }
    return ""
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetSegmentPath() string {
    return "access-interface"
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summaries" {
        return &accessInterface.Summaries
    }
    return nil
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summaries"] = &accessInterface.Summaries
    return children
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetYangName() string { return "access-interface" }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) SetParent(parent types.Entity) { accessInterface.parent = parent }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetParent() types.Entity { return accessInterface.parent }

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetParentYangName() string { return "node" }

// Pppoe_Nodes_Node_AccessInterface_Summaries
// PPPoE access interface summary information
type Pppoe_Nodes_Node_AccessInterface_Summaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary information for a PPPoE-enabled access interface. The type is slice
    // of Pppoe_Nodes_Node_AccessInterface_Summaries_Summary.
    Summary []Pppoe_Nodes_Node_AccessInterface_Summaries_Summary
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetFilter() yfilter.YFilter { return summaries.YFilter }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) SetFilter(yf yfilter.YFilter) { summaries.YFilter = yf }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    return ""
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetSegmentPath() string {
    return "summaries"
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        for _, c := range summaries.Summary {
            if summaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_Nodes_Node_AccessInterface_Summaries_Summary{}
        summaries.Summary = append(summaries.Summary, child)
        return &summaries.Summary[len(summaries.Summary)-1]
    }
    return nil
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summaries.Summary {
        children[summaries.Summary[i].GetSegmentPath()] = &summaries.Summary[i]
    }
    return children
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetBundleName() string { return "cisco_ios_xr" }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetYangName() string { return "summaries" }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) SetParent(parent types.Entity) { summaries.parent = parent }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetParent() types.Entity { return summaries.parent }

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetParentYangName() string { return "access-interface" }

// Pppoe_Nodes_Node_AccessInterface_Summaries_Summary
// Summary information for a PPPoE-enabled
// access interface
type Pppoe_Nodes_Node_AccessInterface_Summaries_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. PPPoE Access Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // Interface State. The type is interface{} with range: 0..4294967295.
    InterfaceState interface{}

    // Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // BBA Group. The type is string.
    BbaGroupName interface{}

    // Is Ready. The type is interface{} with range: -2147483648..2147483647.
    IsReady interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    Sessions interface{}

    // Incomplete Session Count. The type is interface{} with range:
    // 0..4294967295.
    IncompleteSessions interface{}
}

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface-state" { return "InterfaceState" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "bba-group-name" { return "BbaGroupName" }
    if yname == "is-ready" { return "IsReady" }
    if yname == "sessions" { return "Sessions" }
    if yname == "incomplete-sessions" { return "IncompleteSessions" }
    return ""
}

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetSegmentPath() string {
    return "summary" + "[interface-name='" + fmt.Sprintf("%v", summary.InterfaceName) + "']"
}

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = summary.InterfaceName
    leafs["interface-name-xr"] = summary.InterfaceNameXr
    leafs["interface-state"] = summary.InterfaceState
    leafs["mac-address"] = summary.MacAddress
    leafs["bba-group-name"] = summary.BbaGroupName
    leafs["is-ready"] = summary.IsReady
    leafs["sessions"] = summary.Sessions
    leafs["incomplete-sessions"] = summary.IncompleteSessions
    return leafs
}

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetYangName() string { return "summary" }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetParentYangName() string { return "summaries" }

// Pppoe_Nodes_Node_Interfaces
// Per interface PPPoE operational data
type Pppoe_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for a PPPoE interface. The type is slice of
    // Pppoe_Nodes_Node_Interfaces_Interface.
    Interface []Pppoe_Nodes_Node_Interfaces_Interface
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Pppoe_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Pppoe_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Pppoe_Nodes_Node_Interfaces_Interface
// Data for a PPPoE interface
type Pppoe_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. PPPoE Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // Access Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    AccessInterfaceName interface{}

    // BBA Group. The type is string.
    BbaGroupName interface{}

    // Session ID. The type is interface{} with range: 0..65535.
    SessionId interface{}

    // Local Mac-Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    LocalMacAddress interface{}

    // Peer Mac-Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Is Complete. The type is interface{} with range: -2147483648..2147483647.
    IsComplete interface{}

    // VLAN Outer ID. The type is interface{} with range: 0..65535.
    VlanOuterId interface{}

    // VLAN Inner ID. The type is interface{} with range: 0..65535.
    VlanInnerId interface{}

    // SRG state. The type is PppoeMaSessionIdbSrgState.
    SrgState interface{}

    // Tags.
    Tags Pppoe_Nodes_Node_Interfaces_Interface_Tags
}

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "access-interface-name" { return "AccessInterfaceName" }
    if yname == "bba-group-name" { return "BbaGroupName" }
    if yname == "session-id" { return "SessionId" }
    if yname == "local-mac-address" { return "LocalMacAddress" }
    if yname == "peer-mac-address" { return "PeerMacAddress" }
    if yname == "is-complete" { return "IsComplete" }
    if yname == "vlan-outer-id" { return "VlanOuterId" }
    if yname == "vlan-inner-id" { return "VlanInnerId" }
    if yname == "srg-state" { return "SrgState" }
    if yname == "tags" { return "Tags" }
    return ""
}

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tags" {
        return &self.Tags
    }
    return nil
}

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tags"] = &self.Tags
    return children
}

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-name-xr"] = self.InterfaceNameXr
    leafs["access-interface-name"] = self.AccessInterfaceName
    leafs["bba-group-name"] = self.BbaGroupName
    leafs["session-id"] = self.SessionId
    leafs["local-mac-address"] = self.LocalMacAddress
    leafs["peer-mac-address"] = self.PeerMacAddress
    leafs["is-complete"] = self.IsComplete
    leafs["vlan-outer-id"] = self.VlanOuterId
    leafs["vlan-inner-id"] = self.VlanInnerId
    leafs["srg-state"] = self.SrgState
    return leafs
}

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Pppoe_Nodes_Node_Interfaces_Interface_Tags
// Tags
type Pppoe_Nodes_Node_Interfaces_Interface_Tags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service Name. The type is string.
    ServiceName interface{}

    // Max Payload. The type is interface{} with range: 0..65535.
    MaxPayload interface{}

    // Host Uniq. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HostUniq interface{}

    // Relay Session ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    RelaySessionId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Circuit ID. The type is string.
    CircuitId interface{}

    // Is IWF. The type is interface{} with range: -2147483648..2147483647.
    IsIwf interface{}

    // DSL Actual Up. The type is interface{} with range: 0..4294967295.
    DslActualUp interface{}

    // DSL Actual Down. The type is interface{} with range: 0..4294967295.
    DslActualDown interface{}

    // DSL Min Up. The type is interface{} with range: 0..4294967295.
    DslMinUp interface{}

    // DSL Min Down. The type is interface{} with range: 0..4294967295.
    DslMinDown interface{}

    // DSL Attain Up. The type is interface{} with range: 0..4294967295.
    DslAttainUp interface{}

    // DSL Attain Down. The type is interface{} with range: 0..4294967295.
    DslAttainDown interface{}

    // DSL Max Up. The type is interface{} with range: 0..4294967295.
    DslMaxUp interface{}

    // DSL Max Down. The type is interface{} with range: 0..4294967295.
    DslMaxDown interface{}

    // DSL Min Up Low. The type is interface{} with range: 0..4294967295.
    DslMinUpLow interface{}

    // DSL Min Down Low. The type is interface{} with range: 0..4294967295.
    DslMinDownLow interface{}

    // DSL Max Delay Up. The type is interface{} with range: 0..4294967295.
    DslMaxDelayUp interface{}

    // DSL Actual Delay Up. The type is interface{} with range: 0..4294967295.
    DslActualDelayUp interface{}

    // DSL Max Delay Down. The type is interface{} with range: 0..4294967295.
    DslMaxDelayDown interface{}

    // DSL Actual Delay Down. The type is interface{} with range: 0..4294967295.
    DslActualDelayDown interface{}

    // Access Loop Encapsulation.
    AccessLoopEncapsulation Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation
}

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetFilter() yfilter.YFilter { return tags.YFilter }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) SetFilter(yf yfilter.YFilter) { tags.YFilter = yf }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetGoName(yname string) string {
    if yname == "service-name" { return "ServiceName" }
    if yname == "max-payload" { return "MaxPayload" }
    if yname == "host-uniq" { return "HostUniq" }
    if yname == "relay-session-id" { return "RelaySessionId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "is-iwf" { return "IsIwf" }
    if yname == "dsl-actual-up" { return "DslActualUp" }
    if yname == "dsl-actual-down" { return "DslActualDown" }
    if yname == "dsl-min-up" { return "DslMinUp" }
    if yname == "dsl-min-down" { return "DslMinDown" }
    if yname == "dsl-attain-up" { return "DslAttainUp" }
    if yname == "dsl-attain-down" { return "DslAttainDown" }
    if yname == "dsl-max-up" { return "DslMaxUp" }
    if yname == "dsl-max-down" { return "DslMaxDown" }
    if yname == "dsl-min-up-low" { return "DslMinUpLow" }
    if yname == "dsl-min-down-low" { return "DslMinDownLow" }
    if yname == "dsl-max-delay-up" { return "DslMaxDelayUp" }
    if yname == "dsl-actual-delay-up" { return "DslActualDelayUp" }
    if yname == "dsl-max-delay-down" { return "DslMaxDelayDown" }
    if yname == "dsl-actual-delay-down" { return "DslActualDelayDown" }
    if yname == "access-loop-encapsulation" { return "AccessLoopEncapsulation" }
    return ""
}

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetSegmentPath() string {
    return "tags"
}

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-loop-encapsulation" {
        return &tags.AccessLoopEncapsulation
    }
    return nil
}

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-loop-encapsulation"] = &tags.AccessLoopEncapsulation
    return children
}

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service-name"] = tags.ServiceName
    leafs["max-payload"] = tags.MaxPayload
    leafs["host-uniq"] = tags.HostUniq
    leafs["relay-session-id"] = tags.RelaySessionId
    leafs["remote-id"] = tags.RemoteId
    leafs["circuit-id"] = tags.CircuitId
    leafs["is-iwf"] = tags.IsIwf
    leafs["dsl-actual-up"] = tags.DslActualUp
    leafs["dsl-actual-down"] = tags.DslActualDown
    leafs["dsl-min-up"] = tags.DslMinUp
    leafs["dsl-min-down"] = tags.DslMinDown
    leafs["dsl-attain-up"] = tags.DslAttainUp
    leafs["dsl-attain-down"] = tags.DslAttainDown
    leafs["dsl-max-up"] = tags.DslMaxUp
    leafs["dsl-max-down"] = tags.DslMaxDown
    leafs["dsl-min-up-low"] = tags.DslMinUpLow
    leafs["dsl-min-down-low"] = tags.DslMinDownLow
    leafs["dsl-max-delay-up"] = tags.DslMaxDelayUp
    leafs["dsl-actual-delay-up"] = tags.DslActualDelayUp
    leafs["dsl-max-delay-down"] = tags.DslMaxDelayDown
    leafs["dsl-actual-delay-down"] = tags.DslActualDelayDown
    return leafs
}

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetBundleName() string { return "cisco_ios_xr" }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetYangName() string { return "tags" }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) SetParent(parent types.Entity) { tags.parent = parent }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetParent() types.Entity { return tags.parent }

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetParentYangName() string { return "interface" }

// Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation
// Access Loop Encapsulation
type Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data Link. The type is interface{} with range: 0..255.
    DataLink interface{}

    // Encaps 1. The type is interface{} with range: 0..255.
    Encaps1 interface{}

    // Encaps 2. The type is interface{} with range: 0..255.
    Encaps2 interface{}
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetFilter() yfilter.YFilter { return accessLoopEncapsulation.YFilter }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) SetFilter(yf yfilter.YFilter) { accessLoopEncapsulation.YFilter = yf }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetGoName(yname string) string {
    if yname == "data-link" { return "DataLink" }
    if yname == "encaps1" { return "Encaps1" }
    if yname == "encaps2" { return "Encaps2" }
    return ""
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetSegmentPath() string {
    return "access-loop-encapsulation"
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-link"] = accessLoopEncapsulation.DataLink
    leafs["encaps1"] = accessLoopEncapsulation.Encaps1
    leafs["encaps2"] = accessLoopEncapsulation.Encaps2
    return leafs
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetBundleName() string { return "cisco_ios_xr" }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetYangName() string { return "access-loop-encapsulation" }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) SetParent(parent types.Entity) { accessLoopEncapsulation.parent = parent }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetParent() types.Entity { return accessLoopEncapsulation.parent }

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetParentYangName() string { return "tags" }

// Pppoe_Nodes_Node_BbaGroups
// PPPoE BBA-Group information
type Pppoe_Nodes_Node_BbaGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE BBA-Group information. The type is slice of
    // Pppoe_Nodes_Node_BbaGroups_BbaGroup.
    BbaGroup []Pppoe_Nodes_Node_BbaGroups_BbaGroup
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetFilter() yfilter.YFilter { return bbaGroups.YFilter }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) SetFilter(yf yfilter.YFilter) { bbaGroups.YFilter = yf }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetGoName(yname string) string {
    if yname == "bba-group" { return "BbaGroup" }
    return ""
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetSegmentPath() string {
    return "bba-groups"
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bba-group" {
        for _, c := range bbaGroups.BbaGroup {
            if bbaGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_Nodes_Node_BbaGroups_BbaGroup{}
        bbaGroups.BbaGroup = append(bbaGroups.BbaGroup, child)
        return &bbaGroups.BbaGroup[len(bbaGroups.BbaGroup)-1]
    }
    return nil
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bbaGroups.BbaGroup {
        children[bbaGroups.BbaGroup[i].GetSegmentPath()] = &bbaGroups.BbaGroup[i]
    }
    return children
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetBundleName() string { return "cisco_ios_xr" }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetYangName() string { return "bba-groups" }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) SetParent(parent types.Entity) { bbaGroups.parent = parent }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetParent() types.Entity { return bbaGroups.parent }

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetParentYangName() string { return "node" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup
// PPPoE BBA-Group information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BBA Group. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    BbaGroupName interface{}

    // BBA-Group limit configuration information.
    LimitConfig Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig

    // PPPoE session limit information.
    Limits Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits

    // PPPoE throttle information.
    Throttles Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles

    // BBA-Group throttle configuration information.
    ThrottleConfig Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig
}

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetFilter() yfilter.YFilter { return bbaGroup.YFilter }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) SetFilter(yf yfilter.YFilter) { bbaGroup.YFilter = yf }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetGoName(yname string) string {
    if yname == "bba-group-name" { return "BbaGroupName" }
    if yname == "limit-config" { return "LimitConfig" }
    if yname == "limits" { return "Limits" }
    if yname == "throttles" { return "Throttles" }
    if yname == "throttle-config" { return "ThrottleConfig" }
    return ""
}

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetSegmentPath() string {
    return "bba-group" + "[bba-group-name='" + fmt.Sprintf("%v", bbaGroup.BbaGroupName) + "']"
}

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "limit-config" {
        return &bbaGroup.LimitConfig
    }
    if childYangName == "limits" {
        return &bbaGroup.Limits
    }
    if childYangName == "throttles" {
        return &bbaGroup.Throttles
    }
    if childYangName == "throttle-config" {
        return &bbaGroup.ThrottleConfig
    }
    return nil
}

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["limit-config"] = &bbaGroup.LimitConfig
    children["limits"] = &bbaGroup.Limits
    children["throttles"] = &bbaGroup.Throttles
    children["throttle-config"] = &bbaGroup.ThrottleConfig
    return children
}

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bba-group-name"] = bbaGroup.BbaGroupName
    return leafs
}

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetBundleName() string { return "cisco_ios_xr" }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetYangName() string { return "bba-group" }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) SetParent(parent types.Entity) { bbaGroup.parent = parent }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetParent() types.Entity { return bbaGroup.parent }

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetParentYangName() string { return "bba-groups" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig
// BBA-Group limit configuration information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Card.
    Card Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card

    // Access Interface.
    AccessIntf Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf

    // MAC.
    Mac Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac

    // MAC IWF.
    MacIwf Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf

    // MAC Access Interface.
    MacAccessInterface Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface

    // MAC IWF Access Interface.
    MacIwfAccessInterface Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface

    // Circuit ID.
    CircuitId Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId

    // Remote ID.
    RemoteId Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId

    // Circuit ID and Remote ID.
    CircuitIdAndRemoteId Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId

    // Outer VLAN ID.
    OuterVlanId Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId

    // Inner VLAN ID.
    InnerVlanId Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId

    // VLAN ID.
    VlanId Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId
}

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetFilter() yfilter.YFilter { return limitConfig.YFilter }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) SetFilter(yf yfilter.YFilter) { limitConfig.YFilter = yf }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetGoName(yname string) string {
    if yname == "card" { return "Card" }
    if yname == "access-intf" { return "AccessIntf" }
    if yname == "mac" { return "Mac" }
    if yname == "mac-iwf" { return "MacIwf" }
    if yname == "mac-access-interface" { return "MacAccessInterface" }
    if yname == "mac-iwf-access-interface" { return "MacIwfAccessInterface" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "circuit-id-and-remote-id" { return "CircuitIdAndRemoteId" }
    if yname == "outer-vlan-id" { return "OuterVlanId" }
    if yname == "inner-vlan-id" { return "InnerVlanId" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetSegmentPath() string {
    return "limit-config"
}

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "card" {
        return &limitConfig.Card
    }
    if childYangName == "access-intf" {
        return &limitConfig.AccessIntf
    }
    if childYangName == "mac" {
        return &limitConfig.Mac
    }
    if childYangName == "mac-iwf" {
        return &limitConfig.MacIwf
    }
    if childYangName == "mac-access-interface" {
        return &limitConfig.MacAccessInterface
    }
    if childYangName == "mac-iwf-access-interface" {
        return &limitConfig.MacIwfAccessInterface
    }
    if childYangName == "circuit-id" {
        return &limitConfig.CircuitId
    }
    if childYangName == "remote-id" {
        return &limitConfig.RemoteId
    }
    if childYangName == "circuit-id-and-remote-id" {
        return &limitConfig.CircuitIdAndRemoteId
    }
    if childYangName == "outer-vlan-id" {
        return &limitConfig.OuterVlanId
    }
    if childYangName == "inner-vlan-id" {
        return &limitConfig.InnerVlanId
    }
    if childYangName == "vlan-id" {
        return &limitConfig.VlanId
    }
    return nil
}

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["card"] = &limitConfig.Card
    children["access-intf"] = &limitConfig.AccessIntf
    children["mac"] = &limitConfig.Mac
    children["mac-iwf"] = &limitConfig.MacIwf
    children["mac-access-interface"] = &limitConfig.MacAccessInterface
    children["mac-iwf-access-interface"] = &limitConfig.MacIwfAccessInterface
    children["circuit-id"] = &limitConfig.CircuitId
    children["remote-id"] = &limitConfig.RemoteId
    children["circuit-id-and-remote-id"] = &limitConfig.CircuitIdAndRemoteId
    children["outer-vlan-id"] = &limitConfig.OuterVlanId
    children["inner-vlan-id"] = &limitConfig.InnerVlanId
    children["vlan-id"] = &limitConfig.VlanId
    return children
}

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetBundleName() string { return "cisco_ios_xr" }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetYangName() string { return "limit-config" }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) SetParent(parent types.Entity) { limitConfig.parent = parent }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetParent() types.Entity { return limitConfig.parent }

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetParentYangName() string { return "bba-group" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card
// Card
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetFilter() yfilter.YFilter { return card.YFilter }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) SetFilter(yf yfilter.YFilter) { card.YFilter = yf }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetSegmentPath() string {
    return "card"
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = card.MaxLimit
    leafs["threshold"] = card.Threshold
    leafs["radius-override-enabled"] = card.RadiusOverrideEnabled
    return leafs
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetBundleName() string { return "cisco_ios_xr" }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetYangName() string { return "card" }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) SetParent(parent types.Entity) { card.parent = parent }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetParent() types.Entity { return card.parent }

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf
// Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetFilter() yfilter.YFilter { return accessIntf.YFilter }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) SetFilter(yf yfilter.YFilter) { accessIntf.YFilter = yf }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetSegmentPath() string {
    return "access-intf"
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = accessIntf.MaxLimit
    leafs["threshold"] = accessIntf.Threshold
    leafs["radius-override-enabled"] = accessIntf.RadiusOverrideEnabled
    return leafs
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetBundleName() string { return "cisco_ios_xr" }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetYangName() string { return "access-intf" }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) SetParent(parent types.Entity) { accessIntf.parent = parent }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetParent() types.Entity { return accessIntf.parent }

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac
// MAC
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = mac.MaxLimit
    leafs["threshold"] = mac.Threshold
    leafs["radius-override-enabled"] = mac.RadiusOverrideEnabled
    return leafs
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetYangName() string { return "mac" }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf
// MAC IWF
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetFilter() yfilter.YFilter { return macIwf.YFilter }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) SetFilter(yf yfilter.YFilter) { macIwf.YFilter = yf }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetSegmentPath() string {
    return "mac-iwf"
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = macIwf.MaxLimit
    leafs["threshold"] = macIwf.Threshold
    leafs["radius-override-enabled"] = macIwf.RadiusOverrideEnabled
    return leafs
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetBundleName() string { return "cisco_ios_xr" }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetYangName() string { return "mac-iwf" }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) SetParent(parent types.Entity) { macIwf.parent = parent }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetParent() types.Entity { return macIwf.parent }

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface
// MAC Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetFilter() yfilter.YFilter { return macAccessInterface.YFilter }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) SetFilter(yf yfilter.YFilter) { macAccessInterface.YFilter = yf }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetSegmentPath() string {
    return "mac-access-interface"
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = macAccessInterface.MaxLimit
    leafs["threshold"] = macAccessInterface.Threshold
    leafs["radius-override-enabled"] = macAccessInterface.RadiusOverrideEnabled
    return leafs
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetYangName() string { return "mac-access-interface" }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) SetParent(parent types.Entity) { macAccessInterface.parent = parent }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetParent() types.Entity { return macAccessInterface.parent }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface
// MAC IWF Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetFilter() yfilter.YFilter { return macIwfAccessInterface.YFilter }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) SetFilter(yf yfilter.YFilter) { macIwfAccessInterface.YFilter = yf }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetSegmentPath() string {
    return "mac-iwf-access-interface"
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = macIwfAccessInterface.MaxLimit
    leafs["threshold"] = macIwfAccessInterface.Threshold
    leafs["radius-override-enabled"] = macIwfAccessInterface.RadiusOverrideEnabled
    return leafs
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetYangName() string { return "mac-iwf-access-interface" }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) SetParent(parent types.Entity) { macIwfAccessInterface.parent = parent }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetParent() types.Entity { return macIwfAccessInterface.parent }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId
// Circuit ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetFilter() yfilter.YFilter { return circuitId.YFilter }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) SetFilter(yf yfilter.YFilter) { circuitId.YFilter = yf }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetSegmentPath() string {
    return "circuit-id"
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = circuitId.MaxLimit
    leafs["threshold"] = circuitId.Threshold
    leafs["radius-override-enabled"] = circuitId.RadiusOverrideEnabled
    return leafs
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetYangName() string { return "circuit-id" }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) SetParent(parent types.Entity) { circuitId.parent = parent }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetParent() types.Entity { return circuitId.parent }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId
// Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetFilter() yfilter.YFilter { return remoteId.YFilter }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) SetFilter(yf yfilter.YFilter) { remoteId.YFilter = yf }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetSegmentPath() string {
    return "remote-id"
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = remoteId.MaxLimit
    leafs["threshold"] = remoteId.Threshold
    leafs["radius-override-enabled"] = remoteId.RadiusOverrideEnabled
    return leafs
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetBundleName() string { return "cisco_ios_xr" }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetYangName() string { return "remote-id" }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) SetParent(parent types.Entity) { remoteId.parent = parent }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetParent() types.Entity { return remoteId.parent }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId
// Circuit ID and Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetFilter() yfilter.YFilter { return circuitIdAndRemoteId.YFilter }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) SetFilter(yf yfilter.YFilter) { circuitIdAndRemoteId.YFilter = yf }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetSegmentPath() string {
    return "circuit-id-and-remote-id"
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = circuitIdAndRemoteId.MaxLimit
    leafs["threshold"] = circuitIdAndRemoteId.Threshold
    leafs["radius-override-enabled"] = circuitIdAndRemoteId.RadiusOverrideEnabled
    return leafs
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetYangName() string { return "circuit-id-and-remote-id" }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) SetParent(parent types.Entity) { circuitIdAndRemoteId.parent = parent }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetParent() types.Entity { return circuitIdAndRemoteId.parent }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId
// Outer VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetFilter() yfilter.YFilter { return outerVlanId.YFilter }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) SetFilter(yf yfilter.YFilter) { outerVlanId.YFilter = yf }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetSegmentPath() string {
    return "outer-vlan-id"
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = outerVlanId.MaxLimit
    leafs["threshold"] = outerVlanId.Threshold
    leafs["radius-override-enabled"] = outerVlanId.RadiusOverrideEnabled
    return leafs
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetBundleName() string { return "cisco_ios_xr" }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetYangName() string { return "outer-vlan-id" }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) SetParent(parent types.Entity) { outerVlanId.parent = parent }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetParent() types.Entity { return outerVlanId.parent }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId
// Inner VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetFilter() yfilter.YFilter { return innerVlanId.YFilter }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) SetFilter(yf yfilter.YFilter) { innerVlanId.YFilter = yf }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetSegmentPath() string {
    return "inner-vlan-id"
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = innerVlanId.MaxLimit
    leafs["threshold"] = innerVlanId.Threshold
    leafs["radius-override-enabled"] = innerVlanId.RadiusOverrideEnabled
    return leafs
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetBundleName() string { return "cisco_ios_xr" }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetYangName() string { return "inner-vlan-id" }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) SetParent(parent types.Entity) { innerVlanId.parent = parent }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetParent() types.Entity { return innerVlanId.parent }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId
// VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetFilter() yfilter.YFilter { return vlanId.YFilter }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) SetFilter(yf yfilter.YFilter) { vlanId.YFilter = yf }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "threshold" { return "Threshold" }
    if yname == "radius-override-enabled" { return "RadiusOverrideEnabled" }
    return ""
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetSegmentPath() string {
    return "vlan-id"
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = vlanId.MaxLimit
    leafs["threshold"] = vlanId.Threshold
    leafs["radius-override-enabled"] = vlanId.RadiusOverrideEnabled
    return leafs
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetBundleName() string { return "cisco_ios_xr" }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetYangName() string { return "vlan-id" }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) SetParent(parent types.Entity) { vlanId.parent = parent }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetParent() types.Entity { return vlanId.parent }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetParentYangName() string { return "limit-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits
// PPPoE session limit information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE session limit state. The type is slice of
    // Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit.
    Limit []Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetFilter() yfilter.YFilter { return limits.YFilter }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) SetFilter(yf yfilter.YFilter) { limits.YFilter = yf }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    return ""
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetSegmentPath() string {
    return "limits"
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "limit" {
        for _, c := range limits.Limit {
            if limits.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit{}
        limits.Limit = append(limits.Limit, child)
        return &limits.Limit[len(limits.Limit)-1]
    }
    return nil
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range limits.Limit {
        children[limits.Limit[i].GetSegmentPath()] = &limits.Limit[i]
    }
    return children
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetBundleName() string { return "cisco_ios_xr" }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetYangName() string { return "limits" }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) SetParent(parent types.Entity) { limits.parent = parent }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetParent() types.Entity { return limits.parent }

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetParentYangName() string { return "bba-group" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit
// PPPoE session limit state
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IWF flag. The type is bool.
    Iwf interface{}

    // Circuit ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CircuitId interface{}

    // Remote ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    RemoteId interface{}

    // Outer VLAN ID. The type is interface{} with range: 0..4095.
    OuterVlanId interface{}

    // Inner VLAN ID. The type is interface{} with range: 0..4095.
    InnerVlanId interface{}

    // State. The type is PppoeMaLimitState.
    State interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Overridden limit has been set. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideSet interface{}

    // Overridden limit if set. The type is interface{} with range: 0..4294967295.
    OverrideLimit interface{}
}

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetFilter() yfilter.YFilter { return limit.YFilter }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) SetFilter(yf yfilter.YFilter) { limit.YFilter = yf }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "iwf" { return "Iwf" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "outer-vlan-id" { return "OuterVlanId" }
    if yname == "inner-vlan-id" { return "InnerVlanId" }
    if yname == "state" { return "State" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "radius-override-set" { return "RadiusOverrideSet" }
    if yname == "override-limit" { return "OverrideLimit" }
    return ""
}

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetSegmentPath() string {
    return "limit"
}

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = limit.InterfaceName
    leafs["mac-address"] = limit.MacAddress
    leafs["iwf"] = limit.Iwf
    leafs["circuit-id"] = limit.CircuitId
    leafs["remote-id"] = limit.RemoteId
    leafs["outer-vlan-id"] = limit.OuterVlanId
    leafs["inner-vlan-id"] = limit.InnerVlanId
    leafs["state"] = limit.State
    leafs["session-count"] = limit.SessionCount
    leafs["radius-override-set"] = limit.RadiusOverrideSet
    leafs["override-limit"] = limit.OverrideLimit
    return leafs
}

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetBundleName() string { return "cisco_ios_xr" }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetYangName() string { return "limit" }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) SetParent(parent types.Entity) { limit.parent = parent }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetParent() types.Entity { return limit.parent }

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetParentYangName() string { return "limits" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles
// PPPoE throttle information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE session throttle state. The type is slice of
    // Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle.
    Throttle []Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetFilter() yfilter.YFilter { return throttles.YFilter }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) SetFilter(yf yfilter.YFilter) { throttles.YFilter = yf }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    return ""
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetSegmentPath() string {
    return "throttles"
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle" {
        for _, c := range throttles.Throttle {
            if throttles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle{}
        throttles.Throttle = append(throttles.Throttle, child)
        return &throttles.Throttle[len(throttles.Throttle)-1]
    }
    return nil
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range throttles.Throttle {
        children[throttles.Throttle[i].GetSegmentPath()] = &throttles.Throttle[i]
    }
    return children
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetBundleName() string { return "cisco_ios_xr" }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetYangName() string { return "throttles" }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) SetParent(parent types.Entity) { throttles.parent = parent }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetParent() types.Entity { return throttles.parent }

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetParentYangName() string { return "bba-group" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle
// PPPoE session throttle state
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IWF flag. The type is bool.
    Iwf interface{}

    // Circuit ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CircuitId interface{}

    // Remote ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    RemoteId interface{}

    // Outer VLAN ID. The type is interface{} with range: 0..4095.
    OuterVlanId interface{}

    // Inner VLAN ID. The type is interface{} with range: 0..4095.
    InnerVlanId interface{}

    // State. The type is PppoeMaThrottleState.
    State interface{}

    // Time left in seconds. The type is interface{} with range: 0..4294967295.
    // Units are second.
    TimeLeft interface{}

    // Number of seconds since counters reset. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SinceReset interface{}

    // PADI Count. The type is interface{} with range: 0..4294967295.
    PadiCount interface{}

    // PADR Count. The type is interface{} with range: 0..4294967295.
    PadrCount interface{}
}

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "iwf" { return "Iwf" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "outer-vlan-id" { return "OuterVlanId" }
    if yname == "inner-vlan-id" { return "InnerVlanId" }
    if yname == "state" { return "State" }
    if yname == "time-left" { return "TimeLeft" }
    if yname == "since-reset" { return "SinceReset" }
    if yname == "padi-count" { return "PadiCount" }
    if yname == "padr-count" { return "PadrCount" }
    return ""
}

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = throttle.InterfaceName
    leafs["mac-address"] = throttle.MacAddress
    leafs["iwf"] = throttle.Iwf
    leafs["circuit-id"] = throttle.CircuitId
    leafs["remote-id"] = throttle.RemoteId
    leafs["outer-vlan-id"] = throttle.OuterVlanId
    leafs["inner-vlan-id"] = throttle.InnerVlanId
    leafs["state"] = throttle.State
    leafs["time-left"] = throttle.TimeLeft
    leafs["since-reset"] = throttle.SinceReset
    leafs["padi-count"] = throttle.PadiCount
    leafs["padr-count"] = throttle.PadrCount
    return leafs
}

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetYangName() string { return "throttle" }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetParentYangName() string { return "throttles" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig
// BBA-Group throttle configuration information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC.
    Mac Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac

    // MAC Access Interface.
    MacAccessInterface Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface

    // MAC IWF Access Interface.
    MacIwfAccessInterface Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface

    // Circuit ID.
    CircuitId Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId

    // Remote ID.
    RemoteId Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId

    // Circuit ID and Remote ID.
    CircuitIdAndRemoteId Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId

    // Outer VLAN ID.
    OuterVlanId Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId

    // Inner VLAN ID.
    InnerVlanId Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId

    // VLAN ID.
    VlanId Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId
}

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetFilter() yfilter.YFilter { return throttleConfig.YFilter }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) SetFilter(yf yfilter.YFilter) { throttleConfig.YFilter = yf }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    if yname == "mac-access-interface" { return "MacAccessInterface" }
    if yname == "mac-iwf-access-interface" { return "MacIwfAccessInterface" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "circuit-id-and-remote-id" { return "CircuitIdAndRemoteId" }
    if yname == "outer-vlan-id" { return "OuterVlanId" }
    if yname == "inner-vlan-id" { return "InnerVlanId" }
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetSegmentPath() string {
    return "throttle-config"
}

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        return &throttleConfig.Mac
    }
    if childYangName == "mac-access-interface" {
        return &throttleConfig.MacAccessInterface
    }
    if childYangName == "mac-iwf-access-interface" {
        return &throttleConfig.MacIwfAccessInterface
    }
    if childYangName == "circuit-id" {
        return &throttleConfig.CircuitId
    }
    if childYangName == "remote-id" {
        return &throttleConfig.RemoteId
    }
    if childYangName == "circuit-id-and-remote-id" {
        return &throttleConfig.CircuitIdAndRemoteId
    }
    if childYangName == "outer-vlan-id" {
        return &throttleConfig.OuterVlanId
    }
    if childYangName == "inner-vlan-id" {
        return &throttleConfig.InnerVlanId
    }
    if childYangName == "vlan-id" {
        return &throttleConfig.VlanId
    }
    return nil
}

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac"] = &throttleConfig.Mac
    children["mac-access-interface"] = &throttleConfig.MacAccessInterface
    children["mac-iwf-access-interface"] = &throttleConfig.MacIwfAccessInterface
    children["circuit-id"] = &throttleConfig.CircuitId
    children["remote-id"] = &throttleConfig.RemoteId
    children["circuit-id-and-remote-id"] = &throttleConfig.CircuitIdAndRemoteId
    children["outer-vlan-id"] = &throttleConfig.OuterVlanId
    children["inner-vlan-id"] = &throttleConfig.InnerVlanId
    children["vlan-id"] = &throttleConfig.VlanId
    return children
}

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetBundleName() string { return "cisco_ios_xr" }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetYangName() string { return "throttle-config" }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) SetParent(parent types.Entity) { throttleConfig.parent = parent }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetParent() types.Entity { return throttleConfig.parent }

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetParentYangName() string { return "bba-group" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac
// MAC
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = mac.Limit
    leafs["request-period"] = mac.RequestPeriod
    leafs["blocking-period"] = mac.BlockingPeriod
    return leafs
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetYangName() string { return "mac" }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface
// MAC Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetFilter() yfilter.YFilter { return macAccessInterface.YFilter }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) SetFilter(yf yfilter.YFilter) { macAccessInterface.YFilter = yf }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetSegmentPath() string {
    return "mac-access-interface"
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = macAccessInterface.Limit
    leafs["request-period"] = macAccessInterface.RequestPeriod
    leafs["blocking-period"] = macAccessInterface.BlockingPeriod
    return leafs
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetYangName() string { return "mac-access-interface" }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) SetParent(parent types.Entity) { macAccessInterface.parent = parent }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetParent() types.Entity { return macAccessInterface.parent }

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface
// MAC IWF Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetFilter() yfilter.YFilter { return macIwfAccessInterface.YFilter }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) SetFilter(yf yfilter.YFilter) { macIwfAccessInterface.YFilter = yf }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetSegmentPath() string {
    return "mac-iwf-access-interface"
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = macIwfAccessInterface.Limit
    leafs["request-period"] = macIwfAccessInterface.RequestPeriod
    leafs["blocking-period"] = macIwfAccessInterface.BlockingPeriod
    return leafs
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetYangName() string { return "mac-iwf-access-interface" }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) SetParent(parent types.Entity) { macIwfAccessInterface.parent = parent }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetParent() types.Entity { return macIwfAccessInterface.parent }

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId
// Circuit ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetFilter() yfilter.YFilter { return circuitId.YFilter }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) SetFilter(yf yfilter.YFilter) { circuitId.YFilter = yf }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetSegmentPath() string {
    return "circuit-id"
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = circuitId.Limit
    leafs["request-period"] = circuitId.RequestPeriod
    leafs["blocking-period"] = circuitId.BlockingPeriod
    return leafs
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetYangName() string { return "circuit-id" }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) SetParent(parent types.Entity) { circuitId.parent = parent }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetParent() types.Entity { return circuitId.parent }

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId
// Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetFilter() yfilter.YFilter { return remoteId.YFilter }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) SetFilter(yf yfilter.YFilter) { remoteId.YFilter = yf }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetSegmentPath() string {
    return "remote-id"
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = remoteId.Limit
    leafs["request-period"] = remoteId.RequestPeriod
    leafs["blocking-period"] = remoteId.BlockingPeriod
    return leafs
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetBundleName() string { return "cisco_ios_xr" }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetYangName() string { return "remote-id" }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) SetParent(parent types.Entity) { remoteId.parent = parent }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetParent() types.Entity { return remoteId.parent }

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId
// Circuit ID and Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetFilter() yfilter.YFilter { return circuitIdAndRemoteId.YFilter }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) SetFilter(yf yfilter.YFilter) { circuitIdAndRemoteId.YFilter = yf }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetSegmentPath() string {
    return "circuit-id-and-remote-id"
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = circuitIdAndRemoteId.Limit
    leafs["request-period"] = circuitIdAndRemoteId.RequestPeriod
    leafs["blocking-period"] = circuitIdAndRemoteId.BlockingPeriod
    return leafs
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetYangName() string { return "circuit-id-and-remote-id" }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) SetParent(parent types.Entity) { circuitIdAndRemoteId.parent = parent }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetParent() types.Entity { return circuitIdAndRemoteId.parent }

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId
// Outer VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetFilter() yfilter.YFilter { return outerVlanId.YFilter }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) SetFilter(yf yfilter.YFilter) { outerVlanId.YFilter = yf }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetSegmentPath() string {
    return "outer-vlan-id"
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = outerVlanId.Limit
    leafs["request-period"] = outerVlanId.RequestPeriod
    leafs["blocking-period"] = outerVlanId.BlockingPeriod
    return leafs
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetBundleName() string { return "cisco_ios_xr" }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetYangName() string { return "outer-vlan-id" }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) SetParent(parent types.Entity) { outerVlanId.parent = parent }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetParent() types.Entity { return outerVlanId.parent }

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId
// Inner VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetFilter() yfilter.YFilter { return innerVlanId.YFilter }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) SetFilter(yf yfilter.YFilter) { innerVlanId.YFilter = yf }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetSegmentPath() string {
    return "inner-vlan-id"
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = innerVlanId.Limit
    leafs["request-period"] = innerVlanId.RequestPeriod
    leafs["blocking-period"] = innerVlanId.BlockingPeriod
    return leafs
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetBundleName() string { return "cisco_ios_xr" }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetYangName() string { return "inner-vlan-id" }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) SetParent(parent types.Entity) { innerVlanId.parent = parent }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetParent() types.Entity { return innerVlanId.parent }

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId
// VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetFilter() yfilter.YFilter { return vlanId.YFilter }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) SetFilter(yf yfilter.YFilter) { vlanId.YFilter = yf }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetSegmentPath() string {
    return "vlan-id"
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = vlanId.Limit
    leafs["request-period"] = vlanId.RequestPeriod
    leafs["blocking-period"] = vlanId.BlockingPeriod
    return leafs
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetBundleName() string { return "cisco_ios_xr" }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetYangName() string { return "vlan-id" }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) SetParent(parent types.Entity) { vlanId.parent = parent }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetParent() types.Entity { return vlanId.parent }

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetParentYangName() string { return "throttle-config" }

// Pppoe_Nodes_Node_SummaryTotal
// PPPoE statistics for a given node
type Pppoe_Nodes_Node_SummaryTotal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ready Access Interface Count. The type is interface{} with range:
    // 0..4294967295.
    ReadyAccessInterfaces interface{}

    // Not Ready Access Interface Count. The type is interface{} with range:
    // 0..4294967295.
    NotReadyAccessInterfaces interface{}

    // Complete Session Count. The type is interface{} with range: 0..4294967295.
    CompleteSessions interface{}

    // Incomplete Session Count. The type is interface{} with range:
    // 0..4294967295.
    IncompleteSessions interface{}

    // Flow Control credit limit. The type is interface{} with range:
    // 0..4294967295.
    FlowControlLimit interface{}

    // Flow Control In-Flight Count. The type is interface{} with range:
    // 0..4294967295.
    FlowControlInFlightSessions interface{}

    // Flow Control Drop Count. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowControlDroppedSessions interface{}

    // Flow Control Disconnected Count. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowControlDisconnectedSessions interface{}

    // Flow Control Success Count, sessions completing call flow. The type is
    // interface{} with range: 0..18446744073709551615.
    FlowControlSuccessfulSessions interface{}

    // PPPoEMASubscriberInfraFlowControl. The type is interface{} with range:
    // 0..4294967295.
    PppoemaSubscriberInfraFlowControl interface{}
}

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetFilter() yfilter.YFilter { return summaryTotal.YFilter }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) SetFilter(yf yfilter.YFilter) { summaryTotal.YFilter = yf }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetGoName(yname string) string {
    if yname == "ready-access-interfaces" { return "ReadyAccessInterfaces" }
    if yname == "not-ready-access-interfaces" { return "NotReadyAccessInterfaces" }
    if yname == "complete-sessions" { return "CompleteSessions" }
    if yname == "incomplete-sessions" { return "IncompleteSessions" }
    if yname == "flow-control-limit" { return "FlowControlLimit" }
    if yname == "flow-control-in-flight-sessions" { return "FlowControlInFlightSessions" }
    if yname == "flow-control-dropped-sessions" { return "FlowControlDroppedSessions" }
    if yname == "flow-control-disconnected-sessions" { return "FlowControlDisconnectedSessions" }
    if yname == "flow-control-successful-sessions" { return "FlowControlSuccessfulSessions" }
    if yname == "pppoema-subscriber-infra-flow-control" { return "PppoemaSubscriberInfraFlowControl" }
    return ""
}

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetSegmentPath() string {
    return "summary-total"
}

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ready-access-interfaces"] = summaryTotal.ReadyAccessInterfaces
    leafs["not-ready-access-interfaces"] = summaryTotal.NotReadyAccessInterfaces
    leafs["complete-sessions"] = summaryTotal.CompleteSessions
    leafs["incomplete-sessions"] = summaryTotal.IncompleteSessions
    leafs["flow-control-limit"] = summaryTotal.FlowControlLimit
    leafs["flow-control-in-flight-sessions"] = summaryTotal.FlowControlInFlightSessions
    leafs["flow-control-dropped-sessions"] = summaryTotal.FlowControlDroppedSessions
    leafs["flow-control-disconnected-sessions"] = summaryTotal.FlowControlDisconnectedSessions
    leafs["flow-control-successful-sessions"] = summaryTotal.FlowControlSuccessfulSessions
    leafs["pppoema-subscriber-infra-flow-control"] = summaryTotal.PppoemaSubscriberInfraFlowControl
    return leafs
}

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetBundleName() string { return "cisco_ios_xr" }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetYangName() string { return "summary-total" }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) SetParent(parent types.Entity) { summaryTotal.parent = parent }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetParent() types.Entity { return summaryTotal.parent }

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetParentYangName() string { return "node" }

