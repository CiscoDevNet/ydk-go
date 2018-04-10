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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2vpn-cfg l2vpn}", reflect.TypeOf(L2Vpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2vpn-cfg:l2vpn", reflect.TypeOf(L2Vpn{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2vpn-cfg generic-interface-lists}", reflect.TypeOf(GenericInterfaceLists{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2vpn-cfg:generic-interface-lists", reflect.TypeOf(GenericInterfaceLists{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2vpn-cfg evpn}", reflect.TypeOf(Evpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2vpn-cfg:evpn", reflect.TypeOf(Evpn{}))
}

// EvpnEncapsulation represents Evpn encapsulation
type EvpnEncapsulation string

const (
    // VXLAN Encapsulation
    EvpnEncapsulation_evpn_encapsulationvxlan EvpnEncapsulation = "evpn-encapsulationvxlan"

    // MPLS Encapsulation
    EvpnEncapsulation_evpn_encapsulation_mpls EvpnEncapsulation = "evpn-encapsulation-mpls"
)

// Interworking represents Interworking
type Interworking string

const (
    // Ethernet interworking
    Interworking_ethernet Interworking = "ethernet"

    // IPv4 interworking
    Interworking_ipv4 Interworking = "ipv4"
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

// EvpnSide represents Evpn side
type EvpnSide string

const (
    // EVPN Instance side defined as stitching
    EvpnSide_evpn_side_stitching EvpnSide = "evpn-side-stitching"
)

// BridgeDomainTransportMode represents Bridge domain transport mode
type BridgeDomainTransportMode string

const (
    // Vlan tagged passthrough mode
    BridgeDomainTransportMode_vlan_passthrough BridgeDomainTransportMode = "vlan-passthrough"
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

// VccvVerification represents Vccv verification
type VccvVerification string

const (
    // No connectivity verification over VCCV
    VccvVerification_none VccvVerification = "none"

    // LSP Ping over VCCV
    VccvVerification_lsp_ping VccvVerification = "lsp-ping"
)

// MacWithdrawBehavior represents Mac withdraw behavior
type MacWithdrawBehavior string

const (
    // MAC Withdrawal sent on state-down (legacy)
    MacWithdrawBehavior_legacy MacWithdrawBehavior = "legacy"

    // Optimized MAC Withdrawal
    MacWithdrawBehavior_optimized MacWithdrawBehavior = "optimized"
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

// TypeOfServiceMode represents Type of service mode
type TypeOfServiceMode string

const (
    // Do not reflect the type of service
    TypeOfServiceMode_none TypeOfServiceMode = "none"

    // Reflect the type of service
    TypeOfServiceMode_reflect TypeOfServiceMode = "reflect"
)

// MplsSignalingProtocol represents Mpls signaling protocol
type MplsSignalingProtocol string

const (
    // No signaling
    MplsSignalingProtocol_none MplsSignalingProtocol = "none"

    // LDP
    MplsSignalingProtocol_ldp MplsSignalingProtocol = "ldp"
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

// BackupDisable represents Backup disable
type BackupDisable string

const (
    // Never
    BackupDisable_never BackupDisable = "never"

    // Delay seconds
    BackupDisable_delay BackupDisable = "delay"
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

// FlowLabelTlvCode represents Flow label tlv code
type FlowLabelTlvCode string

const (
    // Set Flow Label Legacy TLV code (DEPRECATED)
    FlowLabelTlvCode_Y_17 FlowLabelTlvCode = "17"

    // Disable Sending Flow Label Legacy TLV
    FlowLabelTlvCode_disable FlowLabelTlvCode = "disable"
)

// BgpRouteTarget represents Bgp route target
type BgpRouteTarget string

const (
    // RT is default type
    BgpRouteTarget_no_stitching BgpRouteTarget = "no-stitching"

    // RT is for stitching (Golf-L2)
    BgpRouteTarget_stitching BgpRouteTarget = "stitching"
)

// InterfaceProfile represents Interface profile
type InterfaceProfile string

const (
    // Set the snooping
    InterfaceProfile_snoop InterfaceProfile = "snoop"

    // disable DHCP protocol
    InterfaceProfile_dhcp_protocol InterfaceProfile = "dhcp-protocol"
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

// PwSwitchingPointTlv represents Pw switching point tlv
type PwSwitchingPointTlv string

const (
    // Hide TLV
    PwSwitchingPointTlv_hide PwSwitchingPointTlv = "hide"
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

// L2tpv3Sequencing represents L2tpv3 sequencing
type L2tpv3Sequencing string

const (
    // Sequencing is off
    L2tpv3Sequencing_off L2tpv3Sequencing = "off"

    // Sequencing on both transmit and receive side
    L2tpv3Sequencing_both L2tpv3Sequencing = "both"
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

// L2tpSignalingProtocol represents L2tp signaling protocol
type L2tpSignalingProtocol string

const (
    // No signaling
    L2tpSignalingProtocol_none L2tpSignalingProtocol = "none"

    // L2TPv3
    L2tpSignalingProtocol_l2tpv3 L2tpSignalingProtocol = "l2tpv3"
)

// BdmacLearn represents Bdmac learn
type BdmacLearn string

const (
    // Disable Learning
    BdmacLearn_disable_learning BdmacLearn = "disable-learning"
)

// L2vpnVerification represents L2vpn verification
type L2vpnVerification string

const (
    // enable verification
    L2vpnVerification_enable L2vpnVerification = "enable"

    // disable verification
    L2vpnVerification_disable L2vpnVerification = "disable"
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

// L2Encapsulation represents L2 encapsulation
type L2Encapsulation string

const (
    // Vlan tagged mode
    L2Encapsulation_vlan L2Encapsulation = "vlan"

    // Ethernet port mode
    L2Encapsulation_ethernet L2Encapsulation = "ethernet"
)

// L2vpnLogging represents L2vpn logging
type L2vpnLogging string

const (
    // enable logging
    L2vpnLogging_enable L2vpnLogging = "enable"

    // disable logging
    L2vpnLogging_disable L2vpnLogging = "disable"
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

// ErpPort1 represents Erp port1
type ErpPort1 string

const (
    // ERP main port 0
    ErpPort1_port0 ErpPort1 = "port0"

    // ERP main port 1
    ErpPort1_port1 ErpPort1 = "port1"
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

// ControlWord represents Control word
type ControlWord string

const (
    // Enable control word
    ControlWord_enable ControlWord = "enable"

    // Disable control word
    ControlWord_disable ControlWord = "disable"
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

// LdpVplsId represents Ldp vpls id
type LdpVplsId string

const (
    // VPLS-ID in 2 byte AS:nn format
    LdpVplsId_two_byte_as LdpVplsId = "two-byte-as"

    // VPLS-ID in IPv4 IP:nn format
    LdpVplsId_ipv4_address LdpVplsId = "ipv4-address"
)

// MacAging represents Mac aging
type MacAging string

const (
    // Absolute aging type
    MacAging_absolute MacAging = "absolute"

    // Inactivity aging type
    MacAging_inactivity MacAging = "inactivity"
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

    // Insert/Delete  Flow Label on transmit/receive
    // side
    FlowLabelLoadBalance_both FlowLabelLoadBalance = "both"
)

// L2Vpn
// L2VPN configuration
type L2Vpn struct {
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    L2VpnRouterId interface{}

    // Pseudowire-routing attributes.
    PwRouting L2Vpn_PwRouting

    // L2VPN neighbor submode.
    Neighbor L2Vpn_Neighbor

    // L2VPN databases.
    Database L2Vpn_Database

    // L2VPN PBB Global.
    Pbb L2Vpn_Pbb

    // Global auto-discovery attributes.
    AutoDiscovery L2Vpn_AutoDiscovery

    // L2VPN utilities.
    Utility L2Vpn_Utility

    // SNMP related configuration.
    Snmp L2Vpn_Snmp
}

func (l2Vpn *L2Vpn) GetEntityData() *types.CommonEntityData {
    l2Vpn.EntityData.YFilter = l2Vpn.YFilter
    l2Vpn.EntityData.YangName = "l2vpn"
    l2Vpn.EntityData.BundleName = "cisco_ios_xr"
    l2Vpn.EntityData.ParentYangName = "Cisco-IOS-XR-l2vpn-cfg"
    l2Vpn.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:l2vpn"
    l2Vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Vpn.EntityData.Children = make(map[string]types.YChild)
    l2Vpn.EntityData.Children["pw-routing"] = types.YChild{"PwRouting", &l2Vpn.PwRouting}
    l2Vpn.EntityData.Children["neighbor"] = types.YChild{"Neighbor", &l2Vpn.Neighbor}
    l2Vpn.EntityData.Children["database"] = types.YChild{"Database", &l2Vpn.Database}
    l2Vpn.EntityData.Children["pbb"] = types.YChild{"Pbb", &l2Vpn.Pbb}
    l2Vpn.EntityData.Children["auto-discovery"] = types.YChild{"AutoDiscovery", &l2Vpn.AutoDiscovery}
    l2Vpn.EntityData.Children["utility"] = types.YChild{"Utility", &l2Vpn.Utility}
    l2Vpn.EntityData.Children["snmp"] = types.YChild{"Snmp", &l2Vpn.Snmp}
    l2Vpn.EntityData.Leafs = make(map[string]types.YLeaf)
    l2Vpn.EntityData.Leafs["nsr"] = types.YLeaf{"Nsr", l2Vpn.Nsr}
    l2Vpn.EntityData.Leafs["mtu-mismatch-ignore"] = types.YLeaf{"MtuMismatchIgnore", l2Vpn.MtuMismatchIgnore}
    l2Vpn.EntityData.Leafs["tcn-propagation"] = types.YLeaf{"TcnPropagation", l2Vpn.TcnPropagation}
    l2Vpn.EntityData.Leafs["pwoam-refresh"] = types.YLeaf{"PwoamRefresh", l2Vpn.PwoamRefresh}
    l2Vpn.EntityData.Leafs["load-balance"] = types.YLeaf{"LoadBalance", l2Vpn.LoadBalance}
    l2Vpn.EntityData.Leafs["mspw-description"] = types.YLeaf{"MspwDescription", l2Vpn.MspwDescription}
    l2Vpn.EntityData.Leafs["mac-limit-threshold"] = types.YLeaf{"MacLimitThreshold", l2Vpn.MacLimitThreshold}
    l2Vpn.EntityData.Leafs["pw-status-disable"] = types.YLeaf{"PwStatusDisable", l2Vpn.PwStatusDisable}
    l2Vpn.EntityData.Leafs["enable"] = types.YLeaf{"Enable", l2Vpn.Enable}
    l2Vpn.EntityData.Leafs["pw-grouping"] = types.YLeaf{"PwGrouping", l2Vpn.PwGrouping}
    l2Vpn.EntityData.Leafs["capability"] = types.YLeaf{"Capability", l2Vpn.Capability}
    l2Vpn.EntityData.Leafs["l2vpn-router-id"] = types.YLeaf{"L2VpnRouterId", l2Vpn.L2VpnRouterId}
    return &(l2Vpn.EntityData)
}

// L2Vpn_PwRouting
// Pseudowire-routing attributes
type L2Vpn_PwRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire-routing Global ID. The type is interface{} with range:
    // 1..4294967295.
    PwRoutingGlobalId interface{}

    // Enable Autodiscovery BGP Pseudowire-routing BGP.
    PwRoutingBgp L2Vpn_PwRouting_PwRoutingBgp
}

func (pwRouting *L2Vpn_PwRouting) GetEntityData() *types.CommonEntityData {
    pwRouting.EntityData.YFilter = pwRouting.YFilter
    pwRouting.EntityData.YangName = "pw-routing"
    pwRouting.EntityData.BundleName = "cisco_ios_xr"
    pwRouting.EntityData.ParentYangName = "l2vpn"
    pwRouting.EntityData.SegmentPath = "pw-routing"
    pwRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pwRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pwRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pwRouting.EntityData.Children = make(map[string]types.YChild)
    pwRouting.EntityData.Children["pw-routing-bgp"] = types.YChild{"PwRoutingBgp", &pwRouting.PwRoutingBgp}
    pwRouting.EntityData.Leafs = make(map[string]types.YLeaf)
    pwRouting.EntityData.Leafs["pw-routing-global-id"] = types.YLeaf{"PwRoutingGlobalId", pwRouting.PwRoutingGlobalId}
    return &(pwRouting.EntityData)
}

// L2Vpn_PwRouting_PwRoutingBgp
// Enable Autodiscovery BGP Pseudowire-routing BGP
type L2Vpn_PwRouting_PwRoutingBgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Route Distinguisher.
    EvpnRouteDistinguisher L2Vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher
}

func (pwRoutingBgp *L2Vpn_PwRouting_PwRoutingBgp) GetEntityData() *types.CommonEntityData {
    pwRoutingBgp.EntityData.YFilter = pwRoutingBgp.YFilter
    pwRoutingBgp.EntityData.YangName = "pw-routing-bgp"
    pwRoutingBgp.EntityData.BundleName = "cisco_ios_xr"
    pwRoutingBgp.EntityData.ParentYangName = "pw-routing"
    pwRoutingBgp.EntityData.SegmentPath = "pw-routing-bgp"
    pwRoutingBgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pwRoutingBgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pwRoutingBgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pwRoutingBgp.EntityData.Children = make(map[string]types.YChild)
    pwRoutingBgp.EntityData.Children["evpn-route-distinguisher"] = types.YChild{"EvpnRouteDistinguisher", &pwRoutingBgp.EvpnRouteDistinguisher}
    pwRoutingBgp.EntityData.Leafs = make(map[string]types.YLeaf)
    pwRoutingBgp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pwRoutingBgp.Enable}
    return &(pwRoutingBgp.EntityData)
}

// L2Vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher
// Route Distinguisher
type L2Vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type_ interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (evpnRouteDistinguisher *L2Vpn_PwRouting_PwRoutingBgp_EvpnRouteDistinguisher) GetEntityData() *types.CommonEntityData {
    evpnRouteDistinguisher.EntityData.YFilter = evpnRouteDistinguisher.YFilter
    evpnRouteDistinguisher.EntityData.YangName = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteDistinguisher.EntityData.ParentYangName = "pw-routing-bgp"
    evpnRouteDistinguisher.EntityData.SegmentPath = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteDistinguisher.EntityData.Children = make(map[string]types.YChild)
    evpnRouteDistinguisher.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteDistinguisher.EntityData.Leafs["type"] = types.YLeaf{"Type_", evpnRouteDistinguisher.Type_}
    evpnRouteDistinguisher.EntityData.Leafs["as"] = types.YLeaf{"As", evpnRouteDistinguisher.As}
    evpnRouteDistinguisher.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex}
    evpnRouteDistinguisher.EntityData.Leafs["address"] = types.YLeaf{"Address", evpnRouteDistinguisher.Address}
    evpnRouteDistinguisher.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex}
    return &(evpnRouteDistinguisher.EntityData)
}

// L2Vpn_Neighbor
// L2VPN neighbor submode
type L2Vpn_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable targetted LDP session flap action. The type is interface{}.
    LdpFlap interface{}
}

func (neighbor *L2Vpn_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "l2vpn"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["ldp-flap"] = types.YLeaf{"LdpFlap", neighbor.LdpFlap}
    return &(neighbor.EntityData)
}

// L2Vpn_Database
// L2VPN databases
type L2Vpn_Database struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of G8032 Ring.
    G8032Rings L2Vpn_Database_G8032Rings

    // List of xconnect groups.
    XconnectGroups L2Vpn_Database_XconnectGroups

    // List of bridge  groups.
    BridgeDomainGroups L2Vpn_Database_BridgeDomainGroups

    // List of pseudowire classes.
    PseudowireClasses L2Vpn_Database_PseudowireClasses

    // List of Flexible XConnect Services.
    FlexibleXconnectServiceTable L2Vpn_Database_FlexibleXconnectServiceTable

    // Redundancy groups.
    Redundancy L2Vpn_Database_Redundancy
}

func (database *L2Vpn_Database) GetEntityData() *types.CommonEntityData {
    database.EntityData.YFilter = database.YFilter
    database.EntityData.YangName = "database"
    database.EntityData.BundleName = "cisco_ios_xr"
    database.EntityData.ParentYangName = "l2vpn"
    database.EntityData.SegmentPath = "database"
    database.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    database.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    database.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    database.EntityData.Children = make(map[string]types.YChild)
    database.EntityData.Children["g8032-rings"] = types.YChild{"G8032Rings", &database.G8032Rings}
    database.EntityData.Children["xconnect-groups"] = types.YChild{"XconnectGroups", &database.XconnectGroups}
    database.EntityData.Children["bridge-domain-groups"] = types.YChild{"BridgeDomainGroups", &database.BridgeDomainGroups}
    database.EntityData.Children["pseudowire-classes"] = types.YChild{"PseudowireClasses", &database.PseudowireClasses}
    database.EntityData.Children["flexible-xconnect-service-table"] = types.YChild{"FlexibleXconnectServiceTable", &database.FlexibleXconnectServiceTable}
    database.EntityData.Children["redundancy"] = types.YChild{"Redundancy", &database.Redundancy}
    database.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(database.EntityData)
}

// L2Vpn_Database_G8032Rings
// List of G8032 Ring
type L2Vpn_Database_G8032Rings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // G8032 Ring. The type is slice of L2Vpn_Database_G8032Rings_G8032Ring.
    G8032Ring []L2Vpn_Database_G8032Rings_G8032Ring
}

func (g8032Rings *L2Vpn_Database_G8032Rings) GetEntityData() *types.CommonEntityData {
    g8032Rings.EntityData.YFilter = g8032Rings.YFilter
    g8032Rings.EntityData.YangName = "g8032-rings"
    g8032Rings.EntityData.BundleName = "cisco_ios_xr"
    g8032Rings.EntityData.ParentYangName = "database"
    g8032Rings.EntityData.SegmentPath = "g8032-rings"
    g8032Rings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    g8032Rings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    g8032Rings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    g8032Rings.EntityData.Children = make(map[string]types.YChild)
    g8032Rings.EntityData.Children["g8032-ring"] = types.YChild{"G8032Ring", nil}
    for i := range g8032Rings.G8032Ring {
        g8032Rings.EntityData.Children[types.GetSegmentPath(&g8032Rings.G8032Ring[i])] = types.YChild{"G8032Ring", &g8032Rings.G8032Ring[i]}
    }
    g8032Rings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(g8032Rings.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring
// G8032 Ring
type L2Vpn_Database_G8032Rings_G8032Ring struct {
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
    ErpPort0S L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S

    // List of ethernet ring protection instance.
    ErpInstances L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances

    // Ethernet ring protection port0.
    ErpPort1S L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S
}

func (g8032Ring *L2Vpn_Database_G8032Rings_G8032Ring) GetEntityData() *types.CommonEntityData {
    g8032Ring.EntityData.YFilter = g8032Ring.YFilter
    g8032Ring.EntityData.YangName = "g8032-ring"
    g8032Ring.EntityData.BundleName = "cisco_ios_xr"
    g8032Ring.EntityData.ParentYangName = "g8032-rings"
    g8032Ring.EntityData.SegmentPath = "g8032-ring" + "[g8032-ring-name='" + fmt.Sprintf("%v", g8032Ring.G8032RingName) + "']"
    g8032Ring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    g8032Ring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    g8032Ring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    g8032Ring.EntityData.Children = make(map[string]types.YChild)
    g8032Ring.EntityData.Children["erp-port0s"] = types.YChild{"ErpPort0S", &g8032Ring.ErpPort0S}
    g8032Ring.EntityData.Children["erp-instances"] = types.YChild{"ErpInstances", &g8032Ring.ErpInstances}
    g8032Ring.EntityData.Children["erp-port1s"] = types.YChild{"ErpPort1S", &g8032Ring.ErpPort1S}
    g8032Ring.EntityData.Leafs = make(map[string]types.YLeaf)
    g8032Ring.EntityData.Leafs["g8032-ring-name"] = types.YLeaf{"G8032RingName", g8032Ring.G8032RingName}
    g8032Ring.EntityData.Leafs["open-ring"] = types.YLeaf{"OpenRing", g8032Ring.OpenRing}
    g8032Ring.EntityData.Leafs["exclusion-list"] = types.YLeaf{"ExclusionList", g8032Ring.ExclusionList}
    g8032Ring.EntityData.Leafs["erp-provider-bridge"] = types.YLeaf{"ErpProviderBridge", g8032Ring.ErpProviderBridge}
    return &(g8032Ring.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S
// Ethernet ring protection port0
type L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure ERP main port0. The type is slice of
    // L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S_ErpPort0.
    ErpPort0 []L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S_ErpPort0
}

func (erpPort0S *L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S) GetEntityData() *types.CommonEntityData {
    erpPort0S.EntityData.YFilter = erpPort0S.YFilter
    erpPort0S.EntityData.YangName = "erp-port0s"
    erpPort0S.EntityData.BundleName = "cisco_ios_xr"
    erpPort0S.EntityData.ParentYangName = "g8032-ring"
    erpPort0S.EntityData.SegmentPath = "erp-port0s"
    erpPort0S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort0S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort0S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort0S.EntityData.Children = make(map[string]types.YChild)
    erpPort0S.EntityData.Children["erp-port0"] = types.YChild{"ErpPort0", nil}
    for i := range erpPort0S.ErpPort0 {
        erpPort0S.EntityData.Children[types.GetSegmentPath(&erpPort0S.ErpPort0[i])] = types.YChild{"ErpPort0", &erpPort0S.ErpPort0[i]}
    }
    erpPort0S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(erpPort0S.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S_ErpPort0
// Configure ERP main port0
type L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S_ErpPort0 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port0 interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Ethernet ring protection port0 monitor. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Monitor interface{}
}

func (erpPort0 *L2Vpn_Database_G8032Rings_G8032Ring_ErpPort0S_ErpPort0) GetEntityData() *types.CommonEntityData {
    erpPort0.EntityData.YFilter = erpPort0.YFilter
    erpPort0.EntityData.YangName = "erp-port0"
    erpPort0.EntityData.BundleName = "cisco_ios_xr"
    erpPort0.EntityData.ParentYangName = "erp-port0s"
    erpPort0.EntityData.SegmentPath = "erp-port0" + "[interface-name='" + fmt.Sprintf("%v", erpPort0.InterfaceName) + "']"
    erpPort0.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort0.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort0.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort0.EntityData.Children = make(map[string]types.YChild)
    erpPort0.EntityData.Leafs = make(map[string]types.YLeaf)
    erpPort0.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", erpPort0.InterfaceName}
    erpPort0.EntityData.Leafs["monitor"] = types.YLeaf{"Monitor", erpPort0.Monitor}
    return &(erpPort0.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances
// List of ethernet ring protection instance
type L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet ring protection instance. The type is slice of
    // L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance.
    ErpInstance []L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance
}

func (erpInstances *L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances) GetEntityData() *types.CommonEntityData {
    erpInstances.EntityData.YFilter = erpInstances.YFilter
    erpInstances.EntityData.YangName = "erp-instances"
    erpInstances.EntityData.BundleName = "cisco_ios_xr"
    erpInstances.EntityData.ParentYangName = "g8032-ring"
    erpInstances.EntityData.SegmentPath = "erp-instances"
    erpInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpInstances.EntityData.Children = make(map[string]types.YChild)
    erpInstances.EntityData.Children["erp-instance"] = types.YChild{"ErpInstance", nil}
    for i := range erpInstances.ErpInstance {
        erpInstances.EntityData.Children[types.GetSegmentPath(&erpInstances.ErpInstance[i])] = types.YChild{"ErpInstance", &erpInstances.ErpInstance[i]}
    }
    erpInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(erpInstances.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance
// Ethernet ring protection instance
type L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance struct {
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
    Rpl L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl

    // Automatic protection switching.
    Aps L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps
}

func (erpInstance *L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance) GetEntityData() *types.CommonEntityData {
    erpInstance.EntityData.YFilter = erpInstance.YFilter
    erpInstance.EntityData.YangName = "erp-instance"
    erpInstance.EntityData.BundleName = "cisco_ios_xr"
    erpInstance.EntityData.ParentYangName = "erp-instances"
    erpInstance.EntityData.SegmentPath = "erp-instance" + "[erp-instance-id='" + fmt.Sprintf("%v", erpInstance.ErpInstanceId) + "']"
    erpInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpInstance.EntityData.Children = make(map[string]types.YChild)
    erpInstance.EntityData.Children["rpl"] = types.YChild{"Rpl", &erpInstance.Rpl}
    erpInstance.EntityData.Children["aps"] = types.YChild{"Aps", &erpInstance.Aps}
    erpInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    erpInstance.EntityData.Leafs["erp-instance-id"] = types.YLeaf{"ErpInstanceId", erpInstance.ErpInstanceId}
    erpInstance.EntityData.Leafs["description"] = types.YLeaf{"Description", erpInstance.Description}
    erpInstance.EntityData.Leafs["inclusion-list"] = types.YLeaf{"InclusionList", erpInstance.InclusionList}
    erpInstance.EntityData.Leafs["profile"] = types.YLeaf{"Profile", erpInstance.Profile}
    return &(erpInstance.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl
// Ring protection link
type L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ERP main port number. The type is ErpPort1.
    Port interface{}

    // RPL role. The type is RplRole.
    Role interface{}
}

func (rpl *L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Rpl) GetEntityData() *types.CommonEntityData {
    rpl.EntityData.YFilter = rpl.YFilter
    rpl.EntityData.YangName = "rpl"
    rpl.EntityData.BundleName = "cisco_ios_xr"
    rpl.EntityData.ParentYangName = "erp-instance"
    rpl.EntityData.SegmentPath = "rpl"
    rpl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpl.EntityData.Children = make(map[string]types.YChild)
    rpl.EntityData.Leafs = make(map[string]types.YLeaf)
    rpl.EntityData.Leafs["port"] = types.YLeaf{"Port", rpl.Port}
    rpl.EntityData.Leafs["role"] = types.YLeaf{"Role", rpl.Role}
    return &(rpl.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps
// Automatic protection switching
type L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps struct {
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
    Port1 L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1
}

func (aps *L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps) GetEntityData() *types.CommonEntityData {
    aps.EntityData.YFilter = aps.YFilter
    aps.EntityData.YangName = "aps"
    aps.EntityData.BundleName = "cisco_ios_xr"
    aps.EntityData.ParentYangName = "erp-instance"
    aps.EntityData.SegmentPath = "aps"
    aps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aps.EntityData.Children = make(map[string]types.YChild)
    aps.EntityData.Children["port1"] = types.YChild{"Port1", &aps.Port1}
    aps.EntityData.Leafs = make(map[string]types.YLeaf)
    aps.EntityData.Leafs["port0"] = types.YLeaf{"Port0", aps.Port0}
    aps.EntityData.Leafs["enable"] = types.YLeaf{"Enable", aps.Enable}
    aps.EntityData.Leafs["level"] = types.YLeaf{"Level", aps.Level}
    return &(aps.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1
// APS channel for ERP port1
type L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port1 APS type. The type is Erpaps.
    ApsType interface{}

    // Port1 APS channel in the format of InterfaceName, BDName or XconnectName.
    // The type is string.
    ApsChannel interface{}
}

func (port1 *L2Vpn_Database_G8032Rings_G8032Ring_ErpInstances_ErpInstance_Aps_Port1) GetEntityData() *types.CommonEntityData {
    port1.EntityData.YFilter = port1.YFilter
    port1.EntityData.YangName = "port1"
    port1.EntityData.BundleName = "cisco_ios_xr"
    port1.EntityData.ParentYangName = "aps"
    port1.EntityData.SegmentPath = "port1"
    port1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port1.EntityData.Children = make(map[string]types.YChild)
    port1.EntityData.Leafs = make(map[string]types.YLeaf)
    port1.EntityData.Leafs["aps-type"] = types.YLeaf{"ApsType", port1.ApsType}
    port1.EntityData.Leafs["aps-channel"] = types.YLeaf{"ApsChannel", port1.ApsChannel}
    return &(port1.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S
// Ethernet ring protection port0
type L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet ring protection port1. The type is slice of
    // L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1.
    ErpPort1 []L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1
}

func (erpPort1S *L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S) GetEntityData() *types.CommonEntityData {
    erpPort1S.EntityData.YFilter = erpPort1S.YFilter
    erpPort1S.EntityData.YangName = "erp-port1s"
    erpPort1S.EntityData.BundleName = "cisco_ios_xr"
    erpPort1S.EntityData.ParentYangName = "g8032-ring"
    erpPort1S.EntityData.SegmentPath = "erp-port1s"
    erpPort1S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort1S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort1S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort1S.EntityData.Children = make(map[string]types.YChild)
    erpPort1S.EntityData.Children["erp-port1"] = types.YChild{"ErpPort1", nil}
    for i := range erpPort1S.ErpPort1 {
        erpPort1S.EntityData.Children[types.GetSegmentPath(&erpPort1S.ErpPort1[i])] = types.YChild{"ErpPort1", &erpPort1S.ErpPort1[i]}
    }
    erpPort1S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(erpPort1S.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1
// Ethernet ring protection port1
type L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port1 type. The type is ErpPort.
    ErpPortType interface{}

    // none or virtual.
    NoneOrVirtual L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_NoneOrVirtual

    // interface. The type is slice of
    // L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_Interface_.
    Interface_ []L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_Interface
}

func (erpPort1 *L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1) GetEntityData() *types.CommonEntityData {
    erpPort1.EntityData.YFilter = erpPort1.YFilter
    erpPort1.EntityData.YangName = "erp-port1"
    erpPort1.EntityData.BundleName = "cisco_ios_xr"
    erpPort1.EntityData.ParentYangName = "erp-port1s"
    erpPort1.EntityData.SegmentPath = "erp-port1" + "[erp-port-type='" + fmt.Sprintf("%v", erpPort1.ErpPortType) + "']"
    erpPort1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erpPort1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erpPort1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erpPort1.EntityData.Children = make(map[string]types.YChild)
    erpPort1.EntityData.Children["none-or-virtual"] = types.YChild{"NoneOrVirtual", &erpPort1.NoneOrVirtual}
    erpPort1.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range erpPort1.Interface_ {
        erpPort1.EntityData.Children[types.GetSegmentPath(&erpPort1.Interface_[i])] = types.YChild{"Interface_", &erpPort1.Interface_[i]}
    }
    erpPort1.EntityData.Leafs = make(map[string]types.YLeaf)
    erpPort1.EntityData.Leafs["erp-port-type"] = types.YLeaf{"ErpPortType", erpPort1.ErpPortType}
    return &(erpPort1.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_NoneOrVirtual
// none or virtual
// This type is a presence type.
type L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_NoneOrVirtual struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet ring protection port1 monitor. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Monitor interface{}
}

func (noneOrVirtual *L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_NoneOrVirtual) GetEntityData() *types.CommonEntityData {
    noneOrVirtual.EntityData.YFilter = noneOrVirtual.YFilter
    noneOrVirtual.EntityData.YangName = "none-or-virtual"
    noneOrVirtual.EntityData.BundleName = "cisco_ios_xr"
    noneOrVirtual.EntityData.ParentYangName = "erp-port1"
    noneOrVirtual.EntityData.SegmentPath = "none-or-virtual"
    noneOrVirtual.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    noneOrVirtual.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    noneOrVirtual.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    noneOrVirtual.EntityData.Children = make(map[string]types.YChild)
    noneOrVirtual.EntityData.Leafs = make(map[string]types.YLeaf)
    noneOrVirtual.EntityData.Leafs["monitor"] = types.YLeaf{"Monitor", noneOrVirtual.Monitor}
    return &(noneOrVirtual.EntityData)
}

// L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_Interface
// interface
type L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port1 interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Ethernet ring protection port1 monitor. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Monitor interface{}
}

func (self *L2Vpn_Database_G8032Rings_G8032Ring_ErpPort1S_ErpPort1_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "erp-port1"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["monitor"] = types.YLeaf{"Monitor", self.Monitor}
    return &(self.EntityData)
}

// L2Vpn_Database_XconnectGroups
// List of xconnect groups
type L2Vpn_Database_XconnectGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Xconnect group. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup.
    XconnectGroup []L2Vpn_Database_XconnectGroups_XconnectGroup
}

func (xconnectGroups *L2Vpn_Database_XconnectGroups) GetEntityData() *types.CommonEntityData {
    xconnectGroups.EntityData.YFilter = xconnectGroups.YFilter
    xconnectGroups.EntityData.YangName = "xconnect-groups"
    xconnectGroups.EntityData.BundleName = "cisco_ios_xr"
    xconnectGroups.EntityData.ParentYangName = "database"
    xconnectGroups.EntityData.SegmentPath = "xconnect-groups"
    xconnectGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnectGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnectGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnectGroups.EntityData.Children = make(map[string]types.YChild)
    xconnectGroups.EntityData.Children["xconnect-group"] = types.YChild{"XconnectGroup", nil}
    for i := range xconnectGroups.XconnectGroup {
        xconnectGroups.EntityData.Children[types.GetSegmentPath(&xconnectGroups.XconnectGroup[i])] = types.YChild{"XconnectGroup", &xconnectGroups.XconnectGroup[i]}
    }
    xconnectGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(xconnectGroups.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup
// Xconnect group
type L2Vpn_Database_XconnectGroups_XconnectGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the xconnect group. The type is string
    // with length: 1..32.
    Name interface{}

    // List of point to point xconnects.
    P2PXconnects L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects

    // List of multi point to multi point xconnects.
    Mp2MpXconnects L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects
}

func (xconnectGroup *L2Vpn_Database_XconnectGroups_XconnectGroup) GetEntityData() *types.CommonEntityData {
    xconnectGroup.EntityData.YFilter = xconnectGroup.YFilter
    xconnectGroup.EntityData.YangName = "xconnect-group"
    xconnectGroup.EntityData.BundleName = "cisco_ios_xr"
    xconnectGroup.EntityData.ParentYangName = "xconnect-groups"
    xconnectGroup.EntityData.SegmentPath = "xconnect-group" + "[name='" + fmt.Sprintf("%v", xconnectGroup.Name) + "']"
    xconnectGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xconnectGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xconnectGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xconnectGroup.EntityData.Children = make(map[string]types.YChild)
    xconnectGroup.EntityData.Children["p2p-xconnects"] = types.YChild{"P2PXconnects", &xconnectGroup.P2PXconnects}
    xconnectGroup.EntityData.Children["mp2mp-xconnects"] = types.YChild{"Mp2MpXconnects", &xconnectGroup.Mp2MpXconnects}
    xconnectGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    xconnectGroup.EntityData.Leafs["name"] = types.YLeaf{"Name", xconnectGroup.Name}
    return &(xconnectGroup.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects
// List of point to point xconnects
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Point to point xconnect. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect.
    P2PXconnect []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect
}

func (p2PXconnects *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects) GetEntityData() *types.CommonEntityData {
    p2PXconnects.EntityData.YFilter = p2PXconnects.YFilter
    p2PXconnects.EntityData.YangName = "p2p-xconnects"
    p2PXconnects.EntityData.BundleName = "cisco_ios_xr"
    p2PXconnects.EntityData.ParentYangName = "xconnect-group"
    p2PXconnects.EntityData.SegmentPath = "p2p-xconnects"
    p2PXconnects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2PXconnects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2PXconnects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2PXconnects.EntityData.Children = make(map[string]types.YChild)
    p2PXconnects.EntityData.Children["p2p-xconnect"] = types.YChild{"P2PXconnect", nil}
    for i := range p2PXconnects.P2PXconnect {
        p2PXconnects.EntityData.Children[types.GetSegmentPath(&p2PXconnects.P2PXconnect[i])] = types.YChild{"P2PXconnect", &p2PXconnects.P2PXconnect[i]}
    }
    p2PXconnects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(p2PXconnects.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect
// Point to point xconnect
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the point to point xconnect. The type is
    // string with length: 1..38.
    Name interface{}

    // cross connect description Name. The type is string with length: 1..64.
    P2PDescription interface{}

    // Interworking. The type is Interworking.
    Interworking interface{}

    // List of backup attachment circuits.
    BackupAttachmentCircuits L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits

    // List of EVPN Services.
    PseudowireEvpns L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns

    // List of pseudowires.
    Pseudowires L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires

    // List of Monitor session segments.
    MonitorSessions L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions

    // List of pseudowire-routed.
    PseudowireRouteds L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds

    // List of attachment circuits.
    AttachmentCircuits L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits
}

func (p2PXconnect *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect) GetEntityData() *types.CommonEntityData {
    p2PXconnect.EntityData.YFilter = p2PXconnect.YFilter
    p2PXconnect.EntityData.YangName = "p2p-xconnect"
    p2PXconnect.EntityData.BundleName = "cisco_ios_xr"
    p2PXconnect.EntityData.ParentYangName = "p2p-xconnects"
    p2PXconnect.EntityData.SegmentPath = "p2p-xconnect" + "[name='" + fmt.Sprintf("%v", p2PXconnect.Name) + "']"
    p2PXconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2PXconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2PXconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2PXconnect.EntityData.Children = make(map[string]types.YChild)
    p2PXconnect.EntityData.Children["backup-attachment-circuits"] = types.YChild{"BackupAttachmentCircuits", &p2PXconnect.BackupAttachmentCircuits}
    p2PXconnect.EntityData.Children["pseudowire-evpns"] = types.YChild{"PseudowireEvpns", &p2PXconnect.PseudowireEvpns}
    p2PXconnect.EntityData.Children["pseudowires"] = types.YChild{"Pseudowires", &p2PXconnect.Pseudowires}
    p2PXconnect.EntityData.Children["monitor-sessions"] = types.YChild{"MonitorSessions", &p2PXconnect.MonitorSessions}
    p2PXconnect.EntityData.Children["pseudowire-routeds"] = types.YChild{"PseudowireRouteds", &p2PXconnect.PseudowireRouteds}
    p2PXconnect.EntityData.Children["attachment-circuits"] = types.YChild{"AttachmentCircuits", &p2PXconnect.AttachmentCircuits}
    p2PXconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    p2PXconnect.EntityData.Leafs["name"] = types.YLeaf{"Name", p2PXconnect.Name}
    p2PXconnect.EntityData.Leafs["p2p-description"] = types.YLeaf{"P2PDescription", p2PXconnect.P2PDescription}
    p2PXconnect.EntityData.Leafs["interworking"] = types.YLeaf{"Interworking", p2PXconnect.Interworking}
    return &(p2PXconnect.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits
// List of backup attachment circuits
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup attachment circuit. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit.
    BackupAttachmentCircuit []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit
}

func (backupAttachmentCircuits *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    backupAttachmentCircuits.EntityData.YFilter = backupAttachmentCircuits.YFilter
    backupAttachmentCircuits.EntityData.YangName = "backup-attachment-circuits"
    backupAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    backupAttachmentCircuits.EntityData.ParentYangName = "p2p-xconnect"
    backupAttachmentCircuits.EntityData.SegmentPath = "backup-attachment-circuits"
    backupAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupAttachmentCircuits.EntityData.Children = make(map[string]types.YChild)
    backupAttachmentCircuits.EntityData.Children["backup-attachment-circuit"] = types.YChild{"BackupAttachmentCircuit", nil}
    for i := range backupAttachmentCircuits.BackupAttachmentCircuit {
        backupAttachmentCircuits.EntityData.Children[types.GetSegmentPath(&backupAttachmentCircuits.BackupAttachmentCircuit[i])] = types.YChild{"BackupAttachmentCircuit", &backupAttachmentCircuits.BackupAttachmentCircuit[i]}
    }
    backupAttachmentCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backupAttachmentCircuits.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit
// Backup attachment circuit
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (backupAttachmentCircuit *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_BackupAttachmentCircuits_BackupAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    backupAttachmentCircuit.EntityData.YFilter = backupAttachmentCircuit.YFilter
    backupAttachmentCircuit.EntityData.YangName = "backup-attachment-circuit"
    backupAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    backupAttachmentCircuit.EntityData.ParentYangName = "backup-attachment-circuits"
    backupAttachmentCircuit.EntityData.SegmentPath = "backup-attachment-circuit" + "[interface-name='" + fmt.Sprintf("%v", backupAttachmentCircuit.InterfaceName) + "']"
    backupAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupAttachmentCircuit.EntityData.Children = make(map[string]types.YChild)
    backupAttachmentCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    backupAttachmentCircuit.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", backupAttachmentCircuit.InterfaceName}
    return &(backupAttachmentCircuit.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns
// List of EVPN Services
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN P2P Service Configuration. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns_PseudowireEvpn.
    PseudowireEvpn []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns_PseudowireEvpn
}

func (pseudowireEvpns *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns) GetEntityData() *types.CommonEntityData {
    pseudowireEvpns.EntityData.YFilter = pseudowireEvpns.YFilter
    pseudowireEvpns.EntityData.YangName = "pseudowire-evpns"
    pseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEvpns.EntityData.ParentYangName = "p2p-xconnect"
    pseudowireEvpns.EntityData.SegmentPath = "pseudowire-evpns"
    pseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEvpns.EntityData.Children = make(map[string]types.YChild)
    pseudowireEvpns.EntityData.Children["pseudowire-evpn"] = types.YChild{"PseudowireEvpn", nil}
    for i := range pseudowireEvpns.PseudowireEvpn {
        pseudowireEvpns.EntityData.Children[types.GetSegmentPath(&pseudowireEvpns.PseudowireEvpn[i])] = types.YChild{"PseudowireEvpn", &pseudowireEvpns.PseudowireEvpn[i]}
    }
    pseudowireEvpns.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pseudowireEvpns.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns_PseudowireEvpn
// EVPN P2P Service Configuration
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns_PseudowireEvpn struct {
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

func (pseudowireEvpn *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireEvpns_PseudowireEvpn) GetEntityData() *types.CommonEntityData {
    pseudowireEvpn.EntityData.YFilter = pseudowireEvpn.YFilter
    pseudowireEvpn.EntityData.YangName = "pseudowire-evpn"
    pseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEvpn.EntityData.ParentYangName = "pseudowire-evpns"
    pseudowireEvpn.EntityData.SegmentPath = "pseudowire-evpn" + "[eviid='" + fmt.Sprintf("%v", pseudowireEvpn.Eviid) + "']" + "[remote-acid='" + fmt.Sprintf("%v", pseudowireEvpn.RemoteAcid) + "']" + "[source-acid='" + fmt.Sprintf("%v", pseudowireEvpn.SourceAcid) + "']"
    pseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEvpn.EntityData.Children = make(map[string]types.YChild)
    pseudowireEvpn.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireEvpn.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", pseudowireEvpn.Eviid}
    pseudowireEvpn.EntityData.Leafs["remote-acid"] = types.YLeaf{"RemoteAcid", pseudowireEvpn.RemoteAcid}
    pseudowireEvpn.EntityData.Leafs["source-acid"] = types.YLeaf{"SourceAcid", pseudowireEvpn.SourceAcid}
    pseudowireEvpn.EntityData.Leafs["class"] = types.YLeaf{"Class", pseudowireEvpn.Class}
    return &(pseudowireEvpn.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires
// List of pseudowires
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire.
    Pseudowire []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire
}

func (pseudowires *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires) GetEntityData() *types.CommonEntityData {
    pseudowires.EntityData.YFilter = pseudowires.YFilter
    pseudowires.EntityData.YangName = "pseudowires"
    pseudowires.EntityData.BundleName = "cisco_ios_xr"
    pseudowires.EntityData.ParentYangName = "p2p-xconnect"
    pseudowires.EntityData.SegmentPath = "pseudowires"
    pseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowires.EntityData.Children = make(map[string]types.YChild)
    pseudowires.EntityData.Children["pseudowire"] = types.YChild{"Pseudowire", nil}
    for i := range pseudowires.Pseudowire {
        pseudowires.EntityData.Children[types.GetSegmentPath(&pseudowires.Pseudowire[i])] = types.YChild{"Pseudowire", &pseudowires.Pseudowire[i]}
    }
    pseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pseudowires.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire
// Pseudowire configuration
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // keys: neighbor. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor.
    Neighbor []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor

    // keys: pseudowire-address. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress.
    PseudowireAddress []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress
}

func (pseudowire *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire) GetEntityData() *types.CommonEntityData {
    pseudowire.EntityData.YFilter = pseudowire.YFilter
    pseudowire.EntityData.YangName = "pseudowire"
    pseudowire.EntityData.BundleName = "cisco_ios_xr"
    pseudowire.EntityData.ParentYangName = "pseudowires"
    pseudowire.EntityData.SegmentPath = "pseudowire" + "[pseudowire-id='" + fmt.Sprintf("%v", pseudowire.PseudowireId) + "']"
    pseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowire.EntityData.Children = make(map[string]types.YChild)
    pseudowire.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range pseudowire.Neighbor {
        pseudowire.EntityData.Children[types.GetSegmentPath(&pseudowire.Neighbor[i])] = types.YChild{"Neighbor", &pseudowire.Neighbor[i]}
    }
    pseudowire.EntityData.Children["pseudowire-address"] = types.YChild{"PseudowireAddress", nil}
    for i := range pseudowire.PseudowireAddress {
        pseudowire.EntityData.Children[types.GetSegmentPath(&pseudowire.PseudowireAddress[i])] = types.YChild{"PseudowireAddress", &pseudowire.PseudowireAddress[i]}
    }
    pseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", pseudowire.PseudowireId}
    return &(pseudowire.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor
// keys: neighbor
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Pseudowire IPv4 address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}

    // Tag Impose vlan tagged mode. The type is interface{} with range: 1..4094.
    TagImpose interface{}

    // Name of the pseudowire class. The type is string with length: 1..32.
    Class interface{}

    // Value of the Pseudowire source address. Must be IPv6 only. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Pseudowire Bandwidth. The type is interface{} with range: 0..4294967295.
    Bandwidth interface{}

    // MPLS static labels.
    MplsStaticLabels L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels

    // List of pseudowires.
    BackupPseudowires L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires

    // L2TP Static Attributes.
    L2TpStaticAttributes L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes

    // Pseudowire L2TPv3 static configuration.
    L2TpStatic L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStatic
}

func (neighbor *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "pseudowire"
    neighbor.EntityData.SegmentPath = "neighbor" + "[neighbor='" + fmt.Sprintf("%v", neighbor.Neighbor) + "']"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Children["mpls-static-labels"] = types.YChild{"MplsStaticLabels", &neighbor.MplsStaticLabels}
    neighbor.EntityData.Children["backup-pseudowires"] = types.YChild{"BackupPseudowires", &neighbor.BackupPseudowires}
    neighbor.EntityData.Children["l2tp-static-attributes"] = types.YChild{"L2TpStaticAttributes", &neighbor.L2TpStaticAttributes}
    neighbor.EntityData.Children["l2tp-static"] = types.YChild{"L2TpStatic", &neighbor.L2TpStatic}
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", neighbor.Neighbor}
    neighbor.EntityData.Leafs["tag-impose"] = types.YLeaf{"TagImpose", neighbor.TagImpose}
    neighbor.EntityData.Leafs["class"] = types.YLeaf{"Class", neighbor.Class}
    neighbor.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", neighbor.SourceAddress}
    neighbor.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", neighbor.Bandwidth}
    return &(neighbor.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels
// MPLS static labels
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (mplsStaticLabels *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_MplsStaticLabels) GetEntityData() *types.CommonEntityData {
    mplsStaticLabels.EntityData.YFilter = mplsStaticLabels.YFilter
    mplsStaticLabels.EntityData.YangName = "mpls-static-labels"
    mplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    mplsStaticLabels.EntityData.ParentYangName = "neighbor"
    mplsStaticLabels.EntityData.SegmentPath = "mpls-static-labels"
    mplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsStaticLabels.EntityData.Children = make(map[string]types.YChild)
    mplsStaticLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsStaticLabels.EntityData.Leafs["local-static-label"] = types.YLeaf{"LocalStaticLabel", mplsStaticLabels.LocalStaticLabel}
    mplsStaticLabels.EntityData.Leafs["remote-static-label"] = types.YLeaf{"RemoteStaticLabel", mplsStaticLabels.RemoteStaticLabel}
    return &(mplsStaticLabels.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires
// List of pseudowires
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire for the cross connect. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire.
    BackupPseudowire []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire
}

func (backupPseudowires *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires) GetEntityData() *types.CommonEntityData {
    backupPseudowires.EntityData.YFilter = backupPseudowires.YFilter
    backupPseudowires.EntityData.YangName = "backup-pseudowires"
    backupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowires.EntityData.ParentYangName = "neighbor"
    backupPseudowires.EntityData.SegmentPath = "backup-pseudowires"
    backupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowires.EntityData.Children = make(map[string]types.YChild)
    backupPseudowires.EntityData.Children["backup-pseudowire"] = types.YChild{"BackupPseudowire", nil}
    for i := range backupPseudowires.BackupPseudowire {
        backupPseudowires.EntityData.Children[types.GetSegmentPath(&backupPseudowires.BackupPseudowire[i])] = types.YChild{"BackupPseudowire", &backupPseudowires.BackupPseudowire[i]}
    }
    backupPseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backupPseudowires.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire
// Backup pseudowire for the cross connect
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for the backup PW. The type is string with
    // length: 1..32.
    BackupPwClass interface{}

    // MPLS static labels.
    BackupMplsStaticLabels L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
}

func (backupPseudowire *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire) GetEntityData() *types.CommonEntityData {
    backupPseudowire.EntityData.YFilter = backupPseudowire.YFilter
    backupPseudowire.EntityData.YangName = "backup-pseudowire"
    backupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowire.EntityData.ParentYangName = "backup-pseudowires"
    backupPseudowire.EntityData.SegmentPath = "backup-pseudowire" + "[neighbor='" + fmt.Sprintf("%v", backupPseudowire.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", backupPseudowire.PseudowireId) + "']"
    backupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowire.EntityData.Children = make(map[string]types.YChild)
    backupPseudowire.EntityData.Children["backup-mpls-static-labels"] = types.YChild{"BackupMplsStaticLabels", &backupPseudowire.BackupMplsStaticLabels}
    backupPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    backupPseudowire.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", backupPseudowire.Neighbor}
    backupPseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", backupPseudowire.PseudowireId}
    backupPseudowire.EntityData.Leafs["backup-pw-class"] = types.YLeaf{"BackupPwClass", backupPseudowire.BackupPwClass}
    return &(backupPseudowire.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
// MPLS static labels
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (backupMplsStaticLabels *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    backupMplsStaticLabels.EntityData.YFilter = backupMplsStaticLabels.YFilter
    backupMplsStaticLabels.EntityData.YangName = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    backupMplsStaticLabels.EntityData.ParentYangName = "backup-pseudowire"
    backupMplsStaticLabels.EntityData.SegmentPath = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupMplsStaticLabels.EntityData.Children = make(map[string]types.YChild)
    backupMplsStaticLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    backupMplsStaticLabels.EntityData.Leafs["local-static-label"] = types.YLeaf{"LocalStaticLabel", backupMplsStaticLabels.LocalStaticLabel}
    backupMplsStaticLabels.EntityData.Leafs["remote-static-label"] = types.YLeaf{"RemoteStaticLabel", backupMplsStaticLabels.RemoteStaticLabel}
    return &(backupMplsStaticLabels.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes
// L2TP Static Attributes
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP remote session ID. The type is interface{} with range: 1..65535.
    L2TpRemoteSessionId interface{}

    // L2TP local session ID. The type is interface{} with range: 1..65535.
    L2TpLocalSessionId interface{}

    // L2TP remote cookie.
    L2TpRemoteCookie L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpRemoteCookie

    // L2TP secondary local cookie.
    L2TpSecondaryLocalCookie L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpSecondaryLocalCookie

    // L2TP local cookie.
    L2TpLocalCookie L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpLocalCookie
}

func (l2TpStaticAttributes *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes) GetEntityData() *types.CommonEntityData {
    l2TpStaticAttributes.EntityData.YFilter = l2TpStaticAttributes.YFilter
    l2TpStaticAttributes.EntityData.YangName = "l2tp-static-attributes"
    l2TpStaticAttributes.EntityData.BundleName = "cisco_ios_xr"
    l2TpStaticAttributes.EntityData.ParentYangName = "neighbor"
    l2TpStaticAttributes.EntityData.SegmentPath = "l2tp-static-attributes"
    l2TpStaticAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpStaticAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpStaticAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpStaticAttributes.EntityData.Children = make(map[string]types.YChild)
    l2TpStaticAttributes.EntityData.Children["l2tp-remote-cookie"] = types.YChild{"L2TpRemoteCookie", &l2TpStaticAttributes.L2TpRemoteCookie}
    l2TpStaticAttributes.EntityData.Children["l2tp-secondary-local-cookie"] = types.YChild{"L2TpSecondaryLocalCookie", &l2TpStaticAttributes.L2TpSecondaryLocalCookie}
    l2TpStaticAttributes.EntityData.Children["l2tp-local-cookie"] = types.YChild{"L2TpLocalCookie", &l2TpStaticAttributes.L2TpLocalCookie}
    l2TpStaticAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpStaticAttributes.EntityData.Leafs["l2tp-remote-session-id"] = types.YLeaf{"L2TpRemoteSessionId", l2TpStaticAttributes.L2TpRemoteSessionId}
    l2TpStaticAttributes.EntityData.Leafs["l2tp-local-session-id"] = types.YLeaf{"L2TpLocalSessionId", l2TpStaticAttributes.L2TpLocalSessionId}
    return &(l2TpStaticAttributes.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpRemoteCookie
// L2TP remote cookie
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpRemoteCookie struct {
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

func (l2TpRemoteCookie *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpRemoteCookie) GetEntityData() *types.CommonEntityData {
    l2TpRemoteCookie.EntityData.YFilter = l2TpRemoteCookie.YFilter
    l2TpRemoteCookie.EntityData.YangName = "l2tp-remote-cookie"
    l2TpRemoteCookie.EntityData.BundleName = "cisco_ios_xr"
    l2TpRemoteCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2TpRemoteCookie.EntityData.SegmentPath = "l2tp-remote-cookie"
    l2TpRemoteCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpRemoteCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpRemoteCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpRemoteCookie.EntityData.Children = make(map[string]types.YChild)
    l2TpRemoteCookie.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpRemoteCookie.EntityData.Leafs["size"] = types.YLeaf{"Size", l2TpRemoteCookie.Size}
    l2TpRemoteCookie.EntityData.Leafs["lower-value"] = types.YLeaf{"LowerValue", l2TpRemoteCookie.LowerValue}
    l2TpRemoteCookie.EntityData.Leafs["higher-value"] = types.YLeaf{"HigherValue", l2TpRemoteCookie.HigherValue}
    return &(l2TpRemoteCookie.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpSecondaryLocalCookie
// L2TP secondary local cookie
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpSecondaryLocalCookie struct {
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

func (l2TpSecondaryLocalCookie *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpSecondaryLocalCookie) GetEntityData() *types.CommonEntityData {
    l2TpSecondaryLocalCookie.EntityData.YFilter = l2TpSecondaryLocalCookie.YFilter
    l2TpSecondaryLocalCookie.EntityData.YangName = "l2tp-secondary-local-cookie"
    l2TpSecondaryLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2TpSecondaryLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2TpSecondaryLocalCookie.EntityData.SegmentPath = "l2tp-secondary-local-cookie"
    l2TpSecondaryLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpSecondaryLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpSecondaryLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpSecondaryLocalCookie.EntityData.Children = make(map[string]types.YChild)
    l2TpSecondaryLocalCookie.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpSecondaryLocalCookie.EntityData.Leafs["size"] = types.YLeaf{"Size", l2TpSecondaryLocalCookie.Size}
    l2TpSecondaryLocalCookie.EntityData.Leafs["lower-value"] = types.YLeaf{"LowerValue", l2TpSecondaryLocalCookie.LowerValue}
    l2TpSecondaryLocalCookie.EntityData.Leafs["higher-value"] = types.YLeaf{"HigherValue", l2TpSecondaryLocalCookie.HigherValue}
    return &(l2TpSecondaryLocalCookie.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpLocalCookie
// L2TP local cookie
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpLocalCookie struct {
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

func (l2TpLocalCookie *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStaticAttributes_L2TpLocalCookie) GetEntityData() *types.CommonEntityData {
    l2TpLocalCookie.EntityData.YFilter = l2TpLocalCookie.YFilter
    l2TpLocalCookie.EntityData.YangName = "l2tp-local-cookie"
    l2TpLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2TpLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2TpLocalCookie.EntityData.SegmentPath = "l2tp-local-cookie"
    l2TpLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpLocalCookie.EntityData.Children = make(map[string]types.YChild)
    l2TpLocalCookie.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpLocalCookie.EntityData.Leafs["size"] = types.YLeaf{"Size", l2TpLocalCookie.Size}
    l2TpLocalCookie.EntityData.Leafs["lower-value"] = types.YLeaf{"LowerValue", l2TpLocalCookie.LowerValue}
    l2TpLocalCookie.EntityData.Leafs["higher-value"] = types.YLeaf{"HigherValue", l2TpLocalCookie.HigherValue}
    return &(l2TpLocalCookie.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStatic
// Pseudowire L2TPv3 static configuration
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pseudowire L2TPv3 static configuration. The type is interface{}.
    Enable interface{}
}

func (l2TpStatic *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_Neighbor_L2TpStatic) GetEntityData() *types.CommonEntityData {
    l2TpStatic.EntityData.YFilter = l2TpStatic.YFilter
    l2TpStatic.EntityData.YangName = "l2tp-static"
    l2TpStatic.EntityData.BundleName = "cisco_ios_xr"
    l2TpStatic.EntityData.ParentYangName = "neighbor"
    l2TpStatic.EntityData.SegmentPath = "l2tp-static"
    l2TpStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpStatic.EntityData.Children = make(map[string]types.YChild)
    l2TpStatic.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpStatic.EntityData.Leafs["enable"] = types.YLeaf{"Enable", l2TpStatic.Enable}
    return &(l2TpStatic.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress
// keys: pseudowire-address
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Pseudowire IPv6 address. A pseudowire can have
    // only one address: IPv4 or IPv6. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PseudowireAddress interface{}

    // Tag Impose vlan tagged mode. The type is interface{} with range: 1..4094.
    TagImpose interface{}

    // Name of the pseudowire class. The type is string with length: 1..32.
    Class interface{}

    // Value of the Pseudowire source address. Must be IPv6 only. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Pseudowire Bandwidth. The type is interface{} with range: 0..4294967295.
    Bandwidth interface{}

    // MPLS static labels.
    MplsStaticLabels L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels

    // List of pseudowires.
    BackupPseudowires L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires

    // L2TP Static Attributes.
    L2TpStaticAttributes L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes

    // Pseudowire L2TPv3 static configuration.
    L2TpStatic L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStatic
}

func (pseudowireAddress *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress) GetEntityData() *types.CommonEntityData {
    pseudowireAddress.EntityData.YFilter = pseudowireAddress.YFilter
    pseudowireAddress.EntityData.YangName = "pseudowire-address"
    pseudowireAddress.EntityData.BundleName = "cisco_ios_xr"
    pseudowireAddress.EntityData.ParentYangName = "pseudowire"
    pseudowireAddress.EntityData.SegmentPath = "pseudowire-address" + "[pseudowire-address='" + fmt.Sprintf("%v", pseudowireAddress.PseudowireAddress) + "']"
    pseudowireAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireAddress.EntityData.Children = make(map[string]types.YChild)
    pseudowireAddress.EntityData.Children["mpls-static-labels"] = types.YChild{"MplsStaticLabels", &pseudowireAddress.MplsStaticLabels}
    pseudowireAddress.EntityData.Children["backup-pseudowires"] = types.YChild{"BackupPseudowires", &pseudowireAddress.BackupPseudowires}
    pseudowireAddress.EntityData.Children["l2tp-static-attributes"] = types.YChild{"L2TpStaticAttributes", &pseudowireAddress.L2TpStaticAttributes}
    pseudowireAddress.EntityData.Children["l2tp-static"] = types.YChild{"L2TpStatic", &pseudowireAddress.L2TpStatic}
    pseudowireAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireAddress.EntityData.Leafs["pseudowire-address"] = types.YLeaf{"PseudowireAddress", pseudowireAddress.PseudowireAddress}
    pseudowireAddress.EntityData.Leafs["tag-impose"] = types.YLeaf{"TagImpose", pseudowireAddress.TagImpose}
    pseudowireAddress.EntityData.Leafs["class"] = types.YLeaf{"Class", pseudowireAddress.Class}
    pseudowireAddress.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", pseudowireAddress.SourceAddress}
    pseudowireAddress.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", pseudowireAddress.Bandwidth}
    return &(pseudowireAddress.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels
// MPLS static labels
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (mplsStaticLabels *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_MplsStaticLabels) GetEntityData() *types.CommonEntityData {
    mplsStaticLabels.EntityData.YFilter = mplsStaticLabels.YFilter
    mplsStaticLabels.EntityData.YangName = "mpls-static-labels"
    mplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    mplsStaticLabels.EntityData.ParentYangName = "pseudowire-address"
    mplsStaticLabels.EntityData.SegmentPath = "mpls-static-labels"
    mplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsStaticLabels.EntityData.Children = make(map[string]types.YChild)
    mplsStaticLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsStaticLabels.EntityData.Leafs["local-static-label"] = types.YLeaf{"LocalStaticLabel", mplsStaticLabels.LocalStaticLabel}
    mplsStaticLabels.EntityData.Leafs["remote-static-label"] = types.YLeaf{"RemoteStaticLabel", mplsStaticLabels.RemoteStaticLabel}
    return &(mplsStaticLabels.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires
// List of pseudowires
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire for the cross connect. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire.
    BackupPseudowire []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire
}

func (backupPseudowires *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires) GetEntityData() *types.CommonEntityData {
    backupPseudowires.EntityData.YFilter = backupPseudowires.YFilter
    backupPseudowires.EntityData.YangName = "backup-pseudowires"
    backupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowires.EntityData.ParentYangName = "pseudowire-address"
    backupPseudowires.EntityData.SegmentPath = "backup-pseudowires"
    backupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowires.EntityData.Children = make(map[string]types.YChild)
    backupPseudowires.EntityData.Children["backup-pseudowire"] = types.YChild{"BackupPseudowire", nil}
    for i := range backupPseudowires.BackupPseudowire {
        backupPseudowires.EntityData.Children[types.GetSegmentPath(&backupPseudowires.BackupPseudowire[i])] = types.YChild{"BackupPseudowire", &backupPseudowires.BackupPseudowire[i]}
    }
    backupPseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backupPseudowires.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire
// Backup pseudowire for the cross connect
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for the backup PW. The type is string with
    // length: 1..32.
    BackupPwClass interface{}

    // MPLS static labels.
    BackupMplsStaticLabels L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
}

func (backupPseudowire *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire) GetEntityData() *types.CommonEntityData {
    backupPseudowire.EntityData.YFilter = backupPseudowire.YFilter
    backupPseudowire.EntityData.YangName = "backup-pseudowire"
    backupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    backupPseudowire.EntityData.ParentYangName = "backup-pseudowires"
    backupPseudowire.EntityData.SegmentPath = "backup-pseudowire" + "[neighbor='" + fmt.Sprintf("%v", backupPseudowire.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", backupPseudowire.PseudowireId) + "']"
    backupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPseudowire.EntityData.Children = make(map[string]types.YChild)
    backupPseudowire.EntityData.Children["backup-mpls-static-labels"] = types.YChild{"BackupMplsStaticLabels", &backupPseudowire.BackupMplsStaticLabels}
    backupPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    backupPseudowire.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", backupPseudowire.Neighbor}
    backupPseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", backupPseudowire.PseudowireId}
    backupPseudowire.EntityData.Leafs["backup-pw-class"] = types.YLeaf{"BackupPwClass", backupPseudowire.BackupPwClass}
    return &(backupPseudowire.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels
// MPLS static labels
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (backupMplsStaticLabels *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_BackupPseudowires_BackupPseudowire_BackupMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    backupMplsStaticLabels.EntityData.YFilter = backupMplsStaticLabels.YFilter
    backupMplsStaticLabels.EntityData.YangName = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    backupMplsStaticLabels.EntityData.ParentYangName = "backup-pseudowire"
    backupMplsStaticLabels.EntityData.SegmentPath = "backup-mpls-static-labels"
    backupMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupMplsStaticLabels.EntityData.Children = make(map[string]types.YChild)
    backupMplsStaticLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    backupMplsStaticLabels.EntityData.Leafs["local-static-label"] = types.YLeaf{"LocalStaticLabel", backupMplsStaticLabels.LocalStaticLabel}
    backupMplsStaticLabels.EntityData.Leafs["remote-static-label"] = types.YLeaf{"RemoteStaticLabel", backupMplsStaticLabels.RemoteStaticLabel}
    return &(backupMplsStaticLabels.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes
// L2TP Static Attributes
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TP remote session ID. The type is interface{} with range: 1..65535.
    L2TpRemoteSessionId interface{}

    // L2TP local session ID. The type is interface{} with range: 1..65535.
    L2TpLocalSessionId interface{}

    // L2TP remote cookie.
    L2TpRemoteCookie L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpRemoteCookie

    // L2TP secondary local cookie.
    L2TpSecondaryLocalCookie L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpSecondaryLocalCookie

    // L2TP local cookie.
    L2TpLocalCookie L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpLocalCookie
}

func (l2TpStaticAttributes *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes) GetEntityData() *types.CommonEntityData {
    l2TpStaticAttributes.EntityData.YFilter = l2TpStaticAttributes.YFilter
    l2TpStaticAttributes.EntityData.YangName = "l2tp-static-attributes"
    l2TpStaticAttributes.EntityData.BundleName = "cisco_ios_xr"
    l2TpStaticAttributes.EntityData.ParentYangName = "pseudowire-address"
    l2TpStaticAttributes.EntityData.SegmentPath = "l2tp-static-attributes"
    l2TpStaticAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpStaticAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpStaticAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpStaticAttributes.EntityData.Children = make(map[string]types.YChild)
    l2TpStaticAttributes.EntityData.Children["l2tp-remote-cookie"] = types.YChild{"L2TpRemoteCookie", &l2TpStaticAttributes.L2TpRemoteCookie}
    l2TpStaticAttributes.EntityData.Children["l2tp-secondary-local-cookie"] = types.YChild{"L2TpSecondaryLocalCookie", &l2TpStaticAttributes.L2TpSecondaryLocalCookie}
    l2TpStaticAttributes.EntityData.Children["l2tp-local-cookie"] = types.YChild{"L2TpLocalCookie", &l2TpStaticAttributes.L2TpLocalCookie}
    l2TpStaticAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpStaticAttributes.EntityData.Leafs["l2tp-remote-session-id"] = types.YLeaf{"L2TpRemoteSessionId", l2TpStaticAttributes.L2TpRemoteSessionId}
    l2TpStaticAttributes.EntityData.Leafs["l2tp-local-session-id"] = types.YLeaf{"L2TpLocalSessionId", l2TpStaticAttributes.L2TpLocalSessionId}
    return &(l2TpStaticAttributes.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpRemoteCookie
// L2TP remote cookie
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpRemoteCookie struct {
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

func (l2TpRemoteCookie *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpRemoteCookie) GetEntityData() *types.CommonEntityData {
    l2TpRemoteCookie.EntityData.YFilter = l2TpRemoteCookie.YFilter
    l2TpRemoteCookie.EntityData.YangName = "l2tp-remote-cookie"
    l2TpRemoteCookie.EntityData.BundleName = "cisco_ios_xr"
    l2TpRemoteCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2TpRemoteCookie.EntityData.SegmentPath = "l2tp-remote-cookie"
    l2TpRemoteCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpRemoteCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpRemoteCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpRemoteCookie.EntityData.Children = make(map[string]types.YChild)
    l2TpRemoteCookie.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpRemoteCookie.EntityData.Leafs["size"] = types.YLeaf{"Size", l2TpRemoteCookie.Size}
    l2TpRemoteCookie.EntityData.Leafs["lower-value"] = types.YLeaf{"LowerValue", l2TpRemoteCookie.LowerValue}
    l2TpRemoteCookie.EntityData.Leafs["higher-value"] = types.YLeaf{"HigherValue", l2TpRemoteCookie.HigherValue}
    return &(l2TpRemoteCookie.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpSecondaryLocalCookie
// L2TP secondary local cookie
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpSecondaryLocalCookie struct {
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

func (l2TpSecondaryLocalCookie *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpSecondaryLocalCookie) GetEntityData() *types.CommonEntityData {
    l2TpSecondaryLocalCookie.EntityData.YFilter = l2TpSecondaryLocalCookie.YFilter
    l2TpSecondaryLocalCookie.EntityData.YangName = "l2tp-secondary-local-cookie"
    l2TpSecondaryLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2TpSecondaryLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2TpSecondaryLocalCookie.EntityData.SegmentPath = "l2tp-secondary-local-cookie"
    l2TpSecondaryLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpSecondaryLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpSecondaryLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpSecondaryLocalCookie.EntityData.Children = make(map[string]types.YChild)
    l2TpSecondaryLocalCookie.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpSecondaryLocalCookie.EntityData.Leafs["size"] = types.YLeaf{"Size", l2TpSecondaryLocalCookie.Size}
    l2TpSecondaryLocalCookie.EntityData.Leafs["lower-value"] = types.YLeaf{"LowerValue", l2TpSecondaryLocalCookie.LowerValue}
    l2TpSecondaryLocalCookie.EntityData.Leafs["higher-value"] = types.YLeaf{"HigherValue", l2TpSecondaryLocalCookie.HigherValue}
    return &(l2TpSecondaryLocalCookie.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpLocalCookie
// L2TP local cookie
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpLocalCookie struct {
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

func (l2TpLocalCookie *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStaticAttributes_L2TpLocalCookie) GetEntityData() *types.CommonEntityData {
    l2TpLocalCookie.EntityData.YFilter = l2TpLocalCookie.YFilter
    l2TpLocalCookie.EntityData.YangName = "l2tp-local-cookie"
    l2TpLocalCookie.EntityData.BundleName = "cisco_ios_xr"
    l2TpLocalCookie.EntityData.ParentYangName = "l2tp-static-attributes"
    l2TpLocalCookie.EntityData.SegmentPath = "l2tp-local-cookie"
    l2TpLocalCookie.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpLocalCookie.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpLocalCookie.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpLocalCookie.EntityData.Children = make(map[string]types.YChild)
    l2TpLocalCookie.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpLocalCookie.EntityData.Leafs["size"] = types.YLeaf{"Size", l2TpLocalCookie.Size}
    l2TpLocalCookie.EntityData.Leafs["lower-value"] = types.YLeaf{"LowerValue", l2TpLocalCookie.LowerValue}
    l2TpLocalCookie.EntityData.Leafs["higher-value"] = types.YLeaf{"HigherValue", l2TpLocalCookie.HigherValue}
    return &(l2TpLocalCookie.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStatic
// Pseudowire L2TPv3 static configuration
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pseudowire L2TPv3 static configuration. The type is interface{}.
    Enable interface{}
}

func (l2TpStatic *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_Pseudowires_Pseudowire_PseudowireAddress_L2TpStatic) GetEntityData() *types.CommonEntityData {
    l2TpStatic.EntityData.YFilter = l2TpStatic.YFilter
    l2TpStatic.EntityData.YangName = "l2tp-static"
    l2TpStatic.EntityData.BundleName = "cisco_ios_xr"
    l2TpStatic.EntityData.ParentYangName = "pseudowire-address"
    l2TpStatic.EntityData.SegmentPath = "l2tp-static"
    l2TpStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2TpStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2TpStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2TpStatic.EntityData.Children = make(map[string]types.YChild)
    l2TpStatic.EntityData.Leafs = make(map[string]types.YLeaf)
    l2TpStatic.EntityData.Leafs["enable"] = types.YLeaf{"Enable", l2TpStatic.Enable}
    return &(l2TpStatic.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions
// List of Monitor session segments
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor session segment. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions_MonitorSession.
    MonitorSession []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions_MonitorSession
}

func (monitorSessions *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions) GetEntityData() *types.CommonEntityData {
    monitorSessions.EntityData.YFilter = monitorSessions.YFilter
    monitorSessions.EntityData.YangName = "monitor-sessions"
    monitorSessions.EntityData.BundleName = "cisco_ios_xr"
    monitorSessions.EntityData.ParentYangName = "p2p-xconnect"
    monitorSessions.EntityData.SegmentPath = "monitor-sessions"
    monitorSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorSessions.EntityData.Children = make(map[string]types.YChild)
    monitorSessions.EntityData.Children["monitor-session"] = types.YChild{"MonitorSession", nil}
    for i := range monitorSessions.MonitorSession {
        monitorSessions.EntityData.Children[types.GetSegmentPath(&monitorSessions.MonitorSession[i])] = types.YChild{"MonitorSession", &monitorSessions.MonitorSession[i]}
    }
    monitorSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(monitorSessions.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions_MonitorSession
// Monitor session segment
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions_MonitorSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the monitor session. The type is string
    // with length: 1..64.
    Name interface{}

    // Enable monitor session segment . The type is interface{}.
    Enable interface{}
}

func (monitorSession *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_MonitorSessions_MonitorSession) GetEntityData() *types.CommonEntityData {
    monitorSession.EntityData.YFilter = monitorSession.YFilter
    monitorSession.EntityData.YangName = "monitor-session"
    monitorSession.EntityData.BundleName = "cisco_ios_xr"
    monitorSession.EntityData.ParentYangName = "monitor-sessions"
    monitorSession.EntityData.SegmentPath = "monitor-session" + "[name='" + fmt.Sprintf("%v", monitorSession.Name) + "']"
    monitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorSession.EntityData.Children = make(map[string]types.YChild)
    monitorSession.EntityData.Leafs = make(map[string]types.YLeaf)
    monitorSession.EntityData.Leafs["name"] = types.YLeaf{"Name", monitorSession.Name}
    monitorSession.EntityData.Leafs["enable"] = types.YLeaf{"Enable", monitorSession.Enable}
    return &(monitorSession.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds
// List of pseudowire-routed
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds_PseudowireRouted.
    PseudowireRouted []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds_PseudowireRouted
}

func (pseudowireRouteds *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds) GetEntityData() *types.CommonEntityData {
    pseudowireRouteds.EntityData.YFilter = pseudowireRouteds.YFilter
    pseudowireRouteds.EntityData.YangName = "pseudowire-routeds"
    pseudowireRouteds.EntityData.BundleName = "cisco_ios_xr"
    pseudowireRouteds.EntityData.ParentYangName = "p2p-xconnect"
    pseudowireRouteds.EntityData.SegmentPath = "pseudowire-routeds"
    pseudowireRouteds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireRouteds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireRouteds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireRouteds.EntityData.Children = make(map[string]types.YChild)
    pseudowireRouteds.EntityData.Children["pseudowire-routed"] = types.YChild{"PseudowireRouted", nil}
    for i := range pseudowireRouteds.PseudowireRouted {
        pseudowireRouteds.EntityData.Children[types.GetSegmentPath(&pseudowireRouteds.PseudowireRouted[i])] = types.YChild{"PseudowireRouted", &pseudowireRouteds.PseudowireRouted[i]}
    }
    pseudowireRouteds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pseudowireRouteds.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds_PseudowireRouted
// Pseudowire configuration
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds_PseudowireRouted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Target Global ID. The type is interface{} with
    // range: 1..4294967295.
    GlobalId interface{}

    // This attribute is a key. Target Prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (pseudowireRouted *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_PseudowireRouteds_PseudowireRouted) GetEntityData() *types.CommonEntityData {
    pseudowireRouted.EntityData.YFilter = pseudowireRouted.YFilter
    pseudowireRouted.EntityData.YangName = "pseudowire-routed"
    pseudowireRouted.EntityData.BundleName = "cisco_ios_xr"
    pseudowireRouted.EntityData.ParentYangName = "pseudowire-routeds"
    pseudowireRouted.EntityData.SegmentPath = "pseudowire-routed" + "[global-id='" + fmt.Sprintf("%v", pseudowireRouted.GlobalId) + "']" + "[prefix='" + fmt.Sprintf("%v", pseudowireRouted.Prefix) + "']" + "[acid='" + fmt.Sprintf("%v", pseudowireRouted.Acid) + "']" + "[sacid='" + fmt.Sprintf("%v", pseudowireRouted.Sacid) + "']"
    pseudowireRouted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireRouted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireRouted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireRouted.EntityData.Children = make(map[string]types.YChild)
    pseudowireRouted.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireRouted.EntityData.Leafs["global-id"] = types.YLeaf{"GlobalId", pseudowireRouted.GlobalId}
    pseudowireRouted.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", pseudowireRouted.Prefix}
    pseudowireRouted.EntityData.Leafs["acid"] = types.YLeaf{"Acid", pseudowireRouted.Acid}
    pseudowireRouted.EntityData.Leafs["sacid"] = types.YLeaf{"Sacid", pseudowireRouted.Sacid}
    pseudowireRouted.EntityData.Leafs["tag-impose"] = types.YLeaf{"TagImpose", pseudowireRouted.TagImpose}
    pseudowireRouted.EntityData.Leafs["class"] = types.YLeaf{"Class", pseudowireRouted.Class}
    return &(pseudowireRouted.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits
// List of attachment circuits
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits_AttachmentCircuit.
    AttachmentCircuit []L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits_AttachmentCircuit
}

func (attachmentCircuits *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits) GetEntityData() *types.CommonEntityData {
    attachmentCircuits.EntityData.YFilter = attachmentCircuits.YFilter
    attachmentCircuits.EntityData.YangName = "attachment-circuits"
    attachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    attachmentCircuits.EntityData.ParentYangName = "p2p-xconnect"
    attachmentCircuits.EntityData.SegmentPath = "attachment-circuits"
    attachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachmentCircuits.EntityData.Children = make(map[string]types.YChild)
    attachmentCircuits.EntityData.Children["attachment-circuit"] = types.YChild{"AttachmentCircuit", nil}
    for i := range attachmentCircuits.AttachmentCircuit {
        attachmentCircuits.EntityData.Children[types.GetSegmentPath(&attachmentCircuits.AttachmentCircuit[i])] = types.YChild{"AttachmentCircuit", &attachmentCircuits.AttachmentCircuit[i]}
    }
    attachmentCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attachmentCircuits.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits_AttachmentCircuit
// Attachment circuit interface
type L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits_AttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Enable attachment circuit interface. The type is interface{}.
    Enable interface{}
}

func (attachmentCircuit *L2Vpn_Database_XconnectGroups_XconnectGroup_P2PXconnects_P2PXconnect_AttachmentCircuits_AttachmentCircuit) GetEntityData() *types.CommonEntityData {
    attachmentCircuit.EntityData.YFilter = attachmentCircuit.YFilter
    attachmentCircuit.EntityData.YangName = "attachment-circuit"
    attachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    attachmentCircuit.EntityData.ParentYangName = "attachment-circuits"
    attachmentCircuit.EntityData.SegmentPath = "attachment-circuit" + "[name='" + fmt.Sprintf("%v", attachmentCircuit.Name) + "']"
    attachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachmentCircuit.EntityData.Children = make(map[string]types.YChild)
    attachmentCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    attachmentCircuit.EntityData.Leafs["name"] = types.YLeaf{"Name", attachmentCircuit.Name}
    attachmentCircuit.EntityData.Leafs["enable"] = types.YLeaf{"Enable", attachmentCircuit.Enable}
    return &(attachmentCircuit.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects
// List of multi point to multi point xconnects
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multi point to multi point xconnect. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect.
    Mp2MpXconnect []L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect
}

func (mp2MpXconnects *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects) GetEntityData() *types.CommonEntityData {
    mp2MpXconnects.EntityData.YFilter = mp2MpXconnects.YFilter
    mp2MpXconnects.EntityData.YangName = "mp2mp-xconnects"
    mp2MpXconnects.EntityData.BundleName = "cisco_ios_xr"
    mp2MpXconnects.EntityData.ParentYangName = "xconnect-group"
    mp2MpXconnects.EntityData.SegmentPath = "mp2mp-xconnects"
    mp2MpXconnects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpXconnects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpXconnects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpXconnects.EntityData.Children = make(map[string]types.YChild)
    mp2MpXconnects.EntityData.Children["mp2mp-xconnect"] = types.YChild{"Mp2MpXconnect", nil}
    for i := range mp2MpXconnects.Mp2MpXconnect {
        mp2MpXconnects.EntityData.Children[types.GetSegmentPath(&mp2MpXconnects.Mp2MpXconnect[i])] = types.YChild{"Mp2MpXconnect", &mp2MpXconnects.Mp2MpXconnect[i]}
    }
    mp2MpXconnects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mp2MpXconnects.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect
// Multi point to multi point xconnect
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the multi point to multi point xconnect.
    // The type is string with length: 1..26.
    Name interface{}

    // Maximum transmission unit for this MP2MP VPWS instance. The type is
    // interface{} with range: 64..65535. Units are byte.
    Mp2Mpmtu interface{}

    // Disable control word. The type is interface{}.
    Mp2MpControlWord interface{}

    // Configure Layer 2 Encapsulation. The type is L2Encapsulation.
    Mp2Mpl2Encapsulation interface{}

    // Interworking. The type is Interworking.
    Mp2MpInterworking interface{}

    // shutdown this MP2MP VPWS instance. The type is interface{}.
    Mp2MpShutdown interface{}

    // VPN Identifier. The type is interface{} with range: 1..4294967295.
    Mp2MpvpnId interface{}

    // auto-discovery in this MP2MP.
    Mp2MpAutoDiscovery L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery
}

func (mp2MpXconnect *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect) GetEntityData() *types.CommonEntityData {
    mp2MpXconnect.EntityData.YFilter = mp2MpXconnect.YFilter
    mp2MpXconnect.EntityData.YangName = "mp2mp-xconnect"
    mp2MpXconnect.EntityData.BundleName = "cisco_ios_xr"
    mp2MpXconnect.EntityData.ParentYangName = "mp2mp-xconnects"
    mp2MpXconnect.EntityData.SegmentPath = "mp2mp-xconnect" + "[name='" + fmt.Sprintf("%v", mp2MpXconnect.Name) + "']"
    mp2MpXconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpXconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpXconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpXconnect.EntityData.Children = make(map[string]types.YChild)
    mp2MpXconnect.EntityData.Children["mp2mp-auto-discovery"] = types.YChild{"Mp2MpAutoDiscovery", &mp2MpXconnect.Mp2MpAutoDiscovery}
    mp2MpXconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    mp2MpXconnect.EntityData.Leafs["name"] = types.YLeaf{"Name", mp2MpXconnect.Name}
    mp2MpXconnect.EntityData.Leafs["mp2mpmtu"] = types.YLeaf{"Mp2Mpmtu", mp2MpXconnect.Mp2Mpmtu}
    mp2MpXconnect.EntityData.Leafs["mp2mp-control-word"] = types.YLeaf{"Mp2MpControlWord", mp2MpXconnect.Mp2MpControlWord}
    mp2MpXconnect.EntityData.Leafs["mp2mpl2-encapsulation"] = types.YLeaf{"Mp2Mpl2Encapsulation", mp2MpXconnect.Mp2Mpl2Encapsulation}
    mp2MpXconnect.EntityData.Leafs["mp2mp-interworking"] = types.YLeaf{"Mp2MpInterworking", mp2MpXconnect.Mp2MpInterworking}
    mp2MpXconnect.EntityData.Leafs["mp2mp-shutdown"] = types.YLeaf{"Mp2MpShutdown", mp2MpXconnect.Mp2MpShutdown}
    mp2MpXconnect.EntityData.Leafs["mp2mpvpn-id"] = types.YLeaf{"Mp2MpvpnId", mp2MpXconnect.Mp2MpvpnId}
    return &(mp2MpXconnect.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery
// auto-discovery in this MP2MP
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable auto-discovery. The type is interface{}.
    Enable interface{}

    // Route Distinguisher.
    RouteDistinguisher L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_RouteDistinguisher

    // Route policy.
    Mp2MpRoutePolicy L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRoutePolicy

    // Route Target.
    Mp2MpRouteTargets L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets

    // signaling protocol in this MP2MP.
    Mp2MpSignalingProtocol L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol
}

func (mp2MpAutoDiscovery *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    mp2MpAutoDiscovery.EntityData.YFilter = mp2MpAutoDiscovery.YFilter
    mp2MpAutoDiscovery.EntityData.YangName = "mp2mp-auto-discovery"
    mp2MpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    mp2MpAutoDiscovery.EntityData.ParentYangName = "mp2mp-xconnect"
    mp2MpAutoDiscovery.EntityData.SegmentPath = "mp2mp-auto-discovery"
    mp2MpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    mp2MpAutoDiscovery.EntityData.Children["route-distinguisher"] = types.YChild{"RouteDistinguisher", &mp2MpAutoDiscovery.RouteDistinguisher}
    mp2MpAutoDiscovery.EntityData.Children["mp2mp-route-policy"] = types.YChild{"Mp2MpRoutePolicy", &mp2MpAutoDiscovery.Mp2MpRoutePolicy}
    mp2MpAutoDiscovery.EntityData.Children["mp2mp-route-targets"] = types.YChild{"Mp2MpRouteTargets", &mp2MpAutoDiscovery.Mp2MpRouteTargets}
    mp2MpAutoDiscovery.EntityData.Children["mp2mp-signaling-protocol"] = types.YChild{"Mp2MpSignalingProtocol", &mp2MpAutoDiscovery.Mp2MpSignalingProtocol}
    mp2MpAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    mp2MpAutoDiscovery.EntityData.Leafs["enable"] = types.YLeaf{"Enable", mp2MpAutoDiscovery.Enable}
    return &(mp2MpAutoDiscovery.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_RouteDistinguisher
// Route Distinguisher
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_RouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router distinguisher type. The type is BgpRouteDistinguisher.
    Type_ interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (routeDistinguisher *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_RouteDistinguisher) GetEntityData() *types.CommonEntityData {
    routeDistinguisher.EntityData.YFilter = routeDistinguisher.YFilter
    routeDistinguisher.EntityData.YangName = "route-distinguisher"
    routeDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    routeDistinguisher.EntityData.ParentYangName = "mp2mp-auto-discovery"
    routeDistinguisher.EntityData.SegmentPath = "route-distinguisher"
    routeDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeDistinguisher.EntityData.Children = make(map[string]types.YChild)
    routeDistinguisher.EntityData.Leafs = make(map[string]types.YLeaf)
    routeDistinguisher.EntityData.Leafs["type"] = types.YLeaf{"Type_", routeDistinguisher.Type_}
    routeDistinguisher.EntityData.Leafs["as"] = types.YLeaf{"As", routeDistinguisher.As}
    routeDistinguisher.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", routeDistinguisher.AsIndex}
    routeDistinguisher.EntityData.Leafs["address"] = types.YLeaf{"Address", routeDistinguisher.Address}
    routeDistinguisher.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", routeDistinguisher.AddrIndex}
    return &(routeDistinguisher.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRoutePolicy
// Route policy
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Export route policy. The type is string.
    Export interface{}
}

func (mp2MpRoutePolicy *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRoutePolicy) GetEntityData() *types.CommonEntityData {
    mp2MpRoutePolicy.EntityData.YFilter = mp2MpRoutePolicy.YFilter
    mp2MpRoutePolicy.EntityData.YangName = "mp2mp-route-policy"
    mp2MpRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    mp2MpRoutePolicy.EntityData.ParentYangName = "mp2mp-auto-discovery"
    mp2MpRoutePolicy.EntityData.SegmentPath = "mp2mp-route-policy"
    mp2MpRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpRoutePolicy.EntityData.Children = make(map[string]types.YChild)
    mp2MpRoutePolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    mp2MpRoutePolicy.EntityData.Leafs["export"] = types.YLeaf{"Export", mp2MpRoutePolicy.Export}
    return &(mp2MpRoutePolicy.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets
// Route Target
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget.
    Mp2MpRouteTarget []L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget
}

func (mp2MpRouteTargets *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets) GetEntityData() *types.CommonEntityData {
    mp2MpRouteTargets.EntityData.YFilter = mp2MpRouteTargets.YFilter
    mp2MpRouteTargets.EntityData.YangName = "mp2mp-route-targets"
    mp2MpRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    mp2MpRouteTargets.EntityData.ParentYangName = "mp2mp-auto-discovery"
    mp2MpRouteTargets.EntityData.SegmentPath = "mp2mp-route-targets"
    mp2MpRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpRouteTargets.EntityData.Children = make(map[string]types.YChild)
    mp2MpRouteTargets.EntityData.Children["mp2mp-route-target"] = types.YChild{"Mp2MpRouteTarget", nil}
    for i := range mp2MpRouteTargets.Mp2MpRouteTarget {
        mp2MpRouteTargets.EntityData.Children[types.GetSegmentPath(&mp2MpRouteTargets.Mp2MpRouteTarget[i])] = types.YChild{"Mp2MpRouteTarget", &mp2MpRouteTargets.Mp2MpRouteTarget[i]}
    }
    mp2MpRouteTargets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mp2MpRouteTargets.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget
// Name of the Route Target
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // two byte as or four byte as. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_TwoByteAsOrFourByteAs.
    TwoByteAsOrFourByteAs []L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_TwoByteAsOrFourByteAs

    // ipv4 address. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_Ipv4Address.
    Ipv4Address []L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_Ipv4Address
}

func (mp2MpRouteTarget *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget) GetEntityData() *types.CommonEntityData {
    mp2MpRouteTarget.EntityData.YFilter = mp2MpRouteTarget.YFilter
    mp2MpRouteTarget.EntityData.YangName = "mp2mp-route-target"
    mp2MpRouteTarget.EntityData.BundleName = "cisco_ios_xr"
    mp2MpRouteTarget.EntityData.ParentYangName = "mp2mp-route-targets"
    mp2MpRouteTarget.EntityData.SegmentPath = "mp2mp-route-target" + "[role='" + fmt.Sprintf("%v", mp2MpRouteTarget.Role) + "']" + "[format='" + fmt.Sprintf("%v", mp2MpRouteTarget.Format) + "']"
    mp2MpRouteTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpRouteTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpRouteTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpRouteTarget.EntityData.Children = make(map[string]types.YChild)
    mp2MpRouteTarget.EntityData.Children["two-byte-as-or-four-byte-as"] = types.YChild{"TwoByteAsOrFourByteAs", nil}
    for i := range mp2MpRouteTarget.TwoByteAsOrFourByteAs {
        mp2MpRouteTarget.EntityData.Children[types.GetSegmentPath(&mp2MpRouteTarget.TwoByteAsOrFourByteAs[i])] = types.YChild{"TwoByteAsOrFourByteAs", &mp2MpRouteTarget.TwoByteAsOrFourByteAs[i]}
    }
    mp2MpRouteTarget.EntityData.Children["ipv4-address"] = types.YChild{"Ipv4Address", nil}
    for i := range mp2MpRouteTarget.Ipv4Address {
        mp2MpRouteTarget.EntityData.Children[types.GetSegmentPath(&mp2MpRouteTarget.Ipv4Address[i])] = types.YChild{"Ipv4Address", &mp2MpRouteTarget.Ipv4Address[i]}
    }
    mp2MpRouteTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    mp2MpRouteTarget.EntityData.Leafs["role"] = types.YLeaf{"Role", mp2MpRouteTarget.Role}
    mp2MpRouteTarget.EntityData.Leafs["format"] = types.YLeaf{"Format", mp2MpRouteTarget.Format}
    return &(mp2MpRouteTarget.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_TwoByteAsOrFourByteAs
// two byte as or four byte as
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_TwoByteAsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}
}

func (twoByteAsOrFourByteAs *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_TwoByteAsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAsOrFourByteAs.EntityData.YFilter = twoByteAsOrFourByteAs.YFilter
    twoByteAsOrFourByteAs.EntityData.YangName = "two-byte-as-or-four-byte-as"
    twoByteAsOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAsOrFourByteAs.EntityData.ParentYangName = "mp2mp-route-target"
    twoByteAsOrFourByteAs.EntityData.SegmentPath = "two-byte-as-or-four-byte-as" + "[as='" + fmt.Sprintf("%v", twoByteAsOrFourByteAs.As) + "']" + "[as-index='" + fmt.Sprintf("%v", twoByteAsOrFourByteAs.AsIndex) + "']"
    twoByteAsOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAsOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAsOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAsOrFourByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAsOrFourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAsOrFourByteAs.EntityData.Leafs["as"] = types.YLeaf{"As", twoByteAsOrFourByteAs.As}
    twoByteAsOrFourByteAs.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", twoByteAsOrFourByteAs.AsIndex}
    return &(twoByteAsOrFourByteAs.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_Ipv4Address
// ipv4 address
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}
}

func (ipv4Address *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpRouteTargets_Mp2MpRouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "mp2mp-route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + "[address='" + fmt.Sprintf("%v", ipv4Address.Address) + "']" + "[addr-index='" + fmt.Sprintf("%v", ipv4Address.AddrIndex) + "']"
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = make(map[string]types.YChild)
    ipv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4Address.Address}
    ipv4Address.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", ipv4Address.AddrIndex}
    return &(ipv4Address.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol
// signaling protocol in this MP2MP
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Customer Edge Identifier. The type is interface{} with range:
    // 11..100.
    CeRange interface{}

    // Enable signaling protocol. The type is interface{}.
    Enable interface{}

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_FlowLabelLoadBalance

    // Local Customer Edge Identifier Table.
    Ceids L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids
}

func (mp2MpSignalingProtocol *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    mp2MpSignalingProtocol.EntityData.YFilter = mp2MpSignalingProtocol.YFilter
    mp2MpSignalingProtocol.EntityData.YangName = "mp2mp-signaling-protocol"
    mp2MpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    mp2MpSignalingProtocol.EntityData.ParentYangName = "mp2mp-auto-discovery"
    mp2MpSignalingProtocol.EntityData.SegmentPath = "mp2mp-signaling-protocol"
    mp2MpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mp2MpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mp2MpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mp2MpSignalingProtocol.EntityData.Children = make(map[string]types.YChild)
    mp2MpSignalingProtocol.EntityData.Children["flow-label-load-balance"] = types.YChild{"FlowLabelLoadBalance", &mp2MpSignalingProtocol.FlowLabelLoadBalance}
    mp2MpSignalingProtocol.EntityData.Children["ceids"] = types.YChild{"Ceids", &mp2MpSignalingProtocol.Ceids}
    mp2MpSignalingProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    mp2MpSignalingProtocol.EntityData.Leafs["ce-range"] = types.YLeaf{"CeRange", mp2MpSignalingProtocol.CeRange}
    mp2MpSignalingProtocol.EntityData.Leafs["enable"] = types.YLeaf{"Enable", mp2MpSignalingProtocol.Enable}
    return &(mp2MpSignalingProtocol.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "mp2mp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = make(map[string]types.YChild)
    flowLabelLoadBalance.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabelLoadBalance.EntityData.Leafs["flow-label"] = types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel}
    flowLabelLoadBalance.EntityData.Leafs["static"] = types.YLeaf{"Static", flowLabelLoadBalance.Static}
    return &(flowLabelLoadBalance.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids
// Local Customer Edge Identifier Table
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Customer Edge Identifier . The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid.
    Ceid []L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid
}

func (ceids *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids) GetEntityData() *types.CommonEntityData {
    ceids.EntityData.YFilter = ceids.YFilter
    ceids.EntityData.YangName = "ceids"
    ceids.EntityData.BundleName = "cisco_ios_xr"
    ceids.EntityData.ParentYangName = "mp2mp-signaling-protocol"
    ceids.EntityData.SegmentPath = "ceids"
    ceids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ceids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ceids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ceids.EntityData.Children = make(map[string]types.YChild)
    ceids.EntityData.Children["ceid"] = types.YChild{"Ceid", nil}
    for i := range ceids.Ceid {
        ceids.EntityData.Children[types.GetSegmentPath(&ceids.Ceid[i])] = types.YChild{"Ceid", &ceids.Ceid[i]}
    }
    ceids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceids.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid
// Local Customer Edge Identifier 
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local Customer Edge Identifier. The type is
    // interface{} with range: 1..16384.
    CeId interface{}

    // AC And Remote Customer Edge Identifier Table.
    RemoteCeidAttachmentCircuits L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits
}

func (ceid *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid) GetEntityData() *types.CommonEntityData {
    ceid.EntityData.YFilter = ceid.YFilter
    ceid.EntityData.YangName = "ceid"
    ceid.EntityData.BundleName = "cisco_ios_xr"
    ceid.EntityData.ParentYangName = "ceids"
    ceid.EntityData.SegmentPath = "ceid" + "[ce-id='" + fmt.Sprintf("%v", ceid.CeId) + "']"
    ceid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ceid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ceid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ceid.EntityData.Children = make(map[string]types.YChild)
    ceid.EntityData.Children["remote-ceid-attachment-circuits"] = types.YChild{"RemoteCeidAttachmentCircuits", &ceid.RemoteCeidAttachmentCircuits}
    ceid.EntityData.Leafs = make(map[string]types.YLeaf)
    ceid.EntityData.Leafs["ce-id"] = types.YLeaf{"CeId", ceid.CeId}
    return &(ceid.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits
// AC And Remote Customer Edge Identifier
// Table
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AC And Remote Customer Edge Identifier. The type is slice of
    // L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit.
    RemoteCeidAttachmentCircuit []L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit
}

func (remoteCeidAttachmentCircuits *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    remoteCeidAttachmentCircuits.EntityData.YFilter = remoteCeidAttachmentCircuits.YFilter
    remoteCeidAttachmentCircuits.EntityData.YangName = "remote-ceid-attachment-circuits"
    remoteCeidAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    remoteCeidAttachmentCircuits.EntityData.ParentYangName = "ceid"
    remoteCeidAttachmentCircuits.EntityData.SegmentPath = "remote-ceid-attachment-circuits"
    remoteCeidAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteCeidAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteCeidAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteCeidAttachmentCircuits.EntityData.Children = make(map[string]types.YChild)
    remoteCeidAttachmentCircuits.EntityData.Children["remote-ceid-attachment-circuit"] = types.YChild{"RemoteCeidAttachmentCircuit", nil}
    for i := range remoteCeidAttachmentCircuits.RemoteCeidAttachmentCircuit {
        remoteCeidAttachmentCircuits.EntityData.Children[types.GetSegmentPath(&remoteCeidAttachmentCircuits.RemoteCeidAttachmentCircuit[i])] = types.YChild{"RemoteCeidAttachmentCircuit", &remoteCeidAttachmentCircuits.RemoteCeidAttachmentCircuit[i]}
    }
    remoteCeidAttachmentCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(remoteCeidAttachmentCircuits.EntityData)
}

// L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit
// AC And Remote Customer Edge Identifier
type L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Attachment Circuit. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // This attribute is a key. Remote Customer Edge Identifier. The type is
    // interface{} with range: 1..16384.
    RemoteCeId interface{}
}

func (remoteCeidAttachmentCircuit *L2Vpn_Database_XconnectGroups_XconnectGroup_Mp2MpXconnects_Mp2MpXconnect_Mp2MpAutoDiscovery_Mp2MpSignalingProtocol_Ceids_Ceid_RemoteCeidAttachmentCircuits_RemoteCeidAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    remoteCeidAttachmentCircuit.EntityData.YFilter = remoteCeidAttachmentCircuit.YFilter
    remoteCeidAttachmentCircuit.EntityData.YangName = "remote-ceid-attachment-circuit"
    remoteCeidAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    remoteCeidAttachmentCircuit.EntityData.ParentYangName = "remote-ceid-attachment-circuits"
    remoteCeidAttachmentCircuit.EntityData.SegmentPath = "remote-ceid-attachment-circuit" + "[name='" + fmt.Sprintf("%v", remoteCeidAttachmentCircuit.Name) + "']" + "[remote-ce-id='" + fmt.Sprintf("%v", remoteCeidAttachmentCircuit.RemoteCeId) + "']"
    remoteCeidAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteCeidAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteCeidAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteCeidAttachmentCircuit.EntityData.Children = make(map[string]types.YChild)
    remoteCeidAttachmentCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteCeidAttachmentCircuit.EntityData.Leafs["name"] = types.YLeaf{"Name", remoteCeidAttachmentCircuit.Name}
    remoteCeidAttachmentCircuit.EntityData.Leafs["remote-ce-id"] = types.YLeaf{"RemoteCeId", remoteCeidAttachmentCircuit.RemoteCeId}
    return &(remoteCeidAttachmentCircuit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups
// List of bridge  groups
type L2Vpn_Database_BridgeDomainGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge group. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup.
    BridgeDomainGroup []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup
}

func (bridgeDomainGroups *L2Vpn_Database_BridgeDomainGroups) GetEntityData() *types.CommonEntityData {
    bridgeDomainGroups.EntityData.YFilter = bridgeDomainGroups.YFilter
    bridgeDomainGroups.EntityData.YangName = "bridge-domain-groups"
    bridgeDomainGroups.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainGroups.EntityData.ParentYangName = "database"
    bridgeDomainGroups.EntityData.SegmentPath = "bridge-domain-groups"
    bridgeDomainGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainGroups.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainGroups.EntityData.Children["bridge-domain-group"] = types.YChild{"BridgeDomainGroup", nil}
    for i := range bridgeDomainGroups.BridgeDomainGroup {
        bridgeDomainGroups.EntityData.Children[types.GetSegmentPath(&bridgeDomainGroups.BridgeDomainGroup[i])] = types.YChild{"BridgeDomainGroup", &bridgeDomainGroups.BridgeDomainGroup[i]}
    }
    bridgeDomainGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomainGroups.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup
// Bridge group
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Bridge group. The type is string with
    // length: 1..32.
    Name interface{}

    // List of Bridge Domain.
    BridgeDomains L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains
}

func (bridgeDomainGroup *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup) GetEntityData() *types.CommonEntityData {
    bridgeDomainGroup.EntityData.YFilter = bridgeDomainGroup.YFilter
    bridgeDomainGroup.EntityData.YangName = "bridge-domain-group"
    bridgeDomainGroup.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainGroup.EntityData.ParentYangName = "bridge-domain-groups"
    bridgeDomainGroup.EntityData.SegmentPath = "bridge-domain-group" + "[name='" + fmt.Sprintf("%v", bridgeDomainGroup.Name) + "']"
    bridgeDomainGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainGroup.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainGroup.EntityData.Children["bridge-domains"] = types.YChild{"BridgeDomains", &bridgeDomainGroup.BridgeDomains}
    bridgeDomainGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomainGroup.EntityData.Leafs["name"] = types.YLeaf{"Name", bridgeDomainGroup.Name}
    return &(bridgeDomainGroup.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains
// List of Bridge Domain
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bridge domain. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain.
    BridgeDomain []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain
}

func (bridgeDomains *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomains.EntityData.ParentYangName = "bridge-domain-group"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomains.EntityData.Children = make(map[string]types.YChild)
    bridgeDomains.EntityData.Children["bridge-domain"] = types.YChild{"BridgeDomain", nil}
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children[types.GetSegmentPath(&bridgeDomains.BridgeDomain[i])] = types.YChild{"BridgeDomain", &bridgeDomains.BridgeDomain[i]}
    }
    bridgeDomains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomains.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain
// bridge domain
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain struct {
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
    BdStormControls L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls

    // Bridge Domain VxLAN Network Identifier Table.
    MemberVnis L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis

    // MAC configuration commands.
    BridgeDomainMac L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac

    // nV Satellite.
    NvSatellite L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite

    // Bridge Domain PBB.
    BridgeDomainPbb L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb

    // Bridge Domain EVI Table.
    BridgeDomainEvis L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis

    // Specify the access virtual forwarding interface name.
    AccessVfis L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis

    // List of pseudowires.
    BdPseudowires L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires

    // Specify the virtual forwarding interface name.
    Vfis L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis

    // Attachment Circuit table.
    BdAttachmentCircuits L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits

    // List of EVPN pseudowires.
    BdPseudowireEvpns L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns

    // IP Source Guard.
    IpSourceGuard L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard

    // Dynamic ARP Inspection.
    Dai L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai

    // Bridge Domain Routed Interface Table.
    RoutedInterfaces L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces
}

func (bridgeDomain *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomain.EntityData.ParentYangName = "bridge-domains"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + "[name='" + fmt.Sprintf("%v", bridgeDomain.Name) + "']"
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomain.EntityData.Children = make(map[string]types.YChild)
    bridgeDomain.EntityData.Children["bd-storm-controls"] = types.YChild{"BdStormControls", &bridgeDomain.BdStormControls}
    bridgeDomain.EntityData.Children["member-vnis"] = types.YChild{"MemberVnis", &bridgeDomain.MemberVnis}
    bridgeDomain.EntityData.Children["bridge-domain-mac"] = types.YChild{"BridgeDomainMac", &bridgeDomain.BridgeDomainMac}
    bridgeDomain.EntityData.Children["nv-satellite"] = types.YChild{"NvSatellite", &bridgeDomain.NvSatellite}
    bridgeDomain.EntityData.Children["bridge-domain-pbb"] = types.YChild{"BridgeDomainPbb", &bridgeDomain.BridgeDomainPbb}
    bridgeDomain.EntityData.Children["bridge-domain-evis"] = types.YChild{"BridgeDomainEvis", &bridgeDomain.BridgeDomainEvis}
    bridgeDomain.EntityData.Children["access-vfis"] = types.YChild{"AccessVfis", &bridgeDomain.AccessVfis}
    bridgeDomain.EntityData.Children["bd-pseudowires"] = types.YChild{"BdPseudowires", &bridgeDomain.BdPseudowires}
    bridgeDomain.EntityData.Children["vfis"] = types.YChild{"Vfis", &bridgeDomain.Vfis}
    bridgeDomain.EntityData.Children["bd-attachment-circuits"] = types.YChild{"BdAttachmentCircuits", &bridgeDomain.BdAttachmentCircuits}
    bridgeDomain.EntityData.Children["bd-pseudowire-evpns"] = types.YChild{"BdPseudowireEvpns", &bridgeDomain.BdPseudowireEvpns}
    bridgeDomain.EntityData.Children["ip-source-guard"] = types.YChild{"IpSourceGuard", &bridgeDomain.IpSourceGuard}
    bridgeDomain.EntityData.Children["dai"] = types.YChild{"Dai", &bridgeDomain.Dai}
    bridgeDomain.EntityData.Children["routed-interfaces"] = types.YChild{"RoutedInterfaces", &bridgeDomain.RoutedInterfaces}
    bridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomain.EntityData.Leafs["name"] = types.YLeaf{"Name", bridgeDomain.Name}
    bridgeDomain.EntityData.Leafs["coupled-mode"] = types.YLeaf{"CoupledMode", bridgeDomain.CoupledMode}
    bridgeDomain.EntityData.Leafs["shutdown"] = types.YLeaf{"Shutdown", bridgeDomain.Shutdown}
    bridgeDomain.EntityData.Leafs["flooding-unknown-unicast"] = types.YLeaf{"FloodingUnknownUnicast", bridgeDomain.FloodingUnknownUnicast}
    bridgeDomain.EntityData.Leafs["igmp-snooping-disable"] = types.YLeaf{"IgmpSnoopingDisable", bridgeDomain.IgmpSnoopingDisable}
    bridgeDomain.EntityData.Leafs["transport-mode"] = types.YLeaf{"TransportMode", bridgeDomain.TransportMode}
    bridgeDomain.EntityData.Leafs["mld-snooping"] = types.YLeaf{"MldSnooping", bridgeDomain.MldSnooping}
    bridgeDomain.EntityData.Leafs["bridge-domain-mtu"] = types.YLeaf{"BridgeDomainMtu", bridgeDomain.BridgeDomainMtu}
    bridgeDomain.EntityData.Leafs["dhcp"] = types.YLeaf{"Dhcp", bridgeDomain.Dhcp}
    bridgeDomain.EntityData.Leafs["bridge-description"] = types.YLeaf{"BridgeDescription", bridgeDomain.BridgeDescription}
    bridgeDomain.EntityData.Leafs["igmp-snooping"] = types.YLeaf{"IgmpSnooping", bridgeDomain.IgmpSnooping}
    bridgeDomain.EntityData.Leafs["flooding"] = types.YLeaf{"Flooding", bridgeDomain.Flooding}
    return &(bridgeDomain.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls
// Storm Control
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl.
    BdStormControl []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl
}

func (bdStormControls *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls) GetEntityData() *types.CommonEntityData {
    bdStormControls.EntityData.YFilter = bdStormControls.YFilter
    bdStormControls.EntityData.YangName = "bd-storm-controls"
    bdStormControls.EntityData.BundleName = "cisco_ios_xr"
    bdStormControls.EntityData.ParentYangName = "bridge-domain"
    bdStormControls.EntityData.SegmentPath = "bd-storm-controls"
    bdStormControls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdStormControls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdStormControls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdStormControls.EntityData.Children = make(map[string]types.YChild)
    bdStormControls.EntityData.Children["bd-storm-control"] = types.YChild{"BdStormControl", nil}
    for i := range bdStormControls.BdStormControl {
        bdStormControls.EntityData.Children[types.GetSegmentPath(&bdStormControls.BdStormControl[i])] = types.YChild{"BdStormControl", &bdStormControls.BdStormControl[i]}
    }
    bdStormControls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdStormControls.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl
// Storm Control Type
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit
}

func (bdStormControl *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl) GetEntityData() *types.CommonEntityData {
    bdStormControl.EntityData.YFilter = bdStormControl.YFilter
    bdStormControl.EntityData.YangName = "bd-storm-control"
    bdStormControl.EntityData.BundleName = "cisco_ios_xr"
    bdStormControl.EntityData.ParentYangName = "bd-storm-controls"
    bdStormControl.EntityData.SegmentPath = "bd-storm-control" + "[sctype='" + fmt.Sprintf("%v", bdStormControl.Sctype) + "']"
    bdStormControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdStormControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdStormControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdStormControl.EntityData.Children = make(map[string]types.YChild)
    bdStormControl.EntityData.Children["storm-control-unit"] = types.YChild{"StormControlUnit", &bdStormControl.StormControlUnit}
    bdStormControl.EntityData.Leafs = make(map[string]types.YLeaf)
    bdStormControl.EntityData.Leafs["sctype"] = types.YLeaf{"Sctype", bdStormControl.Sctype}
    return &(bdStormControl.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit
// Specify units for Storm Control Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit struct {
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

func (stormControlUnit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdStormControls_BdStormControl_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bd-storm-control"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = make(map[string]types.YChild)
    stormControlUnit.EntityData.Leafs = make(map[string]types.YLeaf)
    stormControlUnit.EntityData.Leafs["kbits-per-sec"] = types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec}
    stormControlUnit.EntityData.Leafs["pkts-per-sec"] = types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec}
    return &(stormControlUnit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis
// Bridge Domain VxLAN Network Identifier
// Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain Member VxLAN Network Identifier . The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni.
    MemberVni []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni
}

func (memberVnis *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis) GetEntityData() *types.CommonEntityData {
    memberVnis.EntityData.YFilter = memberVnis.YFilter
    memberVnis.EntityData.YangName = "member-vnis"
    memberVnis.EntityData.BundleName = "cisco_ios_xr"
    memberVnis.EntityData.ParentYangName = "bridge-domain"
    memberVnis.EntityData.SegmentPath = "member-vnis"
    memberVnis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVnis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVnis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVnis.EntityData.Children = make(map[string]types.YChild)
    memberVnis.EntityData.Children["member-vni"] = types.YChild{"MemberVni", nil}
    for i := range memberVnis.MemberVni {
        memberVnis.EntityData.Children[types.GetSegmentPath(&memberVnis.MemberVni[i])] = types.YChild{"MemberVni", &memberVnis.MemberVni[i]}
    }
    memberVnis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memberVnis.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni
// Bridge Domain Member VxLAN Network
// Identifier 
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VxLAN Network Identifier number. The type is
    // interface{} with range: 1..16777215.
    Vni interface{}

    // Static Mac Address Table.
    MemberVniStaticMacAddresses L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses
}

func (memberVni *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni) GetEntityData() *types.CommonEntityData {
    memberVni.EntityData.YFilter = memberVni.YFilter
    memberVni.EntityData.YangName = "member-vni"
    memberVni.EntityData.BundleName = "cisco_ios_xr"
    memberVni.EntityData.ParentYangName = "member-vnis"
    memberVni.EntityData.SegmentPath = "member-vni" + "[vni='" + fmt.Sprintf("%v", memberVni.Vni) + "']"
    memberVni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVni.EntityData.Children = make(map[string]types.YChild)
    memberVni.EntityData.Children["member-vni-static-mac-addresses"] = types.YChild{"MemberVniStaticMacAddresses", &memberVni.MemberVniStaticMacAddresses}
    memberVni.EntityData.Leafs = make(map[string]types.YLeaf)
    memberVni.EntityData.Leafs["vni"] = types.YLeaf{"Vni", memberVni.Vni}
    return &(memberVni.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses
// Static Mac Address Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress.
    MemberVniStaticMacAddress []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress
}

func (memberVniStaticMacAddresses *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    memberVniStaticMacAddresses.EntityData.YFilter = memberVniStaticMacAddresses.YFilter
    memberVniStaticMacAddresses.EntityData.YangName = "member-vni-static-mac-addresses"
    memberVniStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    memberVniStaticMacAddresses.EntityData.ParentYangName = "member-vni"
    memberVniStaticMacAddresses.EntityData.SegmentPath = "member-vni-static-mac-addresses"
    memberVniStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVniStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVniStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVniStaticMacAddresses.EntityData.Children = make(map[string]types.YChild)
    memberVniStaticMacAddresses.EntityData.Children["member-vni-static-mac-address"] = types.YChild{"MemberVniStaticMacAddress", nil}
    for i := range memberVniStaticMacAddresses.MemberVniStaticMacAddress {
        memberVniStaticMacAddresses.EntityData.Children[types.GetSegmentPath(&memberVniStaticMacAddresses.MemberVniStaticMacAddress[i])] = types.YChild{"MemberVniStaticMacAddress", &memberVniStaticMacAddresses.MemberVniStaticMacAddress[i]}
    }
    memberVniStaticMacAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memberVniStaticMacAddresses.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress
// Static Mac Address Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Enable Static Mac Address Configuration. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopIp interface{}
}

func (memberVniStaticMacAddress *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_MemberVnis_MemberVni_MemberVniStaticMacAddresses_MemberVniStaticMacAddress) GetEntityData() *types.CommonEntityData {
    memberVniStaticMacAddress.EntityData.YFilter = memberVniStaticMacAddress.YFilter
    memberVniStaticMacAddress.EntityData.YangName = "member-vni-static-mac-address"
    memberVniStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    memberVniStaticMacAddress.EntityData.ParentYangName = "member-vni-static-mac-addresses"
    memberVniStaticMacAddress.EntityData.SegmentPath = "member-vni-static-mac-address" + "[mac-address='" + fmt.Sprintf("%v", memberVniStaticMacAddress.MacAddress) + "']"
    memberVniStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberVniStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberVniStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberVniStaticMacAddress.EntityData.Children = make(map[string]types.YChild)
    memberVniStaticMacAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    memberVniStaticMacAddress.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", memberVniStaticMacAddress.MacAddress}
    memberVniStaticMacAddress.EntityData.Leafs["next-hop-ip"] = types.YLeaf{"NextHopIp", memberVniStaticMacAddress.NextHopIp}
    return &(memberVniStaticMacAddress.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac
// MAC configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac struct {
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
    BdMacLimit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit

    // Filter Mac Address.
    BdMacFilters L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters

    // MAC Secure.
    MacSecure L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure

    // MAC-Aging configuration commands.
    BdMacAging L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging
}

func (bridgeDomainMac *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac) GetEntityData() *types.CommonEntityData {
    bridgeDomainMac.EntityData.YFilter = bridgeDomainMac.YFilter
    bridgeDomainMac.EntityData.YangName = "bridge-domain-mac"
    bridgeDomainMac.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainMac.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainMac.EntityData.SegmentPath = "bridge-domain-mac"
    bridgeDomainMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainMac.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainMac.EntityData.Children["bd-mac-limit"] = types.YChild{"BdMacLimit", &bridgeDomainMac.BdMacLimit}
    bridgeDomainMac.EntityData.Children["bd-mac-filters"] = types.YChild{"BdMacFilters", &bridgeDomainMac.BdMacFilters}
    bridgeDomainMac.EntityData.Children["mac-secure"] = types.YChild{"MacSecure", &bridgeDomainMac.MacSecure}
    bridgeDomainMac.EntityData.Children["bd-mac-aging"] = types.YChild{"BdMacAging", &bridgeDomainMac.BdMacAging}
    bridgeDomainMac.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomainMac.EntityData.Leafs["bd-mac-withdraw-relay"] = types.YLeaf{"BdMacWithdrawRelay", bridgeDomainMac.BdMacWithdrawRelay}
    bridgeDomainMac.EntityData.Leafs["bd-mac-withdraw-access-pw-disable"] = types.YLeaf{"BdMacWithdrawAccessPwDisable", bridgeDomainMac.BdMacWithdrawAccessPwDisable}
    bridgeDomainMac.EntityData.Leafs["bd-mac-port-down-flush"] = types.YLeaf{"BdMacPortDownFlush", bridgeDomainMac.BdMacPortDownFlush}
    bridgeDomainMac.EntityData.Leafs["bd-mac-withdraw"] = types.YLeaf{"BdMacWithdraw", bridgeDomainMac.BdMacWithdraw}
    bridgeDomainMac.EntityData.Leafs["bd-mac-withdraw-behavior"] = types.YLeaf{"BdMacWithdrawBehavior", bridgeDomainMac.BdMacWithdrawBehavior}
    bridgeDomainMac.EntityData.Leafs["bd-mac-learn"] = types.YLeaf{"BdMacLearn", bridgeDomainMac.BdMacLearn}
    return &(bridgeDomainMac.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit
// MAC-Limit configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit struct {
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

func (bdMacLimit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacLimit) GetEntityData() *types.CommonEntityData {
    bdMacLimit.EntityData.YFilter = bdMacLimit.YFilter
    bdMacLimit.EntityData.YangName = "bd-mac-limit"
    bdMacLimit.EntityData.BundleName = "cisco_ios_xr"
    bdMacLimit.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacLimit.EntityData.SegmentPath = "bd-mac-limit"
    bdMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacLimit.EntityData.Children = make(map[string]types.YChild)
    bdMacLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    bdMacLimit.EntityData.Leafs["bd-mac-limit-action"] = types.YLeaf{"BdMacLimitAction", bdMacLimit.BdMacLimitAction}
    bdMacLimit.EntityData.Leafs["bd-mac-limit-notif"] = types.YLeaf{"BdMacLimitNotif", bdMacLimit.BdMacLimitNotif}
    bdMacLimit.EntityData.Leafs["bd-mac-limit-max"] = types.YLeaf{"BdMacLimitMax", bdMacLimit.BdMacLimitMax}
    return &(bdMacLimit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters
// Filter Mac Address
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static MAC address. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter.
    BdMacFilter []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter
}

func (bdMacFilters *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters) GetEntityData() *types.CommonEntityData {
    bdMacFilters.EntityData.YFilter = bdMacFilters.YFilter
    bdMacFilters.EntityData.YangName = "bd-mac-filters"
    bdMacFilters.EntityData.BundleName = "cisco_ios_xr"
    bdMacFilters.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacFilters.EntityData.SegmentPath = "bd-mac-filters"
    bdMacFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacFilters.EntityData.Children = make(map[string]types.YChild)
    bdMacFilters.EntityData.Children["bd-mac-filter"] = types.YChild{"BdMacFilter", nil}
    for i := range bdMacFilters.BdMacFilter {
        bdMacFilters.EntityData.Children[types.GetSegmentPath(&bdMacFilters.BdMacFilter[i])] = types.YChild{"BdMacFilter", &bdMacFilters.BdMacFilter[i]}
    }
    bdMacFilters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdMacFilters.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter
// Static MAC address
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}

    // MAC address for filtering. The type is interface{}.
    Drop interface{}
}

func (bdMacFilter *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacFilters_BdMacFilter) GetEntityData() *types.CommonEntityData {
    bdMacFilter.EntityData.YFilter = bdMacFilter.YFilter
    bdMacFilter.EntityData.YangName = "bd-mac-filter"
    bdMacFilter.EntityData.BundleName = "cisco_ios_xr"
    bdMacFilter.EntityData.ParentYangName = "bd-mac-filters"
    bdMacFilter.EntityData.SegmentPath = "bd-mac-filter" + "[address='" + fmt.Sprintf("%v", bdMacFilter.Address) + "']"
    bdMacFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacFilter.EntityData.Children = make(map[string]types.YChild)
    bdMacFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    bdMacFilter.EntityData.Leafs["address"] = types.YLeaf{"Address", bdMacFilter.Address}
    bdMacFilter.EntityData.Leafs["drop"] = types.YLeaf{"Drop", bdMacFilter.Drop}
    return &(bdMacFilter.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure
// MAC Secure
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure struct {
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

func (macSecure *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_MacSecure) GetEntityData() *types.CommonEntityData {
    macSecure.EntityData.YFilter = macSecure.YFilter
    macSecure.EntityData.YangName = "mac-secure"
    macSecure.EntityData.BundleName = "cisco_ios_xr"
    macSecure.EntityData.ParentYangName = "bridge-domain-mac"
    macSecure.EntityData.SegmentPath = "mac-secure"
    macSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSecure.EntityData.Children = make(map[string]types.YChild)
    macSecure.EntityData.Leafs = make(map[string]types.YLeaf)
    macSecure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", macSecure.Logging}
    macSecure.EntityData.Leafs["action"] = types.YLeaf{"Action", macSecure.Action}
    macSecure.EntityData.Leafs["enable"] = types.YLeaf{"Enable", macSecure.Enable}
    macSecure.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", macSecure.Threshold}
    return &(macSecure.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging
// MAC-Aging configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    BdMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    BdMacAgingTime interface{}
}

func (bdMacAging *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainMac_BdMacAging) GetEntityData() *types.CommonEntityData {
    bdMacAging.EntityData.YFilter = bdMacAging.YFilter
    bdMacAging.EntityData.YangName = "bd-mac-aging"
    bdMacAging.EntityData.BundleName = "cisco_ios_xr"
    bdMacAging.EntityData.ParentYangName = "bridge-domain-mac"
    bdMacAging.EntityData.SegmentPath = "bd-mac-aging"
    bdMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdMacAging.EntityData.Children = make(map[string]types.YChild)
    bdMacAging.EntityData.Leafs = make(map[string]types.YLeaf)
    bdMacAging.EntityData.Leafs["bd-mac-aging-type"] = types.YLeaf{"BdMacAgingType", bdMacAging.BdMacAgingType}
    bdMacAging.EntityData.Leafs["bd-mac-aging-time"] = types.YLeaf{"BdMacAgingTime", bdMacAging.BdMacAgingTime}
    return &(bdMacAging.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite
// nV Satellite
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPv4 Multicast Offload to Satellite Nodes. The type is interface{}.
    OffloadIpv4MulticastEnable interface{}

    // Enable nV Satellite Settings. The type is interface{}.
    Enable interface{}
}

func (nvSatellite *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_NvSatellite) GetEntityData() *types.CommonEntityData {
    nvSatellite.EntityData.YFilter = nvSatellite.YFilter
    nvSatellite.EntityData.YangName = "nv-satellite"
    nvSatellite.EntityData.BundleName = "cisco_ios_xr"
    nvSatellite.EntityData.ParentYangName = "bridge-domain"
    nvSatellite.EntityData.SegmentPath = "nv-satellite"
    nvSatellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellite.EntityData.Children = make(map[string]types.YChild)
    nvSatellite.EntityData.Leafs = make(map[string]types.YLeaf)
    nvSatellite.EntityData.Leafs["offload-ipv4-multicast-enable"] = types.YLeaf{"OffloadIpv4MulticastEnable", nvSatellite.OffloadIpv4MulticastEnable}
    nvSatellite.EntityData.Leafs["enable"] = types.YLeaf{"Enable", nvSatellite.Enable}
    return &(nvSatellite.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb
// Bridge Domain PBB
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Edge.
    PbbEdges L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges

    // PBB Core.
    PbbCore L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore
}

func (bridgeDomainPbb *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb) GetEntityData() *types.CommonEntityData {
    bridgeDomainPbb.EntityData.YFilter = bridgeDomainPbb.YFilter
    bridgeDomainPbb.EntityData.YangName = "bridge-domain-pbb"
    bridgeDomainPbb.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainPbb.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainPbb.EntityData.SegmentPath = "bridge-domain-pbb"
    bridgeDomainPbb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainPbb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainPbb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainPbb.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainPbb.EntityData.Children["pbb-edges"] = types.YChild{"PbbEdges", &bridgeDomainPbb.PbbEdges}
    bridgeDomainPbb.EntityData.Children["pbb-core"] = types.YChild{"PbbCore", &bridgeDomainPbb.PbbCore}
    bridgeDomainPbb.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomainPbb.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges
// PBB Edge
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure BD as PBB Edge with ISID and associated PBB Core BD. The type is
    // slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge.
    PbbEdge []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge
}

func (pbbEdges *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges) GetEntityData() *types.CommonEntityData {
    pbbEdges.EntityData.YFilter = pbbEdges.YFilter
    pbbEdges.EntityData.YangName = "pbb-edges"
    pbbEdges.EntityData.BundleName = "cisco_ios_xr"
    pbbEdges.EntityData.ParentYangName = "bridge-domain-pbb"
    pbbEdges.EntityData.SegmentPath = "pbb-edges"
    pbbEdges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdges.EntityData.Children = make(map[string]types.YChild)
    pbbEdges.EntityData.Children["pbb-edge"] = types.YChild{"PbbEdge", nil}
    for i := range pbbEdges.PbbEdge {
        pbbEdges.EntityData.Children[types.GetSegmentPath(&pbbEdges.PbbEdge[i])] = types.YChild{"PbbEdge", &pbbEdges.PbbEdge[i]}
    }
    pbbEdges.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pbbEdges.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge
// Configure BD as PBB Edge with ISID and
// associated PBB Core BD
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge struct {
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
    // string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    UnknownUnicastBmac interface{}

    // Split Horizon Group.
    PbbEdgeSplitHorizonGroup L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup

    // PBB Static Mac Address Mapping Table.
    PbbStaticMacMappings L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings

    // Attach a DHCP profile.
    PbbEdgeDhcpProfile L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile

    // MAC configuration commands.
    PbbEdgeMac L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac
}

func (pbbEdge *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge) GetEntityData() *types.CommonEntityData {
    pbbEdge.EntityData.YFilter = pbbEdge.YFilter
    pbbEdge.EntityData.YangName = "pbb-edge"
    pbbEdge.EntityData.BundleName = "cisco_ios_xr"
    pbbEdge.EntityData.ParentYangName = "pbb-edges"
    pbbEdge.EntityData.SegmentPath = "pbb-edge" + "[isid='" + fmt.Sprintf("%v", pbbEdge.Isid) + "']" + "[core-bd-name='" + fmt.Sprintf("%v", pbbEdge.CoreBdName) + "']"
    pbbEdge.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdge.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdge.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdge.EntityData.Children = make(map[string]types.YChild)
    pbbEdge.EntityData.Children["pbb-edge-split-horizon-group"] = types.YChild{"PbbEdgeSplitHorizonGroup", &pbbEdge.PbbEdgeSplitHorizonGroup}
    pbbEdge.EntityData.Children["pbb-static-mac-mappings"] = types.YChild{"PbbStaticMacMappings", &pbbEdge.PbbStaticMacMappings}
    pbbEdge.EntityData.Children["pbb-edge-dhcp-profile"] = types.YChild{"PbbEdgeDhcpProfile", &pbbEdge.PbbEdgeDhcpProfile}
    pbbEdge.EntityData.Children["pbb-edge-mac"] = types.YChild{"PbbEdgeMac", &pbbEdge.PbbEdgeMac}
    pbbEdge.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdge.EntityData.Leafs["isid"] = types.YLeaf{"Isid", pbbEdge.Isid}
    pbbEdge.EntityData.Leafs["core-bd-name"] = types.YLeaf{"CoreBdName", pbbEdge.CoreBdName}
    pbbEdge.EntityData.Leafs["pbb-edge-igmp-profile"] = types.YLeaf{"PbbEdgeIgmpProfile", pbbEdge.PbbEdgeIgmpProfile}
    pbbEdge.EntityData.Leafs["unknown-unicast-bmac"] = types.YLeaf{"UnknownUnicastBmac", pbbEdge.UnknownUnicastBmac}
    return &(pbbEdge.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup
// Split Horizon Group
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable split horizon group. The type is interface{}.
    Disable interface{}
}

func (pbbEdgeSplitHorizonGroup *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    pbbEdgeSplitHorizonGroup.EntityData.YFilter = pbbEdgeSplitHorizonGroup.YFilter
    pbbEdgeSplitHorizonGroup.EntityData.YangName = "pbb-edge-split-horizon-group"
    pbbEdgeSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeSplitHorizonGroup.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeSplitHorizonGroup.EntityData.SegmentPath = "pbb-edge-split-horizon-group"
    pbbEdgeSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeSplitHorizonGroup.EntityData.Children = make(map[string]types.YChild)
    pbbEdgeSplitHorizonGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdgeSplitHorizonGroup.EntityData.Leafs["disable"] = types.YLeaf{"Disable", pbbEdgeSplitHorizonGroup.Disable}
    return &(pbbEdgeSplitHorizonGroup.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings
// PBB Static Mac Address Mapping Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Static Mac Address Mapping Configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping.
    PbbStaticMacMapping []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping
}

func (pbbStaticMacMappings *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings) GetEntityData() *types.CommonEntityData {
    pbbStaticMacMappings.EntityData.YFilter = pbbStaticMacMappings.YFilter
    pbbStaticMacMappings.EntityData.YangName = "pbb-static-mac-mappings"
    pbbStaticMacMappings.EntityData.BundleName = "cisco_ios_xr"
    pbbStaticMacMappings.EntityData.ParentYangName = "pbb-edge"
    pbbStaticMacMappings.EntityData.SegmentPath = "pbb-static-mac-mappings"
    pbbStaticMacMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbStaticMacMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbStaticMacMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbStaticMacMappings.EntityData.Children = make(map[string]types.YChild)
    pbbStaticMacMappings.EntityData.Children["pbb-static-mac-mapping"] = types.YChild{"PbbStaticMacMapping", nil}
    for i := range pbbStaticMacMappings.PbbStaticMacMapping {
        pbbStaticMacMappings.EntityData.Children[types.GetSegmentPath(&pbbStaticMacMappings.PbbStaticMacMapping[i])] = types.YChild{"PbbStaticMacMapping", &pbbStaticMacMappings.PbbStaticMacMapping[i]}
    }
    pbbStaticMacMappings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pbbStaticMacMappings.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping
// PBB Static Mac Address Mapping
// Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}

    // Static backbone MAC address to map with. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PbbStaticMacMappingBmac interface{}
}

func (pbbStaticMacMapping *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbStaticMacMappings_PbbStaticMacMapping) GetEntityData() *types.CommonEntityData {
    pbbStaticMacMapping.EntityData.YFilter = pbbStaticMacMapping.YFilter
    pbbStaticMacMapping.EntityData.YangName = "pbb-static-mac-mapping"
    pbbStaticMacMapping.EntityData.BundleName = "cisco_ios_xr"
    pbbStaticMacMapping.EntityData.ParentYangName = "pbb-static-mac-mappings"
    pbbStaticMacMapping.EntityData.SegmentPath = "pbb-static-mac-mapping" + "[address='" + fmt.Sprintf("%v", pbbStaticMacMapping.Address) + "']"
    pbbStaticMacMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbStaticMacMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbStaticMacMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbStaticMacMapping.EntityData.Children = make(map[string]types.YChild)
    pbbStaticMacMapping.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbStaticMacMapping.EntityData.Leafs["address"] = types.YLeaf{"Address", pbbStaticMacMapping.Address}
    pbbStaticMacMapping.EntityData.Leafs["pbb-static-mac-mapping-bmac"] = types.YLeaf{"PbbStaticMacMappingBmac", pbbStaticMacMapping.PbbStaticMacMappingBmac}
    return &(pbbStaticMacMapping.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile
// Attach a DHCP profile
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pbbEdgeDhcpProfile *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeDhcpProfile) GetEntityData() *types.CommonEntityData {
    pbbEdgeDhcpProfile.EntityData.YFilter = pbbEdgeDhcpProfile.YFilter
    pbbEdgeDhcpProfile.EntityData.YangName = "pbb-edge-dhcp-profile"
    pbbEdgeDhcpProfile.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeDhcpProfile.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeDhcpProfile.EntityData.SegmentPath = "pbb-edge-dhcp-profile"
    pbbEdgeDhcpProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeDhcpProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeDhcpProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeDhcpProfile.EntityData.Children = make(map[string]types.YChild)
    pbbEdgeDhcpProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdgeDhcpProfile.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pbbEdgeDhcpProfile.ProfileId}
    pbbEdgeDhcpProfile.EntityData.Leafs["dhcp-snooping-id"] = types.YLeaf{"DhcpSnoopingId", pbbEdgeDhcpProfile.DhcpSnoopingId}
    return &(pbbEdgeDhcpProfile.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac
// MAC configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Mac Learning. The type is MacLearn.
    PbbEdgeMacLearning interface{}

    // MAC-Limit configuration commands.
    PbbEdgeMacLimit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit

    // MAC-Aging configuration commands.
    PbbEdgeMacAging L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging

    // MAC Secure.
    PbbEdgeMacSecure L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure
}

func (pbbEdgeMac *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac) GetEntityData() *types.CommonEntityData {
    pbbEdgeMac.EntityData.YFilter = pbbEdgeMac.YFilter
    pbbEdgeMac.EntityData.YangName = "pbb-edge-mac"
    pbbEdgeMac.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMac.EntityData.ParentYangName = "pbb-edge"
    pbbEdgeMac.EntityData.SegmentPath = "pbb-edge-mac"
    pbbEdgeMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMac.EntityData.Children = make(map[string]types.YChild)
    pbbEdgeMac.EntityData.Children["pbb-edge-mac-limit"] = types.YChild{"PbbEdgeMacLimit", &pbbEdgeMac.PbbEdgeMacLimit}
    pbbEdgeMac.EntityData.Children["pbb-edge-mac-aging"] = types.YChild{"PbbEdgeMacAging", &pbbEdgeMac.PbbEdgeMacAging}
    pbbEdgeMac.EntityData.Children["pbb-edge-mac-secure"] = types.YChild{"PbbEdgeMacSecure", &pbbEdgeMac.PbbEdgeMacSecure}
    pbbEdgeMac.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdgeMac.EntityData.Leafs["pbb-edge-mac-learning"] = types.YLeaf{"PbbEdgeMacLearning", pbbEdgeMac.PbbEdgeMacLearning}
    return &(pbbEdgeMac.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit
// MAC-Limit configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit struct {
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

func (pbbEdgeMacLimit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacLimit) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacLimit.EntityData.YFilter = pbbEdgeMacLimit.YFilter
    pbbEdgeMacLimit.EntityData.YangName = "pbb-edge-mac-limit"
    pbbEdgeMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacLimit.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacLimit.EntityData.SegmentPath = "pbb-edge-mac-limit"
    pbbEdgeMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacLimit.EntityData.Children = make(map[string]types.YChild)
    pbbEdgeMacLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdgeMacLimit.EntityData.Leafs["pbb-edge-mac-limit-action"] = types.YLeaf{"PbbEdgeMacLimitAction", pbbEdgeMacLimit.PbbEdgeMacLimitAction}
    pbbEdgeMacLimit.EntityData.Leafs["pbb-edge-mac-limit-max"] = types.YLeaf{"PbbEdgeMacLimitMax", pbbEdgeMacLimit.PbbEdgeMacLimitMax}
    pbbEdgeMacLimit.EntityData.Leafs["pbb-edge-mac-limit-notif"] = types.YLeaf{"PbbEdgeMacLimitNotif", pbbEdgeMacLimit.PbbEdgeMacLimitNotif}
    return &(pbbEdgeMacLimit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging
// MAC-Aging configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PbbEdgeMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    PbbEdgeMacAgingTime interface{}
}

func (pbbEdgeMacAging *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacAging) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacAging.EntityData.YFilter = pbbEdgeMacAging.YFilter
    pbbEdgeMacAging.EntityData.YangName = "pbb-edge-mac-aging"
    pbbEdgeMacAging.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacAging.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacAging.EntityData.SegmentPath = "pbb-edge-mac-aging"
    pbbEdgeMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacAging.EntityData.Children = make(map[string]types.YChild)
    pbbEdgeMacAging.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdgeMacAging.EntityData.Leafs["pbb-edge-mac-aging-type"] = types.YLeaf{"PbbEdgeMacAgingType", pbbEdgeMacAging.PbbEdgeMacAgingType}
    pbbEdgeMacAging.EntityData.Leafs["pbb-edge-mac-aging-time"] = types.YLeaf{"PbbEdgeMacAgingTime", pbbEdgeMacAging.PbbEdgeMacAgingTime}
    return &(pbbEdgeMacAging.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure
// MAC Secure
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure struct {
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

func (pbbEdgeMacSecure *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbEdges_PbbEdge_PbbEdgeMac_PbbEdgeMacSecure) GetEntityData() *types.CommonEntityData {
    pbbEdgeMacSecure.EntityData.YFilter = pbbEdgeMacSecure.YFilter
    pbbEdgeMacSecure.EntityData.YangName = "pbb-edge-mac-secure"
    pbbEdgeMacSecure.EntityData.BundleName = "cisco_ios_xr"
    pbbEdgeMacSecure.EntityData.ParentYangName = "pbb-edge-mac"
    pbbEdgeMacSecure.EntityData.SegmentPath = "pbb-edge-mac-secure"
    pbbEdgeMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbEdgeMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbEdgeMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbEdgeMacSecure.EntityData.Children = make(map[string]types.YChild)
    pbbEdgeMacSecure.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbEdgeMacSecure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", pbbEdgeMacSecure.Logging}
    pbbEdgeMacSecure.EntityData.Leafs["disable"] = types.YLeaf{"Disable", pbbEdgeMacSecure.Disable}
    pbbEdgeMacSecure.EntityData.Leafs["action"] = types.YLeaf{"Action", pbbEdgeMacSecure.Action}
    pbbEdgeMacSecure.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pbbEdgeMacSecure.Enable}
    pbbEdgeMacSecure.EntityData.Leafs["accept-shutdown"] = types.YLeaf{"AcceptShutdown", pbbEdgeMacSecure.AcceptShutdown}
    return &(pbbEdgeMacSecure.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore
// PBB Core
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore struct {
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
    PbbCoreMac L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac

    // PBB Core EVI Table.
    PbbCoreEvis L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis

    // Attach a DHCP profile.
    PbbCoreDhcpProfile L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile
}

func (pbbCore *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore) GetEntityData() *types.CommonEntityData {
    pbbCore.EntityData.YFilter = pbbCore.YFilter
    pbbCore.EntityData.YangName = "pbb-core"
    pbbCore.EntityData.BundleName = "cisco_ios_xr"
    pbbCore.EntityData.ParentYangName = "bridge-domain-pbb"
    pbbCore.EntityData.SegmentPath = "pbb-core"
    pbbCore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCore.EntityData.Children = make(map[string]types.YChild)
    pbbCore.EntityData.Children["pbb-core-mac"] = types.YChild{"PbbCoreMac", &pbbCore.PbbCoreMac}
    pbbCore.EntityData.Children["pbb-core-evis"] = types.YChild{"PbbCoreEvis", &pbbCore.PbbCoreEvis}
    pbbCore.EntityData.Children["pbb-core-dhcp-profile"] = types.YChild{"PbbCoreDhcpProfile", &pbbCore.PbbCoreDhcpProfile}
    pbbCore.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbCore.EntityData.Leafs["pbb-core-mmrp-flood-optimization"] = types.YLeaf{"PbbCoreMmrpFloodOptimization", pbbCore.PbbCoreMmrpFloodOptimization}
    pbbCore.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", pbbCore.VlanId}
    pbbCore.EntityData.Leafs["pbb-core-igmp-profile"] = types.YLeaf{"PbbCoreIgmpProfile", pbbCore.PbbCoreIgmpProfile}
    pbbCore.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pbbCore.Enable}
    return &(pbbCore.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac
// MAC configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Mac Learning. The type is MacLearn.
    PbbCoreMacLearning interface{}

    // MAC-Aging configuration commands.
    PbbCoreMacAging L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging

    // MAC-Limit configuration commands.
    PbbCoreMacLimit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit
}

func (pbbCoreMac *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac) GetEntityData() *types.CommonEntityData {
    pbbCoreMac.EntityData.YFilter = pbbCoreMac.YFilter
    pbbCoreMac.EntityData.YangName = "pbb-core-mac"
    pbbCoreMac.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMac.EntityData.ParentYangName = "pbb-core"
    pbbCoreMac.EntityData.SegmentPath = "pbb-core-mac"
    pbbCoreMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMac.EntityData.Children = make(map[string]types.YChild)
    pbbCoreMac.EntityData.Children["pbb-core-mac-aging"] = types.YChild{"PbbCoreMacAging", &pbbCoreMac.PbbCoreMacAging}
    pbbCoreMac.EntityData.Children["pbb-core-mac-limit"] = types.YChild{"PbbCoreMacLimit", &pbbCoreMac.PbbCoreMacLimit}
    pbbCoreMac.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbCoreMac.EntityData.Leafs["pbb-core-mac-learning"] = types.YLeaf{"PbbCoreMacLearning", pbbCoreMac.PbbCoreMacLearning}
    return &(pbbCoreMac.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging
// MAC-Aging configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PbbCoreMacAgingType interface{}

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    PbbCoreMacAgingTime interface{}
}

func (pbbCoreMacAging *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacAging) GetEntityData() *types.CommonEntityData {
    pbbCoreMacAging.EntityData.YFilter = pbbCoreMacAging.YFilter
    pbbCoreMacAging.EntityData.YangName = "pbb-core-mac-aging"
    pbbCoreMacAging.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMacAging.EntityData.ParentYangName = "pbb-core-mac"
    pbbCoreMacAging.EntityData.SegmentPath = "pbb-core-mac-aging"
    pbbCoreMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMacAging.EntityData.Children = make(map[string]types.YChild)
    pbbCoreMacAging.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbCoreMacAging.EntityData.Leafs["pbb-core-mac-aging-type"] = types.YLeaf{"PbbCoreMacAgingType", pbbCoreMacAging.PbbCoreMacAgingType}
    pbbCoreMacAging.EntityData.Leafs["pbb-core-mac-aging-time"] = types.YLeaf{"PbbCoreMacAgingTime", pbbCoreMacAging.PbbCoreMacAgingTime}
    return &(pbbCoreMacAging.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit
// MAC-Limit configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit struct {
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

func (pbbCoreMacLimit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreMac_PbbCoreMacLimit) GetEntityData() *types.CommonEntityData {
    pbbCoreMacLimit.EntityData.YFilter = pbbCoreMacLimit.YFilter
    pbbCoreMacLimit.EntityData.YangName = "pbb-core-mac-limit"
    pbbCoreMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreMacLimit.EntityData.ParentYangName = "pbb-core-mac"
    pbbCoreMacLimit.EntityData.SegmentPath = "pbb-core-mac-limit"
    pbbCoreMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreMacLimit.EntityData.Children = make(map[string]types.YChild)
    pbbCoreMacLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbCoreMacLimit.EntityData.Leafs["pbb-core-mac-limit-max"] = types.YLeaf{"PbbCoreMacLimitMax", pbbCoreMacLimit.PbbCoreMacLimitMax}
    pbbCoreMacLimit.EntityData.Leafs["pbb-core-mac-limit-notif"] = types.YLeaf{"PbbCoreMacLimitNotif", pbbCoreMacLimit.PbbCoreMacLimitNotif}
    pbbCoreMacLimit.EntityData.Leafs["pbb-core-mac-limit-action"] = types.YLeaf{"PbbCoreMacLimitAction", pbbCoreMacLimit.PbbCoreMacLimitAction}
    return &(pbbCoreMacLimit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis
// PBB Core EVI Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBB Core EVI. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi.
    PbbCoreEvi []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi
}

func (pbbCoreEvis *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis) GetEntityData() *types.CommonEntityData {
    pbbCoreEvis.EntityData.YFilter = pbbCoreEvis.YFilter
    pbbCoreEvis.EntityData.YangName = "pbb-core-evis"
    pbbCoreEvis.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreEvis.EntityData.ParentYangName = "pbb-core"
    pbbCoreEvis.EntityData.SegmentPath = "pbb-core-evis"
    pbbCoreEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreEvis.EntityData.Children = make(map[string]types.YChild)
    pbbCoreEvis.EntityData.Children["pbb-core-evi"] = types.YChild{"PbbCoreEvi", nil}
    for i := range pbbCoreEvis.PbbCoreEvi {
        pbbCoreEvis.EntityData.Children[types.GetSegmentPath(&pbbCoreEvis.PbbCoreEvi[i])] = types.YChild{"PbbCoreEvi", &pbbCoreEvis.PbbCoreEvi[i]}
    }
    pbbCoreEvis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pbbCoreEvis.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi
// PBB Core EVI
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..4294967295.
    Eviid interface{}
}

func (pbbCoreEvi *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreEvis_PbbCoreEvi) GetEntityData() *types.CommonEntityData {
    pbbCoreEvi.EntityData.YFilter = pbbCoreEvi.YFilter
    pbbCoreEvi.EntityData.YangName = "pbb-core-evi"
    pbbCoreEvi.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreEvi.EntityData.ParentYangName = "pbb-core-evis"
    pbbCoreEvi.EntityData.SegmentPath = "pbb-core-evi" + "[eviid='" + fmt.Sprintf("%v", pbbCoreEvi.Eviid) + "']"
    pbbCoreEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreEvi.EntityData.Children = make(map[string]types.YChild)
    pbbCoreEvi.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbCoreEvi.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", pbbCoreEvi.Eviid}
    return &(pbbCoreEvi.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile
// Attach a DHCP profile
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pbbCoreDhcpProfile *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainPbb_PbbCore_PbbCoreDhcpProfile) GetEntityData() *types.CommonEntityData {
    pbbCoreDhcpProfile.EntityData.YFilter = pbbCoreDhcpProfile.YFilter
    pbbCoreDhcpProfile.EntityData.YangName = "pbb-core-dhcp-profile"
    pbbCoreDhcpProfile.EntityData.BundleName = "cisco_ios_xr"
    pbbCoreDhcpProfile.EntityData.ParentYangName = "pbb-core"
    pbbCoreDhcpProfile.EntityData.SegmentPath = "pbb-core-dhcp-profile"
    pbbCoreDhcpProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbbCoreDhcpProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbbCoreDhcpProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbbCoreDhcpProfile.EntityData.Children = make(map[string]types.YChild)
    pbbCoreDhcpProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    pbbCoreDhcpProfile.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pbbCoreDhcpProfile.ProfileId}
    pbbCoreDhcpProfile.EntityData.Leafs["dhcp-snooping-id"] = types.YLeaf{"DhcpSnoopingId", pbbCoreDhcpProfile.DhcpSnoopingId}
    return &(pbbCoreDhcpProfile.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis
// Bridge Domain EVI Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain EVI. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi.
    BridgeDomainEvi []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi
}

func (bridgeDomainEvis *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis) GetEntityData() *types.CommonEntityData {
    bridgeDomainEvis.EntityData.YFilter = bridgeDomainEvis.YFilter
    bridgeDomainEvis.EntityData.YangName = "bridge-domain-evis"
    bridgeDomainEvis.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainEvis.EntityData.ParentYangName = "bridge-domain"
    bridgeDomainEvis.EntityData.SegmentPath = "bridge-domain-evis"
    bridgeDomainEvis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainEvis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainEvis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainEvis.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainEvis.EntityData.Children["bridge-domain-evi"] = types.YChild{"BridgeDomainEvi", nil}
    for i := range bridgeDomainEvis.BridgeDomainEvi {
        bridgeDomainEvis.EntityData.Children[types.GetSegmentPath(&bridgeDomainEvis.BridgeDomainEvi[i])] = types.YChild{"BridgeDomainEvi", &bridgeDomainEvis.BridgeDomainEvi[i]}
    }
    bridgeDomainEvis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomainEvis.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi
// Bridge Domain EVI
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}
}

func (bridgeDomainEvi *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BridgeDomainEvis_BridgeDomainEvi) GetEntityData() *types.CommonEntityData {
    bridgeDomainEvi.EntityData.YFilter = bridgeDomainEvi.YFilter
    bridgeDomainEvi.EntityData.YangName = "bridge-domain-evi"
    bridgeDomainEvi.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainEvi.EntityData.ParentYangName = "bridge-domain-evis"
    bridgeDomainEvi.EntityData.SegmentPath = "bridge-domain-evi" + "[eviid='" + fmt.Sprintf("%v", bridgeDomainEvi.Eviid) + "']"
    bridgeDomainEvi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainEvi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainEvi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainEvi.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainEvi.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomainEvi.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", bridgeDomainEvi.Eviid}
    return &(bridgeDomainEvi.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis
// Specify the access virtual forwarding
// interface name
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Acess Virtual Forwarding Interface. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi.
    AccessVfi []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi
}

func (accessVfis *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis) GetEntityData() *types.CommonEntityData {
    accessVfis.EntityData.YFilter = accessVfis.YFilter
    accessVfis.EntityData.YangName = "access-vfis"
    accessVfis.EntityData.BundleName = "cisco_ios_xr"
    accessVfis.EntityData.ParentYangName = "bridge-domain"
    accessVfis.EntityData.SegmentPath = "access-vfis"
    accessVfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfis.EntityData.Children = make(map[string]types.YChild)
    accessVfis.EntityData.Children["access-vfi"] = types.YChild{"AccessVfi", nil}
    for i := range accessVfis.AccessVfi {
        accessVfis.EntityData.Children[types.GetSegmentPath(&accessVfis.AccessVfi[i])] = types.YChild{"AccessVfi", &accessVfis.AccessVfi[i]}
    }
    accessVfis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessVfis.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi
// Name of the Acess Virtual Forwarding
// Interface
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the AccessVirtual Forwarding Interface.
    // The type is string with length: 1..32.
    Name interface{}

    // shutdown the AccessVfi. The type is interface{}.
    AccessVfiShutdown interface{}

    // List of pseudowires.
    AccessVfiPseudowires L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires
}

func (accessVfi *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi) GetEntityData() *types.CommonEntityData {
    accessVfi.EntityData.YFilter = accessVfi.YFilter
    accessVfi.EntityData.YangName = "access-vfi"
    accessVfi.EntityData.BundleName = "cisco_ios_xr"
    accessVfi.EntityData.ParentYangName = "access-vfis"
    accessVfi.EntityData.SegmentPath = "access-vfi" + "[name='" + fmt.Sprintf("%v", accessVfi.Name) + "']"
    accessVfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfi.EntityData.Children = make(map[string]types.YChild)
    accessVfi.EntityData.Children["access-vfi-pseudowires"] = types.YChild{"AccessVfiPseudowires", &accessVfi.AccessVfiPseudowires}
    accessVfi.EntityData.Leafs = make(map[string]types.YLeaf)
    accessVfi.EntityData.Leafs["name"] = types.YLeaf{"Name", accessVfi.Name}
    accessVfi.EntityData.Leafs["access-vfi-shutdown"] = types.YLeaf{"AccessVfiShutdown", accessVfi.AccessVfiShutdown}
    return &(accessVfi.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires
// List of pseudowires
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire.
    AccessVfiPseudowire []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire
}

func (accessVfiPseudowires *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowires.EntityData.YFilter = accessVfiPseudowires.YFilter
    accessVfiPseudowires.EntityData.YangName = "access-vfi-pseudowires"
    accessVfiPseudowires.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowires.EntityData.ParentYangName = "access-vfi"
    accessVfiPseudowires.EntityData.SegmentPath = "access-vfi-pseudowires"
    accessVfiPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowires.EntityData.Children = make(map[string]types.YChild)
    accessVfiPseudowires.EntityData.Children["access-vfi-pseudowire"] = types.YChild{"AccessVfiPseudowire", nil}
    for i := range accessVfiPseudowires.AccessVfiPseudowire {
        accessVfiPseudowires.EntityData.Children[types.GetSegmentPath(&accessVfiPseudowires.AccessVfiPseudowire[i])] = types.YChild{"AccessVfiPseudowire", &accessVfiPseudowires.AccessVfiPseudowire[i]}
    }
    accessVfiPseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessVfiPseudowires.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire
// Pseudowire configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // Pseudowire class template name to use for this pseudowire. The type is
    // string with length: 1..32.
    AccessVfiPwClass interface{}

    // Static Mac Address Table.
    AccessVfiPseudowireStaticMacAddresses L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses
}

func (accessVfiPseudowire *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowire.EntityData.YFilter = accessVfiPseudowire.YFilter
    accessVfiPseudowire.EntityData.YangName = "access-vfi-pseudowire"
    accessVfiPseudowire.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowire.EntityData.ParentYangName = "access-vfi-pseudowires"
    accessVfiPseudowire.EntityData.SegmentPath = "access-vfi-pseudowire" + "[neighbor='" + fmt.Sprintf("%v", accessVfiPseudowire.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", accessVfiPseudowire.PseudowireId) + "']"
    accessVfiPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowire.EntityData.Children = make(map[string]types.YChild)
    accessVfiPseudowire.EntityData.Children["access-vfi-pseudowire-static-mac-addresses"] = types.YChild{"AccessVfiPseudowireStaticMacAddresses", &accessVfiPseudowire.AccessVfiPseudowireStaticMacAddresses}
    accessVfiPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    accessVfiPseudowire.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", accessVfiPseudowire.Neighbor}
    accessVfiPseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", accessVfiPseudowire.PseudowireId}
    accessVfiPseudowire.EntityData.Leafs["access-vfi-pw-class"] = types.YLeaf{"AccessVfiPwClass", accessVfiPseudowire.AccessVfiPwClass}
    return &(accessVfiPseudowire.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses
// Static Mac Address Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress.
    AccessVfiPseudowireStaticMacAddress []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress
}

func (accessVfiPseudowireStaticMacAddresses *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowireStaticMacAddresses.EntityData.YFilter = accessVfiPseudowireStaticMacAddresses.YFilter
    accessVfiPseudowireStaticMacAddresses.EntityData.YangName = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowireStaticMacAddresses.EntityData.ParentYangName = "access-vfi-pseudowire"
    accessVfiPseudowireStaticMacAddresses.EntityData.SegmentPath = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowireStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowireStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowireStaticMacAddresses.EntityData.Children = make(map[string]types.YChild)
    accessVfiPseudowireStaticMacAddresses.EntityData.Children["access-vfi-pseudowire-static-mac-address"] = types.YChild{"AccessVfiPseudowireStaticMacAddress", nil}
    for i := range accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress {
        accessVfiPseudowireStaticMacAddresses.EntityData.Children[types.GetSegmentPath(&accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress[i])] = types.YChild{"AccessVfiPseudowireStaticMacAddress", &accessVfiPseudowireStaticMacAddresses.AccessVfiPseudowireStaticMacAddress[i]}
    }
    accessVfiPseudowireStaticMacAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessVfiPseudowireStaticMacAddresses.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress
// Static Mac Address Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}
}

func (accessVfiPseudowireStaticMacAddress *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_AccessVfis_AccessVfi_AccessVfiPseudowires_AccessVfiPseudowire_AccessVfiPseudowireStaticMacAddresses_AccessVfiPseudowireStaticMacAddress) GetEntityData() *types.CommonEntityData {
    accessVfiPseudowireStaticMacAddress.EntityData.YFilter = accessVfiPseudowireStaticMacAddress.YFilter
    accessVfiPseudowireStaticMacAddress.EntityData.YangName = "access-vfi-pseudowire-static-mac-address"
    accessVfiPseudowireStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    accessVfiPseudowireStaticMacAddress.EntityData.ParentYangName = "access-vfi-pseudowire-static-mac-addresses"
    accessVfiPseudowireStaticMacAddress.EntityData.SegmentPath = "access-vfi-pseudowire-static-mac-address" + "[address='" + fmt.Sprintf("%v", accessVfiPseudowireStaticMacAddress.Address) + "']"
    accessVfiPseudowireStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessVfiPseudowireStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessVfiPseudowireStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessVfiPseudowireStaticMacAddress.EntityData.Children = make(map[string]types.YChild)
    accessVfiPseudowireStaticMacAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    accessVfiPseudowireStaticMacAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", accessVfiPseudowireStaticMacAddress.Address}
    return &(accessVfiPseudowireStaticMacAddress.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires
// List of pseudowires
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire.
    BdPseudowire []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire
}

func (bdPseudowires *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires) GetEntityData() *types.CommonEntityData {
    bdPseudowires.EntityData.YFilter = bdPseudowires.YFilter
    bdPseudowires.EntityData.YangName = "bd-pseudowires"
    bdPseudowires.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowires.EntityData.ParentYangName = "bridge-domain"
    bdPseudowires.EntityData.SegmentPath = "bd-pseudowires"
    bdPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowires.EntityData.Children = make(map[string]types.YChild)
    bdPseudowires.EntityData.Children["bd-pseudowire"] = types.YChild{"BdPseudowire", nil}
    for i := range bdPseudowires.BdPseudowire {
        bdPseudowires.EntityData.Children[types.GetSegmentPath(&bdPseudowires.BdPseudowire[i])] = types.YChild{"BdPseudowire", &bdPseudowires.BdPseudowire[i]}
    }
    bdPseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdPseudowires.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire
// Pseudowire configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    PseudowireDai L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai

    // Storm Control.
    BdpwStormControlTypes L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes

    // Attach a DHCP profile.
    PseudowireProfile L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile

    // Static Mac Address Table.
    BdPwStaticMacAddresses L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses

    // IP Source Guard.
    PseudowireIpSourceGuard L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard

    // Bridge-domain Pseudowire MAC configuration commands.
    PseudowireMac L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac

    // Split Horizon.
    BdPwSplitHorizon L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon

    // MPLS static labels.
    BdPwMplsStaticLabels L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels

    // List of pseudowires.
    BridgeDomainBackupPseudowires L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires
}

func (bdPseudowire *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire) GetEntityData() *types.CommonEntityData {
    bdPseudowire.EntityData.YFilter = bdPseudowire.YFilter
    bdPseudowire.EntityData.YangName = "bd-pseudowire"
    bdPseudowire.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowire.EntityData.ParentYangName = "bd-pseudowires"
    bdPseudowire.EntityData.SegmentPath = "bd-pseudowire" + "[neighbor='" + fmt.Sprintf("%v", bdPseudowire.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", bdPseudowire.PseudowireId) + "']"
    bdPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowire.EntityData.Children = make(map[string]types.YChild)
    bdPseudowire.EntityData.Children["pseudowire-dai"] = types.YChild{"PseudowireDai", &bdPseudowire.PseudowireDai}
    bdPseudowire.EntityData.Children["bdpw-storm-control-types"] = types.YChild{"BdpwStormControlTypes", &bdPseudowire.BdpwStormControlTypes}
    bdPseudowire.EntityData.Children["pseudowire-profile"] = types.YChild{"PseudowireProfile", &bdPseudowire.PseudowireProfile}
    bdPseudowire.EntityData.Children["bd-pw-static-mac-addresses"] = types.YChild{"BdPwStaticMacAddresses", &bdPseudowire.BdPwStaticMacAddresses}
    bdPseudowire.EntityData.Children["pseudowire-ip-source-guard"] = types.YChild{"PseudowireIpSourceGuard", &bdPseudowire.PseudowireIpSourceGuard}
    bdPseudowire.EntityData.Children["pseudowire-mac"] = types.YChild{"PseudowireMac", &bdPseudowire.PseudowireMac}
    bdPseudowire.EntityData.Children["bd-pw-split-horizon"] = types.YChild{"BdPwSplitHorizon", &bdPseudowire.BdPwSplitHorizon}
    bdPseudowire.EntityData.Children["bd-pw-mpls-static-labels"] = types.YChild{"BdPwMplsStaticLabels", &bdPseudowire.BdPwMplsStaticLabels}
    bdPseudowire.EntityData.Children["bridge-domain-backup-pseudowires"] = types.YChild{"BridgeDomainBackupPseudowires", &bdPseudowire.BridgeDomainBackupPseudowires}
    bdPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    bdPseudowire.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", bdPseudowire.Neighbor}
    bdPseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", bdPseudowire.PseudowireId}
    bdPseudowire.EntityData.Leafs["pseudowire-mld-snoop"] = types.YLeaf{"PseudowireMldSnoop", bdPseudowire.PseudowireMldSnoop}
    bdPseudowire.EntityData.Leafs["pseudowire-igmp-snoop"] = types.YLeaf{"PseudowireIgmpSnoop", bdPseudowire.PseudowireIgmpSnoop}
    bdPseudowire.EntityData.Leafs["pseudowire-flooding"] = types.YLeaf{"PseudowireFlooding", bdPseudowire.PseudowireFlooding}
    bdPseudowire.EntityData.Leafs["bd-pw-class"] = types.YLeaf{"BdPwClass", bdPseudowire.BdPwClass}
    bdPseudowire.EntityData.Leafs["pseudowire-flooding-unknown-unicast"] = types.YLeaf{"PseudowireFloodingUnknownUnicast", bdPseudowire.PseudowireFloodingUnknownUnicast}
    return &(bdPseudowire.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai
// Access Pseudowire Dynamic ARP Inspection
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable Dynamic ARP Inspection. The type is interface{}.
    Disable interface{}

    // Enable Access Pseudowire Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    PseudowireDaiAddressValidation L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation
}

func (pseudowireDai *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai) GetEntityData() *types.CommonEntityData {
    pseudowireDai.EntityData.YFilter = pseudowireDai.YFilter
    pseudowireDai.EntityData.YangName = "pseudowire-dai"
    pseudowireDai.EntityData.BundleName = "cisco_ios_xr"
    pseudowireDai.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireDai.EntityData.SegmentPath = "pseudowire-dai"
    pseudowireDai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireDai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireDai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireDai.EntityData.Children = make(map[string]types.YChild)
    pseudowireDai.EntityData.Children["pseudowire-dai-address-validation"] = types.YChild{"PseudowireDaiAddressValidation", &pseudowireDai.PseudowireDaiAddressValidation}
    pseudowireDai.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireDai.EntityData.Leafs["logging"] = types.YLeaf{"Logging", pseudowireDai.Logging}
    pseudowireDai.EntityData.Leafs["disable"] = types.YLeaf{"Disable", pseudowireDai.Disable}
    pseudowireDai.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pseudowireDai.Enable}
    return &(pseudowireDai.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation
// Address Validation
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Verification. The type is L2vpnVerification.
    Ipv4Verification interface{}

    // Destination MAC Verification. The type is L2vpnVerification.
    DestinationMacVerification interface{}

    // Source MAC Verification. The type is L2vpnVerification.
    SourceMacVerification interface{}
}

func (pseudowireDaiAddressValidation *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireDai_PseudowireDaiAddressValidation) GetEntityData() *types.CommonEntityData {
    pseudowireDaiAddressValidation.EntityData.YFilter = pseudowireDaiAddressValidation.YFilter
    pseudowireDaiAddressValidation.EntityData.YangName = "pseudowire-dai-address-validation"
    pseudowireDaiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    pseudowireDaiAddressValidation.EntityData.ParentYangName = "pseudowire-dai"
    pseudowireDaiAddressValidation.EntityData.SegmentPath = "pseudowire-dai-address-validation"
    pseudowireDaiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireDaiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireDaiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireDaiAddressValidation.EntityData.Children = make(map[string]types.YChild)
    pseudowireDaiAddressValidation.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireDaiAddressValidation.EntityData.Leafs["ipv4-verification"] = types.YLeaf{"Ipv4Verification", pseudowireDaiAddressValidation.Ipv4Verification}
    pseudowireDaiAddressValidation.EntityData.Leafs["destination-mac-verification"] = types.YLeaf{"DestinationMacVerification", pseudowireDaiAddressValidation.DestinationMacVerification}
    pseudowireDaiAddressValidation.EntityData.Leafs["source-mac-verification"] = types.YLeaf{"SourceMacVerification", pseudowireDaiAddressValidation.SourceMacVerification}
    return &(pseudowireDaiAddressValidation.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes
// Storm Control
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType.
    BdpwStormControlType []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType
}

func (bdpwStormControlTypes *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes) GetEntityData() *types.CommonEntityData {
    bdpwStormControlTypes.EntityData.YFilter = bdpwStormControlTypes.YFilter
    bdpwStormControlTypes.EntityData.YangName = "bdpw-storm-control-types"
    bdpwStormControlTypes.EntityData.BundleName = "cisco_ios_xr"
    bdpwStormControlTypes.EntityData.ParentYangName = "bd-pseudowire"
    bdpwStormControlTypes.EntityData.SegmentPath = "bdpw-storm-control-types"
    bdpwStormControlTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdpwStormControlTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdpwStormControlTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdpwStormControlTypes.EntityData.Children = make(map[string]types.YChild)
    bdpwStormControlTypes.EntityData.Children["bdpw-storm-control-type"] = types.YChild{"BdpwStormControlType", nil}
    for i := range bdpwStormControlTypes.BdpwStormControlType {
        bdpwStormControlTypes.EntityData.Children[types.GetSegmentPath(&bdpwStormControlTypes.BdpwStormControlType[i])] = types.YChild{"BdpwStormControlType", &bdpwStormControlTypes.BdpwStormControlType[i]}
    }
    bdpwStormControlTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdpwStormControlTypes.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType
// Storm Control Type
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit
}

func (bdpwStormControlType *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType) GetEntityData() *types.CommonEntityData {
    bdpwStormControlType.EntityData.YFilter = bdpwStormControlType.YFilter
    bdpwStormControlType.EntityData.YangName = "bdpw-storm-control-type"
    bdpwStormControlType.EntityData.BundleName = "cisco_ios_xr"
    bdpwStormControlType.EntityData.ParentYangName = "bdpw-storm-control-types"
    bdpwStormControlType.EntityData.SegmentPath = "bdpw-storm-control-type" + "[sctype='" + fmt.Sprintf("%v", bdpwStormControlType.Sctype) + "']"
    bdpwStormControlType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdpwStormControlType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdpwStormControlType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdpwStormControlType.EntityData.Children = make(map[string]types.YChild)
    bdpwStormControlType.EntityData.Children["storm-control-unit"] = types.YChild{"StormControlUnit", &bdpwStormControlType.StormControlUnit}
    bdpwStormControlType.EntityData.Leafs = make(map[string]types.YLeaf)
    bdpwStormControlType.EntityData.Leafs["sctype"] = types.YLeaf{"Sctype", bdpwStormControlType.Sctype}
    return &(bdpwStormControlType.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit
// Specify units for Storm Control Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit struct {
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

func (stormControlUnit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdpwStormControlTypes_BdpwStormControlType_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bdpw-storm-control-type"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = make(map[string]types.YChild)
    stormControlUnit.EntityData.Leafs = make(map[string]types.YLeaf)
    stormControlUnit.EntityData.Leafs["kbits-per-sec"] = types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec}
    stormControlUnit.EntityData.Leafs["pkts-per-sec"] = types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec}
    return &(stormControlUnit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile
// Attach a DHCP profile
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (pseudowireProfile *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireProfile) GetEntityData() *types.CommonEntityData {
    pseudowireProfile.EntityData.YFilter = pseudowireProfile.YFilter
    pseudowireProfile.EntityData.YangName = "pseudowire-profile"
    pseudowireProfile.EntityData.BundleName = "cisco_ios_xr"
    pseudowireProfile.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireProfile.EntityData.SegmentPath = "pseudowire-profile"
    pseudowireProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireProfile.EntityData.Children = make(map[string]types.YChild)
    pseudowireProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireProfile.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pseudowireProfile.ProfileId}
    pseudowireProfile.EntityData.Leafs["dhcp-snooping-id"] = types.YLeaf{"DhcpSnoopingId", pseudowireProfile.DhcpSnoopingId}
    return &(pseudowireProfile.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses
// Static Mac Address Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress.
    BdPwStaticMacAddress []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress
}

func (bdPwStaticMacAddresses *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    bdPwStaticMacAddresses.EntityData.YFilter = bdPwStaticMacAddresses.YFilter
    bdPwStaticMacAddresses.EntityData.YangName = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    bdPwStaticMacAddresses.EntityData.ParentYangName = "bd-pseudowire"
    bdPwStaticMacAddresses.EntityData.SegmentPath = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwStaticMacAddresses.EntityData.Children = make(map[string]types.YChild)
    bdPwStaticMacAddresses.EntityData.Children["bd-pw-static-mac-address"] = types.YChild{"BdPwStaticMacAddress", nil}
    for i := range bdPwStaticMacAddresses.BdPwStaticMacAddress {
        bdPwStaticMacAddresses.EntityData.Children[types.GetSegmentPath(&bdPwStaticMacAddresses.BdPwStaticMacAddress[i])] = types.YChild{"BdPwStaticMacAddress", &bdPwStaticMacAddresses.BdPwStaticMacAddress[i]}
    }
    bdPwStaticMacAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdPwStaticMacAddresses.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress
// Static Mac Address Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}
}

func (bdPwStaticMacAddress *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwStaticMacAddresses_BdPwStaticMacAddress) GetEntityData() *types.CommonEntityData {
    bdPwStaticMacAddress.EntityData.YFilter = bdPwStaticMacAddress.YFilter
    bdPwStaticMacAddress.EntityData.YangName = "bd-pw-static-mac-address"
    bdPwStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    bdPwStaticMacAddress.EntityData.ParentYangName = "bd-pw-static-mac-addresses"
    bdPwStaticMacAddress.EntityData.SegmentPath = "bd-pw-static-mac-address" + "[address='" + fmt.Sprintf("%v", bdPwStaticMacAddress.Address) + "']"
    bdPwStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwStaticMacAddress.EntityData.Children = make(map[string]types.YChild)
    bdPwStaticMacAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    bdPwStaticMacAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", bdPwStaticMacAddress.Address}
    return &(bdPwStaticMacAddress.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard
// IP Source Guard
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable Dynamic IP source guard. The type is interface{}.
    Disable interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (pseudowireIpSourceGuard *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireIpSourceGuard) GetEntityData() *types.CommonEntityData {
    pseudowireIpSourceGuard.EntityData.YFilter = pseudowireIpSourceGuard.YFilter
    pseudowireIpSourceGuard.EntityData.YangName = "pseudowire-ip-source-guard"
    pseudowireIpSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    pseudowireIpSourceGuard.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireIpSourceGuard.EntityData.SegmentPath = "pseudowire-ip-source-guard"
    pseudowireIpSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireIpSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireIpSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireIpSourceGuard.EntityData.Children = make(map[string]types.YChild)
    pseudowireIpSourceGuard.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireIpSourceGuard.EntityData.Leafs["logging"] = types.YLeaf{"Logging", pseudowireIpSourceGuard.Logging}
    pseudowireIpSourceGuard.EntityData.Leafs["disable"] = types.YLeaf{"Disable", pseudowireIpSourceGuard.Disable}
    pseudowireIpSourceGuard.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pseudowireIpSourceGuard.Enable}
    return &(pseudowireIpSourceGuard.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac
// Bridge-domain Pseudowire MAC
// configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MAC Flush When Port goes down. The type is PortDownFlush.
    PseudowireMacPortDownFlush interface{}

    // Bridge-domain Pseudowire MAC configuration mode. The type is interface{}.
    Enable interface{}

    // Enable MAC Learning. The type is MacLearn.
    PseudowireMacLearning interface{}

    // MAC Secure.
    PseudowireMacSecure L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure

    // MAC-Aging configuration commands.
    PseudowireMacAging L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging

    // MAC-Limit configuration commands.
    PseudowireMacLimit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit
}

func (pseudowireMac *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac) GetEntityData() *types.CommonEntityData {
    pseudowireMac.EntityData.YFilter = pseudowireMac.YFilter
    pseudowireMac.EntityData.YangName = "pseudowire-mac"
    pseudowireMac.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMac.EntityData.ParentYangName = "bd-pseudowire"
    pseudowireMac.EntityData.SegmentPath = "pseudowire-mac"
    pseudowireMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMac.EntityData.Children = make(map[string]types.YChild)
    pseudowireMac.EntityData.Children["pseudowire-mac-secure"] = types.YChild{"PseudowireMacSecure", &pseudowireMac.PseudowireMacSecure}
    pseudowireMac.EntityData.Children["pseudowire-mac-aging"] = types.YChild{"PseudowireMacAging", &pseudowireMac.PseudowireMacAging}
    pseudowireMac.EntityData.Children["pseudowire-mac-limit"] = types.YChild{"PseudowireMacLimit", &pseudowireMac.PseudowireMacLimit}
    pseudowireMac.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireMac.EntityData.Leafs["pseudowire-mac-port-down-flush"] = types.YLeaf{"PseudowireMacPortDownFlush", pseudowireMac.PseudowireMacPortDownFlush}
    pseudowireMac.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pseudowireMac.Enable}
    pseudowireMac.EntityData.Leafs["pseudowire-mac-learning"] = types.YLeaf{"PseudowireMacLearning", pseudowireMac.PseudowireMacLearning}
    return &(pseudowireMac.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure
// MAC Secure
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure struct {
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

func (pseudowireMacSecure *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacSecure) GetEntityData() *types.CommonEntityData {
    pseudowireMacSecure.EntityData.YFilter = pseudowireMacSecure.YFilter
    pseudowireMacSecure.EntityData.YangName = "pseudowire-mac-secure"
    pseudowireMacSecure.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacSecure.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacSecure.EntityData.SegmentPath = "pseudowire-mac-secure"
    pseudowireMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacSecure.EntityData.Children = make(map[string]types.YChild)
    pseudowireMacSecure.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireMacSecure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", pseudowireMacSecure.Logging}
    pseudowireMacSecure.EntityData.Leafs["disable"] = types.YLeaf{"Disable", pseudowireMacSecure.Disable}
    pseudowireMacSecure.EntityData.Leafs["action"] = types.YLeaf{"Action", pseudowireMacSecure.Action}
    pseudowireMacSecure.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pseudowireMacSecure.Enable}
    return &(pseudowireMacSecure.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging
// MAC-Aging configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address aging type. The type is MacAging.
    PseudowireMacAgingType interface{}

    // MAC Aging Time. The type is interface{} with range: 300..30000.
    PseudowireMacAgingTime interface{}
}

func (pseudowireMacAging *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacAging) GetEntityData() *types.CommonEntityData {
    pseudowireMacAging.EntityData.YFilter = pseudowireMacAging.YFilter
    pseudowireMacAging.EntityData.YangName = "pseudowire-mac-aging"
    pseudowireMacAging.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacAging.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacAging.EntityData.SegmentPath = "pseudowire-mac-aging"
    pseudowireMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacAging.EntityData.Children = make(map[string]types.YChild)
    pseudowireMacAging.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireMacAging.EntityData.Leafs["pseudowire-mac-aging-type"] = types.YLeaf{"PseudowireMacAgingType", pseudowireMacAging.PseudowireMacAgingType}
    pseudowireMacAging.EntityData.Leafs["pseudowire-mac-aging-time"] = types.YLeaf{"PseudowireMacAgingTime", pseudowireMacAging.PseudowireMacAgingTime}
    return &(pseudowireMacAging.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit
// MAC-Limit configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit struct {
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

func (pseudowireMacLimit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_PseudowireMac_PseudowireMacLimit) GetEntityData() *types.CommonEntityData {
    pseudowireMacLimit.EntityData.YFilter = pseudowireMacLimit.YFilter
    pseudowireMacLimit.EntityData.YangName = "pseudowire-mac-limit"
    pseudowireMacLimit.EntityData.BundleName = "cisco_ios_xr"
    pseudowireMacLimit.EntityData.ParentYangName = "pseudowire-mac"
    pseudowireMacLimit.EntityData.SegmentPath = "pseudowire-mac-limit"
    pseudowireMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireMacLimit.EntityData.Children = make(map[string]types.YChild)
    pseudowireMacLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireMacLimit.EntityData.Leafs["pseudowire-mac-limit-action"] = types.YLeaf{"PseudowireMacLimitAction", pseudowireMacLimit.PseudowireMacLimitAction}
    pseudowireMacLimit.EntityData.Leafs["pseudowire-mac-limit-notif"] = types.YLeaf{"PseudowireMacLimitNotif", pseudowireMacLimit.PseudowireMacLimitNotif}
    pseudowireMacLimit.EntityData.Leafs["pseudowire-mac-limit-max"] = types.YLeaf{"PseudowireMacLimitMax", pseudowireMacLimit.PseudowireMacLimitMax}
    return &(pseudowireMacLimit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon
// Split Horizon
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon Group.
    BdPwSplitHorizonGroup L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup
}

func (bdPwSplitHorizon *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon) GetEntityData() *types.CommonEntityData {
    bdPwSplitHorizon.EntityData.YFilter = bdPwSplitHorizon.YFilter
    bdPwSplitHorizon.EntityData.YangName = "bd-pw-split-horizon"
    bdPwSplitHorizon.EntityData.BundleName = "cisco_ios_xr"
    bdPwSplitHorizon.EntityData.ParentYangName = "bd-pseudowire"
    bdPwSplitHorizon.EntityData.SegmentPath = "bd-pw-split-horizon"
    bdPwSplitHorizon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwSplitHorizon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwSplitHorizon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwSplitHorizon.EntityData.Children = make(map[string]types.YChild)
    bdPwSplitHorizon.EntityData.Children["bd-pw-split-horizon-group"] = types.YChild{"BdPwSplitHorizonGroup", &bdPwSplitHorizon.BdPwSplitHorizonGroup}
    bdPwSplitHorizon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdPwSplitHorizon.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup
// Split Horizon Group
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable split horizon group. The type is interface{}.
    Enable interface{}
}

func (bdPwSplitHorizonGroup *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwSplitHorizon_BdPwSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    bdPwSplitHorizonGroup.EntityData.YFilter = bdPwSplitHorizonGroup.YFilter
    bdPwSplitHorizonGroup.EntityData.YangName = "bd-pw-split-horizon-group"
    bdPwSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    bdPwSplitHorizonGroup.EntityData.ParentYangName = "bd-pw-split-horizon"
    bdPwSplitHorizonGroup.EntityData.SegmentPath = "bd-pw-split-horizon-group"
    bdPwSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwSplitHorizonGroup.EntityData.Children = make(map[string]types.YChild)
    bdPwSplitHorizonGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    bdPwSplitHorizonGroup.EntityData.Leafs["enable"] = types.YLeaf{"Enable", bdPwSplitHorizonGroup.Enable}
    return &(bdPwSplitHorizonGroup.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels
// MPLS static labels
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (bdPwMplsStaticLabels *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BdPwMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    bdPwMplsStaticLabels.EntityData.YFilter = bdPwMplsStaticLabels.YFilter
    bdPwMplsStaticLabels.EntityData.YangName = "bd-pw-mpls-static-labels"
    bdPwMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    bdPwMplsStaticLabels.EntityData.ParentYangName = "bd-pseudowire"
    bdPwMplsStaticLabels.EntityData.SegmentPath = "bd-pw-mpls-static-labels"
    bdPwMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPwMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPwMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPwMplsStaticLabels.EntityData.Children = make(map[string]types.YChild)
    bdPwMplsStaticLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    bdPwMplsStaticLabels.EntityData.Leafs["local-static-label"] = types.YLeaf{"LocalStaticLabel", bdPwMplsStaticLabels.LocalStaticLabel}
    bdPwMplsStaticLabels.EntityData.Leafs["remote-static-label"] = types.YLeaf{"RemoteStaticLabel", bdPwMplsStaticLabels.RemoteStaticLabel}
    return &(bdPwMplsStaticLabels.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires
// List of pseudowires
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backup pseudowire configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire.
    BridgeDomainBackupPseudowire []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire
}

func (bridgeDomainBackupPseudowires *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires) GetEntityData() *types.CommonEntityData {
    bridgeDomainBackupPseudowires.EntityData.YFilter = bridgeDomainBackupPseudowires.YFilter
    bridgeDomainBackupPseudowires.EntityData.YangName = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowires.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainBackupPseudowires.EntityData.ParentYangName = "bd-pseudowire"
    bridgeDomainBackupPseudowires.EntityData.SegmentPath = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainBackupPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainBackupPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainBackupPseudowires.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainBackupPseudowires.EntityData.Children["bridge-domain-backup-pseudowire"] = types.YChild{"BridgeDomainBackupPseudowire", nil}
    for i := range bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire {
        bridgeDomainBackupPseudowires.EntityData.Children[types.GetSegmentPath(&bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire[i])] = types.YChild{"BridgeDomainBackupPseudowire", &bridgeDomainBackupPseudowires.BridgeDomainBackupPseudowire[i]}
    }
    bridgeDomainBackupPseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomainBackupPseudowires.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire
// Backup pseudowire configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}

    // This attribute is a key. Pseudowire ID. The type is interface{} with range:
    // 1..4294967295.
    PseudowireId interface{}

    // PW class template name to use for this pseudowire. The type is string with
    // length: 1..32.
    BridgeDomainBackupPwClass interface{}
}

func (bridgeDomainBackupPseudowire *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowires_BdPseudowire_BridgeDomainBackupPseudowires_BridgeDomainBackupPseudowire) GetEntityData() *types.CommonEntityData {
    bridgeDomainBackupPseudowire.EntityData.YFilter = bridgeDomainBackupPseudowire.YFilter
    bridgeDomainBackupPseudowire.EntityData.YangName = "bridge-domain-backup-pseudowire"
    bridgeDomainBackupPseudowire.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainBackupPseudowire.EntityData.ParentYangName = "bridge-domain-backup-pseudowires"
    bridgeDomainBackupPseudowire.EntityData.SegmentPath = "bridge-domain-backup-pseudowire" + "[neighbor='" + fmt.Sprintf("%v", bridgeDomainBackupPseudowire.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", bridgeDomainBackupPseudowire.PseudowireId) + "']"
    bridgeDomainBackupPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainBackupPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainBackupPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainBackupPseudowire.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainBackupPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomainBackupPseudowire.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", bridgeDomainBackupPseudowire.Neighbor}
    bridgeDomainBackupPseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", bridgeDomainBackupPseudowire.PseudowireId}
    bridgeDomainBackupPseudowire.EntityData.Leafs["bridge-domain-backup-pw-class"] = types.YLeaf{"BridgeDomainBackupPwClass", bridgeDomainBackupPseudowire.BridgeDomainBackupPwClass}
    return &(bridgeDomainBackupPseudowire.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis
// Specify the virtual forwarding interface
// name
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Virtual Forwarding Interface. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi.
    Vfi []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi
}

func (vfis *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis) GetEntityData() *types.CommonEntityData {
    vfis.EntityData.YFilter = vfis.YFilter
    vfis.EntityData.YangName = "vfis"
    vfis.EntityData.BundleName = "cisco_ios_xr"
    vfis.EntityData.ParentYangName = "bridge-domain"
    vfis.EntityData.SegmentPath = "vfis"
    vfis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfis.EntityData.Children = make(map[string]types.YChild)
    vfis.EntityData.Children["vfi"] = types.YChild{"Vfi", nil}
    for i := range vfis.Vfi {
        vfis.EntityData.Children[types.GetSegmentPath(&vfis.Vfi[i])] = types.YChild{"Vfi", &vfis.Vfi[i]}
    }
    vfis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vfis.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi
// Name of the Virtual Forwarding Interface
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi struct {
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
    MulticastP2Mp L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp

    // List of pseudowires.
    VfiPseudowires L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires

    // Enable Autodiscovery BGP in this VFI.
    BgpAutoDiscovery L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery
}

func (vfi *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi) GetEntityData() *types.CommonEntityData {
    vfi.EntityData.YFilter = vfi.YFilter
    vfi.EntityData.YangName = "vfi"
    vfi.EntityData.BundleName = "cisco_ios_xr"
    vfi.EntityData.ParentYangName = "vfis"
    vfi.EntityData.SegmentPath = "vfi" + "[name='" + fmt.Sprintf("%v", vfi.Name) + "']"
    vfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfi.EntityData.Children = make(map[string]types.YChild)
    vfi.EntityData.Children["multicast-p2mp"] = types.YChild{"MulticastP2Mp", &vfi.MulticastP2Mp}
    vfi.EntityData.Children["vfi-pseudowires"] = types.YChild{"VfiPseudowires", &vfi.VfiPseudowires}
    vfi.EntityData.Children["bgp-auto-discovery"] = types.YChild{"BgpAutoDiscovery", &vfi.BgpAutoDiscovery}
    vfi.EntityData.Leafs = make(map[string]types.YLeaf)
    vfi.EntityData.Leafs["name"] = types.YLeaf{"Name", vfi.Name}
    vfi.EntityData.Leafs["vfi-shutdown"] = types.YLeaf{"VfiShutdown", vfi.VfiShutdown}
    vfi.EntityData.Leafs["vpnid"] = types.YLeaf{"Vpnid", vfi.Vpnid}
    return &(vfi.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp
// Enable Multicast P2MP in this VFI
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery P2MP. The type is interface{}.
    Enable interface{}

    // Multicast P2MP Transport.
    Transports L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports

    // Multicast P2MP Signaling Type.
    Signalings L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings
}

func (multicastP2Mp *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp) GetEntityData() *types.CommonEntityData {
    multicastP2Mp.EntityData.YFilter = multicastP2Mp.YFilter
    multicastP2Mp.EntityData.YangName = "multicast-p2mp"
    multicastP2Mp.EntityData.BundleName = "cisco_ios_xr"
    multicastP2Mp.EntityData.ParentYangName = "vfi"
    multicastP2Mp.EntityData.SegmentPath = "multicast-p2mp"
    multicastP2Mp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastP2Mp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastP2Mp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastP2Mp.EntityData.Children = make(map[string]types.YChild)
    multicastP2Mp.EntityData.Children["transports"] = types.YChild{"Transports", &multicastP2Mp.Transports}
    multicastP2Mp.EntityData.Children["signalings"] = types.YChild{"Signalings", &multicastP2Mp.Signalings}
    multicastP2Mp.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastP2Mp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", multicastP2Mp.Enable}
    return &(multicastP2Mp.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports
// Multicast P2MP Transport
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multicast P2MP Transport Type. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports_Transport.
    Transport []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports_Transport
}

func (transports *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports) GetEntityData() *types.CommonEntityData {
    transports.EntityData.YFilter = transports.YFilter
    transports.EntityData.YangName = "transports"
    transports.EntityData.BundleName = "cisco_ios_xr"
    transports.EntityData.ParentYangName = "multicast-p2mp"
    transports.EntityData.SegmentPath = "transports"
    transports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transports.EntityData.Children = make(map[string]types.YChild)
    transports.EntityData.Children["transport"] = types.YChild{"Transport", nil}
    for i := range transports.Transport {
        transports.EntityData.Children[types.GetSegmentPath(&transports.Transport[i])] = types.YChild{"Transport", &transports.Transport[i]}
    }
    transports.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(transports.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports_Transport
// Multicast P2MP Transport Type
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Transport Type. The type is string with pattern:
    // b'(RSVP_TE)'.
    TransportName interface{}

    // Multicast P2MP TE Attribute Set Name. The type is string with length:
    // 1..64.
    AttributeSetName interface{}
}

func (transport *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Transports_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xr"
    transport.EntityData.ParentYangName = "transports"
    transport.EntityData.SegmentPath = "transport" + "[transport-name='" + fmt.Sprintf("%v", transport.TransportName) + "']"
    transport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transport.EntityData.Children = make(map[string]types.YChild)
    transport.EntityData.Leafs = make(map[string]types.YLeaf)
    transport.EntityData.Leafs["transport-name"] = types.YLeaf{"TransportName", transport.TransportName}
    transport.EntityData.Leafs["attribute-set-name"] = types.YLeaf{"AttributeSetName", transport.AttributeSetName}
    return &(transport.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings
// Multicast P2MP Signaling Type
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multicast P2MP Signaling Type. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings_Signaling.
    Signaling []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings_Signaling
}

func (signalings *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings) GetEntityData() *types.CommonEntityData {
    signalings.EntityData.YFilter = signalings.YFilter
    signalings.EntityData.YangName = "signalings"
    signalings.EntityData.BundleName = "cisco_ios_xr"
    signalings.EntityData.ParentYangName = "multicast-p2mp"
    signalings.EntityData.SegmentPath = "signalings"
    signalings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalings.EntityData.Children = make(map[string]types.YChild)
    signalings.EntityData.Children["signaling"] = types.YChild{"Signaling", nil}
    for i := range signalings.Signaling {
        signalings.EntityData.Children[types.GetSegmentPath(&signalings.Signaling[i])] = types.YChild{"Signaling", &signalings.Signaling[i]}
    }
    signalings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(signalings.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings_Signaling
// Multicast P2MP Signaling Type
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings_Signaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Signaling Type. The type is string with pattern:
    // b'(BGP)'.
    SignalingName interface{}
}

func (signaling *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_MulticastP2Mp_Signalings_Signaling) GetEntityData() *types.CommonEntityData {
    signaling.EntityData.YFilter = signaling.YFilter
    signaling.EntityData.YangName = "signaling"
    signaling.EntityData.BundleName = "cisco_ios_xr"
    signaling.EntityData.ParentYangName = "signalings"
    signaling.EntityData.SegmentPath = "signaling" + "[signaling-name='" + fmt.Sprintf("%v", signaling.SignalingName) + "']"
    signaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaling.EntityData.Children = make(map[string]types.YChild)
    signaling.EntityData.Leafs = make(map[string]types.YLeaf)
    signaling.EntityData.Leafs["signaling-name"] = types.YLeaf{"SignalingName", signaling.SignalingName}
    return &(signaling.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires
// List of pseudowires
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire.
    VfiPseudowire []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire
}

func (vfiPseudowires *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires) GetEntityData() *types.CommonEntityData {
    vfiPseudowires.EntityData.YFilter = vfiPseudowires.YFilter
    vfiPseudowires.EntityData.YangName = "vfi-pseudowires"
    vfiPseudowires.EntityData.BundleName = "cisco_ios_xr"
    vfiPseudowires.EntityData.ParentYangName = "vfi"
    vfiPseudowires.EntityData.SegmentPath = "vfi-pseudowires"
    vfiPseudowires.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPseudowires.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPseudowires.EntityData.Children = make(map[string]types.YChild)
    vfiPseudowires.EntityData.Children["vfi-pseudowire"] = types.YChild{"VfiPseudowire", nil}
    for i := range vfiPseudowires.VfiPseudowire {
        vfiPseudowires.EntityData.Children[types.GetSegmentPath(&vfiPseudowires.VfiPseudowire[i])] = types.YChild{"VfiPseudowire", &vfiPseudowires.VfiPseudowire[i]}
    }
    vfiPseudowires.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vfiPseudowires.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire
// Pseudowire configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    VfiPwDhcpSnoop L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop

    // MPLS static labels.
    VfiPwMplsStaticLabels L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels

    // Static Mac Address Table.
    PseudowireStaticMacAddresses L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses
}

func (vfiPseudowire *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire) GetEntityData() *types.CommonEntityData {
    vfiPseudowire.EntityData.YFilter = vfiPseudowire.YFilter
    vfiPseudowire.EntityData.YangName = "vfi-pseudowire"
    vfiPseudowire.EntityData.BundleName = "cisco_ios_xr"
    vfiPseudowire.EntityData.ParentYangName = "vfi-pseudowires"
    vfiPseudowire.EntityData.SegmentPath = "vfi-pseudowire" + "[neighbor='" + fmt.Sprintf("%v", vfiPseudowire.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", vfiPseudowire.PseudowireId) + "']"
    vfiPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPseudowire.EntityData.Children = make(map[string]types.YChild)
    vfiPseudowire.EntityData.Children["vfi-pw-dhcp-snoop"] = types.YChild{"VfiPwDhcpSnoop", &vfiPseudowire.VfiPwDhcpSnoop}
    vfiPseudowire.EntityData.Children["vfi-pw-mpls-static-labels"] = types.YChild{"VfiPwMplsStaticLabels", &vfiPseudowire.VfiPwMplsStaticLabels}
    vfiPseudowire.EntityData.Children["pseudowire-static-mac-addresses"] = types.YChild{"PseudowireStaticMacAddresses", &vfiPseudowire.PseudowireStaticMacAddresses}
    vfiPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    vfiPseudowire.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", vfiPseudowire.Neighbor}
    vfiPseudowire.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", vfiPseudowire.PseudowireId}
    vfiPseudowire.EntityData.Leafs["vfi-pw-class"] = types.YLeaf{"VfiPwClass", vfiPseudowire.VfiPwClass}
    vfiPseudowire.EntityData.Leafs["vfi-pw-igmp-snoop"] = types.YLeaf{"VfiPwIgmpSnoop", vfiPseudowire.VfiPwIgmpSnoop}
    vfiPseudowire.EntityData.Leafs["vfi-pw-mld-snoop"] = types.YLeaf{"VfiPwMldSnoop", vfiPseudowire.VfiPwMldSnoop}
    return &(vfiPseudowire.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop
// Attach a DHCP Snooping profile
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (vfiPwDhcpSnoop *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwDhcpSnoop) GetEntityData() *types.CommonEntityData {
    vfiPwDhcpSnoop.EntityData.YFilter = vfiPwDhcpSnoop.YFilter
    vfiPwDhcpSnoop.EntityData.YangName = "vfi-pw-dhcp-snoop"
    vfiPwDhcpSnoop.EntityData.BundleName = "cisco_ios_xr"
    vfiPwDhcpSnoop.EntityData.ParentYangName = "vfi-pseudowire"
    vfiPwDhcpSnoop.EntityData.SegmentPath = "vfi-pw-dhcp-snoop"
    vfiPwDhcpSnoop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPwDhcpSnoop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPwDhcpSnoop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPwDhcpSnoop.EntityData.Children = make(map[string]types.YChild)
    vfiPwDhcpSnoop.EntityData.Leafs = make(map[string]types.YLeaf)
    vfiPwDhcpSnoop.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", vfiPwDhcpSnoop.ProfileId}
    vfiPwDhcpSnoop.EntityData.Leafs["dhcp-snooping-id"] = types.YLeaf{"DhcpSnoopingId", vfiPwDhcpSnoop.DhcpSnoopingId}
    return &(vfiPwDhcpSnoop.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels
// MPLS static labels
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire local static label. The type is interface{} with range:
    // 16..1048575.
    LocalStaticLabel interface{}

    // Pseudowire remote static label. The type is interface{} with range:
    // 16..1048575.
    RemoteStaticLabel interface{}
}

func (vfiPwMplsStaticLabels *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_VfiPwMplsStaticLabels) GetEntityData() *types.CommonEntityData {
    vfiPwMplsStaticLabels.EntityData.YFilter = vfiPwMplsStaticLabels.YFilter
    vfiPwMplsStaticLabels.EntityData.YangName = "vfi-pw-mpls-static-labels"
    vfiPwMplsStaticLabels.EntityData.BundleName = "cisco_ios_xr"
    vfiPwMplsStaticLabels.EntityData.ParentYangName = "vfi-pseudowire"
    vfiPwMplsStaticLabels.EntityData.SegmentPath = "vfi-pw-mpls-static-labels"
    vfiPwMplsStaticLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vfiPwMplsStaticLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vfiPwMplsStaticLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vfiPwMplsStaticLabels.EntityData.Children = make(map[string]types.YChild)
    vfiPwMplsStaticLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    vfiPwMplsStaticLabels.EntityData.Leafs["local-static-label"] = types.YLeaf{"LocalStaticLabel", vfiPwMplsStaticLabels.LocalStaticLabel}
    vfiPwMplsStaticLabels.EntityData.Leafs["remote-static-label"] = types.YLeaf{"RemoteStaticLabel", vfiPwMplsStaticLabels.RemoteStaticLabel}
    return &(vfiPwMplsStaticLabels.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses
// Static Mac Address Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress.
    PseudowireStaticMacAddress []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress
}

func (pseudowireStaticMacAddresses *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses) GetEntityData() *types.CommonEntityData {
    pseudowireStaticMacAddresses.EntityData.YFilter = pseudowireStaticMacAddresses.YFilter
    pseudowireStaticMacAddresses.EntityData.YangName = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    pseudowireStaticMacAddresses.EntityData.ParentYangName = "vfi-pseudowire"
    pseudowireStaticMacAddresses.EntityData.SegmentPath = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireStaticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireStaticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireStaticMacAddresses.EntityData.Children = make(map[string]types.YChild)
    pseudowireStaticMacAddresses.EntityData.Children["pseudowire-static-mac-address"] = types.YChild{"PseudowireStaticMacAddress", nil}
    for i := range pseudowireStaticMacAddresses.PseudowireStaticMacAddress {
        pseudowireStaticMacAddresses.EntityData.Children[types.GetSegmentPath(&pseudowireStaticMacAddresses.PseudowireStaticMacAddress[i])] = types.YChild{"PseudowireStaticMacAddress", &pseudowireStaticMacAddresses.PseudowireStaticMacAddress[i]}
    }
    pseudowireStaticMacAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pseudowireStaticMacAddresses.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress
// Static Mac Address Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}
}

func (pseudowireStaticMacAddress *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_VfiPseudowires_VfiPseudowire_PseudowireStaticMacAddresses_PseudowireStaticMacAddress) GetEntityData() *types.CommonEntityData {
    pseudowireStaticMacAddress.EntityData.YFilter = pseudowireStaticMacAddress.YFilter
    pseudowireStaticMacAddress.EntityData.YangName = "pseudowire-static-mac-address"
    pseudowireStaticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    pseudowireStaticMacAddress.EntityData.ParentYangName = "pseudowire-static-mac-addresses"
    pseudowireStaticMacAddress.EntityData.SegmentPath = "pseudowire-static-mac-address" + "[address='" + fmt.Sprintf("%v", pseudowireStaticMacAddress.Address) + "']"
    pseudowireStaticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireStaticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireStaticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireStaticMacAddress.EntityData.Children = make(map[string]types.YChild)
    pseudowireStaticMacAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireStaticMacAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", pseudowireStaticMacAddress.Address}
    return &(pseudowireStaticMacAddress.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery
// Enable Autodiscovery BGP in this VFI
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery struct {
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
    LdpSignalingProtocol L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol

    // Route policy.
    BgpRoutePolicy L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy

    // Route Distinguisher.
    RouteDistinguisher L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher

    // Enable Signaling Protocol BGP in this VFI.
    BgpSignalingProtocol L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol

    // Route Target.
    RouteTargets L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets
}

func (bgpAutoDiscovery *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    bgpAutoDiscovery.EntityData.YFilter = bgpAutoDiscovery.YFilter
    bgpAutoDiscovery.EntityData.YangName = "bgp-auto-discovery"
    bgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    bgpAutoDiscovery.EntityData.ParentYangName = "vfi"
    bgpAutoDiscovery.EntityData.SegmentPath = "bgp-auto-discovery"
    bgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    bgpAutoDiscovery.EntityData.Children["ldp-signaling-protocol"] = types.YChild{"LdpSignalingProtocol", &bgpAutoDiscovery.LdpSignalingProtocol}
    bgpAutoDiscovery.EntityData.Children["bgp-route-policy"] = types.YChild{"BgpRoutePolicy", &bgpAutoDiscovery.BgpRoutePolicy}
    bgpAutoDiscovery.EntityData.Children["route-distinguisher"] = types.YChild{"RouteDistinguisher", &bgpAutoDiscovery.RouteDistinguisher}
    bgpAutoDiscovery.EntityData.Children["bgp-signaling-protocol"] = types.YChild{"BgpSignalingProtocol", &bgpAutoDiscovery.BgpSignalingProtocol}
    bgpAutoDiscovery.EntityData.Children["route-targets"] = types.YChild{"RouteTargets", &bgpAutoDiscovery.RouteTargets}
    bgpAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpAutoDiscovery.EntityData.Leafs["table-policy"] = types.YLeaf{"TablePolicy", bgpAutoDiscovery.TablePolicy}
    bgpAutoDiscovery.EntityData.Leafs["ad-control-word"] = types.YLeaf{"AdControlWord", bgpAutoDiscovery.AdControlWord}
    bgpAutoDiscovery.EntityData.Leafs["enable"] = types.YLeaf{"Enable", bgpAutoDiscovery.Enable}
    return &(bgpAutoDiscovery.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol
// Signaling Protocol LDP in this VFI
// configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable LDP as Signaling Protocol .Deletion of this object also causes
    // deletion of all objects under LDPSignalingProtocol. The type is
    // interface{}.
    Enable interface{}

    // VPLS ID.
    VplsId L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance
}

func (ldpSignalingProtocol *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    ldpSignalingProtocol.EntityData.YFilter = ldpSignalingProtocol.YFilter
    ldpSignalingProtocol.EntityData.YangName = "ldp-signaling-protocol"
    ldpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    ldpSignalingProtocol.EntityData.ParentYangName = "bgp-auto-discovery"
    ldpSignalingProtocol.EntityData.SegmentPath = "ldp-signaling-protocol"
    ldpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpSignalingProtocol.EntityData.Children = make(map[string]types.YChild)
    ldpSignalingProtocol.EntityData.Children["vpls-id"] = types.YChild{"VplsId", &ldpSignalingProtocol.VplsId}
    ldpSignalingProtocol.EntityData.Children["flow-label-load-balance"] = types.YChild{"FlowLabelLoadBalance", &ldpSignalingProtocol.FlowLabelLoadBalance}
    ldpSignalingProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    ldpSignalingProtocol.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ldpSignalingProtocol.Enable}
    return &(ldpSignalingProtocol.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId
// VPLS ID
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPLS-ID Type. The type is LdpVplsId.
    Type_ interface{}

    // Two byte AS number. The type is interface{} with range: 1..65535.
    As interface{}

    // AS index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Address index. The type is interface{} with range: 0..32767.
    AddressIndex interface{}
}

func (vplsId *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_VplsId) GetEntityData() *types.CommonEntityData {
    vplsId.EntityData.YFilter = vplsId.YFilter
    vplsId.EntityData.YangName = "vpls-id"
    vplsId.EntityData.BundleName = "cisco_ios_xr"
    vplsId.EntityData.ParentYangName = "ldp-signaling-protocol"
    vplsId.EntityData.SegmentPath = "vpls-id"
    vplsId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vplsId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vplsId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vplsId.EntityData.Children = make(map[string]types.YChild)
    vplsId.EntityData.Leafs = make(map[string]types.YLeaf)
    vplsId.EntityData.Leafs["type"] = types.YLeaf{"Type_", vplsId.Type_}
    vplsId.EntityData.Leafs["as"] = types.YLeaf{"As", vplsId.As}
    vplsId.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", vplsId.AsIndex}
    vplsId.EntityData.Leafs["address"] = types.YLeaf{"Address", vplsId.Address}
    vplsId.EntityData.Leafs["address-index"] = types.YLeaf{"AddressIndex", vplsId.AddressIndex}
    return &(vplsId.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_LdpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "ldp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = make(map[string]types.YChild)
    flowLabelLoadBalance.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabelLoadBalance.EntityData.Leafs["flow-label"] = types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel}
    flowLabelLoadBalance.EntityData.Leafs["static"] = types.YLeaf{"Static", flowLabelLoadBalance.Static}
    return &(flowLabelLoadBalance.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy
// Route policy
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Export route policy. The type is string.
    Export interface{}
}

func (bgpRoutePolicy *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpRoutePolicy) GetEntityData() *types.CommonEntityData {
    bgpRoutePolicy.EntityData.YFilter = bgpRoutePolicy.YFilter
    bgpRoutePolicy.EntityData.YangName = "bgp-route-policy"
    bgpRoutePolicy.EntityData.BundleName = "cisco_ios_xr"
    bgpRoutePolicy.EntityData.ParentYangName = "bgp-auto-discovery"
    bgpRoutePolicy.EntityData.SegmentPath = "bgp-route-policy"
    bgpRoutePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpRoutePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpRoutePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpRoutePolicy.EntityData.Children = make(map[string]types.YChild)
    bgpRoutePolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRoutePolicy.EntityData.Leafs["export"] = types.YLeaf{"Export", bgpRoutePolicy.Export}
    return &(bgpRoutePolicy.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher
// Route Distinguisher
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type_ interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (routeDistinguisher *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteDistinguisher) GetEntityData() *types.CommonEntityData {
    routeDistinguisher.EntityData.YFilter = routeDistinguisher.YFilter
    routeDistinguisher.EntityData.YangName = "route-distinguisher"
    routeDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    routeDistinguisher.EntityData.ParentYangName = "bgp-auto-discovery"
    routeDistinguisher.EntityData.SegmentPath = "route-distinguisher"
    routeDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeDistinguisher.EntityData.Children = make(map[string]types.YChild)
    routeDistinguisher.EntityData.Leafs = make(map[string]types.YLeaf)
    routeDistinguisher.EntityData.Leafs["type"] = types.YLeaf{"Type_", routeDistinguisher.Type_}
    routeDistinguisher.EntityData.Leafs["as"] = types.YLeaf{"As", routeDistinguisher.As}
    routeDistinguisher.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", routeDistinguisher.AsIndex}
    routeDistinguisher.EntityData.Leafs["address"] = types.YLeaf{"Address", routeDistinguisher.Address}
    routeDistinguisher.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", routeDistinguisher.AddrIndex}
    return &(routeDistinguisher.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol
// Enable Signaling Protocol BGP in this
// VFI
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol struct {
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
    FlowLabelLoadBalance L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance
}

func (bgpSignalingProtocol *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol) GetEntityData() *types.CommonEntityData {
    bgpSignalingProtocol.EntityData.YFilter = bgpSignalingProtocol.YFilter
    bgpSignalingProtocol.EntityData.YangName = "bgp-signaling-protocol"
    bgpSignalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    bgpSignalingProtocol.EntityData.ParentYangName = "bgp-auto-discovery"
    bgpSignalingProtocol.EntityData.SegmentPath = "bgp-signaling-protocol"
    bgpSignalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpSignalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpSignalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpSignalingProtocol.EntityData.Children = make(map[string]types.YChild)
    bgpSignalingProtocol.EntityData.Children["flow-label-load-balance"] = types.YChild{"FlowLabelLoadBalance", &bgpSignalingProtocol.FlowLabelLoadBalance}
    bgpSignalingProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpSignalingProtocol.EntityData.Leafs["ve-range"] = types.YLeaf{"VeRange", bgpSignalingProtocol.VeRange}
    bgpSignalingProtocol.EntityData.Leafs["veid"] = types.YLeaf{"Veid", bgpSignalingProtocol.Veid}
    bgpSignalingProtocol.EntityData.Leafs["enable"] = types.YLeaf{"Enable", bgpSignalingProtocol.Enable}
    return &(bgpSignalingProtocol.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_BgpSignalingProtocol_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "bgp-signaling-protocol"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = make(map[string]types.YChild)
    flowLabelLoadBalance.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabelLoadBalance.EntityData.Leafs["flow-label"] = types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel}
    flowLabelLoadBalance.EntityData.Leafs["static"] = types.YLeaf{"Static", flowLabelLoadBalance.Static}
    return &(flowLabelLoadBalance.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets
// Route Target
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget.
    RouteTarget []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget
}

func (routeTargets *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "bgp-auto-discovery"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = make(map[string]types.YChild)
    routeTargets.EntityData.Children["route-target"] = types.YChild{"RouteTarget", nil}
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children[types.GetSegmentPath(&routeTargets.RouteTarget[i])] = types.YChild{"RouteTarget", &routeTargets.RouteTarget[i]}
    }
    routeTargets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routeTargets.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget
// Name of the Route Target
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // two byte as or four byte as. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs.
    TwoByteAsOrFourByteAs []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs

    // ipv4 address. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address.
    Ipv4Address []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address
}

func (routeTarget *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target" + "[role='" + fmt.Sprintf("%v", routeTarget.Role) + "']" + "[format='" + fmt.Sprintf("%v", routeTarget.Format) + "']"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = make(map[string]types.YChild)
    routeTarget.EntityData.Children["two-byte-as-or-four-byte-as"] = types.YChild{"TwoByteAsOrFourByteAs", nil}
    for i := range routeTarget.TwoByteAsOrFourByteAs {
        routeTarget.EntityData.Children[types.GetSegmentPath(&routeTarget.TwoByteAsOrFourByteAs[i])] = types.YChild{"TwoByteAsOrFourByteAs", &routeTarget.TwoByteAsOrFourByteAs[i]}
    }
    routeTarget.EntityData.Children["ipv4-address"] = types.YChild{"Ipv4Address", nil}
    for i := range routeTarget.Ipv4Address {
        routeTarget.EntityData.Children[types.GetSegmentPath(&routeTarget.Ipv4Address[i])] = types.YChild{"Ipv4Address", &routeTarget.Ipv4Address[i]}
    }
    routeTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget.EntityData.Leafs["role"] = types.YLeaf{"Role", routeTarget.Role}
    routeTarget.EntityData.Leafs["format"] = types.YLeaf{"Format", routeTarget.Format}
    return &(routeTarget.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs
// two byte as or four byte as
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Two byte or 4 byte AS number. The type is
    // interface{} with range: 1..4294967295.
    As interface{}

    // This attribute is a key. AS:nn (hex or decimal format). The type is
    // interface{} with range: 0..4294967295.
    AsIndex interface{}
}

func (twoByteAsOrFourByteAs *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_TwoByteAsOrFourByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAsOrFourByteAs.EntityData.YFilter = twoByteAsOrFourByteAs.YFilter
    twoByteAsOrFourByteAs.EntityData.YangName = "two-byte-as-or-four-byte-as"
    twoByteAsOrFourByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAsOrFourByteAs.EntityData.ParentYangName = "route-target"
    twoByteAsOrFourByteAs.EntityData.SegmentPath = "two-byte-as-or-four-byte-as" + "[as='" + fmt.Sprintf("%v", twoByteAsOrFourByteAs.As) + "']" + "[as-index='" + fmt.Sprintf("%v", twoByteAsOrFourByteAs.AsIndex) + "']"
    twoByteAsOrFourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAsOrFourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAsOrFourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAsOrFourByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAsOrFourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAsOrFourByteAs.EntityData.Leafs["as"] = types.YLeaf{"As", twoByteAsOrFourByteAs.As}
    twoByteAsOrFourByteAs.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", twoByteAsOrFourByteAs.AsIndex}
    return &(twoByteAsOrFourByteAs.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address
// ipv4 address
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}
}

func (ipv4Address *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Vfis_Vfi_BgpAutoDiscovery_RouteTargets_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + "[address='" + fmt.Sprintf("%v", ipv4Address.Address) + "']" + "[addr-index='" + fmt.Sprintf("%v", ipv4Address.AddrIndex) + "']"
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = make(map[string]types.YChild)
    ipv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4Address.Address}
    ipv4Address.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", ipv4Address.AddrIndex}
    return &(ipv4Address.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits
// Attachment Circuit table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Attachment Circuit. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit.
    BdAttachmentCircuit []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit
}

func (bdAttachmentCircuits *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    bdAttachmentCircuits.EntityData.YFilter = bdAttachmentCircuits.YFilter
    bdAttachmentCircuits.EntityData.YangName = "bd-attachment-circuits"
    bdAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    bdAttachmentCircuits.EntityData.ParentYangName = "bridge-domain"
    bdAttachmentCircuits.EntityData.SegmentPath = "bd-attachment-circuits"
    bdAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdAttachmentCircuits.EntityData.Children = make(map[string]types.YChild)
    bdAttachmentCircuits.EntityData.Children["bd-attachment-circuit"] = types.YChild{"BdAttachmentCircuit", nil}
    for i := range bdAttachmentCircuits.BdAttachmentCircuit {
        bdAttachmentCircuits.EntityData.Children[types.GetSegmentPath(&bdAttachmentCircuits.BdAttachmentCircuit[i])] = types.YChild{"BdAttachmentCircuit", &bdAttachmentCircuits.BdAttachmentCircuit[i]}
    }
    bdAttachmentCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdAttachmentCircuits.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit
// Name of the Attachment Circuit
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Attachment Circuit. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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
    InterfaceIpSourceGuard L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard

    // L2 Interface Dynamic ARP Inspection.
    InterfaceDai L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai

    // Attach a DHCP profile.
    InterfaceProfile L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile

    // Storm Control.
    BdacStormControlTypes L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes

    // Split Horizon.
    SplitHorizon L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon

    // Static Mac Address Table.
    StaticMacAddresses L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses

    // MAC configuration commands.
    InterfaceMac L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac
}

func (bdAttachmentCircuit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    bdAttachmentCircuit.EntityData.YFilter = bdAttachmentCircuit.YFilter
    bdAttachmentCircuit.EntityData.YangName = "bd-attachment-circuit"
    bdAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    bdAttachmentCircuit.EntityData.ParentYangName = "bd-attachment-circuits"
    bdAttachmentCircuit.EntityData.SegmentPath = "bd-attachment-circuit" + "[name='" + fmt.Sprintf("%v", bdAttachmentCircuit.Name) + "']"
    bdAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdAttachmentCircuit.EntityData.Children = make(map[string]types.YChild)
    bdAttachmentCircuit.EntityData.Children["interface-ip-source-guard"] = types.YChild{"InterfaceIpSourceGuard", &bdAttachmentCircuit.InterfaceIpSourceGuard}
    bdAttachmentCircuit.EntityData.Children["interface-dai"] = types.YChild{"InterfaceDai", &bdAttachmentCircuit.InterfaceDai}
    bdAttachmentCircuit.EntityData.Children["interface-profile"] = types.YChild{"InterfaceProfile", &bdAttachmentCircuit.InterfaceProfile}
    bdAttachmentCircuit.EntityData.Children["bdac-storm-control-types"] = types.YChild{"BdacStormControlTypes", &bdAttachmentCircuit.BdacStormControlTypes}
    bdAttachmentCircuit.EntityData.Children["split-horizon"] = types.YChild{"SplitHorizon", &bdAttachmentCircuit.SplitHorizon}
    bdAttachmentCircuit.EntityData.Children["static-mac-addresses"] = types.YChild{"StaticMacAddresses", &bdAttachmentCircuit.StaticMacAddresses}
    bdAttachmentCircuit.EntityData.Children["interface-mac"] = types.YChild{"InterfaceMac", &bdAttachmentCircuit.InterfaceMac}
    bdAttachmentCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    bdAttachmentCircuit.EntityData.Leafs["name"] = types.YLeaf{"Name", bdAttachmentCircuit.Name}
    bdAttachmentCircuit.EntityData.Leafs["interface-flooding"] = types.YLeaf{"InterfaceFlooding", bdAttachmentCircuit.InterfaceFlooding}
    bdAttachmentCircuit.EntityData.Leafs["interface-igmp-snoop"] = types.YLeaf{"InterfaceIgmpSnoop", bdAttachmentCircuit.InterfaceIgmpSnoop}
    bdAttachmentCircuit.EntityData.Leafs["interface-flooding-unknown-unicast"] = types.YLeaf{"InterfaceFloodingUnknownUnicast", bdAttachmentCircuit.InterfaceFloodingUnknownUnicast}
    bdAttachmentCircuit.EntityData.Leafs["interface-mld-snoop"] = types.YLeaf{"InterfaceMldSnoop", bdAttachmentCircuit.InterfaceMldSnoop}
    return &(bdAttachmentCircuit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard
// IP Source Guard
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface Dynamic IP source guard. The type is interface{}.
    Disable interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (interfaceIpSourceGuard *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceIpSourceGuard) GetEntityData() *types.CommonEntityData {
    interfaceIpSourceGuard.EntityData.YFilter = interfaceIpSourceGuard.YFilter
    interfaceIpSourceGuard.EntityData.YangName = "interface-ip-source-guard"
    interfaceIpSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    interfaceIpSourceGuard.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceIpSourceGuard.EntityData.SegmentPath = "interface-ip-source-guard"
    interfaceIpSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIpSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIpSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIpSourceGuard.EntityData.Children = make(map[string]types.YChild)
    interfaceIpSourceGuard.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceIpSourceGuard.EntityData.Leafs["logging"] = types.YLeaf{"Logging", interfaceIpSourceGuard.Logging}
    interfaceIpSourceGuard.EntityData.Leafs["disable"] = types.YLeaf{"Disable", interfaceIpSourceGuard.Disable}
    interfaceIpSourceGuard.EntityData.Leafs["enable"] = types.YLeaf{"Enable", interfaceIpSourceGuard.Enable}
    return &(interfaceIpSourceGuard.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai
// L2 Interface Dynamic ARP Inspection
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging Type. The type is L2vpnLogging.
    Logging interface{}

    // Disable L2 Interface Dynamic ARP Inspection. The type is interface{}.
    Disable interface{}

    // Enable L2 Interface Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    InterfaceDaiAddressValidation L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation
}

func (interfaceDai *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai) GetEntityData() *types.CommonEntityData {
    interfaceDai.EntityData.YFilter = interfaceDai.YFilter
    interfaceDai.EntityData.YangName = "interface-dai"
    interfaceDai.EntityData.BundleName = "cisco_ios_xr"
    interfaceDai.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceDai.EntityData.SegmentPath = "interface-dai"
    interfaceDai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDai.EntityData.Children = make(map[string]types.YChild)
    interfaceDai.EntityData.Children["interface-dai-address-validation"] = types.YChild{"InterfaceDaiAddressValidation", &interfaceDai.InterfaceDaiAddressValidation}
    interfaceDai.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDai.EntityData.Leafs["logging"] = types.YLeaf{"Logging", interfaceDai.Logging}
    interfaceDai.EntityData.Leafs["disable"] = types.YLeaf{"Disable", interfaceDai.Disable}
    interfaceDai.EntityData.Leafs["enable"] = types.YLeaf{"Enable", interfaceDai.Enable}
    return &(interfaceDai.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation
// Address Validation
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation struct {
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

func (interfaceDaiAddressValidation *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceDai_InterfaceDaiAddressValidation) GetEntityData() *types.CommonEntityData {
    interfaceDaiAddressValidation.EntityData.YFilter = interfaceDaiAddressValidation.YFilter
    interfaceDaiAddressValidation.EntityData.YangName = "interface-dai-address-validation"
    interfaceDaiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    interfaceDaiAddressValidation.EntityData.ParentYangName = "interface-dai"
    interfaceDaiAddressValidation.EntityData.SegmentPath = "interface-dai-address-validation"
    interfaceDaiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDaiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDaiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDaiAddressValidation.EntityData.Children = make(map[string]types.YChild)
    interfaceDaiAddressValidation.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDaiAddressValidation.EntityData.Leafs["ipv4-verification"] = types.YLeaf{"Ipv4Verification", interfaceDaiAddressValidation.Ipv4Verification}
    interfaceDaiAddressValidation.EntityData.Leafs["destination-mac-verification"] = types.YLeaf{"DestinationMacVerification", interfaceDaiAddressValidation.DestinationMacVerification}
    interfaceDaiAddressValidation.EntityData.Leafs["source-mac-verification"] = types.YLeaf{"SourceMacVerification", interfaceDaiAddressValidation.SourceMacVerification}
    interfaceDaiAddressValidation.EntityData.Leafs["enable"] = types.YLeaf{"Enable", interfaceDaiAddressValidation.Enable}
    return &(interfaceDaiAddressValidation.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile
// Attach a DHCP profile
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the snooping profile. The type is InterfaceProfile.
    ProfileId interface{}

    // Disable DHCP snooping. The type is string.
    DhcpSnoopingId interface{}
}

func (interfaceProfile *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceProfile) GetEntityData() *types.CommonEntityData {
    interfaceProfile.EntityData.YFilter = interfaceProfile.YFilter
    interfaceProfile.EntityData.YangName = "interface-profile"
    interfaceProfile.EntityData.BundleName = "cisco_ios_xr"
    interfaceProfile.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceProfile.EntityData.SegmentPath = "interface-profile"
    interfaceProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProfile.EntityData.Children = make(map[string]types.YChild)
    interfaceProfile.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceProfile.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", interfaceProfile.ProfileId}
    interfaceProfile.EntityData.Leafs["dhcp-snooping-id"] = types.YLeaf{"DhcpSnoopingId", interfaceProfile.DhcpSnoopingId}
    return &(interfaceProfile.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes
// Storm Control
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Storm Control Type. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType.
    BdacStormControlType []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType
}

func (bdacStormControlTypes *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes) GetEntityData() *types.CommonEntityData {
    bdacStormControlTypes.EntityData.YFilter = bdacStormControlTypes.YFilter
    bdacStormControlTypes.EntityData.YangName = "bdac-storm-control-types"
    bdacStormControlTypes.EntityData.BundleName = "cisco_ios_xr"
    bdacStormControlTypes.EntityData.ParentYangName = "bd-attachment-circuit"
    bdacStormControlTypes.EntityData.SegmentPath = "bdac-storm-control-types"
    bdacStormControlTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdacStormControlTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdacStormControlTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdacStormControlTypes.EntityData.Children = make(map[string]types.YChild)
    bdacStormControlTypes.EntityData.Children["bdac-storm-control-type"] = types.YChild{"BdacStormControlType", nil}
    for i := range bdacStormControlTypes.BdacStormControlType {
        bdacStormControlTypes.EntityData.Children[types.GetSegmentPath(&bdacStormControlTypes.BdacStormControlType[i])] = types.YChild{"BdacStormControlType", &bdacStormControlTypes.BdacStormControlType[i]}
    }
    bdacStormControlTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdacStormControlTypes.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType
// Storm Control Type
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Storm Control Type. The type is StormControl.
    Sctype interface{}

    // Specify units for Storm Control Configuration.
    StormControlUnit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit
}

func (bdacStormControlType *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType) GetEntityData() *types.CommonEntityData {
    bdacStormControlType.EntityData.YFilter = bdacStormControlType.YFilter
    bdacStormControlType.EntityData.YangName = "bdac-storm-control-type"
    bdacStormControlType.EntityData.BundleName = "cisco_ios_xr"
    bdacStormControlType.EntityData.ParentYangName = "bdac-storm-control-types"
    bdacStormControlType.EntityData.SegmentPath = "bdac-storm-control-type" + "[sctype='" + fmt.Sprintf("%v", bdacStormControlType.Sctype) + "']"
    bdacStormControlType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdacStormControlType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdacStormControlType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdacStormControlType.EntityData.Children = make(map[string]types.YChild)
    bdacStormControlType.EntityData.Children["storm-control-unit"] = types.YChild{"StormControlUnit", &bdacStormControlType.StormControlUnit}
    bdacStormControlType.EntityData.Leafs = make(map[string]types.YLeaf)
    bdacStormControlType.EntityData.Leafs["sctype"] = types.YLeaf{"Sctype", bdacStormControlType.Sctype}
    return &(bdacStormControlType.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit
// Specify units for Storm Control Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit struct {
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

func (stormControlUnit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_BdacStormControlTypes_BdacStormControlType_StormControlUnit) GetEntityData() *types.CommonEntityData {
    stormControlUnit.EntityData.YFilter = stormControlUnit.YFilter
    stormControlUnit.EntityData.YangName = "storm-control-unit"
    stormControlUnit.EntityData.BundleName = "cisco_ios_xr"
    stormControlUnit.EntityData.ParentYangName = "bdac-storm-control-type"
    stormControlUnit.EntityData.SegmentPath = "storm-control-unit"
    stormControlUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stormControlUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stormControlUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stormControlUnit.EntityData.Children = make(map[string]types.YChild)
    stormControlUnit.EntityData.Leafs = make(map[string]types.YLeaf)
    stormControlUnit.EntityData.Leafs["kbits-per-sec"] = types.YLeaf{"KbitsPerSec", stormControlUnit.KbitsPerSec}
    stormControlUnit.EntityData.Leafs["pkts-per-sec"] = types.YLeaf{"PktsPerSec", stormControlUnit.PktsPerSec}
    return &(stormControlUnit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon
// Split Horizon
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon Group ID.
    SplitHorizonGroupId L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId
}

func (splitHorizon *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon) GetEntityData() *types.CommonEntityData {
    splitHorizon.EntityData.YFilter = splitHorizon.YFilter
    splitHorizon.EntityData.YangName = "split-horizon"
    splitHorizon.EntityData.BundleName = "cisco_ios_xr"
    splitHorizon.EntityData.ParentYangName = "bd-attachment-circuit"
    splitHorizon.EntityData.SegmentPath = "split-horizon"
    splitHorizon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    splitHorizon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    splitHorizon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    splitHorizon.EntityData.Children = make(map[string]types.YChild)
    splitHorizon.EntityData.Children["split-horizon-group-id"] = types.YChild{"SplitHorizonGroupId", &splitHorizon.SplitHorizonGroupId}
    splitHorizon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(splitHorizon.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId
// Split Horizon Group ID
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable split horizon group. The type is interface{}.
    Enable interface{}
}

func (splitHorizonGroupId *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_SplitHorizon_SplitHorizonGroupId) GetEntityData() *types.CommonEntityData {
    splitHorizonGroupId.EntityData.YFilter = splitHorizonGroupId.YFilter
    splitHorizonGroupId.EntityData.YangName = "split-horizon-group-id"
    splitHorizonGroupId.EntityData.BundleName = "cisco_ios_xr"
    splitHorizonGroupId.EntityData.ParentYangName = "split-horizon"
    splitHorizonGroupId.EntityData.SegmentPath = "split-horizon-group-id"
    splitHorizonGroupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    splitHorizonGroupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    splitHorizonGroupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    splitHorizonGroupId.EntityData.Children = make(map[string]types.YChild)
    splitHorizonGroupId.EntityData.Leafs = make(map[string]types.YLeaf)
    splitHorizonGroupId.EntityData.Leafs["enable"] = types.YLeaf{"Enable", splitHorizonGroupId.Enable}
    return &(splitHorizonGroupId.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses
// Static Mac Address Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Mac Address Configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress.
    StaticMacAddress []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress
}

func (staticMacAddresses *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses) GetEntityData() *types.CommonEntityData {
    staticMacAddresses.EntityData.YFilter = staticMacAddresses.YFilter
    staticMacAddresses.EntityData.YangName = "static-mac-addresses"
    staticMacAddresses.EntityData.BundleName = "cisco_ios_xr"
    staticMacAddresses.EntityData.ParentYangName = "bd-attachment-circuit"
    staticMacAddresses.EntityData.SegmentPath = "static-mac-addresses"
    staticMacAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMacAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMacAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMacAddresses.EntityData.Children = make(map[string]types.YChild)
    staticMacAddresses.EntityData.Children["static-mac-address"] = types.YChild{"StaticMacAddress", nil}
    for i := range staticMacAddresses.StaticMacAddress {
        staticMacAddresses.EntityData.Children[types.GetSegmentPath(&staticMacAddresses.StaticMacAddress[i])] = types.YChild{"StaticMacAddress", &staticMacAddresses.StaticMacAddress[i]}
    }
    staticMacAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticMacAddresses.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress
// Static Mac Address Configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}
}

func (staticMacAddress *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_StaticMacAddresses_StaticMacAddress) GetEntityData() *types.CommonEntityData {
    staticMacAddress.EntityData.YFilter = staticMacAddress.YFilter
    staticMacAddress.EntityData.YangName = "static-mac-address"
    staticMacAddress.EntityData.BundleName = "cisco_ios_xr"
    staticMacAddress.EntityData.ParentYangName = "static-mac-addresses"
    staticMacAddress.EntityData.SegmentPath = "static-mac-address" + "[address='" + fmt.Sprintf("%v", staticMacAddress.Address) + "']"
    staticMacAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMacAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMacAddress.EntityData.Children = make(map[string]types.YChild)
    staticMacAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    staticMacAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", staticMacAddress.Address}
    return &(staticMacAddress.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac
// MAC configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MAC Flush When Port goes down. The type is PortDownFlush.
    InterfaceMacPortDownFlush interface{}

    // Enable Mac Learning. The type is MacLearn.
    InterfaceMacLearning interface{}

    // MAC-Aging configuration commands.
    InterfaceMacAging L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging

    // MAC Secure.
    InterfaceMacSecure L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure

    // MAC-Limit configuration commands.
    InterfaceMacLimit L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit
}

func (interfaceMac *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac) GetEntityData() *types.CommonEntityData {
    interfaceMac.EntityData.YFilter = interfaceMac.YFilter
    interfaceMac.EntityData.YangName = "interface-mac"
    interfaceMac.EntityData.BundleName = "cisco_ios_xr"
    interfaceMac.EntityData.ParentYangName = "bd-attachment-circuit"
    interfaceMac.EntityData.SegmentPath = "interface-mac"
    interfaceMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMac.EntityData.Children = make(map[string]types.YChild)
    interfaceMac.EntityData.Children["interface-mac-aging"] = types.YChild{"InterfaceMacAging", &interfaceMac.InterfaceMacAging}
    interfaceMac.EntityData.Children["interface-mac-secure"] = types.YChild{"InterfaceMacSecure", &interfaceMac.InterfaceMacSecure}
    interfaceMac.EntityData.Children["interface-mac-limit"] = types.YChild{"InterfaceMacLimit", &interfaceMac.InterfaceMacLimit}
    interfaceMac.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceMac.EntityData.Leafs["interface-mac-port-down-flush"] = types.YLeaf{"InterfaceMacPortDownFlush", interfaceMac.InterfaceMacPortDownFlush}
    interfaceMac.EntityData.Leafs["interface-mac-learning"] = types.YLeaf{"InterfaceMacLearning", interfaceMac.InterfaceMacLearning}
    return &(interfaceMac.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging
// MAC-Aging configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mac Aging Time. The type is interface{} with range: 300..30000.
    InterfaceMacAgingTime interface{}

    // MAC address aging type. The type is MacAging.
    InterfaceMacAgingType interface{}
}

func (interfaceMacAging *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacAging) GetEntityData() *types.CommonEntityData {
    interfaceMacAging.EntityData.YFilter = interfaceMacAging.YFilter
    interfaceMacAging.EntityData.YangName = "interface-mac-aging"
    interfaceMacAging.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacAging.EntityData.ParentYangName = "interface-mac"
    interfaceMacAging.EntityData.SegmentPath = "interface-mac-aging"
    interfaceMacAging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacAging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacAging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacAging.EntityData.Children = make(map[string]types.YChild)
    interfaceMacAging.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceMacAging.EntityData.Leafs["interface-mac-aging-time"] = types.YLeaf{"InterfaceMacAgingTime", interfaceMacAging.InterfaceMacAgingTime}
    interfaceMacAging.EntityData.Leafs["interface-mac-aging-type"] = types.YLeaf{"InterfaceMacAgingType", interfaceMacAging.InterfaceMacAgingType}
    return &(interfaceMacAging.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure
// MAC Secure
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure struct {
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

func (interfaceMacSecure *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacSecure) GetEntityData() *types.CommonEntityData {
    interfaceMacSecure.EntityData.YFilter = interfaceMacSecure.YFilter
    interfaceMacSecure.EntityData.YangName = "interface-mac-secure"
    interfaceMacSecure.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacSecure.EntityData.ParentYangName = "interface-mac"
    interfaceMacSecure.EntityData.SegmentPath = "interface-mac-secure"
    interfaceMacSecure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacSecure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacSecure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacSecure.EntityData.Children = make(map[string]types.YChild)
    interfaceMacSecure.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceMacSecure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", interfaceMacSecure.Logging}
    interfaceMacSecure.EntityData.Leafs["disable"] = types.YLeaf{"Disable", interfaceMacSecure.Disable}
    interfaceMacSecure.EntityData.Leafs["action"] = types.YLeaf{"Action", interfaceMacSecure.Action}
    interfaceMacSecure.EntityData.Leafs["enable"] = types.YLeaf{"Enable", interfaceMacSecure.Enable}
    return &(interfaceMacSecure.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit
// MAC-Limit configuration commands
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit struct {
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

func (interfaceMacLimit *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdAttachmentCircuits_BdAttachmentCircuit_InterfaceMac_InterfaceMacLimit) GetEntityData() *types.CommonEntityData {
    interfaceMacLimit.EntityData.YFilter = interfaceMacLimit.YFilter
    interfaceMacLimit.EntityData.YangName = "interface-mac-limit"
    interfaceMacLimit.EntityData.BundleName = "cisco_ios_xr"
    interfaceMacLimit.EntityData.ParentYangName = "interface-mac"
    interfaceMacLimit.EntityData.SegmentPath = "interface-mac-limit"
    interfaceMacLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMacLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMacLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMacLimit.EntityData.Children = make(map[string]types.YChild)
    interfaceMacLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceMacLimit.EntityData.Leafs["interface-mac-limit-max"] = types.YLeaf{"InterfaceMacLimitMax", interfaceMacLimit.InterfaceMacLimitMax}
    interfaceMacLimit.EntityData.Leafs["interface-mac-limit-notif"] = types.YLeaf{"InterfaceMacLimitNotif", interfaceMacLimit.InterfaceMacLimitNotif}
    interfaceMacLimit.EntityData.Leafs["interface-mac-limit-action"] = types.YLeaf{"InterfaceMacLimitAction", interfaceMacLimit.InterfaceMacLimitAction}
    return &(interfaceMacLimit.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns
// List of EVPN pseudowires
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Pseudowire configuration. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn.
    BdPseudowireEvpn []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn
}

func (bdPseudowireEvpns *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns) GetEntityData() *types.CommonEntityData {
    bdPseudowireEvpns.EntityData.YFilter = bdPseudowireEvpns.YFilter
    bdPseudowireEvpns.EntityData.YangName = "bd-pseudowire-evpns"
    bdPseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowireEvpns.EntityData.ParentYangName = "bridge-domain"
    bdPseudowireEvpns.EntityData.SegmentPath = "bd-pseudowire-evpns"
    bdPseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowireEvpns.EntityData.Children = make(map[string]types.YChild)
    bdPseudowireEvpns.EntityData.Children["bd-pseudowire-evpn"] = types.YChild{"BdPseudowireEvpn", nil}
    for i := range bdPseudowireEvpns.BdPseudowireEvpn {
        bdPseudowireEvpns.EntityData.Children[types.GetSegmentPath(&bdPseudowireEvpns.BdPseudowireEvpn[i])] = types.YChild{"BdPseudowireEvpn", &bdPseudowireEvpns.BdPseudowireEvpn[i]}
    }
    bdPseudowireEvpns.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bdPseudowireEvpns.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn
// EVPN Pseudowire configuration
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // This attribute is a key. AC ID. The type is interface{} with range:
    // 1..4294967295.
    Acid interface{}
}

func (bdPseudowireEvpn *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_BdPseudowireEvpns_BdPseudowireEvpn) GetEntityData() *types.CommonEntityData {
    bdPseudowireEvpn.EntityData.YFilter = bdPseudowireEvpn.YFilter
    bdPseudowireEvpn.EntityData.YangName = "bd-pseudowire-evpn"
    bdPseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    bdPseudowireEvpn.EntityData.ParentYangName = "bd-pseudowire-evpns"
    bdPseudowireEvpn.EntityData.SegmentPath = "bd-pseudowire-evpn" + "[eviid='" + fmt.Sprintf("%v", bdPseudowireEvpn.Eviid) + "']" + "[acid='" + fmt.Sprintf("%v", bdPseudowireEvpn.Acid) + "']"
    bdPseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdPseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdPseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdPseudowireEvpn.EntityData.Children = make(map[string]types.YChild)
    bdPseudowireEvpn.EntityData.Leafs = make(map[string]types.YLeaf)
    bdPseudowireEvpn.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", bdPseudowireEvpn.Eviid}
    bdPseudowireEvpn.EntityData.Leafs["acid"] = types.YLeaf{"Acid", bdPseudowireEvpn.Acid}
    return &(bdPseudowireEvpn.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard
// IP Source Guard
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Logging. The type is interface{}.
    Logging interface{}

    // Enable IP Source Guard. The type is interface{}.
    Enable interface{}
}

func (ipSourceGuard *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_IpSourceGuard) GetEntityData() *types.CommonEntityData {
    ipSourceGuard.EntityData.YFilter = ipSourceGuard.YFilter
    ipSourceGuard.EntityData.YangName = "ip-source-guard"
    ipSourceGuard.EntityData.BundleName = "cisco_ios_xr"
    ipSourceGuard.EntityData.ParentYangName = "bridge-domain"
    ipSourceGuard.EntityData.SegmentPath = "ip-source-guard"
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSourceGuard.EntityData.Children = make(map[string]types.YChild)
    ipSourceGuard.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSourceGuard.EntityData.Leafs["logging"] = types.YLeaf{"Logging", ipSourceGuard.Logging}
    ipSourceGuard.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ipSourceGuard.Enable}
    return &(ipSourceGuard.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai
// Dynamic ARP Inspection
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Logging. The type is interface{}.
    Logging interface{}

    // Enable Dynamic ARP Inspection. The type is interface{}.
    Enable interface{}

    // Address Validation.
    DaiAddressValidation L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation
}

func (dai *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai) GetEntityData() *types.CommonEntityData {
    dai.EntityData.YFilter = dai.YFilter
    dai.EntityData.YangName = "dai"
    dai.EntityData.BundleName = "cisco_ios_xr"
    dai.EntityData.ParentYangName = "bridge-domain"
    dai.EntityData.SegmentPath = "dai"
    dai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dai.EntityData.Children = make(map[string]types.YChild)
    dai.EntityData.Children["dai-address-validation"] = types.YChild{"DaiAddressValidation", &dai.DaiAddressValidation}
    dai.EntityData.Leafs = make(map[string]types.YLeaf)
    dai.EntityData.Leafs["logging"] = types.YLeaf{"Logging", dai.Logging}
    dai.EntityData.Leafs["enable"] = types.YLeaf{"Enable", dai.Enable}
    return &(dai.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation
// Address Validation
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation struct {
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

func (daiAddressValidation *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_Dai_DaiAddressValidation) GetEntityData() *types.CommonEntityData {
    daiAddressValidation.EntityData.YFilter = daiAddressValidation.YFilter
    daiAddressValidation.EntityData.YangName = "dai-address-validation"
    daiAddressValidation.EntityData.BundleName = "cisco_ios_xr"
    daiAddressValidation.EntityData.ParentYangName = "dai"
    daiAddressValidation.EntityData.SegmentPath = "dai-address-validation"
    daiAddressValidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    daiAddressValidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    daiAddressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    daiAddressValidation.EntityData.Children = make(map[string]types.YChild)
    daiAddressValidation.EntityData.Leafs = make(map[string]types.YLeaf)
    daiAddressValidation.EntityData.Leafs["ipv4-verification"] = types.YLeaf{"Ipv4Verification", daiAddressValidation.Ipv4Verification}
    daiAddressValidation.EntityData.Leafs["destination-mac-verification"] = types.YLeaf{"DestinationMacVerification", daiAddressValidation.DestinationMacVerification}
    daiAddressValidation.EntityData.Leafs["source-mac-verification"] = types.YLeaf{"SourceMacVerification", daiAddressValidation.SourceMacVerification}
    daiAddressValidation.EntityData.Leafs["enable"] = types.YLeaf{"Enable", daiAddressValidation.Enable}
    return &(daiAddressValidation.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces
// Bridge Domain Routed Interface Table
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge Domain Routed Interface. The type is slice of
    // L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface.
    RoutedInterface []L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface
}

func (routedInterfaces *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces) GetEntityData() *types.CommonEntityData {
    routedInterfaces.EntityData.YFilter = routedInterfaces.YFilter
    routedInterfaces.EntityData.YangName = "routed-interfaces"
    routedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaces.EntityData.ParentYangName = "bridge-domain"
    routedInterfaces.EntityData.SegmentPath = "routed-interfaces"
    routedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaces.EntityData.Children = make(map[string]types.YChild)
    routedInterfaces.EntityData.Children["routed-interface"] = types.YChild{"RoutedInterface", nil}
    for i := range routedInterfaces.RoutedInterface {
        routedInterfaces.EntityData.Children[types.GetSegmentPath(&routedInterfaces.RoutedInterface[i])] = types.YChild{"RoutedInterface", &routedInterfaces.RoutedInterface[i]}
    }
    routedInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routedInterfaces.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface
// Bridge Domain Routed Interface
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Routed Interface. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Routed interface split horizon group.
    RoutedInterfaceSplitHorizonGroup L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup
}

func (routedInterface *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface) GetEntityData() *types.CommonEntityData {
    routedInterface.EntityData.YFilter = routedInterface.YFilter
    routedInterface.EntityData.YangName = "routed-interface"
    routedInterface.EntityData.BundleName = "cisco_ios_xr"
    routedInterface.EntityData.ParentYangName = "routed-interfaces"
    routedInterface.EntityData.SegmentPath = "routed-interface" + "[interface-name='" + fmt.Sprintf("%v", routedInterface.InterfaceName) + "']"
    routedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterface.EntityData.Children = make(map[string]types.YChild)
    routedInterface.EntityData.Children["routed-interface-split-horizon-group"] = types.YChild{"RoutedInterfaceSplitHorizonGroup", &routedInterface.RoutedInterfaceSplitHorizonGroup}
    routedInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    routedInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", routedInterface.InterfaceName}
    return &(routedInterface.EntityData)
}

// L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup
// Routed interface split horizon group
type L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure BVI under SHG 1. The type is interface{}.
    RoutedInterfaceSplitHorizonGroupCore interface{}
}

func (routedInterfaceSplitHorizonGroup *L2Vpn_Database_BridgeDomainGroups_BridgeDomainGroup_BridgeDomains_BridgeDomain_RoutedInterfaces_RoutedInterface_RoutedInterfaceSplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    routedInterfaceSplitHorizonGroup.EntityData.YFilter = routedInterfaceSplitHorizonGroup.YFilter
    routedInterfaceSplitHorizonGroup.EntityData.YangName = "routed-interface-split-horizon-group"
    routedInterfaceSplitHorizonGroup.EntityData.BundleName = "cisco_ios_xr"
    routedInterfaceSplitHorizonGroup.EntityData.ParentYangName = "routed-interface"
    routedInterfaceSplitHorizonGroup.EntityData.SegmentPath = "routed-interface-split-horizon-group"
    routedInterfaceSplitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routedInterfaceSplitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routedInterfaceSplitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routedInterfaceSplitHorizonGroup.EntityData.Children = make(map[string]types.YChild)
    routedInterfaceSplitHorizonGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    routedInterfaceSplitHorizonGroup.EntityData.Leafs["routed-interface-split-horizon-group-core"] = types.YLeaf{"RoutedInterfaceSplitHorizonGroupCore", routedInterfaceSplitHorizonGroup.RoutedInterfaceSplitHorizonGroupCore}
    return &(routedInterfaceSplitHorizonGroup.EntityData)
}

// L2Vpn_Database_PseudowireClasses
// List of pseudowire classes
type L2Vpn_Database_PseudowireClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire class. The type is slice of
    // L2Vpn_Database_PseudowireClasses_PseudowireClass.
    PseudowireClass []L2Vpn_Database_PseudowireClasses_PseudowireClass
}

func (pseudowireClasses *L2Vpn_Database_PseudowireClasses) GetEntityData() *types.CommonEntityData {
    pseudowireClasses.EntityData.YFilter = pseudowireClasses.YFilter
    pseudowireClasses.EntityData.YangName = "pseudowire-classes"
    pseudowireClasses.EntityData.BundleName = "cisco_ios_xr"
    pseudowireClasses.EntityData.ParentYangName = "database"
    pseudowireClasses.EntityData.SegmentPath = "pseudowire-classes"
    pseudowireClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireClasses.EntityData.Children = make(map[string]types.YChild)
    pseudowireClasses.EntityData.Children["pseudowire-class"] = types.YChild{"PseudowireClass", nil}
    for i := range pseudowireClasses.PseudowireClass {
        pseudowireClasses.EntityData.Children[types.GetSegmentPath(&pseudowireClasses.PseudowireClass[i])] = types.YChild{"PseudowireClass", &pseudowireClasses.PseudowireClass[i]}
    }
    pseudowireClasses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pseudowireClasses.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass
// Pseudowire class
type L2Vpn_Database_PseudowireClasses_PseudowireClass struct {
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
    L2Tpv3Encapsulation L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation

    // Back Up Pseudowire class.
    BackupDisableDelay L2Vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay

    // MPLS encapsulation.
    MplsEncapsulation L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation
}

func (pseudowireClass *L2Vpn_Database_PseudowireClasses_PseudowireClass) GetEntityData() *types.CommonEntityData {
    pseudowireClass.EntityData.YFilter = pseudowireClass.YFilter
    pseudowireClass.EntityData.YangName = "pseudowire-class"
    pseudowireClass.EntityData.BundleName = "cisco_ios_xr"
    pseudowireClass.EntityData.ParentYangName = "pseudowire-classes"
    pseudowireClass.EntityData.SegmentPath = "pseudowire-class" + "[name='" + fmt.Sprintf("%v", pseudowireClass.Name) + "']"
    pseudowireClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireClass.EntityData.Children = make(map[string]types.YChild)
    pseudowireClass.EntityData.Children["l2tpv3-encapsulation"] = types.YChild{"L2Tpv3Encapsulation", &pseudowireClass.L2Tpv3Encapsulation}
    pseudowireClass.EntityData.Children["backup-disable-delay"] = types.YChild{"BackupDisableDelay", &pseudowireClass.BackupDisableDelay}
    pseudowireClass.EntityData.Children["mpls-encapsulation"] = types.YChild{"MplsEncapsulation", &pseudowireClass.MplsEncapsulation}
    pseudowireClass.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireClass.EntityData.Leafs["name"] = types.YLeaf{"Name", pseudowireClass.Name}
    pseudowireClass.EntityData.Leafs["mac-withdraw"] = types.YLeaf{"MacWithdraw", pseudowireClass.MacWithdraw}
    pseudowireClass.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pseudowireClass.Enable}
    return &(pseudowireClass.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation
// L2TPv3 encapsulation
type L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the do not fragment bit to 1. The type is interface{}.
    DfBitSet interface{}

    // Cookie size. The type is L2tpCookieSize. The default value is zero.
    CookieSize interface{}

    // Source IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Transport mode. The type is TransportMode.
    TransportMode interface{}

    // Enable L2TPv3 encapsulation. The type is interface{}.
    Enable interface{}

    // Time to live. The type is interface{} with range: 1..255.
    TimeToLive interface{}

    // Sequencing.
    Sequencing L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_Sequencing

    // Type of service.
    TypeOfService L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_TypeOfService

    // L2TPv3 signaling protocol.
    SignalingProtocol L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_SignalingProtocol

    // Path maximum transmission unit.
    PathMtu L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_PathMtu
}

func (l2Tpv3Encapsulation *L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation) GetEntityData() *types.CommonEntityData {
    l2Tpv3Encapsulation.EntityData.YFilter = l2Tpv3Encapsulation.YFilter
    l2Tpv3Encapsulation.EntityData.YangName = "l2tpv3-encapsulation"
    l2Tpv3Encapsulation.EntityData.BundleName = "cisco_ios_xr"
    l2Tpv3Encapsulation.EntityData.ParentYangName = "pseudowire-class"
    l2Tpv3Encapsulation.EntityData.SegmentPath = "l2tpv3-encapsulation"
    l2Tpv3Encapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Tpv3Encapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Tpv3Encapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Tpv3Encapsulation.EntityData.Children = make(map[string]types.YChild)
    l2Tpv3Encapsulation.EntityData.Children["sequencing"] = types.YChild{"Sequencing", &l2Tpv3Encapsulation.Sequencing}
    l2Tpv3Encapsulation.EntityData.Children["type-of-service"] = types.YChild{"TypeOfService", &l2Tpv3Encapsulation.TypeOfService}
    l2Tpv3Encapsulation.EntityData.Children["signaling-protocol"] = types.YChild{"SignalingProtocol", &l2Tpv3Encapsulation.SignalingProtocol}
    l2Tpv3Encapsulation.EntityData.Children["path-mtu"] = types.YChild{"PathMtu", &l2Tpv3Encapsulation.PathMtu}
    l2Tpv3Encapsulation.EntityData.Leafs = make(map[string]types.YLeaf)
    l2Tpv3Encapsulation.EntityData.Leafs["df-bit-set"] = types.YLeaf{"DfBitSet", l2Tpv3Encapsulation.DfBitSet}
    l2Tpv3Encapsulation.EntityData.Leafs["cookie-size"] = types.YLeaf{"CookieSize", l2Tpv3Encapsulation.CookieSize}
    l2Tpv3Encapsulation.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", l2Tpv3Encapsulation.SourceAddress}
    l2Tpv3Encapsulation.EntityData.Leafs["transport-mode"] = types.YLeaf{"TransportMode", l2Tpv3Encapsulation.TransportMode}
    l2Tpv3Encapsulation.EntityData.Leafs["enable"] = types.YLeaf{"Enable", l2Tpv3Encapsulation.Enable}
    l2Tpv3Encapsulation.EntityData.Leafs["time-to-live"] = types.YLeaf{"TimeToLive", l2Tpv3Encapsulation.TimeToLive}
    return &(l2Tpv3Encapsulation.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_Sequencing
// Sequencing
type L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_Sequencing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequencing. The type is L2tpv3Sequencing. The default value is off.
    Sequencing interface{}

    // Out of sequence threshold. The type is interface{} with range: 5..65535.
    // The default value is 5.
    ResyncThreshold interface{}
}

func (sequencing *L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_Sequencing) GetEntityData() *types.CommonEntityData {
    sequencing.EntityData.YFilter = sequencing.YFilter
    sequencing.EntityData.YangName = "sequencing"
    sequencing.EntityData.BundleName = "cisco_ios_xr"
    sequencing.EntityData.ParentYangName = "l2tpv3-encapsulation"
    sequencing.EntityData.SegmentPath = "sequencing"
    sequencing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sequencing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sequencing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sequencing.EntityData.Children = make(map[string]types.YChild)
    sequencing.EntityData.Leafs = make(map[string]types.YLeaf)
    sequencing.EntityData.Leafs["sequencing"] = types.YLeaf{"Sequencing", sequencing.Sequencing}
    sequencing.EntityData.Leafs["resync-threshold"] = types.YLeaf{"ResyncThreshold", sequencing.ResyncThreshold}
    return &(sequencing.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_TypeOfService
// Type of service
type L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_TypeOfService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of service value. The type is interface{} with range: 0..255.
    TypeOfServiceValue interface{}

    // Type of service mode. The type is TypeOfServiceMode.
    TypeOfServiceMode interface{}
}

func (typeOfService *L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_TypeOfService) GetEntityData() *types.CommonEntityData {
    typeOfService.EntityData.YFilter = typeOfService.YFilter
    typeOfService.EntityData.YangName = "type-of-service"
    typeOfService.EntityData.BundleName = "cisco_ios_xr"
    typeOfService.EntityData.ParentYangName = "l2tpv3-encapsulation"
    typeOfService.EntityData.SegmentPath = "type-of-service"
    typeOfService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeOfService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeOfService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeOfService.EntityData.Children = make(map[string]types.YChild)
    typeOfService.EntityData.Leafs = make(map[string]types.YLeaf)
    typeOfService.EntityData.Leafs["type-of-service-value"] = types.YLeaf{"TypeOfServiceValue", typeOfService.TypeOfServiceValue}
    typeOfService.EntityData.Leafs["type-of-service-mode"] = types.YLeaf{"TypeOfServiceMode", typeOfService.TypeOfServiceMode}
    return &(typeOfService.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_SignalingProtocol
// L2TPv3 signaling protocol
type L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_SignalingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2TPv3 signaling protocol. The type is L2tpSignalingProtocol. The default
    // value is l2tpv3.
    Protocol interface{}

    // Name of the L2TPv3 class name. The type is string with length: 1..32.
    L2Tpv3ClassName interface{}
}

func (signalingProtocol *L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_SignalingProtocol) GetEntityData() *types.CommonEntityData {
    signalingProtocol.EntityData.YFilter = signalingProtocol.YFilter
    signalingProtocol.EntityData.YangName = "signaling-protocol"
    signalingProtocol.EntityData.BundleName = "cisco_ios_xr"
    signalingProtocol.EntityData.ParentYangName = "l2tpv3-encapsulation"
    signalingProtocol.EntityData.SegmentPath = "signaling-protocol"
    signalingProtocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalingProtocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalingProtocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalingProtocol.EntityData.Children = make(map[string]types.YChild)
    signalingProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    signalingProtocol.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", signalingProtocol.Protocol}
    signalingProtocol.EntityData.Leafs["l2tpv3-class-name"] = types.YLeaf{"L2Tpv3ClassName", signalingProtocol.L2Tpv3ClassName}
    return &(signalingProtocol.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_PathMtu
// Path maximum transmission unit
type L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_PathMtu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path MTU. The type is interface{}.
    Enable interface{}

    // Maximum path maximum transmission unit. The type is interface{} with range:
    // 68..65535.
    MaxPathMtu interface{}
}

func (pathMtu *L2Vpn_Database_PseudowireClasses_PseudowireClass_L2Tpv3Encapsulation_PathMtu) GetEntityData() *types.CommonEntityData {
    pathMtu.EntityData.YFilter = pathMtu.YFilter
    pathMtu.EntityData.YangName = "path-mtu"
    pathMtu.EntityData.BundleName = "cisco_ios_xr"
    pathMtu.EntityData.ParentYangName = "l2tpv3-encapsulation"
    pathMtu.EntityData.SegmentPath = "path-mtu"
    pathMtu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathMtu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathMtu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathMtu.EntityData.Children = make(map[string]types.YChild)
    pathMtu.EntityData.Leafs = make(map[string]types.YLeaf)
    pathMtu.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pathMtu.Enable}
    pathMtu.EntityData.Leafs["max-path-mtu"] = types.YLeaf{"MaxPathMtu", pathMtu.MaxPathMtu}
    return &(pathMtu.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay
// Back Up Pseudowire class
type L2Vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay or Never. The type is BackupDisable.
    Type_ interface{}

    // Disable backup delay. The type is interface{} with range: 0..180.
    DisableBackup interface{}
}

func (backupDisableDelay *L2Vpn_Database_PseudowireClasses_PseudowireClass_BackupDisableDelay) GetEntityData() *types.CommonEntityData {
    backupDisableDelay.EntityData.YFilter = backupDisableDelay.YFilter
    backupDisableDelay.EntityData.YangName = "backup-disable-delay"
    backupDisableDelay.EntityData.BundleName = "cisco_ios_xr"
    backupDisableDelay.EntityData.ParentYangName = "pseudowire-class"
    backupDisableDelay.EntityData.SegmentPath = "backup-disable-delay"
    backupDisableDelay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupDisableDelay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupDisableDelay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupDisableDelay.EntityData.Children = make(map[string]types.YChild)
    backupDisableDelay.EntityData.Leafs = make(map[string]types.YLeaf)
    backupDisableDelay.EntityData.Leafs["type"] = types.YLeaf{"Type_", backupDisableDelay.Type_}
    backupDisableDelay.EntityData.Leafs["disable-backup"] = types.YLeaf{"DisableBackup", backupDisableDelay.DisableBackup}
    return &(backupDisableDelay.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation
// MPLS encapsulation
type L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation struct {
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Transport mode. The type is TransportMode.
    TransportMode interface{}

    // Enable MPLS encapsulation. The type is interface{}.
    Enable interface{}

    // Enable control word. The type is ControlWord.
    ControlWord interface{}

    // Sequencing.
    Sequencing L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing

    // Redundancy options for MPLS encapsulation.
    MplsRedundancy L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy

    // Preferred path.
    PreferredPath L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath

    // Load Balancing.
    LoadBalanceGroup L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup
}

func (mplsEncapsulation *L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation) GetEntityData() *types.CommonEntityData {
    mplsEncapsulation.EntityData.YFilter = mplsEncapsulation.YFilter
    mplsEncapsulation.EntityData.YangName = "mpls-encapsulation"
    mplsEncapsulation.EntityData.BundleName = "cisco_ios_xr"
    mplsEncapsulation.EntityData.ParentYangName = "pseudowire-class"
    mplsEncapsulation.EntityData.SegmentPath = "mpls-encapsulation"
    mplsEncapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsEncapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsEncapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsEncapsulation.EntityData.Children = make(map[string]types.YChild)
    mplsEncapsulation.EntityData.Children["sequencing"] = types.YChild{"Sequencing", &mplsEncapsulation.Sequencing}
    mplsEncapsulation.EntityData.Children["mpls-redundancy"] = types.YChild{"MplsRedundancy", &mplsEncapsulation.MplsRedundancy}
    mplsEncapsulation.EntityData.Children["preferred-path"] = types.YChild{"PreferredPath", &mplsEncapsulation.PreferredPath}
    mplsEncapsulation.EntityData.Children["load-balance-group"] = types.YChild{"LoadBalanceGroup", &mplsEncapsulation.LoadBalanceGroup}
    mplsEncapsulation.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsEncapsulation.EntityData.Leafs["pw-switching-tlv"] = types.YLeaf{"PwSwitchingTlv", mplsEncapsulation.PwSwitchingTlv}
    mplsEncapsulation.EntityData.Leafs["static-tag-rewrite"] = types.YLeaf{"StaticTagRewrite", mplsEncapsulation.StaticTagRewrite}
    mplsEncapsulation.EntityData.Leafs["signaling-protocol"] = types.YLeaf{"SignalingProtocol", mplsEncapsulation.SignalingProtocol}
    mplsEncapsulation.EntityData.Leafs["vccv-type"] = types.YLeaf{"VccvType", mplsEncapsulation.VccvType}
    mplsEncapsulation.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", mplsEncapsulation.SourceAddress}
    mplsEncapsulation.EntityData.Leafs["transport-mode"] = types.YLeaf{"TransportMode", mplsEncapsulation.TransportMode}
    mplsEncapsulation.EntityData.Leafs["enable"] = types.YLeaf{"Enable", mplsEncapsulation.Enable}
    mplsEncapsulation.EntityData.Leafs["control-word"] = types.YLeaf{"ControlWord", mplsEncapsulation.ControlWord}
    return &(mplsEncapsulation.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing
// Sequencing
type L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequencing. The type is MplsSequencing. The default value is off.
    Sequencing interface{}

    // Out of sequence threshold. The type is interface{} with range: 5..65535.
    // The default value is 5.
    ResyncThreshold interface{}
}

func (sequencing *L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_Sequencing) GetEntityData() *types.CommonEntityData {
    sequencing.EntityData.YFilter = sequencing.YFilter
    sequencing.EntityData.YangName = "sequencing"
    sequencing.EntityData.BundleName = "cisco_ios_xr"
    sequencing.EntityData.ParentYangName = "mpls-encapsulation"
    sequencing.EntityData.SegmentPath = "sequencing"
    sequencing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sequencing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sequencing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sequencing.EntityData.Children = make(map[string]types.YChild)
    sequencing.EntityData.Leafs = make(map[string]types.YLeaf)
    sequencing.EntityData.Leafs["sequencing"] = types.YLeaf{"Sequencing", sequencing.Sequencing}
    sequencing.EntityData.Leafs["resync-threshold"] = types.YLeaf{"ResyncThreshold", sequencing.ResyncThreshold}
    return &(sequencing.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy
// Redundancy options for MPLS encapsulation
type L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force one-way PW redundancy behavior in Redundancy Group. The type is
    // interface{}.
    RedundancyOneWay interface{}

    // Initial delay before activating the redundant PW, in seconds. The type is
    // interface{} with range: 0..120. Units are second.
    RedundancyInitialDelay interface{}
}

func (mplsRedundancy *L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_MplsRedundancy) GetEntityData() *types.CommonEntityData {
    mplsRedundancy.EntityData.YFilter = mplsRedundancy.YFilter
    mplsRedundancy.EntityData.YangName = "mpls-redundancy"
    mplsRedundancy.EntityData.BundleName = "cisco_ios_xr"
    mplsRedundancy.EntityData.ParentYangName = "mpls-encapsulation"
    mplsRedundancy.EntityData.SegmentPath = "mpls-redundancy"
    mplsRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsRedundancy.EntityData.Children = make(map[string]types.YChild)
    mplsRedundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsRedundancy.EntityData.Leafs["redundancy-one-way"] = types.YLeaf{"RedundancyOneWay", mplsRedundancy.RedundancyOneWay}
    mplsRedundancy.EntityData.Leafs["redundancy-initial-delay"] = types.YLeaf{"RedundancyInitialDelay", mplsRedundancy.RedundancyInitialDelay}
    return &(mplsRedundancy.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath
// Preferred path
type L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Preferred Path Type. The type is PreferredPath.
    Type_ interface{}

    // Interface Tunnel number for preferred path. The type is interface{} with
    // range: 0..65535.
    InterfaceTunnelNumber interface{}

    // Fallback disable. The type is interface{}.
    FallbackDisable interface{}

    // Name of the SR TE Policy. The type is string with length: 1..60.
    SrtePolicy interface{}
}

func (preferredPath *L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_PreferredPath) GetEntityData() *types.CommonEntityData {
    preferredPath.EntityData.YFilter = preferredPath.YFilter
    preferredPath.EntityData.YangName = "preferred-path"
    preferredPath.EntityData.BundleName = "cisco_ios_xr"
    preferredPath.EntityData.ParentYangName = "mpls-encapsulation"
    preferredPath.EntityData.SegmentPath = "preferred-path"
    preferredPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preferredPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preferredPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preferredPath.EntityData.Children = make(map[string]types.YChild)
    preferredPath.EntityData.Leafs = make(map[string]types.YLeaf)
    preferredPath.EntityData.Leafs["type"] = types.YLeaf{"Type_", preferredPath.Type_}
    preferredPath.EntityData.Leafs["interface-tunnel-number"] = types.YLeaf{"InterfaceTunnelNumber", preferredPath.InterfaceTunnelNumber}
    preferredPath.EntityData.Leafs["fallback-disable"] = types.YLeaf{"FallbackDisable", preferredPath.FallbackDisable}
    preferredPath.EntityData.Leafs["srte-policy"] = types.YLeaf{"SrtePolicy", preferredPath.SrtePolicy}
    return &(preferredPath.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup
// Load Balancing
type L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Legacy Flow Label TLV code. The type is FlowLabelTlvCode.
    FlowLabelLoadBalanceCode interface{}

    // Enable PW Label based Load Balancing. The type is LoadBalance.
    PwLabelLoadBalance interface{}

    // Enable Flow Label based load balancing.
    FlowLabelLoadBalance L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance
}

func (loadBalanceGroup *L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup) GetEntityData() *types.CommonEntityData {
    loadBalanceGroup.EntityData.YFilter = loadBalanceGroup.YFilter
    loadBalanceGroup.EntityData.YangName = "load-balance-group"
    loadBalanceGroup.EntityData.BundleName = "cisco_ios_xr"
    loadBalanceGroup.EntityData.ParentYangName = "mpls-encapsulation"
    loadBalanceGroup.EntityData.SegmentPath = "load-balance-group"
    loadBalanceGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalanceGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalanceGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalanceGroup.EntityData.Children = make(map[string]types.YChild)
    loadBalanceGroup.EntityData.Children["flow-label-load-balance"] = types.YChild{"FlowLabelLoadBalance", &loadBalanceGroup.FlowLabelLoadBalance}
    loadBalanceGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    loadBalanceGroup.EntityData.Leafs["flow-label-load-balance-code"] = types.YLeaf{"FlowLabelLoadBalanceCode", loadBalanceGroup.FlowLabelLoadBalanceCode}
    loadBalanceGroup.EntityData.Leafs["pw-label-load-balance"] = types.YLeaf{"PwLabelLoadBalance", loadBalanceGroup.PwLabelLoadBalance}
    return &(loadBalanceGroup.EntityData)
}

// L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance
// Enable Flow Label based load balancing
type L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Label load balance type. The type is FlowLabelLoadBalance.
    FlowLabel interface{}

    // Static Flow Label. The type is interface{}.
    Static interface{}
}

func (flowLabelLoadBalance *L2Vpn_Database_PseudowireClasses_PseudowireClass_MplsEncapsulation_LoadBalanceGroup_FlowLabelLoadBalance) GetEntityData() *types.CommonEntityData {
    flowLabelLoadBalance.EntityData.YFilter = flowLabelLoadBalance.YFilter
    flowLabelLoadBalance.EntityData.YangName = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    flowLabelLoadBalance.EntityData.ParentYangName = "load-balance-group"
    flowLabelLoadBalance.EntityData.SegmentPath = "flow-label-load-balance"
    flowLabelLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabelLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabelLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabelLoadBalance.EntityData.Children = make(map[string]types.YChild)
    flowLabelLoadBalance.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabelLoadBalance.EntityData.Leafs["flow-label"] = types.YLeaf{"FlowLabel", flowLabelLoadBalance.FlowLabel}
    flowLabelLoadBalance.EntityData.Leafs["static"] = types.YLeaf{"Static", flowLabelLoadBalance.Static}
    return &(flowLabelLoadBalance.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable
// List of Flexible XConnect Services
type L2Vpn_Database_FlexibleXconnectServiceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Vlan-Unaware Flexible XConnect Services.
    VlanUnawareFlexibleXconnectServices L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices

    // List of Vlan-Aware Flexible XConnect Services.
    VlanAwareFlexibleXconnectServices L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices
}

func (flexibleXconnectServiceTable *L2Vpn_Database_FlexibleXconnectServiceTable) GetEntityData() *types.CommonEntityData {
    flexibleXconnectServiceTable.EntityData.YFilter = flexibleXconnectServiceTable.YFilter
    flexibleXconnectServiceTable.EntityData.YangName = "flexible-xconnect-service-table"
    flexibleXconnectServiceTable.EntityData.BundleName = "cisco_ios_xr"
    flexibleXconnectServiceTable.EntityData.ParentYangName = "database"
    flexibleXconnectServiceTable.EntityData.SegmentPath = "flexible-xconnect-service-table"
    flexibleXconnectServiceTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flexibleXconnectServiceTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flexibleXconnectServiceTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flexibleXconnectServiceTable.EntityData.Children = make(map[string]types.YChild)
    flexibleXconnectServiceTable.EntityData.Children["vlan-unaware-flexible-xconnect-services"] = types.YChild{"VlanUnawareFlexibleXconnectServices", &flexibleXconnectServiceTable.VlanUnawareFlexibleXconnectServices}
    flexibleXconnectServiceTable.EntityData.Children["vlan-aware-flexible-xconnect-services"] = types.YChild{"VlanAwareFlexibleXconnectServices", &flexibleXconnectServiceTable.VlanAwareFlexibleXconnectServices}
    flexibleXconnectServiceTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flexibleXconnectServiceTable.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices
// List of Vlan-Unaware Flexible XConnect
// Services
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flexible XConnect Service. The type is slice of
    // L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService.
    VlanUnawareFlexibleXconnectService []L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService
}

func (vlanUnawareFlexibleXconnectServices *L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices) GetEntityData() *types.CommonEntityData {
    vlanUnawareFlexibleXconnectServices.EntityData.YFilter = vlanUnawareFlexibleXconnectServices.YFilter
    vlanUnawareFlexibleXconnectServices.EntityData.YangName = "vlan-unaware-flexible-xconnect-services"
    vlanUnawareFlexibleXconnectServices.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFlexibleXconnectServices.EntityData.ParentYangName = "flexible-xconnect-service-table"
    vlanUnawareFlexibleXconnectServices.EntityData.SegmentPath = "vlan-unaware-flexible-xconnect-services"
    vlanUnawareFlexibleXconnectServices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFlexibleXconnectServices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFlexibleXconnectServices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFlexibleXconnectServices.EntityData.Children = make(map[string]types.YChild)
    vlanUnawareFlexibleXconnectServices.EntityData.Children["vlan-unaware-flexible-xconnect-service"] = types.YChild{"VlanUnawareFlexibleXconnectService", nil}
    for i := range vlanUnawareFlexibleXconnectServices.VlanUnawareFlexibleXconnectService {
        vlanUnawareFlexibleXconnectServices.EntityData.Children[types.GetSegmentPath(&vlanUnawareFlexibleXconnectServices.VlanUnawareFlexibleXconnectService[i])] = types.YChild{"VlanUnawareFlexibleXconnectService", &vlanUnawareFlexibleXconnectServices.VlanUnawareFlexibleXconnectService[i]}
    }
    vlanUnawareFlexibleXconnectServices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlanUnawareFlexibleXconnectServices.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService
// Flexible XConnect Service
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Flexible XConnect Service. The type is
    // string with length: 1..23.
    Name interface{}

    // List of attachment circuits.
    VlanUnawareFxcAttachmentCircuits L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits

    // List of EVPN Services.
    VlanUnawareFxcPseudowireEvpns L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns
}

func (vlanUnawareFlexibleXconnectService *L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService) GetEntityData() *types.CommonEntityData {
    vlanUnawareFlexibleXconnectService.EntityData.YFilter = vlanUnawareFlexibleXconnectService.YFilter
    vlanUnawareFlexibleXconnectService.EntityData.YangName = "vlan-unaware-flexible-xconnect-service"
    vlanUnawareFlexibleXconnectService.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFlexibleXconnectService.EntityData.ParentYangName = "vlan-unaware-flexible-xconnect-services"
    vlanUnawareFlexibleXconnectService.EntityData.SegmentPath = "vlan-unaware-flexible-xconnect-service" + "[name='" + fmt.Sprintf("%v", vlanUnawareFlexibleXconnectService.Name) + "']"
    vlanUnawareFlexibleXconnectService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFlexibleXconnectService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFlexibleXconnectService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFlexibleXconnectService.EntityData.Children = make(map[string]types.YChild)
    vlanUnawareFlexibleXconnectService.EntityData.Children["vlan-unaware-fxc-attachment-circuits"] = types.YChild{"VlanUnawareFxcAttachmentCircuits", &vlanUnawareFlexibleXconnectService.VlanUnawareFxcAttachmentCircuits}
    vlanUnawareFlexibleXconnectService.EntityData.Children["vlan-unaware-fxc-pseudowire-evpns"] = types.YChild{"VlanUnawareFxcPseudowireEvpns", &vlanUnawareFlexibleXconnectService.VlanUnawareFxcPseudowireEvpns}
    vlanUnawareFlexibleXconnectService.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanUnawareFlexibleXconnectService.EntityData.Leafs["name"] = types.YLeaf{"Name", vlanUnawareFlexibleXconnectService.Name}
    return &(vlanUnawareFlexibleXconnectService.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits
// List of attachment circuits
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit.
    VlanUnawareFxcAttachmentCircuit []L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit
}

func (vlanUnawareFxcAttachmentCircuits *L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcAttachmentCircuits.EntityData.YFilter = vlanUnawareFxcAttachmentCircuits.YFilter
    vlanUnawareFxcAttachmentCircuits.EntityData.YangName = "vlan-unaware-fxc-attachment-circuits"
    vlanUnawareFxcAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcAttachmentCircuits.EntityData.ParentYangName = "vlan-unaware-flexible-xconnect-service"
    vlanUnawareFxcAttachmentCircuits.EntityData.SegmentPath = "vlan-unaware-fxc-attachment-circuits"
    vlanUnawareFxcAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcAttachmentCircuits.EntityData.Children = make(map[string]types.YChild)
    vlanUnawareFxcAttachmentCircuits.EntityData.Children["vlan-unaware-fxc-attachment-circuit"] = types.YChild{"VlanUnawareFxcAttachmentCircuit", nil}
    for i := range vlanUnawareFxcAttachmentCircuits.VlanUnawareFxcAttachmentCircuit {
        vlanUnawareFxcAttachmentCircuits.EntityData.Children[types.GetSegmentPath(&vlanUnawareFxcAttachmentCircuits.VlanUnawareFxcAttachmentCircuit[i])] = types.YChild{"VlanUnawareFxcAttachmentCircuit", &vlanUnawareFxcAttachmentCircuits.VlanUnawareFxcAttachmentCircuit[i]}
    }
    vlanUnawareFxcAttachmentCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlanUnawareFxcAttachmentCircuits.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit
// Attachment circuit interface
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: b'[a-zA-Z0-9./-]+'.
    Name interface{}
}

func (vlanUnawareFxcAttachmentCircuit *L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcAttachmentCircuits_VlanUnawareFxcAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcAttachmentCircuit.EntityData.YFilter = vlanUnawareFxcAttachmentCircuit.YFilter
    vlanUnawareFxcAttachmentCircuit.EntityData.YangName = "vlan-unaware-fxc-attachment-circuit"
    vlanUnawareFxcAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcAttachmentCircuit.EntityData.ParentYangName = "vlan-unaware-fxc-attachment-circuits"
    vlanUnawareFxcAttachmentCircuit.EntityData.SegmentPath = "vlan-unaware-fxc-attachment-circuit" + "[name='" + fmt.Sprintf("%v", vlanUnawareFxcAttachmentCircuit.Name) + "']"
    vlanUnawareFxcAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcAttachmentCircuit.EntityData.Children = make(map[string]types.YChild)
    vlanUnawareFxcAttachmentCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanUnawareFxcAttachmentCircuit.EntityData.Leafs["name"] = types.YLeaf{"Name", vlanUnawareFxcAttachmentCircuit.Name}
    return &(vlanUnawareFxcAttachmentCircuit.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns
// List of EVPN Services
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN FXC Service Configuration. The type is slice of
    // L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn.
    VlanUnawareFxcPseudowireEvpn []L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn
}

func (vlanUnawareFxcPseudowireEvpns *L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcPseudowireEvpns.EntityData.YFilter = vlanUnawareFxcPseudowireEvpns.YFilter
    vlanUnawareFxcPseudowireEvpns.EntityData.YangName = "vlan-unaware-fxc-pseudowire-evpns"
    vlanUnawareFxcPseudowireEvpns.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcPseudowireEvpns.EntityData.ParentYangName = "vlan-unaware-flexible-xconnect-service"
    vlanUnawareFxcPseudowireEvpns.EntityData.SegmentPath = "vlan-unaware-fxc-pseudowire-evpns"
    vlanUnawareFxcPseudowireEvpns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcPseudowireEvpns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcPseudowireEvpns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcPseudowireEvpns.EntityData.Children = make(map[string]types.YChild)
    vlanUnawareFxcPseudowireEvpns.EntityData.Children["vlan-unaware-fxc-pseudowire-evpn"] = types.YChild{"VlanUnawareFxcPseudowireEvpn", nil}
    for i := range vlanUnawareFxcPseudowireEvpns.VlanUnawareFxcPseudowireEvpn {
        vlanUnawareFxcPseudowireEvpns.EntityData.Children[types.GetSegmentPath(&vlanUnawareFxcPseudowireEvpns.VlanUnawareFxcPseudowireEvpn[i])] = types.YChild{"VlanUnawareFxcPseudowireEvpn", &vlanUnawareFxcPseudowireEvpns.VlanUnawareFxcPseudowireEvpn[i]}
    }
    vlanUnawareFxcPseudowireEvpns.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlanUnawareFxcPseudowireEvpns.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn
// EVPN FXC Service Configuration
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // This attribute is a key. AC ID. The type is interface{} with range:
    // 1..4294967295.
    Acid interface{}
}

func (vlanUnawareFxcPseudowireEvpn *L2Vpn_Database_FlexibleXconnectServiceTable_VlanUnawareFlexibleXconnectServices_VlanUnawareFlexibleXconnectService_VlanUnawareFxcPseudowireEvpns_VlanUnawareFxcPseudowireEvpn) GetEntityData() *types.CommonEntityData {
    vlanUnawareFxcPseudowireEvpn.EntityData.YFilter = vlanUnawareFxcPseudowireEvpn.YFilter
    vlanUnawareFxcPseudowireEvpn.EntityData.YangName = "vlan-unaware-fxc-pseudowire-evpn"
    vlanUnawareFxcPseudowireEvpn.EntityData.BundleName = "cisco_ios_xr"
    vlanUnawareFxcPseudowireEvpn.EntityData.ParentYangName = "vlan-unaware-fxc-pseudowire-evpns"
    vlanUnawareFxcPseudowireEvpn.EntityData.SegmentPath = "vlan-unaware-fxc-pseudowire-evpn" + "[eviid='" + fmt.Sprintf("%v", vlanUnawareFxcPseudowireEvpn.Eviid) + "']" + "[acid='" + fmt.Sprintf("%v", vlanUnawareFxcPseudowireEvpn.Acid) + "']"
    vlanUnawareFxcPseudowireEvpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanUnawareFxcPseudowireEvpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanUnawareFxcPseudowireEvpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanUnawareFxcPseudowireEvpn.EntityData.Children = make(map[string]types.YChild)
    vlanUnawareFxcPseudowireEvpn.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanUnawareFxcPseudowireEvpn.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", vlanUnawareFxcPseudowireEvpn.Eviid}
    vlanUnawareFxcPseudowireEvpn.EntityData.Leafs["acid"] = types.YLeaf{"Acid", vlanUnawareFxcPseudowireEvpn.Acid}
    return &(vlanUnawareFxcPseudowireEvpn.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices
// List of Vlan-Aware Flexible XConnect Services
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flexible XConnect Service. The type is slice of
    // L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService.
    VlanAwareFlexibleXconnectService []L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService
}

func (vlanAwareFlexibleXconnectServices *L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices) GetEntityData() *types.CommonEntityData {
    vlanAwareFlexibleXconnectServices.EntityData.YFilter = vlanAwareFlexibleXconnectServices.YFilter
    vlanAwareFlexibleXconnectServices.EntityData.YangName = "vlan-aware-flexible-xconnect-services"
    vlanAwareFlexibleXconnectServices.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFlexibleXconnectServices.EntityData.ParentYangName = "flexible-xconnect-service-table"
    vlanAwareFlexibleXconnectServices.EntityData.SegmentPath = "vlan-aware-flexible-xconnect-services"
    vlanAwareFlexibleXconnectServices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFlexibleXconnectServices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFlexibleXconnectServices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFlexibleXconnectServices.EntityData.Children = make(map[string]types.YChild)
    vlanAwareFlexibleXconnectServices.EntityData.Children["vlan-aware-flexible-xconnect-service"] = types.YChild{"VlanAwareFlexibleXconnectService", nil}
    for i := range vlanAwareFlexibleXconnectServices.VlanAwareFlexibleXconnectService {
        vlanAwareFlexibleXconnectServices.EntityData.Children[types.GetSegmentPath(&vlanAwareFlexibleXconnectServices.VlanAwareFlexibleXconnectService[i])] = types.YChild{"VlanAwareFlexibleXconnectService", &vlanAwareFlexibleXconnectServices.VlanAwareFlexibleXconnectService[i]}
    }
    vlanAwareFlexibleXconnectServices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlanAwareFlexibleXconnectServices.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService
// Flexible XConnect Service
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet VPN ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

    // List of attachment circuits.
    VlanAwareFxcAttachmentCircuits L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits
}

func (vlanAwareFlexibleXconnectService *L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService) GetEntityData() *types.CommonEntityData {
    vlanAwareFlexibleXconnectService.EntityData.YFilter = vlanAwareFlexibleXconnectService.YFilter
    vlanAwareFlexibleXconnectService.EntityData.YangName = "vlan-aware-flexible-xconnect-service"
    vlanAwareFlexibleXconnectService.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFlexibleXconnectService.EntityData.ParentYangName = "vlan-aware-flexible-xconnect-services"
    vlanAwareFlexibleXconnectService.EntityData.SegmentPath = "vlan-aware-flexible-xconnect-service" + "[eviid='" + fmt.Sprintf("%v", vlanAwareFlexibleXconnectService.Eviid) + "']"
    vlanAwareFlexibleXconnectService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFlexibleXconnectService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFlexibleXconnectService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFlexibleXconnectService.EntityData.Children = make(map[string]types.YChild)
    vlanAwareFlexibleXconnectService.EntityData.Children["vlan-aware-fxc-attachment-circuits"] = types.YChild{"VlanAwareFxcAttachmentCircuits", &vlanAwareFlexibleXconnectService.VlanAwareFxcAttachmentCircuits}
    vlanAwareFlexibleXconnectService.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanAwareFlexibleXconnectService.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", vlanAwareFlexibleXconnectService.Eviid}
    return &(vlanAwareFlexibleXconnectService.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits
// List of attachment circuits
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit.
    VlanAwareFxcAttachmentCircuit []L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit
}

func (vlanAwareFxcAttachmentCircuits *L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits) GetEntityData() *types.CommonEntityData {
    vlanAwareFxcAttachmentCircuits.EntityData.YFilter = vlanAwareFxcAttachmentCircuits.YFilter
    vlanAwareFxcAttachmentCircuits.EntityData.YangName = "vlan-aware-fxc-attachment-circuits"
    vlanAwareFxcAttachmentCircuits.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFxcAttachmentCircuits.EntityData.ParentYangName = "vlan-aware-flexible-xconnect-service"
    vlanAwareFxcAttachmentCircuits.EntityData.SegmentPath = "vlan-aware-fxc-attachment-circuits"
    vlanAwareFxcAttachmentCircuits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFxcAttachmentCircuits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFxcAttachmentCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFxcAttachmentCircuits.EntityData.Children = make(map[string]types.YChild)
    vlanAwareFxcAttachmentCircuits.EntityData.Children["vlan-aware-fxc-attachment-circuit"] = types.YChild{"VlanAwareFxcAttachmentCircuit", nil}
    for i := range vlanAwareFxcAttachmentCircuits.VlanAwareFxcAttachmentCircuit {
        vlanAwareFxcAttachmentCircuits.EntityData.Children[types.GetSegmentPath(&vlanAwareFxcAttachmentCircuits.VlanAwareFxcAttachmentCircuit[i])] = types.YChild{"VlanAwareFxcAttachmentCircuit", &vlanAwareFxcAttachmentCircuits.VlanAwareFxcAttachmentCircuit[i]}
    }
    vlanAwareFxcAttachmentCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlanAwareFxcAttachmentCircuits.EntityData)
}

// L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit
// Attachment circuit interface
type L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: b'[a-zA-Z0-9./-]+'.
    Name interface{}
}

func (vlanAwareFxcAttachmentCircuit *L2Vpn_Database_FlexibleXconnectServiceTable_VlanAwareFlexibleXconnectServices_VlanAwareFlexibleXconnectService_VlanAwareFxcAttachmentCircuits_VlanAwareFxcAttachmentCircuit) GetEntityData() *types.CommonEntityData {
    vlanAwareFxcAttachmentCircuit.EntityData.YFilter = vlanAwareFxcAttachmentCircuit.YFilter
    vlanAwareFxcAttachmentCircuit.EntityData.YangName = "vlan-aware-fxc-attachment-circuit"
    vlanAwareFxcAttachmentCircuit.EntityData.BundleName = "cisco_ios_xr"
    vlanAwareFxcAttachmentCircuit.EntityData.ParentYangName = "vlan-aware-fxc-attachment-circuits"
    vlanAwareFxcAttachmentCircuit.EntityData.SegmentPath = "vlan-aware-fxc-attachment-circuit" + "[name='" + fmt.Sprintf("%v", vlanAwareFxcAttachmentCircuit.Name) + "']"
    vlanAwareFxcAttachmentCircuit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanAwareFxcAttachmentCircuit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanAwareFxcAttachmentCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanAwareFxcAttachmentCircuit.EntityData.Children = make(map[string]types.YChild)
    vlanAwareFxcAttachmentCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanAwareFxcAttachmentCircuit.EntityData.Leafs["name"] = types.YLeaf{"Name", vlanAwareFxcAttachmentCircuit.Name}
    return &(vlanAwareFxcAttachmentCircuit.EntityData)
}

// L2Vpn_Database_Redundancy
// Redundancy groups
type L2Vpn_Database_Redundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable redundancy groups. The type is interface{}.
    Enable interface{}

    // List of Inter-Chassis Communication Protocol redundancy groups.
    IccpRedundancyGroups L2Vpn_Database_Redundancy_IccpRedundancyGroups
}

func (redundancy *L2Vpn_Database_Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "database"
    redundancy.EntityData.SegmentPath = "redundancy"
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = make(map[string]types.YChild)
    redundancy.EntityData.Children["iccp-redundancy-groups"] = types.YChild{"IccpRedundancyGroups", &redundancy.IccpRedundancyGroups}
    redundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    redundancy.EntityData.Leafs["enable"] = types.YLeaf{"Enable", redundancy.Enable}
    return &(redundancy.EntityData)
}

// L2Vpn_Database_Redundancy_IccpRedundancyGroups
// List of Inter-Chassis Communication Protocol
// redundancy groups
type L2Vpn_Database_Redundancy_IccpRedundancyGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICCP Redundancy group. The type is slice of
    // L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup.
    IccpRedundancyGroup []L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup
}

func (iccpRedundancyGroups *L2Vpn_Database_Redundancy_IccpRedundancyGroups) GetEntityData() *types.CommonEntityData {
    iccpRedundancyGroups.EntityData.YFilter = iccpRedundancyGroups.YFilter
    iccpRedundancyGroups.EntityData.YangName = "iccp-redundancy-groups"
    iccpRedundancyGroups.EntityData.BundleName = "cisco_ios_xr"
    iccpRedundancyGroups.EntityData.ParentYangName = "redundancy"
    iccpRedundancyGroups.EntityData.SegmentPath = "iccp-redundancy-groups"
    iccpRedundancyGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpRedundancyGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpRedundancyGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpRedundancyGroups.EntityData.Children = make(map[string]types.YChild)
    iccpRedundancyGroups.EntityData.Children["iccp-redundancy-group"] = types.YChild{"IccpRedundancyGroup", nil}
    for i := range iccpRedundancyGroups.IccpRedundancyGroup {
        iccpRedundancyGroups.EntityData.Children[types.GetSegmentPath(&iccpRedundancyGroups.IccpRedundancyGroup[i])] = types.YChild{"IccpRedundancyGroup", &iccpRedundancyGroups.IccpRedundancyGroup[i]}
    }
    iccpRedundancyGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iccpRedundancyGroups.EntityData)
}

// L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup
// ICCP Redundancy group
type L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..4294967295.
    GroupId interface{}

    // ICCP-based service multi-homing node ID. The type is interface{} with
    // range: 0..254.
    MultiHomingNodeId interface{}

    // List of interfaces.
    IccpInterfaces L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces
}

func (iccpRedundancyGroup *L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup) GetEntityData() *types.CommonEntityData {
    iccpRedundancyGroup.EntityData.YFilter = iccpRedundancyGroup.YFilter
    iccpRedundancyGroup.EntityData.YangName = "iccp-redundancy-group"
    iccpRedundancyGroup.EntityData.BundleName = "cisco_ios_xr"
    iccpRedundancyGroup.EntityData.ParentYangName = "iccp-redundancy-groups"
    iccpRedundancyGroup.EntityData.SegmentPath = "iccp-redundancy-group" + "[group-id='" + fmt.Sprintf("%v", iccpRedundancyGroup.GroupId) + "']"
    iccpRedundancyGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpRedundancyGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpRedundancyGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpRedundancyGroup.EntityData.Children = make(map[string]types.YChild)
    iccpRedundancyGroup.EntityData.Children["iccp-interfaces"] = types.YChild{"IccpInterfaces", &iccpRedundancyGroup.IccpInterfaces}
    iccpRedundancyGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    iccpRedundancyGroup.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", iccpRedundancyGroup.GroupId}
    iccpRedundancyGroup.EntityData.Leafs["multi-homing-node-id"] = types.YLeaf{"MultiHomingNodeId", iccpRedundancyGroup.MultiHomingNodeId}
    return &(iccpRedundancyGroup.EntityData)
}

// L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces
// List of interfaces
type L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface.
    IccpInterface []L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface
}

func (iccpInterfaces *L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces) GetEntityData() *types.CommonEntityData {
    iccpInterfaces.EntityData.YFilter = iccpInterfaces.YFilter
    iccpInterfaces.EntityData.YangName = "iccp-interfaces"
    iccpInterfaces.EntityData.BundleName = "cisco_ios_xr"
    iccpInterfaces.EntityData.ParentYangName = "iccp-redundancy-group"
    iccpInterfaces.EntityData.SegmentPath = "iccp-interfaces"
    iccpInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpInterfaces.EntityData.Children = make(map[string]types.YChild)
    iccpInterfaces.EntityData.Children["iccp-interface"] = types.YChild{"IccpInterface", nil}
    for i := range iccpInterfaces.IccpInterface {
        iccpInterfaces.EntityData.Children[types.GetSegmentPath(&iccpInterfaces.IccpInterface[i])] = types.YChild{"IccpInterface", &iccpInterfaces.IccpInterface[i]}
    }
    iccpInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iccpInterfaces.EntityData)
}

// L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface
// Interface name
type L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (iccpInterface *L2Vpn_Database_Redundancy_IccpRedundancyGroups_IccpRedundancyGroup_IccpInterfaces_IccpInterface) GetEntityData() *types.CommonEntityData {
    iccpInterface.EntityData.YFilter = iccpInterface.YFilter
    iccpInterface.EntityData.YangName = "iccp-interface"
    iccpInterface.EntityData.BundleName = "cisco_ios_xr"
    iccpInterface.EntityData.ParentYangName = "iccp-interfaces"
    iccpInterface.EntityData.SegmentPath = "iccp-interface" + "[interface-name='" + fmt.Sprintf("%v", iccpInterface.InterfaceName) + "']"
    iccpInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iccpInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iccpInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iccpInterface.EntityData.Children = make(map[string]types.YChild)
    iccpInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    iccpInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", iccpInterface.InterfaceName}
    iccpInterface.EntityData.Leafs["secondary-vlan-range"] = types.YLeaf{"SecondaryVlanRange", iccpInterface.SecondaryVlanRange}
    iccpInterface.EntityData.Leafs["recovery-delay"] = types.YLeaf{"RecoveryDelay", iccpInterface.RecoveryDelay}
    iccpInterface.EntityData.Leafs["primary-vlan-range"] = types.YLeaf{"PrimaryVlanRange", iccpInterface.PrimaryVlanRange}
    iccpInterface.EntityData.Leafs["mac-flush-tcn"] = types.YLeaf{"MacFlushTcn", iccpInterface.MacFlushTcn}
    return &(iccpInterface.EntityData)
}

// L2Vpn_Pbb
// L2VPN PBB Global
type L2Vpn_Pbb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backbone Source MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    BackboneSourceMac interface{}
}

func (pbb *L2Vpn_Pbb) GetEntityData() *types.CommonEntityData {
    pbb.EntityData.YFilter = pbb.YFilter
    pbb.EntityData.YangName = "pbb"
    pbb.EntityData.BundleName = "cisco_ios_xr"
    pbb.EntityData.ParentYangName = "l2vpn"
    pbb.EntityData.SegmentPath = "pbb"
    pbb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbb.EntityData.Children = make(map[string]types.YChild)
    pbb.EntityData.Leafs = make(map[string]types.YLeaf)
    pbb.EntityData.Leafs["backbone-source-mac"] = types.YLeaf{"BackboneSourceMac", pbb.BackboneSourceMac}
    return &(pbb.EntityData)
}

// L2Vpn_AutoDiscovery
// Global auto-discovery attributes
type L2Vpn_AutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global bgp signaling attributes.
    BgpSignaling L2Vpn_AutoDiscovery_BgpSignaling
}

func (autoDiscovery *L2Vpn_AutoDiscovery) GetEntityData() *types.CommonEntityData {
    autoDiscovery.EntityData.YFilter = autoDiscovery.YFilter
    autoDiscovery.EntityData.YangName = "auto-discovery"
    autoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    autoDiscovery.EntityData.ParentYangName = "l2vpn"
    autoDiscovery.EntityData.SegmentPath = "auto-discovery"
    autoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoDiscovery.EntityData.Children = make(map[string]types.YChild)
    autoDiscovery.EntityData.Children["bgp-signaling"] = types.YChild{"BgpSignaling", &autoDiscovery.BgpSignaling}
    autoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(autoDiscovery.EntityData)
}

// L2Vpn_AutoDiscovery_BgpSignaling
// Global bgp signaling attributes
type L2Vpn_AutoDiscovery_BgpSignaling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore MTU mismatch for auto-discovered pseudowires. The type is
    // interface{}.
    MtuMismatchIgnore interface{}
}

func (bgpSignaling *L2Vpn_AutoDiscovery_BgpSignaling) GetEntityData() *types.CommonEntityData {
    bgpSignaling.EntityData.YFilter = bgpSignaling.YFilter
    bgpSignaling.EntityData.YangName = "bgp-signaling"
    bgpSignaling.EntityData.BundleName = "cisco_ios_xr"
    bgpSignaling.EntityData.ParentYangName = "auto-discovery"
    bgpSignaling.EntityData.SegmentPath = "bgp-signaling"
    bgpSignaling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpSignaling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpSignaling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpSignaling.EntityData.Children = make(map[string]types.YChild)
    bgpSignaling.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpSignaling.EntityData.Leafs["mtu-mismatch-ignore"] = types.YLeaf{"MtuMismatchIgnore", bgpSignaling.MtuMismatchIgnore}
    return &(bgpSignaling.EntityData)
}

// L2Vpn_Utility
// L2VPN utilities
type L2Vpn_Utility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN logging utility.
    Logging L2Vpn_Utility_Logging
}

func (utility *L2Vpn_Utility) GetEntityData() *types.CommonEntityData {
    utility.EntityData.YFilter = utility.YFilter
    utility.EntityData.YangName = "utility"
    utility.EntityData.BundleName = "cisco_ios_xr"
    utility.EntityData.ParentYangName = "l2vpn"
    utility.EntityData.SegmentPath = "utility"
    utility.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utility.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utility.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utility.EntityData.Children = make(map[string]types.YChild)
    utility.EntityData.Children["logging"] = types.YChild{"Logging", &utility.Logging}
    utility.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(utility.EntityData)
}

// L2Vpn_Utility_Logging
// L2VPN logging utility
type L2Vpn_Utility_Logging struct {
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

func (logging *L2Vpn_Utility_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "utility"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    logging.EntityData.Leafs["bridge-domain-state-change"] = types.YLeaf{"BridgeDomainStateChange", logging.BridgeDomainStateChange}
    logging.EntityData.Leafs["pseudowire-state-change"] = types.YLeaf{"PseudowireStateChange", logging.PseudowireStateChange}
    logging.EntityData.Leafs["vfi"] = types.YLeaf{"Vfi", logging.Vfi}
    logging.EntityData.Leafs["nsr-state-change"] = types.YLeaf{"NsrStateChange", logging.NsrStateChange}
    logging.EntityData.Leafs["pwhe-replication-state-change"] = types.YLeaf{"PwheReplicationStateChange", logging.PwheReplicationStateChange}
    return &(logging.EntityData)
}

// L2Vpn_Snmp
// SNMP related configuration
type L2Vpn_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MIB related configuration.
    Mib L2Vpn_Snmp_Mib
}

func (snmp *L2Vpn_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "l2vpn"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = make(map[string]types.YChild)
    snmp.EntityData.Children["mib"] = types.YChild{"Mib", &snmp.Mib}
    snmp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snmp.EntityData)
}

// L2Vpn_Snmp_Mib
// MIB related configuration
type L2Vpn_Snmp_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface related configuration for MIB.
    MibInterface L2Vpn_Snmp_Mib_MibInterface

    // Pseudowire related configuration for MIB.
    MibPseudowire L2Vpn_Snmp_Mib_MibPseudowire
}

func (mib *L2Vpn_Snmp_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "snmp"
    mib.EntityData.SegmentPath = "mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = make(map[string]types.YChild)
    mib.EntityData.Children["mib-interface"] = types.YChild{"MibInterface", &mib.MibInterface}
    mib.EntityData.Children["mib-pseudowire"] = types.YChild{"MibPseudowire", &mib.MibPseudowire}
    mib.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mib.EntityData)
}

// L2Vpn_Snmp_Mib_MibInterface
// Interface related configuration for MIB
type L2Vpn_Snmp_Mib_MibInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MIB interface name output format.
    Format L2Vpn_Snmp_Mib_MibInterface_Format
}

func (mibInterface *L2Vpn_Snmp_Mib_MibInterface) GetEntityData() *types.CommonEntityData {
    mibInterface.EntityData.YFilter = mibInterface.YFilter
    mibInterface.EntityData.YangName = "mib-interface"
    mibInterface.EntityData.BundleName = "cisco_ios_xr"
    mibInterface.EntityData.ParentYangName = "mib"
    mibInterface.EntityData.SegmentPath = "mib-interface"
    mibInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibInterface.EntityData.Children = make(map[string]types.YChild)
    mibInterface.EntityData.Children["format"] = types.YChild{"Format", &mibInterface.Format}
    mibInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mibInterface.EntityData)
}

// L2Vpn_Snmp_Mib_MibInterface_Format
// MIB interface name output format
type L2Vpn_Snmp_Mib_MibInterface_Format struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set MIB interface name output in slash format (/). The type is interface{}.
    ExternalInterfaceFormat interface{}
}

func (format *L2Vpn_Snmp_Mib_MibInterface_Format) GetEntityData() *types.CommonEntityData {
    format.EntityData.YFilter = format.YFilter
    format.EntityData.YangName = "format"
    format.EntityData.BundleName = "cisco_ios_xr"
    format.EntityData.ParentYangName = "mib-interface"
    format.EntityData.SegmentPath = "format"
    format.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    format.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    format.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    format.EntityData.Children = make(map[string]types.YChild)
    format.EntityData.Leafs = make(map[string]types.YLeaf)
    format.EntityData.Leafs["external-interface-format"] = types.YLeaf{"ExternalInterfaceFormat", format.ExternalInterfaceFormat}
    return &(format.EntityData)
}

// L2Vpn_Snmp_Mib_MibPseudowire
// Pseudowire related configuration for MIB
type L2Vpn_Snmp_Mib_MibPseudowire struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pseudowire statistics in MIB output. The type is interface{}.
    Statistics interface{}
}

func (mibPseudowire *L2Vpn_Snmp_Mib_MibPseudowire) GetEntityData() *types.CommonEntityData {
    mibPseudowire.EntityData.YFilter = mibPseudowire.YFilter
    mibPseudowire.EntityData.YangName = "mib-pseudowire"
    mibPseudowire.EntityData.BundleName = "cisco_ios_xr"
    mibPseudowire.EntityData.ParentYangName = "mib"
    mibPseudowire.EntityData.SegmentPath = "mib-pseudowire"
    mibPseudowire.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibPseudowire.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibPseudowire.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibPseudowire.EntityData.Children = make(map[string]types.YChild)
    mibPseudowire.EntityData.Leafs = make(map[string]types.YLeaf)
    mibPseudowire.EntityData.Leafs["statistics"] = types.YLeaf{"Statistics", mibPseudowire.Statistics}
    return &(mibPseudowire.EntityData)
}

// GenericInterfaceLists
// generic interface lists
type GenericInterfaceLists struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic interface list. The type is slice of
    // GenericInterfaceLists_GenericInterfaceList.
    GenericInterfaceList []GenericInterfaceLists_GenericInterfaceList
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

    genericInterfaceLists.EntityData.Children = make(map[string]types.YChild)
    genericInterfaceLists.EntityData.Children["generic-interface-list"] = types.YChild{"GenericInterfaceList", nil}
    for i := range genericInterfaceLists.GenericInterfaceList {
        genericInterfaceLists.EntityData.Children[types.GetSegmentPath(&genericInterfaceLists.GenericInterfaceList[i])] = types.YChild{"GenericInterfaceList", &genericInterfaceLists.GenericInterfaceList[i]}
    }
    genericInterfaceLists.EntityData.Leafs = make(map[string]types.YLeaf)
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
    genericInterfaceList.EntityData.SegmentPath = "generic-interface-list" + "[generic-interface-list-name='" + fmt.Sprintf("%v", genericInterfaceList.GenericInterfaceListName) + "']"
    genericInterfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericInterfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericInterfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericInterfaceList.EntityData.Children = make(map[string]types.YChild)
    genericInterfaceList.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &genericInterfaceList.Interfaces}
    genericInterfaceList.EntityData.Leafs = make(map[string]types.YLeaf)
    genericInterfaceList.EntityData.Leafs["generic-interface-list-name"] = types.YLeaf{"GenericInterfaceListName", genericInterfaceList.GenericInterfaceListName}
    genericInterfaceList.EntityData.Leafs["enable"] = types.YLeaf{"Enable", genericInterfaceList.Enable}
    return &(genericInterfaceList.EntityData)
}

// GenericInterfaceLists_GenericInterfaceList_Interfaces
// Interface table
type GenericInterfaceLists_GenericInterfaceList_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is slice of
    // GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface_.
    Interface_ []GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface
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

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface
// Interface
type GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Enable interface. The type is interface{}.
    Enable interface{}
}

func (self *GenericInterfaceLists_GenericInterfaceList_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["enable"] = types.YLeaf{"Enable", self.Enable}
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

    evpn.EntityData.Children = make(map[string]types.YChild)
    evpn.EntityData.Children["evpn-tables"] = types.YChild{"EvpnTables", &evpn.EvpnTables}
    evpn.EntityData.Leafs = make(map[string]types.YLeaf)
    evpn.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpn.Enable}
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
    // string with pattern: b'[a-zA-Z0-9./-]+'.
    EvpnSourceInterface interface{}

    // Cost-in node after given time (seconds) on startup timer. The type is
    // interface{} with range: 30..86400. Units are second.
    EvpnCostInStartup interface{}

    // Enter EVPN timers configuration submode.
    EvpnTimers Evpn_EvpnTables_EvpnTimers

    // EVPN MAC Configuration.
    Evpnmac Evpn_EvpnTables_Evpnmac

    // Enter EVPN Instance configuration submode.
    Evpnevis Evpn_EvpnTables_Evpnevis

    // Virtual Access VFI interfaces.
    EvpnVirtualAccessVfis Evpn_EvpnTables_EvpnVirtualAccessVfis

    // Enter EVPN Loadbalancing configuration submode.
    EvpnLoadBalancing Evpn_EvpnTables_EvpnLoadBalancing

    // Enable Autodiscovery BGP in EVPN.
    EvpnBgpAutoDiscovery Evpn_EvpnTables_EvpnBgpAutoDiscovery

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

    evpnTables.EntityData.Children = make(map[string]types.YChild)
    evpnTables.EntityData.Children["evpn-timers"] = types.YChild{"EvpnTimers", &evpnTables.EvpnTimers}
    evpnTables.EntityData.Children["evpnmac"] = types.YChild{"Evpnmac", &evpnTables.Evpnmac}
    evpnTables.EntityData.Children["evpnevis"] = types.YChild{"Evpnevis", &evpnTables.Evpnevis}
    evpnTables.EntityData.Children["evpn-virtual-access-vfis"] = types.YChild{"EvpnVirtualAccessVfis", &evpnTables.EvpnVirtualAccessVfis}
    evpnTables.EntityData.Children["evpn-load-balancing"] = types.YChild{"EvpnLoadBalancing", &evpnTables.EvpnLoadBalancing}
    evpnTables.EntityData.Children["evpn-bgp-auto-discovery"] = types.YChild{"EvpnBgpAutoDiscovery", &evpnTables.EvpnBgpAutoDiscovery}
    evpnTables.EntityData.Children["evpn-instances"] = types.YChild{"EvpnInstances", &evpnTables.EvpnInstances}
    evpnTables.EntityData.Children["evpn-logging"] = types.YChild{"EvpnLogging", &evpnTables.EvpnLogging}
    evpnTables.EntityData.Children["evpn-interfaces"] = types.YChild{"EvpnInterfaces", &evpnTables.EvpnInterfaces}
    evpnTables.EntityData.Children["evpn-virtual-access-pws"] = types.YChild{"EvpnVirtualAccessPws", &evpnTables.EvpnVirtualAccessPws}
    evpnTables.EntityData.Children["evpn-ethernet-segment"] = types.YChild{"EvpnEthernetSegment", &evpnTables.EvpnEthernetSegment}
    evpnTables.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnTables.EntityData.Leafs["evi-cost-out"] = types.YLeaf{"EviCostOut", evpnTables.EviCostOut}
    evpnTables.EntityData.Leafs["evpn-source-interface"] = types.YLeaf{"EvpnSourceInterface", evpnTables.EvpnSourceInterface}
    evpnTables.EntityData.Leafs["evpn-cost-in-startup"] = types.YLeaf{"EvpnCostInStartup", evpnTables.EvpnCostInStartup}
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

    evpnTimers.EntityData.Children = make(map[string]types.YChild)
    evpnTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnTimers.EntityData.Leafs["evpn-carving"] = types.YLeaf{"EvpnCarving", evpnTimers.EvpnCarving}
    evpnTimers.EntityData.Leafs["evpn-recovery"] = types.YLeaf{"EvpnRecovery", evpnTimers.EvpnRecovery}
    evpnTimers.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnTimers.Enable}
    evpnTimers.EntityData.Leafs["evpn-peering"] = types.YLeaf{"EvpnPeering", evpnTimers.EvpnPeering}
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

    evpnmac.EntityData.Children = make(map[string]types.YChild)
    evpnmac.EntityData.Children["evpnmac-secure"] = types.YChild{"EvpnmacSecure", &evpnmac.EvpnmacSecure}
    evpnmac.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnmac.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnmac.Enable}
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

    evpnmacSecure.EntityData.Children = make(map[string]types.YChild)
    evpnmacSecure.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnmacSecure.EntityData.Leafs["evpnmac-secure-freeze-time"] = types.YLeaf{"EvpnmacSecureFreezeTime", evpnmacSecure.EvpnmacSecureFreezeTime}
    evpnmacSecure.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnmacSecure.Enable}
    evpnmacSecure.EntityData.Leafs["evpnmac-secure-retry-count"] = types.YLeaf{"EvpnmacSecureRetryCount", evpnmacSecure.EvpnmacSecureRetryCount}
    evpnmacSecure.EntityData.Leafs["evpnmac-secure-move-count"] = types.YLeaf{"EvpnmacSecureMoveCount", evpnmacSecure.EvpnmacSecureMoveCount}
    evpnmacSecure.EntityData.Leafs["evpnmac-secure-move-interval"] = types.YLeaf{"EvpnmacSecureMoveInterval", evpnmacSecure.EvpnmacSecureMoveInterval}
    return &(evpnmacSecure.EntityData)
}

// Evpn_EvpnTables_Evpnevis
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_Evpnevis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter EVPN Instance configuration submode. The type is slice of
    // Evpn_EvpnTables_Evpnevis_Evpnevi.
    Evpnevi []Evpn_EvpnTables_Evpnevis_Evpnevi
}

func (evpnevis *Evpn_EvpnTables_Evpnevis) GetEntityData() *types.CommonEntityData {
    evpnevis.EntityData.YFilter = evpnevis.YFilter
    evpnevis.EntityData.YangName = "evpnevis"
    evpnevis.EntityData.BundleName = "cisco_ios_xr"
    evpnevis.EntityData.ParentYangName = "evpn-tables"
    evpnevis.EntityData.SegmentPath = "evpnevis"
    evpnevis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnevis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnevis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnevis.EntityData.Children = make(map[string]types.YChild)
    evpnevis.EntityData.Children["evpnevi"] = types.YChild{"Evpnevi", nil}
    for i := range evpnevis.Evpnevi {
        evpnevis.EntityData.Children[types.GetSegmentPath(&evpnevis.Evpnevi[i])] = types.YChild{"Evpnevi", &evpnevis.Evpnevi[i]}
    }
    evpnevis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evpnevis.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_Evpnevis_Evpnevi struct {
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
    EvpneviDescription interface{}

    // Disable ECMP on the EVI. The type is interface{}.
    EviEcmpDisable interface{}

    // Disable Unknown Unicast Flooding on this EVI. The type is interface{}.
    EviUnknownUnicastFloodingDisable interface{}

    // CW disable for EVPN EVI. The type is interface{}.
    EvpnEviCwDisable interface{}

    // Enter Loadbalancing configuration submode.
    EviLoadBalancing Evpn_EvpnTables_Evpnevis_Evpnevi_EviLoadBalancing

    // Enable Autodiscovery BGP in EVPN Instance.
    EvpnevIbgpAutoDiscovery Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery

    // Enter Advertise local MAC-only routes configuration submode.
    EviAdvertiseMac Evpn_EvpnTables_Evpnevis_Evpnevi_EviAdvertiseMac
}

func (evpnevi *Evpn_EvpnTables_Evpnevis_Evpnevi) GetEntityData() *types.CommonEntityData {
    evpnevi.EntityData.YFilter = evpnevi.YFilter
    evpnevi.EntityData.YangName = "evpnevi"
    evpnevi.EntityData.BundleName = "cisco_ios_xr"
    evpnevi.EntityData.ParentYangName = "evpnevis"
    evpnevi.EntityData.SegmentPath = "evpnevi" + "[eviid='" + fmt.Sprintf("%v", evpnevi.Eviid) + "']"
    evpnevi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnevi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnevi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnevi.EntityData.Children = make(map[string]types.YChild)
    evpnevi.EntityData.Children["evi-load-balancing"] = types.YChild{"EviLoadBalancing", &evpnevi.EviLoadBalancing}
    evpnevi.EntityData.Children["evpnev-ibgp-auto-discovery"] = types.YChild{"EvpnevIbgpAutoDiscovery", &evpnevi.EvpnevIbgpAutoDiscovery}
    evpnevi.EntityData.Children["evi-advertise-mac"] = types.YChild{"EviAdvertiseMac", &evpnevi.EviAdvertiseMac}
    evpnevi.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnevi.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", evpnevi.Eviid}
    evpnevi.EntityData.Leafs["evi-reorig-disable"] = types.YLeaf{"EviReorigDisable", evpnevi.EviReorigDisable}
    evpnevi.EntityData.Leafs["evi-advertise-mac-deprecated"] = types.YLeaf{"EviAdvertiseMacDeprecated", evpnevi.EviAdvertiseMacDeprecated}
    evpnevi.EntityData.Leafs["evpnevi-description"] = types.YLeaf{"EvpneviDescription", evpnevi.EvpneviDescription}
    evpnevi.EntityData.Leafs["evi-ecmp-disable"] = types.YLeaf{"EviEcmpDisable", evpnevi.EviEcmpDisable}
    evpnevi.EntityData.Leafs["evi-unknown-unicast-flooding-disable"] = types.YLeaf{"EviUnknownUnicastFloodingDisable", evpnevi.EviUnknownUnicastFloodingDisable}
    evpnevi.EntityData.Leafs["evpn-evi-cw-disable"] = types.YLeaf{"EvpnEviCwDisable", evpnevi.EvpnEviCwDisable}
    return &(evpnevi.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EviLoadBalancing
// Enter Loadbalancing configuration submode
type Evpn_EvpnTables_Evpnevis_Evpnevi_EviLoadBalancing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Loadbalancing. The type is interface{}.
    Enable interface{}

    // Enable Static Flow Label based load balancing. The type is interface{}.
    EviStaticFlowLabel interface{}
}

func (eviLoadBalancing *Evpn_EvpnTables_Evpnevis_Evpnevi_EviLoadBalancing) GetEntityData() *types.CommonEntityData {
    eviLoadBalancing.EntityData.YFilter = eviLoadBalancing.YFilter
    eviLoadBalancing.EntityData.YangName = "evi-load-balancing"
    eviLoadBalancing.EntityData.BundleName = "cisco_ios_xr"
    eviLoadBalancing.EntityData.ParentYangName = "evpnevi"
    eviLoadBalancing.EntityData.SegmentPath = "evi-load-balancing"
    eviLoadBalancing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviLoadBalancing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviLoadBalancing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviLoadBalancing.EntityData.Children = make(map[string]types.YChild)
    eviLoadBalancing.EntityData.Leafs = make(map[string]types.YLeaf)
    eviLoadBalancing.EntityData.Leafs["enable"] = types.YLeaf{"Enable", eviLoadBalancing.Enable}
    eviLoadBalancing.EntityData.Leafs["evi-static-flow-label"] = types.YLeaf{"EviStaticFlowLabel", eviLoadBalancing.EviStaticFlowLabel}
    return &(eviLoadBalancing.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery
// Enable Autodiscovery BGP in EVPN Instance
type Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Autodiscovery BGP. The type is interface{}.
    Enable interface{}

    // Table Policy for installation of forwarding data to L2FIB. The type is
    // string.
    TablePolicy interface{}

    // Route Distinguisher.
    EvpnRouteDistinguisher Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteDistinguisher

    // Route Target.
    EvpnRouteTargets Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets
}

func (evpnevIbgpAutoDiscovery *Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery) GetEntityData() *types.CommonEntityData {
    evpnevIbgpAutoDiscovery.EntityData.YFilter = evpnevIbgpAutoDiscovery.YFilter
    evpnevIbgpAutoDiscovery.EntityData.YangName = "evpnev-ibgp-auto-discovery"
    evpnevIbgpAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    evpnevIbgpAutoDiscovery.EntityData.ParentYangName = "evpnevi"
    evpnevIbgpAutoDiscovery.EntityData.SegmentPath = "evpnev-ibgp-auto-discovery"
    evpnevIbgpAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnevIbgpAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnevIbgpAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnevIbgpAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    evpnevIbgpAutoDiscovery.EntityData.Children["evpn-route-distinguisher"] = types.YChild{"EvpnRouteDistinguisher", &evpnevIbgpAutoDiscovery.EvpnRouteDistinguisher}
    evpnevIbgpAutoDiscovery.EntityData.Children["evpn-route-targets"] = types.YChild{"EvpnRouteTargets", &evpnevIbgpAutoDiscovery.EvpnRouteTargets}
    evpnevIbgpAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnevIbgpAutoDiscovery.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnevIbgpAutoDiscovery.Enable}
    evpnevIbgpAutoDiscovery.EntityData.Leafs["table-policy"] = types.YLeaf{"TablePolicy", evpnevIbgpAutoDiscovery.TablePolicy}
    return &(evpnevIbgpAutoDiscovery.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteDistinguisher
// Route Distinguisher
type Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type_ interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Addr index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}
}

func (evpnRouteDistinguisher *Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteDistinguisher) GetEntityData() *types.CommonEntityData {
    evpnRouteDistinguisher.EntityData.YFilter = evpnRouteDistinguisher.YFilter
    evpnRouteDistinguisher.EntityData.YangName = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteDistinguisher.EntityData.ParentYangName = "evpnev-ibgp-auto-discovery"
    evpnRouteDistinguisher.EntityData.SegmentPath = "evpn-route-distinguisher"
    evpnRouteDistinguisher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteDistinguisher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteDistinguisher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteDistinguisher.EntityData.Children = make(map[string]types.YChild)
    evpnRouteDistinguisher.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteDistinguisher.EntityData.Leafs["type"] = types.YLeaf{"Type_", evpnRouteDistinguisher.Type_}
    evpnRouteDistinguisher.EntityData.Leafs["as"] = types.YLeaf{"As", evpnRouteDistinguisher.As}
    evpnRouteDistinguisher.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex}
    evpnRouteDistinguisher.EntityData.Leafs["address"] = types.YLeaf{"Address", evpnRouteDistinguisher.Address}
    evpnRouteDistinguisher.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex}
    return &(evpnRouteDistinguisher.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets
// Route Target
type Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs.
    EvpnRouteTargetAs []Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone.
    EvpnRouteTargetNone []Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address.
    EvpnRouteTargetIpv4Address []Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
}

func (evpnRouteTargets *Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets) GetEntityData() *types.CommonEntityData {
    evpnRouteTargets.EntityData.YFilter = evpnRouteTargets.YFilter
    evpnRouteTargets.EntityData.YangName = "evpn-route-targets"
    evpnRouteTargets.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargets.EntityData.ParentYangName = "evpnev-ibgp-auto-discovery"
    evpnRouteTargets.EntityData.SegmentPath = "evpn-route-targets"
    evpnRouteTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargets.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargets.EntityData.Children["evpn-route-target-as"] = types.YChild{"EvpnRouteTargetAs", nil}
    for i := range evpnRouteTargets.EvpnRouteTargetAs {
        evpnRouteTargets.EntityData.Children[types.GetSegmentPath(&evpnRouteTargets.EvpnRouteTargetAs[i])] = types.YChild{"EvpnRouteTargetAs", &evpnRouteTargets.EvpnRouteTargetAs[i]}
    }
    evpnRouteTargets.EntityData.Children["evpn-route-target-none"] = types.YChild{"EvpnRouteTargetNone", nil}
    for i := range evpnRouteTargets.EvpnRouteTargetNone {
        evpnRouteTargets.EntityData.Children[types.GetSegmentPath(&evpnRouteTargets.EvpnRouteTargetNone[i])] = types.YChild{"EvpnRouteTargetNone", &evpnRouteTargets.EvpnRouteTargetNone[i]}
    }
    evpnRouteTargets.EntityData.Children["evpn-route-target-ipv4-address"] = types.YChild{"EvpnRouteTargetIpv4Address", nil}
    for i := range evpnRouteTargets.EvpnRouteTargetIpv4Address {
        evpnRouteTargets.EntityData.Children[types.GetSegmentPath(&evpnRouteTargets.EvpnRouteTargetIpv4Address[i])] = types.YChild{"EvpnRouteTargetIpv4Address", &evpnRouteTargets.EvpnRouteTargetIpv4Address[i]}
    }
    evpnRouteTargets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evpnRouteTargets.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs
// Name of the Route Target
type Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs struct {
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

func (evpnRouteTargetAs *Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetAs.EntityData.YFilter = evpnRouteTargetAs.YFilter
    evpnRouteTargetAs.EntityData.YangName = "evpn-route-target-as"
    evpnRouteTargetAs.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetAs.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetAs.EntityData.SegmentPath = "evpn-route-target-as" + "[format='" + fmt.Sprintf("%v", evpnRouteTargetAs.Format) + "']" + "[role='" + fmt.Sprintf("%v", evpnRouteTargetAs.Role) + "']" + "[as='" + fmt.Sprintf("%v", evpnRouteTargetAs.As) + "']" + "[as-index='" + fmt.Sprintf("%v", evpnRouteTargetAs.AsIndex) + "']" + "[stitching='" + fmt.Sprintf("%v", evpnRouteTargetAs.Stitching) + "']"
    evpnRouteTargetAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetAs.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargetAs.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteTargetAs.EntityData.Leafs["format"] = types.YLeaf{"Format", evpnRouteTargetAs.Format}
    evpnRouteTargetAs.EntityData.Leafs["role"] = types.YLeaf{"Role", evpnRouteTargetAs.Role}
    evpnRouteTargetAs.EntityData.Leafs["as"] = types.YLeaf{"As", evpnRouteTargetAs.As}
    evpnRouteTargetAs.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", evpnRouteTargetAs.AsIndex}
    evpnRouteTargetAs.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", evpnRouteTargetAs.Stitching}
    return &(evpnRouteTargetAs.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone
// Name of the Route Target
type Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone struct {
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

func (evpnRouteTargetNone *Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetNone.EntityData.YFilter = evpnRouteTargetNone.YFilter
    evpnRouteTargetNone.EntityData.YangName = "evpn-route-target-none"
    evpnRouteTargetNone.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetNone.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetNone.EntityData.SegmentPath = "evpn-route-target-none" + "[format='" + fmt.Sprintf("%v", evpnRouteTargetNone.Format) + "']" + "[role='" + fmt.Sprintf("%v", evpnRouteTargetNone.Role) + "']" + "[stitching='" + fmt.Sprintf("%v", evpnRouteTargetNone.Stitching) + "']"
    evpnRouteTargetNone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetNone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetNone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetNone.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargetNone.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteTargetNone.EntityData.Leafs["format"] = types.YLeaf{"Format", evpnRouteTargetNone.Format}
    evpnRouteTargetNone.EntityData.Leafs["role"] = types.YLeaf{"Role", evpnRouteTargetNone.Role}
    evpnRouteTargetNone.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", evpnRouteTargetNone.Stitching}
    return &(evpnRouteTargetNone.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
// Name of the Route Target
type Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Format of the route target. The type is
    // BgpRouteTargetFormat.
    Format interface{}

    // This attribute is a key. Role of the router target type. The type is
    // BgpRouteTargetRole.
    Role interface{}

    // This attribute is a key. IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Addr index. The type is interface{} with range:
    // 0..65535.
    AddrIndex interface{}

    // This attribute is a key. whether RT is Stitching RT. The type is
    // BgpRouteTarget.
    Stitching interface{}
}

func (evpnRouteTargetIpv4Address *Evpn_EvpnTables_Evpnevis_Evpnevi_EvpnevIbgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address) GetEntityData() *types.CommonEntityData {
    evpnRouteTargetIpv4Address.EntityData.YFilter = evpnRouteTargetIpv4Address.YFilter
    evpnRouteTargetIpv4Address.EntityData.YangName = "evpn-route-target-ipv4-address"
    evpnRouteTargetIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    evpnRouteTargetIpv4Address.EntityData.ParentYangName = "evpn-route-targets"
    evpnRouteTargetIpv4Address.EntityData.SegmentPath = "evpn-route-target-ipv4-address" + "[format='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Format) + "']" + "[role='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Role) + "']" + "[address='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Address) + "']" + "[addr-index='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.AddrIndex) + "']" + "[stitching='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Stitching) + "']"
    evpnRouteTargetIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetIpv4Address.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargetIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteTargetIpv4Address.EntityData.Leafs["format"] = types.YLeaf{"Format", evpnRouteTargetIpv4Address.Format}
    evpnRouteTargetIpv4Address.EntityData.Leafs["role"] = types.YLeaf{"Role", evpnRouteTargetIpv4Address.Role}
    evpnRouteTargetIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", evpnRouteTargetIpv4Address.Address}
    evpnRouteTargetIpv4Address.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", evpnRouteTargetIpv4Address.AddrIndex}
    evpnRouteTargetIpv4Address.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", evpnRouteTargetIpv4Address.Stitching}
    return &(evpnRouteTargetIpv4Address.EntityData)
}

// Evpn_EvpnTables_Evpnevis_Evpnevi_EviAdvertiseMac
// Enter Advertise local MAC-only routes
// configuration submode
type Evpn_EvpnTables_Evpnevis_Evpnevi_EviAdvertiseMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Advertise local MAC-only routes. The type is interface{}.
    Enable interface{}

    // Advertise local MAC-only and BVI MAC routes. The type is interface{}.
    EviAdvertiseMacBvi interface{}
}

func (eviAdvertiseMac *Evpn_EvpnTables_Evpnevis_Evpnevi_EviAdvertiseMac) GetEntityData() *types.CommonEntityData {
    eviAdvertiseMac.EntityData.YFilter = eviAdvertiseMac.YFilter
    eviAdvertiseMac.EntityData.YangName = "evi-advertise-mac"
    eviAdvertiseMac.EntityData.BundleName = "cisco_ios_xr"
    eviAdvertiseMac.EntityData.ParentYangName = "evpnevi"
    eviAdvertiseMac.EntityData.SegmentPath = "evi-advertise-mac"
    eviAdvertiseMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviAdvertiseMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviAdvertiseMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviAdvertiseMac.EntityData.Children = make(map[string]types.YChild)
    eviAdvertiseMac.EntityData.Leafs = make(map[string]types.YLeaf)
    eviAdvertiseMac.EntityData.Leafs["enable"] = types.YLeaf{"Enable", eviAdvertiseMac.Enable}
    eviAdvertiseMac.EntityData.Leafs["evi-advertise-mac-bvi"] = types.YLeaf{"EviAdvertiseMacBvi", eviAdvertiseMac.EviAdvertiseMacBvi}
    return &(eviAdvertiseMac.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis
// Virtual Access VFI interfaces
type Evpn_EvpnTables_EvpnVirtualAccessVfis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual Access VFI. The type is slice of
    // Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi.
    EvpnVirtualAccessVfi []Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi
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

    evpnVirtualAccessVfis.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualAccessVfis.EntityData.Children["evpn-virtual-access-vfi"] = types.YChild{"EvpnVirtualAccessVfi", nil}
    for i := range evpnVirtualAccessVfis.EvpnVirtualAccessVfi {
        evpnVirtualAccessVfis.EntityData.Children[types.GetSegmentPath(&evpnVirtualAccessVfis.EvpnVirtualAccessVfi[i])] = types.YChild{"EvpnVirtualAccessVfi", &evpnVirtualAccessVfis.EvpnVirtualAccessVfi[i]}
    }
    evpnVirtualAccessVfis.EntityData.Leafs = make(map[string]types.YLeaf)
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
    evpnVirtualAccessVfi.EntityData.SegmentPath = "evpn-virtual-access-vfi" + "[name='" + fmt.Sprintf("%v", evpnVirtualAccessVfi.Name) + "']"
    evpnVirtualAccessVfi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessVfi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessVfi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessVfi.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualAccessVfi.EntityData.Children["evpn-virtual-access-vfi-timers"] = types.YChild{"EvpnVirtualAccessVfiTimers", &evpnVirtualAccessVfi.EvpnVirtualAccessVfiTimers}
    evpnVirtualAccessVfi.EntityData.Children["evpn-virtual-ethernet-segment"] = types.YChild{"EvpnVirtualEthernetSegment", &evpnVirtualAccessVfi.EvpnVirtualEthernetSegment}
    evpnVirtualAccessVfi.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnVirtualAccessVfi.EntityData.Leafs["name"] = types.YLeaf{"Name", evpnVirtualAccessVfi.Name}
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

    evpnVirtualAccessVfiTimers.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualAccessVfiTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnVirtualAccessVfiTimers.EntityData.Leafs["evpn-virtual-access-vfi-recovery"] = types.YLeaf{"EvpnVirtualAccessVfiRecovery", evpnVirtualAccessVfiTimers.EvpnVirtualAccessVfiRecovery}
    evpnVirtualAccessVfiTimers.EntityData.Leafs["evpn-virtual-access-vfi-peering"] = types.YLeaf{"EvpnVirtualAccessVfiPeering", evpnVirtualAccessVfiTimers.EvpnVirtualAccessVfiPeering}
    evpnVirtualAccessVfiTimers.EntityData.Leafs["evpn-virtual-access-vfi-carving"] = types.YLeaf{"EvpnVirtualAccessVfiCarving", evpnVirtualAccessVfiTimers.EvpnVirtualAccessVfiCarving}
    evpnVirtualAccessVfiTimers.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnVirtualAccessVfiTimers.Enable}
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    EsImportRouteTarget interface{}

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

    evpnVirtualEthernetSegment.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualEthernetSegment.EntityData.Children["identifier"] = types.YChild{"Identifier", &evpnVirtualEthernetSegment.Identifier}
    evpnVirtualEthernetSegment.EntityData.Children["manual-service-carving"] = types.YChild{"ManualServiceCarving", &evpnVirtualEthernetSegment.ManualServiceCarving}
    evpnVirtualEthernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnVirtualEthernetSegment.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnVirtualEthernetSegment.Enable}
    evpnVirtualEthernetSegment.EntityData.Leafs["es-import-route-target"] = types.YLeaf{"EsImportRouteTarget", evpnVirtualEthernetSegment.EsImportRouteTarget}
    return &(evpnVirtualEthernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_Identifier
// Ethernet segment identifier
// This type is a presence type.
type Evpn_EvpnTables_EvpnVirtualAccessVfis_EvpnVirtualAccessVfi_EvpnVirtualEthernetSegment_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type 0's 1st Byte or Type Byte and 1st Byte. The type is string with
    // pattern: b'[0-9a-fA-F]{1,8}'. This attribute is mandatory.
    Bytes01 interface{}

    // 2nd and 3rd Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes23 interface{}

    // 4th and 5th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes45 interface{}

    // 6th and 7th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes67 interface{}

    // 8th and 9th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes89 interface{}

    // Ethernet segment identifier type. The type is EthernetSegmentIdentifier.
    // This attribute is mandatory.
    Type_ interface{}
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

    identifier.EntityData.Children = make(map[string]types.YChild)
    identifier.EntityData.Leafs = make(map[string]types.YLeaf)
    identifier.EntityData.Leafs["bytes01"] = types.YLeaf{"Bytes01", identifier.Bytes01}
    identifier.EntityData.Leafs["bytes23"] = types.YLeaf{"Bytes23", identifier.Bytes23}
    identifier.EntityData.Leafs["bytes45"] = types.YLeaf{"Bytes45", identifier.Bytes45}
    identifier.EntityData.Leafs["bytes67"] = types.YLeaf{"Bytes67", identifier.Bytes67}
    identifier.EntityData.Leafs["bytes89"] = types.YLeaf{"Bytes89", identifier.Bytes89}
    identifier.EntityData.Leafs["type"] = types.YLeaf{"Type_", identifier.Type_}
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

    manualServiceCarving.EntityData.Children = make(map[string]types.YChild)
    manualServiceCarving.EntityData.Children["service-list"] = types.YChild{"ServiceList", &manualServiceCarving.ServiceList}
    manualServiceCarving.EntityData.Leafs = make(map[string]types.YLeaf)
    manualServiceCarving.EntityData.Leafs["enable"] = types.YLeaf{"Enable", manualServiceCarving.Enable}
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

    serviceList.EntityData.Children = make(map[string]types.YChild)
    serviceList.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceList.EntityData.Leafs["primary"] = types.YLeaf{"Primary", serviceList.Primary}
    serviceList.EntityData.Leafs["secondary"] = types.YLeaf{"Secondary", serviceList.Secondary}
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

    evpnLoadBalancing.EntityData.Children = make(map[string]types.YChild)
    evpnLoadBalancing.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnLoadBalancing.EntityData.Leafs["evpn-static-flow-label"] = types.YLeaf{"EvpnStaticFlowLabel", evpnLoadBalancing.EvpnStaticFlowLabel}
    evpnLoadBalancing.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnLoadBalancing.Enable}
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

    evpnBgpAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    evpnBgpAutoDiscovery.EntityData.Children["evpn-route-distinguisher"] = types.YChild{"EvpnRouteDistinguisher", &evpnBgpAutoDiscovery.EvpnRouteDistinguisher}
    evpnBgpAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnBgpAutoDiscovery.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnBgpAutoDiscovery.Enable}
    return &(evpnBgpAutoDiscovery.EntityData)
}

// Evpn_EvpnTables_EvpnBgpAutoDiscovery_EvpnRouteDistinguisher
// Route Distinguisher
type Evpn_EvpnTables_EvpnBgpAutoDiscovery_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type_ interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    evpnRouteDistinguisher.EntityData.Children = make(map[string]types.YChild)
    evpnRouteDistinguisher.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteDistinguisher.EntityData.Leafs["type"] = types.YLeaf{"Type_", evpnRouteDistinguisher.Type_}
    evpnRouteDistinguisher.EntityData.Leafs["as"] = types.YLeaf{"As", evpnRouteDistinguisher.As}
    evpnRouteDistinguisher.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex}
    evpnRouteDistinguisher.EntityData.Leafs["address"] = types.YLeaf{"Address", evpnRouteDistinguisher.Address}
    evpnRouteDistinguisher.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex}
    return &(evpnRouteDistinguisher.EntityData)
}

// Evpn_EvpnTables_EvpnInstances
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_EvpnInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter EVPN Instance configuration submode. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance.
    EvpnInstance []Evpn_EvpnTables_EvpnInstances_EvpnInstance
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

    evpnInstances.EntityData.Children = make(map[string]types.YChild)
    evpnInstances.EntityData.Children["evpn-instance"] = types.YChild{"EvpnInstance", nil}
    for i := range evpnInstances.EvpnInstance {
        evpnInstances.EntityData.Children[types.GetSegmentPath(&evpnInstances.EvpnInstance[i])] = types.YChild{"EvpnInstance", &evpnInstances.EvpnInstance[i]}
    }
    evpnInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evpnInstances.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance
// Enter EVPN Instance configuration submode
type Evpn_EvpnTables_EvpnInstances_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN Instance ID. The type is interface{} with
    // range: 1..65534.
    Eviid interface{}

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
    EvpneviDescription interface{}

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

    // Enter Loadbalancing configuration submode.
    EvpnInstanceLoadBalancing Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceLoadBalancing
}

func (evpnInstance *Evpn_EvpnTables_EvpnInstances_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "evpn-instances"
    evpnInstance.EntityData.SegmentPath = "evpn-instance" + "[eviid='" + fmt.Sprintf("%v", evpnInstance.Eviid) + "']" + "[encapsulation='" + fmt.Sprintf("%v", evpnInstance.Encapsulation) + "']" + "[side='" + fmt.Sprintf("%v", evpnInstance.Side) + "']"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = make(map[string]types.YChild)
    evpnInstance.EntityData.Children["evpn-instance-bgp-auto-discovery"] = types.YChild{"EvpnInstanceBgpAutoDiscovery", &evpnInstance.EvpnInstanceBgpAutoDiscovery}
    evpnInstance.EntityData.Children["evpn-instance-advertise-mac"] = types.YChild{"EvpnInstanceAdvertiseMac", &evpnInstance.EvpnInstanceAdvertiseMac}
    evpnInstance.EntityData.Children["evpn-instance-load-balancing"] = types.YChild{"EvpnInstanceLoadBalancing", &evpnInstance.EvpnInstanceLoadBalancing}
    evpnInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnInstance.EntityData.Leafs["eviid"] = types.YLeaf{"Eviid", evpnInstance.Eviid}
    evpnInstance.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", evpnInstance.Encapsulation}
    evpnInstance.EntityData.Leafs["side"] = types.YLeaf{"Side", evpnInstance.Side}
    evpnInstance.EntityData.Leafs["evi-reorig-disable"] = types.YLeaf{"EviReorigDisable", evpnInstance.EviReorigDisable}
    evpnInstance.EntityData.Leafs["evi-advertise-mac-deprecated"] = types.YLeaf{"EviAdvertiseMacDeprecated", evpnInstance.EviAdvertiseMacDeprecated}
    evpnInstance.EntityData.Leafs["evpnevi-description"] = types.YLeaf{"EvpneviDescription", evpnInstance.EvpneviDescription}
    evpnInstance.EntityData.Leafs["evi-ecmp-disable"] = types.YLeaf{"EviEcmpDisable", evpnInstance.EviEcmpDisable}
    evpnInstance.EntityData.Leafs["evi-unknown-unicast-flooding-disable"] = types.YLeaf{"EviUnknownUnicastFloodingDisable", evpnInstance.EviUnknownUnicastFloodingDisable}
    evpnInstance.EntityData.Leafs["evpn-evi-cw-disable"] = types.YLeaf{"EvpnEviCwDisable", evpnInstance.EvpnEviCwDisable}
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

    evpnInstanceBgpAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    evpnInstanceBgpAutoDiscovery.EntityData.Children["evpn-route-distinguisher"] = types.YChild{"EvpnRouteDistinguisher", &evpnInstanceBgpAutoDiscovery.EvpnRouteDistinguisher}
    evpnInstanceBgpAutoDiscovery.EntityData.Children["evpn-route-targets"] = types.YChild{"EvpnRouteTargets", &evpnInstanceBgpAutoDiscovery.EvpnRouteTargets}
    evpnInstanceBgpAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnInstanceBgpAutoDiscovery.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnInstanceBgpAutoDiscovery.Enable}
    evpnInstanceBgpAutoDiscovery.EntityData.Leafs["table-policy"] = types.YLeaf{"TablePolicy", evpnInstanceBgpAutoDiscovery.TablePolicy}
    return &(evpnInstanceBgpAutoDiscovery.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteDistinguisher
// Route Distinguisher
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteDistinguisher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Distinguisher Type. The type is BgpRouteDistinguisher.
    Type_ interface{}

    // Two byte or 4 byte AS number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // AS:nn (hex or decimal format). The type is interface{} with range:
    // 0..4294967295.
    AsIndex interface{}

    // IPV4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    evpnRouteDistinguisher.EntityData.Children = make(map[string]types.YChild)
    evpnRouteDistinguisher.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteDistinguisher.EntityData.Leafs["type"] = types.YLeaf{"Type_", evpnRouteDistinguisher.Type_}
    evpnRouteDistinguisher.EntityData.Leafs["as"] = types.YLeaf{"As", evpnRouteDistinguisher.As}
    evpnRouteDistinguisher.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", evpnRouteDistinguisher.AsIndex}
    evpnRouteDistinguisher.EntityData.Leafs["address"] = types.YLeaf{"Address", evpnRouteDistinguisher.Address}
    evpnRouteDistinguisher.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", evpnRouteDistinguisher.AddrIndex}
    return &(evpnRouteDistinguisher.EntityData)
}

// Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets
// Route Target
type Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs.
    EvpnRouteTargetAs []Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetAs

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone.
    EvpnRouteTargetNone []Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetNone

    // Name of the Route Target. The type is slice of
    // Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address.
    EvpnRouteTargetIpv4Address []Evpn_EvpnTables_EvpnInstances_EvpnInstance_EvpnInstanceBgpAutoDiscovery_EvpnRouteTargets_EvpnRouteTargetIpv4Address
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

    evpnRouteTargets.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargets.EntityData.Children["evpn-route-target-as"] = types.YChild{"EvpnRouteTargetAs", nil}
    for i := range evpnRouteTargets.EvpnRouteTargetAs {
        evpnRouteTargets.EntityData.Children[types.GetSegmentPath(&evpnRouteTargets.EvpnRouteTargetAs[i])] = types.YChild{"EvpnRouteTargetAs", &evpnRouteTargets.EvpnRouteTargetAs[i]}
    }
    evpnRouteTargets.EntityData.Children["evpn-route-target-none"] = types.YChild{"EvpnRouteTargetNone", nil}
    for i := range evpnRouteTargets.EvpnRouteTargetNone {
        evpnRouteTargets.EntityData.Children[types.GetSegmentPath(&evpnRouteTargets.EvpnRouteTargetNone[i])] = types.YChild{"EvpnRouteTargetNone", &evpnRouteTargets.EvpnRouteTargetNone[i]}
    }
    evpnRouteTargets.EntityData.Children["evpn-route-target-ipv4-address"] = types.YChild{"EvpnRouteTargetIpv4Address", nil}
    for i := range evpnRouteTargets.EvpnRouteTargetIpv4Address {
        evpnRouteTargets.EntityData.Children[types.GetSegmentPath(&evpnRouteTargets.EvpnRouteTargetIpv4Address[i])] = types.YChild{"EvpnRouteTargetIpv4Address", &evpnRouteTargets.EvpnRouteTargetIpv4Address[i]}
    }
    evpnRouteTargets.EntityData.Leafs = make(map[string]types.YLeaf)
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
    evpnRouteTargetAs.EntityData.SegmentPath = "evpn-route-target-as" + "[format='" + fmt.Sprintf("%v", evpnRouteTargetAs.Format) + "']" + "[role='" + fmt.Sprintf("%v", evpnRouteTargetAs.Role) + "']" + "[as='" + fmt.Sprintf("%v", evpnRouteTargetAs.As) + "']" + "[as-index='" + fmt.Sprintf("%v", evpnRouteTargetAs.AsIndex) + "']" + "[stitching='" + fmt.Sprintf("%v", evpnRouteTargetAs.Stitching) + "']"
    evpnRouteTargetAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetAs.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargetAs.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteTargetAs.EntityData.Leafs["format"] = types.YLeaf{"Format", evpnRouteTargetAs.Format}
    evpnRouteTargetAs.EntityData.Leafs["role"] = types.YLeaf{"Role", evpnRouteTargetAs.Role}
    evpnRouteTargetAs.EntityData.Leafs["as"] = types.YLeaf{"As", evpnRouteTargetAs.As}
    evpnRouteTargetAs.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", evpnRouteTargetAs.AsIndex}
    evpnRouteTargetAs.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", evpnRouteTargetAs.Stitching}
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
    evpnRouteTargetNone.EntityData.SegmentPath = "evpn-route-target-none" + "[format='" + fmt.Sprintf("%v", evpnRouteTargetNone.Format) + "']" + "[role='" + fmt.Sprintf("%v", evpnRouteTargetNone.Role) + "']" + "[stitching='" + fmt.Sprintf("%v", evpnRouteTargetNone.Stitching) + "']"
    evpnRouteTargetNone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetNone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetNone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetNone.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargetNone.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteTargetNone.EntityData.Leafs["format"] = types.YLeaf{"Format", evpnRouteTargetNone.Format}
    evpnRouteTargetNone.EntityData.Leafs["role"] = types.YLeaf{"Role", evpnRouteTargetNone.Role}
    evpnRouteTargetNone.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", evpnRouteTargetNone.Stitching}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    evpnRouteTargetIpv4Address.EntityData.SegmentPath = "evpn-route-target-ipv4-address" + "[format='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Format) + "']" + "[role='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Role) + "']" + "[address='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Address) + "']" + "[addr-index='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.AddrIndex) + "']" + "[stitching='" + fmt.Sprintf("%v", evpnRouteTargetIpv4Address.Stitching) + "']"
    evpnRouteTargetIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnRouteTargetIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnRouteTargetIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnRouteTargetIpv4Address.EntityData.Children = make(map[string]types.YChild)
    evpnRouteTargetIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnRouteTargetIpv4Address.EntityData.Leafs["format"] = types.YLeaf{"Format", evpnRouteTargetIpv4Address.Format}
    evpnRouteTargetIpv4Address.EntityData.Leafs["role"] = types.YLeaf{"Role", evpnRouteTargetIpv4Address.Role}
    evpnRouteTargetIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", evpnRouteTargetIpv4Address.Address}
    evpnRouteTargetIpv4Address.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", evpnRouteTargetIpv4Address.AddrIndex}
    evpnRouteTargetIpv4Address.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", evpnRouteTargetIpv4Address.Stitching}
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

    evpnInstanceAdvertiseMac.EntityData.Children = make(map[string]types.YChild)
    evpnInstanceAdvertiseMac.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnInstanceAdvertiseMac.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnInstanceAdvertiseMac.Enable}
    evpnInstanceAdvertiseMac.EntityData.Leafs["evi-advertise-mac-bvi"] = types.YLeaf{"EviAdvertiseMacBvi", evpnInstanceAdvertiseMac.EviAdvertiseMacBvi}
    return &(evpnInstanceAdvertiseMac.EntityData)
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

    evpnInstanceLoadBalancing.EntityData.Children = make(map[string]types.YChild)
    evpnInstanceLoadBalancing.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnInstanceLoadBalancing.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnInstanceLoadBalancing.Enable}
    evpnInstanceLoadBalancing.EntityData.Leafs["evi-static-flow-label"] = types.YLeaf{"EviStaticFlowLabel", evpnInstanceLoadBalancing.EviStaticFlowLabel}
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

    evpnLogging.EntityData.Children = make(map[string]types.YChild)
    evpnLogging.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnLogging.EntityData.Leafs["evpn-df-election"] = types.YLeaf{"EvpnDfElection", evpnLogging.EvpnDfElection}
    evpnLogging.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnLogging.Enable}
    return &(evpnLogging.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces
// Attachment Circuit interfaces
type Evpn_EvpnTables_EvpnInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment circuit interface. The type is slice of
    // Evpn_EvpnTables_EvpnInterfaces_EvpnInterface.
    EvpnInterface []Evpn_EvpnTables_EvpnInterfaces_EvpnInterface
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

    evpnInterfaces.EntityData.Children = make(map[string]types.YChild)
    evpnInterfaces.EntityData.Children["evpn-interface"] = types.YChild{"EvpnInterface", nil}
    for i := range evpnInterfaces.EvpnInterface {
        evpnInterfaces.EntityData.Children[types.GetSegmentPath(&evpnInterfaces.EvpnInterface[i])] = types.YChild{"EvpnInterface", &evpnInterfaces.EvpnInterface[i]}
    }
    evpnInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evpnInterfaces.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface
// Attachment circuit interface
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the attachment circuit interface. The type
    // is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

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
    evpnInterface.EntityData.SegmentPath = "evpn-interface" + "[interface-name='" + fmt.Sprintf("%v", evpnInterface.InterfaceName) + "']"
    evpnInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInterface.EntityData.Children = make(map[string]types.YChild)
    evpnInterface.EntityData.Children["evpnac-timers"] = types.YChild{"EvpnacTimers", &evpnInterface.EvpnacTimers}
    evpnInterface.EntityData.Children["ethernet-segment"] = types.YChild{"EthernetSegment", &evpnInterface.EthernetSegment}
    evpnInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", evpnInterface.InterfaceName}
    evpnInterface.EntityData.Leafs["mac-flush"] = types.YLeaf{"MacFlush", evpnInterface.MacFlush}
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

    evpnacTimers.EntityData.Children = make(map[string]types.YChild)
    evpnacTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnacTimers.EntityData.Leafs["evpnac-peering"] = types.YLeaf{"EvpnacPeering", evpnacTimers.EvpnacPeering}
    evpnacTimers.EntityData.Leafs["evpnac-carving"] = types.YLeaf{"EvpnacCarving", evpnacTimers.EvpnacCarving}
    evpnacTimers.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnacTimers.Enable}
    evpnacTimers.EntityData.Leafs["evpnac-recovery"] = types.YLeaf{"EvpnacRecovery", evpnacTimers.EvpnacRecovery}
    return &(evpnacTimers.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment
// Enter Ethernet Segment configuration submode
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force ethernet segment to remain single-homed. The type is interface{}.
    ForceSingleHomed interface{}

    // Enable single-active load balancing mode. The type is interface{}.
    LoadBalancingSingleActive interface{}

    // Enable Ethernet Segment. The type is interface{}.
    Enable interface{}

    // Backbone Source MAC. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    BackboneSourceMac interface{}

    // ES-Import Route Target. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    EsImportRouteTarget interface{}

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

    ethernetSegment.EntityData.Children = make(map[string]types.YChild)
    ethernetSegment.EntityData.Children["identifier"] = types.YChild{"Identifier", &ethernetSegment.Identifier}
    ethernetSegment.EntityData.Children["manual-service-carving"] = types.YChild{"ManualServiceCarving", &ethernetSegment.ManualServiceCarving}
    ethernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegment.EntityData.Leafs["force-single-homed"] = types.YLeaf{"ForceSingleHomed", ethernetSegment.ForceSingleHomed}
    ethernetSegment.EntityData.Leafs["load-balancing-single-active"] = types.YLeaf{"LoadBalancingSingleActive", ethernetSegment.LoadBalancingSingleActive}
    ethernetSegment.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ethernetSegment.Enable}
    ethernetSegment.EntityData.Leafs["backbone-source-mac"] = types.YLeaf{"BackboneSourceMac", ethernetSegment.BackboneSourceMac}
    ethernetSegment.EntityData.Leafs["es-import-route-target"] = types.YLeaf{"EsImportRouteTarget", ethernetSegment.EsImportRouteTarget}
    return &(ethernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_Identifier
// Ethernet segment identifier
// This type is a presence type.
type Evpn_EvpnTables_EvpnInterfaces_EvpnInterface_EthernetSegment_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type 0's 1st Byte or Type Byte and 1st Byte. The type is string with
    // pattern: b'[0-9a-fA-F]{1,8}'. This attribute is mandatory.
    Bytes01 interface{}

    // 2nd and 3rd Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes23 interface{}

    // 4th and 5th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes45 interface{}

    // 6th and 7th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes67 interface{}

    // 8th and 9th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes89 interface{}

    // Ethernet segment identifier type. The type is EthernetSegmentIdentifier.
    // This attribute is mandatory.
    Type_ interface{}
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

    identifier.EntityData.Children = make(map[string]types.YChild)
    identifier.EntityData.Leafs = make(map[string]types.YLeaf)
    identifier.EntityData.Leafs["bytes01"] = types.YLeaf{"Bytes01", identifier.Bytes01}
    identifier.EntityData.Leafs["bytes23"] = types.YLeaf{"Bytes23", identifier.Bytes23}
    identifier.EntityData.Leafs["bytes45"] = types.YLeaf{"Bytes45", identifier.Bytes45}
    identifier.EntityData.Leafs["bytes67"] = types.YLeaf{"Bytes67", identifier.Bytes67}
    identifier.EntityData.Leafs["bytes89"] = types.YLeaf{"Bytes89", identifier.Bytes89}
    identifier.EntityData.Leafs["type"] = types.YLeaf{"Type_", identifier.Type_}
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

    manualServiceCarving.EntityData.Children = make(map[string]types.YChild)
    manualServiceCarving.EntityData.Children["service-list"] = types.YChild{"ServiceList", &manualServiceCarving.ServiceList}
    manualServiceCarving.EntityData.Leafs = make(map[string]types.YLeaf)
    manualServiceCarving.EntityData.Leafs["enable"] = types.YLeaf{"Enable", manualServiceCarving.Enable}
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

    serviceList.EntityData.Children = make(map[string]types.YChild)
    serviceList.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceList.EntityData.Leafs["primary"] = types.YLeaf{"Primary", serviceList.Primary}
    serviceList.EntityData.Leafs["secondary"] = types.YLeaf{"Secondary", serviceList.Secondary}
    return &(serviceList.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws
// Virtual Access Pseudowire interfaces
type Evpn_EvpnTables_EvpnVirtualAccessPws struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual Access Pseudowire. The type is slice of
    // Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw.
    EvpnVirtualAccessPw []Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw
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

    evpnVirtualAccessPws.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualAccessPws.EntityData.Children["evpn-virtual-access-pw"] = types.YChild{"EvpnVirtualAccessPw", nil}
    for i := range evpnVirtualAccessPws.EvpnVirtualAccessPw {
        evpnVirtualAccessPws.EntityData.Children[types.GetSegmentPath(&evpnVirtualAccessPws.EvpnVirtualAccessPw[i])] = types.YChild{"EvpnVirtualAccessPw", &evpnVirtualAccessPws.EvpnVirtualAccessPw[i]}
    }
    evpnVirtualAccessPws.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evpnVirtualAccessPws.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw
// Virtual Access Pseudowire
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    evpnVirtualAccessPw.EntityData.SegmentPath = "evpn-virtual-access-pw" + "[neighbor='" + fmt.Sprintf("%v", evpnVirtualAccessPw.Neighbor) + "']" + "[pseudowire-id='" + fmt.Sprintf("%v", evpnVirtualAccessPw.PseudowireId) + "']"
    evpnVirtualAccessPw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnVirtualAccessPw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnVirtualAccessPw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnVirtualAccessPw.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualAccessPw.EntityData.Children["evpn-virtual-access-pw-timers"] = types.YChild{"EvpnVirtualAccessPwTimers", &evpnVirtualAccessPw.EvpnVirtualAccessPwTimers}
    evpnVirtualAccessPw.EntityData.Children["evpn-virtual-ethernet-segment"] = types.YChild{"EvpnVirtualEthernetSegment", &evpnVirtualAccessPw.EvpnVirtualEthernetSegment}
    evpnVirtualAccessPw.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnVirtualAccessPw.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", evpnVirtualAccessPw.Neighbor}
    evpnVirtualAccessPw.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", evpnVirtualAccessPw.PseudowireId}
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

    evpnVirtualAccessPwTimers.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualAccessPwTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnVirtualAccessPwTimers.EntityData.Leafs["evpn-virtual-access-pw-recovery"] = types.YLeaf{"EvpnVirtualAccessPwRecovery", evpnVirtualAccessPwTimers.EvpnVirtualAccessPwRecovery}
    evpnVirtualAccessPwTimers.EntityData.Leafs["evpn-virtual-access-pw-peering"] = types.YLeaf{"EvpnVirtualAccessPwPeering", evpnVirtualAccessPwTimers.EvpnVirtualAccessPwPeering}
    evpnVirtualAccessPwTimers.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnVirtualAccessPwTimers.Enable}
    evpnVirtualAccessPwTimers.EntityData.Leafs["evpn-virtual-access-pw-carving"] = types.YLeaf{"EvpnVirtualAccessPwCarving", evpnVirtualAccessPwTimers.EvpnVirtualAccessPwCarving}
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    EsImportRouteTarget interface{}

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

    evpnVirtualEthernetSegment.EntityData.Children = make(map[string]types.YChild)
    evpnVirtualEthernetSegment.EntityData.Children["identifier"] = types.YChild{"Identifier", &evpnVirtualEthernetSegment.Identifier}
    evpnVirtualEthernetSegment.EntityData.Children["manual-service-carving"] = types.YChild{"ManualServiceCarving", &evpnVirtualEthernetSegment.ManualServiceCarving}
    evpnVirtualEthernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnVirtualEthernetSegment.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnVirtualEthernetSegment.Enable}
    evpnVirtualEthernetSegment.EntityData.Leafs["es-import-route-target"] = types.YLeaf{"EsImportRouteTarget", evpnVirtualEthernetSegment.EsImportRouteTarget}
    return &(evpnVirtualEthernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_Identifier
// Ethernet segment identifier
// This type is a presence type.
type Evpn_EvpnTables_EvpnVirtualAccessPws_EvpnVirtualAccessPw_EvpnVirtualEthernetSegment_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type 0's 1st Byte or Type Byte and 1st Byte. The type is string with
    // pattern: b'[0-9a-fA-F]{1,8}'. This attribute is mandatory.
    Bytes01 interface{}

    // 2nd and 3rd Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes23 interface{}

    // 4th and 5th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes45 interface{}

    // 6th and 7th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes67 interface{}

    // 8th and 9th Bytes. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    // This attribute is mandatory. Units are byte.
    Bytes89 interface{}

    // Ethernet segment identifier type. The type is EthernetSegmentIdentifier.
    // This attribute is mandatory.
    Type_ interface{}
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

    identifier.EntityData.Children = make(map[string]types.YChild)
    identifier.EntityData.Leafs = make(map[string]types.YLeaf)
    identifier.EntityData.Leafs["bytes01"] = types.YLeaf{"Bytes01", identifier.Bytes01}
    identifier.EntityData.Leafs["bytes23"] = types.YLeaf{"Bytes23", identifier.Bytes23}
    identifier.EntityData.Leafs["bytes45"] = types.YLeaf{"Bytes45", identifier.Bytes45}
    identifier.EntityData.Leafs["bytes67"] = types.YLeaf{"Bytes67", identifier.Bytes67}
    identifier.EntityData.Leafs["bytes89"] = types.YLeaf{"Bytes89", identifier.Bytes89}
    identifier.EntityData.Leafs["type"] = types.YLeaf{"Type_", identifier.Type_}
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

    manualServiceCarving.EntityData.Children = make(map[string]types.YChild)
    manualServiceCarving.EntityData.Children["service-list"] = types.YChild{"ServiceList", &manualServiceCarving.ServiceList}
    manualServiceCarving.EntityData.Leafs = make(map[string]types.YLeaf)
    manualServiceCarving.EntityData.Leafs["enable"] = types.YLeaf{"Enable", manualServiceCarving.Enable}
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

    serviceList.EntityData.Children = make(map[string]types.YChild)
    serviceList.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceList.EntityData.Leafs["primary"] = types.YLeaf{"Primary", serviceList.Primary}
    serviceList.EntityData.Leafs["secondary"] = types.YLeaf{"Secondary", serviceList.Secondary}
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

    evpnEthernetSegment.EntityData.Children = make(map[string]types.YChild)
    evpnEthernetSegment.EntityData.Children["evpn-esi-types"] = types.YChild{"EvpnEsiTypes", &evpnEthernetSegment.EvpnEsiTypes}
    evpnEthernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnEthernetSegment.EntityData.Leafs["enable"] = types.YLeaf{"Enable", evpnEthernetSegment.Enable}
    return &(evpnEthernetSegment.EntityData)
}

// Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes
// EVPN ESI type table
type Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ESI type. The type is slice of
    // Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType.
    EvpnEsiType []Evpn_EvpnTables_EvpnEthernetSegment_EvpnEsiTypes_EvpnEsiType
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

    evpnEsiTypes.EntityData.Children = make(map[string]types.YChild)
    evpnEsiTypes.EntityData.Children["evpn-esi-type"] = types.YChild{"EvpnEsiType", nil}
    for i := range evpnEsiTypes.EvpnEsiType {
        evpnEsiTypes.EntityData.Children[types.GetSegmentPath(&evpnEsiTypes.EvpnEsiType[i])] = types.YChild{"EvpnEsiType", &evpnEsiTypes.EvpnEsiType[i]}
    }
    evpnEsiTypes.EntityData.Leafs = make(map[string]types.YLeaf)
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
    evpnEsiType.EntityData.SegmentPath = "evpn-esi-type" + "[esi-type='" + fmt.Sprintf("%v", evpnEsiType.EsiType) + "']"
    evpnEsiType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnEsiType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnEsiType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnEsiType.EntityData.Children = make(map[string]types.YChild)
    evpnEsiType.EntityData.Leafs = make(map[string]types.YLeaf)
    evpnEsiType.EntityData.Leafs["esi-type"] = types.YLeaf{"EsiType", evpnEsiType.EsiType}
    evpnEsiType.EntityData.Leafs["disable-auto-generation"] = types.YLeaf{"DisableAutoGeneration", evpnEsiType.DisableAutoGeneration}
    return &(evpnEsiType.EntityData)
}

