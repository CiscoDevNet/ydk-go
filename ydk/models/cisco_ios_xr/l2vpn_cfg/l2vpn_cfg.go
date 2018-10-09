// This module contains a collection of YANG definitions
// for Cisco IOS-XR l2vpn package configuration.
// 
// This module contains definitions
// for the following management objects:
//   l2vpn: L2VPN configuration
//   generic-interface-lists: generic interface lists
//   evpn: evpn
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-snmp-agent-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package l2vpn_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package l2vpn_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2vpn-cfg l2vpn}", reflect.TypeOf(L2vpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2vpn-cfg:l2vpn", reflect.TypeOf(L2vpn{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2vpn-cfg generic-interface-lists}", reflect.TypeOf(GenericInterfaceLists{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2vpn-cfg:generic-interface-lists", reflect.TypeOf(GenericInterfaceLists{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2vpn-cfg evpn}", reflect.TypeOf(Evpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2vpn-cfg:evpn", reflect.TypeOf(Evpn{}))
}

// FlowLabelTlvCode represents Flow label tlv code
type FlowLabelTlvCode string

const (
    // Set Flow Label Legacy TLV code (DEPRECATED)
    FlowLabelTlvCode_Y_17 FlowLabelTlvCode = "17"

    // Disable Sending Flow Label Legacy TLV
    FlowLabelTlvCode_disable FlowLabelTlvCode = "disable"
)

// MacAging represents Mac aging
type MacAging string

const (
    // Absolute aging type
    MacAging_absolute MacAging = "absolute"

    // Inactivity aging type
    MacAging_inactivity MacAging = "inactivity"
)

// MacLimitAction represents Mac limit action
type MacLimitAction string

const (
    // No action
    MacLimitAction_none MacLimitAction = "none"

    // Flood Mac Limit Action
    MacLimitAction_flood MacLimitAction = "flood"

    // NoFlood Mac Limit Action
    MacLimitAction_no_flood MacLimitAction = "no-flood"

    // Shutdown Mac Limit Action
    MacLimitAction_shutdown MacLimitAction = "shutdown"
)

// BdmacLearn represents Bdmac learn
type BdmacLearn string

const (
    // Disable Learning
    BdmacLearn_disable_learning BdmacLearn = "disable-learning"
)

// Interworking represents Interworking
type Interworking string

const (
    // Ethernet interworking
    Interworking_ethernet Interworking = "ethernet"

    // IPv4 interworking
    Interworking_ipv4 Interworking = "ipv4"
)

// PwSwitchingPointTlv represents Pw switching point tlv
type PwSwitchingPointTlv string

const (
    // Hide TLV
    PwSwitchingPointTlv_hide PwSwitchingPointTlv = "hide"
)

// L2tpv3Sequencing represents L2tpv3 sequencing
type L2tpv3Sequencing string

const (
    // Sequencing is off
    L2tpv3Sequencing_off L2tpv3Sequencing = "off"

    // Sequencing on both transmit and receive side
    L2tpv3Sequencing_both L2tpv3Sequencing = "both"
)

// InterfaceProfile represents Interface profile
type InterfaceProfile string

const (
    // Set the snooping
    InterfaceProfile_snoop InterfaceProfile = "snoop"

    // disable DHCP protocol
    InterfaceProfile_dhcp_protocol InterfaceProfile = "dhcp-protocol"
)

// BgpRouteTargetRole represents Bgp route target role
type BgpRouteTargetRole string

const (
    // Both Import and export roles
    BgpRouteTargetRole_both BgpRouteTargetRole = "both"

    // Import role
    BgpRouteTargetRole_import_ BgpRouteTargetRole = "import"

    // Export role
    BgpRouteTargetRole_export BgpRouteTargetRole = "export"
)

// ErpPort represents Erp port
type ErpPort string

const (
    // ERP port type none
    ErpPort_none ErpPort = "none"

    // ERP port type virtual
    ErpPort_virtual ErpPort = "virtual"

    // ERP port type interface
    ErpPort_interface_ ErpPort = "interface"
)

// BgpRouteTarget represents Bgp route target
type BgpRouteTarget string

const (
    // RT is default type
    BgpRouteTarget_no_stitching BgpRouteTarget = "no-stitching"

    // RT is for stitching (Golf-L2) (DEPRECATED)
    BgpRouteTarget_stitching BgpRouteTarget = "stitching"
)

// FlowLabelLoadBalance represents Flow label load balance
type FlowLabelLoadBalance string

const (
    // Flow Label load balance is off
    FlowLabelLoadBalance_off FlowLabelLoadBalance = "off"

    // Delete Flow Label on receive side
    FlowLabelLoadBalance_receive FlowLabelLoadBalance = "receive"

    // Insert Flow Label on transmit side
    FlowLabelLoadBalance_transmit FlowLabelLoadBalance = "transmit"

    // Insert/Delete Flow Label on transmit/receive
    // side
    FlowLabelLoadBalance_both FlowLabelLoadBalance = "both"
)

// L2vpnVerification represents L2vpn verification
type L2vpnVerification string

const (
    // enable verification
    L2vpnVerification_enable L2vpnVerification = "enable"

    // disable verification
    L2vpnVerification_disable L2vpnVerification = "disable"
)

// MacLearn represents Mac learn
type MacLearn string

const (
    // Mac Learning
    MacLearn_default_learning MacLearn = "default-learning"

    // Enable Learning
    MacLearn_enable_learning MacLearn = "enable-learning"

    // Disable Learning
    MacLearn_disable_learning MacLearn = "disable-learning"
)

// Erpaps represents Erpaps
type Erpaps string

const (
    // ERP APS type interface
    Erpaps_interface_ Erpaps = "interface"

    // ERP APS type bridge domain
    Erpaps_bridge_domain Erpaps = "bridge-domain"

    // ERP APS type xconnect
    Erpaps_xconnect Erpaps = "xconnect"

    // ERP APS type none
    Erpaps_none Erpaps = "none"
)

// VccvVerification represents Vccv verification
type VccvVerification string

const (
    // No connectivity verification over VCCV
    VccvVerification_none VccvVerification = "none"

    // LSP Ping over VCCV
    VccvVerification_lsp_ping VccvVerification = "lsp-ping"
)

// TransportMode represents Transport mode
type TransportMode string

const (
    // Ethernet port mode
    TransportMode_ethernet TransportMode = "ethernet"

    // Vlan tagged mode
    TransportMode_vlan TransportMode = "vlan"

    // Vlan tagged passthrough mode
    TransportMode_vlan_passthrough TransportMode = "vlan-passthrough"
)

// BackupDisable represents Backup disable
type BackupDisable string

const (
    // Never
    BackupDisable_never BackupDisable = "never"

    // Delay seconds
    BackupDisable_delay BackupDisable = "delay"
)

// LoadBalance represents Load balance
type LoadBalance string

const (
    // Source and Destination MAC hashing
    LoadBalance_source_dest_mac LoadBalance = "source-dest-mac"

    // Source and Destination IP hashing
    LoadBalance_source_dest_ip LoadBalance = "source-dest-ip"

    // PW Label hashing
    LoadBalance_pseudowire_label LoadBalance = "pseudowire-label"
)

// ErpPort1 represents Erp port1
type ErpPort1 string

const (
    // ERP main port 0
    ErpPort1_port0 ErpPort1 = "port0"

    // ERP main port 1
    ErpPort1_port1 ErpPort1 = "port1"
)

// InterfaceTrafficFlood represents Interface traffic flood
type InterfaceTrafficFlood string

const (
    // Traffic flooding
    InterfaceTrafficFlood_traffic_flooding InterfaceTrafficFlood = "traffic-flooding"

    // Enable Flooding
    InterfaceTrafficFlood_enable_flooding InterfaceTrafficFlood = "enable-flooding"

    // Disable flooding
    InterfaceTrafficFlood_disable_flooding InterfaceTrafficFlood = "disable-flooding"
)

// MacFlushMode represents Mac flush mode
type MacFlushMode string

const (
    // MVRP MAC Flushing
    MacFlushMode_mvrp MacFlushMode = "mvrp"
)

// L2tpCookieSize represents L2tp cookie size
type L2tpCookieSize string

const (
    // Cookie size is zero bytes
    L2tpCookieSize_zero L2tpCookieSize = "zero"

    // Cookie size is four bytes
    L2tpCookieSize_four L2tpCookieSize = "four"

    // Cookie size is eight bytes
    L2tpCookieSize_eight L2tpCookieSize = "eight"
)

// StormControl represents Storm control
type StormControl string

const (
    // Unknown-unicast Storm Control
    StormControl_unicast StormControl = "unicast"

    // Multicast Storm Control
    StormControl_multicast StormControl = "multicast"

    // Broadcast Storm Control
    StormControl_broadcast StormControl = "broadcast"
)

// BgpRouteDistinguisher represents Bgp route distinguisher
type BgpRouteDistinguisher string

const (
    // RD automatically assigned
    BgpRouteDistinguisher_auto BgpRouteDistinguisher = "auto"

    // RD in 2 byte AS:nn format
    BgpRouteDistinguisher_two_byte_as BgpRouteDistinguisher = "two-byte-as"

    // RD in 4 byte AS:nn format
    BgpRouteDistinguisher_four_byte_as BgpRouteDistinguisher = "four-byte-as"

    // RD in IpV4address
    BgpRouteDistinguisher_ipv4_address BgpRouteDistinguisher = "ipv4-address"
)

// MacNotification represents Mac notification
type MacNotification string

const (
    // No_Notification Trap
    MacNotification_no_notif MacNotification = "no-notif"

    // syslog message
    MacNotification_syslog MacNotification = "syslog"

    // Snmp Trap
    MacNotification_trap MacNotification = "trap"

    // Syslog_snmp Trap
    MacNotification_syslog_snmp MacNotification = "syslog-snmp"
)

// BgpRouteTargetFormat represents Bgp route target format
type BgpRouteTargetFormat string

const (
    // No route target
    BgpRouteTargetFormat_none BgpRouteTargetFormat = "none"

    // 2 Byte AS:nn format
    BgpRouteTargetFormat_two_byte_as BgpRouteTargetFormat = "two-byte-as"

    // 4 byte AS:nn format
    BgpRouteTargetFormat_four_byte_as BgpRouteTargetFormat = "four-byte-as"

    // IP:nn format
    BgpRouteTargetFormat_ipv4_address BgpRouteTargetFormat = "ipv4-address"

    // a.a.i format
    BgpRouteTargetFormat_es_import BgpRouteTargetFormat = "es-import"
)

// MplsSignalingProtocol represents Mpls signaling protocol
type MplsSignalingProtocol string

const (
    // No signaling
    MplsSignalingProtocol_none MplsSignalingProtocol = "none"

    // LDP
    MplsSignalingProtocol_ldp MplsSignalingProtocol = "ldp"
)

// EvpnSide represents Evpn side
type EvpnSide string

const (
    // EVPN Instance side defined as regular
    EvpnSide_evpn_side_regular EvpnSide = "evpn-side-regular"

    // EVPN Instance side defined as stitching
    EvpnSide_evpn_side_stitching EvpnSide = "evpn-side-stitching"
)

// ControlWord represents Control word
type ControlWord string

const (
    // Enable control word
    ControlWord_enable ControlWord = "enable"

    // Disable control word
    ControlWord_disable ControlWord = "disable"
)

// PreferredPath represents Preferred path
type PreferredPath string

const (
    // TE Tunnel
    PreferredPath_te_tunnel PreferredPath = "te-tunnel"

    // IP Tunnel
    PreferredPath_ip_tunnel PreferredPath = "ip-tunnel"

    // TP Tunnel
    PreferredPath_tp_tunnel PreferredPath = "tp-tunnel"

    // SR TE Policy
    PreferredPath_sr_te_policy PreferredPath = "sr-te-policy"
)

// EvpnEncapsulation represents Evpn encapsulation
type EvpnEncapsulation string

const (
    // VXLAN Encapsulation
    EvpnEncapsulation_evpn_encapsulationvxlan EvpnEncapsulation = "evpn-encapsulationvxlan"

    // MPLS Encapsulation
    EvpnEncapsulation_evpn_encapsulation_mpls EvpnEncapsulation = "evpn-encapsulation-mpls"
)

// MplsSequencing represents Mpls sequencing
type MplsSequencing string

const (
    // Sequencing is off
    MplsSequencing_off MplsSequencing = "off"

    // Sequencing on transmit side
    MplsSequencing_transmit MplsSequencing = "transmit"

    // Sequencing on receive side
    MplsSequencing_receive MplsSequencing = "receive"

    // Sequencing on both transmit and receive side
    MplsSequencing_both MplsSequencing = "both"
)

// EthernetSegmentLoadBalance represents Ethernet segment load balance
type EthernetSegmentLoadBalance string

const (
    // Single Active
    EthernetSegmentLoadBalance_single_active EthernetSegmentLoadBalance = "single-active"
)

// L2tpSignalingProtocol represents L2tp signaling protocol
type L2tpSignalingProtocol string

const (
    // No signaling
    L2tpSignalingProtocol_none L2tpSignalingProtocol = "none"

    // L2TPv3
    L2tpSignalingProtocol_l2tpv3 L2tpSignalingProtocol = "l2tpv3"
)

// EthernetSegmentIdentifier represents Ethernet segment identifier
type EthernetSegmentIdentifier string

const (
    // ESI type 0
    EthernetSegmentIdentifier_type0 EthernetSegmentIdentifier = "type0"

    // Legacy ESI type
    EthernetSegmentIdentifier_legacy EthernetSegmentIdentifier = "legacy"

    // Override ESI type
    EthernetSegmentIdentifier_override EthernetSegmentIdentifier = "override"
)

// BridgeDomainTransportMode represents Bridge domain transport mode
type BridgeDomainTransportMode string

const (
    // Vlan tagged passthrough mode
    BridgeDomainTransportMode_vlan_passthrough BridgeDomainTransportMode = "vlan-passthrough"
)

// LdpVplsId represents Ldp vpls id
type LdpVplsId string

const (
    // VPLS-ID in 2 byte AS:nn format
    LdpVplsId_two_byte_as LdpVplsId = "two-byte-as"

    // VPLS-ID in IPv4 IP:nn format
    LdpVplsId_ipv4_address LdpVplsId = "ipv4-address"
)

// L2Encapsulation represents L2 encapsulation
type L2Encapsulation string

const (
    // Vlan tagged mode
    L2Encapsulation_vlan L2Encapsulation = "vlan"

    // Ethernet port mode
    L2Encapsulation_ethernet L2Encapsulation = "ethernet"
)

// EthernetSegmentServiceCarving represents Ethernet segment service carving
type EthernetSegmentServiceCarving string

const (
    // HRW
    EthernetSegmentServiceCarving_hrw EthernetSegmentServiceCarving = "hrw"
)

// L2vpnLogging represents L2vpn logging
type L2vpnLogging string

const (
    // enable logging
    L2vpnLogging_enable L2vpnLogging = "enable"

    // disable logging
    L2vpnLogging_disable L2vpnLogging = "disable"
)

// MacWithdrawBehavior represents Mac withdraw behavior
type MacWithdrawBehavior string

const (
    // MAC Withdrawal sent on state-down (legacy)
    MacWithdrawBehavior_legacy MacWithdrawBehavior = "legacy"

    // Optimized MAC Withdrawal
    MacWithdrawBehavior_optimized MacWithdrawBehavior = "optimized"
)

// RplRole represents Rpl role
type RplRole string

const (
    // ERP RPL owner
    RplRole_owner RplRole = "owner"

    // ERP RPL neighbor
    RplRole_neighbor RplRole = "neighbor"

    // ERP RPL next neighbor
    RplRole_next_neighbor RplRole = "next-neighbor"
)

// TypeOfServiceMode represents Type of service mode
type TypeOfServiceMode string

const (
    // Do not reflect the type of service
    TypeOfServiceMode_none TypeOfServiceMode = "none"

    // Reflect the type of service
    TypeOfServiceMode_reflect TypeOfServiceMode = "reflect"
)

// PortDownFlush represents Port down flush
type PortDownFlush string

const (
    // MAC Port Down Flush
    PortDownFlush_port_down_flush PortDownFlush = "port-down-flush"

    // Enable Port Down Flush
    PortDownFlush_enable_port_down_flush PortDownFlush = "enable-port-down-flush"

    // Disable Port Down Flush
    PortDownFlush_disable_port_down_flush PortDownFlush = "disable-port-down-flush"
)

// L2vpnCapabilityMode represents L2vpn capability mode
type L2vpnCapabilityMode string

const (
    // Compute global capability as the highest node
    // capability
    L2vpnCapabilityMode_high_mode L2vpnCapabilityMode = "high-mode"

    // Disable global capability re-computation
    L2vpnCapabilityMode_single_mode L2vpnCapabilityMode = "single-mode"
)

// MacSecureAction represents Mac secure action
type MacSecureAction string

const (
    // MAC Secure Action Restrict
    MacSecureAction_restrict MacSecureAction = "restrict"

    // No Action
    MacSecureAction_none MacSecureAction = "none"

    // MAC Secure Action Shutdown
    MacSecureAction_shutdown MacSecureAction = "shutdown"
)

// L2vpn
// L2VPN configuration
type L2vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Non-Stop Routing. The type is interface{}.
    Nsr interface{}

    // Ignore MTU Mismatch for XCs. The type is interface{}.
    MtuMismatchIgnore interface{}

    // Topology change notification propagation. The type is interface{}.
    TcnPropagation interface{}

    // Configure PW OAM refresh interval. The type is interface{} with range:
    // 1..4095. Units are second.
    PwoamRefresh interface{}

    // Enable flow load balancing on l2vpn bridges. The type is LoadBalance.
    LoadBalance interface{}

    // MS-PW global description. The type is string with length: 1..64.
    MspwDescription interface{}

    // Configure MAC limit threshold percent. The type is interface{} with range:
    // 1..100. Units are percentage.
    MacLimitThreshold interface{}

    // Disable PW status. The type is interface{}.
    PwStatusDisable interface{}

    // Enable L2VPN feature. The type is interface{}.
    Enable interface{}

    // Enable PW grouping. The type is interface{}.
    PwGrouping interface{}

    // L2VPN Capability Mode. The type is L2vpnCapabilityMode.
    Capability interface{}

    // Global L2VPN Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    L2vpnRouterId interface{}

    // Pseudowire-routing attributes.
    PwRouting L2vpn_PwRouting

    // L2VPN neighbor submode.
    Neighbor L2vpn_Neighbor

    // L2VPN databases.
    Database L2vpn_Database

    // L2VPN PBB Global.
    Pbb L2vpn_Pbb

    // Global auto-discovery attributes.
    AutoDiscovery L2vpn_AutoDiscovery

    // L2VPN utilities.
    Utility L2vpn_Utility

    // SNMP related configuration.
    Snmp L2vpn_Snmp
}

func (l2vpn *L2vpn) GetEntityData() *types.CommonEntityData {
    l2vpn.EntityData.YFilter = l2vpn.YFilter
    l2vpn.EntityData.YangName = "l2vpn"
    l2vpn.EntityData.BundleName = "cisco_ios_xr"
    l2vpn.EntityData.ParentYangName = "Cisco-IOS-XR-l2vpn-cfg"
    l2vpn.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:l2vpn"
    l2vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2vpn.EntityData.Children = types.NewOrderedMap()
    l2vpn.EntityData.Children.Append("pw-routing", types.YChild{"PwRouting", &l2vpn.PwRouting})
    l2vpn.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &l2vpn.Neighbor})
    l2vpn.EntityData.Children.Append("database", types.YChild{"Database", &l2vpn.Database})
    l2vpn.EntityData.Children.Append("pbb", types.YChild{"Pbb", &l2vpn.Pbb})
    l2vpn.EntityData.Children.Append("auto-discovery", types.YChild{"AutoDiscovery", &l2vpn.AutoDiscovery})
    l2vpn.EntityData.Children.Append("utility", types.YChild{"Utility", &l2vpn.Utility})
    l2vpn.EntityData.Children.Append("snmp", types.YChild{"Snmp", &l2vpn.Snmp})
    l2vpn.EntityData.Leafs = types.NewOrderedMap()
    l2vpn.EntityData.Leafs.Append("nsr", types.YLeaf{"Nsr", l2vpn.Nsr})
    l2vpn.EntityData.Leafs.Append("mtu-mismatch-ignore", types.YLeaf{"MtuMismatchIgnore", l2vpn.MtuMismatchIgnore})
    l2vpn.EntityData.Leafs.Append("tcn-propagation", types.YLeaf{"TcnPropagation", l2vpn.TcnPropagation})
    l2vpn.EntityData.Leafs.Append("pwoam-refresh", types.YLeaf{"PwoamRefresh", l2vpn.PwoamRefresh})
    l2vpn.EntityData.Leafs.Append("load-balance", types.YLeaf{"LoadBalance", l2vpn.LoadBalance})
    l2vpn.EntityData.Leafs.Append("mspw-description", types.YLeaf{"MspwDescription", l2vpn.MspwDescription})
    l2vpn.EntityData.Leafs.Append("mac-limit-threshold", types.YLeaf{"MacLimitThreshold", l2vpn.MacLimitThreshold})
    l2vpn.EntityData.Leafs.Append("pw-status-disable", types.YLeaf{"PwStatusDisable", l2vpn.PwStatusDisable})
    l2vpn.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", l2vpn.Enable})
    l2vpn.EntityData.Leafs.Append("pw-grouping", types.YLeaf{"PwGrouping", l2vpn.PwGrouping})
    l2vpn.EntityData.Leafs.Append("capability", types.YLeaf{"Capability", l2vpn.Capability})
    l2vpn.EntityData.Leafs.Append("l2vpn-router-id", types.YLeaf{"L2vpnRouterId", l2vpn.L2vpnRouterId})

    l2vpn.EntityData.YListKeys = []string {}

    return &(l2vpn.EntityData)
}

// L2vpn_PwRouting
// Pseudowire-routing attributes
type L2vpn_PwRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire-routing Global ID. The type is interface{} with range:
    // 1..4294967295.
    PwRoutingGlobalId interface{}

    // Enable Autodiscovery BGP Pseudowire-routing BGP.
    PwRoutingBgp L2vpn_PwRouting_PwRoutingBgp
}

func (pwRouting *L2vpn_PwRouting) GetEntityData() *types.CommonEntityData {
    pwRouting.EntityData.YFilter = pwRouting.YFilter
    pwRouting.EntityData.YangName = "pw-routing"
    pwRouting.EntityData.BundleName = "cisco_ios_xr"
    pwRouting.EntityData.ParentYangName = "l2vpn"
    pwRouting.EntityData.SegmentPath = "pw-routing"
    pwRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pwRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pwRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pwRouting.EntityData.Children = types.NewOrderedMap()
    pwRouting.EntityData.Children.Append("pw-routing-bgp", types.YChild{"PwRoutingBgp", &pwRouting.PwRoutingBgp})
    pwRouting.EntityData.Leafs = types.NewOrderedMap()
    pwRouting.EntityData.Leafs.Append("pw-routing-global-id", types.YLeaf{"PwRoutingGlobalId", pwRouting.PwRoutingGlobalId})

    pwRouting.EntityData.YListKeys = []string {}

    return &(pwRouting.EntityData)
}

// L2vpn_PwRouting_PwRoutingBgp
// Enable Autodiscovery BGP Pseudowire-routing BGP
type L2vpn_PwRouting_PwRoutingBgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Route Distinguisher.
    EvpnRouteDistinguisher L2vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher
}

func (pwRoutingBgp *L2vpn_PwRouting_PwRoutingBgp) GetEntityData() *types.CommonEntityData {
    pwRoutingBgp.EntityData.YFilter = pwRoutingBgp.YFilter
    pwRoutingBgp.EntityData.YangName = "pw-routing-bgp"
    pwRoutingBgp.EntityData.BundleName = "cisco_ios_xr"
    pwRoutingBgp.EntityData.ParentYangName = "pw-routing"
    pwRoutingBgp.EntityData.SegmentPath = "pw-routing-bgp"
    pwRoutingBgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pwRoutingBgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pwRoutingBgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pwRoutingBgp.EntityData.Children = types.NewOrderedMap()
    pwRoutingBgp.EntityData.Children.Append("evpn-route-distinguisher", types.YChild{"EvpnRouteDistinguisher", &pwRoutingBgp.EvpnRouteDistinguisher})
    pwRoutingBgp.EntityData.Leafs = types.NewOrderedMap()
    pwRoutingBgp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pwRoutingBgp.Enable})

    pwRoutingBgp.EntityData.YListKeys = []string {}

    return &(pwRoutingBgp.EntityData)
}

// L2vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher
// Route Distinguisher
type L2vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (evpnRouteDistinguisher *L2vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher) GetEntityData() *types.CommonEntityData {
    evpnRouteDistinguisher.EntityData.YFilter = evpnRouteDistinguisher.YFilter
    evpnRouteDistinguisher.EntityData.YangName = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteDistinguisher.EntityData.ParentYangName = "pw-routing-bgp"
    evpnRouteDistinguisher.EntityData.SegmentPath = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteDistinguisher.EntityData.Children = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnRouteDistinguisher.Type})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", evpnRouteDistinguisher.As})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex})
    evpnRouteDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", evpnRouteDistinguisher.Address})
    evpnRouteDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex})

    evpnRouteDistinguisher.EntityData.YListKeys = []string {}

    return &(evpnRouteDistinguisher.EntityData)
}

// L2vpn_Neighbor
// L2VPN neighbor submode
type L2vpn_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable targetted LDP session flap action. The type is interface{}.
    LdpFlap interface{}
}

func (neighbor *L2vpn_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "l2vpn"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("ldp-flap", types.YLeaf{"LdpFlap", neighbor.LdpFlap})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// L2vpn_Database
// L2VPN databases
type L2vpn_Database struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of G8032 Ring.
    G8032Rings L2vpn_Database_G8032Rings

    // List of xconnect groups.
    XconnectGroups L2vpn_Database_XconnectGroups

    // List of bridge groups.
    BridgeDomainGroups L2vpn_Database_BridgeDomainGroups

    // List of pseudowire classes.
    PseudowireClasses L2vpn_Database_PseudowireClasses

    // List of VLAN Switches.
    VlanSwitches L2vpn_Database_VlanSwitches

    // List of Flexible XConnect Services.
    FlexibleXconnectServiceTable L2vpn_Database_FlexibleXconnectServiceTable

    // Redundancy groups.
    Redundancy L2vpn_Database_Redundancy
}

func (database *L2vpn_Database) GetEntityData() *types.CommonEntityData {
    database.EntityData.YFilter = database.YFilter
    database.EntityData.YangName = "database"
    database.EntityData.BundleName = "cisco_ios_xr"
    database.EntityData.ParentYangName = "l2vpn"
    database.EntityData.SegmentPath = "database"
    database.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    database.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    database.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    database.EntityData.Children = types.NewOrderedMap()
    database.EntityData.Children.Append("g8032-rings", types.YChild{"G8032Rings", &database.G8032Rings})
    database.EntityData.Children.Append("xconnect-groups", types.YChild{"XconnectGroups", &database.XconnectGroups})
    database.EntityData.Children.Append("bridge-domain-groups", types.YChild{"BridgeDomainGroups", &database.BridgeDomainGroups})
    database.EntityData.Children.Append("pseudowire-classes", types.YChild{"PseudowireClasses", &database.PseudowireClasses})
    database.EntityData.Children.Append("vlan-switches", types.YChild{"VlanSwitches", &database.VlanSwitches})
    database.EntityData.Children.Append("flexible-xconnect-service-table", types.YChild{"FlexibleXconnectServiceTable", &database.FlexibleXconnectServiceTable})
    database.EntityData.Children.Append("redundancy", types.YChild{"Redundancy", &database.Redundancy})
    database.EntityData.Leafs = types.NewOrderedMap()

    database.EntityData.YListKeys = []string {}

    return &(database.EntityData)
}

// L2vpn_Database_G8032Rings
// List of G8032 Ring
type L2vpn_Database_G8032Rings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // G8032 Ring. The type is slice of L2vpn_Database_G8032Rings_G8032Ring.
    G8032Ring []*L2vpn_Database_G8032Rings_G8032Ring
}

func (g8032Rings *L2vpn_Database_G8032Rings) GetEntityData() *types.CommonEntityData {
    g8032Rings.EntityData.YFilter = g8032Rings.YFilter
    g8032Rings.EntityData.YangName = "g8032-rings"
    g8032Rings.EntityData.BundleName = "cisco_ios_xr"
    g8032Rings.EntityData.ParentYangName = "database"
    g8032Rings.EntityData.SegmentPath = "g8032-rings"
    g8032Rings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    g8032Rings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    g8032Rings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    g8032Rings.EntityData.Children = types.NewOrderedMap()
    g8032Rings.EntityData.Children.Append("g8032-ring", types.YChild{"G8032Ring", nil})
    for i := range g8032Rings.G8032Ring {
        g8032Rings.EntityData.Children.Append(types.GetSegmentPath(g8032Rings.G8032Ring[i]), types.YChild{"G8032Ring", g8032Rings.G8032Ring[i]})
    }
    g8032Rings.EntityData.Leafs = types.NewOrderedMap()

    g8032Rings.EntityData.YListKeys = []string {}

    return &(g8032Rings.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring
// G8032 Ring
type L2vpn_Database_G8032Rings_G8032Ring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the G8032 ring. The type is string with
    // length: 1..32.
    G8032RingName interface{}

    // Specify the G.8032 instance as open ring. The type is interface{}.
    OpenRing interface{}

    // Vlan IDs in the format of a-b,c,d,e-f,g ,untagged. The type is string.
    ExclusionList interface{}

    // Ethernet ring protection provider bridge. The type is interface{}.
    ErpProviderBridge interface{}

    // Ethernet ring protection port0.
    ErpPort0s L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s

    // List of ethernet ring protection instance.
    ErpInstances L2vpn_Database_G8032Rings_G8032Ring_ErpInstances

    // Ethernet ring protection port0.
    ErpPort1s L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s
}

func (g8032Ring *L2vpn_Database_G8032Rings_G8032Ring) GetEntityData() *types.CommonEntityData {
    g8032Ring.EntityData.YFilter = g8032Ring.YFilter
    g8032Ring.EntityData.YangName = "g8032-ring"
    g8032Ring.EntityData.BundleName = "cisco_ios_xr"
    g8032Ring.EntityData.ParentYangName = "g8032-rings"
    g8032Ring.EntityData.SegmentPath = "g8032-ring" + types.AddKeyToken(g8032Ring.G8032RingName, "g8032-ring-name")
    g8032Ring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    g8032Ring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    g8032Ring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    g8032Ring.EntityData.Children = types.NewOrderedMap()
    g8032Ring.EntityData.Children.Append("erp-port0s", types.YChild{"ErpPort0s", &g8032Ring.ErpPort0s})
    g8032Ring.EntityData.Children.Append("erp-instances", types.YChild{"ErpInstances", &g8032Ring.ErpInstances})
    g8032Ring.EntityData.Children.Append("erp-port1s", types.YChild{"ErpPort1s", &g8032Ring.ErpPort1s})
    g8032Ring.EntityData.Leafs = types.NewOrderedMap()
    g8032Ring.EntityData.Leafs.Append("g8032-ring-name", types.YLeaf{"G8032RingName", g8032Ring.G8032RingName})
    g8032Ring.EntityData.Leafs.Append("open-ring", types.YLeaf{"OpenRing", g8032Ring.OpenRing})
    g8032Ring.EntityData.Leafs.Append("exclusion-list", types.YLeaf{"ExclusionList", g8032Ring.ExclusionList})
    g8032Ring.EntityData.Leafs.Append("erp-provider-bridge", types.YLeaf{"ErpProviderBridge", g8032Ring.ErpProviderBridge})

    g8032Ring.EntityData.YListKeys = []string {"G8032RingName"}

    return &(g8032Ring.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s
// Ethernet ring protection port0
type L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure ERP main port0. The type is slice of
    // L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s_ErpPort0.
    ErpPort0 []*L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s_ErpPort0
}

func (erpPort0s *L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s) GetEntityData() *types.CommonEntityData {
    erpPort0s.EntityData.YFilter = erpPort0s.YFilter
    erpPort0s.EntityData.YangName = "erp-port0s"
    erpPort0s.EntityData.BundleName = "cisco_ios_xr"
    erpPort0s.EntityData.ParentYangName = "g8032-ring"
    erpPort0s.EntityData.SegmentPath = "erp-port0s"
    erpPort0s.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort0s.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort0s.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort0s.EntityData.Children = types.NewOrderedMap()
    erpPort0s.EntityData.Children.Append("erp-port0", types.YChild{"ErpPort0", nil})
    for i := range erpPort0s.ErpPort0 {
        erpPort0s.EntityData.Children.Append(types.GetSegmentPath(erpPort0s.ErpPort0[i]), types.YChild{"ErpPort0", erpPort0s.ErpPort0[i]})
    }
    erpPort0s.EntityData.Leafs = types.NewOrderedMap()

    erpPort0s.EntityData.YListKeys = []string {}

    return &(erpPort0s.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s_ErpPort0
// Configure ERP main port0
type L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s_ErpPort0 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port0 interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Ethernet ring protection port0 monitor. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Monitor interface{}
}

func (erpPort0 *L2vpn_Database_G8032Rings_G8032Ring_ErpPort0s_ErpPort0) GetEntityData() *types.CommonEntityData {
    erpPort0.EntityData.YFilter = erpPort0.YFilter
    erpPort0.EntityData.YangName = "erp-port0"
    erpPort0.EntityData.BundleName = "cisco_ios_xr"
    erpPort0.EntityData.ParentYangName = "erp-port0s"
    erpPort0.EntityData.SegmentPath = "erp-port0" + types.AddKeyToken(erpPort0.InterfaceName, "interface-name")
    erpPort0.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort0.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort0.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort0.EntityData.Children = types.NewOrderedMap()
    erpPort0.EntityData.Leafs = types.NewOrderedMap()
    erpPort0.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", erpPort0.InterfaceName})
    erpPort0.EntityData.Leafs.Append("monitor", types.YLeaf{"Monitor", erpPort0.Monitor})

    erpPort0.EntityData.YListKeys = []string {"InterfaceName"}

    return &(erpPort0.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpInstances
// List of ethernet ring protection instance
type L2vpn_Database_G8032Rings_G8032Ring_ErpInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet ring protection instance. The type is slice of
    // L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance.
    ErpInstance []*L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance
}

func (erpInstances *L2vpn_Database_G8032Rings_G8032Ring_ErpInstances) GetEntityData() *types.CommonEntityData {
    erpInstances.EntityData.YFilter = erpInstances.YFilter
    erpInstances.EntityData.YangName = "erp-instances"
    erpInstances.EntityData.BundleName = "cisco_ios_xr"
    erpInstances.EntityData.ParentYangName = "g8032-ring"
    erpInstances.EntityData.SegmentPath = "erp-instances"
    erpInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpInstances.EntityData.Children = types.NewOrderedMap()
    erpInstances.EntityData.Children.Append("erp-instance", types.YChild{"ErpInstance", nil})
    for i := range erpInstances.ErpInstance {
        erpInstances.EntityData.Children.Append(types.GetSegmentPath(erpInstances.ErpInstance[i]), types.YChild{"ErpInstance", erpInstances.ErpInstance[i]})
    }
    erpInstances.EntityData.Leafs = types.NewOrderedMap()

    erpInstances.EntityData.YListKeys = []string {}

    return &(erpInstances.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance
// Ethernet ring protection instance
type L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ERP instance number. The type is interface{} with
    // range: 1..2.
    ErpInstanceId interface{}

    // Ethernet ring protection instance description. The type is string with
    // length: 1..32.
    Description interface{}

    // Associates a set of VLAN IDs with the G .8032 instance. The type is string.
    InclusionList interface{}

    // Ethernet ring protection instance profile. The type is string with length:
    // 1..32.
    Profile interface{}

    // Ring protection link.
    Rpl L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl

    // Automatic protection switching.
    Aps L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps
}

func (erpInstance *L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance) GetEntityData() *types.CommonEntityData {
    erpInstance.EntityData.YFilter = erpInstance.YFilter
    erpInstance.EntityData.YangName = "erp-instance"
    erpInstance.EntityData.BundleName = "cisco_ios_xr"
    erpInstance.EntityData.ParentYangName = "erp-instances"
    erpInstance.EntityData.SegmentPath = "erp-instance" + types.AddKeyToken(erpInstance.ErpInstanceId, "erp-instance-id")
    erpInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpInstance.EntityData.Children = types.NewOrderedMap()
    erpInstance.EntityData.Children.Append("rpl", types.YChild{"Rpl", &erpInstance.Rpl})
    erpInstance.EntityData.Children.Append("aps", types.YChild{"Aps", &erpInstance.Aps})
    erpInstance.EntityData.Leafs = types.NewOrderedMap()
    erpInstance.EntityData.Leafs.Append("erp-instance-id", types.YLeaf{"ErpInstanceId", erpInstance.ErpInstanceId})
    erpInstance.EntityData.Leafs.Append("description", types.YLeaf{"Description", erpInstance.Description})
    erpInstance.EntityData.Leafs.Append("inclusion-list", types.YLeaf{"InclusionList", erpInstance.InclusionList})
    erpInstance.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", erpInstance.Profile})

    erpInstance.EntityData.YListKeys = []string {"ErpInstanceId"}

    return &(erpInstance.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl
// Ring protection link
type L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ERP main port number. The type is ErpPort1.
    Port interface{}

    // RPL role. The type is RplRole.
    Role interface{}
}

func (rpl *L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl) GetEntityData() *types.CommonEntityData {
    rpl.EntityData.YFilter = rpl.YFilter
    rpl.EntityData.YangName = "rpl"
    rpl.EntityData.BundleName = "cisco_ios_xr"
    rpl.EntityData.ParentYangName = "erp-instance"
    rpl.EntityData.SegmentPath = "rpl"
    rpl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpl.EntityData.Children = types.NewOrderedMap()
    rpl.EntityData.Leafs = types.NewOrderedMap()
    rpl.EntityData.Leafs.Append("port", types.YLeaf{"Port", rpl.Port})
    rpl.EntityData.Leafs.Append("role", types.YLeaf{"Role", rpl.Role})

    rpl.EntityData.YListKeys = []string {}

    return &(rpl.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps
// Automatic protection switching
type L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port0 APS channel in the format of InterfaceName. The type is string.
    Port0 interface{}

    // Enable automatic protection switching. The type is interface{}.
    Enable interface{}

    // Automatic protection switching level. The type is interface{} with range:
    // 0..7.
    Level interface{}

    // APS channel for ERP port1.
    Port1 L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1
}

func (aps *L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps) GetEntityData() *types.CommonEntityData {
    aps.EntityData.YFilter = aps.YFilter
    aps.EntityData.YangName = "aps"
    aps.EntityData.BundleName = "cisco_ios_xr"
    aps.EntityData.ParentYangName = "erp-instance"
    aps.EntityData.SegmentPath = "aps"
    aps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aps.EntityData.Children = types.NewOrderedMap()
    aps.EntityData.Children.Append("port1", types.YChild{"Port1", &aps.Port1})
    aps.EntityData.Leafs = types.NewOrderedMap()
    aps.EntityData.Leafs.Append("port0", types.YLeaf{"Port0", aps.Port0})
    aps.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", aps.Enable})
    aps.EntityData.Leafs.Append("level", types.YLeaf{"Level", aps.Level})

    aps.EntityData.YListKeys = []string {}

    return &(aps.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1
// APS channel for ERP port1
type L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port1 APS type. The type is Erpaps.
    ApsType interface{}

    // Port1 APS channel in the format of InterfaceName, BDName or XconnectName.
    // The type is string.
    ApsChannel interface{}
}

func (port1 *L2vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1) GetEntityData() *types.CommonEntityData {
    port1.EntityData.YFilter = port1.YFilter
    port1.EntityData.YangName = "port1"
    port1.EntityData.BundleName = "cisco_ios_xr"
    port1.EntityData.ParentYangName = "aps"
    port1.EntityData.SegmentPath = "port1"
    port1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port1.EntityData.Children = types.NewOrderedMap()
    port1.EntityData.Leafs = types.NewOrderedMap()
    port1.EntityData.Leafs.Append("aps-type", types.YLeaf{"ApsType", port1.ApsType})
    port1.EntityData.Leafs.Append("aps-channel", types.YLeaf{"ApsChannel", port1.ApsChannel})

    port1.EntityData.YListKeys = []string {}

    return &(port1.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s
// Ethernet ring protection port0
type L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet ring protection port1. The type is slice of
    // L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1.
    ErpPort1 []*L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1
}

func (erpPort1s *L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s) GetEntityData() *types.CommonEntityData {
    erpPort1s.EntityData.YFilter = erpPort1s.YFilter
    erpPort1s.EntityData.YangName = "erp-port1s"
    erpPort1s.EntityData.BundleName = "cisco_ios_xr"
    erpPort1s.EntityData.ParentYangName = "g8032-ring"
    erpPort1s.EntityData.SegmentPath = "erp-port1s"
    erpPort1s.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort1s.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort1s.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort1s.EntityData.Children = types.NewOrderedMap()
    erpPort1s.EntityData.Children.Append("erp-port1", types.YChild{"ErpPort1", nil})
    for i := range erpPort1s.ErpPort1 {
        erpPort1s.EntityData.Children.Append(types.GetSegmentPath(erpPort1s.ErpPort1[i]), types.YChild{"ErpPort1", erpPort1s.ErpPort1[i]})
    }
    erpPort1s.EntityData.Leafs = types.NewOrderedMap()

    erpPort1s.EntityData.YListKeys = []string {}

    return &(erpPort1s.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1
// Ethernet ring protection port1
type L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port1 type. The type is ErpPort.
    ErpPortType interface{}

    // none or virtual.
    NoneOrVirtual L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_NoneOrVirtual

    // interface. The type is slice of
    // L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_Interface.
    Interface []*L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_Interface
}

func (erpPort1 *L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1) GetEntityData() *types.CommonEntityData {
    erpPort1.EntityData.YFilter = erpPort1.YFilter
    erpPort1.EntityData.YangName = "erp-port1"
    erpPort1.EntityData.BundleName = "cisco_ios_xr"
    erpPort1.EntityData.ParentYangName = "erp-port1s"
    erpPort1.EntityData.SegmentPath = "erp-port1" + types.AddKeyToken(erpPort1.ErpPortType, "erp-port-type")
    erpPort1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort1.EntityData.Children = types.NewOrderedMap()
    erpPort1.EntityData.Children.Append("none-or-virtual", types.YChild{"NoneOrVirtual", &erpPort1.NoneOrVirtual})
    erpPort1.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range erpPort1.Interface {
        erpPort1.EntityData.Children.Append(types.GetSegmentPath(erpPort1.Interface[i]), types.YChild{"Interface", erpPort1.Interface[i]})
    }
    erpPort1.EntityData.Leafs = types.NewOrderedMap()
    erpPort1.EntityData.Leafs.Append("erp-port-type", types.YLeaf{"ErpPortType", erpPort1.ErpPortType})

    erpPort1.EntityData.YListKeys = []string {"ErpPortType"}

    return &(erpPort1.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_NoneOrVirtual
// none or virtual
// This type is a presence type.
type L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_NoneOrVirtual struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Ethernet ring protection port1 monitor. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Monitor interface{}
}

func (noneOrVirtual *L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_NoneOrVirtual) GetEntityData() *types.CommonEntityData {
    noneOrVirtual.EntityData.YFilter = noneOrVirtual.YFilter
    noneOrVirtual.EntityData.YangName = "none-or-virtual"
    noneOrVirtual.EntityData.BundleName = "cisco_ios_xr"
    noneOrVirtual.EntityData.ParentYangName = "erp-port1"
    noneOrVirtual.EntityData.SegmentPath = "none-or-virtual"
    noneOrVirtual.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    noneOrVirtual.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    noneOrVirtual.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    noneOrVirtual.EntityData.Children = types.NewOrderedMap()
    noneOrVirtual.EntityData.Leafs = types.NewOrderedMap()
    noneOrVirtual.EntityData.Leafs.Append("monitor", types.YLeaf{"Monitor", noneOrVirtual.Monitor})

    noneOrVirtual.EntityData.YListKeys = []string {}

    return &(noneOrVirtual.EntityData)
}

// L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_Interface
// interface
type L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port1 interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Ethernet ring protection port1 monitor. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Monitor interface{}
}

func (self *L2vpn_Database_G8032Rings_G8032Ring_ErpPort1s_ErpPort1_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "erp-port1"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("monitor", types.YLeaf{"Monitor", self.Monitor})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// L2vpn_Database_XconnectGroups
// List of xconnect groups
type L2vpn_Database_XconnectGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Xconnect group. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup.
    XconnectGroup []*L2vpn_Database_XconnectGroups_XconnectGroup
}

func (xconnectGroups *L2vpn_Database_XconnectGroups) GetEntityData() *types.CommonEntityData {
    xconnectGroups.EntityData.YFilter = xconnectGroups.YFilter
    xconnectGroups.EntityData.YangName = "xconnect-groups"
    xconnectGroups.EntityData.BundleName = "cisco_ios_xr"
    xconnectGroups.EntityData.ParentYangName = "database"
    xconnectGroups.EntityData.SegmentPath = "xconnect-groups"
    xconnectGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnectGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnectGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnectGroups.EntityData.Children = types.NewOrderedMap()
    xconnectGroups.EntityData.Children.Append("xconnect-group", types.YChild{"XconnectGroup", nil})
    for i := range xconnectGroups.XconnectGroup {
        xconnectGroups.EntityData.Children.Append(types.GetSegmentPath(xconnectGroups.XconnectGroup[i]), types.YChild{"XconnectGroup", xconnectGroups.XconnectGroup[i]})
    }
    xconnectGroups.EntityData.Leafs = types.NewOrderedMap()

    xconnectGroups.EntityData.YListKeys = []string {}

    return &(xconnectGroups.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup
// Xconnect group
type L2vpn_Database_XconnectGroups_XconnectGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the xconnect group. The type is string
    // with length: 1..32.
    Name interface{}

    // List of point to point xconnects.
    P2pXconnects L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects

    // List of multi point to multi point xconnects.
    Mp2mpXconnects L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects
}

func (xconnectGroup *L2vpn_Database_XconnectGroups_XconnectGroup) GetEntityData() *types.CommonEntityData {
    xconnectGroup.EntityData.YFilter = xconnectGroup.YFilter
    xconnectGroup.EntityData.YangName = "xconnect-group"
    xconnectGroup.EntityData.BundleName = "cisco_ios_xr"
    xconnectGroup.EntityData.ParentYangName = "xconnect-groups"
    xconnectGroup.EntityData.SegmentPath = "xconnect-group" + types.AddKeyToken(xconnectGroup.Name, "name")
    xconnectGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnectGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnectGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnectGroup.EntityData.Children = types.NewOrderedMap()
    xconnectGroup.EntityData.Children.Append("p2p-xconnects", types.YChild{"P2pXconnects", &xconnectGroup.P2pXconnects})
    xconnectGroup.EntityData.Children.Append("mp2mp-xconnects", types.YChild{"Mp2mpXconnects", &xconnectGroup.Mp2mpXconnects})
    xconnectGroup.EntityData.Leafs = types.NewOrderedMap()
    xconnectGroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", xconnectGroup.Name})

    xconnectGroup.EntityData.YListKeys = []string {"Name"}

    return &(xconnectGroup.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects
// List of point to point xconnects
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Point to point xconnect. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect.
    P2pXconnect []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect
}

func (p2pXconnects *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects) GetEntityData() *types.CommonEntityData {
    p2pXconnects.EntityData.YFilter = p2pXconnects.YFilter
    p2pXconnects.EntityData.YangName = "p2p-xconnects"
    p2pXconnects.EntityData.BundleName = "cisco_ios_xr"
    p2pXconnects.EntityData.ParentYangName = "xconnect-group"
    p2pXconnects.EntityData.SegmentPath = "p2p-xconnects"
    p2pXconnects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2pXconnects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2pXconnects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2pXconnects.EntityData.Children = types.NewOrderedMap()
    p2pXconnects.EntityData.Children.Append("p2p-xconnect", types.YChild{"P2pXconnect", nil})
    for i := range p2pXconnects.P2pXconnect {
        p2pXconnects.EntityData.Children.Append(types.GetSegmentPath(p2pXconnects.P2pXconnect[i]), types.YChild{"P2pXconnect", p2pXconnects.P2pXconnect[i]})
    }
    p2pXconnects.EntityData.Leafs = types.NewOrderedMap()

    p2pXconnects.EntityData.YListKeys = []string {}

    return &(p2pXconnects.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect
// Point to point xconnect
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the point to point xconnect. The type is
    // string with length: 1..38.
    Name interface{}

    // cross connect description Name. The type is string with length: 1..64.
    P2pDescription interface{}

    // Interworking. The type is Interworking.
    Interworking interface{}

    // List of backup attachment circuits.
    BackupAttachmentCircuits L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits

    // List of EVPN Services.
    PseudowireEvpns L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns

    // List of pseudowires.
    Pseudowires L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires

    // List of Monitor session segments.
    MonitorSessions L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions

    // List of pseudowire-routed.
    PseudowireRouteds L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds

    // List of attachment circuits.
    AttachmentCircuits L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits
}

func (p2pXconnect *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect) GetEntityData() *types.CommonEntityData {
    p2pXconnect.EntityData.YFilter = p2pXconnect.YFilter
    p2pXconnect.EntityData.YangName = "p2p-xconnect"
    p2pXconnect.EntityData.BundleName = "cisco_ios_xr"
    p2pXconnect.EntityData.ParentYangName = "p2p-xconnects"
    p2pXconnect.EntityData.SegmentPath = "p2p-xconnect" + types.AddKeyToken(p2pXconnect.Name, "name")
    p2pXconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2pXconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2pXconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2pXconnect.EntityData.Children = types.NewOrderedMap()
    p2pXconnect.EntityData.Children.Append("backup-attachment-circuits", types.YChild{"BackupAttachmentCircuits", &p2pXconnect.BackupAttachmentCircuits})
    p2pXconnect.EntityData.Children.Append("pseudowire-evpns", types.YChild{"PseudowireEvpns", &p2pXconnect.PseudowireEvpns})
    p2pXconnect.EntityData.Children.Append("pseudowires", types.YChild{"Pseudowires", &p2pXconnect.Pseudowires})
    p2pXconnect.EntityData.Children.Append("monitor-sessions", types.YChild{"MonitorSessions", &p2pXconnect.MonitorSessions})
    p2pXconnect.EntityData.Children.Append("pseudowire-routeds", types.YChild{"PseudowireRouteds", &p2pXconnect.PseudowireRouteds})
    p2pXconnect.EntityData.Children.Append("attachment-circuits", types.YChild{"AttachmentCircuits", &p2pXconnect.AttachmentCircuits})
    p2pXconnect.EntityData.Leafs = types.NewOrderedMap()
    p2pXconnect.EntityData.Leafs.Append("name", types.YLeaf{"Name", p2pXconnect.Name})
    p2pXconnect.EntityData.Leafs.Append("p2p-description", types.YLeaf{"P2pDescription", p2pXconnect.P2pDescription})
    p2pXconnect.EntityData.Leafs.Append("interworking", types.YLeaf{"Interworking", p2pXconnect.Interworking})

    p2pXconnect.EntityData.YListKeys = []string {"Name"}

    return &(p2pXconnect.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits
// List of backup attachment circuits
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup attachment circuit. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit.
    BackupAttachmentCircuit []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit
}

func (backupAttachmentCircuits *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    backupAttachmentCircuits.EntityData.YFilter = backupAttachmentCircuits.YFilter
    backupAttachmentCircuits.EntityData.YangName = "backup-attachment-circuits"
    backupAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    backupAttachmentCircuits.EntityData.ParentYangName = "p2p-xconnect"
    backupAttachmentCircuits.EntityData.SegmentPath = "backup-attachment-circuits"
    backupAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupAttachmentCircuits.EntityData.Children = types.NewOrderedMap()
    backupAttachmentCircuits.EntityData.Children.Append("backup-attachment-circuit", types.YChild{"BackupAttachmentCircuit", nil})
    for i := range backupAttachmentCircuits.BackupAttachmentCircuit {
        backupAttachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(backupAttachmentCircuits.BackupAttachmentCircuit[i]), types.YChild{"BackupAttachmentCircuit", backupAttachmentCircuits.BackupAttachmentCircuit[i]})
    }
    backupAttachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    backupAttachmentCircuits.EntityData.YListKeys = []string {}

    return &(backupAttachmentCircuits.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit
// Backup attachment circuit
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (backupAttachmentCircuit *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    backupAttachmentCircuit.EntityData.YFilter = backupAttachmentCircuit.YFilter
    backupAttachmentCircuit.EntityData.YangName = "backup-attachment-circuit"
    backupAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    backupAttachmentCircuit.EntityData.ParentYangName = "backup-attachment-circuits"
    backupAttachmentCircuit.EntityData.SegmentPath = "backup-attachment-circuit" + types.AddKeyToken(backupAttachmentCircuit.InterfaceName, "interface-name")
    backupAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupAttachmentCircuit.EntityData.Children = types.NewOrderedMap()
    backupAttachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    backupAttachmentCircuit.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", backupAttachmentCircuit.InterfaceName})

    backupAttachmentCircuit.EntityData.YListKeys = []string {"InterfaceName"}

    return &(backupAttachmentCircuit.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns
// List of EVPN Services
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN P2P Service Configuration. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns_PseudowireEvpn.
    PseudowireEvpn []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns_PseudowireEvpn
}

func (pseudowireEvpns *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns) GetEntityData() *types.CommonEntityData {
    pseudowireEvpns.EntityData.YFilter = pseudowireEvpns.YFilter
    pseudowireEvpns.EntityData.YangName = "pseudowire-evpns"
    pseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEvpns.EntityData.ParentYangName = "p2p-xconnect"
    pseudowireEvpns.EntityData.SegmentPath = "pseudowire-evpns"
    pseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEvpns.EntityData.Children = types.NewOrderedMap()
    pseudowireEvpns.EntityData.Children.Append("pseudowire-evpn", types.YChild{"PseudowireEvpn", nil})
    for i := range pseudowireEvpns.PseudowireEvpn {
        pseudowireEvpns.EntityData.Children.Append(types.GetSegmentPath(pseudowireEvpns.PseudowireEvpn[i]), types.YChild{"PseudowireEvpn", pseudowireEvpns.PseudowireEvpn[i]})
    }
    pseudowireEvpns.EntityData.Leafs = types.NewOrderedMap()

    pseudowireEvpns.EntityData.YListKeys = []string {}

    return &(pseudowireEvpns.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns_PseudowireEvpn
// EVPN P2P Service Configuration
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns_PseudowireEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // This attribute is a key. Remote AC ID. The type is interface{} with range:
    // 1..16777215.
    RemoteAcid interface{}

    // This attribute is a key. Source AC ID. The type is interface{} with range:
    // 1..16777215.
    SourceAcid interface{}

    // Name of the pseudowire class. The type is string with length: 1..32.
    Class interface{}
}

func (pseudowireEvpn *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireEvpns_PseudowireEvpn) GetEntityData() *types.CommonEntityData {
    pseudowireEvpn.EntityData.YFilter = pseudowireEvpn.YFilter
    pseudowireEvpn.EntityData.YangName = "pseudowire-evpn"
    pseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEvpn.EntityData.ParentYangName = "pseudowire-evpns"
    pseudowireEvpn.EntityData.SegmentPath = "pseudowire-evpn" + types.AddKeyToken(pseudowireEvpn.Eviid, "eviid") + types.AddKeyToken(pseudowireEvpn.RemoteAcid, "remote-acid") + types.AddKeyToken(pseudowireEvpn.SourceAcid, "source-acid")
    pseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEvpn.EntityData.Children = types.NewOrderedMap()
    pseudowireEvpn.EntityData.Leafs = types.NewOrderedMap()
    pseudowireEvpn.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", pseudowireEvpn.Eviid})
    pseudowireEvpn.EntityData.Leafs.Append("remote-acid", types.YLeaf{"RemoteAcid", pseudowireEvpn.RemoteAcid})
    pseudowireEvpn.EntityData.Leafs.Append("source-acid", types.YLeaf{"SourceAcid", pseudowireEvpn.SourceAcid})
    pseudowireEvpn.EntityData.Leafs.Append("class", types.YLeaf{"Class", pseudowireEvpn.Class})

    pseudowireEvpn.EntityData.YListKeys = []string {"Eviid", "RemoteAcid", "SourceAcid"}

    return &(pseudowireEvpn.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires
// List of pseudowires
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire.
    Pseudowire []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire
}

func (pseudowires *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires) GetEntityData() *types.CommonEntityData {
    pseudowires.EntityData.YFilter = pseudowires.YFilter
    pseudowires.EntityData.YangName = "pseudowires"
    pseudowires.EntityData.BundleName = "cisco_ios_xr"
    pseudowires.EntityData.ParentYangName = "p2p-xconnect"
    pseudowires.EntityData.SegmentPath = "pseudowires"
    pseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowires.EntityData.Children = types.NewOrderedMap()
    pseudowires.EntityData.Children.Append("pseudowire", types.YChild{"Pseudowire", nil})
    for i := range pseudowires.Pseudowire {
        pseudowires.EntityData.Children.Append(types.GetSegmentPath(pseudowires.Pseudowire[i]), types.YChild{"Pseudowire", pseudowires.Pseudowire[i]})
    }
    pseudowires.EntityData.Leafs = types.NewOrderedMap()

    pseudowires.EntityData.YListKeys = []string {}

    return &(pseudowires.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire
// Pseudowire configuration
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // keys: neighbor. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor.
    Neighbor []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor

    // keys: pseudowire-address. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress.
    PseudowireAddress []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress
}

func (pseudowire *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire) GetEntityData() *types.CommonEntityData {
    pseudowire.EntityData.YFilter = pseudowire.YFilter
    pseudowire.EntityData.YangName = "pseudowire"
    pseudowire.EntityData.BundleName = "cisco_ios_xr"
    pseudowire.EntityData.ParentYangName = "pseudowires"
    pseudowire.EntityData.SegmentPath = "pseudowire" + types.AddKeyToken(pseudowire.PseudowireId, "pseudowire-id")
    pseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowire.EntityData.Children = types.NewOrderedMap()
    pseudowire.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range pseudowire.Neighbor {
        pseudowire.EntityData.Children.Append(types.GetSegmentPath(pseudowire.Neighbor[i]), types.YChild{"Neighbor", pseudowire.Neighbor[i]})
    }
    pseudowire.EntityData.Children.Append("pseudowire-address", types.YChild{"PseudowireAddress", nil})
    for i := range pseudowire.PseudowireAddress {
        pseudowire.EntityData.Children.Append(types.GetSegmentPath(pseudowire.PseudowireAddress[i]), types.YChild{"PseudowireAddress", pseudowire.PseudowireAddress[i]})
    }
    pseudowire.EntityData.Leafs = types.NewOrderedMap()
    pseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", pseudowire.PseudowireId})

    pseudowire.EntityData.YListKeys = []string {"PseudowireId"}

    return &(pseudowire.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor
// keys: neighbor
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Pseudowire IPv4 address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // Tag Impose vlan tagged mode. The type is interface{} with range: 1..4094.
    TagImpose interface{}

    // Name of the pseudowire class. The type is string with length: 1..32.
    Class interface{}

    // Value of the Pseudowire source address. Must be IPv6 only. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Pseudowire Bandwidth. The type is interface{} with range: 0..4294967295.
    Bandwidth interface{}

    // MPLS static labels.
    MplsStaticLabels L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels

    // List of pseudowires.
    BackupPseudowires L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires

    // L2TP Static Attributes.
    L2tpStaticAttributes L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes

    // Pseudowire L2TPv3 static configuration.
    L2tpStatic L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStatic
}

func (neighbor *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "pseudowire"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Neighbor, "neighbor")
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("mpls-static-labels", types.YChild{"MplsStaticLabels", &neighbor.MplsStaticLabels})
    neighbor.EntityData.Children.Append("backup-pseudowires", types.YChild{"BackupPseudowires", &neighbor.BackupPseudowires})
    neighbor.EntityData.Children.Append("l2tp-static-attributes", types.YChild{"L2tpStaticAttributes", &neighbor.L2tpStaticAttributes})
    neighbor.EntityData.Children.Append("l2tp-static", types.YChild{"L2tpStatic", &neighbor.L2tpStatic})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", neighbor.Neighbor})
    neighbor.EntityData.Leafs.Append("tag-impose", types.YLeaf{"TagImpose", neighbor.TagImpose})
    neighbor.EntityData.Leafs.Append("class", types.YLeaf{"Class", neighbor.Class})
    neighbor.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", neighbor.SourceAddress})
    neighbor.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", neighbor.Bandwidth})

    neighbor.EntityData.YListKeys = []string {"Neighbor"}

    return &(neighbor.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels
// MPLS static labels
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (mplsStaticLabels *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels) GetEntityData() *types.CommonEntityData {
    mplsStaticLabels.EntityData.YFilter = mplsStaticLabels.YFilter
    mplsStaticLabels.EntityData.YangName = "mpls-static-labels"
    mplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    mplsStaticLabels.EntityData.ParentYangName = "neighbor"
    mplsStaticLabels.EntityData.SegmentPath = "mpls-static-labels"
    mplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    mplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    mplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", mplsStaticLabels.LocalStaticLabel})
    mplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", mplsStaticLabels.RemoteStaticLabel})

    mplsStaticLabels.EntityData.YListKeys = []string {}

    return &(mplsStaticLabels.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires
// List of pseudowires
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire for the cross connect. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire.
    BackupPseudowire []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire
}

func (backupPseudowires *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires) GetEntityData() *types.CommonEntityData {
    backupPseudowires.EntityData.YFilter = backupPseudowires.YFilter
    backupPseudowires.EntityData.YangName = "backup-pseudowires"
    backupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowires.EntityData.ParentYangName = "neighbor"
    backupPseudowires.EntityData.SegmentPath = "backup-pseudowires"
    backupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowires.EntityData.Children = types.NewOrderedMap()
    backupPseudowires.EntityData.Children.Append("backup-pseudowire", types.YChild{"BackupPseudowire", nil})
    for i := range backupPseudowires.BackupPseudowire {
        backupPseudowires.EntityData.Children.Append(types.GetSegmentPath(backupPseudowires.BackupPseudowire[i]), types.YChild{"BackupPseudowire", backupPseudowires.BackupPseudowire[i]})
    }
    backupPseudowires.EntityData.Leafs = types.NewOrderedMap()

    backupPseudowires.EntityData.YListKeys = []string {}

    return &(backupPseudowires.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire
// Backup pseudowire for the cross connect
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for the backup PW. The type is string with
    // length: 1..32.
    BackupPwClass interface{}

    // MPLS static labels.
    BackupMplsStaticLabels L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
}

func (backupPseudowire *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire) GetEntityData() *types.CommonEntityData {
    backupPseudowire.EntityData.YFilter = backupPseudowire.YFilter
    backupPseudowire.EntityData.YangName = "backup-pseudowire"
    backupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowire.EntityData.ParentYangName = "backup-pseudowires"
    backupPseudowire.EntityData.SegmentPath = "backup-pseudowire" + types.AddKeyToken(backupPseudowire.Neighbor, "neighbor") + types.AddKeyToken(backupPseudowire.PseudowireId, "pseudowire-id")
    backupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowire.EntityData.Children = types.NewOrderedMap()
    backupPseudowire.EntityData.Children.Append("backup-mpls-static-labels", types.YChild{"BackupMplsStaticLabels", &backupPseudowire.BackupMplsStaticLabels})
    backupPseudowire.EntityData.Leafs = types.NewOrderedMap()
    backupPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", backupPseudowire.Neighbor})
    backupPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", backupPseudowire.PseudowireId})
    backupPseudowire.EntityData.Leafs.Append("backup-pw-class", types.YLeaf{"BackupPwClass", backupPseudowire.BackupPwClass})

    backupPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(backupPseudowire.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
// MPLS static labels
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (backupMplsStaticLabels *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    backupMplsStaticLabels.EntityData.YFilter = backupMplsStaticLabels.YFilter
    backupMplsStaticLabels.EntityData.YangName = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    backupMplsStaticLabels.EntityData.ParentYangName = "backup-pseudowire"
    backupMplsStaticLabels.EntityData.SegmentPath = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupMplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    backupMplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    backupMplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", backupMplsStaticLabels.LocalStaticLabel})
    backupMplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", backupMplsStaticLabels.RemoteStaticLabel})

    backupMplsStaticLabels.EntityData.YListKeys = []string {}

    return &(backupMplsStaticLabels.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes
// L2TP Static Attributes
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP remote session ID. The type is interface{} with range: 1..65535.
    L2tpRemoteSessionId interface{}

    // L2TP local session ID. The type is interface{} with range: 1..65535.
    L2tpLocalSessionId interface{}

    // L2TP remote cookie.
    L2tpRemoteCookie L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpRemoteCookie

    // L2TP secondary local cookie.
    L2tpSecondaryLocalCookie L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpSecondaryLocalCookie

    // L2TP local cookie.
    L2tpLocalCookie L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpLocalCookie
}

func (l2tpStaticAttributes *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes) GetEntityData() *types.CommonEntityData {
    l2tpStaticAttributes.EntityData.YFilter = l2tpStaticAttributes.YFilter
    l2tpStaticAttributes.EntityData.YangName = "l2tp-static-attributes"
    l2tpStaticAttributes.EntityData.BundleName = "cisco_ios_xr"
    l2tpStaticAttributes.EntityData.ParentYangName = "neighbor"
    l2tpStaticAttributes.EntityData.SegmentPath = "l2tp-static-attributes"
    l2tpStaticAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpStaticAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpStaticAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpStaticAttributes.EntityData.Children = types.NewOrderedMap()
    l2tpStaticAttributes.EntityData.Children.Append("l2tp-remote-cookie", types.YChild{"L2tpRemoteCookie", &l2tpStaticAttributes.L2tpRemoteCookie})
    l2tpStaticAttributes.EntityData.Children.Append("l2tp-secondary-local-cookie", types.YChild{"L2tpSecondaryLocalCookie", &l2tpStaticAttributes.L2tpSecondaryLocalCookie})
    l2tpStaticAttributes.EntityData.Children.Append("l2tp-local-cookie", types.YChild{"L2tpLocalCookie", &l2tpStaticAttributes.L2tpLocalCookie})
    l2tpStaticAttributes.EntityData.Leafs = types.NewOrderedMap()
    l2tpStaticAttributes.EntityData.Leafs.Append("l2tp-remote-session-id", types.YLeaf{"L2tpRemoteSessionId", l2tpStaticAttributes.L2tpRemoteSessionId})
    l2tpStaticAttributes.EntityData.Leafs.Append("l2tp-local-session-id", types.YLeaf{"L2tpLocalSessionId", l2tpStaticAttributes.L2tpLocalSessionId})

    l2tpStaticAttributes.EntityData.YListKeys = []string {}

    return &(l2tpStaticAttributes.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpRemoteCookie
// L2TP remote cookie
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpRemoteCookie struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote cookie size. The type is L2tpCookieSize.
    Size interface{}

    // Lower remote cookie value. The type is interface{} with range:
    // 0..4294967295.
    LowerValue interface{}

    // Higher remote cookie value. The type is interface{} with range:
    // 0..4294967295.
    HigherValue interface{}
}

func (l2tpRemoteCookie *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpRemoteCookie) GetEntityData() *types.CommonEntityData {
    l2tpRemoteCookie.EntityData.YFilter = l2tpRemoteCookie.YFilter
    l2tpRemoteCookie.EntityData.YangName = "l2tp-remote-cookie"
    l2tpRemoteCookie.EntityData.BundleName = "cisco_ios_xr"
    l2tpRemoteCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2tpRemoteCookie.EntityData.SegmentPath = "l2tp-remote-cookie"
    l2tpRemoteCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpRemoteCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpRemoteCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpRemoteCookie.EntityData.Children = types.NewOrderedMap()
    l2tpRemoteCookie.EntityData.Leafs = types.NewOrderedMap()
    l2tpRemoteCookie.EntityData.Leafs.Append("size", types.YLeaf{"Size", l2tpRemoteCookie.Size})
    l2tpRemoteCookie.EntityData.Leafs.Append("lower-value", types.YLeaf{"LowerValue", l2tpRemoteCookie.LowerValue})
    l2tpRemoteCookie.EntityData.Leafs.Append("higher-value", types.YLeaf{"HigherValue", l2tpRemoteCookie.HigherValue})

    l2tpRemoteCookie.EntityData.YListKeys = []string {}

    return &(l2tpRemoteCookie.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpSecondaryLocalCookie
// L2TP secondary local cookie
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpSecondaryLocalCookie struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local cookie size. The type is L2tpCookieSize.
    Size interface{}

    // Lower local cookie value. The type is interface{} with range:
    // 0..4294967295.
    LowerValue interface{}

    // Higher local cookie value. The type is interface{} with range:
    // 0..4294967295.
    HigherValue interface{}
}

func (l2tpSecondaryLocalCookie *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpSecondaryLocalCookie) GetEntityData() *types.CommonEntityData {
    l2tpSecondaryLocalCookie.EntityData.YFilter = l2tpSecondaryLocalCookie.YFilter
    l2tpSecondaryLocalCookie.EntityData.YangName = "l2tp-secondary-local-cookie"
    l2tpSecondaryLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2tpSecondaryLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2tpSecondaryLocalCookie.EntityData.SegmentPath = "l2tp-secondary-local-cookie"
    l2tpSecondaryLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpSecondaryLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpSecondaryLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpSecondaryLocalCookie.EntityData.Children = types.NewOrderedMap()
    l2tpSecondaryLocalCookie.EntityData.Leafs = types.NewOrderedMap()
    l2tpSecondaryLocalCookie.EntityData.Leafs.Append("size", types.YLeaf{"Size", l2tpSecondaryLocalCookie.Size})
    l2tpSecondaryLocalCookie.EntityData.Leafs.Append("lower-value", types.YLeaf{"LowerValue", l2tpSecondaryLocalCookie.LowerValue})
    l2tpSecondaryLocalCookie.EntityData.Leafs.Append("higher-value", types.YLeaf{"HigherValue", l2tpSecondaryLocalCookie.HigherValue})

    l2tpSecondaryLocalCookie.EntityData.YListKeys = []string {}

    return &(l2tpSecondaryLocalCookie.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpLocalCookie
// L2TP local cookie
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpLocalCookie struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local cookie size. The type is L2tpCookieSize.
    Size interface{}

    // Lower local cookie value. The type is interface{} with range:
    // 0..4294967295.
    LowerValue interface{}

    // Higher local cookie value. The type is interface{} with range:
    // 0..4294967295.
    HigherValue interface{}
}

func (l2tpLocalCookie *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStaticAttributes_L2tpLocalCookie) GetEntityData() *types.CommonEntityData {
    l2tpLocalCookie.EntityData.YFilter = l2tpLocalCookie.YFilter
    l2tpLocalCookie.EntityData.YangName = "l2tp-local-cookie"
    l2tpLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2tpLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2tpLocalCookie.EntityData.SegmentPath = "l2tp-local-cookie"
    l2tpLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpLocalCookie.EntityData.Children = types.NewOrderedMap()
    l2tpLocalCookie.EntityData.Leafs = types.NewOrderedMap()
    l2tpLocalCookie.EntityData.Leafs.Append("size", types.YLeaf{"Size", l2tpLocalCookie.Size})
    l2tpLocalCookie.EntityData.Leafs.Append("lower-value", types.YLeaf{"LowerValue", l2tpLocalCookie.LowerValue})
    l2tpLocalCookie.EntityData.Leafs.Append("higher-value", types.YLeaf{"HigherValue", l2tpLocalCookie.HigherValue})

    l2tpLocalCookie.EntityData.YListKeys = []string {}

    return &(l2tpLocalCookie.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStatic
// Pseudowire L2TPv3 static configuration
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pseudowire L2TPv3 static configuration. The type is interface{}.
    Enable interface{}
}

func (l2tpStatic *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_Neighbor_L2tpStatic) GetEntityData() *types.CommonEntityData {
    l2tpStatic.EntityData.YFilter = l2tpStatic.YFilter
    l2tpStatic.EntityData.YangName = "l2tp-static"
    l2tpStatic.EntityData.BundleName = "cisco_ios_xr"
    l2tpStatic.EntityData.ParentYangName = "neighbor"
    l2tpStatic.EntityData.SegmentPath = "l2tp-static"
    l2tpStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpStatic.EntityData.Children = types.NewOrderedMap()
    l2tpStatic.EntityData.Leafs = types.NewOrderedMap()
    l2tpStatic.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", l2tpStatic.Enable})

    l2tpStatic.EntityData.YListKeys = []string {}

    return &(l2tpStatic.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress
// keys: pseudowire-address
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Pseudowire IPv6 address. A pseudowire can have
    // only one address: IPv4 or IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PseudowireAddress interface{}

    // Tag Impose vlan tagged mode. The type is interface{} with range: 1..4094.
    TagImpose interface{}

    // Name of the pseudowire class. The type is string with length: 1..32.
    Class interface{}

    // Value of the Pseudowire source address. Must be IPv6 only. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Pseudowire Bandwidth. The type is interface{} with range: 0..4294967295.
    Bandwidth interface{}

    // MPLS static labels.
    MplsStaticLabels L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels

    // List of pseudowires.
    BackupPseudowires L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires

    // L2TP Static Attributes.
    L2tpStaticAttributes L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes

    // Pseudowire L2TPv3 static configuration.
    L2tpStatic L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStatic
}

func (pseudowireAddress *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress) GetEntityData() *types.CommonEntityData {
    pseudowireAddress.EntityData.YFilter = pseudowireAddress.YFilter
    pseudowireAddress.EntityData.YangName = "pseudowire-address"
    pseudowireAddress.EntityData.BundleName = "cisco_ios_xr"
    pseudowireAddress.EntityData.ParentYangName = "pseudowire"
    pseudowireAddress.EntityData.SegmentPath = "pseudowire-address" + types.AddKeyToken(pseudowireAddress.PseudowireAddress, "pseudowire-address")
    pseudowireAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireAddress.EntityData.Children = types.NewOrderedMap()
    pseudowireAddress.EntityData.Children.Append("mpls-static-labels", types.YChild{"MplsStaticLabels", &pseudowireAddress.MplsStaticLabels})
    pseudowireAddress.EntityData.Children.Append("backup-pseudowires", types.YChild{"BackupPseudowires", &pseudowireAddress.BackupPseudowires})
    pseudowireAddress.EntityData.Children.Append("l2tp-static-attributes", types.YChild{"L2tpStaticAttributes", &pseudowireAddress.L2tpStaticAttributes})
    pseudowireAddress.EntityData.Children.Append("l2tp-static", types.YChild{"L2tpStatic", &pseudowireAddress.L2tpStatic})
    pseudowireAddress.EntityData.Leafs = types.NewOrderedMap()
    pseudowireAddress.EntityData.Leafs.Append("pseudowire-address", types.YLeaf{"PseudowireAddress", pseudowireAddress.PseudowireAddress})
    pseudowireAddress.EntityData.Leafs.Append("tag-impose", types.YLeaf{"TagImpose", pseudowireAddress.TagImpose})
    pseudowireAddress.EntityData.Leafs.Append("class", types.YLeaf{"Class", pseudowireAddress.Class})
    pseudowireAddress.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", pseudowireAddress.SourceAddress})
    pseudowireAddress.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", pseudowireAddress.Bandwidth})

    pseudowireAddress.EntityData.YListKeys = []string {"PseudowireAddress"}

    return &(pseudowireAddress.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels
// MPLS static labels
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (mplsStaticLabels *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels) GetEntityData() *types.CommonEntityData {
    mplsStaticLabels.EntityData.YFilter = mplsStaticLabels.YFilter
    mplsStaticLabels.EntityData.YangName = "mpls-static-labels"
    mplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    mplsStaticLabels.EntityData.ParentYangName = "pseudowire-address"
    mplsStaticLabels.EntityData.SegmentPath = "mpls-static-labels"
    mplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    mplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    mplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", mplsStaticLabels.LocalStaticLabel})
    mplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", mplsStaticLabels.RemoteStaticLabel})

    mplsStaticLabels.EntityData.YListKeys = []string {}

    return &(mplsStaticLabels.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires
// List of pseudowires
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire for the cross connect. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire.
    BackupPseudowire []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire
}

func (backupPseudowires *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires) GetEntityData() *types.CommonEntityData {
    backupPseudowires.EntityData.YFilter = backupPseudowires.YFilter
    backupPseudowires.EntityData.YangName = "backup-pseudowires"
    backupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowires.EntityData.ParentYangName = "pseudowire-address"
    backupPseudowires.EntityData.SegmentPath = "backup-pseudowires"
    backupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowires.EntityData.Children = types.NewOrderedMap()
    backupPseudowires.EntityData.Children.Append("backup-pseudowire", types.YChild{"BackupPseudowire", nil})
    for i := range backupPseudowires.BackupPseudowire {
        backupPseudowires.EntityData.Children.Append(types.GetSegmentPath(backupPseudowires.BackupPseudowire[i]), types.YChild{"BackupPseudowire", backupPseudowires.BackupPseudowire[i]})
    }
    backupPseudowires.EntityData.Leafs = types.NewOrderedMap()

    backupPseudowires.EntityData.YListKeys = []string {}

    return &(backupPseudowires.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire
// Backup pseudowire for the cross connect
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for the backup PW. The type is string with
    // length: 1..32.
    BackupPwClass interface{}

    // MPLS static labels.
    BackupMplsStaticLabels L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
}

func (backupPseudowire *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire) GetEntityData() *types.CommonEntityData {
    backupPseudowire.EntityData.YFilter = backupPseudowire.YFilter
    backupPseudowire.EntityData.YangName = "backup-pseudowire"
    backupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowire.EntityData.ParentYangName = "backup-pseudowires"
    backupPseudowire.EntityData.SegmentPath = "backup-pseudowire" + types.AddKeyToken(backupPseudowire.Neighbor, "neighbor") + types.AddKeyToken(backupPseudowire.PseudowireId, "pseudowire-id")
    backupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowire.EntityData.Children = types.NewOrderedMap()
    backupPseudowire.EntityData.Children.Append("backup-mpls-static-labels", types.YChild{"BackupMplsStaticLabels", &backupPseudowire.BackupMplsStaticLabels})
    backupPseudowire.EntityData.Leafs = types.NewOrderedMap()
    backupPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", backupPseudowire.Neighbor})
    backupPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", backupPseudowire.PseudowireId})
    backupPseudowire.EntityData.Leafs.Append("backup-pw-class", types.YLeaf{"BackupPwClass", backupPseudowire.BackupPwClass})

    backupPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(backupPseudowire.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
// MPLS static labels
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (backupMplsStaticLabels *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    backupMplsStaticLabels.EntityData.YFilter = backupMplsStaticLabels.YFilter
    backupMplsStaticLabels.EntityData.YangName = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    backupMplsStaticLabels.EntityData.ParentYangName = "backup-pseudowire"
    backupMplsStaticLabels.EntityData.SegmentPath = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupMplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    backupMplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    backupMplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", backupMplsStaticLabels.LocalStaticLabel})
    backupMplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", backupMplsStaticLabels.RemoteStaticLabel})

    backupMplsStaticLabels.EntityData.YListKeys = []string {}

    return &(backupMplsStaticLabels.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes
// L2TP Static Attributes
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP remote session ID. The type is interface{} with range: 1..65535.
    L2tpRemoteSessionId interface{}

    // L2TP local session ID. The type is interface{} with range: 1..65535.
    L2tpLocalSessionId interface{}

    // L2TP remote cookie.
    L2tpRemoteCookie L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpRemoteCookie

    // L2TP secondary local cookie.
    L2tpSecondaryLocalCookie L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpSecondaryLocalCookie

    // L2TP local cookie.
    L2tpLocalCookie L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpLocalCookie
}

func (l2tpStaticAttributes *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes) GetEntityData() *types.CommonEntityData {
    l2tpStaticAttributes.EntityData.YFilter = l2tpStaticAttributes.YFilter
    l2tpStaticAttributes.EntityData.YangName = "l2tp-static-attributes"
    l2tpStaticAttributes.EntityData.BundleName = "cisco_ios_xr"
    l2tpStaticAttributes.EntityData.ParentYangName = "pseudowire-address"
    l2tpStaticAttributes.EntityData.SegmentPath = "l2tp-static-attributes"
    l2tpStaticAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpStaticAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpStaticAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpStaticAttributes.EntityData.Children = types.NewOrderedMap()
    l2tpStaticAttributes.EntityData.Children.Append("l2tp-remote-cookie", types.YChild{"L2tpRemoteCookie", &l2tpStaticAttributes.L2tpRemoteCookie})
    l2tpStaticAttributes.EntityData.Children.Append("l2tp-secondary-local-cookie", types.YChild{"L2tpSecondaryLocalCookie", &l2tpStaticAttributes.L2tpSecondaryLocalCookie})
    l2tpStaticAttributes.EntityData.Children.Append("l2tp-local-cookie", types.YChild{"L2tpLocalCookie", &l2tpStaticAttributes.L2tpLocalCookie})
    l2tpStaticAttributes.EntityData.Leafs = types.NewOrderedMap()
    l2tpStaticAttributes.EntityData.Leafs.Append("l2tp-remote-session-id", types.YLeaf{"L2tpRemoteSessionId", l2tpStaticAttributes.L2tpRemoteSessionId})
    l2tpStaticAttributes.EntityData.Leafs.Append("l2tp-local-session-id", types.YLeaf{"L2tpLocalSessionId", l2tpStaticAttributes.L2tpLocalSessionId})

    l2tpStaticAttributes.EntityData.YListKeys = []string {}

    return &(l2tpStaticAttributes.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpRemoteCookie
// L2TP remote cookie
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpRemoteCookie struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote cookie size. The type is L2tpCookieSize.
    Size interface{}

    // Lower remote cookie value. The type is interface{} with range:
    // 0..4294967295.
    LowerValue interface{}

    // Higher remote cookie value. The type is interface{} with range:
    // 0..4294967295.
    HigherValue interface{}
}

func (l2tpRemoteCookie *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpRemoteCookie) GetEntityData() *types.CommonEntityData {
    l2tpRemoteCookie.EntityData.YFilter = l2tpRemoteCookie.YFilter
    l2tpRemoteCookie.EntityData.YangName = "l2tp-remote-cookie"
    l2tpRemoteCookie.EntityData.BundleName = "cisco_ios_xr"
    l2tpRemoteCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2tpRemoteCookie.EntityData.SegmentPath = "l2tp-remote-cookie"
    l2tpRemoteCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpRemoteCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpRemoteCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpRemoteCookie.EntityData.Children = types.NewOrderedMap()
    l2tpRemoteCookie.EntityData.Leafs = types.NewOrderedMap()
    l2tpRemoteCookie.EntityData.Leafs.Append("size", types.YLeaf{"Size", l2tpRemoteCookie.Size})
    l2tpRemoteCookie.EntityData.Leafs.Append("lower-value", types.YLeaf{"LowerValue", l2tpRemoteCookie.LowerValue})
    l2tpRemoteCookie.EntityData.Leafs.Append("higher-value", types.YLeaf{"HigherValue", l2tpRemoteCookie.HigherValue})

    l2tpRemoteCookie.EntityData.YListKeys = []string {}

    return &(l2tpRemoteCookie.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpSecondaryLocalCookie
// L2TP secondary local cookie
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpSecondaryLocalCookie struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local cookie size. The type is L2tpCookieSize.
    Size interface{}

    // Lower local cookie value. The type is interface{} with range:
    // 0..4294967295.
    LowerValue interface{}

    // Higher local cookie value. The type is interface{} with range:
    // 0..4294967295.
    HigherValue interface{}
}

func (l2tpSecondaryLocalCookie *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpSecondaryLocalCookie) GetEntityData() *types.CommonEntityData {
    l2tpSecondaryLocalCookie.EntityData.YFilter = l2tpSecondaryLocalCookie.YFilter
    l2tpSecondaryLocalCookie.EntityData.YangName = "l2tp-secondary-local-cookie"
    l2tpSecondaryLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2tpSecondaryLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2tpSecondaryLocalCookie.EntityData.SegmentPath = "l2tp-secondary-local-cookie"
    l2tpSecondaryLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpSecondaryLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpSecondaryLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpSecondaryLocalCookie.EntityData.Children = types.NewOrderedMap()
    l2tpSecondaryLocalCookie.EntityData.Leafs = types.NewOrderedMap()
    l2tpSecondaryLocalCookie.EntityData.Leafs.Append("size", types.YLeaf{"Size", l2tpSecondaryLocalCookie.Size})
    l2tpSecondaryLocalCookie.EntityData.Leafs.Append("lower-value", types.YLeaf{"LowerValue", l2tpSecondaryLocalCookie.LowerValue})
    l2tpSecondaryLocalCookie.EntityData.Leafs.Append("higher-value", types.YLeaf{"HigherValue", l2tpSecondaryLocalCookie.HigherValue})

    l2tpSecondaryLocalCookie.EntityData.YListKeys = []string {}

    return &(l2tpSecondaryLocalCookie.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpLocalCookie
// L2TP local cookie
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpLocalCookie struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local cookie size. The type is L2tpCookieSize.
    Size interface{}

    // Lower local cookie value. The type is interface{} with range:
    // 0..4294967295.
    LowerValue interface{}

    // Higher local cookie value. The type is interface{} with range:
    // 0..4294967295.
    HigherValue interface{}
}

func (l2tpLocalCookie *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStaticAttributes_L2tpLocalCookie) GetEntityData() *types.CommonEntityData {
    l2tpLocalCookie.EntityData.YFilter = l2tpLocalCookie.YFilter
    l2tpLocalCookie.EntityData.YangName = "l2tp-local-cookie"
    l2tpLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2tpLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2tpLocalCookie.EntityData.SegmentPath = "l2tp-local-cookie"
    l2tpLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpLocalCookie.EntityData.Children = types.NewOrderedMap()
    l2tpLocalCookie.EntityData.Leafs = types.NewOrderedMap()
    l2tpLocalCookie.EntityData.Leafs.Append("size", types.YLeaf{"Size", l2tpLocalCookie.Size})
    l2tpLocalCookie.EntityData.Leafs.Append("lower-value", types.YLeaf{"LowerValue", l2tpLocalCookie.LowerValue})
    l2tpLocalCookie.EntityData.Leafs.Append("higher-value", types.YLeaf{"HigherValue", l2tpLocalCookie.HigherValue})

    l2tpLocalCookie.EntityData.YListKeys = []string {}

    return &(l2tpLocalCookie.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStatic
// Pseudowire L2TPv3 static configuration
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pseudowire L2TPv3 static configuration. The type is interface{}.
    Enable interface{}
}

func (l2tpStatic *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2tpStatic) GetEntityData() *types.CommonEntityData {
    l2tpStatic.EntityData.YFilter = l2tpStatic.YFilter
    l2tpStatic.EntityData.YangName = "l2tp-static"
    l2tpStatic.EntityData.BundleName = "cisco_ios_xr"
    l2tpStatic.EntityData.ParentYangName = "pseudowire-address"
    l2tpStatic.EntityData.SegmentPath = "l2tp-static"
    l2tpStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpStatic.EntityData.Children = types.NewOrderedMap()
    l2tpStatic.EntityData.Leafs = types.NewOrderedMap()
    l2tpStatic.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", l2tpStatic.Enable})

    l2tpStatic.EntityData.YListKeys = []string {}

    return &(l2tpStatic.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions
// List of Monitor session segments
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor session segment. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions_MonitorSession.
    MonitorSession []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions_MonitorSession
}

func (monitorSessions *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions) GetEntityData() *types.CommonEntityData {
    monitorSessions.EntityData.YFilter = monitorSessions.YFilter
    monitorSessions.EntityData.YangName = "monitor-sessions"
    monitorSessions.EntityData.BundleName = "cisco_ios_xr"
    monitorSessions.EntityData.ParentYangName = "p2p-xconnect"
    monitorSessions.EntityData.SegmentPath = "monitor-sessions"
    monitorSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorSessions.EntityData.Children = types.NewOrderedMap()
    monitorSessions.EntityData.Children.Append("monitor-session", types.YChild{"MonitorSession", nil})
    for i := range monitorSessions.MonitorSession {
        monitorSessions.EntityData.Children.Append(types.GetSegmentPath(monitorSessions.MonitorSession[i]), types.YChild{"MonitorSession", monitorSessions.MonitorSession[i]})
    }
    monitorSessions.EntityData.Leafs = types.NewOrderedMap()

    monitorSessions.EntityData.YListKeys = []string {}

    return &(monitorSessions.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions_MonitorSession
// Monitor session segment
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions_MonitorSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the monitor session. The type is string
    // with length: 1..64.
    Name interface{}

    // Enable monitor session segment . The type is interface{}.
    Enable interface{}
}

func (monitorSession *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_MonitorSessions_MonitorSession) GetEntityData() *types.CommonEntityData {
    monitorSession.EntityData.YFilter = monitorSession.YFilter
    monitorSession.EntityData.YangName = "monitor-session"
    monitorSession.EntityData.BundleName = "cisco_ios_xr"
    monitorSession.EntityData.ParentYangName = "monitor-sessions"
    monitorSession.EntityData.SegmentPath = "monitor-session" + types.AddKeyToken(monitorSession.Name, "name")
    monitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorSession.EntityData.Children = types.NewOrderedMap()
    monitorSession.EntityData.Leafs = types.NewOrderedMap()
    monitorSession.EntityData.Leafs.Append("name", types.YLeaf{"Name", monitorSession.Name})
    monitorSession.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", monitorSession.Enable})

    monitorSession.EntityData.YListKeys = []string {"Name"}

    return &(monitorSession.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds
// List of pseudowire-routed
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds_PseudowireRouted.
    PseudowireRouted []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds_PseudowireRouted
}

func (pseudowireRouteds *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds) GetEntityData() *types.CommonEntityData {
    pseudowireRouteds.EntityData.YFilter = pseudowireRouteds.YFilter
    pseudowireRouteds.EntityData.YangName = "pseudowire-routeds"
    pseudowireRouteds.EntityData.BundleName = "cisco_ios_xr"
    pseudowireRouteds.EntityData.ParentYangName = "p2p-xconnect"
    pseudowireRouteds.EntityData.SegmentPath = "pseudowire-routeds"
    pseudowireRouteds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireRouteds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireRouteds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireRouteds.EntityData.Children = types.NewOrderedMap()
    pseudowireRouteds.EntityData.Children.Append("pseudowire-routed", types.YChild{"PseudowireRouted", nil})
    for i := range pseudowireRouteds.PseudowireRouted {
        pseudowireRouteds.EntityData.Children.Append(types.GetSegmentPath(pseudowireRouteds.PseudowireRouted[i]), types.YChild{"PseudowireRouted", pseudowireRouteds.PseudowireRouted[i]})
    }
    pseudowireRouteds.EntityData.Leafs = types.NewOrderedMap()

    pseudowireRouteds.EntityData.YListKeys = []string {}

    return &(pseudowireRouteds.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds_PseudowireRouted
// Pseudowire configuration
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds_PseudowireRouted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Target Global ID. The type is interface{} with
    // range: 1..4294967295.
    GlobalId interface{}

    // This attribute is a key. Target Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // This attribute is a key. Target AC ID. The type is interface{} with range:
    // 1..4294967295.
    Acid interface{}

    // This attribute is a key. Source AC ID. The type is interface{} with range:
    // 1..4294967295.
    Sacid interface{}

    // Tag Impose vlan tagged mode. The type is interface{} with range: 1..4094.
    TagImpose interface{}

    // Name of the pseudowire class. The type is string with length: 1..32.
    Class interface{}
}

func (pseudowireRouted *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_PseudowireRouteds_PseudowireRouted) GetEntityData() *types.CommonEntityData {
    pseudowireRouted.EntityData.YFilter = pseudowireRouted.YFilter
    pseudowireRouted.EntityData.YangName = "pseudowire-routed"
    pseudowireRouted.EntityData.BundleName = "cisco_ios_xr"
    pseudowireRouted.EntityData.ParentYangName = "pseudowire-routeds"
    pseudowireRouted.EntityData.SegmentPath = "pseudowire-routed" + types.AddKeyToken(pseudowireRouted.GlobalId, "global-id") + types.AddKeyToken(pseudowireRouted.Prefix, "prefix") + types.AddKeyToken(pseudowireRouted.Acid, "acid") + types.AddKeyToken(pseudowireRouted.Sacid, "sacid")
    pseudowireRouted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireRouted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireRouted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireRouted.EntityData.Children = types.NewOrderedMap()
    pseudowireRouted.EntityData.Leafs = types.NewOrderedMap()
    pseudowireRouted.EntityData.Leafs.Append("global-id", types.YLeaf{"GlobalId", pseudowireRouted.GlobalId})
    pseudowireRouted.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", pseudowireRouted.Prefix})
    pseudowireRouted.EntityData.Leafs.Append("acid", types.YLeaf{"Acid", pseudowireRouted.Acid})
    pseudowireRouted.EntityData.Leafs.Append("sacid", types.YLeaf{"Sacid", pseudowireRouted.Sacid})
    pseudowireRouted.EntityData.Leafs.Append("tag-impose", types.YLeaf{"TagImpose", pseudowireRouted.TagImpose})
    pseudowireRouted.EntityData.Leafs.Append("class", types.YLeaf{"Class", pseudowireRouted.Class})

    pseudowireRouted.EntityData.YListKeys = []string {"GlobalId", "Prefix", "Acid", "Sacid"}

    return &(pseudowireRouted.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits
// List of attachment circuits
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits_AttachmentCircuit.
    AttachmentCircuit []*L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits_AttachmentCircuit
}

func (attachmentCircuits *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits) GetEntityData() *types.CommonEntityData {
    attachmentCircuits.EntityData.YFilter = attachmentCircuits.YFilter
    attachmentCircuits.EntityData.YangName = "attachment-circuits"
    attachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    attachmentCircuits.EntityData.ParentYangName = "p2p-xconnect"
    attachmentCircuits.EntityData.SegmentPath = "attachment-circuits"
    attachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachmentCircuits.EntityData.Children = types.NewOrderedMap()
    attachmentCircuits.EntityData.Children.Append("attachment-circuit", types.YChild{"AttachmentCircuit", nil})
    for i := range attachmentCircuits.AttachmentCircuit {
        attachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(attachmentCircuits.AttachmentCircuit[i]), types.YChild{"AttachmentCircuit", attachmentCircuits.AttachmentCircuit[i]})
    }
    attachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    attachmentCircuits.EntityData.YListKeys = []string {}

    return &(attachmentCircuits.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits_AttachmentCircuit
// Attachment circuit interface
type L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits_AttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: [a-zA-Z0-9._/-]+.
    Name interface{}

    // Enable attachment circuit interface. The type is interface{}.
    Enable interface{}
}

func (attachmentCircuit *L2vpn_Database_XconnectGroups_XconnectGroup_P2pXconnects_P2pXconnect_AttachmentCircuits_AttachmentCircuit) GetEntityData() *types.CommonEntityData {
    attachmentCircuit.EntityData.YFilter = attachmentCircuit.YFilter
    attachmentCircuit.EntityData.YangName = "attachment-circuit"
    attachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    attachmentCircuit.EntityData.ParentYangName = "attachment-circuits"
    attachmentCircuit.EntityData.SegmentPath = "attachment-circuit" + types.AddKeyToken(attachmentCircuit.Name, "name")
    attachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachmentCircuit.EntityData.Children = types.NewOrderedMap()
    attachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    attachmentCircuit.EntityData.Leafs.Append("name", types.YLeaf{"Name", attachmentCircuit.Name})
    attachmentCircuit.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", attachmentCircuit.Enable})

    attachmentCircuit.EntityData.YListKeys = []string {"Name"}

    return &(attachmentCircuit.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects
// List of multi point to multi point xconnects
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multi point to multi point xconnect. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect.
    Mp2mpXconnect []*L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect
}

func (mp2mpXconnects *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects) GetEntityData() *types.CommonEntityData {
    mp2mpXconnects.EntityData.YFilter = mp2mpXconnects.YFilter
    mp2mpXconnects.EntityData.YangName = "mp2mp-xconnects"
    mp2mpXconnects.EntityData.BundleName = "cisco_ios_xr"
    mp2mpXconnects.EntityData.ParentYangName = "xconnect-group"
    mp2mpXconnects.EntityData.SegmentPath = "mp2mp-xconnects"
    mp2mpXconnects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpXconnects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpXconnects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpXconnects.EntityData.Children = types.NewOrderedMap()
    mp2mpXconnects.EntityData.Children.Append("mp2mp-xconnect", types.YChild{"Mp2mpXconnect", nil})
    for i := range mp2mpXconnects.Mp2mpXconnect {
        mp2mpXconnects.EntityData.Children.Append(types.GetSegmentPath(mp2mpXconnects.Mp2mpXconnect[i]), types.YChild{"Mp2mpXconnect", mp2mpXconnects.Mp2mpXconnect[i]})
    }
    mp2mpXconnects.EntityData.Leafs = types.NewOrderedMap()

    mp2mpXconnects.EntityData.YListKeys = []string {}

    return &(mp2mpXconnects.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect
// Multi point to multi point xconnect
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the multi point to multi point xconnect.
    // The type is string with length: 1..26.
    Name interface{}

    // Maximum transmission unit for this MP2MP VPWS instance. The type is
    // interface{} with range: 64..65535. Units are byte.
    Mp2mpmtu interface{}

    // Disable control word. The type is interface{}.
    Mp2mpControlWord interface{}

    // Configure Layer 2 Encapsulation. The type is L2Encapsulation.
    Mp2mpl2Encapsulation interface{}

    // Interworking. The type is Interworking.
    Mp2mpInterworking interface{}

    // shutdown this MP2MP VPWS instance. The type is interface{}.
    Mp2mpShutdown interface{}

    // VPN Identifier. The type is interface{} with range: 1..4294967295.
    Mp2mpvpnId interface{}

    // auto-discovery in this MP2MP.
    Mp2mpAutoDiscovery L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery
}

func (mp2mpXconnect *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect) GetEntityData() *types.CommonEntityData {
    mp2mpXconnect.EntityData.YFilter = mp2mpXconnect.YFilter
    mp2mpXconnect.EntityData.YangName = "mp2mp-xconnect"
    mp2mpXconnect.EntityData.BundleName = "cisco_ios_xr"
    mp2mpXconnect.EntityData.ParentYangName = "mp2mp-xconnects"
    mp2mpXconnect.EntityData.SegmentPath = "mp2mp-xconnect" + types.AddKeyToken(mp2mpXconnect.Name, "name")
    mp2mpXconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpXconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpXconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpXconnect.EntityData.Children = types.NewOrderedMap()
    mp2mpXconnect.EntityData.Children.Append("mp2mp-auto-discovery", types.YChild{"Mp2mpAutoDiscovery", &mp2mpXconnect.Mp2mpAutoDiscovery})
    mp2mpXconnect.EntityData.Leafs = types.NewOrderedMap()
    mp2mpXconnect.EntityData.Leafs.Append("name", types.YLeaf{"Name", mp2mpXconnect.Name})
    mp2mpXconnect.EntityData.Leafs.Append("mp2mpmtu", types.YLeaf{"Mp2mpmtu", mp2mpXconnect.Mp2mpmtu})
    mp2mpXconnect.EntityData.Leafs.Append("mp2mp-control-word", types.YLeaf{"Mp2mpControlWord", mp2mpXconnect.Mp2mpControlWord})
    mp2mpXconnect.EntityData.Leafs.Append("mp2mpl2-encapsulation", types.YLeaf{"Mp2mpl2Encapsulation", mp2mpXconnect.Mp2mpl2Encapsulation})
    mp2mpXconnect.EntityData.Leafs.Append("mp2mp-interworking", types.YLeaf{"Mp2mpInterworking", mp2mpXconnect.Mp2mpInterworking})
    mp2mpXconnect.EntityData.Leafs.Append("mp2mp-shutdown", types.YLeaf{"Mp2mpShutdown", mp2mpXconnect.Mp2mpShutdown})
    mp2mpXconnect.EntityData.Leafs.Append("mp2mpvpn-id", types.YLeaf{"Mp2mpvpnId", mp2mpXconnect.Mp2mpvpnId})

    mp2mpXconnect.EntityData.YListKeys = []string {"Name"}

    return &(mp2mpXconnect.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery
// auto-discovery in this MP2MP
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable auto-discovery. The type is interface{}.
    Enable interface{}

    // Route Distinguisher.
    RouteDistinguisher L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_RouteDistinguisher

    // Route policy.
    Mp2mpRoutePolicy L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRoutePolicy

    // Route Target.
    Mp2mpRouteTargets L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets

    // signaling protocol in this MP2MP.
    Mp2mpSignalingProtocol L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol
}

func (mp2mpAutoDiscovery *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    mp2mpAutoDiscovery.EntityData.YFilter = mp2mpAutoDiscovery.YFilter
    mp2mpAutoDiscovery.EntityData.YangName = "mp2mp-auto-discovery"
    mp2mpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    mp2mpAutoDiscovery.EntityData.ParentYangName = "mp2mp-xconnect"
    mp2mpAutoDiscovery.EntityData.SegmentPath = "mp2mp-auto-discovery"
    mp2mpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    mp2mpAutoDiscovery.EntityData.Children.Append("route-distinguisher", types.YChild{"RouteDistinguisher", &mp2mpAutoDiscovery.RouteDistinguisher})
    mp2mpAutoDiscovery.EntityData.Children.Append("mp2mp-route-policy", types.YChild{"Mp2mpRoutePolicy", &mp2mpAutoDiscovery.Mp2mpRoutePolicy})
    mp2mpAutoDiscovery.EntityData.Children.Append("mp2mp-route-targets", types.YChild{"Mp2mpRouteTargets", &mp2mpAutoDiscovery.Mp2mpRouteTargets})
    mp2mpAutoDiscovery.EntityData.Children.Append("mp2mp-signaling-protocol", types.YChild{"Mp2mpSignalingProtocol", &mp2mpAutoDiscovery.Mp2mpSignalingProtocol})
    mp2mpAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    mp2mpAutoDiscovery.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mp2mpAutoDiscovery.Enable})

    mp2mpAutoDiscovery.EntityData.YListKeys = []string {}

    return &(mp2mpAutoDiscovery.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_RouteDistinguisher
// Route Distinguisher
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_RouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router distinguisher type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (routeDistinguisher *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_RouteDistinguisher) GetEntityData() *types.CommonEntityData {
    routeDistinguisher.EntityData.YFilter = routeDistinguisher.YFilter
    routeDistinguisher.EntityData.YangName = "route-distinguisher"
    routeDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    routeDistinguisher.EntityData.ParentYangName = "mp2mp-auto-discovery"
    routeDistinguisher.EntityData.SegmentPath = "route-distinguisher"
    routeDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeDistinguisher.EntityData.Children = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeDistinguisher.Type})
    routeDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", routeDistinguisher.As})
    routeDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeDistinguisher.AsIndex})
    routeDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeDistinguisher.Address})
    routeDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", routeDistinguisher.AddrIndex})

    routeDistinguisher.EntityData.YListKeys = []string {}

    return &(routeDistinguisher.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRoutePolicy
// Route policy
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Export route policy. The type is string.
    Export interface{}
}

func (mp2mpRoutePolicy *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRoutePolicy) GetEntityData() *types.CommonEntityData {
    mp2mpRoutePolicy.EntityData.YFilter = mp2mpRoutePolicy.YFilter
    mp2mpRoutePolicy.EntityData.YangName = "mp2mp-route-policy"
    mp2mpRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    mp2mpRoutePolicy.EntityData.ParentYangName = "mp2mp-auto-discovery"
    mp2mpRoutePolicy.EntityData.SegmentPath = "mp2mp-route-policy"
    mp2mpRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpRoutePolicy.EntityData.Children = types.NewOrderedMap()
    mp2mpRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    mp2mpRoutePolicy.EntityData.Leafs.Append("export", types.YLeaf{"Export", mp2mpRoutePolicy.Export})

    mp2mpRoutePolicy.EntityData.YListKeys = []string {}

    return &(mp2mpRoutePolicy.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets
// Route Target
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget.
    Mp2mpRouteTarget []*L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget
}

func (mp2mpRouteTargets *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets) GetEntityData() *types.CommonEntityData {
    mp2mpRouteTargets.EntityData.YFilter = mp2mpRouteTargets.YFilter
    mp2mpRouteTargets.EntityData.YangName = "mp2mp-route-targets"
    mp2mpRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    mp2mpRouteTargets.EntityData.ParentYangName = "mp2mp-auto-discovery"
    mp2mpRouteTargets.EntityData.SegmentPath = "mp2mp-route-targets"
    mp2mpRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpRouteTargets.EntityData.Children = types.NewOrderedMap()
    mp2mpRouteTargets.EntityData.Children.Append("mp2mp-route-target", types.YChild{"Mp2mpRouteTarget", nil})
    for i := range mp2mpRouteTargets.Mp2mpRouteTarget {
        mp2mpRouteTargets.EntityData.Children.Append(types.GetSegmentPath(mp2mpRouteTargets.Mp2mpRouteTarget[i]), types.YChild{"Mp2mpRouteTarget", mp2mpRouteTargets.Mp2mpRouteTarget[i]})
    }
    mp2mpRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    mp2mpRouteTargets.EntityData.YListKeys = []string {}

    return &(mp2mpRouteTargets.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget
// Name of the Route Target
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // two byte as or four byte as. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_TwoByteAsOrFourByteAs.
    TwoByteAsOrFourByteAs []*L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_TwoByteAsOrFourByteAs

    // ipv4 address. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_Ipv4Address.
    Ipv4Address []*L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_Ipv4Address
}

func (mp2mpRouteTarget *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget) GetEntityData() *types.CommonEntityData {
    mp2mpRouteTarget.EntityData.YFilter = mp2mpRouteTarget.YFilter
    mp2mpRouteTarget.EntityData.YangName = "mp2mp-route-target"
    mp2mpRouteTarget.EntityData.BundleName = "cisco_ios_xr"
    mp2mpRouteTarget.EntityData.ParentYangName = "mp2mp-route-targets"
    mp2mpRouteTarget.EntityData.SegmentPath = "mp2mp-route-target" + types.AddKeyToken(mp2mpRouteTarget.Role, "role") + types.AddKeyToken(mp2mpRouteTarget.Format, "format")
    mp2mpRouteTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpRouteTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpRouteTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpRouteTarget.EntityData.Children = types.NewOrderedMap()
    mp2mpRouteTarget.EntityData.Children.Append("two-byte-as-or-four-byte-as", types.YChild{"TwoByteAsOrFourByteAs", nil})
    for i := range mp2mpRouteTarget.TwoByteAsOrFourByteAs {
        mp2mpRouteTarget.EntityData.Children.Append(types.GetSegmentPath(mp2mpRouteTarget.TwoByteAsOrFourByteAs[i]), types.YChild{"TwoByteAsOrFourByteAs", mp2mpRouteTarget.TwoByteAsOrFourByteAs[i]})
    }
    mp2mpRouteTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range mp2mpRouteTarget.Ipv4Address {
        mp2mpRouteTarget.EntityData.Children.Append(types.GetSegmentPath(mp2mpRouteTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", mp2mpRouteTarget.Ipv4Address[i]})
    }
    mp2mpRouteTarget.EntityData.Leafs = types.NewOrderedMap()
    mp2mpRouteTarget.EntityData.Leafs.Append("role", types.YLeaf{"Role", mp2mpRouteTarget.Role})
    mp2mpRouteTarget.EntityData.Leafs.Append("format", types.YLeaf{"Format", mp2mpRouteTarget.Format})

    mp2mpRouteTarget.EntityData.YListKeys = []string {"Role", "Format"}

    return &(mp2mpRouteTarget.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_TwoByteAsOrFourByteAs
// two byte as or four byte as
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_TwoByteAsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}
}

func (twoByteAsOrFourByteAs *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_TwoByteAsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAsOrFourByteAs.EntityData.YFilter = twoByteAsOrFourByteAs.YFilter
    twoByteAsOrFourByteAs.EntityData.YangName = "two-byte-as-or-four-byte-as"
    twoByteAsOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAsOrFourByteAs.EntityData.ParentYangName = "mp2mp-route-target"
    twoByteAsOrFourByteAs.EntityData.SegmentPath = "two-byte-as-or-four-byte-as" + types.AddKeyToken(twoByteAsOrFourByteAs.As, "as") + types.AddKeyToken(twoByteAsOrFourByteAs.AsIndex, "as-index")
    twoByteAsOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAsOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAsOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAsOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAsOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAsOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", twoByteAsOrFourByteAs.As})
    twoByteAsOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", twoByteAsOrFourByteAs.AsIndex})

    twoByteAsOrFourByteAs.EntityData.YListKeys = []string {"As", "AsIndex"}

    return &(twoByteAsOrFourByteAs.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_Ipv4Address
// ipv4 address
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}
}

func (ipv4Address *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpRouteTargets_Mp2mpRouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "mp2mp-route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddrIndex, "addr-index")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", ipv4Address.AddrIndex})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddrIndex"}

    return &(ipv4Address.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol
// signaling protocol in this MP2MP
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Customer Edge Identifier. The type is interface{} with range:
    // 11..100.
    CeRange interface{}

    // Enable signaling protocol. The type is interface{}.
    Enable interface{}

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_FlowLabelLoadBalance

    // Local Customer Edge Identifier Table.
    Ceids L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids
}

func (mp2mpSignalingProtocol *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    mp2mpSignalingProtocol.EntityData.YFilter = mp2mpSignalingProtocol.YFilter
    mp2mpSignalingProtocol.EntityData.YangName = "mp2mp-signaling-protocol"
    mp2mpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    mp2mpSignalingProtocol.EntityData.ParentYangName = "mp2mp-auto-discovery"
    mp2mpSignalingProtocol.EntityData.SegmentPath = "mp2mp-signaling-protocol"
    mp2mpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2mpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2mpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2mpSignalingProtocol.EntityData.Children = types.NewOrderedMap()
    mp2mpSignalingProtocol.EntityData.Children.Append("flow-label-load-balance", types.YChild{"FlowLabelLoadBalance", &mp2mpSignalingProtocol.FlowLabelLoadBalance})
    mp2mpSignalingProtocol.EntityData.Children.Append("ceids", types.YChild{"Ceids", &mp2mpSignalingProtocol.Ceids})
    mp2mpSignalingProtocol.EntityData.Leafs = types.NewOrderedMap()
    mp2mpSignalingProtocol.EntityData.Leafs.Append("ce-range", types.YLeaf{"CeRange", mp2mpSignalingProtocol.CeRange})
    mp2mpSignalingProtocol.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mp2mpSignalingProtocol.Enable})

    mp2mpSignalingProtocol.EntityData.YListKeys = []string {}

    return &(mp2mpSignalingProtocol.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "mp2mp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs.Append("flow-label", types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel})
    flowLabelLoadBalance.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabelLoadBalance.Static})

    flowLabelLoadBalance.EntityData.YListKeys = []string {}

    return &(flowLabelLoadBalance.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids
// Local Customer Edge Identifier Table
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Customer Edge Identifier . The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid.
    Ceid []*L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid
}

func (ceids *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids) GetEntityData() *types.CommonEntityData {
    ceids.EntityData.YFilter = ceids.YFilter
    ceids.EntityData.YangName = "ceids"
    ceids.EntityData.BundleName = "cisco_ios_xr"
    ceids.EntityData.ParentYangName = "mp2mp-signaling-protocol"
    ceids.EntityData.SegmentPath = "ceids"
    ceids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ceids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ceids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ceids.EntityData.Children = types.NewOrderedMap()
    ceids.EntityData.Children.Append("ceid", types.YChild{"Ceid", nil})
    for i := range ceids.Ceid {
        ceids.EntityData.Children.Append(types.GetSegmentPath(ceids.Ceid[i]), types.YChild{"Ceid", ceids.Ceid[i]})
    }
    ceids.EntityData.Leafs = types.NewOrderedMap()

    ceids.EntityData.YListKeys = []string {}

    return &(ceids.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid
// Local Customer Edge Identifier 
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local Customer Edge Identifier. The type is
    // interface{} with range: 1..16384.
    CeId interface{}

    // AC And Remote Customer Edge Identifier Table.
    RemoteCeidAttachmentCircuits L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits
}

func (ceid *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid) GetEntityData() *types.CommonEntityData {
    ceid.EntityData.YFilter = ceid.YFilter
    ceid.EntityData.YangName = "ceid"
    ceid.EntityData.BundleName = "cisco_ios_xr"
    ceid.EntityData.ParentYangName = "ceids"
    ceid.EntityData.SegmentPath = "ceid" + types.AddKeyToken(ceid.CeId, "ce-id")
    ceid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ceid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ceid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ceid.EntityData.Children = types.NewOrderedMap()
    ceid.EntityData.Children.Append("remote-ceid-attachment-circuits", types.YChild{"RemoteCeidAttachmentCircuits", &ceid.RemoteCeidAttachmentCircuits})
    ceid.EntityData.Leafs = types.NewOrderedMap()
    ceid.EntityData.Leafs.Append("ce-id", types.YLeaf{"CeId", ceid.CeId})

    ceid.EntityData.YListKeys = []string {"CeId"}

    return &(ceid.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits
// AC And Remote Customer Edge Identifier
// Table
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AC And Remote Customer Edge Identifier. The type is slice of
    // L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit.
    RemoteCeidAttachmentCircuit []*L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit
}

func (remoteCeidAttachmentCircuits *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    remoteCeidAttachmentCircuits.EntityData.YFilter = remoteCeidAttachmentCircuits.YFilter
    remoteCeidAttachmentCircuits.EntityData.YangName = "remote-ceid-attachment-circuits"
    remoteCeidAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    remoteCeidAttachmentCircuits.EntityData.ParentYangName = "ceid"
    remoteCeidAttachmentCircuits.EntityData.SegmentPath = "remote-ceid-attachment-circuits"
    remoteCeidAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteCeidAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteCeidAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteCeidAttachmentCircuits.EntityData.Children = types.NewOrderedMap()
    remoteCeidAttachmentCircuits.EntityData.Children.Append("remote-ceid-attachment-circuit", types.YChild{"RemoteCeidAttachmentCircuit", nil})
    for i := range remoteCeidAttachmentCircuits.RemoteCeidAttachmentCircuit {
        remoteCeidAttachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(remoteCeidAttachmentCircuits.RemoteCeidAttachmentCircuit[i]), types.YChild{"RemoteCeidAttachmentCircuit", remoteCeidAttachmentCircuits.RemoteCeidAttachmentCircuit[i]})
    }
    remoteCeidAttachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    remoteCeidAttachmentCircuits.EntityData.YListKeys = []string {}

    return &(remoteCeidAttachmentCircuits.EntityData)
}

// L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit
// AC And Remote Customer Edge Identifier
type L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Attachment Circuit. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    Name interface{}

    // This attribute is a key. Remote Customer Edge Identifier. The type is
    // interface{} with range: 1..16384.
    RemoteCeId interface{}
}

func (remoteCeidAttachmentCircuit *L2vpn_Database_XconnectGroups_XconnectGroup_Mp2mpXconnects_Mp2mpXconnect_Mp2mpAutoDiscovery_Mp2mpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    remoteCeidAttachmentCircuit.EntityData.YFilter = remoteCeidAttachmentCircuit.YFilter
    remoteCeidAttachmentCircuit.EntityData.YangName = "remote-ceid-attachment-circuit"
    remoteCeidAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    remoteCeidAttachmentCircuit.EntityData.ParentYangName = "remote-ceid-attachment-circuits"
    remoteCeidAttachmentCircuit.EntityData.SegmentPath = "remote-ceid-attachment-circuit" + types.AddKeyToken(remoteCeidAttachmentCircuit.Name, "name") + types.AddKeyToken(remoteCeidAttachmentCircuit.RemoteCeId, "remote-ce-id")
    remoteCeidAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteCeidAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteCeidAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteCeidAttachmentCircuit.EntityData.Children = types.NewOrderedMap()
    remoteCeidAttachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    remoteCeidAttachmentCircuit.EntityData.Leafs.Append("name", types.YLeaf{"Name", remoteCeidAttachmentCircuit.Name})
    remoteCeidAttachmentCircuit.EntityData.Leafs.Append("remote-ce-id", types.YLeaf{"RemoteCeId", remoteCeidAttachmentCircuit.RemoteCeId})

    remoteCeidAttachmentCircuit.EntityData.YListKeys = []string {"Name", "RemoteCeId"}

    return &(remoteCeidAttachmentCircuit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups
// List of bridge groups
type L2vpn_Database_BridgeDomainGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge group. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup.
    BridgeDomainGroup []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup
}

func (bridgeDomainGroups *L2vpn_Database_BridgeDomainGroups) GetEntityData() *types.CommonEntityData {
    bridgeDomainGroups.EntityData.YFilter = bridgeDomainGroups.YFilter
    bridgeDomainGroups.EntityData.YangName = "bridge-domain-groups"
    bridgeDomainGroups.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainGroups.EntityData.ParentYangName = "database"
    bridgeDomainGroups.EntityData.SegmentPath = "bridge-domain-groups"
    bridgeDomainGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainGroups.EntityData.Children = types.NewOrderedMap()
    bridgeDomainGroups.EntityData.Children.Append("bridge-domain-group", types.YChild{"BridgeDomainGroup", nil})
    for i := range bridgeDomainGroups.BridgeDomainGroup {
        bridgeDomainGroups.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainGroups.BridgeDomainGroup[i]), types.YChild{"BridgeDomainGroup", bridgeDomainGroups.BridgeDomainGroup[i]})
    }
    bridgeDomainGroups.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainGroups.EntityData.YListKeys = []string {}

    return &(bridgeDomainGroups.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup
// Bridge group
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Bridge group. The type is string with
    // length: 1..32.
    Name interface{}

    // List of Bridge Domain.
    BridgeDomains L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains
}

func (bridgeDomainGroup *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup) GetEntityData() *types.CommonEntityData {
    bridgeDomainGroup.EntityData.YFilter = bridgeDomainGroup.YFilter
    bridgeDomainGroup.EntityData.YangName = "bridge-domain-group"
    bridgeDomainGroup.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainGroup.EntityData.ParentYangName = "bridge-domain-groups"
    bridgeDomainGroup.EntityData.SegmentPath = "bridge-domain-group" + types.AddKeyToken(bridgeDomainGroup.Name, "name")
    bridgeDomainGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainGroup.EntityData.Children = types.NewOrderedMap()
    bridgeDomainGroup.EntityData.Children.Append("bridge-domains", types.YChild{"BridgeDomains", &bridgeDomainGroup.BridgeDomains})
    bridgeDomainGroup.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainGroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", bridgeDomainGroup.Name})

    bridgeDomainGroup.EntityData.YListKeys = []string {"Name"}

    return &(bridgeDomainGroup.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains
// List of Bridge Domain
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bridge domain. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain.
    BridgeDomain []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain
}

func (bridgeDomains *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomains.EntityData.ParentYangName = "bridge-domain-group"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomains.EntityData.Children = types.NewOrderedMap()
    bridgeDomains.EntityData.Children.Append("bridge-domain", types.YChild{"BridgeDomain", nil})
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children.Append(types.GetSegmentPath(bridgeDomains.BridgeDomain[i]), types.YChild{"BridgeDomain", bridgeDomains.BridgeDomain[i]})
    }
    bridgeDomains.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomains.EntityData.YListKeys = []string {}

    return &(bridgeDomains.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain
// bridge domain
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the bridge domain. The type is string with
    // length: 1..27.
    Name interface{}

    // Coupled-mode configuration. The type is interface{}.
    CoupledMode interface{}

    // shutdown the Bridge Domain. The type is interface{}.
    Shutdown interface{}

    // Disable Unknown Unicast flooding. The type is interface{}.
    FloodingUnknownUnicast interface{}

    // Disable IGMP Snooping. The type is interface{}.
    IgmpSnoopingDisable interface{}

    // Bridge Domain Transport mode. The type is BridgeDomainTransportMode.
    TransportMode interface{}

    // Attach MLD Snooping Profile Name. The type is string with length: 1..32.
    MldSnooping interface{}

    // Maximum transmission unit for this Bridge Domain. The type is interface{}
    // with range: 46..65535. Units are byte.
    BridgeDomainMtu interface{}

    // DHCPv4 Snooping profile name. The type is string with length: 1..32.
    Dhcp interface{}

    // Bridge-domain description Name. The type is string with length: 1..64.
    BridgeDescription interface{}

    // Attach IGMP Snooping Profile Name. The type is string with length: 1..32.
    IgmpSnooping interface{}

    // Disable flooding. The type is interface{}.
    Flooding interface{}

    // Storm Control.
    BdStormControls L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls

    // Bridge Domain VxLAN Network Identifier Table.
    MemberVnis L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis

    // MAC configuration commands.
    BridgeDomainMac L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac

    // nV Satellite.
    NvSatellite L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite

    // Bridge Domain PBB.
    BridgeDomainPbb L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb

    // Bridge Domain EVI Table.
    BridgeDomainEvis L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis

    // Specify the access virtual forwarding interface name.
    AccessVfis L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis

    // List of pseudowires.
    BdPseudowires L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires

    // Specify the virtual forwarding interface name.
    Vfis L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis

    // Bridge Domain EVPN VxLAN Network Identifier Table.
    BridgeDomainvnis L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis

    // Attachment Circuit table.
    BdAttachmentCircuits L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits

    // List of EVPN pseudowires.
    BdPseudowireEvpns L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns

    // IP Source Guard.
    IpSourceGuard L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard

    // Dynamic ARP Inspection.
    Dai L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai

    // Bridge Domain Routed Interface Table.
    RoutedInterfaces L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces
}

func (bridgeDomain *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomain.EntityData.ParentYangName = "bridge-domains"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + types.AddKeyToken(bridgeDomain.Name, "name")
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomain.EntityData.Children = types.NewOrderedMap()
    bridgeDomain.EntityData.Children.Append("bd-storm-controls", types.YChild{"BdStormControls", &bridgeDomain.BdStormControls})
    bridgeDomain.EntityData.Children.Append("member-vnis", types.YChild{"MemberVnis", &bridgeDomain.MemberVnis})
    bridgeDomain.EntityData.Children.Append("bridge-domain-mac", types.YChild{"BridgeDomainMac", &bridgeDomain.BridgeDomainMac})
    bridgeDomain.EntityData.Children.Append("nv-satellite", types.YChild{"NvSatellite", &bridgeDomain.NvSatellite})
    bridgeDomain.EntityData.Children.Append("bridge-domain-pbb", types.YChild{"BridgeDomainPbb", &bridgeDomain.BridgeDomainPbb})
    bridgeDomain.EntityData.Children.Append("bridge-domain-evis", types.YChild{"BridgeDomainEvis", &bridgeDomain.BridgeDomainEvis})
    bridgeDomain.EntityData.Children.Append("access-vfis", types.YChild{"AccessVfis", &bridgeDomain.AccessVfis})
    bridgeDomain.EntityData.Children.Append("bd-pseudowires", types.YChild{"BdPseudowires", &bridgeDomain.BdPseudowires})
    bridgeDomain.EntityData.Children.Append("vfis", types.YChild{"Vfis", &bridgeDomain.Vfis})
    bridgeDomain.EntityData.Children.Append("bridge-domainvnis", types.YChild{"BridgeDomainvnis", &bridgeDomain.BridgeDomainvnis})
    bridgeDomain.EntityData.Children.Append("bd-attachment-circuits", types.YChild{"BdAttachmentCircuits", &bridgeDomain.BdAttachmentCircuits})
    bridgeDomain.EntityData.Children.Append("bd-pseudowire-evpns", types.YChild{"BdPseudowireEvpns", &bridgeDomain.BdPseudowireEvpns})
    bridgeDomain.EntityData.Children.Append("ip-source-guard", types.YChild{"IpSourceGuard", &bridgeDomain.IpSourceGuard})
    bridgeDomain.EntityData.Children.Append("dai", types.YChild{"Dai", &bridgeDomain.Dai})
    bridgeDomain.EntityData.Children.Append("routed-interfaces", types.YChild{"RoutedInterfaces", &bridgeDomain.RoutedInterfaces})
    bridgeDomain.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomain.EntityData.Leafs.Append("name", types.YLeaf{"Name", bridgeDomain.Name})
    bridgeDomain.EntityData.Leafs.Append("coupled-mode", types.YLeaf{"CoupledMode", bridgeDomain.CoupledMode})
    bridgeDomain.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", bridgeDomain.Shutdown})
    bridgeDomain.EntityData.Leafs.Append("flooding-unknown-unicast", types.YLeaf{"FloodingUnknownUnicast", bridgeDomain.FloodingUnknownUnicast})
    bridgeDomain.EntityData.Leafs.Append("igmp-snooping-disable", types.YLeaf{"IgmpSnoopingDisable", bridgeDomain.IgmpSnoopingDisable})
    bridgeDomain.EntityData.Leafs.Append("transport-mode", types.YLeaf{"TransportMode", bridgeDomain.TransportMode})
    bridgeDomain.EntityData.Leafs.Append("mld-snooping", types.YLeaf{"MldSnooping", bridgeDomain.MldSnooping})
    bridgeDomain.EntityData.Leafs.Append("bridge-domain-mtu", types.YLeaf{"BridgeDomainMtu", bridgeDomain.BridgeDomainMtu})
    bridgeDomain.EntityData.Leafs.Append("dhcp", types.YLeaf{"Dhcp", bridgeDomain.Dhcp})
    bridgeDomain.EntityData.Leafs.Append("bridge-description", types.YLeaf{"BridgeDescription", bridgeDomain.BridgeDescription})
    bridgeDomain.EntityData.Leafs.Append("igmp-snooping", types.YLeaf{"IgmpSnooping", bridgeDomain.IgmpSnooping})
    bridgeDomain.EntityData.Leafs.Append("flooding", types.YLeaf{"Flooding", bridgeDomain.Flooding})

    bridgeDomain.EntityData.YListKeys = []string {"Name"}

    return &(bridgeDomain.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls
// Storm Control
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl.
    BdStormControl []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl
}

func (bdStormControls *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls) GetEntityData() *types.CommonEntityData {
    bdStormControls.EntityData.YFilter = bdStormControls.YFilter
    bdStormControls.EntityData.YangName = "bd-storm-controls"
    bdStormControls.EntityData.BundleName = "cisco_ios_xr"
    bdStormControls.EntityData.ParentYangName = "bridge-domain"
    bdStormControls.EntityData.SegmentPath = "bd-storm-controls"
    bdStormControls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdStormControls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdStormControls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdStormControls.EntityData.Children = types.NewOrderedMap()
    bdStormControls.EntityData.Children.Append("bd-storm-control", types.YChild{"BdStormControl", nil})
    for i := range bdStormControls.BdStormControl {
        bdStormControls.EntityData.Children.Append(types.GetSegmentPath(bdStormControls.BdStormControl[i]), types.YChild{"BdStormControl", bdStormControls.BdStormControl[i]})
    }
    bdStormControls.EntityData.Leafs = types.NewOrderedMap()

    bdStormControls.EntityData.YListKeys = []string {}

    return &(bdStormControls.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl
// Storm Control Type
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit
}

func (bdStormControl *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl) GetEntityData() *types.CommonEntityData {
    bdStormControl.EntityData.YFilter = bdStormControl.YFilter
    bdStormControl.EntityData.YangName = "bd-storm-control"
    bdStormControl.EntityData.BundleName = "cisco_ios_xr"
    bdStormControl.EntityData.ParentYangName = "bd-storm-controls"
    bdStormControl.EntityData.SegmentPath = "bd-storm-control" + types.AddKeyToken(bdStormControl.Sctype, "sctype")
    bdStormControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdStormControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdStormControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdStormControl.EntityData.Children = types.NewOrderedMap()
    bdStormControl.EntityData.Children.Append("storm-control-unit", types.YChild{"StormControlUnit", &bdStormControl.StormControlUnit})
    bdStormControl.EntityData.Leafs = types.NewOrderedMap()
    bdStormControl.EntityData.Leafs.Append("sctype", types.YLeaf{"Sctype", bdStormControl.Sctype})

    bdStormControl.EntityData.YListKeys = []string {"Sctype"}

    return &(bdStormControl.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit
// Specify units for Storm Control Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Kilobits Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 64..1280000. Units are
    // kbit/s.
    KbitsPerSec interface{}

    // Packets Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 1..160000. Units are
    // packet/s.
    PktsPerSec interface{}
}

func (stormControlUnit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bd-storm-control"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs.Append("kbits-per-sec", types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec})
    stormControlUnit.EntityData.Leafs.Append("pkts-per-sec", types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec})

    stormControlUnit.EntityData.YListKeys = []string {}

    return &(stormControlUnit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis
// Bridge Domain VxLAN Network Identifier Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain Member VxLAN Network Identifier. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni.
    MemberVni []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni
}

func (memberVnis *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis) GetEntityData() *types.CommonEntityData {
    memberVnis.EntityData.YFilter = memberVnis.YFilter
    memberVnis.EntityData.YangName = "member-vnis"
    memberVnis.EntityData.BundleName = "cisco_ios_xr"
    memberVnis.EntityData.ParentYangName = "bridge-domain"
    memberVnis.EntityData.SegmentPath = "member-vnis"
    memberVnis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVnis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVnis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVnis.EntityData.Children = types.NewOrderedMap()
    memberVnis.EntityData.Children.Append("member-vni", types.YChild{"MemberVni", nil})
    for i := range memberVnis.MemberVni {
        memberVnis.EntityData.Children.Append(types.GetSegmentPath(memberVnis.MemberVni[i]), types.YChild{"MemberVni", memberVnis.MemberVni[i]})
    }
    memberVnis.EntityData.Leafs = types.NewOrderedMap()

    memberVnis.EntityData.YListKeys = []string {}

    return &(memberVnis.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni
// Bridge Domain Member VxLAN Network Identifier
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VxLAN Network Identifier number. The type is
    // interface{} with range: 1..16777215.
    Vni interface{}

    // Static Mac Address Table.
    MemberVniStaticMacAddresses L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses
}

func (memberVni *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni) GetEntityData() *types.CommonEntityData {
    memberVni.EntityData.YFilter = memberVni.YFilter
    memberVni.EntityData.YangName = "member-vni"
    memberVni.EntityData.BundleName = "cisco_ios_xr"
    memberVni.EntityData.ParentYangName = "member-vnis"
    memberVni.EntityData.SegmentPath = "member-vni" + types.AddKeyToken(memberVni.Vni, "vni")
    memberVni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVni.EntityData.Children = types.NewOrderedMap()
    memberVni.EntityData.Children.Append("member-vni-static-mac-addresses", types.YChild{"MemberVniStaticMacAddresses", &memberVni.MemberVniStaticMacAddresses})
    memberVni.EntityData.Leafs = types.NewOrderedMap()
    memberVni.EntityData.Leafs.Append("vni", types.YLeaf{"Vni", memberVni.Vni})

    memberVni.EntityData.YListKeys = []string {"Vni"}

    return &(memberVni.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress.
    MemberVniStaticMacAddress []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress
}

func (memberVniStaticMacAddresses *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    memberVniStaticMacAddresses.EntityData.YFilter = memberVniStaticMacAddresses.YFilter
    memberVniStaticMacAddresses.EntityData.YangName = "member-vni-static-mac-addresses"
    memberVniStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    memberVniStaticMacAddresses.EntityData.ParentYangName = "member-vni"
    memberVniStaticMacAddresses.EntityData.SegmentPath = "member-vni-static-mac-addresses"
    memberVniStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVniStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVniStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVniStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    memberVniStaticMacAddresses.EntityData.Children.Append("member-vni-static-mac-address", types.YChild{"MemberVniStaticMacAddress", nil})
    for i := range memberVniStaticMacAddresses.MemberVniStaticMacAddress {
        memberVniStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(memberVniStaticMacAddresses.MemberVniStaticMacAddress[i]), types.YChild{"MemberVniStaticMacAddress", memberVniStaticMacAddresses.MemberVniStaticMacAddress[i]})
    }
    memberVniStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    memberVniStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(memberVniStaticMacAddresses.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Enable Static Mac Address Configuration. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopIp interface{}
}

func (memberVniStaticMacAddress *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress) GetEntityData() *types.CommonEntityData {
    memberVniStaticMacAddress.EntityData.YFilter = memberVniStaticMacAddress.YFilter
    memberVniStaticMacAddress.EntityData.YangName = "member-vni-static-mac-address"
    memberVniStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    memberVniStaticMacAddress.EntityData.ParentYangName = "member-vni-static-mac-addresses"
    memberVniStaticMacAddress.EntityData.SegmentPath = "member-vni-static-mac-address" + types.AddKeyToken(memberVniStaticMacAddress.MacAddress, "mac-address")
    memberVniStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVniStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVniStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVniStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    memberVniStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    memberVniStaticMacAddress.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", memberVniStaticMacAddress.MacAddress})
    memberVniStaticMacAddress.EntityData.Leafs.Append("next-hop-ip", types.YLeaf{"NextHopIp", memberVniStaticMacAddress.NextHopIp})

    memberVniStaticMacAddress.EntityData.YListKeys = []string {"MacAddress"}

    return &(memberVniStaticMacAddress.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac
// MAC configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mac withdraw sent from access PW to access PW. The type is interface{}.
    BdMacWithdrawRelay interface{}

    // MAC withdraw on Access PW. The type is interface{}.
    BdMacWithdrawAccessPwDisable interface{}

    // Disable MAC Flush when Port goes Down. The type is interface{}.
    BdMacPortDownFlush interface{}

    // Disable Mac Withdraw. The type is interface{}.
    BdMacWithdraw interface{}

    // MAC withdraw sent on bridge port down. The type is MacWithdrawBehavior.
    BdMacWithdrawBehavior interface{}

    // Mac Learning Type. The type is BdmacLearn.
    BdMacLearn interface{}

    // MAC-Limit configuration commands.
    BdMacLimit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit

    // Filter Mac Address.
    BdMacFilters L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters

    // MAC Secure.
    MacSecure L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure

    // MAC-Aging configuration commands.
    BdMacAging L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging
}

func (bridgeDomainMac *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac) GetEntityData() *types.CommonEntityData {
    bridgeDomainMac.EntityData.YFilter = bridgeDomainMac.YFilter
    bridgeDomainMac.EntityData.YangName = "bridge-domain-mac"
    bridgeDomainMac.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainMac.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainMac.EntityData.SegmentPath = "bridge-domain-mac"
    bridgeDomainMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainMac.EntityData.Children = types.NewOrderedMap()
    bridgeDomainMac.EntityData.Children.Append("bd-mac-limit", types.YChild{"BdMacLimit", &bridgeDomainMac.BdMacLimit})
    bridgeDomainMac.EntityData.Children.Append("bd-mac-filters", types.YChild{"BdMacFilters", &bridgeDomainMac.BdMacFilters})
    bridgeDomainMac.EntityData.Children.Append("mac-secure", types.YChild{"MacSecure", &bridgeDomainMac.MacSecure})
    bridgeDomainMac.EntityData.Children.Append("bd-mac-aging", types.YChild{"BdMacAging", &bridgeDomainMac.BdMacAging})
    bridgeDomainMac.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw-relay", types.YLeaf{"BdMacWithdrawRelay", bridgeDomainMac.BdMacWithdrawRelay})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw-access-pw-disable", types.YLeaf{"BdMacWithdrawAccessPwDisable", bridgeDomainMac.BdMacWithdrawAccessPwDisable})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-port-down-flush", types.YLeaf{"BdMacPortDownFlush", bridgeDomainMac.BdMacPortDownFlush})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw", types.YLeaf{"BdMacWithdraw", bridgeDomainMac.BdMacWithdraw})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw-behavior", types.YLeaf{"BdMacWithdrawBehavior", bridgeDomainMac.BdMacWithdrawBehavior})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-learn", types.YLeaf{"BdMacLearn", bridgeDomainMac.BdMacLearn})

    bridgeDomainMac.EntityData.YListKeys = []string {}

    return &(bridgeDomainMac.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address limit enforcement action. The type is MacLimitAction.
    BdMacLimitAction interface{}

    // Mac Address Limit Notification. The type is MacNotification.
    BdMacLimitNotif interface{}

    // Number of MAC addresses after which MAC limit action is taken. The type is
    // interface{} with range: 0..4294967295.
    BdMacLimitMax interface{}
}

func (bdMacLimit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit) GetEntityData() *types.CommonEntityData {
    bdMacLimit.EntityData.YFilter = bdMacLimit.YFilter
    bdMacLimit.EntityData.YangName = "bd-mac-limit"
    bdMacLimit.EntityData.BundleName = "cisco_ios_xr"
    bdMacLimit.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacLimit.EntityData.SegmentPath = "bd-mac-limit"
    bdMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacLimit.EntityData.Children = types.NewOrderedMap()
    bdMacLimit.EntityData.Leafs = types.NewOrderedMap()
    bdMacLimit.EntityData.Leafs.Append("bd-mac-limit-action", types.YLeaf{"BdMacLimitAction", bdMacLimit.BdMacLimitAction})
    bdMacLimit.EntityData.Leafs.Append("bd-mac-limit-notif", types.YLeaf{"BdMacLimitNotif", bdMacLimit.BdMacLimitNotif})
    bdMacLimit.EntityData.Leafs.Append("bd-mac-limit-max", types.YLeaf{"BdMacLimitMax", bdMacLimit.BdMacLimitMax})

    bdMacLimit.EntityData.YListKeys = []string {}

    return &(bdMacLimit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters
// Filter Mac Address
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static MAC address. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter.
    BdMacFilter []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter
}

func (bdMacFilters *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters) GetEntityData() *types.CommonEntityData {
    bdMacFilters.EntityData.YFilter = bdMacFilters.YFilter
    bdMacFilters.EntityData.YangName = "bd-mac-filters"
    bdMacFilters.EntityData.BundleName = "cisco_ios_xr"
    bdMacFilters.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacFilters.EntityData.SegmentPath = "bd-mac-filters"
    bdMacFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacFilters.EntityData.Children = types.NewOrderedMap()
    bdMacFilters.EntityData.Children.Append("bd-mac-filter", types.YChild{"BdMacFilter", nil})
    for i := range bdMacFilters.BdMacFilter {
        bdMacFilters.EntityData.Children.Append(types.GetSegmentPath(bdMacFilters.BdMacFilter[i]), types.YChild{"BdMacFilter", bdMacFilters.BdMacFilter[i]})
    }
    bdMacFilters.EntityData.Leafs = types.NewOrderedMap()

    bdMacFilters.EntityData.YListKeys = []string {}

    return &(bdMacFilters.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter
// Static MAC address
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}

    // MAC address for filtering. The type is interface{}.
    Drop interface{}
}

func (bdMacFilter *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter) GetEntityData() *types.CommonEntityData {
    bdMacFilter.EntityData.YFilter = bdMacFilter.YFilter
    bdMacFilter.EntityData.YangName = "bd-mac-filter"
    bdMacFilter.EntityData.BundleName = "cisco_ios_xr"
    bdMacFilter.EntityData.ParentYangName = "bd-mac-filters"
    bdMacFilter.EntityData.SegmentPath = "bd-mac-filter" + types.AddKeyToken(bdMacFilter.Address, "address")
    bdMacFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacFilter.EntityData.Children = types.NewOrderedMap()
    bdMacFilter.EntityData.Leafs = types.NewOrderedMap()
    bdMacFilter.EntityData.Leafs.Append("address", types.YLeaf{"Address", bdMacFilter.Address})
    bdMacFilter.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", bdMacFilter.Drop})

    bdMacFilter.EntityData.YListKeys = []string {"Address"}

    return &(bdMacFilter.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure
// MAC Secure
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is interface{}.
    Logging interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}

    // MAC Secure Threshold. The type is interface{}.
    Threshold interface{}
}

func (macSecure *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure) GetEntityData() *types.CommonEntityData {
    macSecure.EntityData.YFilter = macSecure.YFilter
    macSecure.EntityData.YangName = "mac-secure"
    macSecure.EntityData.BundleName = "cisco_ios_xr"
    macSecure.EntityData.ParentYangName = "bridge-domain-mac"
    macSecure.EntityData.SegmentPath = "mac-secure"
    macSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSecure.EntityData.Children = types.NewOrderedMap()
    macSecure.EntityData.Leafs = types.NewOrderedMap()
    macSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", macSecure.Logging})
    macSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", macSecure.Action})
    macSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", macSecure.Enable})
    macSecure.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", macSecure.Threshold})

    macSecure.EntityData.YListKeys = []string {}

    return &(macSecure.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging
// MAC-Aging configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    BdMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    BdMacAgingTime interface{}
}

func (bdMacAging *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging) GetEntityData() *types.CommonEntityData {
    bdMacAging.EntityData.YFilter = bdMacAging.YFilter
    bdMacAging.EntityData.YangName = "bd-mac-aging"
    bdMacAging.EntityData.BundleName = "cisco_ios_xr"
    bdMacAging.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacAging.EntityData.SegmentPath = "bd-mac-aging"
    bdMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacAging.EntityData.Children = types.NewOrderedMap()
    bdMacAging.EntityData.Leafs = types.NewOrderedMap()
    bdMacAging.EntityData.Leafs.Append("bd-mac-aging-type", types.YLeaf{"BdMacAgingType", bdMacAging.BdMacAgingType})
    bdMacAging.EntityData.Leafs.Append("bd-mac-aging-time", types.YLeaf{"BdMacAgingTime", bdMacAging.BdMacAgingTime})

    bdMacAging.EntityData.YListKeys = []string {}

    return &(bdMacAging.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite
// nV Satellite
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPv4 Multicast Offload to Satellite Nodes. The type is interface{}.
    OffloadIpv4MulticastEnable interface{}

    // Enable nV Satellite Settings. The type is interface{}.
    Enable interface{}
}

func (nvSatellite *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite) GetEntityData() *types.CommonEntityData {
    nvSatellite.EntityData.YFilter = nvSatellite.YFilter
    nvSatellite.EntityData.YangName = "nv-satellite"
    nvSatellite.EntityData.BundleName = "cisco_ios_xr"
    nvSatellite.EntityData.ParentYangName = "bridge-domain"
    nvSatellite.EntityData.SegmentPath = "nv-satellite"
    nvSatellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellite.EntityData.Children = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs.Append("offload-ipv4-multicast-enable", types.YLeaf{"OffloadIpv4MulticastEnable", nvSatellite.OffloadIpv4MulticastEnable})
    nvSatellite.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", nvSatellite.Enable})

    nvSatellite.EntityData.YListKeys = []string {}

    return &(nvSatellite.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb
// Bridge Domain PBB
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Edge.
    PbbEdges L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges

    // PBB Core.
    PbbCore L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore
}

func (bridgeDomainPbb *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb) GetEntityData() *types.CommonEntityData {
    bridgeDomainPbb.EntityData.YFilter = bridgeDomainPbb.YFilter
    bridgeDomainPbb.EntityData.YangName = "bridge-domain-pbb"
    bridgeDomainPbb.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainPbb.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainPbb.EntityData.SegmentPath = "bridge-domain-pbb"
    bridgeDomainPbb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainPbb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainPbb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainPbb.EntityData.Children = types.NewOrderedMap()
    bridgeDomainPbb.EntityData.Children.Append("pbb-edges", types.YChild{"PbbEdges", &bridgeDomainPbb.PbbEdges})
    bridgeDomainPbb.EntityData.Children.Append("pbb-core", types.YChild{"PbbCore", &bridgeDomainPbb.PbbCore})
    bridgeDomainPbb.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainPbb.EntityData.YListKeys = []string {}

    return &(bridgeDomainPbb.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges
// PBB Edge
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure BD as PBB Edge with ISID and associated PBB Core BD. The type is
    // slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge.
    PbbEdge []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge
}

func (pbbEdges *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges) GetEntityData() *types.CommonEntityData {
    pbbEdges.EntityData.YFilter = pbbEdges.YFilter
    pbbEdges.EntityData.YangName = "pbb-edges"
    pbbEdges.EntityData.BundleName = "cisco_ios_xr"
    pbbEdges.EntityData.ParentYangName = "bridge-domain-pbb"
    pbbEdges.EntityData.SegmentPath = "pbb-edges"
    pbbEdges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdges.EntityData.Children = types.NewOrderedMap()
    pbbEdges.EntityData.Children.Append("pbb-edge", types.YChild{"PbbEdge", nil})
    for i := range pbbEdges.PbbEdge {
        pbbEdges.EntityData.Children.Append(types.GetSegmentPath(pbbEdges.PbbEdge[i]), types.YChild{"PbbEdge", pbbEdges.PbbEdge[i]})
    }
    pbbEdges.EntityData.Leafs = types.NewOrderedMap()

    pbbEdges.EntityData.YListKeys = []string {}

    return &(pbbEdges.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge
// Configure BD as PBB Edge with ISID and
// associated PBB Core BD
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ISID. The type is interface{} with range:
    // 256..16777214.
    Isid interface{}

    // This attribute is a key. Core BD Name. The type is string with length:
    // 1..27.
    CoreBdName interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    PbbEdgeIgmpProfile interface{}

    // Configure Unknown Unicast BMAC address for PBB Edge Port. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    UnknownUnicastBmac interface{}

    // Split Horizon Group.
    PbbEdgeSplitHorizonGroup L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup

    // PBB Static Mac Address Mapping Table.
    PbbStaticMacMappings L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings

    // Attach a DHCP profile.
    PbbEdgeDhcpProfile L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile

    // MAC configuration commands.
    PbbEdgeMac L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac
}

func (pbbEdge *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge) GetEntityData() *types.CommonEntityData {
    pbbEdge.EntityData.YFilter = pbbEdge.YFilter
    pbbEdge.EntityData.YangName = "pbb-edge"
    pbbEdge.EntityData.BundleName = "cisco_ios_xr"
    pbbEdge.EntityData.ParentYangName = "pbb-edges"
    pbbEdge.EntityData.SegmentPath = "pbb-edge" + types.AddKeyToken(pbbEdge.Isid, "isid") + types.AddKeyToken(pbbEdge.CoreBdName, "core-bd-name")
    pbbEdge.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdge.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdge.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdge.EntityData.Children = types.NewOrderedMap()
    pbbEdge.EntityData.Children.Append("pbb-edge-split-horizon-group", types.YChild{"PbbEdgeSplitHorizonGroup", &pbbEdge.PbbEdgeSplitHorizonGroup})
    pbbEdge.EntityData.Children.Append("pbb-static-mac-mappings", types.YChild{"PbbStaticMacMappings", &pbbEdge.PbbStaticMacMappings})
    pbbEdge.EntityData.Children.Append("pbb-edge-dhcp-profile", types.YChild{"PbbEdgeDhcpProfile", &pbbEdge.PbbEdgeDhcpProfile})
    pbbEdge.EntityData.Children.Append("pbb-edge-mac", types.YChild{"PbbEdgeMac", &pbbEdge.PbbEdgeMac})
    pbbEdge.EntityData.Leafs = types.NewOrderedMap()
    pbbEdge.EntityData.Leafs.Append("isid", types.YLeaf{"Isid", pbbEdge.Isid})
    pbbEdge.EntityData.Leafs.Append("core-bd-name", types.YLeaf{"CoreBdName", pbbEdge.CoreBdName})
    pbbEdge.EntityData.Leafs.Append("pbb-edge-igmp-profile", types.YLeaf{"PbbEdgeIgmpProfile", pbbEdge.PbbEdgeIgmpProfile})
    pbbEdge.EntityData.Leafs.Append("unknown-unicast-bmac", types.YLeaf{"UnknownUnicastBmac", pbbEdge.UnknownUnicastBmac})

    pbbEdge.EntityData.YListKeys = []string {"Isid", "CoreBdName"}

    return &(pbbEdge.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup
// Split Horizon Group
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable split horizon group. The type is interface{}.
    Disable interface{}
}

func (pbbEdgeSplitHorizonGroup *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    pbbEdgeSplitHorizonGroup.EntityData.YFilter = pbbEdgeSplitHorizonGroup.YFilter
    pbbEdgeSplitHorizonGroup.EntityData.YangName = "pbb-edge-split-horizon-group"
    pbbEdgeSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeSplitHorizonGroup.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeSplitHorizonGroup.EntityData.SegmentPath = "pbb-edge-split-horizon-group"
    pbbEdgeSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeSplitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    pbbEdgeSplitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeSplitHorizonGroup.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pbbEdgeSplitHorizonGroup.Disable})

    pbbEdgeSplitHorizonGroup.EntityData.YListKeys = []string {}

    return &(pbbEdgeSplitHorizonGroup.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings
// PBB Static Mac Address Mapping Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Static Mac Address Mapping Configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping.
    PbbStaticMacMapping []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping
}

func (pbbStaticMacMappings *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings) GetEntityData() *types.CommonEntityData {
    pbbStaticMacMappings.EntityData.YFilter = pbbStaticMacMappings.YFilter
    pbbStaticMacMappings.EntityData.YangName = "pbb-static-mac-mappings"
    pbbStaticMacMappings.EntityData.BundleName = "cisco_ios_xr"
    pbbStaticMacMappings.EntityData.ParentYangName = "pbb-edge"
    pbbStaticMacMappings.EntityData.SegmentPath = "pbb-static-mac-mappings"
    pbbStaticMacMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbStaticMacMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbStaticMacMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbStaticMacMappings.EntityData.Children = types.NewOrderedMap()
    pbbStaticMacMappings.EntityData.Children.Append("pbb-static-mac-mapping", types.YChild{"PbbStaticMacMapping", nil})
    for i := range pbbStaticMacMappings.PbbStaticMacMapping {
        pbbStaticMacMappings.EntityData.Children.Append(types.GetSegmentPath(pbbStaticMacMappings.PbbStaticMacMapping[i]), types.YChild{"PbbStaticMacMapping", pbbStaticMacMappings.PbbStaticMacMapping[i]})
    }
    pbbStaticMacMappings.EntityData.Leafs = types.NewOrderedMap()

    pbbStaticMacMappings.EntityData.YListKeys = []string {}

    return &(pbbStaticMacMappings.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping
// PBB Static Mac Address Mapping
// Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}

    // Static backbone MAC address to map with. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PbbStaticMacMappingBmac interface{}
}

func (pbbStaticMacMapping *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping) GetEntityData() *types.CommonEntityData {
    pbbStaticMacMapping.EntityData.YFilter = pbbStaticMacMapping.YFilter
    pbbStaticMacMapping.EntityData.YangName = "pbb-static-mac-mapping"
    pbbStaticMacMapping.EntityData.BundleName = "cisco_ios_xr"
    pbbStaticMacMapping.EntityData.ParentYangName = "pbb-static-mac-mappings"
    pbbStaticMacMapping.EntityData.SegmentPath = "pbb-static-mac-mapping" + types.AddKeyToken(pbbStaticMacMapping.Address, "address")
    pbbStaticMacMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbStaticMacMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbStaticMacMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbStaticMacMapping.EntityData.Children = types.NewOrderedMap()
    pbbStaticMacMapping.EntityData.Leafs = types.NewOrderedMap()
    pbbStaticMacMapping.EntityData.Leafs.Append("address", types.YLeaf{"Address", pbbStaticMacMapping.Address})
    pbbStaticMacMapping.EntityData.Leafs.Append("pbb-static-mac-mapping-bmac", types.YLeaf{"PbbStaticMacMappingBmac", pbbStaticMacMapping.PbbStaticMacMappingBmac})

    pbbStaticMacMapping.EntityData.YListKeys = []string {"Address"}

    return &(pbbStaticMacMapping.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile
// Attach a DHCP profile
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pbbEdgeDhcpProfile *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile) GetEntityData() *types.CommonEntityData {
    pbbEdgeDhcpProfile.EntityData.YFilter = pbbEdgeDhcpProfile.YFilter
    pbbEdgeDhcpProfile.EntityData.YangName = "pbb-edge-dhcp-profile"
    pbbEdgeDhcpProfile.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeDhcpProfile.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeDhcpProfile.EntityData.SegmentPath = "pbb-edge-dhcp-profile"
    pbbEdgeDhcpProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeDhcpProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeDhcpProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeDhcpProfile.EntityData.Children = types.NewOrderedMap()
    pbbEdgeDhcpProfile.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeDhcpProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", pbbEdgeDhcpProfile.ProfileId})
    pbbEdgeDhcpProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", pbbEdgeDhcpProfile.DhcpSnoopingId})

    pbbEdgeDhcpProfile.EntityData.YListKeys = []string {}

    return &(pbbEdgeDhcpProfile.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac
// MAC configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Mac Learning. The type is MacLearn.
    PbbEdgeMacLearning interface{}

    // MAC-Limit configuration commands.
    PbbEdgeMacLimit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit

    // MAC-Aging configuration commands.
    PbbEdgeMacAging L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging

    // MAC Secure.
    PbbEdgeMacSecure L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure
}

func (pbbEdgeMac *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac) GetEntityData() *types.CommonEntityData {
    pbbEdgeMac.EntityData.YFilter = pbbEdgeMac.YFilter
    pbbEdgeMac.EntityData.YangName = "pbb-edge-mac"
    pbbEdgeMac.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMac.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeMac.EntityData.SegmentPath = "pbb-edge-mac"
    pbbEdgeMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMac.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMac.EntityData.Children.Append("pbb-edge-mac-limit", types.YChild{"PbbEdgeMacLimit", &pbbEdgeMac.PbbEdgeMacLimit})
    pbbEdgeMac.EntityData.Children.Append("pbb-edge-mac-aging", types.YChild{"PbbEdgeMacAging", &pbbEdgeMac.PbbEdgeMacAging})
    pbbEdgeMac.EntityData.Children.Append("pbb-edge-mac-secure", types.YChild{"PbbEdgeMacSecure", &pbbEdgeMac.PbbEdgeMacSecure})
    pbbEdgeMac.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMac.EntityData.Leafs.Append("pbb-edge-mac-learning", types.YLeaf{"PbbEdgeMacLearning", pbbEdgeMac.PbbEdgeMacLearning})

    pbbEdgeMac.EntityData.YListKeys = []string {}

    return &(pbbEdgeMac.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address limit enforcement action. The type is MacLimitAction.
    PbbEdgeMacLimitAction interface{}

    // Number of MAC addresses after which MAC limit action is taken. The type is
    // interface{} with range: 0..4294967295.
    PbbEdgeMacLimitMax interface{}

    // MAC address limit notification action. The type is MacNotification.
    PbbEdgeMacLimitNotif interface{}
}

func (pbbEdgeMacLimit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacLimit.EntityData.YFilter = pbbEdgeMacLimit.YFilter
    pbbEdgeMacLimit.EntityData.YangName = "pbb-edge-mac-limit"
    pbbEdgeMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacLimit.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacLimit.EntityData.SegmentPath = "pbb-edge-mac-limit"
    pbbEdgeMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacLimit.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMacLimit.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMacLimit.EntityData.Leafs.Append("pbb-edge-mac-limit-action", types.YLeaf{"PbbEdgeMacLimitAction", pbbEdgeMacLimit.PbbEdgeMacLimitAction})
    pbbEdgeMacLimit.EntityData.Leafs.Append("pbb-edge-mac-limit-max", types.YLeaf{"PbbEdgeMacLimitMax", pbbEdgeMacLimit.PbbEdgeMacLimitMax})
    pbbEdgeMacLimit.EntityData.Leafs.Append("pbb-edge-mac-limit-notif", types.YLeaf{"PbbEdgeMacLimitNotif", pbbEdgeMacLimit.PbbEdgeMacLimitNotif})

    pbbEdgeMacLimit.EntityData.YListKeys = []string {}

    return &(pbbEdgeMacLimit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging
// MAC-Aging configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PbbEdgeMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    PbbEdgeMacAgingTime interface{}
}

func (pbbEdgeMacAging *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacAging.EntityData.YFilter = pbbEdgeMacAging.YFilter
    pbbEdgeMacAging.EntityData.YangName = "pbb-edge-mac-aging"
    pbbEdgeMacAging.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacAging.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacAging.EntityData.SegmentPath = "pbb-edge-mac-aging"
    pbbEdgeMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacAging.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMacAging.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMacAging.EntityData.Leafs.Append("pbb-edge-mac-aging-type", types.YLeaf{"PbbEdgeMacAgingType", pbbEdgeMacAging.PbbEdgeMacAgingType})
    pbbEdgeMacAging.EntityData.Leafs.Append("pbb-edge-mac-aging-time", types.YLeaf{"PbbEdgeMacAgingTime", pbbEdgeMacAging.PbbEdgeMacAgingTime})

    pbbEdgeMacAging.EntityData.YListKeys = []string {}

    return &(pbbEdgeMacAging.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure
// MAC Secure
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is L2vpnLogging.
    Logging interface{}

    // Disable Virtual instance port MAC Secure. The type is interface{}.
    Disable interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}

    // Accept Virtual instance port to be shutdown on mac violation. The type is
    // interface{}.
    AcceptShutdown interface{}
}

func (pbbEdgeMacSecure *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacSecure.EntityData.YFilter = pbbEdgeMacSecure.YFilter
    pbbEdgeMacSecure.EntityData.YangName = "pbb-edge-mac-secure"
    pbbEdgeMacSecure.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacSecure.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacSecure.EntityData.SegmentPath = "pbb-edge-mac-secure"
    pbbEdgeMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacSecure.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMacSecure.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMacSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pbbEdgeMacSecure.Logging})
    pbbEdgeMacSecure.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pbbEdgeMacSecure.Disable})
    pbbEdgeMacSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", pbbEdgeMacSecure.Action})
    pbbEdgeMacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pbbEdgeMacSecure.Enable})
    pbbEdgeMacSecure.EntityData.Leafs.Append("accept-shutdown", types.YLeaf{"AcceptShutdown", pbbEdgeMacSecure.AcceptShutdown})

    pbbEdgeMacSecure.EntityData.YListKeys = []string {}

    return &(pbbEdgeMacSecure.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore
// PBB Core
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enabling MMRP PBB-VPLS Flood Optimization. The type is interface{}.
    PbbCoreMmrpFloodOptimization interface{}

    // VLAN ID to push. The type is interface{} with range: 1..4094.
    VlanId interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    PbbCoreIgmpProfile interface{}

    // Enable Bridge Domain PBB Core Configuration. The type is interface{}.
    Enable interface{}

    // MAC configuration commands.
    PbbCoreMac L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac

    // PBB Core EVI Table.
    PbbCoreEvis L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis

    // Attach a DHCP profile.
    PbbCoreDhcpProfile L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile
}

func (pbbCore *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore) GetEntityData() *types.CommonEntityData {
    pbbCore.EntityData.YFilter = pbbCore.YFilter
    pbbCore.EntityData.YangName = "pbb-core"
    pbbCore.EntityData.BundleName = "cisco_ios_xr"
    pbbCore.EntityData.ParentYangName = "bridge-domain-pbb"
    pbbCore.EntityData.SegmentPath = "pbb-core"
    pbbCore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCore.EntityData.Children = types.NewOrderedMap()
    pbbCore.EntityData.Children.Append("pbb-core-mac", types.YChild{"PbbCoreMac", &pbbCore.PbbCoreMac})
    pbbCore.EntityData.Children.Append("pbb-core-evis", types.YChild{"PbbCoreEvis", &pbbCore.PbbCoreEvis})
    pbbCore.EntityData.Children.Append("pbb-core-dhcp-profile", types.YChild{"PbbCoreDhcpProfile", &pbbCore.PbbCoreDhcpProfile})
    pbbCore.EntityData.Leafs = types.NewOrderedMap()
    pbbCore.EntityData.Leafs.Append("pbb-core-mmrp-flood-optimization", types.YLeaf{"PbbCoreMmrpFloodOptimization", pbbCore.PbbCoreMmrpFloodOptimization})
    pbbCore.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pbbCore.VlanId})
    pbbCore.EntityData.Leafs.Append("pbb-core-igmp-profile", types.YLeaf{"PbbCoreIgmpProfile", pbbCore.PbbCoreIgmpProfile})
    pbbCore.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pbbCore.Enable})

    pbbCore.EntityData.YListKeys = []string {}

    return &(pbbCore.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac
// MAC configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Mac Learning. The type is MacLearn.
    PbbCoreMacLearning interface{}

    // MAC-Aging configuration commands.
    PbbCoreMacAging L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging

    // MAC-Limit configuration commands.
    PbbCoreMacLimit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit
}

func (pbbCoreMac *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac) GetEntityData() *types.CommonEntityData {
    pbbCoreMac.EntityData.YFilter = pbbCoreMac.YFilter
    pbbCoreMac.EntityData.YangName = "pbb-core-mac"
    pbbCoreMac.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMac.EntityData.ParentYangName = "pbb-core"
    pbbCoreMac.EntityData.SegmentPath = "pbb-core-mac"
    pbbCoreMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMac.EntityData.Children = types.NewOrderedMap()
    pbbCoreMac.EntityData.Children.Append("pbb-core-mac-aging", types.YChild{"PbbCoreMacAging", &pbbCoreMac.PbbCoreMacAging})
    pbbCoreMac.EntityData.Children.Append("pbb-core-mac-limit", types.YChild{"PbbCoreMacLimit", &pbbCoreMac.PbbCoreMacLimit})
    pbbCoreMac.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreMac.EntityData.Leafs.Append("pbb-core-mac-learning", types.YLeaf{"PbbCoreMacLearning", pbbCoreMac.PbbCoreMacLearning})

    pbbCoreMac.EntityData.YListKeys = []string {}

    return &(pbbCoreMac.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging
// MAC-Aging configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PbbCoreMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    PbbCoreMacAgingTime interface{}
}

func (pbbCoreMacAging *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging) GetEntityData() *types.CommonEntityData {
    pbbCoreMacAging.EntityData.YFilter = pbbCoreMacAging.YFilter
    pbbCoreMacAging.EntityData.YangName = "pbb-core-mac-aging"
    pbbCoreMacAging.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMacAging.EntityData.ParentYangName = "pbb-core-mac"
    pbbCoreMacAging.EntityData.SegmentPath = "pbb-core-mac-aging"
    pbbCoreMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMacAging.EntityData.Children = types.NewOrderedMap()
    pbbCoreMacAging.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreMacAging.EntityData.Leafs.Append("pbb-core-mac-aging-type", types.YLeaf{"PbbCoreMacAgingType", pbbCoreMacAging.PbbCoreMacAgingType})
    pbbCoreMacAging.EntityData.Leafs.Append("pbb-core-mac-aging-time", types.YLeaf{"PbbCoreMacAgingTime", pbbCoreMacAging.PbbCoreMacAgingTime})

    pbbCoreMacAging.EntityData.YListKeys = []string {}

    return &(pbbCoreMacAging.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of MAC addresses after which MAC limit action is taken. The type is
    // interface{} with range: 0..4294967295.
    PbbCoreMacLimitMax interface{}

    // MAC address limit notification action. The type is MacNotification.
    PbbCoreMacLimitNotif interface{}

    // MAC address limit enforcement action. The type is MacLimitAction.
    PbbCoreMacLimitAction interface{}
}

func (pbbCoreMacLimit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit) GetEntityData() *types.CommonEntityData {
    pbbCoreMacLimit.EntityData.YFilter = pbbCoreMacLimit.YFilter
    pbbCoreMacLimit.EntityData.YangName = "pbb-core-mac-limit"
    pbbCoreMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMacLimit.EntityData.ParentYangName = "pbb-core-mac"
    pbbCoreMacLimit.EntityData.SegmentPath = "pbb-core-mac-limit"
    pbbCoreMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMacLimit.EntityData.Children = types.NewOrderedMap()
    pbbCoreMacLimit.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreMacLimit.EntityData.Leafs.Append("pbb-core-mac-limit-max", types.YLeaf{"PbbCoreMacLimitMax", pbbCoreMacLimit.PbbCoreMacLimitMax})
    pbbCoreMacLimit.EntityData.Leafs.Append("pbb-core-mac-limit-notif", types.YLeaf{"PbbCoreMacLimitNotif", pbbCoreMacLimit.PbbCoreMacLimitNotif})
    pbbCoreMacLimit.EntityData.Leafs.Append("pbb-core-mac-limit-action", types.YLeaf{"PbbCoreMacLimitAction", pbbCoreMacLimit.PbbCoreMacLimitAction})

    pbbCoreMacLimit.EntityData.YListKeys = []string {}

    return &(pbbCoreMacLimit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis
// PBB Core EVI Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Core EVI. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi.
    PbbCoreEvi []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi
}

func (pbbCoreEvis *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis) GetEntityData() *types.CommonEntityData {
    pbbCoreEvis.EntityData.YFilter = pbbCoreEvis.YFilter
    pbbCoreEvis.EntityData.YangName = "pbb-core-evis"
    pbbCoreEvis.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreEvis.EntityData.ParentYangName = "pbb-core"
    pbbCoreEvis.EntityData.SegmentPath = "pbb-core-evis"
    pbbCoreEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreEvis.EntityData.Children = types.NewOrderedMap()
    pbbCoreEvis.EntityData.Children.Append("pbb-core-evi", types.YChild{"PbbCoreEvi", nil})
    for i := range pbbCoreEvis.PbbCoreEvi {
        pbbCoreEvis.EntityData.Children.Append(types.GetSegmentPath(pbbCoreEvis.PbbCoreEvi[i]), types.YChild{"PbbCoreEvi", pbbCoreEvis.PbbCoreEvi[i]})
    }
    pbbCoreEvis.EntityData.Leafs = types.NewOrderedMap()

    pbbCoreEvis.EntityData.YListKeys = []string {}

    return &(pbbCoreEvis.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi
// PBB Core EVI
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..4294967295.
    Eviid interface{}
}

func (pbbCoreEvi *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi) GetEntityData() *types.CommonEntityData {
    pbbCoreEvi.EntityData.YFilter = pbbCoreEvi.YFilter
    pbbCoreEvi.EntityData.YangName = "pbb-core-evi"
    pbbCoreEvi.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreEvi.EntityData.ParentYangName = "pbb-core-evis"
    pbbCoreEvi.EntityData.SegmentPath = "pbb-core-evi" + types.AddKeyToken(pbbCoreEvi.Eviid, "eviid")
    pbbCoreEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreEvi.EntityData.Children = types.NewOrderedMap()
    pbbCoreEvi.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreEvi.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", pbbCoreEvi.Eviid})

    pbbCoreEvi.EntityData.YListKeys = []string {"Eviid"}

    return &(pbbCoreEvi.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile
// Attach a DHCP profile
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pbbCoreDhcpProfile *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile) GetEntityData() *types.CommonEntityData {
    pbbCoreDhcpProfile.EntityData.YFilter = pbbCoreDhcpProfile.YFilter
    pbbCoreDhcpProfile.EntityData.YangName = "pbb-core-dhcp-profile"
    pbbCoreDhcpProfile.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreDhcpProfile.EntityData.ParentYangName = "pbb-core"
    pbbCoreDhcpProfile.EntityData.SegmentPath = "pbb-core-dhcp-profile"
    pbbCoreDhcpProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreDhcpProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreDhcpProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreDhcpProfile.EntityData.Children = types.NewOrderedMap()
    pbbCoreDhcpProfile.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreDhcpProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", pbbCoreDhcpProfile.ProfileId})
    pbbCoreDhcpProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", pbbCoreDhcpProfile.DhcpSnoopingId})

    pbbCoreDhcpProfile.EntityData.YListKeys = []string {}

    return &(pbbCoreDhcpProfile.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis
// Bridge Domain EVI Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain MPLS EVPN. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi.
    BridgeDomainEvi []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi
}

func (bridgeDomainEvis *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis) GetEntityData() *types.CommonEntityData {
    bridgeDomainEvis.EntityData.YFilter = bridgeDomainEvis.YFilter
    bridgeDomainEvis.EntityData.YangName = "bridge-domain-evis"
    bridgeDomainEvis.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainEvis.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainEvis.EntityData.SegmentPath = "bridge-domain-evis"
    bridgeDomainEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainEvis.EntityData.Children = types.NewOrderedMap()
    bridgeDomainEvis.EntityData.Children.Append("bridge-domain-evi", types.YChild{"BridgeDomainEvi", nil})
    for i := range bridgeDomainEvis.BridgeDomainEvi {
        bridgeDomainEvis.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainEvis.BridgeDomainEvi[i]), types.YChild{"BridgeDomainEvi", bridgeDomainEvis.BridgeDomainEvi[i]})
    }
    bridgeDomainEvis.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainEvis.EntityData.YListKeys = []string {}

    return &(bridgeDomainEvis.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi
// Bridge Domain MPLS EVPN
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MPLS Ethernet VPN-ID. The type is interface{} with
    // range: 1..65534.
    VpnId interface{}
}

func (bridgeDomainEvi *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi) GetEntityData() *types.CommonEntityData {
    bridgeDomainEvi.EntityData.YFilter = bridgeDomainEvi.YFilter
    bridgeDomainEvi.EntityData.YangName = "bridge-domain-evi"
    bridgeDomainEvi.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainEvi.EntityData.ParentYangName = "bridge-domain-evis"
    bridgeDomainEvi.EntityData.SegmentPath = "bridge-domain-evi" + types.AddKeyToken(bridgeDomainEvi.VpnId, "vpn-id")
    bridgeDomainEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainEvi.EntityData.Children = types.NewOrderedMap()
    bridgeDomainEvi.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainEvi.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", bridgeDomainEvi.VpnId})

    bridgeDomainEvi.EntityData.YListKeys = []string {"VpnId"}

    return &(bridgeDomainEvi.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis
// Specify the access virtual forwarding
// interface name
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Acess Virtual Forwarding Interface. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi.
    AccessVfi []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi
}

func (accessVfis *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis) GetEntityData() *types.CommonEntityData {
    accessVfis.EntityData.YFilter = accessVfis.YFilter
    accessVfis.EntityData.YangName = "access-vfis"
    accessVfis.EntityData.BundleName = "cisco_ios_xr"
    accessVfis.EntityData.ParentYangName = "bridge-domain"
    accessVfis.EntityData.SegmentPath = "access-vfis"
    accessVfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfis.EntityData.Children = types.NewOrderedMap()
    accessVfis.EntityData.Children.Append("access-vfi", types.YChild{"AccessVfi", nil})
    for i := range accessVfis.AccessVfi {
        accessVfis.EntityData.Children.Append(types.GetSegmentPath(accessVfis.AccessVfi[i]), types.YChild{"AccessVfi", accessVfis.AccessVfi[i]})
    }
    accessVfis.EntityData.Leafs = types.NewOrderedMap()

    accessVfis.EntityData.YListKeys = []string {}

    return &(accessVfis.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi
// Name of the Acess Virtual Forwarding
// Interface
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the AccessVirtual Forwarding Interface.
    // The type is string with length: 1..32.
    Name interface{}

    // shutdown the AccessVfi. The type is interface{}.
    AccessVfiShutdown interface{}

    // List of pseudowires.
    AccessVfiPseudowires L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires
}

func (accessVfi *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi) GetEntityData() *types.CommonEntityData {
    accessVfi.EntityData.YFilter = accessVfi.YFilter
    accessVfi.EntityData.YangName = "access-vfi"
    accessVfi.EntityData.BundleName = "cisco_ios_xr"
    accessVfi.EntityData.ParentYangName = "access-vfis"
    accessVfi.EntityData.SegmentPath = "access-vfi" + types.AddKeyToken(accessVfi.Name, "name")
    accessVfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfi.EntityData.Children = types.NewOrderedMap()
    accessVfi.EntityData.Children.Append("access-vfi-pseudowires", types.YChild{"AccessVfiPseudowires", &accessVfi.AccessVfiPseudowires})
    accessVfi.EntityData.Leafs = types.NewOrderedMap()
    accessVfi.EntityData.Leafs.Append("name", types.YLeaf{"Name", accessVfi.Name})
    accessVfi.EntityData.Leafs.Append("access-vfi-shutdown", types.YLeaf{"AccessVfiShutdown", accessVfi.AccessVfiShutdown})

    accessVfi.EntityData.YListKeys = []string {"Name"}

    return &(accessVfi.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires
// List of pseudowires
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire.
    AccessVfiPseudowire []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire
}

func (accessVfiPseudowires *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowires.EntityData.YFilter = accessVfiPseudowires.YFilter
    accessVfiPseudowires.EntityData.YangName = "access-vfi-pseudowires"
    accessVfiPseudowires.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowires.EntityData.ParentYangName = "access-vfi"
    accessVfiPseudowires.EntityData.SegmentPath = "access-vfi-pseudowires"
    accessVfiPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowires.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowires.EntityData.Children.Append("access-vfi-pseudowire", types.YChild{"AccessVfiPseudowire", nil})
    for i := range accessVfiPseudowires.AccessVfiPseudowire {
        accessVfiPseudowires.EntityData.Children.Append(types.GetSegmentPath(accessVfiPseudowires.AccessVfiPseudowire[i]), types.YChild{"AccessVfiPseudowire", accessVfiPseudowires.AccessVfiPseudowire[i]})
    }
    accessVfiPseudowires.EntityData.Leafs = types.NewOrderedMap()

    accessVfiPseudowires.EntityData.YListKeys = []string {}

    return &(accessVfiPseudowires.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire
// Pseudowire configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // Pseudowire class template name to use for this pseudowire. The type is
    // string with length: 1..32.
    AccessVfiPwClass interface{}

    // Static Mac Address Table.
    AccessVfiPseudowireStaticMacAddresses L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses
}

func (accessVfiPseudowire *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowire.EntityData.YFilter = accessVfiPseudowire.YFilter
    accessVfiPseudowire.EntityData.YangName = "access-vfi-pseudowire"
    accessVfiPseudowire.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowire.EntityData.ParentYangName = "access-vfi-pseudowires"
    accessVfiPseudowire.EntityData.SegmentPath = "access-vfi-pseudowire" + types.AddKeyToken(accessVfiPseudowire.Neighbor, "neighbor") + types.AddKeyToken(accessVfiPseudowire.PseudowireId, "pseudowire-id")
    accessVfiPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowire.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowire.EntityData.Children.Append("access-vfi-pseudowire-static-mac-addresses", types.YChild{"AccessVfiPseudowireStaticMacAddresses", &accessVfiPseudowire.AccessVfiPseudowireStaticMacAddresses})
    accessVfiPseudowire.EntityData.Leafs = types.NewOrderedMap()
    accessVfiPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", accessVfiPseudowire.Neighbor})
    accessVfiPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", accessVfiPseudowire.PseudowireId})
    accessVfiPseudowire.EntityData.Leafs.Append("access-vfi-pw-class", types.YLeaf{"AccessVfiPwClass", accessVfiPseudowire.AccessVfiPwClass})

    accessVfiPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(accessVfiPseudowire.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress.
    AccessVfiPseudowireStaticMacAddress []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress
}

func (accessVfiPseudowireStaticMacAddresses *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowireStaticMacAddresses.EntityData.YFilter = accessVfiPseudowireStaticMacAddresses.YFilter
    accessVfiPseudowireStaticMacAddresses.EntityData.YangName = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowireStaticMacAddresses.EntityData.ParentYangName = "access-vfi-pseudowire"
    accessVfiPseudowireStaticMacAddresses.EntityData.SegmentPath = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowireStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowireStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowireStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowireStaticMacAddresses.EntityData.Children.Append("access-vfi-pseudowire-static-mac-address", types.YChild{"AccessVfiPseudowireStaticMacAddress", nil})
    for i := range accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress {
        accessVfiPseudowireStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress[i]), types.YChild{"AccessVfiPseudowireStaticMacAddress", accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress[i]})
    }
    accessVfiPseudowireStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    accessVfiPseudowireStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(accessVfiPseudowireStaticMacAddresses.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (accessVfiPseudowireStaticMacAddress *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowireStaticMacAddress.EntityData.YFilter = accessVfiPseudowireStaticMacAddress.YFilter
    accessVfiPseudowireStaticMacAddress.EntityData.YangName = "access-vfi-pseudowire-static-mac-address"
    accessVfiPseudowireStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowireStaticMacAddress.EntityData.ParentYangName = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddress.EntityData.SegmentPath = "access-vfi-pseudowire-static-mac-address" + types.AddKeyToken(accessVfiPseudowireStaticMacAddress.Address, "address")
    accessVfiPseudowireStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowireStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowireStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowireStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowireStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    accessVfiPseudowireStaticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", accessVfiPseudowireStaticMacAddress.Address})

    accessVfiPseudowireStaticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(accessVfiPseudowireStaticMacAddress.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires
// List of pseudowires
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire.
    BdPseudowire []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire
}

func (bdPseudowires *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires) GetEntityData() *types.CommonEntityData {
    bdPseudowires.EntityData.YFilter = bdPseudowires.YFilter
    bdPseudowires.EntityData.YangName = "bd-pseudowires"
    bdPseudowires.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowires.EntityData.ParentYangName = "bridge-domain"
    bdPseudowires.EntityData.SegmentPath = "bd-pseudowires"
    bdPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowires.EntityData.Children = types.NewOrderedMap()
    bdPseudowires.EntityData.Children.Append("bd-pseudowire", types.YChild{"BdPseudowire", nil})
    for i := range bdPseudowires.BdPseudowire {
        bdPseudowires.EntityData.Children.Append(types.GetSegmentPath(bdPseudowires.BdPseudowire[i]), types.YChild{"BdPseudowire", bdPseudowires.BdPseudowire[i]})
    }
    bdPseudowires.EntityData.Leafs = types.NewOrderedMap()

    bdPseudowires.EntityData.YListKeys = []string {}

    return &(bdPseudowires.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire
// Pseudowire configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // Attach a MLD Snooping profile. The type is string with length: 1..32.
    PseudowireMldSnoop interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    PseudowireIgmpSnoop interface{}

    // Bridge-domain Pseudowire flooding. The type is InterfaceTrafficFlood.
    PseudowireFlooding interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    BdPwClass interface{}

    // Bridge-domain Pseudowire flooding Unknown Unicast. The type is
    // InterfaceTrafficFlood.
    PseudowireFloodingUnknownUnicast interface{}

    // Access Pseudowire Dynamic ARP Inspection.
    PseudowireDai L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai

    // Storm Control.
    BdpwStormControlTypes L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes

    // Attach a DHCP profile.
    PseudowireProfile L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile

    // Static Mac Address Table.
    BdPwStaticMacAddresses L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses

    // IP Source Guard.
    PseudowireIpSourceGuard L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard

    // Bridge-domain Pseudowire MAC configuration commands.
    PseudowireMac L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac

    // Split Horizon.
    BdPwSplitHorizon L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon

    // MPLS static labels.
    BdPwMplsStaticLabels L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels

    // List of pseudowires.
    BridgeDomainBackupPseudowires L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires
}

func (bdPseudowire *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire) GetEntityData() *types.CommonEntityData {
    bdPseudowire.EntityData.YFilter = bdPseudowire.YFilter
    bdPseudowire.EntityData.YangName = "bd-pseudowire"
    bdPseudowire.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowire.EntityData.ParentYangName = "bd-pseudowires"
    bdPseudowire.EntityData.SegmentPath = "bd-pseudowire" + types.AddKeyToken(bdPseudowire.Neighbor, "neighbor") + types.AddKeyToken(bdPseudowire.PseudowireId, "pseudowire-id")
    bdPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowire.EntityData.Children = types.NewOrderedMap()
    bdPseudowire.EntityData.Children.Append("pseudowire-dai", types.YChild{"PseudowireDai", &bdPseudowire.PseudowireDai})
    bdPseudowire.EntityData.Children.Append("bdpw-storm-control-types", types.YChild{"BdpwStormControlTypes", &bdPseudowire.BdpwStormControlTypes})
    bdPseudowire.EntityData.Children.Append("pseudowire-profile", types.YChild{"PseudowireProfile", &bdPseudowire.PseudowireProfile})
    bdPseudowire.EntityData.Children.Append("bd-pw-static-mac-addresses", types.YChild{"BdPwStaticMacAddresses", &bdPseudowire.BdPwStaticMacAddresses})
    bdPseudowire.EntityData.Children.Append("pseudowire-ip-source-guard", types.YChild{"PseudowireIpSourceGuard", &bdPseudowire.PseudowireIpSourceGuard})
    bdPseudowire.EntityData.Children.Append("pseudowire-mac", types.YChild{"PseudowireMac", &bdPseudowire.PseudowireMac})
    bdPseudowire.EntityData.Children.Append("bd-pw-split-horizon", types.YChild{"BdPwSplitHorizon", &bdPseudowire.BdPwSplitHorizon})
    bdPseudowire.EntityData.Children.Append("bd-pw-mpls-static-labels", types.YChild{"BdPwMplsStaticLabels", &bdPseudowire.BdPwMplsStaticLabels})
    bdPseudowire.EntityData.Children.Append("bridge-domain-backup-pseudowires", types.YChild{"BridgeDomainBackupPseudowires", &bdPseudowire.BridgeDomainBackupPseudowires})
    bdPseudowire.EntityData.Leafs = types.NewOrderedMap()
    bdPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", bdPseudowire.Neighbor})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", bdPseudowire.PseudowireId})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-mld-snoop", types.YLeaf{"PseudowireMldSnoop", bdPseudowire.PseudowireMldSnoop})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-igmp-snoop", types.YLeaf{"PseudowireIgmpSnoop", bdPseudowire.PseudowireIgmpSnoop})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-flooding", types.YLeaf{"PseudowireFlooding", bdPseudowire.PseudowireFlooding})
    bdPseudowire.EntityData.Leafs.Append("bd-pw-class", types.YLeaf{"BdPwClass", bdPseudowire.BdPwClass})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-flooding-unknown-unicast", types.YLeaf{"PseudowireFloodingUnknownUnicast", bdPseudowire.PseudowireFloodingUnknownUnicast})

    bdPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(bdPseudowire.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai
// Access Pseudowire Dynamic ARP Inspection
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable Dynamic ARP Inspection. The type is interface{}.
    Disable interface{}

    // Enable Access Pseudowire Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    PseudowireDaiAddressValidation L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation
}

func (pseudowireDai *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai) GetEntityData() *types.CommonEntityData {
    pseudowireDai.EntityData.YFilter = pseudowireDai.YFilter
    pseudowireDai.EntityData.YangName = "pseudowire-dai"
    pseudowireDai.EntityData.BundleName = "cisco_ios_xr"
    pseudowireDai.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireDai.EntityData.SegmentPath = "pseudowire-dai"
    pseudowireDai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireDai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireDai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireDai.EntityData.Children = types.NewOrderedMap()
    pseudowireDai.EntityData.Children.Append("pseudowire-dai-address-validation", types.YChild{"PseudowireDaiAddressValidation", &pseudowireDai.PseudowireDaiAddressValidation})
    pseudowireDai.EntityData.Leafs = types.NewOrderedMap()
    pseudowireDai.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pseudowireDai.Logging})
    pseudowireDai.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pseudowireDai.Disable})
    pseudowireDai.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireDai.Enable})

    pseudowireDai.EntityData.YListKeys = []string {}

    return &(pseudowireDai.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation
// Address Validation
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Verification. The type is L2vpnVerification.
    Ipv4Verification interface{}

    // Destination MAC Verification. The type is L2vpnVerification.
    DestinationMacVerification interface{}

    // Source MAC Verification. The type is L2vpnVerification.
    SourceMacVerification interface{}
}

func (pseudowireDaiAddressValidation *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation) GetEntityData() *types.CommonEntityData {
    pseudowireDaiAddressValidation.EntityData.YFilter = pseudowireDaiAddressValidation.YFilter
    pseudowireDaiAddressValidation.EntityData.YangName = "pseudowire-dai-address-validation"
    pseudowireDaiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    pseudowireDaiAddressValidation.EntityData.ParentYangName = "pseudowire-dai"
    pseudowireDaiAddressValidation.EntityData.SegmentPath = "pseudowire-dai-address-validation"
    pseudowireDaiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireDaiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireDaiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireDaiAddressValidation.EntityData.Children = types.NewOrderedMap()
    pseudowireDaiAddressValidation.EntityData.Leafs = types.NewOrderedMap()
    pseudowireDaiAddressValidation.EntityData.Leafs.Append("ipv4-verification", types.YLeaf{"Ipv4Verification", pseudowireDaiAddressValidation.Ipv4Verification})
    pseudowireDaiAddressValidation.EntityData.Leafs.Append("destination-mac-verification", types.YLeaf{"DestinationMacVerification", pseudowireDaiAddressValidation.DestinationMacVerification})
    pseudowireDaiAddressValidation.EntityData.Leafs.Append("source-mac-verification", types.YLeaf{"SourceMacVerification", pseudowireDaiAddressValidation.SourceMacVerification})

    pseudowireDaiAddressValidation.EntityData.YListKeys = []string {}

    return &(pseudowireDaiAddressValidation.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes
// Storm Control
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType.
    BdpwStormControlType []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType
}

func (bdpwStormControlTypes *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes) GetEntityData() *types.CommonEntityData {
    bdpwStormControlTypes.EntityData.YFilter = bdpwStormControlTypes.YFilter
    bdpwStormControlTypes.EntityData.YangName = "bdpw-storm-control-types"
    bdpwStormControlTypes.EntityData.BundleName = "cisco_ios_xr"
    bdpwStormControlTypes.EntityData.ParentYangName = "bd-pseudowire"
    bdpwStormControlTypes.EntityData.SegmentPath = "bdpw-storm-control-types"
    bdpwStormControlTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdpwStormControlTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdpwStormControlTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdpwStormControlTypes.EntityData.Children = types.NewOrderedMap()
    bdpwStormControlTypes.EntityData.Children.Append("bdpw-storm-control-type", types.YChild{"BdpwStormControlType", nil})
    for i := range bdpwStormControlTypes.BdpwStormControlType {
        bdpwStormControlTypes.EntityData.Children.Append(types.GetSegmentPath(bdpwStormControlTypes.BdpwStormControlType[i]), types.YChild{"BdpwStormControlType", bdpwStormControlTypes.BdpwStormControlType[i]})
    }
    bdpwStormControlTypes.EntityData.Leafs = types.NewOrderedMap()

    bdpwStormControlTypes.EntityData.YListKeys = []string {}

    return &(bdpwStormControlTypes.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType
// Storm Control Type
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit
}

func (bdpwStormControlType *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType) GetEntityData() *types.CommonEntityData {
    bdpwStormControlType.EntityData.YFilter = bdpwStormControlType.YFilter
    bdpwStormControlType.EntityData.YangName = "bdpw-storm-control-type"
    bdpwStormControlType.EntityData.BundleName = "cisco_ios_xr"
    bdpwStormControlType.EntityData.ParentYangName = "bdpw-storm-control-types"
    bdpwStormControlType.EntityData.SegmentPath = "bdpw-storm-control-type" + types.AddKeyToken(bdpwStormControlType.Sctype, "sctype")
    bdpwStormControlType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdpwStormControlType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdpwStormControlType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdpwStormControlType.EntityData.Children = types.NewOrderedMap()
    bdpwStormControlType.EntityData.Children.Append("storm-control-unit", types.YChild{"StormControlUnit", &bdpwStormControlType.StormControlUnit})
    bdpwStormControlType.EntityData.Leafs = types.NewOrderedMap()
    bdpwStormControlType.EntityData.Leafs.Append("sctype", types.YLeaf{"Sctype", bdpwStormControlType.Sctype})

    bdpwStormControlType.EntityData.YListKeys = []string {"Sctype"}

    return &(bdpwStormControlType.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit
// Specify units for Storm Control Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Kilobits Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 64..1280000. Units are
    // kbit/s.
    KbitsPerSec interface{}

    // Packets Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 1..160000. Units are
    // packet/s.
    PktsPerSec interface{}
}

func (stormControlUnit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bdpw-storm-control-type"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs.Append("kbits-per-sec", types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec})
    stormControlUnit.EntityData.Leafs.Append("pkts-per-sec", types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec})

    stormControlUnit.EntityData.YListKeys = []string {}

    return &(stormControlUnit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile
// Attach a DHCP profile
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pseudowireProfile *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile) GetEntityData() *types.CommonEntityData {
    pseudowireProfile.EntityData.YFilter = pseudowireProfile.YFilter
    pseudowireProfile.EntityData.YangName = "pseudowire-profile"
    pseudowireProfile.EntityData.BundleName = "cisco_ios_xr"
    pseudowireProfile.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireProfile.EntityData.SegmentPath = "pseudowire-profile"
    pseudowireProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireProfile.EntityData.Children = types.NewOrderedMap()
    pseudowireProfile.EntityData.Leafs = types.NewOrderedMap()
    pseudowireProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", pseudowireProfile.ProfileId})
    pseudowireProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", pseudowireProfile.DhcpSnoopingId})

    pseudowireProfile.EntityData.YListKeys = []string {}

    return &(pseudowireProfile.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress.
    BdPwStaticMacAddress []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress
}

func (bdPwStaticMacAddresses *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    bdPwStaticMacAddresses.EntityData.YFilter = bdPwStaticMacAddresses.YFilter
    bdPwStaticMacAddresses.EntityData.YangName = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    bdPwStaticMacAddresses.EntityData.ParentYangName = "bd-pseudowire"
    bdPwStaticMacAddresses.EntityData.SegmentPath = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    bdPwStaticMacAddresses.EntityData.Children.Append("bd-pw-static-mac-address", types.YChild{"BdPwStaticMacAddress", nil})
    for i := range bdPwStaticMacAddresses.BdPwStaticMacAddress {
        bdPwStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(bdPwStaticMacAddresses.BdPwStaticMacAddress[i]), types.YChild{"BdPwStaticMacAddress", bdPwStaticMacAddresses.BdPwStaticMacAddress[i]})
    }
    bdPwStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    bdPwStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(bdPwStaticMacAddresses.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (bdPwStaticMacAddress *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress) GetEntityData() *types.CommonEntityData {
    bdPwStaticMacAddress.EntityData.YFilter = bdPwStaticMacAddress.YFilter
    bdPwStaticMacAddress.EntityData.YangName = "bd-pw-static-mac-address"
    bdPwStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    bdPwStaticMacAddress.EntityData.ParentYangName = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddress.EntityData.SegmentPath = "bd-pw-static-mac-address" + types.AddKeyToken(bdPwStaticMacAddress.Address, "address")
    bdPwStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    bdPwStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    bdPwStaticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", bdPwStaticMacAddress.Address})

    bdPwStaticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(bdPwStaticMacAddress.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard
// IP Source Guard
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable Dynamic IP source guard. The type is interface{}.
    Disable interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (pseudowireIpSourceGuard *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard) GetEntityData() *types.CommonEntityData {
    pseudowireIpSourceGuard.EntityData.YFilter = pseudowireIpSourceGuard.YFilter
    pseudowireIpSourceGuard.EntityData.YangName = "pseudowire-ip-source-guard"
    pseudowireIpSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    pseudowireIpSourceGuard.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireIpSourceGuard.EntityData.SegmentPath = "pseudowire-ip-source-guard"
    pseudowireIpSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireIpSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireIpSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireIpSourceGuard.EntityData.Children = types.NewOrderedMap()
    pseudowireIpSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    pseudowireIpSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pseudowireIpSourceGuard.Logging})
    pseudowireIpSourceGuard.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pseudowireIpSourceGuard.Disable})
    pseudowireIpSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireIpSourceGuard.Enable})

    pseudowireIpSourceGuard.EntityData.YListKeys = []string {}

    return &(pseudowireIpSourceGuard.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac
// Bridge-domain Pseudowire MAC configuration
// commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MAC Flush When Port goes down. The type is PortDownFlush.
    PseudowireMacPortDownFlush interface{}

    // Bridge-domain Pseudowire MAC configuration mode. The type is interface{}.
    Enable interface{}

    // Enable MAC Learning. The type is MacLearn.
    PseudowireMacLearning interface{}

    // MAC Secure.
    PseudowireMacSecure L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure

    // MAC-Aging configuration commands.
    PseudowireMacAging L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging

    // MAC-Limit configuration commands.
    PseudowireMacLimit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit
}

func (pseudowireMac *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac) GetEntityData() *types.CommonEntityData {
    pseudowireMac.EntityData.YFilter = pseudowireMac.YFilter
    pseudowireMac.EntityData.YangName = "pseudowire-mac"
    pseudowireMac.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMac.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireMac.EntityData.SegmentPath = "pseudowire-mac"
    pseudowireMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMac.EntityData.Children = types.NewOrderedMap()
    pseudowireMac.EntityData.Children.Append("pseudowire-mac-secure", types.YChild{"PseudowireMacSecure", &pseudowireMac.PseudowireMacSecure})
    pseudowireMac.EntityData.Children.Append("pseudowire-mac-aging", types.YChild{"PseudowireMacAging", &pseudowireMac.PseudowireMacAging})
    pseudowireMac.EntityData.Children.Append("pseudowire-mac-limit", types.YChild{"PseudowireMacLimit", &pseudowireMac.PseudowireMacLimit})
    pseudowireMac.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMac.EntityData.Leafs.Append("pseudowire-mac-port-down-flush", types.YLeaf{"PseudowireMacPortDownFlush", pseudowireMac.PseudowireMacPortDownFlush})
    pseudowireMac.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireMac.Enable})
    pseudowireMac.EntityData.Leafs.Append("pseudowire-mac-learning", types.YLeaf{"PseudowireMacLearning", pseudowireMac.PseudowireMacLearning})

    pseudowireMac.EntityData.YListKeys = []string {}

    return &(pseudowireMac.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure
// MAC Secure
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Pseudowire MAC Secure. The type is interface{}.
    Disable interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}
}

func (pseudowireMacSecure *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure) GetEntityData() *types.CommonEntityData {
    pseudowireMacSecure.EntityData.YFilter = pseudowireMacSecure.YFilter
    pseudowireMacSecure.EntityData.YangName = "pseudowire-mac-secure"
    pseudowireMacSecure.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacSecure.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacSecure.EntityData.SegmentPath = "pseudowire-mac-secure"
    pseudowireMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacSecure.EntityData.Children = types.NewOrderedMap()
    pseudowireMacSecure.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMacSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pseudowireMacSecure.Logging})
    pseudowireMacSecure.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pseudowireMacSecure.Disable})
    pseudowireMacSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", pseudowireMacSecure.Action})
    pseudowireMacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireMacSecure.Enable})

    pseudowireMacSecure.EntityData.YListKeys = []string {}

    return &(pseudowireMacSecure.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging
// MAC-Aging configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PseudowireMacAgingType interface{}

    // MAC Aging Time. The type is interface{} with range: 300..30000.
    PseudowireMacAgingTime interface{}
}

func (pseudowireMacAging *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging) GetEntityData() *types.CommonEntityData {
    pseudowireMacAging.EntityData.YFilter = pseudowireMacAging.YFilter
    pseudowireMacAging.EntityData.YangName = "pseudowire-mac-aging"
    pseudowireMacAging.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacAging.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacAging.EntityData.SegmentPath = "pseudowire-mac-aging"
    pseudowireMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacAging.EntityData.Children = types.NewOrderedMap()
    pseudowireMacAging.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMacAging.EntityData.Leafs.Append("pseudowire-mac-aging-type", types.YLeaf{"PseudowireMacAgingType", pseudowireMacAging.PseudowireMacAgingType})
    pseudowireMacAging.EntityData.Leafs.Append("pseudowire-mac-aging-time", types.YLeaf{"PseudowireMacAgingTime", pseudowireMacAging.PseudowireMacAgingTime})

    pseudowireMacAging.EntityData.YListKeys = []string {}

    return &(pseudowireMacAging.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Access Pseudowire MAC address limit enforcement action. The type is
    // MacLimitAction.
    PseudowireMacLimitAction interface{}

    // MAC address limit notification action in a Bridge Access Pseudowire. The
    // type is MacNotification.
    PseudowireMacLimitNotif interface{}

    // Number of MAC addresses on a Bridge Access Pseudowire after which MAC limit
    // action is taken. The type is interface{} with range: 0..4294967295.
    PseudowireMacLimitMax interface{}
}

func (pseudowireMacLimit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit) GetEntityData() *types.CommonEntityData {
    pseudowireMacLimit.EntityData.YFilter = pseudowireMacLimit.YFilter
    pseudowireMacLimit.EntityData.YangName = "pseudowire-mac-limit"
    pseudowireMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacLimit.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacLimit.EntityData.SegmentPath = "pseudowire-mac-limit"
    pseudowireMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacLimit.EntityData.Children = types.NewOrderedMap()
    pseudowireMacLimit.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMacLimit.EntityData.Leafs.Append("pseudowire-mac-limit-action", types.YLeaf{"PseudowireMacLimitAction", pseudowireMacLimit.PseudowireMacLimitAction})
    pseudowireMacLimit.EntityData.Leafs.Append("pseudowire-mac-limit-notif", types.YLeaf{"PseudowireMacLimitNotif", pseudowireMacLimit.PseudowireMacLimitNotif})
    pseudowireMacLimit.EntityData.Leafs.Append("pseudowire-mac-limit-max", types.YLeaf{"PseudowireMacLimitMax", pseudowireMacLimit.PseudowireMacLimitMax})

    pseudowireMacLimit.EntityData.YListKeys = []string {}

    return &(pseudowireMacLimit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon
// Split Horizon
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon Group.
    BdPwSplitHorizonGroup L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup
}

func (bdPwSplitHorizon *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon) GetEntityData() *types.CommonEntityData {
    bdPwSplitHorizon.EntityData.YFilter = bdPwSplitHorizon.YFilter
    bdPwSplitHorizon.EntityData.YangName = "bd-pw-split-horizon"
    bdPwSplitHorizon.EntityData.BundleName = "cisco_ios_xr"
    bdPwSplitHorizon.EntityData.ParentYangName = "bd-pseudowire"
    bdPwSplitHorizon.EntityData.SegmentPath = "bd-pw-split-horizon"
    bdPwSplitHorizon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwSplitHorizon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwSplitHorizon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwSplitHorizon.EntityData.Children = types.NewOrderedMap()
    bdPwSplitHorizon.EntityData.Children.Append("bd-pw-split-horizon-group", types.YChild{"BdPwSplitHorizonGroup", &bdPwSplitHorizon.BdPwSplitHorizonGroup})
    bdPwSplitHorizon.EntityData.Leafs = types.NewOrderedMap()

    bdPwSplitHorizon.EntityData.YListKeys = []string {}

    return &(bdPwSplitHorizon.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup
// Split Horizon Group
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable split horizon group. The type is interface{}.
    Enable interface{}
}

func (bdPwSplitHorizonGroup *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    bdPwSplitHorizonGroup.EntityData.YFilter = bdPwSplitHorizonGroup.YFilter
    bdPwSplitHorizonGroup.EntityData.YangName = "bd-pw-split-horizon-group"
    bdPwSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    bdPwSplitHorizonGroup.EntityData.ParentYangName = "bd-pw-split-horizon"
    bdPwSplitHorizonGroup.EntityData.SegmentPath = "bd-pw-split-horizon-group"
    bdPwSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwSplitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    bdPwSplitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    bdPwSplitHorizonGroup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bdPwSplitHorizonGroup.Enable})

    bdPwSplitHorizonGroup.EntityData.YListKeys = []string {}

    return &(bdPwSplitHorizonGroup.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels
// MPLS static labels
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (bdPwMplsStaticLabels *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    bdPwMplsStaticLabels.EntityData.YFilter = bdPwMplsStaticLabels.YFilter
    bdPwMplsStaticLabels.EntityData.YangName = "bd-pw-mpls-static-labels"
    bdPwMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    bdPwMplsStaticLabels.EntityData.ParentYangName = "bd-pseudowire"
    bdPwMplsStaticLabels.EntityData.SegmentPath = "bd-pw-mpls-static-labels"
    bdPwMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwMplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    bdPwMplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    bdPwMplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", bdPwMplsStaticLabels.LocalStaticLabel})
    bdPwMplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", bdPwMplsStaticLabels.RemoteStaticLabel})

    bdPwMplsStaticLabels.EntityData.YListKeys = []string {}

    return &(bdPwMplsStaticLabels.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires
// List of pseudowires
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire.
    BridgeDomainBackupPseudowire []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire
}

func (bridgeDomainBackupPseudowires *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires) GetEntityData() *types.CommonEntityData {
    bridgeDomainBackupPseudowires.EntityData.YFilter = bridgeDomainBackupPseudowires.YFilter
    bridgeDomainBackupPseudowires.EntityData.YangName = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainBackupPseudowires.EntityData.ParentYangName = "bd-pseudowire"
    bridgeDomainBackupPseudowires.EntityData.SegmentPath = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainBackupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainBackupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainBackupPseudowires.EntityData.Children = types.NewOrderedMap()
    bridgeDomainBackupPseudowires.EntityData.Children.Append("bridge-domain-backup-pseudowire", types.YChild{"BridgeDomainBackupPseudowire", nil})
    for i := range bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire {
        bridgeDomainBackupPseudowires.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire[i]), types.YChild{"BridgeDomainBackupPseudowire", bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire[i]})
    }
    bridgeDomainBackupPseudowires.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainBackupPseudowires.EntityData.YListKeys = []string {}

    return &(bridgeDomainBackupPseudowires.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire
// Backup pseudowire configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    BridgeDomainBackupPwClass interface{}
}

func (bridgeDomainBackupPseudowire *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire) GetEntityData() *types.CommonEntityData {
    bridgeDomainBackupPseudowire.EntityData.YFilter = bridgeDomainBackupPseudowire.YFilter
    bridgeDomainBackupPseudowire.EntityData.YangName = "bridge-domain-backup-pseudowire"
    bridgeDomainBackupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainBackupPseudowire.EntityData.ParentYangName = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowire.EntityData.SegmentPath = "bridge-domain-backup-pseudowire" + types.AddKeyToken(bridgeDomainBackupPseudowire.Neighbor, "neighbor") + types.AddKeyToken(bridgeDomainBackupPseudowire.PseudowireId, "pseudowire-id")
    bridgeDomainBackupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainBackupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainBackupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainBackupPseudowire.EntityData.Children = types.NewOrderedMap()
    bridgeDomainBackupPseudowire.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainBackupPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", bridgeDomainBackupPseudowire.Neighbor})
    bridgeDomainBackupPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", bridgeDomainBackupPseudowire.PseudowireId})
    bridgeDomainBackupPseudowire.EntityData.Leafs.Append("bridge-domain-backup-pw-class", types.YLeaf{"BridgeDomainBackupPwClass", bridgeDomainBackupPseudowire.BridgeDomainBackupPwClass})

    bridgeDomainBackupPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(bridgeDomainBackupPseudowire.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis
// Specify the virtual forwarding interface name
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Virtual Forwarding Interface. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi.
    Vfi []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi
}

func (vfis *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis) GetEntityData() *types.CommonEntityData {
    vfis.EntityData.YFilter = vfis.YFilter
    vfis.EntityData.YangName = "vfis"
    vfis.EntityData.BundleName = "cisco_ios_xr"
    vfis.EntityData.ParentYangName = "bridge-domain"
    vfis.EntityData.SegmentPath = "vfis"
    vfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfis.EntityData.Children = types.NewOrderedMap()
    vfis.EntityData.Children.Append("vfi", types.YChild{"Vfi", nil})
    for i := range vfis.Vfi {
        vfis.EntityData.Children.Append(types.GetSegmentPath(vfis.Vfi[i]), types.YChild{"Vfi", vfis.Vfi[i]})
    }
    vfis.EntityData.Leafs = types.NewOrderedMap()

    vfis.EntityData.YListKeys = []string {}

    return &(vfis.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi
// Name of the Virtual Forwarding Interface
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Virtual Forwarding Interface. The type
    // is string with length: 1..32.
    Name interface{}

    // Enabling Shutdown. The type is interface{}.
    VfiShutdown interface{}

    // VPN Identifier. The type is interface{} with range: 1..4294967295.
    Vpnid interface{}

    // Enable Multicast P2MP in this VFI.
    MulticastP2mp L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp

    // List of pseudowires.
    VfiPseudowires L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires

    // Enable Autodiscovery BGP in this VFI.
    BgpAutoDiscovery L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery
}

func (vfi *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi) GetEntityData() *types.CommonEntityData {
    vfi.EntityData.YFilter = vfi.YFilter
    vfi.EntityData.YangName = "vfi"
    vfi.EntityData.BundleName = "cisco_ios_xr"
    vfi.EntityData.ParentYangName = "vfis"
    vfi.EntityData.SegmentPath = "vfi" + types.AddKeyToken(vfi.Name, "name")
    vfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfi.EntityData.Children = types.NewOrderedMap()
    vfi.EntityData.Children.Append("multicast-p2mp", types.YChild{"MulticastP2mp", &vfi.MulticastP2mp})
    vfi.EntityData.Children.Append("vfi-pseudowires", types.YChild{"VfiPseudowires", &vfi.VfiPseudowires})
    vfi.EntityData.Children.Append("bgp-auto-discovery", types.YChild{"BgpAutoDiscovery", &vfi.BgpAutoDiscovery})
    vfi.EntityData.Leafs = types.NewOrderedMap()
    vfi.EntityData.Leafs.Append("name", types.YLeaf{"Name", vfi.Name})
    vfi.EntityData.Leafs.Append("vfi-shutdown", types.YLeaf{"VfiShutdown", vfi.VfiShutdown})
    vfi.EntityData.Leafs.Append("vpnid", types.YLeaf{"Vpnid", vfi.Vpnid})

    vfi.EntityData.YListKeys = []string {"Name"}

    return &(vfi.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp
// Enable Multicast P2MP in this VFI
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery P2MP. The type is interface{}.
    Enable interface{}

    // Multicast P2MP Transport.
    Transports L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports

    // Multicast P2MP Signaling Type.
    Signalings L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings
}

func (multicastP2mp *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp) GetEntityData() *types.CommonEntityData {
    multicastP2mp.EntityData.YFilter = multicastP2mp.YFilter
    multicastP2mp.EntityData.YangName = "multicast-p2mp"
    multicastP2mp.EntityData.BundleName = "cisco_ios_xr"
    multicastP2mp.EntityData.ParentYangName = "vfi"
    multicastP2mp.EntityData.SegmentPath = "multicast-p2mp"
    multicastP2mp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastP2mp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastP2mp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastP2mp.EntityData.Children = types.NewOrderedMap()
    multicastP2mp.EntityData.Children.Append("transports", types.YChild{"Transports", &multicastP2mp.Transports})
    multicastP2mp.EntityData.Children.Append("signalings", types.YChild{"Signalings", &multicastP2mp.Signalings})
    multicastP2mp.EntityData.Leafs = types.NewOrderedMap()
    multicastP2mp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", multicastP2mp.Enable})

    multicastP2mp.EntityData.YListKeys = []string {}

    return &(multicastP2mp.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports
// Multicast P2MP Transport
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multicast P2MP Transport Type. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport.
    Transport []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport
}

func (transports *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports) GetEntityData() *types.CommonEntityData {
    transports.EntityData.YFilter = transports.YFilter
    transports.EntityData.YangName = "transports"
    transports.EntityData.BundleName = "cisco_ios_xr"
    transports.EntityData.ParentYangName = "multicast-p2mp"
    transports.EntityData.SegmentPath = "transports"
    transports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transports.EntityData.Children = types.NewOrderedMap()
    transports.EntityData.Children.Append("transport", types.YChild{"Transport", nil})
    for i := range transports.Transport {
        transports.EntityData.Children.Append(types.GetSegmentPath(transports.Transport[i]), types.YChild{"Transport", transports.Transport[i]})
    }
    transports.EntityData.Leafs = types.NewOrderedMap()

    transports.EntityData.YListKeys = []string {}

    return &(transports.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport
// Multicast P2MP Transport Type
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Transport Type. The type is string with pattern:
    // (RSVP_TE).
    TransportName interface{}

    // Multicast P2MP TE Attribute Set Name. The type is string with length:
    // 1..64.
    AttributeSetName interface{}
}

func (transport *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xr"
    transport.EntityData.ParentYangName = "transports"
    transport.EntityData.SegmentPath = "transport" + types.AddKeyToken(transport.TransportName, "transport-name")
    transport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Leafs = types.NewOrderedMap()
    transport.EntityData.Leafs.Append("transport-name", types.YLeaf{"TransportName", transport.TransportName})
    transport.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", transport.AttributeSetName})

    transport.EntityData.YListKeys = []string {"TransportName"}

    return &(transport.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings
// Multicast P2MP Signaling Type
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multicast P2MP Signaling Type. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling.
    Signaling []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling
}

func (signalings *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings) GetEntityData() *types.CommonEntityData {
    signalings.EntityData.YFilter = signalings.YFilter
    signalings.EntityData.YangName = "signalings"
    signalings.EntityData.BundleName = "cisco_ios_xr"
    signalings.EntityData.ParentYangName = "multicast-p2mp"
    signalings.EntityData.SegmentPath = "signalings"
    signalings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalings.EntityData.Children = types.NewOrderedMap()
    signalings.EntityData.Children.Append("signaling", types.YChild{"Signaling", nil})
    for i := range signalings.Signaling {
        signalings.EntityData.Children.Append(types.GetSegmentPath(signalings.Signaling[i]), types.YChild{"Signaling", signalings.Signaling[i]})
    }
    signalings.EntityData.Leafs = types.NewOrderedMap()

    signalings.EntityData.YListKeys = []string {}

    return &(signalings.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling
// Multicast P2MP Signaling Type
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Signaling Type. The type is string with pattern:
    // (BGP).
    SignalingName interface{}
}

func (signaling *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling) GetEntityData() *types.CommonEntityData {
    signaling.EntityData.YFilter = signaling.YFilter
    signaling.EntityData.YangName = "signaling"
    signaling.EntityData.BundleName = "cisco_ios_xr"
    signaling.EntityData.ParentYangName = "signalings"
    signaling.EntityData.SegmentPath = "signaling" + types.AddKeyToken(signaling.SignalingName, "signaling-name")
    signaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaling.EntityData.Children = types.NewOrderedMap()
    signaling.EntityData.Leafs = types.NewOrderedMap()
    signaling.EntityData.Leafs.Append("signaling-name", types.YLeaf{"SignalingName", signaling.SignalingName})

    signaling.EntityData.YListKeys = []string {"SignalingName"}

    return &(signaling.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires
// List of pseudowires
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire.
    VfiPseudowire []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire
}

func (vfiPseudowires *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires) GetEntityData() *types.CommonEntityData {
    vfiPseudowires.EntityData.YFilter = vfiPseudowires.YFilter
    vfiPseudowires.EntityData.YangName = "vfi-pseudowires"
    vfiPseudowires.EntityData.BundleName = "cisco_ios_xr"
    vfiPseudowires.EntityData.ParentYangName = "vfi"
    vfiPseudowires.EntityData.SegmentPath = "vfi-pseudowires"
    vfiPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPseudowires.EntityData.Children = types.NewOrderedMap()
    vfiPseudowires.EntityData.Children.Append("vfi-pseudowire", types.YChild{"VfiPseudowire", nil})
    for i := range vfiPseudowires.VfiPseudowire {
        vfiPseudowires.EntityData.Children.Append(types.GetSegmentPath(vfiPseudowires.VfiPseudowire[i]), types.YChild{"VfiPseudowire", vfiPseudowires.VfiPseudowire[i]})
    }
    vfiPseudowires.EntityData.Leafs = types.NewOrderedMap()

    vfiPseudowires.EntityData.YListKeys = []string {}

    return &(vfiPseudowires.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire
// Pseudowire configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    VfiPwClass interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    VfiPwIgmpSnoop interface{}

    // Attach a MLD Snooping profile. The type is string with length: 1..32.
    VfiPwMldSnoop interface{}

    // Attach a DHCP Snooping profile.
    VfiPwDhcpSnoop L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop

    // MPLS static labels.
    VfiPwMplsStaticLabels L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels

    // Static Mac Address Table.
    PseudowireStaticMacAddresses L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses
}

func (vfiPseudowire *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire) GetEntityData() *types.CommonEntityData {
    vfiPseudowire.EntityData.YFilter = vfiPseudowire.YFilter
    vfiPseudowire.EntityData.YangName = "vfi-pseudowire"
    vfiPseudowire.EntityData.BundleName = "cisco_ios_xr"
    vfiPseudowire.EntityData.ParentYangName = "vfi-pseudowires"
    vfiPseudowire.EntityData.SegmentPath = "vfi-pseudowire" + types.AddKeyToken(vfiPseudowire.Neighbor, "neighbor") + types.AddKeyToken(vfiPseudowire.PseudowireId, "pseudowire-id")
    vfiPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPseudowire.EntityData.Children = types.NewOrderedMap()
    vfiPseudowire.EntityData.Children.Append("vfi-pw-dhcp-snoop", types.YChild{"VfiPwDhcpSnoop", &vfiPseudowire.VfiPwDhcpSnoop})
    vfiPseudowire.EntityData.Children.Append("vfi-pw-mpls-static-labels", types.YChild{"VfiPwMplsStaticLabels", &vfiPseudowire.VfiPwMplsStaticLabels})
    vfiPseudowire.EntityData.Children.Append("pseudowire-static-mac-addresses", types.YChild{"PseudowireStaticMacAddresses", &vfiPseudowire.PseudowireStaticMacAddresses})
    vfiPseudowire.EntityData.Leafs = types.NewOrderedMap()
    vfiPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", vfiPseudowire.Neighbor})
    vfiPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", vfiPseudowire.PseudowireId})
    vfiPseudowire.EntityData.Leafs.Append("vfi-pw-class", types.YLeaf{"VfiPwClass", vfiPseudowire.VfiPwClass})
    vfiPseudowire.EntityData.Leafs.Append("vfi-pw-igmp-snoop", types.YLeaf{"VfiPwIgmpSnoop", vfiPseudowire.VfiPwIgmpSnoop})
    vfiPseudowire.EntityData.Leafs.Append("vfi-pw-mld-snoop", types.YLeaf{"VfiPwMldSnoop", vfiPseudowire.VfiPwMldSnoop})

    vfiPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(vfiPseudowire.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop
// Attach a DHCP Snooping profile
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (vfiPwDhcpSnoop *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop) GetEntityData() *types.CommonEntityData {
    vfiPwDhcpSnoop.EntityData.YFilter = vfiPwDhcpSnoop.YFilter
    vfiPwDhcpSnoop.EntityData.YangName = "vfi-pw-dhcp-snoop"
    vfiPwDhcpSnoop.EntityData.BundleName = "cisco_ios_xr"
    vfiPwDhcpSnoop.EntityData.ParentYangName = "vfi-pseudowire"
    vfiPwDhcpSnoop.EntityData.SegmentPath = "vfi-pw-dhcp-snoop"
    vfiPwDhcpSnoop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPwDhcpSnoop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPwDhcpSnoop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPwDhcpSnoop.EntityData.Children = types.NewOrderedMap()
    vfiPwDhcpSnoop.EntityData.Leafs = types.NewOrderedMap()
    vfiPwDhcpSnoop.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", vfiPwDhcpSnoop.ProfileId})
    vfiPwDhcpSnoop.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", vfiPwDhcpSnoop.DhcpSnoopingId})

    vfiPwDhcpSnoop.EntityData.YListKeys = []string {}

    return &(vfiPwDhcpSnoop.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels
// MPLS static labels
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (vfiPwMplsStaticLabels *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    vfiPwMplsStaticLabels.EntityData.YFilter = vfiPwMplsStaticLabels.YFilter
    vfiPwMplsStaticLabels.EntityData.YangName = "vfi-pw-mpls-static-labels"
    vfiPwMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    vfiPwMplsStaticLabels.EntityData.ParentYangName = "vfi-pseudowire"
    vfiPwMplsStaticLabels.EntityData.SegmentPath = "vfi-pw-mpls-static-labels"
    vfiPwMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPwMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPwMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPwMplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    vfiPwMplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    vfiPwMplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", vfiPwMplsStaticLabels.LocalStaticLabel})
    vfiPwMplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", vfiPwMplsStaticLabels.RemoteStaticLabel})

    vfiPwMplsStaticLabels.EntityData.YListKeys = []string {}

    return &(vfiPwMplsStaticLabels.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress.
    PseudowireStaticMacAddress []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress
}

func (pseudowireStaticMacAddresses *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    pseudowireStaticMacAddresses.EntityData.YFilter = pseudowireStaticMacAddresses.YFilter
    pseudowireStaticMacAddresses.EntityData.YangName = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    pseudowireStaticMacAddresses.EntityData.ParentYangName = "vfi-pseudowire"
    pseudowireStaticMacAddresses.EntityData.SegmentPath = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    pseudowireStaticMacAddresses.EntityData.Children.Append("pseudowire-static-mac-address", types.YChild{"PseudowireStaticMacAddress", nil})
    for i := range pseudowireStaticMacAddresses.PseudowireStaticMacAddress {
        pseudowireStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(pseudowireStaticMacAddresses.PseudowireStaticMacAddress[i]), types.YChild{"PseudowireStaticMacAddress", pseudowireStaticMacAddresses.PseudowireStaticMacAddress[i]})
    }
    pseudowireStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    pseudowireStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(pseudowireStaticMacAddresses.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (pseudowireStaticMacAddress *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress) GetEntityData() *types.CommonEntityData {
    pseudowireStaticMacAddress.EntityData.YFilter = pseudowireStaticMacAddress.YFilter
    pseudowireStaticMacAddress.EntityData.YangName = "pseudowire-static-mac-address"
    pseudowireStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    pseudowireStaticMacAddress.EntityData.ParentYangName = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddress.EntityData.SegmentPath = "pseudowire-static-mac-address" + types.AddKeyToken(pseudowireStaticMacAddress.Address, "address")
    pseudowireStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    pseudowireStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    pseudowireStaticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", pseudowireStaticMacAddress.Address})

    pseudowireStaticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(pseudowireStaticMacAddress.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery
// Enable Autodiscovery BGP in this VFI
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Policy for installation of forwarding data to L2FIB. The type is
    // string.
    TablePolicy interface{}

    // Enable control-word for this VFI. The type is interface{}.
    AdControlWord interface{}

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Signaling Protocol LDP in this VFI configuration.
    LdpSignalingProtocol L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol

    // Route policy.
    BgpRoutePolicy L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy

    // Route Distinguisher.
    RouteDistinguisher L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher

    // Enable Signaling Protocol BGP in this VFI.
    BgpSignalingProtocol L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol

    // Route Target.
    RouteTargets L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets
}

func (bgpAutoDiscovery *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    bgpAutoDiscovery.EntityData.YFilter = bgpAutoDiscovery.YFilter
    bgpAutoDiscovery.EntityData.YangName = "bgp-auto-discovery"
    bgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    bgpAutoDiscovery.EntityData.ParentYangName = "vfi"
    bgpAutoDiscovery.EntityData.SegmentPath = "bgp-auto-discovery"
    bgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    bgpAutoDiscovery.EntityData.Children.Append("ldp-signaling-protocol", types.YChild{"LdpSignalingProtocol", &bgpAutoDiscovery.LdpSignalingProtocol})
    bgpAutoDiscovery.EntityData.Children.Append("bgp-route-policy", types.YChild{"BgpRoutePolicy", &bgpAutoDiscovery.BgpRoutePolicy})
    bgpAutoDiscovery.EntityData.Children.Append("route-distinguisher", types.YChild{"RouteDistinguisher", &bgpAutoDiscovery.RouteDistinguisher})
    bgpAutoDiscovery.EntityData.Children.Append("bgp-signaling-protocol", types.YChild{"BgpSignalingProtocol", &bgpAutoDiscovery.BgpSignalingProtocol})
    bgpAutoDiscovery.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &bgpAutoDiscovery.RouteTargets})
    bgpAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    bgpAutoDiscovery.EntityData.Leafs.Append("table-policy", types.YLeaf{"TablePolicy", bgpAutoDiscovery.TablePolicy})
    bgpAutoDiscovery.EntityData.Leafs.Append("ad-control-word", types.YLeaf{"AdControlWord", bgpAutoDiscovery.AdControlWord})
    bgpAutoDiscovery.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bgpAutoDiscovery.Enable})

    bgpAutoDiscovery.EntityData.YListKeys = []string {}

    return &(bgpAutoDiscovery.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol
// Signaling Protocol LDP in this VFI
// configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable LDP as Signaling Protocol.Deletion of this object also causes
    // deletion of all objects under LDPSignalingProtocol. The type is
    // interface{}.
    Enable interface{}

    // VPLS ID.
    VplsId L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance
}

func (ldpSignalingProtocol *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    ldpSignalingProtocol.EntityData.YFilter = ldpSignalingProtocol.YFilter
    ldpSignalingProtocol.EntityData.YangName = "ldp-signaling-protocol"
    ldpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    ldpSignalingProtocol.EntityData.ParentYangName = "bgp-auto-discovery"
    ldpSignalingProtocol.EntityData.SegmentPath = "ldp-signaling-protocol"
    ldpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpSignalingProtocol.EntityData.Children = types.NewOrderedMap()
    ldpSignalingProtocol.EntityData.Children.Append("vpls-id", types.YChild{"VplsId", &ldpSignalingProtocol.VplsId})
    ldpSignalingProtocol.EntityData.Children.Append("flow-label-load-balance", types.YChild{"FlowLabelLoadBalance", &ldpSignalingProtocol.FlowLabelLoadBalance})
    ldpSignalingProtocol.EntityData.Leafs = types.NewOrderedMap()
    ldpSignalingProtocol.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ldpSignalingProtocol.Enable})

    ldpSignalingProtocol.EntityData.YListKeys = []string {}

    return &(ldpSignalingProtocol.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId
// VPLS ID
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPLS-ID Type. The type is LdpVplsId.
    Type interface{}

    // Two byte AS number. The type is interface{} with range: 1..65535.
    As interface{}

    // AS index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Address index. The type is interface{} with range: 0..32767.
    AddressIndex interface{}
}

func (vplsId *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId) GetEntityData() *types.CommonEntityData {
    vplsId.EntityData.YFilter = vplsId.YFilter
    vplsId.EntityData.YangName = "vpls-id"
    vplsId.EntityData.BundleName = "cisco_ios_xr"
    vplsId.EntityData.ParentYangName = "ldp-signaling-protocol"
    vplsId.EntityData.SegmentPath = "vpls-id"
    vplsId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vplsId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vplsId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vplsId.EntityData.Children = types.NewOrderedMap()
    vplsId.EntityData.Leafs = types.NewOrderedMap()
    vplsId.EntityData.Leafs.Append("type", types.YLeaf{"Type", vplsId.Type})
    vplsId.EntityData.Leafs.Append("as", types.YLeaf{"As", vplsId.As})
    vplsId.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", vplsId.AsIndex})
    vplsId.EntityData.Leafs.Append("address", types.YLeaf{"Address", vplsId.Address})
    vplsId.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", vplsId.AddressIndex})

    vplsId.EntityData.YListKeys = []string {}

    return &(vplsId.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "ldp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs.Append("flow-label", types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel})
    flowLabelLoadBalance.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabelLoadBalance.Static})

    flowLabelLoadBalance.EntityData.YListKeys = []string {}

    return &(flowLabelLoadBalance.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy
// Route policy
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Export route policy. The type is string.
    Export interface{}
}

func (bgpRoutePolicy *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy) GetEntityData() *types.CommonEntityData {
    bgpRoutePolicy.EntityData.YFilter = bgpRoutePolicy.YFilter
    bgpRoutePolicy.EntityData.YangName = "bgp-route-policy"
    bgpRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    bgpRoutePolicy.EntityData.ParentYangName = "bgp-auto-discovery"
    bgpRoutePolicy.EntityData.SegmentPath = "bgp-route-policy"
    bgpRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpRoutePolicy.EntityData.Children = types.NewOrderedMap()
    bgpRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    bgpRoutePolicy.EntityData.Leafs.Append("export", types.YLeaf{"Export", bgpRoutePolicy.Export})

    bgpRoutePolicy.EntityData.YListKeys = []string {}

    return &(bgpRoutePolicy.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher
// Route Distinguisher
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (routeDistinguisher *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher) GetEntityData() *types.CommonEntityData {
    routeDistinguisher.EntityData.YFilter = routeDistinguisher.YFilter
    routeDistinguisher.EntityData.YangName = "route-distinguisher"
    routeDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    routeDistinguisher.EntityData.ParentYangName = "bgp-auto-discovery"
    routeDistinguisher.EntityData.SegmentPath = "route-distinguisher"
    routeDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeDistinguisher.EntityData.Children = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeDistinguisher.Type})
    routeDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", routeDistinguisher.As})
    routeDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeDistinguisher.AsIndex})
    routeDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeDistinguisher.Address})
    routeDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", routeDistinguisher.AddrIndex})

    routeDistinguisher.EntityData.YListKeys = []string {}

    return &(routeDistinguisher.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol
// Enable Signaling Protocol BGP in this VFI
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Virtual Edge Block Configurable Range. The type is interface{} with
    // range: 11..100.
    VeRange interface{}

    // Local Virtual Edge Identifier. The type is interface{} with range:
    // 1..16384.
    Veid interface{}

    // Enable BGP as Signaling Protocol. The type is interface{}.
    Enable interface{}

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance
}

func (bgpSignalingProtocol *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    bgpSignalingProtocol.EntityData.YFilter = bgpSignalingProtocol.YFilter
    bgpSignalingProtocol.EntityData.YangName = "bgp-signaling-protocol"
    bgpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    bgpSignalingProtocol.EntityData.ParentYangName = "bgp-auto-discovery"
    bgpSignalingProtocol.EntityData.SegmentPath = "bgp-signaling-protocol"
    bgpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpSignalingProtocol.EntityData.Children = types.NewOrderedMap()
    bgpSignalingProtocol.EntityData.Children.Append("flow-label-load-balance", types.YChild{"FlowLabelLoadBalance", &bgpSignalingProtocol.FlowLabelLoadBalance})
    bgpSignalingProtocol.EntityData.Leafs = types.NewOrderedMap()
    bgpSignalingProtocol.EntityData.Leafs.Append("ve-range", types.YLeaf{"VeRange", bgpSignalingProtocol.VeRange})
    bgpSignalingProtocol.EntityData.Leafs.Append("veid", types.YLeaf{"Veid", bgpSignalingProtocol.Veid})
    bgpSignalingProtocol.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bgpSignalingProtocol.Enable})

    bgpSignalingProtocol.EntityData.YListKeys = []string {}

    return &(bgpSignalingProtocol.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "bgp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs.Append("flow-label", types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel})
    flowLabelLoadBalance.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabelLoadBalance.Static})

    flowLabelLoadBalance.EntityData.YListKeys = []string {}

    return &(flowLabelLoadBalance.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets
// Route Target
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget.
    RouteTarget []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget
}

func (routeTargets *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "bgp-auto-discovery"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget
// Name of the Route Target
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // two byte as or four byte as. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs.
    TwoByteAsOrFourByteAs []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs

    // ipv4 address. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddKeyToken(routeTarget.Role, "role") + types.AddKeyToken(routeTarget.Format, "format")
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("two-byte-as-or-four-byte-as", types.YChild{"TwoByteAsOrFourByteAs", nil})
    for i := range routeTarget.TwoByteAsOrFourByteAs {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.TwoByteAsOrFourByteAs[i]), types.YChild{"TwoByteAsOrFourByteAs", routeTarget.TwoByteAsOrFourByteAs[i]})
    }
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", routeTarget.Ipv4Address[i]})
    }
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("role", types.YLeaf{"Role", routeTarget.Role})
    routeTarget.EntityData.Leafs.Append("format", types.YLeaf{"Format", routeTarget.Format})

    routeTarget.EntityData.YListKeys = []string {"Role", "Format"}

    return &(routeTarget.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs
// two byte as or four byte as
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}
}

func (twoByteAsOrFourByteAs *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAsOrFourByteAs.EntityData.YFilter = twoByteAsOrFourByteAs.YFilter
    twoByteAsOrFourByteAs.EntityData.YangName = "two-byte-as-or-four-byte-as"
    twoByteAsOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAsOrFourByteAs.EntityData.ParentYangName = "route-target"
    twoByteAsOrFourByteAs.EntityData.SegmentPath = "two-byte-as-or-four-byte-as" + types.AddKeyToken(twoByteAsOrFourByteAs.As, "as") + types.AddKeyToken(twoByteAsOrFourByteAs.AsIndex, "as-index")
    twoByteAsOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAsOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAsOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAsOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAsOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAsOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", twoByteAsOrFourByteAs.As})
    twoByteAsOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", twoByteAsOrFourByteAs.AsIndex})

    twoByteAsOrFourByteAs.EntityData.YListKeys = []string {"As", "AsIndex"}

    return &(twoByteAsOrFourByteAs.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}
}

func (ipv4Address *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddrIndex, "addr-index")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", ipv4Address.AddrIndex})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddrIndex"}

    return &(ipv4Address.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis
// Bridge Domain EVPN VxLAN Network Identifier
// Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain VxLAN EVPN. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni.
    BridgeDomainvni []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni
}

func (bridgeDomainvnis *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis) GetEntityData() *types.CommonEntityData {
    bridgeDomainvnis.EntityData.YFilter = bridgeDomainvnis.YFilter
    bridgeDomainvnis.EntityData.YangName = "bridge-domainvnis"
    bridgeDomainvnis.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainvnis.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainvnis.EntityData.SegmentPath = "bridge-domainvnis"
    bridgeDomainvnis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainvnis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainvnis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainvnis.EntityData.Children = types.NewOrderedMap()
    bridgeDomainvnis.EntityData.Children.Append("bridge-domainvni", types.YChild{"BridgeDomainvni", nil})
    for i := range bridgeDomainvnis.BridgeDomainvni {
        bridgeDomainvnis.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainvnis.BridgeDomainvni[i]), types.YChild{"BridgeDomainvni", bridgeDomainvnis.BridgeDomainvni[i]})
    }
    bridgeDomainvnis.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainvnis.EntityData.YListKeys = []string {}

    return &(bridgeDomainvnis.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni
// Bridge Domain VxLAN EVPN
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VxLAN Ethernet VPN-ID. The type is interface{}
    // with range: 1..16777215.
    VpnId interface{}
}

func (bridgeDomainvni *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni) GetEntityData() *types.CommonEntityData {
    bridgeDomainvni.EntityData.YFilter = bridgeDomainvni.YFilter
    bridgeDomainvni.EntityData.YangName = "bridge-domainvni"
    bridgeDomainvni.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainvni.EntityData.ParentYangName = "bridge-domainvnis"
    bridgeDomainvni.EntityData.SegmentPath = "bridge-domainvni" + types.AddKeyToken(bridgeDomainvni.VpnId, "vpn-id")
    bridgeDomainvni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainvni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainvni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainvni.EntityData.Children = types.NewOrderedMap()
    bridgeDomainvni.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainvni.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", bridgeDomainvni.VpnId})

    bridgeDomainvni.EntityData.YListKeys = []string {"VpnId"}

    return &(bridgeDomainvni.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits
// Attachment Circuit table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Attachment Circuit. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit.
    BdAttachmentCircuit []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit
}

func (bdAttachmentCircuits *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    bdAttachmentCircuits.EntityData.YFilter = bdAttachmentCircuits.YFilter
    bdAttachmentCircuits.EntityData.YangName = "bd-attachment-circuits"
    bdAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    bdAttachmentCircuits.EntityData.ParentYangName = "bridge-domain"
    bdAttachmentCircuits.EntityData.SegmentPath = "bd-attachment-circuits"
    bdAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdAttachmentCircuits.EntityData.Children = types.NewOrderedMap()
    bdAttachmentCircuits.EntityData.Children.Append("bd-attachment-circuit", types.YChild{"BdAttachmentCircuit", nil})
    for i := range bdAttachmentCircuits.BdAttachmentCircuit {
        bdAttachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(bdAttachmentCircuits.BdAttachmentCircuit[i]), types.YChild{"BdAttachmentCircuit", bdAttachmentCircuits.BdAttachmentCircuit[i]})
    }
    bdAttachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    bdAttachmentCircuits.EntityData.YListKeys = []string {}

    return &(bdAttachmentCircuits.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit
// Name of the Attachment Circuit
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Attachment Circuit. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    Name interface{}

    // Enable or Disable Flooding. The type is InterfaceTrafficFlood.
    InterfaceFlooding interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    InterfaceIgmpSnoop interface{}

    // Enable or Disable Unknown Unicast Flooding. The type is
    // InterfaceTrafficFlood.
    InterfaceFloodingUnknownUnicast interface{}

    // Attach a MLD Snooping profile. The type is string with length: 1..32.
    InterfaceMldSnoop interface{}

    // IP Source Guard.
    InterfaceIpSourceGuard L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard

    // L2 Interface Dynamic ARP Inspection.
    InterfaceDai L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai

    // Attach a DHCP profile.
    InterfaceProfile L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile

    // Storm Control.
    BdacStormControlTypes L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes

    // Split Horizon.
    SplitHorizon L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon

    // Static Mac Address Table.
    StaticMacAddresses L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses

    // MAC configuration commands.
    InterfaceMac L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac
}

func (bdAttachmentCircuit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    bdAttachmentCircuit.EntityData.YFilter = bdAttachmentCircuit.YFilter
    bdAttachmentCircuit.EntityData.YangName = "bd-attachment-circuit"
    bdAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    bdAttachmentCircuit.EntityData.ParentYangName = "bd-attachment-circuits"
    bdAttachmentCircuit.EntityData.SegmentPath = "bd-attachment-circuit" + types.AddKeyToken(bdAttachmentCircuit.Name, "name")
    bdAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdAttachmentCircuit.EntityData.Children = types.NewOrderedMap()
    bdAttachmentCircuit.EntityData.Children.Append("interface-ip-source-guard", types.YChild{"InterfaceIpSourceGuard", &bdAttachmentCircuit.InterfaceIpSourceGuard})
    bdAttachmentCircuit.EntityData.Children.Append("interface-dai", types.YChild{"InterfaceDai", &bdAttachmentCircuit.InterfaceDai})
    bdAttachmentCircuit.EntityData.Children.Append("interface-profile", types.YChild{"InterfaceProfile", &bdAttachmentCircuit.InterfaceProfile})
    bdAttachmentCircuit.EntityData.Children.Append("bdac-storm-control-types", types.YChild{"BdacStormControlTypes", &bdAttachmentCircuit.BdacStormControlTypes})
    bdAttachmentCircuit.EntityData.Children.Append("split-horizon", types.YChild{"SplitHorizon", &bdAttachmentCircuit.SplitHorizon})
    bdAttachmentCircuit.EntityData.Children.Append("static-mac-addresses", types.YChild{"StaticMacAddresses", &bdAttachmentCircuit.StaticMacAddresses})
    bdAttachmentCircuit.EntityData.Children.Append("interface-mac", types.YChild{"InterfaceMac", &bdAttachmentCircuit.InterfaceMac})
    bdAttachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    bdAttachmentCircuit.EntityData.Leafs.Append("name", types.YLeaf{"Name", bdAttachmentCircuit.Name})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-flooding", types.YLeaf{"InterfaceFlooding", bdAttachmentCircuit.InterfaceFlooding})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-igmp-snoop", types.YLeaf{"InterfaceIgmpSnoop", bdAttachmentCircuit.InterfaceIgmpSnoop})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-flooding-unknown-unicast", types.YLeaf{"InterfaceFloodingUnknownUnicast", bdAttachmentCircuit.InterfaceFloodingUnknownUnicast})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-mld-snoop", types.YLeaf{"InterfaceMldSnoop", bdAttachmentCircuit.InterfaceMldSnoop})

    bdAttachmentCircuit.EntityData.YListKeys = []string {"Name"}

    return &(bdAttachmentCircuit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard
// IP Source Guard
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface Dynamic IP source guard. The type is interface{}.
    Disable interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (interfaceIpSourceGuard *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard) GetEntityData() *types.CommonEntityData {
    interfaceIpSourceGuard.EntityData.YFilter = interfaceIpSourceGuard.YFilter
    interfaceIpSourceGuard.EntityData.YangName = "interface-ip-source-guard"
    interfaceIpSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    interfaceIpSourceGuard.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceIpSourceGuard.EntityData.SegmentPath = "interface-ip-source-guard"
    interfaceIpSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIpSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIpSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIpSourceGuard.EntityData.Children = types.NewOrderedMap()
    interfaceIpSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    interfaceIpSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", interfaceIpSourceGuard.Logging})
    interfaceIpSourceGuard.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", interfaceIpSourceGuard.Disable})
    interfaceIpSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceIpSourceGuard.Enable})

    interfaceIpSourceGuard.EntityData.YListKeys = []string {}

    return &(interfaceIpSourceGuard.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai
// L2 Interface Dynamic ARP Inspection
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface Dynamic ARP Inspection. The type is interface{}.
    Disable interface{}

    // Enable L2 Interface Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    InterfaceDaiAddressValidation L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation
}

func (interfaceDai *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai) GetEntityData() *types.CommonEntityData {
    interfaceDai.EntityData.YFilter = interfaceDai.YFilter
    interfaceDai.EntityData.YangName = "interface-dai"
    interfaceDai.EntityData.BundleName = "cisco_ios_xr"
    interfaceDai.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceDai.EntityData.SegmentPath = "interface-dai"
    interfaceDai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDai.EntityData.Children = types.NewOrderedMap()
    interfaceDai.EntityData.Children.Append("interface-dai-address-validation", types.YChild{"InterfaceDaiAddressValidation", &interfaceDai.InterfaceDaiAddressValidation})
    interfaceDai.EntityData.Leafs = types.NewOrderedMap()
    interfaceDai.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", interfaceDai.Logging})
    interfaceDai.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", interfaceDai.Disable})
    interfaceDai.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceDai.Enable})

    interfaceDai.EntityData.YListKeys = []string {}

    return &(interfaceDai.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation
// Address Validation
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Verification. The type is L2vpnVerification.
    Ipv4Verification interface{}

    // Destination MAC Verification. The type is L2vpnVerification.
    DestinationMacVerification interface{}

    // Source MAC Verification. The type is L2vpnVerification.
    SourceMacVerification interface{}

    // Enable Address Validation. The type is interface{}.
    Enable interface{}
}

func (interfaceDaiAddressValidation *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation) GetEntityData() *types.CommonEntityData {
    interfaceDaiAddressValidation.EntityData.YFilter = interfaceDaiAddressValidation.YFilter
    interfaceDaiAddressValidation.EntityData.YangName = "interface-dai-address-validation"
    interfaceDaiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    interfaceDaiAddressValidation.EntityData.ParentYangName = "interface-dai"
    interfaceDaiAddressValidation.EntityData.SegmentPath = "interface-dai-address-validation"
    interfaceDaiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDaiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDaiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDaiAddressValidation.EntityData.Children = types.NewOrderedMap()
    interfaceDaiAddressValidation.EntityData.Leafs = types.NewOrderedMap()
    interfaceDaiAddressValidation.EntityData.Leafs.Append("ipv4-verification", types.YLeaf{"Ipv4Verification", interfaceDaiAddressValidation.Ipv4Verification})
    interfaceDaiAddressValidation.EntityData.Leafs.Append("destination-mac-verification", types.YLeaf{"DestinationMacVerification", interfaceDaiAddressValidation.DestinationMacVerification})
    interfaceDaiAddressValidation.EntityData.Leafs.Append("source-mac-verification", types.YLeaf{"SourceMacVerification", interfaceDaiAddressValidation.SourceMacVerification})
    interfaceDaiAddressValidation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceDaiAddressValidation.Enable})

    interfaceDaiAddressValidation.EntityData.YListKeys = []string {}

    return &(interfaceDaiAddressValidation.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile
// Attach a DHCP profile
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (interfaceProfile *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile) GetEntityData() *types.CommonEntityData {
    interfaceProfile.EntityData.YFilter = interfaceProfile.YFilter
    interfaceProfile.EntityData.YangName = "interface-profile"
    interfaceProfile.EntityData.BundleName = "cisco_ios_xr"
    interfaceProfile.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceProfile.EntityData.SegmentPath = "interface-profile"
    interfaceProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProfile.EntityData.Children = types.NewOrderedMap()
    interfaceProfile.EntityData.Leafs = types.NewOrderedMap()
    interfaceProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", interfaceProfile.ProfileId})
    interfaceProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", interfaceProfile.DhcpSnoopingId})

    interfaceProfile.EntityData.YListKeys = []string {}

    return &(interfaceProfile.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes
// Storm Control
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType.
    BdacStormControlType []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType
}

func (bdacStormControlTypes *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes) GetEntityData() *types.CommonEntityData {
    bdacStormControlTypes.EntityData.YFilter = bdacStormControlTypes.YFilter
    bdacStormControlTypes.EntityData.YangName = "bdac-storm-control-types"
    bdacStormControlTypes.EntityData.BundleName = "cisco_ios_xr"
    bdacStormControlTypes.EntityData.ParentYangName = "bd-attachment-circuit"
    bdacStormControlTypes.EntityData.SegmentPath = "bdac-storm-control-types"
    bdacStormControlTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdacStormControlTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdacStormControlTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdacStormControlTypes.EntityData.Children = types.NewOrderedMap()
    bdacStormControlTypes.EntityData.Children.Append("bdac-storm-control-type", types.YChild{"BdacStormControlType", nil})
    for i := range bdacStormControlTypes.BdacStormControlType {
        bdacStormControlTypes.EntityData.Children.Append(types.GetSegmentPath(bdacStormControlTypes.BdacStormControlType[i]), types.YChild{"BdacStormControlType", bdacStormControlTypes.BdacStormControlType[i]})
    }
    bdacStormControlTypes.EntityData.Leafs = types.NewOrderedMap()

    bdacStormControlTypes.EntityData.YListKeys = []string {}

    return &(bdacStormControlTypes.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType
// Storm Control Type
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit
}

func (bdacStormControlType *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType) GetEntityData() *types.CommonEntityData {
    bdacStormControlType.EntityData.YFilter = bdacStormControlType.YFilter
    bdacStormControlType.EntityData.YangName = "bdac-storm-control-type"
    bdacStormControlType.EntityData.BundleName = "cisco_ios_xr"
    bdacStormControlType.EntityData.ParentYangName = "bdac-storm-control-types"
    bdacStormControlType.EntityData.SegmentPath = "bdac-storm-control-type" + types.AddKeyToken(bdacStormControlType.Sctype, "sctype")
    bdacStormControlType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdacStormControlType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdacStormControlType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdacStormControlType.EntityData.Children = types.NewOrderedMap()
    bdacStormControlType.EntityData.Children.Append("storm-control-unit", types.YChild{"StormControlUnit", &bdacStormControlType.StormControlUnit})
    bdacStormControlType.EntityData.Leafs = types.NewOrderedMap()
    bdacStormControlType.EntityData.Leafs.Append("sctype", types.YLeaf{"Sctype", bdacStormControlType.Sctype})

    bdacStormControlType.EntityData.YListKeys = []string {"Sctype"}

    return &(bdacStormControlType.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit
// Specify units for Storm Control Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Kilobits Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 64..1280000. Units are
    // kbit/s.
    KbitsPerSec interface{}

    // Packets Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 1..160000. Units are
    // packet/s.
    PktsPerSec interface{}
}

func (stormControlUnit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bdac-storm-control-type"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs.Append("kbits-per-sec", types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec})
    stormControlUnit.EntityData.Leafs.Append("pkts-per-sec", types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec})

    stormControlUnit.EntityData.YListKeys = []string {}

    return &(stormControlUnit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon
// Split Horizon
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon Group ID.
    SplitHorizonGroupId L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId
}

func (splitHorizon *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon) GetEntityData() *types.CommonEntityData {
    splitHorizon.EntityData.YFilter = splitHorizon.YFilter
    splitHorizon.EntityData.YangName = "split-horizon"
    splitHorizon.EntityData.BundleName = "cisco_ios_xr"
    splitHorizon.EntityData.ParentYangName = "bd-attachment-circuit"
    splitHorizon.EntityData.SegmentPath = "split-horizon"
    splitHorizon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    splitHorizon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    splitHorizon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    splitHorizon.EntityData.Children = types.NewOrderedMap()
    splitHorizon.EntityData.Children.Append("split-horizon-group-id", types.YChild{"SplitHorizonGroupId", &splitHorizon.SplitHorizonGroupId})
    splitHorizon.EntityData.Leafs = types.NewOrderedMap()

    splitHorizon.EntityData.YListKeys = []string {}

    return &(splitHorizon.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId
// Split Horizon Group ID
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable split horizon group. The type is interface{}.
    Enable interface{}
}

func (splitHorizonGroupId *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId) GetEntityData() *types.CommonEntityData {
    splitHorizonGroupId.EntityData.YFilter = splitHorizonGroupId.YFilter
    splitHorizonGroupId.EntityData.YangName = "split-horizon-group-id"
    splitHorizonGroupId.EntityData.BundleName = "cisco_ios_xr"
    splitHorizonGroupId.EntityData.ParentYangName = "split-horizon"
    splitHorizonGroupId.EntityData.SegmentPath = "split-horizon-group-id"
    splitHorizonGroupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    splitHorizonGroupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    splitHorizonGroupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    splitHorizonGroupId.EntityData.Children = types.NewOrderedMap()
    splitHorizonGroupId.EntityData.Leafs = types.NewOrderedMap()
    splitHorizonGroupId.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", splitHorizonGroupId.Enable})

    splitHorizonGroupId.EntityData.YListKeys = []string {}

    return &(splitHorizonGroupId.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress.
    StaticMacAddress []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress
}

func (staticMacAddresses *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses) GetEntityData() *types.CommonEntityData {
    staticMacAddresses.EntityData.YFilter = staticMacAddresses.YFilter
    staticMacAddresses.EntityData.YangName = "static-mac-addresses"
    staticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticMacAddresses.EntityData.ParentYangName = "bd-attachment-circuit"
    staticMacAddresses.EntityData.SegmentPath = "static-mac-addresses"
    staticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMacAddresses.EntityData.Children = types.NewOrderedMap()
    staticMacAddresses.EntityData.Children.Append("static-mac-address", types.YChild{"StaticMacAddress", nil})
    for i := range staticMacAddresses.StaticMacAddress {
        staticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(staticMacAddresses.StaticMacAddress[i]), types.YChild{"StaticMacAddress", staticMacAddresses.StaticMacAddress[i]})
    }
    staticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    staticMacAddresses.EntityData.YListKeys = []string {}

    return &(staticMacAddresses.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (staticMacAddress *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress) GetEntityData() *types.CommonEntityData {
    staticMacAddress.EntityData.YFilter = staticMacAddress.YFilter
    staticMacAddress.EntityData.YangName = "static-mac-address"
    staticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    staticMacAddress.EntityData.ParentYangName = "static-mac-addresses"
    staticMacAddress.EntityData.SegmentPath = "static-mac-address" + types.AddKeyToken(staticMacAddress.Address, "address")
    staticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMacAddress.EntityData.Children = types.NewOrderedMap()
    staticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    staticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", staticMacAddress.Address})

    staticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(staticMacAddress.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac
// MAC configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MAC Flush When Port goes down. The type is PortDownFlush.
    InterfaceMacPortDownFlush interface{}

    // Enable Mac Learning. The type is MacLearn.
    InterfaceMacLearning interface{}

    // MAC-Aging configuration commands.
    InterfaceMacAging L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging

    // MAC Secure.
    InterfaceMacSecure L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure

    // MAC-Limit configuration commands.
    InterfaceMacLimit L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit
}

func (interfaceMac *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac) GetEntityData() *types.CommonEntityData {
    interfaceMac.EntityData.YFilter = interfaceMac.YFilter
    interfaceMac.EntityData.YangName = "interface-mac"
    interfaceMac.EntityData.BundleName = "cisco_ios_xr"
    interfaceMac.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceMac.EntityData.SegmentPath = "interface-mac"
    interfaceMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMac.EntityData.Children = types.NewOrderedMap()
    interfaceMac.EntityData.Children.Append("interface-mac-aging", types.YChild{"InterfaceMacAging", &interfaceMac.InterfaceMacAging})
    interfaceMac.EntityData.Children.Append("interface-mac-secure", types.YChild{"InterfaceMacSecure", &interfaceMac.InterfaceMacSecure})
    interfaceMac.EntityData.Children.Append("interface-mac-limit", types.YChild{"InterfaceMacLimit", &interfaceMac.InterfaceMacLimit})
    interfaceMac.EntityData.Leafs = types.NewOrderedMap()
    interfaceMac.EntityData.Leafs.Append("interface-mac-port-down-flush", types.YLeaf{"InterfaceMacPortDownFlush", interfaceMac.InterfaceMacPortDownFlush})
    interfaceMac.EntityData.Leafs.Append("interface-mac-learning", types.YLeaf{"InterfaceMacLearning", interfaceMac.InterfaceMacLearning})

    interfaceMac.EntityData.YListKeys = []string {}

    return &(interfaceMac.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging
// MAC-Aging configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    InterfaceMacAgingTime interface{}

    // MAC address aging type. The type is MacAging.
    InterfaceMacAgingType interface{}
}

func (interfaceMacAging *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging) GetEntityData() *types.CommonEntityData {
    interfaceMacAging.EntityData.YFilter = interfaceMacAging.YFilter
    interfaceMacAging.EntityData.YangName = "interface-mac-aging"
    interfaceMacAging.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacAging.EntityData.ParentYangName = "interface-mac"
    interfaceMacAging.EntityData.SegmentPath = "interface-mac-aging"
    interfaceMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacAging.EntityData.Children = types.NewOrderedMap()
    interfaceMacAging.EntityData.Leafs = types.NewOrderedMap()
    interfaceMacAging.EntityData.Leafs.Append("interface-mac-aging-time", types.YLeaf{"InterfaceMacAgingTime", interfaceMacAging.InterfaceMacAgingTime})
    interfaceMacAging.EntityData.Leafs.Append("interface-mac-aging-type", types.YLeaf{"InterfaceMacAgingType", interfaceMacAging.InterfaceMacAgingType})

    interfaceMacAging.EntityData.YListKeys = []string {}

    return &(interfaceMacAging.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure
// MAC Secure
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface MAC Secure. The type is interface{}.
    Disable interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}
}

func (interfaceMacSecure *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure) GetEntityData() *types.CommonEntityData {
    interfaceMacSecure.EntityData.YFilter = interfaceMacSecure.YFilter
    interfaceMacSecure.EntityData.YangName = "interface-mac-secure"
    interfaceMacSecure.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacSecure.EntityData.ParentYangName = "interface-mac"
    interfaceMacSecure.EntityData.SegmentPath = "interface-mac-secure"
    interfaceMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacSecure.EntityData.Children = types.NewOrderedMap()
    interfaceMacSecure.EntityData.Leafs = types.NewOrderedMap()
    interfaceMacSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", interfaceMacSecure.Logging})
    interfaceMacSecure.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", interfaceMacSecure.Disable})
    interfaceMacSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", interfaceMacSecure.Action})
    interfaceMacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceMacSecure.Enable})

    interfaceMacSecure.EntityData.YListKeys = []string {}

    return &(interfaceMacSecure.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of MAC addresses on an Interface after which MAC limit action is
    // taken. The type is interface{} with range: 0..4294967295.
    InterfaceMacLimitMax interface{}

    // MAC address limit notification action in a Interface. The type is
    // MacNotification.
    InterfaceMacLimitNotif interface{}

    // Interface MAC address limit enforcement action. The type is MacLimitAction.
    InterfaceMacLimitAction interface{}
}

func (interfaceMacLimit *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit) GetEntityData() *types.CommonEntityData {
    interfaceMacLimit.EntityData.YFilter = interfaceMacLimit.YFilter
    interfaceMacLimit.EntityData.YangName = "interface-mac-limit"
    interfaceMacLimit.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacLimit.EntityData.ParentYangName = "interface-mac"
    interfaceMacLimit.EntityData.SegmentPath = "interface-mac-limit"
    interfaceMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacLimit.EntityData.Children = types.NewOrderedMap()
    interfaceMacLimit.EntityData.Leafs = types.NewOrderedMap()
    interfaceMacLimit.EntityData.Leafs.Append("interface-mac-limit-max", types.YLeaf{"InterfaceMacLimitMax", interfaceMacLimit.InterfaceMacLimitMax})
    interfaceMacLimit.EntityData.Leafs.Append("interface-mac-limit-notif", types.YLeaf{"InterfaceMacLimitNotif", interfaceMacLimit.InterfaceMacLimitNotif})
    interfaceMacLimit.EntityData.Leafs.Append("interface-mac-limit-action", types.YLeaf{"InterfaceMacLimitAction", interfaceMacLimit.InterfaceMacLimitAction})

    interfaceMacLimit.EntityData.YListKeys = []string {}

    return &(interfaceMacLimit.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns
// List of EVPN pseudowires
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Pseudowire configuration. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn.
    BdPseudowireEvpn []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn
}

func (bdPseudowireEvpns *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns) GetEntityData() *types.CommonEntityData {
    bdPseudowireEvpns.EntityData.YFilter = bdPseudowireEvpns.YFilter
    bdPseudowireEvpns.EntityData.YangName = "bd-pseudowire-evpns"
    bdPseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowireEvpns.EntityData.ParentYangName = "bridge-domain"
    bdPseudowireEvpns.EntityData.SegmentPath = "bd-pseudowire-evpns"
    bdPseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowireEvpns.EntityData.Children = types.NewOrderedMap()
    bdPseudowireEvpns.EntityData.Children.Append("bd-pseudowire-evpn", types.YChild{"BdPseudowireEvpn", nil})
    for i := range bdPseudowireEvpns.BdPseudowireEvpn {
        bdPseudowireEvpns.EntityData.Children.Append(types.GetSegmentPath(bdPseudowireEvpns.BdPseudowireEvpn[i]), types.YChild{"BdPseudowireEvpn", bdPseudowireEvpns.BdPseudowireEvpn[i]})
    }
    bdPseudowireEvpns.EntityData.Leafs = types.NewOrderedMap()

    bdPseudowireEvpns.EntityData.YListKeys = []string {}

    return &(bdPseudowireEvpns.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn
// EVPN Pseudowire configuration
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // This attribute is a key. AC ID. The type is interface{} with range:
    // 1..4294967295.
    Acid interface{}
}

func (bdPseudowireEvpn *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn) GetEntityData() *types.CommonEntityData {
    bdPseudowireEvpn.EntityData.YFilter = bdPseudowireEvpn.YFilter
    bdPseudowireEvpn.EntityData.YangName = "bd-pseudowire-evpn"
    bdPseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowireEvpn.EntityData.ParentYangName = "bd-pseudowire-evpns"
    bdPseudowireEvpn.EntityData.SegmentPath = "bd-pseudowire-evpn" + types.AddKeyToken(bdPseudowireEvpn.Eviid, "eviid") + types.AddKeyToken(bdPseudowireEvpn.Acid, "acid")
    bdPseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowireEvpn.EntityData.Children = types.NewOrderedMap()
    bdPseudowireEvpn.EntityData.Leafs = types.NewOrderedMap()
    bdPseudowireEvpn.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", bdPseudowireEvpn.Eviid})
    bdPseudowireEvpn.EntityData.Leafs.Append("acid", types.YLeaf{"Acid", bdPseudowireEvpn.Acid})

    bdPseudowireEvpn.EntityData.YListKeys = []string {"Eviid", "Acid"}

    return &(bdPseudowireEvpn.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard
// IP Source Guard
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Logging. The type is interface{}.
    Logging interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (ipSourceGuard *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard) GetEntityData() *types.CommonEntityData {
    ipSourceGuard.EntityData.YFilter = ipSourceGuard.YFilter
    ipSourceGuard.EntityData.YangName = "ip-source-guard"
    ipSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    ipSourceGuard.EntityData.ParentYangName = "bridge-domain"
    ipSourceGuard.EntityData.SegmentPath = "ip-source-guard"
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSourceGuard.EntityData.Children = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", ipSourceGuard.Logging})
    ipSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipSourceGuard.Enable})

    ipSourceGuard.EntityData.YListKeys = []string {}

    return &(ipSourceGuard.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai
// Dynamic ARP Inspection
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Logging. The type is interface{}.
    Logging interface{}

    // Enable Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    DaiAddressValidation L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation
}

func (dai *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai) GetEntityData() *types.CommonEntityData {
    dai.EntityData.YFilter = dai.YFilter
    dai.EntityData.YangName = "dai"
    dai.EntityData.BundleName = "cisco_ios_xr"
    dai.EntityData.ParentYangName = "bridge-domain"
    dai.EntityData.SegmentPath = "dai"
    dai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dai.EntityData.Children = types.NewOrderedMap()
    dai.EntityData.Children.Append("dai-address-validation", types.YChild{"DaiAddressValidation", &dai.DaiAddressValidation})
    dai.EntityData.Leafs = types.NewOrderedMap()
    dai.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", dai.Logging})
    dai.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", dai.Enable})

    dai.EntityData.YListKeys = []string {}

    return &(dai.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation
// Address Validation
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPv4 Verification. The type is interface{}.
    Ipv4Verification interface{}

    // Enable Destination MAC Verification. The type is interface{}.
    DestinationMacVerification interface{}

    // Enable Source MAC Verification. The type is interface{}.
    SourceMacVerification interface{}

    // Enable Address Validation. The type is interface{}.
    Enable interface{}
}

func (daiAddressValidation *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation) GetEntityData() *types.CommonEntityData {
    daiAddressValidation.EntityData.YFilter = daiAddressValidation.YFilter
    daiAddressValidation.EntityData.YangName = "dai-address-validation"
    daiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    daiAddressValidation.EntityData.ParentYangName = "dai"
    daiAddressValidation.EntityData.SegmentPath = "dai-address-validation"
    daiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    daiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    daiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    daiAddressValidation.EntityData.Children = types.NewOrderedMap()
    daiAddressValidation.EntityData.Leafs = types.NewOrderedMap()
    daiAddressValidation.EntityData.Leafs.Append("ipv4-verification", types.YLeaf{"Ipv4Verification", daiAddressValidation.Ipv4Verification})
    daiAddressValidation.EntityData.Leafs.Append("destination-mac-verification", types.YLeaf{"DestinationMacVerification", daiAddressValidation.DestinationMacVerification})
    daiAddressValidation.EntityData.Leafs.Append("source-mac-verification", types.YLeaf{"SourceMacVerification", daiAddressValidation.SourceMacVerification})
    daiAddressValidation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", daiAddressValidation.Enable})

    daiAddressValidation.EntityData.YListKeys = []string {}

    return &(daiAddressValidation.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces
// Bridge Domain Routed Interface Table
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain Routed Interface. The type is slice of
    // L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface.
    RoutedInterface []*L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface
}

func (routedInterfaces *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces) GetEntityData() *types.CommonEntityData {
    routedInterfaces.EntityData.YFilter = routedInterfaces.YFilter
    routedInterfaces.EntityData.YangName = "routed-interfaces"
    routedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaces.EntityData.ParentYangName = "bridge-domain"
    routedInterfaces.EntityData.SegmentPath = "routed-interfaces"
    routedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaces.EntityData.Children = types.NewOrderedMap()
    routedInterfaces.EntityData.Children.Append("routed-interface", types.YChild{"RoutedInterface", nil})
    for i := range routedInterfaces.RoutedInterface {
        routedInterfaces.EntityData.Children.Append(types.GetSegmentPath(routedInterfaces.RoutedInterface[i]), types.YChild{"RoutedInterface", routedInterfaces.RoutedInterface[i]})
    }
    routedInterfaces.EntityData.Leafs = types.NewOrderedMap()

    routedInterfaces.EntityData.YListKeys = []string {}

    return &(routedInterfaces.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface
// Bridge Domain Routed Interface
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Routed Interface. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Routed interface split horizon group.
    RoutedInterfaceSplitHorizonGroup L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup
}

func (routedInterface *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface) GetEntityData() *types.CommonEntityData {
    routedInterface.EntityData.YFilter = routedInterface.YFilter
    routedInterface.EntityData.YangName = "routed-interface"
    routedInterface.EntityData.BundleName = "cisco_ios_xr"
    routedInterface.EntityData.ParentYangName = "routed-interfaces"
    routedInterface.EntityData.SegmentPath = "routed-interface" + types.AddKeyToken(routedInterface.InterfaceName, "interface-name")
    routedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterface.EntityData.Children = types.NewOrderedMap()
    routedInterface.EntityData.Children.Append("routed-interface-split-horizon-group", types.YChild{"RoutedInterfaceSplitHorizonGroup", &routedInterface.RoutedInterfaceSplitHorizonGroup})
    routedInterface.EntityData.Leafs = types.NewOrderedMap()
    routedInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", routedInterface.InterfaceName})

    routedInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(routedInterface.EntityData)
}

// L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup
// Routed interface split horizon group
type L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure BVI under SHG 1. The type is interface{}.
    RoutedInterfaceSplitHorizonGroupCore interface{}
}

func (routedInterfaceSplitHorizonGroup *L2vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    routedInterfaceSplitHorizonGroup.EntityData.YFilter = routedInterfaceSplitHorizonGroup.YFilter
    routedInterfaceSplitHorizonGroup.EntityData.YangName = "routed-interface-split-horizon-group"
    routedInterfaceSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaceSplitHorizonGroup.EntityData.ParentYangName = "routed-interface"
    routedInterfaceSplitHorizonGroup.EntityData.SegmentPath = "routed-interface-split-horizon-group"
    routedInterfaceSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaceSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaceSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaceSplitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    routedInterfaceSplitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    routedInterfaceSplitHorizonGroup.EntityData.Leafs.Append("routed-interface-split-horizon-group-core", types.YLeaf{"RoutedInterfaceSplitHorizonGroupCore", routedInterfaceSplitHorizonGroup.RoutedInterfaceSplitHorizonGroupCore})

    routedInterfaceSplitHorizonGroup.EntityData.YListKeys = []string {}

    return &(routedInterfaceSplitHorizonGroup.EntityData)
}

// L2vpn_Database_PseudowireClasses
// List of pseudowire classes
type L2vpn_Database_PseudowireClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire class. The type is slice of
    // L2vpn_Database_PseudowireClasses_PseudowireClass.
    PseudowireClass []*L2vpn_Database_PseudowireClasses_PseudowireClass
}

func (pseudowireClasses *L2vpn_Database_PseudowireClasses) GetEntityData() *types.CommonEntityData {
    pseudowireClasses.EntityData.YFilter = pseudowireClasses.YFilter
    pseudowireClasses.EntityData.YangName = "pseudowire-classes"
    pseudowireClasses.EntityData.BundleName = "cisco_ios_xr"
    pseudowireClasses.EntityData.ParentYangName = "database"
    pseudowireClasses.EntityData.SegmentPath = "pseudowire-classes"
    pseudowireClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireClasses.EntityData.Children = types.NewOrderedMap()
    pseudowireClasses.EntityData.Children.Append("pseudowire-class", types.YChild{"PseudowireClass", nil})
    for i := range pseudowireClasses.PseudowireClass {
        pseudowireClasses.EntityData.Children.Append(types.GetSegmentPath(pseudowireClasses.PseudowireClass[i]), types.YChild{"PseudowireClass", pseudowireClasses.PseudowireClass[i]})
    }
    pseudowireClasses.EntityData.Leafs = types.NewOrderedMap()

    pseudowireClasses.EntityData.YListKeys = []string {}

    return &(pseudowireClasses.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass
// Pseudowire class
type L2vpn_Database_PseudowireClasses_PseudowireClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the pseudowire class. The type is string
    // with length: 1..32.
    Name interface{}

    // Enable backup MAC withdraw. The type is interface{}.
    MacWithdraw interface{}

    // Enable pseudowire class. The type is interface{}.
    Enable interface{}

    // L2TPv3 encapsulation.
    L2tpv3Encapsulation L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation

    // Back Up Pseudowire class.
    BackupDisableDelay L2vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay

    // MPLS encapsulation.
    MplsEncapsulation L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation
}

func (pseudowireClass *L2vpn_Database_PseudowireClasses_PseudowireClass) GetEntityData() *types.CommonEntityData {
    pseudowireClass.EntityData.YFilter = pseudowireClass.YFilter
    pseudowireClass.EntityData.YangName = "pseudowire-class"
    pseudowireClass.EntityData.BundleName = "cisco_ios_xr"
    pseudowireClass.EntityData.ParentYangName = "pseudowire-classes"
    pseudowireClass.EntityData.SegmentPath = "pseudowire-class" + types.AddKeyToken(pseudowireClass.Name, "name")
    pseudowireClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireClass.EntityData.Children = types.NewOrderedMap()
    pseudowireClass.EntityData.Children.Append("l2tpv3-encapsulation", types.YChild{"L2tpv3Encapsulation", &pseudowireClass.L2tpv3Encapsulation})
    pseudowireClass.EntityData.Children.Append("backup-disable-delay", types.YChild{"BackupDisableDelay", &pseudowireClass.BackupDisableDelay})
    pseudowireClass.EntityData.Children.Append("mpls-encapsulation", types.YChild{"MplsEncapsulation", &pseudowireClass.MplsEncapsulation})
    pseudowireClass.EntityData.Leafs = types.NewOrderedMap()
    pseudowireClass.EntityData.Leafs.Append("name", types.YLeaf{"Name", pseudowireClass.Name})
    pseudowireClass.EntityData.Leafs.Append("mac-withdraw", types.YLeaf{"MacWithdraw", pseudowireClass.MacWithdraw})
    pseudowireClass.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireClass.Enable})

    pseudowireClass.EntityData.YListKeys = []string {"Name"}

    return &(pseudowireClass.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation
// L2TPv3 encapsulation
type L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the do not fragment bit to 1. The type is interface{}.
    DfBitSet interface{}

    // Cookie size. The type is L2tpCookieSize. The default value is zero.
    CookieSize interface{}

    // Source IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Transport mode. The type is TransportMode.
    TransportMode interface{}

    // Enable L2TPv3 encapsulation. The type is interface{}.
    Enable interface{}

    // Time to live. The type is interface{} with range: 1..255.
    TimeToLive interface{}

    // Sequencing.
    Sequencing L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_Sequencing

    // Type of service.
    TypeOfService L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_TypeOfService

    // L2TPv3 signaling protocol.
    SignalingProtocol L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_SignalingProtocol

    // Path maximum transmission unit.
    PathMtu L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_PathMtu
}

func (l2tpv3Encapsulation *L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation) GetEntityData() *types.CommonEntityData {
    l2tpv3Encapsulation.EntityData.YFilter = l2tpv3Encapsulation.YFilter
    l2tpv3Encapsulation.EntityData.YangName = "l2tpv3-encapsulation"
    l2tpv3Encapsulation.EntityData.BundleName = "cisco_ios_xr"
    l2tpv3Encapsulation.EntityData.ParentYangName = "pseudowire-class"
    l2tpv3Encapsulation.EntityData.SegmentPath = "l2tpv3-encapsulation"
    l2tpv3Encapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tpv3Encapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tpv3Encapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tpv3Encapsulation.EntityData.Children = types.NewOrderedMap()
    l2tpv3Encapsulation.EntityData.Children.Append("sequencing", types.YChild{"Sequencing", &l2tpv3Encapsulation.Sequencing})
    l2tpv3Encapsulation.EntityData.Children.Append("type-of-service", types.YChild{"TypeOfService", &l2tpv3Encapsulation.TypeOfService})
    l2tpv3Encapsulation.EntityData.Children.Append("signaling-protocol", types.YChild{"SignalingProtocol", &l2tpv3Encapsulation.SignalingProtocol})
    l2tpv3Encapsulation.EntityData.Children.Append("path-mtu", types.YChild{"PathMtu", &l2tpv3Encapsulation.PathMtu})
    l2tpv3Encapsulation.EntityData.Leafs = types.NewOrderedMap()
    l2tpv3Encapsulation.EntityData.Leafs.Append("df-bit-set", types.YLeaf{"DfBitSet", l2tpv3Encapsulation.DfBitSet})
    l2tpv3Encapsulation.EntityData.Leafs.Append("cookie-size", types.YLeaf{"CookieSize", l2tpv3Encapsulation.CookieSize})
    l2tpv3Encapsulation.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", l2tpv3Encapsulation.SourceAddress})
    l2tpv3Encapsulation.EntityData.Leafs.Append("transport-mode", types.YLeaf{"TransportMode", l2tpv3Encapsulation.TransportMode})
    l2tpv3Encapsulation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", l2tpv3Encapsulation.Enable})
    l2tpv3Encapsulation.EntityData.Leafs.Append("time-to-live", types.YLeaf{"TimeToLive", l2tpv3Encapsulation.TimeToLive})

    l2tpv3Encapsulation.EntityData.YListKeys = []string {}

    return &(l2tpv3Encapsulation.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_Sequencing
// Sequencing
type L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_Sequencing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequencing. The type is L2tpv3Sequencing. The default value is off.
    Sequencing interface{}

    // Out of sequence threshold. The type is interface{} with range: 5..65535.
    // The default value is 5.
    ResyncThreshold interface{}
}

func (sequencing *L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_Sequencing) GetEntityData() *types.CommonEntityData {
    sequencing.EntityData.YFilter = sequencing.YFilter
    sequencing.EntityData.YangName = "sequencing"
    sequencing.EntityData.BundleName = "cisco_ios_xr"
    sequencing.EntityData.ParentYangName = "l2tpv3-encapsulation"
    sequencing.EntityData.SegmentPath = "sequencing"
    sequencing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sequencing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sequencing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sequencing.EntityData.Children = types.NewOrderedMap()
    sequencing.EntityData.Leafs = types.NewOrderedMap()
    sequencing.EntityData.Leafs.Append("sequencing", types.YLeaf{"Sequencing", sequencing.Sequencing})
    sequencing.EntityData.Leafs.Append("resync-threshold", types.YLeaf{"ResyncThreshold", sequencing.ResyncThreshold})

    sequencing.EntityData.YListKeys = []string {}

    return &(sequencing.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_TypeOfService
// Type of service
type L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_TypeOfService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of service value. The type is interface{} with range: 0..255.
    TypeOfServiceValue interface{}

    // Type of service mode. The type is TypeOfServiceMode.
    TypeOfServiceMode interface{}
}

func (typeOfService *L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_TypeOfService) GetEntityData() *types.CommonEntityData {
    typeOfService.EntityData.YFilter = typeOfService.YFilter
    typeOfService.EntityData.YangName = "type-of-service"
    typeOfService.EntityData.BundleName = "cisco_ios_xr"
    typeOfService.EntityData.ParentYangName = "l2tpv3-encapsulation"
    typeOfService.EntityData.SegmentPath = "type-of-service"
    typeOfService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeOfService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeOfService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeOfService.EntityData.Children = types.NewOrderedMap()
    typeOfService.EntityData.Leafs = types.NewOrderedMap()
    typeOfService.EntityData.Leafs.Append("type-of-service-value", types.YLeaf{"TypeOfServiceValue", typeOfService.TypeOfServiceValue})
    typeOfService.EntityData.Leafs.Append("type-of-service-mode", types.YLeaf{"TypeOfServiceMode", typeOfService.TypeOfServiceMode})

    typeOfService.EntityData.YListKeys = []string {}

    return &(typeOfService.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_SignalingProtocol
// L2TPv3 signaling protocol
type L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_SignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TPv3 signaling protocol. The type is L2tpSignalingProtocol. The default
    // value is l2tpv3.
    Protocol interface{}

    // Name of the L2TPv3 class name. The type is string with length: 1..32.
    L2tpv3ClassName interface{}
}

func (signalingProtocol *L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_SignalingProtocol) GetEntityData() *types.CommonEntityData {
    signalingProtocol.EntityData.YFilter = signalingProtocol.YFilter
    signalingProtocol.EntityData.YangName = "signaling-protocol"
    signalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    signalingProtocol.EntityData.ParentYangName = "l2tpv3-encapsulation"
    signalingProtocol.EntityData.SegmentPath = "signaling-protocol"
    signalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalingProtocol.EntityData.Children = types.NewOrderedMap()
    signalingProtocol.EntityData.Leafs = types.NewOrderedMap()
    signalingProtocol.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", signalingProtocol.Protocol})
    signalingProtocol.EntityData.Leafs.Append("l2tpv3-class-name", types.YLeaf{"L2tpv3ClassName", signalingProtocol.L2tpv3ClassName})

    signalingProtocol.EntityData.YListKeys = []string {}

    return &(signalingProtocol.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_PathMtu
// Path maximum transmission unit
type L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_PathMtu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path MTU. The type is interface{}.
    Enable interface{}

    // Maximum path maximum transmission unit. The type is interface{} with range:
    // 68..65535.
    MaxPathMtu interface{}
}

func (pathMtu *L2vpn_Database_PseudowireClasses_PseudowireClass_L2tpv3Encapsulation_PathMtu) GetEntityData() *types.CommonEntityData {
    pathMtu.EntityData.YFilter = pathMtu.YFilter
    pathMtu.EntityData.YangName = "path-mtu"
    pathMtu.EntityData.BundleName = "cisco_ios_xr"
    pathMtu.EntityData.ParentYangName = "l2tpv3-encapsulation"
    pathMtu.EntityData.SegmentPath = "path-mtu"
    pathMtu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathMtu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathMtu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathMtu.EntityData.Children = types.NewOrderedMap()
    pathMtu.EntityData.Leafs = types.NewOrderedMap()
    pathMtu.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathMtu.Enable})
    pathMtu.EntityData.Leafs.Append("max-path-mtu", types.YLeaf{"MaxPathMtu", pathMtu.MaxPathMtu})

    pathMtu.EntityData.YListKeys = []string {}

    return &(pathMtu.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay
// Back Up Pseudowire class
type L2vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay or Never. The type is BackupDisable.
    Type interface{}

    // Disable backup delay. The type is interface{} with range: 0..180.
    DisableBackup interface{}
}

func (backupDisableDelay *L2vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay) GetEntityData() *types.CommonEntityData {
    backupDisableDelay.EntityData.YFilter = backupDisableDelay.YFilter
    backupDisableDelay.EntityData.YangName = "backup-disable-delay"
    backupDisableDelay.EntityData.BundleName = "cisco_ios_xr"
    backupDisableDelay.EntityData.ParentYangName = "pseudowire-class"
    backupDisableDelay.EntityData.SegmentPath = "backup-disable-delay"
    backupDisableDelay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupDisableDelay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupDisableDelay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupDisableDelay.EntityData.Children = types.NewOrderedMap()
    backupDisableDelay.EntityData.Leafs = types.NewOrderedMap()
    backupDisableDelay.EntityData.Leafs.Append("type", types.YLeaf{"Type", backupDisableDelay.Type})
    backupDisableDelay.EntityData.Leafs.Append("disable-backup", types.YLeaf{"DisableBackup", backupDisableDelay.DisableBackup})

    backupDisableDelay.EntityData.YListKeys = []string {}

    return &(backupDisableDelay.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation
// MPLS encapsulation
type L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire Switching Point Tlv. The type is PwSwitchingPointTlv.
    PwSwitchingTlv interface{}

    // Static Tag rewrite. The type is interface{} with range: 1..4094.
    StaticTagRewrite interface{}

    // MPLS signaling protocol. The type is MplsSignalingProtocol. The default
    // value is ldp.
    SignalingProtocol interface{}

    // VCCV verification type. The type is VccvVerification. The default value is
    // lsp-ping.
    VccvType interface{}

    // Source IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Transport mode. The type is TransportMode.
    TransportMode interface{}

    // Enable MPLS encapsulation. The type is interface{}.
    Enable interface{}

    // Enable control word. The type is ControlWord.
    ControlWord interface{}

    // Sequencing.
    Sequencing L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing

    // Redundancy options for MPLS encapsulation.
    MplsRedundancy L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy

    // Preferred path.
    PreferredPath L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath

    // Load Balancing.
    LoadBalanceGroup L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup
}

func (mplsEncapsulation *L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation) GetEntityData() *types.CommonEntityData {
    mplsEncapsulation.EntityData.YFilter = mplsEncapsulation.YFilter
    mplsEncapsulation.EntityData.YangName = "mpls-encapsulation"
    mplsEncapsulation.EntityData.BundleName = "cisco_ios_xr"
    mplsEncapsulation.EntityData.ParentYangName = "pseudowire-class"
    mplsEncapsulation.EntityData.SegmentPath = "mpls-encapsulation"
    mplsEncapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsEncapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsEncapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsEncapsulation.EntityData.Children = types.NewOrderedMap()
    mplsEncapsulation.EntityData.Children.Append("sequencing", types.YChild{"Sequencing", &mplsEncapsulation.Sequencing})
    mplsEncapsulation.EntityData.Children.Append("mpls-redundancy", types.YChild{"MplsRedundancy", &mplsEncapsulation.MplsRedundancy})
    mplsEncapsulation.EntityData.Children.Append("preferred-path", types.YChild{"PreferredPath", &mplsEncapsulation.PreferredPath})
    mplsEncapsulation.EntityData.Children.Append("load-balance-group", types.YChild{"LoadBalanceGroup", &mplsEncapsulation.LoadBalanceGroup})
    mplsEncapsulation.EntityData.Leafs = types.NewOrderedMap()
    mplsEncapsulation.EntityData.Leafs.Append("pw-switching-tlv", types.YLeaf{"PwSwitchingTlv", mplsEncapsulation.PwSwitchingTlv})
    mplsEncapsulation.EntityData.Leafs.Append("static-tag-rewrite", types.YLeaf{"StaticTagRewrite", mplsEncapsulation.StaticTagRewrite})
    mplsEncapsulation.EntityData.Leafs.Append("signaling-protocol", types.YLeaf{"SignalingProtocol", mplsEncapsulation.SignalingProtocol})
    mplsEncapsulation.EntityData.Leafs.Append("vccv-type", types.YLeaf{"VccvType", mplsEncapsulation.VccvType})
    mplsEncapsulation.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", mplsEncapsulation.SourceAddress})
    mplsEncapsulation.EntityData.Leafs.Append("transport-mode", types.YLeaf{"TransportMode", mplsEncapsulation.TransportMode})
    mplsEncapsulation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mplsEncapsulation.Enable})
    mplsEncapsulation.EntityData.Leafs.Append("control-word", types.YLeaf{"ControlWord", mplsEncapsulation.ControlWord})

    mplsEncapsulation.EntityData.YListKeys = []string {}

    return &(mplsEncapsulation.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing
// Sequencing
type L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequencing. The type is MplsSequencing. The default value is off.
    Sequencing interface{}

    // Out of sequence threshold. The type is interface{} with range: 5..65535.
    // The default value is 5.
    ResyncThreshold interface{}
}

func (sequencing *L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing) GetEntityData() *types.CommonEntityData {
    sequencing.EntityData.YFilter = sequencing.YFilter
    sequencing.EntityData.YangName = "sequencing"
    sequencing.EntityData.BundleName = "cisco_ios_xr"
    sequencing.EntityData.ParentYangName = "mpls-encapsulation"
    sequencing.EntityData.SegmentPath = "sequencing"
    sequencing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sequencing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sequencing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sequencing.EntityData.Children = types.NewOrderedMap()
    sequencing.EntityData.Leafs = types.NewOrderedMap()
    sequencing.EntityData.Leafs.Append("sequencing", types.YLeaf{"Sequencing", sequencing.Sequencing})
    sequencing.EntityData.Leafs.Append("resync-threshold", types.YLeaf{"ResyncThreshold", sequencing.ResyncThreshold})

    sequencing.EntityData.YListKeys = []string {}

    return &(sequencing.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy
// Redundancy options for MPLS encapsulation
type L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force one-way PW redundancy behavior in Redundancy Group. The type is
    // interface{}.
    RedundancyOneWay interface{}

    // Initial delay before activating the redundant PW, in seconds. The type is
    // interface{} with range: 0..120. Units are second.
    RedundancyInitialDelay interface{}
}

func (mplsRedundancy *L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy) GetEntityData() *types.CommonEntityData {
    mplsRedundancy.EntityData.YFilter = mplsRedundancy.YFilter
    mplsRedundancy.EntityData.YangName = "mpls-redundancy"
    mplsRedundancy.EntityData.BundleName = "cisco_ios_xr"
    mplsRedundancy.EntityData.ParentYangName = "mpls-encapsulation"
    mplsRedundancy.EntityData.SegmentPath = "mpls-redundancy"
    mplsRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsRedundancy.EntityData.Children = types.NewOrderedMap()
    mplsRedundancy.EntityData.Leafs = types.NewOrderedMap()
    mplsRedundancy.EntityData.Leafs.Append("redundancy-one-way", types.YLeaf{"RedundancyOneWay", mplsRedundancy.RedundancyOneWay})
    mplsRedundancy.EntityData.Leafs.Append("redundancy-initial-delay", types.YLeaf{"RedundancyInitialDelay", mplsRedundancy.RedundancyInitialDelay})

    mplsRedundancy.EntityData.YListKeys = []string {}

    return &(mplsRedundancy.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath
// Preferred path
type L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Preferred Path Type. The type is PreferredPath.
    Type interface{}

    // Interface Tunnel number for preferred path. The type is interface{} with
    // range: 0..65535.
    InterfaceTunnelNumber interface{}

    // Fallback disable. The type is interface{}.
    FallbackDisable interface{}

    // Name of the SR TE Policy. The type is string with length: 1..60.
    SrtePolicy interface{}
}

func (preferredPath *L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath) GetEntityData() *types.CommonEntityData {
    preferredPath.EntityData.YFilter = preferredPath.YFilter
    preferredPath.EntityData.YangName = "preferred-path"
    preferredPath.EntityData.BundleName = "cisco_ios_xr"
    preferredPath.EntityData.ParentYangName = "mpls-encapsulation"
    preferredPath.EntityData.SegmentPath = "preferred-path"
    preferredPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preferredPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preferredPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preferredPath.EntityData.Children = types.NewOrderedMap()
    preferredPath.EntityData.Leafs = types.NewOrderedMap()
    preferredPath.EntityData.Leafs.Append("type", types.YLeaf{"Type", preferredPath.Type})
    preferredPath.EntityData.Leafs.Append("interface-tunnel-number", types.YLeaf{"InterfaceTunnelNumber", preferredPath.InterfaceTunnelNumber})
    preferredPath.EntityData.Leafs.Append("fallback-disable", types.YLeaf{"FallbackDisable", preferredPath.FallbackDisable})
    preferredPath.EntityData.Leafs.Append("srte-policy", types.YLeaf{"SrtePolicy", preferredPath.SrtePolicy})

    preferredPath.EntityData.YListKeys = []string {}

    return &(preferredPath.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup
// Load Balancing
type L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Legacy Flow Label TLV code. The type is FlowLabelTlvCode.
    FlowLabelLoadBalanceCode interface{}

    // Enable PW Label based Load Balancing. The type is LoadBalance.
    PwLabelLoadBalance interface{}

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance
}

func (loadBalanceGroup *L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup) GetEntityData() *types.CommonEntityData {
    loadBalanceGroup.EntityData.YFilter = loadBalanceGroup.YFilter
    loadBalanceGroup.EntityData.YangName = "load-balance-group"
    loadBalanceGroup.EntityData.BundleName = "cisco_ios_xr"
    loadBalanceGroup.EntityData.ParentYangName = "mpls-encapsulation"
    loadBalanceGroup.EntityData.SegmentPath = "load-balance-group"
    loadBalanceGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalanceGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalanceGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalanceGroup.EntityData.Children = types.NewOrderedMap()
    loadBalanceGroup.EntityData.Children.Append("flow-label-load-balance", types.YChild{"FlowLabelLoadBalance", &loadBalanceGroup.FlowLabelLoadBalance})
    loadBalanceGroup.EntityData.Leafs = types.NewOrderedMap()
    loadBalanceGroup.EntityData.Leafs.Append("flow-label-load-balance-code", types.YLeaf{"FlowLabelLoadBalanceCode", loadBalanceGroup.FlowLabelLoadBalanceCode})
    loadBalanceGroup.EntityData.Leafs.Append("pw-label-load-balance", types.YLeaf{"PwLabelLoadBalance", loadBalanceGroup.PwLabelLoadBalance})

    loadBalanceGroup.EntityData.YListKeys = []string {}

    return &(loadBalanceGroup.EntityData)
}

// L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "load-balance-group"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs.Append("flow-label", types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel})
    flowLabelLoadBalance.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabelLoadBalance.Static})

    flowLabelLoadBalance.EntityData.YListKeys = []string {}

    return &(flowLabelLoadBalance.EntityData)
}

// L2vpn_Database_VlanSwitches
// List of VLAN Switches
type L2vpn_Database_VlanSwitches struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN Switch. The type is slice of L2vpn_Database_VlanSwitches_VlanSwitch.
    VlanSwitch []*L2vpn_Database_VlanSwitches_VlanSwitch
}

func (vlanSwitches *L2vpn_Database_VlanSwitches) GetEntityData() *types.CommonEntityData {
    vlanSwitches.EntityData.YFilter = vlanSwitches.YFilter
    vlanSwitches.EntityData.YangName = "vlan-switches"
    vlanSwitches.EntityData.BundleName = "cisco_ios_xr"
    vlanSwitches.EntityData.ParentYangName = "database"
    vlanSwitches.EntityData.SegmentPath = "vlan-switches"
    vlanSwitches.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanSwitches.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanSwitches.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanSwitches.EntityData.Children = types.NewOrderedMap()
    vlanSwitches.EntityData.Children.Append("vlan-switch", types.YChild{"VlanSwitch", nil})
    for i := range vlanSwitches.VlanSwitch {
        vlanSwitches.EntityData.Children.Append(types.GetSegmentPath(vlanSwitches.VlanSwitch[i]), types.YChild{"VlanSwitch", vlanSwitches.VlanSwitch[i]})
    }
    vlanSwitches.EntityData.Leafs = types.NewOrderedMap()

    vlanSwitches.EntityData.YListKeys = []string {}

    return &(vlanSwitches.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch
// VLAN Switch
type L2vpn_Database_VlanSwitches_VlanSwitch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VLAN Switch. The type is string with
    // length: 1..32.
    Name interface{}

    // List of VLAN Switched Ports.
    VlanSwitchPorts L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts

    // Configure VLAN Switch VxLAN Ethernet VPN-ID ranges.
    VniRanges L2vpn_Database_VlanSwitches_VlanSwitch_VniRanges

    // Configure VLAN Switch VLAN ranges.
    VlanRanges L2vpn_Database_VlanSwitches_VlanSwitch_VlanRanges

    // Configure VLAN Switch Routed BVI Interface ranges.
    RoutedInterfaceRanges L2vpn_Database_VlanSwitches_VlanSwitch_RoutedInterfaceRanges

    // List of Bridge Domain.
    BridgeDomains L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains
}

func (vlanSwitch *L2vpn_Database_VlanSwitches_VlanSwitch) GetEntityData() *types.CommonEntityData {
    vlanSwitch.EntityData.YFilter = vlanSwitch.YFilter
    vlanSwitch.EntityData.YangName = "vlan-switch"
    vlanSwitch.EntityData.BundleName = "cisco_ios_xr"
    vlanSwitch.EntityData.ParentYangName = "vlan-switches"
    vlanSwitch.EntityData.SegmentPath = "vlan-switch" + types.AddKeyToken(vlanSwitch.Name, "name")
    vlanSwitch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanSwitch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanSwitch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanSwitch.EntityData.Children = types.NewOrderedMap()
    vlanSwitch.EntityData.Children.Append("vlan-switch-ports", types.YChild{"VlanSwitchPorts", &vlanSwitch.VlanSwitchPorts})
    vlanSwitch.EntityData.Children.Append("vni-ranges", types.YChild{"VniRanges", &vlanSwitch.VniRanges})
    vlanSwitch.EntityData.Children.Append("vlan-ranges", types.YChild{"VlanRanges", &vlanSwitch.VlanRanges})
    vlanSwitch.EntityData.Children.Append("routed-interface-ranges", types.YChild{"RoutedInterfaceRanges", &vlanSwitch.RoutedInterfaceRanges})
    vlanSwitch.EntityData.Children.Append("bridge-domains", types.YChild{"BridgeDomains", &vlanSwitch.BridgeDomains})
    vlanSwitch.EntityData.Leafs = types.NewOrderedMap()
    vlanSwitch.EntityData.Leafs.Append("name", types.YLeaf{"Name", vlanSwitch.Name})

    vlanSwitch.EntityData.YListKeys = []string {"Name"}

    return &(vlanSwitch.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts
// List of VLAN Switched Ports
type L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN Switched Port. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts_VlanSwitchPort.
    VlanSwitchPort []*L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts_VlanSwitchPort
}

func (vlanSwitchPorts *L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts) GetEntityData() *types.CommonEntityData {
    vlanSwitchPorts.EntityData.YFilter = vlanSwitchPorts.YFilter
    vlanSwitchPorts.EntityData.YangName = "vlan-switch-ports"
    vlanSwitchPorts.EntityData.BundleName = "cisco_ios_xr"
    vlanSwitchPorts.EntityData.ParentYangName = "vlan-switch"
    vlanSwitchPorts.EntityData.SegmentPath = "vlan-switch-ports"
    vlanSwitchPorts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanSwitchPorts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanSwitchPorts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanSwitchPorts.EntityData.Children = types.NewOrderedMap()
    vlanSwitchPorts.EntityData.Children.Append("vlan-switch-port", types.YChild{"VlanSwitchPort", nil})
    for i := range vlanSwitchPorts.VlanSwitchPort {
        vlanSwitchPorts.EntityData.Children.Append(types.GetSegmentPath(vlanSwitchPorts.VlanSwitchPort[i]), types.YChild{"VlanSwitchPort", vlanSwitchPorts.VlanSwitchPort[i]})
    }
    vlanSwitchPorts.EntityData.Leafs = types.NewOrderedMap()

    vlanSwitchPorts.EntityData.YListKeys = []string {}

    return &(vlanSwitchPorts.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts_VlanSwitchPort
// VLAN Switched Port
type L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts_VlanSwitchPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (vlanSwitchPort *L2vpn_Database_VlanSwitches_VlanSwitch_VlanSwitchPorts_VlanSwitchPort) GetEntityData() *types.CommonEntityData {
    vlanSwitchPort.EntityData.YFilter = vlanSwitchPort.YFilter
    vlanSwitchPort.EntityData.YangName = "vlan-switch-port"
    vlanSwitchPort.EntityData.BundleName = "cisco_ios_xr"
    vlanSwitchPort.EntityData.ParentYangName = "vlan-switch-ports"
    vlanSwitchPort.EntityData.SegmentPath = "vlan-switch-port" + types.AddKeyToken(vlanSwitchPort.InterfaceName, "interface-name")
    vlanSwitchPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanSwitchPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanSwitchPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanSwitchPort.EntityData.Children = types.NewOrderedMap()
    vlanSwitchPort.EntityData.Leafs = types.NewOrderedMap()
    vlanSwitchPort.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", vlanSwitchPort.InterfaceName})

    vlanSwitchPort.EntityData.YListKeys = []string {"InterfaceName"}

    return &(vlanSwitchPort.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_VniRanges
// Configure VLAN Switch VxLAN Ethernet VPN-ID
// ranges
type L2vpn_Database_VlanSwitches_VlanSwitch_VniRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value of VNI range #1. The type is interface{} with range:
    // 0..4294967295.
    VniRange1Min interface{}

    // Maximum value of VNI range #1. The type is interface{} with range:
    // 0..4294967295.
    VniRange1Max interface{}

    // Minimum value of VNI range #2. The type is interface{} with range:
    // 0..4294967295.
    VniRange2Min interface{}

    // Maximum value of VNI range #2. The type is interface{} with range:
    // 0..4294967295.
    VniRange2Max interface{}

    // Minimum value of VNI range #3. The type is interface{} with range:
    // 0..4294967295.
    VniRange3Min interface{}

    // Maximum value of VNI range #3. The type is interface{} with range:
    // 0..4294967295.
    VniRange3Max interface{}

    // Minimum value of VNI range #4. The type is interface{} with range:
    // 0..4294967295.
    VniRange4Min interface{}

    // Maximum value of VNI range #4. The type is interface{} with range:
    // 0..4294967295.
    VniRange4Max interface{}

    // Minimum value of VNI range #5. The type is interface{} with range:
    // 0..4294967295.
    VniRange5Min interface{}

    // Maximum value of VNI range #5. The type is interface{} with range:
    // 0..4294967295.
    VniRange5Max interface{}

    // Minimum value of VNI range #6. The type is interface{} with range:
    // 0..4294967295.
    VniRange6Min interface{}

    // Maximum value of VNI range #6. The type is interface{} with range:
    // 0..4294967295.
    VniRange6Max interface{}

    // Minimum value of VNI range #7. The type is interface{} with range:
    // 0..4294967295.
    VniRange7Min interface{}

    // Maximum value of VNI range #7. The type is interface{} with range:
    // 0..4294967295.
    VniRange7Max interface{}

    // Minimum value of VNI range #8. The type is interface{} with range:
    // 0..4294967295.
    VniRange8Min interface{}

    // Maximum value of VNI range #8. The type is interface{} with range:
    // 0..4294967295.
    VniRange8Max interface{}

    // Minimum value of VNI range #9. The type is interface{} with range:
    // 0..4294967295.
    VniRange9Min interface{}

    // Maximum value of VNI range #9. The type is interface{} with range:
    // 0..4294967295.
    VniRange9Max interface{}
}

func (vniRanges *L2vpn_Database_VlanSwitches_VlanSwitch_VniRanges) GetEntityData() *types.CommonEntityData {
    vniRanges.EntityData.YFilter = vniRanges.YFilter
    vniRanges.EntityData.YangName = "vni-ranges"
    vniRanges.EntityData.BundleName = "cisco_ios_xr"
    vniRanges.EntityData.ParentYangName = "vlan-switch"
    vniRanges.EntityData.SegmentPath = "vni-ranges"
    vniRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vniRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vniRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vniRanges.EntityData.Children = types.NewOrderedMap()
    vniRanges.EntityData.Leafs = types.NewOrderedMap()
    vniRanges.EntityData.Leafs.Append("vni-range1-min", types.YLeaf{"VniRange1Min", vniRanges.VniRange1Min})
    vniRanges.EntityData.Leafs.Append("vni-range1-max", types.YLeaf{"VniRange1Max", vniRanges.VniRange1Max})
    vniRanges.EntityData.Leafs.Append("vni-range2-min", types.YLeaf{"VniRange2Min", vniRanges.VniRange2Min})
    vniRanges.EntityData.Leafs.Append("vni-range2-max", types.YLeaf{"VniRange2Max", vniRanges.VniRange2Max})
    vniRanges.EntityData.Leafs.Append("vni-range3-min", types.YLeaf{"VniRange3Min", vniRanges.VniRange3Min})
    vniRanges.EntityData.Leafs.Append("vni-range3-max", types.YLeaf{"VniRange3Max", vniRanges.VniRange3Max})
    vniRanges.EntityData.Leafs.Append("vni-range4-min", types.YLeaf{"VniRange4Min", vniRanges.VniRange4Min})
    vniRanges.EntityData.Leafs.Append("vni-range4-max", types.YLeaf{"VniRange4Max", vniRanges.VniRange4Max})
    vniRanges.EntityData.Leafs.Append("vni-range5-min", types.YLeaf{"VniRange5Min", vniRanges.VniRange5Min})
    vniRanges.EntityData.Leafs.Append("vni-range5-max", types.YLeaf{"VniRange5Max", vniRanges.VniRange5Max})
    vniRanges.EntityData.Leafs.Append("vni-range6-min", types.YLeaf{"VniRange6Min", vniRanges.VniRange6Min})
    vniRanges.EntityData.Leafs.Append("vni-range6-max", types.YLeaf{"VniRange6Max", vniRanges.VniRange6Max})
    vniRanges.EntityData.Leafs.Append("vni-range7-min", types.YLeaf{"VniRange7Min", vniRanges.VniRange7Min})
    vniRanges.EntityData.Leafs.Append("vni-range7-max", types.YLeaf{"VniRange7Max", vniRanges.VniRange7Max})
    vniRanges.EntityData.Leafs.Append("vni-range8-min", types.YLeaf{"VniRange8Min", vniRanges.VniRange8Min})
    vniRanges.EntityData.Leafs.Append("vni-range8-max", types.YLeaf{"VniRange8Max", vniRanges.VniRange8Max})
    vniRanges.EntityData.Leafs.Append("vni-range9-min", types.YLeaf{"VniRange9Min", vniRanges.VniRange9Min})
    vniRanges.EntityData.Leafs.Append("vni-range9-max", types.YLeaf{"VniRange9Max", vniRanges.VniRange9Max})

    vniRanges.EntityData.YListKeys = []string {}

    return &(vniRanges.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_VlanRanges
// Configure VLAN Switch VLAN ranges
type L2vpn_Database_VlanSwitches_VlanSwitch_VlanRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value of VLAN range #1. The type is interface{} with range:
    // 1..4094.
    VlanRange1Min interface{}

    // Maximum value of VLAN range #1. The type is interface{} with range:
    // 1..4094.
    VlanRange1Max interface{}

    // Minimum value of VLAN range #2. The type is interface{} with range:
    // 1..4094.
    VlanRange2Min interface{}

    // Maximum value of VLAN range #2. The type is interface{} with range:
    // 1..4094.
    VlanRange2Max interface{}

    // Minimum value of VLAN range #3. The type is interface{} with range:
    // 1..4094.
    VlanRange3Min interface{}

    // Maximum value of VLAN range #3. The type is interface{} with range:
    // 1..4094.
    VlanRange3Max interface{}

    // Minimum value of VLAN range #4. The type is interface{} with range:
    // 1..4094.
    VlanRange4Min interface{}

    // Maximum value of VLAN range #4. The type is interface{} with range:
    // 1..4094.
    VlanRange4Max interface{}

    // Minimum value of VLAN range #5. The type is interface{} with range:
    // 1..4094.
    VlanRange5Min interface{}

    // Maximum value of VLAN range #5. The type is interface{} with range:
    // 1..4094.
    VlanRange5Max interface{}

    // Minimum value of VLAN range #6. The type is interface{} with range:
    // 1..4094.
    VlanRange6Min interface{}

    // Maximum value of VLAN range #6. The type is interface{} with range:
    // 1..4094.
    VlanRange6Max interface{}

    // Minimum value of VLAN range #7. The type is interface{} with range:
    // 1..4094.
    VlanRange7Min interface{}

    // Maximum value of VLAN range #7. The type is interface{} with range:
    // 1..4094.
    VlanRange7Max interface{}

    // Minimum value of VLAN range #8. The type is interface{} with range:
    // 1..4094.
    VlanRange8Min interface{}

    // Maximum value of VLAN range #8. The type is interface{} with range:
    // 1..4094.
    VlanRange8Max interface{}

    // Minimum value of VLAN range #9. The type is interface{} with range:
    // 1..4094.
    VlanRange9Min interface{}

    // Maximum value of VLAN range #9. The type is interface{} with range:
    // 1..4094.
    VlanRange9Max interface{}
}

func (vlanRanges *L2vpn_Database_VlanSwitches_VlanSwitch_VlanRanges) GetEntityData() *types.CommonEntityData {
    vlanRanges.EntityData.YFilter = vlanRanges.YFilter
    vlanRanges.EntityData.YangName = "vlan-ranges"
    vlanRanges.EntityData.BundleName = "cisco_ios_xr"
    vlanRanges.EntityData.ParentYangName = "vlan-switch"
    vlanRanges.EntityData.SegmentPath = "vlan-ranges"
    vlanRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRanges.EntityData.Children = types.NewOrderedMap()
    vlanRanges.EntityData.Leafs = types.NewOrderedMap()
    vlanRanges.EntityData.Leafs.Append("vlan-range1-min", types.YLeaf{"VlanRange1Min", vlanRanges.VlanRange1Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range1-max", types.YLeaf{"VlanRange1Max", vlanRanges.VlanRange1Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range2-min", types.YLeaf{"VlanRange2Min", vlanRanges.VlanRange2Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range2-max", types.YLeaf{"VlanRange2Max", vlanRanges.VlanRange2Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range3-min", types.YLeaf{"VlanRange3Min", vlanRanges.VlanRange3Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range3-max", types.YLeaf{"VlanRange3Max", vlanRanges.VlanRange3Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range4-min", types.YLeaf{"VlanRange4Min", vlanRanges.VlanRange4Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range4-max", types.YLeaf{"VlanRange4Max", vlanRanges.VlanRange4Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range5-min", types.YLeaf{"VlanRange5Min", vlanRanges.VlanRange5Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range5-max", types.YLeaf{"VlanRange5Max", vlanRanges.VlanRange5Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range6-min", types.YLeaf{"VlanRange6Min", vlanRanges.VlanRange6Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range6-max", types.YLeaf{"VlanRange6Max", vlanRanges.VlanRange6Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range7-min", types.YLeaf{"VlanRange7Min", vlanRanges.VlanRange7Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range7-max", types.YLeaf{"VlanRange7Max", vlanRanges.VlanRange7Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range8-min", types.YLeaf{"VlanRange8Min", vlanRanges.VlanRange8Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range8-max", types.YLeaf{"VlanRange8Max", vlanRanges.VlanRange8Max})
    vlanRanges.EntityData.Leafs.Append("vlan-range9-min", types.YLeaf{"VlanRange9Min", vlanRanges.VlanRange9Min})
    vlanRanges.EntityData.Leafs.Append("vlan-range9-max", types.YLeaf{"VlanRange9Max", vlanRanges.VlanRange9Max})

    vlanRanges.EntityData.YListKeys = []string {}

    return &(vlanRanges.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_RoutedInterfaceRanges
// Configure VLAN Switch Routed BVI Interface
// ranges
type L2vpn_Database_VlanSwitches_VlanSwitch_RoutedInterfaceRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value of Interface range #1. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange1Min interface{}

    // Maximum value of Interface range #1. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange1Max interface{}

    // Minimum value of Interface range #2. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange2Min interface{}

    // Maximum value of Interface range #2. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange2Max interface{}

    // Minimum value of Interface range #3. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange3Min interface{}

    // Maximum value of Interface range #3. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange3Max interface{}

    // Minimum value of Interface range #4. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange4Min interface{}

    // Maximum value of Interface range #4. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange4Max interface{}

    // Minimum value of Interface range #5. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange5Min interface{}

    // Maximum value of Interface range #5. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange5Max interface{}

    // Minimum value of Interface range #6. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange6Min interface{}

    // Maximum value of Interface range #6. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange6Max interface{}

    // Minimum value of Interface range #7. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange7Min interface{}

    // Maximum value of Interface range #7. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange7Max interface{}

    // Minimum value of Interface range #8. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange8Min interface{}

    // Maximum value of Interface range #8. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange8Max interface{}

    // Minimum value of Interface range #9. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange9Min interface{}

    // Maximum value of Interface range #9. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRange9Max interface{}
}

func (routedInterfaceRanges *L2vpn_Database_VlanSwitches_VlanSwitch_RoutedInterfaceRanges) GetEntityData() *types.CommonEntityData {
    routedInterfaceRanges.EntityData.YFilter = routedInterfaceRanges.YFilter
    routedInterfaceRanges.EntityData.YangName = "routed-interface-ranges"
    routedInterfaceRanges.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaceRanges.EntityData.ParentYangName = "vlan-switch"
    routedInterfaceRanges.EntityData.SegmentPath = "routed-interface-ranges"
    routedInterfaceRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaceRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaceRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaceRanges.EntityData.Children = types.NewOrderedMap()
    routedInterfaceRanges.EntityData.Leafs = types.NewOrderedMap()
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range1-min", types.YLeaf{"InterfaceRange1Min", routedInterfaceRanges.InterfaceRange1Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range1-max", types.YLeaf{"InterfaceRange1Max", routedInterfaceRanges.InterfaceRange1Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range2-min", types.YLeaf{"InterfaceRange2Min", routedInterfaceRanges.InterfaceRange2Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range2-max", types.YLeaf{"InterfaceRange2Max", routedInterfaceRanges.InterfaceRange2Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range3-min", types.YLeaf{"InterfaceRange3Min", routedInterfaceRanges.InterfaceRange3Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range3-max", types.YLeaf{"InterfaceRange3Max", routedInterfaceRanges.InterfaceRange3Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range4-min", types.YLeaf{"InterfaceRange4Min", routedInterfaceRanges.InterfaceRange4Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range4-max", types.YLeaf{"InterfaceRange4Max", routedInterfaceRanges.InterfaceRange4Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range5-min", types.YLeaf{"InterfaceRange5Min", routedInterfaceRanges.InterfaceRange5Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range5-max", types.YLeaf{"InterfaceRange5Max", routedInterfaceRanges.InterfaceRange5Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range6-min", types.YLeaf{"InterfaceRange6Min", routedInterfaceRanges.InterfaceRange6Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range6-max", types.YLeaf{"InterfaceRange6Max", routedInterfaceRanges.InterfaceRange6Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range7-min", types.YLeaf{"InterfaceRange7Min", routedInterfaceRanges.InterfaceRange7Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range7-max", types.YLeaf{"InterfaceRange7Max", routedInterfaceRanges.InterfaceRange7Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range8-min", types.YLeaf{"InterfaceRange8Min", routedInterfaceRanges.InterfaceRange8Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range8-max", types.YLeaf{"InterfaceRange8Max", routedInterfaceRanges.InterfaceRange8Max})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range9-min", types.YLeaf{"InterfaceRange9Min", routedInterfaceRanges.InterfaceRange9Min})
    routedInterfaceRanges.EntityData.Leafs.Append("interface-range9-max", types.YLeaf{"InterfaceRange9Max", routedInterfaceRanges.InterfaceRange9Max})

    routedInterfaceRanges.EntityData.YListKeys = []string {}

    return &(routedInterfaceRanges.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains
// List of Bridge Domain
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bridge domain. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain.
    BridgeDomain []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain
}

func (bridgeDomains *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomains.EntityData.ParentYangName = "vlan-switch"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomains.EntityData.Children = types.NewOrderedMap()
    bridgeDomains.EntityData.Children.Append("bridge-domain", types.YChild{"BridgeDomain", nil})
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children.Append(types.GetSegmentPath(bridgeDomains.BridgeDomain[i]), types.YChild{"BridgeDomain", bridgeDomains.BridgeDomain[i]})
    }
    bridgeDomains.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomains.EntityData.YListKeys = []string {}

    return &(bridgeDomains.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain
// bridge domain
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the bridge domain. The type is string with
    // length: 1..27.
    Name interface{}

    // Coupled-mode configuration. The type is interface{}.
    CoupledMode interface{}

    // shutdown the Bridge Domain. The type is interface{}.
    Shutdown interface{}

    // Disable Unknown Unicast flooding. The type is interface{}.
    FloodingUnknownUnicast interface{}

    // Disable IGMP Snooping. The type is interface{}.
    IgmpSnoopingDisable interface{}

    // Bridge Domain Transport mode. The type is BridgeDomainTransportMode.
    TransportMode interface{}

    // Attach MLD Snooping Profile Name. The type is string with length: 1..32.
    MldSnooping interface{}

    // Maximum transmission unit for this Bridge Domain. The type is interface{}
    // with range: 46..65535. Units are byte.
    BridgeDomainMtu interface{}

    // DHCPv4 Snooping profile name. The type is string with length: 1..32.
    Dhcp interface{}

    // Bridge-domain description Name. The type is string with length: 1..64.
    BridgeDescription interface{}

    // Attach IGMP Snooping Profile Name. The type is string with length: 1..32.
    IgmpSnooping interface{}

    // Disable flooding. The type is interface{}.
    Flooding interface{}

    // Storm Control.
    BdStormControls L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls

    // Bridge Domain VxLAN Network Identifier Table.
    MemberVnis L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis

    // MAC configuration commands.
    BridgeDomainMac L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac

    // nV Satellite.
    NvSatellite L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_NvSatellite

    // Bridge Domain PBB.
    BridgeDomainPbb L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb

    // Bridge Domain EVI Table.
    BridgeDomainEvis L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis

    // Specify the access virtual forwarding interface name.
    AccessVfis L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis

    // List of pseudowires.
    BdPseudowires L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires

    // Specify the virtual forwarding interface name.
    Vfis L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis

    // Bridge Domain EVPN VxLAN Network Identifier Table.
    BridgeDomainvnis L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis

    // Attachment Circuit table.
    BdAttachmentCircuits L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits

    // List of EVPN pseudowires.
    BdPseudowireEvpns L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns

    // IP Source Guard.
    IpSourceGuard L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_IpSourceGuard

    // Dynamic ARP Inspection.
    Dai L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai

    // Bridge Domain Routed Interface Table.
    RoutedInterfaces L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces
}

func (bridgeDomain *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomain.EntityData.ParentYangName = "bridge-domains"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + types.AddKeyToken(bridgeDomain.Name, "name")
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomain.EntityData.Children = types.NewOrderedMap()
    bridgeDomain.EntityData.Children.Append("bd-storm-controls", types.YChild{"BdStormControls", &bridgeDomain.BdStormControls})
    bridgeDomain.EntityData.Children.Append("member-vnis", types.YChild{"MemberVnis", &bridgeDomain.MemberVnis})
    bridgeDomain.EntityData.Children.Append("bridge-domain-mac", types.YChild{"BridgeDomainMac", &bridgeDomain.BridgeDomainMac})
    bridgeDomain.EntityData.Children.Append("nv-satellite", types.YChild{"NvSatellite", &bridgeDomain.NvSatellite})
    bridgeDomain.EntityData.Children.Append("bridge-domain-pbb", types.YChild{"BridgeDomainPbb", &bridgeDomain.BridgeDomainPbb})
    bridgeDomain.EntityData.Children.Append("bridge-domain-evis", types.YChild{"BridgeDomainEvis", &bridgeDomain.BridgeDomainEvis})
    bridgeDomain.EntityData.Children.Append("access-vfis", types.YChild{"AccessVfis", &bridgeDomain.AccessVfis})
    bridgeDomain.EntityData.Children.Append("bd-pseudowires", types.YChild{"BdPseudowires", &bridgeDomain.BdPseudowires})
    bridgeDomain.EntityData.Children.Append("vfis", types.YChild{"Vfis", &bridgeDomain.Vfis})
    bridgeDomain.EntityData.Children.Append("bridge-domainvnis", types.YChild{"BridgeDomainvnis", &bridgeDomain.BridgeDomainvnis})
    bridgeDomain.EntityData.Children.Append("bd-attachment-circuits", types.YChild{"BdAttachmentCircuits", &bridgeDomain.BdAttachmentCircuits})
    bridgeDomain.EntityData.Children.Append("bd-pseudowire-evpns", types.YChild{"BdPseudowireEvpns", &bridgeDomain.BdPseudowireEvpns})
    bridgeDomain.EntityData.Children.Append("ip-source-guard", types.YChild{"IpSourceGuard", &bridgeDomain.IpSourceGuard})
    bridgeDomain.EntityData.Children.Append("dai", types.YChild{"Dai", &bridgeDomain.Dai})
    bridgeDomain.EntityData.Children.Append("routed-interfaces", types.YChild{"RoutedInterfaces", &bridgeDomain.RoutedInterfaces})
    bridgeDomain.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomain.EntityData.Leafs.Append("name", types.YLeaf{"Name", bridgeDomain.Name})
    bridgeDomain.EntityData.Leafs.Append("coupled-mode", types.YLeaf{"CoupledMode", bridgeDomain.CoupledMode})
    bridgeDomain.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", bridgeDomain.Shutdown})
    bridgeDomain.EntityData.Leafs.Append("flooding-unknown-unicast", types.YLeaf{"FloodingUnknownUnicast", bridgeDomain.FloodingUnknownUnicast})
    bridgeDomain.EntityData.Leafs.Append("igmp-snooping-disable", types.YLeaf{"IgmpSnoopingDisable", bridgeDomain.IgmpSnoopingDisable})
    bridgeDomain.EntityData.Leafs.Append("transport-mode", types.YLeaf{"TransportMode", bridgeDomain.TransportMode})
    bridgeDomain.EntityData.Leafs.Append("mld-snooping", types.YLeaf{"MldSnooping", bridgeDomain.MldSnooping})
    bridgeDomain.EntityData.Leafs.Append("bridge-domain-mtu", types.YLeaf{"BridgeDomainMtu", bridgeDomain.BridgeDomainMtu})
    bridgeDomain.EntityData.Leafs.Append("dhcp", types.YLeaf{"Dhcp", bridgeDomain.Dhcp})
    bridgeDomain.EntityData.Leafs.Append("bridge-description", types.YLeaf{"BridgeDescription", bridgeDomain.BridgeDescription})
    bridgeDomain.EntityData.Leafs.Append("igmp-snooping", types.YLeaf{"IgmpSnooping", bridgeDomain.IgmpSnooping})
    bridgeDomain.EntityData.Leafs.Append("flooding", types.YLeaf{"Flooding", bridgeDomain.Flooding})

    bridgeDomain.EntityData.YListKeys = []string {"Name"}

    return &(bridgeDomain.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls
// Storm Control
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl.
    BdStormControl []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl
}

func (bdStormControls *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls) GetEntityData() *types.CommonEntityData {
    bdStormControls.EntityData.YFilter = bdStormControls.YFilter
    bdStormControls.EntityData.YangName = "bd-storm-controls"
    bdStormControls.EntityData.BundleName = "cisco_ios_xr"
    bdStormControls.EntityData.ParentYangName = "bridge-domain"
    bdStormControls.EntityData.SegmentPath = "bd-storm-controls"
    bdStormControls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdStormControls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdStormControls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdStormControls.EntityData.Children = types.NewOrderedMap()
    bdStormControls.EntityData.Children.Append("bd-storm-control", types.YChild{"BdStormControl", nil})
    for i := range bdStormControls.BdStormControl {
        bdStormControls.EntityData.Children.Append(types.GetSegmentPath(bdStormControls.BdStormControl[i]), types.YChild{"BdStormControl", bdStormControls.BdStormControl[i]})
    }
    bdStormControls.EntityData.Leafs = types.NewOrderedMap()

    bdStormControls.EntityData.YListKeys = []string {}

    return &(bdStormControls.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl
// Storm Control Type
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit
}

func (bdStormControl *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl) GetEntityData() *types.CommonEntityData {
    bdStormControl.EntityData.YFilter = bdStormControl.YFilter
    bdStormControl.EntityData.YangName = "bd-storm-control"
    bdStormControl.EntityData.BundleName = "cisco_ios_xr"
    bdStormControl.EntityData.ParentYangName = "bd-storm-controls"
    bdStormControl.EntityData.SegmentPath = "bd-storm-control" + types.AddKeyToken(bdStormControl.Sctype, "sctype")
    bdStormControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdStormControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdStormControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdStormControl.EntityData.Children = types.NewOrderedMap()
    bdStormControl.EntityData.Children.Append("storm-control-unit", types.YChild{"StormControlUnit", &bdStormControl.StormControlUnit})
    bdStormControl.EntityData.Leafs = types.NewOrderedMap()
    bdStormControl.EntityData.Leafs.Append("sctype", types.YLeaf{"Sctype", bdStormControl.Sctype})

    bdStormControl.EntityData.YListKeys = []string {"Sctype"}

    return &(bdStormControl.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit
// Specify units for Storm Control Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Kilobits Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 64..1280000. Units are
    // kbit/s.
    KbitsPerSec interface{}

    // Packets Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 1..160000. Units are
    // packet/s.
    PktsPerSec interface{}
}

func (stormControlUnit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bd-storm-control"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs.Append("kbits-per-sec", types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec})
    stormControlUnit.EntityData.Leafs.Append("pkts-per-sec", types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec})

    stormControlUnit.EntityData.YListKeys = []string {}

    return &(stormControlUnit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis
// Bridge Domain VxLAN Network Identifier Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain Member VxLAN Network Identifier. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni.
    MemberVni []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni
}

func (memberVnis *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis) GetEntityData() *types.CommonEntityData {
    memberVnis.EntityData.YFilter = memberVnis.YFilter
    memberVnis.EntityData.YangName = "member-vnis"
    memberVnis.EntityData.BundleName = "cisco_ios_xr"
    memberVnis.EntityData.ParentYangName = "bridge-domain"
    memberVnis.EntityData.SegmentPath = "member-vnis"
    memberVnis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVnis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVnis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVnis.EntityData.Children = types.NewOrderedMap()
    memberVnis.EntityData.Children.Append("member-vni", types.YChild{"MemberVni", nil})
    for i := range memberVnis.MemberVni {
        memberVnis.EntityData.Children.Append(types.GetSegmentPath(memberVnis.MemberVni[i]), types.YChild{"MemberVni", memberVnis.MemberVni[i]})
    }
    memberVnis.EntityData.Leafs = types.NewOrderedMap()

    memberVnis.EntityData.YListKeys = []string {}

    return &(memberVnis.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni
// Bridge Domain Member VxLAN Network Identifier
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VxLAN Network Identifier number. The type is
    // interface{} with range: 1..16777215.
    Vni interface{}

    // Static Mac Address Table.
    MemberVniStaticMacAddresses L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses
}

func (memberVni *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni) GetEntityData() *types.CommonEntityData {
    memberVni.EntityData.YFilter = memberVni.YFilter
    memberVni.EntityData.YangName = "member-vni"
    memberVni.EntityData.BundleName = "cisco_ios_xr"
    memberVni.EntityData.ParentYangName = "member-vnis"
    memberVni.EntityData.SegmentPath = "member-vni" + types.AddKeyToken(memberVni.Vni, "vni")
    memberVni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVni.EntityData.Children = types.NewOrderedMap()
    memberVni.EntityData.Children.Append("member-vni-static-mac-addresses", types.YChild{"MemberVniStaticMacAddresses", &memberVni.MemberVniStaticMacAddresses})
    memberVni.EntityData.Leafs = types.NewOrderedMap()
    memberVni.EntityData.Leafs.Append("vni", types.YLeaf{"Vni", memberVni.Vni})

    memberVni.EntityData.YListKeys = []string {"Vni"}

    return &(memberVni.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress.
    MemberVniStaticMacAddress []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress
}

func (memberVniStaticMacAddresses *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    memberVniStaticMacAddresses.EntityData.YFilter = memberVniStaticMacAddresses.YFilter
    memberVniStaticMacAddresses.EntityData.YangName = "member-vni-static-mac-addresses"
    memberVniStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    memberVniStaticMacAddresses.EntityData.ParentYangName = "member-vni"
    memberVniStaticMacAddresses.EntityData.SegmentPath = "member-vni-static-mac-addresses"
    memberVniStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVniStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVniStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVniStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    memberVniStaticMacAddresses.EntityData.Children.Append("member-vni-static-mac-address", types.YChild{"MemberVniStaticMacAddress", nil})
    for i := range memberVniStaticMacAddresses.MemberVniStaticMacAddress {
        memberVniStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(memberVniStaticMacAddresses.MemberVniStaticMacAddress[i]), types.YChild{"MemberVniStaticMacAddress", memberVniStaticMacAddresses.MemberVniStaticMacAddress[i]})
    }
    memberVniStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    memberVniStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(memberVniStaticMacAddresses.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Enable Static Mac Address Configuration. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopIp interface{}
}

func (memberVniStaticMacAddress *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress) GetEntityData() *types.CommonEntityData {
    memberVniStaticMacAddress.EntityData.YFilter = memberVniStaticMacAddress.YFilter
    memberVniStaticMacAddress.EntityData.YangName = "member-vni-static-mac-address"
    memberVniStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    memberVniStaticMacAddress.EntityData.ParentYangName = "member-vni-static-mac-addresses"
    memberVniStaticMacAddress.EntityData.SegmentPath = "member-vni-static-mac-address" + types.AddKeyToken(memberVniStaticMacAddress.MacAddress, "mac-address")
    memberVniStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVniStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVniStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVniStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    memberVniStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    memberVniStaticMacAddress.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", memberVniStaticMacAddress.MacAddress})
    memberVniStaticMacAddress.EntityData.Leafs.Append("next-hop-ip", types.YLeaf{"NextHopIp", memberVniStaticMacAddress.NextHopIp})

    memberVniStaticMacAddress.EntityData.YListKeys = []string {"MacAddress"}

    return &(memberVniStaticMacAddress.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac
// MAC configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mac withdraw sent from access PW to access PW. The type is interface{}.
    BdMacWithdrawRelay interface{}

    // MAC withdraw on Access PW. The type is interface{}.
    BdMacWithdrawAccessPwDisable interface{}

    // Disable MAC Flush when Port goes Down. The type is interface{}.
    BdMacPortDownFlush interface{}

    // Disable Mac Withdraw. The type is interface{}.
    BdMacWithdraw interface{}

    // MAC withdraw sent on bridge port down. The type is MacWithdrawBehavior.
    BdMacWithdrawBehavior interface{}

    // Mac Learning Type. The type is BdmacLearn.
    BdMacLearn interface{}

    // MAC-Limit configuration commands.
    BdMacLimit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit

    // Filter Mac Address.
    BdMacFilters L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters

    // MAC Secure.
    MacSecure L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure

    // MAC-Aging configuration commands.
    BdMacAging L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging
}

func (bridgeDomainMac *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac) GetEntityData() *types.CommonEntityData {
    bridgeDomainMac.EntityData.YFilter = bridgeDomainMac.YFilter
    bridgeDomainMac.EntityData.YangName = "bridge-domain-mac"
    bridgeDomainMac.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainMac.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainMac.EntityData.SegmentPath = "bridge-domain-mac"
    bridgeDomainMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainMac.EntityData.Children = types.NewOrderedMap()
    bridgeDomainMac.EntityData.Children.Append("bd-mac-limit", types.YChild{"BdMacLimit", &bridgeDomainMac.BdMacLimit})
    bridgeDomainMac.EntityData.Children.Append("bd-mac-filters", types.YChild{"BdMacFilters", &bridgeDomainMac.BdMacFilters})
    bridgeDomainMac.EntityData.Children.Append("mac-secure", types.YChild{"MacSecure", &bridgeDomainMac.MacSecure})
    bridgeDomainMac.EntityData.Children.Append("bd-mac-aging", types.YChild{"BdMacAging", &bridgeDomainMac.BdMacAging})
    bridgeDomainMac.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw-relay", types.YLeaf{"BdMacWithdrawRelay", bridgeDomainMac.BdMacWithdrawRelay})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw-access-pw-disable", types.YLeaf{"BdMacWithdrawAccessPwDisable", bridgeDomainMac.BdMacWithdrawAccessPwDisable})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-port-down-flush", types.YLeaf{"BdMacPortDownFlush", bridgeDomainMac.BdMacPortDownFlush})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw", types.YLeaf{"BdMacWithdraw", bridgeDomainMac.BdMacWithdraw})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-withdraw-behavior", types.YLeaf{"BdMacWithdrawBehavior", bridgeDomainMac.BdMacWithdrawBehavior})
    bridgeDomainMac.EntityData.Leafs.Append("bd-mac-learn", types.YLeaf{"BdMacLearn", bridgeDomainMac.BdMacLearn})

    bridgeDomainMac.EntityData.YListKeys = []string {}

    return &(bridgeDomainMac.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address limit enforcement action. The type is MacLimitAction.
    BdMacLimitAction interface{}

    // Mac Address Limit Notification. The type is MacNotification.
    BdMacLimitNotif interface{}

    // Number of MAC addresses after which MAC limit action is taken. The type is
    // interface{} with range: 0..4294967295.
    BdMacLimitMax interface{}
}

func (bdMacLimit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit) GetEntityData() *types.CommonEntityData {
    bdMacLimit.EntityData.YFilter = bdMacLimit.YFilter
    bdMacLimit.EntityData.YangName = "bd-mac-limit"
    bdMacLimit.EntityData.BundleName = "cisco_ios_xr"
    bdMacLimit.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacLimit.EntityData.SegmentPath = "bd-mac-limit"
    bdMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacLimit.EntityData.Children = types.NewOrderedMap()
    bdMacLimit.EntityData.Leafs = types.NewOrderedMap()
    bdMacLimit.EntityData.Leafs.Append("bd-mac-limit-action", types.YLeaf{"BdMacLimitAction", bdMacLimit.BdMacLimitAction})
    bdMacLimit.EntityData.Leafs.Append("bd-mac-limit-notif", types.YLeaf{"BdMacLimitNotif", bdMacLimit.BdMacLimitNotif})
    bdMacLimit.EntityData.Leafs.Append("bd-mac-limit-max", types.YLeaf{"BdMacLimitMax", bdMacLimit.BdMacLimitMax})

    bdMacLimit.EntityData.YListKeys = []string {}

    return &(bdMacLimit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters
// Filter Mac Address
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static MAC address. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter.
    BdMacFilter []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter
}

func (bdMacFilters *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters) GetEntityData() *types.CommonEntityData {
    bdMacFilters.EntityData.YFilter = bdMacFilters.YFilter
    bdMacFilters.EntityData.YangName = "bd-mac-filters"
    bdMacFilters.EntityData.BundleName = "cisco_ios_xr"
    bdMacFilters.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacFilters.EntityData.SegmentPath = "bd-mac-filters"
    bdMacFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacFilters.EntityData.Children = types.NewOrderedMap()
    bdMacFilters.EntityData.Children.Append("bd-mac-filter", types.YChild{"BdMacFilter", nil})
    for i := range bdMacFilters.BdMacFilter {
        bdMacFilters.EntityData.Children.Append(types.GetSegmentPath(bdMacFilters.BdMacFilter[i]), types.YChild{"BdMacFilter", bdMacFilters.BdMacFilter[i]})
    }
    bdMacFilters.EntityData.Leafs = types.NewOrderedMap()

    bdMacFilters.EntityData.YListKeys = []string {}

    return &(bdMacFilters.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter
// Static MAC address
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}

    // MAC address for filtering. The type is interface{}.
    Drop interface{}
}

func (bdMacFilter *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter) GetEntityData() *types.CommonEntityData {
    bdMacFilter.EntityData.YFilter = bdMacFilter.YFilter
    bdMacFilter.EntityData.YangName = "bd-mac-filter"
    bdMacFilter.EntityData.BundleName = "cisco_ios_xr"
    bdMacFilter.EntityData.ParentYangName = "bd-mac-filters"
    bdMacFilter.EntityData.SegmentPath = "bd-mac-filter" + types.AddKeyToken(bdMacFilter.Address, "address")
    bdMacFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacFilter.EntityData.Children = types.NewOrderedMap()
    bdMacFilter.EntityData.Leafs = types.NewOrderedMap()
    bdMacFilter.EntityData.Leafs.Append("address", types.YLeaf{"Address", bdMacFilter.Address})
    bdMacFilter.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", bdMacFilter.Drop})

    bdMacFilter.EntityData.YListKeys = []string {"Address"}

    return &(bdMacFilter.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure
// MAC Secure
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is interface{}.
    Logging interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}

    // MAC Secure Threshold. The type is interface{}.
    Threshold interface{}
}

func (macSecure *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure) GetEntityData() *types.CommonEntityData {
    macSecure.EntityData.YFilter = macSecure.YFilter
    macSecure.EntityData.YangName = "mac-secure"
    macSecure.EntityData.BundleName = "cisco_ios_xr"
    macSecure.EntityData.ParentYangName = "bridge-domain-mac"
    macSecure.EntityData.SegmentPath = "mac-secure"
    macSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSecure.EntityData.Children = types.NewOrderedMap()
    macSecure.EntityData.Leafs = types.NewOrderedMap()
    macSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", macSecure.Logging})
    macSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", macSecure.Action})
    macSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", macSecure.Enable})
    macSecure.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", macSecure.Threshold})

    macSecure.EntityData.YListKeys = []string {}

    return &(macSecure.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging
// MAC-Aging configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    BdMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    BdMacAgingTime interface{}
}

func (bdMacAging *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging) GetEntityData() *types.CommonEntityData {
    bdMacAging.EntityData.YFilter = bdMacAging.YFilter
    bdMacAging.EntityData.YangName = "bd-mac-aging"
    bdMacAging.EntityData.BundleName = "cisco_ios_xr"
    bdMacAging.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacAging.EntityData.SegmentPath = "bd-mac-aging"
    bdMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacAging.EntityData.Children = types.NewOrderedMap()
    bdMacAging.EntityData.Leafs = types.NewOrderedMap()
    bdMacAging.EntityData.Leafs.Append("bd-mac-aging-type", types.YLeaf{"BdMacAgingType", bdMacAging.BdMacAgingType})
    bdMacAging.EntityData.Leafs.Append("bd-mac-aging-time", types.YLeaf{"BdMacAgingTime", bdMacAging.BdMacAgingTime})

    bdMacAging.EntityData.YListKeys = []string {}

    return &(bdMacAging.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_NvSatellite
// nV Satellite
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_NvSatellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPv4 Multicast Offload to Satellite Nodes. The type is interface{}.
    OffloadIpv4MulticastEnable interface{}

    // Enable nV Satellite Settings. The type is interface{}.
    Enable interface{}
}

func (nvSatellite *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_NvSatellite) GetEntityData() *types.CommonEntityData {
    nvSatellite.EntityData.YFilter = nvSatellite.YFilter
    nvSatellite.EntityData.YangName = "nv-satellite"
    nvSatellite.EntityData.BundleName = "cisco_ios_xr"
    nvSatellite.EntityData.ParentYangName = "bridge-domain"
    nvSatellite.EntityData.SegmentPath = "nv-satellite"
    nvSatellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellite.EntityData.Children = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs.Append("offload-ipv4-multicast-enable", types.YLeaf{"OffloadIpv4MulticastEnable", nvSatellite.OffloadIpv4MulticastEnable})
    nvSatellite.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", nvSatellite.Enable})

    nvSatellite.EntityData.YListKeys = []string {}

    return &(nvSatellite.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb
// Bridge Domain PBB
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Edge.
    PbbEdges L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges

    // PBB Core.
    PbbCore L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore
}

func (bridgeDomainPbb *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb) GetEntityData() *types.CommonEntityData {
    bridgeDomainPbb.EntityData.YFilter = bridgeDomainPbb.YFilter
    bridgeDomainPbb.EntityData.YangName = "bridge-domain-pbb"
    bridgeDomainPbb.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainPbb.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainPbb.EntityData.SegmentPath = "bridge-domain-pbb"
    bridgeDomainPbb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainPbb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainPbb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainPbb.EntityData.Children = types.NewOrderedMap()
    bridgeDomainPbb.EntityData.Children.Append("pbb-edges", types.YChild{"PbbEdges", &bridgeDomainPbb.PbbEdges})
    bridgeDomainPbb.EntityData.Children.Append("pbb-core", types.YChild{"PbbCore", &bridgeDomainPbb.PbbCore})
    bridgeDomainPbb.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainPbb.EntityData.YListKeys = []string {}

    return &(bridgeDomainPbb.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges
// PBB Edge
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure BD as PBB Edge with ISID and associated PBB Core BD. The type is
    // slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge.
    PbbEdge []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge
}

func (pbbEdges *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges) GetEntityData() *types.CommonEntityData {
    pbbEdges.EntityData.YFilter = pbbEdges.YFilter
    pbbEdges.EntityData.YangName = "pbb-edges"
    pbbEdges.EntityData.BundleName = "cisco_ios_xr"
    pbbEdges.EntityData.ParentYangName = "bridge-domain-pbb"
    pbbEdges.EntityData.SegmentPath = "pbb-edges"
    pbbEdges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdges.EntityData.Children = types.NewOrderedMap()
    pbbEdges.EntityData.Children.Append("pbb-edge", types.YChild{"PbbEdge", nil})
    for i := range pbbEdges.PbbEdge {
        pbbEdges.EntityData.Children.Append(types.GetSegmentPath(pbbEdges.PbbEdge[i]), types.YChild{"PbbEdge", pbbEdges.PbbEdge[i]})
    }
    pbbEdges.EntityData.Leafs = types.NewOrderedMap()

    pbbEdges.EntityData.YListKeys = []string {}

    return &(pbbEdges.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge
// Configure BD as PBB Edge with ISID and
// associated PBB Core BD
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ISID. The type is interface{} with range:
    // 256..16777214.
    Isid interface{}

    // This attribute is a key. Core BD Name. The type is string with length:
    // 1..27.
    CoreBdName interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    PbbEdgeIgmpProfile interface{}

    // Configure Unknown Unicast BMAC address for PBB Edge Port. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    UnknownUnicastBmac interface{}

    // Split Horizon Group.
    PbbEdgeSplitHorizonGroup L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup

    // PBB Static Mac Address Mapping Table.
    PbbStaticMacMappings L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings

    // Attach a DHCP profile.
    PbbEdgeDhcpProfile L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile

    // MAC configuration commands.
    PbbEdgeMac L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac
}

func (pbbEdge *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge) GetEntityData() *types.CommonEntityData {
    pbbEdge.EntityData.YFilter = pbbEdge.YFilter
    pbbEdge.EntityData.YangName = "pbb-edge"
    pbbEdge.EntityData.BundleName = "cisco_ios_xr"
    pbbEdge.EntityData.ParentYangName = "pbb-edges"
    pbbEdge.EntityData.SegmentPath = "pbb-edge" + types.AddKeyToken(pbbEdge.Isid, "isid") + types.AddKeyToken(pbbEdge.CoreBdName, "core-bd-name")
    pbbEdge.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdge.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdge.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdge.EntityData.Children = types.NewOrderedMap()
    pbbEdge.EntityData.Children.Append("pbb-edge-split-horizon-group", types.YChild{"PbbEdgeSplitHorizonGroup", &pbbEdge.PbbEdgeSplitHorizonGroup})
    pbbEdge.EntityData.Children.Append("pbb-static-mac-mappings", types.YChild{"PbbStaticMacMappings", &pbbEdge.PbbStaticMacMappings})
    pbbEdge.EntityData.Children.Append("pbb-edge-dhcp-profile", types.YChild{"PbbEdgeDhcpProfile", &pbbEdge.PbbEdgeDhcpProfile})
    pbbEdge.EntityData.Children.Append("pbb-edge-mac", types.YChild{"PbbEdgeMac", &pbbEdge.PbbEdgeMac})
    pbbEdge.EntityData.Leafs = types.NewOrderedMap()
    pbbEdge.EntityData.Leafs.Append("isid", types.YLeaf{"Isid", pbbEdge.Isid})
    pbbEdge.EntityData.Leafs.Append("core-bd-name", types.YLeaf{"CoreBdName", pbbEdge.CoreBdName})
    pbbEdge.EntityData.Leafs.Append("pbb-edge-igmp-profile", types.YLeaf{"PbbEdgeIgmpProfile", pbbEdge.PbbEdgeIgmpProfile})
    pbbEdge.EntityData.Leafs.Append("unknown-unicast-bmac", types.YLeaf{"UnknownUnicastBmac", pbbEdge.UnknownUnicastBmac})

    pbbEdge.EntityData.YListKeys = []string {"Isid", "CoreBdName"}

    return &(pbbEdge.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup
// Split Horizon Group
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable split horizon group. The type is interface{}.
    Disable interface{}
}

func (pbbEdgeSplitHorizonGroup *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    pbbEdgeSplitHorizonGroup.EntityData.YFilter = pbbEdgeSplitHorizonGroup.YFilter
    pbbEdgeSplitHorizonGroup.EntityData.YangName = "pbb-edge-split-horizon-group"
    pbbEdgeSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeSplitHorizonGroup.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeSplitHorizonGroup.EntityData.SegmentPath = "pbb-edge-split-horizon-group"
    pbbEdgeSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeSplitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    pbbEdgeSplitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeSplitHorizonGroup.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pbbEdgeSplitHorizonGroup.Disable})

    pbbEdgeSplitHorizonGroup.EntityData.YListKeys = []string {}

    return &(pbbEdgeSplitHorizonGroup.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings
// PBB Static Mac Address Mapping Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Static Mac Address Mapping Configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping.
    PbbStaticMacMapping []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping
}

func (pbbStaticMacMappings *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings) GetEntityData() *types.CommonEntityData {
    pbbStaticMacMappings.EntityData.YFilter = pbbStaticMacMappings.YFilter
    pbbStaticMacMappings.EntityData.YangName = "pbb-static-mac-mappings"
    pbbStaticMacMappings.EntityData.BundleName = "cisco_ios_xr"
    pbbStaticMacMappings.EntityData.ParentYangName = "pbb-edge"
    pbbStaticMacMappings.EntityData.SegmentPath = "pbb-static-mac-mappings"
    pbbStaticMacMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbStaticMacMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbStaticMacMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbStaticMacMappings.EntityData.Children = types.NewOrderedMap()
    pbbStaticMacMappings.EntityData.Children.Append("pbb-static-mac-mapping", types.YChild{"PbbStaticMacMapping", nil})
    for i := range pbbStaticMacMappings.PbbStaticMacMapping {
        pbbStaticMacMappings.EntityData.Children.Append(types.GetSegmentPath(pbbStaticMacMappings.PbbStaticMacMapping[i]), types.YChild{"PbbStaticMacMapping", pbbStaticMacMappings.PbbStaticMacMapping[i]})
    }
    pbbStaticMacMappings.EntityData.Leafs = types.NewOrderedMap()

    pbbStaticMacMappings.EntityData.YListKeys = []string {}

    return &(pbbStaticMacMappings.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping
// PBB Static Mac Address Mapping
// Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}

    // Static backbone MAC address to map with. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PbbStaticMacMappingBmac interface{}
}

func (pbbStaticMacMapping *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping) GetEntityData() *types.CommonEntityData {
    pbbStaticMacMapping.EntityData.YFilter = pbbStaticMacMapping.YFilter
    pbbStaticMacMapping.EntityData.YangName = "pbb-static-mac-mapping"
    pbbStaticMacMapping.EntityData.BundleName = "cisco_ios_xr"
    pbbStaticMacMapping.EntityData.ParentYangName = "pbb-static-mac-mappings"
    pbbStaticMacMapping.EntityData.SegmentPath = "pbb-static-mac-mapping" + types.AddKeyToken(pbbStaticMacMapping.Address, "address")
    pbbStaticMacMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbStaticMacMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbStaticMacMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbStaticMacMapping.EntityData.Children = types.NewOrderedMap()
    pbbStaticMacMapping.EntityData.Leafs = types.NewOrderedMap()
    pbbStaticMacMapping.EntityData.Leafs.Append("address", types.YLeaf{"Address", pbbStaticMacMapping.Address})
    pbbStaticMacMapping.EntityData.Leafs.Append("pbb-static-mac-mapping-bmac", types.YLeaf{"PbbStaticMacMappingBmac", pbbStaticMacMapping.PbbStaticMacMappingBmac})

    pbbStaticMacMapping.EntityData.YListKeys = []string {"Address"}

    return &(pbbStaticMacMapping.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile
// Attach a DHCP profile
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pbbEdgeDhcpProfile *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile) GetEntityData() *types.CommonEntityData {
    pbbEdgeDhcpProfile.EntityData.YFilter = pbbEdgeDhcpProfile.YFilter
    pbbEdgeDhcpProfile.EntityData.YangName = "pbb-edge-dhcp-profile"
    pbbEdgeDhcpProfile.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeDhcpProfile.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeDhcpProfile.EntityData.SegmentPath = "pbb-edge-dhcp-profile"
    pbbEdgeDhcpProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeDhcpProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeDhcpProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeDhcpProfile.EntityData.Children = types.NewOrderedMap()
    pbbEdgeDhcpProfile.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeDhcpProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", pbbEdgeDhcpProfile.ProfileId})
    pbbEdgeDhcpProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", pbbEdgeDhcpProfile.DhcpSnoopingId})

    pbbEdgeDhcpProfile.EntityData.YListKeys = []string {}

    return &(pbbEdgeDhcpProfile.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac
// MAC configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Mac Learning. The type is MacLearn.
    PbbEdgeMacLearning interface{}

    // MAC-Limit configuration commands.
    PbbEdgeMacLimit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit

    // MAC-Aging configuration commands.
    PbbEdgeMacAging L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging

    // MAC Secure.
    PbbEdgeMacSecure L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure
}

func (pbbEdgeMac *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac) GetEntityData() *types.CommonEntityData {
    pbbEdgeMac.EntityData.YFilter = pbbEdgeMac.YFilter
    pbbEdgeMac.EntityData.YangName = "pbb-edge-mac"
    pbbEdgeMac.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMac.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeMac.EntityData.SegmentPath = "pbb-edge-mac"
    pbbEdgeMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMac.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMac.EntityData.Children.Append("pbb-edge-mac-limit", types.YChild{"PbbEdgeMacLimit", &pbbEdgeMac.PbbEdgeMacLimit})
    pbbEdgeMac.EntityData.Children.Append("pbb-edge-mac-aging", types.YChild{"PbbEdgeMacAging", &pbbEdgeMac.PbbEdgeMacAging})
    pbbEdgeMac.EntityData.Children.Append("pbb-edge-mac-secure", types.YChild{"PbbEdgeMacSecure", &pbbEdgeMac.PbbEdgeMacSecure})
    pbbEdgeMac.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMac.EntityData.Leafs.Append("pbb-edge-mac-learning", types.YLeaf{"PbbEdgeMacLearning", pbbEdgeMac.PbbEdgeMacLearning})

    pbbEdgeMac.EntityData.YListKeys = []string {}

    return &(pbbEdgeMac.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address limit enforcement action. The type is MacLimitAction.
    PbbEdgeMacLimitAction interface{}

    // Number of MAC addresses after which MAC limit action is taken. The type is
    // interface{} with range: 0..4294967295.
    PbbEdgeMacLimitMax interface{}

    // MAC address limit notification action. The type is MacNotification.
    PbbEdgeMacLimitNotif interface{}
}

func (pbbEdgeMacLimit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacLimit.EntityData.YFilter = pbbEdgeMacLimit.YFilter
    pbbEdgeMacLimit.EntityData.YangName = "pbb-edge-mac-limit"
    pbbEdgeMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacLimit.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacLimit.EntityData.SegmentPath = "pbb-edge-mac-limit"
    pbbEdgeMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacLimit.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMacLimit.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMacLimit.EntityData.Leafs.Append("pbb-edge-mac-limit-action", types.YLeaf{"PbbEdgeMacLimitAction", pbbEdgeMacLimit.PbbEdgeMacLimitAction})
    pbbEdgeMacLimit.EntityData.Leafs.Append("pbb-edge-mac-limit-max", types.YLeaf{"PbbEdgeMacLimitMax", pbbEdgeMacLimit.PbbEdgeMacLimitMax})
    pbbEdgeMacLimit.EntityData.Leafs.Append("pbb-edge-mac-limit-notif", types.YLeaf{"PbbEdgeMacLimitNotif", pbbEdgeMacLimit.PbbEdgeMacLimitNotif})

    pbbEdgeMacLimit.EntityData.YListKeys = []string {}

    return &(pbbEdgeMacLimit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging
// MAC-Aging configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PbbEdgeMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    PbbEdgeMacAgingTime interface{}
}

func (pbbEdgeMacAging *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacAging.EntityData.YFilter = pbbEdgeMacAging.YFilter
    pbbEdgeMacAging.EntityData.YangName = "pbb-edge-mac-aging"
    pbbEdgeMacAging.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacAging.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacAging.EntityData.SegmentPath = "pbb-edge-mac-aging"
    pbbEdgeMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacAging.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMacAging.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMacAging.EntityData.Leafs.Append("pbb-edge-mac-aging-type", types.YLeaf{"PbbEdgeMacAgingType", pbbEdgeMacAging.PbbEdgeMacAgingType})
    pbbEdgeMacAging.EntityData.Leafs.Append("pbb-edge-mac-aging-time", types.YLeaf{"PbbEdgeMacAgingTime", pbbEdgeMacAging.PbbEdgeMacAgingTime})

    pbbEdgeMacAging.EntityData.YListKeys = []string {}

    return &(pbbEdgeMacAging.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure
// MAC Secure
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is L2vpnLogging.
    Logging interface{}

    // Disable Virtual instance port MAC Secure. The type is interface{}.
    Disable interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}

    // Accept Virtual instance port to be shutdown on mac violation. The type is
    // interface{}.
    AcceptShutdown interface{}
}

func (pbbEdgeMacSecure *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacSecure.EntityData.YFilter = pbbEdgeMacSecure.YFilter
    pbbEdgeMacSecure.EntityData.YangName = "pbb-edge-mac-secure"
    pbbEdgeMacSecure.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacSecure.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacSecure.EntityData.SegmentPath = "pbb-edge-mac-secure"
    pbbEdgeMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacSecure.EntityData.Children = types.NewOrderedMap()
    pbbEdgeMacSecure.EntityData.Leafs = types.NewOrderedMap()
    pbbEdgeMacSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pbbEdgeMacSecure.Logging})
    pbbEdgeMacSecure.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pbbEdgeMacSecure.Disable})
    pbbEdgeMacSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", pbbEdgeMacSecure.Action})
    pbbEdgeMacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pbbEdgeMacSecure.Enable})
    pbbEdgeMacSecure.EntityData.Leafs.Append("accept-shutdown", types.YLeaf{"AcceptShutdown", pbbEdgeMacSecure.AcceptShutdown})

    pbbEdgeMacSecure.EntityData.YListKeys = []string {}

    return &(pbbEdgeMacSecure.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore
// PBB Core
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enabling MMRP PBB-VPLS Flood Optimization. The type is interface{}.
    PbbCoreMmrpFloodOptimization interface{}

    // VLAN ID to push. The type is interface{} with range: 1..4094.
    VlanId interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    PbbCoreIgmpProfile interface{}

    // Enable Bridge Domain PBB Core Configuration. The type is interface{}.
    Enable interface{}

    // MAC configuration commands.
    PbbCoreMac L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac

    // PBB Core EVI Table.
    PbbCoreEvis L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis

    // Attach a DHCP profile.
    PbbCoreDhcpProfile L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile
}

func (pbbCore *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore) GetEntityData() *types.CommonEntityData {
    pbbCore.EntityData.YFilter = pbbCore.YFilter
    pbbCore.EntityData.YangName = "pbb-core"
    pbbCore.EntityData.BundleName = "cisco_ios_xr"
    pbbCore.EntityData.ParentYangName = "bridge-domain-pbb"
    pbbCore.EntityData.SegmentPath = "pbb-core"
    pbbCore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCore.EntityData.Children = types.NewOrderedMap()
    pbbCore.EntityData.Children.Append("pbb-core-mac", types.YChild{"PbbCoreMac", &pbbCore.PbbCoreMac})
    pbbCore.EntityData.Children.Append("pbb-core-evis", types.YChild{"PbbCoreEvis", &pbbCore.PbbCoreEvis})
    pbbCore.EntityData.Children.Append("pbb-core-dhcp-profile", types.YChild{"PbbCoreDhcpProfile", &pbbCore.PbbCoreDhcpProfile})
    pbbCore.EntityData.Leafs = types.NewOrderedMap()
    pbbCore.EntityData.Leafs.Append("pbb-core-mmrp-flood-optimization", types.YLeaf{"PbbCoreMmrpFloodOptimization", pbbCore.PbbCoreMmrpFloodOptimization})
    pbbCore.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", pbbCore.VlanId})
    pbbCore.EntityData.Leafs.Append("pbb-core-igmp-profile", types.YLeaf{"PbbCoreIgmpProfile", pbbCore.PbbCoreIgmpProfile})
    pbbCore.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pbbCore.Enable})

    pbbCore.EntityData.YListKeys = []string {}

    return &(pbbCore.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac
// MAC configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Mac Learning. The type is MacLearn.
    PbbCoreMacLearning interface{}

    // MAC-Aging configuration commands.
    PbbCoreMacAging L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging

    // MAC-Limit configuration commands.
    PbbCoreMacLimit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit
}

func (pbbCoreMac *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac) GetEntityData() *types.CommonEntityData {
    pbbCoreMac.EntityData.YFilter = pbbCoreMac.YFilter
    pbbCoreMac.EntityData.YangName = "pbb-core-mac"
    pbbCoreMac.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMac.EntityData.ParentYangName = "pbb-core"
    pbbCoreMac.EntityData.SegmentPath = "pbb-core-mac"
    pbbCoreMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMac.EntityData.Children = types.NewOrderedMap()
    pbbCoreMac.EntityData.Children.Append("pbb-core-mac-aging", types.YChild{"PbbCoreMacAging", &pbbCoreMac.PbbCoreMacAging})
    pbbCoreMac.EntityData.Children.Append("pbb-core-mac-limit", types.YChild{"PbbCoreMacLimit", &pbbCoreMac.PbbCoreMacLimit})
    pbbCoreMac.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreMac.EntityData.Leafs.Append("pbb-core-mac-learning", types.YLeaf{"PbbCoreMacLearning", pbbCoreMac.PbbCoreMacLearning})

    pbbCoreMac.EntityData.YListKeys = []string {}

    return &(pbbCoreMac.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging
// MAC-Aging configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PbbCoreMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    PbbCoreMacAgingTime interface{}
}

func (pbbCoreMacAging *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging) GetEntityData() *types.CommonEntityData {
    pbbCoreMacAging.EntityData.YFilter = pbbCoreMacAging.YFilter
    pbbCoreMacAging.EntityData.YangName = "pbb-core-mac-aging"
    pbbCoreMacAging.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMacAging.EntityData.ParentYangName = "pbb-core-mac"
    pbbCoreMacAging.EntityData.SegmentPath = "pbb-core-mac-aging"
    pbbCoreMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMacAging.EntityData.Children = types.NewOrderedMap()
    pbbCoreMacAging.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreMacAging.EntityData.Leafs.Append("pbb-core-mac-aging-type", types.YLeaf{"PbbCoreMacAgingType", pbbCoreMacAging.PbbCoreMacAgingType})
    pbbCoreMacAging.EntityData.Leafs.Append("pbb-core-mac-aging-time", types.YLeaf{"PbbCoreMacAgingTime", pbbCoreMacAging.PbbCoreMacAgingTime})

    pbbCoreMacAging.EntityData.YListKeys = []string {}

    return &(pbbCoreMacAging.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of MAC addresses after which MAC limit action is taken. The type is
    // interface{} with range: 0..4294967295.
    PbbCoreMacLimitMax interface{}

    // MAC address limit notification action. The type is MacNotification.
    PbbCoreMacLimitNotif interface{}

    // MAC address limit enforcement action. The type is MacLimitAction.
    PbbCoreMacLimitAction interface{}
}

func (pbbCoreMacLimit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit) GetEntityData() *types.CommonEntityData {
    pbbCoreMacLimit.EntityData.YFilter = pbbCoreMacLimit.YFilter
    pbbCoreMacLimit.EntityData.YangName = "pbb-core-mac-limit"
    pbbCoreMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMacLimit.EntityData.ParentYangName = "pbb-core-mac"
    pbbCoreMacLimit.EntityData.SegmentPath = "pbb-core-mac-limit"
    pbbCoreMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMacLimit.EntityData.Children = types.NewOrderedMap()
    pbbCoreMacLimit.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreMacLimit.EntityData.Leafs.Append("pbb-core-mac-limit-max", types.YLeaf{"PbbCoreMacLimitMax", pbbCoreMacLimit.PbbCoreMacLimitMax})
    pbbCoreMacLimit.EntityData.Leafs.Append("pbb-core-mac-limit-notif", types.YLeaf{"PbbCoreMacLimitNotif", pbbCoreMacLimit.PbbCoreMacLimitNotif})
    pbbCoreMacLimit.EntityData.Leafs.Append("pbb-core-mac-limit-action", types.YLeaf{"PbbCoreMacLimitAction", pbbCoreMacLimit.PbbCoreMacLimitAction})

    pbbCoreMacLimit.EntityData.YListKeys = []string {}

    return &(pbbCoreMacLimit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis
// PBB Core EVI Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Core EVI. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi.
    PbbCoreEvi []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi
}

func (pbbCoreEvis *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis) GetEntityData() *types.CommonEntityData {
    pbbCoreEvis.EntityData.YFilter = pbbCoreEvis.YFilter
    pbbCoreEvis.EntityData.YangName = "pbb-core-evis"
    pbbCoreEvis.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreEvis.EntityData.ParentYangName = "pbb-core"
    pbbCoreEvis.EntityData.SegmentPath = "pbb-core-evis"
    pbbCoreEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreEvis.EntityData.Children = types.NewOrderedMap()
    pbbCoreEvis.EntityData.Children.Append("pbb-core-evi", types.YChild{"PbbCoreEvi", nil})
    for i := range pbbCoreEvis.PbbCoreEvi {
        pbbCoreEvis.EntityData.Children.Append(types.GetSegmentPath(pbbCoreEvis.PbbCoreEvi[i]), types.YChild{"PbbCoreEvi", pbbCoreEvis.PbbCoreEvi[i]})
    }
    pbbCoreEvis.EntityData.Leafs = types.NewOrderedMap()

    pbbCoreEvis.EntityData.YListKeys = []string {}

    return &(pbbCoreEvis.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi
// PBB Core EVI
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..4294967295.
    Eviid interface{}
}

func (pbbCoreEvi *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi) GetEntityData() *types.CommonEntityData {
    pbbCoreEvi.EntityData.YFilter = pbbCoreEvi.YFilter
    pbbCoreEvi.EntityData.YangName = "pbb-core-evi"
    pbbCoreEvi.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreEvi.EntityData.ParentYangName = "pbb-core-evis"
    pbbCoreEvi.EntityData.SegmentPath = "pbb-core-evi" + types.AddKeyToken(pbbCoreEvi.Eviid, "eviid")
    pbbCoreEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreEvi.EntityData.Children = types.NewOrderedMap()
    pbbCoreEvi.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreEvi.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", pbbCoreEvi.Eviid})

    pbbCoreEvi.EntityData.YListKeys = []string {"Eviid"}

    return &(pbbCoreEvi.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile
// Attach a DHCP profile
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pbbCoreDhcpProfile *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile) GetEntityData() *types.CommonEntityData {
    pbbCoreDhcpProfile.EntityData.YFilter = pbbCoreDhcpProfile.YFilter
    pbbCoreDhcpProfile.EntityData.YangName = "pbb-core-dhcp-profile"
    pbbCoreDhcpProfile.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreDhcpProfile.EntityData.ParentYangName = "pbb-core"
    pbbCoreDhcpProfile.EntityData.SegmentPath = "pbb-core-dhcp-profile"
    pbbCoreDhcpProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreDhcpProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreDhcpProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreDhcpProfile.EntityData.Children = types.NewOrderedMap()
    pbbCoreDhcpProfile.EntityData.Leafs = types.NewOrderedMap()
    pbbCoreDhcpProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", pbbCoreDhcpProfile.ProfileId})
    pbbCoreDhcpProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", pbbCoreDhcpProfile.DhcpSnoopingId})

    pbbCoreDhcpProfile.EntityData.YListKeys = []string {}

    return &(pbbCoreDhcpProfile.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis
// Bridge Domain EVI Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain MPLS EVPN. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi.
    BridgeDomainEvi []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi
}

func (bridgeDomainEvis *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis) GetEntityData() *types.CommonEntityData {
    bridgeDomainEvis.EntityData.YFilter = bridgeDomainEvis.YFilter
    bridgeDomainEvis.EntityData.YangName = "bridge-domain-evis"
    bridgeDomainEvis.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainEvis.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainEvis.EntityData.SegmentPath = "bridge-domain-evis"
    bridgeDomainEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainEvis.EntityData.Children = types.NewOrderedMap()
    bridgeDomainEvis.EntityData.Children.Append("bridge-domain-evi", types.YChild{"BridgeDomainEvi", nil})
    for i := range bridgeDomainEvis.BridgeDomainEvi {
        bridgeDomainEvis.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainEvis.BridgeDomainEvi[i]), types.YChild{"BridgeDomainEvi", bridgeDomainEvis.BridgeDomainEvi[i]})
    }
    bridgeDomainEvis.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainEvis.EntityData.YListKeys = []string {}

    return &(bridgeDomainEvis.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi
// Bridge Domain MPLS EVPN
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MPLS Ethernet VPN-ID. The type is interface{} with
    // range: 1..65534.
    VpnId interface{}
}

func (bridgeDomainEvi *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi) GetEntityData() *types.CommonEntityData {
    bridgeDomainEvi.EntityData.YFilter = bridgeDomainEvi.YFilter
    bridgeDomainEvi.EntityData.YangName = "bridge-domain-evi"
    bridgeDomainEvi.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainEvi.EntityData.ParentYangName = "bridge-domain-evis"
    bridgeDomainEvi.EntityData.SegmentPath = "bridge-domain-evi" + types.AddKeyToken(bridgeDomainEvi.VpnId, "vpn-id")
    bridgeDomainEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainEvi.EntityData.Children = types.NewOrderedMap()
    bridgeDomainEvi.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainEvi.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", bridgeDomainEvi.VpnId})

    bridgeDomainEvi.EntityData.YListKeys = []string {"VpnId"}

    return &(bridgeDomainEvi.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis
// Specify the access virtual forwarding
// interface name
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Acess Virtual Forwarding Interface. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi.
    AccessVfi []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi
}

func (accessVfis *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis) GetEntityData() *types.CommonEntityData {
    accessVfis.EntityData.YFilter = accessVfis.YFilter
    accessVfis.EntityData.YangName = "access-vfis"
    accessVfis.EntityData.BundleName = "cisco_ios_xr"
    accessVfis.EntityData.ParentYangName = "bridge-domain"
    accessVfis.EntityData.SegmentPath = "access-vfis"
    accessVfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfis.EntityData.Children = types.NewOrderedMap()
    accessVfis.EntityData.Children.Append("access-vfi", types.YChild{"AccessVfi", nil})
    for i := range accessVfis.AccessVfi {
        accessVfis.EntityData.Children.Append(types.GetSegmentPath(accessVfis.AccessVfi[i]), types.YChild{"AccessVfi", accessVfis.AccessVfi[i]})
    }
    accessVfis.EntityData.Leafs = types.NewOrderedMap()

    accessVfis.EntityData.YListKeys = []string {}

    return &(accessVfis.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi
// Name of the Acess Virtual Forwarding
// Interface
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the AccessVirtual Forwarding Interface.
    // The type is string with length: 1..32.
    Name interface{}

    // shutdown the AccessVfi. The type is interface{}.
    AccessVfiShutdown interface{}

    // List of pseudowires.
    AccessVfiPseudowires L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires
}

func (accessVfi *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi) GetEntityData() *types.CommonEntityData {
    accessVfi.EntityData.YFilter = accessVfi.YFilter
    accessVfi.EntityData.YangName = "access-vfi"
    accessVfi.EntityData.BundleName = "cisco_ios_xr"
    accessVfi.EntityData.ParentYangName = "access-vfis"
    accessVfi.EntityData.SegmentPath = "access-vfi" + types.AddKeyToken(accessVfi.Name, "name")
    accessVfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfi.EntityData.Children = types.NewOrderedMap()
    accessVfi.EntityData.Children.Append("access-vfi-pseudowires", types.YChild{"AccessVfiPseudowires", &accessVfi.AccessVfiPseudowires})
    accessVfi.EntityData.Leafs = types.NewOrderedMap()
    accessVfi.EntityData.Leafs.Append("name", types.YLeaf{"Name", accessVfi.Name})
    accessVfi.EntityData.Leafs.Append("access-vfi-shutdown", types.YLeaf{"AccessVfiShutdown", accessVfi.AccessVfiShutdown})

    accessVfi.EntityData.YListKeys = []string {"Name"}

    return &(accessVfi.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires
// List of pseudowires
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire.
    AccessVfiPseudowire []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire
}

func (accessVfiPseudowires *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowires.EntityData.YFilter = accessVfiPseudowires.YFilter
    accessVfiPseudowires.EntityData.YangName = "access-vfi-pseudowires"
    accessVfiPseudowires.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowires.EntityData.ParentYangName = "access-vfi"
    accessVfiPseudowires.EntityData.SegmentPath = "access-vfi-pseudowires"
    accessVfiPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowires.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowires.EntityData.Children.Append("access-vfi-pseudowire", types.YChild{"AccessVfiPseudowire", nil})
    for i := range accessVfiPseudowires.AccessVfiPseudowire {
        accessVfiPseudowires.EntityData.Children.Append(types.GetSegmentPath(accessVfiPseudowires.AccessVfiPseudowire[i]), types.YChild{"AccessVfiPseudowire", accessVfiPseudowires.AccessVfiPseudowire[i]})
    }
    accessVfiPseudowires.EntityData.Leafs = types.NewOrderedMap()

    accessVfiPseudowires.EntityData.YListKeys = []string {}

    return &(accessVfiPseudowires.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire
// Pseudowire configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // Pseudowire class template name to use for this pseudowire. The type is
    // string with length: 1..32.
    AccessVfiPwClass interface{}

    // Static Mac Address Table.
    AccessVfiPseudowireStaticMacAddresses L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses
}

func (accessVfiPseudowire *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowire.EntityData.YFilter = accessVfiPseudowire.YFilter
    accessVfiPseudowire.EntityData.YangName = "access-vfi-pseudowire"
    accessVfiPseudowire.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowire.EntityData.ParentYangName = "access-vfi-pseudowires"
    accessVfiPseudowire.EntityData.SegmentPath = "access-vfi-pseudowire" + types.AddKeyToken(accessVfiPseudowire.Neighbor, "neighbor") + types.AddKeyToken(accessVfiPseudowire.PseudowireId, "pseudowire-id")
    accessVfiPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowire.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowire.EntityData.Children.Append("access-vfi-pseudowire-static-mac-addresses", types.YChild{"AccessVfiPseudowireStaticMacAddresses", &accessVfiPseudowire.AccessVfiPseudowireStaticMacAddresses})
    accessVfiPseudowire.EntityData.Leafs = types.NewOrderedMap()
    accessVfiPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", accessVfiPseudowire.Neighbor})
    accessVfiPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", accessVfiPseudowire.PseudowireId})
    accessVfiPseudowire.EntityData.Leafs.Append("access-vfi-pw-class", types.YLeaf{"AccessVfiPwClass", accessVfiPseudowire.AccessVfiPwClass})

    accessVfiPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(accessVfiPseudowire.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress.
    AccessVfiPseudowireStaticMacAddress []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress
}

func (accessVfiPseudowireStaticMacAddresses *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowireStaticMacAddresses.EntityData.YFilter = accessVfiPseudowireStaticMacAddresses.YFilter
    accessVfiPseudowireStaticMacAddresses.EntityData.YangName = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowireStaticMacAddresses.EntityData.ParentYangName = "access-vfi-pseudowire"
    accessVfiPseudowireStaticMacAddresses.EntityData.SegmentPath = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowireStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowireStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowireStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowireStaticMacAddresses.EntityData.Children.Append("access-vfi-pseudowire-static-mac-address", types.YChild{"AccessVfiPseudowireStaticMacAddress", nil})
    for i := range accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress {
        accessVfiPseudowireStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress[i]), types.YChild{"AccessVfiPseudowireStaticMacAddress", accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress[i]})
    }
    accessVfiPseudowireStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    accessVfiPseudowireStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(accessVfiPseudowireStaticMacAddresses.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (accessVfiPseudowireStaticMacAddress *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowireStaticMacAddress.EntityData.YFilter = accessVfiPseudowireStaticMacAddress.YFilter
    accessVfiPseudowireStaticMacAddress.EntityData.YangName = "access-vfi-pseudowire-static-mac-address"
    accessVfiPseudowireStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowireStaticMacAddress.EntityData.ParentYangName = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddress.EntityData.SegmentPath = "access-vfi-pseudowire-static-mac-address" + types.AddKeyToken(accessVfiPseudowireStaticMacAddress.Address, "address")
    accessVfiPseudowireStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowireStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowireStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowireStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    accessVfiPseudowireStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    accessVfiPseudowireStaticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", accessVfiPseudowireStaticMacAddress.Address})

    accessVfiPseudowireStaticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(accessVfiPseudowireStaticMacAddress.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires
// List of pseudowires
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire.
    BdPseudowire []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire
}

func (bdPseudowires *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires) GetEntityData() *types.CommonEntityData {
    bdPseudowires.EntityData.YFilter = bdPseudowires.YFilter
    bdPseudowires.EntityData.YangName = "bd-pseudowires"
    bdPseudowires.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowires.EntityData.ParentYangName = "bridge-domain"
    bdPseudowires.EntityData.SegmentPath = "bd-pseudowires"
    bdPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowires.EntityData.Children = types.NewOrderedMap()
    bdPseudowires.EntityData.Children.Append("bd-pseudowire", types.YChild{"BdPseudowire", nil})
    for i := range bdPseudowires.BdPseudowire {
        bdPseudowires.EntityData.Children.Append(types.GetSegmentPath(bdPseudowires.BdPseudowire[i]), types.YChild{"BdPseudowire", bdPseudowires.BdPseudowire[i]})
    }
    bdPseudowires.EntityData.Leafs = types.NewOrderedMap()

    bdPseudowires.EntityData.YListKeys = []string {}

    return &(bdPseudowires.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire
// Pseudowire configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // Attach a MLD Snooping profile. The type is string with length: 1..32.
    PseudowireMldSnoop interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    PseudowireIgmpSnoop interface{}

    // Bridge-domain Pseudowire flooding. The type is InterfaceTrafficFlood.
    PseudowireFlooding interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    BdPwClass interface{}

    // Bridge-domain Pseudowire flooding Unknown Unicast. The type is
    // InterfaceTrafficFlood.
    PseudowireFloodingUnknownUnicast interface{}

    // Access Pseudowire Dynamic ARP Inspection.
    PseudowireDai L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai

    // Storm Control.
    BdpwStormControlTypes L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes

    // Attach a DHCP profile.
    PseudowireProfile L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile

    // Static Mac Address Table.
    BdPwStaticMacAddresses L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses

    // IP Source Guard.
    PseudowireIpSourceGuard L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard

    // Bridge-domain Pseudowire MAC configuration commands.
    PseudowireMac L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac

    // Split Horizon.
    BdPwSplitHorizon L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon

    // MPLS static labels.
    BdPwMplsStaticLabels L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels

    // List of pseudowires.
    BridgeDomainBackupPseudowires L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires
}

func (bdPseudowire *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire) GetEntityData() *types.CommonEntityData {
    bdPseudowire.EntityData.YFilter = bdPseudowire.YFilter
    bdPseudowire.EntityData.YangName = "bd-pseudowire"
    bdPseudowire.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowire.EntityData.ParentYangName = "bd-pseudowires"
    bdPseudowire.EntityData.SegmentPath = "bd-pseudowire" + types.AddKeyToken(bdPseudowire.Neighbor, "neighbor") + types.AddKeyToken(bdPseudowire.PseudowireId, "pseudowire-id")
    bdPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowire.EntityData.Children = types.NewOrderedMap()
    bdPseudowire.EntityData.Children.Append("pseudowire-dai", types.YChild{"PseudowireDai", &bdPseudowire.PseudowireDai})
    bdPseudowire.EntityData.Children.Append("bdpw-storm-control-types", types.YChild{"BdpwStormControlTypes", &bdPseudowire.BdpwStormControlTypes})
    bdPseudowire.EntityData.Children.Append("pseudowire-profile", types.YChild{"PseudowireProfile", &bdPseudowire.PseudowireProfile})
    bdPseudowire.EntityData.Children.Append("bd-pw-static-mac-addresses", types.YChild{"BdPwStaticMacAddresses", &bdPseudowire.BdPwStaticMacAddresses})
    bdPseudowire.EntityData.Children.Append("pseudowire-ip-source-guard", types.YChild{"PseudowireIpSourceGuard", &bdPseudowire.PseudowireIpSourceGuard})
    bdPseudowire.EntityData.Children.Append("pseudowire-mac", types.YChild{"PseudowireMac", &bdPseudowire.PseudowireMac})
    bdPseudowire.EntityData.Children.Append("bd-pw-split-horizon", types.YChild{"BdPwSplitHorizon", &bdPseudowire.BdPwSplitHorizon})
    bdPseudowire.EntityData.Children.Append("bd-pw-mpls-static-labels", types.YChild{"BdPwMplsStaticLabels", &bdPseudowire.BdPwMplsStaticLabels})
    bdPseudowire.EntityData.Children.Append("bridge-domain-backup-pseudowires", types.YChild{"BridgeDomainBackupPseudowires", &bdPseudowire.BridgeDomainBackupPseudowires})
    bdPseudowire.EntityData.Leafs = types.NewOrderedMap()
    bdPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", bdPseudowire.Neighbor})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", bdPseudowire.PseudowireId})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-mld-snoop", types.YLeaf{"PseudowireMldSnoop", bdPseudowire.PseudowireMldSnoop})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-igmp-snoop", types.YLeaf{"PseudowireIgmpSnoop", bdPseudowire.PseudowireIgmpSnoop})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-flooding", types.YLeaf{"PseudowireFlooding", bdPseudowire.PseudowireFlooding})
    bdPseudowire.EntityData.Leafs.Append("bd-pw-class", types.YLeaf{"BdPwClass", bdPseudowire.BdPwClass})
    bdPseudowire.EntityData.Leafs.Append("pseudowire-flooding-unknown-unicast", types.YLeaf{"PseudowireFloodingUnknownUnicast", bdPseudowire.PseudowireFloodingUnknownUnicast})

    bdPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(bdPseudowire.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai
// Access Pseudowire Dynamic ARP Inspection
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable Dynamic ARP Inspection. The type is interface{}.
    Disable interface{}

    // Enable Access Pseudowire Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    PseudowireDaiAddressValidation L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation
}

func (pseudowireDai *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai) GetEntityData() *types.CommonEntityData {
    pseudowireDai.EntityData.YFilter = pseudowireDai.YFilter
    pseudowireDai.EntityData.YangName = "pseudowire-dai"
    pseudowireDai.EntityData.BundleName = "cisco_ios_xr"
    pseudowireDai.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireDai.EntityData.SegmentPath = "pseudowire-dai"
    pseudowireDai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireDai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireDai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireDai.EntityData.Children = types.NewOrderedMap()
    pseudowireDai.EntityData.Children.Append("pseudowire-dai-address-validation", types.YChild{"PseudowireDaiAddressValidation", &pseudowireDai.PseudowireDaiAddressValidation})
    pseudowireDai.EntityData.Leafs = types.NewOrderedMap()
    pseudowireDai.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pseudowireDai.Logging})
    pseudowireDai.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pseudowireDai.Disable})
    pseudowireDai.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireDai.Enable})

    pseudowireDai.EntityData.YListKeys = []string {}

    return &(pseudowireDai.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation
// Address Validation
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Verification. The type is L2vpnVerification.
    Ipv4Verification interface{}

    // Destination MAC Verification. The type is L2vpnVerification.
    DestinationMacVerification interface{}

    // Source MAC Verification. The type is L2vpnVerification.
    SourceMacVerification interface{}
}

func (pseudowireDaiAddressValidation *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation) GetEntityData() *types.CommonEntityData {
    pseudowireDaiAddressValidation.EntityData.YFilter = pseudowireDaiAddressValidation.YFilter
    pseudowireDaiAddressValidation.EntityData.YangName = "pseudowire-dai-address-validation"
    pseudowireDaiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    pseudowireDaiAddressValidation.EntityData.ParentYangName = "pseudowire-dai"
    pseudowireDaiAddressValidation.EntityData.SegmentPath = "pseudowire-dai-address-validation"
    pseudowireDaiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireDaiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireDaiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireDaiAddressValidation.EntityData.Children = types.NewOrderedMap()
    pseudowireDaiAddressValidation.EntityData.Leafs = types.NewOrderedMap()
    pseudowireDaiAddressValidation.EntityData.Leafs.Append("ipv4-verification", types.YLeaf{"Ipv4Verification", pseudowireDaiAddressValidation.Ipv4Verification})
    pseudowireDaiAddressValidation.EntityData.Leafs.Append("destination-mac-verification", types.YLeaf{"DestinationMacVerification", pseudowireDaiAddressValidation.DestinationMacVerification})
    pseudowireDaiAddressValidation.EntityData.Leafs.Append("source-mac-verification", types.YLeaf{"SourceMacVerification", pseudowireDaiAddressValidation.SourceMacVerification})

    pseudowireDaiAddressValidation.EntityData.YListKeys = []string {}

    return &(pseudowireDaiAddressValidation.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes
// Storm Control
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType.
    BdpwStormControlType []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType
}

func (bdpwStormControlTypes *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes) GetEntityData() *types.CommonEntityData {
    bdpwStormControlTypes.EntityData.YFilter = bdpwStormControlTypes.YFilter
    bdpwStormControlTypes.EntityData.YangName = "bdpw-storm-control-types"
    bdpwStormControlTypes.EntityData.BundleName = "cisco_ios_xr"
    bdpwStormControlTypes.EntityData.ParentYangName = "bd-pseudowire"
    bdpwStormControlTypes.EntityData.SegmentPath = "bdpw-storm-control-types"
    bdpwStormControlTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdpwStormControlTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdpwStormControlTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdpwStormControlTypes.EntityData.Children = types.NewOrderedMap()
    bdpwStormControlTypes.EntityData.Children.Append("bdpw-storm-control-type", types.YChild{"BdpwStormControlType", nil})
    for i := range bdpwStormControlTypes.BdpwStormControlType {
        bdpwStormControlTypes.EntityData.Children.Append(types.GetSegmentPath(bdpwStormControlTypes.BdpwStormControlType[i]), types.YChild{"BdpwStormControlType", bdpwStormControlTypes.BdpwStormControlType[i]})
    }
    bdpwStormControlTypes.EntityData.Leafs = types.NewOrderedMap()

    bdpwStormControlTypes.EntityData.YListKeys = []string {}

    return &(bdpwStormControlTypes.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType
// Storm Control Type
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit
}

func (bdpwStormControlType *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType) GetEntityData() *types.CommonEntityData {
    bdpwStormControlType.EntityData.YFilter = bdpwStormControlType.YFilter
    bdpwStormControlType.EntityData.YangName = "bdpw-storm-control-type"
    bdpwStormControlType.EntityData.BundleName = "cisco_ios_xr"
    bdpwStormControlType.EntityData.ParentYangName = "bdpw-storm-control-types"
    bdpwStormControlType.EntityData.SegmentPath = "bdpw-storm-control-type" + types.AddKeyToken(bdpwStormControlType.Sctype, "sctype")
    bdpwStormControlType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdpwStormControlType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdpwStormControlType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdpwStormControlType.EntityData.Children = types.NewOrderedMap()
    bdpwStormControlType.EntityData.Children.Append("storm-control-unit", types.YChild{"StormControlUnit", &bdpwStormControlType.StormControlUnit})
    bdpwStormControlType.EntityData.Leafs = types.NewOrderedMap()
    bdpwStormControlType.EntityData.Leafs.Append("sctype", types.YLeaf{"Sctype", bdpwStormControlType.Sctype})

    bdpwStormControlType.EntityData.YListKeys = []string {"Sctype"}

    return &(bdpwStormControlType.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit
// Specify units for Storm Control Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Kilobits Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 64..1280000. Units are
    // kbit/s.
    KbitsPerSec interface{}

    // Packets Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 1..160000. Units are
    // packet/s.
    PktsPerSec interface{}
}

func (stormControlUnit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bdpw-storm-control-type"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs.Append("kbits-per-sec", types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec})
    stormControlUnit.EntityData.Leafs.Append("pkts-per-sec", types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec})

    stormControlUnit.EntityData.YListKeys = []string {}

    return &(stormControlUnit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile
// Attach a DHCP profile
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pseudowireProfile *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile) GetEntityData() *types.CommonEntityData {
    pseudowireProfile.EntityData.YFilter = pseudowireProfile.YFilter
    pseudowireProfile.EntityData.YangName = "pseudowire-profile"
    pseudowireProfile.EntityData.BundleName = "cisco_ios_xr"
    pseudowireProfile.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireProfile.EntityData.SegmentPath = "pseudowire-profile"
    pseudowireProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireProfile.EntityData.Children = types.NewOrderedMap()
    pseudowireProfile.EntityData.Leafs = types.NewOrderedMap()
    pseudowireProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", pseudowireProfile.ProfileId})
    pseudowireProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", pseudowireProfile.DhcpSnoopingId})

    pseudowireProfile.EntityData.YListKeys = []string {}

    return &(pseudowireProfile.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress.
    BdPwStaticMacAddress []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress
}

func (bdPwStaticMacAddresses *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    bdPwStaticMacAddresses.EntityData.YFilter = bdPwStaticMacAddresses.YFilter
    bdPwStaticMacAddresses.EntityData.YangName = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    bdPwStaticMacAddresses.EntityData.ParentYangName = "bd-pseudowire"
    bdPwStaticMacAddresses.EntityData.SegmentPath = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    bdPwStaticMacAddresses.EntityData.Children.Append("bd-pw-static-mac-address", types.YChild{"BdPwStaticMacAddress", nil})
    for i := range bdPwStaticMacAddresses.BdPwStaticMacAddress {
        bdPwStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(bdPwStaticMacAddresses.BdPwStaticMacAddress[i]), types.YChild{"BdPwStaticMacAddress", bdPwStaticMacAddresses.BdPwStaticMacAddress[i]})
    }
    bdPwStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    bdPwStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(bdPwStaticMacAddresses.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (bdPwStaticMacAddress *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress) GetEntityData() *types.CommonEntityData {
    bdPwStaticMacAddress.EntityData.YFilter = bdPwStaticMacAddress.YFilter
    bdPwStaticMacAddress.EntityData.YangName = "bd-pw-static-mac-address"
    bdPwStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    bdPwStaticMacAddress.EntityData.ParentYangName = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddress.EntityData.SegmentPath = "bd-pw-static-mac-address" + types.AddKeyToken(bdPwStaticMacAddress.Address, "address")
    bdPwStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    bdPwStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    bdPwStaticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", bdPwStaticMacAddress.Address})

    bdPwStaticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(bdPwStaticMacAddress.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard
// IP Source Guard
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable Dynamic IP source guard. The type is interface{}.
    Disable interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (pseudowireIpSourceGuard *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard) GetEntityData() *types.CommonEntityData {
    pseudowireIpSourceGuard.EntityData.YFilter = pseudowireIpSourceGuard.YFilter
    pseudowireIpSourceGuard.EntityData.YangName = "pseudowire-ip-source-guard"
    pseudowireIpSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    pseudowireIpSourceGuard.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireIpSourceGuard.EntityData.SegmentPath = "pseudowire-ip-source-guard"
    pseudowireIpSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireIpSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireIpSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireIpSourceGuard.EntityData.Children = types.NewOrderedMap()
    pseudowireIpSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    pseudowireIpSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pseudowireIpSourceGuard.Logging})
    pseudowireIpSourceGuard.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pseudowireIpSourceGuard.Disable})
    pseudowireIpSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireIpSourceGuard.Enable})

    pseudowireIpSourceGuard.EntityData.YListKeys = []string {}

    return &(pseudowireIpSourceGuard.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac
// Bridge-domain Pseudowire MAC configuration
// commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MAC Flush When Port goes down. The type is PortDownFlush.
    PseudowireMacPortDownFlush interface{}

    // Bridge-domain Pseudowire MAC configuration mode. The type is interface{}.
    Enable interface{}

    // Enable MAC Learning. The type is MacLearn.
    PseudowireMacLearning interface{}

    // MAC Secure.
    PseudowireMacSecure L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure

    // MAC-Aging configuration commands.
    PseudowireMacAging L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging

    // MAC-Limit configuration commands.
    PseudowireMacLimit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit
}

func (pseudowireMac *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac) GetEntityData() *types.CommonEntityData {
    pseudowireMac.EntityData.YFilter = pseudowireMac.YFilter
    pseudowireMac.EntityData.YangName = "pseudowire-mac"
    pseudowireMac.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMac.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireMac.EntityData.SegmentPath = "pseudowire-mac"
    pseudowireMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMac.EntityData.Children = types.NewOrderedMap()
    pseudowireMac.EntityData.Children.Append("pseudowire-mac-secure", types.YChild{"PseudowireMacSecure", &pseudowireMac.PseudowireMacSecure})
    pseudowireMac.EntityData.Children.Append("pseudowire-mac-aging", types.YChild{"PseudowireMacAging", &pseudowireMac.PseudowireMacAging})
    pseudowireMac.EntityData.Children.Append("pseudowire-mac-limit", types.YChild{"PseudowireMacLimit", &pseudowireMac.PseudowireMacLimit})
    pseudowireMac.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMac.EntityData.Leafs.Append("pseudowire-mac-port-down-flush", types.YLeaf{"PseudowireMacPortDownFlush", pseudowireMac.PseudowireMacPortDownFlush})
    pseudowireMac.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireMac.Enable})
    pseudowireMac.EntityData.Leafs.Append("pseudowire-mac-learning", types.YLeaf{"PseudowireMacLearning", pseudowireMac.PseudowireMacLearning})

    pseudowireMac.EntityData.YListKeys = []string {}

    return &(pseudowireMac.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure
// MAC Secure
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Pseudowire MAC Secure. The type is interface{}.
    Disable interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}
}

func (pseudowireMacSecure *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure) GetEntityData() *types.CommonEntityData {
    pseudowireMacSecure.EntityData.YFilter = pseudowireMacSecure.YFilter
    pseudowireMacSecure.EntityData.YangName = "pseudowire-mac-secure"
    pseudowireMacSecure.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacSecure.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacSecure.EntityData.SegmentPath = "pseudowire-mac-secure"
    pseudowireMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacSecure.EntityData.Children = types.NewOrderedMap()
    pseudowireMacSecure.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMacSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", pseudowireMacSecure.Logging})
    pseudowireMacSecure.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", pseudowireMacSecure.Disable})
    pseudowireMacSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", pseudowireMacSecure.Action})
    pseudowireMacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pseudowireMacSecure.Enable})

    pseudowireMacSecure.EntityData.YListKeys = []string {}

    return &(pseudowireMacSecure.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging
// MAC-Aging configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PseudowireMacAgingType interface{}

    // MAC Aging Time. The type is interface{} with range: 300..30000.
    PseudowireMacAgingTime interface{}
}

func (pseudowireMacAging *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging) GetEntityData() *types.CommonEntityData {
    pseudowireMacAging.EntityData.YFilter = pseudowireMacAging.YFilter
    pseudowireMacAging.EntityData.YangName = "pseudowire-mac-aging"
    pseudowireMacAging.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacAging.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacAging.EntityData.SegmentPath = "pseudowire-mac-aging"
    pseudowireMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacAging.EntityData.Children = types.NewOrderedMap()
    pseudowireMacAging.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMacAging.EntityData.Leafs.Append("pseudowire-mac-aging-type", types.YLeaf{"PseudowireMacAgingType", pseudowireMacAging.PseudowireMacAgingType})
    pseudowireMacAging.EntityData.Leafs.Append("pseudowire-mac-aging-time", types.YLeaf{"PseudowireMacAgingTime", pseudowireMacAging.PseudowireMacAgingTime})

    pseudowireMacAging.EntityData.YListKeys = []string {}

    return &(pseudowireMacAging.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Access Pseudowire MAC address limit enforcement action. The type is
    // MacLimitAction.
    PseudowireMacLimitAction interface{}

    // MAC address limit notification action in a Bridge Access Pseudowire. The
    // type is MacNotification.
    PseudowireMacLimitNotif interface{}

    // Number of MAC addresses on a Bridge Access Pseudowire after which MAC limit
    // action is taken. The type is interface{} with range: 0..4294967295.
    PseudowireMacLimitMax interface{}
}

func (pseudowireMacLimit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit) GetEntityData() *types.CommonEntityData {
    pseudowireMacLimit.EntityData.YFilter = pseudowireMacLimit.YFilter
    pseudowireMacLimit.EntityData.YangName = "pseudowire-mac-limit"
    pseudowireMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacLimit.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacLimit.EntityData.SegmentPath = "pseudowire-mac-limit"
    pseudowireMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacLimit.EntityData.Children = types.NewOrderedMap()
    pseudowireMacLimit.EntityData.Leafs = types.NewOrderedMap()
    pseudowireMacLimit.EntityData.Leafs.Append("pseudowire-mac-limit-action", types.YLeaf{"PseudowireMacLimitAction", pseudowireMacLimit.PseudowireMacLimitAction})
    pseudowireMacLimit.EntityData.Leafs.Append("pseudowire-mac-limit-notif", types.YLeaf{"PseudowireMacLimitNotif", pseudowireMacLimit.PseudowireMacLimitNotif})
    pseudowireMacLimit.EntityData.Leafs.Append("pseudowire-mac-limit-max", types.YLeaf{"PseudowireMacLimitMax", pseudowireMacLimit.PseudowireMacLimitMax})

    pseudowireMacLimit.EntityData.YListKeys = []string {}

    return &(pseudowireMacLimit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon
// Split Horizon
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon Group.
    BdPwSplitHorizonGroup L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup
}

func (bdPwSplitHorizon *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon) GetEntityData() *types.CommonEntityData {
    bdPwSplitHorizon.EntityData.YFilter = bdPwSplitHorizon.YFilter
    bdPwSplitHorizon.EntityData.YangName = "bd-pw-split-horizon"
    bdPwSplitHorizon.EntityData.BundleName = "cisco_ios_xr"
    bdPwSplitHorizon.EntityData.ParentYangName = "bd-pseudowire"
    bdPwSplitHorizon.EntityData.SegmentPath = "bd-pw-split-horizon"
    bdPwSplitHorizon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwSplitHorizon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwSplitHorizon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwSplitHorizon.EntityData.Children = types.NewOrderedMap()
    bdPwSplitHorizon.EntityData.Children.Append("bd-pw-split-horizon-group", types.YChild{"BdPwSplitHorizonGroup", &bdPwSplitHorizon.BdPwSplitHorizonGroup})
    bdPwSplitHorizon.EntityData.Leafs = types.NewOrderedMap()

    bdPwSplitHorizon.EntityData.YListKeys = []string {}

    return &(bdPwSplitHorizon.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup
// Split Horizon Group
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable split horizon group. The type is interface{}.
    Enable interface{}
}

func (bdPwSplitHorizonGroup *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    bdPwSplitHorizonGroup.EntityData.YFilter = bdPwSplitHorizonGroup.YFilter
    bdPwSplitHorizonGroup.EntityData.YangName = "bd-pw-split-horizon-group"
    bdPwSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    bdPwSplitHorizonGroup.EntityData.ParentYangName = "bd-pw-split-horizon"
    bdPwSplitHorizonGroup.EntityData.SegmentPath = "bd-pw-split-horizon-group"
    bdPwSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwSplitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    bdPwSplitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    bdPwSplitHorizonGroup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bdPwSplitHorizonGroup.Enable})

    bdPwSplitHorizonGroup.EntityData.YListKeys = []string {}

    return &(bdPwSplitHorizonGroup.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels
// MPLS static labels
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (bdPwMplsStaticLabels *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    bdPwMplsStaticLabels.EntityData.YFilter = bdPwMplsStaticLabels.YFilter
    bdPwMplsStaticLabels.EntityData.YangName = "bd-pw-mpls-static-labels"
    bdPwMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    bdPwMplsStaticLabels.EntityData.ParentYangName = "bd-pseudowire"
    bdPwMplsStaticLabels.EntityData.SegmentPath = "bd-pw-mpls-static-labels"
    bdPwMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwMplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    bdPwMplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    bdPwMplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", bdPwMplsStaticLabels.LocalStaticLabel})
    bdPwMplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", bdPwMplsStaticLabels.RemoteStaticLabel})

    bdPwMplsStaticLabels.EntityData.YListKeys = []string {}

    return &(bdPwMplsStaticLabels.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires
// List of pseudowires
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire.
    BridgeDomainBackupPseudowire []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire
}

func (bridgeDomainBackupPseudowires *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires) GetEntityData() *types.CommonEntityData {
    bridgeDomainBackupPseudowires.EntityData.YFilter = bridgeDomainBackupPseudowires.YFilter
    bridgeDomainBackupPseudowires.EntityData.YangName = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainBackupPseudowires.EntityData.ParentYangName = "bd-pseudowire"
    bridgeDomainBackupPseudowires.EntityData.SegmentPath = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainBackupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainBackupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainBackupPseudowires.EntityData.Children = types.NewOrderedMap()
    bridgeDomainBackupPseudowires.EntityData.Children.Append("bridge-domain-backup-pseudowire", types.YChild{"BridgeDomainBackupPseudowire", nil})
    for i := range bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire {
        bridgeDomainBackupPseudowires.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire[i]), types.YChild{"BridgeDomainBackupPseudowire", bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire[i]})
    }
    bridgeDomainBackupPseudowires.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainBackupPseudowires.EntityData.YListKeys = []string {}

    return &(bridgeDomainBackupPseudowires.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire
// Backup pseudowire configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    BridgeDomainBackupPwClass interface{}
}

func (bridgeDomainBackupPseudowire *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire) GetEntityData() *types.CommonEntityData {
    bridgeDomainBackupPseudowire.EntityData.YFilter = bridgeDomainBackupPseudowire.YFilter
    bridgeDomainBackupPseudowire.EntityData.YangName = "bridge-domain-backup-pseudowire"
    bridgeDomainBackupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainBackupPseudowire.EntityData.ParentYangName = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowire.EntityData.SegmentPath = "bridge-domain-backup-pseudowire" + types.AddKeyToken(bridgeDomainBackupPseudowire.Neighbor, "neighbor") + types.AddKeyToken(bridgeDomainBackupPseudowire.PseudowireId, "pseudowire-id")
    bridgeDomainBackupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainBackupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainBackupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainBackupPseudowire.EntityData.Children = types.NewOrderedMap()
    bridgeDomainBackupPseudowire.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainBackupPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", bridgeDomainBackupPseudowire.Neighbor})
    bridgeDomainBackupPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", bridgeDomainBackupPseudowire.PseudowireId})
    bridgeDomainBackupPseudowire.EntityData.Leafs.Append("bridge-domain-backup-pw-class", types.YLeaf{"BridgeDomainBackupPwClass", bridgeDomainBackupPseudowire.BridgeDomainBackupPwClass})

    bridgeDomainBackupPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(bridgeDomainBackupPseudowire.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis
// Specify the virtual forwarding interface name
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Virtual Forwarding Interface. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi.
    Vfi []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi
}

func (vfis *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis) GetEntityData() *types.CommonEntityData {
    vfis.EntityData.YFilter = vfis.YFilter
    vfis.EntityData.YangName = "vfis"
    vfis.EntityData.BundleName = "cisco_ios_xr"
    vfis.EntityData.ParentYangName = "bridge-domain"
    vfis.EntityData.SegmentPath = "vfis"
    vfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfis.EntityData.Children = types.NewOrderedMap()
    vfis.EntityData.Children.Append("vfi", types.YChild{"Vfi", nil})
    for i := range vfis.Vfi {
        vfis.EntityData.Children.Append(types.GetSegmentPath(vfis.Vfi[i]), types.YChild{"Vfi", vfis.Vfi[i]})
    }
    vfis.EntityData.Leafs = types.NewOrderedMap()

    vfis.EntityData.YListKeys = []string {}

    return &(vfis.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi
// Name of the Virtual Forwarding Interface
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Virtual Forwarding Interface. The type
    // is string with length: 1..32.
    Name interface{}

    // Enabling Shutdown. The type is interface{}.
    VfiShutdown interface{}

    // VPN Identifier. The type is interface{} with range: 1..4294967295.
    Vpnid interface{}

    // Enable Multicast P2MP in this VFI.
    MulticastP2mp L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp

    // List of pseudowires.
    VfiPseudowires L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires

    // Enable Autodiscovery BGP in this VFI.
    BgpAutoDiscovery L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery
}

func (vfi *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi) GetEntityData() *types.CommonEntityData {
    vfi.EntityData.YFilter = vfi.YFilter
    vfi.EntityData.YangName = "vfi"
    vfi.EntityData.BundleName = "cisco_ios_xr"
    vfi.EntityData.ParentYangName = "vfis"
    vfi.EntityData.SegmentPath = "vfi" + types.AddKeyToken(vfi.Name, "name")
    vfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfi.EntityData.Children = types.NewOrderedMap()
    vfi.EntityData.Children.Append("multicast-p2mp", types.YChild{"MulticastP2mp", &vfi.MulticastP2mp})
    vfi.EntityData.Children.Append("vfi-pseudowires", types.YChild{"VfiPseudowires", &vfi.VfiPseudowires})
    vfi.EntityData.Children.Append("bgp-auto-discovery", types.YChild{"BgpAutoDiscovery", &vfi.BgpAutoDiscovery})
    vfi.EntityData.Leafs = types.NewOrderedMap()
    vfi.EntityData.Leafs.Append("name", types.YLeaf{"Name", vfi.Name})
    vfi.EntityData.Leafs.Append("vfi-shutdown", types.YLeaf{"VfiShutdown", vfi.VfiShutdown})
    vfi.EntityData.Leafs.Append("vpnid", types.YLeaf{"Vpnid", vfi.Vpnid})

    vfi.EntityData.YListKeys = []string {"Name"}

    return &(vfi.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp
// Enable Multicast P2MP in this VFI
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery P2MP. The type is interface{}.
    Enable interface{}

    // Multicast P2MP Transport.
    Transports L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports

    // Multicast P2MP Signaling Type.
    Signalings L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings
}

func (multicastP2mp *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp) GetEntityData() *types.CommonEntityData {
    multicastP2mp.EntityData.YFilter = multicastP2mp.YFilter
    multicastP2mp.EntityData.YangName = "multicast-p2mp"
    multicastP2mp.EntityData.BundleName = "cisco_ios_xr"
    multicastP2mp.EntityData.ParentYangName = "vfi"
    multicastP2mp.EntityData.SegmentPath = "multicast-p2mp"
    multicastP2mp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastP2mp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastP2mp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastP2mp.EntityData.Children = types.NewOrderedMap()
    multicastP2mp.EntityData.Children.Append("transports", types.YChild{"Transports", &multicastP2mp.Transports})
    multicastP2mp.EntityData.Children.Append("signalings", types.YChild{"Signalings", &multicastP2mp.Signalings})
    multicastP2mp.EntityData.Leafs = types.NewOrderedMap()
    multicastP2mp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", multicastP2mp.Enable})

    multicastP2mp.EntityData.YListKeys = []string {}

    return &(multicastP2mp.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports
// Multicast P2MP Transport
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multicast P2MP Transport Type. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport.
    Transport []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport
}

func (transports *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports) GetEntityData() *types.CommonEntityData {
    transports.EntityData.YFilter = transports.YFilter
    transports.EntityData.YangName = "transports"
    transports.EntityData.BundleName = "cisco_ios_xr"
    transports.EntityData.ParentYangName = "multicast-p2mp"
    transports.EntityData.SegmentPath = "transports"
    transports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transports.EntityData.Children = types.NewOrderedMap()
    transports.EntityData.Children.Append("transport", types.YChild{"Transport", nil})
    for i := range transports.Transport {
        transports.EntityData.Children.Append(types.GetSegmentPath(transports.Transport[i]), types.YChild{"Transport", transports.Transport[i]})
    }
    transports.EntityData.Leafs = types.NewOrderedMap()

    transports.EntityData.YListKeys = []string {}

    return &(transports.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport
// Multicast P2MP Transport Type
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Transport Type. The type is string with pattern:
    // (RSVP_TE).
    TransportName interface{}

    // Multicast P2MP TE Attribute Set Name. The type is string with length:
    // 1..64.
    AttributeSetName interface{}
}

func (transport *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Transports_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xr"
    transport.EntityData.ParentYangName = "transports"
    transport.EntityData.SegmentPath = "transport" + types.AddKeyToken(transport.TransportName, "transport-name")
    transport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Leafs = types.NewOrderedMap()
    transport.EntityData.Leafs.Append("transport-name", types.YLeaf{"TransportName", transport.TransportName})
    transport.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", transport.AttributeSetName})

    transport.EntityData.YListKeys = []string {"TransportName"}

    return &(transport.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings
// Multicast P2MP Signaling Type
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multicast P2MP Signaling Type. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling.
    Signaling []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling
}

func (signalings *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings) GetEntityData() *types.CommonEntityData {
    signalings.EntityData.YFilter = signalings.YFilter
    signalings.EntityData.YangName = "signalings"
    signalings.EntityData.BundleName = "cisco_ios_xr"
    signalings.EntityData.ParentYangName = "multicast-p2mp"
    signalings.EntityData.SegmentPath = "signalings"
    signalings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalings.EntityData.Children = types.NewOrderedMap()
    signalings.EntityData.Children.Append("signaling", types.YChild{"Signaling", nil})
    for i := range signalings.Signaling {
        signalings.EntityData.Children.Append(types.GetSegmentPath(signalings.Signaling[i]), types.YChild{"Signaling", signalings.Signaling[i]})
    }
    signalings.EntityData.Leafs = types.NewOrderedMap()

    signalings.EntityData.YListKeys = []string {}

    return &(signalings.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling
// Multicast P2MP Signaling Type
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Signaling Type. The type is string with pattern:
    // (BGP).
    SignalingName interface{}
}

func (signaling *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2mp_Signalings_Signaling) GetEntityData() *types.CommonEntityData {
    signaling.EntityData.YFilter = signaling.YFilter
    signaling.EntityData.YangName = "signaling"
    signaling.EntityData.BundleName = "cisco_ios_xr"
    signaling.EntityData.ParentYangName = "signalings"
    signaling.EntityData.SegmentPath = "signaling" + types.AddKeyToken(signaling.SignalingName, "signaling-name")
    signaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaling.EntityData.Children = types.NewOrderedMap()
    signaling.EntityData.Leafs = types.NewOrderedMap()
    signaling.EntityData.Leafs.Append("signaling-name", types.YLeaf{"SignalingName", signaling.SignalingName})

    signaling.EntityData.YListKeys = []string {"SignalingName"}

    return &(signaling.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires
// List of pseudowires
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire.
    VfiPseudowire []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire
}

func (vfiPseudowires *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires) GetEntityData() *types.CommonEntityData {
    vfiPseudowires.EntityData.YFilter = vfiPseudowires.YFilter
    vfiPseudowires.EntityData.YangName = "vfi-pseudowires"
    vfiPseudowires.EntityData.BundleName = "cisco_ios_xr"
    vfiPseudowires.EntityData.ParentYangName = "vfi"
    vfiPseudowires.EntityData.SegmentPath = "vfi-pseudowires"
    vfiPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPseudowires.EntityData.Children = types.NewOrderedMap()
    vfiPseudowires.EntityData.Children.Append("vfi-pseudowire", types.YChild{"VfiPseudowire", nil})
    for i := range vfiPseudowires.VfiPseudowire {
        vfiPseudowires.EntityData.Children.Append(types.GetSegmentPath(vfiPseudowires.VfiPseudowire[i]), types.YChild{"VfiPseudowire", vfiPseudowires.VfiPseudowire[i]})
    }
    vfiPseudowires.EntityData.Leafs = types.NewOrderedMap()

    vfiPseudowires.EntityData.YListKeys = []string {}

    return &(vfiPseudowires.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire
// Pseudowire configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    VfiPwClass interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    VfiPwIgmpSnoop interface{}

    // Attach a MLD Snooping profile. The type is string with length: 1..32.
    VfiPwMldSnoop interface{}

    // Attach a DHCP Snooping profile.
    VfiPwDhcpSnoop L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop

    // MPLS static labels.
    VfiPwMplsStaticLabels L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels

    // Static Mac Address Table.
    PseudowireStaticMacAddresses L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses
}

func (vfiPseudowire *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire) GetEntityData() *types.CommonEntityData {
    vfiPseudowire.EntityData.YFilter = vfiPseudowire.YFilter
    vfiPseudowire.EntityData.YangName = "vfi-pseudowire"
    vfiPseudowire.EntityData.BundleName = "cisco_ios_xr"
    vfiPseudowire.EntityData.ParentYangName = "vfi-pseudowires"
    vfiPseudowire.EntityData.SegmentPath = "vfi-pseudowire" + types.AddKeyToken(vfiPseudowire.Neighbor, "neighbor") + types.AddKeyToken(vfiPseudowire.PseudowireId, "pseudowire-id")
    vfiPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPseudowire.EntityData.Children = types.NewOrderedMap()
    vfiPseudowire.EntityData.Children.Append("vfi-pw-dhcp-snoop", types.YChild{"VfiPwDhcpSnoop", &vfiPseudowire.VfiPwDhcpSnoop})
    vfiPseudowire.EntityData.Children.Append("vfi-pw-mpls-static-labels", types.YChild{"VfiPwMplsStaticLabels", &vfiPseudowire.VfiPwMplsStaticLabels})
    vfiPseudowire.EntityData.Children.Append("pseudowire-static-mac-addresses", types.YChild{"PseudowireStaticMacAddresses", &vfiPseudowire.PseudowireStaticMacAddresses})
    vfiPseudowire.EntityData.Leafs = types.NewOrderedMap()
    vfiPseudowire.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", vfiPseudowire.Neighbor})
    vfiPseudowire.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", vfiPseudowire.PseudowireId})
    vfiPseudowire.EntityData.Leafs.Append("vfi-pw-class", types.YLeaf{"VfiPwClass", vfiPseudowire.VfiPwClass})
    vfiPseudowire.EntityData.Leafs.Append("vfi-pw-igmp-snoop", types.YLeaf{"VfiPwIgmpSnoop", vfiPseudowire.VfiPwIgmpSnoop})
    vfiPseudowire.EntityData.Leafs.Append("vfi-pw-mld-snoop", types.YLeaf{"VfiPwMldSnoop", vfiPseudowire.VfiPwMldSnoop})

    vfiPseudowire.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(vfiPseudowire.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop
// Attach a DHCP Snooping profile
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (vfiPwDhcpSnoop *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop) GetEntityData() *types.CommonEntityData {
    vfiPwDhcpSnoop.EntityData.YFilter = vfiPwDhcpSnoop.YFilter
    vfiPwDhcpSnoop.EntityData.YangName = "vfi-pw-dhcp-snoop"
    vfiPwDhcpSnoop.EntityData.BundleName = "cisco_ios_xr"
    vfiPwDhcpSnoop.EntityData.ParentYangName = "vfi-pseudowire"
    vfiPwDhcpSnoop.EntityData.SegmentPath = "vfi-pw-dhcp-snoop"
    vfiPwDhcpSnoop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPwDhcpSnoop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPwDhcpSnoop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPwDhcpSnoop.EntityData.Children = types.NewOrderedMap()
    vfiPwDhcpSnoop.EntityData.Leafs = types.NewOrderedMap()
    vfiPwDhcpSnoop.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", vfiPwDhcpSnoop.ProfileId})
    vfiPwDhcpSnoop.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", vfiPwDhcpSnoop.DhcpSnoopingId})

    vfiPwDhcpSnoop.EntityData.YListKeys = []string {}

    return &(vfiPwDhcpSnoop.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels
// MPLS static labels
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (vfiPwMplsStaticLabels *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    vfiPwMplsStaticLabels.EntityData.YFilter = vfiPwMplsStaticLabels.YFilter
    vfiPwMplsStaticLabels.EntityData.YangName = "vfi-pw-mpls-static-labels"
    vfiPwMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    vfiPwMplsStaticLabels.EntityData.ParentYangName = "vfi-pseudowire"
    vfiPwMplsStaticLabels.EntityData.SegmentPath = "vfi-pw-mpls-static-labels"
    vfiPwMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPwMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPwMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPwMplsStaticLabels.EntityData.Children = types.NewOrderedMap()
    vfiPwMplsStaticLabels.EntityData.Leafs = types.NewOrderedMap()
    vfiPwMplsStaticLabels.EntityData.Leafs.Append("local-static-label", types.YLeaf{"LocalStaticLabel", vfiPwMplsStaticLabels.LocalStaticLabel})
    vfiPwMplsStaticLabels.EntityData.Leafs.Append("remote-static-label", types.YLeaf{"RemoteStaticLabel", vfiPwMplsStaticLabels.RemoteStaticLabel})

    vfiPwMplsStaticLabels.EntityData.YListKeys = []string {}

    return &(vfiPwMplsStaticLabels.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress.
    PseudowireStaticMacAddress []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress
}

func (pseudowireStaticMacAddresses *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    pseudowireStaticMacAddresses.EntityData.YFilter = pseudowireStaticMacAddresses.YFilter
    pseudowireStaticMacAddresses.EntityData.YangName = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    pseudowireStaticMacAddresses.EntityData.ParentYangName = "vfi-pseudowire"
    pseudowireStaticMacAddresses.EntityData.SegmentPath = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireStaticMacAddresses.EntityData.Children = types.NewOrderedMap()
    pseudowireStaticMacAddresses.EntityData.Children.Append("pseudowire-static-mac-address", types.YChild{"PseudowireStaticMacAddress", nil})
    for i := range pseudowireStaticMacAddresses.PseudowireStaticMacAddress {
        pseudowireStaticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(pseudowireStaticMacAddresses.PseudowireStaticMacAddress[i]), types.YChild{"PseudowireStaticMacAddress", pseudowireStaticMacAddresses.PseudowireStaticMacAddress[i]})
    }
    pseudowireStaticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    pseudowireStaticMacAddresses.EntityData.YListKeys = []string {}

    return &(pseudowireStaticMacAddresses.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (pseudowireStaticMacAddress *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress) GetEntityData() *types.CommonEntityData {
    pseudowireStaticMacAddress.EntityData.YFilter = pseudowireStaticMacAddress.YFilter
    pseudowireStaticMacAddress.EntityData.YangName = "pseudowire-static-mac-address"
    pseudowireStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    pseudowireStaticMacAddress.EntityData.ParentYangName = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddress.EntityData.SegmentPath = "pseudowire-static-mac-address" + types.AddKeyToken(pseudowireStaticMacAddress.Address, "address")
    pseudowireStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireStaticMacAddress.EntityData.Children = types.NewOrderedMap()
    pseudowireStaticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    pseudowireStaticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", pseudowireStaticMacAddress.Address})

    pseudowireStaticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(pseudowireStaticMacAddress.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery
// Enable Autodiscovery BGP in this VFI
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Policy for installation of forwarding data to L2FIB. The type is
    // string.
    TablePolicy interface{}

    // Enable control-word for this VFI. The type is interface{}.
    AdControlWord interface{}

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Signaling Protocol LDP in this VFI configuration.
    LdpSignalingProtocol L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol

    // Route policy.
    BgpRoutePolicy L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy

    // Route Distinguisher.
    RouteDistinguisher L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher

    // Enable Signaling Protocol BGP in this VFI.
    BgpSignalingProtocol L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol

    // Route Target.
    RouteTargets L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets
}

func (bgpAutoDiscovery *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    bgpAutoDiscovery.EntityData.YFilter = bgpAutoDiscovery.YFilter
    bgpAutoDiscovery.EntityData.YangName = "bgp-auto-discovery"
    bgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    bgpAutoDiscovery.EntityData.ParentYangName = "vfi"
    bgpAutoDiscovery.EntityData.SegmentPath = "bgp-auto-discovery"
    bgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    bgpAutoDiscovery.EntityData.Children.Append("ldp-signaling-protocol", types.YChild{"LdpSignalingProtocol", &bgpAutoDiscovery.LdpSignalingProtocol})
    bgpAutoDiscovery.EntityData.Children.Append("bgp-route-policy", types.YChild{"BgpRoutePolicy", &bgpAutoDiscovery.BgpRoutePolicy})
    bgpAutoDiscovery.EntityData.Children.Append("route-distinguisher", types.YChild{"RouteDistinguisher", &bgpAutoDiscovery.RouteDistinguisher})
    bgpAutoDiscovery.EntityData.Children.Append("bgp-signaling-protocol", types.YChild{"BgpSignalingProtocol", &bgpAutoDiscovery.BgpSignalingProtocol})
    bgpAutoDiscovery.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &bgpAutoDiscovery.RouteTargets})
    bgpAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    bgpAutoDiscovery.EntityData.Leafs.Append("table-policy", types.YLeaf{"TablePolicy", bgpAutoDiscovery.TablePolicy})
    bgpAutoDiscovery.EntityData.Leafs.Append("ad-control-word", types.YLeaf{"AdControlWord", bgpAutoDiscovery.AdControlWord})
    bgpAutoDiscovery.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bgpAutoDiscovery.Enable})

    bgpAutoDiscovery.EntityData.YListKeys = []string {}

    return &(bgpAutoDiscovery.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol
// Signaling Protocol LDP in this VFI
// configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable LDP as Signaling Protocol.Deletion of this object also causes
    // deletion of all objects under LDPSignalingProtocol. The type is
    // interface{}.
    Enable interface{}

    // VPLS ID.
    VplsId L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance
}

func (ldpSignalingProtocol *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    ldpSignalingProtocol.EntityData.YFilter = ldpSignalingProtocol.YFilter
    ldpSignalingProtocol.EntityData.YangName = "ldp-signaling-protocol"
    ldpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    ldpSignalingProtocol.EntityData.ParentYangName = "bgp-auto-discovery"
    ldpSignalingProtocol.EntityData.SegmentPath = "ldp-signaling-protocol"
    ldpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpSignalingProtocol.EntityData.Children = types.NewOrderedMap()
    ldpSignalingProtocol.EntityData.Children.Append("vpls-id", types.YChild{"VplsId", &ldpSignalingProtocol.VplsId})
    ldpSignalingProtocol.EntityData.Children.Append("flow-label-load-balance", types.YChild{"FlowLabelLoadBalance", &ldpSignalingProtocol.FlowLabelLoadBalance})
    ldpSignalingProtocol.EntityData.Leafs = types.NewOrderedMap()
    ldpSignalingProtocol.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ldpSignalingProtocol.Enable})

    ldpSignalingProtocol.EntityData.YListKeys = []string {}

    return &(ldpSignalingProtocol.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId
// VPLS ID
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPLS-ID Type. The type is LdpVplsId.
    Type interface{}

    // Two byte AS number. The type is interface{} with range: 1..65535.
    As interface{}

    // AS index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Address index. The type is interface{} with range: 0..32767.
    AddressIndex interface{}
}

func (vplsId *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId) GetEntityData() *types.CommonEntityData {
    vplsId.EntityData.YFilter = vplsId.YFilter
    vplsId.EntityData.YangName = "vpls-id"
    vplsId.EntityData.BundleName = "cisco_ios_xr"
    vplsId.EntityData.ParentYangName = "ldp-signaling-protocol"
    vplsId.EntityData.SegmentPath = "vpls-id"
    vplsId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vplsId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vplsId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vplsId.EntityData.Children = types.NewOrderedMap()
    vplsId.EntityData.Leafs = types.NewOrderedMap()
    vplsId.EntityData.Leafs.Append("type", types.YLeaf{"Type", vplsId.Type})
    vplsId.EntityData.Leafs.Append("as", types.YLeaf{"As", vplsId.As})
    vplsId.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", vplsId.AsIndex})
    vplsId.EntityData.Leafs.Append("address", types.YLeaf{"Address", vplsId.Address})
    vplsId.EntityData.Leafs.Append("address-index", types.YLeaf{"AddressIndex", vplsId.AddressIndex})

    vplsId.EntityData.YListKeys = []string {}

    return &(vplsId.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "ldp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs.Append("flow-label", types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel})
    flowLabelLoadBalance.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabelLoadBalance.Static})

    flowLabelLoadBalance.EntityData.YListKeys = []string {}

    return &(flowLabelLoadBalance.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy
// Route policy
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Export route policy. The type is string.
    Export interface{}
}

func (bgpRoutePolicy *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy) GetEntityData() *types.CommonEntityData {
    bgpRoutePolicy.EntityData.YFilter = bgpRoutePolicy.YFilter
    bgpRoutePolicy.EntityData.YangName = "bgp-route-policy"
    bgpRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    bgpRoutePolicy.EntityData.ParentYangName = "bgp-auto-discovery"
    bgpRoutePolicy.EntityData.SegmentPath = "bgp-route-policy"
    bgpRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpRoutePolicy.EntityData.Children = types.NewOrderedMap()
    bgpRoutePolicy.EntityData.Leafs = types.NewOrderedMap()
    bgpRoutePolicy.EntityData.Leafs.Append("export", types.YLeaf{"Export", bgpRoutePolicy.Export})

    bgpRoutePolicy.EntityData.YListKeys = []string {}

    return &(bgpRoutePolicy.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher
// Route Distinguisher
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (routeDistinguisher *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher) GetEntityData() *types.CommonEntityData {
    routeDistinguisher.EntityData.YFilter = routeDistinguisher.YFilter
    routeDistinguisher.EntityData.YangName = "route-distinguisher"
    routeDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    routeDistinguisher.EntityData.ParentYangName = "bgp-auto-discovery"
    routeDistinguisher.EntityData.SegmentPath = "route-distinguisher"
    routeDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeDistinguisher.EntityData.Children = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    routeDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", routeDistinguisher.Type})
    routeDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", routeDistinguisher.As})
    routeDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeDistinguisher.AsIndex})
    routeDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeDistinguisher.Address})
    routeDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", routeDistinguisher.AddrIndex})

    routeDistinguisher.EntityData.YListKeys = []string {}

    return &(routeDistinguisher.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol
// Enable Signaling Protocol BGP in this VFI
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Virtual Edge Block Configurable Range. The type is interface{} with
    // range: 11..100.
    VeRange interface{}

    // Local Virtual Edge Identifier. The type is interface{} with range:
    // 1..16384.
    Veid interface{}

    // Enable BGP as Signaling Protocol. The type is interface{}.
    Enable interface{}

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance
}

func (bgpSignalingProtocol *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    bgpSignalingProtocol.EntityData.YFilter = bgpSignalingProtocol.YFilter
    bgpSignalingProtocol.EntityData.YangName = "bgp-signaling-protocol"
    bgpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    bgpSignalingProtocol.EntityData.ParentYangName = "bgp-auto-discovery"
    bgpSignalingProtocol.EntityData.SegmentPath = "bgp-signaling-protocol"
    bgpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpSignalingProtocol.EntityData.Children = types.NewOrderedMap()
    bgpSignalingProtocol.EntityData.Children.Append("flow-label-load-balance", types.YChild{"FlowLabelLoadBalance", &bgpSignalingProtocol.FlowLabelLoadBalance})
    bgpSignalingProtocol.EntityData.Leafs = types.NewOrderedMap()
    bgpSignalingProtocol.EntityData.Leafs.Append("ve-range", types.YLeaf{"VeRange", bgpSignalingProtocol.VeRange})
    bgpSignalingProtocol.EntityData.Leafs.Append("veid", types.YLeaf{"Veid", bgpSignalingProtocol.Veid})
    bgpSignalingProtocol.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bgpSignalingProtocol.Enable})

    bgpSignalingProtocol.EntityData.YListKeys = []string {}

    return &(bgpSignalingProtocol.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "bgp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs = types.NewOrderedMap()
    flowLabelLoadBalance.EntityData.Leafs.Append("flow-label", types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel})
    flowLabelLoadBalance.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabelLoadBalance.Static})

    flowLabelLoadBalance.EntityData.YListKeys = []string {}

    return &(flowLabelLoadBalance.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets
// Route Target
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget.
    RouteTarget []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget
}

func (routeTargets *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "bgp-auto-discovery"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget
// Name of the Route Target
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // two byte as or four byte as. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs.
    TwoByteAsOrFourByteAs []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs

    // ipv4 address. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + types.AddKeyToken(routeTarget.Role, "role") + types.AddKeyToken(routeTarget.Format, "format")
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("two-byte-as-or-four-byte-as", types.YChild{"TwoByteAsOrFourByteAs", nil})
    for i := range routeTarget.TwoByteAsOrFourByteAs {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.TwoByteAsOrFourByteAs[i]), types.YChild{"TwoByteAsOrFourByteAs", routeTarget.TwoByteAsOrFourByteAs[i]})
    }
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children.Append(types.GetSegmentPath(routeTarget.Ipv4Address[i]), types.YChild{"Ipv4Address", routeTarget.Ipv4Address[i]})
    }
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("role", types.YLeaf{"Role", routeTarget.Role})
    routeTarget.EntityData.Leafs.Append("format", types.YLeaf{"Format", routeTarget.Format})

    routeTarget.EntityData.YListKeys = []string {"Role", "Format"}

    return &(routeTarget.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs
// two byte as or four byte as
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}
}

func (twoByteAsOrFourByteAs *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAsOrFourByteAs.EntityData.YFilter = twoByteAsOrFourByteAs.YFilter
    twoByteAsOrFourByteAs.EntityData.YangName = "two-byte-as-or-four-byte-as"
    twoByteAsOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAsOrFourByteAs.EntityData.ParentYangName = "route-target"
    twoByteAsOrFourByteAs.EntityData.SegmentPath = "two-byte-as-or-four-byte-as" + types.AddKeyToken(twoByteAsOrFourByteAs.As, "as") + types.AddKeyToken(twoByteAsOrFourByteAs.AsIndex, "as-index")
    twoByteAsOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAsOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAsOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAsOrFourByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAsOrFourByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAsOrFourByteAs.EntityData.Leafs.Append("as", types.YLeaf{"As", twoByteAsOrFourByteAs.As})
    twoByteAsOrFourByteAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", twoByteAsOrFourByteAs.AsIndex})

    twoByteAsOrFourByteAs.EntityData.YListKeys = []string {"As", "AsIndex"}

    return &(twoByteAsOrFourByteAs.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}
}

func (ipv4Address *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address") + types.AddKeyToken(ipv4Address.AddrIndex, "addr-index")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", ipv4Address.AddrIndex})

    ipv4Address.EntityData.YListKeys = []string {"Address", "AddrIndex"}

    return &(ipv4Address.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis
// Bridge Domain EVPN VxLAN Network Identifier
// Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain VxLAN EVPN. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni.
    BridgeDomainvni []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni
}

func (bridgeDomainvnis *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis) GetEntityData() *types.CommonEntityData {
    bridgeDomainvnis.EntityData.YFilter = bridgeDomainvnis.YFilter
    bridgeDomainvnis.EntityData.YangName = "bridge-domainvnis"
    bridgeDomainvnis.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainvnis.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainvnis.EntityData.SegmentPath = "bridge-domainvnis"
    bridgeDomainvnis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainvnis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainvnis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainvnis.EntityData.Children = types.NewOrderedMap()
    bridgeDomainvnis.EntityData.Children.Append("bridge-domainvni", types.YChild{"BridgeDomainvni", nil})
    for i := range bridgeDomainvnis.BridgeDomainvni {
        bridgeDomainvnis.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainvnis.BridgeDomainvni[i]), types.YChild{"BridgeDomainvni", bridgeDomainvnis.BridgeDomainvni[i]})
    }
    bridgeDomainvnis.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainvnis.EntityData.YListKeys = []string {}

    return &(bridgeDomainvnis.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni
// Bridge Domain VxLAN EVPN
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VxLAN Ethernet VPN-ID. The type is interface{}
    // with range: 1..16777215.
    VpnId interface{}
}

func (bridgeDomainvni *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BridgeDomainvnis_BridgeDomainvni) GetEntityData() *types.CommonEntityData {
    bridgeDomainvni.EntityData.YFilter = bridgeDomainvni.YFilter
    bridgeDomainvni.EntityData.YangName = "bridge-domainvni"
    bridgeDomainvni.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainvni.EntityData.ParentYangName = "bridge-domainvnis"
    bridgeDomainvni.EntityData.SegmentPath = "bridge-domainvni" + types.AddKeyToken(bridgeDomainvni.VpnId, "vpn-id")
    bridgeDomainvni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainvni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainvni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainvni.EntityData.Children = types.NewOrderedMap()
    bridgeDomainvni.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomainvni.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", bridgeDomainvni.VpnId})

    bridgeDomainvni.EntityData.YListKeys = []string {"VpnId"}

    return &(bridgeDomainvni.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits
// Attachment Circuit table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Attachment Circuit. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit.
    BdAttachmentCircuit []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit
}

func (bdAttachmentCircuits *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    bdAttachmentCircuits.EntityData.YFilter = bdAttachmentCircuits.YFilter
    bdAttachmentCircuits.EntityData.YangName = "bd-attachment-circuits"
    bdAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    bdAttachmentCircuits.EntityData.ParentYangName = "bridge-domain"
    bdAttachmentCircuits.EntityData.SegmentPath = "bd-attachment-circuits"
    bdAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdAttachmentCircuits.EntityData.Children = types.NewOrderedMap()
    bdAttachmentCircuits.EntityData.Children.Append("bd-attachment-circuit", types.YChild{"BdAttachmentCircuit", nil})
    for i := range bdAttachmentCircuits.BdAttachmentCircuit {
        bdAttachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(bdAttachmentCircuits.BdAttachmentCircuit[i]), types.YChild{"BdAttachmentCircuit", bdAttachmentCircuits.BdAttachmentCircuit[i]})
    }
    bdAttachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    bdAttachmentCircuits.EntityData.YListKeys = []string {}

    return &(bdAttachmentCircuits.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit
// Name of the Attachment Circuit
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Attachment Circuit. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    Name interface{}

    // Enable or Disable Flooding. The type is InterfaceTrafficFlood.
    InterfaceFlooding interface{}

    // Attach a IGMP Snooping profile. The type is string with length: 1..32.
    InterfaceIgmpSnoop interface{}

    // Enable or Disable Unknown Unicast Flooding. The type is
    // InterfaceTrafficFlood.
    InterfaceFloodingUnknownUnicast interface{}

    // Attach a MLD Snooping profile. The type is string with length: 1..32.
    InterfaceMldSnoop interface{}

    // IP Source Guard.
    InterfaceIpSourceGuard L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard

    // L2 Interface Dynamic ARP Inspection.
    InterfaceDai L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai

    // Attach a DHCP profile.
    InterfaceProfile L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile

    // Storm Control.
    BdacStormControlTypes L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes

    // Split Horizon.
    SplitHorizon L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon

    // Static Mac Address Table.
    StaticMacAddresses L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses

    // MAC configuration commands.
    InterfaceMac L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac
}

func (bdAttachmentCircuit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    bdAttachmentCircuit.EntityData.YFilter = bdAttachmentCircuit.YFilter
    bdAttachmentCircuit.EntityData.YangName = "bd-attachment-circuit"
    bdAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    bdAttachmentCircuit.EntityData.ParentYangName = "bd-attachment-circuits"
    bdAttachmentCircuit.EntityData.SegmentPath = "bd-attachment-circuit" + types.AddKeyToken(bdAttachmentCircuit.Name, "name")
    bdAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdAttachmentCircuit.EntityData.Children = types.NewOrderedMap()
    bdAttachmentCircuit.EntityData.Children.Append("interface-ip-source-guard", types.YChild{"InterfaceIpSourceGuard", &bdAttachmentCircuit.InterfaceIpSourceGuard})
    bdAttachmentCircuit.EntityData.Children.Append("interface-dai", types.YChild{"InterfaceDai", &bdAttachmentCircuit.InterfaceDai})
    bdAttachmentCircuit.EntityData.Children.Append("interface-profile", types.YChild{"InterfaceProfile", &bdAttachmentCircuit.InterfaceProfile})
    bdAttachmentCircuit.EntityData.Children.Append("bdac-storm-control-types", types.YChild{"BdacStormControlTypes", &bdAttachmentCircuit.BdacStormControlTypes})
    bdAttachmentCircuit.EntityData.Children.Append("split-horizon", types.YChild{"SplitHorizon", &bdAttachmentCircuit.SplitHorizon})
    bdAttachmentCircuit.EntityData.Children.Append("static-mac-addresses", types.YChild{"StaticMacAddresses", &bdAttachmentCircuit.StaticMacAddresses})
    bdAttachmentCircuit.EntityData.Children.Append("interface-mac", types.YChild{"InterfaceMac", &bdAttachmentCircuit.InterfaceMac})
    bdAttachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    bdAttachmentCircuit.EntityData.Leafs.Append("name", types.YLeaf{"Name", bdAttachmentCircuit.Name})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-flooding", types.YLeaf{"InterfaceFlooding", bdAttachmentCircuit.InterfaceFlooding})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-igmp-snoop", types.YLeaf{"InterfaceIgmpSnoop", bdAttachmentCircuit.InterfaceIgmpSnoop})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-flooding-unknown-unicast", types.YLeaf{"InterfaceFloodingUnknownUnicast", bdAttachmentCircuit.InterfaceFloodingUnknownUnicast})
    bdAttachmentCircuit.EntityData.Leafs.Append("interface-mld-snoop", types.YLeaf{"InterfaceMldSnoop", bdAttachmentCircuit.InterfaceMldSnoop})

    bdAttachmentCircuit.EntityData.YListKeys = []string {"Name"}

    return &(bdAttachmentCircuit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard
// IP Source Guard
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface Dynamic IP source guard. The type is interface{}.
    Disable interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (interfaceIpSourceGuard *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard) GetEntityData() *types.CommonEntityData {
    interfaceIpSourceGuard.EntityData.YFilter = interfaceIpSourceGuard.YFilter
    interfaceIpSourceGuard.EntityData.YangName = "interface-ip-source-guard"
    interfaceIpSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    interfaceIpSourceGuard.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceIpSourceGuard.EntityData.SegmentPath = "interface-ip-source-guard"
    interfaceIpSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIpSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIpSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIpSourceGuard.EntityData.Children = types.NewOrderedMap()
    interfaceIpSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    interfaceIpSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", interfaceIpSourceGuard.Logging})
    interfaceIpSourceGuard.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", interfaceIpSourceGuard.Disable})
    interfaceIpSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceIpSourceGuard.Enable})

    interfaceIpSourceGuard.EntityData.YListKeys = []string {}

    return &(interfaceIpSourceGuard.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai
// L2 Interface Dynamic ARP Inspection
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface Dynamic ARP Inspection. The type is interface{}.
    Disable interface{}

    // Enable L2 Interface Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    InterfaceDaiAddressValidation L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation
}

func (interfaceDai *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai) GetEntityData() *types.CommonEntityData {
    interfaceDai.EntityData.YFilter = interfaceDai.YFilter
    interfaceDai.EntityData.YangName = "interface-dai"
    interfaceDai.EntityData.BundleName = "cisco_ios_xr"
    interfaceDai.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceDai.EntityData.SegmentPath = "interface-dai"
    interfaceDai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDai.EntityData.Children = types.NewOrderedMap()
    interfaceDai.EntityData.Children.Append("interface-dai-address-validation", types.YChild{"InterfaceDaiAddressValidation", &interfaceDai.InterfaceDaiAddressValidation})
    interfaceDai.EntityData.Leafs = types.NewOrderedMap()
    interfaceDai.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", interfaceDai.Logging})
    interfaceDai.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", interfaceDai.Disable})
    interfaceDai.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceDai.Enable})

    interfaceDai.EntityData.YListKeys = []string {}

    return &(interfaceDai.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation
// Address Validation
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Verification. The type is L2vpnVerification.
    Ipv4Verification interface{}

    // Destination MAC Verification. The type is L2vpnVerification.
    DestinationMacVerification interface{}

    // Source MAC Verification. The type is L2vpnVerification.
    SourceMacVerification interface{}

    // Enable Address Validation. The type is interface{}.
    Enable interface{}
}

func (interfaceDaiAddressValidation *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation) GetEntityData() *types.CommonEntityData {
    interfaceDaiAddressValidation.EntityData.YFilter = interfaceDaiAddressValidation.YFilter
    interfaceDaiAddressValidation.EntityData.YangName = "interface-dai-address-validation"
    interfaceDaiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    interfaceDaiAddressValidation.EntityData.ParentYangName = "interface-dai"
    interfaceDaiAddressValidation.EntityData.SegmentPath = "interface-dai-address-validation"
    interfaceDaiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDaiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDaiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDaiAddressValidation.EntityData.Children = types.NewOrderedMap()
    interfaceDaiAddressValidation.EntityData.Leafs = types.NewOrderedMap()
    interfaceDaiAddressValidation.EntityData.Leafs.Append("ipv4-verification", types.YLeaf{"Ipv4Verification", interfaceDaiAddressValidation.Ipv4Verification})
    interfaceDaiAddressValidation.EntityData.Leafs.Append("destination-mac-verification", types.YLeaf{"DestinationMacVerification", interfaceDaiAddressValidation.DestinationMacVerification})
    interfaceDaiAddressValidation.EntityData.Leafs.Append("source-mac-verification", types.YLeaf{"SourceMacVerification", interfaceDaiAddressValidation.SourceMacVerification})
    interfaceDaiAddressValidation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceDaiAddressValidation.Enable})

    interfaceDaiAddressValidation.EntityData.YListKeys = []string {}

    return &(interfaceDaiAddressValidation.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile
// Attach a DHCP profile
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (interfaceProfile *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile) GetEntityData() *types.CommonEntityData {
    interfaceProfile.EntityData.YFilter = interfaceProfile.YFilter
    interfaceProfile.EntityData.YangName = "interface-profile"
    interfaceProfile.EntityData.BundleName = "cisco_ios_xr"
    interfaceProfile.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceProfile.EntityData.SegmentPath = "interface-profile"
    interfaceProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProfile.EntityData.Children = types.NewOrderedMap()
    interfaceProfile.EntityData.Leafs = types.NewOrderedMap()
    interfaceProfile.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", interfaceProfile.ProfileId})
    interfaceProfile.EntityData.Leafs.Append("dhcp-snooping-id", types.YLeaf{"DhcpSnoopingId", interfaceProfile.DhcpSnoopingId})

    interfaceProfile.EntityData.YListKeys = []string {}

    return &(interfaceProfile.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes
// Storm Control
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType.
    BdacStormControlType []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType
}

func (bdacStormControlTypes *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes) GetEntityData() *types.CommonEntityData {
    bdacStormControlTypes.EntityData.YFilter = bdacStormControlTypes.YFilter
    bdacStormControlTypes.EntityData.YangName = "bdac-storm-control-types"
    bdacStormControlTypes.EntityData.BundleName = "cisco_ios_xr"
    bdacStormControlTypes.EntityData.ParentYangName = "bd-attachment-circuit"
    bdacStormControlTypes.EntityData.SegmentPath = "bdac-storm-control-types"
    bdacStormControlTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdacStormControlTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdacStormControlTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdacStormControlTypes.EntityData.Children = types.NewOrderedMap()
    bdacStormControlTypes.EntityData.Children.Append("bdac-storm-control-type", types.YChild{"BdacStormControlType", nil})
    for i := range bdacStormControlTypes.BdacStormControlType {
        bdacStormControlTypes.EntityData.Children.Append(types.GetSegmentPath(bdacStormControlTypes.BdacStormControlType[i]), types.YChild{"BdacStormControlType", bdacStormControlTypes.BdacStormControlType[i]})
    }
    bdacStormControlTypes.EntityData.Leafs = types.NewOrderedMap()

    bdacStormControlTypes.EntityData.YListKeys = []string {}

    return &(bdacStormControlTypes.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType
// Storm Control Type
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit
}

func (bdacStormControlType *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType) GetEntityData() *types.CommonEntityData {
    bdacStormControlType.EntityData.YFilter = bdacStormControlType.YFilter
    bdacStormControlType.EntityData.YangName = "bdac-storm-control-type"
    bdacStormControlType.EntityData.BundleName = "cisco_ios_xr"
    bdacStormControlType.EntityData.ParentYangName = "bdac-storm-control-types"
    bdacStormControlType.EntityData.SegmentPath = "bdac-storm-control-type" + types.AddKeyToken(bdacStormControlType.Sctype, "sctype")
    bdacStormControlType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdacStormControlType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdacStormControlType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdacStormControlType.EntityData.Children = types.NewOrderedMap()
    bdacStormControlType.EntityData.Children.Append("storm-control-unit", types.YChild{"StormControlUnit", &bdacStormControlType.StormControlUnit})
    bdacStormControlType.EntityData.Leafs = types.NewOrderedMap()
    bdacStormControlType.EntityData.Leafs.Append("sctype", types.YLeaf{"Sctype", bdacStormControlType.Sctype})

    bdacStormControlType.EntityData.YListKeys = []string {"Sctype"}

    return &(bdacStormControlType.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit
// Specify units for Storm Control Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Kilobits Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 64..1280000. Units are
    // kbit/s.
    KbitsPerSec interface{}

    // Packets Per Second, PktsPerSec and KbitsPerSec cannot be configured
    // together. The type is interface{} with range: 1..160000. Units are
    // packet/s.
    PktsPerSec interface{}
}

func (stormControlUnit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bdac-storm-control-type"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs = types.NewOrderedMap()
    stormControlUnit.EntityData.Leafs.Append("kbits-per-sec", types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec})
    stormControlUnit.EntityData.Leafs.Append("pkts-per-sec", types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec})

    stormControlUnit.EntityData.YListKeys = []string {}

    return &(stormControlUnit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon
// Split Horizon
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon Group ID.
    SplitHorizonGroupId L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId
}

func (splitHorizon *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon) GetEntityData() *types.CommonEntityData {
    splitHorizon.EntityData.YFilter = splitHorizon.YFilter
    splitHorizon.EntityData.YangName = "split-horizon"
    splitHorizon.EntityData.BundleName = "cisco_ios_xr"
    splitHorizon.EntityData.ParentYangName = "bd-attachment-circuit"
    splitHorizon.EntityData.SegmentPath = "split-horizon"
    splitHorizon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    splitHorizon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    splitHorizon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    splitHorizon.EntityData.Children = types.NewOrderedMap()
    splitHorizon.EntityData.Children.Append("split-horizon-group-id", types.YChild{"SplitHorizonGroupId", &splitHorizon.SplitHorizonGroupId})
    splitHorizon.EntityData.Leafs = types.NewOrderedMap()

    splitHorizon.EntityData.YListKeys = []string {}

    return &(splitHorizon.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId
// Split Horizon Group ID
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable split horizon group. The type is interface{}.
    Enable interface{}
}

func (splitHorizonGroupId *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId) GetEntityData() *types.CommonEntityData {
    splitHorizonGroupId.EntityData.YFilter = splitHorizonGroupId.YFilter
    splitHorizonGroupId.EntityData.YangName = "split-horizon-group-id"
    splitHorizonGroupId.EntityData.BundleName = "cisco_ios_xr"
    splitHorizonGroupId.EntityData.ParentYangName = "split-horizon"
    splitHorizonGroupId.EntityData.SegmentPath = "split-horizon-group-id"
    splitHorizonGroupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    splitHorizonGroupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    splitHorizonGroupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    splitHorizonGroupId.EntityData.Children = types.NewOrderedMap()
    splitHorizonGroupId.EntityData.Leafs = types.NewOrderedMap()
    splitHorizonGroupId.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", splitHorizonGroupId.Enable})

    splitHorizonGroupId.EntityData.YListKeys = []string {}

    return &(splitHorizonGroupId.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses
// Static Mac Address Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress.
    StaticMacAddress []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress
}

func (staticMacAddresses *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses) GetEntityData() *types.CommonEntityData {
    staticMacAddresses.EntityData.YFilter = staticMacAddresses.YFilter
    staticMacAddresses.EntityData.YangName = "static-mac-addresses"
    staticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticMacAddresses.EntityData.ParentYangName = "bd-attachment-circuit"
    staticMacAddresses.EntityData.SegmentPath = "static-mac-addresses"
    staticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMacAddresses.EntityData.Children = types.NewOrderedMap()
    staticMacAddresses.EntityData.Children.Append("static-mac-address", types.YChild{"StaticMacAddress", nil})
    for i := range staticMacAddresses.StaticMacAddress {
        staticMacAddresses.EntityData.Children.Append(types.GetSegmentPath(staticMacAddresses.StaticMacAddress[i]), types.YChild{"StaticMacAddress", staticMacAddresses.StaticMacAddress[i]})
    }
    staticMacAddresses.EntityData.Leafs = types.NewOrderedMap()

    staticMacAddresses.EntityData.YListKeys = []string {}

    return &(staticMacAddresses.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress
// Static Mac Address Configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}
}

func (staticMacAddress *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress) GetEntityData() *types.CommonEntityData {
    staticMacAddress.EntityData.YFilter = staticMacAddress.YFilter
    staticMacAddress.EntityData.YangName = "static-mac-address"
    staticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    staticMacAddress.EntityData.ParentYangName = "static-mac-addresses"
    staticMacAddress.EntityData.SegmentPath = "static-mac-address" + types.AddKeyToken(staticMacAddress.Address, "address")
    staticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMacAddress.EntityData.Children = types.NewOrderedMap()
    staticMacAddress.EntityData.Leafs = types.NewOrderedMap()
    staticMacAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", staticMacAddress.Address})

    staticMacAddress.EntityData.YListKeys = []string {"Address"}

    return &(staticMacAddress.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac
// MAC configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MAC Flush When Port goes down. The type is PortDownFlush.
    InterfaceMacPortDownFlush interface{}

    // Enable Mac Learning. The type is MacLearn.
    InterfaceMacLearning interface{}

    // MAC-Aging configuration commands.
    InterfaceMacAging L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging

    // MAC Secure.
    InterfaceMacSecure L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure

    // MAC-Limit configuration commands.
    InterfaceMacLimit L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit
}

func (interfaceMac *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac) GetEntityData() *types.CommonEntityData {
    interfaceMac.EntityData.YFilter = interfaceMac.YFilter
    interfaceMac.EntityData.YangName = "interface-mac"
    interfaceMac.EntityData.BundleName = "cisco_ios_xr"
    interfaceMac.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceMac.EntityData.SegmentPath = "interface-mac"
    interfaceMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMac.EntityData.Children = types.NewOrderedMap()
    interfaceMac.EntityData.Children.Append("interface-mac-aging", types.YChild{"InterfaceMacAging", &interfaceMac.InterfaceMacAging})
    interfaceMac.EntityData.Children.Append("interface-mac-secure", types.YChild{"InterfaceMacSecure", &interfaceMac.InterfaceMacSecure})
    interfaceMac.EntityData.Children.Append("interface-mac-limit", types.YChild{"InterfaceMacLimit", &interfaceMac.InterfaceMacLimit})
    interfaceMac.EntityData.Leafs = types.NewOrderedMap()
    interfaceMac.EntityData.Leafs.Append("interface-mac-port-down-flush", types.YLeaf{"InterfaceMacPortDownFlush", interfaceMac.InterfaceMacPortDownFlush})
    interfaceMac.EntityData.Leafs.Append("interface-mac-learning", types.YLeaf{"InterfaceMacLearning", interfaceMac.InterfaceMacLearning})

    interfaceMac.EntityData.YListKeys = []string {}

    return &(interfaceMac.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging
// MAC-Aging configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    InterfaceMacAgingTime interface{}

    // MAC address aging type. The type is MacAging.
    InterfaceMacAgingType interface{}
}

func (interfaceMacAging *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging) GetEntityData() *types.CommonEntityData {
    interfaceMacAging.EntityData.YFilter = interfaceMacAging.YFilter
    interfaceMacAging.EntityData.YangName = "interface-mac-aging"
    interfaceMacAging.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacAging.EntityData.ParentYangName = "interface-mac"
    interfaceMacAging.EntityData.SegmentPath = "interface-mac-aging"
    interfaceMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacAging.EntityData.Children = types.NewOrderedMap()
    interfaceMacAging.EntityData.Leafs = types.NewOrderedMap()
    interfaceMacAging.EntityData.Leafs.Append("interface-mac-aging-time", types.YLeaf{"InterfaceMacAgingTime", interfaceMacAging.InterfaceMacAgingTime})
    interfaceMacAging.EntityData.Leafs.Append("interface-mac-aging-type", types.YLeaf{"InterfaceMacAgingType", interfaceMacAging.InterfaceMacAgingType})

    interfaceMacAging.EntityData.YListKeys = []string {}

    return &(interfaceMacAging.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure
// MAC Secure
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Secure Logging. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface MAC Secure. The type is interface{}.
    Disable interface{}

    // MAC secure enforcement action. The type is MacSecureAction.
    Action interface{}

    // Enable MAC Secure. The type is interface{}.
    Enable interface{}
}

func (interfaceMacSecure *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure) GetEntityData() *types.CommonEntityData {
    interfaceMacSecure.EntityData.YFilter = interfaceMacSecure.YFilter
    interfaceMacSecure.EntityData.YangName = "interface-mac-secure"
    interfaceMacSecure.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacSecure.EntityData.ParentYangName = "interface-mac"
    interfaceMacSecure.EntityData.SegmentPath = "interface-mac-secure"
    interfaceMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacSecure.EntityData.Children = types.NewOrderedMap()
    interfaceMacSecure.EntityData.Leafs = types.NewOrderedMap()
    interfaceMacSecure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", interfaceMacSecure.Logging})
    interfaceMacSecure.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", interfaceMacSecure.Disable})
    interfaceMacSecure.EntityData.Leafs.Append("action", types.YLeaf{"Action", interfaceMacSecure.Action})
    interfaceMacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceMacSecure.Enable})

    interfaceMacSecure.EntityData.YListKeys = []string {}

    return &(interfaceMacSecure.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit
// MAC-Limit configuration commands
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of MAC addresses on an Interface after which MAC limit action is
    // taken. The type is interface{} with range: 0..4294967295.
    InterfaceMacLimitMax interface{}

    // MAC address limit notification action in a Interface. The type is
    // MacNotification.
    InterfaceMacLimitNotif interface{}

    // Interface MAC address limit enforcement action. The type is MacLimitAction.
    InterfaceMacLimitAction interface{}
}

func (interfaceMacLimit *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit) GetEntityData() *types.CommonEntityData {
    interfaceMacLimit.EntityData.YFilter = interfaceMacLimit.YFilter
    interfaceMacLimit.EntityData.YangName = "interface-mac-limit"
    interfaceMacLimit.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacLimit.EntityData.ParentYangName = "interface-mac"
    interfaceMacLimit.EntityData.SegmentPath = "interface-mac-limit"
    interfaceMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacLimit.EntityData.Children = types.NewOrderedMap()
    interfaceMacLimit.EntityData.Leafs = types.NewOrderedMap()
    interfaceMacLimit.EntityData.Leafs.Append("interface-mac-limit-max", types.YLeaf{"InterfaceMacLimitMax", interfaceMacLimit.InterfaceMacLimitMax})
    interfaceMacLimit.EntityData.Leafs.Append("interface-mac-limit-notif", types.YLeaf{"InterfaceMacLimitNotif", interfaceMacLimit.InterfaceMacLimitNotif})
    interfaceMacLimit.EntityData.Leafs.Append("interface-mac-limit-action", types.YLeaf{"InterfaceMacLimitAction", interfaceMacLimit.InterfaceMacLimitAction})

    interfaceMacLimit.EntityData.YListKeys = []string {}

    return &(interfaceMacLimit.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns
// List of EVPN pseudowires
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Pseudowire configuration. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn.
    BdPseudowireEvpn []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn
}

func (bdPseudowireEvpns *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns) GetEntityData() *types.CommonEntityData {
    bdPseudowireEvpns.EntityData.YFilter = bdPseudowireEvpns.YFilter
    bdPseudowireEvpns.EntityData.YangName = "bd-pseudowire-evpns"
    bdPseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowireEvpns.EntityData.ParentYangName = "bridge-domain"
    bdPseudowireEvpns.EntityData.SegmentPath = "bd-pseudowire-evpns"
    bdPseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowireEvpns.EntityData.Children = types.NewOrderedMap()
    bdPseudowireEvpns.EntityData.Children.Append("bd-pseudowire-evpn", types.YChild{"BdPseudowireEvpn", nil})
    for i := range bdPseudowireEvpns.BdPseudowireEvpn {
        bdPseudowireEvpns.EntityData.Children.Append(types.GetSegmentPath(bdPseudowireEvpns.BdPseudowireEvpn[i]), types.YChild{"BdPseudowireEvpn", bdPseudowireEvpns.BdPseudowireEvpn[i]})
    }
    bdPseudowireEvpns.EntityData.Leafs = types.NewOrderedMap()

    bdPseudowireEvpns.EntityData.YListKeys = []string {}

    return &(bdPseudowireEvpns.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn
// EVPN Pseudowire configuration
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // This attribute is a key. AC ID. The type is interface{} with range:
    // 1..4294967295.
    Acid interface{}
}

func (bdPseudowireEvpn *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn) GetEntityData() *types.CommonEntityData {
    bdPseudowireEvpn.EntityData.YFilter = bdPseudowireEvpn.YFilter
    bdPseudowireEvpn.EntityData.YangName = "bd-pseudowire-evpn"
    bdPseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowireEvpn.EntityData.ParentYangName = "bd-pseudowire-evpns"
    bdPseudowireEvpn.EntityData.SegmentPath = "bd-pseudowire-evpn" + types.AddKeyToken(bdPseudowireEvpn.Eviid, "eviid") + types.AddKeyToken(bdPseudowireEvpn.Acid, "acid")
    bdPseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowireEvpn.EntityData.Children = types.NewOrderedMap()
    bdPseudowireEvpn.EntityData.Leafs = types.NewOrderedMap()
    bdPseudowireEvpn.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", bdPseudowireEvpn.Eviid})
    bdPseudowireEvpn.EntityData.Leafs.Append("acid", types.YLeaf{"Acid", bdPseudowireEvpn.Acid})

    bdPseudowireEvpn.EntityData.YListKeys = []string {"Eviid", "Acid"}

    return &(bdPseudowireEvpn.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_IpSourceGuard
// IP Source Guard
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_IpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Logging. The type is interface{}.
    Logging interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (ipSourceGuard *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_IpSourceGuard) GetEntityData() *types.CommonEntityData {
    ipSourceGuard.EntityData.YFilter = ipSourceGuard.YFilter
    ipSourceGuard.EntityData.YangName = "ip-source-guard"
    ipSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    ipSourceGuard.EntityData.ParentYangName = "bridge-domain"
    ipSourceGuard.EntityData.SegmentPath = "ip-source-guard"
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSourceGuard.EntityData.Children = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", ipSourceGuard.Logging})
    ipSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipSourceGuard.Enable})

    ipSourceGuard.EntityData.YListKeys = []string {}

    return &(ipSourceGuard.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai
// Dynamic ARP Inspection
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Logging. The type is interface{}.
    Logging interface{}

    // Enable Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    DaiAddressValidation L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation
}

func (dai *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai) GetEntityData() *types.CommonEntityData {
    dai.EntityData.YFilter = dai.YFilter
    dai.EntityData.YangName = "dai"
    dai.EntityData.BundleName = "cisco_ios_xr"
    dai.EntityData.ParentYangName = "bridge-domain"
    dai.EntityData.SegmentPath = "dai"
    dai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dai.EntityData.Children = types.NewOrderedMap()
    dai.EntityData.Children.Append("dai-address-validation", types.YChild{"DaiAddressValidation", &dai.DaiAddressValidation})
    dai.EntityData.Leafs = types.NewOrderedMap()
    dai.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", dai.Logging})
    dai.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", dai.Enable})

    dai.EntityData.YListKeys = []string {}

    return &(dai.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation
// Address Validation
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPv4 Verification. The type is interface{}.
    Ipv4Verification interface{}

    // Enable Destination MAC Verification. The type is interface{}.
    DestinationMacVerification interface{}

    // Enable Source MAC Verification. The type is interface{}.
    SourceMacVerification interface{}

    // Enable Address Validation. The type is interface{}.
    Enable interface{}
}

func (daiAddressValidation *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation) GetEntityData() *types.CommonEntityData {
    daiAddressValidation.EntityData.YFilter = daiAddressValidation.YFilter
    daiAddressValidation.EntityData.YangName = "dai-address-validation"
    daiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    daiAddressValidation.EntityData.ParentYangName = "dai"
    daiAddressValidation.EntityData.SegmentPath = "dai-address-validation"
    daiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    daiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    daiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    daiAddressValidation.EntityData.Children = types.NewOrderedMap()
    daiAddressValidation.EntityData.Leafs = types.NewOrderedMap()
    daiAddressValidation.EntityData.Leafs.Append("ipv4-verification", types.YLeaf{"Ipv4Verification", daiAddressValidation.Ipv4Verification})
    daiAddressValidation.EntityData.Leafs.Append("destination-mac-verification", types.YLeaf{"DestinationMacVerification", daiAddressValidation.DestinationMacVerification})
    daiAddressValidation.EntityData.Leafs.Append("source-mac-verification", types.YLeaf{"SourceMacVerification", daiAddressValidation.SourceMacVerification})
    daiAddressValidation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", daiAddressValidation.Enable})

    daiAddressValidation.EntityData.YListKeys = []string {}

    return &(daiAddressValidation.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces
// Bridge Domain Routed Interface Table
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain Routed Interface. The type is slice of
    // L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface.
    RoutedInterface []*L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface
}

func (routedInterfaces *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces) GetEntityData() *types.CommonEntityData {
    routedInterfaces.EntityData.YFilter = routedInterfaces.YFilter
    routedInterfaces.EntityData.YangName = "routed-interfaces"
    routedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaces.EntityData.ParentYangName = "bridge-domain"
    routedInterfaces.EntityData.SegmentPath = "routed-interfaces"
    routedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaces.EntityData.Children = types.NewOrderedMap()
    routedInterfaces.EntityData.Children.Append("routed-interface", types.YChild{"RoutedInterface", nil})
    for i := range routedInterfaces.RoutedInterface {
        routedInterfaces.EntityData.Children.Append(types.GetSegmentPath(routedInterfaces.RoutedInterface[i]), types.YChild{"RoutedInterface", routedInterfaces.RoutedInterface[i]})
    }
    routedInterfaces.EntityData.Leafs = types.NewOrderedMap()

    routedInterfaces.EntityData.YListKeys = []string {}

    return &(routedInterfaces.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface
// Bridge Domain Routed Interface
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Routed Interface. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Routed interface split horizon group.
    RoutedInterfaceSplitHorizonGroup L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup
}

func (routedInterface *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface) GetEntityData() *types.CommonEntityData {
    routedInterface.EntityData.YFilter = routedInterface.YFilter
    routedInterface.EntityData.YangName = "routed-interface"
    routedInterface.EntityData.BundleName = "cisco_ios_xr"
    routedInterface.EntityData.ParentYangName = "routed-interfaces"
    routedInterface.EntityData.SegmentPath = "routed-interface" + types.AddKeyToken(routedInterface.InterfaceName, "interface-name")
    routedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterface.EntityData.Children = types.NewOrderedMap()
    routedInterface.EntityData.Children.Append("routed-interface-split-horizon-group", types.YChild{"RoutedInterfaceSplitHorizonGroup", &routedInterface.RoutedInterfaceSplitHorizonGroup})
    routedInterface.EntityData.Leafs = types.NewOrderedMap()
    routedInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", routedInterface.InterfaceName})

    routedInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(routedInterface.EntityData)
}

// L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup
// Routed interface split horizon group
type L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure BVI under SHG 1. The type is interface{}.
    RoutedInterfaceSplitHorizonGroupCore interface{}
}

func (routedInterfaceSplitHorizonGroup *L2vpn_Database_VlanSwitches_VlanSwitch_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    routedInterfaceSplitHorizonGroup.EntityData.YFilter = routedInterfaceSplitHorizonGroup.YFilter
    routedInterfaceSplitHorizonGroup.EntityData.YangName = "routed-interface-split-horizon-group"
    routedInterfaceSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaceSplitHorizonGroup.EntityData.ParentYangName = "routed-interface"
    routedInterfaceSplitHorizonGroup.EntityData.SegmentPath = "routed-interface-split-horizon-group"
    routedInterfaceSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaceSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaceSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaceSplitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    routedInterfaceSplitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    routedInterfaceSplitHorizonGroup.EntityData.Leafs.Append("routed-interface-split-horizon-group-core", types.YLeaf{"RoutedInterfaceSplitHorizonGroupCore", routedInterfaceSplitHorizonGroup.RoutedInterfaceSplitHorizonGroupCore})

    routedInterfaceSplitHorizonGroup.EntityData.YListKeys = []string {}

    return &(routedInterfaceSplitHorizonGroup.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable
// List of Flexible XConnect Services
type L2vpn_Database_FlexibleXconnectServiceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Vlan-Unaware Flexible XConnect Services.
    VlanUnawareFlexibleXconnectServices L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices

    // List of Vlan-Aware Flexible XConnect Services.
    VlanAwareFlexibleXconnectServices L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices
}

func (flexibleXconnectServiceTable *L2vpn_Database_FlexibleXconnectServiceTable) GetEntityData() *types.CommonEntityData {
    flexibleXconnectServiceTable.EntityData.YFilter = flexibleXconnectServiceTable.YFilter
    flexibleXconnectServiceTable.EntityData.YangName = "flexible-xconnect-service-table"
    flexibleXconnectServiceTable.EntityData.BundleName = "cisco_ios_xr"
    flexibleXconnectServiceTable.EntityData.ParentYangName = "database"
    flexibleXconnectServiceTable.EntityData.SegmentPath = "flexible-xconnect-service-table"
    flexibleXconnectServiceTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flexibleXconnectServiceTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flexibleXconnectServiceTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flexibleXconnectServiceTable.EntityData.Children = types.NewOrderedMap()
    flexibleXconnectServiceTable.EntityData.Children.Append("vlan-unaware-flexible-xconnect-services", types.YChild{"VlanUnawareFlexibleXconnectServices", &flexibleXconnectServiceTable.VlanUnawareFlexibleXconnectServices})
    flexibleXconnectServiceTable.EntityData.Children.Append("vlan-aware-flexible-xconnect-services", types.YChild{"VlanAwareFlexibleXconnectServices", &flexibleXconnectServiceTable.VlanAwareFlexibleXconnectServices})
    flexibleXconnectServiceTable.EntityData.Leafs = types.NewOrderedMap()

    flexibleXconnectServiceTable.EntityData.YListKeys = []string {}

    return &(flexibleXconnectServiceTable.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices
// List of Vlan-Unaware Flexible XConnect
// Services
type L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flexible XConnect Service. The type is slice of
    // L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService.
    VlanUnawareFlexibleXconnectService []*L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService
}

func (vlanUnawareFlexibleXconnectServices *L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices) GetEntityData() *types.CommonEntityData {
    vlanUnawareFlexibleXconnectServices.EntityData.YFilter = vlanUnawareFlexibleXconnectServices.YFilter
    vlanUnawareFlexibleXconnectServices.EntityData.YangName = "vlan-unaware-flexible-xconnect-services"
    vlanUnawareFlexibleXconnectServices.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFlexibleXconnectServices.EntityData.ParentYangName = "flexible-xconnect-service-table"
    vlanUnawareFlexibleXconnectServices.EntityData.SegmentPath = "vlan-unaware-flexible-xconnect-services"
    vlanUnawareFlexibleXconnectServices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFlexibleXconnectServices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFlexibleXconnectServices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFlexibleXconnectServices.EntityData.Children = types.NewOrderedMap()
    vlanUnawareFlexibleXconnectServices.EntityData.Children.Append("vlan-unaware-flexible-xconnect-service", types.YChild{"VlanUnawareFlexibleXconnectService", nil})
    for i := range vlanUnawareFlexibleXconnectServices.VlanUnawareFlexibleXconnectService {
        vlanUnawareFlexibleXconnectServices.EntityData.Children.Append(types.GetSegmentPath(vlanUnawareFlexibleXconnectServices.VlanUnawareFlexibleXconnectService[i]), types.YChild{"VlanUnawareFlexibleXconnectService", vlanUnawareFlexibleXconnectServices.VlanUnawareFlexibleXconnectService[i]})
    }
    vlanUnawareFlexibleXconnectServices.EntityData.Leafs = types.NewOrderedMap()

    vlanUnawareFlexibleXconnectServices.EntityData.YListKeys = []string {}

    return &(vlanUnawareFlexibleXconnectServices.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService
// Flexible XConnect Service
type L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Flexible XConnect Service. The type is
    // string with length: 1..23.
    Name interface{}

    // List of attachment circuits.
    VlanUnawareFxcAttachmentCircuits L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits

    // List of EVPN Services.
    VlanUnawareFxcPseudowireEvpns L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns
}

func (vlanUnawareFlexibleXconnectService *L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService) GetEntityData() *types.CommonEntityData {
    vlanUnawareFlexibleXconnectService.EntityData.YFilter = vlanUnawareFlexibleXconnectService.YFilter
    vlanUnawareFlexibleXconnectService.EntityData.YangName = "vlan-unaware-flexible-xconnect-service"
    vlanUnawareFlexibleXconnectService.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFlexibleXconnectService.EntityData.ParentYangName = "vlan-unaware-flexible-xconnect-services"
    vlanUnawareFlexibleXconnectService.EntityData.SegmentPath = "vlan-unaware-flexible-xconnect-service" + types.AddKeyToken(vlanUnawareFlexibleXconnectService.Name, "name")
    vlanUnawareFlexibleXconnectService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFlexibleXconnectService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFlexibleXconnectService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFlexibleXconnectService.EntityData.Children = types.NewOrderedMap()
    vlanUnawareFlexibleXconnectService.EntityData.Children.Append("vlan-unaware-fxc-attachment-circuits", types.YChild{"VlanUnawareFxcAttachmentCircuits", &vlanUnawareFlexibleXconnectService.VlanUnawareFxcAttachmentCircuits})
    vlanUnawareFlexibleXconnectService.EntityData.Children.Append("vlan-unaware-fxc-pseudowire-evpns", types.YChild{"VlanUnawareFxcPseudowireEvpns", &vlanUnawareFlexibleXconnectService.VlanUnawareFxcPseudowireEvpns})
    vlanUnawareFlexibleXconnectService.EntityData.Leafs = types.NewOrderedMap()
    vlanUnawareFlexibleXconnectService.EntityData.Leafs.Append("name", types.YLeaf{"Name", vlanUnawareFlexibleXconnectService.Name})

    vlanUnawareFlexibleXconnectService.EntityData.YListKeys = []string {"Name"}

    return &(vlanUnawareFlexibleXconnectService.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits
// List of attachment circuits
type L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit.
    VlanUnawareFxcAttachmentCircuit []*L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit
}

func (vlanUnawareFxcAttachmentCircuits *L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcAttachmentCircuits.EntityData.YFilter = vlanUnawareFxcAttachmentCircuits.YFilter
    vlanUnawareFxcAttachmentCircuits.EntityData.YangName = "vlan-unaware-fxc-attachment-circuits"
    vlanUnawareFxcAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcAttachmentCircuits.EntityData.ParentYangName = "vlan-unaware-flexible-xconnect-service"
    vlanUnawareFxcAttachmentCircuits.EntityData.SegmentPath = "vlan-unaware-fxc-attachment-circuits"
    vlanUnawareFxcAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcAttachmentCircuits.EntityData.Children = types.NewOrderedMap()
    vlanUnawareFxcAttachmentCircuits.EntityData.Children.Append("vlan-unaware-fxc-attachment-circuit", types.YChild{"VlanUnawareFxcAttachmentCircuit", nil})
    for i := range vlanUnawareFxcAttachmentCircuits.VlanUnawareFxcAttachmentCircuit {
        vlanUnawareFxcAttachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(vlanUnawareFxcAttachmentCircuits.VlanUnawareFxcAttachmentCircuit[i]), types.YChild{"VlanUnawareFxcAttachmentCircuit", vlanUnawareFxcAttachmentCircuits.VlanUnawareFxcAttachmentCircuit[i]})
    }
    vlanUnawareFxcAttachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    vlanUnawareFxcAttachmentCircuits.EntityData.YListKeys = []string {}

    return &(vlanUnawareFxcAttachmentCircuits.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit
// Attachment circuit interface
type L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: [a-zA-Z0-9._/-]+.
    Name interface{}
}

func (vlanUnawareFxcAttachmentCircuit *L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcAttachmentCircuit.EntityData.YFilter = vlanUnawareFxcAttachmentCircuit.YFilter
    vlanUnawareFxcAttachmentCircuit.EntityData.YangName = "vlan-unaware-fxc-attachment-circuit"
    vlanUnawareFxcAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcAttachmentCircuit.EntityData.ParentYangName = "vlan-unaware-fxc-attachment-circuits"
    vlanUnawareFxcAttachmentCircuit.EntityData.SegmentPath = "vlan-unaware-fxc-attachment-circuit" + types.AddKeyToken(vlanUnawareFxcAttachmentCircuit.Name, "name")
    vlanUnawareFxcAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcAttachmentCircuit.EntityData.Children = types.NewOrderedMap()
    vlanUnawareFxcAttachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    vlanUnawareFxcAttachmentCircuit.EntityData.Leafs.Append("name", types.YLeaf{"Name", vlanUnawareFxcAttachmentCircuit.Name})

    vlanUnawareFxcAttachmentCircuit.EntityData.YListKeys = []string {"Name"}

    return &(vlanUnawareFxcAttachmentCircuit.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns
// List of EVPN Services
type L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN FXC Service Configuration. The type is slice of
    // L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn.
    VlanUnawareFxcPseudowireEvpn []*L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn
}

func (vlanUnawareFxcPseudowireEvpns *L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcPseudowireEvpns.EntityData.YFilter = vlanUnawareFxcPseudowireEvpns.YFilter
    vlanUnawareFxcPseudowireEvpns.EntityData.YangName = "vlan-unaware-fxc-pseudowire-evpns"
    vlanUnawareFxcPseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcPseudowireEvpns.EntityData.ParentYangName = "vlan-unaware-flexible-xconnect-service"
    vlanUnawareFxcPseudowireEvpns.EntityData.SegmentPath = "vlan-unaware-fxc-pseudowire-evpns"
    vlanUnawareFxcPseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcPseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcPseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcPseudowireEvpns.EntityData.Children = types.NewOrderedMap()
    vlanUnawareFxcPseudowireEvpns.EntityData.Children.Append("vlan-unaware-fxc-pseudowire-evpn", types.YChild{"VlanUnawareFxcPseudowireEvpn", nil})
    for i := range vlanUnawareFxcPseudowireEvpns.VlanUnawareFxcPseudowireEvpn {
        vlanUnawareFxcPseudowireEvpns.EntityData.Children.Append(types.GetSegmentPath(vlanUnawareFxcPseudowireEvpns.VlanUnawareFxcPseudowireEvpn[i]), types.YChild{"VlanUnawareFxcPseudowireEvpn", vlanUnawareFxcPseudowireEvpns.VlanUnawareFxcPseudowireEvpn[i]})
    }
    vlanUnawareFxcPseudowireEvpns.EntityData.Leafs = types.NewOrderedMap()

    vlanUnawareFxcPseudowireEvpns.EntityData.YListKeys = []string {}

    return &(vlanUnawareFxcPseudowireEvpns.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn
// EVPN FXC Service Configuration
type L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // This attribute is a key. AC ID. The type is interface{} with range:
    // 1..4294967295.
    Acid interface{}
}

func (vlanUnawareFxcPseudowireEvpn *L2vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcPseudowireEvpn.EntityData.YFilter = vlanUnawareFxcPseudowireEvpn.YFilter
    vlanUnawareFxcPseudowireEvpn.EntityData.YangName = "vlan-unaware-fxc-pseudowire-evpn"
    vlanUnawareFxcPseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcPseudowireEvpn.EntityData.ParentYangName = "vlan-unaware-fxc-pseudowire-evpns"
    vlanUnawareFxcPseudowireEvpn.EntityData.SegmentPath = "vlan-unaware-fxc-pseudowire-evpn" + types.AddKeyToken(vlanUnawareFxcPseudowireEvpn.Eviid, "eviid") + types.AddKeyToken(vlanUnawareFxcPseudowireEvpn.Acid, "acid")
    vlanUnawareFxcPseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcPseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcPseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcPseudowireEvpn.EntityData.Children = types.NewOrderedMap()
    vlanUnawareFxcPseudowireEvpn.EntityData.Leafs = types.NewOrderedMap()
    vlanUnawareFxcPseudowireEvpn.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", vlanUnawareFxcPseudowireEvpn.Eviid})
    vlanUnawareFxcPseudowireEvpn.EntityData.Leafs.Append("acid", types.YLeaf{"Acid", vlanUnawareFxcPseudowireEvpn.Acid})

    vlanUnawareFxcPseudowireEvpn.EntityData.YListKeys = []string {"Eviid", "Acid"}

    return &(vlanUnawareFxcPseudowireEvpn.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices
// List of Vlan-Aware Flexible XConnect Services
type L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flexible XConnect Service. The type is slice of
    // L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService.
    VlanAwareFlexibleXconnectService []*L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService
}

func (vlanAwareFlexibleXconnectServices *L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices) GetEntityData() *types.CommonEntityData {
    vlanAwareFlexibleXconnectServices.EntityData.YFilter = vlanAwareFlexibleXconnectServices.YFilter
    vlanAwareFlexibleXconnectServices.EntityData.YangName = "vlan-aware-flexible-xconnect-services"
    vlanAwareFlexibleXconnectServices.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFlexibleXconnectServices.EntityData.ParentYangName = "flexible-xconnect-service-table"
    vlanAwareFlexibleXconnectServices.EntityData.SegmentPath = "vlan-aware-flexible-xconnect-services"
    vlanAwareFlexibleXconnectServices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFlexibleXconnectServices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFlexibleXconnectServices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFlexibleXconnectServices.EntityData.Children = types.NewOrderedMap()
    vlanAwareFlexibleXconnectServices.EntityData.Children.Append("vlan-aware-flexible-xconnect-service", types.YChild{"VlanAwareFlexibleXconnectService", nil})
    for i := range vlanAwareFlexibleXconnectServices.VlanAwareFlexibleXconnectService {
        vlanAwareFlexibleXconnectServices.EntityData.Children.Append(types.GetSegmentPath(vlanAwareFlexibleXconnectServices.VlanAwareFlexibleXconnectService[i]), types.YChild{"VlanAwareFlexibleXconnectService", vlanAwareFlexibleXconnectServices.VlanAwareFlexibleXconnectService[i]})
    }
    vlanAwareFlexibleXconnectServices.EntityData.Leafs = types.NewOrderedMap()

    vlanAwareFlexibleXconnectServices.EntityData.YListKeys = []string {}

    return &(vlanAwareFlexibleXconnectServices.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService
// Flexible XConnect Service
type L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // List of attachment circuits.
    VlanAwareFxcAttachmentCircuits L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits
}

func (vlanAwareFlexibleXconnectService *L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService) GetEntityData() *types.CommonEntityData {
    vlanAwareFlexibleXconnectService.EntityData.YFilter = vlanAwareFlexibleXconnectService.YFilter
    vlanAwareFlexibleXconnectService.EntityData.YangName = "vlan-aware-flexible-xconnect-service"
    vlanAwareFlexibleXconnectService.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFlexibleXconnectService.EntityData.ParentYangName = "vlan-aware-flexible-xconnect-services"
    vlanAwareFlexibleXconnectService.EntityData.SegmentPath = "vlan-aware-flexible-xconnect-service" + types.AddKeyToken(vlanAwareFlexibleXconnectService.Eviid, "eviid")
    vlanAwareFlexibleXconnectService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFlexibleXconnectService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFlexibleXconnectService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFlexibleXconnectService.EntityData.Children = types.NewOrderedMap()
    vlanAwareFlexibleXconnectService.EntityData.Children.Append("vlan-aware-fxc-attachment-circuits", types.YChild{"VlanAwareFxcAttachmentCircuits", &vlanAwareFlexibleXconnectService.VlanAwareFxcAttachmentCircuits})
    vlanAwareFlexibleXconnectService.EntityData.Leafs = types.NewOrderedMap()
    vlanAwareFlexibleXconnectService.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", vlanAwareFlexibleXconnectService.Eviid})

    vlanAwareFlexibleXconnectService.EntityData.YListKeys = []string {"Eviid"}

    return &(vlanAwareFlexibleXconnectService.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits
// List of attachment circuits
type L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit.
    VlanAwareFxcAttachmentCircuit []*L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit
}

func (vlanAwareFxcAttachmentCircuits *L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    vlanAwareFxcAttachmentCircuits.EntityData.YFilter = vlanAwareFxcAttachmentCircuits.YFilter
    vlanAwareFxcAttachmentCircuits.EntityData.YangName = "vlan-aware-fxc-attachment-circuits"
    vlanAwareFxcAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFxcAttachmentCircuits.EntityData.ParentYangName = "vlan-aware-flexible-xconnect-service"
    vlanAwareFxcAttachmentCircuits.EntityData.SegmentPath = "vlan-aware-fxc-attachment-circuits"
    vlanAwareFxcAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFxcAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFxcAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFxcAttachmentCircuits.EntityData.Children = types.NewOrderedMap()
    vlanAwareFxcAttachmentCircuits.EntityData.Children.Append("vlan-aware-fxc-attachment-circuit", types.YChild{"VlanAwareFxcAttachmentCircuit", nil})
    for i := range vlanAwareFxcAttachmentCircuits.VlanAwareFxcAttachmentCircuit {
        vlanAwareFxcAttachmentCircuits.EntityData.Children.Append(types.GetSegmentPath(vlanAwareFxcAttachmentCircuits.VlanAwareFxcAttachmentCircuit[i]), types.YChild{"VlanAwareFxcAttachmentCircuit", vlanAwareFxcAttachmentCircuits.VlanAwareFxcAttachmentCircuit[i]})
    }
    vlanAwareFxcAttachmentCircuits.EntityData.Leafs = types.NewOrderedMap()

    vlanAwareFxcAttachmentCircuits.EntityData.YListKeys = []string {}

    return &(vlanAwareFxcAttachmentCircuits.EntityData)
}

// L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit
// Attachment circuit interface
type L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: [a-zA-Z0-9._/-]+.
    Name interface{}
}

func (vlanAwareFxcAttachmentCircuit *L2vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    vlanAwareFxcAttachmentCircuit.EntityData.YFilter = vlanAwareFxcAttachmentCircuit.YFilter
    vlanAwareFxcAttachmentCircuit.EntityData.YangName = "vlan-aware-fxc-attachment-circuit"
    vlanAwareFxcAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFxcAttachmentCircuit.EntityData.ParentYangName = "vlan-aware-fxc-attachment-circuits"
    vlanAwareFxcAttachmentCircuit.EntityData.SegmentPath = "vlan-aware-fxc-attachment-circuit" + types.AddKeyToken(vlanAwareFxcAttachmentCircuit.Name, "name")
    vlanAwareFxcAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFxcAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFxcAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFxcAttachmentCircuit.EntityData.Children = types.NewOrderedMap()
    vlanAwareFxcAttachmentCircuit.EntityData.Leafs = types.NewOrderedMap()
    vlanAwareFxcAttachmentCircuit.EntityData.Leafs.Append("name", types.YLeaf{"Name", vlanAwareFxcAttachmentCircuit.Name})

    vlanAwareFxcAttachmentCircuit.EntityData.YListKeys = []string {"Name"}

    return &(vlanAwareFxcAttachmentCircuit.EntityData)
}

// L2vpn_Database_Redundancy
// Redundancy groups
type L2vpn_Database_Redundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable redundancy groups. The type is interface{}.
    Enable interface{}

    // List of Inter-Chassis Communication Protocol redundancy groups.
    IccpRedundancyGroups L2vpn_Database_Redundancy_IccpRedundancyGroups
}

func (redundancy *L2vpn_Database_Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "database"
    redundancy.EntityData.SegmentPath = "redundancy"
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = types.NewOrderedMap()
    redundancy.EntityData.Children.Append("iccp-redundancy-groups", types.YChild{"IccpRedundancyGroups", &redundancy.IccpRedundancyGroups})
    redundancy.EntityData.Leafs = types.NewOrderedMap()
    redundancy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", redundancy.Enable})

    redundancy.EntityData.YListKeys = []string {}

    return &(redundancy.EntityData)
}

// L2vpn_Database_Redundancy_IccpRedundancyGroups
// List of Inter-Chassis Communication Protocol
// redundancy groups
type L2vpn_Database_Redundancy_IccpRedundancyGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICCP Redundancy group. The type is slice of
    // L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup.
    IccpRedundancyGroup []*L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup
}

func (iccpRedundancyGroups *L2vpn_Database_Redundancy_IccpRedundancyGroups) GetEntityData() *types.CommonEntityData {
    iccpRedundancyGroups.EntityData.YFilter = iccpRedundancyGroups.YFilter
    iccpRedundancyGroups.EntityData.YangName = "iccp-redundancy-groups"
    iccpRedundancyGroups.EntityData.BundleName = "cisco_ios_xr"
    iccpRedundancyGroups.EntityData.ParentYangName = "redundancy"
    iccpRedundancyGroups.EntityData.SegmentPath = "iccp-redundancy-groups"
    iccpRedundancyGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpRedundancyGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpRedundancyGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpRedundancyGroups.EntityData.Children = types.NewOrderedMap()
    iccpRedundancyGroups.EntityData.Children.Append("iccp-redundancy-group", types.YChild{"IccpRedundancyGroup", nil})
    for i := range iccpRedundancyGroups.IccpRedundancyGroup {
        iccpRedundancyGroups.EntityData.Children.Append(types.GetSegmentPath(iccpRedundancyGroups.IccpRedundancyGroup[i]), types.YChild{"IccpRedundancyGroup", iccpRedundancyGroups.IccpRedundancyGroup[i]})
    }
    iccpRedundancyGroups.EntityData.Leafs = types.NewOrderedMap()

    iccpRedundancyGroups.EntityData.YListKeys = []string {}

    return &(iccpRedundancyGroups.EntityData)
}

// L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup
// ICCP Redundancy group
type L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..4294967295.
    GroupId interface{}

    // ICCP-based service multi-homing node ID. The type is interface{} with
    // range: 0..254.
    MultiHomingNodeId interface{}

    // List of interfaces.
    IccpInterfaces L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces
}

func (iccpRedundancyGroup *L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup) GetEntityData() *types.CommonEntityData {
    iccpRedundancyGroup.EntityData.YFilter = iccpRedundancyGroup.YFilter
    iccpRedundancyGroup.EntityData.YangName = "iccp-redundancy-group"
    iccpRedundancyGroup.EntityData.BundleName = "cisco_ios_xr"
    iccpRedundancyGroup.EntityData.ParentYangName = "iccp-redundancy-groups"
    iccpRedundancyGroup.EntityData.SegmentPath = "iccp-redundancy-group" + types.AddKeyToken(iccpRedundancyGroup.GroupId, "group-id")
    iccpRedundancyGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpRedundancyGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpRedundancyGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpRedundancyGroup.EntityData.Children = types.NewOrderedMap()
    iccpRedundancyGroup.EntityData.Children.Append("iccp-interfaces", types.YChild{"IccpInterfaces", &iccpRedundancyGroup.IccpInterfaces})
    iccpRedundancyGroup.EntityData.Leafs = types.NewOrderedMap()
    iccpRedundancyGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", iccpRedundancyGroup.GroupId})
    iccpRedundancyGroup.EntityData.Leafs.Append("multi-homing-node-id", types.YLeaf{"MultiHomingNodeId", iccpRedundancyGroup.MultiHomingNodeId})

    iccpRedundancyGroup.EntityData.YListKeys = []string {"GroupId"}

    return &(iccpRedundancyGroup.EntityData)
}

// L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces
// List of interfaces
type L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface.
    IccpInterface []*L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface
}

func (iccpInterfaces *L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces) GetEntityData() *types.CommonEntityData {
    iccpInterfaces.EntityData.YFilter = iccpInterfaces.YFilter
    iccpInterfaces.EntityData.YangName = "iccp-interfaces"
    iccpInterfaces.EntityData.BundleName = "cisco_ios_xr"
    iccpInterfaces.EntityData.ParentYangName = "iccp-redundancy-group"
    iccpInterfaces.EntityData.SegmentPath = "iccp-interfaces"
    iccpInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpInterfaces.EntityData.Children = types.NewOrderedMap()
    iccpInterfaces.EntityData.Children.Append("iccp-interface", types.YChild{"IccpInterface", nil})
    for i := range iccpInterfaces.IccpInterface {
        iccpInterfaces.EntityData.Children.Append(types.GetSegmentPath(iccpInterfaces.IccpInterface[i]), types.YChild{"IccpInterface", iccpInterfaces.IccpInterface[i]})
    }
    iccpInterfaces.EntityData.Leafs = types.NewOrderedMap()

    iccpInterfaces.EntityData.YListKeys = []string {}

    return &(iccpInterfaces.EntityData)
}

// L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface
// Interface name
type L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Secondary VLAN range, in the form of 1-3,5 ,8-11. The type is string.
    SecondaryVlanRange interface{}

    // Failure clear recovery delay. The type is interface{} with range: 30..3600.
    // The default value is 180.
    RecoveryDelay interface{}

    // Primary VLAN range, in the form of 1-3,5 ,8-11. The type is string.
    PrimaryVlanRange interface{}

    // Enable STP-TCN MAC flushing. The type is interface{}.
    MacFlushTcn interface{}
}

func (iccpInterface *L2vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface) GetEntityData() *types.CommonEntityData {
    iccpInterface.EntityData.YFilter = iccpInterface.YFilter
    iccpInterface.EntityData.YangName = "iccp-interface"
    iccpInterface.EntityData.BundleName = "cisco_ios_xr"
    iccpInterface.EntityData.ParentYangName = "iccp-interfaces"
    iccpInterface.EntityData.SegmentPath = "iccp-interface" + types.AddKeyToken(iccpInterface.InterfaceName, "interface-name")
    iccpInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpInterface.EntityData.Children = types.NewOrderedMap()
    iccpInterface.EntityData.Leafs = types.NewOrderedMap()
    iccpInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", iccpInterface.InterfaceName})
    iccpInterface.EntityData.Leafs.Append("secondary-vlan-range", types.YLeaf{"SecondaryVlanRange", iccpInterface.SecondaryVlanRange})
    iccpInterface.EntityData.Leafs.Append("recovery-delay", types.YLeaf{"RecoveryDelay", iccpInterface.RecoveryDelay})
    iccpInterface.EntityData.Leafs.Append("primary-vlan-range", types.YLeaf{"PrimaryVlanRange", iccpInterface.PrimaryVlanRange})
    iccpInterface.EntityData.Leafs.Append("mac-flush-tcn", types.YLeaf{"MacFlushTcn", iccpInterface.MacFlushTcn})

    iccpInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(iccpInterface.EntityData)
}

// L2vpn_Pbb
// L2VPN PBB Global
type L2vpn_Pbb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backbone Source MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BackboneSourceMac interface{}
}

func (pbb *L2vpn_Pbb) GetEntityData() *types.CommonEntityData {
    pbb.EntityData.YFilter = pbb.YFilter
    pbb.EntityData.YangName = "pbb"
    pbb.EntityData.BundleName = "cisco_ios_xr"
    pbb.EntityData.ParentYangName = "l2vpn"
    pbb.EntityData.SegmentPath = "pbb"
    pbb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbb.EntityData.Children = types.NewOrderedMap()
    pbb.EntityData.Leafs = types.NewOrderedMap()
    pbb.EntityData.Leafs.Append("backbone-source-mac", types.YLeaf{"BackboneSourceMac", pbb.BackboneSourceMac})

    pbb.EntityData.YListKeys = []string {}

    return &(pbb.EntityData)
}

// L2vpn_AutoDiscovery
// Global auto-discovery attributes
type L2vpn_AutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global bgp signaling attributes.
    BgpSignaling L2vpn_AutoDiscovery_BgpSignaling
}

func (autoDiscovery *L2vpn_AutoDiscovery) GetEntityData() *types.CommonEntityData {
    autoDiscovery.EntityData.YFilter = autoDiscovery.YFilter
    autoDiscovery.EntityData.YangName = "auto-discovery"
    autoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    autoDiscovery.EntityData.ParentYangName = "l2vpn"
    autoDiscovery.EntityData.SegmentPath = "auto-discovery"
    autoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoDiscovery.EntityData.Children = types.NewOrderedMap()
    autoDiscovery.EntityData.Children.Append("bgp-signaling", types.YChild{"BgpSignaling", &autoDiscovery.BgpSignaling})
    autoDiscovery.EntityData.Leafs = types.NewOrderedMap()

    autoDiscovery.EntityData.YListKeys = []string {}

    return &(autoDiscovery.EntityData)
}

// L2vpn_AutoDiscovery_BgpSignaling
// Global bgp signaling attributes
type L2vpn_AutoDiscovery_BgpSignaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore MTU mismatch for auto-discovered pseudowires. The type is
    // interface{}.
    MtuMismatchIgnore interface{}
}

func (bgpSignaling *L2vpn_AutoDiscovery_BgpSignaling) GetEntityData() *types.CommonEntityData {
    bgpSignaling.EntityData.YFilter = bgpSignaling.YFilter
    bgpSignaling.EntityData.YangName = "bgp-signaling"
    bgpSignaling.EntityData.BundleName = "cisco_ios_xr"
    bgpSignaling.EntityData.ParentYangName = "auto-discovery"
    bgpSignaling.EntityData.SegmentPath = "bgp-signaling"
    bgpSignaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpSignaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpSignaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpSignaling.EntityData.Children = types.NewOrderedMap()
    bgpSignaling.EntityData.Leafs = types.NewOrderedMap()
    bgpSignaling.EntityData.Leafs.Append("mtu-mismatch-ignore", types.YLeaf{"MtuMismatchIgnore", bgpSignaling.MtuMismatchIgnore})

    bgpSignaling.EntityData.YListKeys = []string {}

    return &(bgpSignaling.EntityData)
}

// L2vpn_Utility
// L2VPN utilities
type L2vpn_Utility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN logging utility.
    Logging L2vpn_Utility_Logging
}

func (utility *L2vpn_Utility) GetEntityData() *types.CommonEntityData {
    utility.EntityData.YFilter = utility.YFilter
    utility.EntityData.YangName = "utility"
    utility.EntityData.BundleName = "cisco_ios_xr"
    utility.EntityData.ParentYangName = "l2vpn"
    utility.EntityData.SegmentPath = "utility"
    utility.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utility.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utility.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utility.EntityData.Children = types.NewOrderedMap()
    utility.EntityData.Children.Append("logging", types.YChild{"Logging", &utility.Logging})
    utility.EntityData.Leafs = types.NewOrderedMap()

    utility.EntityData.YListKeys = []string {}

    return &(utility.EntityData)
}

// L2vpn_Utility_Logging
// L2VPN logging utility
type L2vpn_Utility_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Bridge Domain state change logging. The type is interface{}.
    BridgeDomainStateChange interface{}

    // Enable pseudowire state change logging. The type is interface{}.
    PseudowireStateChange interface{}

    // Enable VFI state change logging. The type is interface{}.
    Vfi interface{}

    // Enable Non Stop Routing state change logging. The type is interface{}.
    NsrStateChange interface{}

    // Enable PW-HE Replication state change logging. The type is interface{}.
    PwheReplicationStateChange interface{}
}

func (logging *L2vpn_Utility_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "utility"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("bridge-domain-state-change", types.YLeaf{"BridgeDomainStateChange", logging.BridgeDomainStateChange})
    logging.EntityData.Leafs.Append("pseudowire-state-change", types.YLeaf{"PseudowireStateChange", logging.PseudowireStateChange})
    logging.EntityData.Leafs.Append("vfi", types.YLeaf{"Vfi", logging.Vfi})
    logging.EntityData.Leafs.Append("nsr-state-change", types.YLeaf{"NsrStateChange", logging.NsrStateChange})
    logging.EntityData.Leafs.Append("pwhe-replication-state-change", types.YLeaf{"PwheReplicationStateChange", logging.PwheReplicationStateChange})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// L2vpn_Snmp
// SNMP related configuration
type L2vpn_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MIB related configuration.
    Mib L2vpn_Snmp_Mib
}

func (snmp *L2vpn_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "l2vpn"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Children.Append("mib", types.YChild{"Mib", &snmp.Mib})
    snmp.EntityData.Leafs = types.NewOrderedMap()

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// L2vpn_Snmp_Mib
// MIB related configuration
type L2vpn_Snmp_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface related configuration for MIB.
    MibInterface L2vpn_Snmp_Mib_MibInterface

    // Pseudowire related configuration for MIB.
    MibPseudowire L2vpn_Snmp_Mib_MibPseudowire
}

func (mib *L2vpn_Snmp_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "snmp"
    mib.EntityData.SegmentPath = "mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Children.Append("mib-interface", types.YChild{"MibInterface", &mib.MibInterface})
    mib.EntityData.Children.Append("mib-pseudowire", types.YChild{"MibPseudowire", &mib.MibPseudowire})
    mib.EntityData.Leafs = types.NewOrderedMap()

    mib.EntityData.YListKeys = []string {}

    return &(mib.EntityData)
}

// L2vpn_Snmp_Mib_MibInterface
// Interface related configuration for MIB
type L2vpn_Snmp_Mib_MibInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MIB interface name output format.
    Format L2vpn_Snmp_Mib_MibInterface_Format
}

func (mibInterface *L2vpn_Snmp_Mib_MibInterface) GetEntityData() *types.CommonEntityData {
    mibInterface.EntityData.YFilter = mibInterface.YFilter
    mibInterface.EntityData.YangName = "mib-interface"
    mibInterface.EntityData.BundleName = "cisco_ios_xr"
    mibInterface.EntityData.ParentYangName = "mib"
    mibInterface.EntityData.SegmentPath = "mib-interface"
    mibInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibInterface.EntityData.Children = types.NewOrderedMap()
    mibInterface.EntityData.Children.Append("format", types.YChild{"Format", &mibInterface.Format})
    mibInterface.EntityData.Leafs = types.NewOrderedMap()

    mibInterface.EntityData.YListKeys = []string {}

    return &(mibInterface.EntityData)
}

// L2vpn_Snmp_Mib_MibInterface_Format
// MIB interface name output format
type L2vpn_Snmp_Mib_MibInterface_Format struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set MIB interface name output in slash format (/). The type is interface{}.
    ExternalInterfaceFormat interface{}
}

func (format *L2vpn_Snmp_Mib_MibInterface_Format) GetEntityData() *types.CommonEntityData {
    format.EntityData.YFilter = format.YFilter
    format.EntityData.YangName = "format"
    format.EntityData.BundleName = "cisco_ios_xr"
    format.EntityData.ParentYangName = "mib-interface"
    format.EntityData.SegmentPath = "format"
    format.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    format.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    format.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    format.EntityData.Children = types.NewOrderedMap()
    format.EntityData.Leafs = types.NewOrderedMap()
    format.EntityData.Leafs.Append("external-interface-format", types.YLeaf{"ExternalInterfaceFormat", format.ExternalInterfaceFormat})

    format.EntityData.YListKeys = []string {}

    return &(format.EntityData)
}

// L2vpn_Snmp_Mib_MibPseudowire
// Pseudowire related configuration for MIB
type L2vpn_Snmp_Mib_MibPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pseudowire statistics in MIB output. The type is interface{}.
    Statistics interface{}
}

func (mibPseudowire *L2vpn_Snmp_Mib_MibPseudowire) GetEntityData() *types.CommonEntityData {
    mibPseudowire.EntityData.YFilter = mibPseudowire.YFilter
    mibPseudowire.EntityData.YangName = "mib-pseudowire"
    mibPseudowire.EntityData.BundleName = "cisco_ios_xr"
    mibPseudowire.EntityData.ParentYangName = "mib"
    mibPseudowire.EntityData.SegmentPath = "mib-pseudowire"
    mibPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibPseudowire.EntityData.Children = types.NewOrderedMap()
    mibPseudowire.EntityData.Leafs = types.NewOrderedMap()
    mibPseudowire.EntityData.Leafs.Append("statistics", types.YLeaf{"Statistics", mibPseudowire.Statistics})

    mibPseudowire.EntityData.YListKeys = []string {}

    return &(mibPseudowire.EntityData)
}

// GenericInterfaceLists
// generic interface lists
type GenericInterfaceLists struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic interface list. The type is slice of
    // GenericInterfaceLists_GenericInterfaceList.
    GenericInterfaceList []*GenericInterfaceLists_GenericInterfaceList
}

func (genericInterfaceLists *GenericInterfaceLists) GetEntityData() *types.CommonEntityData {
    genericInterfaceLists.EntityData.YFilter = genericInterfaceLists.YFilter
    genericInterfaceLists.EntityData.YangName = "generic-interface-lists"
    genericInterfaceLists.EntityData.BundleName = "cisco_ios_xr"
    genericInterfaceLists.EntityData.ParentYangName = "Cisco-IOS-XR-l2vpn-cfg"
    genericInterfaceLists.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:generic-interface-lists"
    genericInterfaceLists.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericInterfaceLists.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericInterfaceLists.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericInterfaceLists.EntityData.Children = types.NewOrderedMap()
    genericInterfaceLists.EntityData.Children.Append("generic-interface-list", types.YChild{"GenericInterfaceList", nil})
    for i := range genericInterfaceLists.GenericInterfaceList {
        genericInterfaceLists.EntityData.Children.Append(types.GetSegmentPath(genericInterfaceLists.GenericInterfaceList[i]), types.YChild{"GenericInterfaceList", genericInterfaceLists.GenericInterfaceList[i]})
    }
    genericInterfaceLists.EntityData.Leafs = types.NewOrderedMap()

    genericInterfaceLists.EntityData.YListKeys = []string {}

    return &(genericInterfaceLists.EntityData)
}

// GenericInterfaceLists_GenericInterfaceList
// Generic interface list
type GenericInterfaceLists_GenericInterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface list. The type is string
    // with length: 1..32.
    GenericInterfaceListName interface{}

    // Enable interface list. The type is interface{}.
    Enable interface{}

    // Interface table.
    Interfaces GenericInterfaceLists_GenericInterfaceList_Interfaces
}

func (genericInterfaceList *GenericInterfaceLists_GenericInterfaceList) GetEntityData() *types.CommonEntityData {
    genericInterfaceList.EntityData.YFilter = genericInterfaceList.YFilter
    genericInterfaceList.EntityData.YangName = "generic-interface-list"
    genericInterfaceList.EntityData.BundleName = "cisco_ios_xr"
    genericInterfaceList.EntityData.ParentYangName = "generic-interface-lists"
    genericInterfaceList.EntityData.SegmentPath = "generic-interface-list" + types.AddKeyToken(genericInterfaceList.GenericInterfaceListName, "generic-interface-list-name")
    genericInterfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericInterfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericInterfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericInterfaceList.EntityData.Children = types.NewOrderedMap()
    genericInterfaceList.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &genericInterfaceList.Interfaces})
    genericInterfaceList.EntityData.Leafs = types.NewOrderedMap()
    genericInterfaceList.EntityData.Leafs.Append("generic-interface-list-name", types.YLeaf{"GenericInterfaceListName", genericInterfaceList.GenericInterfaceListName})
    genericInterfaceList.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", genericInterfaceList.Enable})

    genericInterfaceList.EntityData.YListKeys = []string {"GenericInterfaceListName"}

    return &(genericInterfaceList.EntityData)
}

// GenericInterfaceLists_GenericInterfaceList_Interfaces
// Interface table
type GenericInterfaceLists_GenericInterfaceList_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is slice of
    // GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface.
    Interface []*GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface
}

func (interfaces *GenericInterfaceLists_GenericInterfaceList_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "generic-interface-list"
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

// GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface
// Interface
type GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Enable interface. The type is interface{}.
    Enable interface{}
}

func (self *GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Evpn
// evpn
type Evpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable EVPN feature. The type is interface{}.
    Enable interface{}

    // EVPN submodes.
    EvpnTables Evpn_EvpnTables
}

func (evpn *Evpn) GetEntityData() *types.CommonEntityData {
    evpn.EntityData.YFilter = evpn.YFilter
    evpn.EntityData.YangName = "evpn"
    evpn.EntityData.BundleName = "cisco_ios_xr"
    evpn.EntityData.ParentYangName = "Cisco-IOS-XR-l2vpn-cfg"
    evpn.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:evpn"
    evpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpn.EntityData.Children = types.NewOrderedMap()
    evpn.EntityData.Children.Append("evpn-tables", types.YChild{"EvpnTables", &evpn.EvpnTables})
    evpn.EntityData.Leafs = types.NewOrderedMap()
    evpn.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpn.Enable})

    evpn.EntityData.YListKeys = []string {}

    return &(evpn.EntityData)
}

// Evpn_EvpnTables
// EVPN submodes
type Evpn_EvpnTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure node to cost-out. The type is interface{}.
    EviCostOut interface{}

    // Configure EVPN router-id implicitly through Loopback Interface. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    EvpnSourceInterface interface{}

    // Cost-in node after given time (seconds) on startup timer. The type is
    // interface{} with range: 30..86400. Units are second.
    EvpnCostInStartup interface{}

    // Enter EVPN timers configuration submode.
    EvpnTimers Evpn_EvpnTables_EvpnTimers

    // EVPN MAC Configuration.
    Evpnmac Evpn_EvpnTables_Evpnmac

    // Enter EVPN Instance configuration submode.
    EvpnEvis Evpn_EvpnTables_EvpnEvis

    // Virtual Access VFI interfaces.
    EvpnVirtualAccessVfis Evpn_EvpnTables_EvpnVirtualAccessVfis

    // Enter EVPN Loadbalancing configuration submode.
    EvpnLoadBalancing Evpn_EvpnTables_EvpnLoadBalancing

    // Enable Autodiscovery BGP in EVPN.
    EvpnBgpAutoDiscovery Evpn_EvpnTables_EvpnBgpAutoDiscovery

    // Enter EVPN Group Table submode.
    EvpnGroups Evpn_EvpnTables_EvpnGroups

    // Enter EVPN Instance configuration submode.
    EvpnInstances Evpn_EvpnTables_EvpnInstances

    // Enter EVPN Logging configuration submode.
    EvpnLogging Evpn_EvpnTables_EvpnLogging

    // Attachment Circuit interfaces.
    EvpnInterfaces Evpn_EvpnTables_EvpnInterfaces

    // Virtual Access Pseudowire interfaces.
    EvpnVirtualAccessPws Evpn_EvpnTables_EvpnVirtualAccessPws

    // EVPN Global Ethernet Segment submode.
    EvpnEthernetSegment Evpn_EvpnTables_EvpnEthernetSegment
}

func (evpnTables *Evpn_EvpnTables) GetEntityData() *types.CommonEntityData {
    evpnTables.EntityData.YFilter = evpnTables.YFilter
    evpnTables.EntityData.YangName = "evpn-tables"
    evpnTables.EntityData.BundleName = "cisco_ios_xr"
    evpnTables.EntityData.ParentYangName = "evpn"
    evpnTables.EntityData.SegmentPath = "evpn-tables"
    evpnTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnTables.EntityData.Children = types.NewOrderedMap()
    evpnTables.EntityData.Children.Append("evpn-timers", types.YChild{"EvpnTimers", &evpnTables.EvpnTimers})
    evpnTables.EntityData.Children.Append("evpnmac", types.YChild{"Evpnmac", &evpnTables.Evpnmac})
    evpnTables.EntityData.Children.Append("evpn-evis", types.YChild{"EvpnEvis", &evpnTables.EvpnEvis})
    evpnTables.EntityData.Children.Append("evpn-virtual-access-vfis", types.YChild{"EvpnVirtualAccessVfis", &evpnTables.EvpnVirtualAccessVfis})
    evpnTables.EntityData.Children.Append("evpn-load-balancing", types.YChild{"EvpnLoadBalancing", &evpnTables.EvpnLoadBalancing})
    evpnTables.EntityData.Children.Append("evpn-bgp-auto-discovery", types.YChild{"EvpnBgpAutoDiscovery", &evpnTables.EvpnBgpAutoDiscovery})
    evpnTables.EntityData.Children.Append("evpn-groups", types.YChild{"EvpnGroups", &evpnTables.EvpnGroups})
    evpnTables.EntityData.Children.Append("evpn-instances", types.YChild{"EvpnInstances", &evpnTables.EvpnInstances})
    evpnTables.EntityData.Children.Append("evpn-logging", types.YChild{"EvpnLogging", &evpnTables.EvpnLogging})
    evpnTables.EntityData.Children.Append("evpn-interfaces", types.YChild{"EvpnInterfaces", &evpnTables.EvpnInterfaces})
    evpnTables.EntityData.Children.Append("evpn-virtual-access-pws", types.YChild{"EvpnVirtualAccessPws", &evpnTables.EvpnVirtualAccessPws})
    evpnTables.EntityData.Children.Append("evpn-ethernet-segment", types.YChild{"EvpnEthernetSegment", &evpnTables.EvpnEthernetSegment})
    evpnTables.EntityData.Leafs = types.NewOrderedMap()
    evpnTables.EntityData.Leafs.Append("evi-cost-out", types.YLeaf{"EviCostOut", evpnTables.EviCostOut})
    evpnTables.EntityData.Leafs.Append("evpn-source-interface", types.YLeaf{"EvpnSourceInterface", evpnTables.EvpnSourceInterface})
    evpnTables.EntityData.Leafs.Append("evpn-cost-in-startup", types.YLeaf{"EvpnCostInStartup", evpnTables.EvpnCostInStartup})

    evpnTables.EntityData.YListKeys = []string {}

    return &(evpnTables.EntityData)
}

// Evpn_EvpnTables_EvpnTimers
// Enter EVPN timers configuration submode
type Evpn_EvpnTables_EvpnTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global Carving timer. The type is interface{} with range: 0..10. The
    // default value is 0.
    EvpnCarving interface{}

    // Global Recovery timer. The type is interface{} with range: 20..3600. The
    // default value is 30.
    EvpnRecovery interface{}

    // Enable EVPN timers. The type is interface{}.
    Enable interface{}

    // Global Peering timer. The type is interface{} with range: 0..300. The
    // default value is 3.
    EvpnPeering interface{}
}

func (evpnTimers *Evpn_EvpnTables_EvpnTimers) GetEntityData() *types.CommonEntityData {
    evpnTimers.EntityData.YFilter = evpnTimers.YFilter
    evpnTimers.EntityData.YangName = "evpn-timers"
    evpnTimers.EntityData.BundleName = "cisco_ios_xr"
    evpnTimers.EntityData.ParentYangName = "evpn-tables"
    evpnTimers.EntityData.SegmentPath = "evpn-timers"
    evpnTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnTimers.EntityData.Children = types.NewOrderedMap()
    evpnTimers.EntityData.Leafs = types.NewOrderedMap()
    evpnTimers.EntityData.Leafs.Append("evpn-carving", types.YLeaf{"EvpnCarving", evpnTimers.EvpnCarving})
    evpnTimers.EntityData.Leafs.Append("evpn-recovery", types.YLeaf{"EvpnRecovery", evpnTimers.EvpnRecovery})
    evpnTimers.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnTimers.Enable})
    evpnTimers.EntityData.Leafs.Append("evpn-peering", types.YLeaf{"EvpnPeering", evpnTimers.EvpnPeering})

    evpnTimers.EntityData.YListKeys = []string {}

    return &(evpnTimers.EntityData)
}

// Evpn_EvpnTables_Evpnmac
// EVPN MAC Configuration
type Evpn_EvpnTables_Evpnmac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable EVPN MAC Configuration. The type is interface{}.
    Enable interface{}

    // EVPN MAC Secure Configuration.
    EvpnmacSecure Evpn_EvpnTables_Evpnmac_EvpnmacSecure
}

func (evpnmac *Evpn_EvpnTables_Evpnmac) GetEntityData() *types.CommonEntityData {
    evpnmac.EntityData.YFilter = evpnmac.YFilter
    evpnmac.EntityData.YangName = "evpnmac"
    evpnmac.EntityData.BundleName = "cisco_ios_xr"
    evpnmac.EntityData.ParentYangName = "evpn-tables"
    evpnmac.EntityData.SegmentPath = "evpnmac"
    evpnmac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnmac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnmac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnmac.EntityData.Children = types.NewOrderedMap()
    evpnmac.EntityData.Children.Append("evpnmac-secure", types.YChild{"EvpnmacSecure", &evpnmac.EvpnmacSecure})
    evpnmac.EntityData.Leafs = types.NewOrderedMap()
    evpnmac.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnmac.Enable})

    evpnmac.EntityData.YListKeys = []string {}

    return &(evpnmac.EntityData)
}

// Evpn_EvpnTables_Evpnmac_EvpnmacSecure
// EVPN MAC Secure Configuration
type Evpn_EvpnTables_Evpnmac_EvpnmacSecure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Length of time to lock the MAC after a MAC security violation. The type is
    // interface{} with range: 5..3600.
    EvpnmacSecureFreezeTime interface{}

    // Enable EVPN MAC Secure Configuration. The type is interface{}.
    Enable interface{}

    // Number of times to unfreeze a MAC before permanently freezing it. The type
    // is interface{} with range: 0..1000.
    EvpnmacSecureRetryCount interface{}

    // Number of moves to occur within the move interval before locking the MAC.
    // The type is interface{} with range: 1..1000.
    EvpnmacSecureMoveCount interface{}

    // Interval to watch for subsequent MAC moves before locking the MAC. The type
    // is interface{} with range: 5..3600.
    EvpnmacSecureMoveInterval interface{}
}

func (evpnmacSecure *Evpn_EvpnTables_Evpnmac_EvpnmacSecure) GetEntityData() *types.CommonEntityData {
    evpnmacSecure.EntityData.YFilter = evpnmacSecure.YFilter
    evpnmacSecure.EntityData.YangName = "evpnmac-secure"
    evpnmacSecure.EntityData.BundleName = "cisco_ios_xr"
    evpnmacSecure.EntityData.ParentYangName = "evpnmac"
    evpnmacSecure.EntityData.SegmentPath = "evpnmac-secure"
    evpnmacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnmacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnmacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnmacSecure.EntityData.Children = types.NewOrderedMap()
    evpnmacSecure.EntityData.Leafs = types.NewOrderedMap()
    evpnmacSecure.EntityData.Leafs.Append("evpnmac-secure-freeze-time", types.YLeaf{"EvpnmacSecureFreezeTime", evpnmacSecure.EvpnmacSecureFreezeTime})
    evpnmacSecure.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnmacSecure.Enable})
    evpnmacSecure.EntityData.Leafs.Append("evpnmac-secure-retry-count", types.YLeaf{"EvpnmacSecureRetryCount", evpnmacSecure.EvpnmacSecureRetryCount})
    evpnmacSecure.EntityData.Leafs.Append("evpnmac-secure-move-count", types.YLeaf{"EvpnmacSecureMoveCount", evpnmacSecure.EvpnmacSecureMoveCount})
    evpnmacSecure.EntityData.Leafs.Append("evpnmac-secure-move-interval", types.YLeaf{"EvpnmacSecureMoveInterval", evpnmacSecure.EvpnmacSecureMoveInterval})

    evpnmacSecure.EntityData.YListKeys = []string {}

    return &(evpnmacSecure.EntityData)
}

// Evpn_EvpnTables_EvpnEvis
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_EvpnEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter EVPN Instance configuration submode. The type is slice of
    // Evpn_EvpnTables_EvpnEvis_EvpnEvi.
    EvpnEvi []*Evpn_EvpnTables_EvpnEvis_EvpnEvi
}

func (evpnEvis *Evpn_EvpnTables_EvpnEvis) GetEntityData() *types.CommonEntityData {
    evpnEvis.EntityData.YFilter = evpnEvis.YFilter
    evpnEvis.EntityData.YangName = "evpn-evis"
    evpnEvis.EntityData.BundleName = "cisco_ios_xr"
    evpnEvis.EntityData.ParentYangName = "evpn-tables"
    evpnEvis.EntityData.SegmentPath = "evpn-evis"
    evpnEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEvis.EntityData.Children = types.NewOrderedMap()
    evpnEvis.EntityData.Children.Append("evpn-evi", types.YChild{"EvpnEvi", nil})
    for i := range evpnEvis.EvpnEvi {
        evpnEvis.EntityData.Children.Append(types.GetSegmentPath(evpnEvis.EvpnEvi[i]), types.YChild{"EvpnEvi", evpnEvis.EvpnEvi[i]})
    }
    evpnEvis.EntityData.Leafs = types.NewOrderedMap()

    evpnEvis.EntityData.YListKeys = []string {}

    return &(evpnEvis.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_EvpnEvis_EvpnEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVI ID. The type is interface{} with range:
    // 1..65534.
    Eviid interface{}

    // Disable route re-origination. The type is interface{}.
    EviReorigDisable interface{}

    // DEPRECATED: Advertise local MAC-only and BVI MAC routes. The type is
    // interface{}.
    EviAdvertiseMacDeprecated interface{}

    // EVPN Instance description. The type is string with length: 1..64.
    EvpnEviDescription interface{}

    // Disable ECMP on the EVI. The type is interface{}.
    EviEcmpDisable interface{}

    // Disable Unknown Unicast Flooding on this EVI. The type is interface{}.
    EviUnknownUnicastFloodingDisable interface{}

    // CW disable for EVPN EVI. The type is interface{}.
    EvpnEviCwDisable interface{}

    // Enter Loadbalancing configuration submode.
    EviLoadBalancing Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviLoadBalancing

    // Enable Autodiscovery BGP in EVPN Instance.
    EvpnEviBgpAutoDiscovery Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery

    // Enter Multicast configuration submode.
    EviMulticast Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviMulticast

    // Enter Advertise local MAC-only routes configuration submode.
    EviAdvertiseMac Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviAdvertiseMac
}

func (evpnEvi *Evpn_EvpnTables_EvpnEvis_EvpnEvi) GetEntityData() *types.CommonEntityData {
    evpnEvi.EntityData.YFilter = evpnEvi.YFilter
    evpnEvi.EntityData.YangName = "evpn-evi"
    evpnEvi.EntityData.BundleName = "cisco_ios_xr"
    evpnEvi.EntityData.ParentYangName = "evpn-evis"
    evpnEvi.EntityData.SegmentPath = "evpn-evi" + types.AddKeyToken(evpnEvi.Eviid, "eviid")
    evpnEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEvi.EntityData.Children = types.NewOrderedMap()
    evpnEvi.EntityData.Children.Append("evi-load-balancing", types.YChild{"EviLoadBalancing", &evpnEvi.EviLoadBalancing})
    evpnEvi.EntityData.Children.Append("evpn-evi-bgp-auto-discovery", types.YChild{"EvpnEviBgpAutoDiscovery", &evpnEvi.EvpnEviBgpAutoDiscovery})
    evpnEvi.EntityData.Children.Append("evi-multicast", types.YChild{"EviMulticast", &evpnEvi.EviMulticast})
    evpnEvi.EntityData.Children.Append("evi-advertise-mac", types.YChild{"EviAdvertiseMac", &evpnEvi.EviAdvertiseMac})
    evpnEvi.EntityData.Leafs = types.NewOrderedMap()
    evpnEvi.EntityData.Leafs.Append("eviid", types.YLeaf{"Eviid", evpnEvi.Eviid})
    evpnEvi.EntityData.Leafs.Append("evi-reorig-disable", types.YLeaf{"EviReorigDisable", evpnEvi.EviReorigDisable})
    evpnEvi.EntityData.Leafs.Append("evi-advertise-mac-deprecated", types.YLeaf{"EviAdvertiseMacDeprecated", evpnEvi.EviAdvertiseMacDeprecated})
    evpnEvi.EntityData.Leafs.Append("evpn-evi-description", types.YLeaf{"EvpnEviDescription", evpnEvi.EvpnEviDescription})
    evpnEvi.EntityData.Leafs.Append("evi-ecmp-disable", types.YLeaf{"EviEcmpDisable", evpnEvi.EviEcmpDisable})
    evpnEvi.EntityData.Leafs.Append("evi-unknown-unicast-flooding-disable", types.YLeaf{"EviUnknownUnicastFloodingDisable", evpnEvi.EviUnknownUnicastFloodingDisable})
    evpnEvi.EntityData.Leafs.Append("evpn-evi-cw-disable", types.YLeaf{"EvpnEviCwDisable", evpnEvi.EvpnEviCwDisable})

    evpnEvi.EntityData.YListKeys = []string {"Eviid"}

    return &(evpnEvi.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviLoadBalancing
// Enter Loadbalancing configuration submode
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviLoadBalancing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Loadbalancing. The type is interface{}.
    Enable interface{}

    // Enable Static Flow Label based load balancing. The type is interface{}.
    EviStaticFlowLabel interface{}
}

func (eviLoadBalancing *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviLoadBalancing) GetEntityData() *types.CommonEntityData {
    eviLoadBalancing.EntityData.YFilter = eviLoadBalancing.YFilter
    eviLoadBalancing.EntityData.YangName = "evi-load-balancing"
    eviLoadBalancing.EntityData.BundleName = "cisco_ios_xr"
    eviLoadBalancing.EntityData.ParentYangName = "evpn-evi"
    eviLoadBalancing.EntityData.SegmentPath = "evi-load-balancing"
    eviLoadBalancing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviLoadBalancing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviLoadBalancing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviLoadBalancing.EntityData.Children = types.NewOrderedMap()
    eviLoadBalancing.EntityData.Leafs = types.NewOrderedMap()
    eviLoadBalancing.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", eviLoadBalancing.Enable})
    eviLoadBalancing.EntityData.Leafs.Append("evi-static-flow-label", types.YLeaf{"EviStaticFlowLabel", eviLoadBalancing.EviStaticFlowLabel})

    eviLoadBalancing.EntityData.YListKeys = []string {}

    return &(eviLoadBalancing.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery
// Enable Autodiscovery BGP in EVPN Instance
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Table Policy for installation of forwarding data to L2FIB. The type is
    // string.
    TablePolicy interface{}

    // Route Distinguisher.
    EvpnRouteDistinguisher Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteDistinguisher

    // Route Target.
    EvpnRouteTargets Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets
}

func (evpnEviBgpAutoDiscovery *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    evpnEviBgpAutoDiscovery.EntityData.YFilter = evpnEviBgpAutoDiscovery.YFilter
    evpnEviBgpAutoDiscovery.EntityData.YangName = "evpn-evi-bgp-auto-discovery"
    evpnEviBgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    evpnEviBgpAutoDiscovery.EntityData.ParentYangName = "evpn-evi"
    evpnEviBgpAutoDiscovery.EntityData.SegmentPath = "evpn-evi-bgp-auto-discovery"
    evpnEviBgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEviBgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEviBgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEviBgpAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    evpnEviBgpAutoDiscovery.EntityData.Children.Append("evpn-route-distinguisher", types.YChild{"EvpnRouteDistinguisher", &evpnEviBgpAutoDiscovery.EvpnRouteDistinguisher})
    evpnEviBgpAutoDiscovery.EntityData.Children.Append("evpn-route-targets", types.YChild{"EvpnRouteTargets", &evpnEviBgpAutoDiscovery.EvpnRouteTargets})
    evpnEviBgpAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    evpnEviBgpAutoDiscovery.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnEviBgpAutoDiscovery.Enable})
    evpnEviBgpAutoDiscovery.EntityData.Leafs.Append("table-policy", types.YLeaf{"TablePolicy", evpnEviBgpAutoDiscovery.TablePolicy})

    evpnEviBgpAutoDiscovery.EntityData.YListKeys = []string {}

    return &(evpnEviBgpAutoDiscovery.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteDistinguisher
// Route Distinguisher
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (evpnRouteDistinguisher *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteDistinguisher) GetEntityData() *types.CommonEntityData {
    evpnRouteDistinguisher.EntityData.YFilter = evpnRouteDistinguisher.YFilter
    evpnRouteDistinguisher.EntityData.YangName = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteDistinguisher.EntityData.ParentYangName = "evpn-evi-bgp-auto-discovery"
    evpnRouteDistinguisher.EntityData.SegmentPath = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteDistinguisher.EntityData.Children = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnRouteDistinguisher.Type})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", evpnRouteDistinguisher.As})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex})
    evpnRouteDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", evpnRouteDistinguisher.Address})
    evpnRouteDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex})

    evpnRouteDistinguisher.EntityData.YListKeys = []string {}

    return &(evpnRouteDistinguisher.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets
// Route Target
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs.
    EvpnRouteTargetAs []*Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone.
    EvpnRouteTargetNone []*Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address.
    EvpnRouteTargetIpv4Address []*Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
}

func (evpnRouteTargets *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets) GetEntityData() *types.CommonEntityData {
    evpnRouteTargets.EntityData.YFilter = evpnRouteTargets.YFilter
    evpnRouteTargets.EntityData.YangName = "evpn-route-targets"
    evpnRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargets.EntityData.ParentYangName = "evpn-evi-bgp-auto-discovery"
    evpnRouteTargets.EntityData.SegmentPath = "evpn-route-targets"
    evpnRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargets.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargets.EntityData.Children.Append("evpn-route-target-as", types.YChild{"EvpnRouteTargetAs", nil})
    for i := range evpnRouteTargets.EvpnRouteTargetAs {
        evpnRouteTargets.EntityData.Children.Append(types.GetSegmentPath(evpnRouteTargets.EvpnRouteTargetAs[i]), types.YChild{"EvpnRouteTargetAs", evpnRouteTargets.EvpnRouteTargetAs[i]})
    }
    evpnRouteTargets.EntityData.Children.Append("evpn-route-target-none", types.YChild{"EvpnRouteTargetNone", nil})
    for i := range evpnRouteTargets.EvpnRouteTargetNone {
        evpnRouteTargets.EntityData.Children.Append(types.GetSegmentPath(evpnRouteTargets.EvpnRouteTargetNone[i]), types.YChild{"EvpnRouteTargetNone", evpnRouteTargets.EvpnRouteTargetNone[i]})
    }
    evpnRouteTargets.EntityData.Children.Append("evpn-route-target-ipv4-address", types.YChild{"EvpnRouteTargetIpv4Address", nil})
    for i := range evpnRouteTargets.EvpnRouteTargetIpv4Address {
        evpnRouteTargets.EntityData.Children.Append(types.GetSegmentPath(evpnRouteTargets.EvpnRouteTargetIpv4Address[i]), types.YChild{"EvpnRouteTargetIpv4Address", evpnRouteTargets.EvpnRouteTargetIpv4Address[i]})
    }
    evpnRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    evpnRouteTargets.EntityData.YListKeys = []string {}

    return &(evpnRouteTargets.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs
// Name of the Route Target
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetAs *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetAs.EntityData.YFilter = evpnRouteTargetAs.YFilter
    evpnRouteTargetAs.EntityData.YangName = "evpn-route-target-as"
    evpnRouteTargetAs.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetAs.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetAs.EntityData.SegmentPath = "evpn-route-target-as" + types.AddKeyToken(evpnRouteTargetAs.Format, "format") + types.AddKeyToken(evpnRouteTargetAs.Role, "role") + types.AddKeyToken(evpnRouteTargetAs.As, "as") + types.AddKeyToken(evpnRouteTargetAs.AsIndex, "as-index") + types.AddKeyToken(evpnRouteTargetAs.Stitching, "stitching")
    evpnRouteTargetAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetAs.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargetAs.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteTargetAs.EntityData.Leafs.Append("format", types.YLeaf{"Format", evpnRouteTargetAs.Format})
    evpnRouteTargetAs.EntityData.Leafs.Append("role", types.YLeaf{"Role", evpnRouteTargetAs.Role})
    evpnRouteTargetAs.EntityData.Leafs.Append("as", types.YLeaf{"As", evpnRouteTargetAs.As})
    evpnRouteTargetAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", evpnRouteTargetAs.AsIndex})
    evpnRouteTargetAs.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", evpnRouteTargetAs.Stitching})

    evpnRouteTargetAs.EntityData.YListKeys = []string {"Format", "Role", "As", "AsIndex", "Stitching"}

    return &(evpnRouteTargetAs.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone
// Name of the Route Target
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetNone *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetNone.EntityData.YFilter = evpnRouteTargetNone.YFilter
    evpnRouteTargetNone.EntityData.YangName = "evpn-route-target-none"
    evpnRouteTargetNone.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetNone.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetNone.EntityData.SegmentPath = "evpn-route-target-none" + types.AddKeyToken(evpnRouteTargetNone.Format, "format") + types.AddKeyToken(evpnRouteTargetNone.Role, "role") + types.AddKeyToken(evpnRouteTargetNone.Stitching, "stitching")
    evpnRouteTargetNone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetNone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetNone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetNone.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargetNone.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteTargetNone.EntityData.Leafs.Append("format", types.YLeaf{"Format", evpnRouteTargetNone.Format})
    evpnRouteTargetNone.EntityData.Leafs.Append("role", types.YLeaf{"Role", evpnRouteTargetNone.Role})
    evpnRouteTargetNone.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", evpnRouteTargetNone.Stitching})

    evpnRouteTargetNone.EntityData.YListKeys = []string {"Format", "Role", "Stitching"}

    return &(evpnRouteTargetNone.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
// Name of the Route Target
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetIpv4Address *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EvpnEviBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetIpv4Address.EntityData.YFilter = evpnRouteTargetIpv4Address.YFilter
    evpnRouteTargetIpv4Address.EntityData.YangName = "evpn-route-target-ipv4-address"
    evpnRouteTargetIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetIpv4Address.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetIpv4Address.EntityData.SegmentPath = "evpn-route-target-ipv4-address" + types.AddKeyToken(evpnRouteTargetIpv4Address.Format, "format") + types.AddKeyToken(evpnRouteTargetIpv4Address.Role, "role") + types.AddKeyToken(evpnRouteTargetIpv4Address.Address, "address") + types.AddKeyToken(evpnRouteTargetIpv4Address.AddrIndex, "addr-index") + types.AddKeyToken(evpnRouteTargetIpv4Address.Stitching, "stitching")
    evpnRouteTargetIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetIpv4Address.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargetIpv4Address.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("format", types.YLeaf{"Format", evpnRouteTargetIpv4Address.Format})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("role", types.YLeaf{"Role", evpnRouteTargetIpv4Address.Role})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", evpnRouteTargetIpv4Address.Address})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", evpnRouteTargetIpv4Address.AddrIndex})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", evpnRouteTargetIpv4Address.Stitching})

    evpnRouteTargetIpv4Address.EntityData.YListKeys = []string {"Format", "Role", "Address", "AddrIndex", "Stitching"}

    return &(evpnRouteTargetIpv4Address.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviMulticast
// Enter Multicast configuration submode
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Multicast. The type is interface{}.
    Enable interface{}

    // Enable Multicast source connectivity. The type is interface{}.
    EviMcastSourceConnected interface{}
}

func (eviMulticast *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviMulticast) GetEntityData() *types.CommonEntityData {
    eviMulticast.EntityData.YFilter = eviMulticast.YFilter
    eviMulticast.EntityData.YangName = "evi-multicast"
    eviMulticast.EntityData.BundleName = "cisco_ios_xr"
    eviMulticast.EntityData.ParentYangName = "evpn-evi"
    eviMulticast.EntityData.SegmentPath = "evi-multicast"
    eviMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviMulticast.EntityData.Children = types.NewOrderedMap()
    eviMulticast.EntityData.Leafs = types.NewOrderedMap()
    eviMulticast.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", eviMulticast.Enable})
    eviMulticast.EntityData.Leafs.Append("evi-mcast-source-connected", types.YLeaf{"EviMcastSourceConnected", eviMulticast.EviMcastSourceConnected})

    eviMulticast.EntityData.YListKeys = []string {}

    return &(eviMulticast.EntityData)
}

// Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviAdvertiseMac
// Enter Advertise local MAC-only routes
// configuration submode
type Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviAdvertiseMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Advertise local MAC-only routes. The type is interface{}.
    Enable interface{}

    // Advertise local MAC-only and BVI MAC routes. The type is interface{}.
    EviAdvertiseMacBvi interface{}
}

func (eviAdvertiseMac *Evpn_EvpnTables_EvpnEvis_EvpnEvi_EviAdvertiseMac) GetEntityData() *types.CommonEntityData {
    eviAdvertiseMac.EntityData.YFilter = eviAdvertiseMac.YFilter
    eviAdvertiseMac.EntityData.YangName = "evi-advertise-mac"
    eviAdvertiseMac.EntityData.BundleName = "cisco_ios_xr"
    eviAdvertiseMac.EntityData.ParentYangName = "evpn-evi"
    eviAdvertiseMac.EntityData.SegmentPath = "evi-advertise-mac"
    eviAdvertiseMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviAdvertiseMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviAdvertiseMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviAdvertiseMac.EntityData.Children = types.NewOrderedMap()
    eviAdvertiseMac.EntityData.Leafs = types.NewOrderedMap()
    eviAdvertiseMac.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", eviAdvertiseMac.Enable})
    eviAdvertiseMac.EntityData.Leafs.Append("evi-advertise-mac-bvi", types.YLeaf{"EviAdvertiseMacBvi", eviAdvertiseMac.EviAdvertiseMacBvi})

    eviAdvertiseMac.EntityData.YListKeys = []string {}

    return &(eviAdvertiseMac.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis
// Virtual Access VFI interfaces
type Evpn_EvpnTables_EvpnVirtualAccessVfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual Access VFI. The type is slice of
    // Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi.
    EvpnVirtualAccessVfi []*Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi
}

func (evpnVirtualAccessVfis *Evpn_EvpnTables_EvpnVirtualAccessVfis) GetEntityData() *types.CommonEntityData {
    evpnVirtualAccessVfis.EntityData.YFilter = evpnVirtualAccessVfis.YFilter
    evpnVirtualAccessVfis.EntityData.YangName = "evpn-virtual-access-vfis"
    evpnVirtualAccessVfis.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualAccessVfis.EntityData.ParentYangName = "evpn-tables"
    evpnVirtualAccessVfis.EntityData.SegmentPath = "evpn-virtual-access-vfis"
    evpnVirtualAccessVfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessVfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessVfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessVfis.EntityData.Children = types.NewOrderedMap()
    evpnVirtualAccessVfis.EntityData.Children.Append("evpn-virtual-access-vfi", types.YChild{"EvpnVirtualAccessVfi", nil})
    for i := range evpnVirtualAccessVfis.EvpnVirtualAccessVfi {
        evpnVirtualAccessVfis.EntityData.Children.Append(types.GetSegmentPath(evpnVirtualAccessVfis.EvpnVirtualAccessVfi[i]), types.YChild{"EvpnVirtualAccessVfi", evpnVirtualAccessVfis.EvpnVirtualAccessVfi[i]})
    }
    evpnVirtualAccessVfis.EntityData.Leafs = types.NewOrderedMap()

    evpnVirtualAccessVfis.EntityData.YListKeys = []string {}

    return &(evpnVirtualAccessVfis.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi
// Virtual Access VFI
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Virtual Access VFI. The type is string
    // with length: 1..32.
    Name interface{}

    // Enter Virtual Forwarding Interface timers configuration submode.
    EvpnVirtualAccessVfiTimers Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualAccessVfiTimers

    // Enter Ethernet Segment configuration submode.
    EvpnVirtualEthernetSegment Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment
}

func (evpnVirtualAccessVfi *Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi) GetEntityData() *types.CommonEntityData {
    evpnVirtualAccessVfi.EntityData.YFilter = evpnVirtualAccessVfi.YFilter
    evpnVirtualAccessVfi.EntityData.YangName = "evpn-virtual-access-vfi"
    evpnVirtualAccessVfi.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualAccessVfi.EntityData.ParentYangName = "evpn-virtual-access-vfis"
    evpnVirtualAccessVfi.EntityData.SegmentPath = "evpn-virtual-access-vfi" + types.AddKeyToken(evpnVirtualAccessVfi.Name, "name")
    evpnVirtualAccessVfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessVfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessVfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessVfi.EntityData.Children = types.NewOrderedMap()
    evpnVirtualAccessVfi.EntityData.Children.Append("evpn-virtual-access-vfi-timers", types.YChild{"EvpnVirtualAccessVfiTimers", &evpnVirtualAccessVfi.EvpnVirtualAccessVfiTimers})
    evpnVirtualAccessVfi.EntityData.Children.Append("evpn-virtual-ethernet-segment", types.YChild{"EvpnVirtualEthernetSegment", &evpnVirtualAccessVfi.EvpnVirtualEthernetSegment})
    evpnVirtualAccessVfi.EntityData.Leafs = types.NewOrderedMap()
    evpnVirtualAccessVfi.EntityData.Leafs.Append("name", types.YLeaf{"Name", evpnVirtualAccessVfi.Name})

    evpnVirtualAccessVfi.EntityData.YListKeys = []string {"Name"}

    return &(evpnVirtualAccessVfi.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualAccessVfiTimers
// Enter Virtual Forwarding Interface timers
// configuration submode
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualAccessVfiTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual Forwarding Interface-specific Recovery timer. The type is
    // interface{} with range: 20..3600. The default value is 30.
    EvpnVirtualAccessVfiRecovery interface{}

    // Virtual Forwarding Interface-specific Peering timer. The type is
    // interface{} with range: 0..300. The default value is 3.
    EvpnVirtualAccessVfiPeering interface{}

    // Virtual Forwarding Interface-specific Carving timer. The type is
    // interface{} with range: 0..10. The default value is 0.
    EvpnVirtualAccessVfiCarving interface{}

    // Enable Virtual Forwarding Interface timers. The type is interface{}.
    Enable interface{}
}

func (evpnVirtualAccessVfiTimers *Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualAccessVfiTimers) GetEntityData() *types.CommonEntityData {
    evpnVirtualAccessVfiTimers.EntityData.YFilter = evpnVirtualAccessVfiTimers.YFilter
    evpnVirtualAccessVfiTimers.EntityData.YangName = "evpn-virtual-access-vfi-timers"
    evpnVirtualAccessVfiTimers.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualAccessVfiTimers.EntityData.ParentYangName = "evpn-virtual-access-vfi"
    evpnVirtualAccessVfiTimers.EntityData.SegmentPath = "evpn-virtual-access-vfi-timers"
    evpnVirtualAccessVfiTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessVfiTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessVfiTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessVfiTimers.EntityData.Children = types.NewOrderedMap()
    evpnVirtualAccessVfiTimers.EntityData.Leafs = types.NewOrderedMap()
    evpnVirtualAccessVfiTimers.EntityData.Leafs.Append("evpn-virtual-access-vfi-recovery", types.YLeaf{"EvpnVirtualAccessVfiRecovery", evpnVirtualAccessVfiTimers.EvpnVirtualAccessVfiRecovery})
    evpnVirtualAccessVfiTimers.EntityData.Leafs.Append("evpn-virtual-access-vfi-peering", types.YLeaf{"EvpnVirtualAccessVfiPeering", evpnVirtualAccessVfiTimers.EvpnVirtualAccessVfiPeering})
    evpnVirtualAccessVfiTimers.EntityData.Leafs.Append("evpn-virtual-access-vfi-carving", types.YLeaf{"EvpnVirtualAccessVfiCarving", evpnVirtualAccessVfiTimers.EvpnVirtualAccessVfiCarving})
    evpnVirtualAccessVfiTimers.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnVirtualAccessVfiTimers.Enable})

    evpnVirtualAccessVfiTimers.EntityData.YListKeys = []string {}

    return &(evpnVirtualAccessVfiTimers.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment
// Enter Ethernet Segment configuration submode
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Ethernet Segment. The type is interface{}.
    Enable interface{}

    // ES-Import Route Target. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    EsImportRouteTarget interface{}

    // Ethernet-Segment Service Carving mode. The type is
    // EthernetSegmentServiceCarving.
    ServiceCarvingType interface{}

    // Ethernet segment identifier.
    Identifier Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_Identifier

    // Enter Manual service carving configuration submode.
    ManualServiceCarving Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving
}

func (evpnVirtualEthernetSegment *Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment) GetEntityData() *types.CommonEntityData {
    evpnVirtualEthernetSegment.EntityData.YFilter = evpnVirtualEthernetSegment.YFilter
    evpnVirtualEthernetSegment.EntityData.YangName = "evpn-virtual-ethernet-segment"
    evpnVirtualEthernetSegment.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualEthernetSegment.EntityData.ParentYangName = "evpn-virtual-access-vfi"
    evpnVirtualEthernetSegment.EntityData.SegmentPath = "evpn-virtual-ethernet-segment"
    evpnVirtualEthernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualEthernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualEthernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualEthernetSegment.EntityData.Children = types.NewOrderedMap()
    evpnVirtualEthernetSegment.EntityData.Children.Append("identifier", types.YChild{"Identifier", &evpnVirtualEthernetSegment.Identifier})
    evpnVirtualEthernetSegment.EntityData.Children.Append("manual-service-carving", types.YChild{"ManualServiceCarving", &evpnVirtualEthernetSegment.ManualServiceCarving})
    evpnVirtualEthernetSegment.EntityData.Leafs = types.NewOrderedMap()
    evpnVirtualEthernetSegment.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnVirtualEthernetSegment.Enable})
    evpnVirtualEthernetSegment.EntityData.Leafs.Append("es-import-route-target", types.YLeaf{"EsImportRouteTarget", evpnVirtualEthernetSegment.EsImportRouteTarget})
    evpnVirtualEthernetSegment.EntityData.Leafs.Append("service-carving-type", types.YLeaf{"ServiceCarvingType", evpnVirtualEthernetSegment.ServiceCarvingType})

    evpnVirtualEthernetSegment.EntityData.YListKeys = []string {}

    return &(evpnVirtualEthernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_Identifier
// Ethernet segment identifier
// This type is a presence type.
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Type 0's 1st Byte or Type Byte and 1st Byte. The type is string with
    // pattern: [0-9a-fA-F]{1,8}. This attribute is mandatory.
    Bytes01 interface{}

    // 2nd and 3rd Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes23 interface{}

    // 4th and 5th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes45 interface{}

    // 6th and 7th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes67 interface{}

    // 8th and 9th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes89 interface{}

    // Ethernet segment identifier type. The type is EthernetSegmentIdentifier.
    // This attribute is mandatory.
    Type interface{}
}

func (identifier *Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_Identifier) GetEntityData() *types.CommonEntityData {
    identifier.EntityData.YFilter = identifier.YFilter
    identifier.EntityData.YangName = "identifier"
    identifier.EntityData.BundleName = "cisco_ios_xr"
    identifier.EntityData.ParentYangName = "evpn-virtual-ethernet-segment"
    identifier.EntityData.SegmentPath = "identifier"
    identifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifier.EntityData.Children = types.NewOrderedMap()
    identifier.EntityData.Leafs = types.NewOrderedMap()
    identifier.EntityData.Leafs.Append("bytes01", types.YLeaf{"Bytes01", identifier.Bytes01})
    identifier.EntityData.Leafs.Append("bytes23", types.YLeaf{"Bytes23", identifier.Bytes23})
    identifier.EntityData.Leafs.Append("bytes45", types.YLeaf{"Bytes45", identifier.Bytes45})
    identifier.EntityData.Leafs.Append("bytes67", types.YLeaf{"Bytes67", identifier.Bytes67})
    identifier.EntityData.Leafs.Append("bytes89", types.YLeaf{"Bytes89", identifier.Bytes89})
    identifier.EntityData.Leafs.Append("type", types.YLeaf{"Type", identifier.Type})

    identifier.EntityData.YListKeys = []string {}

    return &(identifier.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving
// Enter Manual service carving configuration
// submode
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Manual service carving. The type is interface{}.
    Enable interface{}

    // Manual service carving primary,secondary lists.
    ServiceList Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList
}

func (manualServiceCarving *Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving) GetEntityData() *types.CommonEntityData {
    manualServiceCarving.EntityData.YFilter = manualServiceCarving.YFilter
    manualServiceCarving.EntityData.YangName = "manual-service-carving"
    manualServiceCarving.EntityData.BundleName = "cisco_ios_xr"
    manualServiceCarving.EntityData.ParentYangName = "evpn-virtual-ethernet-segment"
    manualServiceCarving.EntityData.SegmentPath = "manual-service-carving"
    manualServiceCarving.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualServiceCarving.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualServiceCarving.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualServiceCarving.EntityData.Children = types.NewOrderedMap()
    manualServiceCarving.EntityData.Children.Append("service-list", types.YChild{"ServiceList", &manualServiceCarving.ServiceList})
    manualServiceCarving.EntityData.Leafs = types.NewOrderedMap()
    manualServiceCarving.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", manualServiceCarving.Enable})

    manualServiceCarving.EntityData.YListKeys = []string {}

    return &(manualServiceCarving.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList
// Manual service carving primary,secondary lists
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary services list. The type is string with length: 1..150.
    Primary interface{}

    // Secondary services list. The type is string with length: 1..150.
    Secondary interface{}
}

func (serviceList *Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList) GetEntityData() *types.CommonEntityData {
    serviceList.EntityData.YFilter = serviceList.YFilter
    serviceList.EntityData.YangName = "service-list"
    serviceList.EntityData.BundleName = "cisco_ios_xr"
    serviceList.EntityData.ParentYangName = "manual-service-carving"
    serviceList.EntityData.SegmentPath = "service-list"
    serviceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceList.EntityData.Children = types.NewOrderedMap()
    serviceList.EntityData.Leafs = types.NewOrderedMap()
    serviceList.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", serviceList.Primary})
    serviceList.EntityData.Leafs.Append("secondary", types.YLeaf{"Secondary", serviceList.Secondary})

    serviceList.EntityData.YListKeys = []string {}

    return &(serviceList.EntityData)
}

// Evpn_EvpnTables_EvpnLoadBalancing
// Enter EVPN Loadbalancing configuration submode
type Evpn_EvpnTables_EvpnLoadBalancing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Static Flow Label based load balancing. The type is interface{}.
    EvpnStaticFlowLabel interface{}

    // Enable EVPN Loadbalancing. The type is interface{}.
    Enable interface{}
}

func (evpnLoadBalancing *Evpn_EvpnTables_EvpnLoadBalancing) GetEntityData() *types.CommonEntityData {
    evpnLoadBalancing.EntityData.YFilter = evpnLoadBalancing.YFilter
    evpnLoadBalancing.EntityData.YangName = "evpn-load-balancing"
    evpnLoadBalancing.EntityData.BundleName = "cisco_ios_xr"
    evpnLoadBalancing.EntityData.ParentYangName = "evpn-tables"
    evpnLoadBalancing.EntityData.SegmentPath = "evpn-load-balancing"
    evpnLoadBalancing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnLoadBalancing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnLoadBalancing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnLoadBalancing.EntityData.Children = types.NewOrderedMap()
    evpnLoadBalancing.EntityData.Leafs = types.NewOrderedMap()
    evpnLoadBalancing.EntityData.Leafs.Append("evpn-static-flow-label", types.YLeaf{"EvpnStaticFlowLabel", evpnLoadBalancing.EvpnStaticFlowLabel})
    evpnLoadBalancing.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnLoadBalancing.Enable})

    evpnLoadBalancing.EntityData.YListKeys = []string {}

    return &(evpnLoadBalancing.EntityData)
}

// Evpn_EvpnTables_EvpnBgpAutoDiscovery
// Enable Autodiscovery BGP in EVPN
type Evpn_EvpnTables_EvpnBgpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Route Distinguisher.
    EvpnRouteDistinguisher Evpn_EvpnTables_EvpnBgpAutoDiscovery_EvpnRouteDistinguisher
}

func (evpnBgpAutoDiscovery *Evpn_EvpnTables_EvpnBgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    evpnBgpAutoDiscovery.EntityData.YFilter = evpnBgpAutoDiscovery.YFilter
    evpnBgpAutoDiscovery.EntityData.YangName = "evpn-bgp-auto-discovery"
    evpnBgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    evpnBgpAutoDiscovery.EntityData.ParentYangName = "evpn-tables"
    evpnBgpAutoDiscovery.EntityData.SegmentPath = "evpn-bgp-auto-discovery"
    evpnBgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnBgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnBgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnBgpAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    evpnBgpAutoDiscovery.EntityData.Children.Append("evpn-route-distinguisher", types.YChild{"EvpnRouteDistinguisher", &evpnBgpAutoDiscovery.EvpnRouteDistinguisher})
    evpnBgpAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    evpnBgpAutoDiscovery.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnBgpAutoDiscovery.Enable})

    evpnBgpAutoDiscovery.EntityData.YListKeys = []string {}

    return &(evpnBgpAutoDiscovery.EntityData)
}

// Evpn_EvpnTables_EvpnBgpAutoDiscovery_EvpnRouteDistinguisher
// Route Distinguisher
type Evpn_EvpnTables_EvpnBgpAutoDiscovery_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (evpnRouteDistinguisher *Evpn_EvpnTables_EvpnBgpAutoDiscovery_EvpnRouteDistinguisher) GetEntityData() *types.CommonEntityData {
    evpnRouteDistinguisher.EntityData.YFilter = evpnRouteDistinguisher.YFilter
    evpnRouteDistinguisher.EntityData.YangName = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteDistinguisher.EntityData.ParentYangName = "evpn-bgp-auto-discovery"
    evpnRouteDistinguisher.EntityData.SegmentPath = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteDistinguisher.EntityData.Children = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnRouteDistinguisher.Type})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", evpnRouteDistinguisher.As})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex})
    evpnRouteDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", evpnRouteDistinguisher.Address})
    evpnRouteDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex})

    evpnRouteDistinguisher.EntityData.YListKeys = []string {}

    return &(evpnRouteDistinguisher.EntityData)
}

// Evpn_EvpnTables_EvpnGroups
// Enter EVPN Group Table submode
type Evpn_EvpnTables_EvpnGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter EVPN Group submode. The type is slice of
    // Evpn_EvpnTables_EvpnGroups_EvpnGroup.
    EvpnGroup []*Evpn_EvpnTables_EvpnGroups_EvpnGroup
}

func (evpnGroups *Evpn_EvpnTables_EvpnGroups) GetEntityData() *types.CommonEntityData {
    evpnGroups.EntityData.YFilter = evpnGroups.YFilter
    evpnGroups.EntityData.YangName = "evpn-groups"
    evpnGroups.EntityData.BundleName = "cisco_ios_xr"
    evpnGroups.EntityData.ParentYangName = "evpn-tables"
    evpnGroups.EntityData.SegmentPath = "evpn-groups"
    evpnGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroups.EntityData.Children = types.NewOrderedMap()
    evpnGroups.EntityData.Children.Append("evpn-group", types.YChild{"EvpnGroup", nil})
    for i := range evpnGroups.EvpnGroup {
        evpnGroups.EntityData.Children.Append(types.GetSegmentPath(evpnGroups.EvpnGroup[i]), types.YChild{"EvpnGroup", evpnGroups.EvpnGroup[i]})
    }
    evpnGroups.EntityData.Leafs = types.NewOrderedMap()

    evpnGroups.EntityData.YListKeys = []string {}

    return &(evpnGroups.EntityData)
}

// Evpn_EvpnTables_EvpnGroups_EvpnGroup
// Enter EVPN Group submode
type Evpn_EvpnTables_EvpnGroups_EvpnGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..4294967295.
    GroupId interface{}

    // EVPN Group core interfaces.
    EvpnGroupCoreInterfaces Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces
}

func (evpnGroup *Evpn_EvpnTables_EvpnGroups_EvpnGroup) GetEntityData() *types.CommonEntityData {
    evpnGroup.EntityData.YFilter = evpnGroup.YFilter
    evpnGroup.EntityData.YangName = "evpn-group"
    evpnGroup.EntityData.BundleName = "cisco_ios_xr"
    evpnGroup.EntityData.ParentYangName = "evpn-groups"
    evpnGroup.EntityData.SegmentPath = "evpn-group" + types.AddKeyToken(evpnGroup.GroupId, "group-id")
    evpnGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroup.EntityData.Children = types.NewOrderedMap()
    evpnGroup.EntityData.Children.Append("evpn-group-core-interfaces", types.YChild{"EvpnGroupCoreInterfaces", &evpnGroup.EvpnGroupCoreInterfaces})
    evpnGroup.EntityData.Leafs = types.NewOrderedMap()
    evpnGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", evpnGroup.GroupId})

    evpnGroup.EntityData.YListKeys = []string {"GroupId"}

    return &(evpnGroup.EntityData)
}

// Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces
// EVPN Group core interfaces
type Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Group Core interface. The type is slice of
    // Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces_EvpnGroupCoreInterface.
    EvpnGroupCoreInterface []*Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces_EvpnGroupCoreInterface
}

func (evpnGroupCoreInterfaces *Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces) GetEntityData() *types.CommonEntityData {
    evpnGroupCoreInterfaces.EntityData.YFilter = evpnGroupCoreInterfaces.YFilter
    evpnGroupCoreInterfaces.EntityData.YangName = "evpn-group-core-interfaces"
    evpnGroupCoreInterfaces.EntityData.BundleName = "cisco_ios_xr"
    evpnGroupCoreInterfaces.EntityData.ParentYangName = "evpn-group"
    evpnGroupCoreInterfaces.EntityData.SegmentPath = "evpn-group-core-interfaces"
    evpnGroupCoreInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroupCoreInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroupCoreInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroupCoreInterfaces.EntityData.Children = types.NewOrderedMap()
    evpnGroupCoreInterfaces.EntityData.Children.Append("evpn-group-core-interface", types.YChild{"EvpnGroupCoreInterface", nil})
    for i := range evpnGroupCoreInterfaces.EvpnGroupCoreInterface {
        evpnGroupCoreInterfaces.EntityData.Children.Append(types.GetSegmentPath(evpnGroupCoreInterfaces.EvpnGroupCoreInterface[i]), types.YChild{"EvpnGroupCoreInterface", evpnGroupCoreInterfaces.EvpnGroupCoreInterface[i]})
    }
    evpnGroupCoreInterfaces.EntityData.Leafs = types.NewOrderedMap()

    evpnGroupCoreInterfaces.EntityData.YListKeys = []string {}

    return &(evpnGroupCoreInterfaces.EntityData)
}

// Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces_EvpnGroupCoreInterface
// EVPN Group Core interface
type Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces_EvpnGroupCoreInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the EVPN Group core interface. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (evpnGroupCoreInterface *Evpn_EvpnTables_EvpnGroups_EvpnGroup_EvpnGroupCoreInterfaces_EvpnGroupCoreInterface) GetEntityData() *types.CommonEntityData {
    evpnGroupCoreInterface.EntityData.YFilter = evpnGroupCoreInterface.YFilter
    evpnGroupCoreInterface.EntityData.YangName = "evpn-group-core-interface"
    evpnGroupCoreInterface.EntityData.BundleName = "cisco_ios_xr"
    evpnGroupCoreInterface.EntityData.ParentYangName = "evpn-group-core-interfaces"
    evpnGroupCoreInterface.EntityData.SegmentPath = "evpn-group-core-interface" + types.AddKeyToken(evpnGroupCoreInterface.InterfaceName, "interface-name")
    evpnGroupCoreInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroupCoreInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroupCoreInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroupCoreInterface.EntityData.Children = types.NewOrderedMap()
    evpnGroupCoreInterface.EntityData.Leafs = types.NewOrderedMap()
    evpnGroupCoreInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", evpnGroupCoreInterface.InterfaceName})

    evpnGroupCoreInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(evpnGroupCoreInterface.EntityData)
}

// Evpn_EvpnTables_EvpnInstances
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_EvpnInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter EVPN Instance configuration submode. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance.
    EvpnInstance []*Evpn_EvpnTables_EvpnInstances_EvpnInstance
}

func (evpnInstances *Evpn_EvpnTables_EvpnInstances) GetEntityData() *types.CommonEntityData {
    evpnInstances.EntityData.YFilter = evpnInstances.YFilter
    evpnInstances.EntityData.YangName = "evpn-instances"
    evpnInstances.EntityData.BundleName = "cisco_ios_xr"
    evpnInstances.EntityData.ParentYangName = "evpn-tables"
    evpnInstances.EntityData.SegmentPath = "evpn-instances"
    evpnInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstances.EntityData.Children = types.NewOrderedMap()
    evpnInstances.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", nil})
    for i := range evpnInstances.EvpnInstance {
        evpnInstances.EntityData.Children.Append(types.GetSegmentPath(evpnInstances.EvpnInstance[i]), types.YChild{"EvpnInstance", evpnInstances.EvpnInstance[i]})
    }
    evpnInstances.EntityData.Leafs = types.NewOrderedMap()

    evpnInstances.EntityData.YListKeys = []string {}

    return &(evpnInstances.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_EvpnInstances_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN Instance ID. The type is interface{} with
    // range: 1..4294967295.
    VpnId interface{}

    // This attribute is a key. EVPN Instance Encapsulation. The type is
    // EvpnEncapsulation.
    Encapsulation interface{}

    // This attribute is a key. EVPN Instance Side. The type is EvpnSide.
    Side interface{}

    // Disable route re-origination. The type is interface{}.
    EviReorigDisable interface{}

    // DEPRECATED: Advertise local MAC-only and BVI MAC routes. The type is
    // interface{}.
    EviAdvertiseMacDeprecated interface{}

    // EVPN Instance description. The type is string with length: 1..64.
    EvpnEviDescription interface{}

    // Disable ECMP on the EVI. The type is interface{}.
    EviEcmpDisable interface{}

    // Disable Unknown Unicast Flooding on this EVI. The type is interface{}.
    EviUnknownUnicastFloodingDisable interface{}

    // CW disable for EVPN EVI. The type is interface{}.
    EvpnEviCwDisable interface{}

    // Enable Autodiscovery BGP in EVPN Instance.
    EvpnInstanceBgpAutoDiscovery Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery

    // Enter Advertise local MAC-only routes configuration submode.
    EvpnInstanceAdvertiseMac Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceAdvertiseMac

    // Enter Multicast configuration submode.
    EvpnInstanceMulticast Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceMulticast

    // Enter Loadbalancing configuration submode.
    EvpnInstanceLoadBalancing Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceLoadBalancing
}

func (evpnInstance *Evpn_EvpnTables_EvpnInstances_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "evpn-instances"
    evpnInstance.EntityData.SegmentPath = "evpn-instance" + types.AddKeyToken(evpnInstance.VpnId, "vpn-id") + types.AddKeyToken(evpnInstance.Encapsulation, "encapsulation") + types.AddKeyToken(evpnInstance.Side, "side")
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Children.Append("evpn-instance-bgp-auto-discovery", types.YChild{"EvpnInstanceBgpAutoDiscovery", &evpnInstance.EvpnInstanceBgpAutoDiscovery})
    evpnInstance.EntityData.Children.Append("evpn-instance-advertise-mac", types.YChild{"EvpnInstanceAdvertiseMac", &evpnInstance.EvpnInstanceAdvertiseMac})
    evpnInstance.EntityData.Children.Append("evpn-instance-multicast", types.YChild{"EvpnInstanceMulticast", &evpnInstance.EvpnInstanceMulticast})
    evpnInstance.EntityData.Children.Append("evpn-instance-load-balancing", types.YChild{"EvpnInstanceLoadBalancing", &evpnInstance.EvpnInstanceLoadBalancing})
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", evpnInstance.VpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", evpnInstance.Encapsulation})
    evpnInstance.EntityData.Leafs.Append("side", types.YLeaf{"Side", evpnInstance.Side})
    evpnInstance.EntityData.Leafs.Append("evi-reorig-disable", types.YLeaf{"EviReorigDisable", evpnInstance.EviReorigDisable})
    evpnInstance.EntityData.Leafs.Append("evi-advertise-mac-deprecated", types.YLeaf{"EviAdvertiseMacDeprecated", evpnInstance.EviAdvertiseMacDeprecated})
    evpnInstance.EntityData.Leafs.Append("evpn-evi-description", types.YLeaf{"EvpnEviDescription", evpnInstance.EvpnEviDescription})
    evpnInstance.EntityData.Leafs.Append("evi-ecmp-disable", types.YLeaf{"EviEcmpDisable", evpnInstance.EviEcmpDisable})
    evpnInstance.EntityData.Leafs.Append("evi-unknown-unicast-flooding-disable", types.YLeaf{"EviUnknownUnicastFloodingDisable", evpnInstance.EviUnknownUnicastFloodingDisable})
    evpnInstance.EntityData.Leafs.Append("evpn-evi-cw-disable", types.YLeaf{"EvpnEviCwDisable", evpnInstance.EvpnEviCwDisable})

    evpnInstance.EntityData.YListKeys = []string {"VpnId", "Encapsulation", "Side"}

    return &(evpnInstance.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery
// Enable Autodiscovery BGP in EVPN Instance
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Table Policy for installation of forwarding data to L2FIB. The type is
    // string.
    TablePolicy interface{}

    // Route Distinguisher.
    EvpnRouteDistinguisher Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteDistinguisher

    // Route Target.
    EvpnRouteTargets Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets
}

func (evpnInstanceBgpAutoDiscovery *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    evpnInstanceBgpAutoDiscovery.EntityData.YFilter = evpnInstanceBgpAutoDiscovery.YFilter
    evpnInstanceBgpAutoDiscovery.EntityData.YangName = "evpn-instance-bgp-auto-discovery"
    evpnInstanceBgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    evpnInstanceBgpAutoDiscovery.EntityData.ParentYangName = "evpn-instance"
    evpnInstanceBgpAutoDiscovery.EntityData.SegmentPath = "evpn-instance-bgp-auto-discovery"
    evpnInstanceBgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstanceBgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstanceBgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstanceBgpAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    evpnInstanceBgpAutoDiscovery.EntityData.Children.Append("evpn-route-distinguisher", types.YChild{"EvpnRouteDistinguisher", &evpnInstanceBgpAutoDiscovery.EvpnRouteDistinguisher})
    evpnInstanceBgpAutoDiscovery.EntityData.Children.Append("evpn-route-targets", types.YChild{"EvpnRouteTargets", &evpnInstanceBgpAutoDiscovery.EvpnRouteTargets})
    evpnInstanceBgpAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    evpnInstanceBgpAutoDiscovery.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnInstanceBgpAutoDiscovery.Enable})
    evpnInstanceBgpAutoDiscovery.EntityData.Leafs.Append("table-policy", types.YLeaf{"TablePolicy", evpnInstanceBgpAutoDiscovery.TablePolicy})

    evpnInstanceBgpAutoDiscovery.EntityData.YListKeys = []string {}

    return &(evpnInstanceBgpAutoDiscovery.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteDistinguisher
// Route Distinguisher
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (evpnRouteDistinguisher *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteDistinguisher) GetEntityData() *types.CommonEntityData {
    evpnRouteDistinguisher.EntityData.YFilter = evpnRouteDistinguisher.YFilter
    evpnRouteDistinguisher.EntityData.YangName = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteDistinguisher.EntityData.ParentYangName = "evpn-instance-bgp-auto-discovery"
    evpnRouteDistinguisher.EntityData.SegmentPath = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteDistinguisher.EntityData.Children = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteDistinguisher.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnRouteDistinguisher.Type})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as", types.YLeaf{"As", evpnRouteDistinguisher.As})
    evpnRouteDistinguisher.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex})
    evpnRouteDistinguisher.EntityData.Leafs.Append("address", types.YLeaf{"Address", evpnRouteDistinguisher.Address})
    evpnRouteDistinguisher.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex})

    evpnRouteDistinguisher.EntityData.YListKeys = []string {}

    return &(evpnRouteDistinguisher.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets
// Route Target
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs.
    EvpnRouteTargetAs []*Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone.
    EvpnRouteTargetNone []*Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address.
    EvpnRouteTargetIpv4Address []*Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
}

func (evpnRouteTargets *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets) GetEntityData() *types.CommonEntityData {
    evpnRouteTargets.EntityData.YFilter = evpnRouteTargets.YFilter
    evpnRouteTargets.EntityData.YangName = "evpn-route-targets"
    evpnRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargets.EntityData.ParentYangName = "evpn-instance-bgp-auto-discovery"
    evpnRouteTargets.EntityData.SegmentPath = "evpn-route-targets"
    evpnRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargets.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargets.EntityData.Children.Append("evpn-route-target-as", types.YChild{"EvpnRouteTargetAs", nil})
    for i := range evpnRouteTargets.EvpnRouteTargetAs {
        evpnRouteTargets.EntityData.Children.Append(types.GetSegmentPath(evpnRouteTargets.EvpnRouteTargetAs[i]), types.YChild{"EvpnRouteTargetAs", evpnRouteTargets.EvpnRouteTargetAs[i]})
    }
    evpnRouteTargets.EntityData.Children.Append("evpn-route-target-none", types.YChild{"EvpnRouteTargetNone", nil})
    for i := range evpnRouteTargets.EvpnRouteTargetNone {
        evpnRouteTargets.EntityData.Children.Append(types.GetSegmentPath(evpnRouteTargets.EvpnRouteTargetNone[i]), types.YChild{"EvpnRouteTargetNone", evpnRouteTargets.EvpnRouteTargetNone[i]})
    }
    evpnRouteTargets.EntityData.Children.Append("evpn-route-target-ipv4-address", types.YChild{"EvpnRouteTargetIpv4Address", nil})
    for i := range evpnRouteTargets.EvpnRouteTargetIpv4Address {
        evpnRouteTargets.EntityData.Children.Append(types.GetSegmentPath(evpnRouteTargets.EvpnRouteTargetIpv4Address[i]), types.YChild{"EvpnRouteTargetIpv4Address", evpnRouteTargets.EvpnRouteTargetIpv4Address[i]})
    }
    evpnRouteTargets.EntityData.Leafs = types.NewOrderedMap()

    evpnRouteTargets.EntityData.YListKeys = []string {}

    return &(evpnRouteTargets.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs
// Name of the Route Target
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetAs *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetAs.EntityData.YFilter = evpnRouteTargetAs.YFilter
    evpnRouteTargetAs.EntityData.YangName = "evpn-route-target-as"
    evpnRouteTargetAs.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetAs.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetAs.EntityData.SegmentPath = "evpn-route-target-as" + types.AddKeyToken(evpnRouteTargetAs.Format, "format") + types.AddKeyToken(evpnRouteTargetAs.Role, "role") + types.AddKeyToken(evpnRouteTargetAs.As, "as") + types.AddKeyToken(evpnRouteTargetAs.AsIndex, "as-index") + types.AddKeyToken(evpnRouteTargetAs.Stitching, "stitching")
    evpnRouteTargetAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetAs.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargetAs.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteTargetAs.EntityData.Leafs.Append("format", types.YLeaf{"Format", evpnRouteTargetAs.Format})
    evpnRouteTargetAs.EntityData.Leafs.Append("role", types.YLeaf{"Role", evpnRouteTargetAs.Role})
    evpnRouteTargetAs.EntityData.Leafs.Append("as", types.YLeaf{"As", evpnRouteTargetAs.As})
    evpnRouteTargetAs.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", evpnRouteTargetAs.AsIndex})
    evpnRouteTargetAs.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", evpnRouteTargetAs.Stitching})

    evpnRouteTargetAs.EntityData.YListKeys = []string {"Format", "Role", "As", "AsIndex", "Stitching"}

    return &(evpnRouteTargetAs.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone
// Name of the Route Target
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetNone *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetNone.EntityData.YFilter = evpnRouteTargetNone.YFilter
    evpnRouteTargetNone.EntityData.YangName = "evpn-route-target-none"
    evpnRouteTargetNone.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetNone.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetNone.EntityData.SegmentPath = "evpn-route-target-none" + types.AddKeyToken(evpnRouteTargetNone.Format, "format") + types.AddKeyToken(evpnRouteTargetNone.Role, "role") + types.AddKeyToken(evpnRouteTargetNone.Stitching, "stitching")
    evpnRouteTargetNone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetNone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetNone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetNone.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargetNone.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteTargetNone.EntityData.Leafs.Append("format", types.YLeaf{"Format", evpnRouteTargetNone.Format})
    evpnRouteTargetNone.EntityData.Leafs.Append("role", types.YLeaf{"Role", evpnRouteTargetNone.Role})
    evpnRouteTargetNone.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", evpnRouteTargetNone.Stitching})

    evpnRouteTargetNone.EntityData.YListKeys = []string {"Format", "Role", "Stitching"}

    return &(evpnRouteTargetNone.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
// Name of the Route Target
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetIpv4Address *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetIpv4Address.EntityData.YFilter = evpnRouteTargetIpv4Address.YFilter
    evpnRouteTargetIpv4Address.EntityData.YangName = "evpn-route-target-ipv4-address"
    evpnRouteTargetIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetIpv4Address.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetIpv4Address.EntityData.SegmentPath = "evpn-route-target-ipv4-address" + types.AddKeyToken(evpnRouteTargetIpv4Address.Format, "format") + types.AddKeyToken(evpnRouteTargetIpv4Address.Role, "role") + types.AddKeyToken(evpnRouteTargetIpv4Address.Address, "address") + types.AddKeyToken(evpnRouteTargetIpv4Address.AddrIndex, "addr-index") + types.AddKeyToken(evpnRouteTargetIpv4Address.Stitching, "stitching")
    evpnRouteTargetIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetIpv4Address.EntityData.Children = types.NewOrderedMap()
    evpnRouteTargetIpv4Address.EntityData.Leafs = types.NewOrderedMap()
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("format", types.YLeaf{"Format", evpnRouteTargetIpv4Address.Format})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("role", types.YLeaf{"Role", evpnRouteTargetIpv4Address.Role})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", evpnRouteTargetIpv4Address.Address})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", evpnRouteTargetIpv4Address.AddrIndex})
    evpnRouteTargetIpv4Address.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", evpnRouteTargetIpv4Address.Stitching})

    evpnRouteTargetIpv4Address.EntityData.YListKeys = []string {"Format", "Role", "Address", "AddrIndex", "Stitching"}

    return &(evpnRouteTargetIpv4Address.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceAdvertiseMac
// Enter Advertise local MAC-only routes
// configuration submode
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceAdvertiseMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Advertise local MAC-only routes. The type is interface{}.
    Enable interface{}

    // Advertise local MAC-only and BVI MAC routes. The type is interface{}.
    EviAdvertiseMacBvi interface{}
}

func (evpnInstanceAdvertiseMac *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceAdvertiseMac) GetEntityData() *types.CommonEntityData {
    evpnInstanceAdvertiseMac.EntityData.YFilter = evpnInstanceAdvertiseMac.YFilter
    evpnInstanceAdvertiseMac.EntityData.YangName = "evpn-instance-advertise-mac"
    evpnInstanceAdvertiseMac.EntityData.BundleName = "cisco_ios_xr"
    evpnInstanceAdvertiseMac.EntityData.ParentYangName = "evpn-instance"
    evpnInstanceAdvertiseMac.EntityData.SegmentPath = "evpn-instance-advertise-mac"
    evpnInstanceAdvertiseMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstanceAdvertiseMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstanceAdvertiseMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstanceAdvertiseMac.EntityData.Children = types.NewOrderedMap()
    evpnInstanceAdvertiseMac.EntityData.Leafs = types.NewOrderedMap()
    evpnInstanceAdvertiseMac.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnInstanceAdvertiseMac.Enable})
    evpnInstanceAdvertiseMac.EntityData.Leafs.Append("evi-advertise-mac-bvi", types.YLeaf{"EviAdvertiseMacBvi", evpnInstanceAdvertiseMac.EviAdvertiseMacBvi})

    evpnInstanceAdvertiseMac.EntityData.YListKeys = []string {}

    return &(evpnInstanceAdvertiseMac.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceMulticast
// Enter Multicast configuration submode
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Multicast. The type is interface{}.
    Enable interface{}

    // Enable Multicast source connectivity. The type is interface{}.
    EviMcastSourceConnected interface{}
}

func (evpnInstanceMulticast *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceMulticast) GetEntityData() *types.CommonEntityData {
    evpnInstanceMulticast.EntityData.YFilter = evpnInstanceMulticast.YFilter
    evpnInstanceMulticast.EntityData.YangName = "evpn-instance-multicast"
    evpnInstanceMulticast.EntityData.BundleName = "cisco_ios_xr"
    evpnInstanceMulticast.EntityData.ParentYangName = "evpn-instance"
    evpnInstanceMulticast.EntityData.SegmentPath = "evpn-instance-multicast"
    evpnInstanceMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstanceMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstanceMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstanceMulticast.EntityData.Children = types.NewOrderedMap()
    evpnInstanceMulticast.EntityData.Leafs = types.NewOrderedMap()
    evpnInstanceMulticast.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnInstanceMulticast.Enable})
    evpnInstanceMulticast.EntityData.Leafs.Append("evi-mcast-source-connected", types.YLeaf{"EviMcastSourceConnected", evpnInstanceMulticast.EviMcastSourceConnected})

    evpnInstanceMulticast.EntityData.YListKeys = []string {}

    return &(evpnInstanceMulticast.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceLoadBalancing
// Enter Loadbalancing configuration submode
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceLoadBalancing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Loadbalancing. The type is interface{}.
    Enable interface{}

    // Enable Static Flow Label based load balancing. The type is interface{}.
    EviStaticFlowLabel interface{}
}

func (evpnInstanceLoadBalancing *Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceLoadBalancing) GetEntityData() *types.CommonEntityData {
    evpnInstanceLoadBalancing.EntityData.YFilter = evpnInstanceLoadBalancing.YFilter
    evpnInstanceLoadBalancing.EntityData.YangName = "evpn-instance-load-balancing"
    evpnInstanceLoadBalancing.EntityData.BundleName = "cisco_ios_xr"
    evpnInstanceLoadBalancing.EntityData.ParentYangName = "evpn-instance"
    evpnInstanceLoadBalancing.EntityData.SegmentPath = "evpn-instance-load-balancing"
    evpnInstanceLoadBalancing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstanceLoadBalancing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstanceLoadBalancing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstanceLoadBalancing.EntityData.Children = types.NewOrderedMap()
    evpnInstanceLoadBalancing.EntityData.Leafs = types.NewOrderedMap()
    evpnInstanceLoadBalancing.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnInstanceLoadBalancing.Enable})
    evpnInstanceLoadBalancing.EntityData.Leafs.Append("evi-static-flow-label", types.YLeaf{"EviStaticFlowLabel", evpnInstanceLoadBalancing.EviStaticFlowLabel})

    evpnInstanceLoadBalancing.EntityData.YListKeys = []string {}

    return &(evpnInstanceLoadBalancing.EntityData)
}

// Evpn_EvpnTables_EvpnLogging
// Enter EVPN Logging configuration submode
type Evpn_EvpnTables_EvpnLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Designated Forwarder election logging. The type is interface{}.
    EvpnDfElection interface{}

    // Enable EVPN Logging. The type is interface{}.
    Enable interface{}
}

func (evpnLogging *Evpn_EvpnTables_EvpnLogging) GetEntityData() *types.CommonEntityData {
    evpnLogging.EntityData.YFilter = evpnLogging.YFilter
    evpnLogging.EntityData.YangName = "evpn-logging"
    evpnLogging.EntityData.BundleName = "cisco_ios_xr"
    evpnLogging.EntityData.ParentYangName = "evpn-tables"
    evpnLogging.EntityData.SegmentPath = "evpn-logging"
    evpnLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnLogging.EntityData.Children = types.NewOrderedMap()
    evpnLogging.EntityData.Leafs = types.NewOrderedMap()
    evpnLogging.EntityData.Leafs.Append("evpn-df-election", types.YLeaf{"EvpnDfElection", evpnLogging.EvpnDfElection})
    evpnLogging.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnLogging.Enable})

    evpnLogging.EntityData.YListKeys = []string {}

    return &(evpnLogging.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces
// Attachment Circuit interfaces
type Evpn_EvpnTables_EvpnInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // Evpn_EvpnTables_EvpnInterfaces_EvpnInterface.
    EvpnInterface []*Evpn_EvpnTables_EvpnInterfaces_EvpnInterface
}

func (evpnInterfaces *Evpn_EvpnTables_EvpnInterfaces) GetEntityData() *types.CommonEntityData {
    evpnInterfaces.EntityData.YFilter = evpnInterfaces.YFilter
    evpnInterfaces.EntityData.YangName = "evpn-interfaces"
    evpnInterfaces.EntityData.BundleName = "cisco_ios_xr"
    evpnInterfaces.EntityData.ParentYangName = "evpn-tables"
    evpnInterfaces.EntityData.SegmentPath = "evpn-interfaces"
    evpnInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInterfaces.EntityData.Children = types.NewOrderedMap()
    evpnInterfaces.EntityData.Children.Append("evpn-interface", types.YChild{"EvpnInterface", nil})
    for i := range evpnInterfaces.EvpnInterface {
        evpnInterfaces.EntityData.Children.Append(types.GetSegmentPath(evpnInterfaces.EvpnInterface[i]), types.YChild{"EvpnInterface", evpnInterfaces.EvpnInterface[i]})
    }
    evpnInterfaces.EntityData.Leafs = types.NewOrderedMap()

    evpnInterfaces.EntityData.YListKeys = []string {}

    return &(evpnInterfaces.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface
// Attachment circuit interface
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Enter EVPN Core Isolation Group ID. The type is interface{} with range:
    // 1..4294967295.
    EvpnCoreIsolationGroup interface{}

    // Enable MAC Flushing. The type is MacFlushMode.
    MacFlush interface{}

    // Enter Interface-specific timers configuration submode.
    EvpnacTimers Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EvpnacTimers

    // Enter Ethernet Segment configuration submode.
    EthernetSegment Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment
}

func (evpnInterface *Evpn_EvpnTables_EvpnInterfaces_EvpnInterface) GetEntityData() *types.CommonEntityData {
    evpnInterface.EntityData.YFilter = evpnInterface.YFilter
    evpnInterface.EntityData.YangName = "evpn-interface"
    evpnInterface.EntityData.BundleName = "cisco_ios_xr"
    evpnInterface.EntityData.ParentYangName = "evpn-interfaces"
    evpnInterface.EntityData.SegmentPath = "evpn-interface" + types.AddKeyToken(evpnInterface.InterfaceName, "interface-name")
    evpnInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInterface.EntityData.Children = types.NewOrderedMap()
    evpnInterface.EntityData.Children.Append("evpnac-timers", types.YChild{"EvpnacTimers", &evpnInterface.EvpnacTimers})
    evpnInterface.EntityData.Children.Append("ethernet-segment", types.YChild{"EthernetSegment", &evpnInterface.EthernetSegment})
    evpnInterface.EntityData.Leafs = types.NewOrderedMap()
    evpnInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", evpnInterface.InterfaceName})
    evpnInterface.EntityData.Leafs.Append("evpn-core-isolation-group", types.YLeaf{"EvpnCoreIsolationGroup", evpnInterface.EvpnCoreIsolationGroup})
    evpnInterface.EntityData.Leafs.Append("mac-flush", types.YLeaf{"MacFlush", evpnInterface.MacFlush})

    evpnInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(evpnInterface.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EvpnacTimers
// Enter Interface-specific timers configuration
// submode
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EvpnacTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-specific Peering timer. The type is interface{} with range:
    // 0..300. The default value is 3.
    EvpnacPeering interface{}

    // Interface-specific Carving timer. The type is interface{} with range:
    // 0..10. The default value is 0.
    EvpnacCarving interface{}

    // Enable Interface-specific timers. The type is interface{}.
    Enable interface{}

    // Interface-specific Recovery timer. The type is interface{} with range:
    // 20..3600. The default value is 30.
    EvpnacRecovery interface{}
}

func (evpnacTimers *Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EvpnacTimers) GetEntityData() *types.CommonEntityData {
    evpnacTimers.EntityData.YFilter = evpnacTimers.YFilter
    evpnacTimers.EntityData.YangName = "evpnac-timers"
    evpnacTimers.EntityData.BundleName = "cisco_ios_xr"
    evpnacTimers.EntityData.ParentYangName = "evpn-interface"
    evpnacTimers.EntityData.SegmentPath = "evpnac-timers"
    evpnacTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnacTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnacTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnacTimers.EntityData.Children = types.NewOrderedMap()
    evpnacTimers.EntityData.Leafs = types.NewOrderedMap()
    evpnacTimers.EntityData.Leafs.Append("evpnac-peering", types.YLeaf{"EvpnacPeering", evpnacTimers.EvpnacPeering})
    evpnacTimers.EntityData.Leafs.Append("evpnac-carving", types.YLeaf{"EvpnacCarving", evpnacTimers.EvpnacCarving})
    evpnacTimers.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnacTimers.Enable})
    evpnacTimers.EntityData.Leafs.Append("evpnac-recovery", types.YLeaf{"EvpnacRecovery", evpnacTimers.EvpnacRecovery})

    evpnacTimers.EntityData.YListKeys = []string {}

    return &(evpnacTimers.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment
// Enter Ethernet Segment configuration submode
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force ethernet segment to remain single-homed. The type is interface{}.
    ForceSingleHomed interface{}

    // Ethernet-Segment Load Balancing mode. The type is
    // EthernetSegmentLoadBalance.
    LoadBalancingMode interface{}

    // Enable Ethernet Segment. The type is interface{}.
    Enable interface{}

    // Backbone Source MAC. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BackboneSourceMac interface{}

    // ES-Import Route Target. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    EsImportRouteTarget interface{}

    // Ethernet-Segment Service Carving mode. The type is
    // EthernetSegmentServiceCarving.
    ServiceCarvingType interface{}

    // Ethernet segment identifier.
    Identifier Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_Identifier

    // Enter Manual service carving configuration submode.
    ManualServiceCarving Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving
}

func (ethernetSegment *Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment) GetEntityData() *types.CommonEntityData {
    ethernetSegment.EntityData.YFilter = ethernetSegment.YFilter
    ethernetSegment.EntityData.YangName = "ethernet-segment"
    ethernetSegment.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegment.EntityData.ParentYangName = "evpn-interface"
    ethernetSegment.EntityData.SegmentPath = "ethernet-segment"
    ethernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegment.EntityData.Children = types.NewOrderedMap()
    ethernetSegment.EntityData.Children.Append("identifier", types.YChild{"Identifier", &ethernetSegment.Identifier})
    ethernetSegment.EntityData.Children.Append("manual-service-carving", types.YChild{"ManualServiceCarving", &ethernetSegment.ManualServiceCarving})
    ethernetSegment.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegment.EntityData.Leafs.Append("force-single-homed", types.YLeaf{"ForceSingleHomed", ethernetSegment.ForceSingleHomed})
    ethernetSegment.EntityData.Leafs.Append("load-balancing-mode", types.YLeaf{"LoadBalancingMode", ethernetSegment.LoadBalancingMode})
    ethernetSegment.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ethernetSegment.Enable})
    ethernetSegment.EntityData.Leafs.Append("backbone-source-mac", types.YLeaf{"BackboneSourceMac", ethernetSegment.BackboneSourceMac})
    ethernetSegment.EntityData.Leafs.Append("es-import-route-target", types.YLeaf{"EsImportRouteTarget", ethernetSegment.EsImportRouteTarget})
    ethernetSegment.EntityData.Leafs.Append("service-carving-type", types.YLeaf{"ServiceCarvingType", ethernetSegment.ServiceCarvingType})

    ethernetSegment.EntityData.YListKeys = []string {}

    return &(ethernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_Identifier
// Ethernet segment identifier
// This type is a presence type.
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Type 0's 1st Byte or Type Byte and 1st Byte. The type is string with
    // pattern: [0-9a-fA-F]{1,8}. This attribute is mandatory.
    Bytes01 interface{}

    // 2nd and 3rd Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes23 interface{}

    // 4th and 5th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes45 interface{}

    // 6th and 7th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes67 interface{}

    // 8th and 9th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes89 interface{}

    // Ethernet segment identifier type. The type is EthernetSegmentIdentifier.
    // This attribute is mandatory.
    Type interface{}
}

func (identifier *Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_Identifier) GetEntityData() *types.CommonEntityData {
    identifier.EntityData.YFilter = identifier.YFilter
    identifier.EntityData.YangName = "identifier"
    identifier.EntityData.BundleName = "cisco_ios_xr"
    identifier.EntityData.ParentYangName = "ethernet-segment"
    identifier.EntityData.SegmentPath = "identifier"
    identifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifier.EntityData.Children = types.NewOrderedMap()
    identifier.EntityData.Leafs = types.NewOrderedMap()
    identifier.EntityData.Leafs.Append("bytes01", types.YLeaf{"Bytes01", identifier.Bytes01})
    identifier.EntityData.Leafs.Append("bytes23", types.YLeaf{"Bytes23", identifier.Bytes23})
    identifier.EntityData.Leafs.Append("bytes45", types.YLeaf{"Bytes45", identifier.Bytes45})
    identifier.EntityData.Leafs.Append("bytes67", types.YLeaf{"Bytes67", identifier.Bytes67})
    identifier.EntityData.Leafs.Append("bytes89", types.YLeaf{"Bytes89", identifier.Bytes89})
    identifier.EntityData.Leafs.Append("type", types.YLeaf{"Type", identifier.Type})

    identifier.EntityData.YListKeys = []string {}

    return &(identifier.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving
// Enter Manual service carving configuration
// submode
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Manual service carving. The type is interface{}.
    Enable interface{}

    // Manual service carving primary,secondary lists.
    ServiceList Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving_ServiceList
}

func (manualServiceCarving *Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving) GetEntityData() *types.CommonEntityData {
    manualServiceCarving.EntityData.YFilter = manualServiceCarving.YFilter
    manualServiceCarving.EntityData.YangName = "manual-service-carving"
    manualServiceCarving.EntityData.BundleName = "cisco_ios_xr"
    manualServiceCarving.EntityData.ParentYangName = "ethernet-segment"
    manualServiceCarving.EntityData.SegmentPath = "manual-service-carving"
    manualServiceCarving.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualServiceCarving.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualServiceCarving.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualServiceCarving.EntityData.Children = types.NewOrderedMap()
    manualServiceCarving.EntityData.Children.Append("service-list", types.YChild{"ServiceList", &manualServiceCarving.ServiceList})
    manualServiceCarving.EntityData.Leafs = types.NewOrderedMap()
    manualServiceCarving.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", manualServiceCarving.Enable})

    manualServiceCarving.EntityData.YListKeys = []string {}

    return &(manualServiceCarving.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving_ServiceList
// Manual service carving primary,secondary lists
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving_ServiceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary services list. The type is string with length: 1..150.
    Primary interface{}

    // Secondary services list. The type is string with length: 1..150.
    Secondary interface{}
}

func (serviceList *Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_ManualServiceCarving_ServiceList) GetEntityData() *types.CommonEntityData {
    serviceList.EntityData.YFilter = serviceList.YFilter
    serviceList.EntityData.YangName = "service-list"
    serviceList.EntityData.BundleName = "cisco_ios_xr"
    serviceList.EntityData.ParentYangName = "manual-service-carving"
    serviceList.EntityData.SegmentPath = "service-list"
    serviceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceList.EntityData.Children = types.NewOrderedMap()
    serviceList.EntityData.Leafs = types.NewOrderedMap()
    serviceList.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", serviceList.Primary})
    serviceList.EntityData.Leafs.Append("secondary", types.YLeaf{"Secondary", serviceList.Secondary})

    serviceList.EntityData.YListKeys = []string {}

    return &(serviceList.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws
// Virtual Access Pseudowire interfaces
type Evpn_EvpnTables_EvpnVirtualAccessPws struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual Access Pseudowire. The type is slice of
    // Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw.
    EvpnVirtualAccessPw []*Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw
}

func (evpnVirtualAccessPws *Evpn_EvpnTables_EvpnVirtualAccessPws) GetEntityData() *types.CommonEntityData {
    evpnVirtualAccessPws.EntityData.YFilter = evpnVirtualAccessPws.YFilter
    evpnVirtualAccessPws.EntityData.YangName = "evpn-virtual-access-pws"
    evpnVirtualAccessPws.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualAccessPws.EntityData.ParentYangName = "evpn-tables"
    evpnVirtualAccessPws.EntityData.SegmentPath = "evpn-virtual-access-pws"
    evpnVirtualAccessPws.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessPws.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessPws.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessPws.EntityData.Children = types.NewOrderedMap()
    evpnVirtualAccessPws.EntityData.Children.Append("evpn-virtual-access-pw", types.YChild{"EvpnVirtualAccessPw", nil})
    for i := range evpnVirtualAccessPws.EvpnVirtualAccessPw {
        evpnVirtualAccessPws.EntityData.Children.Append(types.GetSegmentPath(evpnVirtualAccessPws.EvpnVirtualAccessPw[i]), types.YChild{"EvpnVirtualAccessPw", evpnVirtualAccessPws.EvpnVirtualAccessPw[i]})
    }
    evpnVirtualAccessPws.EntityData.Leafs = types.NewOrderedMap()

    evpnVirtualAccessPws.EntityData.YListKeys = []string {}

    return &(evpnVirtualAccessPws.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw
// Virtual Access Pseudowire
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // Enter Virtual Access Pseudowire-specific timers configuration submode.
    EvpnVirtualAccessPwTimers Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualAccessPwTimers

    // Enter Ethernet Segment configuration submode.
    EvpnVirtualEthernetSegment Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment
}

func (evpnVirtualAccessPw *Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw) GetEntityData() *types.CommonEntityData {
    evpnVirtualAccessPw.EntityData.YFilter = evpnVirtualAccessPw.YFilter
    evpnVirtualAccessPw.EntityData.YangName = "evpn-virtual-access-pw"
    evpnVirtualAccessPw.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualAccessPw.EntityData.ParentYangName = "evpn-virtual-access-pws"
    evpnVirtualAccessPw.EntityData.SegmentPath = "evpn-virtual-access-pw" + types.AddKeyToken(evpnVirtualAccessPw.Neighbor, "neighbor") + types.AddKeyToken(evpnVirtualAccessPw.PseudowireId, "pseudowire-id")
    evpnVirtualAccessPw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessPw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessPw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessPw.EntityData.Children = types.NewOrderedMap()
    evpnVirtualAccessPw.EntityData.Children.Append("evpn-virtual-access-pw-timers", types.YChild{"EvpnVirtualAccessPwTimers", &evpnVirtualAccessPw.EvpnVirtualAccessPwTimers})
    evpnVirtualAccessPw.EntityData.Children.Append("evpn-virtual-ethernet-segment", types.YChild{"EvpnVirtualEthernetSegment", &evpnVirtualAccessPw.EvpnVirtualEthernetSegment})
    evpnVirtualAccessPw.EntityData.Leafs = types.NewOrderedMap()
    evpnVirtualAccessPw.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", evpnVirtualAccessPw.Neighbor})
    evpnVirtualAccessPw.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", evpnVirtualAccessPw.PseudowireId})

    evpnVirtualAccessPw.EntityData.YListKeys = []string {"Neighbor", "PseudowireId"}

    return &(evpnVirtualAccessPw.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualAccessPwTimers
// Enter Virtual Access Pseudowire-specific
// timers configuration submode
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualAccessPwTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual Access Pseudowire-specific Recovery timer. The type is interface{}
    // with range: 20..3600. The default value is 30.
    EvpnVirtualAccessPwRecovery interface{}

    // Virtual Access Pseudowire-specific Peering timer. The type is interface{}
    // with range: 0..300. The default value is 3.
    EvpnVirtualAccessPwPeering interface{}

    // Enable Virtual Access Pseudowire-specific timers. The type is interface{}.
    Enable interface{}

    // Virtual Access Pseudowire-specific Carving timer. The type is interface{}
    // with range: 0..10. The default value is 0.
    EvpnVirtualAccessPwCarving interface{}
}

func (evpnVirtualAccessPwTimers *Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualAccessPwTimers) GetEntityData() *types.CommonEntityData {
    evpnVirtualAccessPwTimers.EntityData.YFilter = evpnVirtualAccessPwTimers.YFilter
    evpnVirtualAccessPwTimers.EntityData.YangName = "evpn-virtual-access-pw-timers"
    evpnVirtualAccessPwTimers.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualAccessPwTimers.EntityData.ParentYangName = "evpn-virtual-access-pw"
    evpnVirtualAccessPwTimers.EntityData.SegmentPath = "evpn-virtual-access-pw-timers"
    evpnVirtualAccessPwTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessPwTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessPwTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessPwTimers.EntityData.Children = types.NewOrderedMap()
    evpnVirtualAccessPwTimers.EntityData.Leafs = types.NewOrderedMap()
    evpnVirtualAccessPwTimers.EntityData.Leafs.Append("evpn-virtual-access-pw-recovery", types.YLeaf{"EvpnVirtualAccessPwRecovery", evpnVirtualAccessPwTimers.EvpnVirtualAccessPwRecovery})
    evpnVirtualAccessPwTimers.EntityData.Leafs.Append("evpn-virtual-access-pw-peering", types.YLeaf{"EvpnVirtualAccessPwPeering", evpnVirtualAccessPwTimers.EvpnVirtualAccessPwPeering})
    evpnVirtualAccessPwTimers.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnVirtualAccessPwTimers.Enable})
    evpnVirtualAccessPwTimers.EntityData.Leafs.Append("evpn-virtual-access-pw-carving", types.YLeaf{"EvpnVirtualAccessPwCarving", evpnVirtualAccessPwTimers.EvpnVirtualAccessPwCarving})

    evpnVirtualAccessPwTimers.EntityData.YListKeys = []string {}

    return &(evpnVirtualAccessPwTimers.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment
// Enter Ethernet Segment configuration submode
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Ethernet Segment. The type is interface{}.
    Enable interface{}

    // ES-Import Route Target. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    EsImportRouteTarget interface{}

    // Ethernet-Segment Service Carving mode. The type is
    // EthernetSegmentServiceCarving.
    ServiceCarvingType interface{}

    // Ethernet segment identifier.
    Identifier Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_Identifier

    // Enter Manual service carving configuration submode.
    ManualServiceCarving Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving
}

func (evpnVirtualEthernetSegment *Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment) GetEntityData() *types.CommonEntityData {
    evpnVirtualEthernetSegment.EntityData.YFilter = evpnVirtualEthernetSegment.YFilter
    evpnVirtualEthernetSegment.EntityData.YangName = "evpn-virtual-ethernet-segment"
    evpnVirtualEthernetSegment.EntityData.BundleName = "cisco_ios_xr"
    evpnVirtualEthernetSegment.EntityData.ParentYangName = "evpn-virtual-access-pw"
    evpnVirtualEthernetSegment.EntityData.SegmentPath = "evpn-virtual-ethernet-segment"
    evpnVirtualEthernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualEthernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualEthernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualEthernetSegment.EntityData.Children = types.NewOrderedMap()
    evpnVirtualEthernetSegment.EntityData.Children.Append("identifier", types.YChild{"Identifier", &evpnVirtualEthernetSegment.Identifier})
    evpnVirtualEthernetSegment.EntityData.Children.Append("manual-service-carving", types.YChild{"ManualServiceCarving", &evpnVirtualEthernetSegment.ManualServiceCarving})
    evpnVirtualEthernetSegment.EntityData.Leafs = types.NewOrderedMap()
    evpnVirtualEthernetSegment.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnVirtualEthernetSegment.Enable})
    evpnVirtualEthernetSegment.EntityData.Leafs.Append("es-import-route-target", types.YLeaf{"EsImportRouteTarget", evpnVirtualEthernetSegment.EsImportRouteTarget})
    evpnVirtualEthernetSegment.EntityData.Leafs.Append("service-carving-type", types.YLeaf{"ServiceCarvingType", evpnVirtualEthernetSegment.ServiceCarvingType})

    evpnVirtualEthernetSegment.EntityData.YListKeys = []string {}

    return &(evpnVirtualEthernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_Identifier
// Ethernet segment identifier
// This type is a presence type.
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Type 0's 1st Byte or Type Byte and 1st Byte. The type is string with
    // pattern: [0-9a-fA-F]{1,8}. This attribute is mandatory.
    Bytes01 interface{}

    // 2nd and 3rd Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes23 interface{}

    // 4th and 5th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes45 interface{}

    // 6th and 7th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes67 interface{}

    // 8th and 9th Bytes. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory. Units are byte.
    Bytes89 interface{}

    // Ethernet segment identifier type. The type is EthernetSegmentIdentifier.
    // This attribute is mandatory.
    Type interface{}
}

func (identifier *Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_Identifier) GetEntityData() *types.CommonEntityData {
    identifier.EntityData.YFilter = identifier.YFilter
    identifier.EntityData.YangName = "identifier"
    identifier.EntityData.BundleName = "cisco_ios_xr"
    identifier.EntityData.ParentYangName = "evpn-virtual-ethernet-segment"
    identifier.EntityData.SegmentPath = "identifier"
    identifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifier.EntityData.Children = types.NewOrderedMap()
    identifier.EntityData.Leafs = types.NewOrderedMap()
    identifier.EntityData.Leafs.Append("bytes01", types.YLeaf{"Bytes01", identifier.Bytes01})
    identifier.EntityData.Leafs.Append("bytes23", types.YLeaf{"Bytes23", identifier.Bytes23})
    identifier.EntityData.Leafs.Append("bytes45", types.YLeaf{"Bytes45", identifier.Bytes45})
    identifier.EntityData.Leafs.Append("bytes67", types.YLeaf{"Bytes67", identifier.Bytes67})
    identifier.EntityData.Leafs.Append("bytes89", types.YLeaf{"Bytes89", identifier.Bytes89})
    identifier.EntityData.Leafs.Append("type", types.YLeaf{"Type", identifier.Type})

    identifier.EntityData.YListKeys = []string {}

    return &(identifier.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving
// Enter Manual service carving configuration
// submode
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Manual service carving. The type is interface{}.
    Enable interface{}

    // Manual service carving primary,secondary lists.
    ServiceList Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList
}

func (manualServiceCarving *Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving) GetEntityData() *types.CommonEntityData {
    manualServiceCarving.EntityData.YFilter = manualServiceCarving.YFilter
    manualServiceCarving.EntityData.YangName = "manual-service-carving"
    manualServiceCarving.EntityData.BundleName = "cisco_ios_xr"
    manualServiceCarving.EntityData.ParentYangName = "evpn-virtual-ethernet-segment"
    manualServiceCarving.EntityData.SegmentPath = "manual-service-carving"
    manualServiceCarving.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualServiceCarving.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualServiceCarving.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualServiceCarving.EntityData.Children = types.NewOrderedMap()
    manualServiceCarving.EntityData.Children.Append("service-list", types.YChild{"ServiceList", &manualServiceCarving.ServiceList})
    manualServiceCarving.EntityData.Leafs = types.NewOrderedMap()
    manualServiceCarving.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", manualServiceCarving.Enable})

    manualServiceCarving.EntityData.YListKeys = []string {}

    return &(manualServiceCarving.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList
// Manual service carving primary,secondary lists
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary services list. The type is string with length: 1..150.
    Primary interface{}

    // Secondary services list. The type is string with length: 1..150.
    Secondary interface{}
}

func (serviceList *Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_ManualServiceCarving_ServiceList) GetEntityData() *types.CommonEntityData {
    serviceList.EntityData.YFilter = serviceList.YFilter
    serviceList.EntityData.YangName = "service-list"
    serviceList.EntityData.BundleName = "cisco_ios_xr"
    serviceList.EntityData.ParentYangName = "manual-service-carving"
    serviceList.EntityData.SegmentPath = "service-list"
    serviceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceList.EntityData.Children = types.NewOrderedMap()
    serviceList.EntityData.Leafs = types.NewOrderedMap()
    serviceList.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", serviceList.Primary})
    serviceList.EntityData.Leafs.Append("secondary", types.YLeaf{"Secondary", serviceList.Secondary})

    serviceList.EntityData.YListKeys = []string {}

    return &(serviceList.EntityData)
}

// Evpn_EvpnTables_EvpnEthernetSegment
// EVPN Global Ethernet Segment submode
type Evpn_EvpnTables_EvpnEthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable EVPN Global Ethernet Segment submode. The type is interface{}.
    Enable interface{}

    // EVPN ESI type table.
    EvpnEsiTypes Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes
}

func (evpnEthernetSegment *Evpn_EvpnTables_EvpnEthernetSegment) GetEntityData() *types.CommonEntityData {
    evpnEthernetSegment.EntityData.YFilter = evpnEthernetSegment.YFilter
    evpnEthernetSegment.EntityData.YangName = "evpn-ethernet-segment"
    evpnEthernetSegment.EntityData.BundleName = "cisco_ios_xr"
    evpnEthernetSegment.EntityData.ParentYangName = "evpn-tables"
    evpnEthernetSegment.EntityData.SegmentPath = "evpn-ethernet-segment"
    evpnEthernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEthernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEthernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEthernetSegment.EntityData.Children = types.NewOrderedMap()
    evpnEthernetSegment.EntityData.Children.Append("evpn-esi-types", types.YChild{"EvpnEsiTypes", &evpnEthernetSegment.EvpnEsiTypes})
    evpnEthernetSegment.EntityData.Leafs = types.NewOrderedMap()
    evpnEthernetSegment.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", evpnEthernetSegment.Enable})

    evpnEthernetSegment.EntityData.YListKeys = []string {}

    return &(evpnEthernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes
// EVPN ESI type table
type Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ESI type. The type is slice of
    // Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType.
    EvpnEsiType []*Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType
}

func (evpnEsiTypes *Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes) GetEntityData() *types.CommonEntityData {
    evpnEsiTypes.EntityData.YFilter = evpnEsiTypes.YFilter
    evpnEsiTypes.EntityData.YangName = "evpn-esi-types"
    evpnEsiTypes.EntityData.BundleName = "cisco_ios_xr"
    evpnEsiTypes.EntityData.ParentYangName = "evpn-ethernet-segment"
    evpnEsiTypes.EntityData.SegmentPath = "evpn-esi-types"
    evpnEsiTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEsiTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEsiTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEsiTypes.EntityData.Children = types.NewOrderedMap()
    evpnEsiTypes.EntityData.Children.Append("evpn-esi-type", types.YChild{"EvpnEsiType", nil})
    for i := range evpnEsiTypes.EvpnEsiType {
        evpnEsiTypes.EntityData.Children.Append(types.GetSegmentPath(evpnEsiTypes.EvpnEsiType[i]), types.YChild{"EvpnEsiType", evpnEsiTypes.EvpnEsiType[i]})
    }
    evpnEsiTypes.EntityData.Leafs = types.NewOrderedMap()

    evpnEsiTypes.EntityData.YListKeys = []string {}

    return &(evpnEsiTypes.EntityData)
}

// Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType
// ESI type
type Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ESI type. The type is interface{} with range:
    // 0..4294967295.
    EsiType interface{}

    // Disable ESI Autogeneration. The type is interface{}.
    DisableAutoGeneration interface{}
}

func (evpnEsiType *Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType) GetEntityData() *types.CommonEntityData {
    evpnEsiType.EntityData.YFilter = evpnEsiType.YFilter
    evpnEsiType.EntityData.YangName = "evpn-esi-type"
    evpnEsiType.EntityData.BundleName = "cisco_ios_xr"
    evpnEsiType.EntityData.ParentYangName = "evpn-esi-types"
    evpnEsiType.EntityData.SegmentPath = "evpn-esi-type" + types.AddKeyToken(evpnEsiType.EsiType, "esi-type")
    evpnEsiType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEsiType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEsiType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEsiType.EntityData.Children = types.NewOrderedMap()
    evpnEsiType.EntityData.Leafs = types.NewOrderedMap()
    evpnEsiType.EntityData.Leafs.Append("esi-type", types.YLeaf{"EsiType", evpnEsiType.EsiType})
    evpnEsiType.EntityData.Leafs.Append("disable-auto-generation", types.YLeaf{"DisableAutoGeneration", evpnEsiType.DisableAutoGeneration})

    evpnEsiType.EntityData.YListKeys = []string {"EsiType"}

    return &(evpnEsiType.EntityData)
}

