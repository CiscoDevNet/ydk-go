// This module contains a collection of YANG definitions
// for Cisco IOS-XR ptp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ptp: PTP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ptp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ptp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ptp-oper ptp}", reflect.TypeOf(Ptp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ptp-oper:ptp", reflect.TypeOf(Ptp{}))
}

// PtpBagRestrictPortState represents Restrict Port State
type PtpBagRestrictPortState string

const (
    // Any
    PtpBagRestrictPortState_any PtpBagRestrictPortState = "any"

    // Slave only
    PtpBagRestrictPortState_slave_only PtpBagRestrictPortState = "slave-only"

    // Master only
    PtpBagRestrictPortState_master_only PtpBagRestrictPortState = "master-only"
)

// PtpBagPortState represents Port State
type PtpBagPortState string

const (
    // Initializing state
    PtpBagPortState_initializing PtpBagPortState = "initializing"

    // Listen state
    PtpBagPortState_listen PtpBagPortState = "listen"

    // Passive state
    PtpBagPortState_passive PtpBagPortState = "passive"

    // Pre-Master state
    PtpBagPortState_pre_master PtpBagPortState = "pre-master"

    // Master state
    PtpBagPortState_master PtpBagPortState = "master"

    // Uncalibrated state
    PtpBagPortState_uncalibrated PtpBagPortState = "uncalibrated"

    // Slave state
    PtpBagPortState_slave PtpBagPortState = "slave"

    // Faulty state
    PtpBagPortState_faulty PtpBagPortState = "faulty"
)

// PtpBagEncap represents Encapsulation
type PtpBagEncap string

const (
    // Unknown encapsulation
    PtpBagEncap_unknown PtpBagEncap = "unknown"

    // Ethernet encapsulation
    PtpBagEncap_ethernet PtpBagEncap = "ethernet"

    // IPv4 encapsulation
    PtpBagEncap_ipv4 PtpBagEncap = "ipv4"

    // IPv6 encapsulation
    PtpBagEncap_ipv6 PtpBagEncap = "ipv6"
)

// PtpBagDelayMechanism represents Delay Mechanism
type PtpBagDelayMechanism string

const (
    // End to end delay mechanism
    PtpBagDelayMechanism_e2e PtpBagDelayMechanism = "e2e"

    // Peer to peer delay mechanism
    PtpBagDelayMechanism_p2p PtpBagDelayMechanism = "p2p"
)

// PtpBagTelecomClock represents Telecom Clock
type PtpBagTelecomClock string

const (
    // Grandmaster
    PtpBagTelecomClock_grandmaster PtpBagTelecomClock = "grandmaster"

    // Boundary
    PtpBagTelecomClock_boundary PtpBagTelecomClock = "boundary"

    // Slave
    PtpBagTelecomClock_slave PtpBagTelecomClock = "slave"
)

// PtpBagCommunicationModel represents Communication Model
type PtpBagCommunicationModel string

const (
    // Unication communication
    PtpBagCommunicationModel_unicast PtpBagCommunicationModel = "unicast"

    // Mixed-mode communication
    PtpBagCommunicationModel_mixed_mode PtpBagCommunicationModel = "mixed-mode"

    // Multicast communication
    PtpBagCommunicationModel_multicast PtpBagCommunicationModel = "multicast"
)

// PtpBagClockTimeSource represents Time source
type PtpBagClockTimeSource string

const (
    // Unknown
    PtpBagClockTimeSource_unknown PtpBagClockTimeSource = "unknown"

    // Atomic clock
    PtpBagClockTimeSource_atomic PtpBagClockTimeSource = "atomic"

    // GPS clock
    PtpBagClockTimeSource_gps PtpBagClockTimeSource = "gps"

    // Terrestrial Radio
    PtpBagClockTimeSource_terrestrial_radio PtpBagClockTimeSource = "terrestrial-radio"

    // Precision Time Protocol
    PtpBagClockTimeSource_ptp PtpBagClockTimeSource = "ptp"

    // Network Time Protocol
    PtpBagClockTimeSource_ntp PtpBagClockTimeSource = "ntp"

    // Hand set
    PtpBagClockTimeSource_hand_set PtpBagClockTimeSource = "hand-set"

    // Other Time Source
    PtpBagClockTimeSource_other PtpBagClockTimeSource = "other"

    // Internal Oscillator
    PtpBagClockTimeSource_internal_oscillator PtpBagClockTimeSource = "internal-oscillator"
)

// PtpBagClockLeapSeconds represents Leap Seconds
type PtpBagClockLeapSeconds string

const (
    // No leap seconds
    PtpBagClockLeapSeconds_none PtpBagClockLeapSeconds = "none"

    // The last minute of the day has 59 seconds
    PtpBagClockLeapSeconds_leap59 PtpBagClockLeapSeconds = "leap59"

    // The last minute of the day has 61 seconds
    PtpBagClockLeapSeconds_leap61 PtpBagClockLeapSeconds = "leap61"
)

// PtpBagProfile represents Profile
type PtpBagProfile string

const (
    // 1588v2 profile (default)
    PtpBagProfile_default_ PtpBagProfile = "default"

    // G.8265.1 profile
    PtpBagProfile_g82651 PtpBagProfile = "g82651"

    // G.8275.1 profile
    PtpBagProfile_g82751 PtpBagProfile = "g82751"

    // G.8275.2 profile
    PtpBagProfile_g82752 PtpBagProfile = "g82752"
)

// PtpBagClockTimescale represents Timescale
type PtpBagClockTimescale string

const (
    // PTP timescale
    PtpBagClockTimescale_ptp PtpBagClockTimescale = "ptp"

    // ARB timescale
    PtpBagClockTimescale_arb PtpBagClockTimescale = "arb"
)

// Ptp
// PTP operational data
type Ptp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table for node-specific operational data.
    Nodes Ptp_Nodes

    // Table for interface configuration error operational data.
    InterfaceConfigurationErrors Ptp_InterfaceConfigurationErrors

    // Table for interface foreign master clock operational data.
    InterfaceForeignMasters Ptp_InterfaceForeignMasters

    // Local clock operational data.
    LocalClock Ptp_LocalClock

    // Table for interface packet counter operational data.
    InterfacePacketCounters Ptp_InterfacePacketCounters

    // Advertised clock operational data.
    AdvertisedClock Ptp_AdvertisedClock

    // Upcoming leap-seconds information.
    LeapSeconds Ptp_LeapSeconds

    // Table for interface operational data.
    Interfaces Ptp_Interfaces

    // Global PTP datasets.
    Dataset Ptp_Dataset

    // Global configuration error operational data.
    GlobalConfigurationError Ptp_GlobalConfigurationError

    // Grandmaster clock operational data.
    Grandmaster Ptp_Grandmaster

    // Table for interface unicast peers operational data.
    InterfaceUnicastPeers Ptp_InterfaceUnicastPeers

    // PTP platform specific data.
    Platform Ptp_Platform
}

func (ptp *Ptp) GetFilter() yfilter.YFilter { return ptp.YFilter }

func (ptp *Ptp) SetFilter(yf yfilter.YFilter) { ptp.YFilter = yf }

func (ptp *Ptp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "interface-configuration-errors" { return "InterfaceConfigurationErrors" }
    if yname == "interface-foreign-masters" { return "InterfaceForeignMasters" }
    if yname == "local-clock" { return "LocalClock" }
    if yname == "interface-packet-counters" { return "InterfacePacketCounters" }
    if yname == "advertised-clock" { return "AdvertisedClock" }
    if yname == "leap-seconds" { return "LeapSeconds" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "dataset" { return "Dataset" }
    if yname == "global-configuration-error" { return "GlobalConfigurationError" }
    if yname == "grandmaster" { return "Grandmaster" }
    if yname == "interface-unicast-peers" { return "InterfaceUnicastPeers" }
    if yname == "Cisco-IOS-XR-ptp-pd-oper:platform" { return "Platform" }
    return ""
}

func (ptp *Ptp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ptp-oper:ptp"
}

func (ptp *Ptp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ptp.Nodes
    }
    if childYangName == "interface-configuration-errors" {
        return &ptp.InterfaceConfigurationErrors
    }
    if childYangName == "interface-foreign-masters" {
        return &ptp.InterfaceForeignMasters
    }
    if childYangName == "local-clock" {
        return &ptp.LocalClock
    }
    if childYangName == "interface-packet-counters" {
        return &ptp.InterfacePacketCounters
    }
    if childYangName == "advertised-clock" {
        return &ptp.AdvertisedClock
    }
    if childYangName == "leap-seconds" {
        return &ptp.LeapSeconds
    }
    if childYangName == "interfaces" {
        return &ptp.Interfaces
    }
    if childYangName == "dataset" {
        return &ptp.Dataset
    }
    if childYangName == "global-configuration-error" {
        return &ptp.GlobalConfigurationError
    }
    if childYangName == "grandmaster" {
        return &ptp.Grandmaster
    }
    if childYangName == "interface-unicast-peers" {
        return &ptp.InterfaceUnicastPeers
    }
    if childYangName == "Cisco-IOS-XR-ptp-pd-oper:platform" {
        return &ptp.Platform
    }
    return nil
}

func (ptp *Ptp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ptp.Nodes
    children["interface-configuration-errors"] = &ptp.InterfaceConfigurationErrors
    children["interface-foreign-masters"] = &ptp.InterfaceForeignMasters
    children["local-clock"] = &ptp.LocalClock
    children["interface-packet-counters"] = &ptp.InterfacePacketCounters
    children["advertised-clock"] = &ptp.AdvertisedClock
    children["leap-seconds"] = &ptp.LeapSeconds
    children["interfaces"] = &ptp.Interfaces
    children["dataset"] = &ptp.Dataset
    children["global-configuration-error"] = &ptp.GlobalConfigurationError
    children["grandmaster"] = &ptp.Grandmaster
    children["interface-unicast-peers"] = &ptp.InterfaceUnicastPeers
    children["Cisco-IOS-XR-ptp-pd-oper:platform"] = &ptp.Platform
    return children
}

func (ptp *Ptp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ptp *Ptp) GetBundleName() string { return "cisco_ios_xr" }

func (ptp *Ptp) GetYangName() string { return "ptp" }

func (ptp *Ptp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ptp *Ptp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ptp *Ptp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ptp *Ptp) SetParent(parent types.Entity) { ptp.parent = parent }

func (ptp *Ptp) GetParent() types.Entity { return ptp.parent }

func (ptp *Ptp) GetParentYangName() string { return "Cisco-IOS-XR-ptp-oper" }

// Ptp_Nodes
// Table for node-specific operational data
type Ptp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific operational data for a given node. The type is slice of
    // Ptp_Nodes_Node.
    Node []Ptp_Nodes_Node
}

func (nodes *Ptp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ptp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ptp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ptp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ptp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ptp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ptp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ptp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ptp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ptp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ptp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ptp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ptp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ptp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ptp_Nodes) GetParentYangName() string { return "ptp" }

// Ptp_Nodes_Node
// Node-specific operational data for a given node
type Ptp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table for node foreign master clock operational data.
    NodeInterfaceForeignMasters Ptp_Nodes_Node_NodeInterfaceForeignMasters

    // Node summary operational data.
    Summary Ptp_Nodes_Node_Summary

    // Table for node interface operational data.
    NodeInterfaces Ptp_Nodes_Node_NodeInterfaces

    // Table for node unicast peers operational data.
    NodeInterfaceUnicastPeers Ptp_Nodes_Node_NodeInterfaceUnicastPeers

    // Node packet counter operational data.
    PacketCounters Ptp_Nodes_Node_PacketCounters
}

func (node *Ptp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ptp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ptp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "node-interface-foreign-masters" { return "NodeInterfaceForeignMasters" }
    if yname == "summary" { return "Summary" }
    if yname == "node-interfaces" { return "NodeInterfaces" }
    if yname == "node-interface-unicast-peers" { return "NodeInterfaceUnicastPeers" }
    if yname == "packet-counters" { return "PacketCounters" }
    return ""
}

func (node *Ptp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ptp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-interface-foreign-masters" {
        return &node.NodeInterfaceForeignMasters
    }
    if childYangName == "summary" {
        return &node.Summary
    }
    if childYangName == "node-interfaces" {
        return &node.NodeInterfaces
    }
    if childYangName == "node-interface-unicast-peers" {
        return &node.NodeInterfaceUnicastPeers
    }
    if childYangName == "packet-counters" {
        return &node.PacketCounters
    }
    return nil
}

func (node *Ptp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-interface-foreign-masters"] = &node.NodeInterfaceForeignMasters
    children["summary"] = &node.Summary
    children["node-interfaces"] = &node.NodeInterfaces
    children["node-interface-unicast-peers"] = &node.NodeInterfaceUnicastPeers
    children["packet-counters"] = &node.PacketCounters
    return children
}

func (node *Ptp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ptp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ptp_Nodes_Node) GetYangName() string { return "node" }

func (node *Ptp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ptp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ptp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ptp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ptp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ptp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters
// Table for node foreign master clock
// operational data
type Ptp_Nodes_Node_NodeInterfaceForeignMasters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node interface foreign master clock operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster.
    NodeInterfaceForeignMaster []Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetFilter() yfilter.YFilter { return nodeInterfaceForeignMasters.YFilter }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) SetFilter(yf yfilter.YFilter) { nodeInterfaceForeignMasters.YFilter = yf }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetGoName(yname string) string {
    if yname == "node-interface-foreign-master" { return "NodeInterfaceForeignMaster" }
    return ""
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetSegmentPath() string {
    return "node-interface-foreign-masters"
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-interface-foreign-master" {
        for _, c := range nodeInterfaceForeignMasters.NodeInterfaceForeignMaster {
            if nodeInterfaceForeignMasters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster{}
        nodeInterfaceForeignMasters.NodeInterfaceForeignMaster = append(nodeInterfaceForeignMasters.NodeInterfaceForeignMaster, child)
        return &nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[len(nodeInterfaceForeignMasters.NodeInterfaceForeignMaster)-1]
    }
    return nil
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaceForeignMasters.NodeInterfaceForeignMaster {
        children[nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[i].GetSegmentPath()] = &nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[i]
    }
    return children
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetYangName() string { return "node-interface-foreign-masters" }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) SetParent(parent types.Entity) { nodeInterfaceForeignMasters.parent = parent }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetParent() types.Entity { return nodeInterfaceForeignMasters.parent }

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetParentYangName() string { return "node" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster
// Node interface foreign master clock
// operational data
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Foreign clocks received on this interface. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock.
    ForeignClock []Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetFilter() yfilter.YFilter { return nodeInterfaceForeignMaster.YFilter }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) SetFilter(yf yfilter.YFilter) { nodeInterfaceForeignMaster.YFilter = yf }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "foreign-clock" { return "ForeignClock" }
    return ""
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetSegmentPath() string {
    return "node-interface-foreign-master" + "[interface-name='" + fmt.Sprintf("%v", nodeInterfaceForeignMaster.InterfaceName) + "']"
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "foreign-clock" {
        for _, c := range nodeInterfaceForeignMaster.ForeignClock {
            if nodeInterfaceForeignMaster.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock{}
        nodeInterfaceForeignMaster.ForeignClock = append(nodeInterfaceForeignMaster.ForeignClock, child)
        return &nodeInterfaceForeignMaster.ForeignClock[len(nodeInterfaceForeignMaster.ForeignClock)-1]
    }
    return nil
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaceForeignMaster.ForeignClock {
        children[nodeInterfaceForeignMaster.ForeignClock[i].GetSegmentPath()] = &nodeInterfaceForeignMaster.ForeignClock[i]
    }
    return children
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = nodeInterfaceForeignMaster.InterfaceName
    leafs["port-number"] = nodeInterfaceForeignMaster.PortNumber
    return leafs
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetYangName() string { return "node-interface-foreign-master" }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) SetParent(parent types.Entity) { nodeInterfaceForeignMaster.parent = parent }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetParent() types.Entity { return nodeInterfaceForeignMaster.parent }

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetParentYangName() string { return "node-interface-foreign-masters" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock
// Foreign clocks received on this interface
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The clock is qualified for best master clock selection. The type is bool.
    IsQualified interface{}

    // This clock is the currently selected grand master clock. The type is bool.
    IsGrandmaster interface{}

    // The communication model configured on this clock. The type is
    // PtpBagCommunicationModel.
    CommunicationModel interface{}

    // This clock is known by this router. The type is bool.
    IsKnown interface{}

    // How long the clock has been known by this router for, in seconds. The type
    // is interface{} with range: 0..4294967295. Units are second.
    TimeKnownFor interface{}

    // The PTP domain that the foreign clock is in. The type is interface{} with
    // range: 0..255.
    ForeignDomain interface{}

    // Priority configured for the clock, if any. The type is interface{} with
    // range: 0..255.
    ConfiguredPriority interface{}

    // Clock class configured for the clock, if any. The type is interface{} with
    // range: 0..255.
    ConfiguredClockClass interface{}

    // Delay asymmetry configured for the clock, if any. The type is interface{}
    // with range: -2147483648..2147483647.
    DelayAsymmetry interface{}

    // Announced messages are not being received from the master. The type is
    // bool.
    PtsfLossAnnounce interface{}

    // Sync messages are not being received from the master. The type is bool.
    PtsfLossSync interface{}

    // Foreign clock information.
    ForeignClock Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock

    // The address of the clock.
    Address Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address

    // Unicast grant information for announce messages.
    AnnounceGrant Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant

    // Unicast grant information for sync messages.
    SyncGrant Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant

    // Unicast grant information for delay-response messages.
    DelayResponseGrant Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetFilter() yfilter.YFilter { return foreignClock.YFilter }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) SetFilter(yf yfilter.YFilter) { foreignClock.YFilter = yf }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetGoName(yname string) string {
    if yname == "is-qualified" { return "IsQualified" }
    if yname == "is-grandmaster" { return "IsGrandmaster" }
    if yname == "communication-model" { return "CommunicationModel" }
    if yname == "is-known" { return "IsKnown" }
    if yname == "time-known-for" { return "TimeKnownFor" }
    if yname == "foreign-domain" { return "ForeignDomain" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "delay-asymmetry" { return "DelayAsymmetry" }
    if yname == "ptsf-loss-announce" { return "PtsfLossAnnounce" }
    if yname == "ptsf-loss-sync" { return "PtsfLossSync" }
    if yname == "foreign-clock" { return "ForeignClock" }
    if yname == "address" { return "Address" }
    if yname == "announce-grant" { return "AnnounceGrant" }
    if yname == "sync-grant" { return "SyncGrant" }
    if yname == "delay-response-grant" { return "DelayResponseGrant" }
    return ""
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetSegmentPath() string {
    return "foreign-clock"
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "foreign-clock" {
        return &foreignClock.ForeignClock
    }
    if childYangName == "address" {
        return &foreignClock.Address
    }
    if childYangName == "announce-grant" {
        return &foreignClock.AnnounceGrant
    }
    if childYangName == "sync-grant" {
        return &foreignClock.SyncGrant
    }
    if childYangName == "delay-response-grant" {
        return &foreignClock.DelayResponseGrant
    }
    return nil
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["foreign-clock"] = &foreignClock.ForeignClock
    children["address"] = &foreignClock.Address
    children["announce-grant"] = &foreignClock.AnnounceGrant
    children["sync-grant"] = &foreignClock.SyncGrant
    children["delay-response-grant"] = &foreignClock.DelayResponseGrant
    return children
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-qualified"] = foreignClock.IsQualified
    leafs["is-grandmaster"] = foreignClock.IsGrandmaster
    leafs["communication-model"] = foreignClock.CommunicationModel
    leafs["is-known"] = foreignClock.IsKnown
    leafs["time-known-for"] = foreignClock.TimeKnownFor
    leafs["foreign-domain"] = foreignClock.ForeignDomain
    leafs["configured-priority"] = foreignClock.ConfiguredPriority
    leafs["configured-clock-class"] = foreignClock.ConfiguredClockClass
    leafs["delay-asymmetry"] = foreignClock.DelayAsymmetry
    leafs["ptsf-loss-announce"] = foreignClock.PtsfLossAnnounce
    leafs["ptsf-loss-sync"] = foreignClock.PtsfLossSync
    return leafs
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetBundleName() string { return "cisco_ios_xr" }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetYangName() string { return "foreign-clock" }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) SetParent(parent types.Entity) { foreignClock.parent = parent }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetParent() types.Entity { return foreignClock.parent }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetParentYangName() string { return "node-interface-foreign-master" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock
// Foreign clock information
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Class. The type is interface{} with range: 0..255.
    Class interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Steps removed. The type is interface{} with range: 0..65535.
    StepsRemoved interface{}

    // Time source. The type is PtpBagClockTimeSource.
    TimeSource interface{}

    // The clock is frequency traceable. The type is bool.
    FrequencyTraceable interface{}

    // The clock is time traceable. The type is bool.
    TimeTraceable interface{}

    // Timescale. The type is PtpBagClockTimescale.
    Timescale interface{}

    // Leap Seconds. The type is PtpBagClockLeapSeconds. Units are second.
    LeapSeconds interface{}

    // The clock is the local clock. The type is bool.
    Local interface{}

    // The configured clock class. The type is interface{} with range: 0..255.
    ConfiguredClockClass interface{}

    // The configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // UTC offset.
    UtcOffset Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset

    // Receiver.
    Receiver Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver

    // Sender.
    Sender Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetFilter() yfilter.YFilter { return foreignClock.YFilter }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) SetFilter(yf yfilter.YFilter) { foreignClock.YFilter = yf }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "priority1" { return "Priority1" }
    if yname == "priority2" { return "Priority2" }
    if yname == "class" { return "Class" }
    if yname == "accuracy" { return "Accuracy" }
    if yname == "offset-log-variance" { return "OffsetLogVariance" }
    if yname == "steps-removed" { return "StepsRemoved" }
    if yname == "time-source" { return "TimeSource" }
    if yname == "frequency-traceable" { return "FrequencyTraceable" }
    if yname == "time-traceable" { return "TimeTraceable" }
    if yname == "timescale" { return "Timescale" }
    if yname == "leap-seconds" { return "LeapSeconds" }
    if yname == "local" { return "Local" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "utc-offset" { return "UtcOffset" }
    if yname == "receiver" { return "Receiver" }
    if yname == "sender" { return "Sender" }
    return ""
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetSegmentPath() string {
    return "foreign-clock"
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "utc-offset" {
        return &foreignClock.UtcOffset
    }
    if childYangName == "receiver" {
        return &foreignClock.Receiver
    }
    if childYangName == "sender" {
        return &foreignClock.Sender
    }
    return nil
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["utc-offset"] = &foreignClock.UtcOffset
    children["receiver"] = &foreignClock.Receiver
    children["sender"] = &foreignClock.Sender
    return children
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = foreignClock.ClockId
    leafs["priority1"] = foreignClock.Priority1
    leafs["priority2"] = foreignClock.Priority2
    leafs["class"] = foreignClock.Class
    leafs["accuracy"] = foreignClock.Accuracy
    leafs["offset-log-variance"] = foreignClock.OffsetLogVariance
    leafs["steps-removed"] = foreignClock.StepsRemoved
    leafs["time-source"] = foreignClock.TimeSource
    leafs["frequency-traceable"] = foreignClock.FrequencyTraceable
    leafs["time-traceable"] = foreignClock.TimeTraceable
    leafs["timescale"] = foreignClock.Timescale
    leafs["leap-seconds"] = foreignClock.LeapSeconds
    leafs["local"] = foreignClock.Local
    leafs["configured-clock-class"] = foreignClock.ConfiguredClockClass
    leafs["configured-priority"] = foreignClock.ConfiguredPriority
    return leafs
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetBundleName() string { return "cisco_ios_xr" }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetYangName() string { return "foreign-clock" }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) SetParent(parent types.Entity) { foreignClock.parent = parent }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetParent() types.Entity { return foreignClock.parent }

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset
// UTC offset
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetFilter() yfilter.YFilter { return utcOffset.YFilter }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) SetFilter(yf yfilter.YFilter) { utcOffset.YFilter = yf }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetGoName(yname string) string {
    if yname == "current-offset" { return "CurrentOffset" }
    if yname == "offset-valid" { return "OffsetValid" }
    return ""
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetSegmentPath() string {
    return "utc-offset"
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-offset"] = utcOffset.CurrentOffset
    leafs["offset-valid"] = utcOffset.OffsetValid
    return leafs
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetBundleName() string { return "cisco_ios_xr" }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetYangName() string { return "utc-offset" }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) SetParent(parent types.Entity) { utcOffset.parent = parent }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetParent() types.Entity { return utcOffset.parent }

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver
// Receiver
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetSegmentPath() string {
    return "receiver"
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = receiver.ClockId
    leafs["port-number"] = receiver.PortNumber
    return leafs
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetBundleName() string { return "cisco_ios_xr" }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetYangName() string { return "receiver" }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender
// Sender
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetFilter() yfilter.YFilter { return sender.YFilter }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) SetFilter(yf yfilter.YFilter) { sender.YFilter = yf }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetSegmentPath() string {
    return "sender"
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = sender.ClockId
    leafs["port-number"] = sender.PortNumber
    return leafs
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetBundleName() string { return "cisco_ios_xr" }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetYangName() string { return "sender" }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) SetParent(parent types.Entity) { sender.parent = parent }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetParent() types.Entity { return sender.parent }

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address
// The address of the clock
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetYangName() string { return "address" }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress
// Ethernet MAC address
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address
// IPv6 address
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetFilter() yfilter.YFilter { return announceGrant.YFilter }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) SetFilter(yf yfilter.YFilter) { announceGrant.YFilter = yf }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetSegmentPath() string {
    return "announce-grant"
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = announceGrant.LogGrantInterval
    leafs["grant-duration"] = announceGrant.GrantDuration
    return leafs
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetBundleName() string { return "cisco_ios_xr" }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetYangName() string { return "announce-grant" }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) SetParent(parent types.Entity) { announceGrant.parent = parent }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetParent() types.Entity { return announceGrant.parent }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant
// Unicast grant information for sync messages
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetFilter() yfilter.YFilter { return syncGrant.YFilter }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) SetFilter(yf yfilter.YFilter) { syncGrant.YFilter = yf }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetSegmentPath() string {
    return "sync-grant"
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = syncGrant.LogGrantInterval
    leafs["grant-duration"] = syncGrant.GrantDuration
    return leafs
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetBundleName() string { return "cisco_ios_xr" }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetYangName() string { return "sync-grant" }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) SetParent(parent types.Entity) { syncGrant.parent = parent }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetParent() types.Entity { return syncGrant.parent }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetFilter() yfilter.YFilter { return delayResponseGrant.YFilter }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) SetFilter(yf yfilter.YFilter) { delayResponseGrant.YFilter = yf }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetSegmentPath() string {
    return "delay-response-grant"
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = delayResponseGrant.LogGrantInterval
    leafs["grant-duration"] = delayResponseGrant.GrantDuration
    return leafs
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetBundleName() string { return "cisco_ios_xr" }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetYangName() string { return "delay-response-grant" }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) SetParent(parent types.Entity) { delayResponseGrant.parent = parent }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetParent() types.Entity { return delayResponseGrant.parent }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetParentYangName() string { return "foreign-clock" }

// Ptp_Nodes_Node_Summary
// Node summary operational data
type Ptp_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of interfaces in 'Init' port state. The type is interface{} with
    // range: 0..4294967295.
    PortStateInitCount interface{}

    // Number of interfaces in 'Listening' port state. The type is interface{}
    // with range: 0..4294967295.
    PortStateListeningCount interface{}

    // Number of interfaces in 'Passive' port state. The type is interface{} with
    // range: 0..4294967295.
    PortStatePassiveCount interface{}

    // Number of interfaces in 'Pre-Master' port state. The type is interface{}
    // with range: 0..4294967295.
    PortStatePreMasterCount interface{}

    // Number of interfaces in 'Master' port state. The type is interface{} with
    // range: 0..4294967295.
    PortStateMasterCount interface{}

    // Number of interfaces in 'Slave' port state. The type is interface{} with
    // range: 0..4294967295.
    PortStateSlaveCount interface{}

    // Number of interfaces in 'Uncalibrated port state. The type is interface{}
    // with range: 0..4294967295.
    PortStateUncalibratedCount interface{}

    // Number of interfaces in 'Faulty' port state. The type is interface{} with
    // range: 0..4294967295.
    PortStateFaultyCount interface{}

    // Total number of interfaces. The type is interface{} with range:
    // 0..4294967295.
    TotalInterfaces interface{}

    // Total number of interfaces with a valid port number. The type is
    // interface{} with range: 0..4294967295.
    TotalInterfacesValidPortNum interface{}
}

func (summary *Ptp_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ptp_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ptp_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "port-state-init-count" { return "PortStateInitCount" }
    if yname == "port-state-listening-count" { return "PortStateListeningCount" }
    if yname == "port-state-passive-count" { return "PortStatePassiveCount" }
    if yname == "port-state-pre-master-count" { return "PortStatePreMasterCount" }
    if yname == "port-state-master-count" { return "PortStateMasterCount" }
    if yname == "port-state-slave-count" { return "PortStateSlaveCount" }
    if yname == "port-state-uncalibrated-count" { return "PortStateUncalibratedCount" }
    if yname == "port-state-faulty-count" { return "PortStateFaultyCount" }
    if yname == "total-interfaces" { return "TotalInterfaces" }
    if yname == "total-interfaces-valid-port-num" { return "TotalInterfacesValidPortNum" }
    return ""
}

func (summary *Ptp_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ptp_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Ptp_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Ptp_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-state-init-count"] = summary.PortStateInitCount
    leafs["port-state-listening-count"] = summary.PortStateListeningCount
    leafs["port-state-passive-count"] = summary.PortStatePassiveCount
    leafs["port-state-pre-master-count"] = summary.PortStatePreMasterCount
    leafs["port-state-master-count"] = summary.PortStateMasterCount
    leafs["port-state-slave-count"] = summary.PortStateSlaveCount
    leafs["port-state-uncalibrated-count"] = summary.PortStateUncalibratedCount
    leafs["port-state-faulty-count"] = summary.PortStateFaultyCount
    leafs["total-interfaces"] = summary.TotalInterfaces
    leafs["total-interfaces-valid-port-num"] = summary.TotalInterfacesValidPortNum
    return leafs
}

func (summary *Ptp_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ptp_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *Ptp_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ptp_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ptp_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ptp_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ptp_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ptp_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// Ptp_Nodes_Node_NodeInterfaces
// Table for node interface operational data
type Ptp_Nodes_Node_NodeInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node interface operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface.
    NodeInterface []Ptp_Nodes_Node_NodeInterfaces_NodeInterface
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetFilter() yfilter.YFilter { return nodeInterfaces.YFilter }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) SetFilter(yf yfilter.YFilter) { nodeInterfaces.YFilter = yf }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetGoName(yname string) string {
    if yname == "node-interface" { return "NodeInterface" }
    return ""
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetSegmentPath() string {
    return "node-interfaces"
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-interface" {
        for _, c := range nodeInterfaces.NodeInterface {
            if nodeInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node_NodeInterfaces_NodeInterface{}
        nodeInterfaces.NodeInterface = append(nodeInterfaces.NodeInterface, child)
        return &nodeInterfaces.NodeInterface[len(nodeInterfaces.NodeInterface)-1]
    }
    return nil
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaces.NodeInterface {
        children[nodeInterfaces.NodeInterface[i].GetSegmentPath()] = &nodeInterfaces.NodeInterface[i]
    }
    return children
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetYangName() string { return "node-interfaces" }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) SetParent(parent types.Entity) { nodeInterfaces.parent = parent }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetParent() types.Entity { return nodeInterfaces.parent }

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetParentYangName() string { return "node" }

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface
// Node interface operational data
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Port state. The type is PtpBagPortState.
    PortState interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Line state. The type is interface{} with range: 0..4294967295.
    LineState interface{}

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Ipv6 address, if IPv6 encapsulation is being used. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // IPv4 address, if IPv4 encapsulation is being used. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Two step delay-request mechanism is being used. The type is bool.
    TwoStep interface{}

    // Communication model configured on the interface. The type is
    // PtpBagCommunicationModel.
    CommunicationModel interface{}

    // Log of the interface's sync interval. The type is interface{} with range:
    // -2147483648..2147483647.
    LogSyncInterval interface{}

    // Log of the interface's announce interval. The type is interface{} with
    // range: -2147483648..2147483647.
    LogAnnounceInterval interface{}

    // Announce timeout. The type is interface{} with range: 0..4294967295.
    AnnounceTimeout interface{}

    // Log of the interface's Minimum delay-request interval. The type is
    // interface{} with range: -2147483648..2147483647.
    LogMinDelayRequestInterval interface{}

    // The configured port state. The type is PtpBagRestrictPortState.
    ConfiguredPortState interface{}

    // The interface supports one-step operation. The type is bool.
    SupportsOneStep interface{}

    // The interface supports two-step operation. The type is bool.
    SupportsTwoStep interface{}

    // The interface supports ethernet transport. The type is bool.
    SupportsEthernet interface{}

    // The interface supports multicast. The type is bool.
    SupportsMulticast interface{}

    // The interface supports IPv6 transport. The type is bool.
    SupportsIpv6 interface{}

    // The interface supports operation in slave mode. The type is bool.
    SupportsSlave interface{}

    // The interface supports source ip configuration. The type is bool.
    SupportsSourceIp interface{}

    // The maximum rate of sync packets on the interface. The type is interface{}
    // with range: 0..255.
    MaxSyncRate interface{}

    // The class of service used on the interface for event messages. The type is
    // interface{} with range: 0..4294967295.
    EventCos interface{}

    // The class of service used on the interface for general messages. The type
    // is interface{} with range: 0..4294967295.
    GeneralCos interface{}

    // The DSCP class used on the interface for event messages. The type is
    // interface{} with range: 0..4294967295.
    EventDscp interface{}

    // The DSCP class used on the interface for general messages. The type is
    // interface{} with range: 0..4294967295.
    GeneralDscp interface{}

    // The number of unicast peers known by the interface. The type is interface{}
    // with range: 0..4294967295.
    UnicastPeers interface{}

    // Local priority, for the G.8275.1 PTP profile. The type is interface{} with
    // range: 0..255.
    LocalPriority interface{}

    // Signal fail status of the interface. The type is bool.
    SignalFail interface{}

    // MAC address, if Ethernet encapsulation is being used.
    MacAddress Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress

    // The interface's master table. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable.
    MasterTable []Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetFilter() yfilter.YFilter { return nodeInterface.YFilter }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) SetFilter(yf yfilter.YFilter) { nodeInterface.YFilter = yf }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "port-state" { return "PortState" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "line-state" { return "LineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-step" { return "TwoStep" }
    if yname == "communication-model" { return "CommunicationModel" }
    if yname == "log-sync-interval" { return "LogSyncInterval" }
    if yname == "log-announce-interval" { return "LogAnnounceInterval" }
    if yname == "announce-timeout" { return "AnnounceTimeout" }
    if yname == "log-min-delay-request-interval" { return "LogMinDelayRequestInterval" }
    if yname == "configured-port-state" { return "ConfiguredPortState" }
    if yname == "supports-one-step" { return "SupportsOneStep" }
    if yname == "supports-two-step" { return "SupportsTwoStep" }
    if yname == "supports-ethernet" { return "SupportsEthernet" }
    if yname == "supports-multicast" { return "SupportsMulticast" }
    if yname == "supports-ipv6" { return "SupportsIpv6" }
    if yname == "supports-slave" { return "SupportsSlave" }
    if yname == "supports-source-ip" { return "SupportsSourceIp" }
    if yname == "max-sync-rate" { return "MaxSyncRate" }
    if yname == "event-cos" { return "EventCos" }
    if yname == "general-cos" { return "GeneralCos" }
    if yname == "event-dscp" { return "EventDscp" }
    if yname == "general-dscp" { return "GeneralDscp" }
    if yname == "unicast-peers" { return "UnicastPeers" }
    if yname == "local-priority" { return "LocalPriority" }
    if yname == "signal-fail" { return "SignalFail" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "master-table" { return "MasterTable" }
    return ""
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetSegmentPath() string {
    return "node-interface" + "[interface-name='" + fmt.Sprintf("%v", nodeInterface.InterfaceName) + "']"
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &nodeInterface.MacAddress
    }
    if childYangName == "master-table" {
        for _, c := range nodeInterface.MasterTable {
            if nodeInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable{}
        nodeInterface.MasterTable = append(nodeInterface.MasterTable, child)
        return &nodeInterface.MasterTable[len(nodeInterface.MasterTable)-1]
    }
    return nil
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &nodeInterface.MacAddress
    for i := range nodeInterface.MasterTable {
        children[nodeInterface.MasterTable[i].GetSegmentPath()] = &nodeInterface.MasterTable[i]
    }
    return children
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = nodeInterface.InterfaceName
    leafs["port-state"] = nodeInterface.PortState
    leafs["port-number"] = nodeInterface.PortNumber
    leafs["line-state"] = nodeInterface.LineState
    leafs["encapsulation"] = nodeInterface.Encapsulation
    leafs["ipv6-address"] = nodeInterface.Ipv6Address
    leafs["ipv4-address"] = nodeInterface.Ipv4Address
    leafs["two-step"] = nodeInterface.TwoStep
    leafs["communication-model"] = nodeInterface.CommunicationModel
    leafs["log-sync-interval"] = nodeInterface.LogSyncInterval
    leafs["log-announce-interval"] = nodeInterface.LogAnnounceInterval
    leafs["announce-timeout"] = nodeInterface.AnnounceTimeout
    leafs["log-min-delay-request-interval"] = nodeInterface.LogMinDelayRequestInterval
    leafs["configured-port-state"] = nodeInterface.ConfiguredPortState
    leafs["supports-one-step"] = nodeInterface.SupportsOneStep
    leafs["supports-two-step"] = nodeInterface.SupportsTwoStep
    leafs["supports-ethernet"] = nodeInterface.SupportsEthernet
    leafs["supports-multicast"] = nodeInterface.SupportsMulticast
    leafs["supports-ipv6"] = nodeInterface.SupportsIpv6
    leafs["supports-slave"] = nodeInterface.SupportsSlave
    leafs["supports-source-ip"] = nodeInterface.SupportsSourceIp
    leafs["max-sync-rate"] = nodeInterface.MaxSyncRate
    leafs["event-cos"] = nodeInterface.EventCos
    leafs["general-cos"] = nodeInterface.GeneralCos
    leafs["event-dscp"] = nodeInterface.EventDscp
    leafs["general-dscp"] = nodeInterface.GeneralDscp
    leafs["unicast-peers"] = nodeInterface.UnicastPeers
    leafs["local-priority"] = nodeInterface.LocalPriority
    leafs["signal-fail"] = nodeInterface.SignalFail
    return leafs
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetYangName() string { return "node-interface" }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) SetParent(parent types.Entity) { nodeInterface.parent = parent }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetParent() types.Entity { return nodeInterface.parent }

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetParentYangName() string { return "node-interfaces" }

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress
// MAC address, if Ethernet encapsulation is being
// used
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetParentYangName() string { return "node-interface" }

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable
// The interface's master table
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The configured communication model of the master clock. The type is
    // PtpBagCommunicationModel.
    CommunicationModel interface{}

    // The priority of the master clock, if it is set. The type is interface{}
    // with range: 0..255.
    Priority interface{}

    // Whether the interface is receiving messages from this master. The type is
    // bool.
    Known interface{}

    // The master is qualified for best master clock selection. The type is bool.
    Qualified interface{}

    // Whether this is the grandmaster. The type is bool.
    IsGrandmaster interface{}

    // Announced messages are not being received from the master. The type is
    // interface{} with range: 0..255.
    PtsfLossAnnounce interface{}

    // Sync messages are not being received from the master. The type is
    // interface{} with range: 0..255.
    PtsfLossSync interface{}

    // Whether this master uses non-negotiated unicast. The type is bool.
    IsNonnegotiated interface{}

    // The address of the master clock.
    Address Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address
}

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetFilter() yfilter.YFilter { return masterTable.YFilter }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) SetFilter(yf yfilter.YFilter) { masterTable.YFilter = yf }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetGoName(yname string) string {
    if yname == "communication-model" { return "CommunicationModel" }
    if yname == "priority" { return "Priority" }
    if yname == "known" { return "Known" }
    if yname == "qualified" { return "Qualified" }
    if yname == "is-grandmaster" { return "IsGrandmaster" }
    if yname == "ptsf-loss-announce" { return "PtsfLossAnnounce" }
    if yname == "ptsf-loss-sync" { return "PtsfLossSync" }
    if yname == "is-nonnegotiated" { return "IsNonnegotiated" }
    if yname == "address" { return "Address" }
    return ""
}

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetSegmentPath() string {
    return "master-table"
}

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &masterTable.Address
    }
    return nil
}

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &masterTable.Address
    return children
}

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["communication-model"] = masterTable.CommunicationModel
    leafs["priority"] = masterTable.Priority
    leafs["known"] = masterTable.Known
    leafs["qualified"] = masterTable.Qualified
    leafs["is-grandmaster"] = masterTable.IsGrandmaster
    leafs["ptsf-loss-announce"] = masterTable.PtsfLossAnnounce
    leafs["ptsf-loss-sync"] = masterTable.PtsfLossSync
    leafs["is-nonnegotiated"] = masterTable.IsNonnegotiated
    return leafs
}

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetBundleName() string { return "cisco_ios_xr" }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetYangName() string { return "master-table" }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) SetParent(parent types.Entity) { masterTable.parent = parent }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetParent() types.Entity { return masterTable.parent }

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetParentYangName() string { return "node-interface" }

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address
// The address of the master clock
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetYangName() string { return "address" }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetParentYangName() string { return "master-table" }

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress
// Ethernet MAC address
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address
// IPv6 address
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers
// Table for node unicast peers operational data
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node interface unicast peers operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer.
    NodeInterfaceUnicastPeer []Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetFilter() yfilter.YFilter { return nodeInterfaceUnicastPeers.YFilter }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) SetFilter(yf yfilter.YFilter) { nodeInterfaceUnicastPeers.YFilter = yf }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetGoName(yname string) string {
    if yname == "node-interface-unicast-peer" { return "NodeInterfaceUnicastPeer" }
    return ""
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetSegmentPath() string {
    return "node-interface-unicast-peers"
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-interface-unicast-peer" {
        for _, c := range nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer {
            if nodeInterfaceUnicastPeers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer{}
        nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer = append(nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer, child)
        return &nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[len(nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer)-1]
    }
    return nil
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer {
        children[nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[i].GetSegmentPath()] = &nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[i]
    }
    return children
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetYangName() string { return "node-interface-unicast-peers" }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) SetParent(parent types.Entity) { nodeInterfaceUnicastPeers.parent = parent }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetParent() types.Entity { return nodeInterfaceUnicastPeers.parent }

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetParentYangName() string { return "node" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer
// Node interface unicast peers operational data
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Unicast Peers. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers.
    Peers []Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetFilter() yfilter.YFilter { return nodeInterfaceUnicastPeer.YFilter }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) SetFilter(yf yfilter.YFilter) { nodeInterfaceUnicastPeer.YFilter = yf }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "name" { return "Name" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "peers" { return "Peers" }
    return ""
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetSegmentPath() string {
    return "node-interface-unicast-peer" + "[interface-name='" + fmt.Sprintf("%v", nodeInterfaceUnicastPeer.InterfaceName) + "']"
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peers" {
        for _, c := range nodeInterfaceUnicastPeer.Peers {
            if nodeInterfaceUnicastPeer.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers{}
        nodeInterfaceUnicastPeer.Peers = append(nodeInterfaceUnicastPeer.Peers, child)
        return &nodeInterfaceUnicastPeer.Peers[len(nodeInterfaceUnicastPeer.Peers)-1]
    }
    return nil
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaceUnicastPeer.Peers {
        children[nodeInterfaceUnicastPeer.Peers[i].GetSegmentPath()] = &nodeInterfaceUnicastPeer.Peers[i]
    }
    return children
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = nodeInterfaceUnicastPeer.InterfaceName
    leafs["name"] = nodeInterfaceUnicastPeer.Name
    leafs["port-number"] = nodeInterfaceUnicastPeer.PortNumber
    return leafs
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetYangName() string { return "node-interface-unicast-peer" }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) SetParent(parent types.Entity) { nodeInterfaceUnicastPeer.parent = parent }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetParent() types.Entity { return nodeInterfaceUnicastPeer.parent }

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetParentYangName() string { return "node-interface-unicast-peers" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers
// Unicast Peers
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The address of the unicast peer.
    Address Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address

    // Unicast grant information for announce messages.
    AnnounceGrant Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant

    // Unicast grant information for sync messages.
    SyncGrant Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant

    // Unicast grant information for delay-response messages.
    DelayResponseGrant Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant
}

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "announce-grant" { return "AnnounceGrant" }
    if yname == "sync-grant" { return "SyncGrant" }
    if yname == "delay-response-grant" { return "DelayResponseGrant" }
    return ""
}

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &peers.Address
    }
    if childYangName == "announce-grant" {
        return &peers.AnnounceGrant
    }
    if childYangName == "sync-grant" {
        return &peers.SyncGrant
    }
    if childYangName == "delay-response-grant" {
        return &peers.DelayResponseGrant
    }
    return nil
}

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &peers.Address
    children["announce-grant"] = &peers.AnnounceGrant
    children["sync-grant"] = &peers.SyncGrant
    children["delay-response-grant"] = &peers.DelayResponseGrant
    return children
}

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetYangName() string { return "peers" }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetParent() types.Entity { return peers.parent }

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetParentYangName() string { return "node-interface-unicast-peer" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address
// The address of the unicast peer
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetYangName() string { return "address" }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetParentYangName() string { return "peers" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress
// Ethernet MAC address
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address
// IPv6 address
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetFilter() yfilter.YFilter { return announceGrant.YFilter }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) SetFilter(yf yfilter.YFilter) { announceGrant.YFilter = yf }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetSegmentPath() string {
    return "announce-grant"
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = announceGrant.LogGrantInterval
    leafs["grant-duration"] = announceGrant.GrantDuration
    return leafs
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetBundleName() string { return "cisco_ios_xr" }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetYangName() string { return "announce-grant" }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) SetParent(parent types.Entity) { announceGrant.parent = parent }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetParent() types.Entity { return announceGrant.parent }

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetParentYangName() string { return "peers" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant
// Unicast grant information for sync messages
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetFilter() yfilter.YFilter { return syncGrant.YFilter }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) SetFilter(yf yfilter.YFilter) { syncGrant.YFilter = yf }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetSegmentPath() string {
    return "sync-grant"
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = syncGrant.LogGrantInterval
    leafs["grant-duration"] = syncGrant.GrantDuration
    return leafs
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetBundleName() string { return "cisco_ios_xr" }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetYangName() string { return "sync-grant" }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) SetParent(parent types.Entity) { syncGrant.parent = parent }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetParent() types.Entity { return syncGrant.parent }

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetParentYangName() string { return "peers" }

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetFilter() yfilter.YFilter { return delayResponseGrant.YFilter }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) SetFilter(yf yfilter.YFilter) { delayResponseGrant.YFilter = yf }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetSegmentPath() string {
    return "delay-response-grant"
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = delayResponseGrant.LogGrantInterval
    leafs["grant-duration"] = delayResponseGrant.GrantDuration
    return leafs
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetBundleName() string { return "cisco_ios_xr" }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetYangName() string { return "delay-response-grant" }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) SetParent(parent types.Entity) { delayResponseGrant.parent = parent }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetParent() types.Entity { return delayResponseGrant.parent }

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetParentYangName() string { return "peers" }

// Ptp_Nodes_Node_PacketCounters
// Node packet counter operational data
type Ptp_Nodes_Node_PacketCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet counters.
    Counters Ptp_Nodes_Node_PacketCounters_Counters

    // Drop reasons.
    DropReasons Ptp_Nodes_Node_PacketCounters_DropReasons
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetFilter() yfilter.YFilter { return packetCounters.YFilter }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) SetFilter(yf yfilter.YFilter) { packetCounters.YFilter = yf }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetGoName(yname string) string {
    if yname == "counters" { return "Counters" }
    if yname == "drop-reasons" { return "DropReasons" }
    return ""
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetSegmentPath() string {
    return "packet-counters"
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &packetCounters.Counters
    }
    if childYangName == "drop-reasons" {
        return &packetCounters.DropReasons
    }
    return nil
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &packetCounters.Counters
    children["drop-reasons"] = &packetCounters.DropReasons
    return children
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetBundleName() string { return "cisco_ios_xr" }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetYangName() string { return "packet-counters" }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) SetParent(parent types.Entity) { packetCounters.parent = parent }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetParent() types.Entity { return packetCounters.parent }

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetParentYangName() string { return "node" }

// Ptp_Nodes_Node_PacketCounters_Counters
// Packet counters
type Ptp_Nodes_Node_PacketCounters_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of announce packets sent. The type is interface{} with range:
    // 0..4294967295.
    AnnounceSent interface{}

    // Number of announce packets received. The type is interface{} with range:
    // 0..4294967295.
    AnnounceReceived interface{}

    // Number of announce packets dropped. The type is interface{} with range:
    // 0..4294967295.
    AnnounceDropped interface{}

    // Number of sync packets sent. The type is interface{} with range:
    // 0..4294967295.
    SyncSent interface{}

    // Number of sync packets received. The type is interface{} with range:
    // 0..4294967295.
    SyncReceived interface{}

    // Number of sync packetsdropped. The type is interface{} with range:
    // 0..4294967295.
    SyncDropped interface{}

    // Number of follow-up packets sent. The type is interface{} with range:
    // 0..4294967295.
    FollowUpSent interface{}

    // Number of follow-up packets received. The type is interface{} with range:
    // 0..4294967295.
    FollowUpReceived interface{}

    // Number of follow-up packets dropped. The type is interface{} with range:
    // 0..4294967295.
    FollowUpDropped interface{}

    // Number of delay-request packets sent. The type is interface{} with range:
    // 0..4294967295.
    DelayRequestSent interface{}

    // Number of delay-request packets received. The type is interface{} with
    // range: 0..4294967295.
    DelayRequestReceived interface{}

    // Number of delay-request packets dropped. The type is interface{} with
    // range: 0..4294967295.
    DelayRequestDropped interface{}

    // Number of delay-response packets sent. The type is interface{} with range:
    // 0..4294967295.
    DelayResponseSent interface{}

    // Number of delay-response packets received. The type is interface{} with
    // range: 0..4294967295.
    DelayResponseReceived interface{}

    // Number of delay-response packets dropped. The type is interface{} with
    // range: 0..4294967295.
    DelayResponseDropped interface{}

    // Number of peer-delay-request packets sent. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestSent interface{}

    // Number of peer-delay-request packets received. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestReceived interface{}

    // Number of peer-delay-request packets dropped. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestDropped interface{}

    // Number of peer-delay-response packets sent. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayResponseSent interface{}

    // Number of peer-delay-response packets received. The type is interface{}
    // with range: 0..4294967295.
    PeerDelayResponseReceived interface{}

    // Number of peer-delay-response packets dropped. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayResponseDropped interface{}

    // Number of peer-delay-response follow-up packets sent. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpSent interface{}

    // Number of peer-delay-response follow-up packets received. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpReceived interface{}

    // Number of peer-delay-response follow-up packets dropped. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpDropped interface{}

    // Number of signaling packets sent. The type is interface{} with range:
    // 0..4294967295.
    SignalingSent interface{}

    // Number of signaling packets received. The type is interface{} with range:
    // 0..4294967295.
    SignalingReceived interface{}

    // Number of signaling packets dropped. The type is interface{} with range:
    // 0..4294967295.
    SignalingDropped interface{}

    // Number of management messages sent. The type is interface{} with range:
    // 0..4294967295.
    ManagementSent interface{}

    // Number of management messages received. The type is interface{} with range:
    // 0..4294967295.
    ManagementReceived interface{}

    // Number of management messages dropped. The type is interface{} with range:
    // 0..4294967295.
    ManagementDropped interface{}

    // Number of other packets sent. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsSent interface{}

    // Number of other packets received. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsReceived interface{}

    // Number of other packets dropped. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsDropped interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsSent interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsReceived interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsDropped interface{}
}

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetGoName(yname string) string {
    if yname == "announce-sent" { return "AnnounceSent" }
    if yname == "announce-received" { return "AnnounceReceived" }
    if yname == "announce-dropped" { return "AnnounceDropped" }
    if yname == "sync-sent" { return "SyncSent" }
    if yname == "sync-received" { return "SyncReceived" }
    if yname == "sync-dropped" { return "SyncDropped" }
    if yname == "follow-up-sent" { return "FollowUpSent" }
    if yname == "follow-up-received" { return "FollowUpReceived" }
    if yname == "follow-up-dropped" { return "FollowUpDropped" }
    if yname == "delay-request-sent" { return "DelayRequestSent" }
    if yname == "delay-request-received" { return "DelayRequestReceived" }
    if yname == "delay-request-dropped" { return "DelayRequestDropped" }
    if yname == "delay-response-sent" { return "DelayResponseSent" }
    if yname == "delay-response-received" { return "DelayResponseReceived" }
    if yname == "delay-response-dropped" { return "DelayResponseDropped" }
    if yname == "peer-delay-request-sent" { return "PeerDelayRequestSent" }
    if yname == "peer-delay-request-received" { return "PeerDelayRequestReceived" }
    if yname == "peer-delay-request-dropped" { return "PeerDelayRequestDropped" }
    if yname == "peer-delay-response-sent" { return "PeerDelayResponseSent" }
    if yname == "peer-delay-response-received" { return "PeerDelayResponseReceived" }
    if yname == "peer-delay-response-dropped" { return "PeerDelayResponseDropped" }
    if yname == "peer-delay-response-follow-up-sent" { return "PeerDelayResponseFollowUpSent" }
    if yname == "peer-delay-response-follow-up-received" { return "PeerDelayResponseFollowUpReceived" }
    if yname == "peer-delay-response-follow-up-dropped" { return "PeerDelayResponseFollowUpDropped" }
    if yname == "signaling-sent" { return "SignalingSent" }
    if yname == "signaling-received" { return "SignalingReceived" }
    if yname == "signaling-dropped" { return "SignalingDropped" }
    if yname == "management-sent" { return "ManagementSent" }
    if yname == "management-received" { return "ManagementReceived" }
    if yname == "management-dropped" { return "ManagementDropped" }
    if yname == "other-packets-sent" { return "OtherPacketsSent" }
    if yname == "other-packets-received" { return "OtherPacketsReceived" }
    if yname == "other-packets-dropped" { return "OtherPacketsDropped" }
    if yname == "total-packets-sent" { return "TotalPacketsSent" }
    if yname == "total-packets-received" { return "TotalPacketsReceived" }
    if yname == "total-packets-dropped" { return "TotalPacketsDropped" }
    return ""
}

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["announce-sent"] = counters.AnnounceSent
    leafs["announce-received"] = counters.AnnounceReceived
    leafs["announce-dropped"] = counters.AnnounceDropped
    leafs["sync-sent"] = counters.SyncSent
    leafs["sync-received"] = counters.SyncReceived
    leafs["sync-dropped"] = counters.SyncDropped
    leafs["follow-up-sent"] = counters.FollowUpSent
    leafs["follow-up-received"] = counters.FollowUpReceived
    leafs["follow-up-dropped"] = counters.FollowUpDropped
    leafs["delay-request-sent"] = counters.DelayRequestSent
    leafs["delay-request-received"] = counters.DelayRequestReceived
    leafs["delay-request-dropped"] = counters.DelayRequestDropped
    leafs["delay-response-sent"] = counters.DelayResponseSent
    leafs["delay-response-received"] = counters.DelayResponseReceived
    leafs["delay-response-dropped"] = counters.DelayResponseDropped
    leafs["peer-delay-request-sent"] = counters.PeerDelayRequestSent
    leafs["peer-delay-request-received"] = counters.PeerDelayRequestReceived
    leafs["peer-delay-request-dropped"] = counters.PeerDelayRequestDropped
    leafs["peer-delay-response-sent"] = counters.PeerDelayResponseSent
    leafs["peer-delay-response-received"] = counters.PeerDelayResponseReceived
    leafs["peer-delay-response-dropped"] = counters.PeerDelayResponseDropped
    leafs["peer-delay-response-follow-up-sent"] = counters.PeerDelayResponseFollowUpSent
    leafs["peer-delay-response-follow-up-received"] = counters.PeerDelayResponseFollowUpReceived
    leafs["peer-delay-response-follow-up-dropped"] = counters.PeerDelayResponseFollowUpDropped
    leafs["signaling-sent"] = counters.SignalingSent
    leafs["signaling-received"] = counters.SignalingReceived
    leafs["signaling-dropped"] = counters.SignalingDropped
    leafs["management-sent"] = counters.ManagementSent
    leafs["management-received"] = counters.ManagementReceived
    leafs["management-dropped"] = counters.ManagementDropped
    leafs["other-packets-sent"] = counters.OtherPacketsSent
    leafs["other-packets-received"] = counters.OtherPacketsReceived
    leafs["other-packets-dropped"] = counters.OtherPacketsDropped
    leafs["total-packets-sent"] = counters.TotalPacketsSent
    leafs["total-packets-received"] = counters.TotalPacketsReceived
    leafs["total-packets-dropped"] = counters.TotalPacketsDropped
    return leafs
}

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetYangName() string { return "counters" }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetParentYangName() string { return "packet-counters" }

// Ptp_Nodes_Node_PacketCounters_DropReasons
// Drop reasons
type Ptp_Nodes_Node_PacketCounters_DropReasons struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not ready for packets. The type is interface{} with range: 0..4294967295.
    NotReady interface{}

    // Wrong domain number. The type is interface{} with range: 0..4294967295.
    WrongDomain interface{}

    // Packet too short. The type is interface{} with range: 0..4294967295.
    TooShort interface{}

    // Local packet received, same port number. The type is interface{} with
    // range: 0..4294967295.
    LoopedSamePort interface{}

    // Local packet received, higher port number. The type is interface{} with
    // range: 0..4294967295.
    LoopedHigherPort interface{}

    // Local packet received, lower port number. The type is interface{} with
    // range: 0..4294967295.
    LoopedLowerPort interface{}

    // No timestamp received with packet. The type is interface{} with range:
    // 0..4294967295.
    NoTimestamp interface{}

    // Zero timestamp received with packet. The type is interface{} with range:
    // 0..4294967295.
    ZeroTimestamp interface{}

    // Invalid TLVs received in packet. The type is interface{} with range:
    // 0..4294967295.
    InvalidTlVs interface{}

    // Packet not for us. The type is interface{} with range: 0..4294967295.
    NotForUs interface{}

    // Not listening for packets on this interface. The type is interface{} with
    // range: 0..4294967295.
    NotListening interface{}

    // Packet from incorrect master. The type is interface{} with range:
    // 0..4294967295.
    WrongMaster interface{}

    // Packet from unknown master. The type is interface{} with range:
    // 0..4294967295.
    UnknownMaster interface{}

    // Packet only handled in Master state. The type is interface{} with range:
    // 0..4294967295.
    NotMaster interface{}

    // Packet only handled in Slave state. The type is interface{} with range:
    // 0..4294967295.
    NotSlave interface{}

    // Packet from peer not granted unicast. The type is interface{} with range:
    // 0..4294967295.
    NotGranted interface{}

    // Packet received too late. The type is interface{} with range:
    // 0..4294967295.
    TooSlow interface{}

    // Invalid packet or packet metadata. The type is interface{} with range:
    // 0..4294967295.
    InvalidPacket interface{}

    // Unexpected sequence ID. The type is interface{} with range: 0..4294967295.
    WrongSequenceId interface{}

    // No offload session. The type is interface{} with range: 0..4294967295.
    NoOffloadSession interface{}

    // PTP packet type not supported. The type is interface{} with range:
    // 0..4294967295.
    NotSupported interface{}

    // Clock class below minimum. The type is interface{} with range:
    // 0..4294967295.
    MinClockClass interface{}

    // Packet not compatible with G.8275.1 profile. The type is interface{} with
    // range: 0..4294967295.
    G82751Incompatible interface{}

    // Packet not compatible with G.8275.2 profile. The type is interface{} with
    // range: 0..4294967295.
    G82752Incompatible interface{}
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetFilter() yfilter.YFilter { return dropReasons.YFilter }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) SetFilter(yf yfilter.YFilter) { dropReasons.YFilter = yf }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetGoName(yname string) string {
    if yname == "not-ready" { return "NotReady" }
    if yname == "wrong-domain" { return "WrongDomain" }
    if yname == "too-short" { return "TooShort" }
    if yname == "looped-same-port" { return "LoopedSamePort" }
    if yname == "looped-higher-port" { return "LoopedHigherPort" }
    if yname == "looped-lower-port" { return "LoopedLowerPort" }
    if yname == "no-timestamp" { return "NoTimestamp" }
    if yname == "zero-timestamp" { return "ZeroTimestamp" }
    if yname == "invalid-tl-vs" { return "InvalidTlVs" }
    if yname == "not-for-us" { return "NotForUs" }
    if yname == "not-listening" { return "NotListening" }
    if yname == "wrong-master" { return "WrongMaster" }
    if yname == "unknown-master" { return "UnknownMaster" }
    if yname == "not-master" { return "NotMaster" }
    if yname == "not-slave" { return "NotSlave" }
    if yname == "not-granted" { return "NotGranted" }
    if yname == "too-slow" { return "TooSlow" }
    if yname == "invalid-packet" { return "InvalidPacket" }
    if yname == "wrong-sequence-id" { return "WrongSequenceId" }
    if yname == "no-offload-session" { return "NoOffloadSession" }
    if yname == "not-supported" { return "NotSupported" }
    if yname == "min-clock-class" { return "MinClockClass" }
    if yname == "g8275-1-incompatible" { return "G82751Incompatible" }
    if yname == "g8275-2-incompatible" { return "G82752Incompatible" }
    return ""
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetSegmentPath() string {
    return "drop-reasons"
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["not-ready"] = dropReasons.NotReady
    leafs["wrong-domain"] = dropReasons.WrongDomain
    leafs["too-short"] = dropReasons.TooShort
    leafs["looped-same-port"] = dropReasons.LoopedSamePort
    leafs["looped-higher-port"] = dropReasons.LoopedHigherPort
    leafs["looped-lower-port"] = dropReasons.LoopedLowerPort
    leafs["no-timestamp"] = dropReasons.NoTimestamp
    leafs["zero-timestamp"] = dropReasons.ZeroTimestamp
    leafs["invalid-tl-vs"] = dropReasons.InvalidTlVs
    leafs["not-for-us"] = dropReasons.NotForUs
    leafs["not-listening"] = dropReasons.NotListening
    leafs["wrong-master"] = dropReasons.WrongMaster
    leafs["unknown-master"] = dropReasons.UnknownMaster
    leafs["not-master"] = dropReasons.NotMaster
    leafs["not-slave"] = dropReasons.NotSlave
    leafs["not-granted"] = dropReasons.NotGranted
    leafs["too-slow"] = dropReasons.TooSlow
    leafs["invalid-packet"] = dropReasons.InvalidPacket
    leafs["wrong-sequence-id"] = dropReasons.WrongSequenceId
    leafs["no-offload-session"] = dropReasons.NoOffloadSession
    leafs["not-supported"] = dropReasons.NotSupported
    leafs["min-clock-class"] = dropReasons.MinClockClass
    leafs["g8275-1-incompatible"] = dropReasons.G82751Incompatible
    leafs["g8275-2-incompatible"] = dropReasons.G82752Incompatible
    return leafs
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetBundleName() string { return "cisco_ios_xr" }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetYangName() string { return "drop-reasons" }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) SetParent(parent types.Entity) { dropReasons.parent = parent }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetParent() types.Entity { return dropReasons.parent }

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetParentYangName() string { return "packet-counters" }

// Ptp_InterfaceConfigurationErrors
// Table for interface configuration error
// operational data
type Ptp_InterfaceConfigurationErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface configuration error operational data. The type is slice of
    // Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError.
    InterfaceConfigurationError []Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetFilter() yfilter.YFilter { return interfaceConfigurationErrors.YFilter }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) SetFilter(yf yfilter.YFilter) { interfaceConfigurationErrors.YFilter = yf }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetGoName(yname string) string {
    if yname == "interface-configuration-error" { return "InterfaceConfigurationError" }
    return ""
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetSegmentPath() string {
    return "interface-configuration-errors"
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-configuration-error" {
        for _, c := range interfaceConfigurationErrors.InterfaceConfigurationError {
            if interfaceConfigurationErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError{}
        interfaceConfigurationErrors.InterfaceConfigurationError = append(interfaceConfigurationErrors.InterfaceConfigurationError, child)
        return &interfaceConfigurationErrors.InterfaceConfigurationError[len(interfaceConfigurationErrors.InterfaceConfigurationError)-1]
    }
    return nil
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceConfigurationErrors.InterfaceConfigurationError {
        children[interfaceConfigurationErrors.InterfaceConfigurationError[i].GetSegmentPath()] = &interfaceConfigurationErrors.InterfaceConfigurationError[i]
    }
    return children
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetYangName() string { return "interface-configuration-errors" }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) SetParent(parent types.Entity) { interfaceConfigurationErrors.parent = parent }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetParent() types.Entity { return interfaceConfigurationErrors.parent }

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetParentYangName() string { return "ptp" }

// Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError
// Interface configuration error operational data
type Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Configuration profile name, if a profile is selected. The type is string.
    ConfigurationProfileName interface{}

    // The clock profile. The type is PtpBagProfile.
    ClockProfile interface{}

    // The telecom clock type. The type is PtpBagTelecomClock.
    TelecomClockType interface{}

    // Restriction on the port state. The type is PtpBagRestrictPortState.
    RestrictPortState interface{}

    // Configuration Errors.
    ConfigurationErrors Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetFilter() yfilter.YFilter { return interfaceConfigurationError.YFilter }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) SetFilter(yf yfilter.YFilter) { interfaceConfigurationError.YFilter = yf }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "configuration-profile-name" { return "ConfigurationProfileName" }
    if yname == "clock-profile" { return "ClockProfile" }
    if yname == "telecom-clock-type" { return "TelecomClockType" }
    if yname == "restrict-port-state" { return "RestrictPortState" }
    if yname == "configuration-errors" { return "ConfigurationErrors" }
    return ""
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetSegmentPath() string {
    return "interface-configuration-error" + "[interface-name='" + fmt.Sprintf("%v", interfaceConfigurationError.InterfaceName) + "']"
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configuration-errors" {
        return &interfaceConfigurationError.ConfigurationErrors
    }
    return nil
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configuration-errors"] = &interfaceConfigurationError.ConfigurationErrors
    return children
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceConfigurationError.InterfaceName
    leafs["configuration-profile-name"] = interfaceConfigurationError.ConfigurationProfileName
    leafs["clock-profile"] = interfaceConfigurationError.ClockProfile
    leafs["telecom-clock-type"] = interfaceConfigurationError.TelecomClockType
    leafs["restrict-port-state"] = interfaceConfigurationError.RestrictPortState
    return leafs
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetYangName() string { return "interface-configuration-error" }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) SetParent(parent types.Entity) { interfaceConfigurationError.parent = parent }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetParent() types.Entity { return interfaceConfigurationError.parent }

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetParentYangName() string { return "interface-configuration-errors" }

// Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors
// Configuration Errors
type Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PTP enabled on interface but not globally. The type is bool.
    GlobalPtp interface{}

    // Ethernet transport configured but not supported. The type is bool.
    EthernetTransport interface{}

    // One step clock operation configured but not supported. The type is bool.
    OneStep interface{}

    // Slave-operation configured but not supported. The type is bool.
    Slave interface{}

    // IPv6 transport configured but not supported. The type is bool.
    Ipv6 interface{}

    // Multicast configured but not supported. The type is bool.
    Multicast interface{}

    // Profile is referenced but not globally configured. The type is bool.
    ProfileNotGlobal interface{}

    // Local priority configuration is not compatible with profile. The type is
    // bool.
    LocalPriority interface{}

    // Ethernet transport is not compatible with profile. The type is bool.
    ProfileEthernet interface{}

    // IPv6 transport is not compatible with profile. The type is bool.
    ProfileIpv4 interface{}

    // IPv6 transport is not compatible with profile. The type is bool.
    ProfileIpv6 interface{}

    // Unicast is not compatible with profile. The type is bool.
    ProfileUnicast interface{}

    // Multicast is not compatible with profile. The type is bool.
    ProfileMulticast interface{}

    // Mixed-mode multicast is not compatible with profile. The type is bool.
    ProfileMixed interface{}

    // Unicast master is not compatible with profile. The type is bool.
    ProfileMasterUnicast interface{}

    // Multicast master is not compatible with profile. The type is bool.
    ProfileMasterMulticast interface{}

    // Mixed-mode multicast master is not compatible with profile. The type is
    // bool.
    ProfileMasterMixed interface{}

    // Ethernet multicast target-address is configured, but transport is IPv4. The
    // type is bool.
    TargetAddressIpv4 interface{}

    // Ethernet multicast target-address is configured, but transport is IPv6. The
    // type is bool.
    TargetAddressIpv6 interface{}

    // Port state restriction is not compatible with telecom clock type. The type
    // is bool.
    ProfilePortState interface{}

    // Announce interval is not compatible with profile. The type is bool.
    ProfileAnnounceInterval interface{}

    // Sync interval is not compatible with profile. The type is bool.
    ProfileSyncInterval interface{}

    // Delay request interval is not compatible with profile. The type is bool.
    ProfileDelayReqInterval interface{}

    // Sync timeout configuration is not compatible with profile. The type is
    // bool.
    ProfileSyncTimeout interface{}

    // Delay response timeout configuration is not compatible with profile. The
    // type is bool.
    ProfileDelayRespTimeout interface{}

    // Reducing invalid unicast grants is not compatible with configured profile.
    // The type is bool.
    InvalidGrantReduction interface{}
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetFilter() yfilter.YFilter { return configurationErrors.YFilter }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) SetFilter(yf yfilter.YFilter) { configurationErrors.YFilter = yf }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetGoName(yname string) string {
    if yname == "global-ptp" { return "GlobalPtp" }
    if yname == "ethernet-transport" { return "EthernetTransport" }
    if yname == "one-step" { return "OneStep" }
    if yname == "slave" { return "Slave" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "multicast" { return "Multicast" }
    if yname == "profile-not-global" { return "ProfileNotGlobal" }
    if yname == "local-priority" { return "LocalPriority" }
    if yname == "profile-ethernet" { return "ProfileEthernet" }
    if yname == "profile-ipv4" { return "ProfileIpv4" }
    if yname == "profile-ipv6" { return "ProfileIpv6" }
    if yname == "profile-unicast" { return "ProfileUnicast" }
    if yname == "profile-multicast" { return "ProfileMulticast" }
    if yname == "profile-mixed" { return "ProfileMixed" }
    if yname == "profile-master-unicast" { return "ProfileMasterUnicast" }
    if yname == "profile-master-multicast" { return "ProfileMasterMulticast" }
    if yname == "profile-master-mixed" { return "ProfileMasterMixed" }
    if yname == "target-address-ipv4" { return "TargetAddressIpv4" }
    if yname == "target-address-ipv6" { return "TargetAddressIpv6" }
    if yname == "profile-port-state" { return "ProfilePortState" }
    if yname == "profile-announce-interval" { return "ProfileAnnounceInterval" }
    if yname == "profile-sync-interval" { return "ProfileSyncInterval" }
    if yname == "profile-delay-req-interval" { return "ProfileDelayReqInterval" }
    if yname == "profile-sync-timeout" { return "ProfileSyncTimeout" }
    if yname == "profile-delay-resp-timeout" { return "ProfileDelayRespTimeout" }
    if yname == "invalid-grant-reduction" { return "InvalidGrantReduction" }
    return ""
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetSegmentPath() string {
    return "configuration-errors"
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-ptp"] = configurationErrors.GlobalPtp
    leafs["ethernet-transport"] = configurationErrors.EthernetTransport
    leafs["one-step"] = configurationErrors.OneStep
    leafs["slave"] = configurationErrors.Slave
    leafs["ipv6"] = configurationErrors.Ipv6
    leafs["multicast"] = configurationErrors.Multicast
    leafs["profile-not-global"] = configurationErrors.ProfileNotGlobal
    leafs["local-priority"] = configurationErrors.LocalPriority
    leafs["profile-ethernet"] = configurationErrors.ProfileEthernet
    leafs["profile-ipv4"] = configurationErrors.ProfileIpv4
    leafs["profile-ipv6"] = configurationErrors.ProfileIpv6
    leafs["profile-unicast"] = configurationErrors.ProfileUnicast
    leafs["profile-multicast"] = configurationErrors.ProfileMulticast
    leafs["profile-mixed"] = configurationErrors.ProfileMixed
    leafs["profile-master-unicast"] = configurationErrors.ProfileMasterUnicast
    leafs["profile-master-multicast"] = configurationErrors.ProfileMasterMulticast
    leafs["profile-master-mixed"] = configurationErrors.ProfileMasterMixed
    leafs["target-address-ipv4"] = configurationErrors.TargetAddressIpv4
    leafs["target-address-ipv6"] = configurationErrors.TargetAddressIpv6
    leafs["profile-port-state"] = configurationErrors.ProfilePortState
    leafs["profile-announce-interval"] = configurationErrors.ProfileAnnounceInterval
    leafs["profile-sync-interval"] = configurationErrors.ProfileSyncInterval
    leafs["profile-delay-req-interval"] = configurationErrors.ProfileDelayReqInterval
    leafs["profile-sync-timeout"] = configurationErrors.ProfileSyncTimeout
    leafs["profile-delay-resp-timeout"] = configurationErrors.ProfileDelayRespTimeout
    leafs["invalid-grant-reduction"] = configurationErrors.InvalidGrantReduction
    return leafs
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetBundleName() string { return "cisco_ios_xr" }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetYangName() string { return "configuration-errors" }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) SetParent(parent types.Entity) { configurationErrors.parent = parent }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetParent() types.Entity { return configurationErrors.parent }

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetParentYangName() string { return "interface-configuration-error" }

// Ptp_InterfaceForeignMasters
// Table for interface foreign master clock
// operational data
type Ptp_InterfaceForeignMasters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface foreign master clock operational data. The type is slice of
    // Ptp_InterfaceForeignMasters_InterfaceForeignMaster.
    InterfaceForeignMaster []Ptp_InterfaceForeignMasters_InterfaceForeignMaster
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetFilter() yfilter.YFilter { return interfaceForeignMasters.YFilter }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) SetFilter(yf yfilter.YFilter) { interfaceForeignMasters.YFilter = yf }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetGoName(yname string) string {
    if yname == "interface-foreign-master" { return "InterfaceForeignMaster" }
    return ""
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetSegmentPath() string {
    return "interface-foreign-masters"
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-foreign-master" {
        for _, c := range interfaceForeignMasters.InterfaceForeignMaster {
            if interfaceForeignMasters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfaceForeignMasters_InterfaceForeignMaster{}
        interfaceForeignMasters.InterfaceForeignMaster = append(interfaceForeignMasters.InterfaceForeignMaster, child)
        return &interfaceForeignMasters.InterfaceForeignMaster[len(interfaceForeignMasters.InterfaceForeignMaster)-1]
    }
    return nil
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceForeignMasters.InterfaceForeignMaster {
        children[interfaceForeignMasters.InterfaceForeignMaster[i].GetSegmentPath()] = &interfaceForeignMasters.InterfaceForeignMaster[i]
    }
    return children
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetYangName() string { return "interface-foreign-masters" }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) SetParent(parent types.Entity) { interfaceForeignMasters.parent = parent }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetParent() types.Entity { return interfaceForeignMasters.parent }

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetParentYangName() string { return "ptp" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster
// Interface foreign master clock operational data
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Foreign clocks received on this interface. The type is slice of
    // Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock.
    ForeignClock []Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetFilter() yfilter.YFilter { return interfaceForeignMaster.YFilter }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) SetFilter(yf yfilter.YFilter) { interfaceForeignMaster.YFilter = yf }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "foreign-clock" { return "ForeignClock" }
    return ""
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetSegmentPath() string {
    return "interface-foreign-master" + "[interface-name='" + fmt.Sprintf("%v", interfaceForeignMaster.InterfaceName) + "']"
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "foreign-clock" {
        for _, c := range interfaceForeignMaster.ForeignClock {
            if interfaceForeignMaster.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock{}
        interfaceForeignMaster.ForeignClock = append(interfaceForeignMaster.ForeignClock, child)
        return &interfaceForeignMaster.ForeignClock[len(interfaceForeignMaster.ForeignClock)-1]
    }
    return nil
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceForeignMaster.ForeignClock {
        children[interfaceForeignMaster.ForeignClock[i].GetSegmentPath()] = &interfaceForeignMaster.ForeignClock[i]
    }
    return children
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceForeignMaster.InterfaceName
    leafs["port-number"] = interfaceForeignMaster.PortNumber
    return leafs
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetYangName() string { return "interface-foreign-master" }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) SetParent(parent types.Entity) { interfaceForeignMaster.parent = parent }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetParent() types.Entity { return interfaceForeignMaster.parent }

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetParentYangName() string { return "interface-foreign-masters" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock
// Foreign clocks received on this interface
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The clock is qualified for best master clock selection. The type is bool.
    IsQualified interface{}

    // This clock is the currently selected grand master clock. The type is bool.
    IsGrandmaster interface{}

    // The communication model configured on this clock. The type is
    // PtpBagCommunicationModel.
    CommunicationModel interface{}

    // This clock is known by this router. The type is bool.
    IsKnown interface{}

    // How long the clock has been known by this router for, in seconds. The type
    // is interface{} with range: 0..4294967295. Units are second.
    TimeKnownFor interface{}

    // The PTP domain that the foreign clock is in. The type is interface{} with
    // range: 0..255.
    ForeignDomain interface{}

    // Priority configured for the clock, if any. The type is interface{} with
    // range: 0..255.
    ConfiguredPriority interface{}

    // Clock class configured for the clock, if any. The type is interface{} with
    // range: 0..255.
    ConfiguredClockClass interface{}

    // Delay asymmetry configured for the clock, if any. The type is interface{}
    // with range: -2147483648..2147483647.
    DelayAsymmetry interface{}

    // Announced messages are not being received from the master. The type is
    // bool.
    PtsfLossAnnounce interface{}

    // Sync messages are not being received from the master. The type is bool.
    PtsfLossSync interface{}

    // Foreign clock information.
    ForeignClock Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock

    // The address of the clock.
    Address Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address

    // Unicast grant information for announce messages.
    AnnounceGrant Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant

    // Unicast grant information for sync messages.
    SyncGrant Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant

    // Unicast grant information for delay-response messages.
    DelayResponseGrant Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetFilter() yfilter.YFilter { return foreignClock.YFilter }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) SetFilter(yf yfilter.YFilter) { foreignClock.YFilter = yf }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetGoName(yname string) string {
    if yname == "is-qualified" { return "IsQualified" }
    if yname == "is-grandmaster" { return "IsGrandmaster" }
    if yname == "communication-model" { return "CommunicationModel" }
    if yname == "is-known" { return "IsKnown" }
    if yname == "time-known-for" { return "TimeKnownFor" }
    if yname == "foreign-domain" { return "ForeignDomain" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "delay-asymmetry" { return "DelayAsymmetry" }
    if yname == "ptsf-loss-announce" { return "PtsfLossAnnounce" }
    if yname == "ptsf-loss-sync" { return "PtsfLossSync" }
    if yname == "foreign-clock" { return "ForeignClock" }
    if yname == "address" { return "Address" }
    if yname == "announce-grant" { return "AnnounceGrant" }
    if yname == "sync-grant" { return "SyncGrant" }
    if yname == "delay-response-grant" { return "DelayResponseGrant" }
    return ""
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetSegmentPath() string {
    return "foreign-clock"
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "foreign-clock" {
        return &foreignClock.ForeignClock
    }
    if childYangName == "address" {
        return &foreignClock.Address
    }
    if childYangName == "announce-grant" {
        return &foreignClock.AnnounceGrant
    }
    if childYangName == "sync-grant" {
        return &foreignClock.SyncGrant
    }
    if childYangName == "delay-response-grant" {
        return &foreignClock.DelayResponseGrant
    }
    return nil
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["foreign-clock"] = &foreignClock.ForeignClock
    children["address"] = &foreignClock.Address
    children["announce-grant"] = &foreignClock.AnnounceGrant
    children["sync-grant"] = &foreignClock.SyncGrant
    children["delay-response-grant"] = &foreignClock.DelayResponseGrant
    return children
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-qualified"] = foreignClock.IsQualified
    leafs["is-grandmaster"] = foreignClock.IsGrandmaster
    leafs["communication-model"] = foreignClock.CommunicationModel
    leafs["is-known"] = foreignClock.IsKnown
    leafs["time-known-for"] = foreignClock.TimeKnownFor
    leafs["foreign-domain"] = foreignClock.ForeignDomain
    leafs["configured-priority"] = foreignClock.ConfiguredPriority
    leafs["configured-clock-class"] = foreignClock.ConfiguredClockClass
    leafs["delay-asymmetry"] = foreignClock.DelayAsymmetry
    leafs["ptsf-loss-announce"] = foreignClock.PtsfLossAnnounce
    leafs["ptsf-loss-sync"] = foreignClock.PtsfLossSync
    return leafs
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetBundleName() string { return "cisco_ios_xr" }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetYangName() string { return "foreign-clock" }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) SetParent(parent types.Entity) { foreignClock.parent = parent }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetParent() types.Entity { return foreignClock.parent }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetParentYangName() string { return "interface-foreign-master" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock
// Foreign clock information
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Class. The type is interface{} with range: 0..255.
    Class interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Steps removed. The type is interface{} with range: 0..65535.
    StepsRemoved interface{}

    // Time source. The type is PtpBagClockTimeSource.
    TimeSource interface{}

    // The clock is frequency traceable. The type is bool.
    FrequencyTraceable interface{}

    // The clock is time traceable. The type is bool.
    TimeTraceable interface{}

    // Timescale. The type is PtpBagClockTimescale.
    Timescale interface{}

    // Leap Seconds. The type is PtpBagClockLeapSeconds. Units are second.
    LeapSeconds interface{}

    // The clock is the local clock. The type is bool.
    Local interface{}

    // The configured clock class. The type is interface{} with range: 0..255.
    ConfiguredClockClass interface{}

    // The configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // UTC offset.
    UtcOffset Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset

    // Receiver.
    Receiver Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver

    // Sender.
    Sender Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetFilter() yfilter.YFilter { return foreignClock.YFilter }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) SetFilter(yf yfilter.YFilter) { foreignClock.YFilter = yf }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "priority1" { return "Priority1" }
    if yname == "priority2" { return "Priority2" }
    if yname == "class" { return "Class" }
    if yname == "accuracy" { return "Accuracy" }
    if yname == "offset-log-variance" { return "OffsetLogVariance" }
    if yname == "steps-removed" { return "StepsRemoved" }
    if yname == "time-source" { return "TimeSource" }
    if yname == "frequency-traceable" { return "FrequencyTraceable" }
    if yname == "time-traceable" { return "TimeTraceable" }
    if yname == "timescale" { return "Timescale" }
    if yname == "leap-seconds" { return "LeapSeconds" }
    if yname == "local" { return "Local" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "utc-offset" { return "UtcOffset" }
    if yname == "receiver" { return "Receiver" }
    if yname == "sender" { return "Sender" }
    return ""
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetSegmentPath() string {
    return "foreign-clock"
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "utc-offset" {
        return &foreignClock.UtcOffset
    }
    if childYangName == "receiver" {
        return &foreignClock.Receiver
    }
    if childYangName == "sender" {
        return &foreignClock.Sender
    }
    return nil
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["utc-offset"] = &foreignClock.UtcOffset
    children["receiver"] = &foreignClock.Receiver
    children["sender"] = &foreignClock.Sender
    return children
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = foreignClock.ClockId
    leafs["priority1"] = foreignClock.Priority1
    leafs["priority2"] = foreignClock.Priority2
    leafs["class"] = foreignClock.Class
    leafs["accuracy"] = foreignClock.Accuracy
    leafs["offset-log-variance"] = foreignClock.OffsetLogVariance
    leafs["steps-removed"] = foreignClock.StepsRemoved
    leafs["time-source"] = foreignClock.TimeSource
    leafs["frequency-traceable"] = foreignClock.FrequencyTraceable
    leafs["time-traceable"] = foreignClock.TimeTraceable
    leafs["timescale"] = foreignClock.Timescale
    leafs["leap-seconds"] = foreignClock.LeapSeconds
    leafs["local"] = foreignClock.Local
    leafs["configured-clock-class"] = foreignClock.ConfiguredClockClass
    leafs["configured-priority"] = foreignClock.ConfiguredPriority
    return leafs
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetBundleName() string { return "cisco_ios_xr" }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetYangName() string { return "foreign-clock" }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) SetParent(parent types.Entity) { foreignClock.parent = parent }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetParent() types.Entity { return foreignClock.parent }

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset
// UTC offset
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetFilter() yfilter.YFilter { return utcOffset.YFilter }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) SetFilter(yf yfilter.YFilter) { utcOffset.YFilter = yf }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetGoName(yname string) string {
    if yname == "current-offset" { return "CurrentOffset" }
    if yname == "offset-valid" { return "OffsetValid" }
    return ""
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetSegmentPath() string {
    return "utc-offset"
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-offset"] = utcOffset.CurrentOffset
    leafs["offset-valid"] = utcOffset.OffsetValid
    return leafs
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetBundleName() string { return "cisco_ios_xr" }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetYangName() string { return "utc-offset" }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) SetParent(parent types.Entity) { utcOffset.parent = parent }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetParent() types.Entity { return utcOffset.parent }

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver
// Receiver
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetSegmentPath() string {
    return "receiver"
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = receiver.ClockId
    leafs["port-number"] = receiver.PortNumber
    return leafs
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetBundleName() string { return "cisco_ios_xr" }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetYangName() string { return "receiver" }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender
// Sender
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetFilter() yfilter.YFilter { return sender.YFilter }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) SetFilter(yf yfilter.YFilter) { sender.YFilter = yf }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetSegmentPath() string {
    return "sender"
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = sender.ClockId
    leafs["port-number"] = sender.PortNumber
    return leafs
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetBundleName() string { return "cisco_ios_xr" }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetYangName() string { return "sender" }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) SetParent(parent types.Entity) { sender.parent = parent }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetParent() types.Entity { return sender.parent }

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address
// The address of the clock
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetYangName() string { return "address" }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address
// IPv6 address
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetFilter() yfilter.YFilter { return announceGrant.YFilter }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) SetFilter(yf yfilter.YFilter) { announceGrant.YFilter = yf }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetSegmentPath() string {
    return "announce-grant"
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = announceGrant.LogGrantInterval
    leafs["grant-duration"] = announceGrant.GrantDuration
    return leafs
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetBundleName() string { return "cisco_ios_xr" }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetYangName() string { return "announce-grant" }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) SetParent(parent types.Entity) { announceGrant.parent = parent }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetParent() types.Entity { return announceGrant.parent }

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant
// Unicast grant information for sync messages
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetFilter() yfilter.YFilter { return syncGrant.YFilter }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) SetFilter(yf yfilter.YFilter) { syncGrant.YFilter = yf }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetSegmentPath() string {
    return "sync-grant"
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = syncGrant.LogGrantInterval
    leafs["grant-duration"] = syncGrant.GrantDuration
    return leafs
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetBundleName() string { return "cisco_ios_xr" }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetYangName() string { return "sync-grant" }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) SetParent(parent types.Entity) { syncGrant.parent = parent }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetParent() types.Entity { return syncGrant.parent }

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetParentYangName() string { return "foreign-clock" }

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetFilter() yfilter.YFilter { return delayResponseGrant.YFilter }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) SetFilter(yf yfilter.YFilter) { delayResponseGrant.YFilter = yf }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetSegmentPath() string {
    return "delay-response-grant"
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = delayResponseGrant.LogGrantInterval
    leafs["grant-duration"] = delayResponseGrant.GrantDuration
    return leafs
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetBundleName() string { return "cisco_ios_xr" }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetYangName() string { return "delay-response-grant" }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) SetParent(parent types.Entity) { delayResponseGrant.parent = parent }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetParent() types.Entity { return delayResponseGrant.parent }

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetParentYangName() string { return "foreign-clock" }

// Ptp_LocalClock
// Local clock operational data
type Ptp_LocalClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The PTP domain of the local clock. The type is interface{} with range:
    // 0..255.
    Domain interface{}

    // Local clock.
    ClockProperties Ptp_LocalClock_ClockProperties
}

func (localClock *Ptp_LocalClock) GetFilter() yfilter.YFilter { return localClock.YFilter }

func (localClock *Ptp_LocalClock) SetFilter(yf yfilter.YFilter) { localClock.YFilter = yf }

func (localClock *Ptp_LocalClock) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "clock-properties" { return "ClockProperties" }
    return ""
}

func (localClock *Ptp_LocalClock) GetSegmentPath() string {
    return "local-clock"
}

func (localClock *Ptp_LocalClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-properties" {
        return &localClock.ClockProperties
    }
    return nil
}

func (localClock *Ptp_LocalClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-properties"] = &localClock.ClockProperties
    return children
}

func (localClock *Ptp_LocalClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = localClock.Domain
    return leafs
}

func (localClock *Ptp_LocalClock) GetBundleName() string { return "cisco_ios_xr" }

func (localClock *Ptp_LocalClock) GetYangName() string { return "local-clock" }

func (localClock *Ptp_LocalClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localClock *Ptp_LocalClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localClock *Ptp_LocalClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localClock *Ptp_LocalClock) SetParent(parent types.Entity) { localClock.parent = parent }

func (localClock *Ptp_LocalClock) GetParent() types.Entity { return localClock.parent }

func (localClock *Ptp_LocalClock) GetParentYangName() string { return "ptp" }

// Ptp_LocalClock_ClockProperties
// Local clock
type Ptp_LocalClock_ClockProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Class. The type is interface{} with range: 0..255.
    Class interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Steps removed. The type is interface{} with range: 0..65535.
    StepsRemoved interface{}

    // Time source. The type is PtpBagClockTimeSource.
    TimeSource interface{}

    // The clock is frequency traceable. The type is bool.
    FrequencyTraceable interface{}

    // The clock is time traceable. The type is bool.
    TimeTraceable interface{}

    // Timescale. The type is PtpBagClockTimescale.
    Timescale interface{}

    // Leap Seconds. The type is PtpBagClockLeapSeconds. Units are second.
    LeapSeconds interface{}

    // The clock is the local clock. The type is bool.
    Local interface{}

    // The configured clock class. The type is interface{} with range: 0..255.
    ConfiguredClockClass interface{}

    // The configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // UTC offset.
    UtcOffset Ptp_LocalClock_ClockProperties_UtcOffset

    // Receiver.
    Receiver Ptp_LocalClock_ClockProperties_Receiver

    // Sender.
    Sender Ptp_LocalClock_ClockProperties_Sender
}

func (clockProperties *Ptp_LocalClock_ClockProperties) GetFilter() yfilter.YFilter { return clockProperties.YFilter }

func (clockProperties *Ptp_LocalClock_ClockProperties) SetFilter(yf yfilter.YFilter) { clockProperties.YFilter = yf }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "priority1" { return "Priority1" }
    if yname == "priority2" { return "Priority2" }
    if yname == "class" { return "Class" }
    if yname == "accuracy" { return "Accuracy" }
    if yname == "offset-log-variance" { return "OffsetLogVariance" }
    if yname == "steps-removed" { return "StepsRemoved" }
    if yname == "time-source" { return "TimeSource" }
    if yname == "frequency-traceable" { return "FrequencyTraceable" }
    if yname == "time-traceable" { return "TimeTraceable" }
    if yname == "timescale" { return "Timescale" }
    if yname == "leap-seconds" { return "LeapSeconds" }
    if yname == "local" { return "Local" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "utc-offset" { return "UtcOffset" }
    if yname == "receiver" { return "Receiver" }
    if yname == "sender" { return "Sender" }
    return ""
}

func (clockProperties *Ptp_LocalClock_ClockProperties) GetSegmentPath() string {
    return "clock-properties"
}

func (clockProperties *Ptp_LocalClock_ClockProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "utc-offset" {
        return &clockProperties.UtcOffset
    }
    if childYangName == "receiver" {
        return &clockProperties.Receiver
    }
    if childYangName == "sender" {
        return &clockProperties.Sender
    }
    return nil
}

func (clockProperties *Ptp_LocalClock_ClockProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["utc-offset"] = &clockProperties.UtcOffset
    children["receiver"] = &clockProperties.Receiver
    children["sender"] = &clockProperties.Sender
    return children
}

func (clockProperties *Ptp_LocalClock_ClockProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = clockProperties.ClockId
    leafs["priority1"] = clockProperties.Priority1
    leafs["priority2"] = clockProperties.Priority2
    leafs["class"] = clockProperties.Class
    leafs["accuracy"] = clockProperties.Accuracy
    leafs["offset-log-variance"] = clockProperties.OffsetLogVariance
    leafs["steps-removed"] = clockProperties.StepsRemoved
    leafs["time-source"] = clockProperties.TimeSource
    leafs["frequency-traceable"] = clockProperties.FrequencyTraceable
    leafs["time-traceable"] = clockProperties.TimeTraceable
    leafs["timescale"] = clockProperties.Timescale
    leafs["leap-seconds"] = clockProperties.LeapSeconds
    leafs["local"] = clockProperties.Local
    leafs["configured-clock-class"] = clockProperties.ConfiguredClockClass
    leafs["configured-priority"] = clockProperties.ConfiguredPriority
    return leafs
}

func (clockProperties *Ptp_LocalClock_ClockProperties) GetBundleName() string { return "cisco_ios_xr" }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetYangName() string { return "clock-properties" }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockProperties *Ptp_LocalClock_ClockProperties) SetParent(parent types.Entity) { clockProperties.parent = parent }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetParent() types.Entity { return clockProperties.parent }

func (clockProperties *Ptp_LocalClock_ClockProperties) GetParentYangName() string { return "local-clock" }

// Ptp_LocalClock_ClockProperties_UtcOffset
// UTC offset
type Ptp_LocalClock_ClockProperties_UtcOffset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetFilter() yfilter.YFilter { return utcOffset.YFilter }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) SetFilter(yf yfilter.YFilter) { utcOffset.YFilter = yf }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetGoName(yname string) string {
    if yname == "current-offset" { return "CurrentOffset" }
    if yname == "offset-valid" { return "OffsetValid" }
    return ""
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetSegmentPath() string {
    return "utc-offset"
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-offset"] = utcOffset.CurrentOffset
    leafs["offset-valid"] = utcOffset.OffsetValid
    return leafs
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetBundleName() string { return "cisco_ios_xr" }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetYangName() string { return "utc-offset" }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) SetParent(parent types.Entity) { utcOffset.parent = parent }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetParent() types.Entity { return utcOffset.parent }

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetParentYangName() string { return "clock-properties" }

// Ptp_LocalClock_ClockProperties_Receiver
// Receiver
type Ptp_LocalClock_ClockProperties_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetSegmentPath() string {
    return "receiver"
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = receiver.ClockId
    leafs["port-number"] = receiver.PortNumber
    return leafs
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetBundleName() string { return "cisco_ios_xr" }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetYangName() string { return "receiver" }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetParentYangName() string { return "clock-properties" }

// Ptp_LocalClock_ClockProperties_Sender
// Sender
type Ptp_LocalClock_ClockProperties_Sender struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetFilter() yfilter.YFilter { return sender.YFilter }

func (sender *Ptp_LocalClock_ClockProperties_Sender) SetFilter(yf yfilter.YFilter) { sender.YFilter = yf }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetSegmentPath() string {
    return "sender"
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = sender.ClockId
    leafs["port-number"] = sender.PortNumber
    return leafs
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetBundleName() string { return "cisco_ios_xr" }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetYangName() string { return "sender" }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sender *Ptp_LocalClock_ClockProperties_Sender) SetParent(parent types.Entity) { sender.parent = parent }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetParent() types.Entity { return sender.parent }

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetParentYangName() string { return "clock-properties" }

// Ptp_InterfacePacketCounters
// Table for interface packet counter operational
// data
type Ptp_InterfacePacketCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface packet counter operational data. The type is slice of
    // Ptp_InterfacePacketCounters_InterfacePacketCounter.
    InterfacePacketCounter []Ptp_InterfacePacketCounters_InterfacePacketCounter
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetFilter() yfilter.YFilter { return interfacePacketCounters.YFilter }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) SetFilter(yf yfilter.YFilter) { interfacePacketCounters.YFilter = yf }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetGoName(yname string) string {
    if yname == "interface-packet-counter" { return "InterfacePacketCounter" }
    return ""
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetSegmentPath() string {
    return "interface-packet-counters"
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-packet-counter" {
        for _, c := range interfacePacketCounters.InterfacePacketCounter {
            if interfacePacketCounters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfacePacketCounters_InterfacePacketCounter{}
        interfacePacketCounters.InterfacePacketCounter = append(interfacePacketCounters.InterfacePacketCounter, child)
        return &interfacePacketCounters.InterfacePacketCounter[len(interfacePacketCounters.InterfacePacketCounter)-1]
    }
    return nil
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfacePacketCounters.InterfacePacketCounter {
        children[interfacePacketCounters.InterfacePacketCounter[i].GetSegmentPath()] = &interfacePacketCounters.InterfacePacketCounter[i]
    }
    return children
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetBundleName() string { return "cisco_ios_xr" }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetYangName() string { return "interface-packet-counters" }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) SetParent(parent types.Entity) { interfacePacketCounters.parent = parent }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetParent() types.Entity { return interfacePacketCounters.parent }

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetParentYangName() string { return "ptp" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter
// Interface packet counter operational data
type Ptp_InterfacePacketCounters_InterfacePacketCounter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Packet counters.
    Counters Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters

    // Packet counters for each peer on this interface. The type is slice of
    // Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter.
    PeerCounter []Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetFilter() yfilter.YFilter { return interfacePacketCounter.YFilter }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) SetFilter(yf yfilter.YFilter) { interfacePacketCounter.YFilter = yf }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "counters" { return "Counters" }
    if yname == "peer-counter" { return "PeerCounter" }
    return ""
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetSegmentPath() string {
    return "interface-packet-counter" + "[interface-name='" + fmt.Sprintf("%v", interfacePacketCounter.InterfaceName) + "']"
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &interfacePacketCounter.Counters
    }
    if childYangName == "peer-counter" {
        for _, c := range interfacePacketCounter.PeerCounter {
            if interfacePacketCounter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter{}
        interfacePacketCounter.PeerCounter = append(interfacePacketCounter.PeerCounter, child)
        return &interfacePacketCounter.PeerCounter[len(interfacePacketCounter.PeerCounter)-1]
    }
    return nil
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &interfacePacketCounter.Counters
    for i := range interfacePacketCounter.PeerCounter {
        children[interfacePacketCounter.PeerCounter[i].GetSegmentPath()] = &interfacePacketCounter.PeerCounter[i]
    }
    return children
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfacePacketCounter.InterfaceName
    return leafs
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetBundleName() string { return "cisco_ios_xr" }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetYangName() string { return "interface-packet-counter" }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) SetParent(parent types.Entity) { interfacePacketCounter.parent = parent }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetParent() types.Entity { return interfacePacketCounter.parent }

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetParentYangName() string { return "interface-packet-counters" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters
// Packet counters
type Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of announce packets sent. The type is interface{} with range:
    // 0..4294967295.
    AnnounceSent interface{}

    // Number of announce packets received. The type is interface{} with range:
    // 0..4294967295.
    AnnounceReceived interface{}

    // Number of announce packets dropped. The type is interface{} with range:
    // 0..4294967295.
    AnnounceDropped interface{}

    // Number of sync packets sent. The type is interface{} with range:
    // 0..4294967295.
    SyncSent interface{}

    // Number of sync packets received. The type is interface{} with range:
    // 0..4294967295.
    SyncReceived interface{}

    // Number of sync packetsdropped. The type is interface{} with range:
    // 0..4294967295.
    SyncDropped interface{}

    // Number of follow-up packets sent. The type is interface{} with range:
    // 0..4294967295.
    FollowUpSent interface{}

    // Number of follow-up packets received. The type is interface{} with range:
    // 0..4294967295.
    FollowUpReceived interface{}

    // Number of follow-up packets dropped. The type is interface{} with range:
    // 0..4294967295.
    FollowUpDropped interface{}

    // Number of delay-request packets sent. The type is interface{} with range:
    // 0..4294967295.
    DelayRequestSent interface{}

    // Number of delay-request packets received. The type is interface{} with
    // range: 0..4294967295.
    DelayRequestReceived interface{}

    // Number of delay-request packets dropped. The type is interface{} with
    // range: 0..4294967295.
    DelayRequestDropped interface{}

    // Number of delay-response packets sent. The type is interface{} with range:
    // 0..4294967295.
    DelayResponseSent interface{}

    // Number of delay-response packets received. The type is interface{} with
    // range: 0..4294967295.
    DelayResponseReceived interface{}

    // Number of delay-response packets dropped. The type is interface{} with
    // range: 0..4294967295.
    DelayResponseDropped interface{}

    // Number of peer-delay-request packets sent. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestSent interface{}

    // Number of peer-delay-request packets received. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestReceived interface{}

    // Number of peer-delay-request packets dropped. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestDropped interface{}

    // Number of peer-delay-response packets sent. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayResponseSent interface{}

    // Number of peer-delay-response packets received. The type is interface{}
    // with range: 0..4294967295.
    PeerDelayResponseReceived interface{}

    // Number of peer-delay-response packets dropped. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayResponseDropped interface{}

    // Number of peer-delay-response follow-up packets sent. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpSent interface{}

    // Number of peer-delay-response follow-up packets received. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpReceived interface{}

    // Number of peer-delay-response follow-up packets dropped. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpDropped interface{}

    // Number of signaling packets sent. The type is interface{} with range:
    // 0..4294967295.
    SignalingSent interface{}

    // Number of signaling packets received. The type is interface{} with range:
    // 0..4294967295.
    SignalingReceived interface{}

    // Number of signaling packets dropped. The type is interface{} with range:
    // 0..4294967295.
    SignalingDropped interface{}

    // Number of management messages sent. The type is interface{} with range:
    // 0..4294967295.
    ManagementSent interface{}

    // Number of management messages received. The type is interface{} with range:
    // 0..4294967295.
    ManagementReceived interface{}

    // Number of management messages dropped. The type is interface{} with range:
    // 0..4294967295.
    ManagementDropped interface{}

    // Number of other packets sent. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsSent interface{}

    // Number of other packets received. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsReceived interface{}

    // Number of other packets dropped. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsDropped interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsSent interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsReceived interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsDropped interface{}
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetGoName(yname string) string {
    if yname == "announce-sent" { return "AnnounceSent" }
    if yname == "announce-received" { return "AnnounceReceived" }
    if yname == "announce-dropped" { return "AnnounceDropped" }
    if yname == "sync-sent" { return "SyncSent" }
    if yname == "sync-received" { return "SyncReceived" }
    if yname == "sync-dropped" { return "SyncDropped" }
    if yname == "follow-up-sent" { return "FollowUpSent" }
    if yname == "follow-up-received" { return "FollowUpReceived" }
    if yname == "follow-up-dropped" { return "FollowUpDropped" }
    if yname == "delay-request-sent" { return "DelayRequestSent" }
    if yname == "delay-request-received" { return "DelayRequestReceived" }
    if yname == "delay-request-dropped" { return "DelayRequestDropped" }
    if yname == "delay-response-sent" { return "DelayResponseSent" }
    if yname == "delay-response-received" { return "DelayResponseReceived" }
    if yname == "delay-response-dropped" { return "DelayResponseDropped" }
    if yname == "peer-delay-request-sent" { return "PeerDelayRequestSent" }
    if yname == "peer-delay-request-received" { return "PeerDelayRequestReceived" }
    if yname == "peer-delay-request-dropped" { return "PeerDelayRequestDropped" }
    if yname == "peer-delay-response-sent" { return "PeerDelayResponseSent" }
    if yname == "peer-delay-response-received" { return "PeerDelayResponseReceived" }
    if yname == "peer-delay-response-dropped" { return "PeerDelayResponseDropped" }
    if yname == "peer-delay-response-follow-up-sent" { return "PeerDelayResponseFollowUpSent" }
    if yname == "peer-delay-response-follow-up-received" { return "PeerDelayResponseFollowUpReceived" }
    if yname == "peer-delay-response-follow-up-dropped" { return "PeerDelayResponseFollowUpDropped" }
    if yname == "signaling-sent" { return "SignalingSent" }
    if yname == "signaling-received" { return "SignalingReceived" }
    if yname == "signaling-dropped" { return "SignalingDropped" }
    if yname == "management-sent" { return "ManagementSent" }
    if yname == "management-received" { return "ManagementReceived" }
    if yname == "management-dropped" { return "ManagementDropped" }
    if yname == "other-packets-sent" { return "OtherPacketsSent" }
    if yname == "other-packets-received" { return "OtherPacketsReceived" }
    if yname == "other-packets-dropped" { return "OtherPacketsDropped" }
    if yname == "total-packets-sent" { return "TotalPacketsSent" }
    if yname == "total-packets-received" { return "TotalPacketsReceived" }
    if yname == "total-packets-dropped" { return "TotalPacketsDropped" }
    return ""
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["announce-sent"] = counters.AnnounceSent
    leafs["announce-received"] = counters.AnnounceReceived
    leafs["announce-dropped"] = counters.AnnounceDropped
    leafs["sync-sent"] = counters.SyncSent
    leafs["sync-received"] = counters.SyncReceived
    leafs["sync-dropped"] = counters.SyncDropped
    leafs["follow-up-sent"] = counters.FollowUpSent
    leafs["follow-up-received"] = counters.FollowUpReceived
    leafs["follow-up-dropped"] = counters.FollowUpDropped
    leafs["delay-request-sent"] = counters.DelayRequestSent
    leafs["delay-request-received"] = counters.DelayRequestReceived
    leafs["delay-request-dropped"] = counters.DelayRequestDropped
    leafs["delay-response-sent"] = counters.DelayResponseSent
    leafs["delay-response-received"] = counters.DelayResponseReceived
    leafs["delay-response-dropped"] = counters.DelayResponseDropped
    leafs["peer-delay-request-sent"] = counters.PeerDelayRequestSent
    leafs["peer-delay-request-received"] = counters.PeerDelayRequestReceived
    leafs["peer-delay-request-dropped"] = counters.PeerDelayRequestDropped
    leafs["peer-delay-response-sent"] = counters.PeerDelayResponseSent
    leafs["peer-delay-response-received"] = counters.PeerDelayResponseReceived
    leafs["peer-delay-response-dropped"] = counters.PeerDelayResponseDropped
    leafs["peer-delay-response-follow-up-sent"] = counters.PeerDelayResponseFollowUpSent
    leafs["peer-delay-response-follow-up-received"] = counters.PeerDelayResponseFollowUpReceived
    leafs["peer-delay-response-follow-up-dropped"] = counters.PeerDelayResponseFollowUpDropped
    leafs["signaling-sent"] = counters.SignalingSent
    leafs["signaling-received"] = counters.SignalingReceived
    leafs["signaling-dropped"] = counters.SignalingDropped
    leafs["management-sent"] = counters.ManagementSent
    leafs["management-received"] = counters.ManagementReceived
    leafs["management-dropped"] = counters.ManagementDropped
    leafs["other-packets-sent"] = counters.OtherPacketsSent
    leafs["other-packets-received"] = counters.OtherPacketsReceived
    leafs["other-packets-dropped"] = counters.OtherPacketsDropped
    leafs["total-packets-sent"] = counters.TotalPacketsSent
    leafs["total-packets-received"] = counters.TotalPacketsReceived
    leafs["total-packets-dropped"] = counters.TotalPacketsDropped
    return leafs
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetYangName() string { return "counters" }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetParentYangName() string { return "interface-packet-counter" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter
// Packet counters for each peer on this interface
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer address.
    Address Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address

    // Packet counters.
    Counters Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetFilter() yfilter.YFilter { return peerCounter.YFilter }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) SetFilter(yf yfilter.YFilter) { peerCounter.YFilter = yf }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetSegmentPath() string {
    return "peer-counter"
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &peerCounter.Address
    }
    if childYangName == "counters" {
        return &peerCounter.Counters
    }
    return nil
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &peerCounter.Address
    children["counters"] = &peerCounter.Counters
    return children
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetBundleName() string { return "cisco_ios_xr" }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetYangName() string { return "peer-counter" }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) SetParent(parent types.Entity) { peerCounter.parent = parent }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetParent() types.Entity { return peerCounter.parent }

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetParentYangName() string { return "interface-packet-counter" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address
// Peer address
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetYangName() string { return "address" }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetParentYangName() string { return "peer-counter" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address
// IPv6 address
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters
// Packet counters
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of announce packets sent. The type is interface{} with range:
    // 0..4294967295.
    AnnounceSent interface{}

    // Number of announce packets received. The type is interface{} with range:
    // 0..4294967295.
    AnnounceReceived interface{}

    // Number of announce packets dropped. The type is interface{} with range:
    // 0..4294967295.
    AnnounceDropped interface{}

    // Number of sync packets sent. The type is interface{} with range:
    // 0..4294967295.
    SyncSent interface{}

    // Number of sync packets received. The type is interface{} with range:
    // 0..4294967295.
    SyncReceived interface{}

    // Number of sync packetsdropped. The type is interface{} with range:
    // 0..4294967295.
    SyncDropped interface{}

    // Number of follow-up packets sent. The type is interface{} with range:
    // 0..4294967295.
    FollowUpSent interface{}

    // Number of follow-up packets received. The type is interface{} with range:
    // 0..4294967295.
    FollowUpReceived interface{}

    // Number of follow-up packets dropped. The type is interface{} with range:
    // 0..4294967295.
    FollowUpDropped interface{}

    // Number of delay-request packets sent. The type is interface{} with range:
    // 0..4294967295.
    DelayRequestSent interface{}

    // Number of delay-request packets received. The type is interface{} with
    // range: 0..4294967295.
    DelayRequestReceived interface{}

    // Number of delay-request packets dropped. The type is interface{} with
    // range: 0..4294967295.
    DelayRequestDropped interface{}

    // Number of delay-response packets sent. The type is interface{} with range:
    // 0..4294967295.
    DelayResponseSent interface{}

    // Number of delay-response packets received. The type is interface{} with
    // range: 0..4294967295.
    DelayResponseReceived interface{}

    // Number of delay-response packets dropped. The type is interface{} with
    // range: 0..4294967295.
    DelayResponseDropped interface{}

    // Number of peer-delay-request packets sent. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestSent interface{}

    // Number of peer-delay-request packets received. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestReceived interface{}

    // Number of peer-delay-request packets dropped. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayRequestDropped interface{}

    // Number of peer-delay-response packets sent. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayResponseSent interface{}

    // Number of peer-delay-response packets received. The type is interface{}
    // with range: 0..4294967295.
    PeerDelayResponseReceived interface{}

    // Number of peer-delay-response packets dropped. The type is interface{} with
    // range: 0..4294967295.
    PeerDelayResponseDropped interface{}

    // Number of peer-delay-response follow-up packets sent. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpSent interface{}

    // Number of peer-delay-response follow-up packets received. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpReceived interface{}

    // Number of peer-delay-response follow-up packets dropped. The type is
    // interface{} with range: 0..4294967295.
    PeerDelayResponseFollowUpDropped interface{}

    // Number of signaling packets sent. The type is interface{} with range:
    // 0..4294967295.
    SignalingSent interface{}

    // Number of signaling packets received. The type is interface{} with range:
    // 0..4294967295.
    SignalingReceived interface{}

    // Number of signaling packets dropped. The type is interface{} with range:
    // 0..4294967295.
    SignalingDropped interface{}

    // Number of management messages sent. The type is interface{} with range:
    // 0..4294967295.
    ManagementSent interface{}

    // Number of management messages received. The type is interface{} with range:
    // 0..4294967295.
    ManagementReceived interface{}

    // Number of management messages dropped. The type is interface{} with range:
    // 0..4294967295.
    ManagementDropped interface{}

    // Number of other packets sent. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsSent interface{}

    // Number of other packets received. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsReceived interface{}

    // Number of other packets dropped. The type is interface{} with range:
    // 0..4294967295.
    OtherPacketsDropped interface{}

    // Total number of packets sent. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsSent interface{}

    // Total number of packets received. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsReceived interface{}

    // Total number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    TotalPacketsDropped interface{}
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetGoName(yname string) string {
    if yname == "announce-sent" { return "AnnounceSent" }
    if yname == "announce-received" { return "AnnounceReceived" }
    if yname == "announce-dropped" { return "AnnounceDropped" }
    if yname == "sync-sent" { return "SyncSent" }
    if yname == "sync-received" { return "SyncReceived" }
    if yname == "sync-dropped" { return "SyncDropped" }
    if yname == "follow-up-sent" { return "FollowUpSent" }
    if yname == "follow-up-received" { return "FollowUpReceived" }
    if yname == "follow-up-dropped" { return "FollowUpDropped" }
    if yname == "delay-request-sent" { return "DelayRequestSent" }
    if yname == "delay-request-received" { return "DelayRequestReceived" }
    if yname == "delay-request-dropped" { return "DelayRequestDropped" }
    if yname == "delay-response-sent" { return "DelayResponseSent" }
    if yname == "delay-response-received" { return "DelayResponseReceived" }
    if yname == "delay-response-dropped" { return "DelayResponseDropped" }
    if yname == "peer-delay-request-sent" { return "PeerDelayRequestSent" }
    if yname == "peer-delay-request-received" { return "PeerDelayRequestReceived" }
    if yname == "peer-delay-request-dropped" { return "PeerDelayRequestDropped" }
    if yname == "peer-delay-response-sent" { return "PeerDelayResponseSent" }
    if yname == "peer-delay-response-received" { return "PeerDelayResponseReceived" }
    if yname == "peer-delay-response-dropped" { return "PeerDelayResponseDropped" }
    if yname == "peer-delay-response-follow-up-sent" { return "PeerDelayResponseFollowUpSent" }
    if yname == "peer-delay-response-follow-up-received" { return "PeerDelayResponseFollowUpReceived" }
    if yname == "peer-delay-response-follow-up-dropped" { return "PeerDelayResponseFollowUpDropped" }
    if yname == "signaling-sent" { return "SignalingSent" }
    if yname == "signaling-received" { return "SignalingReceived" }
    if yname == "signaling-dropped" { return "SignalingDropped" }
    if yname == "management-sent" { return "ManagementSent" }
    if yname == "management-received" { return "ManagementReceived" }
    if yname == "management-dropped" { return "ManagementDropped" }
    if yname == "other-packets-sent" { return "OtherPacketsSent" }
    if yname == "other-packets-received" { return "OtherPacketsReceived" }
    if yname == "other-packets-dropped" { return "OtherPacketsDropped" }
    if yname == "total-packets-sent" { return "TotalPacketsSent" }
    if yname == "total-packets-received" { return "TotalPacketsReceived" }
    if yname == "total-packets-dropped" { return "TotalPacketsDropped" }
    return ""
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["announce-sent"] = counters.AnnounceSent
    leafs["announce-received"] = counters.AnnounceReceived
    leafs["announce-dropped"] = counters.AnnounceDropped
    leafs["sync-sent"] = counters.SyncSent
    leafs["sync-received"] = counters.SyncReceived
    leafs["sync-dropped"] = counters.SyncDropped
    leafs["follow-up-sent"] = counters.FollowUpSent
    leafs["follow-up-received"] = counters.FollowUpReceived
    leafs["follow-up-dropped"] = counters.FollowUpDropped
    leafs["delay-request-sent"] = counters.DelayRequestSent
    leafs["delay-request-received"] = counters.DelayRequestReceived
    leafs["delay-request-dropped"] = counters.DelayRequestDropped
    leafs["delay-response-sent"] = counters.DelayResponseSent
    leafs["delay-response-received"] = counters.DelayResponseReceived
    leafs["delay-response-dropped"] = counters.DelayResponseDropped
    leafs["peer-delay-request-sent"] = counters.PeerDelayRequestSent
    leafs["peer-delay-request-received"] = counters.PeerDelayRequestReceived
    leafs["peer-delay-request-dropped"] = counters.PeerDelayRequestDropped
    leafs["peer-delay-response-sent"] = counters.PeerDelayResponseSent
    leafs["peer-delay-response-received"] = counters.PeerDelayResponseReceived
    leafs["peer-delay-response-dropped"] = counters.PeerDelayResponseDropped
    leafs["peer-delay-response-follow-up-sent"] = counters.PeerDelayResponseFollowUpSent
    leafs["peer-delay-response-follow-up-received"] = counters.PeerDelayResponseFollowUpReceived
    leafs["peer-delay-response-follow-up-dropped"] = counters.PeerDelayResponseFollowUpDropped
    leafs["signaling-sent"] = counters.SignalingSent
    leafs["signaling-received"] = counters.SignalingReceived
    leafs["signaling-dropped"] = counters.SignalingDropped
    leafs["management-sent"] = counters.ManagementSent
    leafs["management-received"] = counters.ManagementReceived
    leafs["management-dropped"] = counters.ManagementDropped
    leafs["other-packets-sent"] = counters.OtherPacketsSent
    leafs["other-packets-received"] = counters.OtherPacketsReceived
    leafs["other-packets-dropped"] = counters.OtherPacketsDropped
    leafs["total-packets-sent"] = counters.TotalPacketsSent
    leafs["total-packets-received"] = counters.TotalPacketsReceived
    leafs["total-packets-dropped"] = counters.TotalPacketsDropped
    return leafs
}

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetYangName() string { return "counters" }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetParentYangName() string { return "peer-counter" }

// Ptp_AdvertisedClock
// Advertised clock operational data
type Ptp_AdvertisedClock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The PTP domain of that the advertised clock is in. The type is interface{}
    // with range: 0..255.
    Domain interface{}

    // Whether the advertised time source is configured. The type is bool.
    TimeSourceConfigured interface{}

    // The time source received from the parent clock. The type is
    // PtpBagClockTimeSource.
    ReceivedTimeSource interface{}

    // Whether the advertised timescale is configured. The type is bool.
    TimescaleConfigured interface{}

    // The timescale received from the parent clock. The type is
    // PtpBagClockTimescale.
    ReceivedTimescale interface{}

    // Advertised Clock.
    ClockProperties Ptp_AdvertisedClock_ClockProperties
}

func (advertisedClock *Ptp_AdvertisedClock) GetFilter() yfilter.YFilter { return advertisedClock.YFilter }

func (advertisedClock *Ptp_AdvertisedClock) SetFilter(yf yfilter.YFilter) { advertisedClock.YFilter = yf }

func (advertisedClock *Ptp_AdvertisedClock) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "time-source-configured" { return "TimeSourceConfigured" }
    if yname == "received-time-source" { return "ReceivedTimeSource" }
    if yname == "timescale-configured" { return "TimescaleConfigured" }
    if yname == "received-timescale" { return "ReceivedTimescale" }
    if yname == "clock-properties" { return "ClockProperties" }
    return ""
}

func (advertisedClock *Ptp_AdvertisedClock) GetSegmentPath() string {
    return "advertised-clock"
}

func (advertisedClock *Ptp_AdvertisedClock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-properties" {
        return &advertisedClock.ClockProperties
    }
    return nil
}

func (advertisedClock *Ptp_AdvertisedClock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-properties"] = &advertisedClock.ClockProperties
    return children
}

func (advertisedClock *Ptp_AdvertisedClock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = advertisedClock.Domain
    leafs["time-source-configured"] = advertisedClock.TimeSourceConfigured
    leafs["received-time-source"] = advertisedClock.ReceivedTimeSource
    leafs["timescale-configured"] = advertisedClock.TimescaleConfigured
    leafs["received-timescale"] = advertisedClock.ReceivedTimescale
    return leafs
}

func (advertisedClock *Ptp_AdvertisedClock) GetBundleName() string { return "cisco_ios_xr" }

func (advertisedClock *Ptp_AdvertisedClock) GetYangName() string { return "advertised-clock" }

func (advertisedClock *Ptp_AdvertisedClock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (advertisedClock *Ptp_AdvertisedClock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (advertisedClock *Ptp_AdvertisedClock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (advertisedClock *Ptp_AdvertisedClock) SetParent(parent types.Entity) { advertisedClock.parent = parent }

func (advertisedClock *Ptp_AdvertisedClock) GetParent() types.Entity { return advertisedClock.parent }

func (advertisedClock *Ptp_AdvertisedClock) GetParentYangName() string { return "ptp" }

// Ptp_AdvertisedClock_ClockProperties
// Advertised Clock
type Ptp_AdvertisedClock_ClockProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Class. The type is interface{} with range: 0..255.
    Class interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Steps removed. The type is interface{} with range: 0..65535.
    StepsRemoved interface{}

    // Time source. The type is PtpBagClockTimeSource.
    TimeSource interface{}

    // The clock is frequency traceable. The type is bool.
    FrequencyTraceable interface{}

    // The clock is time traceable. The type is bool.
    TimeTraceable interface{}

    // Timescale. The type is PtpBagClockTimescale.
    Timescale interface{}

    // Leap Seconds. The type is PtpBagClockLeapSeconds. Units are second.
    LeapSeconds interface{}

    // The clock is the local clock. The type is bool.
    Local interface{}

    // The configured clock class. The type is interface{} with range: 0..255.
    ConfiguredClockClass interface{}

    // The configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // UTC offset.
    UtcOffset Ptp_AdvertisedClock_ClockProperties_UtcOffset

    // Receiver.
    Receiver Ptp_AdvertisedClock_ClockProperties_Receiver

    // Sender.
    Sender Ptp_AdvertisedClock_ClockProperties_Sender
}

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetFilter() yfilter.YFilter { return clockProperties.YFilter }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) SetFilter(yf yfilter.YFilter) { clockProperties.YFilter = yf }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "priority1" { return "Priority1" }
    if yname == "priority2" { return "Priority2" }
    if yname == "class" { return "Class" }
    if yname == "accuracy" { return "Accuracy" }
    if yname == "offset-log-variance" { return "OffsetLogVariance" }
    if yname == "steps-removed" { return "StepsRemoved" }
    if yname == "time-source" { return "TimeSource" }
    if yname == "frequency-traceable" { return "FrequencyTraceable" }
    if yname == "time-traceable" { return "TimeTraceable" }
    if yname == "timescale" { return "Timescale" }
    if yname == "leap-seconds" { return "LeapSeconds" }
    if yname == "local" { return "Local" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "utc-offset" { return "UtcOffset" }
    if yname == "receiver" { return "Receiver" }
    if yname == "sender" { return "Sender" }
    return ""
}

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetSegmentPath() string {
    return "clock-properties"
}

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "utc-offset" {
        return &clockProperties.UtcOffset
    }
    if childYangName == "receiver" {
        return &clockProperties.Receiver
    }
    if childYangName == "sender" {
        return &clockProperties.Sender
    }
    return nil
}

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["utc-offset"] = &clockProperties.UtcOffset
    children["receiver"] = &clockProperties.Receiver
    children["sender"] = &clockProperties.Sender
    return children
}

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = clockProperties.ClockId
    leafs["priority1"] = clockProperties.Priority1
    leafs["priority2"] = clockProperties.Priority2
    leafs["class"] = clockProperties.Class
    leafs["accuracy"] = clockProperties.Accuracy
    leafs["offset-log-variance"] = clockProperties.OffsetLogVariance
    leafs["steps-removed"] = clockProperties.StepsRemoved
    leafs["time-source"] = clockProperties.TimeSource
    leafs["frequency-traceable"] = clockProperties.FrequencyTraceable
    leafs["time-traceable"] = clockProperties.TimeTraceable
    leafs["timescale"] = clockProperties.Timescale
    leafs["leap-seconds"] = clockProperties.LeapSeconds
    leafs["local"] = clockProperties.Local
    leafs["configured-clock-class"] = clockProperties.ConfiguredClockClass
    leafs["configured-priority"] = clockProperties.ConfiguredPriority
    return leafs
}

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetBundleName() string { return "cisco_ios_xr" }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetYangName() string { return "clock-properties" }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) SetParent(parent types.Entity) { clockProperties.parent = parent }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetParent() types.Entity { return clockProperties.parent }

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetParentYangName() string { return "advertised-clock" }

// Ptp_AdvertisedClock_ClockProperties_UtcOffset
// UTC offset
type Ptp_AdvertisedClock_ClockProperties_UtcOffset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetFilter() yfilter.YFilter { return utcOffset.YFilter }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) SetFilter(yf yfilter.YFilter) { utcOffset.YFilter = yf }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetGoName(yname string) string {
    if yname == "current-offset" { return "CurrentOffset" }
    if yname == "offset-valid" { return "OffsetValid" }
    return ""
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetSegmentPath() string {
    return "utc-offset"
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-offset"] = utcOffset.CurrentOffset
    leafs["offset-valid"] = utcOffset.OffsetValid
    return leafs
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetBundleName() string { return "cisco_ios_xr" }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetYangName() string { return "utc-offset" }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) SetParent(parent types.Entity) { utcOffset.parent = parent }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetParent() types.Entity { return utcOffset.parent }

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetParentYangName() string { return "clock-properties" }

// Ptp_AdvertisedClock_ClockProperties_Receiver
// Receiver
type Ptp_AdvertisedClock_ClockProperties_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetSegmentPath() string {
    return "receiver"
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = receiver.ClockId
    leafs["port-number"] = receiver.PortNumber
    return leafs
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetBundleName() string { return "cisco_ios_xr" }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetYangName() string { return "receiver" }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetParentYangName() string { return "clock-properties" }

// Ptp_AdvertisedClock_ClockProperties_Sender
// Sender
type Ptp_AdvertisedClock_ClockProperties_Sender struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetFilter() yfilter.YFilter { return sender.YFilter }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) SetFilter(yf yfilter.YFilter) { sender.YFilter = yf }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetSegmentPath() string {
    return "sender"
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = sender.ClockId
    leafs["port-number"] = sender.PortNumber
    return leafs
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetBundleName() string { return "cisco_ios_xr" }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetYangName() string { return "sender" }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) SetParent(parent types.Entity) { sender.parent = parent }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetParent() types.Entity { return sender.parent }

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetParentYangName() string { return "clock-properties" }

// Ptp_LeapSeconds
// Upcoming leap-seconds information
type Ptp_LeapSeconds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The current UTC offset, in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    CurrentOffset interface{}

    // Is the current UTC offset known to be valid?. The type is bool.
    OffsetValid interface{}

    // The URL of the file containing upcoming leap seconds. The type is string.
    // Units are second.
    SourceFile interface{}

    // Source file expiry timestamp, in seconds since UNIX epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    SourceExpiryDate interface{}

    // Source file polling frequency, in days. The type is interface{} with range:
    // 0..4294967295. Units are day.
    PollingFrequency interface{}

    // The list of upcoming leap second updates. The type is slice of
    // Ptp_LeapSeconds_LeapSecond.
    LeapSecond []Ptp_LeapSeconds_LeapSecond
}

func (leapSeconds *Ptp_LeapSeconds) GetFilter() yfilter.YFilter { return leapSeconds.YFilter }

func (leapSeconds *Ptp_LeapSeconds) SetFilter(yf yfilter.YFilter) { leapSeconds.YFilter = yf }

func (leapSeconds *Ptp_LeapSeconds) GetGoName(yname string) string {
    if yname == "current-offset" { return "CurrentOffset" }
    if yname == "offset-valid" { return "OffsetValid" }
    if yname == "source-file" { return "SourceFile" }
    if yname == "source-expiry-date" { return "SourceExpiryDate" }
    if yname == "polling-frequency" { return "PollingFrequency" }
    if yname == "leap-second" { return "LeapSecond" }
    return ""
}

func (leapSeconds *Ptp_LeapSeconds) GetSegmentPath() string {
    return "leap-seconds"
}

func (leapSeconds *Ptp_LeapSeconds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "leap-second" {
        for _, c := range leapSeconds.LeapSecond {
            if leapSeconds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_LeapSeconds_LeapSecond{}
        leapSeconds.LeapSecond = append(leapSeconds.LeapSecond, child)
        return &leapSeconds.LeapSecond[len(leapSeconds.LeapSecond)-1]
    }
    return nil
}

func (leapSeconds *Ptp_LeapSeconds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range leapSeconds.LeapSecond {
        children[leapSeconds.LeapSecond[i].GetSegmentPath()] = &leapSeconds.LeapSecond[i]
    }
    return children
}

func (leapSeconds *Ptp_LeapSeconds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-offset"] = leapSeconds.CurrentOffset
    leafs["offset-valid"] = leapSeconds.OffsetValid
    leafs["source-file"] = leapSeconds.SourceFile
    leafs["source-expiry-date"] = leapSeconds.SourceExpiryDate
    leafs["polling-frequency"] = leapSeconds.PollingFrequency
    return leafs
}

func (leapSeconds *Ptp_LeapSeconds) GetBundleName() string { return "cisco_ios_xr" }

func (leapSeconds *Ptp_LeapSeconds) GetYangName() string { return "leap-seconds" }

func (leapSeconds *Ptp_LeapSeconds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leapSeconds *Ptp_LeapSeconds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leapSeconds *Ptp_LeapSeconds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leapSeconds *Ptp_LeapSeconds) SetParent(parent types.Entity) { leapSeconds.parent = parent }

func (leapSeconds *Ptp_LeapSeconds) GetParent() types.Entity { return leapSeconds.parent }

func (leapSeconds *Ptp_LeapSeconds) GetParentYangName() string { return "ptp" }

// Ptp_LeapSeconds_LeapSecond
// The list of upcoming leap second updates
type Ptp_LeapSeconds_LeapSecond struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // The UNIX timestamp at which the offset becomes valid. The type is
    // interface{} with range: 0..18446744073709551615.
    OffsetStartDate interface{}

    // Indicates whether the offset has been applied. The type is bool.
    OffsetApplied interface{}
}

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetFilter() yfilter.YFilter { return leapSecond.YFilter }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) SetFilter(yf yfilter.YFilter) { leapSecond.YFilter = yf }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetGoName(yname string) string {
    if yname == "offset" { return "Offset" }
    if yname == "offset-start-date" { return "OffsetStartDate" }
    if yname == "offset-applied" { return "OffsetApplied" }
    return ""
}

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetSegmentPath() string {
    return "leap-second"
}

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["offset"] = leapSecond.Offset
    leafs["offset-start-date"] = leapSecond.OffsetStartDate
    leafs["offset-applied"] = leapSecond.OffsetApplied
    return leafs
}

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetBundleName() string { return "cisco_ios_xr" }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetYangName() string { return "leap-second" }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) SetParent(parent types.Entity) { leapSecond.parent = parent }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetParent() types.Entity { return leapSecond.parent }

func (leapSecond *Ptp_LeapSeconds_LeapSecond) GetParentYangName() string { return "leap-seconds" }

// Ptp_Interfaces
// Table for interface operational data
type Ptp_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface operational data. The type is slice of Ptp_Interfaces_Interface.
    Interface []Ptp_Interfaces_Interface
}

func (interfaces *Ptp_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Ptp_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Ptp_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Ptp_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Ptp_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Ptp_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Ptp_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Ptp_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Ptp_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Ptp_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Ptp_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Ptp_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Ptp_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Ptp_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Ptp_Interfaces) GetParentYangName() string { return "ptp" }

// Ptp_Interfaces_Interface
// Interface operational data
type Ptp_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Port state. The type is PtpBagPortState.
    PortState interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Line state. The type is interface{} with range: 0..4294967295.
    LineState interface{}

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Ipv6 address, if IPv6 encapsulation is being used. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // IPv4 address, if IPv4 encapsulation is being used. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Two step delay-request mechanism is being used. The type is bool.
    TwoStep interface{}

    // Communication model configured on the interface. The type is
    // PtpBagCommunicationModel.
    CommunicationModel interface{}

    // Log of the interface's sync interval. The type is interface{} with range:
    // -2147483648..2147483647.
    LogSyncInterval interface{}

    // Log of the interface's announce interval. The type is interface{} with
    // range: -2147483648..2147483647.
    LogAnnounceInterval interface{}

    // Announce timeout. The type is interface{} with range: 0..4294967295.
    AnnounceTimeout interface{}

    // Log of the interface's Minimum delay-request interval. The type is
    // interface{} with range: -2147483648..2147483647.
    LogMinDelayRequestInterval interface{}

    // The configured port state. The type is PtpBagRestrictPortState.
    ConfiguredPortState interface{}

    // The interface supports one-step operation. The type is bool.
    SupportsOneStep interface{}

    // The interface supports two-step operation. The type is bool.
    SupportsTwoStep interface{}

    // The interface supports ethernet transport. The type is bool.
    SupportsEthernet interface{}

    // The interface supports multicast. The type is bool.
    SupportsMulticast interface{}

    // The interface supports IPv6 transport. The type is bool.
    SupportsIpv6 interface{}

    // The interface supports operation in slave mode. The type is bool.
    SupportsSlave interface{}

    // The interface supports source ip configuration. The type is bool.
    SupportsSourceIp interface{}

    // The maximum rate of sync packets on the interface. The type is interface{}
    // with range: 0..255.
    MaxSyncRate interface{}

    // The class of service used on the interface for event messages. The type is
    // interface{} with range: 0..4294967295.
    EventCos interface{}

    // The class of service used on the interface for general messages. The type
    // is interface{} with range: 0..4294967295.
    GeneralCos interface{}

    // The DSCP class used on the interface for event messages. The type is
    // interface{} with range: 0..4294967295.
    EventDscp interface{}

    // The DSCP class used on the interface for general messages. The type is
    // interface{} with range: 0..4294967295.
    GeneralDscp interface{}

    // The number of unicast peers known by the interface. The type is interface{}
    // with range: 0..4294967295.
    UnicastPeers interface{}

    // Local priority, for the G.8275.1 PTP profile. The type is interface{} with
    // range: 0..255.
    LocalPriority interface{}

    // Signal fail status of the interface. The type is bool.
    SignalFail interface{}

    // MAC address, if Ethernet encapsulation is being used.
    MacAddress Ptp_Interfaces_Interface_MacAddress

    // The interface's master table. The type is slice of
    // Ptp_Interfaces_Interface_MasterTable.
    MasterTable []Ptp_Interfaces_Interface_MasterTable
}

func (self *Ptp_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ptp_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ptp_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "port-state" { return "PortState" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "line-state" { return "LineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-step" { return "TwoStep" }
    if yname == "communication-model" { return "CommunicationModel" }
    if yname == "log-sync-interval" { return "LogSyncInterval" }
    if yname == "log-announce-interval" { return "LogAnnounceInterval" }
    if yname == "announce-timeout" { return "AnnounceTimeout" }
    if yname == "log-min-delay-request-interval" { return "LogMinDelayRequestInterval" }
    if yname == "configured-port-state" { return "ConfiguredPortState" }
    if yname == "supports-one-step" { return "SupportsOneStep" }
    if yname == "supports-two-step" { return "SupportsTwoStep" }
    if yname == "supports-ethernet" { return "SupportsEthernet" }
    if yname == "supports-multicast" { return "SupportsMulticast" }
    if yname == "supports-ipv6" { return "SupportsIpv6" }
    if yname == "supports-slave" { return "SupportsSlave" }
    if yname == "supports-source-ip" { return "SupportsSourceIp" }
    if yname == "max-sync-rate" { return "MaxSyncRate" }
    if yname == "event-cos" { return "EventCos" }
    if yname == "general-cos" { return "GeneralCos" }
    if yname == "event-dscp" { return "EventDscp" }
    if yname == "general-dscp" { return "GeneralDscp" }
    if yname == "unicast-peers" { return "UnicastPeers" }
    if yname == "local-priority" { return "LocalPriority" }
    if yname == "signal-fail" { return "SignalFail" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "master-table" { return "MasterTable" }
    return ""
}

func (self *Ptp_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Ptp_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &self.MacAddress
    }
    if childYangName == "master-table" {
        for _, c := range self.MasterTable {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Interfaces_Interface_MasterTable{}
        self.MasterTable = append(self.MasterTable, child)
        return &self.MasterTable[len(self.MasterTable)-1]
    }
    return nil
}

func (self *Ptp_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &self.MacAddress
    for i := range self.MasterTable {
        children[self.MasterTable[i].GetSegmentPath()] = &self.MasterTable[i]
    }
    return children
}

func (self *Ptp_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["port-state"] = self.PortState
    leafs["port-number"] = self.PortNumber
    leafs["line-state"] = self.LineState
    leafs["encapsulation"] = self.Encapsulation
    leafs["ipv6-address"] = self.Ipv6Address
    leafs["ipv4-address"] = self.Ipv4Address
    leafs["two-step"] = self.TwoStep
    leafs["communication-model"] = self.CommunicationModel
    leafs["log-sync-interval"] = self.LogSyncInterval
    leafs["log-announce-interval"] = self.LogAnnounceInterval
    leafs["announce-timeout"] = self.AnnounceTimeout
    leafs["log-min-delay-request-interval"] = self.LogMinDelayRequestInterval
    leafs["configured-port-state"] = self.ConfiguredPortState
    leafs["supports-one-step"] = self.SupportsOneStep
    leafs["supports-two-step"] = self.SupportsTwoStep
    leafs["supports-ethernet"] = self.SupportsEthernet
    leafs["supports-multicast"] = self.SupportsMulticast
    leafs["supports-ipv6"] = self.SupportsIpv6
    leafs["supports-slave"] = self.SupportsSlave
    leafs["supports-source-ip"] = self.SupportsSourceIp
    leafs["max-sync-rate"] = self.MaxSyncRate
    leafs["event-cos"] = self.EventCos
    leafs["general-cos"] = self.GeneralCos
    leafs["event-dscp"] = self.EventDscp
    leafs["general-dscp"] = self.GeneralDscp
    leafs["unicast-peers"] = self.UnicastPeers
    leafs["local-priority"] = self.LocalPriority
    leafs["signal-fail"] = self.SignalFail
    return leafs
}

func (self *Ptp_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ptp_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Ptp_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ptp_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ptp_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ptp_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ptp_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Ptp_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Ptp_Interfaces_Interface_MacAddress
// MAC address, if Ethernet encapsulation is being
// used
type Ptp_Interfaces_Interface_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetParentYangName() string { return "interface" }

// Ptp_Interfaces_Interface_MasterTable
// The interface's master table
type Ptp_Interfaces_Interface_MasterTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The configured communication model of the master clock. The type is
    // PtpBagCommunicationModel.
    CommunicationModel interface{}

    // The priority of the master clock, if it is set. The type is interface{}
    // with range: 0..255.
    Priority interface{}

    // Whether the interface is receiving messages from this master. The type is
    // bool.
    Known interface{}

    // The master is qualified for best master clock selection. The type is bool.
    Qualified interface{}

    // Whether this is the grandmaster. The type is bool.
    IsGrandmaster interface{}

    // Announced messages are not being received from the master. The type is
    // interface{} with range: 0..255.
    PtsfLossAnnounce interface{}

    // Sync messages are not being received from the master. The type is
    // interface{} with range: 0..255.
    PtsfLossSync interface{}

    // Whether this master uses non-negotiated unicast. The type is bool.
    IsNonnegotiated interface{}

    // The address of the master clock.
    Address Ptp_Interfaces_Interface_MasterTable_Address
}

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetFilter() yfilter.YFilter { return masterTable.YFilter }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) SetFilter(yf yfilter.YFilter) { masterTable.YFilter = yf }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetGoName(yname string) string {
    if yname == "communication-model" { return "CommunicationModel" }
    if yname == "priority" { return "Priority" }
    if yname == "known" { return "Known" }
    if yname == "qualified" { return "Qualified" }
    if yname == "is-grandmaster" { return "IsGrandmaster" }
    if yname == "ptsf-loss-announce" { return "PtsfLossAnnounce" }
    if yname == "ptsf-loss-sync" { return "PtsfLossSync" }
    if yname == "is-nonnegotiated" { return "IsNonnegotiated" }
    if yname == "address" { return "Address" }
    return ""
}

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetSegmentPath() string {
    return "master-table"
}

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &masterTable.Address
    }
    return nil
}

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &masterTable.Address
    return children
}

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["communication-model"] = masterTable.CommunicationModel
    leafs["priority"] = masterTable.Priority
    leafs["known"] = masterTable.Known
    leafs["qualified"] = masterTable.Qualified
    leafs["is-grandmaster"] = masterTable.IsGrandmaster
    leafs["ptsf-loss-announce"] = masterTable.PtsfLossAnnounce
    leafs["ptsf-loss-sync"] = masterTable.PtsfLossSync
    leafs["is-nonnegotiated"] = masterTable.IsNonnegotiated
    return leafs
}

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetBundleName() string { return "cisco_ios_xr" }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetYangName() string { return "master-table" }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) SetParent(parent types.Entity) { masterTable.parent = parent }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetParent() types.Entity { return masterTable.parent }

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetParentYangName() string { return "interface" }

// Ptp_Interfaces_Interface_MasterTable_Address
// The address of the master clock
type Ptp_Interfaces_Interface_MasterTable_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Interfaces_Interface_MasterTable_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetYangName() string { return "address" }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetParentYangName() string { return "master-table" }

// Ptp_Interfaces_Interface_MasterTable_Address_MacAddress
// Ethernet MAC address
type Ptp_Interfaces_Interface_MasterTable_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address
// IPv6 address
type Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_Dataset
// Global PTP datasets
type Ptp_Dataset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // defaultDS information as described in IEEE 1588-2008.
    DefaultDs Ptp_Dataset_DefaultDs

    // currentDS information as described in IEEE 1588-2008.
    CurrentDs Ptp_Dataset_CurrentDs

    // parentDS information as described in IEEE 1588-2008.
    ParentDs Ptp_Dataset_ParentDs

    // Table for portDS information.
    PortDses Ptp_Dataset_PortDses

    // timePropertiesDS information as described in IEEE 1588-2008.
    TimePropertiesDs Ptp_Dataset_TimePropertiesDs
}

func (dataset *Ptp_Dataset) GetFilter() yfilter.YFilter { return dataset.YFilter }

func (dataset *Ptp_Dataset) SetFilter(yf yfilter.YFilter) { dataset.YFilter = yf }

func (dataset *Ptp_Dataset) GetGoName(yname string) string {
    if yname == "default-ds" { return "DefaultDs" }
    if yname == "current-ds" { return "CurrentDs" }
    if yname == "parent-ds" { return "ParentDs" }
    if yname == "port-dses" { return "PortDses" }
    if yname == "time-properties-ds" { return "TimePropertiesDs" }
    return ""
}

func (dataset *Ptp_Dataset) GetSegmentPath() string {
    return "dataset"
}

func (dataset *Ptp_Dataset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-ds" {
        return &dataset.DefaultDs
    }
    if childYangName == "current-ds" {
        return &dataset.CurrentDs
    }
    if childYangName == "parent-ds" {
        return &dataset.ParentDs
    }
    if childYangName == "port-dses" {
        return &dataset.PortDses
    }
    if childYangName == "time-properties-ds" {
        return &dataset.TimePropertiesDs
    }
    return nil
}

func (dataset *Ptp_Dataset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-ds"] = &dataset.DefaultDs
    children["current-ds"] = &dataset.CurrentDs
    children["parent-ds"] = &dataset.ParentDs
    children["port-dses"] = &dataset.PortDses
    children["time-properties-ds"] = &dataset.TimePropertiesDs
    return children
}

func (dataset *Ptp_Dataset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataset *Ptp_Dataset) GetBundleName() string { return "cisco_ios_xr" }

func (dataset *Ptp_Dataset) GetYangName() string { return "dataset" }

func (dataset *Ptp_Dataset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataset *Ptp_Dataset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataset *Ptp_Dataset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataset *Ptp_Dataset) SetParent(parent types.Entity) { dataset.parent = parent }

func (dataset *Ptp_Dataset) GetParent() types.Entity { return dataset.parent }

func (dataset *Ptp_Dataset) GetParentYangName() string { return "ptp" }

// Ptp_Dataset_DefaultDs
// defaultDS information as described in IEEE
// 1588-2008
type Ptp_Dataset_DefaultDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is the twoStepFlag set for this clock. The type is bool.
    TwoStepFlag interface{}

    // The local-clock ID. The type is interface{} with range:
    // 0..18446744073709551615.
    ClockId interface{}

    // The number of active PTP ports on this clock. The type is interface{} with
    // range: 0..4294967295.
    NumberPorts interface{}

    // The clock class of the local-clock. The type is interface{} with range:
    // 0..255.
    ClockClass interface{}

    // The accuracy of the local-clock. The type is interface{} with range:
    // 0..255.
    ClockAccuracy interface{}

    // The offset scaled log variance of the local-clock. The type is interface{}
    // with range: 0..65535.
    Oslv interface{}

    // The priority1 of the local-clock. The type is interface{} with range:
    // 0..255.
    Priority1 interface{}

    // The priority2 of the local-clock. The type is interface{} with range:
    // 0..255.
    Priority2 interface{}

    // The domain of the local-clock. The type is interface{} with range: 0..255.
    DomainNumber interface{}

    // Whether the local-clock is globally configured as slave-only. The type is
    // bool.
    SlaveOnly interface{}

    // Local priority of the local clock. The type is interface{} with range:
    // 0..4294967295.
    LocalPriority interface{}

    // Signal fail status of the local clock. The type is bool.
    SignalFail interface{}
}

func (defaultDs *Ptp_Dataset_DefaultDs) GetFilter() yfilter.YFilter { return defaultDs.YFilter }

func (defaultDs *Ptp_Dataset_DefaultDs) SetFilter(yf yfilter.YFilter) { defaultDs.YFilter = yf }

func (defaultDs *Ptp_Dataset_DefaultDs) GetGoName(yname string) string {
    if yname == "two-step-flag" { return "TwoStepFlag" }
    if yname == "clock-id" { return "ClockId" }
    if yname == "number-ports" { return "NumberPorts" }
    if yname == "clock-class" { return "ClockClass" }
    if yname == "clock-accuracy" { return "ClockAccuracy" }
    if yname == "oslv" { return "Oslv" }
    if yname == "priority1" { return "Priority1" }
    if yname == "priority2" { return "Priority2" }
    if yname == "domain-number" { return "DomainNumber" }
    if yname == "slave-only" { return "SlaveOnly" }
    if yname == "local-priority" { return "LocalPriority" }
    if yname == "signal-fail" { return "SignalFail" }
    return ""
}

func (defaultDs *Ptp_Dataset_DefaultDs) GetSegmentPath() string {
    return "default-ds"
}

func (defaultDs *Ptp_Dataset_DefaultDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultDs *Ptp_Dataset_DefaultDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultDs *Ptp_Dataset_DefaultDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-step-flag"] = defaultDs.TwoStepFlag
    leafs["clock-id"] = defaultDs.ClockId
    leafs["number-ports"] = defaultDs.NumberPorts
    leafs["clock-class"] = defaultDs.ClockClass
    leafs["clock-accuracy"] = defaultDs.ClockAccuracy
    leafs["oslv"] = defaultDs.Oslv
    leafs["priority1"] = defaultDs.Priority1
    leafs["priority2"] = defaultDs.Priority2
    leafs["domain-number"] = defaultDs.DomainNumber
    leafs["slave-only"] = defaultDs.SlaveOnly
    leafs["local-priority"] = defaultDs.LocalPriority
    leafs["signal-fail"] = defaultDs.SignalFail
    return leafs
}

func (defaultDs *Ptp_Dataset_DefaultDs) GetBundleName() string { return "cisco_ios_xr" }

func (defaultDs *Ptp_Dataset_DefaultDs) GetYangName() string { return "default-ds" }

func (defaultDs *Ptp_Dataset_DefaultDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultDs *Ptp_Dataset_DefaultDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultDs *Ptp_Dataset_DefaultDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultDs *Ptp_Dataset_DefaultDs) SetParent(parent types.Entity) { defaultDs.parent = parent }

func (defaultDs *Ptp_Dataset_DefaultDs) GetParent() types.Entity { return defaultDs.parent }

func (defaultDs *Ptp_Dataset_DefaultDs) GetParentYangName() string { return "dataset" }

// Ptp_Dataset_CurrentDs
// currentDS information as described in IEEE
// 1588-2008
type Ptp_Dataset_CurrentDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // How many steps removed this clock is from the GM. The type is interface{}
    // with range: 0..65535.
    StepsRemoved interface{}

    // The UTC offset of the local-clock from the GM. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    OffsetFromMaster interface{}

    // The mean path delay bewteen the foreign-master and the local-clock. The
    // type is interface{} with range: -9223372036854775808..9223372036854775807.
    MeanPathDelay interface{}
}

func (currentDs *Ptp_Dataset_CurrentDs) GetFilter() yfilter.YFilter { return currentDs.YFilter }

func (currentDs *Ptp_Dataset_CurrentDs) SetFilter(yf yfilter.YFilter) { currentDs.YFilter = yf }

func (currentDs *Ptp_Dataset_CurrentDs) GetGoName(yname string) string {
    if yname == "steps-removed" { return "StepsRemoved" }
    if yname == "offset-from-master" { return "OffsetFromMaster" }
    if yname == "mean-path-delay" { return "MeanPathDelay" }
    return ""
}

func (currentDs *Ptp_Dataset_CurrentDs) GetSegmentPath() string {
    return "current-ds"
}

func (currentDs *Ptp_Dataset_CurrentDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currentDs *Ptp_Dataset_CurrentDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currentDs *Ptp_Dataset_CurrentDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["steps-removed"] = currentDs.StepsRemoved
    leafs["offset-from-master"] = currentDs.OffsetFromMaster
    leafs["mean-path-delay"] = currentDs.MeanPathDelay
    return leafs
}

func (currentDs *Ptp_Dataset_CurrentDs) GetBundleName() string { return "cisco_ios_xr" }

func (currentDs *Ptp_Dataset_CurrentDs) GetYangName() string { return "current-ds" }

func (currentDs *Ptp_Dataset_CurrentDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currentDs *Ptp_Dataset_CurrentDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currentDs *Ptp_Dataset_CurrentDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currentDs *Ptp_Dataset_CurrentDs) SetParent(parent types.Entity) { currentDs.parent = parent }

func (currentDs *Ptp_Dataset_CurrentDs) GetParent() types.Entity { return currentDs.parent }

func (currentDs *Ptp_Dataset_CurrentDs) GetParentYangName() string { return "dataset" }

// Ptp_Dataset_ParentDs
// parentDS information as described in IEEE
// 1588-2008
type Ptp_Dataset_ParentDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Clock ID of the parent clock. The type is interface{} with range:
    // 0..18446744073709551615.
    ParentClockId interface{}

    // The port number on which the parent clock is received. The type is
    // interface{} with range: 0..65535.
    ParentPortNumber interface{}

    // Whether the parentStats is valid. The type is bool.
    ParentStats interface{}

    // The observed parent offset scaled log variance. The type is interface{}
    // with range: 0..65535.
    ObservedParentOslv interface{}

    // The observed rate of change of phase of the parent clock. The type is
    // interface{} with range: 0..4294967295.
    ObservedParentClockPhaseChangeRate interface{}

    // The Clock ID of the GM. The type is interface{} with range:
    // 0..18446744073709551615.
    GmClockId interface{}

    // The clock class of the GM. The type is interface{} with range: 0..255.
    GmClockClass interface{}

    // The clock accuracy of the GM. The type is interface{} with range: 0..255.
    GmClockAccuracy interface{}

    // The offset scaled log variance of the GM. The type is interface{} with
    // range: 0..65535.
    Gmoslv interface{}

    // The priority1 of the GM. The type is interface{} with range: 0..255.
    GmPriority1 interface{}

    // The priority2 of the GM. The type is interface{} with range: 0..255.
    GmPriority2 interface{}
}

func (parentDs *Ptp_Dataset_ParentDs) GetFilter() yfilter.YFilter { return parentDs.YFilter }

func (parentDs *Ptp_Dataset_ParentDs) SetFilter(yf yfilter.YFilter) { parentDs.YFilter = yf }

func (parentDs *Ptp_Dataset_ParentDs) GetGoName(yname string) string {
    if yname == "parent-clock-id" { return "ParentClockId" }
    if yname == "parent-port-number" { return "ParentPortNumber" }
    if yname == "parent-stats" { return "ParentStats" }
    if yname == "observed-parent-oslv" { return "ObservedParentOslv" }
    if yname == "observed-parent-clock-phase-change-rate" { return "ObservedParentClockPhaseChangeRate" }
    if yname == "gm-clock-id" { return "GmClockId" }
    if yname == "gm-clock-class" { return "GmClockClass" }
    if yname == "gm-clock-accuracy" { return "GmClockAccuracy" }
    if yname == "gmoslv" { return "Gmoslv" }
    if yname == "gm-priority1" { return "GmPriority1" }
    if yname == "gm-priority2" { return "GmPriority2" }
    return ""
}

func (parentDs *Ptp_Dataset_ParentDs) GetSegmentPath() string {
    return "parent-ds"
}

func (parentDs *Ptp_Dataset_ParentDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (parentDs *Ptp_Dataset_ParentDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (parentDs *Ptp_Dataset_ParentDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["parent-clock-id"] = parentDs.ParentClockId
    leafs["parent-port-number"] = parentDs.ParentPortNumber
    leafs["parent-stats"] = parentDs.ParentStats
    leafs["observed-parent-oslv"] = parentDs.ObservedParentOslv
    leafs["observed-parent-clock-phase-change-rate"] = parentDs.ObservedParentClockPhaseChangeRate
    leafs["gm-clock-id"] = parentDs.GmClockId
    leafs["gm-clock-class"] = parentDs.GmClockClass
    leafs["gm-clock-accuracy"] = parentDs.GmClockAccuracy
    leafs["gmoslv"] = parentDs.Gmoslv
    leafs["gm-priority1"] = parentDs.GmPriority1
    leafs["gm-priority2"] = parentDs.GmPriority2
    return leafs
}

func (parentDs *Ptp_Dataset_ParentDs) GetBundleName() string { return "cisco_ios_xr" }

func (parentDs *Ptp_Dataset_ParentDs) GetYangName() string { return "parent-ds" }

func (parentDs *Ptp_Dataset_ParentDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parentDs *Ptp_Dataset_ParentDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parentDs *Ptp_Dataset_ParentDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parentDs *Ptp_Dataset_ParentDs) SetParent(parent types.Entity) { parentDs.parent = parent }

func (parentDs *Ptp_Dataset_ParentDs) GetParent() types.Entity { return parentDs.parent }

func (parentDs *Ptp_Dataset_ParentDs) GetParentYangName() string { return "dataset" }

// Ptp_Dataset_PortDses
// Table for portDS information
type Ptp_Dataset_PortDses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PortDS information. The type is slice of Ptp_Dataset_PortDses_PortDs.
    PortDs []Ptp_Dataset_PortDses_PortDs
}

func (portDses *Ptp_Dataset_PortDses) GetFilter() yfilter.YFilter { return portDses.YFilter }

func (portDses *Ptp_Dataset_PortDses) SetFilter(yf yfilter.YFilter) { portDses.YFilter = yf }

func (portDses *Ptp_Dataset_PortDses) GetGoName(yname string) string {
    if yname == "port-ds" { return "PortDs" }
    return ""
}

func (portDses *Ptp_Dataset_PortDses) GetSegmentPath() string {
    return "port-dses"
}

func (portDses *Ptp_Dataset_PortDses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-ds" {
        for _, c := range portDses.PortDs {
            if portDses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_Dataset_PortDses_PortDs{}
        portDses.PortDs = append(portDses.PortDs, child)
        return &portDses.PortDs[len(portDses.PortDs)-1]
    }
    return nil
}

func (portDses *Ptp_Dataset_PortDses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portDses.PortDs {
        children[portDses.PortDs[i].GetSegmentPath()] = &portDses.PortDs[i]
    }
    return children
}

func (portDses *Ptp_Dataset_PortDses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portDses *Ptp_Dataset_PortDses) GetBundleName() string { return "cisco_ios_xr" }

func (portDses *Ptp_Dataset_PortDses) GetYangName() string { return "port-dses" }

func (portDses *Ptp_Dataset_PortDses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portDses *Ptp_Dataset_PortDses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portDses *Ptp_Dataset_PortDses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portDses *Ptp_Dataset_PortDses) SetParent(parent types.Entity) { portDses.parent = parent }

func (portDses *Ptp_Dataset_PortDses) GetParent() types.Entity { return portDses.parent }

func (portDses *Ptp_Dataset_PortDses) GetParentYangName() string { return "dataset" }

// Ptp_Dataset_PortDses_PortDs
// PortDS information
type Ptp_Dataset_PortDses_PortDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The ID of the local-clock. The type is interface{} with range:
    // 0..18446744073709551615.
    ClockId interface{}

    // The port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // The port state. The type is PtpBagPortState.
    PortState interface{}

    // The log (base 2) of the minimum delay-request interval. The type is
    // interface{} with range: -32768..32767.
    LogMinDelayReqInterval interface{}

    // The mean path delay between peers. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PeerMeanPathDelay interface{}

    // The log (base 2) of the announce interval. The type is interface{} with
    // range: -32768..32767.
    LogAnnounceInterval interface{}

    // The announce timeout. The type is interface{} with range: 0..4294967295.
    AnnoucneReceiptTimeout interface{}

    // The log (base 2) of the sync interval. The type is interface{} with range:
    // -32768..32767.
    LogSyncInterval interface{}

    // The delay mechanism being used on this port. The type is
    // PtpBagDelayMechanism.
    DelayMechanism interface{}

    // The log (base 2) of the minimum peer-delay-request interval. The type is
    // interface{} with range: -32768..32767.
    LogMinPDelayReqInterval interface{}

    // The version of IEEE 1588 being run. The type is interface{} with range:
    // 0..255.
    VersionNumber interface{}

    // Local priority of the local clock. The type is interface{} with range:
    // 0..4294967295.
    LocalPriority interface{}

    // Is the port master-only?. The type is bool.
    MasterOnly interface{}

    // Signal fail status of the port. The type is bool.
    SignalFail interface{}
}

func (portDs *Ptp_Dataset_PortDses_PortDs) GetFilter() yfilter.YFilter { return portDs.YFilter }

func (portDs *Ptp_Dataset_PortDses_PortDs) SetFilter(yf yfilter.YFilter) { portDs.YFilter = yf }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "port-state" { return "PortState" }
    if yname == "log-min-delay-req-interval" { return "LogMinDelayReqInterval" }
    if yname == "peer-mean-path-delay" { return "PeerMeanPathDelay" }
    if yname == "log-announce-interval" { return "LogAnnounceInterval" }
    if yname == "annoucne-receipt-timeout" { return "AnnoucneReceiptTimeout" }
    if yname == "log-sync-interval" { return "LogSyncInterval" }
    if yname == "delay-mechanism" { return "DelayMechanism" }
    if yname == "log-min-p-delay-req-interval" { return "LogMinPDelayReqInterval" }
    if yname == "version-number" { return "VersionNumber" }
    if yname == "local-priority" { return "LocalPriority" }
    if yname == "master-only" { return "MasterOnly" }
    if yname == "signal-fail" { return "SignalFail" }
    return ""
}

func (portDs *Ptp_Dataset_PortDses_PortDs) GetSegmentPath() string {
    return "port-ds" + "[interface-name='" + fmt.Sprintf("%v", portDs.InterfaceName) + "']"
}

func (portDs *Ptp_Dataset_PortDses_PortDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portDs *Ptp_Dataset_PortDses_PortDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portDs *Ptp_Dataset_PortDses_PortDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = portDs.InterfaceName
    leafs["clock-id"] = portDs.ClockId
    leafs["port-number"] = portDs.PortNumber
    leafs["port-state"] = portDs.PortState
    leafs["log-min-delay-req-interval"] = portDs.LogMinDelayReqInterval
    leafs["peer-mean-path-delay"] = portDs.PeerMeanPathDelay
    leafs["log-announce-interval"] = portDs.LogAnnounceInterval
    leafs["annoucne-receipt-timeout"] = portDs.AnnoucneReceiptTimeout
    leafs["log-sync-interval"] = portDs.LogSyncInterval
    leafs["delay-mechanism"] = portDs.DelayMechanism
    leafs["log-min-p-delay-req-interval"] = portDs.LogMinPDelayReqInterval
    leafs["version-number"] = portDs.VersionNumber
    leafs["local-priority"] = portDs.LocalPriority
    leafs["master-only"] = portDs.MasterOnly
    leafs["signal-fail"] = portDs.SignalFail
    return leafs
}

func (portDs *Ptp_Dataset_PortDses_PortDs) GetBundleName() string { return "cisco_ios_xr" }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetYangName() string { return "port-ds" }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portDs *Ptp_Dataset_PortDses_PortDs) SetParent(parent types.Entity) { portDs.parent = parent }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetParent() types.Entity { return portDs.parent }

func (portDs *Ptp_Dataset_PortDses_PortDs) GetParentYangName() string { return "port-dses" }

// Ptp_Dataset_TimePropertiesDs
// timePropertiesDS information as described in
// IEEE 1588-2008
type Ptp_Dataset_TimePropertiesDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The current UTC offset. The type is interface{} with range: -32768..32767.
    CurrentUtcOffset interface{}

    // Whether the current UTC offset is valid. The type is bool.
    CurrentUtcOffsetValid interface{}

    // Whether the last minute of the day has 59 seconds. The type is bool.
    Leap59 interface{}

    // Whether the last minute of the day has 61 seconds. The type is bool.
    Leap61 interface{}

    // Whether the time-of-day source is traceable. The type is bool.
    TimeTraceable interface{}

    // Whther the frequency source is traceable. The type is bool.
    FrequencyTraceable interface{}

    // Whether the timescale being used is the PTP timescale. The type is bool.
    PtpTimescale interface{}

    // The physical time-source of the GM clock. The type is
    // PtpBagClockTimeSource.
    TimeSource interface{}
}

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetFilter() yfilter.YFilter { return timePropertiesDs.YFilter }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) SetFilter(yf yfilter.YFilter) { timePropertiesDs.YFilter = yf }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetGoName(yname string) string {
    if yname == "current-utc-offset" { return "CurrentUtcOffset" }
    if yname == "current-utc-offset-valid" { return "CurrentUtcOffsetValid" }
    if yname == "leap59" { return "Leap59" }
    if yname == "leap61" { return "Leap61" }
    if yname == "time-traceable" { return "TimeTraceable" }
    if yname == "frequency-traceable" { return "FrequencyTraceable" }
    if yname == "ptp-timescale" { return "PtpTimescale" }
    if yname == "time-source" { return "TimeSource" }
    return ""
}

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetSegmentPath() string {
    return "time-properties-ds"
}

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-utc-offset"] = timePropertiesDs.CurrentUtcOffset
    leafs["current-utc-offset-valid"] = timePropertiesDs.CurrentUtcOffsetValid
    leafs["leap59"] = timePropertiesDs.Leap59
    leafs["leap61"] = timePropertiesDs.Leap61
    leafs["time-traceable"] = timePropertiesDs.TimeTraceable
    leafs["frequency-traceable"] = timePropertiesDs.FrequencyTraceable
    leafs["ptp-timescale"] = timePropertiesDs.PtpTimescale
    leafs["time-source"] = timePropertiesDs.TimeSource
    return leafs
}

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetBundleName() string { return "cisco_ios_xr" }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetYangName() string { return "time-properties-ds" }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) SetParent(parent types.Entity) { timePropertiesDs.parent = parent }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetParent() types.Entity { return timePropertiesDs.parent }

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetParentYangName() string { return "dataset" }

// Ptp_GlobalConfigurationError
// Global configuration error operational data
type Ptp_GlobalConfigurationError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The clock profile. The type is PtpBagProfile.
    ClockProfile interface{}

    // Is the clock profile set. The type is bool.
    ClockProfileSet interface{}

    // Configured telecom clock type. The type is PtpBagTelecomClock.
    TelecomClockType interface{}

    // The PTP domain. The type is interface{} with range: 0..255.
    DomainNumber interface{}

    // The configured priority2 value. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Configuration Errors.
    ConfigurationErrors Ptp_GlobalConfigurationError_ConfigurationErrors
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetFilter() yfilter.YFilter { return globalConfigurationError.YFilter }

func (globalConfigurationError *Ptp_GlobalConfigurationError) SetFilter(yf yfilter.YFilter) { globalConfigurationError.YFilter = yf }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetGoName(yname string) string {
    if yname == "clock-profile" { return "ClockProfile" }
    if yname == "clock-profile-set" { return "ClockProfileSet" }
    if yname == "telecom-clock-type" { return "TelecomClockType" }
    if yname == "domain-number" { return "DomainNumber" }
    if yname == "priority2" { return "Priority2" }
    if yname == "configuration-errors" { return "ConfigurationErrors" }
    return ""
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetSegmentPath() string {
    return "global-configuration-error"
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configuration-errors" {
        return &globalConfigurationError.ConfigurationErrors
    }
    return nil
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configuration-errors"] = &globalConfigurationError.ConfigurationErrors
    return children
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-profile"] = globalConfigurationError.ClockProfile
    leafs["clock-profile-set"] = globalConfigurationError.ClockProfileSet
    leafs["telecom-clock-type"] = globalConfigurationError.TelecomClockType
    leafs["domain-number"] = globalConfigurationError.DomainNumber
    leafs["priority2"] = globalConfigurationError.Priority2
    return leafs
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetBundleName() string { return "cisco_ios_xr" }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetYangName() string { return "global-configuration-error" }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalConfigurationError *Ptp_GlobalConfigurationError) SetParent(parent types.Entity) { globalConfigurationError.parent = parent }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetParent() types.Entity { return globalConfigurationError.parent }

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetParentYangName() string { return "ptp" }

// Ptp_GlobalConfigurationError_ConfigurationErrors
// Configuration Errors
type Ptp_GlobalConfigurationError_ConfigurationErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain not compatible with configured profile. The type is bool.
    Domain interface{}

    // Priority1 configuration is not compatible with configured profile. The type
    // is bool.
    ProfilePriority1Config interface{}

    // Priority2 value is not compatible with configured profile. The type is
    // bool.
    ProfilePriority2Value interface{}
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetFilter() yfilter.YFilter { return configurationErrors.YFilter }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) SetFilter(yf yfilter.YFilter) { configurationErrors.YFilter = yf }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "profile-priority1-config" { return "ProfilePriority1Config" }
    if yname == "profile-priority2-value" { return "ProfilePriority2Value" }
    return ""
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetSegmentPath() string {
    return "configuration-errors"
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = configurationErrors.Domain
    leafs["profile-priority1-config"] = configurationErrors.ProfilePriority1Config
    leafs["profile-priority2-value"] = configurationErrors.ProfilePriority2Value
    return leafs
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetBundleName() string { return "cisco_ios_xr" }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetYangName() string { return "configuration-errors" }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) SetParent(parent types.Entity) { configurationErrors.parent = parent }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetParent() types.Entity { return configurationErrors.parent }

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetParentYangName() string { return "global-configuration-error" }

// Ptp_Grandmaster
// Grandmaster clock operational data
type Ptp_Grandmaster struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the grandmaster is setting time-of-day on the system. The type is
    // bool.
    UsedForTime interface{}

    // Whether the grandmaster is setting frequency on the system. The type is
    // bool.
    UsedForFrequency interface{}

    // How long the clock has been grandmaster for, in seconds. The type is
    // interface{} with range: 0..4294967295. Units are second.
    KnownForTime interface{}

    // The PTP domain that the grandmaster is in. The type is interface{} with
    // range: 0..255.
    Domain interface{}

    // Grandmaster clock.
    ClockProperties Ptp_Grandmaster_ClockProperties

    // The grandmaster's address information.
    Address Ptp_Grandmaster_Address
}

func (grandmaster *Ptp_Grandmaster) GetFilter() yfilter.YFilter { return grandmaster.YFilter }

func (grandmaster *Ptp_Grandmaster) SetFilter(yf yfilter.YFilter) { grandmaster.YFilter = yf }

func (grandmaster *Ptp_Grandmaster) GetGoName(yname string) string {
    if yname == "used-for-time" { return "UsedForTime" }
    if yname == "used-for-frequency" { return "UsedForFrequency" }
    if yname == "known-for-time" { return "KnownForTime" }
    if yname == "domain" { return "Domain" }
    if yname == "clock-properties" { return "ClockProperties" }
    if yname == "address" { return "Address" }
    return ""
}

func (grandmaster *Ptp_Grandmaster) GetSegmentPath() string {
    return "grandmaster"
}

func (grandmaster *Ptp_Grandmaster) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-properties" {
        return &grandmaster.ClockProperties
    }
    if childYangName == "address" {
        return &grandmaster.Address
    }
    return nil
}

func (grandmaster *Ptp_Grandmaster) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-properties"] = &grandmaster.ClockProperties
    children["address"] = &grandmaster.Address
    return children
}

func (grandmaster *Ptp_Grandmaster) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["used-for-time"] = grandmaster.UsedForTime
    leafs["used-for-frequency"] = grandmaster.UsedForFrequency
    leafs["known-for-time"] = grandmaster.KnownForTime
    leafs["domain"] = grandmaster.Domain
    return leafs
}

func (grandmaster *Ptp_Grandmaster) GetBundleName() string { return "cisco_ios_xr" }

func (grandmaster *Ptp_Grandmaster) GetYangName() string { return "grandmaster" }

func (grandmaster *Ptp_Grandmaster) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (grandmaster *Ptp_Grandmaster) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (grandmaster *Ptp_Grandmaster) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (grandmaster *Ptp_Grandmaster) SetParent(parent types.Entity) { grandmaster.parent = parent }

func (grandmaster *Ptp_Grandmaster) GetParent() types.Entity { return grandmaster.parent }

func (grandmaster *Ptp_Grandmaster) GetParentYangName() string { return "ptp" }

// Ptp_Grandmaster_ClockProperties
// Grandmaster clock
type Ptp_Grandmaster_ClockProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Class. The type is interface{} with range: 0..255.
    Class interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Steps removed. The type is interface{} with range: 0..65535.
    StepsRemoved interface{}

    // Time source. The type is PtpBagClockTimeSource.
    TimeSource interface{}

    // The clock is frequency traceable. The type is bool.
    FrequencyTraceable interface{}

    // The clock is time traceable. The type is bool.
    TimeTraceable interface{}

    // Timescale. The type is PtpBagClockTimescale.
    Timescale interface{}

    // Leap Seconds. The type is PtpBagClockLeapSeconds. Units are second.
    LeapSeconds interface{}

    // The clock is the local clock. The type is bool.
    Local interface{}

    // The configured clock class. The type is interface{} with range: 0..255.
    ConfiguredClockClass interface{}

    // The configured priority. The type is interface{} with range: 0..255.
    ConfiguredPriority interface{}

    // UTC offset.
    UtcOffset Ptp_Grandmaster_ClockProperties_UtcOffset

    // Receiver.
    Receiver Ptp_Grandmaster_ClockProperties_Receiver

    // Sender.
    Sender Ptp_Grandmaster_ClockProperties_Sender
}

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetFilter() yfilter.YFilter { return clockProperties.YFilter }

func (clockProperties *Ptp_Grandmaster_ClockProperties) SetFilter(yf yfilter.YFilter) { clockProperties.YFilter = yf }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "priority1" { return "Priority1" }
    if yname == "priority2" { return "Priority2" }
    if yname == "class" { return "Class" }
    if yname == "accuracy" { return "Accuracy" }
    if yname == "offset-log-variance" { return "OffsetLogVariance" }
    if yname == "steps-removed" { return "StepsRemoved" }
    if yname == "time-source" { return "TimeSource" }
    if yname == "frequency-traceable" { return "FrequencyTraceable" }
    if yname == "time-traceable" { return "TimeTraceable" }
    if yname == "timescale" { return "Timescale" }
    if yname == "leap-seconds" { return "LeapSeconds" }
    if yname == "local" { return "Local" }
    if yname == "configured-clock-class" { return "ConfiguredClockClass" }
    if yname == "configured-priority" { return "ConfiguredPriority" }
    if yname == "utc-offset" { return "UtcOffset" }
    if yname == "receiver" { return "Receiver" }
    if yname == "sender" { return "Sender" }
    return ""
}

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetSegmentPath() string {
    return "clock-properties"
}

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "utc-offset" {
        return &clockProperties.UtcOffset
    }
    if childYangName == "receiver" {
        return &clockProperties.Receiver
    }
    if childYangName == "sender" {
        return &clockProperties.Sender
    }
    return nil
}

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["utc-offset"] = &clockProperties.UtcOffset
    children["receiver"] = &clockProperties.Receiver
    children["sender"] = &clockProperties.Sender
    return children
}

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = clockProperties.ClockId
    leafs["priority1"] = clockProperties.Priority1
    leafs["priority2"] = clockProperties.Priority2
    leafs["class"] = clockProperties.Class
    leafs["accuracy"] = clockProperties.Accuracy
    leafs["offset-log-variance"] = clockProperties.OffsetLogVariance
    leafs["steps-removed"] = clockProperties.StepsRemoved
    leafs["time-source"] = clockProperties.TimeSource
    leafs["frequency-traceable"] = clockProperties.FrequencyTraceable
    leafs["time-traceable"] = clockProperties.TimeTraceable
    leafs["timescale"] = clockProperties.Timescale
    leafs["leap-seconds"] = clockProperties.LeapSeconds
    leafs["local"] = clockProperties.Local
    leafs["configured-clock-class"] = clockProperties.ConfiguredClockClass
    leafs["configured-priority"] = clockProperties.ConfiguredPriority
    return leafs
}

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetBundleName() string { return "cisco_ios_xr" }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetYangName() string { return "clock-properties" }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockProperties *Ptp_Grandmaster_ClockProperties) SetParent(parent types.Entity) { clockProperties.parent = parent }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetParent() types.Entity { return clockProperties.parent }

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetParentYangName() string { return "grandmaster" }

// Ptp_Grandmaster_ClockProperties_UtcOffset
// UTC offset
type Ptp_Grandmaster_ClockProperties_UtcOffset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetFilter() yfilter.YFilter { return utcOffset.YFilter }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) SetFilter(yf yfilter.YFilter) { utcOffset.YFilter = yf }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetGoName(yname string) string {
    if yname == "current-offset" { return "CurrentOffset" }
    if yname == "offset-valid" { return "OffsetValid" }
    return ""
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetSegmentPath() string {
    return "utc-offset"
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-offset"] = utcOffset.CurrentOffset
    leafs["offset-valid"] = utcOffset.OffsetValid
    return leafs
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetBundleName() string { return "cisco_ios_xr" }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetYangName() string { return "utc-offset" }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) SetParent(parent types.Entity) { utcOffset.parent = parent }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetParent() types.Entity { return utcOffset.parent }

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetParentYangName() string { return "clock-properties" }

// Ptp_Grandmaster_ClockProperties_Receiver
// Receiver
type Ptp_Grandmaster_ClockProperties_Receiver struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetFilter() yfilter.YFilter { return receiver.YFilter }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) SetFilter(yf yfilter.YFilter) { receiver.YFilter = yf }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetSegmentPath() string {
    return "receiver"
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = receiver.ClockId
    leafs["port-number"] = receiver.PortNumber
    return leafs
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetBundleName() string { return "cisco_ios_xr" }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetYangName() string { return "receiver" }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) SetParent(parent types.Entity) { receiver.parent = parent }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetParent() types.Entity { return receiver.parent }

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetParentYangName() string { return "clock-properties" }

// Ptp_Grandmaster_ClockProperties_Sender
// Sender
type Ptp_Grandmaster_ClockProperties_Sender struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetFilter() yfilter.YFilter { return sender.YFilter }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) SetFilter(yf yfilter.YFilter) { sender.YFilter = yf }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetGoName(yname string) string {
    if yname == "clock-id" { return "ClockId" }
    if yname == "port-number" { return "PortNumber" }
    return ""
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetSegmentPath() string {
    return "sender"
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-id"] = sender.ClockId
    leafs["port-number"] = sender.PortNumber
    return leafs
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetBundleName() string { return "cisco_ios_xr" }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetYangName() string { return "sender" }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) SetParent(parent types.Entity) { sender.parent = parent }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetParent() types.Entity { return sender.parent }

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetParentYangName() string { return "clock-properties" }

// Ptp_Grandmaster_Address
// The grandmaster's address information
type Ptp_Grandmaster_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Grandmaster_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Grandmaster_Address_Ipv6Address
}

func (address *Ptp_Grandmaster_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_Grandmaster_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_Grandmaster_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_Grandmaster_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_Grandmaster_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_Grandmaster_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_Grandmaster_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_Grandmaster_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_Grandmaster_Address) GetYangName() string { return "address" }

func (address *Ptp_Grandmaster_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_Grandmaster_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_Grandmaster_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_Grandmaster_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_Grandmaster_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_Grandmaster_Address) GetParentYangName() string { return "grandmaster" }

// Ptp_Grandmaster_Address_MacAddress
// Ethernet MAC address
type Ptp_Grandmaster_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_Grandmaster_Address_Ipv6Address
// IPv6 address
type Ptp_Grandmaster_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_InterfaceUnicastPeers
// Table for interface unicast peers operational
// data
type Ptp_InterfaceUnicastPeers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface unicast peers operational data. The type is slice of
    // Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer.
    InterfaceUnicastPeer []Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetFilter() yfilter.YFilter { return interfaceUnicastPeers.YFilter }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) SetFilter(yf yfilter.YFilter) { interfaceUnicastPeers.YFilter = yf }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetGoName(yname string) string {
    if yname == "interface-unicast-peer" { return "InterfaceUnicastPeer" }
    return ""
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetSegmentPath() string {
    return "interface-unicast-peers"
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-unicast-peer" {
        for _, c := range interfaceUnicastPeers.InterfaceUnicastPeer {
            if interfaceUnicastPeers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer{}
        interfaceUnicastPeers.InterfaceUnicastPeer = append(interfaceUnicastPeers.InterfaceUnicastPeer, child)
        return &interfaceUnicastPeers.InterfaceUnicastPeer[len(interfaceUnicastPeers.InterfaceUnicastPeer)-1]
    }
    return nil
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceUnicastPeers.InterfaceUnicastPeer {
        children[interfaceUnicastPeers.InterfaceUnicastPeer[i].GetSegmentPath()] = &interfaceUnicastPeers.InterfaceUnicastPeer[i]
    }
    return children
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetYangName() string { return "interface-unicast-peers" }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) SetParent(parent types.Entity) { interfaceUnicastPeers.parent = parent }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetParent() types.Entity { return interfaceUnicastPeers.parent }

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetParentYangName() string { return "ptp" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer
// Interface unicast peers operational data
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Unicast Peers. The type is slice of
    // Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers.
    Peers []Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetFilter() yfilter.YFilter { return interfaceUnicastPeer.YFilter }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) SetFilter(yf yfilter.YFilter) { interfaceUnicastPeer.YFilter = yf }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "name" { return "Name" }
    if yname == "port-number" { return "PortNumber" }
    if yname == "peers" { return "Peers" }
    return ""
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetSegmentPath() string {
    return "interface-unicast-peer" + "[interface-name='" + fmt.Sprintf("%v", interfaceUnicastPeer.InterfaceName) + "']"
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peers" {
        for _, c := range interfaceUnicastPeer.Peers {
            if interfaceUnicastPeer.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers{}
        interfaceUnicastPeer.Peers = append(interfaceUnicastPeer.Peers, child)
        return &interfaceUnicastPeer.Peers[len(interfaceUnicastPeer.Peers)-1]
    }
    return nil
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceUnicastPeer.Peers {
        children[interfaceUnicastPeer.Peers[i].GetSegmentPath()] = &interfaceUnicastPeer.Peers[i]
    }
    return children
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceUnicastPeer.InterfaceName
    leafs["name"] = interfaceUnicastPeer.Name
    leafs["port-number"] = interfaceUnicastPeer.PortNumber
    return leafs
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetYangName() string { return "interface-unicast-peer" }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) SetParent(parent types.Entity) { interfaceUnicastPeer.parent = parent }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetParent() types.Entity { return interfaceUnicastPeer.parent }

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetParentYangName() string { return "interface-unicast-peers" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers
// Unicast Peers
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The address of the unicast peer.
    Address Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address

    // Unicast grant information for announce messages.
    AnnounceGrant Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant

    // Unicast grant information for sync messages.
    SyncGrant Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant

    // Unicast grant information for delay-response messages.
    DelayResponseGrant Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant
}

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "announce-grant" { return "AnnounceGrant" }
    if yname == "sync-grant" { return "SyncGrant" }
    if yname == "delay-response-grant" { return "DelayResponseGrant" }
    return ""
}

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &peers.Address
    }
    if childYangName == "announce-grant" {
        return &peers.AnnounceGrant
    }
    if childYangName == "sync-grant" {
        return &peers.SyncGrant
    }
    if childYangName == "delay-response-grant" {
        return &peers.DelayResponseGrant
    }
    return nil
}

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &peers.Address
    children["announce-grant"] = &peers.AnnounceGrant
    children["sync-grant"] = &peers.SyncGrant
    children["delay-response-grant"] = &peers.DelayResponseGrant
    return children
}

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetYangName() string { return "peers" }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetParent() types.Entity { return peers.parent }

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetParentYangName() string { return "interface-unicast-peer" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address
// The address of the unicast peer
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetGoName(yname string) string {
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "address-unknown" { return "AddressUnknown" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetSegmentPath() string {
    return "address"
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-address" {
        return &address.MacAddress
    }
    if childYangName == "ipv6-address" {
        return &address.Ipv6Address
    }
    return nil
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-address"] = &address.MacAddress
    children["ipv6-address"] = &address.Ipv6Address
    return children
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation"] = address.Encapsulation
    leafs["address-unknown"] = address.AddressUnknown
    leafs["ipv4-address"] = address.Ipv4Address
    return leafs
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetYangName() string { return "address" }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetParent() types.Entity { return address.parent }

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetParentYangName() string { return "peers" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetFilter() yfilter.YFilter { return macAddress.YFilter }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) SetFilter(yf yfilter.YFilter) { macAddress.YFilter = yf }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetGoName(yname string) string {
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetSegmentPath() string {
    return "mac-address"
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["macaddr"] = macAddress.Macaddr
    return leafs
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetBundleName() string { return "cisco_ios_xr" }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetYangName() string { return "mac-address" }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) SetParent(parent types.Entity) { macAddress.parent = parent }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetParent() types.Entity { return macAddress.parent }

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetParentYangName() string { return "address" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address
// IPv6 address
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address"
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6Address.Ipv6Address
    return leafs
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetParentYangName() string { return "address" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetFilter() yfilter.YFilter { return announceGrant.YFilter }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) SetFilter(yf yfilter.YFilter) { announceGrant.YFilter = yf }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetSegmentPath() string {
    return "announce-grant"
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = announceGrant.LogGrantInterval
    leafs["grant-duration"] = announceGrant.GrantDuration
    return leafs
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetBundleName() string { return "cisco_ios_xr" }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetYangName() string { return "announce-grant" }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) SetParent(parent types.Entity) { announceGrant.parent = parent }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetParent() types.Entity { return announceGrant.parent }

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetParentYangName() string { return "peers" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant
// Unicast grant information for sync messages
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetFilter() yfilter.YFilter { return syncGrant.YFilter }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) SetFilter(yf yfilter.YFilter) { syncGrant.YFilter = yf }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetSegmentPath() string {
    return "sync-grant"
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = syncGrant.LogGrantInterval
    leafs["grant-duration"] = syncGrant.GrantDuration
    return leafs
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetBundleName() string { return "cisco_ios_xr" }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetYangName() string { return "sync-grant" }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) SetParent(parent types.Entity) { syncGrant.parent = parent }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetParent() types.Entity { return syncGrant.parent }

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetParentYangName() string { return "peers" }

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetFilter() yfilter.YFilter { return delayResponseGrant.YFilter }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) SetFilter(yf yfilter.YFilter) { delayResponseGrant.YFilter = yf }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetGoName(yname string) string {
    if yname == "log-grant-interval" { return "LogGrantInterval" }
    if yname == "grant-duration" { return "GrantDuration" }
    return ""
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetSegmentPath() string {
    return "delay-response-grant"
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-grant-interval"] = delayResponseGrant.LogGrantInterval
    leafs["grant-duration"] = delayResponseGrant.GrantDuration
    return leafs
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetBundleName() string { return "cisco_ios_xr" }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetYangName() string { return "delay-response-grant" }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) SetParent(parent types.Entity) { delayResponseGrant.parent = parent }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetParent() types.Entity { return delayResponseGrant.parent }

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetParentYangName() string { return "peers" }

// Ptp_Platform
// PTP platform specific data
type Ptp_Platform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PTP servo related parameters.
    Servo Ptp_Platform_Servo
}

func (platform *Ptp_Platform) GetFilter() yfilter.YFilter { return platform.YFilter }

func (platform *Ptp_Platform) SetFilter(yf yfilter.YFilter) { platform.YFilter = yf }

func (platform *Ptp_Platform) GetGoName(yname string) string {
    if yname == "servo" { return "Servo" }
    return ""
}

func (platform *Ptp_Platform) GetSegmentPath() string {
    return "Cisco-IOS-XR-ptp-pd-oper:platform"
}

func (platform *Ptp_Platform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "servo" {
        return &platform.Servo
    }
    return nil
}

func (platform *Ptp_Platform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["servo"] = &platform.Servo
    return children
}

func (platform *Ptp_Platform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platform *Ptp_Platform) GetBundleName() string { return "cisco_ios_xr" }

func (platform *Ptp_Platform) GetYangName() string { return "platform" }

func (platform *Ptp_Platform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platform *Ptp_Platform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platform *Ptp_Platform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platform *Ptp_Platform) SetParent(parent types.Entity) { platform.parent = parent }

func (platform *Ptp_Platform) GetParent() types.Entity { return platform.parent }

func (platform *Ptp_Platform) GetParentYangName() string { return "ptp" }

// Ptp_Platform_Servo
// PTP servo related parameters
type Ptp_Platform_Servo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lock status of device. The type is interface{} with range: 0..65535.
    LockStatus interface{}

    // running status of apr. The type is bool.
    Running interface{}

    // status of device. The type is string with length: 0..50.
    DeviceStatus interface{}

    // log level of apr. The type is interface{} with range: 0..65535.
    LogLevel interface{}

    // last phase alignment accuracy. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PhaseAccuracyLast interface{}

    // number of sync timestamp received. The type is interface{} with range:
    // 0..4294967295.
    NumSyncTimestamp interface{}

    // number of delay timestamp received. The type is interface{} with range:
    // 0..4294967295.
    NumDelayTimestamp interface{}

    // number of setTime() been called. The type is interface{} with range:
    // 0..4294967295.
    NumSetTime interface{}

    // number of stepTime() been called. The type is interface{} with range:
    // 0..4294967295.
    NumStepTime interface{}

    // number of adjustFreq() been called. The type is interface{} with range:
    // 0..4294967295.
    NumAdjustFreq interface{}

    // number of adjustFreqTime() been called. The type is interface{} with range:
    // 0..4294967295.
    NumAdjustFreqTime interface{}

    // last input of adjustFreq. The type is interface{} with range:
    // -2147483648..2147483647.
    LastAdjustFreq interface{}

    // last input of stepTime. The type is interface{} with range:
    // -2147483648..2147483647.
    LastStepTime interface{}

    // number of sync timestamp discarded. The type is interface{} with range:
    // 0..4294967295.
    NumDiscardSyncTimestamp interface{}

    // number of delay timestamp discarded. The type is interface{} with range:
    // 0..4294967295.
    NumDiscardDelayTimestamp interface{}

    // last input flag of setTime. The type is bool.
    FlagofLastSetTime interface{}

    // Time Offset From Master. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    OffsetFromMaster interface{}

    // Mean Path Delay. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    MeanPathDelay interface{}

    // last input of setTime.
    LastSetTime Ptp_Platform_Servo_LastSetTime

    // last T1 timestamp received.
    LastReceivedT1 Ptp_Platform_Servo_LastReceivedT1

    // last T2 timestamp received.
    LastReceivedT2 Ptp_Platform_Servo_LastReceivedT2

    // last T3 timestamp received.
    LastReceivedT3 Ptp_Platform_Servo_LastReceivedT3

    // last T4 timestamp received.
    LastReceivedT4 Ptp_Platform_Servo_LastReceivedT4

    // pre T1 timestamp received.
    PreReceivedT1 Ptp_Platform_Servo_PreReceivedT1

    // pre T2 timestamp received.
    PreReceivedT2 Ptp_Platform_Servo_PreReceivedT2

    // pre T3 timestamp received.
    PreReceivedT3 Ptp_Platform_Servo_PreReceivedT3

    // pre T4 timestamp received.
    PreReceivedT4 Ptp_Platform_Servo_PreReceivedT4
}

func (servo *Ptp_Platform_Servo) GetFilter() yfilter.YFilter { return servo.YFilter }

func (servo *Ptp_Platform_Servo) SetFilter(yf yfilter.YFilter) { servo.YFilter = yf }

func (servo *Ptp_Platform_Servo) GetGoName(yname string) string {
    if yname == "lock-status" { return "LockStatus" }
    if yname == "running" { return "Running" }
    if yname == "device-status" { return "DeviceStatus" }
    if yname == "log-level" { return "LogLevel" }
    if yname == "phase-accuracy-last" { return "PhaseAccuracyLast" }
    if yname == "num-sync-timestamp" { return "NumSyncTimestamp" }
    if yname == "num-delay-timestamp" { return "NumDelayTimestamp" }
    if yname == "num-set-time" { return "NumSetTime" }
    if yname == "num-step-time" { return "NumStepTime" }
    if yname == "num-adjust-freq" { return "NumAdjustFreq" }
    if yname == "num-adjust-freq-time" { return "NumAdjustFreqTime" }
    if yname == "last-adjust-freq" { return "LastAdjustFreq" }
    if yname == "last-step-time" { return "LastStepTime" }
    if yname == "num-discard-sync-timestamp" { return "NumDiscardSyncTimestamp" }
    if yname == "num-discard-delay-timestamp" { return "NumDiscardDelayTimestamp" }
    if yname == "flagof-last-set-time" { return "FlagofLastSetTime" }
    if yname == "offset-from-master" { return "OffsetFromMaster" }
    if yname == "mean-path-delay" { return "MeanPathDelay" }
    if yname == "last-set-time" { return "LastSetTime" }
    if yname == "last-received-t1" { return "LastReceivedT1" }
    if yname == "last-received-t2" { return "LastReceivedT2" }
    if yname == "last-received-t3" { return "LastReceivedT3" }
    if yname == "last-received-t4" { return "LastReceivedT4" }
    if yname == "pre-received-t1" { return "PreReceivedT1" }
    if yname == "pre-received-t2" { return "PreReceivedT2" }
    if yname == "pre-received-t3" { return "PreReceivedT3" }
    if yname == "pre-received-t4" { return "PreReceivedT4" }
    return ""
}

func (servo *Ptp_Platform_Servo) GetSegmentPath() string {
    return "servo"
}

func (servo *Ptp_Platform_Servo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-set-time" {
        return &servo.LastSetTime
    }
    if childYangName == "last-received-t1" {
        return &servo.LastReceivedT1
    }
    if childYangName == "last-received-t2" {
        return &servo.LastReceivedT2
    }
    if childYangName == "last-received-t3" {
        return &servo.LastReceivedT3
    }
    if childYangName == "last-received-t4" {
        return &servo.LastReceivedT4
    }
    if childYangName == "pre-received-t1" {
        return &servo.PreReceivedT1
    }
    if childYangName == "pre-received-t2" {
        return &servo.PreReceivedT2
    }
    if childYangName == "pre-received-t3" {
        return &servo.PreReceivedT3
    }
    if childYangName == "pre-received-t4" {
        return &servo.PreReceivedT4
    }
    return nil
}

func (servo *Ptp_Platform_Servo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-set-time"] = &servo.LastSetTime
    children["last-received-t1"] = &servo.LastReceivedT1
    children["last-received-t2"] = &servo.LastReceivedT2
    children["last-received-t3"] = &servo.LastReceivedT3
    children["last-received-t4"] = &servo.LastReceivedT4
    children["pre-received-t1"] = &servo.PreReceivedT1
    children["pre-received-t2"] = &servo.PreReceivedT2
    children["pre-received-t3"] = &servo.PreReceivedT3
    children["pre-received-t4"] = &servo.PreReceivedT4
    return children
}

func (servo *Ptp_Platform_Servo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lock-status"] = servo.LockStatus
    leafs["running"] = servo.Running
    leafs["device-status"] = servo.DeviceStatus
    leafs["log-level"] = servo.LogLevel
    leafs["phase-accuracy-last"] = servo.PhaseAccuracyLast
    leafs["num-sync-timestamp"] = servo.NumSyncTimestamp
    leafs["num-delay-timestamp"] = servo.NumDelayTimestamp
    leafs["num-set-time"] = servo.NumSetTime
    leafs["num-step-time"] = servo.NumStepTime
    leafs["num-adjust-freq"] = servo.NumAdjustFreq
    leafs["num-adjust-freq-time"] = servo.NumAdjustFreqTime
    leafs["last-adjust-freq"] = servo.LastAdjustFreq
    leafs["last-step-time"] = servo.LastStepTime
    leafs["num-discard-sync-timestamp"] = servo.NumDiscardSyncTimestamp
    leafs["num-discard-delay-timestamp"] = servo.NumDiscardDelayTimestamp
    leafs["flagof-last-set-time"] = servo.FlagofLastSetTime
    leafs["offset-from-master"] = servo.OffsetFromMaster
    leafs["mean-path-delay"] = servo.MeanPathDelay
    return leafs
}

func (servo *Ptp_Platform_Servo) GetBundleName() string { return "cisco_ios_xr" }

func (servo *Ptp_Platform_Servo) GetYangName() string { return "servo" }

func (servo *Ptp_Platform_Servo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servo *Ptp_Platform_Servo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servo *Ptp_Platform_Servo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servo *Ptp_Platform_Servo) SetParent(parent types.Entity) { servo.parent = parent }

func (servo *Ptp_Platform_Servo) GetParent() types.Entity { return servo.parent }

func (servo *Ptp_Platform_Servo) GetParentYangName() string { return "platform" }

// Ptp_Platform_Servo_LastSetTime
// last input of setTime
type Ptp_Platform_Servo_LastSetTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetFilter() yfilter.YFilter { return lastSetTime.YFilter }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) SetFilter(yf yfilter.YFilter) { lastSetTime.YFilter = yf }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetSegmentPath() string {
    return "last-set-time"
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastSetTime.Second
    leafs["nano-second"] = lastSetTime.NanoSecond
    return leafs
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetBundleName() string { return "cisco_ios_xr" }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetYangName() string { return "last-set-time" }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) SetParent(parent types.Entity) { lastSetTime.parent = parent }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetParent() types.Entity { return lastSetTime.parent }

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_LastReceivedT1
// last T1 timestamp received
type Ptp_Platform_Servo_LastReceivedT1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetFilter() yfilter.YFilter { return lastReceivedT1.YFilter }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) SetFilter(yf yfilter.YFilter) { lastReceivedT1.YFilter = yf }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetSegmentPath() string {
    return "last-received-t1"
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT1.Second
    leafs["nano-second"] = lastReceivedT1.NanoSecond
    return leafs
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetYangName() string { return "last-received-t1" }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) SetParent(parent types.Entity) { lastReceivedT1.parent = parent }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetParent() types.Entity { return lastReceivedT1.parent }

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_LastReceivedT2
// last T2 timestamp received
type Ptp_Platform_Servo_LastReceivedT2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetFilter() yfilter.YFilter { return lastReceivedT2.YFilter }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) SetFilter(yf yfilter.YFilter) { lastReceivedT2.YFilter = yf }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetSegmentPath() string {
    return "last-received-t2"
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT2.Second
    leafs["nano-second"] = lastReceivedT2.NanoSecond
    return leafs
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetYangName() string { return "last-received-t2" }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) SetParent(parent types.Entity) { lastReceivedT2.parent = parent }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetParent() types.Entity { return lastReceivedT2.parent }

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_LastReceivedT3
// last T3 timestamp received
type Ptp_Platform_Servo_LastReceivedT3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetFilter() yfilter.YFilter { return lastReceivedT3.YFilter }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) SetFilter(yf yfilter.YFilter) { lastReceivedT3.YFilter = yf }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetSegmentPath() string {
    return "last-received-t3"
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT3.Second
    leafs["nano-second"] = lastReceivedT3.NanoSecond
    return leafs
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetYangName() string { return "last-received-t3" }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) SetParent(parent types.Entity) { lastReceivedT3.parent = parent }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetParent() types.Entity { return lastReceivedT3.parent }

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_LastReceivedT4
// last T4 timestamp received
type Ptp_Platform_Servo_LastReceivedT4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetFilter() yfilter.YFilter { return lastReceivedT4.YFilter }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) SetFilter(yf yfilter.YFilter) { lastReceivedT4.YFilter = yf }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetSegmentPath() string {
    return "last-received-t4"
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT4.Second
    leafs["nano-second"] = lastReceivedT4.NanoSecond
    return leafs
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetYangName() string { return "last-received-t4" }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) SetParent(parent types.Entity) { lastReceivedT4.parent = parent }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetParent() types.Entity { return lastReceivedT4.parent }

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_PreReceivedT1
// pre T1 timestamp received
type Ptp_Platform_Servo_PreReceivedT1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetFilter() yfilter.YFilter { return preReceivedT1.YFilter }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) SetFilter(yf yfilter.YFilter) { preReceivedT1.YFilter = yf }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetSegmentPath() string {
    return "pre-received-t1"
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT1.Second
    leafs["nano-second"] = preReceivedT1.NanoSecond
    return leafs
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetYangName() string { return "pre-received-t1" }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) SetParent(parent types.Entity) { preReceivedT1.parent = parent }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetParent() types.Entity { return preReceivedT1.parent }

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_PreReceivedT2
// pre T2 timestamp received
type Ptp_Platform_Servo_PreReceivedT2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetFilter() yfilter.YFilter { return preReceivedT2.YFilter }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) SetFilter(yf yfilter.YFilter) { preReceivedT2.YFilter = yf }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetSegmentPath() string {
    return "pre-received-t2"
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT2.Second
    leafs["nano-second"] = preReceivedT2.NanoSecond
    return leafs
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetYangName() string { return "pre-received-t2" }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) SetParent(parent types.Entity) { preReceivedT2.parent = parent }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetParent() types.Entity { return preReceivedT2.parent }

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_PreReceivedT3
// pre T3 timestamp received
type Ptp_Platform_Servo_PreReceivedT3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetFilter() yfilter.YFilter { return preReceivedT3.YFilter }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) SetFilter(yf yfilter.YFilter) { preReceivedT3.YFilter = yf }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetSegmentPath() string {
    return "pre-received-t3"
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT3.Second
    leafs["nano-second"] = preReceivedT3.NanoSecond
    return leafs
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetYangName() string { return "pre-received-t3" }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) SetParent(parent types.Entity) { preReceivedT3.parent = parent }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetParent() types.Entity { return preReceivedT3.parent }

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetParentYangName() string { return "servo" }

// Ptp_Platform_Servo_PreReceivedT4
// pre T4 timestamp received
type Ptp_Platform_Servo_PreReceivedT4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetFilter() yfilter.YFilter { return preReceivedT4.YFilter }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) SetFilter(yf yfilter.YFilter) { preReceivedT4.YFilter = yf }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetSegmentPath() string {
    return "pre-received-t4"
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT4.Second
    leafs["nano-second"] = preReceivedT4.NanoSecond
    return leafs
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetYangName() string { return "pre-received-t4" }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) SetParent(parent types.Entity) { preReceivedT4.parent = parent }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetParent() types.Entity { return preReceivedT4.parent }

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetParentYangName() string { return "servo" }

