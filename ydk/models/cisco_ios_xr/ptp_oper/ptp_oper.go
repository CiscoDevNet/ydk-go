// This module contains a collection of YANG definitions
// for Cisco IOS-XR ptp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ptp: PTP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
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

    // Summary operational data.
    Summary Ptp_Summary

    // Table for interface configuration error operational data.
    InterfaceConfigurationErrors Ptp_InterfaceConfigurationErrors

    // Table for interface foreign master clock operational data.
    InterfaceForeignMasters Ptp_InterfaceForeignMasters

    // Table for interface interop operational data.
    InterfaceInterops Ptp_InterfaceInterops

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
}

func (ptp *Ptp) GetEntityData() *types.CommonEntityData {
    ptp.EntityData.YFilter = ptp.YFilter
    ptp.EntityData.YangName = "ptp"
    ptp.EntityData.BundleName = "cisco_ios_xr"
    ptp.EntityData.ParentYangName = "Cisco-IOS-XR-ptp-oper"
    ptp.EntityData.SegmentPath = "Cisco-IOS-XR-ptp-oper:ptp"
    ptp.EntityData.AbsolutePath = ptp.EntityData.SegmentPath
    ptp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptp.EntityData.Children = types.NewOrderedMap()
    ptp.EntityData.Children.Append("nodes", types.YChild{"Nodes", &ptp.Nodes})
    ptp.EntityData.Children.Append("summary", types.YChild{"Summary", &ptp.Summary})
    ptp.EntityData.Children.Append("interface-configuration-errors", types.YChild{"InterfaceConfigurationErrors", &ptp.InterfaceConfigurationErrors})
    ptp.EntityData.Children.Append("interface-foreign-masters", types.YChild{"InterfaceForeignMasters", &ptp.InterfaceForeignMasters})
    ptp.EntityData.Children.Append("interface-interops", types.YChild{"InterfaceInterops", &ptp.InterfaceInterops})
    ptp.EntityData.Children.Append("local-clock", types.YChild{"LocalClock", &ptp.LocalClock})
    ptp.EntityData.Children.Append("interface-packet-counters", types.YChild{"InterfacePacketCounters", &ptp.InterfacePacketCounters})
    ptp.EntityData.Children.Append("advertised-clock", types.YChild{"AdvertisedClock", &ptp.AdvertisedClock})
    ptp.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ptp.Interfaces})
    ptp.EntityData.Children.Append("dataset", types.YChild{"Dataset", &ptp.Dataset})
    ptp.EntityData.Children.Append("global-configuration-error", types.YChild{"GlobalConfigurationError", &ptp.GlobalConfigurationError})
    ptp.EntityData.Children.Append("grandmaster", types.YChild{"Grandmaster", &ptp.Grandmaster})
    ptp.EntityData.Children.Append("interface-unicast-peers", types.YChild{"InterfaceUnicastPeers", &ptp.InterfaceUnicastPeers})
    ptp.EntityData.Children.Append("utc-offset-info", types.YChild{"UtcOffsetInfo", &ptp.UtcOffsetInfo})
    ptp.EntityData.Leafs = types.NewOrderedMap()

    ptp.EntityData.YListKeys = []string {}

    return &(ptp.EntityData)
}

// Ptp_Nodes
// Table for node-specific operational data
type Ptp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific operational data for a given node. The type is slice of
    // Ptp_Nodes_Node.
    Node []*Ptp_Nodes_Node
}

func (nodes *Ptp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ptp"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + nodes.EntityData.SegmentPath
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

// Ptp_Nodes_Node
// Node-specific operational data for a given node
type Ptp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("node-interface-foreign-masters", types.YChild{"NodeInterfaceForeignMasters", &node.NodeInterfaceForeignMasters})
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("node-interfaces", types.YChild{"NodeInterfaces", &node.NodeInterfaces})
    node.EntityData.Children.Append("node-interface-unicast-peers", types.YChild{"NodeInterfaceUnicastPeers", &node.NodeInterfaceUnicastPeers})
    node.EntityData.Children.Append("packet-counters", types.YChild{"PacketCounters", &node.PacketCounters})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

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
    NodeInterfaceForeignMaster []*Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster
}

func (nodeInterfaceForeignMasters *Ptp_Nodes_Node_NodeInterfaceForeignMasters) GetEntityData() *types.CommonEntityData {
    nodeInterfaceForeignMasters.EntityData.YFilter = nodeInterfaceForeignMasters.YFilter
    nodeInterfaceForeignMasters.EntityData.YangName = "node-interface-foreign-masters"
    nodeInterfaceForeignMasters.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceForeignMasters.EntityData.ParentYangName = "node"
    nodeInterfaceForeignMasters.EntityData.SegmentPath = "node-interface-foreign-masters"
    nodeInterfaceForeignMasters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/" + nodeInterfaceForeignMasters.EntityData.SegmentPath
    nodeInterfaceForeignMasters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceForeignMasters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceForeignMasters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceForeignMasters.EntityData.Children = types.NewOrderedMap()
    nodeInterfaceForeignMasters.EntityData.Children.Append("node-interface-foreign-master", types.YChild{"NodeInterfaceForeignMaster", nil})
    for i := range nodeInterfaceForeignMasters.NodeInterfaceForeignMaster {
        nodeInterfaceForeignMasters.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[i]), types.YChild{"NodeInterfaceForeignMaster", nodeInterfaceForeignMasters.NodeInterfaceForeignMaster[i]})
    }
    nodeInterfaceForeignMasters.EntityData.Leafs = types.NewOrderedMap()

    nodeInterfaceForeignMasters.EntityData.YListKeys = []string {}

    return &(nodeInterfaceForeignMasters.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster
// Node interface foreign master clock
// operational data
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Foreign clocks received on this interface. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock.
    ForeignClock []*Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock
}

func (nodeInterfaceForeignMaster *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster) GetEntityData() *types.CommonEntityData {
    nodeInterfaceForeignMaster.EntityData.YFilter = nodeInterfaceForeignMaster.YFilter
    nodeInterfaceForeignMaster.EntityData.YangName = "node-interface-foreign-master"
    nodeInterfaceForeignMaster.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceForeignMaster.EntityData.ParentYangName = "node-interface-foreign-masters"
    nodeInterfaceForeignMaster.EntityData.SegmentPath = "node-interface-foreign-master" + types.AddKeyToken(nodeInterfaceForeignMaster.InterfaceName, "interface-name")
    nodeInterfaceForeignMaster.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/" + nodeInterfaceForeignMaster.EntityData.SegmentPath
    nodeInterfaceForeignMaster.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceForeignMaster.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceForeignMaster.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceForeignMaster.EntityData.Children = types.NewOrderedMap()
    nodeInterfaceForeignMaster.EntityData.Children.Append("foreign-clock", types.YChild{"ForeignClock", nil})
    for i := range nodeInterfaceForeignMaster.ForeignClock {
        types.SetYListKey(nodeInterfaceForeignMaster.ForeignClock[i], i)
        nodeInterfaceForeignMaster.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaceForeignMaster.ForeignClock[i]), types.YChild{"ForeignClock", nodeInterfaceForeignMaster.ForeignClock[i]})
    }
    nodeInterfaceForeignMaster.EntityData.Leafs = types.NewOrderedMap()
    nodeInterfaceForeignMaster.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nodeInterfaceForeignMaster.InterfaceName})
    nodeInterfaceForeignMaster.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", nodeInterfaceForeignMaster.PortNumber})

    nodeInterfaceForeignMaster.EntityData.YListKeys = []string {"InterfaceName"}

    return &(nodeInterfaceForeignMaster.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock
// Foreign clocks received on this interface
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // The clock has clock class corresponding to QL-DNU. The type is bool.
    IsDnu interface{}

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

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock) GetEntityData() *types.CommonEntityData {
    foreignClock.EntityData.YFilter = foreignClock.YFilter
    foreignClock.EntityData.YangName = "foreign-clock"
    foreignClock.EntityData.BundleName = "cisco_ios_xr"
    foreignClock.EntityData.ParentYangName = "node-interface-foreign-master"
    foreignClock.EntityData.SegmentPath = "foreign-clock" + types.AddNoKeyToken(foreignClock)
    foreignClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/" + foreignClock.EntityData.SegmentPath
    foreignClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock.EntityData.Children = types.NewOrderedMap()
    foreignClock.EntityData.Children.Append("foreign-clock", types.YChild{"ForeignClock", &foreignClock.ForeignClock})
    foreignClock.EntityData.Children.Append("address", types.YChild{"Address", &foreignClock.Address})
    foreignClock.EntityData.Children.Append("announce-grant", types.YChild{"AnnounceGrant", &foreignClock.AnnounceGrant})
    foreignClock.EntityData.Children.Append("sync-grant", types.YChild{"SyncGrant", &foreignClock.SyncGrant})
    foreignClock.EntityData.Children.Append("delay-response-grant", types.YChild{"DelayResponseGrant", &foreignClock.DelayResponseGrant})
    foreignClock.EntityData.Leafs = types.NewOrderedMap()
    foreignClock.EntityData.Leafs.Append("is-qualified", types.YLeaf{"IsQualified", foreignClock.IsQualified})
    foreignClock.EntityData.Leafs.Append("is-grandmaster", types.YLeaf{"IsGrandmaster", foreignClock.IsGrandmaster})
    foreignClock.EntityData.Leafs.Append("communication-model", types.YLeaf{"CommunicationModel", foreignClock.CommunicationModel})
    foreignClock.EntityData.Leafs.Append("is-known", types.YLeaf{"IsKnown", foreignClock.IsKnown})
    foreignClock.EntityData.Leafs.Append("time-known-for", types.YLeaf{"TimeKnownFor", foreignClock.TimeKnownFor})
    foreignClock.EntityData.Leafs.Append("foreign-domain", types.YLeaf{"ForeignDomain", foreignClock.ForeignDomain})
    foreignClock.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", foreignClock.ConfiguredPriority})
    foreignClock.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", foreignClock.ConfiguredClockClass})
    foreignClock.EntityData.Leafs.Append("delay-asymmetry", types.YLeaf{"DelayAsymmetry", foreignClock.DelayAsymmetry})
    foreignClock.EntityData.Leafs.Append("ptsf-loss-announce", types.YLeaf{"PtsfLossAnnounce", foreignClock.PtsfLossAnnounce})
    foreignClock.EntityData.Leafs.Append("ptsf-loss-sync", types.YLeaf{"PtsfLossSync", foreignClock.PtsfLossSync})
    foreignClock.EntityData.Leafs.Append("is-dnu", types.YLeaf{"IsDnu", foreignClock.IsDnu})

    foreignClock.EntityData.YListKeys = []string {}

    return &(foreignClock.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock
// Foreign clock information
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock struct {
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
    UtcOffset Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset

    // Receiver.
    Receiver Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver

    // Sender.
    Sender Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender
}

func (foreignClock *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock) GetEntityData() *types.CommonEntityData {
    foreignClock.EntityData.YFilter = foreignClock.YFilter
    foreignClock.EntityData.YangName = "foreign-clock"
    foreignClock.EntityData.BundleName = "cisco_ios_xr"
    foreignClock.EntityData.ParentYangName = "foreign-clock"
    foreignClock.EntityData.SegmentPath = "foreign-clock"
    foreignClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/" + foreignClock.EntityData.SegmentPath
    foreignClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock.EntityData.Children = types.NewOrderedMap()
    foreignClock.EntityData.Children.Append("utc-offset", types.YChild{"UtcOffset", &foreignClock.UtcOffset})
    foreignClock.EntityData.Children.Append("receiver", types.YChild{"Receiver", &foreignClock.Receiver})
    foreignClock.EntityData.Children.Append("sender", types.YChild{"Sender", &foreignClock.Sender})
    foreignClock.EntityData.Leafs = types.NewOrderedMap()
    foreignClock.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", foreignClock.ClockId})
    foreignClock.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", foreignClock.Priority1})
    foreignClock.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", foreignClock.Priority2})
    foreignClock.EntityData.Leafs.Append("class", types.YLeaf{"Class", foreignClock.Class})
    foreignClock.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", foreignClock.Accuracy})
    foreignClock.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", foreignClock.OffsetLogVariance})
    foreignClock.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", foreignClock.StepsRemoved})
    foreignClock.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", foreignClock.TimeSource})
    foreignClock.EntityData.Leafs.Append("frequency-traceable", types.YLeaf{"FrequencyTraceable", foreignClock.FrequencyTraceable})
    foreignClock.EntityData.Leafs.Append("time-traceable", types.YLeaf{"TimeTraceable", foreignClock.TimeTraceable})
    foreignClock.EntityData.Leafs.Append("timescale", types.YLeaf{"Timescale", foreignClock.Timescale})
    foreignClock.EntityData.Leafs.Append("leap-seconds", types.YLeaf{"LeapSeconds", foreignClock.LeapSeconds})
    foreignClock.EntityData.Leafs.Append("local", types.YLeaf{"Local", foreignClock.Local})
    foreignClock.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", foreignClock.ConfiguredClockClass})
    foreignClock.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", foreignClock.ConfiguredPriority})

    foreignClock.EntityData.YListKeys = []string {}

    return &(foreignClock.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset
// UTC offset
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "foreign-clock"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/foreign-clock/" + utcOffset.EntityData.SegmentPath
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = types.NewOrderedMap()
    utcOffset.EntityData.Leafs = types.NewOrderedMap()
    utcOffset.EntityData.Leafs.Append("current-offset", types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset})
    utcOffset.EntityData.Leafs.Append("offset-valid", types.YLeaf{"OffsetValid", utcOffset.OffsetValid})

    utcOffset.EntityData.YListKeys = []string {}

    return &(utcOffset.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver
// Receiver
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "foreign-clock"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/foreign-clock/" + receiver.EntityData.SegmentPath
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", receiver.ClockId})
    receiver.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", receiver.PortNumber})

    receiver.EntityData.YListKeys = []string {}

    return &(receiver.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender
// Sender
type Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_Nodes_Node_NodeInterfaceForeignMasters_NodeInterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "foreign-clock"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/foreign-clock/" + sender.EntityData.SegmentPath
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = types.NewOrderedMap()
    sender.EntityData.Leafs = types.NewOrderedMap()
    sender.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", sender.ClockId})
    sender.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", sender.PortNumber})

    sender.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    announceGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/" + announceGrant.EntityData.SegmentPath
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = types.NewOrderedMap()
    announceGrant.EntityData.Leafs = types.NewOrderedMap()
    announceGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval})
    announceGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", announceGrant.GrantDuration})

    announceGrant.EntityData.YListKeys = []string {}

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
    syncGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/" + syncGrant.EntityData.SegmentPath
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = types.NewOrderedMap()
    syncGrant.EntityData.Leafs = types.NewOrderedMap()
    syncGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval})
    syncGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", syncGrant.GrantDuration})

    syncGrant.EntityData.YListKeys = []string {}

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
    delayResponseGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-foreign-masters/node-interface-foreign-master/foreign-clock/" + delayResponseGrant.EntityData.SegmentPath
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval})
    delayResponseGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration})

    delayResponseGrant.EntityData.YListKeys = []string {}

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
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("port-state-init-count", types.YLeaf{"PortStateInitCount", summary.PortStateInitCount})
    summary.EntityData.Leafs.Append("port-state-listening-count", types.YLeaf{"PortStateListeningCount", summary.PortStateListeningCount})
    summary.EntityData.Leafs.Append("port-state-passive-count", types.YLeaf{"PortStatePassiveCount", summary.PortStatePassiveCount})
    summary.EntityData.Leafs.Append("port-state-pre-master-count", types.YLeaf{"PortStatePreMasterCount", summary.PortStatePreMasterCount})
    summary.EntityData.Leafs.Append("port-state-master-count", types.YLeaf{"PortStateMasterCount", summary.PortStateMasterCount})
    summary.EntityData.Leafs.Append("port-state-slave-count", types.YLeaf{"PortStateSlaveCount", summary.PortStateSlaveCount})
    summary.EntityData.Leafs.Append("port-state-uncalibrated-count", types.YLeaf{"PortStateUncalibratedCount", summary.PortStateUncalibratedCount})
    summary.EntityData.Leafs.Append("port-state-faulty-count", types.YLeaf{"PortStateFaultyCount", summary.PortStateFaultyCount})
    summary.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", summary.TotalInterfaces})
    summary.EntityData.Leafs.Append("total-interfaces-valid-port-num", types.YLeaf{"TotalInterfacesValidPortNum", summary.TotalInterfacesValidPortNum})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces
// Table for node interface operational data
type Ptp_Nodes_Node_NodeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node interface operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface.
    NodeInterface []*Ptp_Nodes_Node_NodeInterfaces_NodeInterface
}

func (nodeInterfaces *Ptp_Nodes_Node_NodeInterfaces) GetEntityData() *types.CommonEntityData {
    nodeInterfaces.EntityData.YFilter = nodeInterfaces.YFilter
    nodeInterfaces.EntityData.YangName = "node-interfaces"
    nodeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaces.EntityData.ParentYangName = "node"
    nodeInterfaces.EntityData.SegmentPath = "node-interfaces"
    nodeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/" + nodeInterfaces.EntityData.SegmentPath
    nodeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaces.EntityData.Children = types.NewOrderedMap()
    nodeInterfaces.EntityData.Children.Append("node-interface", types.YChild{"NodeInterface", nil})
    for i := range nodeInterfaces.NodeInterface {
        nodeInterfaces.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaces.NodeInterface[i]), types.YChild{"NodeInterface", nodeInterfaces.NodeInterface[i]})
    }
    nodeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    nodeInterfaces.EntityData.YListKeys = []string {}

    return &(nodeInterfaces.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface
// Node interface operational data
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Port state. The type is PtpBagPortState.
    PortState interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Line state. The type is ImStateEnum.
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

    // The interface supports unicast. The type is bool.
    SupportsUnicast interface{}

    // The interface supports operation in master mode. The type is bool.
    SupportsMaster interface{}

    // The interface supports one-step operation. The type is bool.
    SupportsOneStep interface{}

    // The interface supports two-step operation. The type is bool.
    SupportsTwoStep interface{}

    // The interface supports ethernet transport. The type is bool.
    SupportsEthernet interface{}

    // The interface supports multicast. The type is bool.
    SupportsMulticast interface{}

    // The interface supports IPv4 transport. The type is bool.
    SupportsIpv4 interface{}

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

    // Indicate whether profile interop is in use. The type is bool.
    ProfileInterop interface{}

    // The PTP domain that is being interoperated with. The type is interface{}
    // with range: 0..255.
    InteropDomain interface{}

    // Profile that is being interoperated with. The type is PtpBagProfile.
    InteropProfile interface{}

    // List of Ipv6 addresses, if IPv6 encapsulation is being used. If a source
    // address is configured, this is the only item in the list.
    Ipv6AddressArray Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv6AddressArray

    // List of IPv4 addresses, if IPv4 encapsulation is being used. The first
    // address is the primary address. If a source address is configured, this is
    // the only item in the list.
    Ipv4AddressArray Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv4AddressArray

    // MAC address, if Ethernet encapsulation is being used.
    MacAddress Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MacAddress

    // Details of any ingress conversion.
    IngressConversion Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion

    // Details of any egress conversion.
    EgressConversion Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion

    // The interface's master table. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable.
    MasterTable []*Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable
}

func (nodeInterface *Ptp_Nodes_Node_NodeInterfaces_NodeInterface) GetEntityData() *types.CommonEntityData {
    nodeInterface.EntityData.YFilter = nodeInterface.YFilter
    nodeInterface.EntityData.YangName = "node-interface"
    nodeInterface.EntityData.BundleName = "cisco_ios_xr"
    nodeInterface.EntityData.ParentYangName = "node-interfaces"
    nodeInterface.EntityData.SegmentPath = "node-interface" + types.AddKeyToken(nodeInterface.InterfaceName, "interface-name")
    nodeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/" + nodeInterface.EntityData.SegmentPath
    nodeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterface.EntityData.Children = types.NewOrderedMap()
    nodeInterface.EntityData.Children.Append("ipv6-address-array", types.YChild{"Ipv6AddressArray", &nodeInterface.Ipv6AddressArray})
    nodeInterface.EntityData.Children.Append("ipv4-address-array", types.YChild{"Ipv4AddressArray", &nodeInterface.Ipv4AddressArray})
    nodeInterface.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &nodeInterface.MacAddress})
    nodeInterface.EntityData.Children.Append("ingress-conversion", types.YChild{"IngressConversion", &nodeInterface.IngressConversion})
    nodeInterface.EntityData.Children.Append("egress-conversion", types.YChild{"EgressConversion", &nodeInterface.EgressConversion})
    nodeInterface.EntityData.Children.Append("master-table", types.YChild{"MasterTable", nil})
    for i := range nodeInterface.MasterTable {
        types.SetYListKey(nodeInterface.MasterTable[i], i)
        nodeInterface.EntityData.Children.Append(types.GetSegmentPath(nodeInterface.MasterTable[i]), types.YChild{"MasterTable", nodeInterface.MasterTable[i]})
    }
    nodeInterface.EntityData.Leafs = types.NewOrderedMap()
    nodeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nodeInterface.InterfaceName})
    nodeInterface.EntityData.Leafs.Append("port-state", types.YLeaf{"PortState", nodeInterface.PortState})
    nodeInterface.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", nodeInterface.PortNumber})
    nodeInterface.EntityData.Leafs.Append("line-state", types.YLeaf{"LineState", nodeInterface.LineState})
    nodeInterface.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", nodeInterface.Encapsulation})
    nodeInterface.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nodeInterface.Ipv6Address})
    nodeInterface.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nodeInterface.Ipv4Address})
    nodeInterface.EntityData.Leafs.Append("two-step", types.YLeaf{"TwoStep", nodeInterface.TwoStep})
    nodeInterface.EntityData.Leafs.Append("communication-model", types.YLeaf{"CommunicationModel", nodeInterface.CommunicationModel})
    nodeInterface.EntityData.Leafs.Append("log-sync-interval", types.YLeaf{"LogSyncInterval", nodeInterface.LogSyncInterval})
    nodeInterface.EntityData.Leafs.Append("log-announce-interval", types.YLeaf{"LogAnnounceInterval", nodeInterface.LogAnnounceInterval})
    nodeInterface.EntityData.Leafs.Append("announce-timeout", types.YLeaf{"AnnounceTimeout", nodeInterface.AnnounceTimeout})
    nodeInterface.EntityData.Leafs.Append("log-min-delay-request-interval", types.YLeaf{"LogMinDelayRequestInterval", nodeInterface.LogMinDelayRequestInterval})
    nodeInterface.EntityData.Leafs.Append("configured-port-state", types.YLeaf{"ConfiguredPortState", nodeInterface.ConfiguredPortState})
    nodeInterface.EntityData.Leafs.Append("supports-unicast", types.YLeaf{"SupportsUnicast", nodeInterface.SupportsUnicast})
    nodeInterface.EntityData.Leafs.Append("supports-master", types.YLeaf{"SupportsMaster", nodeInterface.SupportsMaster})
    nodeInterface.EntityData.Leafs.Append("supports-one-step", types.YLeaf{"SupportsOneStep", nodeInterface.SupportsOneStep})
    nodeInterface.EntityData.Leafs.Append("supports-two-step", types.YLeaf{"SupportsTwoStep", nodeInterface.SupportsTwoStep})
    nodeInterface.EntityData.Leafs.Append("supports-ethernet", types.YLeaf{"SupportsEthernet", nodeInterface.SupportsEthernet})
    nodeInterface.EntityData.Leafs.Append("supports-multicast", types.YLeaf{"SupportsMulticast", nodeInterface.SupportsMulticast})
    nodeInterface.EntityData.Leafs.Append("supports-ipv4", types.YLeaf{"SupportsIpv4", nodeInterface.SupportsIpv4})
    nodeInterface.EntityData.Leafs.Append("supports-ipv6", types.YLeaf{"SupportsIpv6", nodeInterface.SupportsIpv6})
    nodeInterface.EntityData.Leafs.Append("supports-slave", types.YLeaf{"SupportsSlave", nodeInterface.SupportsSlave})
    nodeInterface.EntityData.Leafs.Append("supports-source-ip", types.YLeaf{"SupportsSourceIp", nodeInterface.SupportsSourceIp})
    nodeInterface.EntityData.Leafs.Append("max-sync-rate", types.YLeaf{"MaxSyncRate", nodeInterface.MaxSyncRate})
    nodeInterface.EntityData.Leafs.Append("event-cos", types.YLeaf{"EventCos", nodeInterface.EventCos})
    nodeInterface.EntityData.Leafs.Append("general-cos", types.YLeaf{"GeneralCos", nodeInterface.GeneralCos})
    nodeInterface.EntityData.Leafs.Append("event-dscp", types.YLeaf{"EventDscp", nodeInterface.EventDscp})
    nodeInterface.EntityData.Leafs.Append("general-dscp", types.YLeaf{"GeneralDscp", nodeInterface.GeneralDscp})
    nodeInterface.EntityData.Leafs.Append("unicast-peers", types.YLeaf{"UnicastPeers", nodeInterface.UnicastPeers})
    nodeInterface.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", nodeInterface.LocalPriority})
    nodeInterface.EntityData.Leafs.Append("signal-fail", types.YLeaf{"SignalFail", nodeInterface.SignalFail})
    nodeInterface.EntityData.Leafs.Append("profile-interop", types.YLeaf{"ProfileInterop", nodeInterface.ProfileInterop})
    nodeInterface.EntityData.Leafs.Append("interop-domain", types.YLeaf{"InteropDomain", nodeInterface.InteropDomain})
    nodeInterface.EntityData.Leafs.Append("interop-profile", types.YLeaf{"InteropProfile", nodeInterface.InteropProfile})

    nodeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(nodeInterface.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv6AddressArray
// List of Ipv6 addresses, if IPv6 encapsulation is
// being used. If a source address is configured,
// this is the only item in the list
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv6AddressArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv6 addresses. The type is slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Addr []interface{}
}

func (ipv6AddressArray *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv6AddressArray) GetEntityData() *types.CommonEntityData {
    ipv6AddressArray.EntityData.YFilter = ipv6AddressArray.YFilter
    ipv6AddressArray.EntityData.YangName = "ipv6-address-array"
    ipv6AddressArray.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressArray.EntityData.ParentYangName = "node-interface"
    ipv6AddressArray.EntityData.SegmentPath = "ipv6-address-array"
    ipv6AddressArray.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/" + ipv6AddressArray.EntityData.SegmentPath
    ipv6AddressArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressArray.EntityData.Children = types.NewOrderedMap()
    ipv6AddressArray.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressArray.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", ipv6AddressArray.Addr})

    ipv6AddressArray.EntityData.YListKeys = []string {}

    return &(ipv6AddressArray.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv4AddressArray
// List of IPv4 addresses, if IPv4 encapsulation is
// being used. The first address is the primary
// address. If a source address is configured, this
// is the only item in the list.
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv4AddressArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv4 addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Addr []interface{}
}

func (ipv4AddressArray *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_Ipv4AddressArray) GetEntityData() *types.CommonEntityData {
    ipv4AddressArray.EntityData.YFilter = ipv4AddressArray.YFilter
    ipv4AddressArray.EntityData.YangName = "ipv4-address-array"
    ipv4AddressArray.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressArray.EntityData.ParentYangName = "node-interface"
    ipv4AddressArray.EntityData.SegmentPath = "ipv4-address-array"
    ipv4AddressArray.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/" + ipv4AddressArray.EntityData.SegmentPath
    ipv4AddressArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressArray.EntityData.Children = types.NewOrderedMap()
    ipv4AddressArray.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressArray.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", ipv4AddressArray.Addr})

    ipv4AddressArray.EntityData.YListKeys = []string {}

    return &(ipv4AddressArray.EntityData)
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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

    return &(macAddress.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion
// Details of any ingress conversion
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Class Default. The type is interface{} with range: 0..255.
    ClassDefault interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Class Mapping. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion_ClassMapping.
    ClassMapping []*Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion_ClassMapping
}

func (ingressConversion *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion) GetEntityData() *types.CommonEntityData {
    ingressConversion.EntityData.YFilter = ingressConversion.YFilter
    ingressConversion.EntityData.YangName = "ingress-conversion"
    ingressConversion.EntityData.BundleName = "cisco_ios_xr"
    ingressConversion.EntityData.ParentYangName = "node-interface"
    ingressConversion.EntityData.SegmentPath = "ingress-conversion"
    ingressConversion.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/" + ingressConversion.EntityData.SegmentPath
    ingressConversion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressConversion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressConversion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressConversion.EntityData.Children = types.NewOrderedMap()
    ingressConversion.EntityData.Children.Append("class-mapping", types.YChild{"ClassMapping", nil})
    for i := range ingressConversion.ClassMapping {
        types.SetYListKey(ingressConversion.ClassMapping[i], i)
        ingressConversion.EntityData.Children.Append(types.GetSegmentPath(ingressConversion.ClassMapping[i]), types.YChild{"ClassMapping", ingressConversion.ClassMapping[i]})
    }
    ingressConversion.EntityData.Leafs = types.NewOrderedMap()
    ingressConversion.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", ingressConversion.Priority1})
    ingressConversion.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", ingressConversion.Priority2})
    ingressConversion.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", ingressConversion.Accuracy})
    ingressConversion.EntityData.Leafs.Append("class-default", types.YLeaf{"ClassDefault", ingressConversion.ClassDefault})
    ingressConversion.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", ingressConversion.OffsetLogVariance})

    ingressConversion.EntityData.YListKeys = []string {}

    return &(ingressConversion.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion_ClassMapping
// Class Mapping
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion_ClassMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // From clock class. The type is interface{} with range: 0..255.
    FromClockClass interface{}

    // To clock class. The type is interface{} with range: 0..255.
    ToClockClass interface{}
}

func (classMapping *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_IngressConversion_ClassMapping) GetEntityData() *types.CommonEntityData {
    classMapping.EntityData.YFilter = classMapping.YFilter
    classMapping.EntityData.YangName = "class-mapping"
    classMapping.EntityData.BundleName = "cisco_ios_xr"
    classMapping.EntityData.ParentYangName = "ingress-conversion"
    classMapping.EntityData.SegmentPath = "class-mapping" + types.AddNoKeyToken(classMapping)
    classMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/ingress-conversion/" + classMapping.EntityData.SegmentPath
    classMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapping.EntityData.Children = types.NewOrderedMap()
    classMapping.EntityData.Leafs = types.NewOrderedMap()
    classMapping.EntityData.Leafs.Append("from-clock-class", types.YLeaf{"FromClockClass", classMapping.FromClockClass})
    classMapping.EntityData.Leafs.Append("to-clock-class", types.YLeaf{"ToClockClass", classMapping.ToClockClass})

    classMapping.EntityData.YListKeys = []string {}

    return &(classMapping.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion
// Details of any egress conversion
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Class Default. The type is interface{} with range: 0..255.
    ClassDefault interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Class Mapping. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion_ClassMapping.
    ClassMapping []*Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion_ClassMapping
}

func (egressConversion *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion) GetEntityData() *types.CommonEntityData {
    egressConversion.EntityData.YFilter = egressConversion.YFilter
    egressConversion.EntityData.YangName = "egress-conversion"
    egressConversion.EntityData.BundleName = "cisco_ios_xr"
    egressConversion.EntityData.ParentYangName = "node-interface"
    egressConversion.EntityData.SegmentPath = "egress-conversion"
    egressConversion.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/" + egressConversion.EntityData.SegmentPath
    egressConversion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressConversion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressConversion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressConversion.EntityData.Children = types.NewOrderedMap()
    egressConversion.EntityData.Children.Append("class-mapping", types.YChild{"ClassMapping", nil})
    for i := range egressConversion.ClassMapping {
        types.SetYListKey(egressConversion.ClassMapping[i], i)
        egressConversion.EntityData.Children.Append(types.GetSegmentPath(egressConversion.ClassMapping[i]), types.YChild{"ClassMapping", egressConversion.ClassMapping[i]})
    }
    egressConversion.EntityData.Leafs = types.NewOrderedMap()
    egressConversion.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", egressConversion.Priority1})
    egressConversion.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", egressConversion.Priority2})
    egressConversion.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", egressConversion.Accuracy})
    egressConversion.EntityData.Leafs.Append("class-default", types.YLeaf{"ClassDefault", egressConversion.ClassDefault})
    egressConversion.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", egressConversion.OffsetLogVariance})

    egressConversion.EntityData.YListKeys = []string {}

    return &(egressConversion.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion_ClassMapping
// Class Mapping
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion_ClassMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // From clock class. The type is interface{} with range: 0..255.
    FromClockClass interface{}

    // To clock class. The type is interface{} with range: 0..255.
    ToClockClass interface{}
}

func (classMapping *Ptp_Nodes_Node_NodeInterfaces_NodeInterface_EgressConversion_ClassMapping) GetEntityData() *types.CommonEntityData {
    classMapping.EntityData.YFilter = classMapping.YFilter
    classMapping.EntityData.YangName = "class-mapping"
    classMapping.EntityData.BundleName = "cisco_ios_xr"
    classMapping.EntityData.ParentYangName = "egress-conversion"
    classMapping.EntityData.SegmentPath = "class-mapping" + types.AddNoKeyToken(classMapping)
    classMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/egress-conversion/" + classMapping.EntityData.SegmentPath
    classMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapping.EntityData.Children = types.NewOrderedMap()
    classMapping.EntityData.Leafs = types.NewOrderedMap()
    classMapping.EntityData.Leafs.Append("from-clock-class", types.YLeaf{"FromClockClass", classMapping.FromClockClass})
    classMapping.EntityData.Leafs.Append("to-clock-class", types.YLeaf{"ToClockClass", classMapping.ToClockClass})

    classMapping.EntityData.YListKeys = []string {}

    return &(classMapping.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable
// The interface's master table
type Ptp_Nodes_Node_NodeInterfaces_NodeInterface_MasterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    masterTable.EntityData.SegmentPath = "master-table" + types.AddNoKeyToken(masterTable)
    masterTable.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/" + masterTable.EntityData.SegmentPath
    masterTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterTable.EntityData.Children = types.NewOrderedMap()
    masterTable.EntityData.Children.Append("address", types.YChild{"Address", &masterTable.Address})
    masterTable.EntityData.Leafs = types.NewOrderedMap()
    masterTable.EntityData.Leafs.Append("communication-model", types.YLeaf{"CommunicationModel", masterTable.CommunicationModel})
    masterTable.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", masterTable.Priority})
    masterTable.EntityData.Leafs.Append("known", types.YLeaf{"Known", masterTable.Known})
    masterTable.EntityData.Leafs.Append("qualified", types.YLeaf{"Qualified", masterTable.Qualified})
    masterTable.EntityData.Leafs.Append("is-grandmaster", types.YLeaf{"IsGrandmaster", masterTable.IsGrandmaster})
    masterTable.EntityData.Leafs.Append("ptsf-loss-announce", types.YLeaf{"PtsfLossAnnounce", masterTable.PtsfLossAnnounce})
    masterTable.EntityData.Leafs.Append("ptsf-loss-sync", types.YLeaf{"PtsfLossSync", masterTable.PtsfLossSync})
    masterTable.EntityData.Leafs.Append("is-nonnegotiated", types.YLeaf{"IsNonnegotiated", masterTable.IsNonnegotiated})

    masterTable.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/master-table/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/master-table/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interfaces/node-interface/master-table/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

    return &(ipv6Address.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers
// Table for node unicast peers operational data
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node interface unicast peers operational data. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer.
    NodeInterfaceUnicastPeer []*Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer
}

func (nodeInterfaceUnicastPeers *Ptp_Nodes_Node_NodeInterfaceUnicastPeers) GetEntityData() *types.CommonEntityData {
    nodeInterfaceUnicastPeers.EntityData.YFilter = nodeInterfaceUnicastPeers.YFilter
    nodeInterfaceUnicastPeers.EntityData.YangName = "node-interface-unicast-peers"
    nodeInterfaceUnicastPeers.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceUnicastPeers.EntityData.ParentYangName = "node"
    nodeInterfaceUnicastPeers.EntityData.SegmentPath = "node-interface-unicast-peers"
    nodeInterfaceUnicastPeers.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/" + nodeInterfaceUnicastPeers.EntityData.SegmentPath
    nodeInterfaceUnicastPeers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceUnicastPeers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceUnicastPeers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceUnicastPeers.EntityData.Children = types.NewOrderedMap()
    nodeInterfaceUnicastPeers.EntityData.Children.Append("node-interface-unicast-peer", types.YChild{"NodeInterfaceUnicastPeer", nil})
    for i := range nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer {
        nodeInterfaceUnicastPeers.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[i]), types.YChild{"NodeInterfaceUnicastPeer", nodeInterfaceUnicastPeers.NodeInterfaceUnicastPeer[i]})
    }
    nodeInterfaceUnicastPeers.EntityData.Leafs = types.NewOrderedMap()

    nodeInterfaceUnicastPeers.EntityData.YListKeys = []string {}

    return &(nodeInterfaceUnicastPeers.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer
// Node interface unicast peers operational data
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Unicast Peers. The type is slice of
    // Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers.
    Peers []*Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers
}

func (nodeInterfaceUnicastPeer *Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer) GetEntityData() *types.CommonEntityData {
    nodeInterfaceUnicastPeer.EntityData.YFilter = nodeInterfaceUnicastPeer.YFilter
    nodeInterfaceUnicastPeer.EntityData.YangName = "node-interface-unicast-peer"
    nodeInterfaceUnicastPeer.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceUnicastPeer.EntityData.ParentYangName = "node-interface-unicast-peers"
    nodeInterfaceUnicastPeer.EntityData.SegmentPath = "node-interface-unicast-peer" + types.AddKeyToken(nodeInterfaceUnicastPeer.InterfaceName, "interface-name")
    nodeInterfaceUnicastPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/" + nodeInterfaceUnicastPeer.EntityData.SegmentPath
    nodeInterfaceUnicastPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceUnicastPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceUnicastPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceUnicastPeer.EntityData.Children = types.NewOrderedMap()
    nodeInterfaceUnicastPeer.EntityData.Children.Append("peers", types.YChild{"Peers", nil})
    for i := range nodeInterfaceUnicastPeer.Peers {
        types.SetYListKey(nodeInterfaceUnicastPeer.Peers[i], i)
        nodeInterfaceUnicastPeer.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaceUnicastPeer.Peers[i]), types.YChild{"Peers", nodeInterfaceUnicastPeer.Peers[i]})
    }
    nodeInterfaceUnicastPeer.EntityData.Leafs = types.NewOrderedMap()
    nodeInterfaceUnicastPeer.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nodeInterfaceUnicastPeer.InterfaceName})
    nodeInterfaceUnicastPeer.EntityData.Leafs.Append("name", types.YLeaf{"Name", nodeInterfaceUnicastPeer.Name})
    nodeInterfaceUnicastPeer.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", nodeInterfaceUnicastPeer.PortNumber})

    nodeInterfaceUnicastPeer.EntityData.YListKeys = []string {"InterfaceName"}

    return &(nodeInterfaceUnicastPeer.EntityData)
}

// Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers
// Unicast Peers
type Ptp_Nodes_Node_NodeInterfaceUnicastPeers_NodeInterfaceUnicastPeer_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    peers.EntityData.SegmentPath = "peers" + types.AddNoKeyToken(peers)
    peers.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/" + peers.EntityData.SegmentPath
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("address", types.YChild{"Address", &peers.Address})
    peers.EntityData.Children.Append("announce-grant", types.YChild{"AnnounceGrant", &peers.AnnounceGrant})
    peers.EntityData.Children.Append("sync-grant", types.YChild{"SyncGrant", &peers.SyncGrant})
    peers.EntityData.Children.Append("delay-response-grant", types.YChild{"DelayResponseGrant", &peers.DelayResponseGrant})
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/peers/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/peers/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/peers/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    announceGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/peers/" + announceGrant.EntityData.SegmentPath
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = types.NewOrderedMap()
    announceGrant.EntityData.Leafs = types.NewOrderedMap()
    announceGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval})
    announceGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", announceGrant.GrantDuration})

    announceGrant.EntityData.YListKeys = []string {}

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
    syncGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/peers/" + syncGrant.EntityData.SegmentPath
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = types.NewOrderedMap()
    syncGrant.EntityData.Leafs = types.NewOrderedMap()
    syncGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval})
    syncGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", syncGrant.GrantDuration})

    syncGrant.EntityData.YListKeys = []string {}

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
    delayResponseGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/node-interface-unicast-peers/node-interface-unicast-peer/peers/" + delayResponseGrant.EntityData.SegmentPath
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval})
    delayResponseGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration})

    delayResponseGrant.EntityData.YListKeys = []string {}

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
    packetCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/" + packetCounters.EntityData.SegmentPath
    packetCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetCounters.EntityData.Children = types.NewOrderedMap()
    packetCounters.EntityData.Children.Append("counters", types.YChild{"Counters", &packetCounters.Counters})
    packetCounters.EntityData.Children.Append("drop-reasons", types.YChild{"DropReasons", &packetCounters.DropReasons})
    packetCounters.EntityData.Leafs = types.NewOrderedMap()

    packetCounters.EntityData.YListKeys = []string {}

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
    counters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/packet-counters/" + counters.EntityData.SegmentPath
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("announce-sent", types.YLeaf{"AnnounceSent", counters.AnnounceSent})
    counters.EntityData.Leafs.Append("announce-received", types.YLeaf{"AnnounceReceived", counters.AnnounceReceived})
    counters.EntityData.Leafs.Append("announce-dropped", types.YLeaf{"AnnounceDropped", counters.AnnounceDropped})
    counters.EntityData.Leafs.Append("sync-sent", types.YLeaf{"SyncSent", counters.SyncSent})
    counters.EntityData.Leafs.Append("sync-received", types.YLeaf{"SyncReceived", counters.SyncReceived})
    counters.EntityData.Leafs.Append("sync-dropped", types.YLeaf{"SyncDropped", counters.SyncDropped})
    counters.EntityData.Leafs.Append("follow-up-sent", types.YLeaf{"FollowUpSent", counters.FollowUpSent})
    counters.EntityData.Leafs.Append("follow-up-received", types.YLeaf{"FollowUpReceived", counters.FollowUpReceived})
    counters.EntityData.Leafs.Append("follow-up-dropped", types.YLeaf{"FollowUpDropped", counters.FollowUpDropped})
    counters.EntityData.Leafs.Append("delay-request-sent", types.YLeaf{"DelayRequestSent", counters.DelayRequestSent})
    counters.EntityData.Leafs.Append("delay-request-received", types.YLeaf{"DelayRequestReceived", counters.DelayRequestReceived})
    counters.EntityData.Leafs.Append("delay-request-dropped", types.YLeaf{"DelayRequestDropped", counters.DelayRequestDropped})
    counters.EntityData.Leafs.Append("delay-response-sent", types.YLeaf{"DelayResponseSent", counters.DelayResponseSent})
    counters.EntityData.Leafs.Append("delay-response-received", types.YLeaf{"DelayResponseReceived", counters.DelayResponseReceived})
    counters.EntityData.Leafs.Append("delay-response-dropped", types.YLeaf{"DelayResponseDropped", counters.DelayResponseDropped})
    counters.EntityData.Leafs.Append("peer-delay-request-sent", types.YLeaf{"PeerDelayRequestSent", counters.PeerDelayRequestSent})
    counters.EntityData.Leafs.Append("peer-delay-request-received", types.YLeaf{"PeerDelayRequestReceived", counters.PeerDelayRequestReceived})
    counters.EntityData.Leafs.Append("peer-delay-request-dropped", types.YLeaf{"PeerDelayRequestDropped", counters.PeerDelayRequestDropped})
    counters.EntityData.Leafs.Append("peer-delay-response-sent", types.YLeaf{"PeerDelayResponseSent", counters.PeerDelayResponseSent})
    counters.EntityData.Leafs.Append("peer-delay-response-received", types.YLeaf{"PeerDelayResponseReceived", counters.PeerDelayResponseReceived})
    counters.EntityData.Leafs.Append("peer-delay-response-dropped", types.YLeaf{"PeerDelayResponseDropped", counters.PeerDelayResponseDropped})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-sent", types.YLeaf{"PeerDelayResponseFollowUpSent", counters.PeerDelayResponseFollowUpSent})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-received", types.YLeaf{"PeerDelayResponseFollowUpReceived", counters.PeerDelayResponseFollowUpReceived})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-dropped", types.YLeaf{"PeerDelayResponseFollowUpDropped", counters.PeerDelayResponseFollowUpDropped})
    counters.EntityData.Leafs.Append("signaling-sent", types.YLeaf{"SignalingSent", counters.SignalingSent})
    counters.EntityData.Leafs.Append("signaling-received", types.YLeaf{"SignalingReceived", counters.SignalingReceived})
    counters.EntityData.Leafs.Append("signaling-dropped", types.YLeaf{"SignalingDropped", counters.SignalingDropped})
    counters.EntityData.Leafs.Append("management-sent", types.YLeaf{"ManagementSent", counters.ManagementSent})
    counters.EntityData.Leafs.Append("management-received", types.YLeaf{"ManagementReceived", counters.ManagementReceived})
    counters.EntityData.Leafs.Append("management-dropped", types.YLeaf{"ManagementDropped", counters.ManagementDropped})
    counters.EntityData.Leafs.Append("other-packets-sent", types.YLeaf{"OtherPacketsSent", counters.OtherPacketsSent})
    counters.EntityData.Leafs.Append("other-packets-received", types.YLeaf{"OtherPacketsReceived", counters.OtherPacketsReceived})
    counters.EntityData.Leafs.Append("other-packets-dropped", types.YLeaf{"OtherPacketsDropped", counters.OtherPacketsDropped})
    counters.EntityData.Leafs.Append("total-packets-sent", types.YLeaf{"TotalPacketsSent", counters.TotalPacketsSent})
    counters.EntityData.Leafs.Append("total-packets-received", types.YLeaf{"TotalPacketsReceived", counters.TotalPacketsReceived})
    counters.EntityData.Leafs.Append("total-packets-dropped", types.YLeaf{"TotalPacketsDropped", counters.TotalPacketsDropped})

    counters.EntityData.YListKeys = []string {}

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

    // Illegal clock class (255) in announce messages. The type is interface{}
    // with range: 0..4294967295.
    BadClockClass interface{}

    // Reserved Clock ID. The type is interface{} with range: 0..4294967295.
    ReservedClockId interface{}

    // Steps removed too high. The type is interface{} with range: 0..4294967295.
    StepsRemoved interface{}

    // Packet not compatible with G.8265.1 profile. The type is interface{} with
    // range: 0..4294967295.
    G82651Incompatible interface{}

    // Packet not compatible with G.8275.1 profile. The type is interface{} with
    // range: 0..4294967295.
    G82751Incompatible interface{}

    // Packet not compatible with G.8275.2 profile. The type is interface{} with
    // range: 0..4294967295.
    G82752Incompatible interface{}

    // Packet sent to incorrect address. The type is interface{} with range:
    // 0..4294967295.
    IncorrectAddress interface{}
}

func (dropReasons *Ptp_Nodes_Node_PacketCounters_DropReasons) GetEntityData() *types.CommonEntityData {
    dropReasons.EntityData.YFilter = dropReasons.YFilter
    dropReasons.EntityData.YangName = "drop-reasons"
    dropReasons.EntityData.BundleName = "cisco_ios_xr"
    dropReasons.EntityData.ParentYangName = "packet-counters"
    dropReasons.EntityData.SegmentPath = "drop-reasons"
    dropReasons.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/nodes/node/packet-counters/" + dropReasons.EntityData.SegmentPath
    dropReasons.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropReasons.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropReasons.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropReasons.EntityData.Children = types.NewOrderedMap()
    dropReasons.EntityData.Leafs = types.NewOrderedMap()
    dropReasons.EntityData.Leafs.Append("not-ready", types.YLeaf{"NotReady", dropReasons.NotReady})
    dropReasons.EntityData.Leafs.Append("wrong-domain", types.YLeaf{"WrongDomain", dropReasons.WrongDomain})
    dropReasons.EntityData.Leafs.Append("too-short", types.YLeaf{"TooShort", dropReasons.TooShort})
    dropReasons.EntityData.Leafs.Append("looped-same-port", types.YLeaf{"LoopedSamePort", dropReasons.LoopedSamePort})
    dropReasons.EntityData.Leafs.Append("looped-higher-port", types.YLeaf{"LoopedHigherPort", dropReasons.LoopedHigherPort})
    dropReasons.EntityData.Leafs.Append("looped-lower-port", types.YLeaf{"LoopedLowerPort", dropReasons.LoopedLowerPort})
    dropReasons.EntityData.Leafs.Append("no-timestamp", types.YLeaf{"NoTimestamp", dropReasons.NoTimestamp})
    dropReasons.EntityData.Leafs.Append("zero-timestamp", types.YLeaf{"ZeroTimestamp", dropReasons.ZeroTimestamp})
    dropReasons.EntityData.Leafs.Append("invalid-tl-vs", types.YLeaf{"InvalidTlVs", dropReasons.InvalidTlVs})
    dropReasons.EntityData.Leafs.Append("not-for-us", types.YLeaf{"NotForUs", dropReasons.NotForUs})
    dropReasons.EntityData.Leafs.Append("not-listening", types.YLeaf{"NotListening", dropReasons.NotListening})
    dropReasons.EntityData.Leafs.Append("wrong-master", types.YLeaf{"WrongMaster", dropReasons.WrongMaster})
    dropReasons.EntityData.Leafs.Append("unknown-master", types.YLeaf{"UnknownMaster", dropReasons.UnknownMaster})
    dropReasons.EntityData.Leafs.Append("not-master", types.YLeaf{"NotMaster", dropReasons.NotMaster})
    dropReasons.EntityData.Leafs.Append("not-slave", types.YLeaf{"NotSlave", dropReasons.NotSlave})
    dropReasons.EntityData.Leafs.Append("not-granted", types.YLeaf{"NotGranted", dropReasons.NotGranted})
    dropReasons.EntityData.Leafs.Append("too-slow", types.YLeaf{"TooSlow", dropReasons.TooSlow})
    dropReasons.EntityData.Leafs.Append("invalid-packet", types.YLeaf{"InvalidPacket", dropReasons.InvalidPacket})
    dropReasons.EntityData.Leafs.Append("wrong-sequence-id", types.YLeaf{"WrongSequenceId", dropReasons.WrongSequenceId})
    dropReasons.EntityData.Leafs.Append("no-offload-session", types.YLeaf{"NoOffloadSession", dropReasons.NoOffloadSession})
    dropReasons.EntityData.Leafs.Append("not-supported", types.YLeaf{"NotSupported", dropReasons.NotSupported})
    dropReasons.EntityData.Leafs.Append("min-clock-class", types.YLeaf{"MinClockClass", dropReasons.MinClockClass})
    dropReasons.EntityData.Leafs.Append("bad-clock-class", types.YLeaf{"BadClockClass", dropReasons.BadClockClass})
    dropReasons.EntityData.Leafs.Append("reserved-clock-id", types.YLeaf{"ReservedClockId", dropReasons.ReservedClockId})
    dropReasons.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", dropReasons.StepsRemoved})
    dropReasons.EntityData.Leafs.Append("g8265-1-incompatible", types.YLeaf{"G82651Incompatible", dropReasons.G82651Incompatible})
    dropReasons.EntityData.Leafs.Append("g8275-1-incompatible", types.YLeaf{"G82751Incompatible", dropReasons.G82751Incompatible})
    dropReasons.EntityData.Leafs.Append("g8275-2-incompatible", types.YLeaf{"G82752Incompatible", dropReasons.G82752Incompatible})
    dropReasons.EntityData.Leafs.Append("incorrect-address", types.YLeaf{"IncorrectAddress", dropReasons.IncorrectAddress})

    dropReasons.EntityData.YListKeys = []string {}

    return &(dropReasons.EntityData)
}

// Ptp_Summary
// Summary operational data
type Ptp_Summary struct {
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

func (summary *Ptp_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "ptp"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("port-state-init-count", types.YLeaf{"PortStateInitCount", summary.PortStateInitCount})
    summary.EntityData.Leafs.Append("port-state-listening-count", types.YLeaf{"PortStateListeningCount", summary.PortStateListeningCount})
    summary.EntityData.Leafs.Append("port-state-passive-count", types.YLeaf{"PortStatePassiveCount", summary.PortStatePassiveCount})
    summary.EntityData.Leafs.Append("port-state-pre-master-count", types.YLeaf{"PortStatePreMasterCount", summary.PortStatePreMasterCount})
    summary.EntityData.Leafs.Append("port-state-master-count", types.YLeaf{"PortStateMasterCount", summary.PortStateMasterCount})
    summary.EntityData.Leafs.Append("port-state-slave-count", types.YLeaf{"PortStateSlaveCount", summary.PortStateSlaveCount})
    summary.EntityData.Leafs.Append("port-state-uncalibrated-count", types.YLeaf{"PortStateUncalibratedCount", summary.PortStateUncalibratedCount})
    summary.EntityData.Leafs.Append("port-state-faulty-count", types.YLeaf{"PortStateFaultyCount", summary.PortStateFaultyCount})
    summary.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", summary.TotalInterfaces})
    summary.EntityData.Leafs.Append("total-interfaces-valid-port-num", types.YLeaf{"TotalInterfacesValidPortNum", summary.TotalInterfacesValidPortNum})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Ptp_InterfaceConfigurationErrors
// Table for interface configuration error
// operational data
type Ptp_InterfaceConfigurationErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface configuration error operational data. The type is slice of
    // Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError.
    InterfaceConfigurationError []*Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError
}

func (interfaceConfigurationErrors *Ptp_InterfaceConfigurationErrors) GetEntityData() *types.CommonEntityData {
    interfaceConfigurationErrors.EntityData.YFilter = interfaceConfigurationErrors.YFilter
    interfaceConfigurationErrors.EntityData.YangName = "interface-configuration-errors"
    interfaceConfigurationErrors.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigurationErrors.EntityData.ParentYangName = "ptp"
    interfaceConfigurationErrors.EntityData.SegmentPath = "interface-configuration-errors"
    interfaceConfigurationErrors.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + interfaceConfigurationErrors.EntityData.SegmentPath
    interfaceConfigurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigurationErrors.EntityData.Children = types.NewOrderedMap()
    interfaceConfigurationErrors.EntityData.Children.Append("interface-configuration-error", types.YChild{"InterfaceConfigurationError", nil})
    for i := range interfaceConfigurationErrors.InterfaceConfigurationError {
        interfaceConfigurationErrors.EntityData.Children.Append(types.GetSegmentPath(interfaceConfigurationErrors.InterfaceConfigurationError[i]), types.YChild{"InterfaceConfigurationError", interfaceConfigurationErrors.InterfaceConfigurationError[i]})
    }
    interfaceConfigurationErrors.EntityData.Leafs = types.NewOrderedMap()

    interfaceConfigurationErrors.EntityData.YListKeys = []string {}

    return &(interfaceConfigurationErrors.EntityData)
}

// Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError
// Interface configuration error operational data
type Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Configuration profile name, if a profile is selected. The type is string.
    ConfigurationProfileName interface{}

    // The clock profile. The type is PtpBagProfile.
    ClockProfile interface{}

    // The telecom clock type. The type is PtpBagTelecomClock.
    TelecomClockType interface{}

    // Restriction on the port state. The type is PtpBagRestrictPortState.
    RestrictPortState interface{}

    // The clock profile to interoperate with, if interoperation is configured.
    // The type is PtpBagProfile.
    InteropProfile interface{}

    // Configuration Errors.
    ConfigurationErrors Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors
}

func (interfaceConfigurationError *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError) GetEntityData() *types.CommonEntityData {
    interfaceConfigurationError.EntityData.YFilter = interfaceConfigurationError.YFilter
    interfaceConfigurationError.EntityData.YangName = "interface-configuration-error"
    interfaceConfigurationError.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigurationError.EntityData.ParentYangName = "interface-configuration-errors"
    interfaceConfigurationError.EntityData.SegmentPath = "interface-configuration-error" + types.AddKeyToken(interfaceConfigurationError.InterfaceName, "interface-name")
    interfaceConfigurationError.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-configuration-errors/" + interfaceConfigurationError.EntityData.SegmentPath
    interfaceConfigurationError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigurationError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigurationError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigurationError.EntityData.Children = types.NewOrderedMap()
    interfaceConfigurationError.EntityData.Children.Append("configuration-errors", types.YChild{"ConfigurationErrors", &interfaceConfigurationError.ConfigurationErrors})
    interfaceConfigurationError.EntityData.Leafs = types.NewOrderedMap()
    interfaceConfigurationError.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceConfigurationError.InterfaceName})
    interfaceConfigurationError.EntityData.Leafs.Append("configuration-profile-name", types.YLeaf{"ConfigurationProfileName", interfaceConfigurationError.ConfigurationProfileName})
    interfaceConfigurationError.EntityData.Leafs.Append("clock-profile", types.YLeaf{"ClockProfile", interfaceConfigurationError.ClockProfile})
    interfaceConfigurationError.EntityData.Leafs.Append("telecom-clock-type", types.YLeaf{"TelecomClockType", interfaceConfigurationError.TelecomClockType})
    interfaceConfigurationError.EntityData.Leafs.Append("restrict-port-state", types.YLeaf{"RestrictPortState", interfaceConfigurationError.RestrictPortState})
    interfaceConfigurationError.EntityData.Leafs.Append("interop-profile", types.YLeaf{"InteropProfile", interfaceConfigurationError.InteropProfile})

    interfaceConfigurationError.EntityData.YListKeys = []string {"InterfaceName"}

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

    // IPv4 TTL value configured but transport is not IPv4. The type is bool.
    Ipv4ttl interface{}

    // IPv6 hop limit value configured but transport is not IPv6. The type is
    // bool.
    Ipv6HopLimit interface{}

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

    // Domain is not compatible with configured profile interop. The type is bool.
    InvalidInteropDomain interface{}

    // Ingress conversion clock class default is not compatible with configured
    // profile interop. The type is bool.
    InvalidInteropIngressClockClassDefault interface{}

    // Ingress conversion priority1 is not compatible with configured profile
    // interop. The type is bool.
    InvalidInteropIngressPriority1 interface{}

    // Ingress conversion clock-accuracy is not compatible with configured profile
    // interop. The type is bool.
    InvalidInteropIngressClockAccuracy interface{}

    // Ingress conversion OSLV not compatible with configured profile interop. The
    // type is bool.
    InvalidInteropIngressOslv interface{}

    // Egress conversion clock class default is not compatible with configured
    // profile interop. The type is bool.
    InvalidInteropEgressClockClassDefault interface{}

    // Egress conversion priority1 is not compatible with configured profile
    // interop. The type is bool.
    InvalidInteropEgressPriority1 interface{}

    // Egress conversion priority2 is not compatible with configured profile
    // interop. The type is bool.
    InvalidInteropEgressPriority2 interface{}

    // Egress conversion clock-accuracy is not compatible with configured profile
    // interop. The type is bool.
    InvalidInteropEgressClockAccuracy interface{}

    // Egress conversion OSLV not compatible with configured profile interop. The
    // type is bool.
    InvalidInteropEgressOslv interface{}

    // Master configuration is not compatible with configured clock-type. The type
    // is bool.
    InvalidMasterConfig interface{}

    // Slave configuration is not compatible with configured clock-type. The type
    // is bool.
    InvalidSlaveConfig interface{}

    // List of ingress conversion clock class mapping 'from values' that are not
    // compatible with the configure profile interop. The type is slice of
    // interface{} with range: 0..255.
    InvalidInteropIngressClockClassMapFromVal []interface{}

    // List of ingress conversion clock class mapping 'to values' that are not
    // compatible with the configure profile interop. The type is slice of
    // interface{} with range: 0..255.
    InvalidInteropIngressClockClassMapToVal []interface{}

    // List of egress conversion clock class mapping 'from values' that are not
    // compatible with the configure profile interop. The type is slice of
    // interface{} with range: 0..255.
    InvalidInteropEgressClockClassMapFromVal []interface{}

    // List of egress conversion clock class mapping 'to values' that are not
    // compatible with the configure profile interop. The type is slice of
    // interface{} with range: 0..255.
    InvalidInteropEgressClockClassMapToVal []interface{}
}

func (configurationErrors *Ptp_InterfaceConfigurationErrors_InterfaceConfigurationError_ConfigurationErrors) GetEntityData() *types.CommonEntityData {
    configurationErrors.EntityData.YFilter = configurationErrors.YFilter
    configurationErrors.EntityData.YangName = "configuration-errors"
    configurationErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationErrors.EntityData.ParentYangName = "interface-configuration-error"
    configurationErrors.EntityData.SegmentPath = "configuration-errors"
    configurationErrors.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-configuration-errors/interface-configuration-error/" + configurationErrors.EntityData.SegmentPath
    configurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationErrors.EntityData.Children = types.NewOrderedMap()
    configurationErrors.EntityData.Leafs = types.NewOrderedMap()
    configurationErrors.EntityData.Leafs.Append("global-ptp", types.YLeaf{"GlobalPtp", configurationErrors.GlobalPtp})
    configurationErrors.EntityData.Leafs.Append("ethernet-transport", types.YLeaf{"EthernetTransport", configurationErrors.EthernetTransport})
    configurationErrors.EntityData.Leafs.Append("one-step", types.YLeaf{"OneStep", configurationErrors.OneStep})
    configurationErrors.EntityData.Leafs.Append("slave", types.YLeaf{"Slave", configurationErrors.Slave})
    configurationErrors.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", configurationErrors.Ipv6})
    configurationErrors.EntityData.Leafs.Append("multicast", types.YLeaf{"Multicast", configurationErrors.Multicast})
    configurationErrors.EntityData.Leafs.Append("profile-not-global", types.YLeaf{"ProfileNotGlobal", configurationErrors.ProfileNotGlobal})
    configurationErrors.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", configurationErrors.LocalPriority})
    configurationErrors.EntityData.Leafs.Append("profile-ethernet", types.YLeaf{"ProfileEthernet", configurationErrors.ProfileEthernet})
    configurationErrors.EntityData.Leafs.Append("profile-ipv4", types.YLeaf{"ProfileIpv4", configurationErrors.ProfileIpv4})
    configurationErrors.EntityData.Leafs.Append("profile-ipv6", types.YLeaf{"ProfileIpv6", configurationErrors.ProfileIpv6})
    configurationErrors.EntityData.Leafs.Append("profile-unicast", types.YLeaf{"ProfileUnicast", configurationErrors.ProfileUnicast})
    configurationErrors.EntityData.Leafs.Append("profile-multicast", types.YLeaf{"ProfileMulticast", configurationErrors.ProfileMulticast})
    configurationErrors.EntityData.Leafs.Append("profile-mixed", types.YLeaf{"ProfileMixed", configurationErrors.ProfileMixed})
    configurationErrors.EntityData.Leafs.Append("profile-master-unicast", types.YLeaf{"ProfileMasterUnicast", configurationErrors.ProfileMasterUnicast})
    configurationErrors.EntityData.Leafs.Append("profile-master-multicast", types.YLeaf{"ProfileMasterMulticast", configurationErrors.ProfileMasterMulticast})
    configurationErrors.EntityData.Leafs.Append("profile-master-mixed", types.YLeaf{"ProfileMasterMixed", configurationErrors.ProfileMasterMixed})
    configurationErrors.EntityData.Leafs.Append("target-address-ipv4", types.YLeaf{"TargetAddressIpv4", configurationErrors.TargetAddressIpv4})
    configurationErrors.EntityData.Leafs.Append("target-address-ipv6", types.YLeaf{"TargetAddressIpv6", configurationErrors.TargetAddressIpv6})
    configurationErrors.EntityData.Leafs.Append("ipv4ttl", types.YLeaf{"Ipv4ttl", configurationErrors.Ipv4ttl})
    configurationErrors.EntityData.Leafs.Append("ipv6-hop-limit", types.YLeaf{"Ipv6HopLimit", configurationErrors.Ipv6HopLimit})
    configurationErrors.EntityData.Leafs.Append("profile-port-state", types.YLeaf{"ProfilePortState", configurationErrors.ProfilePortState})
    configurationErrors.EntityData.Leafs.Append("profile-announce-interval", types.YLeaf{"ProfileAnnounceInterval", configurationErrors.ProfileAnnounceInterval})
    configurationErrors.EntityData.Leafs.Append("profile-sync-interval", types.YLeaf{"ProfileSyncInterval", configurationErrors.ProfileSyncInterval})
    configurationErrors.EntityData.Leafs.Append("profile-delay-req-interval", types.YLeaf{"ProfileDelayReqInterval", configurationErrors.ProfileDelayReqInterval})
    configurationErrors.EntityData.Leafs.Append("profile-sync-timeout", types.YLeaf{"ProfileSyncTimeout", configurationErrors.ProfileSyncTimeout})
    configurationErrors.EntityData.Leafs.Append("profile-delay-resp-timeout", types.YLeaf{"ProfileDelayRespTimeout", configurationErrors.ProfileDelayRespTimeout})
    configurationErrors.EntityData.Leafs.Append("invalid-grant-reduction", types.YLeaf{"InvalidGrantReduction", configurationErrors.InvalidGrantReduction})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-domain", types.YLeaf{"InvalidInteropDomain", configurationErrors.InvalidInteropDomain})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-ingress-clock-class-default", types.YLeaf{"InvalidInteropIngressClockClassDefault", configurationErrors.InvalidInteropIngressClockClassDefault})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-ingress-priority1", types.YLeaf{"InvalidInteropIngressPriority1", configurationErrors.InvalidInteropIngressPriority1})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-ingress-clock-accuracy", types.YLeaf{"InvalidInteropIngressClockAccuracy", configurationErrors.InvalidInteropIngressClockAccuracy})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-ingress-oslv", types.YLeaf{"InvalidInteropIngressOslv", configurationErrors.InvalidInteropIngressOslv})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-clock-class-default", types.YLeaf{"InvalidInteropEgressClockClassDefault", configurationErrors.InvalidInteropEgressClockClassDefault})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-priority1", types.YLeaf{"InvalidInteropEgressPriority1", configurationErrors.InvalidInteropEgressPriority1})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-priority2", types.YLeaf{"InvalidInteropEgressPriority2", configurationErrors.InvalidInteropEgressPriority2})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-clock-accuracy", types.YLeaf{"InvalidInteropEgressClockAccuracy", configurationErrors.InvalidInteropEgressClockAccuracy})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-oslv", types.YLeaf{"InvalidInteropEgressOslv", configurationErrors.InvalidInteropEgressOslv})
    configurationErrors.EntityData.Leafs.Append("invalid-master-config", types.YLeaf{"InvalidMasterConfig", configurationErrors.InvalidMasterConfig})
    configurationErrors.EntityData.Leafs.Append("invalid-slave-config", types.YLeaf{"InvalidSlaveConfig", configurationErrors.InvalidSlaveConfig})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-ingress-clock-class-map-from-val", types.YLeaf{"InvalidInteropIngressClockClassMapFromVal", configurationErrors.InvalidInteropIngressClockClassMapFromVal})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-ingress-clock-class-map-to-val", types.YLeaf{"InvalidInteropIngressClockClassMapToVal", configurationErrors.InvalidInteropIngressClockClassMapToVal})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-clock-class-map-from-val", types.YLeaf{"InvalidInteropEgressClockClassMapFromVal", configurationErrors.InvalidInteropEgressClockClassMapFromVal})
    configurationErrors.EntityData.Leafs.Append("invalid-interop-egress-clock-class-map-to-val", types.YLeaf{"InvalidInteropEgressClockClassMapToVal", configurationErrors.InvalidInteropEgressClockClassMapToVal})

    configurationErrors.EntityData.YListKeys = []string {}

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
    InterfaceForeignMaster []*Ptp_InterfaceForeignMasters_InterfaceForeignMaster
}

func (interfaceForeignMasters *Ptp_InterfaceForeignMasters) GetEntityData() *types.CommonEntityData {
    interfaceForeignMasters.EntityData.YFilter = interfaceForeignMasters.YFilter
    interfaceForeignMasters.EntityData.YangName = "interface-foreign-masters"
    interfaceForeignMasters.EntityData.BundleName = "cisco_ios_xr"
    interfaceForeignMasters.EntityData.ParentYangName = "ptp"
    interfaceForeignMasters.EntityData.SegmentPath = "interface-foreign-masters"
    interfaceForeignMasters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + interfaceForeignMasters.EntityData.SegmentPath
    interfaceForeignMasters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceForeignMasters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceForeignMasters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceForeignMasters.EntityData.Children = types.NewOrderedMap()
    interfaceForeignMasters.EntityData.Children.Append("interface-foreign-master", types.YChild{"InterfaceForeignMaster", nil})
    for i := range interfaceForeignMasters.InterfaceForeignMaster {
        interfaceForeignMasters.EntityData.Children.Append(types.GetSegmentPath(interfaceForeignMasters.InterfaceForeignMaster[i]), types.YChild{"InterfaceForeignMaster", interfaceForeignMasters.InterfaceForeignMaster[i]})
    }
    interfaceForeignMasters.EntityData.Leafs = types.NewOrderedMap()

    interfaceForeignMasters.EntityData.YListKeys = []string {}

    return &(interfaceForeignMasters.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster
// Interface foreign master clock operational data
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Foreign clocks received on this interface. The type is slice of
    // Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock.
    ForeignClock []*Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock
}

func (interfaceForeignMaster *Ptp_InterfaceForeignMasters_InterfaceForeignMaster) GetEntityData() *types.CommonEntityData {
    interfaceForeignMaster.EntityData.YFilter = interfaceForeignMaster.YFilter
    interfaceForeignMaster.EntityData.YangName = "interface-foreign-master"
    interfaceForeignMaster.EntityData.BundleName = "cisco_ios_xr"
    interfaceForeignMaster.EntityData.ParentYangName = "interface-foreign-masters"
    interfaceForeignMaster.EntityData.SegmentPath = "interface-foreign-master" + types.AddKeyToken(interfaceForeignMaster.InterfaceName, "interface-name")
    interfaceForeignMaster.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/" + interfaceForeignMaster.EntityData.SegmentPath
    interfaceForeignMaster.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceForeignMaster.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceForeignMaster.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceForeignMaster.EntityData.Children = types.NewOrderedMap()
    interfaceForeignMaster.EntityData.Children.Append("foreign-clock", types.YChild{"ForeignClock", nil})
    for i := range interfaceForeignMaster.ForeignClock {
        types.SetYListKey(interfaceForeignMaster.ForeignClock[i], i)
        interfaceForeignMaster.EntityData.Children.Append(types.GetSegmentPath(interfaceForeignMaster.ForeignClock[i]), types.YChild{"ForeignClock", interfaceForeignMaster.ForeignClock[i]})
    }
    interfaceForeignMaster.EntityData.Leafs = types.NewOrderedMap()
    interfaceForeignMaster.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceForeignMaster.InterfaceName})
    interfaceForeignMaster.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", interfaceForeignMaster.PortNumber})

    interfaceForeignMaster.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceForeignMaster.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock
// Foreign clocks received on this interface
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // The clock has clock class corresponding to QL-DNU. The type is bool.
    IsDnu interface{}

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

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock) GetEntityData() *types.CommonEntityData {
    foreignClock.EntityData.YFilter = foreignClock.YFilter
    foreignClock.EntityData.YangName = "foreign-clock"
    foreignClock.EntityData.BundleName = "cisco_ios_xr"
    foreignClock.EntityData.ParentYangName = "interface-foreign-master"
    foreignClock.EntityData.SegmentPath = "foreign-clock" + types.AddNoKeyToken(foreignClock)
    foreignClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/" + foreignClock.EntityData.SegmentPath
    foreignClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock.EntityData.Children = types.NewOrderedMap()
    foreignClock.EntityData.Children.Append("foreign-clock", types.YChild{"ForeignClock", &foreignClock.ForeignClock})
    foreignClock.EntityData.Children.Append("address", types.YChild{"Address", &foreignClock.Address})
    foreignClock.EntityData.Children.Append("announce-grant", types.YChild{"AnnounceGrant", &foreignClock.AnnounceGrant})
    foreignClock.EntityData.Children.Append("sync-grant", types.YChild{"SyncGrant", &foreignClock.SyncGrant})
    foreignClock.EntityData.Children.Append("delay-response-grant", types.YChild{"DelayResponseGrant", &foreignClock.DelayResponseGrant})
    foreignClock.EntityData.Leafs = types.NewOrderedMap()
    foreignClock.EntityData.Leafs.Append("is-qualified", types.YLeaf{"IsQualified", foreignClock.IsQualified})
    foreignClock.EntityData.Leafs.Append("is-grandmaster", types.YLeaf{"IsGrandmaster", foreignClock.IsGrandmaster})
    foreignClock.EntityData.Leafs.Append("communication-model", types.YLeaf{"CommunicationModel", foreignClock.CommunicationModel})
    foreignClock.EntityData.Leafs.Append("is-known", types.YLeaf{"IsKnown", foreignClock.IsKnown})
    foreignClock.EntityData.Leafs.Append("time-known-for", types.YLeaf{"TimeKnownFor", foreignClock.TimeKnownFor})
    foreignClock.EntityData.Leafs.Append("foreign-domain", types.YLeaf{"ForeignDomain", foreignClock.ForeignDomain})
    foreignClock.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", foreignClock.ConfiguredPriority})
    foreignClock.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", foreignClock.ConfiguredClockClass})
    foreignClock.EntityData.Leafs.Append("delay-asymmetry", types.YLeaf{"DelayAsymmetry", foreignClock.DelayAsymmetry})
    foreignClock.EntityData.Leafs.Append("ptsf-loss-announce", types.YLeaf{"PtsfLossAnnounce", foreignClock.PtsfLossAnnounce})
    foreignClock.EntityData.Leafs.Append("ptsf-loss-sync", types.YLeaf{"PtsfLossSync", foreignClock.PtsfLossSync})
    foreignClock.EntityData.Leafs.Append("is-dnu", types.YLeaf{"IsDnu", foreignClock.IsDnu})

    foreignClock.EntityData.YListKeys = []string {}

    return &(foreignClock.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock
// Foreign clock information
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock struct {
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
    UtcOffset Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset

    // Receiver.
    Receiver Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver

    // Sender.
    Sender Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender
}

func (foreignClock *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock) GetEntityData() *types.CommonEntityData {
    foreignClock.EntityData.YFilter = foreignClock.YFilter
    foreignClock.EntityData.YangName = "foreign-clock"
    foreignClock.EntityData.BundleName = "cisco_ios_xr"
    foreignClock.EntityData.ParentYangName = "foreign-clock"
    foreignClock.EntityData.SegmentPath = "foreign-clock"
    foreignClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/" + foreignClock.EntityData.SegmentPath
    foreignClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignClock.EntityData.Children = types.NewOrderedMap()
    foreignClock.EntityData.Children.Append("utc-offset", types.YChild{"UtcOffset", &foreignClock.UtcOffset})
    foreignClock.EntityData.Children.Append("receiver", types.YChild{"Receiver", &foreignClock.Receiver})
    foreignClock.EntityData.Children.Append("sender", types.YChild{"Sender", &foreignClock.Sender})
    foreignClock.EntityData.Leafs = types.NewOrderedMap()
    foreignClock.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", foreignClock.ClockId})
    foreignClock.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", foreignClock.Priority1})
    foreignClock.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", foreignClock.Priority2})
    foreignClock.EntityData.Leafs.Append("class", types.YLeaf{"Class", foreignClock.Class})
    foreignClock.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", foreignClock.Accuracy})
    foreignClock.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", foreignClock.OffsetLogVariance})
    foreignClock.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", foreignClock.StepsRemoved})
    foreignClock.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", foreignClock.TimeSource})
    foreignClock.EntityData.Leafs.Append("frequency-traceable", types.YLeaf{"FrequencyTraceable", foreignClock.FrequencyTraceable})
    foreignClock.EntityData.Leafs.Append("time-traceable", types.YLeaf{"TimeTraceable", foreignClock.TimeTraceable})
    foreignClock.EntityData.Leafs.Append("timescale", types.YLeaf{"Timescale", foreignClock.Timescale})
    foreignClock.EntityData.Leafs.Append("leap-seconds", types.YLeaf{"LeapSeconds", foreignClock.LeapSeconds})
    foreignClock.EntityData.Leafs.Append("local", types.YLeaf{"Local", foreignClock.Local})
    foreignClock.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", foreignClock.ConfiguredClockClass})
    foreignClock.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", foreignClock.ConfiguredPriority})

    foreignClock.EntityData.YListKeys = []string {}

    return &(foreignClock.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset
// UTC offset
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current offset. The type is interface{} with range: -32768..32767.
    CurrentOffset interface{}

    // The current offset is valid. The type is bool.
    OffsetValid interface{}
}

func (utcOffset *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_UtcOffset) GetEntityData() *types.CommonEntityData {
    utcOffset.EntityData.YFilter = utcOffset.YFilter
    utcOffset.EntityData.YangName = "utc-offset"
    utcOffset.EntityData.BundleName = "cisco_ios_xr"
    utcOffset.EntityData.ParentYangName = "foreign-clock"
    utcOffset.EntityData.SegmentPath = "utc-offset"
    utcOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/foreign-clock/" + utcOffset.EntityData.SegmentPath
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = types.NewOrderedMap()
    utcOffset.EntityData.Leafs = types.NewOrderedMap()
    utcOffset.EntityData.Leafs.Append("current-offset", types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset})
    utcOffset.EntityData.Leafs.Append("offset-valid", types.YLeaf{"OffsetValid", utcOffset.OffsetValid})

    utcOffset.EntityData.YListKeys = []string {}

    return &(utcOffset.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver
// Receiver
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (receiver *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Receiver) GetEntityData() *types.CommonEntityData {
    receiver.EntityData.YFilter = receiver.YFilter
    receiver.EntityData.YangName = "receiver"
    receiver.EntityData.BundleName = "cisco_ios_xr"
    receiver.EntityData.ParentYangName = "foreign-clock"
    receiver.EntityData.SegmentPath = "receiver"
    receiver.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/foreign-clock/" + receiver.EntityData.SegmentPath
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", receiver.ClockId})
    receiver.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", receiver.PortNumber})

    receiver.EntityData.YListKeys = []string {}

    return &(receiver.EntityData)
}

// Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender
// Sender
type Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock ID. The type is interface{} with range: 0..18446744073709551615.
    ClockId interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}
}

func (sender *Ptp_InterfaceForeignMasters_InterfaceForeignMaster_ForeignClock_ForeignClock_Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "foreign-clock"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/foreign-clock/" + sender.EntityData.SegmentPath
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = types.NewOrderedMap()
    sender.EntityData.Leafs = types.NewOrderedMap()
    sender.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", sender.ClockId})
    sender.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", sender.PortNumber})

    sender.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    announceGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/" + announceGrant.EntityData.SegmentPath
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = types.NewOrderedMap()
    announceGrant.EntityData.Leafs = types.NewOrderedMap()
    announceGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval})
    announceGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", announceGrant.GrantDuration})

    announceGrant.EntityData.YListKeys = []string {}

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
    syncGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/" + syncGrant.EntityData.SegmentPath
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = types.NewOrderedMap()
    syncGrant.EntityData.Leafs = types.NewOrderedMap()
    syncGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval})
    syncGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", syncGrant.GrantDuration})

    syncGrant.EntityData.YListKeys = []string {}

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
    delayResponseGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-foreign-masters/interface-foreign-master/foreign-clock/" + delayResponseGrant.EntityData.SegmentPath
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval})
    delayResponseGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration})

    delayResponseGrant.EntityData.YListKeys = []string {}

    return &(delayResponseGrant.EntityData)
}

// Ptp_InterfaceInterops
// Table for interface interop operational data
type Ptp_InterfaceInterops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface interop operational data. The type is slice of
    // Ptp_InterfaceInterops_InterfaceInterop.
    InterfaceInterop []*Ptp_InterfaceInterops_InterfaceInterop
}

func (interfaceInterops *Ptp_InterfaceInterops) GetEntityData() *types.CommonEntityData {
    interfaceInterops.EntityData.YFilter = interfaceInterops.YFilter
    interfaceInterops.EntityData.YangName = "interface-interops"
    interfaceInterops.EntityData.BundleName = "cisco_ios_xr"
    interfaceInterops.EntityData.ParentYangName = "ptp"
    interfaceInterops.EntityData.SegmentPath = "interface-interops"
    interfaceInterops.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + interfaceInterops.EntityData.SegmentPath
    interfaceInterops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInterops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInterops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInterops.EntityData.Children = types.NewOrderedMap()
    interfaceInterops.EntityData.Children.Append("interface-interop", types.YChild{"InterfaceInterop", nil})
    for i := range interfaceInterops.InterfaceInterop {
        interfaceInterops.EntityData.Children.Append(types.GetSegmentPath(interfaceInterops.InterfaceInterop[i]), types.YChild{"InterfaceInterop", interfaceInterops.InterfaceInterop[i]})
    }
    interfaceInterops.EntityData.Leafs = types.NewOrderedMap()

    interfaceInterops.EntityData.YListKeys = []string {}

    return &(interfaceInterops.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop
// Interface interop operational data
type Ptp_InterfaceInterops_InterfaceInterop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // The PTP domain configured for this interface. The type is interface{} with
    // range: 0..255.
    LocalDomain interface{}

    // The PTP domain that is being interoperated with. The type is interface{}
    // with range: 0..255.
    InteropDomain interface{}

    // The PTP Profile configured for this interface. The type is PtpBagProfile.
    LocalProfile interface{}

    // The PTP profile that is being interoperated with. The type is
    // PtpBagProfile.
    InteropProfile interface{}

    // Egress interop information.
    EgressInterop Ptp_InterfaceInterops_InterfaceInterop_EgressInterop

    // Per-peer ingress interop information. The type is slice of
    // Ptp_InterfaceInterops_InterfaceInterop_IngressInterop.
    IngressInterop []*Ptp_InterfaceInterops_InterfaceInterop_IngressInterop
}

func (interfaceInterop *Ptp_InterfaceInterops_InterfaceInterop) GetEntityData() *types.CommonEntityData {
    interfaceInterop.EntityData.YFilter = interfaceInterop.YFilter
    interfaceInterop.EntityData.YangName = "interface-interop"
    interfaceInterop.EntityData.BundleName = "cisco_ios_xr"
    interfaceInterop.EntityData.ParentYangName = "interface-interops"
    interfaceInterop.EntityData.SegmentPath = "interface-interop" + types.AddKeyToken(interfaceInterop.InterfaceName, "interface-name")
    interfaceInterop.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/" + interfaceInterop.EntityData.SegmentPath
    interfaceInterop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInterop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInterop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInterop.EntityData.Children = types.NewOrderedMap()
    interfaceInterop.EntityData.Children.Append("egress-interop", types.YChild{"EgressInterop", &interfaceInterop.EgressInterop})
    interfaceInterop.EntityData.Children.Append("ingress-interop", types.YChild{"IngressInterop", nil})
    for i := range interfaceInterop.IngressInterop {
        types.SetYListKey(interfaceInterop.IngressInterop[i], i)
        interfaceInterop.EntityData.Children.Append(types.GetSegmentPath(interfaceInterop.IngressInterop[i]), types.YChild{"IngressInterop", interfaceInterop.IngressInterop[i]})
    }
    interfaceInterop.EntityData.Leafs = types.NewOrderedMap()
    interfaceInterop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceInterop.InterfaceName})
    interfaceInterop.EntityData.Leafs.Append("local-domain", types.YLeaf{"LocalDomain", interfaceInterop.LocalDomain})
    interfaceInterop.EntityData.Leafs.Append("interop-domain", types.YLeaf{"InteropDomain", interfaceInterop.InteropDomain})
    interfaceInterop.EntityData.Leafs.Append("local-profile", types.YLeaf{"LocalProfile", interfaceInterop.LocalProfile})
    interfaceInterop.EntityData.Leafs.Append("interop-profile", types.YLeaf{"InteropProfile", interfaceInterop.InteropProfile})

    interfaceInterop.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceInterop.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop_EgressInterop
// Egress interop information
type Ptp_InterfaceInterops_InterfaceInterop_EgressInterop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // From Priority 1. The type is interface{} with range: 0..255.
    FromPriority1 interface{}

    // To Priority 1. The type is interface{} with range: 0..255.
    ToPriority1 interface{}

    // From Priority 2. The type is interface{} with range: 0..255.
    FromPriority2 interface{}

    // To Priority 2. The type is interface{} with range: 0..255.
    ToPriority2 interface{}

    // From Accuracy. The type is interface{} with range: 0..255.
    FromAccuracy interface{}

    // To Accuracy. The type is interface{} with range: 0..255.
    ToAccuracy interface{}

    // From Clock Class. The type is interface{} with range: 0..255.
    FromClockClass interface{}

    // To Clock Class. The type is interface{} with range: 0..255.
    ToClockClass interface{}

    // From Offset log variance. The type is interface{} with range: 0..65535.
    FromOffsetLogVariance interface{}

    // To Offset log variance. The type is interface{} with range: 0..65535.
    ToOffsetLogVariance interface{}
}

func (egressInterop *Ptp_InterfaceInterops_InterfaceInterop_EgressInterop) GetEntityData() *types.CommonEntityData {
    egressInterop.EntityData.YFilter = egressInterop.YFilter
    egressInterop.EntityData.YangName = "egress-interop"
    egressInterop.EntityData.BundleName = "cisco_ios_xr"
    egressInterop.EntityData.ParentYangName = "interface-interop"
    egressInterop.EntityData.SegmentPath = "egress-interop"
    egressInterop.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/interface-interop/" + egressInterop.EntityData.SegmentPath
    egressInterop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressInterop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressInterop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressInterop.EntityData.Children = types.NewOrderedMap()
    egressInterop.EntityData.Leafs = types.NewOrderedMap()
    egressInterop.EntityData.Leafs.Append("from-priority1", types.YLeaf{"FromPriority1", egressInterop.FromPriority1})
    egressInterop.EntityData.Leafs.Append("to-priority1", types.YLeaf{"ToPriority1", egressInterop.ToPriority1})
    egressInterop.EntityData.Leafs.Append("from-priority2", types.YLeaf{"FromPriority2", egressInterop.FromPriority2})
    egressInterop.EntityData.Leafs.Append("to-priority2", types.YLeaf{"ToPriority2", egressInterop.ToPriority2})
    egressInterop.EntityData.Leafs.Append("from-accuracy", types.YLeaf{"FromAccuracy", egressInterop.FromAccuracy})
    egressInterop.EntityData.Leafs.Append("to-accuracy", types.YLeaf{"ToAccuracy", egressInterop.ToAccuracy})
    egressInterop.EntityData.Leafs.Append("from-clock-class", types.YLeaf{"FromClockClass", egressInterop.FromClockClass})
    egressInterop.EntityData.Leafs.Append("to-clock-class", types.YLeaf{"ToClockClass", egressInterop.ToClockClass})
    egressInterop.EntityData.Leafs.Append("from-offset-log-variance", types.YLeaf{"FromOffsetLogVariance", egressInterop.FromOffsetLogVariance})
    egressInterop.EntityData.Leafs.Append("to-offset-log-variance", types.YLeaf{"ToOffsetLogVariance", egressInterop.ToOffsetLogVariance})

    egressInterop.EntityData.YListKeys = []string {}

    return &(egressInterop.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop_IngressInterop
// Per-peer ingress interop information
type Ptp_InterfaceInterops_InterfaceInterop_IngressInterop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Peer address.
    Address Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address

    // Interop information.
    Interop Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Interop
}

func (ingressInterop *Ptp_InterfaceInterops_InterfaceInterop_IngressInterop) GetEntityData() *types.CommonEntityData {
    ingressInterop.EntityData.YFilter = ingressInterop.YFilter
    ingressInterop.EntityData.YangName = "ingress-interop"
    ingressInterop.EntityData.BundleName = "cisco_ios_xr"
    ingressInterop.EntityData.ParentYangName = "interface-interop"
    ingressInterop.EntityData.SegmentPath = "ingress-interop" + types.AddNoKeyToken(ingressInterop)
    ingressInterop.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/interface-interop/" + ingressInterop.EntityData.SegmentPath
    ingressInterop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressInterop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressInterop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressInterop.EntityData.Children = types.NewOrderedMap()
    ingressInterop.EntityData.Children.Append("address", types.YChild{"Address", &ingressInterop.Address})
    ingressInterop.EntityData.Children.Append("interop", types.YChild{"Interop", &ingressInterop.Interop})
    ingressInterop.EntityData.Leafs = types.NewOrderedMap()

    ingressInterop.EntityData.YListKeys = []string {}

    return &(ingressInterop.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address
// Peer address
type Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address struct {
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
    MacAddress Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_MacAddress

    // IPv6 address.
    Ipv6Address Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_Ipv6Address
}

func (address *Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "ingress-interop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/interface-interop/ingress-interop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_MacAddress
// Ethernet MAC address
type Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_MacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macaddr. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Macaddr interface{}
}

func (macAddress *Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_MacAddress) GetEntityData() *types.CommonEntityData {
    macAddress.EntityData.YFilter = macAddress.YFilter
    macAddress.EntityData.YangName = "mac-address"
    macAddress.EntityData.BundleName = "cisco_ios_xr"
    macAddress.EntityData.ParentYangName = "address"
    macAddress.EntityData.SegmentPath = "mac-address"
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/interface-interop/ingress-interop/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

    return &(macAddress.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_Ipv6Address
// IPv6 address
type Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (ipv6Address *Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Address_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "address"
    ipv6Address.EntityData.SegmentPath = "ipv6-address"
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/interface-interop/ingress-interop/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

    return &(ipv6Address.EntityData)
}

// Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Interop
// Interop information
type Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Interop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // From Priority 1. The type is interface{} with range: 0..255.
    FromPriority1 interface{}

    // To Priority 1. The type is interface{} with range: 0..255.
    ToPriority1 interface{}

    // From Priority 2. The type is interface{} with range: 0..255.
    FromPriority2 interface{}

    // To Priority 2. The type is interface{} with range: 0..255.
    ToPriority2 interface{}

    // From Accuracy. The type is interface{} with range: 0..255.
    FromAccuracy interface{}

    // To Accuracy. The type is interface{} with range: 0..255.
    ToAccuracy interface{}

    // From Clock Class. The type is interface{} with range: 0..255.
    FromClockClass interface{}

    // To Clock Class. The type is interface{} with range: 0..255.
    ToClockClass interface{}

    // From Offset log variance. The type is interface{} with range: 0..65535.
    FromOffsetLogVariance interface{}

    // To Offset log variance. The type is interface{} with range: 0..65535.
    ToOffsetLogVariance interface{}
}

func (interop *Ptp_InterfaceInterops_InterfaceInterop_IngressInterop_Interop) GetEntityData() *types.CommonEntityData {
    interop.EntityData.YFilter = interop.YFilter
    interop.EntityData.YangName = "interop"
    interop.EntityData.BundleName = "cisco_ios_xr"
    interop.EntityData.ParentYangName = "ingress-interop"
    interop.EntityData.SegmentPath = "interop"
    interop.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-interops/interface-interop/ingress-interop/" + interop.EntityData.SegmentPath
    interop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interop.EntityData.Children = types.NewOrderedMap()
    interop.EntityData.Leafs = types.NewOrderedMap()
    interop.EntityData.Leafs.Append("from-priority1", types.YLeaf{"FromPriority1", interop.FromPriority1})
    interop.EntityData.Leafs.Append("to-priority1", types.YLeaf{"ToPriority1", interop.ToPriority1})
    interop.EntityData.Leafs.Append("from-priority2", types.YLeaf{"FromPriority2", interop.FromPriority2})
    interop.EntityData.Leafs.Append("to-priority2", types.YLeaf{"ToPriority2", interop.ToPriority2})
    interop.EntityData.Leafs.Append("from-accuracy", types.YLeaf{"FromAccuracy", interop.FromAccuracy})
    interop.EntityData.Leafs.Append("to-accuracy", types.YLeaf{"ToAccuracy", interop.ToAccuracy})
    interop.EntityData.Leafs.Append("from-clock-class", types.YLeaf{"FromClockClass", interop.FromClockClass})
    interop.EntityData.Leafs.Append("to-clock-class", types.YLeaf{"ToClockClass", interop.ToClockClass})
    interop.EntityData.Leafs.Append("from-offset-log-variance", types.YLeaf{"FromOffsetLogVariance", interop.FromOffsetLogVariance})
    interop.EntityData.Leafs.Append("to-offset-log-variance", types.YLeaf{"ToOffsetLogVariance", interop.ToOffsetLogVariance})

    interop.EntityData.YListKeys = []string {}

    return &(interop.EntityData)
}

// Ptp_LocalClock
// Local clock operational data
type Ptp_LocalClock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The PTP domain of the local clock. The type is interface{} with range:
    // 0..255.
    Domain interface{}

    // Whether the local clock is the grandmaster. The type is bool.
    Grandmaster interface{}

    // Local clock.
    ClockProperties Ptp_LocalClock_ClockProperties

    // Virtual port.
    VirtualPort Ptp_LocalClock_VirtualPort
}

func (localClock *Ptp_LocalClock) GetEntityData() *types.CommonEntityData {
    localClock.EntityData.YFilter = localClock.YFilter
    localClock.EntityData.YangName = "local-clock"
    localClock.EntityData.BundleName = "cisco_ios_xr"
    localClock.EntityData.ParentYangName = "ptp"
    localClock.EntityData.SegmentPath = "local-clock"
    localClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + localClock.EntityData.SegmentPath
    localClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localClock.EntityData.Children = types.NewOrderedMap()
    localClock.EntityData.Children.Append("clock-properties", types.YChild{"ClockProperties", &localClock.ClockProperties})
    localClock.EntityData.Children.Append("virtual-port", types.YChild{"VirtualPort", &localClock.VirtualPort})
    localClock.EntityData.Leafs = types.NewOrderedMap()
    localClock.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", localClock.Domain})
    localClock.EntityData.Leafs.Append("grandmaster", types.YLeaf{"Grandmaster", localClock.Grandmaster})

    localClock.EntityData.YListKeys = []string {}

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
    clockProperties.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/local-clock/" + clockProperties.EntityData.SegmentPath
    clockProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockProperties.EntityData.Children = types.NewOrderedMap()
    clockProperties.EntityData.Children.Append("utc-offset", types.YChild{"UtcOffset", &clockProperties.UtcOffset})
    clockProperties.EntityData.Children.Append("receiver", types.YChild{"Receiver", &clockProperties.Receiver})
    clockProperties.EntityData.Children.Append("sender", types.YChild{"Sender", &clockProperties.Sender})
    clockProperties.EntityData.Leafs = types.NewOrderedMap()
    clockProperties.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", clockProperties.ClockId})
    clockProperties.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", clockProperties.Priority1})
    clockProperties.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", clockProperties.Priority2})
    clockProperties.EntityData.Leafs.Append("class", types.YLeaf{"Class", clockProperties.Class})
    clockProperties.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", clockProperties.Accuracy})
    clockProperties.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", clockProperties.OffsetLogVariance})
    clockProperties.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", clockProperties.StepsRemoved})
    clockProperties.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", clockProperties.TimeSource})
    clockProperties.EntityData.Leafs.Append("frequency-traceable", types.YLeaf{"FrequencyTraceable", clockProperties.FrequencyTraceable})
    clockProperties.EntityData.Leafs.Append("time-traceable", types.YLeaf{"TimeTraceable", clockProperties.TimeTraceable})
    clockProperties.EntityData.Leafs.Append("timescale", types.YLeaf{"Timescale", clockProperties.Timescale})
    clockProperties.EntityData.Leafs.Append("leap-seconds", types.YLeaf{"LeapSeconds", clockProperties.LeapSeconds})
    clockProperties.EntityData.Leafs.Append("local", types.YLeaf{"Local", clockProperties.Local})
    clockProperties.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", clockProperties.ConfiguredClockClass})
    clockProperties.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", clockProperties.ConfiguredPriority})

    clockProperties.EntityData.YListKeys = []string {}

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
    utcOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/local-clock/clock-properties/" + utcOffset.EntityData.SegmentPath
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = types.NewOrderedMap()
    utcOffset.EntityData.Leafs = types.NewOrderedMap()
    utcOffset.EntityData.Leafs.Append("current-offset", types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset})
    utcOffset.EntityData.Leafs.Append("offset-valid", types.YLeaf{"OffsetValid", utcOffset.OffsetValid})

    utcOffset.EntityData.YListKeys = []string {}

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
    receiver.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/local-clock/clock-properties/" + receiver.EntityData.SegmentPath
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", receiver.ClockId})
    receiver.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", receiver.PortNumber})

    receiver.EntityData.YListKeys = []string {}

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
    sender.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/local-clock/clock-properties/" + sender.EntityData.SegmentPath
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = types.NewOrderedMap()
    sender.EntityData.Leafs = types.NewOrderedMap()
    sender.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", sender.ClockId})
    sender.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", sender.PortNumber})

    sender.EntityData.YListKeys = []string {}

    return &(sender.EntityData)
}

// Ptp_LocalClock_VirtualPort
// Virtual port
type Ptp_LocalClock_VirtualPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured. The type is bool.
    Configured interface{}

    // Connected. The type is bool.
    Connected interface{}

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

    // The local priority. The type is interface{} with range: 0..255.
    LocalPriority interface{}
}

func (virtualPort *Ptp_LocalClock_VirtualPort) GetEntityData() *types.CommonEntityData {
    virtualPort.EntityData.YFilter = virtualPort.YFilter
    virtualPort.EntityData.YangName = "virtual-port"
    virtualPort.EntityData.BundleName = "cisco_ios_xr"
    virtualPort.EntityData.ParentYangName = "local-clock"
    virtualPort.EntityData.SegmentPath = "virtual-port"
    virtualPort.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/local-clock/" + virtualPort.EntityData.SegmentPath
    virtualPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualPort.EntityData.Children = types.NewOrderedMap()
    virtualPort.EntityData.Leafs = types.NewOrderedMap()
    virtualPort.EntityData.Leafs.Append("configured", types.YLeaf{"Configured", virtualPort.Configured})
    virtualPort.EntityData.Leafs.Append("connected", types.YLeaf{"Connected", virtualPort.Connected})
    virtualPort.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", virtualPort.Priority1})
    virtualPort.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", virtualPort.Priority2})
    virtualPort.EntityData.Leafs.Append("class", types.YLeaf{"Class", virtualPort.Class})
    virtualPort.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", virtualPort.Accuracy})
    virtualPort.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", virtualPort.OffsetLogVariance})
    virtualPort.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", virtualPort.LocalPriority})

    virtualPort.EntityData.YListKeys = []string {}

    return &(virtualPort.EntityData)
}

// Ptp_InterfacePacketCounters
// Table for interface packet counter operational
// data
type Ptp_InterfacePacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface packet counter operational data. The type is slice of
    // Ptp_InterfacePacketCounters_InterfacePacketCounter.
    InterfacePacketCounter []*Ptp_InterfacePacketCounters_InterfacePacketCounter
}

func (interfacePacketCounters *Ptp_InterfacePacketCounters) GetEntityData() *types.CommonEntityData {
    interfacePacketCounters.EntityData.YFilter = interfacePacketCounters.YFilter
    interfacePacketCounters.EntityData.YangName = "interface-packet-counters"
    interfacePacketCounters.EntityData.BundleName = "cisco_ios_xr"
    interfacePacketCounters.EntityData.ParentYangName = "ptp"
    interfacePacketCounters.EntityData.SegmentPath = "interface-packet-counters"
    interfacePacketCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + interfacePacketCounters.EntityData.SegmentPath
    interfacePacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacePacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacePacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacePacketCounters.EntityData.Children = types.NewOrderedMap()
    interfacePacketCounters.EntityData.Children.Append("interface-packet-counter", types.YChild{"InterfacePacketCounter", nil})
    for i := range interfacePacketCounters.InterfacePacketCounter {
        interfacePacketCounters.EntityData.Children.Append(types.GetSegmentPath(interfacePacketCounters.InterfacePacketCounter[i]), types.YChild{"InterfacePacketCounter", interfacePacketCounters.InterfacePacketCounter[i]})
    }
    interfacePacketCounters.EntityData.Leafs = types.NewOrderedMap()

    interfacePacketCounters.EntityData.YListKeys = []string {}

    return &(interfacePacketCounters.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter
// Interface packet counter operational data
type Ptp_InterfacePacketCounters_InterfacePacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Packet counters.
    Counters Ptp_InterfacePacketCounters_InterfacePacketCounter_Counters

    // Packet counters for each peer on this interface. The type is slice of
    // Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter.
    PeerCounter []*Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter
}

func (interfacePacketCounter *Ptp_InterfacePacketCounters_InterfacePacketCounter) GetEntityData() *types.CommonEntityData {
    interfacePacketCounter.EntityData.YFilter = interfacePacketCounter.YFilter
    interfacePacketCounter.EntityData.YangName = "interface-packet-counter"
    interfacePacketCounter.EntityData.BundleName = "cisco_ios_xr"
    interfacePacketCounter.EntityData.ParentYangName = "interface-packet-counters"
    interfacePacketCounter.EntityData.SegmentPath = "interface-packet-counter" + types.AddKeyToken(interfacePacketCounter.InterfaceName, "interface-name")
    interfacePacketCounter.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/" + interfacePacketCounter.EntityData.SegmentPath
    interfacePacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfacePacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfacePacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfacePacketCounter.EntityData.Children = types.NewOrderedMap()
    interfacePacketCounter.EntityData.Children.Append("counters", types.YChild{"Counters", &interfacePacketCounter.Counters})
    interfacePacketCounter.EntityData.Children.Append("peer-counter", types.YChild{"PeerCounter", nil})
    for i := range interfacePacketCounter.PeerCounter {
        types.SetYListKey(interfacePacketCounter.PeerCounter[i], i)
        interfacePacketCounter.EntityData.Children.Append(types.GetSegmentPath(interfacePacketCounter.PeerCounter[i]), types.YChild{"PeerCounter", interfacePacketCounter.PeerCounter[i]})
    }
    interfacePacketCounter.EntityData.Leafs = types.NewOrderedMap()
    interfacePacketCounter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfacePacketCounter.InterfaceName})

    interfacePacketCounter.EntityData.YListKeys = []string {"InterfaceName"}

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
    counters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/interface-packet-counter/" + counters.EntityData.SegmentPath
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("announce-sent", types.YLeaf{"AnnounceSent", counters.AnnounceSent})
    counters.EntityData.Leafs.Append("announce-received", types.YLeaf{"AnnounceReceived", counters.AnnounceReceived})
    counters.EntityData.Leafs.Append("announce-dropped", types.YLeaf{"AnnounceDropped", counters.AnnounceDropped})
    counters.EntityData.Leafs.Append("sync-sent", types.YLeaf{"SyncSent", counters.SyncSent})
    counters.EntityData.Leafs.Append("sync-received", types.YLeaf{"SyncReceived", counters.SyncReceived})
    counters.EntityData.Leafs.Append("sync-dropped", types.YLeaf{"SyncDropped", counters.SyncDropped})
    counters.EntityData.Leafs.Append("follow-up-sent", types.YLeaf{"FollowUpSent", counters.FollowUpSent})
    counters.EntityData.Leafs.Append("follow-up-received", types.YLeaf{"FollowUpReceived", counters.FollowUpReceived})
    counters.EntityData.Leafs.Append("follow-up-dropped", types.YLeaf{"FollowUpDropped", counters.FollowUpDropped})
    counters.EntityData.Leafs.Append("delay-request-sent", types.YLeaf{"DelayRequestSent", counters.DelayRequestSent})
    counters.EntityData.Leafs.Append("delay-request-received", types.YLeaf{"DelayRequestReceived", counters.DelayRequestReceived})
    counters.EntityData.Leafs.Append("delay-request-dropped", types.YLeaf{"DelayRequestDropped", counters.DelayRequestDropped})
    counters.EntityData.Leafs.Append("delay-response-sent", types.YLeaf{"DelayResponseSent", counters.DelayResponseSent})
    counters.EntityData.Leafs.Append("delay-response-received", types.YLeaf{"DelayResponseReceived", counters.DelayResponseReceived})
    counters.EntityData.Leafs.Append("delay-response-dropped", types.YLeaf{"DelayResponseDropped", counters.DelayResponseDropped})
    counters.EntityData.Leafs.Append("peer-delay-request-sent", types.YLeaf{"PeerDelayRequestSent", counters.PeerDelayRequestSent})
    counters.EntityData.Leafs.Append("peer-delay-request-received", types.YLeaf{"PeerDelayRequestReceived", counters.PeerDelayRequestReceived})
    counters.EntityData.Leafs.Append("peer-delay-request-dropped", types.YLeaf{"PeerDelayRequestDropped", counters.PeerDelayRequestDropped})
    counters.EntityData.Leafs.Append("peer-delay-response-sent", types.YLeaf{"PeerDelayResponseSent", counters.PeerDelayResponseSent})
    counters.EntityData.Leafs.Append("peer-delay-response-received", types.YLeaf{"PeerDelayResponseReceived", counters.PeerDelayResponseReceived})
    counters.EntityData.Leafs.Append("peer-delay-response-dropped", types.YLeaf{"PeerDelayResponseDropped", counters.PeerDelayResponseDropped})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-sent", types.YLeaf{"PeerDelayResponseFollowUpSent", counters.PeerDelayResponseFollowUpSent})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-received", types.YLeaf{"PeerDelayResponseFollowUpReceived", counters.PeerDelayResponseFollowUpReceived})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-dropped", types.YLeaf{"PeerDelayResponseFollowUpDropped", counters.PeerDelayResponseFollowUpDropped})
    counters.EntityData.Leafs.Append("signaling-sent", types.YLeaf{"SignalingSent", counters.SignalingSent})
    counters.EntityData.Leafs.Append("signaling-received", types.YLeaf{"SignalingReceived", counters.SignalingReceived})
    counters.EntityData.Leafs.Append("signaling-dropped", types.YLeaf{"SignalingDropped", counters.SignalingDropped})
    counters.EntityData.Leafs.Append("management-sent", types.YLeaf{"ManagementSent", counters.ManagementSent})
    counters.EntityData.Leafs.Append("management-received", types.YLeaf{"ManagementReceived", counters.ManagementReceived})
    counters.EntityData.Leafs.Append("management-dropped", types.YLeaf{"ManagementDropped", counters.ManagementDropped})
    counters.EntityData.Leafs.Append("other-packets-sent", types.YLeaf{"OtherPacketsSent", counters.OtherPacketsSent})
    counters.EntityData.Leafs.Append("other-packets-received", types.YLeaf{"OtherPacketsReceived", counters.OtherPacketsReceived})
    counters.EntityData.Leafs.Append("other-packets-dropped", types.YLeaf{"OtherPacketsDropped", counters.OtherPacketsDropped})
    counters.EntityData.Leafs.Append("total-packets-sent", types.YLeaf{"TotalPacketsSent", counters.TotalPacketsSent})
    counters.EntityData.Leafs.Append("total-packets-received", types.YLeaf{"TotalPacketsReceived", counters.TotalPacketsReceived})
    counters.EntityData.Leafs.Append("total-packets-dropped", types.YLeaf{"TotalPacketsDropped", counters.TotalPacketsDropped})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter
// Packet counters for each peer on this interface
type Ptp_InterfacePacketCounters_InterfacePacketCounter_PeerCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    peerCounter.EntityData.SegmentPath = "peer-counter" + types.AddNoKeyToken(peerCounter)
    peerCounter.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/interface-packet-counter/" + peerCounter.EntityData.SegmentPath
    peerCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerCounter.EntityData.Children = types.NewOrderedMap()
    peerCounter.EntityData.Children.Append("address", types.YChild{"Address", &peerCounter.Address})
    peerCounter.EntityData.Children.Append("counters", types.YChild{"Counters", &peerCounter.Counters})
    peerCounter.EntityData.Leafs = types.NewOrderedMap()

    peerCounter.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/interface-packet-counter/peer-counter/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/interface-packet-counter/peer-counter/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/interface-packet-counter/peer-counter/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    counters.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-packet-counters/interface-packet-counter/peer-counter/" + counters.EntityData.SegmentPath
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("announce-sent", types.YLeaf{"AnnounceSent", counters.AnnounceSent})
    counters.EntityData.Leafs.Append("announce-received", types.YLeaf{"AnnounceReceived", counters.AnnounceReceived})
    counters.EntityData.Leafs.Append("announce-dropped", types.YLeaf{"AnnounceDropped", counters.AnnounceDropped})
    counters.EntityData.Leafs.Append("sync-sent", types.YLeaf{"SyncSent", counters.SyncSent})
    counters.EntityData.Leafs.Append("sync-received", types.YLeaf{"SyncReceived", counters.SyncReceived})
    counters.EntityData.Leafs.Append("sync-dropped", types.YLeaf{"SyncDropped", counters.SyncDropped})
    counters.EntityData.Leafs.Append("follow-up-sent", types.YLeaf{"FollowUpSent", counters.FollowUpSent})
    counters.EntityData.Leafs.Append("follow-up-received", types.YLeaf{"FollowUpReceived", counters.FollowUpReceived})
    counters.EntityData.Leafs.Append("follow-up-dropped", types.YLeaf{"FollowUpDropped", counters.FollowUpDropped})
    counters.EntityData.Leafs.Append("delay-request-sent", types.YLeaf{"DelayRequestSent", counters.DelayRequestSent})
    counters.EntityData.Leafs.Append("delay-request-received", types.YLeaf{"DelayRequestReceived", counters.DelayRequestReceived})
    counters.EntityData.Leafs.Append("delay-request-dropped", types.YLeaf{"DelayRequestDropped", counters.DelayRequestDropped})
    counters.EntityData.Leafs.Append("delay-response-sent", types.YLeaf{"DelayResponseSent", counters.DelayResponseSent})
    counters.EntityData.Leafs.Append("delay-response-received", types.YLeaf{"DelayResponseReceived", counters.DelayResponseReceived})
    counters.EntityData.Leafs.Append("delay-response-dropped", types.YLeaf{"DelayResponseDropped", counters.DelayResponseDropped})
    counters.EntityData.Leafs.Append("peer-delay-request-sent", types.YLeaf{"PeerDelayRequestSent", counters.PeerDelayRequestSent})
    counters.EntityData.Leafs.Append("peer-delay-request-received", types.YLeaf{"PeerDelayRequestReceived", counters.PeerDelayRequestReceived})
    counters.EntityData.Leafs.Append("peer-delay-request-dropped", types.YLeaf{"PeerDelayRequestDropped", counters.PeerDelayRequestDropped})
    counters.EntityData.Leafs.Append("peer-delay-response-sent", types.YLeaf{"PeerDelayResponseSent", counters.PeerDelayResponseSent})
    counters.EntityData.Leafs.Append("peer-delay-response-received", types.YLeaf{"PeerDelayResponseReceived", counters.PeerDelayResponseReceived})
    counters.EntityData.Leafs.Append("peer-delay-response-dropped", types.YLeaf{"PeerDelayResponseDropped", counters.PeerDelayResponseDropped})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-sent", types.YLeaf{"PeerDelayResponseFollowUpSent", counters.PeerDelayResponseFollowUpSent})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-received", types.YLeaf{"PeerDelayResponseFollowUpReceived", counters.PeerDelayResponseFollowUpReceived})
    counters.EntityData.Leafs.Append("peer-delay-response-follow-up-dropped", types.YLeaf{"PeerDelayResponseFollowUpDropped", counters.PeerDelayResponseFollowUpDropped})
    counters.EntityData.Leafs.Append("signaling-sent", types.YLeaf{"SignalingSent", counters.SignalingSent})
    counters.EntityData.Leafs.Append("signaling-received", types.YLeaf{"SignalingReceived", counters.SignalingReceived})
    counters.EntityData.Leafs.Append("signaling-dropped", types.YLeaf{"SignalingDropped", counters.SignalingDropped})
    counters.EntityData.Leafs.Append("management-sent", types.YLeaf{"ManagementSent", counters.ManagementSent})
    counters.EntityData.Leafs.Append("management-received", types.YLeaf{"ManagementReceived", counters.ManagementReceived})
    counters.EntityData.Leafs.Append("management-dropped", types.YLeaf{"ManagementDropped", counters.ManagementDropped})
    counters.EntityData.Leafs.Append("other-packets-sent", types.YLeaf{"OtherPacketsSent", counters.OtherPacketsSent})
    counters.EntityData.Leafs.Append("other-packets-received", types.YLeaf{"OtherPacketsReceived", counters.OtherPacketsReceived})
    counters.EntityData.Leafs.Append("other-packets-dropped", types.YLeaf{"OtherPacketsDropped", counters.OtherPacketsDropped})
    counters.EntityData.Leafs.Append("total-packets-sent", types.YLeaf{"TotalPacketsSent", counters.TotalPacketsSent})
    counters.EntityData.Leafs.Append("total-packets-received", types.YLeaf{"TotalPacketsReceived", counters.TotalPacketsReceived})
    counters.EntityData.Leafs.Append("total-packets-dropped", types.YLeaf{"TotalPacketsDropped", counters.TotalPacketsDropped})

    counters.EntityData.YListKeys = []string {}

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
    advertisedClock.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + advertisedClock.EntityData.SegmentPath
    advertisedClock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertisedClock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertisedClock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertisedClock.EntityData.Children = types.NewOrderedMap()
    advertisedClock.EntityData.Children.Append("clock-properties", types.YChild{"ClockProperties", &advertisedClock.ClockProperties})
    advertisedClock.EntityData.Leafs = types.NewOrderedMap()
    advertisedClock.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", advertisedClock.Domain})
    advertisedClock.EntityData.Leafs.Append("time-source-configured", types.YLeaf{"TimeSourceConfigured", advertisedClock.TimeSourceConfigured})
    advertisedClock.EntityData.Leafs.Append("received-time-source", types.YLeaf{"ReceivedTimeSource", advertisedClock.ReceivedTimeSource})
    advertisedClock.EntityData.Leafs.Append("timescale-configured", types.YLeaf{"TimescaleConfigured", advertisedClock.TimescaleConfigured})
    advertisedClock.EntityData.Leafs.Append("received-timescale", types.YLeaf{"ReceivedTimescale", advertisedClock.ReceivedTimescale})

    advertisedClock.EntityData.YListKeys = []string {}

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
    clockProperties.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/advertised-clock/" + clockProperties.EntityData.SegmentPath
    clockProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockProperties.EntityData.Children = types.NewOrderedMap()
    clockProperties.EntityData.Children.Append("utc-offset", types.YChild{"UtcOffset", &clockProperties.UtcOffset})
    clockProperties.EntityData.Children.Append("receiver", types.YChild{"Receiver", &clockProperties.Receiver})
    clockProperties.EntityData.Children.Append("sender", types.YChild{"Sender", &clockProperties.Sender})
    clockProperties.EntityData.Leafs = types.NewOrderedMap()
    clockProperties.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", clockProperties.ClockId})
    clockProperties.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", clockProperties.Priority1})
    clockProperties.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", clockProperties.Priority2})
    clockProperties.EntityData.Leafs.Append("class", types.YLeaf{"Class", clockProperties.Class})
    clockProperties.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", clockProperties.Accuracy})
    clockProperties.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", clockProperties.OffsetLogVariance})
    clockProperties.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", clockProperties.StepsRemoved})
    clockProperties.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", clockProperties.TimeSource})
    clockProperties.EntityData.Leafs.Append("frequency-traceable", types.YLeaf{"FrequencyTraceable", clockProperties.FrequencyTraceable})
    clockProperties.EntityData.Leafs.Append("time-traceable", types.YLeaf{"TimeTraceable", clockProperties.TimeTraceable})
    clockProperties.EntityData.Leafs.Append("timescale", types.YLeaf{"Timescale", clockProperties.Timescale})
    clockProperties.EntityData.Leafs.Append("leap-seconds", types.YLeaf{"LeapSeconds", clockProperties.LeapSeconds})
    clockProperties.EntityData.Leafs.Append("local", types.YLeaf{"Local", clockProperties.Local})
    clockProperties.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", clockProperties.ConfiguredClockClass})
    clockProperties.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", clockProperties.ConfiguredPriority})

    clockProperties.EntityData.YListKeys = []string {}

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
    utcOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/advertised-clock/clock-properties/" + utcOffset.EntityData.SegmentPath
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = types.NewOrderedMap()
    utcOffset.EntityData.Leafs = types.NewOrderedMap()
    utcOffset.EntityData.Leafs.Append("current-offset", types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset})
    utcOffset.EntityData.Leafs.Append("offset-valid", types.YLeaf{"OffsetValid", utcOffset.OffsetValid})

    utcOffset.EntityData.YListKeys = []string {}

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
    receiver.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/advertised-clock/clock-properties/" + receiver.EntityData.SegmentPath
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", receiver.ClockId})
    receiver.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", receiver.PortNumber})

    receiver.EntityData.YListKeys = []string {}

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
    sender.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/advertised-clock/clock-properties/" + sender.EntityData.SegmentPath
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = types.NewOrderedMap()
    sender.EntityData.Leafs = types.NewOrderedMap()
    sender.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", sender.ClockId})
    sender.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", sender.PortNumber})

    sender.EntityData.YListKeys = []string {}

    return &(sender.EntityData)
}

// Ptp_Interfaces
// Table for interface operational data
type Ptp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface operational data. The type is slice of Ptp_Interfaces_Interface.
    Interface []*Ptp_Interfaces_Interface
}

func (interfaces *Ptp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ptp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + interfaces.EntityData.SegmentPath
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

// Ptp_Interfaces_Interface
// Interface operational data
type Ptp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Port state. The type is PtpBagPortState.
    PortState interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Line state. The type is ImStateEnum.
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

    // The interface supports unicast. The type is bool.
    SupportsUnicast interface{}

    // The interface supports operation in master mode. The type is bool.
    SupportsMaster interface{}

    // The interface supports one-step operation. The type is bool.
    SupportsOneStep interface{}

    // The interface supports two-step operation. The type is bool.
    SupportsTwoStep interface{}

    // The interface supports ethernet transport. The type is bool.
    SupportsEthernet interface{}

    // The interface supports multicast. The type is bool.
    SupportsMulticast interface{}

    // The interface supports IPv4 transport. The type is bool.
    SupportsIpv4 interface{}

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

    // Indicate whether profile interop is in use. The type is bool.
    ProfileInterop interface{}

    // The PTP domain that is being interoperated with. The type is interface{}
    // with range: 0..255.
    InteropDomain interface{}

    // Profile that is being interoperated with. The type is PtpBagProfile.
    InteropProfile interface{}

    // List of Ipv6 addresses, if IPv6 encapsulation is being used. If a source
    // address is configured, this is the only item in the list.
    Ipv6AddressArray Ptp_Interfaces_Interface_Ipv6AddressArray

    // List of IPv4 addresses, if IPv4 encapsulation is being used. The first
    // address is the primary address. If a source address is configured, this is
    // the only item in the list.
    Ipv4AddressArray Ptp_Interfaces_Interface_Ipv4AddressArray

    // MAC address, if Ethernet encapsulation is being used.
    MacAddress Ptp_Interfaces_Interface_MacAddress

    // Details of any ingress conversion.
    IngressConversion Ptp_Interfaces_Interface_IngressConversion

    // Details of any egress conversion.
    EgressConversion Ptp_Interfaces_Interface_EgressConversion

    // The interface's master table. The type is slice of
    // Ptp_Interfaces_Interface_MasterTable.
    MasterTable []*Ptp_Interfaces_Interface_MasterTable
}

func (self *Ptp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("ipv6-address-array", types.YChild{"Ipv6AddressArray", &self.Ipv6AddressArray})
    self.EntityData.Children.Append("ipv4-address-array", types.YChild{"Ipv4AddressArray", &self.Ipv4AddressArray})
    self.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &self.MacAddress})
    self.EntityData.Children.Append("ingress-conversion", types.YChild{"IngressConversion", &self.IngressConversion})
    self.EntityData.Children.Append("egress-conversion", types.YChild{"EgressConversion", &self.EgressConversion})
    self.EntityData.Children.Append("master-table", types.YChild{"MasterTable", nil})
    for i := range self.MasterTable {
        types.SetYListKey(self.MasterTable[i], i)
        self.EntityData.Children.Append(types.GetSegmentPath(self.MasterTable[i]), types.YChild{"MasterTable", self.MasterTable[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("port-state", types.YLeaf{"PortState", self.PortState})
    self.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", self.PortNumber})
    self.EntityData.Leafs.Append("line-state", types.YLeaf{"LineState", self.LineState})
    self.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", self.Encapsulation})
    self.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", self.Ipv6Address})
    self.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", self.Ipv4Address})
    self.EntityData.Leafs.Append("two-step", types.YLeaf{"TwoStep", self.TwoStep})
    self.EntityData.Leafs.Append("communication-model", types.YLeaf{"CommunicationModel", self.CommunicationModel})
    self.EntityData.Leafs.Append("log-sync-interval", types.YLeaf{"LogSyncInterval", self.LogSyncInterval})
    self.EntityData.Leafs.Append("log-announce-interval", types.YLeaf{"LogAnnounceInterval", self.LogAnnounceInterval})
    self.EntityData.Leafs.Append("announce-timeout", types.YLeaf{"AnnounceTimeout", self.AnnounceTimeout})
    self.EntityData.Leafs.Append("log-min-delay-request-interval", types.YLeaf{"LogMinDelayRequestInterval", self.LogMinDelayRequestInterval})
    self.EntityData.Leafs.Append("configured-port-state", types.YLeaf{"ConfiguredPortState", self.ConfiguredPortState})
    self.EntityData.Leafs.Append("supports-unicast", types.YLeaf{"SupportsUnicast", self.SupportsUnicast})
    self.EntityData.Leafs.Append("supports-master", types.YLeaf{"SupportsMaster", self.SupportsMaster})
    self.EntityData.Leafs.Append("supports-one-step", types.YLeaf{"SupportsOneStep", self.SupportsOneStep})
    self.EntityData.Leafs.Append("supports-two-step", types.YLeaf{"SupportsTwoStep", self.SupportsTwoStep})
    self.EntityData.Leafs.Append("supports-ethernet", types.YLeaf{"SupportsEthernet", self.SupportsEthernet})
    self.EntityData.Leafs.Append("supports-multicast", types.YLeaf{"SupportsMulticast", self.SupportsMulticast})
    self.EntityData.Leafs.Append("supports-ipv4", types.YLeaf{"SupportsIpv4", self.SupportsIpv4})
    self.EntityData.Leafs.Append("supports-ipv6", types.YLeaf{"SupportsIpv6", self.SupportsIpv6})
    self.EntityData.Leafs.Append("supports-slave", types.YLeaf{"SupportsSlave", self.SupportsSlave})
    self.EntityData.Leafs.Append("supports-source-ip", types.YLeaf{"SupportsSourceIp", self.SupportsSourceIp})
    self.EntityData.Leafs.Append("max-sync-rate", types.YLeaf{"MaxSyncRate", self.MaxSyncRate})
    self.EntityData.Leafs.Append("event-cos", types.YLeaf{"EventCos", self.EventCos})
    self.EntityData.Leafs.Append("general-cos", types.YLeaf{"GeneralCos", self.GeneralCos})
    self.EntityData.Leafs.Append("event-dscp", types.YLeaf{"EventDscp", self.EventDscp})
    self.EntityData.Leafs.Append("general-dscp", types.YLeaf{"GeneralDscp", self.GeneralDscp})
    self.EntityData.Leafs.Append("unicast-peers", types.YLeaf{"UnicastPeers", self.UnicastPeers})
    self.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", self.LocalPriority})
    self.EntityData.Leafs.Append("signal-fail", types.YLeaf{"SignalFail", self.SignalFail})
    self.EntityData.Leafs.Append("profile-interop", types.YLeaf{"ProfileInterop", self.ProfileInterop})
    self.EntityData.Leafs.Append("interop-domain", types.YLeaf{"InteropDomain", self.InteropDomain})
    self.EntityData.Leafs.Append("interop-profile", types.YLeaf{"InteropProfile", self.InteropProfile})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ptp_Interfaces_Interface_Ipv6AddressArray
// List of Ipv6 addresses, if IPv6 encapsulation is
// being used. If a source address is configured,
// this is the only item in the list
type Ptp_Interfaces_Interface_Ipv6AddressArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv6 addresses. The type is slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Addr []interface{}
}

func (ipv6AddressArray *Ptp_Interfaces_Interface_Ipv6AddressArray) GetEntityData() *types.CommonEntityData {
    ipv6AddressArray.EntityData.YFilter = ipv6AddressArray.YFilter
    ipv6AddressArray.EntityData.YangName = "ipv6-address-array"
    ipv6AddressArray.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressArray.EntityData.ParentYangName = "interface"
    ipv6AddressArray.EntityData.SegmentPath = "ipv6-address-array"
    ipv6AddressArray.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/" + ipv6AddressArray.EntityData.SegmentPath
    ipv6AddressArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressArray.EntityData.Children = types.NewOrderedMap()
    ipv6AddressArray.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressArray.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", ipv6AddressArray.Addr})

    ipv6AddressArray.EntityData.YListKeys = []string {}

    return &(ipv6AddressArray.EntityData)
}

// Ptp_Interfaces_Interface_Ipv4AddressArray
// List of IPv4 addresses, if IPv4 encapsulation is
// being used. The first address is the primary
// address. If a source address is configured, this
// is the only item in the list.
type Ptp_Interfaces_Interface_Ipv4AddressArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv4 addresses. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Addr []interface{}
}

func (ipv4AddressArray *Ptp_Interfaces_Interface_Ipv4AddressArray) GetEntityData() *types.CommonEntityData {
    ipv4AddressArray.EntityData.YFilter = ipv4AddressArray.YFilter
    ipv4AddressArray.EntityData.YangName = "ipv4-address-array"
    ipv4AddressArray.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressArray.EntityData.ParentYangName = "interface"
    ipv4AddressArray.EntityData.SegmentPath = "ipv4-address-array"
    ipv4AddressArray.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/" + ipv4AddressArray.EntityData.SegmentPath
    ipv4AddressArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressArray.EntityData.Children = types.NewOrderedMap()
    ipv4AddressArray.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressArray.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", ipv4AddressArray.Addr})

    ipv4AddressArray.EntityData.YListKeys = []string {}

    return &(ipv4AddressArray.EntityData)
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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

    return &(macAddress.EntityData)
}

// Ptp_Interfaces_Interface_IngressConversion
// Details of any ingress conversion
type Ptp_Interfaces_Interface_IngressConversion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Class Default. The type is interface{} with range: 0..255.
    ClassDefault interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Class Mapping. The type is slice of
    // Ptp_Interfaces_Interface_IngressConversion_ClassMapping.
    ClassMapping []*Ptp_Interfaces_Interface_IngressConversion_ClassMapping
}

func (ingressConversion *Ptp_Interfaces_Interface_IngressConversion) GetEntityData() *types.CommonEntityData {
    ingressConversion.EntityData.YFilter = ingressConversion.YFilter
    ingressConversion.EntityData.YangName = "ingress-conversion"
    ingressConversion.EntityData.BundleName = "cisco_ios_xr"
    ingressConversion.EntityData.ParentYangName = "interface"
    ingressConversion.EntityData.SegmentPath = "ingress-conversion"
    ingressConversion.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/" + ingressConversion.EntityData.SegmentPath
    ingressConversion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressConversion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressConversion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressConversion.EntityData.Children = types.NewOrderedMap()
    ingressConversion.EntityData.Children.Append("class-mapping", types.YChild{"ClassMapping", nil})
    for i := range ingressConversion.ClassMapping {
        types.SetYListKey(ingressConversion.ClassMapping[i], i)
        ingressConversion.EntityData.Children.Append(types.GetSegmentPath(ingressConversion.ClassMapping[i]), types.YChild{"ClassMapping", ingressConversion.ClassMapping[i]})
    }
    ingressConversion.EntityData.Leafs = types.NewOrderedMap()
    ingressConversion.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", ingressConversion.Priority1})
    ingressConversion.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", ingressConversion.Priority2})
    ingressConversion.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", ingressConversion.Accuracy})
    ingressConversion.EntityData.Leafs.Append("class-default", types.YLeaf{"ClassDefault", ingressConversion.ClassDefault})
    ingressConversion.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", ingressConversion.OffsetLogVariance})

    ingressConversion.EntityData.YListKeys = []string {}

    return &(ingressConversion.EntityData)
}

// Ptp_Interfaces_Interface_IngressConversion_ClassMapping
// Class Mapping
type Ptp_Interfaces_Interface_IngressConversion_ClassMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // From clock class. The type is interface{} with range: 0..255.
    FromClockClass interface{}

    // To clock class. The type is interface{} with range: 0..255.
    ToClockClass interface{}
}

func (classMapping *Ptp_Interfaces_Interface_IngressConversion_ClassMapping) GetEntityData() *types.CommonEntityData {
    classMapping.EntityData.YFilter = classMapping.YFilter
    classMapping.EntityData.YangName = "class-mapping"
    classMapping.EntityData.BundleName = "cisco_ios_xr"
    classMapping.EntityData.ParentYangName = "ingress-conversion"
    classMapping.EntityData.SegmentPath = "class-mapping" + types.AddNoKeyToken(classMapping)
    classMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/ingress-conversion/" + classMapping.EntityData.SegmentPath
    classMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapping.EntityData.Children = types.NewOrderedMap()
    classMapping.EntityData.Leafs = types.NewOrderedMap()
    classMapping.EntityData.Leafs.Append("from-clock-class", types.YLeaf{"FromClockClass", classMapping.FromClockClass})
    classMapping.EntityData.Leafs.Append("to-clock-class", types.YLeaf{"ToClockClass", classMapping.ToClockClass})

    classMapping.EntityData.YListKeys = []string {}

    return &(classMapping.EntityData)
}

// Ptp_Interfaces_Interface_EgressConversion
// Details of any egress conversion
type Ptp_Interfaces_Interface_EgressConversion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority 1. The type is interface{} with range: 0..255.
    Priority1 interface{}

    // Priority 2. The type is interface{} with range: 0..255.
    Priority2 interface{}

    // Accuracy. The type is interface{} with range: 0..255.
    Accuracy interface{}

    // Class Default. The type is interface{} with range: 0..255.
    ClassDefault interface{}

    // Offset log variance. The type is interface{} with range: 0..65535.
    OffsetLogVariance interface{}

    // Class Mapping. The type is slice of
    // Ptp_Interfaces_Interface_EgressConversion_ClassMapping.
    ClassMapping []*Ptp_Interfaces_Interface_EgressConversion_ClassMapping
}

func (egressConversion *Ptp_Interfaces_Interface_EgressConversion) GetEntityData() *types.CommonEntityData {
    egressConversion.EntityData.YFilter = egressConversion.YFilter
    egressConversion.EntityData.YangName = "egress-conversion"
    egressConversion.EntityData.BundleName = "cisco_ios_xr"
    egressConversion.EntityData.ParentYangName = "interface"
    egressConversion.EntityData.SegmentPath = "egress-conversion"
    egressConversion.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/" + egressConversion.EntityData.SegmentPath
    egressConversion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressConversion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressConversion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressConversion.EntityData.Children = types.NewOrderedMap()
    egressConversion.EntityData.Children.Append("class-mapping", types.YChild{"ClassMapping", nil})
    for i := range egressConversion.ClassMapping {
        types.SetYListKey(egressConversion.ClassMapping[i], i)
        egressConversion.EntityData.Children.Append(types.GetSegmentPath(egressConversion.ClassMapping[i]), types.YChild{"ClassMapping", egressConversion.ClassMapping[i]})
    }
    egressConversion.EntityData.Leafs = types.NewOrderedMap()
    egressConversion.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", egressConversion.Priority1})
    egressConversion.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", egressConversion.Priority2})
    egressConversion.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", egressConversion.Accuracy})
    egressConversion.EntityData.Leafs.Append("class-default", types.YLeaf{"ClassDefault", egressConversion.ClassDefault})
    egressConversion.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", egressConversion.OffsetLogVariance})

    egressConversion.EntityData.YListKeys = []string {}

    return &(egressConversion.EntityData)
}

// Ptp_Interfaces_Interface_EgressConversion_ClassMapping
// Class Mapping
type Ptp_Interfaces_Interface_EgressConversion_ClassMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // From clock class. The type is interface{} with range: 0..255.
    FromClockClass interface{}

    // To clock class. The type is interface{} with range: 0..255.
    ToClockClass interface{}
}

func (classMapping *Ptp_Interfaces_Interface_EgressConversion_ClassMapping) GetEntityData() *types.CommonEntityData {
    classMapping.EntityData.YFilter = classMapping.YFilter
    classMapping.EntityData.YangName = "class-mapping"
    classMapping.EntityData.BundleName = "cisco_ios_xr"
    classMapping.EntityData.ParentYangName = "egress-conversion"
    classMapping.EntityData.SegmentPath = "class-mapping" + types.AddNoKeyToken(classMapping)
    classMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/egress-conversion/" + classMapping.EntityData.SegmentPath
    classMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMapping.EntityData.Children = types.NewOrderedMap()
    classMapping.EntityData.Leafs = types.NewOrderedMap()
    classMapping.EntityData.Leafs.Append("from-clock-class", types.YLeaf{"FromClockClass", classMapping.FromClockClass})
    classMapping.EntityData.Leafs.Append("to-clock-class", types.YLeaf{"ToClockClass", classMapping.ToClockClass})

    classMapping.EntityData.YListKeys = []string {}

    return &(classMapping.EntityData)
}

// Ptp_Interfaces_Interface_MasterTable
// The interface's master table
type Ptp_Interfaces_Interface_MasterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    masterTable.EntityData.SegmentPath = "master-table" + types.AddNoKeyToken(masterTable)
    masterTable.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/" + masterTable.EntityData.SegmentPath
    masterTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    masterTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    masterTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    masterTable.EntityData.Children = types.NewOrderedMap()
    masterTable.EntityData.Children.Append("address", types.YChild{"Address", &masterTable.Address})
    masterTable.EntityData.Leafs = types.NewOrderedMap()
    masterTable.EntityData.Leafs.Append("communication-model", types.YLeaf{"CommunicationModel", masterTable.CommunicationModel})
    masterTable.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", masterTable.Priority})
    masterTable.EntityData.Leafs.Append("known", types.YLeaf{"Known", masterTable.Known})
    masterTable.EntityData.Leafs.Append("qualified", types.YLeaf{"Qualified", masterTable.Qualified})
    masterTable.EntityData.Leafs.Append("is-grandmaster", types.YLeaf{"IsGrandmaster", masterTable.IsGrandmaster})
    masterTable.EntityData.Leafs.Append("ptsf-loss-announce", types.YLeaf{"PtsfLossAnnounce", masterTable.PtsfLossAnnounce})
    masterTable.EntityData.Leafs.Append("ptsf-loss-sync", types.YLeaf{"PtsfLossSync", masterTable.PtsfLossSync})
    masterTable.EntityData.Leafs.Append("is-nonnegotiated", types.YLeaf{"IsNonnegotiated", masterTable.IsNonnegotiated})

    masterTable.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/master-table/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/master-table/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interfaces/interface/master-table/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    dataset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + dataset.EntityData.SegmentPath
    dataset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataset.EntityData.Children = types.NewOrderedMap()
    dataset.EntityData.Children.Append("default-ds", types.YChild{"DefaultDs", &dataset.DefaultDs})
    dataset.EntityData.Children.Append("current-ds", types.YChild{"CurrentDs", &dataset.CurrentDs})
    dataset.EntityData.Children.Append("parent-ds", types.YChild{"ParentDs", &dataset.ParentDs})
    dataset.EntityData.Children.Append("port-dses", types.YChild{"PortDses", &dataset.PortDses})
    dataset.EntityData.Children.Append("time-properties-ds", types.YChild{"TimePropertiesDs", &dataset.TimePropertiesDs})
    dataset.EntityData.Leafs = types.NewOrderedMap()

    dataset.EntityData.YListKeys = []string {}

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
    defaultDs.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/dataset/" + defaultDs.EntityData.SegmentPath
    defaultDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultDs.EntityData.Children = types.NewOrderedMap()
    defaultDs.EntityData.Leafs = types.NewOrderedMap()
    defaultDs.EntityData.Leafs.Append("two-step-flag", types.YLeaf{"TwoStepFlag", defaultDs.TwoStepFlag})
    defaultDs.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", defaultDs.ClockId})
    defaultDs.EntityData.Leafs.Append("number-ports", types.YLeaf{"NumberPorts", defaultDs.NumberPorts})
    defaultDs.EntityData.Leafs.Append("clock-class", types.YLeaf{"ClockClass", defaultDs.ClockClass})
    defaultDs.EntityData.Leafs.Append("clock-accuracy", types.YLeaf{"ClockAccuracy", defaultDs.ClockAccuracy})
    defaultDs.EntityData.Leafs.Append("oslv", types.YLeaf{"Oslv", defaultDs.Oslv})
    defaultDs.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", defaultDs.Priority1})
    defaultDs.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", defaultDs.Priority2})
    defaultDs.EntityData.Leafs.Append("domain-number", types.YLeaf{"DomainNumber", defaultDs.DomainNumber})
    defaultDs.EntityData.Leafs.Append("slave-only", types.YLeaf{"SlaveOnly", defaultDs.SlaveOnly})
    defaultDs.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", defaultDs.LocalPriority})
    defaultDs.EntityData.Leafs.Append("signal-fail", types.YLeaf{"SignalFail", defaultDs.SignalFail})

    defaultDs.EntityData.YListKeys = []string {}

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
    currentDs.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/dataset/" + currentDs.EntityData.SegmentPath
    currentDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentDs.EntityData.Children = types.NewOrderedMap()
    currentDs.EntityData.Leafs = types.NewOrderedMap()
    currentDs.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", currentDs.StepsRemoved})
    currentDs.EntityData.Leafs.Append("offset-from-master", types.YLeaf{"OffsetFromMaster", currentDs.OffsetFromMaster})
    currentDs.EntityData.Leafs.Append("mean-path-delay", types.YLeaf{"MeanPathDelay", currentDs.MeanPathDelay})

    currentDs.EntityData.YListKeys = []string {}

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
    parentDs.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/dataset/" + parentDs.EntityData.SegmentPath
    parentDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentDs.EntityData.Children = types.NewOrderedMap()
    parentDs.EntityData.Leafs = types.NewOrderedMap()
    parentDs.EntityData.Leafs.Append("parent-clock-id", types.YLeaf{"ParentClockId", parentDs.ParentClockId})
    parentDs.EntityData.Leafs.Append("parent-port-number", types.YLeaf{"ParentPortNumber", parentDs.ParentPortNumber})
    parentDs.EntityData.Leafs.Append("parent-stats", types.YLeaf{"ParentStats", parentDs.ParentStats})
    parentDs.EntityData.Leafs.Append("observed-parent-oslv", types.YLeaf{"ObservedParentOslv", parentDs.ObservedParentOslv})
    parentDs.EntityData.Leafs.Append("observed-parent-clock-phase-change-rate", types.YLeaf{"ObservedParentClockPhaseChangeRate", parentDs.ObservedParentClockPhaseChangeRate})
    parentDs.EntityData.Leafs.Append("gm-clock-id", types.YLeaf{"GmClockId", parentDs.GmClockId})
    parentDs.EntityData.Leafs.Append("gm-clock-class", types.YLeaf{"GmClockClass", parentDs.GmClockClass})
    parentDs.EntityData.Leafs.Append("gm-clock-accuracy", types.YLeaf{"GmClockAccuracy", parentDs.GmClockAccuracy})
    parentDs.EntityData.Leafs.Append("gmoslv", types.YLeaf{"Gmoslv", parentDs.Gmoslv})
    parentDs.EntityData.Leafs.Append("gm-priority1", types.YLeaf{"GmPriority1", parentDs.GmPriority1})
    parentDs.EntityData.Leafs.Append("gm-priority2", types.YLeaf{"GmPriority2", parentDs.GmPriority2})

    parentDs.EntityData.YListKeys = []string {}

    return &(parentDs.EntityData)
}

// Ptp_Dataset_PortDses
// Table for portDS information
type Ptp_Dataset_PortDses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortDS information. The type is slice of Ptp_Dataset_PortDses_PortDs.
    PortDs []*Ptp_Dataset_PortDses_PortDs
}

func (portDses *Ptp_Dataset_PortDses) GetEntityData() *types.CommonEntityData {
    portDses.EntityData.YFilter = portDses.YFilter
    portDses.EntityData.YangName = "port-dses"
    portDses.EntityData.BundleName = "cisco_ios_xr"
    portDses.EntityData.ParentYangName = "dataset"
    portDses.EntityData.SegmentPath = "port-dses"
    portDses.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/dataset/" + portDses.EntityData.SegmentPath
    portDses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portDses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portDses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portDses.EntityData.Children = types.NewOrderedMap()
    portDses.EntityData.Children.Append("port-ds", types.YChild{"PortDs", nil})
    for i := range portDses.PortDs {
        portDses.EntityData.Children.Append(types.GetSegmentPath(portDses.PortDs[i]), types.YChild{"PortDs", portDses.PortDs[i]})
    }
    portDses.EntityData.Leafs = types.NewOrderedMap()

    portDses.EntityData.YListKeys = []string {}

    return &(portDses.EntityData)
}

// Ptp_Dataset_PortDses_PortDs
// PortDS information
type Ptp_Dataset_PortDses_PortDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
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
    portDs.EntityData.SegmentPath = "port-ds" + types.AddKeyToken(portDs.InterfaceName, "interface-name")
    portDs.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/dataset/port-dses/" + portDs.EntityData.SegmentPath
    portDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portDs.EntityData.Children = types.NewOrderedMap()
    portDs.EntityData.Leafs = types.NewOrderedMap()
    portDs.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", portDs.InterfaceName})
    portDs.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", portDs.ClockId})
    portDs.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", portDs.PortNumber})
    portDs.EntityData.Leafs.Append("port-state", types.YLeaf{"PortState", portDs.PortState})
    portDs.EntityData.Leafs.Append("log-min-delay-req-interval", types.YLeaf{"LogMinDelayReqInterval", portDs.LogMinDelayReqInterval})
    portDs.EntityData.Leafs.Append("peer-mean-path-delay", types.YLeaf{"PeerMeanPathDelay", portDs.PeerMeanPathDelay})
    portDs.EntityData.Leafs.Append("log-announce-interval", types.YLeaf{"LogAnnounceInterval", portDs.LogAnnounceInterval})
    portDs.EntityData.Leafs.Append("annoucne-receipt-timeout", types.YLeaf{"AnnoucneReceiptTimeout", portDs.AnnoucneReceiptTimeout})
    portDs.EntityData.Leafs.Append("log-sync-interval", types.YLeaf{"LogSyncInterval", portDs.LogSyncInterval})
    portDs.EntityData.Leafs.Append("delay-mechanism", types.YLeaf{"DelayMechanism", portDs.DelayMechanism})
    portDs.EntityData.Leafs.Append("log-min-p-delay-req-interval", types.YLeaf{"LogMinPDelayReqInterval", portDs.LogMinPDelayReqInterval})
    portDs.EntityData.Leafs.Append("version-number", types.YLeaf{"VersionNumber", portDs.VersionNumber})
    portDs.EntityData.Leafs.Append("local-priority", types.YLeaf{"LocalPriority", portDs.LocalPriority})
    portDs.EntityData.Leafs.Append("master-only", types.YLeaf{"MasterOnly", portDs.MasterOnly})
    portDs.EntityData.Leafs.Append("signal-fail", types.YLeaf{"SignalFail", portDs.SignalFail})

    portDs.EntityData.YListKeys = []string {"InterfaceName"}

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
    timePropertiesDs.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/dataset/" + timePropertiesDs.EntityData.SegmentPath
    timePropertiesDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timePropertiesDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timePropertiesDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timePropertiesDs.EntityData.Children = types.NewOrderedMap()
    timePropertiesDs.EntityData.Leafs = types.NewOrderedMap()
    timePropertiesDs.EntityData.Leafs.Append("current-utc-offset", types.YLeaf{"CurrentUtcOffset", timePropertiesDs.CurrentUtcOffset})
    timePropertiesDs.EntityData.Leafs.Append("current-utc-offset-valid", types.YLeaf{"CurrentUtcOffsetValid", timePropertiesDs.CurrentUtcOffsetValid})
    timePropertiesDs.EntityData.Leafs.Append("leap59", types.YLeaf{"Leap59", timePropertiesDs.Leap59})
    timePropertiesDs.EntityData.Leafs.Append("leap61", types.YLeaf{"Leap61", timePropertiesDs.Leap61})
    timePropertiesDs.EntityData.Leafs.Append("time-traceable", types.YLeaf{"TimeTraceable", timePropertiesDs.TimeTraceable})
    timePropertiesDs.EntityData.Leafs.Append("frequency-traceable", types.YLeaf{"FrequencyTraceable", timePropertiesDs.FrequencyTraceable})
    timePropertiesDs.EntityData.Leafs.Append("ptp-timescale", types.YLeaf{"PtpTimescale", timePropertiesDs.PtpTimescale})
    timePropertiesDs.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", timePropertiesDs.TimeSource})

    timePropertiesDs.EntityData.YListKeys = []string {}

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

    // The configured priority2 value of the virtual port. The type is interface{}
    // with range: 0..255.
    VirtualPortPriority2 interface{}

    // The configured clock class of the virtual port. The type is interface{}
    // with range: 0..255.
    VirtualPortClockClass interface{}

    // The configured clock accuracy of the virtual port. The type is interface{}
    // with range: 0..255.
    VirtualPortClockAccuracy interface{}

    // The configured oslv of the virtual port. The type is interface{} with
    // range: 0..65535.
    VirtualPortOslv interface{}

    // Configuration Errors.
    ConfigurationErrors Ptp_GlobalConfigurationError_ConfigurationErrors
}

func (globalConfigurationError *Ptp_GlobalConfigurationError) GetEntityData() *types.CommonEntityData {
    globalConfigurationError.EntityData.YFilter = globalConfigurationError.YFilter
    globalConfigurationError.EntityData.YangName = "global-configuration-error"
    globalConfigurationError.EntityData.BundleName = "cisco_ios_xr"
    globalConfigurationError.EntityData.ParentYangName = "ptp"
    globalConfigurationError.EntityData.SegmentPath = "global-configuration-error"
    globalConfigurationError.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + globalConfigurationError.EntityData.SegmentPath
    globalConfigurationError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalConfigurationError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalConfigurationError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalConfigurationError.EntityData.Children = types.NewOrderedMap()
    globalConfigurationError.EntityData.Children.Append("configuration-errors", types.YChild{"ConfigurationErrors", &globalConfigurationError.ConfigurationErrors})
    globalConfigurationError.EntityData.Leafs = types.NewOrderedMap()
    globalConfigurationError.EntityData.Leafs.Append("clock-profile", types.YLeaf{"ClockProfile", globalConfigurationError.ClockProfile})
    globalConfigurationError.EntityData.Leafs.Append("clock-profile-set", types.YLeaf{"ClockProfileSet", globalConfigurationError.ClockProfileSet})
    globalConfigurationError.EntityData.Leafs.Append("telecom-clock-type", types.YLeaf{"TelecomClockType", globalConfigurationError.TelecomClockType})
    globalConfigurationError.EntityData.Leafs.Append("domain-number", types.YLeaf{"DomainNumber", globalConfigurationError.DomainNumber})
    globalConfigurationError.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", globalConfigurationError.Priority2})
    globalConfigurationError.EntityData.Leafs.Append("virtual-port-priority2", types.YLeaf{"VirtualPortPriority2", globalConfigurationError.VirtualPortPriority2})
    globalConfigurationError.EntityData.Leafs.Append("virtual-port-clock-class", types.YLeaf{"VirtualPortClockClass", globalConfigurationError.VirtualPortClockClass})
    globalConfigurationError.EntityData.Leafs.Append("virtual-port-clock-accuracy", types.YLeaf{"VirtualPortClockAccuracy", globalConfigurationError.VirtualPortClockAccuracy})
    globalConfigurationError.EntityData.Leafs.Append("virtual-port-oslv", types.YLeaf{"VirtualPortOslv", globalConfigurationError.VirtualPortOslv})

    globalConfigurationError.EntityData.YListKeys = []string {}

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

    // Physical Layer Frequency configuration is not compatible with G.8265.1
    // profile. The type is bool.
    PhysicalLayerFrequency interface{}

    // Virtual Port configuration is not compatible with default profile. The type
    // is bool.
    ProfileVirtualPort interface{}

    // Virtual Port priority1 configuration is not compatible with configured
    // profile. The type is bool.
    VirtualPortPriority1Config interface{}

    // Virtual Port priority2 value is not compatible with configured profile. The
    // type is bool.
    VirtualPortPriority2Value interface{}

    // Virtual port clock class is not compatible with configured profile. The
    // type is bool.
    VirtualPortProfileClockClass interface{}

    // Virtual port clock accuracy is not compatible with configured profile. The
    // type is bool.
    VirtualPortClockAccuracy interface{}

    // Virtual port OSLV is not compatible with configured profile. The type is
    // bool.
    VirtualPortOslv interface{}

    // Virtual port local priority configuration is not compatible with configured
    // profile. The type is bool.
    VirtualPortLocalPriority interface{}
}

func (configurationErrors *Ptp_GlobalConfigurationError_ConfigurationErrors) GetEntityData() *types.CommonEntityData {
    configurationErrors.EntityData.YFilter = configurationErrors.YFilter
    configurationErrors.EntityData.YangName = "configuration-errors"
    configurationErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationErrors.EntityData.ParentYangName = "global-configuration-error"
    configurationErrors.EntityData.SegmentPath = "configuration-errors"
    configurationErrors.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/global-configuration-error/" + configurationErrors.EntityData.SegmentPath
    configurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationErrors.EntityData.Children = types.NewOrderedMap()
    configurationErrors.EntityData.Leafs = types.NewOrderedMap()
    configurationErrors.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", configurationErrors.Domain})
    configurationErrors.EntityData.Leafs.Append("profile-priority1-config", types.YLeaf{"ProfilePriority1Config", configurationErrors.ProfilePriority1Config})
    configurationErrors.EntityData.Leafs.Append("profile-priority2-value", types.YLeaf{"ProfilePriority2Value", configurationErrors.ProfilePriority2Value})
    configurationErrors.EntityData.Leafs.Append("utc-offset-change", types.YLeaf{"UtcOffsetChange", configurationErrors.UtcOffsetChange})
    configurationErrors.EntityData.Leafs.Append("physical-layer-frequency", types.YLeaf{"PhysicalLayerFrequency", configurationErrors.PhysicalLayerFrequency})
    configurationErrors.EntityData.Leafs.Append("profile-virtual-port", types.YLeaf{"ProfileVirtualPort", configurationErrors.ProfileVirtualPort})
    configurationErrors.EntityData.Leafs.Append("virtual-port-priority1-config", types.YLeaf{"VirtualPortPriority1Config", configurationErrors.VirtualPortPriority1Config})
    configurationErrors.EntityData.Leafs.Append("virtual-port-priority2-value", types.YLeaf{"VirtualPortPriority2Value", configurationErrors.VirtualPortPriority2Value})
    configurationErrors.EntityData.Leafs.Append("virtual-port-profile-clock-class", types.YLeaf{"VirtualPortProfileClockClass", configurationErrors.VirtualPortProfileClockClass})
    configurationErrors.EntityData.Leafs.Append("virtual-port-clock-accuracy", types.YLeaf{"VirtualPortClockAccuracy", configurationErrors.VirtualPortClockAccuracy})
    configurationErrors.EntityData.Leafs.Append("virtual-port-oslv", types.YLeaf{"VirtualPortOslv", configurationErrors.VirtualPortOslv})
    configurationErrors.EntityData.Leafs.Append("virtual-port-local-priority", types.YLeaf{"VirtualPortLocalPriority", configurationErrors.VirtualPortLocalPriority})

    configurationErrors.EntityData.YListKeys = []string {}

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
    grandmaster.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + grandmaster.EntityData.SegmentPath
    grandmaster.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grandmaster.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grandmaster.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grandmaster.EntityData.Children = types.NewOrderedMap()
    grandmaster.EntityData.Children.Append("clock-properties", types.YChild{"ClockProperties", &grandmaster.ClockProperties})
    grandmaster.EntityData.Children.Append("address", types.YChild{"Address", &grandmaster.Address})
    grandmaster.EntityData.Leafs = types.NewOrderedMap()
    grandmaster.EntityData.Leafs.Append("used-for-time", types.YLeaf{"UsedForTime", grandmaster.UsedForTime})
    grandmaster.EntityData.Leafs.Append("used-for-frequency", types.YLeaf{"UsedForFrequency", grandmaster.UsedForFrequency})
    grandmaster.EntityData.Leafs.Append("known-for-time", types.YLeaf{"KnownForTime", grandmaster.KnownForTime})
    grandmaster.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", grandmaster.Domain})

    grandmaster.EntityData.YListKeys = []string {}

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
    clockProperties.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/" + clockProperties.EntityData.SegmentPath
    clockProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockProperties.EntityData.Children = types.NewOrderedMap()
    clockProperties.EntityData.Children.Append("utc-offset", types.YChild{"UtcOffset", &clockProperties.UtcOffset})
    clockProperties.EntityData.Children.Append("receiver", types.YChild{"Receiver", &clockProperties.Receiver})
    clockProperties.EntityData.Children.Append("sender", types.YChild{"Sender", &clockProperties.Sender})
    clockProperties.EntityData.Leafs = types.NewOrderedMap()
    clockProperties.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", clockProperties.ClockId})
    clockProperties.EntityData.Leafs.Append("priority1", types.YLeaf{"Priority1", clockProperties.Priority1})
    clockProperties.EntityData.Leafs.Append("priority2", types.YLeaf{"Priority2", clockProperties.Priority2})
    clockProperties.EntityData.Leafs.Append("class", types.YLeaf{"Class", clockProperties.Class})
    clockProperties.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", clockProperties.Accuracy})
    clockProperties.EntityData.Leafs.Append("offset-log-variance", types.YLeaf{"OffsetLogVariance", clockProperties.OffsetLogVariance})
    clockProperties.EntityData.Leafs.Append("steps-removed", types.YLeaf{"StepsRemoved", clockProperties.StepsRemoved})
    clockProperties.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", clockProperties.TimeSource})
    clockProperties.EntityData.Leafs.Append("frequency-traceable", types.YLeaf{"FrequencyTraceable", clockProperties.FrequencyTraceable})
    clockProperties.EntityData.Leafs.Append("time-traceable", types.YLeaf{"TimeTraceable", clockProperties.TimeTraceable})
    clockProperties.EntityData.Leafs.Append("timescale", types.YLeaf{"Timescale", clockProperties.Timescale})
    clockProperties.EntityData.Leafs.Append("leap-seconds", types.YLeaf{"LeapSeconds", clockProperties.LeapSeconds})
    clockProperties.EntityData.Leafs.Append("local", types.YLeaf{"Local", clockProperties.Local})
    clockProperties.EntityData.Leafs.Append("configured-clock-class", types.YLeaf{"ConfiguredClockClass", clockProperties.ConfiguredClockClass})
    clockProperties.EntityData.Leafs.Append("configured-priority", types.YLeaf{"ConfiguredPriority", clockProperties.ConfiguredPriority})

    clockProperties.EntityData.YListKeys = []string {}

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
    utcOffset.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/clock-properties/" + utcOffset.EntityData.SegmentPath
    utcOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffset.EntityData.Children = types.NewOrderedMap()
    utcOffset.EntityData.Leafs = types.NewOrderedMap()
    utcOffset.EntityData.Leafs.Append("current-offset", types.YLeaf{"CurrentOffset", utcOffset.CurrentOffset})
    utcOffset.EntityData.Leafs.Append("offset-valid", types.YLeaf{"OffsetValid", utcOffset.OffsetValid})

    utcOffset.EntityData.YListKeys = []string {}

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
    receiver.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/clock-properties/" + receiver.EntityData.SegmentPath
    receiver.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiver.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiver.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiver.EntityData.Children = types.NewOrderedMap()
    receiver.EntityData.Leafs = types.NewOrderedMap()
    receiver.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", receiver.ClockId})
    receiver.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", receiver.PortNumber})

    receiver.EntityData.YListKeys = []string {}

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
    sender.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/clock-properties/" + sender.EntityData.SegmentPath
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = types.NewOrderedMap()
    sender.EntityData.Leafs = types.NewOrderedMap()
    sender.EntityData.Leafs.Append("clock-id", types.YLeaf{"ClockId", sender.ClockId})
    sender.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", sender.PortNumber})

    sender.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/grandmaster/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    InterfaceUnicastPeer []*Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer
}

func (interfaceUnicastPeers *Ptp_InterfaceUnicastPeers) GetEntityData() *types.CommonEntityData {
    interfaceUnicastPeers.EntityData.YFilter = interfaceUnicastPeers.YFilter
    interfaceUnicastPeers.EntityData.YangName = "interface-unicast-peers"
    interfaceUnicastPeers.EntityData.BundleName = "cisco_ios_xr"
    interfaceUnicastPeers.EntityData.ParentYangName = "ptp"
    interfaceUnicastPeers.EntityData.SegmentPath = "interface-unicast-peers"
    interfaceUnicastPeers.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + interfaceUnicastPeers.EntityData.SegmentPath
    interfaceUnicastPeers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceUnicastPeers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceUnicastPeers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceUnicastPeers.EntityData.Children = types.NewOrderedMap()
    interfaceUnicastPeers.EntityData.Children.Append("interface-unicast-peer", types.YChild{"InterfaceUnicastPeer", nil})
    for i := range interfaceUnicastPeers.InterfaceUnicastPeer {
        interfaceUnicastPeers.EntityData.Children.Append(types.GetSegmentPath(interfaceUnicastPeers.InterfaceUnicastPeer[i]), types.YChild{"InterfaceUnicastPeer", interfaceUnicastPeers.InterfaceUnicastPeer[i]})
    }
    interfaceUnicastPeers.EntityData.Leafs = types.NewOrderedMap()

    interfaceUnicastPeers.EntityData.YListKeys = []string {}

    return &(interfaceUnicastPeers.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer
// Interface unicast peers operational data
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Port number. The type is interface{} with range: 0..65535.
    PortNumber interface{}

    // Unicast Peers. The type is slice of
    // Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers.
    Peers []*Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers
}

func (interfaceUnicastPeer *Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer) GetEntityData() *types.CommonEntityData {
    interfaceUnicastPeer.EntityData.YFilter = interfaceUnicastPeer.YFilter
    interfaceUnicastPeer.EntityData.YangName = "interface-unicast-peer"
    interfaceUnicastPeer.EntityData.BundleName = "cisco_ios_xr"
    interfaceUnicastPeer.EntityData.ParentYangName = "interface-unicast-peers"
    interfaceUnicastPeer.EntityData.SegmentPath = "interface-unicast-peer" + types.AddKeyToken(interfaceUnicastPeer.InterfaceName, "interface-name")
    interfaceUnicastPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/" + interfaceUnicastPeer.EntityData.SegmentPath
    interfaceUnicastPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceUnicastPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceUnicastPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceUnicastPeer.EntityData.Children = types.NewOrderedMap()
    interfaceUnicastPeer.EntityData.Children.Append("peers", types.YChild{"Peers", nil})
    for i := range interfaceUnicastPeer.Peers {
        types.SetYListKey(interfaceUnicastPeer.Peers[i], i)
        interfaceUnicastPeer.EntityData.Children.Append(types.GetSegmentPath(interfaceUnicastPeer.Peers[i]), types.YChild{"Peers", interfaceUnicastPeer.Peers[i]})
    }
    interfaceUnicastPeer.EntityData.Leafs = types.NewOrderedMap()
    interfaceUnicastPeer.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceUnicastPeer.InterfaceName})
    interfaceUnicastPeer.EntityData.Leafs.Append("name", types.YLeaf{"Name", interfaceUnicastPeer.Name})
    interfaceUnicastPeer.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", interfaceUnicastPeer.PortNumber})

    interfaceUnicastPeer.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceUnicastPeer.EntityData)
}

// Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers
// Unicast Peers
type Ptp_InterfaceUnicastPeers_InterfaceUnicastPeer_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    peers.EntityData.SegmentPath = "peers" + types.AddNoKeyToken(peers)
    peers.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/" + peers.EntityData.SegmentPath
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("address", types.YChild{"Address", &peers.Address})
    peers.EntityData.Children.Append("announce-grant", types.YChild{"AnnounceGrant", &peers.AnnounceGrant})
    peers.EntityData.Children.Append("sync-grant", types.YChild{"SyncGrant", &peers.SyncGrant})
    peers.EntityData.Children.Append("delay-response-grant", types.YChild{"DelayResponseGrant", &peers.DelayResponseGrant})
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

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
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/peers/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("mac-address", types.YChild{"MacAddress", &address.MacAddress})
    address.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", &address.Ipv6Address})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", address.Encapsulation})
    address.EntityData.Leafs.Append("address-unknown", types.YLeaf{"AddressUnknown", address.AddressUnknown})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})

    address.EntityData.YListKeys = []string {}

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
    macAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/peers/address/" + macAddress.EntityData.SegmentPath
    macAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAddress.EntityData.Children = types.NewOrderedMap()
    macAddress.EntityData.Leafs = types.NewOrderedMap()
    macAddress.EntityData.Leafs.Append("macaddr", types.YLeaf{"Macaddr", macAddress.Macaddr})

    macAddress.EntityData.YListKeys = []string {}

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
    ipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/peers/address/" + ipv6Address.EntityData.SegmentPath
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6Address.Ipv6Address})

    ipv6Address.EntityData.YListKeys = []string {}

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
    announceGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/peers/" + announceGrant.EntityData.SegmentPath
    announceGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announceGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announceGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announceGrant.EntityData.Children = types.NewOrderedMap()
    announceGrant.EntityData.Leafs = types.NewOrderedMap()
    announceGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", announceGrant.LogGrantInterval})
    announceGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", announceGrant.GrantDuration})

    announceGrant.EntityData.YListKeys = []string {}

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
    syncGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/peers/" + syncGrant.EntityData.SegmentPath
    syncGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncGrant.EntityData.Children = types.NewOrderedMap()
    syncGrant.EntityData.Leafs = types.NewOrderedMap()
    syncGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", syncGrant.LogGrantInterval})
    syncGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", syncGrant.GrantDuration})

    syncGrant.EntityData.YListKeys = []string {}

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
    delayResponseGrant.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/interface-unicast-peers/interface-unicast-peer/peers/" + delayResponseGrant.EntityData.SegmentPath
    delayResponseGrant.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayResponseGrant.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayResponseGrant.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayResponseGrant.EntityData.Children = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs = types.NewOrderedMap()
    delayResponseGrant.EntityData.Leafs.Append("log-grant-interval", types.YLeaf{"LogGrantInterval", delayResponseGrant.LogGrantInterval})
    delayResponseGrant.EntityData.Leafs.Append("grant-duration", types.YLeaf{"GrantDuration", delayResponseGrant.GrantDuration})

    delayResponseGrant.EntityData.YListKeys = []string {}

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
    ConfiguredLeapSecond []*Ptp_UtcOffsetInfo_ConfiguredLeapSecond
}

func (utcOffsetInfo *Ptp_UtcOffsetInfo) GetEntityData() *types.CommonEntityData {
    utcOffsetInfo.EntityData.YFilter = utcOffsetInfo.YFilter
    utcOffsetInfo.EntityData.YangName = "utc-offset-info"
    utcOffsetInfo.EntityData.BundleName = "cisco_ios_xr"
    utcOffsetInfo.EntityData.ParentYangName = "ptp"
    utcOffsetInfo.EntityData.SegmentPath = "utc-offset-info"
    utcOffsetInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/" + utcOffsetInfo.EntityData.SegmentPath
    utcOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utcOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utcOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utcOffsetInfo.EntityData.Children = types.NewOrderedMap()
    utcOffsetInfo.EntityData.Children.Append("current-offset-info", types.YChild{"CurrentOffsetInfo", &utcOffsetInfo.CurrentOffsetInfo})
    utcOffsetInfo.EntityData.Children.Append("current-gm-offset-info", types.YChild{"CurrentGmOffsetInfo", &utcOffsetInfo.CurrentGmOffsetInfo})
    utcOffsetInfo.EntityData.Children.Append("configured-offset-info", types.YChild{"ConfiguredOffsetInfo", &utcOffsetInfo.ConfiguredOffsetInfo})
    utcOffsetInfo.EntityData.Children.Append("previous-gm-offset-info", types.YChild{"PreviousGmOffsetInfo", &utcOffsetInfo.PreviousGmOffsetInfo})
    utcOffsetInfo.EntityData.Children.Append("hardware-offset-info", types.YChild{"HardwareOffsetInfo", &utcOffsetInfo.HardwareOffsetInfo})
    utcOffsetInfo.EntityData.Children.Append("gm-leap-second", types.YChild{"GmLeapSecond", &utcOffsetInfo.GmLeapSecond})
    utcOffsetInfo.EntityData.Children.Append("configured-leap-second", types.YChild{"ConfiguredLeapSecond", nil})
    for i := range utcOffsetInfo.ConfiguredLeapSecond {
        types.SetYListKey(utcOffsetInfo.ConfiguredLeapSecond[i], i)
        utcOffsetInfo.EntityData.Children.Append(types.GetSegmentPath(utcOffsetInfo.ConfiguredLeapSecond[i]), types.YChild{"ConfiguredLeapSecond", utcOffsetInfo.ConfiguredLeapSecond[i]})
    }
    utcOffsetInfo.EntityData.Leafs = types.NewOrderedMap()
    utcOffsetInfo.EntityData.Leafs.Append("source-type", types.YLeaf{"SourceType", utcOffsetInfo.SourceType})
    utcOffsetInfo.EntityData.Leafs.Append("source-file", types.YLeaf{"SourceFile", utcOffsetInfo.SourceFile})
    utcOffsetInfo.EntityData.Leafs.Append("source-expiry-date", types.YLeaf{"SourceExpiryDate", utcOffsetInfo.SourceExpiryDate})
    utcOffsetInfo.EntityData.Leafs.Append("polling-frequency", types.YLeaf{"PollingFrequency", utcOffsetInfo.PollingFrequency})

    utcOffsetInfo.EntityData.YListKeys = []string {}

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
    currentOffsetInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + currentOffsetInfo.EntityData.SegmentPath
    currentOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentOffsetInfo.EntityData.Children = types.NewOrderedMap()
    currentOffsetInfo.EntityData.Leafs = types.NewOrderedMap()
    currentOffsetInfo.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", currentOffsetInfo.Offset})
    currentOffsetInfo.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", currentOffsetInfo.Valid})
    currentOffsetInfo.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", currentOffsetInfo.Flag})

    currentOffsetInfo.EntityData.YListKeys = []string {}

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
    currentGmOffsetInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + currentGmOffsetInfo.EntityData.SegmentPath
    currentGmOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentGmOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentGmOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentGmOffsetInfo.EntityData.Children = types.NewOrderedMap()
    currentGmOffsetInfo.EntityData.Leafs = types.NewOrderedMap()
    currentGmOffsetInfo.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", currentGmOffsetInfo.Offset})
    currentGmOffsetInfo.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", currentGmOffsetInfo.Valid})
    currentGmOffsetInfo.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", currentGmOffsetInfo.Flag})

    currentGmOffsetInfo.EntityData.YListKeys = []string {}

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
    configuredOffsetInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + configuredOffsetInfo.EntityData.SegmentPath
    configuredOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOffsetInfo.EntityData.Children = types.NewOrderedMap()
    configuredOffsetInfo.EntityData.Leafs = types.NewOrderedMap()
    configuredOffsetInfo.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", configuredOffsetInfo.Offset})
    configuredOffsetInfo.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", configuredOffsetInfo.Valid})
    configuredOffsetInfo.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", configuredOffsetInfo.Flag})

    configuredOffsetInfo.EntityData.YListKeys = []string {}

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
    previousGmOffsetInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + previousGmOffsetInfo.EntityData.SegmentPath
    previousGmOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    previousGmOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    previousGmOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    previousGmOffsetInfo.EntityData.Children = types.NewOrderedMap()
    previousGmOffsetInfo.EntityData.Leafs = types.NewOrderedMap()
    previousGmOffsetInfo.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", previousGmOffsetInfo.Offset})
    previousGmOffsetInfo.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", previousGmOffsetInfo.Valid})
    previousGmOffsetInfo.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", previousGmOffsetInfo.Flag})

    previousGmOffsetInfo.EntityData.YListKeys = []string {}

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
    hardwareOffsetInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + hardwareOffsetInfo.EntityData.SegmentPath
    hardwareOffsetInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareOffsetInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareOffsetInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareOffsetInfo.EntityData.Children = types.NewOrderedMap()
    hardwareOffsetInfo.EntityData.Leafs = types.NewOrderedMap()
    hardwareOffsetInfo.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", hardwareOffsetInfo.Offset})
    hardwareOffsetInfo.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", hardwareOffsetInfo.Valid})
    hardwareOffsetInfo.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", hardwareOffsetInfo.Flag})

    hardwareOffsetInfo.EntityData.YListKeys = []string {}

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
    gmLeapSecond.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + gmLeapSecond.EntityData.SegmentPath
    gmLeapSecond.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmLeapSecond.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmLeapSecond.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmLeapSecond.EntityData.Children = types.NewOrderedMap()
    gmLeapSecond.EntityData.Leafs = types.NewOrderedMap()
    gmLeapSecond.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", gmLeapSecond.Offset})
    gmLeapSecond.EntityData.Leafs.Append("offset-start-date", types.YLeaf{"OffsetStartDate", gmLeapSecond.OffsetStartDate})
    gmLeapSecond.EntityData.Leafs.Append("offset-change", types.YLeaf{"OffsetChange", gmLeapSecond.OffsetChange})
    gmLeapSecond.EntityData.Leafs.Append("offset-applied", types.YLeaf{"OffsetApplied", gmLeapSecond.OffsetApplied})

    gmLeapSecond.EntityData.YListKeys = []string {}

    return &(gmLeapSecond.EntityData)
}

// Ptp_UtcOffsetInfo_ConfiguredLeapSecond
// The list of upcoming configured leap second
// updates
type Ptp_UtcOffsetInfo_ConfiguredLeapSecond struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    configuredLeapSecond.EntityData.SegmentPath = "configured-leap-second" + types.AddNoKeyToken(configuredLeapSecond)
    configuredLeapSecond.EntityData.AbsolutePath = "Cisco-IOS-XR-ptp-oper:ptp/utc-offset-info/" + configuredLeapSecond.EntityData.SegmentPath
    configuredLeapSecond.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredLeapSecond.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredLeapSecond.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredLeapSecond.EntityData.Children = types.NewOrderedMap()
    configuredLeapSecond.EntityData.Leafs = types.NewOrderedMap()
    configuredLeapSecond.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", configuredLeapSecond.Offset})
    configuredLeapSecond.EntityData.Leafs.Append("offset-start-date", types.YLeaf{"OffsetStartDate", configuredLeapSecond.OffsetStartDate})
    configuredLeapSecond.EntityData.Leafs.Append("offset-change", types.YLeaf{"OffsetChange", configuredLeapSecond.OffsetChange})
    configuredLeapSecond.EntityData.Leafs.Append("offset-applied", types.YLeaf{"OffsetApplied", configuredLeapSecond.OffsetApplied})

    configuredLeapSecond.EntityData.YListKeys = []string {}

    return &(configuredLeapSecond.EntityData)
}

