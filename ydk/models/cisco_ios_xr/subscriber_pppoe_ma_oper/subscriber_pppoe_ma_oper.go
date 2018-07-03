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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE access interface statistics information.
    AccessInterfaceStatistics Pppoe_AccessInterfaceStatistics

    // Per-node PPPoE operational data.
    Nodes Pppoe_Nodes
}

func (pppoe *Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-pppoe-ma-oper"
    pppoe.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-pppoe-ma-oper:pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Children.Append("access-interface-statistics", types.YChild{"AccessInterfaceStatistics", &pppoe.AccessInterfaceStatistics})
    pppoe.EntityData.Children.Append("nodes", types.YChild{"Nodes", &pppoe.Nodes})
    pppoe.EntityData.Leafs = types.NewOrderedMap()

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Pppoe_AccessInterfaceStatistics
// PPPoE access interface statistics information
type Pppoe_AccessInterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics information for a PPPoE-enabled access interface. The type is
    // slice of Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic.
    AccessInterfaceStatistic []*Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic
}

func (accessInterfaceStatistics *Pppoe_AccessInterfaceStatistics) GetEntityData() *types.CommonEntityData {
    accessInterfaceStatistics.EntityData.YFilter = accessInterfaceStatistics.YFilter
    accessInterfaceStatistics.EntityData.YangName = "access-interface-statistics"
    accessInterfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceStatistics.EntityData.ParentYangName = "pppoe"
    accessInterfaceStatistics.EntityData.SegmentPath = "access-interface-statistics"
    accessInterfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceStatistics.EntityData.Children = types.NewOrderedMap()
    accessInterfaceStatistics.EntityData.Children.Append("access-interface-statistic", types.YChild{"AccessInterfaceStatistic", nil})
    for i := range accessInterfaceStatistics.AccessInterfaceStatistic {
        accessInterfaceStatistics.EntityData.Children.Append(types.GetSegmentPath(accessInterfaceStatistics.AccessInterfaceStatistic[i]), types.YChild{"AccessInterfaceStatistic", accessInterfaceStatistics.AccessInterfaceStatistic[i]})
    }
    accessInterfaceStatistics.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaceStatistics.EntityData.YListKeys = []string {}

    return &(accessInterfaceStatistics.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic
// Statistics information for a PPPoE-enabled
// access interface
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. PPPoE Access Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Packet Counts.
    PacketCounts Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts
}

func (accessInterfaceStatistic *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic) GetEntityData() *types.CommonEntityData {
    accessInterfaceStatistic.EntityData.YFilter = accessInterfaceStatistic.YFilter
    accessInterfaceStatistic.EntityData.YangName = "access-interface-statistic"
    accessInterfaceStatistic.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceStatistic.EntityData.ParentYangName = "access-interface-statistics"
    accessInterfaceStatistic.EntityData.SegmentPath = "access-interface-statistic" + types.AddKeyToken(accessInterfaceStatistic.InterfaceName, "interface-name")
    accessInterfaceStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceStatistic.EntityData.Children = types.NewOrderedMap()
    accessInterfaceStatistic.EntityData.Children.Append("packet-counts", types.YChild{"PacketCounts", &accessInterfaceStatistic.PacketCounts})
    accessInterfaceStatistic.EntityData.Leafs = types.NewOrderedMap()
    accessInterfaceStatistic.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterfaceStatistic.InterfaceName})

    accessInterfaceStatistic.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterfaceStatistic.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts
// Packet Counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts struct {
    EntityData types.CommonEntityData
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

func (packetCounts *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts) GetEntityData() *types.CommonEntityData {
    packetCounts.EntityData.YFilter = packetCounts.YFilter
    packetCounts.EntityData.YangName = "packet-counts"
    packetCounts.EntityData.BundleName = "cisco_ios_xr"
    packetCounts.EntityData.ParentYangName = "access-interface-statistic"
    packetCounts.EntityData.SegmentPath = "packet-counts"
    packetCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetCounts.EntityData.Children = types.NewOrderedMap()
    packetCounts.EntityData.Children.Append("padi", types.YChild{"Padi", &packetCounts.Padi})
    packetCounts.EntityData.Children.Append("pado", types.YChild{"Pado", &packetCounts.Pado})
    packetCounts.EntityData.Children.Append("padr", types.YChild{"Padr", &packetCounts.Padr})
    packetCounts.EntityData.Children.Append("pads-success", types.YChild{"PadsSuccess", &packetCounts.PadsSuccess})
    packetCounts.EntityData.Children.Append("pads-error", types.YChild{"PadsError", &packetCounts.PadsError})
    packetCounts.EntityData.Children.Append("padt", types.YChild{"Padt", &packetCounts.Padt})
    packetCounts.EntityData.Children.Append("session-state", types.YChild{"SessionState", &packetCounts.SessionState})
    packetCounts.EntityData.Children.Append("other", types.YChild{"Other", &packetCounts.Other})
    packetCounts.EntityData.Leafs = types.NewOrderedMap()

    packetCounts.EntityData.YListKeys = []string {}

    return &(packetCounts.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi
// PADI counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padi *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padi) GetEntityData() *types.CommonEntityData {
    padi.EntityData.YFilter = padi.YFilter
    padi.EntityData.YangName = "padi"
    padi.EntityData.BundleName = "cisco_ios_xr"
    padi.EntityData.ParentYangName = "packet-counts"
    padi.EntityData.SegmentPath = "padi"
    padi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padi.EntityData.Children = types.NewOrderedMap()
    padi.EntityData.Leafs = types.NewOrderedMap()
    padi.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padi.Sent})
    padi.EntityData.Leafs.Append("received", types.YLeaf{"Received", padi.Received})
    padi.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padi.Dropped})

    padi.EntityData.YListKeys = []string {}

    return &(padi.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado
// PADO counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (pado *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Pado) GetEntityData() *types.CommonEntityData {
    pado.EntityData.YFilter = pado.YFilter
    pado.EntityData.YangName = "pado"
    pado.EntityData.BundleName = "cisco_ios_xr"
    pado.EntityData.ParentYangName = "packet-counts"
    pado.EntityData.SegmentPath = "pado"
    pado.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pado.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pado.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pado.EntityData.Children = types.NewOrderedMap()
    pado.EntityData.Leafs = types.NewOrderedMap()
    pado.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", pado.Sent})
    pado.EntityData.Leafs.Append("received", types.YLeaf{"Received", pado.Received})
    pado.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", pado.Dropped})

    pado.EntityData.YListKeys = []string {}

    return &(pado.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr
// PADR counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padr *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padr) GetEntityData() *types.CommonEntityData {
    padr.EntityData.YFilter = padr.YFilter
    padr.EntityData.YangName = "padr"
    padr.EntityData.BundleName = "cisco_ios_xr"
    padr.EntityData.ParentYangName = "packet-counts"
    padr.EntityData.SegmentPath = "padr"
    padr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padr.EntityData.Children = types.NewOrderedMap()
    padr.EntityData.Leafs = types.NewOrderedMap()
    padr.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padr.Sent})
    padr.EntityData.Leafs.Append("received", types.YLeaf{"Received", padr.Received})
    padr.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padr.Dropped})

    padr.EntityData.YListKeys = []string {}

    return &(padr.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess
// PADS Success counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsSuccess *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsSuccess) GetEntityData() *types.CommonEntityData {
    padsSuccess.EntityData.YFilter = padsSuccess.YFilter
    padsSuccess.EntityData.YangName = "pads-success"
    padsSuccess.EntityData.BundleName = "cisco_ios_xr"
    padsSuccess.EntityData.ParentYangName = "packet-counts"
    padsSuccess.EntityData.SegmentPath = "pads-success"
    padsSuccess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padsSuccess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padsSuccess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padsSuccess.EntityData.Children = types.NewOrderedMap()
    padsSuccess.EntityData.Leafs = types.NewOrderedMap()
    padsSuccess.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padsSuccess.Sent})
    padsSuccess.EntityData.Leafs.Append("received", types.YLeaf{"Received", padsSuccess.Received})
    padsSuccess.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padsSuccess.Dropped})

    padsSuccess.EntityData.YListKeys = []string {}

    return &(padsSuccess.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError
// PADS Error counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsError *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_PadsError) GetEntityData() *types.CommonEntityData {
    padsError.EntityData.YFilter = padsError.YFilter
    padsError.EntityData.YangName = "pads-error"
    padsError.EntityData.BundleName = "cisco_ios_xr"
    padsError.EntityData.ParentYangName = "packet-counts"
    padsError.EntityData.SegmentPath = "pads-error"
    padsError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padsError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padsError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padsError.EntityData.Children = types.NewOrderedMap()
    padsError.EntityData.Leafs = types.NewOrderedMap()
    padsError.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padsError.Sent})
    padsError.EntityData.Leafs.Append("received", types.YLeaf{"Received", padsError.Received})
    padsError.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padsError.Dropped})

    padsError.EntityData.YListKeys = []string {}

    return &(padsError.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt
// PADT counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padt *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Padt) GetEntityData() *types.CommonEntityData {
    padt.EntityData.YFilter = padt.YFilter
    padt.EntityData.YangName = "padt"
    padt.EntityData.BundleName = "cisco_ios_xr"
    padt.EntityData.ParentYangName = "packet-counts"
    padt.EntityData.SegmentPath = "padt"
    padt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padt.EntityData.Children = types.NewOrderedMap()
    padt.EntityData.Leafs = types.NewOrderedMap()
    padt.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padt.Sent})
    padt.EntityData.Leafs.Append("received", types.YLeaf{"Received", padt.Received})
    padt.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padt.Dropped})

    padt.EntityData.YListKeys = []string {}

    return &(padt.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState
// Session Stage counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (sessionState *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "packet-counts"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", sessionState.Sent})
    sessionState.EntityData.Leafs.Append("received", types.YLeaf{"Received", sessionState.Received})
    sessionState.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", sessionState.Dropped})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other
// Other counts
type Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (other *Pppoe_AccessInterfaceStatistics_AccessInterfaceStatistic_PacketCounts_Other) GetEntityData() *types.CommonEntityData {
    other.EntityData.YFilter = other.YFilter
    other.EntityData.YangName = "other"
    other.EntityData.BundleName = "cisco_ios_xr"
    other.EntityData.ParentYangName = "packet-counts"
    other.EntityData.SegmentPath = "other"
    other.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    other.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    other.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    other.EntityData.Children = types.NewOrderedMap()
    other.EntityData.Leafs = types.NewOrderedMap()
    other.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", other.Sent})
    other.EntityData.Leafs.Append("received", types.YLeaf{"Received", other.Received})
    other.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", other.Dropped})

    other.EntityData.YListKeys = []string {}

    return &(other.EntityData)
}

// Pppoe_Nodes
// Per-node PPPoE operational data
type Pppoe_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE operational data for a particular node. The type is slice of
    // Pppoe_Nodes_Node.
    Node []*Pppoe_Nodes_Node
}

func (nodes *Pppoe_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "pppoe"
    nodes.EntityData.SegmentPath = "nodes"
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

// Pppoe_Nodes_Node
// PPPoE operational data for a particular node
type Pppoe_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Pppoe_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", &node.AccessInterface})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("bba-groups", types.YChild{"BbaGroups", &node.BbaGroups})
    node.EntityData.Children.Append("summary-total", types.YChild{"SummaryTotal", &node.SummaryTotal})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Pppoe_Nodes_Node_Statistics
// PPPoE statistics for a given node
type Pppoe_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet Counts.
    PacketCounts Pppoe_Nodes_Node_Statistics_PacketCounts

    // Packet Error Counts.
    PacketErrorCounts Pppoe_Nodes_Node_Statistics_PacketErrorCounts
}

func (statistics *Pppoe_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("packet-counts", types.YChild{"PacketCounts", &statistics.PacketCounts})
    statistics.EntityData.Children.Append("packet-error-counts", types.YChild{"PacketErrorCounts", &statistics.PacketErrorCounts})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts
// Packet Counts
type Pppoe_Nodes_Node_Statistics_PacketCounts struct {
    EntityData types.CommonEntityData
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

func (packetCounts *Pppoe_Nodes_Node_Statistics_PacketCounts) GetEntityData() *types.CommonEntityData {
    packetCounts.EntityData.YFilter = packetCounts.YFilter
    packetCounts.EntityData.YangName = "packet-counts"
    packetCounts.EntityData.BundleName = "cisco_ios_xr"
    packetCounts.EntityData.ParentYangName = "statistics"
    packetCounts.EntityData.SegmentPath = "packet-counts"
    packetCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetCounts.EntityData.Children = types.NewOrderedMap()
    packetCounts.EntityData.Children.Append("padi", types.YChild{"Padi", &packetCounts.Padi})
    packetCounts.EntityData.Children.Append("pado", types.YChild{"Pado", &packetCounts.Pado})
    packetCounts.EntityData.Children.Append("padr", types.YChild{"Padr", &packetCounts.Padr})
    packetCounts.EntityData.Children.Append("pads-success", types.YChild{"PadsSuccess", &packetCounts.PadsSuccess})
    packetCounts.EntityData.Children.Append("pads-error", types.YChild{"PadsError", &packetCounts.PadsError})
    packetCounts.EntityData.Children.Append("padt", types.YChild{"Padt", &packetCounts.Padt})
    packetCounts.EntityData.Children.Append("session-state", types.YChild{"SessionState", &packetCounts.SessionState})
    packetCounts.EntityData.Children.Append("other", types.YChild{"Other", &packetCounts.Other})
    packetCounts.EntityData.Leafs = types.NewOrderedMap()

    packetCounts.EntityData.YListKeys = []string {}

    return &(packetCounts.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_Padi
// PADI counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Padi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padi *Pppoe_Nodes_Node_Statistics_PacketCounts_Padi) GetEntityData() *types.CommonEntityData {
    padi.EntityData.YFilter = padi.YFilter
    padi.EntityData.YangName = "padi"
    padi.EntityData.BundleName = "cisco_ios_xr"
    padi.EntityData.ParentYangName = "packet-counts"
    padi.EntityData.SegmentPath = "padi"
    padi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padi.EntityData.Children = types.NewOrderedMap()
    padi.EntityData.Leafs = types.NewOrderedMap()
    padi.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padi.Sent})
    padi.EntityData.Leafs.Append("received", types.YLeaf{"Received", padi.Received})
    padi.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padi.Dropped})

    padi.EntityData.YListKeys = []string {}

    return &(padi.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_Pado
// PADO counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Pado struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (pado *Pppoe_Nodes_Node_Statistics_PacketCounts_Pado) GetEntityData() *types.CommonEntityData {
    pado.EntityData.YFilter = pado.YFilter
    pado.EntityData.YangName = "pado"
    pado.EntityData.BundleName = "cisco_ios_xr"
    pado.EntityData.ParentYangName = "packet-counts"
    pado.EntityData.SegmentPath = "pado"
    pado.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pado.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pado.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pado.EntityData.Children = types.NewOrderedMap()
    pado.EntityData.Leafs = types.NewOrderedMap()
    pado.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", pado.Sent})
    pado.EntityData.Leafs.Append("received", types.YLeaf{"Received", pado.Received})
    pado.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", pado.Dropped})

    pado.EntityData.YListKeys = []string {}

    return &(pado.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_Padr
// PADR counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Padr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padr *Pppoe_Nodes_Node_Statistics_PacketCounts_Padr) GetEntityData() *types.CommonEntityData {
    padr.EntityData.YFilter = padr.YFilter
    padr.EntityData.YangName = "padr"
    padr.EntityData.BundleName = "cisco_ios_xr"
    padr.EntityData.ParentYangName = "packet-counts"
    padr.EntityData.SegmentPath = "padr"
    padr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padr.EntityData.Children = types.NewOrderedMap()
    padr.EntityData.Leafs = types.NewOrderedMap()
    padr.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padr.Sent})
    padr.EntityData.Leafs.Append("received", types.YLeaf{"Received", padr.Received})
    padr.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padr.Dropped})

    padr.EntityData.YListKeys = []string {}

    return &(padr.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess
// PADS Success counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsSuccess *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsSuccess) GetEntityData() *types.CommonEntityData {
    padsSuccess.EntityData.YFilter = padsSuccess.YFilter
    padsSuccess.EntityData.YangName = "pads-success"
    padsSuccess.EntityData.BundleName = "cisco_ios_xr"
    padsSuccess.EntityData.ParentYangName = "packet-counts"
    padsSuccess.EntityData.SegmentPath = "pads-success"
    padsSuccess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padsSuccess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padsSuccess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padsSuccess.EntityData.Children = types.NewOrderedMap()
    padsSuccess.EntityData.Leafs = types.NewOrderedMap()
    padsSuccess.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padsSuccess.Sent})
    padsSuccess.EntityData.Leafs.Append("received", types.YLeaf{"Received", padsSuccess.Received})
    padsSuccess.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padsSuccess.Dropped})

    padsSuccess.EntityData.YListKeys = []string {}

    return &(padsSuccess.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError
// PADS Error counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padsError *Pppoe_Nodes_Node_Statistics_PacketCounts_PadsError) GetEntityData() *types.CommonEntityData {
    padsError.EntityData.YFilter = padsError.YFilter
    padsError.EntityData.YangName = "pads-error"
    padsError.EntityData.BundleName = "cisco_ios_xr"
    padsError.EntityData.ParentYangName = "packet-counts"
    padsError.EntityData.SegmentPath = "pads-error"
    padsError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padsError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padsError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padsError.EntityData.Children = types.NewOrderedMap()
    padsError.EntityData.Leafs = types.NewOrderedMap()
    padsError.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padsError.Sent})
    padsError.EntityData.Leafs.Append("received", types.YLeaf{"Received", padsError.Received})
    padsError.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padsError.Dropped})

    padsError.EntityData.YListKeys = []string {}

    return &(padsError.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_Padt
// PADT counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Padt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (padt *Pppoe_Nodes_Node_Statistics_PacketCounts_Padt) GetEntityData() *types.CommonEntityData {
    padt.EntityData.YFilter = padt.YFilter
    padt.EntityData.YangName = "padt"
    padt.EntityData.BundleName = "cisco_ios_xr"
    padt.EntityData.ParentYangName = "packet-counts"
    padt.EntityData.SegmentPath = "padt"
    padt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padt.EntityData.Children = types.NewOrderedMap()
    padt.EntityData.Leafs = types.NewOrderedMap()
    padt.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", padt.Sent})
    padt.EntityData.Leafs.Append("received", types.YLeaf{"Received", padt.Received})
    padt.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", padt.Dropped})

    padt.EntityData.YListKeys = []string {}

    return &(padt.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState
// Session Stage counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (sessionState *Pppoe_Nodes_Node_Statistics_PacketCounts_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "packet-counts"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", sessionState.Sent})
    sessionState.EntityData.Leafs.Append("received", types.YLeaf{"Received", sessionState.Received})
    sessionState.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", sessionState.Dropped})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketCounts_Other
// Other counts
type Pppoe_Nodes_Node_Statistics_PacketCounts_Other struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sent. The type is interface{} with range: 0..4294967295.
    Sent interface{}

    // Received. The type is interface{} with range: 0..4294967295.
    Received interface{}

    // Dropped. The type is interface{} with range: 0..4294967295.
    Dropped interface{}
}

func (other *Pppoe_Nodes_Node_Statistics_PacketCounts_Other) GetEntityData() *types.CommonEntityData {
    other.EntityData.YFilter = other.YFilter
    other.EntityData.YangName = "other"
    other.EntityData.BundleName = "cisco_ios_xr"
    other.EntityData.ParentYangName = "packet-counts"
    other.EntityData.SegmentPath = "other"
    other.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    other.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    other.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    other.EntityData.Children = types.NewOrderedMap()
    other.EntityData.Leafs = types.NewOrderedMap()
    other.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", other.Sent})
    other.EntityData.Leafs.Append("received", types.YLeaf{"Received", other.Received})
    other.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", other.Dropped})

    other.EntityData.YListKeys = []string {}

    return &(other.EntityData)
}

// Pppoe_Nodes_Node_Statistics_PacketErrorCounts
// Packet Error Counts
type Pppoe_Nodes_Node_Statistics_PacketErrorCounts struct {
    EntityData types.CommonEntityData
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

func (packetErrorCounts *Pppoe_Nodes_Node_Statistics_PacketErrorCounts) GetEntityData() *types.CommonEntityData {
    packetErrorCounts.EntityData.YFilter = packetErrorCounts.YFilter
    packetErrorCounts.EntityData.YangName = "packet-error-counts"
    packetErrorCounts.EntityData.BundleName = "cisco_ios_xr"
    packetErrorCounts.EntityData.ParentYangName = "statistics"
    packetErrorCounts.EntityData.SegmentPath = "packet-error-counts"
    packetErrorCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetErrorCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetErrorCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetErrorCounts.EntityData.Children = types.NewOrderedMap()
    packetErrorCounts.EntityData.Leafs = types.NewOrderedMap()
    packetErrorCounts.EntityData.Leafs.Append("no-interface-handle", types.YLeaf{"NoInterfaceHandle", packetErrorCounts.NoInterfaceHandle})
    packetErrorCounts.EntityData.Leafs.Append("no-packet-payload", types.YLeaf{"NoPacketPayload", packetErrorCounts.NoPacketPayload})
    packetErrorCounts.EntityData.Leafs.Append("no-packet-mac-address", types.YLeaf{"NoPacketMacAddress", packetErrorCounts.NoPacketMacAddress})
    packetErrorCounts.EntityData.Leafs.Append("invalid-version-type-value", types.YLeaf{"InvalidVersionTypeValue", packetErrorCounts.InvalidVersionTypeValue})
    packetErrorCounts.EntityData.Leafs.Append("bad-packet-length", types.YLeaf{"BadPacketLength", packetErrorCounts.BadPacketLength})
    packetErrorCounts.EntityData.Leafs.Append("unknown-interface", types.YLeaf{"UnknownInterface", packetErrorCounts.UnknownInterface})
    packetErrorCounts.EntityData.Leafs.Append("pado-received", types.YLeaf{"PadoReceived", packetErrorCounts.PadoReceived})
    packetErrorCounts.EntityData.Leafs.Append("pads-received", types.YLeaf{"PadsReceived", packetErrorCounts.PadsReceived})
    packetErrorCounts.EntityData.Leafs.Append("unknown-packet-type-received", types.YLeaf{"UnknownPacketTypeReceived", packetErrorCounts.UnknownPacketTypeReceived})
    packetErrorCounts.EntityData.Leafs.Append("unexpected-session-id-in-packet", types.YLeaf{"UnexpectedSessionIdInPacket", packetErrorCounts.UnexpectedSessionIdInPacket})
    packetErrorCounts.EntityData.Leafs.Append("no-service-name-tag", types.YLeaf{"NoServiceNameTag", packetErrorCounts.NoServiceNameTag})
    packetErrorCounts.EntityData.Leafs.Append("padt-for-unknown-session", types.YLeaf{"PadtForUnknownSession", packetErrorCounts.PadtForUnknownSession})
    packetErrorCounts.EntityData.Leafs.Append("padt-with-wrong-peer-mac", types.YLeaf{"PadtWithWrongPeerMac", packetErrorCounts.PadtWithWrongPeerMac})
    packetErrorCounts.EntityData.Leafs.Append("padt-with-wrong-vlan-tags", types.YLeaf{"PadtWithWrongVlanTags", packetErrorCounts.PadtWithWrongVlanTags})
    packetErrorCounts.EntityData.Leafs.Append("zero-length-host-uniq", types.YLeaf{"ZeroLengthHostUniq", packetErrorCounts.ZeroLengthHostUniq})
    packetErrorCounts.EntityData.Leafs.Append("padt-before-pads-sent", types.YLeaf{"PadtBeforePadsSent", packetErrorCounts.PadtBeforePadsSent})
    packetErrorCounts.EntityData.Leafs.Append("session-stage-packet-for-unknown-session", types.YLeaf{"SessionStagePacketForUnknownSession", packetErrorCounts.SessionStagePacketForUnknownSession})
    packetErrorCounts.EntityData.Leafs.Append("session-stage-packet-with-wrong-mac", types.YLeaf{"SessionStagePacketWithWrongMac", packetErrorCounts.SessionStagePacketWithWrongMac})
    packetErrorCounts.EntityData.Leafs.Append("session-stage-packet-with-wrong-vlan-tags", types.YLeaf{"SessionStagePacketWithWrongVlanTags", packetErrorCounts.SessionStagePacketWithWrongVlanTags})
    packetErrorCounts.EntityData.Leafs.Append("session-stage-packet-with-no-error", types.YLeaf{"SessionStagePacketWithNoError", packetErrorCounts.SessionStagePacketWithNoError})
    packetErrorCounts.EntityData.Leafs.Append("tag-too-short", types.YLeaf{"TagTooShort", packetErrorCounts.TagTooShort})
    packetErrorCounts.EntityData.Leafs.Append("bad-tag-length-field", types.YLeaf{"BadTagLengthField", packetErrorCounts.BadTagLengthField})
    packetErrorCounts.EntityData.Leafs.Append("multiple-service-name-tags", types.YLeaf{"MultipleServiceNameTags", packetErrorCounts.MultipleServiceNameTags})
    packetErrorCounts.EntityData.Leafs.Append("multiple-max-payload-tags", types.YLeaf{"MultipleMaxPayloadTags", packetErrorCounts.MultipleMaxPayloadTags})
    packetErrorCounts.EntityData.Leafs.Append("invalid-max-payload-tag", types.YLeaf{"InvalidMaxPayloadTag", packetErrorCounts.InvalidMaxPayloadTag})
    packetErrorCounts.EntityData.Leafs.Append("multiple-vendor-specific-tags", types.YLeaf{"MultipleVendorSpecificTags", packetErrorCounts.MultipleVendorSpecificTags})
    packetErrorCounts.EntityData.Leafs.Append("unexpected-ac-name-tag", types.YLeaf{"UnexpectedAcNameTag", packetErrorCounts.UnexpectedAcNameTag})
    packetErrorCounts.EntityData.Leafs.Append("unexpected-error-tags", types.YLeaf{"UnexpectedErrorTags", packetErrorCounts.UnexpectedErrorTags})
    packetErrorCounts.EntityData.Leafs.Append("unknown-tag-received", types.YLeaf{"UnknownTagReceived", packetErrorCounts.UnknownTagReceived})
    packetErrorCounts.EntityData.Leafs.Append("no-iana-code-invendor-tag", types.YLeaf{"NoIanaCodeInvendorTag", packetErrorCounts.NoIanaCodeInvendorTag})
    packetErrorCounts.EntityData.Leafs.Append("invalid-iana-code-invendor-tag", types.YLeaf{"InvalidIanaCodeInvendorTag", packetErrorCounts.InvalidIanaCodeInvendorTag})
    packetErrorCounts.EntityData.Leafs.Append("vendor-tag-too-short", types.YLeaf{"VendorTagTooShort", packetErrorCounts.VendorTagTooShort})
    packetErrorCounts.EntityData.Leafs.Append("bad-vendor-tag-length-field", types.YLeaf{"BadVendorTagLengthField", packetErrorCounts.BadVendorTagLengthField})
    packetErrorCounts.EntityData.Leafs.Append("multiple-host-uniq-tags", types.YLeaf{"MultipleHostUniqTags", packetErrorCounts.MultipleHostUniqTags})
    packetErrorCounts.EntityData.Leafs.Append("multiple-relay-session-id-tags", types.YLeaf{"MultipleRelaySessionIdTags", packetErrorCounts.MultipleRelaySessionIdTags})
    packetErrorCounts.EntityData.Leafs.Append("multiple-circuit-id-tags", types.YLeaf{"MultipleCircuitIdTags", packetErrorCounts.MultipleCircuitIdTags})
    packetErrorCounts.EntityData.Leafs.Append("multiple-remote-id-tags", types.YLeaf{"MultipleRemoteIdTags", packetErrorCounts.MultipleRemoteIdTags})
    packetErrorCounts.EntityData.Leafs.Append("invalid-dsl-tag", types.YLeaf{"InvalidDslTag", packetErrorCounts.InvalidDslTag})
    packetErrorCounts.EntityData.Leafs.Append("multiple-of-the-same-dsl-tag", types.YLeaf{"MultipleOfTheSameDslTag", packetErrorCounts.MultipleOfTheSameDslTag})
    packetErrorCounts.EntityData.Leafs.Append("invalid-iwf-tag", types.YLeaf{"InvalidIwfTag", packetErrorCounts.InvalidIwfTag})
    packetErrorCounts.EntityData.Leafs.Append("multiple-iwf-tags", types.YLeaf{"MultipleIwfTags", packetErrorCounts.MultipleIwfTags})
    packetErrorCounts.EntityData.Leafs.Append("unknownvendor-tag", types.YLeaf{"UnknownvendorTag", packetErrorCounts.UnknownvendorTag})
    packetErrorCounts.EntityData.Leafs.Append("no-space-left-in-packet", types.YLeaf{"NoSpaceLeftInPacket", packetErrorCounts.NoSpaceLeftInPacket})
    packetErrorCounts.EntityData.Leafs.Append("duplicate-host-uniq-tag-received", types.YLeaf{"DuplicateHostUniqTagReceived", packetErrorCounts.DuplicateHostUniqTagReceived})
    packetErrorCounts.EntityData.Leafs.Append("duplicate-relay-session-id-tag-received", types.YLeaf{"DuplicateRelaySessionIdTagReceived", packetErrorCounts.DuplicateRelaySessionIdTagReceived})
    packetErrorCounts.EntityData.Leafs.Append("packet-too-long", types.YLeaf{"PacketTooLong", packetErrorCounts.PacketTooLong})
    packetErrorCounts.EntityData.Leafs.Append("invalid-ale-tag", types.YLeaf{"InvalidAleTag", packetErrorCounts.InvalidAleTag})
    packetErrorCounts.EntityData.Leafs.Append("multiple-ale-tags", types.YLeaf{"MultipleAleTags", packetErrorCounts.MultipleAleTags})
    packetErrorCounts.EntityData.Leafs.Append("invalid-service-name", types.YLeaf{"InvalidServiceName", packetErrorCounts.InvalidServiceName})
    packetErrorCounts.EntityData.Leafs.Append("invalid-peer-mac", types.YLeaf{"InvalidPeerMac", packetErrorCounts.InvalidPeerMac})
    packetErrorCounts.EntityData.Leafs.Append("invalid-vlan-tags", types.YLeaf{"InvalidVlanTags", packetErrorCounts.InvalidVlanTags})
    packetErrorCounts.EntityData.Leafs.Append("packet-on-srg-slave", types.YLeaf{"PacketOnSrgSlave", packetErrorCounts.PacketOnSrgSlave})

    packetErrorCounts.EntityData.YListKeys = []string {}

    return &(packetErrorCounts.EntityData)
}

// Pppoe_Nodes_Node_AccessInterface
// PPPoE access interface information
type Pppoe_Nodes_Node_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE access interface summary information.
    Summaries Pppoe_Nodes_Node_AccessInterface_Summaries
}

func (accessInterface *Pppoe_Nodes_Node_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "node"
    accessInterface.EntityData.SegmentPath = "access-interface"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Children.Append("summaries", types.YChild{"Summaries", &accessInterface.Summaries})
    accessInterface.EntityData.Leafs = types.NewOrderedMap()

    accessInterface.EntityData.YListKeys = []string {}

    return &(accessInterface.EntityData)
}

// Pppoe_Nodes_Node_AccessInterface_Summaries
// PPPoE access interface summary information
type Pppoe_Nodes_Node_AccessInterface_Summaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for a PPPoE-enabled access interface. The type is slice
    // of Pppoe_Nodes_Node_AccessInterface_Summaries_Summary.
    Summary []*Pppoe_Nodes_Node_AccessInterface_Summaries_Summary
}

func (summaries *Pppoe_Nodes_Node_AccessInterface_Summaries) GetEntityData() *types.CommonEntityData {
    summaries.EntityData.YFilter = summaries.YFilter
    summaries.EntityData.YangName = "summaries"
    summaries.EntityData.BundleName = "cisco_ios_xr"
    summaries.EntityData.ParentYangName = "access-interface"
    summaries.EntityData.SegmentPath = "summaries"
    summaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaries.EntityData.Children = types.NewOrderedMap()
    summaries.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range summaries.Summary {
        summaries.EntityData.Children.Append(types.GetSegmentPath(summaries.Summary[i]), types.YChild{"Summary", summaries.Summary[i]})
    }
    summaries.EntityData.Leafs = types.NewOrderedMap()

    summaries.EntityData.YListKeys = []string {}

    return &(summaries.EntityData)
}

// Pppoe_Nodes_Node_AccessInterface_Summaries_Summary
// Summary information for a PPPoE-enabled
// access interface
type Pppoe_Nodes_Node_AccessInterface_Summaries_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Pppoe_Nodes_Node_AccessInterface_Summaries_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "summaries"
    summary.EntityData.SegmentPath = "summary" + types.AddKeyToken(summary.InterfaceName, "interface-name")
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", summary.InterfaceName})
    summary.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", summary.InterfaceNameXr})
    summary.EntityData.Leafs.Append("interface-state", types.YLeaf{"InterfaceState", summary.InterfaceState})
    summary.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", summary.MacAddress})
    summary.EntityData.Leafs.Append("bba-group-name", types.YLeaf{"BbaGroupName", summary.BbaGroupName})
    summary.EntityData.Leafs.Append("is-ready", types.YLeaf{"IsReady", summary.IsReady})
    summary.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", summary.Sessions})
    summary.EntityData.Leafs.Append("incomplete-sessions", types.YLeaf{"IncompleteSessions", summary.IncompleteSessions})

    summary.EntityData.YListKeys = []string {"InterfaceName"}

    return &(summary.EntityData)
}

// Pppoe_Nodes_Node_Interfaces
// Per interface PPPoE operational data
type Pppoe_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for a PPPoE interface. The type is slice of
    // Pppoe_Nodes_Node_Interfaces_Interface.
    Interface []*Pppoe_Nodes_Node_Interfaces_Interface
}

func (interfaces *Pppoe_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Pppoe_Nodes_Node_Interfaces_Interface
// Data for a PPPoE interface
type Pppoe_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *Pppoe_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("tags", types.YChild{"Tags", &self.Tags})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr})
    self.EntityData.Leafs.Append("access-interface-name", types.YLeaf{"AccessInterfaceName", self.AccessInterfaceName})
    self.EntityData.Leafs.Append("bba-group-name", types.YLeaf{"BbaGroupName", self.BbaGroupName})
    self.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", self.SessionId})
    self.EntityData.Leafs.Append("local-mac-address", types.YLeaf{"LocalMacAddress", self.LocalMacAddress})
    self.EntityData.Leafs.Append("peer-mac-address", types.YLeaf{"PeerMacAddress", self.PeerMacAddress})
    self.EntityData.Leafs.Append("is-complete", types.YLeaf{"IsComplete", self.IsComplete})
    self.EntityData.Leafs.Append("vlan-outer-id", types.YLeaf{"VlanOuterId", self.VlanOuterId})
    self.EntityData.Leafs.Append("vlan-inner-id", types.YLeaf{"VlanInnerId", self.VlanInnerId})
    self.EntityData.Leafs.Append("srg-state", types.YLeaf{"SrgState", self.SrgState})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Pppoe_Nodes_Node_Interfaces_Interface_Tags
// Tags
type Pppoe_Nodes_Node_Interfaces_Interface_Tags struct {
    EntityData types.CommonEntityData
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

func (tags *Pppoe_Nodes_Node_Interfaces_Interface_Tags) GetEntityData() *types.CommonEntityData {
    tags.EntityData.YFilter = tags.YFilter
    tags.EntityData.YangName = "tags"
    tags.EntityData.BundleName = "cisco_ios_xr"
    tags.EntityData.ParentYangName = "interface"
    tags.EntityData.SegmentPath = "tags"
    tags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tags.EntityData.Children = types.NewOrderedMap()
    tags.EntityData.Children.Append("access-loop-encapsulation", types.YChild{"AccessLoopEncapsulation", &tags.AccessLoopEncapsulation})
    tags.EntityData.Leafs = types.NewOrderedMap()
    tags.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", tags.ServiceName})
    tags.EntityData.Leafs.Append("max-payload", types.YLeaf{"MaxPayload", tags.MaxPayload})
    tags.EntityData.Leafs.Append("host-uniq", types.YLeaf{"HostUniq", tags.HostUniq})
    tags.EntityData.Leafs.Append("relay-session-id", types.YLeaf{"RelaySessionId", tags.RelaySessionId})
    tags.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", tags.RemoteId})
    tags.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", tags.CircuitId})
    tags.EntityData.Leafs.Append("is-iwf", types.YLeaf{"IsIwf", tags.IsIwf})
    tags.EntityData.Leafs.Append("dsl-actual-up", types.YLeaf{"DslActualUp", tags.DslActualUp})
    tags.EntityData.Leafs.Append("dsl-actual-down", types.YLeaf{"DslActualDown", tags.DslActualDown})
    tags.EntityData.Leafs.Append("dsl-min-up", types.YLeaf{"DslMinUp", tags.DslMinUp})
    tags.EntityData.Leafs.Append("dsl-min-down", types.YLeaf{"DslMinDown", tags.DslMinDown})
    tags.EntityData.Leafs.Append("dsl-attain-up", types.YLeaf{"DslAttainUp", tags.DslAttainUp})
    tags.EntityData.Leafs.Append("dsl-attain-down", types.YLeaf{"DslAttainDown", tags.DslAttainDown})
    tags.EntityData.Leafs.Append("dsl-max-up", types.YLeaf{"DslMaxUp", tags.DslMaxUp})
    tags.EntityData.Leafs.Append("dsl-max-down", types.YLeaf{"DslMaxDown", tags.DslMaxDown})
    tags.EntityData.Leafs.Append("dsl-min-up-low", types.YLeaf{"DslMinUpLow", tags.DslMinUpLow})
    tags.EntityData.Leafs.Append("dsl-min-down-low", types.YLeaf{"DslMinDownLow", tags.DslMinDownLow})
    tags.EntityData.Leafs.Append("dsl-max-delay-up", types.YLeaf{"DslMaxDelayUp", tags.DslMaxDelayUp})
    tags.EntityData.Leafs.Append("dsl-actual-delay-up", types.YLeaf{"DslActualDelayUp", tags.DslActualDelayUp})
    tags.EntityData.Leafs.Append("dsl-max-delay-down", types.YLeaf{"DslMaxDelayDown", tags.DslMaxDelayDown})
    tags.EntityData.Leafs.Append("dsl-actual-delay-down", types.YLeaf{"DslActualDelayDown", tags.DslActualDelayDown})

    tags.EntityData.YListKeys = []string {}

    return &(tags.EntityData)
}

// Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation
// Access Loop Encapsulation
type Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data Link. The type is interface{} with range: 0..255.
    DataLink interface{}

    // Encaps 1. The type is interface{} with range: 0..255.
    Encaps1 interface{}

    // Encaps 2. The type is interface{} with range: 0..255.
    Encaps2 interface{}
}

func (accessLoopEncapsulation *Pppoe_Nodes_Node_Interfaces_Interface_Tags_AccessLoopEncapsulation) GetEntityData() *types.CommonEntityData {
    accessLoopEncapsulation.EntityData.YFilter = accessLoopEncapsulation.YFilter
    accessLoopEncapsulation.EntityData.YangName = "access-loop-encapsulation"
    accessLoopEncapsulation.EntityData.BundleName = "cisco_ios_xr"
    accessLoopEncapsulation.EntityData.ParentYangName = "tags"
    accessLoopEncapsulation.EntityData.SegmentPath = "access-loop-encapsulation"
    accessLoopEncapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessLoopEncapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessLoopEncapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessLoopEncapsulation.EntityData.Children = types.NewOrderedMap()
    accessLoopEncapsulation.EntityData.Leafs = types.NewOrderedMap()
    accessLoopEncapsulation.EntityData.Leafs.Append("data-link", types.YLeaf{"DataLink", accessLoopEncapsulation.DataLink})
    accessLoopEncapsulation.EntityData.Leafs.Append("encaps1", types.YLeaf{"Encaps1", accessLoopEncapsulation.Encaps1})
    accessLoopEncapsulation.EntityData.Leafs.Append("encaps2", types.YLeaf{"Encaps2", accessLoopEncapsulation.Encaps2})

    accessLoopEncapsulation.EntityData.YListKeys = []string {}

    return &(accessLoopEncapsulation.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups
// PPPoE BBA-Group information
type Pppoe_Nodes_Node_BbaGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE BBA-Group information. The type is slice of
    // Pppoe_Nodes_Node_BbaGroups_BbaGroup.
    BbaGroup []*Pppoe_Nodes_Node_BbaGroups_BbaGroup
}

func (bbaGroups *Pppoe_Nodes_Node_BbaGroups) GetEntityData() *types.CommonEntityData {
    bbaGroups.EntityData.YFilter = bbaGroups.YFilter
    bbaGroups.EntityData.YangName = "bba-groups"
    bbaGroups.EntityData.BundleName = "cisco_ios_xr"
    bbaGroups.EntityData.ParentYangName = "node"
    bbaGroups.EntityData.SegmentPath = "bba-groups"
    bbaGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bbaGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bbaGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bbaGroups.EntityData.Children = types.NewOrderedMap()
    bbaGroups.EntityData.Children.Append("bba-group", types.YChild{"BbaGroup", nil})
    for i := range bbaGroups.BbaGroup {
        bbaGroups.EntityData.Children.Append(types.GetSegmentPath(bbaGroups.BbaGroup[i]), types.YChild{"BbaGroup", bbaGroups.BbaGroup[i]})
    }
    bbaGroups.EntityData.Leafs = types.NewOrderedMap()

    bbaGroups.EntityData.YListKeys = []string {}

    return &(bbaGroups.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup
// PPPoE BBA-Group information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup struct {
    EntityData types.CommonEntityData
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

func (bbaGroup *Pppoe_Nodes_Node_BbaGroups_BbaGroup) GetEntityData() *types.CommonEntityData {
    bbaGroup.EntityData.YFilter = bbaGroup.YFilter
    bbaGroup.EntityData.YangName = "bba-group"
    bbaGroup.EntityData.BundleName = "cisco_ios_xr"
    bbaGroup.EntityData.ParentYangName = "bba-groups"
    bbaGroup.EntityData.SegmentPath = "bba-group" + types.AddKeyToken(bbaGroup.BbaGroupName, "bba-group-name")
    bbaGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bbaGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bbaGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bbaGroup.EntityData.Children = types.NewOrderedMap()
    bbaGroup.EntityData.Children.Append("limit-config", types.YChild{"LimitConfig", &bbaGroup.LimitConfig})
    bbaGroup.EntityData.Children.Append("limits", types.YChild{"Limits", &bbaGroup.Limits})
    bbaGroup.EntityData.Children.Append("throttles", types.YChild{"Throttles", &bbaGroup.Throttles})
    bbaGroup.EntityData.Children.Append("throttle-config", types.YChild{"ThrottleConfig", &bbaGroup.ThrottleConfig})
    bbaGroup.EntityData.Leafs = types.NewOrderedMap()
    bbaGroup.EntityData.Leafs.Append("bba-group-name", types.YLeaf{"BbaGroupName", bbaGroup.BbaGroupName})

    bbaGroup.EntityData.YListKeys = []string {"BbaGroupName"}

    return &(bbaGroup.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig
// BBA-Group limit configuration information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig struct {
    EntityData types.CommonEntityData
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

func (limitConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig) GetEntityData() *types.CommonEntityData {
    limitConfig.EntityData.YFilter = limitConfig.YFilter
    limitConfig.EntityData.YangName = "limit-config"
    limitConfig.EntityData.BundleName = "cisco_ios_xr"
    limitConfig.EntityData.ParentYangName = "bba-group"
    limitConfig.EntityData.SegmentPath = "limit-config"
    limitConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    limitConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    limitConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    limitConfig.EntityData.Children = types.NewOrderedMap()
    limitConfig.EntityData.Children.Append("card", types.YChild{"Card", &limitConfig.Card})
    limitConfig.EntityData.Children.Append("access-intf", types.YChild{"AccessIntf", &limitConfig.AccessIntf})
    limitConfig.EntityData.Children.Append("mac", types.YChild{"Mac", &limitConfig.Mac})
    limitConfig.EntityData.Children.Append("mac-iwf", types.YChild{"MacIwf", &limitConfig.MacIwf})
    limitConfig.EntityData.Children.Append("mac-access-interface", types.YChild{"MacAccessInterface", &limitConfig.MacAccessInterface})
    limitConfig.EntityData.Children.Append("mac-iwf-access-interface", types.YChild{"MacIwfAccessInterface", &limitConfig.MacIwfAccessInterface})
    limitConfig.EntityData.Children.Append("circuit-id", types.YChild{"CircuitId", &limitConfig.CircuitId})
    limitConfig.EntityData.Children.Append("remote-id", types.YChild{"RemoteId", &limitConfig.RemoteId})
    limitConfig.EntityData.Children.Append("circuit-id-and-remote-id", types.YChild{"CircuitIdAndRemoteId", &limitConfig.CircuitIdAndRemoteId})
    limitConfig.EntityData.Children.Append("outer-vlan-id", types.YChild{"OuterVlanId", &limitConfig.OuterVlanId})
    limitConfig.EntityData.Children.Append("inner-vlan-id", types.YChild{"InnerVlanId", &limitConfig.InnerVlanId})
    limitConfig.EntityData.Children.Append("vlan-id", types.YChild{"VlanId", &limitConfig.VlanId})
    limitConfig.EntityData.Leafs = types.NewOrderedMap()

    limitConfig.EntityData.YListKeys = []string {}

    return &(limitConfig.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card
// Card
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (card *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "limit-config"
    card.EntityData.SegmentPath = "card"
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Leafs = types.NewOrderedMap()
    card.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", card.MaxLimit})
    card.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", card.Threshold})
    card.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", card.RadiusOverrideEnabled})

    card.EntityData.YListKeys = []string {}

    return &(card.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf
// Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (accessIntf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_AccessIntf) GetEntityData() *types.CommonEntityData {
    accessIntf.EntityData.YFilter = accessIntf.YFilter
    accessIntf.EntityData.YangName = "access-intf"
    accessIntf.EntityData.BundleName = "cisco_ios_xr"
    accessIntf.EntityData.ParentYangName = "limit-config"
    accessIntf.EntityData.SegmentPath = "access-intf"
    accessIntf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessIntf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessIntf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessIntf.EntityData.Children = types.NewOrderedMap()
    accessIntf.EntityData.Leafs = types.NewOrderedMap()
    accessIntf.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", accessIntf.MaxLimit})
    accessIntf.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", accessIntf.Threshold})
    accessIntf.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", accessIntf.RadiusOverrideEnabled})

    accessIntf.EntityData.YListKeys = []string {}

    return &(accessIntf.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac
// MAC
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "limit-config"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", mac.MaxLimit})
    mac.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", mac.Threshold})
    mac.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", mac.RadiusOverrideEnabled})

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf
// MAC IWF
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (macIwf *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwf) GetEntityData() *types.CommonEntityData {
    macIwf.EntityData.YFilter = macIwf.YFilter
    macIwf.EntityData.YangName = "mac-iwf"
    macIwf.EntityData.BundleName = "cisco_ios_xr"
    macIwf.EntityData.ParentYangName = "limit-config"
    macIwf.EntityData.SegmentPath = "mac-iwf"
    macIwf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIwf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIwf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIwf.EntityData.Children = types.NewOrderedMap()
    macIwf.EntityData.Leafs = types.NewOrderedMap()
    macIwf.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", macIwf.MaxLimit})
    macIwf.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", macIwf.Threshold})
    macIwf.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", macIwf.RadiusOverrideEnabled})

    macIwf.EntityData.YListKeys = []string {}

    return &(macIwf.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface
// MAC Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacAccessInterface) GetEntityData() *types.CommonEntityData {
    macAccessInterface.EntityData.YFilter = macAccessInterface.YFilter
    macAccessInterface.EntityData.YangName = "mac-access-interface"
    macAccessInterface.EntityData.BundleName = "cisco_ios_xr"
    macAccessInterface.EntityData.ParentYangName = "limit-config"
    macAccessInterface.EntityData.SegmentPath = "mac-access-interface"
    macAccessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAccessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAccessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAccessInterface.EntityData.Children = types.NewOrderedMap()
    macAccessInterface.EntityData.Leafs = types.NewOrderedMap()
    macAccessInterface.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", macAccessInterface.MaxLimit})
    macAccessInterface.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", macAccessInterface.Threshold})
    macAccessInterface.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", macAccessInterface.RadiusOverrideEnabled})

    macAccessInterface.EntityData.YListKeys = []string {}

    return &(macAccessInterface.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface
// MAC IWF Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_MacIwfAccessInterface) GetEntityData() *types.CommonEntityData {
    macIwfAccessInterface.EntityData.YFilter = macIwfAccessInterface.YFilter
    macIwfAccessInterface.EntityData.YangName = "mac-iwf-access-interface"
    macIwfAccessInterface.EntityData.BundleName = "cisco_ios_xr"
    macIwfAccessInterface.EntityData.ParentYangName = "limit-config"
    macIwfAccessInterface.EntityData.SegmentPath = "mac-iwf-access-interface"
    macIwfAccessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIwfAccessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIwfAccessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIwfAccessInterface.EntityData.Children = types.NewOrderedMap()
    macIwfAccessInterface.EntityData.Leafs = types.NewOrderedMap()
    macIwfAccessInterface.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", macIwfAccessInterface.MaxLimit})
    macIwfAccessInterface.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", macIwfAccessInterface.Threshold})
    macIwfAccessInterface.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", macIwfAccessInterface.RadiusOverrideEnabled})

    macIwfAccessInterface.EntityData.YListKeys = []string {}

    return &(macIwfAccessInterface.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId
// Circuit ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitId) GetEntityData() *types.CommonEntityData {
    circuitId.EntityData.YFilter = circuitId.YFilter
    circuitId.EntityData.YangName = "circuit-id"
    circuitId.EntityData.BundleName = "cisco_ios_xr"
    circuitId.EntityData.ParentYangName = "limit-config"
    circuitId.EntityData.SegmentPath = "circuit-id"
    circuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitId.EntityData.Children = types.NewOrderedMap()
    circuitId.EntityData.Leafs = types.NewOrderedMap()
    circuitId.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", circuitId.MaxLimit})
    circuitId.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", circuitId.Threshold})
    circuitId.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", circuitId.RadiusOverrideEnabled})

    circuitId.EntityData.YListKeys = []string {}

    return &(circuitId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId
// Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_RemoteId) GetEntityData() *types.CommonEntityData {
    remoteId.EntityData.YFilter = remoteId.YFilter
    remoteId.EntityData.YangName = "remote-id"
    remoteId.EntityData.BundleName = "cisco_ios_xr"
    remoteId.EntityData.ParentYangName = "limit-config"
    remoteId.EntityData.SegmentPath = "remote-id"
    remoteId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteId.EntityData.Children = types.NewOrderedMap()
    remoteId.EntityData.Leafs = types.NewOrderedMap()
    remoteId.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", remoteId.MaxLimit})
    remoteId.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", remoteId.Threshold})
    remoteId.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", remoteId.RadiusOverrideEnabled})

    remoteId.EntityData.YListKeys = []string {}

    return &(remoteId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId
// Circuit ID and Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_CircuitIdAndRemoteId) GetEntityData() *types.CommonEntityData {
    circuitIdAndRemoteId.EntityData.YFilter = circuitIdAndRemoteId.YFilter
    circuitIdAndRemoteId.EntityData.YangName = "circuit-id-and-remote-id"
    circuitIdAndRemoteId.EntityData.BundleName = "cisco_ios_xr"
    circuitIdAndRemoteId.EntityData.ParentYangName = "limit-config"
    circuitIdAndRemoteId.EntityData.SegmentPath = "circuit-id-and-remote-id"
    circuitIdAndRemoteId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdAndRemoteId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdAndRemoteId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdAndRemoteId.EntityData.Children = types.NewOrderedMap()
    circuitIdAndRemoteId.EntityData.Leafs = types.NewOrderedMap()
    circuitIdAndRemoteId.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", circuitIdAndRemoteId.MaxLimit})
    circuitIdAndRemoteId.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", circuitIdAndRemoteId.Threshold})
    circuitIdAndRemoteId.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", circuitIdAndRemoteId.RadiusOverrideEnabled})

    circuitIdAndRemoteId.EntityData.YListKeys = []string {}

    return &(circuitIdAndRemoteId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId
// Outer VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_OuterVlanId) GetEntityData() *types.CommonEntityData {
    outerVlanId.EntityData.YFilter = outerVlanId.YFilter
    outerVlanId.EntityData.YangName = "outer-vlan-id"
    outerVlanId.EntityData.BundleName = "cisco_ios_xr"
    outerVlanId.EntityData.ParentYangName = "limit-config"
    outerVlanId.EntityData.SegmentPath = "outer-vlan-id"
    outerVlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outerVlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outerVlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outerVlanId.EntityData.Children = types.NewOrderedMap()
    outerVlanId.EntityData.Leafs = types.NewOrderedMap()
    outerVlanId.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", outerVlanId.MaxLimit})
    outerVlanId.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", outerVlanId.Threshold})
    outerVlanId.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", outerVlanId.RadiusOverrideEnabled})

    outerVlanId.EntityData.YListKeys = []string {}

    return &(outerVlanId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId
// Inner VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_InnerVlanId) GetEntityData() *types.CommonEntityData {
    innerVlanId.EntityData.YFilter = innerVlanId.YFilter
    innerVlanId.EntityData.YangName = "inner-vlan-id"
    innerVlanId.EntityData.BundleName = "cisco_ios_xr"
    innerVlanId.EntityData.ParentYangName = "limit-config"
    innerVlanId.EntityData.SegmentPath = "inner-vlan-id"
    innerVlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    innerVlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    innerVlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    innerVlanId.EntityData.Children = types.NewOrderedMap()
    innerVlanId.EntityData.Leafs = types.NewOrderedMap()
    innerVlanId.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", innerVlanId.MaxLimit})
    innerVlanId.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", innerVlanId.Threshold})
    innerVlanId.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", innerVlanId.RadiusOverrideEnabled})

    innerVlanId.EntityData.YListKeys = []string {}

    return &(innerVlanId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId
// VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Limit. The type is interface{} with range: 0..4294967295.
    MaxLimit interface{}

    // Threshold. The type is interface{} with range: 0..4294967295.
    Threshold interface{}

    // Radius override is enabled. The type is interface{} with range:
    // -2147483648..2147483647.
    RadiusOverrideEnabled interface{}
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_LimitConfig_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "limit-config"
    vlanId.EntityData.SegmentPath = "vlan-id"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = types.NewOrderedMap()
    vlanId.EntityData.Leafs = types.NewOrderedMap()
    vlanId.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", vlanId.MaxLimit})
    vlanId.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", vlanId.Threshold})
    vlanId.EntityData.Leafs.Append("radius-override-enabled", types.YLeaf{"RadiusOverrideEnabled", vlanId.RadiusOverrideEnabled})

    vlanId.EntityData.YListKeys = []string {}

    return &(vlanId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits
// PPPoE session limit information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE session limit state. The type is slice of
    // Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit.
    Limit []*Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit
}

func (limits *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits) GetEntityData() *types.CommonEntityData {
    limits.EntityData.YFilter = limits.YFilter
    limits.EntityData.YangName = "limits"
    limits.EntityData.BundleName = "cisco_ios_xr"
    limits.EntityData.ParentYangName = "bba-group"
    limits.EntityData.SegmentPath = "limits"
    limits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    limits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    limits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    limits.EntityData.Children = types.NewOrderedMap()
    limits.EntityData.Children.Append("limit", types.YChild{"Limit", nil})
    for i := range limits.Limit {
        limits.EntityData.Children.Append(types.GetSegmentPath(limits.Limit[i]), types.YChild{"Limit", limits.Limit[i]})
    }
    limits.EntityData.Leafs = types.NewOrderedMap()

    limits.EntityData.YListKeys = []string {}

    return &(limits.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit
// PPPoE session limit state
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit struct {
    EntityData types.CommonEntityData
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

func (limit *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Limits_Limit) GetEntityData() *types.CommonEntityData {
    limit.EntityData.YFilter = limit.YFilter
    limit.EntityData.YangName = "limit"
    limit.EntityData.BundleName = "cisco_ios_xr"
    limit.EntityData.ParentYangName = "limits"
    limit.EntityData.SegmentPath = "limit"
    limit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    limit.EntityData.Children = types.NewOrderedMap()
    limit.EntityData.Leafs = types.NewOrderedMap()
    limit.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", limit.InterfaceName})
    limit.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", limit.MacAddress})
    limit.EntityData.Leafs.Append("iwf", types.YLeaf{"Iwf", limit.Iwf})
    limit.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", limit.CircuitId})
    limit.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", limit.RemoteId})
    limit.EntityData.Leafs.Append("outer-vlan-id", types.YLeaf{"OuterVlanId", limit.OuterVlanId})
    limit.EntityData.Leafs.Append("inner-vlan-id", types.YLeaf{"InnerVlanId", limit.InnerVlanId})
    limit.EntityData.Leafs.Append("state", types.YLeaf{"State", limit.State})
    limit.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", limit.SessionCount})
    limit.EntityData.Leafs.Append("radius-override-set", types.YLeaf{"RadiusOverrideSet", limit.RadiusOverrideSet})
    limit.EntityData.Leafs.Append("override-limit", types.YLeaf{"OverrideLimit", limit.OverrideLimit})

    limit.EntityData.YListKeys = []string {}

    return &(limit.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles
// PPPoE throttle information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE session throttle state. The type is slice of
    // Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle.
    Throttle []*Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle
}

func (throttles *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles) GetEntityData() *types.CommonEntityData {
    throttles.EntityData.YFilter = throttles.YFilter
    throttles.EntityData.YangName = "throttles"
    throttles.EntityData.BundleName = "cisco_ios_xr"
    throttles.EntityData.ParentYangName = "bba-group"
    throttles.EntityData.SegmentPath = "throttles"
    throttles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttles.EntityData.Children = types.NewOrderedMap()
    throttles.EntityData.Children.Append("throttle", types.YChild{"Throttle", nil})
    for i := range throttles.Throttle {
        throttles.EntityData.Children.Append(types.GetSegmentPath(throttles.Throttle[i]), types.YChild{"Throttle", throttles.Throttle[i]})
    }
    throttles.EntityData.Leafs = types.NewOrderedMap()

    throttles.EntityData.YListKeys = []string {}

    return &(throttles.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle
// PPPoE session throttle state
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle struct {
    EntityData types.CommonEntityData
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

func (throttle *Pppoe_Nodes_Node_BbaGroups_BbaGroup_Throttles_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "throttles"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", throttle.InterfaceName})
    throttle.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", throttle.MacAddress})
    throttle.EntityData.Leafs.Append("iwf", types.YLeaf{"Iwf", throttle.Iwf})
    throttle.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", throttle.CircuitId})
    throttle.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", throttle.RemoteId})
    throttle.EntityData.Leafs.Append("outer-vlan-id", types.YLeaf{"OuterVlanId", throttle.OuterVlanId})
    throttle.EntityData.Leafs.Append("inner-vlan-id", types.YLeaf{"InnerVlanId", throttle.InnerVlanId})
    throttle.EntityData.Leafs.Append("state", types.YLeaf{"State", throttle.State})
    throttle.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", throttle.TimeLeft})
    throttle.EntityData.Leafs.Append("since-reset", types.YLeaf{"SinceReset", throttle.SinceReset})
    throttle.EntityData.Leafs.Append("padi-count", types.YLeaf{"PadiCount", throttle.PadiCount})
    throttle.EntityData.Leafs.Append("padr-count", types.YLeaf{"PadrCount", throttle.PadrCount})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig
// BBA-Group throttle configuration information
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig struct {
    EntityData types.CommonEntityData
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

func (throttleConfig *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig) GetEntityData() *types.CommonEntityData {
    throttleConfig.EntityData.YFilter = throttleConfig.YFilter
    throttleConfig.EntityData.YangName = "throttle-config"
    throttleConfig.EntityData.BundleName = "cisco_ios_xr"
    throttleConfig.EntityData.ParentYangName = "bba-group"
    throttleConfig.EntityData.SegmentPath = "throttle-config"
    throttleConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleConfig.EntityData.Children = types.NewOrderedMap()
    throttleConfig.EntityData.Children.Append("mac", types.YChild{"Mac", &throttleConfig.Mac})
    throttleConfig.EntityData.Children.Append("mac-access-interface", types.YChild{"MacAccessInterface", &throttleConfig.MacAccessInterface})
    throttleConfig.EntityData.Children.Append("mac-iwf-access-interface", types.YChild{"MacIwfAccessInterface", &throttleConfig.MacIwfAccessInterface})
    throttleConfig.EntityData.Children.Append("circuit-id", types.YChild{"CircuitId", &throttleConfig.CircuitId})
    throttleConfig.EntityData.Children.Append("remote-id", types.YChild{"RemoteId", &throttleConfig.RemoteId})
    throttleConfig.EntityData.Children.Append("circuit-id-and-remote-id", types.YChild{"CircuitIdAndRemoteId", &throttleConfig.CircuitIdAndRemoteId})
    throttleConfig.EntityData.Children.Append("outer-vlan-id", types.YChild{"OuterVlanId", &throttleConfig.OuterVlanId})
    throttleConfig.EntityData.Children.Append("inner-vlan-id", types.YChild{"InnerVlanId", &throttleConfig.InnerVlanId})
    throttleConfig.EntityData.Children.Append("vlan-id", types.YChild{"VlanId", &throttleConfig.VlanId})
    throttleConfig.EntityData.Leafs = types.NewOrderedMap()

    throttleConfig.EntityData.YListKeys = []string {}

    return &(throttleConfig.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac
// MAC
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (mac *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "throttle-config"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", mac.Limit})
    mac.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", mac.RequestPeriod})
    mac.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", mac.BlockingPeriod})

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface
// MAC Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (macAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacAccessInterface) GetEntityData() *types.CommonEntityData {
    macAccessInterface.EntityData.YFilter = macAccessInterface.YFilter
    macAccessInterface.EntityData.YangName = "mac-access-interface"
    macAccessInterface.EntityData.BundleName = "cisco_ios_xr"
    macAccessInterface.EntityData.ParentYangName = "throttle-config"
    macAccessInterface.EntityData.SegmentPath = "mac-access-interface"
    macAccessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAccessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAccessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAccessInterface.EntityData.Children = types.NewOrderedMap()
    macAccessInterface.EntityData.Leafs = types.NewOrderedMap()
    macAccessInterface.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", macAccessInterface.Limit})
    macAccessInterface.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", macAccessInterface.RequestPeriod})
    macAccessInterface.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", macAccessInterface.BlockingPeriod})

    macAccessInterface.EntityData.YListKeys = []string {}

    return &(macAccessInterface.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface
// MAC IWF Access Interface
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (macIwfAccessInterface *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_MacIwfAccessInterface) GetEntityData() *types.CommonEntityData {
    macIwfAccessInterface.EntityData.YFilter = macIwfAccessInterface.YFilter
    macIwfAccessInterface.EntityData.YangName = "mac-iwf-access-interface"
    macIwfAccessInterface.EntityData.BundleName = "cisco_ios_xr"
    macIwfAccessInterface.EntityData.ParentYangName = "throttle-config"
    macIwfAccessInterface.EntityData.SegmentPath = "mac-iwf-access-interface"
    macIwfAccessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIwfAccessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIwfAccessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIwfAccessInterface.EntityData.Children = types.NewOrderedMap()
    macIwfAccessInterface.EntityData.Leafs = types.NewOrderedMap()
    macIwfAccessInterface.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", macIwfAccessInterface.Limit})
    macIwfAccessInterface.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", macIwfAccessInterface.RequestPeriod})
    macIwfAccessInterface.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", macIwfAccessInterface.BlockingPeriod})

    macIwfAccessInterface.EntityData.YListKeys = []string {}

    return &(macIwfAccessInterface.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId
// Circuit ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (circuitId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitId) GetEntityData() *types.CommonEntityData {
    circuitId.EntityData.YFilter = circuitId.YFilter
    circuitId.EntityData.YangName = "circuit-id"
    circuitId.EntityData.BundleName = "cisco_ios_xr"
    circuitId.EntityData.ParentYangName = "throttle-config"
    circuitId.EntityData.SegmentPath = "circuit-id"
    circuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitId.EntityData.Children = types.NewOrderedMap()
    circuitId.EntityData.Leafs = types.NewOrderedMap()
    circuitId.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", circuitId.Limit})
    circuitId.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", circuitId.RequestPeriod})
    circuitId.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", circuitId.BlockingPeriod})

    circuitId.EntityData.YListKeys = []string {}

    return &(circuitId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId
// Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (remoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_RemoteId) GetEntityData() *types.CommonEntityData {
    remoteId.EntityData.YFilter = remoteId.YFilter
    remoteId.EntityData.YangName = "remote-id"
    remoteId.EntityData.BundleName = "cisco_ios_xr"
    remoteId.EntityData.ParentYangName = "throttle-config"
    remoteId.EntityData.SegmentPath = "remote-id"
    remoteId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteId.EntityData.Children = types.NewOrderedMap()
    remoteId.EntityData.Leafs = types.NewOrderedMap()
    remoteId.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", remoteId.Limit})
    remoteId.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", remoteId.RequestPeriod})
    remoteId.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", remoteId.BlockingPeriod})

    remoteId.EntityData.YListKeys = []string {}

    return &(remoteId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId
// Circuit ID and Remote ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (circuitIdAndRemoteId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_CircuitIdAndRemoteId) GetEntityData() *types.CommonEntityData {
    circuitIdAndRemoteId.EntityData.YFilter = circuitIdAndRemoteId.YFilter
    circuitIdAndRemoteId.EntityData.YangName = "circuit-id-and-remote-id"
    circuitIdAndRemoteId.EntityData.BundleName = "cisco_ios_xr"
    circuitIdAndRemoteId.EntityData.ParentYangName = "throttle-config"
    circuitIdAndRemoteId.EntityData.SegmentPath = "circuit-id-and-remote-id"
    circuitIdAndRemoteId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdAndRemoteId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdAndRemoteId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdAndRemoteId.EntityData.Children = types.NewOrderedMap()
    circuitIdAndRemoteId.EntityData.Leafs = types.NewOrderedMap()
    circuitIdAndRemoteId.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", circuitIdAndRemoteId.Limit})
    circuitIdAndRemoteId.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", circuitIdAndRemoteId.RequestPeriod})
    circuitIdAndRemoteId.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", circuitIdAndRemoteId.BlockingPeriod})

    circuitIdAndRemoteId.EntityData.YListKeys = []string {}

    return &(circuitIdAndRemoteId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId
// Outer VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (outerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_OuterVlanId) GetEntityData() *types.CommonEntityData {
    outerVlanId.EntityData.YFilter = outerVlanId.YFilter
    outerVlanId.EntityData.YangName = "outer-vlan-id"
    outerVlanId.EntityData.BundleName = "cisco_ios_xr"
    outerVlanId.EntityData.ParentYangName = "throttle-config"
    outerVlanId.EntityData.SegmentPath = "outer-vlan-id"
    outerVlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outerVlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outerVlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outerVlanId.EntityData.Children = types.NewOrderedMap()
    outerVlanId.EntityData.Leafs = types.NewOrderedMap()
    outerVlanId.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", outerVlanId.Limit})
    outerVlanId.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", outerVlanId.RequestPeriod})
    outerVlanId.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", outerVlanId.BlockingPeriod})

    outerVlanId.EntityData.YListKeys = []string {}

    return &(outerVlanId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId
// Inner VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (innerVlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_InnerVlanId) GetEntityData() *types.CommonEntityData {
    innerVlanId.EntityData.YFilter = innerVlanId.YFilter
    innerVlanId.EntityData.YangName = "inner-vlan-id"
    innerVlanId.EntityData.BundleName = "cisco_ios_xr"
    innerVlanId.EntityData.ParentYangName = "throttle-config"
    innerVlanId.EntityData.SegmentPath = "inner-vlan-id"
    innerVlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    innerVlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    innerVlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    innerVlanId.EntityData.Children = types.NewOrderedMap()
    innerVlanId.EntityData.Leafs = types.NewOrderedMap()
    innerVlanId.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", innerVlanId.Limit})
    innerVlanId.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", innerVlanId.RequestPeriod})
    innerVlanId.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", innerVlanId.BlockingPeriod})

    innerVlanId.EntityData.YListKeys = []string {}

    return &(innerVlanId.EntityData)
}

// Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId
// VLAN ID
type Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit. The type is interface{} with range: 0..4294967295.
    Limit interface{}

    // Request Period. The type is interface{} with range: 0..4294967295.
    RequestPeriod interface{}

    // Blocking Period. The type is interface{} with range: 0..4294967295.
    BlockingPeriod interface{}
}

func (vlanId *Pppoe_Nodes_Node_BbaGroups_BbaGroup_ThrottleConfig_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "throttle-config"
    vlanId.EntityData.SegmentPath = "vlan-id"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = types.NewOrderedMap()
    vlanId.EntityData.Leafs = types.NewOrderedMap()
    vlanId.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", vlanId.Limit})
    vlanId.EntityData.Leafs.Append("request-period", types.YLeaf{"RequestPeriod", vlanId.RequestPeriod})
    vlanId.EntityData.Leafs.Append("blocking-period", types.YLeaf{"BlockingPeriod", vlanId.BlockingPeriod})

    vlanId.EntityData.YListKeys = []string {}

    return &(vlanId.EntityData)
}

// Pppoe_Nodes_Node_SummaryTotal
// PPPoE statistics for a given node
type Pppoe_Nodes_Node_SummaryTotal struct {
    EntityData types.CommonEntityData
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

func (summaryTotal *Pppoe_Nodes_Node_SummaryTotal) GetEntityData() *types.CommonEntityData {
    summaryTotal.EntityData.YFilter = summaryTotal.YFilter
    summaryTotal.EntityData.YangName = "summary-total"
    summaryTotal.EntityData.BundleName = "cisco_ios_xr"
    summaryTotal.EntityData.ParentYangName = "node"
    summaryTotal.EntityData.SegmentPath = "summary-total"
    summaryTotal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryTotal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryTotal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryTotal.EntityData.Children = types.NewOrderedMap()
    summaryTotal.EntityData.Leafs = types.NewOrderedMap()
    summaryTotal.EntityData.Leafs.Append("ready-access-interfaces", types.YLeaf{"ReadyAccessInterfaces", summaryTotal.ReadyAccessInterfaces})
    summaryTotal.EntityData.Leafs.Append("not-ready-access-interfaces", types.YLeaf{"NotReadyAccessInterfaces", summaryTotal.NotReadyAccessInterfaces})
    summaryTotal.EntityData.Leafs.Append("complete-sessions", types.YLeaf{"CompleteSessions", summaryTotal.CompleteSessions})
    summaryTotal.EntityData.Leafs.Append("incomplete-sessions", types.YLeaf{"IncompleteSessions", summaryTotal.IncompleteSessions})
    summaryTotal.EntityData.Leafs.Append("flow-control-limit", types.YLeaf{"FlowControlLimit", summaryTotal.FlowControlLimit})
    summaryTotal.EntityData.Leafs.Append("flow-control-in-flight-sessions", types.YLeaf{"FlowControlInFlightSessions", summaryTotal.FlowControlInFlightSessions})
    summaryTotal.EntityData.Leafs.Append("flow-control-dropped-sessions", types.YLeaf{"FlowControlDroppedSessions", summaryTotal.FlowControlDroppedSessions})
    summaryTotal.EntityData.Leafs.Append("flow-control-disconnected-sessions", types.YLeaf{"FlowControlDisconnectedSessions", summaryTotal.FlowControlDisconnectedSessions})
    summaryTotal.EntityData.Leafs.Append("flow-control-successful-sessions", types.YLeaf{"FlowControlSuccessfulSessions", summaryTotal.FlowControlSuccessfulSessions})
    summaryTotal.EntityData.Leafs.Append("pppoema-subscriber-infra-flow-control", types.YLeaf{"PppoemaSubscriberInfraFlowControl", summaryTotal.PppoemaSubscriberInfraFlowControl})

    summaryTotal.EntityData.YListKeys = []string {}

    return &(summaryTotal.EntityData)
}

