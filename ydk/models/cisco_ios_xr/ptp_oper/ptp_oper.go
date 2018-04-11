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

// PtpBagClockLeapSeconds represents Leap Seconds
type PtpBagClockLeapSeconds string

const (
    // No leap second
    PtpBagClockLeapSeconds_none PtpBagClockLeapSeconds = "none"

    // The last minute of the day has 59 seconds
    PtpBagClockLeapSeconds_leap59 PtpBagClockLeapSeconds = "leap59"

    // The last minute of the day has 61 seconds
    PtpBagClockLeapSeconds_leap61 PtpBagClockLeapSeconds = "leap61"
)

// PtpBagClockTimescale represents Timescale
type PtpBagClockTimescale string

const (
    // PTP timescale
    PtpBagClockTimescale_ptp PtpBagClockTimescale = "ptp"

    // ARB timescale
    PtpBagClockTimescale_arb PtpBagClockTimescale = "arb"
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

// Ptp
// PTP operational data
type Ptp struct {
    EntityData types.CommonEntityData
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

    // UTC offset information.
    UtcOffsetInfo Ptp_UtcOffsetInfo

    // PTP platform specific data.
    Platform Ptp_Platform
}

func (ptp *Ptp) GetEntityData() *types.CommonEntityData {
    ptp.EntityData.YFilter = ptp.YFilter
    ptp.EntityData.YangName = "ptp"
    ptp.EntityData.BundleName = "cisco_ios_xr"
    ptp.EntityData.ParentYangName = "Cisco-IOS-XR-ptp-oper"
    ptp.EntityData.SegmentPath = "Cisco-IOS-XR-ptp-oper:ptp"
    ptp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptp.EntityData.Children = make(map[string]types.YChild)
    ptp.EntityData.Children["nodes"] = types.YChild{"Nodes", &ptp.Nodes}
    ptp.EntityData.Children["interface-configuration-errors"] = types.YChild{"InterfaceConfigurationErrors", &ptp.InterfaceConfigurationErrors}
    ptp.EntityData.Children["interface-foreign-masters"] = types.YChild{"InterfaceForeignMasters", &ptp.InterfaceForeignMasters}
    ptp.EntityData.Children["local-clock"] = types.YChild{"LocalClock", &ptp.LocalClock}
    ptp.EntityData.Children["interface-packet-counters"] = types.YChild{"InterfacePacketCounters", &ptp.InterfacePacketCounters}
    ptp.EntityData.Children["advertised-clock"] = types.YChild{"AdvertisedClock", &ptp.AdvertisedClock}
    ptp.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &ptp.Interfaces}
    ptp.EntityData.Children["dataset"] = types.YChild{"Dataset", &ptp.Dataset}
    ptp.EntityData.Children["global-configuration-error"] = types.YChild{"GlobalConfigurationError", &ptp.GlobalConfigurationError}
    ptp.EntityData.Children["grandmaster"] = types.YChild{"Grandmaster", &ptp.Grandmaster}
    ptp.EntityData.Children["interface-unicast-peers"] = types.YChild{"InterfaceUnicastPeers", &ptp.InterfaceUnicastPeers}
    ptp.EntityData.Children["utc-offset-info"] = types.YChild{"UtcOffsetInfo", &ptp.UtcOffsetInfo}
    ptp.EntityData.Children["Cisco-IOS-XR-ptp-pd-oper:platform"] = types.YChild{"Platform", &ptp.Platform}
    ptp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ptp.EntityData)
}

// Ptp_Nodes
// Table for node-specific operational data
type Ptp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific operational data for a given node. The type is slice of
    // Ptp_Nodes_Node.
    Node []Ptp_Nodes_Node
}

func (nodes *Ptp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ptp"
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

// Ptp_Nodes_Node
// Node-specific operational data for a given node
type Ptp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Ptp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["node-interface-foreign-masters"] = types.YChild{"NodeInterfaceForeignMasters", &node.NodeInterfaceForeignMasters}
    node.EntityData.Children["summary"] = types.YChild{"Summary", &node.Summary}
    node.EntityData.Children["node-interfaces"] = types.YChild{"NodeInterfaces", &node.NodeInterfaces}
    node.EntityData.Children["node-interface-unicast-peers"] = types.YChild{"NodeInterfaceUnicastPeers", &node.NodeInterfaceUnicastPeers}
    node.EntityData.Children["packet-counters"] = types.YChild{"PacketCounters", &node.PacketCounters}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters
// Table for node foreign master clock
// operational data
type Ptp_Nodes_Node_NodeInterfaceForeignMasters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node interface foreign master clock operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster.
    NodeInterfaceForeignMaster []Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetEntityData() *types.CommonEntityData {
    nodeInterfaceForeignMasters.EntityData.YFilter = nodeInterfaceForeignMasters.YFilter
    nodeInterfaceForeignMasters.EntityData.YangName = "node-interface-foreign-masters"
    nodeInterfaceForeignMasters.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceForeignMasters.EntityData.ParentYangName = "node"
    nodeInterfaceForeignMasters.EntityData.SegmentPath = "node-interface-foreign-masters"
    nodeInterfaceForeignMasters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceForeignMasters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceForeignMasters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceForeignMasters.EntityData.Children = make(map[string]types.YChild)
    nodeInterfaceForeignMasters.EntityData.Children["node-interface-foreign-master"] = types.YChild{"NodeInterfaceForeignMaster", nil}
    for i := range nodeInterfaceForeignMasters.NodeInterfaceForeignMaster {
        nodeInterfaceForeignMasters.EntityData.Children[types.GetSegmentPath(&nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[i])] = types.YChild{"NodeInterfaceForeignMaster", &nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[i]}
    }
    nodeInterfaceForeignMasters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodeInterfaceForeignMasters.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster
// Node interface foreign master clock
// operational data
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Foreign clocks received on this interface. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock.
    ForeignClock []Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetEntityData() *types.CommonEntityData {
    nodeInterfaceForeignMaster.EntityData.YFilter = nodeInterfaceForeignMaster.YFilter
    nodeInterfaceForeignMaster.EntityData.YangName = "node-interface-foreign-master"
    nodeInterfaceForeignMaster.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceForeignMaster.EntityData.ParentYangName = "node-interface-foreign-masters"
    nodeInterfaceForeignMaster.EntityData.SegmentPath = "node-interface-foreign-master" + "[interface-name='" + fmt.Sprintf("%v", nodeInterfaceForeignMaster.InterfaceName) + "']"
    nodeInterfaceForeignMaster.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceForeignMaster.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceForeignMaster.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceForeignMaster.EntityData.Children = make(map[string]types.YChild)
    nodeInterfaceForeignMaster.EntityData.Children["foreign-clock"] = types.YChild{"ForeignClock", nil}
    for i := range nodeInterfaceForeignMaster.ForeignClock {
        nodeInterfaceForeignMaster.EntityData.Children[types.GetSegmentPath(&nodeInterfaceForeignMaster.ForeignClock[i])] = types.YChild{"ForeignClock", &nodeInterfaceForeignMaster.ForeignClock[i]}
    }
    nodeInterfaceForeignMaster.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeInterfaceForeignMaster.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", nodeInterfaceForeignMaster.InterfaceName}
    nodeInterfaceForeignMaster.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", nodeInterfaceForeignMaster.PortNumber}
    return &(nodeInterfaceForeignMaster.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock
// Foreign clocks received on this interface
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock struct {
    EntityData types.CommonEntityData
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
    ForeignClock Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_

    // The address of the clock.
    Address Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address

    // Unicast grant information for announce messages.
    AnnounceGrant Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant

    // Unicast grant information for sync messages.
    SyncGrant Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant

    // Unicast grant information for delay-response messages.
    DelayResponseGrant Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetEntityData() *types.CommonEntityData {
    foreignClock.EntityData.YFilter = foreignClock.YFilter
    foreignClock.EntityData.YangName = "foreign-clock"
    foreignClock.EntityData.BundleName = "cisco_ios_xr"
    foreignClock.EntityData.ParentYangName = "node-interface-foreign-master"
    foreignClock.EntityData.SegmentPath = "foreign-clock"
    foreignClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock.EntityData.Children = make(map[string]types.YChild)
    foreignClock.EntityData.Children["foreign-clock"] = types.YChild{"ForeignClock", &foreignClock.ForeignClock}
    foreignClock.EntityData.Children["address"] = types.YChild{"Address", &foreignClock.Address}
    foreignClock.EntityData.Children["announce-grant"] = types.YChild{"AnnounceGrant", &foreignClock.AnnounceGrant}
    foreignClock.EntityData.Children["sync-grant"] = types.YChild{"SyncGrant", &foreignClock.SyncGrant}
    foreignClock.EntityData.Children["delay-response-grant"] = types.YChild{"DelayResponseGrant", &foreignClock.DelayResponseGrant}
    foreignClock.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignClock.EntityData.Leafs["is-qualified"] = types.YLeaf{"IsQualified", foreignClock.IsQualified}
    foreignClock.EntityData.Leafs["is-grandmaster"] = types.YLeaf{"IsGrandmaster", foreignClock.IsGrandmaster}
    foreignClock.EntityData.Leafs["communication-model"] = types.YLeaf{"CommunicationModel", foreignClock.CommunicationModel}
    foreignClock.EntityData.Leafs["is-known"] = types.YLeaf{"IsKnown", foreignClock.IsKnown}
    foreignClock.EntityData.Leafs["time-known-for"] = types.YLeaf{"TimeKnownFor", foreignClock.TimeKnownFor}
    foreignClock.EntityData.Leafs["foreign-domain"] = types.YLeaf{"ForeignDomain", foreignClock.ForeignDomain}
    foreignClock.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", foreignClock.ConfiguredPriority}
    foreignClock.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", foreignClock.ConfiguredClockClass}
    foreignClock.EntityData.Leafs["delay-asymmetry"] = types.YLeaf{"DelayAsymmetry", foreignClock.DelayAsymmetry}
    foreignClock.EntityData.Leafs["ptsf-loss-announce"] = types.YLeaf{"PtsfLossAnnounce", foreignClock.PtsfLossAnnounce}
    foreignClock.EntityData.Leafs["ptsf-loss-sync"] = types.YLeaf{"PtsfLossSync", foreignClock.PtsfLossSync}
    return &(foreignClock.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_
// Foreign clock information
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_ struct {
    EntityData types.CommonEntityData
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
    UtcOffset Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset

    // Receiver.
    Receiver Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Receiver

    // Sender.
    Sender Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Sender
}

func (foreignClock_ *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_) GetEntityData() *types.CommonEntityData {
    foreignClock_.EntityData.YFilter = foreignClock_.YFilter
    foreignClock_.EntityData.YangName = "foreign-clock"
    foreignClock_.EntityData.BundleName = "cisco_ios_xr"
    foreignClock_.EntityData.ParentYangName = "foreign-clock"
    foreignClock_.EntityData.SegmentPath = "foreign-clock"
    foreignClock_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock_.EntityData.Children = make(map[string]types.YChild)
    foreignClock_.EntityData.Children["utc-offset"] = types.YChild{"UtcOffset", &foreignClock_.UtcOffset}
    foreignClock_.EntityData.Children["receiver"] = types.YChild{"Receiver", &foreignClock_.Receiver}
    foreignClock_.EntityData.Children["sender"] = types.YChild{"Sender", &foreignClock_.Sender}
    foreignClock_.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignClock_.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", foreignClock_.ClockId}
    foreignClock_.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", foreignClock_.Priority1}
    foreignClock_.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", foreignClock_.Priority2}
    foreignClock_.EntityData.Leafs["class"] = types.YLeaf{"Class", foreignClock_.Class}
    foreignClock_.EntityData.Leafs["accuracy"] = types.YLeaf{"Accuracy", foreignClock_.Accuracy}
    foreignClock_.EntityData.Leafs["offset-log-variance"] = types.YLeaf{"OffsetLogVariance", foreignClock_.OffsetLogVariance}
    foreignClock_.EntityData.Leafs["steps-removed"] = types.YLeaf{"StepsRemoved", foreignClock_.StepsRemoved}
    foreignClock_.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", foreignClock_.TimeSource}
    foreignClock_.EntityData.Leafs["frequency-traceable"] = types.YLeaf{"FrequencyTraceable", foreignClock_.FrequencyTraceable}
    foreignClock_.EntityData.Leafs["time-traceable"] = types.YLeaf{"TimeTraceable", foreignClock_.TimeTraceable}
    foreignClock_.EntityData.Leafs["timescale"] = types.YLeaf{"Timescale", foreignClock_.Timescale}
    foreignClock_.EntityData.Leafs["leap-seconds"] = types.YLeaf{"LeapSeconds", foreignClock_.LeapSeconds}
    foreignClock_.EntityData.Leafs["local"] = types.YLeaf{"Local", foreignClock_.Local}
    foreignClock_.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", foreignClock_.ConfiguredClockClass}
    foreignClock_.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", foreignClock_.ConfiguredPriority}
    return &(foreignClock_.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset
// UTC offset
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "foreign-clock"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = make(map[string]types.YChild)
    utcOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffset.EntityData.Leafs["current-offset"] = types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset}
    utcOffset.EntityData.Leafs["offset-valid"] = types.YLeaf{"OffsetValid", utcOffset.OffsetValid}
    return &(utcOffset.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Receiver
// Receiver
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "foreign-clock"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = make(map[string]types.YChild)
    receiver.EntityData.Leafs = make(map[string]types.YLeaf)
    receiver.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", receiver.ClockId}
    receiver.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", receiver.PortNumber}
    return &(receiver.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Sender
// Sender
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock__Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "foreign-clock"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = make(map[string]types.YChild)
    sender.EntityData.Leafs = make(map[string]types.YLeaf)
    sender.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", sender.ClockId}
    sender.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", sender.PortNumber}
    return &(sender.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address
// The address of the clock
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address
}

func (address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "foreign-clock"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress
// Ethernet MAC address
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address
// IPv6 address
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_AnnounceGrant) GetEntityData() *types.CommonEntityData {
    announceGrant.EntityData.YFilter = announceGrant.YFilter
    announceGrant.EntityData.YangName = "announce-grant"
    announceGrant.EntityData.BundleName = "cisco_ios_xr"
    announceGrant.EntityData.ParentYangName = "foreign-clock"
    announceGrant.EntityData.SegmentPath = "announce-grant"
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = make(map[string]types.YChild)
    announceGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    announceGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval}
    announceGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", announceGrant.GrantDuration}
    return &(announceGrant.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant
// Unicast grant information for sync messages
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_SyncGrant) GetEntityData() *types.CommonEntityData {
    syncGrant.EntityData.YFilter = syncGrant.YFilter
    syncGrant.EntityData.YangName = "sync-grant"
    syncGrant.EntityData.BundleName = "cisco_ios_xr"
    syncGrant.EntityData.ParentYangName = "foreign-clock"
    syncGrant.EntityData.SegmentPath = "sync-grant"
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = make(map[string]types.YChild)
    syncGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    syncGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval}
    syncGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", syncGrant.GrantDuration}
    return &(syncGrant.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetEntityData() *types.CommonEntityData {
    delayResponseGrant.EntityData.YFilter = delayResponseGrant.YFilter
    delayResponseGrant.EntityData.YangName = "delay-response-grant"
    delayResponseGrant.EntityData.BundleName = "cisco_ios_xr"
    delayResponseGrant.EntityData.ParentYangName = "foreign-clock"
    delayResponseGrant.EntityData.SegmentPath = "delay-response-grant"
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = make(map[string]types.YChild)
    delayResponseGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    delayResponseGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval}
    delayResponseGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration}
    return &(delayResponseGrant.EntityData)
}

// Ptp_Nodes_Node_Summary
// Node summary operational data
type Ptp_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *Ptp_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["port-state-init-count"] = types.YLeaf{"PortStateInitCount", summary.PortStateInitCount}
    summary.EntityData.Leafs["port-state-listening-count"] = types.YLeaf{"PortStateListeningCount", summary.PortStateListeningCount}
    summary.EntityData.Leafs["port-state-passive-count"] = types.YLeaf{"PortStatePassiveCount", summary.PortStatePassiveCount}
    summary.EntityData.Leafs["port-state-pre-master-count"] = types.YLeaf{"PortStatePreMasterCount", summary.PortStatePreMasterCount}
    summary.EntityData.Leafs["port-state-master-count"] = types.YLeaf{"PortStateMasterCount", summary.PortStateMasterCount}
    summary.EntityData.Leafs["port-state-slave-count"] = types.YLeaf{"PortStateSlaveCount", summary.PortStateSlaveCount}
    summary.EntityData.Leafs["port-state-uncalibrated-count"] = types.YLeaf{"PortStateUncalibratedCount", summary.PortStateUncalibratedCount}
    summary.EntityData.Leafs["port-state-faulty-count"] = types.YLeaf{"PortStateFaultyCount", summary.PortStateFaultyCount}
    summary.EntityData.Leafs["total-interfaces"] = types.YLeaf{"TotalInterfaces", summary.TotalInterfaces}
    summary.EntityData.Leafs["total-interfaces-valid-port-num"] = types.YLeaf{"TotalInterfacesValidPortNum", summary.TotalInterfacesValidPortNum}
    return &(summary.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces
// Table for node interface operational data
type Ptp_Nodes_Node_NodeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node interface operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface.
    NodeInterface []Ptp_Nodes_Node_NodeInterfaces_NodeInterface
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetEntityData() *types.CommonEntityData {
    nodeInterfaces.EntityData.YFilter = nodeInterfaces.YFilter
    nodeInterfaces.EntityData.YangName = "node-interfaces"
    nodeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaces.EntityData.ParentYangName = "node"
    nodeInterfaces.EntityData.SegmentPath = "node-interfaces"
    nodeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaces.EntityData.Children = make(map[string]types.YChild)
    nodeInterfaces.EntityData.Children["node-interface"] = types.YChild{"NodeInterface", nil}
    for i := range nodeInterfaces.NodeInterface {
        nodeInterfaces.EntityData.Children[types.GetSegmentPath(&nodeInterfaces.NodeInterface[i])] = types.YChild{"NodeInterface", &nodeInterfaces.NodeInterface[i]}
    }
    nodeInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodeInterfaces.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface
// Node interface operational data
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // IPv4 address, if IPv4 encapsulation is being used. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetEntityData() *types.CommonEntityData {
    nodeInterface.EntityData.YFilter = nodeInterface.YFilter
    nodeInterface.EntityData.YangName = "node-interface"
    nodeInterface.EntityData.BundleName = "cisco_ios_xr"
    nodeInterface.EntityData.ParentYangName = "node-interfaces"
    nodeInterface.EntityData.SegmentPath = "node-interface" + "[interface-name='" + fmt.Sprintf("%v", nodeInterface.InterfaceName) + "']"
    nodeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterface.EntityData.Children = make(map[string]types.YChild)
    nodeInterface.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &nodeInterface.MacAddress}
    nodeInterface.EntityData.Children["master-table"] = types.YChild{"MasterTable", nil}
    for i := range nodeInterface.MasterTable {
        nodeInterface.EntityData.Children[types.GetSegmentPath(&nodeInterface.MasterTable[i])] = types.YChild{"MasterTable", &nodeInterface.MasterTable[i]}
    }
    nodeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", nodeInterface.InterfaceName}
    nodeInterface.EntityData.Leafs["port-state"] = types.YLeaf{"PortState", nodeInterface.PortState}
    nodeInterface.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", nodeInterface.PortNumber}
    nodeInterface.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", nodeInterface.LineState}
    nodeInterface.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", nodeInterface.Encapsulation}
    nodeInterface.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nodeInterface.Ipv6Address}
    nodeInterface.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nodeInterface.Ipv4Address}
    nodeInterface.EntityData.Leafs["two-step"] = types.YLeaf{"TwoStep", nodeInterface.TwoStep}
    nodeInterface.EntityData.Leafs["communication-model"] = types.YLeaf{"CommunicationModel", nodeInterface.CommunicationModel}
    nodeInterface.EntityData.Leafs["log-sync-interval"] = types.YLeaf{"LogSyncInterval", nodeInterface.LogSyncInterval}
    nodeInterface.EntityData.Leafs["log-announce-interval"] = types.YLeaf{"LogAnnounceInterval", nodeInterface.LogAnnounceInterval}
    nodeInterface.EntityData.Leafs["announce-timeout"] = types.YLeaf{"AnnounceTimeout", nodeInterface.AnnounceTimeout}
    nodeInterface.EntityData.Leafs["log-min-delay-request-interval"] = types.YLeaf{"LogMinDelayRequestInterval", nodeInterface.LogMinDelayRequestInterval}
    nodeInterface.EntityData.Leafs["configured-port-state"] = types.YLeaf{"ConfiguredPortState", nodeInterface.ConfiguredPortState}
    nodeInterface.EntityData.Leafs["supports-one-step"] = types.YLeaf{"SupportsOneStep", nodeInterface.SupportsOneStep}
    nodeInterface.EntityData.Leafs["supports-two-step"] = types.YLeaf{"SupportsTwoStep", nodeInterface.SupportsTwoStep}
    nodeInterface.EntityData.Leafs["supports-ethernet"] = types.YLeaf{"SupportsEthernet", nodeInterface.SupportsEthernet}
    nodeInterface.EntityData.Leafs["supports-multicast"] = types.YLeaf{"SupportsMulticast", nodeInterface.SupportsMulticast}
    nodeInterface.EntityData.Leafs["supports-ipv6"] = types.YLeaf{"SupportsIpv6", nodeInterface.SupportsIpv6}
    nodeInterface.EntityData.Leafs["supports-slave"] = types.YLeaf{"SupportsSlave", nodeInterface.SupportsSlave}
    nodeInterface.EntityData.Leafs["supports-source-ip"] = types.YLeaf{"SupportsSourceIp", nodeInterface.SupportsSourceIp}
    nodeInterface.EntityData.Leafs["max-sync-rate"] = types.YLeaf{"MaxSyncRate", nodeInterface.MaxSyncRate}
    nodeInterface.EntityData.Leafs["event-cos"] = types.YLeaf{"EventCos", nodeInterface.EventCos}
    nodeInterface.EntityData.Leafs["general-cos"] = types.YLeaf{"GeneralCos", nodeInterface.GeneralCos}
    nodeInterface.EntityData.Leafs["event-dscp"] = types.YLeaf{"EventDscp", nodeInterface.EventDscp}
    nodeInterface.EntityData.Leafs["general-dscp"] = types.YLeaf{"GeneralDscp", nodeInterface.GeneralDscp}
    nodeInterface.EntityData.Leafs["unicast-peers"] = types.YLeaf{"UnicastPeers", nodeInterface.UnicastPeers}
    nodeInterface.EntityData.Leafs["local-priority"] = types.YLeaf{"LocalPriority", nodeInterface.LocalPriority}
    nodeInterface.EntityData.Leafs["signal-fail"] = types.YLeaf{"SignalFail", nodeInterface.SignalFail}
    return &(nodeInterface.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress
// MAC address, if Ethernet encapsulation is being
// used
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "node-interface"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable
// The interface's master table
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable struct {
    EntityData types.CommonEntityData
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

func (masterTable *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable) GetEntityData() *types.CommonEntityData {
    masterTable.EntityData.YFilter = masterTable.YFilter
    masterTable.EntityData.YangName = "master-table"
    masterTable.EntityData.BundleName = "cisco_ios_xr"
    masterTable.EntityData.ParentYangName = "node-interface"
    masterTable.EntityData.SegmentPath = "master-table"
    masterTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterTable.EntityData.Children = make(map[string]types.YChild)
    masterTable.EntityData.Children["address"] = types.YChild{"Address", &masterTable.Address}
    masterTable.EntityData.Leafs = make(map[string]types.YLeaf)
    masterTable.EntityData.Leafs["communication-model"] = types.YLeaf{"CommunicationModel", masterTable.CommunicationModel}
    masterTable.EntityData.Leafs["priority"] = types.YLeaf{"Priority", masterTable.Priority}
    masterTable.EntityData.Leafs["known"] = types.YLeaf{"Known", masterTable.Known}
    masterTable.EntityData.Leafs["qualified"] = types.YLeaf{"Qualified", masterTable.Qualified}
    masterTable.EntityData.Leafs["is-grandmaster"] = types.YLeaf{"IsGrandmaster", masterTable.IsGrandmaster}
    masterTable.EntityData.Leafs["ptsf-loss-announce"] = types.YLeaf{"PtsfLossAnnounce", masterTable.PtsfLossAnnounce}
    masterTable.EntityData.Leafs["ptsf-loss-sync"] = types.YLeaf{"PtsfLossSync", masterTable.PtsfLossSync}
    masterTable.EntityData.Leafs["is-nonnegotiated"] = types.YLeaf{"IsNonnegotiated", masterTable.IsNonnegotiated}
    return &(masterTable.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address
// The address of the master clock
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address
}

func (address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "master-table"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress
// Ethernet MAC address
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address
// IPv6 address
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers
// Table for node unicast peers operational data
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node interface unicast peers operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer.
    NodeInterfaceUnicastPeer []Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetEntityData() *types.CommonEntityData {
    nodeInterfaceUnicastPeers.EntityData.YFilter = nodeInterfaceUnicastPeers.YFilter
    nodeInterfaceUnicastPeers.EntityData.YangName = "node-interface-unicast-peers"
    nodeInterfaceUnicastPeers.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceUnicastPeers.EntityData.ParentYangName = "node"
    nodeInterfaceUnicastPeers.EntityData.SegmentPath = "node-interface-unicast-peers"
    nodeInterfaceUnicastPeers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceUnicastPeers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceUnicastPeers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceUnicastPeers.EntityData.Children = make(map[string]types.YChild)
    nodeInterfaceUnicastPeers.EntityData.Children["node-interface-unicast-peer"] = types.YChild{"NodeInterfaceUnicastPeer", nil}
    for i := range nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer {
        nodeInterfaceUnicastPeers.EntityData.Children[types.GetSegmentPath(&nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[i])] = types.YChild{"NodeInterfaceUnicastPeer", &nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[i]}
    }
    nodeInterfaceUnicastPeers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodeInterfaceUnicastPeers.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer
// Node interface unicast peers operational data
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Unicast Peers. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers.
    Peers []Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetEntityData() *types.CommonEntityData {
    nodeInterfaceUnicastPeer.EntityData.YFilter = nodeInterfaceUnicastPeer.YFilter
    nodeInterfaceUnicastPeer.EntityData.YangName = "node-interface-unicast-peer"
    nodeInterfaceUnicastPeer.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceUnicastPeer.EntityData.ParentYangName = "node-interface-unicast-peers"
    nodeInterfaceUnicastPeer.EntityData.SegmentPath = "node-interface-unicast-peer" + "[interface-name='" + fmt.Sprintf("%v", nodeInterfaceUnicastPeer.InterfaceName) + "']"
    nodeInterfaceUnicastPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceUnicastPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceUnicastPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceUnicastPeer.EntityData.Children = make(map[string]types.YChild)
    nodeInterfaceUnicastPeer.EntityData.Children["peers"] = types.YChild{"Peers", nil}
    for i := range nodeInterfaceUnicastPeer.Peers {
        nodeInterfaceUnicastPeer.EntityData.Children[types.GetSegmentPath(&nodeInterfaceUnicastPeer.Peers[i])] = types.YChild{"Peers", &nodeInterfaceUnicastPeer.Peers[i]}
    }
    nodeInterfaceUnicastPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeInterfaceUnicastPeer.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", nodeInterfaceUnicastPeer.InterfaceName}
    nodeInterfaceUnicastPeer.EntityData.Leafs["name"] = types.YLeaf{"Name", nodeInterfaceUnicastPeer.Name}
    nodeInterfaceUnicastPeer.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", nodeInterfaceUnicastPeer.PortNumber}
    return &(nodeInterfaceUnicastPeer.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers
// Unicast Peers
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers struct {
    EntityData types.CommonEntityData
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

func (peers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "node-interface-unicast-peer"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = make(map[string]types.YChild)
    peers.EntityData.Children["address"] = types.YChild{"Address", &peers.Address}
    peers.EntityData.Children["announce-grant"] = types.YChild{"AnnounceGrant", &peers.AnnounceGrant}
    peers.EntityData.Children["sync-grant"] = types.YChild{"SyncGrant", &peers.SyncGrant}
    peers.EntityData.Children["delay-response-grant"] = types.YChild{"DelayResponseGrant", &peers.DelayResponseGrant}
    peers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peers.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address
// The address of the unicast peer
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address
}

func (address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "peers"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress
// Ethernet MAC address
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address
// IPv6 address
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_AnnounceGrant) GetEntityData() *types.CommonEntityData {
    announceGrant.EntityData.YFilter = announceGrant.YFilter
    announceGrant.EntityData.YangName = "announce-grant"
    announceGrant.EntityData.BundleName = "cisco_ios_xr"
    announceGrant.EntityData.ParentYangName = "peers"
    announceGrant.EntityData.SegmentPath = "announce-grant"
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = make(map[string]types.YChild)
    announceGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    announceGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval}
    announceGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", announceGrant.GrantDuration}
    return &(announceGrant.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant
// Unicast grant information for sync messages
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_SyncGrant) GetEntityData() *types.CommonEntityData {
    syncGrant.EntityData.YFilter = syncGrant.YFilter
    syncGrant.EntityData.YangName = "sync-grant"
    syncGrant.EntityData.BundleName = "cisco_ios_xr"
    syncGrant.EntityData.ParentYangName = "peers"
    syncGrant.EntityData.SegmentPath = "sync-grant"
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = make(map[string]types.YChild)
    syncGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    syncGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval}
    syncGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", syncGrant.GrantDuration}
    return &(syncGrant.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers_DelayResponseGrant) GetEntityData() *types.CommonEntityData {
    delayResponseGrant.EntityData.YFilter = delayResponseGrant.YFilter
    delayResponseGrant.EntityData.YangName = "delay-response-grant"
    delayResponseGrant.EntityData.BundleName = "cisco_ios_xr"
    delayResponseGrant.EntityData.ParentYangName = "peers"
    delayResponseGrant.EntityData.SegmentPath = "delay-response-grant"
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = make(map[string]types.YChild)
    delayResponseGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    delayResponseGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval}
    delayResponseGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration}
    return &(delayResponseGrant.EntityData)
}

// Ptp_Nodes_Node_PacketCounters
// Node packet counter operational data
type Ptp_Nodes_Node_PacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counters.
    Counters Ptp_Nodes_Node_PacketCounters_Counters

    // Drop reasons.
    DropReasons Ptp_Nodes_Node_PacketCounters_DropReasons
}

func (packetCounters *Ptp_Nodes_Node_PacketCounters) GetEntityData() *types.CommonEntityData {
    packetCounters.EntityData.YFilter = packetCounters.YFilter
    packetCounters.EntityData.YangName = "packet-counters"
    packetCounters.EntityData.BundleName = "cisco_ios_xr"
    packetCounters.EntityData.ParentYangName = "node"
    packetCounters.EntityData.SegmentPath = "packet-counters"
    packetCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetCounters.EntityData.Children = make(map[string]types.YChild)
    packetCounters.EntityData.Children["counters"] = types.YChild{"Counters", &packetCounters.Counters}
    packetCounters.EntityData.Children["drop-reasons"] = types.YChild{"DropReasons", &packetCounters.DropReasons}
    packetCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(packetCounters.EntityData)
}

// Ptp_Nodes_Node_PacketCounters_Counters
// Packet counters
type Ptp_Nodes_Node_PacketCounters_Counters struct {
    EntityData types.CommonEntityData
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

func (counters *Ptp_Nodes_Node_PacketCounters_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "packet-counters"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["announce-sent"] = types.YLeaf{"AnnounceSent", counters.AnnounceSent}
    counters.EntityData.Leafs["announce-received"] = types.YLeaf{"AnnounceReceived", counters.AnnounceReceived}
    counters.EntityData.Leafs["announce-dropped"] = types.YLeaf{"AnnounceDropped", counters.AnnounceDropped}
    counters.EntityData.Leafs["sync-sent"] = types.YLeaf{"SyncSent", counters.SyncSent}
    counters.EntityData.Leafs["sync-received"] = types.YLeaf{"SyncReceived", counters.SyncReceived}
    counters.EntityData.Leafs["sync-dropped"] = types.YLeaf{"SyncDropped", counters.SyncDropped}
    counters.EntityData.Leafs["follow-up-sent"] = types.YLeaf{"FollowUpSent", counters.FollowUpSent}
    counters.EntityData.Leafs["follow-up-received"] = types.YLeaf{"FollowUpReceived", counters.FollowUpReceived}
    counters.EntityData.Leafs["follow-up-dropped"] = types.YLeaf{"FollowUpDropped", counters.FollowUpDropped}
    counters.EntityData.Leafs["delay-request-sent"] = types.YLeaf{"DelayRequestSent", counters.DelayRequestSent}
    counters.EntityData.Leafs["delay-request-received"] = types.YLeaf{"DelayRequestReceived", counters.DelayRequestReceived}
    counters.EntityData.Leafs["delay-request-dropped"] = types.YLeaf{"DelayRequestDropped", counters.DelayRequestDropped}
    counters.EntityData.Leafs["delay-response-sent"] = types.YLeaf{"DelayResponseSent", counters.DelayResponseSent}
    counters.EntityData.Leafs["delay-response-received"] = types.YLeaf{"DelayResponseReceived", counters.DelayResponseReceived}
    counters.EntityData.Leafs["delay-response-dropped"] = types.YLeaf{"DelayResponseDropped", counters.DelayResponseDropped}
    counters.EntityData.Leafs["peer-delay-request-sent"] = types.YLeaf{"PeerDelayRequestSent", counters.PeerDelayRequestSent}
    counters.EntityData.Leafs["peer-delay-request-received"] = types.YLeaf{"PeerDelayRequestReceived", counters.PeerDelayRequestReceived}
    counters.EntityData.Leafs["peer-delay-request-dropped"] = types.YLeaf{"PeerDelayRequestDropped", counters.PeerDelayRequestDropped}
    counters.EntityData.Leafs["peer-delay-response-sent"] = types.YLeaf{"PeerDelayResponseSent", counters.PeerDelayResponseSent}
    counters.EntityData.Leafs["peer-delay-response-received"] = types.YLeaf{"PeerDelayResponseReceived", counters.PeerDelayResponseReceived}
    counters.EntityData.Leafs["peer-delay-response-dropped"] = types.YLeaf{"PeerDelayResponseDropped", counters.PeerDelayResponseDropped}
    counters.EntityData.Leafs["peer-delay-response-follow-up-sent"] = types.YLeaf{"PeerDelayResponseFollowUpSent", counters.PeerDelayResponseFollowUpSent}
    counters.EntityData.Leafs["peer-delay-response-follow-up-received"] = types.YLeaf{"PeerDelayResponseFollowUpReceived", counters.PeerDelayResponseFollowUpReceived}
    counters.EntityData.Leafs["peer-delay-response-follow-up-dropped"] = types.YLeaf{"PeerDelayResponseFollowUpDropped", counters.PeerDelayResponseFollowUpDropped}
    counters.EntityData.Leafs["signaling-sent"] = types.YLeaf{"SignalingSent", counters.SignalingSent}
    counters.EntityData.Leafs["signaling-received"] = types.YLeaf{"SignalingReceived", counters.SignalingReceived}
    counters.EntityData.Leafs["signaling-dropped"] = types.YLeaf{"SignalingDropped", counters.SignalingDropped}
    counters.EntityData.Leafs["management-sent"] = types.YLeaf{"ManagementSent", counters.ManagementSent}
    counters.EntityData.Leafs["management-received"] = types.YLeaf{"ManagementReceived", counters.ManagementReceived}
    counters.EntityData.Leafs["management-dropped"] = types.YLeaf{"ManagementDropped", counters.ManagementDropped}
    counters.EntityData.Leafs["other-packets-sent"] = types.YLeaf{"OtherPacketsSent", counters.OtherPacketsSent}
    counters.EntityData.Leafs["other-packets-received"] = types.YLeaf{"OtherPacketsReceived", counters.OtherPacketsReceived}
    counters.EntityData.Leafs["other-packets-dropped"] = types.YLeaf{"OtherPacketsDropped", counters.OtherPacketsDropped}
    counters.EntityData.Leafs["total-packets-sent"] = types.YLeaf{"TotalPacketsSent", counters.TotalPacketsSent}
    counters.EntityData.Leafs["total-packets-received"] = types.YLeaf{"TotalPacketsReceived", counters.TotalPacketsReceived}
    counters.EntityData.Leafs["total-packets-dropped"] = types.YLeaf{"TotalPacketsDropped", counters.TotalPacketsDropped}
    return &(counters.EntityData)
}

// Ptp_Nodes_Node_PacketCounters_DropReasons
// Drop reasons
type Ptp_Nodes_Node_PacketCounters_DropReasons struct {
    EntityData types.CommonEntityData
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

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetEntityData() *types.CommonEntityData {
    dropReasons.EntityData.YFilter = dropReasons.YFilter
    dropReasons.EntityData.YangName = "drop-reasons"
    dropReasons.EntityData.BundleName = "cisco_ios_xr"
    dropReasons.EntityData.ParentYangName = "packet-counters"
    dropReasons.EntityData.SegmentPath = "drop-reasons"
    dropReasons.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropReasons.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropReasons.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropReasons.EntityData.Children = make(map[string]types.YChild)
    dropReasons.EntityData.Leafs = make(map[string]types.YLeaf)
    dropReasons.EntityData.Leafs["not-ready"] = types.YLeaf{"NotReady", dropReasons.NotReady}
    dropReasons.EntityData.Leafs["wrong-domain"] = types.YLeaf{"WrongDomain", dropReasons.WrongDomain}
    dropReasons.EntityData.Leafs["too-short"] = types.YLeaf{"TooShort", dropReasons.TooShort}
    dropReasons.EntityData.Leafs["looped-same-port"] = types.YLeaf{"LoopedSamePort", dropReasons.LoopedSamePort}
    dropReasons.EntityData.Leafs["looped-higher-port"] = types.YLeaf{"LoopedHigherPort", dropReasons.LoopedHigherPort}
    dropReasons.EntityData.Leafs["looped-lower-port"] = types.YLeaf{"LoopedLowerPort", dropReasons.LoopedLowerPort}
    dropReasons.EntityData.Leafs["no-timestamp"] = types.YLeaf{"NoTimestamp", dropReasons.NoTimestamp}
    dropReasons.EntityData.Leafs["zero-timestamp"] = types.YLeaf{"ZeroTimestamp", dropReasons.ZeroTimestamp}
    dropReasons.EntityData.Leafs["invalid-tl-vs"] = types.YLeaf{"InvalidTlVs", dropReasons.InvalidTlVs}
    dropReasons.EntityData.Leafs["not-for-us"] = types.YLeaf{"NotForUs", dropReasons.NotForUs}
    dropReasons.EntityData.Leafs["not-listening"] = types.YLeaf{"NotListening", dropReasons.NotListening}
    dropReasons.EntityData.Leafs["wrong-master"] = types.YLeaf{"WrongMaster", dropReasons.WrongMaster}
    dropReasons.EntityData.Leafs["unknown-master"] = types.YLeaf{"UnknownMaster", dropReasons.UnknownMaster}
    dropReasons.EntityData.Leafs["not-master"] = types.YLeaf{"NotMaster", dropReasons.NotMaster}
    dropReasons.EntityData.Leafs["not-slave"] = types.YLeaf{"NotSlave", dropReasons.NotSlave}
    dropReasons.EntityData.Leafs["not-granted"] = types.YLeaf{"NotGranted", dropReasons.NotGranted}
    dropReasons.EntityData.Leafs["too-slow"] = types.YLeaf{"TooSlow", dropReasons.TooSlow}
    dropReasons.EntityData.Leafs["invalid-packet"] = types.YLeaf{"InvalidPacket", dropReasons.InvalidPacket}
    dropReasons.EntityData.Leafs["wrong-sequence-id"] = types.YLeaf{"WrongSequenceId", dropReasons.WrongSequenceId}
    dropReasons.EntityData.Leafs["no-offload-session"] = types.YLeaf{"NoOffloadSession", dropReasons.NoOffloadSession}
    dropReasons.EntityData.Leafs["not-supported"] = types.YLeaf{"NotSupported", dropReasons.NotSupported}
    dropReasons.EntityData.Leafs["min-clock-class"] = types.YLeaf{"MinClockClass", dropReasons.MinClockClass}
    dropReasons.EntityData.Leafs["g8275-1-incompatible"] = types.YLeaf{"G82751Incompatible", dropReasons.G82751Incompatible}
    dropReasons.EntityData.Leafs["g8275-2-incompatible"] = types.YLeaf{"G82752Incompatible", dropReasons.G82752Incompatible}
    return &(dropReasons.EntityData)
}

// Ptp_InterfaceConfigurationErrors
// Table for interface configuration error
// operational data
type Ptp_InterfaceConfigurationErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface configuration error operational data. The type is slice of
    // Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError.
    InterfaceConfigurationError []Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetEntityData() *types.CommonEntityData {
    interfaceConfigurationErrors.EntityData.YFilter = interfaceConfigurationErrors.YFilter
    interfaceConfigurationErrors.EntityData.YangName = "interface-configuration-errors"
    interfaceConfigurationErrors.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigurationErrors.EntityData.ParentYangName = "ptp"
    interfaceConfigurationErrors.EntityData.SegmentPath = "interface-configuration-errors"
    interfaceConfigurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigurationErrors.EntityData.Children = make(map[string]types.YChild)
    interfaceConfigurationErrors.EntityData.Children["interface-configuration-error"] = types.YChild{"InterfaceConfigurationError", nil}
    for i := range interfaceConfigurationErrors.InterfaceConfigurationError {
        interfaceConfigurationErrors.EntityData.Children[types.GetSegmentPath(&interfaceConfigurationErrors.InterfaceConfigurationError[i])] = types.YChild{"InterfaceConfigurationError", &interfaceConfigurationErrors.InterfaceConfigurationError[i]}
    }
    interfaceConfigurationErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceConfigurationErrors.EntityData)
}

// Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError
// Interface configuration error operational data
type Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetEntityData() *types.CommonEntityData {
    interfaceConfigurationError.EntityData.YFilter = interfaceConfigurationError.YFilter
    interfaceConfigurationError.EntityData.YangName = "interface-configuration-error"
    interfaceConfigurationError.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigurationError.EntityData.ParentYangName = "interface-configuration-errors"
    interfaceConfigurationError.EntityData.SegmentPath = "interface-configuration-error" + "[interface-name='" + fmt.Sprintf("%v", interfaceConfigurationError.InterfaceName) + "']"
    interfaceConfigurationError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigurationError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigurationError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigurationError.EntityData.Children = make(map[string]types.YChild)
    interfaceConfigurationError.EntityData.Children["configuration-errors"] = types.YChild{"ConfigurationErrors", &interfaceConfigurationError.ConfigurationErrors}
    interfaceConfigurationError.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConfigurationError.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceConfigurationError.InterfaceName}
    interfaceConfigurationError.EntityData.Leafs["configuration-profile-name"] = types.YLeaf{"ConfigurationProfileName", interfaceConfigurationError.ConfigurationProfileName}
    interfaceConfigurationError.EntityData.Leafs["clock-profile"] = types.YLeaf{"ClockProfile", interfaceConfigurationError.ClockProfile}
    interfaceConfigurationError.EntityData.Leafs["telecom-clock-type"] = types.YLeaf{"TelecomClockType", interfaceConfigurationError.TelecomClockType}
    interfaceConfigurationError.EntityData.Leafs["restrict-port-state"] = types.YLeaf{"RestrictPortState", interfaceConfigurationError.RestrictPortState}
    return &(interfaceConfigurationError.EntityData)
}

// Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors
// Configuration Errors
type Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors struct {
    EntityData types.CommonEntityData
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

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetEntityData() *types.CommonEntityData {
    configurationErrors.EntityData.YFilter = configurationErrors.YFilter
    configurationErrors.EntityData.YangName = "configuration-errors"
    configurationErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationErrors.EntityData.ParentYangName = "interface-configuration-error"
    configurationErrors.EntityData.SegmentPath = "configuration-errors"
    configurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationErrors.EntityData.Children = make(map[string]types.YChild)
    configurationErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    configurationErrors.EntityData.Leafs["global-ptp"] = types.YLeaf{"GlobalPtp", configurationErrors.GlobalPtp}
    configurationErrors.EntityData.Leafs["ethernet-transport"] = types.YLeaf{"EthernetTransport", configurationErrors.EthernetTransport}
    configurationErrors.EntityData.Leafs["one-step"] = types.YLeaf{"OneStep", configurationErrors.OneStep}
    configurationErrors.EntityData.Leafs["slave"] = types.YLeaf{"Slave", configurationErrors.Slave}
    configurationErrors.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", configurationErrors.Ipv6}
    configurationErrors.EntityData.Leafs["multicast"] = types.YLeaf{"Multicast", configurationErrors.Multicast}
    configurationErrors.EntityData.Leafs["profile-not-global"] = types.YLeaf{"ProfileNotGlobal", configurationErrors.ProfileNotGlobal}
    configurationErrors.EntityData.Leafs["local-priority"] = types.YLeaf{"LocalPriority", configurationErrors.LocalPriority}
    configurationErrors.EntityData.Leafs["profile-ethernet"] = types.YLeaf{"ProfileEthernet", configurationErrors.ProfileEthernet}
    configurationErrors.EntityData.Leafs["profile-ipv4"] = types.YLeaf{"ProfileIpv4", configurationErrors.ProfileIpv4}
    configurationErrors.EntityData.Leafs["profile-ipv6"] = types.YLeaf{"ProfileIpv6", configurationErrors.ProfileIpv6}
    configurationErrors.EntityData.Leafs["profile-unicast"] = types.YLeaf{"ProfileUnicast", configurationErrors.ProfileUnicast}
    configurationErrors.EntityData.Leafs["profile-multicast"] = types.YLeaf{"ProfileMulticast", configurationErrors.ProfileMulticast}
    configurationErrors.EntityData.Leafs["profile-mixed"] = types.YLeaf{"ProfileMixed", configurationErrors.ProfileMixed}
    configurationErrors.EntityData.Leafs["profile-master-unicast"] = types.YLeaf{"ProfileMasterUnicast", configurationErrors.ProfileMasterUnicast}
    configurationErrors.EntityData.Leafs["profile-master-multicast"] = types.YLeaf{"ProfileMasterMulticast", configurationErrors.ProfileMasterMulticast}
    configurationErrors.EntityData.Leafs["profile-master-mixed"] = types.YLeaf{"ProfileMasterMixed", configurationErrors.ProfileMasterMixed}
    configurationErrors.EntityData.Leafs["target-address-ipv4"] = types.YLeaf{"TargetAddressIpv4", configurationErrors.TargetAddressIpv4}
    configurationErrors.EntityData.Leafs["target-address-ipv6"] = types.YLeaf{"TargetAddressIpv6", configurationErrors.TargetAddressIpv6}
    configurationErrors.EntityData.Leafs["profile-port-state"] = types.YLeaf{"ProfilePortState", configurationErrors.ProfilePortState}
    configurationErrors.EntityData.Leafs["profile-announce-interval"] = types.YLeaf{"ProfileAnnounceInterval", configurationErrors.ProfileAnnounceInterval}
    configurationErrors.EntityData.Leafs["profile-sync-interval"] = types.YLeaf{"ProfileSyncInterval", configurationErrors.ProfileSyncInterval}
    configurationErrors.EntityData.Leafs["profile-delay-req-interval"] = types.YLeaf{"ProfileDelayReqInterval", configurationErrors.ProfileDelayReqInterval}
    configurationErrors.EntityData.Leafs["profile-sync-timeout"] = types.YLeaf{"ProfileSyncTimeout", configurationErrors.ProfileSyncTimeout}
    configurationErrors.EntityData.Leafs["profile-delay-resp-timeout"] = types.YLeaf{"ProfileDelayRespTimeout", configurationErrors.ProfileDelayRespTimeout}
    configurationErrors.EntityData.Leafs["invalid-grant-reduction"] = types.YLeaf{"InvalidGrantReduction", configurationErrors.InvalidGrantReduction}
    return &(configurationErrors.EntityData)
}

// Ptp_InterfaceForeignMasters
// Table for interface foreign master clock
// operational data
type Ptp_InterfaceForeignMasters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface foreign master clock operational data. The type is slice of
    // Ptp_InterfaceForeignMasters_InterfaceForeignMaster.
    InterfaceForeignMaster []Ptp_InterfaceForeignMasters_InterfaceForeignMaster
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetEntityData() *types.CommonEntityData {
    interfaceForeignMasters.EntityData.YFilter = interfaceForeignMasters.YFilter
    interfaceForeignMasters.EntityData.YangName = "interface-foreign-masters"
    interfaceForeignMasters.EntityData.BundleName = "cisco_ios_xr"
    interfaceForeignMasters.EntityData.ParentYangName = "ptp"
    interfaceForeignMasters.EntityData.SegmentPath = "interface-foreign-masters"
    interfaceForeignMasters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceForeignMasters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceForeignMasters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceForeignMasters.EntityData.Children = make(map[string]types.YChild)
    interfaceForeignMasters.EntityData.Children["interface-foreign-master"] = types.YChild{"InterfaceForeignMaster", nil}
    for i := range interfaceForeignMasters.InterfaceForeignMaster {
        interfaceForeignMasters.EntityData.Children[types.GetSegmentPath(&interfaceForeignMasters.InterfaceForeignMaster[i])] = types.YChild{"InterfaceForeignMaster", &interfaceForeignMasters.InterfaceForeignMaster[i]}
    }
    interfaceForeignMasters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceForeignMasters.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster
// Interface foreign master clock operational data
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Foreign clocks received on this interface. The type is slice of
    // Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock.
    ForeignClock []Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetEntityData() *types.CommonEntityData {
    interfaceForeignMaster.EntityData.YFilter = interfaceForeignMaster.YFilter
    interfaceForeignMaster.EntityData.YangName = "interface-foreign-master"
    interfaceForeignMaster.EntityData.BundleName = "cisco_ios_xr"
    interfaceForeignMaster.EntityData.ParentYangName = "interface-foreign-masters"
    interfaceForeignMaster.EntityData.SegmentPath = "interface-foreign-master" + "[interface-name='" + fmt.Sprintf("%v", interfaceForeignMaster.InterfaceName) + "']"
    interfaceForeignMaster.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceForeignMaster.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceForeignMaster.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceForeignMaster.EntityData.Children = make(map[string]types.YChild)
    interfaceForeignMaster.EntityData.Children["foreign-clock"] = types.YChild{"ForeignClock", nil}
    for i := range interfaceForeignMaster.ForeignClock {
        interfaceForeignMaster.EntityData.Children[types.GetSegmentPath(&interfaceForeignMaster.ForeignClock[i])] = types.YChild{"ForeignClock", &interfaceForeignMaster.ForeignClock[i]}
    }
    interfaceForeignMaster.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceForeignMaster.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceForeignMaster.InterfaceName}
    interfaceForeignMaster.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", interfaceForeignMaster.PortNumber}
    return &(interfaceForeignMaster.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock
// Foreign clocks received on this interface
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock struct {
    EntityData types.CommonEntityData
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
    ForeignClock Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_

    // The address of the clock.
    Address Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address

    // Unicast grant information for announce messages.
    AnnounceGrant Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant

    // Unicast grant information for sync messages.
    SyncGrant Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant

    // Unicast grant information for delay-response messages.
    DelayResponseGrant Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetEntityData() *types.CommonEntityData {
    foreignClock.EntityData.YFilter = foreignClock.YFilter
    foreignClock.EntityData.YangName = "foreign-clock"
    foreignClock.EntityData.BundleName = "cisco_ios_xr"
    foreignClock.EntityData.ParentYangName = "interface-foreign-master"
    foreignClock.EntityData.SegmentPath = "foreign-clock"
    foreignClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock.EntityData.Children = make(map[string]types.YChild)
    foreignClock.EntityData.Children["foreign-clock"] = types.YChild{"ForeignClock", &foreignClock.ForeignClock}
    foreignClock.EntityData.Children["address"] = types.YChild{"Address", &foreignClock.Address}
    foreignClock.EntityData.Children["announce-grant"] = types.YChild{"AnnounceGrant", &foreignClock.AnnounceGrant}
    foreignClock.EntityData.Children["sync-grant"] = types.YChild{"SyncGrant", &foreignClock.SyncGrant}
    foreignClock.EntityData.Children["delay-response-grant"] = types.YChild{"DelayResponseGrant", &foreignClock.DelayResponseGrant}
    foreignClock.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignClock.EntityData.Leafs["is-qualified"] = types.YLeaf{"IsQualified", foreignClock.IsQualified}
    foreignClock.EntityData.Leafs["is-grandmaster"] = types.YLeaf{"IsGrandmaster", foreignClock.IsGrandmaster}
    foreignClock.EntityData.Leafs["communication-model"] = types.YLeaf{"CommunicationModel", foreignClock.CommunicationModel}
    foreignClock.EntityData.Leafs["is-known"] = types.YLeaf{"IsKnown", foreignClock.IsKnown}
    foreignClock.EntityData.Leafs["time-known-for"] = types.YLeaf{"TimeKnownFor", foreignClock.TimeKnownFor}
    foreignClock.EntityData.Leafs["foreign-domain"] = types.YLeaf{"ForeignDomain", foreignClock.ForeignDomain}
    foreignClock.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", foreignClock.ConfiguredPriority}
    foreignClock.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", foreignClock.ConfiguredClockClass}
    foreignClock.EntityData.Leafs["delay-asymmetry"] = types.YLeaf{"DelayAsymmetry", foreignClock.DelayAsymmetry}
    foreignClock.EntityData.Leafs["ptsf-loss-announce"] = types.YLeaf{"PtsfLossAnnounce", foreignClock.PtsfLossAnnounce}
    foreignClock.EntityData.Leafs["ptsf-loss-sync"] = types.YLeaf{"PtsfLossSync", foreignClock.PtsfLossSync}
    return &(foreignClock.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_
// Foreign clock information
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_ struct {
    EntityData types.CommonEntityData
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
    UtcOffset Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset

    // Receiver.
    Receiver Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Receiver

    // Sender.
    Sender Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Sender
}

func (foreignClock_ *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_) GetEntityData() *types.CommonEntityData {
    foreignClock_.EntityData.YFilter = foreignClock_.YFilter
    foreignClock_.EntityData.YangName = "foreign-clock"
    foreignClock_.EntityData.BundleName = "cisco_ios_xr"
    foreignClock_.EntityData.ParentYangName = "foreign-clock"
    foreignClock_.EntityData.SegmentPath = "foreign-clock"
    foreignClock_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock_.EntityData.Children = make(map[string]types.YChild)
    foreignClock_.EntityData.Children["utc-offset"] = types.YChild{"UtcOffset", &foreignClock_.UtcOffset}
    foreignClock_.EntityData.Children["receiver"] = types.YChild{"Receiver", &foreignClock_.Receiver}
    foreignClock_.EntityData.Children["sender"] = types.YChild{"Sender", &foreignClock_.Sender}
    foreignClock_.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignClock_.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", foreignClock_.ClockId}
    foreignClock_.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", foreignClock_.Priority1}
    foreignClock_.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", foreignClock_.Priority2}
    foreignClock_.EntityData.Leafs["class"] = types.YLeaf{"Class", foreignClock_.Class}
    foreignClock_.EntityData.Leafs["accuracy"] = types.YLeaf{"Accuracy", foreignClock_.Accuracy}
    foreignClock_.EntityData.Leafs["offset-log-variance"] = types.YLeaf{"OffsetLogVariance", foreignClock_.OffsetLogVariance}
    foreignClock_.EntityData.Leafs["steps-removed"] = types.YLeaf{"StepsRemoved", foreignClock_.StepsRemoved}
    foreignClock_.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", foreignClock_.TimeSource}
    foreignClock_.EntityData.Leafs["frequency-traceable"] = types.YLeaf{"FrequencyTraceable", foreignClock_.FrequencyTraceable}
    foreignClock_.EntityData.Leafs["time-traceable"] = types.YLeaf{"TimeTraceable", foreignClock_.TimeTraceable}
    foreignClock_.EntityData.Leafs["timescale"] = types.YLeaf{"Timescale", foreignClock_.Timescale}
    foreignClock_.EntityData.Leafs["leap-seconds"] = types.YLeaf{"LeapSeconds", foreignClock_.LeapSeconds}
    foreignClock_.EntityData.Leafs["local"] = types.YLeaf{"Local", foreignClock_.Local}
    foreignClock_.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", foreignClock_.ConfiguredClockClass}
    foreignClock_.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", foreignClock_.ConfiguredPriority}
    return &(foreignClock_.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset
// UTC offset
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "foreign-clock"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = make(map[string]types.YChild)
    utcOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffset.EntityData.Leafs["current-offset"] = types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset}
    utcOffset.EntityData.Leafs["offset-valid"] = types.YLeaf{"OffsetValid", utcOffset.OffsetValid}
    return &(utcOffset.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Receiver
// Receiver
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "foreign-clock"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = make(map[string]types.YChild)
    receiver.EntityData.Leafs = make(map[string]types.YLeaf)
    receiver.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", receiver.ClockId}
    receiver.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", receiver.PortNumber}
    return &(receiver.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Sender
// Sender
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock__Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "foreign-clock"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = make(map[string]types.YChild)
    sender.EntityData.Leafs = make(map[string]types.YLeaf)
    sender.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", sender.ClockId}
    sender.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", sender.PortNumber}
    return &(sender.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address
// The address of the clock
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address
}

func (address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "foreign-clock"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address
// IPv6 address
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_AnnounceGrant) GetEntityData() *types.CommonEntityData {
    announceGrant.EntityData.YFilter = announceGrant.YFilter
    announceGrant.EntityData.YangName = "announce-grant"
    announceGrant.EntityData.BundleName = "cisco_ios_xr"
    announceGrant.EntityData.ParentYangName = "foreign-clock"
    announceGrant.EntityData.SegmentPath = "announce-grant"
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = make(map[string]types.YChild)
    announceGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    announceGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval}
    announceGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", announceGrant.GrantDuration}
    return &(announceGrant.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant
// Unicast grant information for sync messages
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_SyncGrant) GetEntityData() *types.CommonEntityData {
    syncGrant.EntityData.YFilter = syncGrant.YFilter
    syncGrant.EntityData.YangName = "sync-grant"
    syncGrant.EntityData.BundleName = "cisco_ios_xr"
    syncGrant.EntityData.ParentYangName = "foreign-clock"
    syncGrant.EntityData.SegmentPath = "sync-grant"
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = make(map[string]types.YChild)
    syncGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    syncGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval}
    syncGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", syncGrant.GrantDuration}
    return &(syncGrant.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_DelayResponseGrant) GetEntityData() *types.CommonEntityData {
    delayResponseGrant.EntityData.YFilter = delayResponseGrant.YFilter
    delayResponseGrant.EntityData.YangName = "delay-response-grant"
    delayResponseGrant.EntityData.BundleName = "cisco_ios_xr"
    delayResponseGrant.EntityData.ParentYangName = "foreign-clock"
    delayResponseGrant.EntityData.SegmentPath = "delay-response-grant"
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = make(map[string]types.YChild)
    delayResponseGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    delayResponseGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval}
    delayResponseGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration}
    return &(delayResponseGrant.EntityData)
}

// Ptp_LocalClock
// Local clock operational data
type Ptp_LocalClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The PTP domain of the local clock. The type is interface{} with range:
    // 0..255.
    Domain interface{}

    // Local clock.
    ClockProperties Ptp_LocalClock_ClockProperties
}

func (localClock *Ptp_LocalClock) GetEntityData() *types.CommonEntityData {
    localClock.EntityData.YFilter = localClock.YFilter
    localClock.EntityData.YangName = "local-clock"
    localClock.EntityData.BundleName = "cisco_ios_xr"
    localClock.EntityData.ParentYangName = "ptp"
    localClock.EntityData.SegmentPath = "local-clock"
    localClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localClock.EntityData.Children = make(map[string]types.YChild)
    localClock.EntityData.Children["clock-properties"] = types.YChild{"ClockProperties", &localClock.ClockProperties}
    localClock.EntityData.Leafs = make(map[string]types.YLeaf)
    localClock.EntityData.Leafs["domain"] = types.YLeaf{"Domain", localClock.Domain}
    return &(localClock.EntityData)
}

// Ptp_LocalClock_ClockProperties
// Local clock
type Ptp_LocalClock_ClockProperties struct {
    EntityData types.CommonEntityData
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

func (clockProperties *Ptp_LocalClock_ClockProperties) GetEntityData() *types.CommonEntityData {
    clockProperties.EntityData.YFilter = clockProperties.YFilter
    clockProperties.EntityData.YangName = "clock-properties"
    clockProperties.EntityData.BundleName = "cisco_ios_xr"
    clockProperties.EntityData.ParentYangName = "local-clock"
    clockProperties.EntityData.SegmentPath = "clock-properties"
    clockProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockProperties.EntityData.Children = make(map[string]types.YChild)
    clockProperties.EntityData.Children["utc-offset"] = types.YChild{"UtcOffset", &clockProperties.UtcOffset}
    clockProperties.EntityData.Children["receiver"] = types.YChild{"Receiver", &clockProperties.Receiver}
    clockProperties.EntityData.Children["sender"] = types.YChild{"Sender", &clockProperties.Sender}
    clockProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    clockProperties.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", clockProperties.ClockId}
    clockProperties.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", clockProperties.Priority1}
    clockProperties.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", clockProperties.Priority2}
    clockProperties.EntityData.Leafs["class"] = types.YLeaf{"Class", clockProperties.Class}
    clockProperties.EntityData.Leafs["accuracy"] = types.YLeaf{"Accuracy", clockProperties.Accuracy}
    clockProperties.EntityData.Leafs["offset-log-variance"] = types.YLeaf{"OffsetLogVariance", clockProperties.OffsetLogVariance}
    clockProperties.EntityData.Leafs["steps-removed"] = types.YLeaf{"StepsRemoved", clockProperties.StepsRemoved}
    clockProperties.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", clockProperties.TimeSource}
    clockProperties.EntityData.Leafs["frequency-traceable"] = types.YLeaf{"FrequencyTraceable", clockProperties.FrequencyTraceable}
    clockProperties.EntityData.Leafs["time-traceable"] = types.YLeaf{"TimeTraceable", clockProperties.TimeTraceable}
    clockProperties.EntityData.Leafs["timescale"] = types.YLeaf{"Timescale", clockProperties.Timescale}
    clockProperties.EntityData.Leafs["leap-seconds"] = types.YLeaf{"LeapSeconds", clockProperties.LeapSeconds}
    clockProperties.EntityData.Leafs["local"] = types.YLeaf{"Local", clockProperties.Local}
    clockProperties.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", clockProperties.ConfiguredClockClass}
    clockProperties.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", clockProperties.ConfiguredPriority}
    return &(clockProperties.EntityData)
}

// Ptp_LocalClock_ClockProperties_UtcOffset
// UTC offset
type Ptp_LocalClock_ClockProperties_UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_LocalClock_ClockProperties_UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "clock-properties"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = make(map[string]types.YChild)
    utcOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffset.EntityData.Leafs["current-offset"] = types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset}
    utcOffset.EntityData.Leafs["offset-valid"] = types.YLeaf{"OffsetValid", utcOffset.OffsetValid}
    return &(utcOffset.EntityData)
}

// Ptp_LocalClock_ClockProperties_Receiver
// Receiver
type Ptp_LocalClock_ClockProperties_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_LocalClock_ClockProperties_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "clock-properties"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = make(map[string]types.YChild)
    receiver.EntityData.Leafs = make(map[string]types.YLeaf)
    receiver.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", receiver.ClockId}
    receiver.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", receiver.PortNumber}
    return &(receiver.EntityData)
}

// Ptp_LocalClock_ClockProperties_Sender
// Sender
type Ptp_LocalClock_ClockProperties_Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_LocalClock_ClockProperties_Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "clock-properties"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = make(map[string]types.YChild)
    sender.EntityData.Leafs = make(map[string]types.YLeaf)
    sender.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", sender.ClockId}
    sender.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", sender.PortNumber}
    return &(sender.EntityData)
}

// Ptp_InterfacePacketCounters
// Table for interface packet counter operational
// data
type Ptp_InterfacePacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface packet counter operational data. The type is slice of
    // Ptp_InterfacePacketCounters_InterfacePacketCounter.
    InterfacePacketCounter []Ptp_InterfacePacketCounters_InterfacePacketCounter
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetEntityData() *types.CommonEntityData {
    interfacePacketCounters.EntityData.YFilter = interfacePacketCounters.YFilter
    interfacePacketCounters.EntityData.YangName = "interface-packet-counters"
    interfacePacketCounters.EntityData.BundleName = "cisco_ios_xr"
    interfacePacketCounters.EntityData.ParentYangName = "ptp"
    interfacePacketCounters.EntityData.SegmentPath = "interface-packet-counters"
    interfacePacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacePacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacePacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacePacketCounters.EntityData.Children = make(map[string]types.YChild)
    interfacePacketCounters.EntityData.Children["interface-packet-counter"] = types.YChild{"InterfacePacketCounter", nil}
    for i := range interfacePacketCounters.InterfacePacketCounter {
        interfacePacketCounters.EntityData.Children[types.GetSegmentPath(&interfacePacketCounters.InterfacePacketCounter[i])] = types.YChild{"InterfacePacketCounter", &interfacePacketCounters.InterfacePacketCounter[i]}
    }
    interfacePacketCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfacePacketCounters.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter
// Interface packet counter operational data
type Ptp_InterfacePacketCounters_InterfacePacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Packet counters.
    Counters Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters

    // Packet counters for each peer on this interface. The type is slice of
    // Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter.
    PeerCounter []Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetEntityData() *types.CommonEntityData {
    interfacePacketCounter.EntityData.YFilter = interfacePacketCounter.YFilter
    interfacePacketCounter.EntityData.YangName = "interface-packet-counter"
    interfacePacketCounter.EntityData.BundleName = "cisco_ios_xr"
    interfacePacketCounter.EntityData.ParentYangName = "interface-packet-counters"
    interfacePacketCounter.EntityData.SegmentPath = "interface-packet-counter" + "[interface-name='" + fmt.Sprintf("%v", interfacePacketCounter.InterfaceName) + "']"
    interfacePacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacePacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacePacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacePacketCounter.EntityData.Children = make(map[string]types.YChild)
    interfacePacketCounter.EntityData.Children["counters"] = types.YChild{"Counters", &interfacePacketCounter.Counters}
    interfacePacketCounter.EntityData.Children["peer-counter"] = types.YChild{"PeerCounter", nil}
    for i := range interfacePacketCounter.PeerCounter {
        interfacePacketCounter.EntityData.Children[types.GetSegmentPath(&interfacePacketCounter.PeerCounter[i])] = types.YChild{"PeerCounter", &interfacePacketCounter.PeerCounter[i]}
    }
    interfacePacketCounter.EntityData.Leafs = make(map[string]types.YLeaf)
    interfacePacketCounter.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfacePacketCounter.InterfaceName}
    return &(interfacePacketCounter.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters
// Packet counters
type Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters struct {
    EntityData types.CommonEntityData
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

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "interface-packet-counter"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["announce-sent"] = types.YLeaf{"AnnounceSent", counters.AnnounceSent}
    counters.EntityData.Leafs["announce-received"] = types.YLeaf{"AnnounceReceived", counters.AnnounceReceived}
    counters.EntityData.Leafs["announce-dropped"] = types.YLeaf{"AnnounceDropped", counters.AnnounceDropped}
    counters.EntityData.Leafs["sync-sent"] = types.YLeaf{"SyncSent", counters.SyncSent}
    counters.EntityData.Leafs["sync-received"] = types.YLeaf{"SyncReceived", counters.SyncReceived}
    counters.EntityData.Leafs["sync-dropped"] = types.YLeaf{"SyncDropped", counters.SyncDropped}
    counters.EntityData.Leafs["follow-up-sent"] = types.YLeaf{"FollowUpSent", counters.FollowUpSent}
    counters.EntityData.Leafs["follow-up-received"] = types.YLeaf{"FollowUpReceived", counters.FollowUpReceived}
    counters.EntityData.Leafs["follow-up-dropped"] = types.YLeaf{"FollowUpDropped", counters.FollowUpDropped}
    counters.EntityData.Leafs["delay-request-sent"] = types.YLeaf{"DelayRequestSent", counters.DelayRequestSent}
    counters.EntityData.Leafs["delay-request-received"] = types.YLeaf{"DelayRequestReceived", counters.DelayRequestReceived}
    counters.EntityData.Leafs["delay-request-dropped"] = types.YLeaf{"DelayRequestDropped", counters.DelayRequestDropped}
    counters.EntityData.Leafs["delay-response-sent"] = types.YLeaf{"DelayResponseSent", counters.DelayResponseSent}
    counters.EntityData.Leafs["delay-response-received"] = types.YLeaf{"DelayResponseReceived", counters.DelayResponseReceived}
    counters.EntityData.Leafs["delay-response-dropped"] = types.YLeaf{"DelayResponseDropped", counters.DelayResponseDropped}
    counters.EntityData.Leafs["peer-delay-request-sent"] = types.YLeaf{"PeerDelayRequestSent", counters.PeerDelayRequestSent}
    counters.EntityData.Leafs["peer-delay-request-received"] = types.YLeaf{"PeerDelayRequestReceived", counters.PeerDelayRequestReceived}
    counters.EntityData.Leafs["peer-delay-request-dropped"] = types.YLeaf{"PeerDelayRequestDropped", counters.PeerDelayRequestDropped}
    counters.EntityData.Leafs["peer-delay-response-sent"] = types.YLeaf{"PeerDelayResponseSent", counters.PeerDelayResponseSent}
    counters.EntityData.Leafs["peer-delay-response-received"] = types.YLeaf{"PeerDelayResponseReceived", counters.PeerDelayResponseReceived}
    counters.EntityData.Leafs["peer-delay-response-dropped"] = types.YLeaf{"PeerDelayResponseDropped", counters.PeerDelayResponseDropped}
    counters.EntityData.Leafs["peer-delay-response-follow-up-sent"] = types.YLeaf{"PeerDelayResponseFollowUpSent", counters.PeerDelayResponseFollowUpSent}
    counters.EntityData.Leafs["peer-delay-response-follow-up-received"] = types.YLeaf{"PeerDelayResponseFollowUpReceived", counters.PeerDelayResponseFollowUpReceived}
    counters.EntityData.Leafs["peer-delay-response-follow-up-dropped"] = types.YLeaf{"PeerDelayResponseFollowUpDropped", counters.PeerDelayResponseFollowUpDropped}
    counters.EntityData.Leafs["signaling-sent"] = types.YLeaf{"SignalingSent", counters.SignalingSent}
    counters.EntityData.Leafs["signaling-received"] = types.YLeaf{"SignalingReceived", counters.SignalingReceived}
    counters.EntityData.Leafs["signaling-dropped"] = types.YLeaf{"SignalingDropped", counters.SignalingDropped}
    counters.EntityData.Leafs["management-sent"] = types.YLeaf{"ManagementSent", counters.ManagementSent}
    counters.EntityData.Leafs["management-received"] = types.YLeaf{"ManagementReceived", counters.ManagementReceived}
    counters.EntityData.Leafs["management-dropped"] = types.YLeaf{"ManagementDropped", counters.ManagementDropped}
    counters.EntityData.Leafs["other-packets-sent"] = types.YLeaf{"OtherPacketsSent", counters.OtherPacketsSent}
    counters.EntityData.Leafs["other-packets-received"] = types.YLeaf{"OtherPacketsReceived", counters.OtherPacketsReceived}
    counters.EntityData.Leafs["other-packets-dropped"] = types.YLeaf{"OtherPacketsDropped", counters.OtherPacketsDropped}
    counters.EntityData.Leafs["total-packets-sent"] = types.YLeaf{"TotalPacketsSent", counters.TotalPacketsSent}
    counters.EntityData.Leafs["total-packets-received"] = types.YLeaf{"TotalPacketsReceived", counters.TotalPacketsReceived}
    counters.EntityData.Leafs["total-packets-dropped"] = types.YLeaf{"TotalPacketsDropped", counters.TotalPacketsDropped}
    return &(counters.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter
// Packet counters for each peer on this interface
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer address.
    Address Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address

    // Packet counters.
    Counters Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters
}

func (peerCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter) GetEntityData() *types.CommonEntityData {
    peerCounter.EntityData.YFilter = peerCounter.YFilter
    peerCounter.EntityData.YangName = "peer-counter"
    peerCounter.EntityData.BundleName = "cisco_ios_xr"
    peerCounter.EntityData.ParentYangName = "interface-packet-counter"
    peerCounter.EntityData.SegmentPath = "peer-counter"
    peerCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerCounter.EntityData.Children = make(map[string]types.YChild)
    peerCounter.EntityData.Children["address"] = types.YChild{"Address", &peerCounter.Address}
    peerCounter.EntityData.Children["counters"] = types.YChild{"Counters", &peerCounter.Counters}
    peerCounter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerCounter.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address
// Peer address
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address
}

func (address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "peer-counter"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address
// IPv6 address
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters
// Packet counters
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters struct {
    EntityData types.CommonEntityData
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

func (counters *Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "peer-counter"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["announce-sent"] = types.YLeaf{"AnnounceSent", counters.AnnounceSent}
    counters.EntityData.Leafs["announce-received"] = types.YLeaf{"AnnounceReceived", counters.AnnounceReceived}
    counters.EntityData.Leafs["announce-dropped"] = types.YLeaf{"AnnounceDropped", counters.AnnounceDropped}
    counters.EntityData.Leafs["sync-sent"] = types.YLeaf{"SyncSent", counters.SyncSent}
    counters.EntityData.Leafs["sync-received"] = types.YLeaf{"SyncReceived", counters.SyncReceived}
    counters.EntityData.Leafs["sync-dropped"] = types.YLeaf{"SyncDropped", counters.SyncDropped}
    counters.EntityData.Leafs["follow-up-sent"] = types.YLeaf{"FollowUpSent", counters.FollowUpSent}
    counters.EntityData.Leafs["follow-up-received"] = types.YLeaf{"FollowUpReceived", counters.FollowUpReceived}
    counters.EntityData.Leafs["follow-up-dropped"] = types.YLeaf{"FollowUpDropped", counters.FollowUpDropped}
    counters.EntityData.Leafs["delay-request-sent"] = types.YLeaf{"DelayRequestSent", counters.DelayRequestSent}
    counters.EntityData.Leafs["delay-request-received"] = types.YLeaf{"DelayRequestReceived", counters.DelayRequestReceived}
    counters.EntityData.Leafs["delay-request-dropped"] = types.YLeaf{"DelayRequestDropped", counters.DelayRequestDropped}
    counters.EntityData.Leafs["delay-response-sent"] = types.YLeaf{"DelayResponseSent", counters.DelayResponseSent}
    counters.EntityData.Leafs["delay-response-received"] = types.YLeaf{"DelayResponseReceived", counters.DelayResponseReceived}
    counters.EntityData.Leafs["delay-response-dropped"] = types.YLeaf{"DelayResponseDropped", counters.DelayResponseDropped}
    counters.EntityData.Leafs["peer-delay-request-sent"] = types.YLeaf{"PeerDelayRequestSent", counters.PeerDelayRequestSent}
    counters.EntityData.Leafs["peer-delay-request-received"] = types.YLeaf{"PeerDelayRequestReceived", counters.PeerDelayRequestReceived}
    counters.EntityData.Leafs["peer-delay-request-dropped"] = types.YLeaf{"PeerDelayRequestDropped", counters.PeerDelayRequestDropped}
    counters.EntityData.Leafs["peer-delay-response-sent"] = types.YLeaf{"PeerDelayResponseSent", counters.PeerDelayResponseSent}
    counters.EntityData.Leafs["peer-delay-response-received"] = types.YLeaf{"PeerDelayResponseReceived", counters.PeerDelayResponseReceived}
    counters.EntityData.Leafs["peer-delay-response-dropped"] = types.YLeaf{"PeerDelayResponseDropped", counters.PeerDelayResponseDropped}
    counters.EntityData.Leafs["peer-delay-response-follow-up-sent"] = types.YLeaf{"PeerDelayResponseFollowUpSent", counters.PeerDelayResponseFollowUpSent}
    counters.EntityData.Leafs["peer-delay-response-follow-up-received"] = types.YLeaf{"PeerDelayResponseFollowUpReceived", counters.PeerDelayResponseFollowUpReceived}
    counters.EntityData.Leafs["peer-delay-response-follow-up-dropped"] = types.YLeaf{"PeerDelayResponseFollowUpDropped", counters.PeerDelayResponseFollowUpDropped}
    counters.EntityData.Leafs["signaling-sent"] = types.YLeaf{"SignalingSent", counters.SignalingSent}
    counters.EntityData.Leafs["signaling-received"] = types.YLeaf{"SignalingReceived", counters.SignalingReceived}
    counters.EntityData.Leafs["signaling-dropped"] = types.YLeaf{"SignalingDropped", counters.SignalingDropped}
    counters.EntityData.Leafs["management-sent"] = types.YLeaf{"ManagementSent", counters.ManagementSent}
    counters.EntityData.Leafs["management-received"] = types.YLeaf{"ManagementReceived", counters.ManagementReceived}
    counters.EntityData.Leafs["management-dropped"] = types.YLeaf{"ManagementDropped", counters.ManagementDropped}
    counters.EntityData.Leafs["other-packets-sent"] = types.YLeaf{"OtherPacketsSent", counters.OtherPacketsSent}
    counters.EntityData.Leafs["other-packets-received"] = types.YLeaf{"OtherPacketsReceived", counters.OtherPacketsReceived}
    counters.EntityData.Leafs["other-packets-dropped"] = types.YLeaf{"OtherPacketsDropped", counters.OtherPacketsDropped}
    counters.EntityData.Leafs["total-packets-sent"] = types.YLeaf{"TotalPacketsSent", counters.TotalPacketsSent}
    counters.EntityData.Leafs["total-packets-received"] = types.YLeaf{"TotalPacketsReceived", counters.TotalPacketsReceived}
    counters.EntityData.Leafs["total-packets-dropped"] = types.YLeaf{"TotalPacketsDropped", counters.TotalPacketsDropped}
    return &(counters.EntityData)
}

// Ptp_AdvertisedClock
// Advertised clock operational data
type Ptp_AdvertisedClock struct {
    EntityData types.CommonEntityData
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

func (advertisedClock *Ptp_AdvertisedClock) GetEntityData() *types.CommonEntityData {
    advertisedClock.EntityData.YFilter = advertisedClock.YFilter
    advertisedClock.EntityData.YangName = "advertised-clock"
    advertisedClock.EntityData.BundleName = "cisco_ios_xr"
    advertisedClock.EntityData.ParentYangName = "ptp"
    advertisedClock.EntityData.SegmentPath = "advertised-clock"
    advertisedClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertisedClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertisedClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertisedClock.EntityData.Children = make(map[string]types.YChild)
    advertisedClock.EntityData.Children["clock-properties"] = types.YChild{"ClockProperties", &advertisedClock.ClockProperties}
    advertisedClock.EntityData.Leafs = make(map[string]types.YLeaf)
    advertisedClock.EntityData.Leafs["domain"] = types.YLeaf{"Domain", advertisedClock.Domain}
    advertisedClock.EntityData.Leafs["time-source-configured"] = types.YLeaf{"TimeSourceConfigured", advertisedClock.TimeSourceConfigured}
    advertisedClock.EntityData.Leafs["received-time-source"] = types.YLeaf{"ReceivedTimeSource", advertisedClock.ReceivedTimeSource}
    advertisedClock.EntityData.Leafs["timescale-configured"] = types.YLeaf{"TimescaleConfigured", advertisedClock.TimescaleConfigured}
    advertisedClock.EntityData.Leafs["received-timescale"] = types.YLeaf{"ReceivedTimescale", advertisedClock.ReceivedTimescale}
    return &(advertisedClock.EntityData)
}

// Ptp_AdvertisedClock_ClockProperties
// Advertised Clock
type Ptp_AdvertisedClock_ClockProperties struct {
    EntityData types.CommonEntityData
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

func (clockProperties *Ptp_AdvertisedClock_ClockProperties) GetEntityData() *types.CommonEntityData {
    clockProperties.EntityData.YFilter = clockProperties.YFilter
    clockProperties.EntityData.YangName = "clock-properties"
    clockProperties.EntityData.BundleName = "cisco_ios_xr"
    clockProperties.EntityData.ParentYangName = "advertised-clock"
    clockProperties.EntityData.SegmentPath = "clock-properties"
    clockProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockProperties.EntityData.Children = make(map[string]types.YChild)
    clockProperties.EntityData.Children["utc-offset"] = types.YChild{"UtcOffset", &clockProperties.UtcOffset}
    clockProperties.EntityData.Children["receiver"] = types.YChild{"Receiver", &clockProperties.Receiver}
    clockProperties.EntityData.Children["sender"] = types.YChild{"Sender", &clockProperties.Sender}
    clockProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    clockProperties.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", clockProperties.ClockId}
    clockProperties.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", clockProperties.Priority1}
    clockProperties.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", clockProperties.Priority2}
    clockProperties.EntityData.Leafs["class"] = types.YLeaf{"Class", clockProperties.Class}
    clockProperties.EntityData.Leafs["accuracy"] = types.YLeaf{"Accuracy", clockProperties.Accuracy}
    clockProperties.EntityData.Leafs["offset-log-variance"] = types.YLeaf{"OffsetLogVariance", clockProperties.OffsetLogVariance}
    clockProperties.EntityData.Leafs["steps-removed"] = types.YLeaf{"StepsRemoved", clockProperties.StepsRemoved}
    clockProperties.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", clockProperties.TimeSource}
    clockProperties.EntityData.Leafs["frequency-traceable"] = types.YLeaf{"FrequencyTraceable", clockProperties.FrequencyTraceable}
    clockProperties.EntityData.Leafs["time-traceable"] = types.YLeaf{"TimeTraceable", clockProperties.TimeTraceable}
    clockProperties.EntityData.Leafs["timescale"] = types.YLeaf{"Timescale", clockProperties.Timescale}
    clockProperties.EntityData.Leafs["leap-seconds"] = types.YLeaf{"LeapSeconds", clockProperties.LeapSeconds}
    clockProperties.EntityData.Leafs["local"] = types.YLeaf{"Local", clockProperties.Local}
    clockProperties.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", clockProperties.ConfiguredClockClass}
    clockProperties.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", clockProperties.ConfiguredPriority}
    return &(clockProperties.EntityData)
}

// Ptp_AdvertisedClock_ClockProperties_UtcOffset
// UTC offset
type Ptp_AdvertisedClock_ClockProperties_UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_AdvertisedClock_ClockProperties_UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "clock-properties"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = make(map[string]types.YChild)
    utcOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffset.EntityData.Leafs["current-offset"] = types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset}
    utcOffset.EntityData.Leafs["offset-valid"] = types.YLeaf{"OffsetValid", utcOffset.OffsetValid}
    return &(utcOffset.EntityData)
}

// Ptp_AdvertisedClock_ClockProperties_Receiver
// Receiver
type Ptp_AdvertisedClock_ClockProperties_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_AdvertisedClock_ClockProperties_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "clock-properties"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = make(map[string]types.YChild)
    receiver.EntityData.Leafs = make(map[string]types.YLeaf)
    receiver.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", receiver.ClockId}
    receiver.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", receiver.PortNumber}
    return &(receiver.EntityData)
}

// Ptp_AdvertisedClock_ClockProperties_Sender
// Sender
type Ptp_AdvertisedClock_ClockProperties_Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_AdvertisedClock_ClockProperties_Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "clock-properties"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = make(map[string]types.YChild)
    sender.EntityData.Leafs = make(map[string]types.YLeaf)
    sender.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", sender.ClockId}
    sender.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", sender.PortNumber}
    return &(sender.EntityData)
}

// Ptp_Interfaces
// Table for interface operational data
type Ptp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface operational data. The type is slice of Ptp_Interfaces_Interface_.
    Interface_ []Ptp_Interfaces_Interface
}

func (interfaces *Ptp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ptp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Ptp_Interfaces_Interface
// Interface operational data
type Ptp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // IPv4 address, if IPv4 encapsulation is being used. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (self *Ptp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &self.MacAddress}
    self.EntityData.Children["master-table"] = types.YChild{"MasterTable", nil}
    for i := range self.MasterTable {
        self.EntityData.Children[types.GetSegmentPath(&self.MasterTable[i])] = types.YChild{"MasterTable", &self.MasterTable[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["port-state"] = types.YLeaf{"PortState", self.PortState}
    self.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", self.PortNumber}
    self.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", self.LineState}
    self.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", self.Encapsulation}
    self.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", self.Ipv6Address}
    self.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", self.Ipv4Address}
    self.EntityData.Leafs["two-step"] = types.YLeaf{"TwoStep", self.TwoStep}
    self.EntityData.Leafs["communication-model"] = types.YLeaf{"CommunicationModel", self.CommunicationModel}
    self.EntityData.Leafs["log-sync-interval"] = types.YLeaf{"LogSyncInterval", self.LogSyncInterval}
    self.EntityData.Leafs["log-announce-interval"] = types.YLeaf{"LogAnnounceInterval", self.LogAnnounceInterval}
    self.EntityData.Leafs["announce-timeout"] = types.YLeaf{"AnnounceTimeout", self.AnnounceTimeout}
    self.EntityData.Leafs["log-min-delay-request-interval"] = types.YLeaf{"LogMinDelayRequestInterval", self.LogMinDelayRequestInterval}
    self.EntityData.Leafs["configured-port-state"] = types.YLeaf{"ConfiguredPortState", self.ConfiguredPortState}
    self.EntityData.Leafs["supports-one-step"] = types.YLeaf{"SupportsOneStep", self.SupportsOneStep}
    self.EntityData.Leafs["supports-two-step"] = types.YLeaf{"SupportsTwoStep", self.SupportsTwoStep}
    self.EntityData.Leafs["supports-ethernet"] = types.YLeaf{"SupportsEthernet", self.SupportsEthernet}
    self.EntityData.Leafs["supports-multicast"] = types.YLeaf{"SupportsMulticast", self.SupportsMulticast}
    self.EntityData.Leafs["supports-ipv6"] = types.YLeaf{"SupportsIpv6", self.SupportsIpv6}
    self.EntityData.Leafs["supports-slave"] = types.YLeaf{"SupportsSlave", self.SupportsSlave}
    self.EntityData.Leafs["supports-source-ip"] = types.YLeaf{"SupportsSourceIp", self.SupportsSourceIp}
    self.EntityData.Leafs["max-sync-rate"] = types.YLeaf{"MaxSyncRate", self.MaxSyncRate}
    self.EntityData.Leafs["event-cos"] = types.YLeaf{"EventCos", self.EventCos}
    self.EntityData.Leafs["general-cos"] = types.YLeaf{"GeneralCos", self.GeneralCos}
    self.EntityData.Leafs["event-dscp"] = types.YLeaf{"EventDscp", self.EventDscp}
    self.EntityData.Leafs["general-dscp"] = types.YLeaf{"GeneralDscp", self.GeneralDscp}
    self.EntityData.Leafs["unicast-peers"] = types.YLeaf{"UnicastPeers", self.UnicastPeers}
    self.EntityData.Leafs["local-priority"] = types.YLeaf{"LocalPriority", self.LocalPriority}
    self.EntityData.Leafs["signal-fail"] = types.YLeaf{"SignalFail", self.SignalFail}
    return &(self.EntityData)
}

// Ptp_Interfaces_Interface_MacAddress
// MAC address, if Ethernet encapsulation is being
// used
type Ptp_Interfaces_Interface_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Interfaces_Interface_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "interface"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Interfaces_Interface_MasterTable
// The interface's master table
type Ptp_Interfaces_Interface_MasterTable struct {
    EntityData types.CommonEntityData
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

func (masterTable *Ptp_Interfaces_Interface_MasterTable) GetEntityData() *types.CommonEntityData {
    masterTable.EntityData.YFilter = masterTable.YFilter
    masterTable.EntityData.YangName = "master-table"
    masterTable.EntityData.BundleName = "cisco_ios_xr"
    masterTable.EntityData.ParentYangName = "interface"
    masterTable.EntityData.SegmentPath = "master-table"
    masterTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterTable.EntityData.Children = make(map[string]types.YChild)
    masterTable.EntityData.Children["address"] = types.YChild{"Address", &masterTable.Address}
    masterTable.EntityData.Leafs = make(map[string]types.YLeaf)
    masterTable.EntityData.Leafs["communication-model"] = types.YLeaf{"CommunicationModel", masterTable.CommunicationModel}
    masterTable.EntityData.Leafs["priority"] = types.YLeaf{"Priority", masterTable.Priority}
    masterTable.EntityData.Leafs["known"] = types.YLeaf{"Known", masterTable.Known}
    masterTable.EntityData.Leafs["qualified"] = types.YLeaf{"Qualified", masterTable.Qualified}
    masterTable.EntityData.Leafs["is-grandmaster"] = types.YLeaf{"IsGrandmaster", masterTable.IsGrandmaster}
    masterTable.EntityData.Leafs["ptsf-loss-announce"] = types.YLeaf{"PtsfLossAnnounce", masterTable.PtsfLossAnnounce}
    masterTable.EntityData.Leafs["ptsf-loss-sync"] = types.YLeaf{"PtsfLossSync", masterTable.PtsfLossSync}
    masterTable.EntityData.Leafs["is-nonnegotiated"] = types.YLeaf{"IsNonnegotiated", masterTable.IsNonnegotiated}
    return &(masterTable.EntityData)
}

// Ptp_Interfaces_Interface_MasterTable_Address
// The address of the master clock
type Ptp_Interfaces_Interface_MasterTable_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Interfaces_Interface_MasterTable_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address
}

func (address *Ptp_Interfaces_Interface_MasterTable_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "master-table"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_Interfaces_Interface_MasterTable_Address_MacAddress
// Ethernet MAC address
type Ptp_Interfaces_Interface_MasterTable_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Interfaces_Interface_MasterTable_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address
// IPv6 address
type Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Interfaces_Interface_MasterTable_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_Dataset
// Global PTP datasets
type Ptp_Dataset struct {
    EntityData types.CommonEntityData
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

func (dataset *Ptp_Dataset) GetEntityData() *types.CommonEntityData {
    dataset.EntityData.YFilter = dataset.YFilter
    dataset.EntityData.YangName = "dataset"
    dataset.EntityData.BundleName = "cisco_ios_xr"
    dataset.EntityData.ParentYangName = "ptp"
    dataset.EntityData.SegmentPath = "dataset"
    dataset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataset.EntityData.Children = make(map[string]types.YChild)
    dataset.EntityData.Children["default-ds"] = types.YChild{"DefaultDs", &dataset.DefaultDs}
    dataset.EntityData.Children["current-ds"] = types.YChild{"CurrentDs", &dataset.CurrentDs}
    dataset.EntityData.Children["parent-ds"] = types.YChild{"ParentDs", &dataset.ParentDs}
    dataset.EntityData.Children["port-dses"] = types.YChild{"PortDses", &dataset.PortDses}
    dataset.EntityData.Children["time-properties-ds"] = types.YChild{"TimePropertiesDs", &dataset.TimePropertiesDs}
    dataset.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataset.EntityData)
}

// Ptp_Dataset_DefaultDs
// defaultDS information as described in IEEE
// 1588-2008
type Ptp_Dataset_DefaultDs struct {
    EntityData types.CommonEntityData
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

func (defaultDs *Ptp_Dataset_DefaultDs) GetEntityData() *types.CommonEntityData {
    defaultDs.EntityData.YFilter = defaultDs.YFilter
    defaultDs.EntityData.YangName = "default-ds"
    defaultDs.EntityData.BundleName = "cisco_ios_xr"
    defaultDs.EntityData.ParentYangName = "dataset"
    defaultDs.EntityData.SegmentPath = "default-ds"
    defaultDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultDs.EntityData.Children = make(map[string]types.YChild)
    defaultDs.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultDs.EntityData.Leafs["two-step-flag"] = types.YLeaf{"TwoStepFlag", defaultDs.TwoStepFlag}
    defaultDs.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", defaultDs.ClockId}
    defaultDs.EntityData.Leafs["number-ports"] = types.YLeaf{"NumberPorts", defaultDs.NumberPorts}
    defaultDs.EntityData.Leafs["clock-class"] = types.YLeaf{"ClockClass", defaultDs.ClockClass}
    defaultDs.EntityData.Leafs["clock-accuracy"] = types.YLeaf{"ClockAccuracy", defaultDs.ClockAccuracy}
    defaultDs.EntityData.Leafs["oslv"] = types.YLeaf{"Oslv", defaultDs.Oslv}
    defaultDs.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", defaultDs.Priority1}
    defaultDs.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", defaultDs.Priority2}
    defaultDs.EntityData.Leafs["domain-number"] = types.YLeaf{"DomainNumber", defaultDs.DomainNumber}
    defaultDs.EntityData.Leafs["slave-only"] = types.YLeaf{"SlaveOnly", defaultDs.SlaveOnly}
    defaultDs.EntityData.Leafs["local-priority"] = types.YLeaf{"LocalPriority", defaultDs.LocalPriority}
    defaultDs.EntityData.Leafs["signal-fail"] = types.YLeaf{"SignalFail", defaultDs.SignalFail}
    return &(defaultDs.EntityData)
}

// Ptp_Dataset_CurrentDs
// currentDS information as described in IEEE
// 1588-2008
type Ptp_Dataset_CurrentDs struct {
    EntityData types.CommonEntityData
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

func (currentDs *Ptp_Dataset_CurrentDs) GetEntityData() *types.CommonEntityData {
    currentDs.EntityData.YFilter = currentDs.YFilter
    currentDs.EntityData.YangName = "current-ds"
    currentDs.EntityData.BundleName = "cisco_ios_xr"
    currentDs.EntityData.ParentYangName = "dataset"
    currentDs.EntityData.SegmentPath = "current-ds"
    currentDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentDs.EntityData.Children = make(map[string]types.YChild)
    currentDs.EntityData.Leafs = make(map[string]types.YLeaf)
    currentDs.EntityData.Leafs["steps-removed"] = types.YLeaf{"StepsRemoved", currentDs.StepsRemoved}
    currentDs.EntityData.Leafs["offset-from-master"] = types.YLeaf{"OffsetFromMaster", currentDs.OffsetFromMaster}
    currentDs.EntityData.Leafs["mean-path-delay"] = types.YLeaf{"MeanPathDelay", currentDs.MeanPathDelay}
    return &(currentDs.EntityData)
}

// Ptp_Dataset_ParentDs
// parentDS information as described in IEEE
// 1588-2008
type Ptp_Dataset_ParentDs struct {
    EntityData types.CommonEntityData
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

func (parentDs *Ptp_Dataset_ParentDs) GetEntityData() *types.CommonEntityData {
    parentDs.EntityData.YFilter = parentDs.YFilter
    parentDs.EntityData.YangName = "parent-ds"
    parentDs.EntityData.BundleName = "cisco_ios_xr"
    parentDs.EntityData.ParentYangName = "dataset"
    parentDs.EntityData.SegmentPath = "parent-ds"
    parentDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentDs.EntityData.Children = make(map[string]types.YChild)
    parentDs.EntityData.Leafs = make(map[string]types.YLeaf)
    parentDs.EntityData.Leafs["parent-clock-id"] = types.YLeaf{"ParentClockId", parentDs.ParentClockId}
    parentDs.EntityData.Leafs["parent-port-number"] = types.YLeaf{"ParentPortNumber", parentDs.ParentPortNumber}
    parentDs.EntityData.Leafs["parent-stats"] = types.YLeaf{"ParentStats", parentDs.ParentStats}
    parentDs.EntityData.Leafs["observed-parent-oslv"] = types.YLeaf{"ObservedParentOslv", parentDs.ObservedParentOslv}
    parentDs.EntityData.Leafs["observed-parent-clock-phase-change-rate"] = types.YLeaf{"ObservedParentClockPhaseChangeRate", parentDs.ObservedParentClockPhaseChangeRate}
    parentDs.EntityData.Leafs["gm-clock-id"] = types.YLeaf{"GmClockId", parentDs.GmClockId}
    parentDs.EntityData.Leafs["gm-clock-class"] = types.YLeaf{"GmClockClass", parentDs.GmClockClass}
    parentDs.EntityData.Leafs["gm-clock-accuracy"] = types.YLeaf{"GmClockAccuracy", parentDs.GmClockAccuracy}
    parentDs.EntityData.Leafs["gmoslv"] = types.YLeaf{"Gmoslv", parentDs.Gmoslv}
    parentDs.EntityData.Leafs["gm-priority1"] = types.YLeaf{"GmPriority1", parentDs.GmPriority1}
    parentDs.EntityData.Leafs["gm-priority2"] = types.YLeaf{"GmPriority2", parentDs.GmPriority2}
    return &(parentDs.EntityData)
}

// Ptp_Dataset_PortDses
// Table for portDS information
type Ptp_Dataset_PortDses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortDS information. The type is slice of Ptp_Dataset_PortDses_PortDs.
    PortDs []Ptp_Dataset_PortDses_PortDs
}

func (portDses *Ptp_Dataset_PortDses) GetEntityData() *types.CommonEntityData {
    portDses.EntityData.YFilter = portDses.YFilter
    portDses.EntityData.YangName = "port-dses"
    portDses.EntityData.BundleName = "cisco_ios_xr"
    portDses.EntityData.ParentYangName = "dataset"
    portDses.EntityData.SegmentPath = "port-dses"
    portDses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portDses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portDses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portDses.EntityData.Children = make(map[string]types.YChild)
    portDses.EntityData.Children["port-ds"] = types.YChild{"PortDs", nil}
    for i := range portDses.PortDs {
        portDses.EntityData.Children[types.GetSegmentPath(&portDses.PortDs[i])] = types.YChild{"PortDs", &portDses.PortDs[i]}
    }
    portDses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portDses.EntityData)
}

// Ptp_Dataset_PortDses_PortDs
// PortDS information
type Ptp_Dataset_PortDses_PortDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (portDs *Ptp_Dataset_PortDses_PortDs) GetEntityData() *types.CommonEntityData {
    portDs.EntityData.YFilter = portDs.YFilter
    portDs.EntityData.YangName = "port-ds"
    portDs.EntityData.BundleName = "cisco_ios_xr"
    portDs.EntityData.ParentYangName = "port-dses"
    portDs.EntityData.SegmentPath = "port-ds" + "[interface-name='" + fmt.Sprintf("%v", portDs.InterfaceName) + "']"
    portDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portDs.EntityData.Children = make(map[string]types.YChild)
    portDs.EntityData.Leafs = make(map[string]types.YLeaf)
    portDs.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", portDs.InterfaceName}
    portDs.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", portDs.ClockId}
    portDs.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", portDs.PortNumber}
    portDs.EntityData.Leafs["port-state"] = types.YLeaf{"PortState", portDs.PortState}
    portDs.EntityData.Leafs["log-min-delay-req-interval"] = types.YLeaf{"LogMinDelayReqInterval", portDs.LogMinDelayReqInterval}
    portDs.EntityData.Leafs["peer-mean-path-delay"] = types.YLeaf{"PeerMeanPathDelay", portDs.PeerMeanPathDelay}
    portDs.EntityData.Leafs["log-announce-interval"] = types.YLeaf{"LogAnnounceInterval", portDs.LogAnnounceInterval}
    portDs.EntityData.Leafs["annoucne-receipt-timeout"] = types.YLeaf{"AnnoucneReceiptTimeout", portDs.AnnoucneReceiptTimeout}
    portDs.EntityData.Leafs["log-sync-interval"] = types.YLeaf{"LogSyncInterval", portDs.LogSyncInterval}
    portDs.EntityData.Leafs["delay-mechanism"] = types.YLeaf{"DelayMechanism", portDs.DelayMechanism}
    portDs.EntityData.Leafs["log-min-p-delay-req-interval"] = types.YLeaf{"LogMinPDelayReqInterval", portDs.LogMinPDelayReqInterval}
    portDs.EntityData.Leafs["version-number"] = types.YLeaf{"VersionNumber", portDs.VersionNumber}
    portDs.EntityData.Leafs["local-priority"] = types.YLeaf{"LocalPriority", portDs.LocalPriority}
    portDs.EntityData.Leafs["master-only"] = types.YLeaf{"MasterOnly", portDs.MasterOnly}
    portDs.EntityData.Leafs["signal-fail"] = types.YLeaf{"SignalFail", portDs.SignalFail}
    return &(portDs.EntityData)
}

// Ptp_Dataset_TimePropertiesDs
// timePropertiesDS information as described in
// IEEE 1588-2008
type Ptp_Dataset_TimePropertiesDs struct {
    EntityData types.CommonEntityData
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

func (timePropertiesDs *Ptp_Dataset_TimePropertiesDs) GetEntityData() *types.CommonEntityData {
    timePropertiesDs.EntityData.YFilter = timePropertiesDs.YFilter
    timePropertiesDs.EntityData.YangName = "time-properties-ds"
    timePropertiesDs.EntityData.BundleName = "cisco_ios_xr"
    timePropertiesDs.EntityData.ParentYangName = "dataset"
    timePropertiesDs.EntityData.SegmentPath = "time-properties-ds"
    timePropertiesDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timePropertiesDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timePropertiesDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timePropertiesDs.EntityData.Children = make(map[string]types.YChild)
    timePropertiesDs.EntityData.Leafs = make(map[string]types.YLeaf)
    timePropertiesDs.EntityData.Leafs["current-utc-offset"] = types.YLeaf{"CurrentUtcOffset", timePropertiesDs.CurrentUtcOffset}
    timePropertiesDs.EntityData.Leafs["current-utc-offset-valid"] = types.YLeaf{"CurrentUtcOffsetValid", timePropertiesDs.CurrentUtcOffsetValid}
    timePropertiesDs.EntityData.Leafs["leap59"] = types.YLeaf{"Leap59", timePropertiesDs.Leap59}
    timePropertiesDs.EntityData.Leafs["leap61"] = types.YLeaf{"Leap61", timePropertiesDs.Leap61}
    timePropertiesDs.EntityData.Leafs["time-traceable"] = types.YLeaf{"TimeTraceable", timePropertiesDs.TimeTraceable}
    timePropertiesDs.EntityData.Leafs["frequency-traceable"] = types.YLeaf{"FrequencyTraceable", timePropertiesDs.FrequencyTraceable}
    timePropertiesDs.EntityData.Leafs["ptp-timescale"] = types.YLeaf{"PtpTimescale", timePropertiesDs.PtpTimescale}
    timePropertiesDs.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", timePropertiesDs.TimeSource}
    return &(timePropertiesDs.EntityData)
}

// Ptp_GlobalConfigurationError
// Global configuration error operational data
type Ptp_GlobalConfigurationError struct {
    EntityData types.CommonEntityData
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

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetEntityData() *types.CommonEntityData {
    globalConfigurationError.EntityData.YFilter = globalConfigurationError.YFilter
    globalConfigurationError.EntityData.YangName = "global-configuration-error"
    globalConfigurationError.EntityData.BundleName = "cisco_ios_xr"
    globalConfigurationError.EntityData.ParentYangName = "ptp"
    globalConfigurationError.EntityData.SegmentPath = "global-configuration-error"
    globalConfigurationError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalConfigurationError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalConfigurationError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalConfigurationError.EntityData.Children = make(map[string]types.YChild)
    globalConfigurationError.EntityData.Children["configuration-errors"] = types.YChild{"ConfigurationErrors", &globalConfigurationError.ConfigurationErrors}
    globalConfigurationError.EntityData.Leafs = make(map[string]types.YLeaf)
    globalConfigurationError.EntityData.Leafs["clock-profile"] = types.YLeaf{"ClockProfile", globalConfigurationError.ClockProfile}
    globalConfigurationError.EntityData.Leafs["clock-profile-set"] = types.YLeaf{"ClockProfileSet", globalConfigurationError.ClockProfileSet}
    globalConfigurationError.EntityData.Leafs["telecom-clock-type"] = types.YLeaf{"TelecomClockType", globalConfigurationError.TelecomClockType}
    globalConfigurationError.EntityData.Leafs["domain-number"] = types.YLeaf{"DomainNumber", globalConfigurationError.DomainNumber}
    globalConfigurationError.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", globalConfigurationError.Priority2}
    return &(globalConfigurationError.EntityData)
}

// Ptp_GlobalConfigurationError_ConfigurationErrors
// Configuration Errors
type Ptp_GlobalConfigurationError_ConfigurationErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain not compatible with configured profile. The type is bool.
    Domain interface{}

    // Priority1 configuration is not compatible with configured profile. The type
    // is bool.
    ProfilePriority1Config interface{}

    // Priority2 value is not compatible with configured profile. The type is
    // bool.
    ProfilePriority2Value interface{}

    // Leap seconds configuration contains an invalid UTC offset change. The type
    // is bool.
    UtcOffsetChange interface{}
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetEntityData() *types.CommonEntityData {
    configurationErrors.EntityData.YFilter = configurationErrors.YFilter
    configurationErrors.EntityData.YangName = "configuration-errors"
    configurationErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationErrors.EntityData.ParentYangName = "global-configuration-error"
    configurationErrors.EntityData.SegmentPath = "configuration-errors"
    configurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationErrors.EntityData.Children = make(map[string]types.YChild)
    configurationErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    configurationErrors.EntityData.Leafs["domain"] = types.YLeaf{"Domain", configurationErrors.Domain}
    configurationErrors.EntityData.Leafs["profile-priority1-config"] = types.YLeaf{"ProfilePriority1Config", configurationErrors.ProfilePriority1Config}
    configurationErrors.EntityData.Leafs["profile-priority2-value"] = types.YLeaf{"ProfilePriority2Value", configurationErrors.ProfilePriority2Value}
    configurationErrors.EntityData.Leafs["utc-offset-change"] = types.YLeaf{"UtcOffsetChange", configurationErrors.UtcOffsetChange}
    return &(configurationErrors.EntityData)
}

// Ptp_Grandmaster
// Grandmaster clock operational data
type Ptp_Grandmaster struct {
    EntityData types.CommonEntityData
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

func (grandmaster *Ptp_Grandmaster) GetEntityData() *types.CommonEntityData {
    grandmaster.EntityData.YFilter = grandmaster.YFilter
    grandmaster.EntityData.YangName = "grandmaster"
    grandmaster.EntityData.BundleName = "cisco_ios_xr"
    grandmaster.EntityData.ParentYangName = "ptp"
    grandmaster.EntityData.SegmentPath = "grandmaster"
    grandmaster.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grandmaster.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grandmaster.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grandmaster.EntityData.Children = make(map[string]types.YChild)
    grandmaster.EntityData.Children["clock-properties"] = types.YChild{"ClockProperties", &grandmaster.ClockProperties}
    grandmaster.EntityData.Children["address"] = types.YChild{"Address", &grandmaster.Address}
    grandmaster.EntityData.Leafs = make(map[string]types.YLeaf)
    grandmaster.EntityData.Leafs["used-for-time"] = types.YLeaf{"UsedForTime", grandmaster.UsedForTime}
    grandmaster.EntityData.Leafs["used-for-frequency"] = types.YLeaf{"UsedForFrequency", grandmaster.UsedForFrequency}
    grandmaster.EntityData.Leafs["known-for-time"] = types.YLeaf{"KnownForTime", grandmaster.KnownForTime}
    grandmaster.EntityData.Leafs["domain"] = types.YLeaf{"Domain", grandmaster.Domain}
    return &(grandmaster.EntityData)
}

// Ptp_Grandmaster_ClockProperties
// Grandmaster clock
type Ptp_Grandmaster_ClockProperties struct {
    EntityData types.CommonEntityData
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

func (clockProperties *Ptp_Grandmaster_ClockProperties) GetEntityData() *types.CommonEntityData {
    clockProperties.EntityData.YFilter = clockProperties.YFilter
    clockProperties.EntityData.YangName = "clock-properties"
    clockProperties.EntityData.BundleName = "cisco_ios_xr"
    clockProperties.EntityData.ParentYangName = "grandmaster"
    clockProperties.EntityData.SegmentPath = "clock-properties"
    clockProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockProperties.EntityData.Children = make(map[string]types.YChild)
    clockProperties.EntityData.Children["utc-offset"] = types.YChild{"UtcOffset", &clockProperties.UtcOffset}
    clockProperties.EntityData.Children["receiver"] = types.YChild{"Receiver", &clockProperties.Receiver}
    clockProperties.EntityData.Children["sender"] = types.YChild{"Sender", &clockProperties.Sender}
    clockProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    clockProperties.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", clockProperties.ClockId}
    clockProperties.EntityData.Leafs["priority1"] = types.YLeaf{"Priority1", clockProperties.Priority1}
    clockProperties.EntityData.Leafs["priority2"] = types.YLeaf{"Priority2", clockProperties.Priority2}
    clockProperties.EntityData.Leafs["class"] = types.YLeaf{"Class", clockProperties.Class}
    clockProperties.EntityData.Leafs["accuracy"] = types.YLeaf{"Accuracy", clockProperties.Accuracy}
    clockProperties.EntityData.Leafs["offset-log-variance"] = types.YLeaf{"OffsetLogVariance", clockProperties.OffsetLogVariance}
    clockProperties.EntityData.Leafs["steps-removed"] = types.YLeaf{"StepsRemoved", clockProperties.StepsRemoved}
    clockProperties.EntityData.Leafs["time-source"] = types.YLeaf{"TimeSource", clockProperties.TimeSource}
    clockProperties.EntityData.Leafs["frequency-traceable"] = types.YLeaf{"FrequencyTraceable", clockProperties.FrequencyTraceable}
    clockProperties.EntityData.Leafs["time-traceable"] = types.YLeaf{"TimeTraceable", clockProperties.TimeTraceable}
    clockProperties.EntityData.Leafs["timescale"] = types.YLeaf{"Timescale", clockProperties.Timescale}
    clockProperties.EntityData.Leafs["leap-seconds"] = types.YLeaf{"LeapSeconds", clockProperties.LeapSeconds}
    clockProperties.EntityData.Leafs["local"] = types.YLeaf{"Local", clockProperties.Local}
    clockProperties.EntityData.Leafs["configured-clock-class"] = types.YLeaf{"ConfiguredClockClass", clockProperties.ConfiguredClockClass}
    clockProperties.EntityData.Leafs["configured-priority"] = types.YLeaf{"ConfiguredPriority", clockProperties.ConfiguredPriority}
    return &(clockProperties.EntityData)
}

// Ptp_Grandmaster_ClockProperties_UtcOffset
// UTC offset
type Ptp_Grandmaster_ClockProperties_UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_Grandmaster_ClockProperties_UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "clock-properties"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = make(map[string]types.YChild)
    utcOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffset.EntityData.Leafs["current-offset"] = types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset}
    utcOffset.EntityData.Leafs["offset-valid"] = types.YLeaf{"OffsetValid", utcOffset.OffsetValid}
    return &(utcOffset.EntityData)
}

// Ptp_Grandmaster_ClockProperties_Receiver
// Receiver
type Ptp_Grandmaster_ClockProperties_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_Grandmaster_ClockProperties_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "clock-properties"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = make(map[string]types.YChild)
    receiver.EntityData.Leafs = make(map[string]types.YLeaf)
    receiver.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", receiver.ClockId}
    receiver.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", receiver.PortNumber}
    return &(receiver.EntityData)
}

// Ptp_Grandmaster_ClockProperties_Sender
// Sender
type Ptp_Grandmaster_ClockProperties_Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_Grandmaster_ClockProperties_Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "clock-properties"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = make(map[string]types.YChild)
    sender.EntityData.Leafs = make(map[string]types.YLeaf)
    sender.EntityData.Leafs["clock-id"] = types.YLeaf{"ClockId", sender.ClockId}
    sender.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", sender.PortNumber}
    return &(sender.EntityData)
}

// Ptp_Grandmaster_Address
// The grandmaster's address information
type Ptp_Grandmaster_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_Grandmaster_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_Grandmaster_Address_Ipv6Address
}

func (address *Ptp_Grandmaster_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "grandmaster"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_Grandmaster_Address_MacAddress
// Ethernet MAC address
type Ptp_Grandmaster_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_Grandmaster_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_Grandmaster_Address_Ipv6Address
// IPv6 address
type Ptp_Grandmaster_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_Grandmaster_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_InterfaceUnicastPeers
// Table for interface unicast peers operational
// data
type Ptp_InterfaceUnicastPeers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface unicast peers operational data. The type is slice of
    // Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer.
    InterfaceUnicastPeer []Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetEntityData() *types.CommonEntityData {
    interfaceUnicastPeers.EntityData.YFilter = interfaceUnicastPeers.YFilter
    interfaceUnicastPeers.EntityData.YangName = "interface-unicast-peers"
    interfaceUnicastPeers.EntityData.BundleName = "cisco_ios_xr"
    interfaceUnicastPeers.EntityData.ParentYangName = "ptp"
    interfaceUnicastPeers.EntityData.SegmentPath = "interface-unicast-peers"
    interfaceUnicastPeers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceUnicastPeers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceUnicastPeers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceUnicastPeers.EntityData.Children = make(map[string]types.YChild)
    interfaceUnicastPeers.EntityData.Children["interface-unicast-peer"] = types.YChild{"InterfaceUnicastPeer", nil}
    for i := range interfaceUnicastPeers.InterfaceUnicastPeer {
        interfaceUnicastPeers.EntityData.Children[types.GetSegmentPath(&interfaceUnicastPeers.InterfaceUnicastPeer[i])] = types.YChild{"InterfaceUnicastPeer", &interfaceUnicastPeers.InterfaceUnicastPeer[i]}
    }
    interfaceUnicastPeers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceUnicastPeers.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer
// Interface unicast peers operational data
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Unicast Peers. The type is slice of
    // Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers.
    Peers []Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetEntityData() *types.CommonEntityData {
    interfaceUnicastPeer.EntityData.YFilter = interfaceUnicastPeer.YFilter
    interfaceUnicastPeer.EntityData.YangName = "interface-unicast-peer"
    interfaceUnicastPeer.EntityData.BundleName = "cisco_ios_xr"
    interfaceUnicastPeer.EntityData.ParentYangName = "interface-unicast-peers"
    interfaceUnicastPeer.EntityData.SegmentPath = "interface-unicast-peer" + "[interface-name='" + fmt.Sprintf("%v", interfaceUnicastPeer.InterfaceName) + "']"
    interfaceUnicastPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceUnicastPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceUnicastPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceUnicastPeer.EntityData.Children = make(map[string]types.YChild)
    interfaceUnicastPeer.EntityData.Children["peers"] = types.YChild{"Peers", nil}
    for i := range interfaceUnicastPeer.Peers {
        interfaceUnicastPeer.EntityData.Children[types.GetSegmentPath(&interfaceUnicastPeer.Peers[i])] = types.YChild{"Peers", &interfaceUnicastPeer.Peers[i]}
    }
    interfaceUnicastPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceUnicastPeer.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceUnicastPeer.InterfaceName}
    interfaceUnicastPeer.EntityData.Leafs["name"] = types.YLeaf{"Name", interfaceUnicastPeer.Name}
    interfaceUnicastPeer.EntityData.Leafs["port-number"] = types.YLeaf{"PortNumber", interfaceUnicastPeer.PortNumber}
    return &(interfaceUnicastPeer.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers
// Unicast Peers
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers struct {
    EntityData types.CommonEntityData
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

func (peers *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "interface-unicast-peer"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = make(map[string]types.YChild)
    peers.EntityData.Children["address"] = types.YChild{"Address", &peers.Address}
    peers.EntityData.Children["announce-grant"] = types.YChild{"AnnounceGrant", &peers.AnnounceGrant}
    peers.EntityData.Children["sync-grant"] = types.YChild{"SyncGrant", &peers.SyncGrant}
    peers.EntityData.Children["delay-response-grant"] = types.YChild{"DelayResponseGrant", &peers.DelayResponseGrant}
    peers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peers.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address
// The address of the unicast peer
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation. The type is PtpBagEncap.
    Encapsulation interface{}

    // Unknown address type. The type is bool.
    AddressUnknown interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // Ethernet MAC address.
    MacAddress Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address
}

func (address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "peers"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["mac-address"] = types.YChild{"MacAddress", &address.MacAddress}
    address.EntityData.Children["ipv6-address"] = types.YChild{"Ipv6Address", &address.Ipv6Address}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", address.Encapsulation}
    address.EntityData.Leafs["address-unknown"] = types.YLeaf{"AddressUnknown", address.AddressUnknown}
    address.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", address.Ipv4Address}
    return &(address.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = make(map[string]types.YChild)
    macAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddress.EntityData.Leafs["macaddr"] = types.YLeaf{"Macaddr", macAddress.Macaddr}
    return &(macAddress.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address
// IPv6 address
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = make(map[string]types.YChild)
    ipv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Address.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address}
    return &(ipv6Address.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant
// Unicast grant information for announce messages
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (announceGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_AnnounceGrant) GetEntityData() *types.CommonEntityData {
    announceGrant.EntityData.YFilter = announceGrant.YFilter
    announceGrant.EntityData.YangName = "announce-grant"
    announceGrant.EntityData.BundleName = "cisco_ios_xr"
    announceGrant.EntityData.ParentYangName = "peers"
    announceGrant.EntityData.SegmentPath = "announce-grant"
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = make(map[string]types.YChild)
    announceGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    announceGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval}
    announceGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", announceGrant.GrantDuration}
    return &(announceGrant.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant
// Unicast grant information for sync messages
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (syncGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_SyncGrant) GetEntityData() *types.CommonEntityData {
    syncGrant.EntityData.YFilter = syncGrant.YFilter
    syncGrant.EntityData.YangName = "sync-grant"
    syncGrant.EntityData.BundleName = "cisco_ios_xr"
    syncGrant.EntityData.ParentYangName = "peers"
    syncGrant.EntityData.SegmentPath = "sync-grant"
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = make(map[string]types.YChild)
    syncGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    syncGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval}
    syncGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", syncGrant.GrantDuration}
    return &(syncGrant.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant
// Unicast grant information for delay-response
// messages
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log of the interval which has been granted. The type is interface{} with
    // range: -128..127.
    LogGrantInterval interface{}

    // Duraction of the grant. The type is interface{} with range: 0..4294967295.
    GrantDuration interface{}
}

func (delayResponseGrant *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers_DelayResponseGrant) GetEntityData() *types.CommonEntityData {
    delayResponseGrant.EntityData.YFilter = delayResponseGrant.YFilter
    delayResponseGrant.EntityData.YangName = "delay-response-grant"
    delayResponseGrant.EntityData.BundleName = "cisco_ios_xr"
    delayResponseGrant.EntityData.ParentYangName = "peers"
    delayResponseGrant.EntityData.SegmentPath = "delay-response-grant"
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = make(map[string]types.YChild)
    delayResponseGrant.EntityData.Leafs = make(map[string]types.YLeaf)
    delayResponseGrant.EntityData.Leafs["log-grant-interval"] = types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval}
    delayResponseGrant.EntityData.Leafs["grant-duration"] = types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration}
    return &(delayResponseGrant.EntityData)
}

// Ptp_UtcOffsetInfo
// UTC offset information
type Ptp_UtcOffsetInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The current source of the UTC offset information. The type is interface{}
    // with range: 0..255.
    SourceType interface{}

    // The URL of the file containing leap second information. The type is string.
    SourceFile interface{}

    // Source file expiry timestamp, in seconds since UNIX epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    SourceExpiryDate interface{}

    // Source file polling frequency, in days. The type is interface{} with range:
    // 0..4294967295. Units are day.
    PollingFrequency interface{}

    // The current UTC offset information.
    CurrentOffsetInfo Ptp_UtcOffsetInfo_CurrentOffsetInfo

    // The UTC offset information recovered from the current grandmaster.
    CurrentGmOffsetInfo Ptp_UtcOffsetInfo_CurrentGmOffsetInfo

    // The currently configured UTC offset information.
    ConfiguredOffsetInfo Ptp_UtcOffsetInfo_ConfiguredOffsetInfo

    // The UTC offset information recovered from the prevous grandmaster.
    PreviousGmOffsetInfo Ptp_UtcOffsetInfo_PreviousGmOffsetInfo

    // The UTC offset information taken from the hardware.
    HardwareOffsetInfo Ptp_UtcOffsetInfo_HardwareOffsetInfo

    // The upcoming leap second advertised by the grandmaster (if there is one).
    GmLeapSecond Ptp_UtcOffsetInfo_GmLeapSecond

    // The list of upcoming configured leap second updates. The type is slice of
    // Ptp_UtcOffsetInfo_ConfiguredLeapSecond.
    ConfiguredLeapSecond []Ptp_UtcOffsetInfo_ConfiguredLeapSecond
}

func (utcOffsetInfo *Ptp_UtcOffsetInfo) GetEntityData() *types.CommonEntityData {
    utcOffsetInfo.EntityData.YFilter = utcOffsetInfo.YFilter
    utcOffsetInfo.EntityData.YangName = "utc-offset-info"
    utcOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    utcOffsetInfo.EntityData.ParentYangName = "ptp"
    utcOffsetInfo.EntityData.SegmentPath = "utc-offset-info"
    utcOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffsetInfo.EntityData.Children = make(map[string]types.YChild)
    utcOffsetInfo.EntityData.Children["current-offset-info"] = types.YChild{"CurrentOffsetInfo", &utcOffsetInfo.CurrentOffsetInfo}
    utcOffsetInfo.EntityData.Children["current-gm-offset-info"] = types.YChild{"CurrentGmOffsetInfo", &utcOffsetInfo.CurrentGmOffsetInfo}
    utcOffsetInfo.EntityData.Children["configured-offset-info"] = types.YChild{"ConfiguredOffsetInfo", &utcOffsetInfo.ConfiguredOffsetInfo}
    utcOffsetInfo.EntityData.Children["previous-gm-offset-info"] = types.YChild{"PreviousGmOffsetInfo", &utcOffsetInfo.PreviousGmOffsetInfo}
    utcOffsetInfo.EntityData.Children["hardware-offset-info"] = types.YChild{"HardwareOffsetInfo", &utcOffsetInfo.HardwareOffsetInfo}
    utcOffsetInfo.EntityData.Children["gm-leap-second"] = types.YChild{"GmLeapSecond", &utcOffsetInfo.GmLeapSecond}
    utcOffsetInfo.EntityData.Children["configured-leap-second"] = types.YChild{"ConfiguredLeapSecond", nil}
    for i := range utcOffsetInfo.ConfiguredLeapSecond {
        utcOffsetInfo.EntityData.Children[types.GetSegmentPath(&utcOffsetInfo.ConfiguredLeapSecond[i])] = types.YChild{"ConfiguredLeapSecond", &utcOffsetInfo.ConfiguredLeapSecond[i]}
    }
    utcOffsetInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    utcOffsetInfo.EntityData.Leafs["source-type"] = types.YLeaf{"SourceType", utcOffsetInfo.SourceType}
    utcOffsetInfo.EntityData.Leafs["source-file"] = types.YLeaf{"SourceFile", utcOffsetInfo.SourceFile}
    utcOffsetInfo.EntityData.Leafs["source-expiry-date"] = types.YLeaf{"SourceExpiryDate", utcOffsetInfo.SourceExpiryDate}
    utcOffsetInfo.EntityData.Leafs["polling-frequency"] = types.YLeaf{"PollingFrequency", utcOffsetInfo.PollingFrequency}
    return &(utcOffsetInfo.EntityData)
}

// Ptp_UtcOffsetInfo_CurrentOffsetInfo
// The current UTC offset information
type Ptp_UtcOffsetInfo_CurrentOffsetInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // Is the UTC offset valid?. The type is bool.
    Valid interface{}

    // Indicates the duration of the final minute of the current day. The type is
    // interface{} with range: 0..255.
    Flag interface{}
}

func (currentOffsetInfo *Ptp_UtcOffsetInfo_CurrentOffsetInfo) GetEntityData() *types.CommonEntityData {
    currentOffsetInfo.EntityData.YFilter = currentOffsetInfo.YFilter
    currentOffsetInfo.EntityData.YangName = "current-offset-info"
    currentOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    currentOffsetInfo.EntityData.ParentYangName = "utc-offset-info"
    currentOffsetInfo.EntityData.SegmentPath = "current-offset-info"
    currentOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentOffsetInfo.EntityData.Children = make(map[string]types.YChild)
    currentOffsetInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    currentOffsetInfo.EntityData.Leafs["offset"] = types.YLeaf{"Offset", currentOffsetInfo.Offset}
    currentOffsetInfo.EntityData.Leafs["valid"] = types.YLeaf{"Valid", currentOffsetInfo.Valid}
    currentOffsetInfo.EntityData.Leafs["flag"] = types.YLeaf{"Flag", currentOffsetInfo.Flag}
    return &(currentOffsetInfo.EntityData)
}

// Ptp_UtcOffsetInfo_CurrentGmOffsetInfo
// The UTC offset information recovered from the
// current grandmaster
type Ptp_UtcOffsetInfo_CurrentGmOffsetInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // Is the UTC offset valid?. The type is bool.
    Valid interface{}

    // Indicates the duration of the final minute of the current day. The type is
    // interface{} with range: 0..255.
    Flag interface{}
}

func (currentGmOffsetInfo *Ptp_UtcOffsetInfo_CurrentGmOffsetInfo) GetEntityData() *types.CommonEntityData {
    currentGmOffsetInfo.EntityData.YFilter = currentGmOffsetInfo.YFilter
    currentGmOffsetInfo.EntityData.YangName = "current-gm-offset-info"
    currentGmOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    currentGmOffsetInfo.EntityData.ParentYangName = "utc-offset-info"
    currentGmOffsetInfo.EntityData.SegmentPath = "current-gm-offset-info"
    currentGmOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentGmOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentGmOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentGmOffsetInfo.EntityData.Children = make(map[string]types.YChild)
    currentGmOffsetInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    currentGmOffsetInfo.EntityData.Leafs["offset"] = types.YLeaf{"Offset", currentGmOffsetInfo.Offset}
    currentGmOffsetInfo.EntityData.Leafs["valid"] = types.YLeaf{"Valid", currentGmOffsetInfo.Valid}
    currentGmOffsetInfo.EntityData.Leafs["flag"] = types.YLeaf{"Flag", currentGmOffsetInfo.Flag}
    return &(currentGmOffsetInfo.EntityData)
}

// Ptp_UtcOffsetInfo_ConfiguredOffsetInfo
// The currently configured UTC offset information
type Ptp_UtcOffsetInfo_ConfiguredOffsetInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // Is the UTC offset valid?. The type is bool.
    Valid interface{}

    // Indicates the duration of the final minute of the current day. The type is
    // interface{} with range: 0..255.
    Flag interface{}
}

func (configuredOffsetInfo *Ptp_UtcOffsetInfo_ConfiguredOffsetInfo) GetEntityData() *types.CommonEntityData {
    configuredOffsetInfo.EntityData.YFilter = configuredOffsetInfo.YFilter
    configuredOffsetInfo.EntityData.YangName = "configured-offset-info"
    configuredOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    configuredOffsetInfo.EntityData.ParentYangName = "utc-offset-info"
    configuredOffsetInfo.EntityData.SegmentPath = "configured-offset-info"
    configuredOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOffsetInfo.EntityData.Children = make(map[string]types.YChild)
    configuredOffsetInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    configuredOffsetInfo.EntityData.Leafs["offset"] = types.YLeaf{"Offset", configuredOffsetInfo.Offset}
    configuredOffsetInfo.EntityData.Leafs["valid"] = types.YLeaf{"Valid", configuredOffsetInfo.Valid}
    configuredOffsetInfo.EntityData.Leafs["flag"] = types.YLeaf{"Flag", configuredOffsetInfo.Flag}
    return &(configuredOffsetInfo.EntityData)
}

// Ptp_UtcOffsetInfo_PreviousGmOffsetInfo
// The UTC offset information recovered from the
// prevous grandmaster
type Ptp_UtcOffsetInfo_PreviousGmOffsetInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // Is the UTC offset valid?. The type is bool.
    Valid interface{}

    // Indicates the duration of the final minute of the current day. The type is
    // interface{} with range: 0..255.
    Flag interface{}
}

func (previousGmOffsetInfo *Ptp_UtcOffsetInfo_PreviousGmOffsetInfo) GetEntityData() *types.CommonEntityData {
    previousGmOffsetInfo.EntityData.YFilter = previousGmOffsetInfo.YFilter
    previousGmOffsetInfo.EntityData.YangName = "previous-gm-offset-info"
    previousGmOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    previousGmOffsetInfo.EntityData.ParentYangName = "utc-offset-info"
    previousGmOffsetInfo.EntityData.SegmentPath = "previous-gm-offset-info"
    previousGmOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    previousGmOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    previousGmOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    previousGmOffsetInfo.EntityData.Children = make(map[string]types.YChild)
    previousGmOffsetInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    previousGmOffsetInfo.EntityData.Leafs["offset"] = types.YLeaf{"Offset", previousGmOffsetInfo.Offset}
    previousGmOffsetInfo.EntityData.Leafs["valid"] = types.YLeaf{"Valid", previousGmOffsetInfo.Valid}
    previousGmOffsetInfo.EntityData.Leafs["flag"] = types.YLeaf{"Flag", previousGmOffsetInfo.Flag}
    return &(previousGmOffsetInfo.EntityData)
}

// Ptp_UtcOffsetInfo_HardwareOffsetInfo
// The UTC offset information taken from the
// hardware
type Ptp_UtcOffsetInfo_HardwareOffsetInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // Is the UTC offset valid?. The type is bool.
    Valid interface{}

    // Indicates the duration of the final minute of the current day. The type is
    // interface{} with range: 0..255.
    Flag interface{}
}

func (hardwareOffsetInfo *Ptp_UtcOffsetInfo_HardwareOffsetInfo) GetEntityData() *types.CommonEntityData {
    hardwareOffsetInfo.EntityData.YFilter = hardwareOffsetInfo.YFilter
    hardwareOffsetInfo.EntityData.YangName = "hardware-offset-info"
    hardwareOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    hardwareOffsetInfo.EntityData.ParentYangName = "utc-offset-info"
    hardwareOffsetInfo.EntityData.SegmentPath = "hardware-offset-info"
    hardwareOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareOffsetInfo.EntityData.Children = make(map[string]types.YChild)
    hardwareOffsetInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareOffsetInfo.EntityData.Leafs["offset"] = types.YLeaf{"Offset", hardwareOffsetInfo.Offset}
    hardwareOffsetInfo.EntityData.Leafs["valid"] = types.YLeaf{"Valid", hardwareOffsetInfo.Valid}
    hardwareOffsetInfo.EntityData.Leafs["flag"] = types.YLeaf{"Flag", hardwareOffsetInfo.Flag}
    return &(hardwareOffsetInfo.EntityData)
}

// Ptp_UtcOffsetInfo_GmLeapSecond
// The upcoming leap second advertised by the
// grandmaster (if there is one)
type Ptp_UtcOffsetInfo_GmLeapSecond struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // The UNIX timestamp at which the offset becomes valid. The type is
    // interface{} with range: 0..18446744073709551615.
    OffsetStartDate interface{}

    // The change in UTC offset on applying this offset. The type is interface{}
    // with range: -32768..32767.
    OffsetChange interface{}

    // Indicates whether the offset has been applied. The type is bool.
    OffsetApplied interface{}
}

func (gmLeapSecond *Ptp_UtcOffsetInfo_GmLeapSecond) GetEntityData() *types.CommonEntityData {
    gmLeapSecond.EntityData.YFilter = gmLeapSecond.YFilter
    gmLeapSecond.EntityData.YangName = "gm-leap-second"
    gmLeapSecond.EntityData.BundleName = "cisco_ios_xr"
    gmLeapSecond.EntityData.ParentYangName = "utc-offset-info"
    gmLeapSecond.EntityData.SegmentPath = "gm-leap-second"
    gmLeapSecond.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmLeapSecond.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmLeapSecond.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmLeapSecond.EntityData.Children = make(map[string]types.YChild)
    gmLeapSecond.EntityData.Leafs = make(map[string]types.YLeaf)
    gmLeapSecond.EntityData.Leafs["offset"] = types.YLeaf{"Offset", gmLeapSecond.Offset}
    gmLeapSecond.EntityData.Leafs["offset-start-date"] = types.YLeaf{"OffsetStartDate", gmLeapSecond.OffsetStartDate}
    gmLeapSecond.EntityData.Leafs["offset-change"] = types.YLeaf{"OffsetChange", gmLeapSecond.OffsetChange}
    gmLeapSecond.EntityData.Leafs["offset-applied"] = types.YLeaf{"OffsetApplied", gmLeapSecond.OffsetApplied}
    return &(gmLeapSecond.EntityData)
}

// Ptp_UtcOffsetInfo_ConfiguredLeapSecond
// The list of upcoming configured leap second
// updates
type Ptp_UtcOffsetInfo_ConfiguredLeapSecond struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The UTC offset (TAI - UTC), in seconds. The type is interface{} with range:
    // -32768..32767. Units are second.
    Offset interface{}

    // The UNIX timestamp at which the offset becomes valid. The type is
    // interface{} with range: 0..18446744073709551615.
    OffsetStartDate interface{}

    // The change in UTC offset on applying this offset. The type is interface{}
    // with range: -32768..32767.
    OffsetChange interface{}

    // Indicates whether the offset has been applied. The type is bool.
    OffsetApplied interface{}
}

func (configuredLeapSecond *Ptp_UtcOffsetInfo_ConfiguredLeapSecond) GetEntityData() *types.CommonEntityData {
    configuredLeapSecond.EntityData.YFilter = configuredLeapSecond.YFilter
    configuredLeapSecond.EntityData.YangName = "configured-leap-second"
    configuredLeapSecond.EntityData.BundleName = "cisco_ios_xr"
    configuredLeapSecond.EntityData.ParentYangName = "utc-offset-info"
    configuredLeapSecond.EntityData.SegmentPath = "configured-leap-second"
    configuredLeapSecond.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredLeapSecond.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredLeapSecond.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredLeapSecond.EntityData.Children = make(map[string]types.YChild)
    configuredLeapSecond.EntityData.Leafs = make(map[string]types.YLeaf)
    configuredLeapSecond.EntityData.Leafs["offset"] = types.YLeaf{"Offset", configuredLeapSecond.Offset}
    configuredLeapSecond.EntityData.Leafs["offset-start-date"] = types.YLeaf{"OffsetStartDate", configuredLeapSecond.OffsetStartDate}
    configuredLeapSecond.EntityData.Leafs["offset-change"] = types.YLeaf{"OffsetChange", configuredLeapSecond.OffsetChange}
    configuredLeapSecond.EntityData.Leafs["offset-applied"] = types.YLeaf{"OffsetApplied", configuredLeapSecond.OffsetApplied}
    return &(configuredLeapSecond.EntityData)
}

// Ptp_Platform
// PTP platform specific data
type Ptp_Platform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PTP servo related parameters.
    Servo Ptp_Platform_Servo
}

func (platform *Ptp_Platform) GetEntityData() *types.CommonEntityData {
    platform.EntityData.YFilter = platform.YFilter
    platform.EntityData.YangName = "platform"
    platform.EntityData.BundleName = "cisco_ios_xr"
    platform.EntityData.ParentYangName = "ptp"
    platform.EntityData.SegmentPath = "Cisco-IOS-XR-ptp-pd-oper:platform"
    platform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platform.EntityData.Children = make(map[string]types.YChild)
    platform.EntityData.Children["servo"] = types.YChild{"Servo", &platform.Servo}
    platform.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platform.EntityData)
}

// Ptp_Platform_Servo
// PTP servo related parameters
type Ptp_Platform_Servo struct {
    EntityData types.CommonEntityData
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

func (servo *Ptp_Platform_Servo) GetEntityData() *types.CommonEntityData {
    servo.EntityData.YFilter = servo.YFilter
    servo.EntityData.YangName = "servo"
    servo.EntityData.BundleName = "cisco_ios_xr"
    servo.EntityData.ParentYangName = "platform"
    servo.EntityData.SegmentPath = "servo"
    servo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servo.EntityData.Children = make(map[string]types.YChild)
    servo.EntityData.Children["last-set-time"] = types.YChild{"LastSetTime", &servo.LastSetTime}
    servo.EntityData.Children["last-received-t1"] = types.YChild{"LastReceivedT1", &servo.LastReceivedT1}
    servo.EntityData.Children["last-received-t2"] = types.YChild{"LastReceivedT2", &servo.LastReceivedT2}
    servo.EntityData.Children["last-received-t3"] = types.YChild{"LastReceivedT3", &servo.LastReceivedT3}
    servo.EntityData.Children["last-received-t4"] = types.YChild{"LastReceivedT4", &servo.LastReceivedT4}
    servo.EntityData.Children["pre-received-t1"] = types.YChild{"PreReceivedT1", &servo.PreReceivedT1}
    servo.EntityData.Children["pre-received-t2"] = types.YChild{"PreReceivedT2", &servo.PreReceivedT2}
    servo.EntityData.Children["pre-received-t3"] = types.YChild{"PreReceivedT3", &servo.PreReceivedT3}
    servo.EntityData.Children["pre-received-t4"] = types.YChild{"PreReceivedT4", &servo.PreReceivedT4}
    servo.EntityData.Leafs = make(map[string]types.YLeaf)
    servo.EntityData.Leafs["lock-status"] = types.YLeaf{"LockStatus", servo.LockStatus}
    servo.EntityData.Leafs["running"] = types.YLeaf{"Running", servo.Running}
    servo.EntityData.Leafs["device-status"] = types.YLeaf{"DeviceStatus", servo.DeviceStatus}
    servo.EntityData.Leafs["log-level"] = types.YLeaf{"LogLevel", servo.LogLevel}
    servo.EntityData.Leafs["phase-accuracy-last"] = types.YLeaf{"PhaseAccuracyLast", servo.PhaseAccuracyLast}
    servo.EntityData.Leafs["num-sync-timestamp"] = types.YLeaf{"NumSyncTimestamp", servo.NumSyncTimestamp}
    servo.EntityData.Leafs["num-delay-timestamp"] = types.YLeaf{"NumDelayTimestamp", servo.NumDelayTimestamp}
    servo.EntityData.Leafs["num-set-time"] = types.YLeaf{"NumSetTime", servo.NumSetTime}
    servo.EntityData.Leafs["num-step-time"] = types.YLeaf{"NumStepTime", servo.NumStepTime}
    servo.EntityData.Leafs["num-adjust-freq"] = types.YLeaf{"NumAdjustFreq", servo.NumAdjustFreq}
    servo.EntityData.Leafs["num-adjust-freq-time"] = types.YLeaf{"NumAdjustFreqTime", servo.NumAdjustFreqTime}
    servo.EntityData.Leafs["last-adjust-freq"] = types.YLeaf{"LastAdjustFreq", servo.LastAdjustFreq}
    servo.EntityData.Leafs["last-step-time"] = types.YLeaf{"LastStepTime", servo.LastStepTime}
    servo.EntityData.Leafs["num-discard-sync-timestamp"] = types.YLeaf{"NumDiscardSyncTimestamp", servo.NumDiscardSyncTimestamp}
    servo.EntityData.Leafs["num-discard-delay-timestamp"] = types.YLeaf{"NumDiscardDelayTimestamp", servo.NumDiscardDelayTimestamp}
    servo.EntityData.Leafs["flagof-last-set-time"] = types.YLeaf{"FlagofLastSetTime", servo.FlagofLastSetTime}
    servo.EntityData.Leafs["offset-from-master"] = types.YLeaf{"OffsetFromMaster", servo.OffsetFromMaster}
    servo.EntityData.Leafs["mean-path-delay"] = types.YLeaf{"MeanPathDelay", servo.MeanPathDelay}
    return &(servo.EntityData)
}

// Ptp_Platform_Servo_LastSetTime
// last input of setTime
type Ptp_Platform_Servo_LastSetTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastSetTime *Ptp_Platform_Servo_LastSetTime) GetEntityData() *types.CommonEntityData {
    lastSetTime.EntityData.YFilter = lastSetTime.YFilter
    lastSetTime.EntityData.YangName = "last-set-time"
    lastSetTime.EntityData.BundleName = "cisco_ios_xr"
    lastSetTime.EntityData.ParentYangName = "servo"
    lastSetTime.EntityData.SegmentPath = "last-set-time"
    lastSetTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastSetTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastSetTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastSetTime.EntityData.Children = make(map[string]types.YChild)
    lastSetTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lastSetTime.EntityData.Leafs["second"] = types.YLeaf{"Second", lastSetTime.Second}
    lastSetTime.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastSetTime.NanoSecond}
    return &(lastSetTime.EntityData)
}

// Ptp_Platform_Servo_LastReceivedT1
// last T1 timestamp received
type Ptp_Platform_Servo_LastReceivedT1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT1 *Ptp_Platform_Servo_LastReceivedT1) GetEntityData() *types.CommonEntityData {
    lastReceivedT1.EntityData.YFilter = lastReceivedT1.YFilter
    lastReceivedT1.EntityData.YangName = "last-received-t1"
    lastReceivedT1.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT1.EntityData.ParentYangName = "servo"
    lastReceivedT1.EntityData.SegmentPath = "last-received-t1"
    lastReceivedT1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT1.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT1.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT1.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT1.Second}
    lastReceivedT1.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT1.NanoSecond}
    return &(lastReceivedT1.EntityData)
}

// Ptp_Platform_Servo_LastReceivedT2
// last T2 timestamp received
type Ptp_Platform_Servo_LastReceivedT2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT2 *Ptp_Platform_Servo_LastReceivedT2) GetEntityData() *types.CommonEntityData {
    lastReceivedT2.EntityData.YFilter = lastReceivedT2.YFilter
    lastReceivedT2.EntityData.YangName = "last-received-t2"
    lastReceivedT2.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT2.EntityData.ParentYangName = "servo"
    lastReceivedT2.EntityData.SegmentPath = "last-received-t2"
    lastReceivedT2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT2.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT2.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT2.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT2.Second}
    lastReceivedT2.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT2.NanoSecond}
    return &(lastReceivedT2.EntityData)
}

// Ptp_Platform_Servo_LastReceivedT3
// last T3 timestamp received
type Ptp_Platform_Servo_LastReceivedT3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT3 *Ptp_Platform_Servo_LastReceivedT3) GetEntityData() *types.CommonEntityData {
    lastReceivedT3.EntityData.YFilter = lastReceivedT3.YFilter
    lastReceivedT3.EntityData.YangName = "last-received-t3"
    lastReceivedT3.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT3.EntityData.ParentYangName = "servo"
    lastReceivedT3.EntityData.SegmentPath = "last-received-t3"
    lastReceivedT3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT3.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT3.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT3.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT3.Second}
    lastReceivedT3.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT3.NanoSecond}
    return &(lastReceivedT3.EntityData)
}

// Ptp_Platform_Servo_LastReceivedT4
// last T4 timestamp received
type Ptp_Platform_Servo_LastReceivedT4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT4 *Ptp_Platform_Servo_LastReceivedT4) GetEntityData() *types.CommonEntityData {
    lastReceivedT4.EntityData.YFilter = lastReceivedT4.YFilter
    lastReceivedT4.EntityData.YangName = "last-received-t4"
    lastReceivedT4.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT4.EntityData.ParentYangName = "servo"
    lastReceivedT4.EntityData.SegmentPath = "last-received-t4"
    lastReceivedT4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT4.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT4.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT4.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT4.Second}
    lastReceivedT4.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT4.NanoSecond}
    return &(lastReceivedT4.EntityData)
}

// Ptp_Platform_Servo_PreReceivedT1
// pre T1 timestamp received
type Ptp_Platform_Servo_PreReceivedT1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT1 *Ptp_Platform_Servo_PreReceivedT1) GetEntityData() *types.CommonEntityData {
    preReceivedT1.EntityData.YFilter = preReceivedT1.YFilter
    preReceivedT1.EntityData.YangName = "pre-received-t1"
    preReceivedT1.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT1.EntityData.ParentYangName = "servo"
    preReceivedT1.EntityData.SegmentPath = "pre-received-t1"
    preReceivedT1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT1.EntityData.Children = make(map[string]types.YChild)
    preReceivedT1.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT1.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT1.Second}
    preReceivedT1.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT1.NanoSecond}
    return &(preReceivedT1.EntityData)
}

// Ptp_Platform_Servo_PreReceivedT2
// pre T2 timestamp received
type Ptp_Platform_Servo_PreReceivedT2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT2 *Ptp_Platform_Servo_PreReceivedT2) GetEntityData() *types.CommonEntityData {
    preReceivedT2.EntityData.YFilter = preReceivedT2.YFilter
    preReceivedT2.EntityData.YangName = "pre-received-t2"
    preReceivedT2.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT2.EntityData.ParentYangName = "servo"
    preReceivedT2.EntityData.SegmentPath = "pre-received-t2"
    preReceivedT2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT2.EntityData.Children = make(map[string]types.YChild)
    preReceivedT2.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT2.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT2.Second}
    preReceivedT2.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT2.NanoSecond}
    return &(preReceivedT2.EntityData)
}

// Ptp_Platform_Servo_PreReceivedT3
// pre T3 timestamp received
type Ptp_Platform_Servo_PreReceivedT3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT3 *Ptp_Platform_Servo_PreReceivedT3) GetEntityData() *types.CommonEntityData {
    preReceivedT3.EntityData.YFilter = preReceivedT3.YFilter
    preReceivedT3.EntityData.YangName = "pre-received-t3"
    preReceivedT3.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT3.EntityData.ParentYangName = "servo"
    preReceivedT3.EntityData.SegmentPath = "pre-received-t3"
    preReceivedT3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT3.EntityData.Children = make(map[string]types.YChild)
    preReceivedT3.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT3.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT3.Second}
    preReceivedT3.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT3.NanoSecond}
    return &(preReceivedT3.EntityData)
}

// Ptp_Platform_Servo_PreReceivedT4
// pre T4 timestamp received
type Ptp_Platform_Servo_PreReceivedT4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT4 *Ptp_Platform_Servo_PreReceivedT4) GetEntityData() *types.CommonEntityData {
    preReceivedT4.EntityData.YFilter = preReceivedT4.YFilter
    preReceivedT4.EntityData.YangName = "pre-received-t4"
    preReceivedT4.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT4.EntityData.ParentYangName = "servo"
    preReceivedT4.EntityData.SegmentPath = "pre-received-t4"
    preReceivedT4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT4.EntityData.Children = make(map[string]types.YChild)
    preReceivedT4.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT4.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT4.Second}
    preReceivedT4.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT4.NanoSecond}
    return &(preReceivedT4.EntityData)
}

