// This module contains a collection of YANG definitions
// for bgp operational data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package bgp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-bgp-oper bgp-state-data}", reflect.TypeOf(BgpStateData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-bgp-oper:bgp-state-data", reflect.TypeOf(BgpStateData{}))
}

// BgpFsmState represents BGP FSM State
type BgpFsmState string

const (
    // neighbor is in Idle state
    BgpFsmState_fsm_idle BgpFsmState = "fsm-idle"

    // neighbor is in Connect state
    BgpFsmState_fsm_connect BgpFsmState = "fsm-connect"

    // neighbor is in Active state
    BgpFsmState_fsm_active BgpFsmState = "fsm-active"

    // neighbor is in OpenSent state
    BgpFsmState_fsm_opensent BgpFsmState = "fsm-opensent"

    // neighbor is in OpenConfirm state
    BgpFsmState_fsm_openconfirm BgpFsmState = "fsm-openconfirm"

    // neighbor is in Established state
    BgpFsmState_fsm_established BgpFsmState = "fsm-established"

    // neighbor is Non Negotiated
    BgpFsmState_fsm_nonnegotiated BgpFsmState = "fsm-nonnegotiated"
)

// BgpLink represents Operational state relevent to bgp global neighbor
type BgpLink string

const (
    // iBGP neighbors
    BgpLink_internal BgpLink = "internal"

    // eBGP neighbors
    BgpLink_external BgpLink = "external"
)

// BgpMode represents BGP mode
type BgpMode string

const (
    // active connection
    BgpMode_mode_active BgpMode = "mode-active"

    // passive connection
    BgpMode_mode_passive BgpMode = "mode-passive"
)

// BgpStateData
// BGP operational state data
type BgpStateData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP neighbor information.
    Neighbors BgpStateData_Neighbors

    // BGP address family.
    AddressFamilies BgpStateData_AddressFamilies

    // BGP VRFs.
    BgpRouteVrfs BgpStateData_BgpRouteVrfs
}

func (bgpStateData *BgpStateData) GetFilter() yfilter.YFilter { return bgpStateData.YFilter }

func (bgpStateData *BgpStateData) SetFilter(yf yfilter.YFilter) { bgpStateData.YFilter = yf }

func (bgpStateData *BgpStateData) GetGoName(yname string) string {
    if yname == "neighbors" { return "Neighbors" }
    if yname == "address-families" { return "AddressFamilies" }
    if yname == "bgp-route-vrfs" { return "BgpRouteVrfs" }
    return ""
}

func (bgpStateData *BgpStateData) GetSegmentPath() string {
    return "Cisco-IOS-XE-bgp-oper:bgp-state-data"
}

func (bgpStateData *BgpStateData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &bgpStateData.Neighbors
    }
    if childYangName == "address-families" {
        return &bgpStateData.AddressFamilies
    }
    if childYangName == "bgp-route-vrfs" {
        return &bgpStateData.BgpRouteVrfs
    }
    return nil
}

func (bgpStateData *BgpStateData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &bgpStateData.Neighbors
    children["address-families"] = &bgpStateData.AddressFamilies
    children["bgp-route-vrfs"] = &bgpStateData.BgpRouteVrfs
    return children
}

func (bgpStateData *BgpStateData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpStateData *BgpStateData) GetBundleName() string { return "cisco_ios_xe" }

func (bgpStateData *BgpStateData) GetYangName() string { return "bgp-state-data" }

func (bgpStateData *BgpStateData) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpStateData *BgpStateData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpStateData *BgpStateData) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpStateData *BgpStateData) SetParent(parent types.Entity) { bgpStateData.parent = parent }

func (bgpStateData *BgpStateData) GetParent() types.Entity { return bgpStateData.parent }

func (bgpStateData *BgpStateData) GetParentYangName() string { return "Cisco-IOS-XE-bgp-oper" }

// BgpStateData_Neighbors
// BGP neighbor information
type BgpStateData_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP neighbors. The type is slice of
    // BgpStateData_Neighbors_Neighbor.
    Neighbor []BgpStateData_Neighbors_Neighbor
}

func (neighbors *BgpStateData_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *BgpStateData_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *BgpStateData_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *BgpStateData_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *BgpStateData_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *BgpStateData_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *BgpStateData_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *BgpStateData_Neighbors) GetBundleName() string { return "cisco_ios_xe" }

func (neighbors *BgpStateData_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *BgpStateData_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (neighbors *BgpStateData_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (neighbors *BgpStateData_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (neighbors *BgpStateData_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *BgpStateData_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *BgpStateData_Neighbors) GetParentYangName() string { return "bgp-state-data" }

// BgpStateData_Neighbors_Neighbor
// List of BGP neighbors
type BgpStateData_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Afi-safi key. The type is AfiSafi.
    AfiSafi interface{}

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // This attribute is a key. Neighbor identifier. The type is string.
    NeighborId interface{}

    // Neighbor description string. The type is string.
    Description interface{}

    // BGP version being used to communicate with the remote router. The type is
    // interface{} with range: 0..65535.
    BgpVersion interface{}

    // Neighbor link type. The type is BgpLink.
    Link interface{}

    // Amout of time the bgp session has been up since being established. The type
    // is string.
    UpTime interface{}

    // Time since BGP last sent a message to the neighbor. The type is string.
    LastWrite interface{}

    // Time since BGP last received a message from the neighbor. The type is
    // string.
    LastRead interface{}

    // The number of installed prefixes. The type is interface{} with range:
    // 0..4294967295.
    InstalledPrefixes interface{}

    // BGP neighbor session state. The type is BgpFsmState.
    SessionState interface{}

    // Negotiated capabilities for neighbor session. The type is slice of string.
    NegotiatedCap []interface{}

    // Negotiated keepalive timers status of BGP neighbor.
    NegotiatedKeepaliveTimers BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers

    // BGP neighbor session counters.
    BgpNeighborCounters BgpStateData_Neighbors_Neighbor_BgpNeighborCounters

    // BGP neighbor connection.
    Connection BgpStateData_Neighbors_Neighbor_Connection

    // BGP neighbor transport.
    Transport BgpStateData_Neighbors_Neighbor_Transport

    // BGP neighbor activity.
    PrefixActivity BgpStateData_Neighbors_Neighbor_PrefixActivity
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *BgpStateData_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "description" { return "Description" }
    if yname == "bgp-version" { return "BgpVersion" }
    if yname == "link" { return "Link" }
    if yname == "up-time" { return "UpTime" }
    if yname == "last-write" { return "LastWrite" }
    if yname == "last-read" { return "LastRead" }
    if yname == "installed-prefixes" { return "InstalledPrefixes" }
    if yname == "session-state" { return "SessionState" }
    if yname == "negotiated-cap" { return "NegotiatedCap" }
    if yname == "negotiated-keepalive-timers" { return "NegotiatedKeepaliveTimers" }
    if yname == "bgp-neighbor-counters" { return "BgpNeighborCounters" }
    if yname == "connection" { return "Connection" }
    if yname == "transport" { return "Transport" }
    if yname == "prefix-activity" { return "PrefixActivity" }
    return ""
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[afi-safi='" + fmt.Sprintf("%v", neighbor.AfiSafi) + "']" + "[vrf-name='" + fmt.Sprintf("%v", neighbor.VrfName) + "']" + "[neighbor-id='" + fmt.Sprintf("%v", neighbor.NeighborId) + "']"
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "negotiated-keepalive-timers" {
        return &neighbor.NegotiatedKeepaliveTimers
    }
    if childYangName == "bgp-neighbor-counters" {
        return &neighbor.BgpNeighborCounters
    }
    if childYangName == "connection" {
        return &neighbor.Connection
    }
    if childYangName == "transport" {
        return &neighbor.Transport
    }
    if childYangName == "prefix-activity" {
        return &neighbor.PrefixActivity
    }
    return nil
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["negotiated-keepalive-timers"] = &neighbor.NegotiatedKeepaliveTimers
    children["bgp-neighbor-counters"] = &neighbor.BgpNeighborCounters
    children["connection"] = &neighbor.Connection
    children["transport"] = &neighbor.Transport
    children["prefix-activity"] = &neighbor.PrefixActivity
    return children
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi"] = neighbor.AfiSafi
    leafs["vrf-name"] = neighbor.VrfName
    leafs["neighbor-id"] = neighbor.NeighborId
    leafs["description"] = neighbor.Description
    leafs["bgp-version"] = neighbor.BgpVersion
    leafs["link"] = neighbor.Link
    leafs["up-time"] = neighbor.UpTime
    leafs["last-write"] = neighbor.LastWrite
    leafs["last-read"] = neighbor.LastRead
    leafs["installed-prefixes"] = neighbor.InstalledPrefixes
    leafs["session-state"] = neighbor.SessionState
    leafs["negotiated-cap"] = neighbor.NegotiatedCap
    return leafs
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xe" }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (neighbor *BgpStateData_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *BgpStateData_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers
// Negotiated keepalive timers status of BGP neighbor
type BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Keepalive interval. The type is interface{} with range: 0..65535.
    KeepaliveInterval interface{}
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetFilter() yfilter.YFilter { return negotiatedKeepaliveTimers.YFilter }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) SetFilter(yf yfilter.YFilter) { negotiatedKeepaliveTimers.YFilter = yf }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetGoName(yname string) string {
    if yname == "hold-time" { return "HoldTime" }
    if yname == "keepalive-interval" { return "KeepaliveInterval" }
    return ""
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetSegmentPath() string {
    return "negotiated-keepalive-timers"
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hold-time"] = negotiatedKeepaliveTimers.HoldTime
    leafs["keepalive-interval"] = negotiatedKeepaliveTimers.KeepaliveInterval
    return leafs
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetBundleName() string { return "cisco_ios_xe" }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetYangName() string { return "negotiated-keepalive-timers" }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) SetParent(parent types.Entity) { negotiatedKeepaliveTimers.parent = parent }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetParent() types.Entity { return negotiatedKeepaliveTimers.parent }

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetParentYangName() string { return "neighbor" }

// BgpStateData_Neighbors_Neighbor_BgpNeighborCounters
// BGP neighbor session counters
type BgpStateData_Neighbors_Neighbor_BgpNeighborCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Input Q depth. The type is interface{} with range: 0..4294967295.
    InqDepth interface{}

    // Output Q depth. The type is interface{} with range: 0..4294967295.
    OutqDepth interface{}

    // Number of mesaged sent.
    Sent BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent

    // Number of mesaged received.
    Received BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetFilter() yfilter.YFilter { return bgpNeighborCounters.YFilter }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) SetFilter(yf yfilter.YFilter) { bgpNeighborCounters.YFilter = yf }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetGoName(yname string) string {
    if yname == "inq-depth" { return "InqDepth" }
    if yname == "outq-depth" { return "OutqDepth" }
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    return ""
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetSegmentPath() string {
    return "bgp-neighbor-counters"
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sent" {
        return &bgpNeighborCounters.Sent
    }
    if childYangName == "received" {
        return &bgpNeighborCounters.Received
    }
    return nil
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sent"] = &bgpNeighborCounters.Sent
    children["received"] = &bgpNeighborCounters.Received
    return children
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inq-depth"] = bgpNeighborCounters.InqDepth
    leafs["outq-depth"] = bgpNeighborCounters.OutqDepth
    return leafs
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetBundleName() string { return "cisco_ios_xe" }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetYangName() string { return "bgp-neighbor-counters" }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) SetParent(parent types.Entity) { bgpNeighborCounters.parent = parent }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetParent() types.Entity { return bgpNeighborCounters.parent }

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetParentYangName() string { return "neighbor" }

// BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent
// Number of mesaged sent
type BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OPEN message count. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // UPDATE message count. The type is interface{} with range: 0..4294967295.
    Updates interface{}

    // NOTIFICATION message count. The type is interface{} with range:
    // 0..4294967295.
    Notifications interface{}

    // KEEPALIVE message count. The type is interface{} with range: 0..4294967295.
    Keepalives interface{}

    // Route refresh message count. The type is interface{} with range:
    // 0..4294967295.
    RouteRefreshes interface{}
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetFilter() yfilter.YFilter { return sent.YFilter }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) SetFilter(yf yfilter.YFilter) { sent.YFilter = yf }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetGoName(yname string) string {
    if yname == "opens" { return "Opens" }
    if yname == "updates" { return "Updates" }
    if yname == "notifications" { return "Notifications" }
    if yname == "keepalives" { return "Keepalives" }
    if yname == "route-refreshes" { return "RouteRefreshes" }
    return ""
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetSegmentPath() string {
    return "sent"
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["opens"] = sent.Opens
    leafs["updates"] = sent.Updates
    leafs["notifications"] = sent.Notifications
    leafs["keepalives"] = sent.Keepalives
    leafs["route-refreshes"] = sent.RouteRefreshes
    return leafs
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetBundleName() string { return "cisco_ios_xe" }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetYangName() string { return "sent" }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) SetParent(parent types.Entity) { sent.parent = parent }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetParent() types.Entity { return sent.parent }

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetParentYangName() string { return "bgp-neighbor-counters" }

// BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received
// Number of mesaged received
type BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OPEN message count. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // UPDATE message count. The type is interface{} with range: 0..4294967295.
    Updates interface{}

    // NOTIFICATION message count. The type is interface{} with range:
    // 0..4294967295.
    Notifications interface{}

    // KEEPALIVE message count. The type is interface{} with range: 0..4294967295.
    Keepalives interface{}

    // Route refresh message count. The type is interface{} with range:
    // 0..4294967295.
    RouteRefreshes interface{}
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetGoName(yname string) string {
    if yname == "opens" { return "Opens" }
    if yname == "updates" { return "Updates" }
    if yname == "notifications" { return "Notifications" }
    if yname == "keepalives" { return "Keepalives" }
    if yname == "route-refreshes" { return "RouteRefreshes" }
    return ""
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetSegmentPath() string {
    return "received"
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["opens"] = received.Opens
    leafs["updates"] = received.Updates
    leafs["notifications"] = received.Notifications
    leafs["keepalives"] = received.Keepalives
    leafs["route-refreshes"] = received.RouteRefreshes
    return leafs
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetBundleName() string { return "cisco_ios_xe" }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetYangName() string { return "received" }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetParent() types.Entity { return received.parent }

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetParentYangName() string { return "bgp-neighbor-counters" }

// BgpStateData_Neighbors_Neighbor_Connection
// BGP neighbor connection
type BgpStateData_Neighbors_Neighbor_Connection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP FSM state. The type is TcpFsmState.
    State interface{}

    // BGP transport connection mode. The type is BgpMode.
    Mode interface{}

    // The number of times a TCP and BGP  connection has been successfully
    // established. The type is interface{} with range: 0..4294967295.
    TotalEstablished interface{}

    // The number of times that a valid session has failed or been taken down. The
    // type is interface{} with range: 0..4294967295.
    TotalDropped interface{}

    // Time since peering session was last reset. The type is string.
    LastReset interface{}

    // The reason for the last reset. The type is string.
    ResetReason interface{}
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetFilter() yfilter.YFilter { return connection.YFilter }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) SetFilter(yf yfilter.YFilter) { connection.YFilter = yf }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "mode" { return "Mode" }
    if yname == "total-established" { return "TotalEstablished" }
    if yname == "total-dropped" { return "TotalDropped" }
    if yname == "last-reset" { return "LastReset" }
    if yname == "reset-reason" { return "ResetReason" }
    return ""
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetSegmentPath() string {
    return "connection"
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = connection.State
    leafs["mode"] = connection.Mode
    leafs["total-established"] = connection.TotalEstablished
    leafs["total-dropped"] = connection.TotalDropped
    leafs["last-reset"] = connection.LastReset
    leafs["reset-reason"] = connection.ResetReason
    return leafs
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetBundleName() string { return "cisco_ios_xe" }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetYangName() string { return "connection" }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) SetParent(parent types.Entity) { connection.parent = parent }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetParent() types.Entity { return connection.parent }

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetParentYangName() string { return "neighbor" }

// BgpStateData_Neighbors_Neighbor_Transport
// BGP neighbor transport
type BgpStateData_Neighbors_Neighbor_Transport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indication whether path MTU discovrey is enabled. The type is bool.
    PathMtuDiscovery interface{}

    // Local TCP port used for TCP session. The type is interface{} with range:
    // 0..4294967295.
    LocalPort interface{}

    // Local address used for the TCP session. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalHost interface{}

    // Remote port used by the peer for the TCP session. The type is interface{}
    // with range: 0..4294967295.
    ForeignPort interface{}

    // Remote address of the BGP session. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForeignHost interface{}

    // Maximum Data segment size. The type is interface{} with range:
    // 0..4294967295.
    Mss interface{}
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetFilter() yfilter.YFilter { return transport.YFilter }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) SetFilter(yf yfilter.YFilter) { transport.YFilter = yf }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetGoName(yname string) string {
    if yname == "path-mtu-discovery" { return "PathMtuDiscovery" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "local-host" { return "LocalHost" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "foreign-host" { return "ForeignHost" }
    if yname == "mss" { return "Mss" }
    return ""
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetSegmentPath() string {
    return "transport"
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-mtu-discovery"] = transport.PathMtuDiscovery
    leafs["local-port"] = transport.LocalPort
    leafs["local-host"] = transport.LocalHost
    leafs["foreign-port"] = transport.ForeignPort
    leafs["foreign-host"] = transport.ForeignHost
    leafs["mss"] = transport.Mss
    return leafs
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetBundleName() string { return "cisco_ios_xe" }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetYangName() string { return "transport" }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) SetParent(parent types.Entity) { transport.parent = parent }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetParent() types.Entity { return transport.parent }

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetParentYangName() string { return "neighbor" }

// BgpStateData_Neighbors_Neighbor_PrefixActivity
// BGP neighbor activity
type BgpStateData_Neighbors_Neighbor_PrefixActivity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of prefixes sent.
    Sent BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent

    // Number of prefixes received.
    Received BgpStateData_Neighbors_Neighbor_PrefixActivity_Received
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetFilter() yfilter.YFilter { return prefixActivity.YFilter }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) SetFilter(yf yfilter.YFilter) { prefixActivity.YFilter = yf }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    return ""
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetSegmentPath() string {
    return "prefix-activity"
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sent" {
        return &prefixActivity.Sent
    }
    if childYangName == "received" {
        return &prefixActivity.Received
    }
    return nil
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sent"] = &prefixActivity.Sent
    children["received"] = &prefixActivity.Received
    return children
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetBundleName() string { return "cisco_ios_xe" }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetYangName() string { return "prefix-activity" }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) SetParent(parent types.Entity) { prefixActivity.parent = parent }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetParent() types.Entity { return prefixActivity.parent }

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetParentYangName() string { return "neighbor" }

// BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent
// Number of prefixes sent
type BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The current number of accepted prefixes. The type is interface{} with
    // range: 0..18446744073709551615.
    CurrentPrefixes interface{}

    // The total number of accepted prefixes. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalPrefixes interface{}

    // The number of times a prefix has been withdrawn and readvertised. The type
    // is interface{} with range: 0..18446744073709551615.
    ImplicitWithdraw interface{}

    // The number of times a prefix has been withdrawn because it is no longer
    // feasible. The type is interface{} with range: 0..18446744073709551615.
    ExplicitWithdraw interface{}

    // The number of received prefixes installed as best paths. The type is
    // interface{} with range: 0..18446744073709551615.
    Bestpaths interface{}

    // The number of received prefixes installed as multipaths. The type is
    // interface{} with range: 0..18446744073709551615.
    Multipaths interface{}
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetFilter() yfilter.YFilter { return sent.YFilter }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) SetFilter(yf yfilter.YFilter) { sent.YFilter = yf }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetGoName(yname string) string {
    if yname == "current-prefixes" { return "CurrentPrefixes" }
    if yname == "total-prefixes" { return "TotalPrefixes" }
    if yname == "implicit-withdraw" { return "ImplicitWithdraw" }
    if yname == "explicit-withdraw" { return "ExplicitWithdraw" }
    if yname == "bestpaths" { return "Bestpaths" }
    if yname == "multipaths" { return "Multipaths" }
    return ""
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetSegmentPath() string {
    return "sent"
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-prefixes"] = sent.CurrentPrefixes
    leafs["total-prefixes"] = sent.TotalPrefixes
    leafs["implicit-withdraw"] = sent.ImplicitWithdraw
    leafs["explicit-withdraw"] = sent.ExplicitWithdraw
    leafs["bestpaths"] = sent.Bestpaths
    leafs["multipaths"] = sent.Multipaths
    return leafs
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetBundleName() string { return "cisco_ios_xe" }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetYangName() string { return "sent" }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) SetParent(parent types.Entity) { sent.parent = parent }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetParent() types.Entity { return sent.parent }

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetParentYangName() string { return "prefix-activity" }

// BgpStateData_Neighbors_Neighbor_PrefixActivity_Received
// Number of prefixes received
type BgpStateData_Neighbors_Neighbor_PrefixActivity_Received struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The current number of accepted prefixes. The type is interface{} with
    // range: 0..18446744073709551615.
    CurrentPrefixes interface{}

    // The total number of accepted prefixes. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalPrefixes interface{}

    // The number of times a prefix has been withdrawn and readvertised. The type
    // is interface{} with range: 0..18446744073709551615.
    ImplicitWithdraw interface{}

    // The number of times a prefix has been withdrawn because it is no longer
    // feasible. The type is interface{} with range: 0..18446744073709551615.
    ExplicitWithdraw interface{}

    // The number of received prefixes installed as best paths. The type is
    // interface{} with range: 0..18446744073709551615.
    Bestpaths interface{}

    // The number of received prefixes installed as multipaths. The type is
    // interface{} with range: 0..18446744073709551615.
    Multipaths interface{}
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetGoName(yname string) string {
    if yname == "current-prefixes" { return "CurrentPrefixes" }
    if yname == "total-prefixes" { return "TotalPrefixes" }
    if yname == "implicit-withdraw" { return "ImplicitWithdraw" }
    if yname == "explicit-withdraw" { return "ExplicitWithdraw" }
    if yname == "bestpaths" { return "Bestpaths" }
    if yname == "multipaths" { return "Multipaths" }
    return ""
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetSegmentPath() string {
    return "received"
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-prefixes"] = received.CurrentPrefixes
    leafs["total-prefixes"] = received.TotalPrefixes
    leafs["implicit-withdraw"] = received.ImplicitWithdraw
    leafs["explicit-withdraw"] = received.ExplicitWithdraw
    leafs["bestpaths"] = received.Bestpaths
    leafs["multipaths"] = received.Multipaths
    return leafs
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetBundleName() string { return "cisco_ios_xe" }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetYangName() string { return "received" }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetParent() types.Entity { return received.parent }

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetParentYangName() string { return "prefix-activity" }

// BgpStateData_AddressFamilies
// BGP address family
type BgpStateData_AddressFamilies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP address families. The type is slice of
    // BgpStateData_AddressFamilies_AddressFamily.
    AddressFamily []BgpStateData_AddressFamilies_AddressFamily
}

func (addressFamilies *BgpStateData_AddressFamilies) GetFilter() yfilter.YFilter { return addressFamilies.YFilter }

func (addressFamilies *BgpStateData_AddressFamilies) SetFilter(yf yfilter.YFilter) { addressFamilies.YFilter = yf }

func (addressFamilies *BgpStateData_AddressFamilies) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    return ""
}

func (addressFamilies *BgpStateData_AddressFamilies) GetSegmentPath() string {
    return "address-families"
}

func (addressFamilies *BgpStateData_AddressFamilies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-family" {
        for _, c := range addressFamilies.AddressFamily {
            if addressFamilies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_AddressFamilies_AddressFamily{}
        addressFamilies.AddressFamily = append(addressFamilies.AddressFamily, child)
        return &addressFamilies.AddressFamily[len(addressFamilies.AddressFamily)-1]
    }
    return nil
}

func (addressFamilies *BgpStateData_AddressFamilies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressFamilies.AddressFamily {
        children[addressFamilies.AddressFamily[i].GetSegmentPath()] = &addressFamilies.AddressFamily[i]
    }
    return children
}

func (addressFamilies *BgpStateData_AddressFamilies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressFamilies *BgpStateData_AddressFamilies) GetBundleName() string { return "cisco_ios_xe" }

func (addressFamilies *BgpStateData_AddressFamilies) GetYangName() string { return "address-families" }

func (addressFamilies *BgpStateData_AddressFamilies) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressFamilies *BgpStateData_AddressFamilies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressFamilies *BgpStateData_AddressFamilies) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressFamilies *BgpStateData_AddressFamilies) SetParent(parent types.Entity) { addressFamilies.parent = parent }

func (addressFamilies *BgpStateData_AddressFamilies) GetParent() types.Entity { return addressFamilies.parent }

func (addressFamilies *BgpStateData_AddressFamilies) GetParentYangName() string { return "bgp-state-data" }

// BgpStateData_AddressFamilies_AddressFamily
// List of BGP address families
type BgpStateData_AddressFamilies_AddressFamily struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Afi-safi value. The type is AfiSafi.
    AfiSafi interface{}

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // Router ID. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // BGP table version number. The type is interface{} with range:
    // 0..18446744073709551615.
    BgpTableVersion interface{}

    // Routing table version number. The type is interface{} with range:
    // 0..18446744073709551615.
    RoutingTableVersion interface{}

    // Total memory in use. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalMemory interface{}

    // Prefix entry statistics.
    Prefixes BgpStateData_AddressFamilies_AddressFamily_Prefixes

    // Path entry statistics.
    Path BgpStateData_AddressFamilies_AddressFamily_Path

    // AS path entry statistics.
    AsPath BgpStateData_AddressFamilies_AddressFamily_AsPath

    // Route map entry statistics.
    RouteMap BgpStateData_AddressFamilies_AddressFamily_RouteMap

    // Filter list entry statistics.
    FilterList BgpStateData_AddressFamilies_AddressFamily_FilterList

    // BGP activity information.
    Activities BgpStateData_AddressFamilies_AddressFamily_Activities

    // Neighbor summary.
    BgpNeighborSummaries BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetFilter() yfilter.YFilter { return addressFamily.YFilter }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) SetFilter(yf yfilter.YFilter) { addressFamily.YFilter = yf }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "router-id" { return "RouterId" }
    if yname == "bgp-table-version" { return "BgpTableVersion" }
    if yname == "routing-table-version" { return "RoutingTableVersion" }
    if yname == "total-memory" { return "TotalMemory" }
    if yname == "prefixes" { return "Prefixes" }
    if yname == "path" { return "Path" }
    if yname == "as-path" { return "AsPath" }
    if yname == "route-map" { return "RouteMap" }
    if yname == "filter-list" { return "FilterList" }
    if yname == "activities" { return "Activities" }
    if yname == "bgp-neighbor-summaries" { return "BgpNeighborSummaries" }
    return ""
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetSegmentPath() string {
    return "address-family" + "[afi-safi='" + fmt.Sprintf("%v", addressFamily.AfiSafi) + "']" + "[vrf-name='" + fmt.Sprintf("%v", addressFamily.VrfName) + "']"
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &addressFamily.Prefixes
    }
    if childYangName == "path" {
        return &addressFamily.Path
    }
    if childYangName == "as-path" {
        return &addressFamily.AsPath
    }
    if childYangName == "route-map" {
        return &addressFamily.RouteMap
    }
    if childYangName == "filter-list" {
        return &addressFamily.FilterList
    }
    if childYangName == "activities" {
        return &addressFamily.Activities
    }
    if childYangName == "bgp-neighbor-summaries" {
        return &addressFamily.BgpNeighborSummaries
    }
    return nil
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &addressFamily.Prefixes
    children["path"] = &addressFamily.Path
    children["as-path"] = &addressFamily.AsPath
    children["route-map"] = &addressFamily.RouteMap
    children["filter-list"] = &addressFamily.FilterList
    children["activities"] = &addressFamily.Activities
    children["bgp-neighbor-summaries"] = &addressFamily.BgpNeighborSummaries
    return children
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi"] = addressFamily.AfiSafi
    leafs["vrf-name"] = addressFamily.VrfName
    leafs["router-id"] = addressFamily.RouterId
    leafs["bgp-table-version"] = addressFamily.BgpTableVersion
    leafs["routing-table-version"] = addressFamily.RoutingTableVersion
    leafs["total-memory"] = addressFamily.TotalMemory
    return leafs
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetBundleName() string { return "cisco_ios_xe" }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetYangName() string { return "address-family" }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) SetParent(parent types.Entity) { addressFamily.parent = parent }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetParent() types.Entity { return addressFamily.parent }

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetParentYangName() string { return "address-families" }

// BgpStateData_AddressFamilies_AddressFamily_Prefixes
// Prefix entry statistics
type BgpStateData_AddressFamilies_AddressFamily_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetGoName(yname string) string {
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "memory-usage" { return "MemoryUsage" }
    return ""
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-entries"] = prefixes.TotalEntries
    leafs["memory-usage"] = prefixes.MemoryUsage
    return leafs
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetBundleName() string { return "cisco_ios_xe" }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_Path
// Path entry statistics
type BgpStateData_AddressFamilies_AddressFamily_Path struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetGoName(yname string) string {
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "memory-usage" { return "MemoryUsage" }
    return ""
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetSegmentPath() string {
    return "path"
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-entries"] = path.TotalEntries
    leafs["memory-usage"] = path.MemoryUsage
    return leafs
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetBundleName() string { return "cisco_ios_xe" }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetYangName() string { return "path" }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetParent() types.Entity { return path.parent }

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_AsPath
// AS path entry statistics
type BgpStateData_AddressFamilies_AddressFamily_AsPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetFilter() yfilter.YFilter { return asPath.YFilter }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) SetFilter(yf yfilter.YFilter) { asPath.YFilter = yf }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetGoName(yname string) string {
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "memory-usage" { return "MemoryUsage" }
    return ""
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetSegmentPath() string {
    return "as-path"
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-entries"] = asPath.TotalEntries
    leafs["memory-usage"] = asPath.MemoryUsage
    return leafs
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetBundleName() string { return "cisco_ios_xe" }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetYangName() string { return "as-path" }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) SetParent(parent types.Entity) { asPath.parent = parent }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetParent() types.Entity { return asPath.parent }

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_RouteMap
// Route map entry statistics
type BgpStateData_AddressFamilies_AddressFamily_RouteMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetFilter() yfilter.YFilter { return routeMap.YFilter }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) SetFilter(yf yfilter.YFilter) { routeMap.YFilter = yf }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetGoName(yname string) string {
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "memory-usage" { return "MemoryUsage" }
    return ""
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetSegmentPath() string {
    return "route-map"
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-entries"] = routeMap.TotalEntries
    leafs["memory-usage"] = routeMap.MemoryUsage
    return leafs
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetBundleName() string { return "cisco_ios_xe" }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetYangName() string { return "route-map" }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) SetParent(parent types.Entity) { routeMap.parent = parent }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetParent() types.Entity { return routeMap.parent }

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_FilterList
// Filter list entry statistics
type BgpStateData_AddressFamilies_AddressFamily_FilterList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetFilter() yfilter.YFilter { return filterList.YFilter }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) SetFilter(yf yfilter.YFilter) { filterList.YFilter = yf }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetGoName(yname string) string {
    if yname == "total-entries" { return "TotalEntries" }
    if yname == "memory-usage" { return "MemoryUsage" }
    return ""
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetSegmentPath() string {
    return "filter-list"
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-entries"] = filterList.TotalEntries
    leafs["memory-usage"] = filterList.MemoryUsage
    return leafs
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetBundleName() string { return "cisco_ios_xe" }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetYangName() string { return "filter-list" }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) SetParent(parent types.Entity) { filterList.parent = parent }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetParent() types.Entity { return filterList.parent }

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_Activities
// BGP activity information
type BgpStateData_AddressFamilies_AddressFamily_Activities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of prefixes. The type is interface{} with range:
    // 0..18446744073709551615.
    Prefixes interface{}

    // Total number of paths. The type is interface{} with range:
    // 0..18446744073709551615.
    Paths interface{}

    // Scan interval in seconds. The type is string.
    ScanInterval interface{}
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetFilter() yfilter.YFilter { return activities.YFilter }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) SetFilter(yf yfilter.YFilter) { activities.YFilter = yf }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetGoName(yname string) string {
    if yname == "prefixes" { return "Prefixes" }
    if yname == "paths" { return "Paths" }
    if yname == "scan-interval" { return "ScanInterval" }
    return ""
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetSegmentPath() string {
    return "activities"
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefixes"] = activities.Prefixes
    leafs["paths"] = activities.Paths
    leafs["scan-interval"] = activities.ScanInterval
    return leafs
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetBundleName() string { return "cisco_ios_xe" }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetYangName() string { return "activities" }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) SetParent(parent types.Entity) { activities.parent = parent }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetParent() types.Entity { return activities.parent }

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries
// Neighbor summary
type BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of neighbor summaries. The type is slice of
    // BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary.
    BgpNeighborSummary []BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetFilter() yfilter.YFilter { return bgpNeighborSummaries.YFilter }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) SetFilter(yf yfilter.YFilter) { bgpNeighborSummaries.YFilter = yf }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetGoName(yname string) string {
    if yname == "bgp-neighbor-summary" { return "BgpNeighborSummary" }
    return ""
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetSegmentPath() string {
    return "bgp-neighbor-summaries"
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-neighbor-summary" {
        for _, c := range bgpNeighborSummaries.BgpNeighborSummary {
            if bgpNeighborSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary{}
        bgpNeighborSummaries.BgpNeighborSummary = append(bgpNeighborSummaries.BgpNeighborSummary, child)
        return &bgpNeighborSummaries.BgpNeighborSummary[len(bgpNeighborSummaries.BgpNeighborSummary)-1]
    }
    return nil
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpNeighborSummaries.BgpNeighborSummary {
        children[bgpNeighborSummaries.BgpNeighborSummary[i].GetSegmentPath()] = &bgpNeighborSummaries.BgpNeighborSummary[i]
    }
    return children
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetBundleName() string { return "cisco_ios_xe" }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetYangName() string { return "bgp-neighbor-summaries" }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) SetParent(parent types.Entity) { bgpNeighborSummaries.parent = parent }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetParent() types.Entity { return bgpNeighborSummaries.parent }

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetParentYangName() string { return "address-family" }

// BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary
// List of neighbor summaries
type BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor address. The type is string.
    Id interface{}

    // BGP protocol version. The type is interface{} with range: 0..4294967295.
    BgpVersion interface{}

    // Number of messages received from this neighbor. The type is interface{}
    // with range: 0..18446744073709551615.
    MessagesReceived interface{}

    // Number of messages sent to this neighbor. The type is interface{} with
    // range: 0..18446744073709551615.
    MessagesSent interface{}

    // BGP table version. The type is interface{} with range:
    // 0..18446744073709551615.
    TableVersion interface{}

    // Number of messages in input queue. The type is interface{} with range:
    // 0..18446744073709551615.
    InputQueue interface{}

    // Number of messages in output queue. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputQueue interface{}

    // Neighbor session uptime. The type is string.
    UpTime interface{}

    // BGP session state. The type is BgpFsmState.
    State interface{}

    // Number of prefixes received from the neighbor. The type is interface{} with
    // range: 0..18446744073709551615.
    PrefixesReceived interface{}

    // Indication of whether the neighbor was dyanmically configured. The type is
    // bool.
    DynamicallyConfigured interface{}
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetFilter() yfilter.YFilter { return bgpNeighborSummary.YFilter }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) SetFilter(yf yfilter.YFilter) { bgpNeighborSummary.YFilter = yf }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "bgp-version" { return "BgpVersion" }
    if yname == "messages-received" { return "MessagesReceived" }
    if yname == "messages-sent" { return "MessagesSent" }
    if yname == "table-version" { return "TableVersion" }
    if yname == "input-queue" { return "InputQueue" }
    if yname == "output-queue" { return "OutputQueue" }
    if yname == "up-time" { return "UpTime" }
    if yname == "state" { return "State" }
    if yname == "prefixes-received" { return "PrefixesReceived" }
    if yname == "dynamically-configured" { return "DynamicallyConfigured" }
    return ""
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetSegmentPath() string {
    return "bgp-neighbor-summary" + "[id='" + fmt.Sprintf("%v", bgpNeighborSummary.Id) + "']"
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = bgpNeighborSummary.Id
    leafs["bgp-version"] = bgpNeighborSummary.BgpVersion
    leafs["messages-received"] = bgpNeighborSummary.MessagesReceived
    leafs["messages-sent"] = bgpNeighborSummary.MessagesSent
    leafs["table-version"] = bgpNeighborSummary.TableVersion
    leafs["input-queue"] = bgpNeighborSummary.InputQueue
    leafs["output-queue"] = bgpNeighborSummary.OutputQueue
    leafs["up-time"] = bgpNeighborSummary.UpTime
    leafs["state"] = bgpNeighborSummary.State
    leafs["prefixes-received"] = bgpNeighborSummary.PrefixesReceived
    leafs["dynamically-configured"] = bgpNeighborSummary.DynamicallyConfigured
    return leafs
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetBundleName() string { return "cisco_ios_xe" }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetYangName() string { return "bgp-neighbor-summary" }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) SetParent(parent types.Entity) { bgpNeighborSummary.parent = parent }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetParent() types.Entity { return bgpNeighborSummary.parent }

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetParentYangName() string { return "bgp-neighbor-summaries" }

// BgpStateData_BgpRouteVrfs
// BGP VRFs
type BgpStateData_BgpRouteVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP VRFs. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf.
    BgpRouteVrf []BgpStateData_BgpRouteVrfs_BgpRouteVrf
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetFilter() yfilter.YFilter { return bgpRouteVrfs.YFilter }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) SetFilter(yf yfilter.YFilter) { bgpRouteVrfs.YFilter = yf }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetGoName(yname string) string {
    if yname == "bgp-route-vrf" { return "BgpRouteVrf" }
    return ""
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetSegmentPath() string {
    return "bgp-route-vrfs"
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-vrf" {
        for _, c := range bgpRouteVrfs.BgpRouteVrf {
            if bgpRouteVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_BgpRouteVrfs_BgpRouteVrf{}
        bgpRouteVrfs.BgpRouteVrf = append(bgpRouteVrfs.BgpRouteVrf, child)
        return &bgpRouteVrfs.BgpRouteVrf[len(bgpRouteVrfs.BgpRouteVrf)-1]
    }
    return nil
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpRouteVrfs.BgpRouteVrf {
        children[bgpRouteVrfs.BgpRouteVrf[i].GetSegmentPath()] = &bgpRouteVrfs.BgpRouteVrf[i]
    }
    return children
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetYangName() string { return "bgp-route-vrfs" }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) SetParent(parent types.Entity) { bgpRouteVrfs.parent = parent }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetParent() types.Entity { return bgpRouteVrfs.parent }

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetParentYangName() string { return "bgp-state-data" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf
// List of BGP VRFs
type BgpStateData_BgpRouteVrfs_BgpRouteVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BGP vrf. The type is string.
    Vrf interface{}

    // BGP address families.
    BgpRouteAfs BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetFilter() yfilter.YFilter { return bgpRouteVrf.YFilter }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) SetFilter(yf yfilter.YFilter) { bgpRouteVrf.YFilter = yf }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "bgp-route-afs" { return "BgpRouteAfs" }
    return ""
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetSegmentPath() string {
    return "bgp-route-vrf" + "[vrf='" + fmt.Sprintf("%v", bgpRouteVrf.Vrf) + "']"
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-afs" {
        return &bgpRouteVrf.BgpRouteAfs
    }
    return nil
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-route-afs"] = &bgpRouteVrf.BgpRouteAfs
    return children
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = bgpRouteVrf.Vrf
    return leafs
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetYangName() string { return "bgp-route-vrf" }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) SetParent(parent types.Entity) { bgpRouteVrf.parent = parent }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetParent() types.Entity { return bgpRouteVrf.parent }

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetParentYangName() string { return "bgp-route-vrfs" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs
// BGP address families
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP address families. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf.
    BgpRouteAf []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetFilter() yfilter.YFilter { return bgpRouteAfs.YFilter }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) SetFilter(yf yfilter.YFilter) { bgpRouteAfs.YFilter = yf }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetGoName(yname string) string {
    if yname == "bgp-route-af" { return "BgpRouteAf" }
    return ""
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetSegmentPath() string {
    return "bgp-route-afs"
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-af" {
        for _, c := range bgpRouteAfs.BgpRouteAf {
            if bgpRouteAfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf{}
        bgpRouteAfs.BgpRouteAf = append(bgpRouteAfs.BgpRouteAf, child)
        return &bgpRouteAfs.BgpRouteAf[len(bgpRouteAfs.BgpRouteAf)-1]
    }
    return nil
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpRouteAfs.BgpRouteAf {
        children[bgpRouteAfs.BgpRouteAf[i].GetSegmentPath()] = &bgpRouteAfs.BgpRouteAf[i]
    }
    return children
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetYangName() string { return "bgp-route-afs" }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) SetParent(parent types.Entity) { bgpRouteAfs.parent = parent }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetParent() types.Entity { return bgpRouteAfs.parent }

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetParentYangName() string { return "bgp-route-vrf" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf
// List of BGP address families
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BGP address family. The type is AfiSafi.
    AfiSafi interface{}

    // BGP route filters.
    BgpRouteFilters BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetFilter() yfilter.YFilter { return bgpRouteAf.YFilter }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) SetFilter(yf yfilter.YFilter) { bgpRouteAf.YFilter = yf }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    if yname == "bgp-route-filters" { return "BgpRouteFilters" }
    return ""
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetSegmentPath() string {
    return "bgp-route-af" + "[afi-safi='" + fmt.Sprintf("%v", bgpRouteAf.AfiSafi) + "']"
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-filters" {
        return &bgpRouteAf.BgpRouteFilters
    }
    return nil
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-route-filters"] = &bgpRouteAf.BgpRouteFilters
    return children
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi"] = bgpRouteAf.AfiSafi
    return leafs
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetYangName() string { return "bgp-route-af" }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) SetParent(parent types.Entity) { bgpRouteAf.parent = parent }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetParent() types.Entity { return bgpRouteAf.parent }

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetParentYangName() string { return "bgp-route-afs" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters
// BGP route filters
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP route filters. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter.
    BgpRouteFilter []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetFilter() yfilter.YFilter { return bgpRouteFilters.YFilter }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) SetFilter(yf yfilter.YFilter) { bgpRouteFilters.YFilter = yf }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetGoName(yname string) string {
    if yname == "bgp-route-filter" { return "BgpRouteFilter" }
    return ""
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetSegmentPath() string {
    return "bgp-route-filters"
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-filter" {
        for _, c := range bgpRouteFilters.BgpRouteFilter {
            if bgpRouteFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter{}
        bgpRouteFilters.BgpRouteFilter = append(bgpRouteFilters.BgpRouteFilter, child)
        return &bgpRouteFilters.BgpRouteFilter[len(bgpRouteFilters.BgpRouteFilter)-1]
    }
    return nil
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpRouteFilters.BgpRouteFilter {
        children[bgpRouteFilters.BgpRouteFilter[i].GetSegmentPath()] = &bgpRouteFilters.BgpRouteFilter[i]
    }
    return children
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetYangName() string { return "bgp-route-filters" }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) SetParent(parent types.Entity) { bgpRouteFilters.parent = parent }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetParent() types.Entity { return bgpRouteFilters.parent }

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetParentYangName() string { return "bgp-route-af" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter
// List of BGP route filters
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BGP route filter. The type is BgpRouteFilters.
    RouteFilter interface{}

    // BGP route entries.
    BgpRouteEntries BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetFilter() yfilter.YFilter { return bgpRouteFilter.YFilter }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) SetFilter(yf yfilter.YFilter) { bgpRouteFilter.YFilter = yf }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetGoName(yname string) string {
    if yname == "route-filter" { return "RouteFilter" }
    if yname == "bgp-route-entries" { return "BgpRouteEntries" }
    return ""
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetSegmentPath() string {
    return "bgp-route-filter" + "[route-filter='" + fmt.Sprintf("%v", bgpRouteFilter.RouteFilter) + "']"
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-entries" {
        return &bgpRouteFilter.BgpRouteEntries
    }
    return nil
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-route-entries"] = &bgpRouteFilter.BgpRouteEntries
    return children
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-filter"] = bgpRouteFilter.RouteFilter
    return leafs
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetYangName() string { return "bgp-route-filter" }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) SetParent(parent types.Entity) { bgpRouteFilter.parent = parent }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetParent() types.Entity { return bgpRouteFilter.parent }

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetParentYangName() string { return "bgp-route-filters" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries
// BGP route entries
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP route entries. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry.
    BgpRouteEntry []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetFilter() yfilter.YFilter { return bgpRouteEntries.YFilter }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) SetFilter(yf yfilter.YFilter) { bgpRouteEntries.YFilter = yf }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetGoName(yname string) string {
    if yname == "bgp-route-entry" { return "BgpRouteEntry" }
    return ""
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetSegmentPath() string {
    return "bgp-route-entries"
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-route-entry" {
        for _, c := range bgpRouteEntries.BgpRouteEntry {
            if bgpRouteEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry{}
        bgpRouteEntries.BgpRouteEntry = append(bgpRouteEntries.BgpRouteEntry, child)
        return &bgpRouteEntries.BgpRouteEntry[len(bgpRouteEntries.BgpRouteEntry)-1]
    }
    return nil
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpRouteEntries.BgpRouteEntry {
        children[bgpRouteEntries.BgpRouteEntry[i].GetSegmentPath()] = &bgpRouteEntries.BgpRouteEntry[i]
    }
    return children
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetYangName() string { return "bgp-route-entries" }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) SetParent(parent types.Entity) { bgpRouteEntries.parent = parent }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetParent() types.Entity { return bgpRouteEntries.parent }

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetParentYangName() string { return "bgp-route-filter" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry
// List of BGP route entries
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Routing table entry prefix. The type is string.
    Prefix interface{}

    // Routing table version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Number of paths available for BGP prefix. The type is interface{} with
    // range: 0..4294967295.
    AvailablePaths interface{}

    // Whom is thie prefix advertized to. The type is string.
    AdvertisedTo interface{}

    // Prefix next hop details.
    BgpPathEntries BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetFilter() yfilter.YFilter { return bgpRouteEntry.YFilter }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) SetFilter(yf yfilter.YFilter) { bgpRouteEntry.YFilter = yf }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "version" { return "Version" }
    if yname == "available-paths" { return "AvailablePaths" }
    if yname == "advertised-to" { return "AdvertisedTo" }
    if yname == "bgp-path-entries" { return "BgpPathEntries" }
    return ""
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetSegmentPath() string {
    return "bgp-route-entry" + "[prefix='" + fmt.Sprintf("%v", bgpRouteEntry.Prefix) + "']"
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-path-entries" {
        return &bgpRouteEntry.BgpPathEntries
    }
    return nil
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-path-entries"] = &bgpRouteEntry.BgpPathEntries
    return children
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = bgpRouteEntry.Prefix
    leafs["version"] = bgpRouteEntry.Version
    leafs["available-paths"] = bgpRouteEntry.AvailablePaths
    leafs["advertised-to"] = bgpRouteEntry.AdvertisedTo
    return leafs
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetYangName() string { return "bgp-route-entry" }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) SetParent(parent types.Entity) { bgpRouteEntry.parent = parent }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetParent() types.Entity { return bgpRouteEntry.parent }

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetParentYangName() string { return "bgp-route-entries" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries
// Prefix next hop details
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of prefix next hop details. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry.
    BgpPathEntry []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetFilter() yfilter.YFilter { return bgpPathEntries.YFilter }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) SetFilter(yf yfilter.YFilter) { bgpPathEntries.YFilter = yf }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetGoName(yname string) string {
    if yname == "bgp-path-entry" { return "BgpPathEntry" }
    return ""
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetSegmentPath() string {
    return "bgp-path-entries"
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-path-entry" {
        for _, c := range bgpPathEntries.BgpPathEntry {
            if bgpPathEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry{}
        bgpPathEntries.BgpPathEntry = append(bgpPathEntries.BgpPathEntry, child)
        return &bgpPathEntries.BgpPathEntry[len(bgpPathEntries.BgpPathEntry)-1]
    }
    return nil
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpPathEntries.BgpPathEntry {
        children[bgpPathEntries.BgpPathEntry[i].GetSegmentPath()] = &bgpPathEntries.BgpPathEntry[i]
    }
    return children
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetBundleName() string { return "cisco_ios_xe" }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetYangName() string { return "bgp-path-entries" }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) SetParent(parent types.Entity) { bgpPathEntries.parent = parent }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetParent() types.Entity { return bgpPathEntries.parent }

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetParentYangName() string { return "bgp-route-entry" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry
// List of prefix next hop details
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Next hop for this path. The type is string.
    Nexthop interface{}

    // Metric associated with the path. The type is interface{} with range:
    // 0..4294967295.
    Metric interface{}

    // Local preference for this path. The type is interface{} with range:
    // 0..4294967295.
    LocalPref interface{}

    // Path weight. The type is interface{} with range: 0..4294967295.
    Weight interface{}

    // AS path. The type is string.
    AsPath interface{}

    // Path origin. The type is BgpOriginCode.
    Origin interface{}

    // RPKI path status. The type is BgpRpkiStatus.
    RpkiStatus interface{}

    // Community label for the path. The type is string.
    Community interface{}

    // Path status.
    PathStatus BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetFilter() yfilter.YFilter { return bgpPathEntry.YFilter }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) SetFilter(yf yfilter.YFilter) { bgpPathEntry.YFilter = yf }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetGoName(yname string) string {
    if yname == "nexthop" { return "Nexthop" }
    if yname == "metric" { return "Metric" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "weight" { return "Weight" }
    if yname == "as-path" { return "AsPath" }
    if yname == "origin" { return "Origin" }
    if yname == "rpki-status" { return "RpkiStatus" }
    if yname == "community" { return "Community" }
    if yname == "path-status" { return "PathStatus" }
    return ""
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetSegmentPath() string {
    return "bgp-path-entry" + "[nexthop='" + fmt.Sprintf("%v", bgpPathEntry.Nexthop) + "']"
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-status" {
        return &bgpPathEntry.PathStatus
    }
    return nil
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-status"] = &bgpPathEntry.PathStatus
    return children
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nexthop"] = bgpPathEntry.Nexthop
    leafs["metric"] = bgpPathEntry.Metric
    leafs["local-pref"] = bgpPathEntry.LocalPref
    leafs["weight"] = bgpPathEntry.Weight
    leafs["as-path"] = bgpPathEntry.AsPath
    leafs["origin"] = bgpPathEntry.Origin
    leafs["rpki-status"] = bgpPathEntry.RpkiStatus
    leafs["community"] = bgpPathEntry.Community
    return leafs
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetBundleName() string { return "cisco_ios_xe" }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetYangName() string { return "bgp-path-entry" }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) SetParent(parent types.Entity) { bgpPathEntry.parent = parent }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetParent() types.Entity { return bgpPathEntry.parent }

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetParentYangName() string { return "bgp-path-entries" }

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus
// Path status
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Suppressed path. The type is interface{}.
    Suppressed interface{}

    // Damped path. The type is interface{}.
    Damped interface{}

    // History path. The type is interface{}.
    History interface{}

    // Valid path. The type is interface{}.
    Valid interface{}

    // Sourced path. The type is interface{}.
    Sourced interface{}

    // Best path. The type is interface{}.
    Bestpath interface{}

    // Internal path. The type is interface{}.
    Internal interface{}

    // RIB-fail path. The type is interface{}.
    RibFail interface{}

    // Stale path. The type is interface{}.
    Stale interface{}

    // Multipath path. The type is interface{}.
    Multipath interface{}

    // Backup path. The type is interface{}.
    BackupPath interface{}

    // RT filter path. The type is interface{}.
    RtFilter interface{}

    // Best externa path. The type is interface{}.
    BestExternal interface{}

    // Additional path. The type is interface{}.
    AdditionalPath interface{}

    // RIB compressed path. The type is interface{}.
    RibCompressed interface{}
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetFilter() yfilter.YFilter { return pathStatus.YFilter }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) SetFilter(yf yfilter.YFilter) { pathStatus.YFilter = yf }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetGoName(yname string) string {
    if yname == "suppressed" { return "Suppressed" }
    if yname == "damped" { return "Damped" }
    if yname == "history" { return "History" }
    if yname == "valid" { return "Valid" }
    if yname == "sourced" { return "Sourced" }
    if yname == "bestpath" { return "Bestpath" }
    if yname == "internal" { return "Internal" }
    if yname == "rib-fail" { return "RibFail" }
    if yname == "stale" { return "Stale" }
    if yname == "multipath" { return "Multipath" }
    if yname == "backup-path" { return "BackupPath" }
    if yname == "rt-filter" { return "RtFilter" }
    if yname == "best-external" { return "BestExternal" }
    if yname == "additional-path" { return "AdditionalPath" }
    if yname == "rib-compressed" { return "RibCompressed" }
    return ""
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetSegmentPath() string {
    return "path-status"
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["suppressed"] = pathStatus.Suppressed
    leafs["damped"] = pathStatus.Damped
    leafs["history"] = pathStatus.History
    leafs["valid"] = pathStatus.Valid
    leafs["sourced"] = pathStatus.Sourced
    leafs["bestpath"] = pathStatus.Bestpath
    leafs["internal"] = pathStatus.Internal
    leafs["rib-fail"] = pathStatus.RibFail
    leafs["stale"] = pathStatus.Stale
    leafs["multipath"] = pathStatus.Multipath
    leafs["backup-path"] = pathStatus.BackupPath
    leafs["rt-filter"] = pathStatus.RtFilter
    leafs["best-external"] = pathStatus.BestExternal
    leafs["additional-path"] = pathStatus.AdditionalPath
    leafs["rib-compressed"] = pathStatus.RibCompressed
    return leafs
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetBundleName() string { return "cisco_ios_xe" }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetYangName() string { return "path-status" }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) SetParent(parent types.Entity) { pathStatus.parent = parent }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetParent() types.Entity { return pathStatus.parent }

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetParentYangName() string { return "bgp-path-entry" }

