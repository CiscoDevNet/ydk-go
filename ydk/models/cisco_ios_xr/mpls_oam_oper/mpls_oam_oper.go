// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-oam package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpls-oam: MPLS OAM operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_oam_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_oam_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-oam-oper mpls-oam}", reflect.TypeOf(MplsOam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-oam-oper:mpls-oam", reflect.TypeOf(MplsOam{}))
}

// LspvBagInterfaceState represents LSPV interface state
type LspvBagInterfaceState string

const (
    // Not ready
    LspvBagInterfaceState_not_ready LspvBagInterfaceState = "not-ready"

    // Admin down
    LspvBagInterfaceState_admin_down LspvBagInterfaceState = "admin-down"

    // Down
    LspvBagInterfaceState_down LspvBagInterfaceState = "down"

    // UP
    LspvBagInterfaceState_up LspvBagInterfaceState = "up"

    // Shutdown
    LspvBagInterfaceState_shutdown LspvBagInterfaceState = "shutdown"

    // Error disable
    LspvBagInterfaceState_error_disable LspvBagInterfaceState = "error-disable"

    // Immediate down
    LspvBagInterfaceState_down_immediate LspvBagInterfaceState = "down-immediate"

    // Immediate admin
    LspvBagInterfaceState_admin_immediate LspvBagInterfaceState = "admin-immediate"

    // Graceful down
    LspvBagInterfaceState_graceful_down LspvBagInterfaceState = "graceful-down"

    // Begin shutdown
    LspvBagInterfaceState_begin_shutdown LspvBagInterfaceState = "begin-shutdown"

    // End shutdown
    LspvBagInterfaceState_end_shutdown LspvBagInterfaceState = "end-shutdown"

    // Begin error disable
    LspvBagInterfaceState_begin_error_disable LspvBagInterfaceState = "begin-error-disable"

    // End error disable
    LspvBagInterfaceState_end_error_disable LspvBagInterfaceState = "end-error-disable"

    // Begin graceful down
    LspvBagInterfaceState_begin_graceful_down LspvBagInterfaceState = "begin-graceful-down"

    // Reset
    LspvBagInterfaceState_reset LspvBagInterfaceState = "reset"

    // Operational
    LspvBagInterfaceState_operational LspvBagInterfaceState = "operational"

    // Not operational
    LspvBagInterfaceState_not_operational LspvBagInterfaceState = "not-operational"

    // Unknown
    LspvBagInterfaceState_not_known LspvBagInterfaceState = "not-known"
)

// MplsOam
// MPLS OAM operational data
type MplsOam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS OAM interface operational data.
    Interface_ MplsOam_Interface

    // LSPV packet counters operational data.
    Packet MplsOam_Packet

    // LSPV global counters operational data.
    Global MplsOam_Global
}

func (mplsOam *MplsOam) GetEntityData() *types.CommonEntityData {
    mplsOam.EntityData.YFilter = mplsOam.YFilter
    mplsOam.EntityData.YangName = "mpls-oam"
    mplsOam.EntityData.BundleName = "cisco_ios_xr"
    mplsOam.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-oam-oper"
    mplsOam.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-oam-oper:mpls-oam"
    mplsOam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsOam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsOam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsOam.EntityData.Children = make(map[string]types.YChild)
    mplsOam.EntityData.Children["interface"] = types.YChild{"Interface_", &mplsOam.Interface_}
    mplsOam.EntityData.Children["packet"] = types.YChild{"Packet", &mplsOam.Packet}
    mplsOam.EntityData.Children["global"] = types.YChild{"Global", &mplsOam.Global}
    mplsOam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsOam.EntityData)
}

// MplsOam_Interface
// MPLS OAM interface operational data
type MplsOam_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS OAM interface detail data.
    Briefs MplsOam_Interface_Briefs

    // MPLS OAM interface detail data.
    Details MplsOam_Interface_Details
}

func (self *MplsOam_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "mpls-oam"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["briefs"] = types.YChild{"Briefs", &self.Briefs}
    self.EntityData.Children["details"] = types.YChild{"Details", &self.Details}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// MplsOam_Interface_Briefs
// MPLS OAM interface detail data
type MplsOam_Interface_Briefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS OAM interface operational data. The type is slice of
    // MplsOam_Interface_Briefs_Brief.
    Brief []MplsOam_Interface_Briefs_Brief
}

func (briefs *MplsOam_Interface_Briefs) GetEntityData() *types.CommonEntityData {
    briefs.EntityData.YFilter = briefs.YFilter
    briefs.EntityData.YangName = "briefs"
    briefs.EntityData.BundleName = "cisco_ios_xr"
    briefs.EntityData.ParentYangName = "interface"
    briefs.EntityData.SegmentPath = "briefs"
    briefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefs.EntityData.Children = make(map[string]types.YChild)
    briefs.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range briefs.Brief {
        briefs.EntityData.Children[types.GetSegmentPath(&briefs.Brief[i])] = types.YChild{"Brief", &briefs.Brief[i]}
    }
    briefs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(briefs.EntityData)
}

// MplsOam_Interface_Briefs_Brief
// MPLS OAM interface operational data
type MplsOam_Interface_Briefs_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceNameXr interface{}

    // Interface state. The type is LspvBagInterfaceState.
    State interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Prefix length (IPv4). The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Prefix length (IPv6). The type is interface{} with range: 0..4294967295.
    PrefixLengthV6 interface{}

    // Primary interface address (IPv4). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddress interface{}

    // Primary interface address (IPv6). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddressV6 interface{}
}

func (brief *MplsOam_Interface_Briefs_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "briefs"
    brief.EntityData.SegmentPath = "brief" + "[interface-name='" + fmt.Sprintf("%v", brief.InterfaceName) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", brief.InterfaceName}
    brief.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", brief.InterfaceNameXr}
    brief.EntityData.Leafs["state"] = types.YLeaf{"State", brief.State}
    brief.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", brief.Mtu}
    brief.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", brief.PrefixLength}
    brief.EntityData.Leafs["prefix-length-v6"] = types.YLeaf{"PrefixLengthV6", brief.PrefixLengthV6}
    brief.EntityData.Leafs["primary-address"] = types.YLeaf{"PrimaryAddress", brief.PrimaryAddress}
    brief.EntityData.Leafs["primary-address-v6"] = types.YLeaf{"PrimaryAddressV6", brief.PrimaryAddressV6}
    return &(brief.EntityData)
}

// MplsOam_Interface_Details
// MPLS OAM interface detail data
type MplsOam_Interface_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS OAM interface operational data. The type is slice of
    // MplsOam_Interface_Details_Detail.
    Detail []MplsOam_Interface_Details_Detail
}

func (details *MplsOam_Interface_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "interface"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range details.Detail {
        details.EntityData.Children[types.GetSegmentPath(&details.Detail[i])] = types.YChild{"Detail", &details.Detail[i]}
    }
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// MplsOam_Interface_Details_Detail
// MPLS OAM interface operational data
type MplsOam_Interface_Details_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface brief.
    InterfaceBrief MplsOam_Interface_Details_Detail_InterfaceBrief

    // Packet statistics.
    PacketStatistics MplsOam_Interface_Details_Detail_PacketStatistics
}

func (detail *MplsOam_Interface_Details_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "details"
    detail.EntityData.SegmentPath = "detail" + "[interface-name='" + fmt.Sprintf("%v", detail.InterfaceName) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["interface-brief"] = types.YChild{"InterfaceBrief", &detail.InterfaceBrief}
    detail.EntityData.Children["packet-statistics"] = types.YChild{"PacketStatistics", &detail.PacketStatistics}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", detail.InterfaceName}
    return &(detail.EntityData)
}

// MplsOam_Interface_Details_Detail_InterfaceBrief
// Interface brief
type MplsOam_Interface_Details_Detail_InterfaceBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceNameXr interface{}

    // Interface state. The type is LspvBagInterfaceState.
    State interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Prefix length (IPv4). The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Prefix length (IPv6). The type is interface{} with range: 0..4294967295.
    PrefixLengthV6 interface{}

    // Primary interface address (IPv4). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddress interface{}

    // Primary interface address (IPv6). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrimaryAddressV6 interface{}
}

func (interfaceBrief *MplsOam_Interface_Details_Detail_InterfaceBrief) GetEntityData() *types.CommonEntityData {
    interfaceBrief.EntityData.YFilter = interfaceBrief.YFilter
    interfaceBrief.EntityData.YangName = "interface-brief"
    interfaceBrief.EntityData.BundleName = "cisco_ios_xr"
    interfaceBrief.EntityData.ParentYangName = "detail"
    interfaceBrief.EntityData.SegmentPath = "interface-brief"
    interfaceBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceBrief.EntityData.Children = make(map[string]types.YChild)
    interfaceBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceBrief.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", interfaceBrief.InterfaceNameXr}
    interfaceBrief.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceBrief.State}
    interfaceBrief.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", interfaceBrief.Mtu}
    interfaceBrief.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", interfaceBrief.PrefixLength}
    interfaceBrief.EntityData.Leafs["prefix-length-v6"] = types.YLeaf{"PrefixLengthV6", interfaceBrief.PrefixLengthV6}
    interfaceBrief.EntityData.Leafs["primary-address"] = types.YLeaf{"PrimaryAddress", interfaceBrief.PrimaryAddress}
    interfaceBrief.EntityData.Leafs["primary-address-v6"] = types.YLeaf{"PrimaryAddressV6", interfaceBrief.PrimaryAddressV6}
    return &(interfaceBrief.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics
// Packet statistics
type MplsOam_Interface_Details_Detail_PacketStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet reception counts.
    Received MplsOam_Interface_Details_Detail_PacketStatistics_Received

    // Packet transmit counts.
    Sent MplsOam_Interface_Details_Detail_PacketStatistics_Sent

    // Working Request Packet transmit counts.
    WorkingReqSent MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent

    // Working Reply Packet transmit counts.
    WorkingRepSent MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent

    // Protect Request Packet transmit counts.
    ProtectReqSent MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent

    // Protect Reply Packet transmit counts.
    ProtectRepSent MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent
}

func (packetStatistics *MplsOam_Interface_Details_Detail_PacketStatistics) GetEntityData() *types.CommonEntityData {
    packetStatistics.EntityData.YFilter = packetStatistics.YFilter
    packetStatistics.EntityData.YangName = "packet-statistics"
    packetStatistics.EntityData.BundleName = "cisco_ios_xr"
    packetStatistics.EntityData.ParentYangName = "detail"
    packetStatistics.EntityData.SegmentPath = "packet-statistics"
    packetStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetStatistics.EntityData.Children = make(map[string]types.YChild)
    packetStatistics.EntityData.Children["received"] = types.YChild{"Received", &packetStatistics.Received}
    packetStatistics.EntityData.Children["sent"] = types.YChild{"Sent", &packetStatistics.Sent}
    packetStatistics.EntityData.Children["working-req-sent"] = types.YChild{"WorkingReqSent", &packetStatistics.WorkingReqSent}
    packetStatistics.EntityData.Children["working-rep-sent"] = types.YChild{"WorkingRepSent", &packetStatistics.WorkingRepSent}
    packetStatistics.EntityData.Children["protect-req-sent"] = types.YChild{"ProtectReqSent", &packetStatistics.ProtectReqSent}
    packetStatistics.EntityData.Children["protect-rep-sent"] = types.YChild{"ProtectRepSent", &packetStatistics.ProtectRepSent}
    packetStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(packetStatistics.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received
// Packet reception counts
type MplsOam_Interface_Details_Detail_PacketStatistics_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received good request.
    ReceivedGoodRequest MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest

    // Received good reply.
    ReceivedGoodReply MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply

    // Received unknown packets.
    ReceivedUnknown MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown

    // IP header error.
    ReceivedErrorIpHeader MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader

    // UDP header error.
    ReceivedErrorUdpHeader MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader

    // RUNT error.
    ReceivedErrorRunt MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt

    // Dropped queue full.
    ReceivedErrorQueueFull MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull

    // General error.
    ReceivedErrorGeneral MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral

    // Error no Interfaces.
    ReceivedErrorNoInterface MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface

    // Error no memory.
    ReceivedErrorNoMemory MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory

    // Protect Protocol Received good request.
    ProtectProtocolReceivedGoodRequest MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest

    // Protect Protocol Received good reply.
    ProtectProtocolReceivedGoodReply MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply

    // Received Reqeust with BFD TLV.
    ReceivedGoodBfdRequest MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest

    // Received Reply with BFD TLV.
    ReceivedGoodBfdReply MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply
}

func (received *MplsOam_Interface_Details_Detail_PacketStatistics_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "packet-statistics"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Children["received-good-request"] = types.YChild{"ReceivedGoodRequest", &received.ReceivedGoodRequest}
    received.EntityData.Children["received-good-reply"] = types.YChild{"ReceivedGoodReply", &received.ReceivedGoodReply}
    received.EntityData.Children["received-unknown"] = types.YChild{"ReceivedUnknown", &received.ReceivedUnknown}
    received.EntityData.Children["received-error-ip-header"] = types.YChild{"ReceivedErrorIpHeader", &received.ReceivedErrorIpHeader}
    received.EntityData.Children["received-error-udp-header"] = types.YChild{"ReceivedErrorUdpHeader", &received.ReceivedErrorUdpHeader}
    received.EntityData.Children["received-error-runt"] = types.YChild{"ReceivedErrorRunt", &received.ReceivedErrorRunt}
    received.EntityData.Children["received-error-queue-full"] = types.YChild{"ReceivedErrorQueueFull", &received.ReceivedErrorQueueFull}
    received.EntityData.Children["received-error-general"] = types.YChild{"ReceivedErrorGeneral", &received.ReceivedErrorGeneral}
    received.EntityData.Children["received-error-no-interface"] = types.YChild{"ReceivedErrorNoInterface", &received.ReceivedErrorNoInterface}
    received.EntityData.Children["received-error-no-memory"] = types.YChild{"ReceivedErrorNoMemory", &received.ReceivedErrorNoMemory}
    received.EntityData.Children["protect-protocol-received-good-request"] = types.YChild{"ProtectProtocolReceivedGoodRequest", &received.ProtectProtocolReceivedGoodRequest}
    received.EntityData.Children["protect-protocol-received-good-reply"] = types.YChild{"ProtectProtocolReceivedGoodReply", &received.ProtectProtocolReceivedGoodReply}
    received.EntityData.Children["received-good-bfd-request"] = types.YChild{"ReceivedGoodBfdRequest", &received.ReceivedGoodBfdRequest}
    received.EntityData.Children["received-good-bfd-reply"] = types.YChild{"ReceivedGoodBfdReply", &received.ReceivedGoodBfdReply}
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(received.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest
// Received good request
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodRequest) GetEntityData() *types.CommonEntityData {
    receivedGoodRequest.EntityData.YFilter = receivedGoodRequest.YFilter
    receivedGoodRequest.EntityData.YangName = "received-good-request"
    receivedGoodRequest.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodRequest.EntityData.ParentYangName = "received"
    receivedGoodRequest.EntityData.SegmentPath = "received-good-request"
    receivedGoodRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodRequest.EntityData.Children = make(map[string]types.YChild)
    receivedGoodRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodRequest.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodRequest.Packets}
    receivedGoodRequest.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodRequest.Bytes}
    return &(receivedGoodRequest.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply
// Received good reply
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodReply) GetEntityData() *types.CommonEntityData {
    receivedGoodReply.EntityData.YFilter = receivedGoodReply.YFilter
    receivedGoodReply.EntityData.YangName = "received-good-reply"
    receivedGoodReply.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodReply.EntityData.ParentYangName = "received"
    receivedGoodReply.EntityData.SegmentPath = "received-good-reply"
    receivedGoodReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodReply.EntityData.Children = make(map[string]types.YChild)
    receivedGoodReply.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodReply.Packets}
    receivedGoodReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodReply.Bytes}
    return &(receivedGoodReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown
// Received unknown packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedUnknown *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedUnknown) GetEntityData() *types.CommonEntityData {
    receivedUnknown.EntityData.YFilter = receivedUnknown.YFilter
    receivedUnknown.EntityData.YangName = "received-unknown"
    receivedUnknown.EntityData.BundleName = "cisco_ios_xr"
    receivedUnknown.EntityData.ParentYangName = "received"
    receivedUnknown.EntityData.SegmentPath = "received-unknown"
    receivedUnknown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedUnknown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedUnknown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedUnknown.EntityData.Children = make(map[string]types.YChild)
    receivedUnknown.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedUnknown.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedUnknown.Packets}
    receivedUnknown.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedUnknown.Bytes}
    return &(receivedUnknown.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader
// IP header error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorIpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorIpHeader) GetEntityData() *types.CommonEntityData {
    receivedErrorIpHeader.EntityData.YFilter = receivedErrorIpHeader.YFilter
    receivedErrorIpHeader.EntityData.YangName = "received-error-ip-header"
    receivedErrorIpHeader.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorIpHeader.EntityData.ParentYangName = "received"
    receivedErrorIpHeader.EntityData.SegmentPath = "received-error-ip-header"
    receivedErrorIpHeader.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorIpHeader.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorIpHeader.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorIpHeader.EntityData.Children = make(map[string]types.YChild)
    receivedErrorIpHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorIpHeader.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorIpHeader.Packets}
    receivedErrorIpHeader.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorIpHeader.Bytes}
    return &(receivedErrorIpHeader.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader
// UDP header error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorUdpHeader *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorUdpHeader) GetEntityData() *types.CommonEntityData {
    receivedErrorUdpHeader.EntityData.YFilter = receivedErrorUdpHeader.YFilter
    receivedErrorUdpHeader.EntityData.YangName = "received-error-udp-header"
    receivedErrorUdpHeader.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorUdpHeader.EntityData.ParentYangName = "received"
    receivedErrorUdpHeader.EntityData.SegmentPath = "received-error-udp-header"
    receivedErrorUdpHeader.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorUdpHeader.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorUdpHeader.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorUdpHeader.EntityData.Children = make(map[string]types.YChild)
    receivedErrorUdpHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorUdpHeader.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorUdpHeader.Packets}
    receivedErrorUdpHeader.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorUdpHeader.Bytes}
    return &(receivedErrorUdpHeader.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt
// RUNT error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorRunt *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorRunt) GetEntityData() *types.CommonEntityData {
    receivedErrorRunt.EntityData.YFilter = receivedErrorRunt.YFilter
    receivedErrorRunt.EntityData.YangName = "received-error-runt"
    receivedErrorRunt.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorRunt.EntityData.ParentYangName = "received"
    receivedErrorRunt.EntityData.SegmentPath = "received-error-runt"
    receivedErrorRunt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorRunt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorRunt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorRunt.EntityData.Children = make(map[string]types.YChild)
    receivedErrorRunt.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorRunt.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorRunt.Packets}
    receivedErrorRunt.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorRunt.Bytes}
    return &(receivedErrorRunt.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull
// Dropped queue full
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorQueueFull *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorQueueFull) GetEntityData() *types.CommonEntityData {
    receivedErrorQueueFull.EntityData.YFilter = receivedErrorQueueFull.YFilter
    receivedErrorQueueFull.EntityData.YangName = "received-error-queue-full"
    receivedErrorQueueFull.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorQueueFull.EntityData.ParentYangName = "received"
    receivedErrorQueueFull.EntityData.SegmentPath = "received-error-queue-full"
    receivedErrorQueueFull.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorQueueFull.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorQueueFull.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorQueueFull.EntityData.Children = make(map[string]types.YChild)
    receivedErrorQueueFull.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorQueueFull.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorQueueFull.Packets}
    receivedErrorQueueFull.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorQueueFull.Bytes}
    return &(receivedErrorQueueFull.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral
// General error
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorGeneral *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorGeneral) GetEntityData() *types.CommonEntityData {
    receivedErrorGeneral.EntityData.YFilter = receivedErrorGeneral.YFilter
    receivedErrorGeneral.EntityData.YangName = "received-error-general"
    receivedErrorGeneral.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorGeneral.EntityData.ParentYangName = "received"
    receivedErrorGeneral.EntityData.SegmentPath = "received-error-general"
    receivedErrorGeneral.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorGeneral.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorGeneral.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorGeneral.EntityData.Children = make(map[string]types.YChild)
    receivedErrorGeneral.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorGeneral.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorGeneral.Packets}
    receivedErrorGeneral.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorGeneral.Bytes}
    return &(receivedErrorGeneral.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface
// Error no Interfaces
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoInterface *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoInterface) GetEntityData() *types.CommonEntityData {
    receivedErrorNoInterface.EntityData.YFilter = receivedErrorNoInterface.YFilter
    receivedErrorNoInterface.EntityData.YangName = "received-error-no-interface"
    receivedErrorNoInterface.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorNoInterface.EntityData.ParentYangName = "received"
    receivedErrorNoInterface.EntityData.SegmentPath = "received-error-no-interface"
    receivedErrorNoInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorNoInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorNoInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorNoInterface.EntityData.Children = make(map[string]types.YChild)
    receivedErrorNoInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorNoInterface.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorNoInterface.Packets}
    receivedErrorNoInterface.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorNoInterface.Bytes}
    return &(receivedErrorNoInterface.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory
// Error no memory
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoMemory *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedErrorNoMemory) GetEntityData() *types.CommonEntityData {
    receivedErrorNoMemory.EntityData.YFilter = receivedErrorNoMemory.YFilter
    receivedErrorNoMemory.EntityData.YangName = "received-error-no-memory"
    receivedErrorNoMemory.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorNoMemory.EntityData.ParentYangName = "received"
    receivedErrorNoMemory.EntityData.SegmentPath = "received-error-no-memory"
    receivedErrorNoMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorNoMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorNoMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorNoMemory.EntityData.Children = make(map[string]types.YChild)
    receivedErrorNoMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorNoMemory.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorNoMemory.Packets}
    receivedErrorNoMemory.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorNoMemory.Bytes}
    return &(receivedErrorNoMemory.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest
// Protect Protocol Received good request
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodRequest) GetEntityData() *types.CommonEntityData {
    protectProtocolReceivedGoodRequest.EntityData.YFilter = protectProtocolReceivedGoodRequest.YFilter
    protectProtocolReceivedGoodRequest.EntityData.YangName = "protect-protocol-received-good-request"
    protectProtocolReceivedGoodRequest.EntityData.BundleName = "cisco_ios_xr"
    protectProtocolReceivedGoodRequest.EntityData.ParentYangName = "received"
    protectProtocolReceivedGoodRequest.EntityData.SegmentPath = "protect-protocol-received-good-request"
    protectProtocolReceivedGoodRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectProtocolReceivedGoodRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectProtocolReceivedGoodRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectProtocolReceivedGoodRequest.EntityData.Children = make(map[string]types.YChild)
    protectProtocolReceivedGoodRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    protectProtocolReceivedGoodRequest.EntityData.Leafs["packets"] = types.YLeaf{"Packets", protectProtocolReceivedGoodRequest.Packets}
    protectProtocolReceivedGoodRequest.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", protectProtocolReceivedGoodRequest.Bytes}
    return &(protectProtocolReceivedGoodRequest.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply
// Protect Protocol Received good reply
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ProtectProtocolReceivedGoodReply) GetEntityData() *types.CommonEntityData {
    protectProtocolReceivedGoodReply.EntityData.YFilter = protectProtocolReceivedGoodReply.YFilter
    protectProtocolReceivedGoodReply.EntityData.YangName = "protect-protocol-received-good-reply"
    protectProtocolReceivedGoodReply.EntityData.BundleName = "cisco_ios_xr"
    protectProtocolReceivedGoodReply.EntityData.ParentYangName = "received"
    protectProtocolReceivedGoodReply.EntityData.SegmentPath = "protect-protocol-received-good-reply"
    protectProtocolReceivedGoodReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectProtocolReceivedGoodReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectProtocolReceivedGoodReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectProtocolReceivedGoodReply.EntityData.Children = make(map[string]types.YChild)
    protectProtocolReceivedGoodReply.EntityData.Leafs = make(map[string]types.YLeaf)
    protectProtocolReceivedGoodReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", protectProtocolReceivedGoodReply.Packets}
    protectProtocolReceivedGoodReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", protectProtocolReceivedGoodReply.Bytes}
    return &(protectProtocolReceivedGoodReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest
// Received Reqeust with BFD TLV
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdRequest *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdRequest) GetEntityData() *types.CommonEntityData {
    receivedGoodBfdRequest.EntityData.YFilter = receivedGoodBfdRequest.YFilter
    receivedGoodBfdRequest.EntityData.YangName = "received-good-bfd-request"
    receivedGoodBfdRequest.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodBfdRequest.EntityData.ParentYangName = "received"
    receivedGoodBfdRequest.EntityData.SegmentPath = "received-good-bfd-request"
    receivedGoodBfdRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodBfdRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodBfdRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodBfdRequest.EntityData.Children = make(map[string]types.YChild)
    receivedGoodBfdRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodBfdRequest.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodBfdRequest.Packets}
    receivedGoodBfdRequest.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodBfdRequest.Bytes}
    return &(receivedGoodBfdRequest.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply
// Received Reply with BFD TLV
type MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdReply *MplsOam_Interface_Details_Detail_PacketStatistics_Received_ReceivedGoodBfdReply) GetEntityData() *types.CommonEntityData {
    receivedGoodBfdReply.EntityData.YFilter = receivedGoodBfdReply.YFilter
    receivedGoodBfdReply.EntityData.YangName = "received-good-bfd-reply"
    receivedGoodBfdReply.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodBfdReply.EntityData.ParentYangName = "received"
    receivedGoodBfdReply.EntityData.SegmentPath = "received-good-bfd-reply"
    receivedGoodBfdReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodBfdReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodBfdReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodBfdReply.EntityData.Children = make(map[string]types.YChild)
    receivedGoodBfdReply.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodBfdReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodBfdReply.Packets}
    receivedGoodBfdReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodBfdReply.Bytes}
    return &(receivedGoodBfdReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent
// Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply
}

func (sent *MplsOam_Interface_Details_Detail_PacketStatistics_Sent) GetEntityData() *types.CommonEntityData {
    sent.EntityData.YFilter = sent.YFilter
    sent.EntityData.YangName = "sent"
    sent.EntityData.BundleName = "cisco_ios_xr"
    sent.EntityData.ParentYangName = "packet-statistics"
    sent.EntityData.SegmentPath = "sent"
    sent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sent.EntityData.Children = make(map[string]types.YChild)
    sent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &sent.TransmitGood}
    sent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &sent.TransmitDrop}
    sent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &sent.TransmitBfdGood}
    sent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &sent.BfdNoReply}
    sent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sent.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_Sent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent
// Working Request Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply
}

func (workingReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent) GetEntityData() *types.CommonEntityData {
    workingReqSent.EntityData.YFilter = workingReqSent.YFilter
    workingReqSent.EntityData.YangName = "working-req-sent"
    workingReqSent.EntityData.BundleName = "cisco_ios_xr"
    workingReqSent.EntityData.ParentYangName = "packet-statistics"
    workingReqSent.EntityData.SegmentPath = "working-req-sent"
    workingReqSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    workingReqSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    workingReqSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    workingReqSent.EntityData.Children = make(map[string]types.YChild)
    workingReqSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &workingReqSent.TransmitGood}
    workingReqSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &workingReqSent.TransmitDrop}
    workingReqSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &workingReqSent.TransmitBfdGood}
    workingReqSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &workingReqSent.BfdNoReply}
    workingReqSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(workingReqSent.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "working-req-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "working-req-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "working-req-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingReqSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "working-req-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent
// Working Reply Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply
}

func (workingRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent) GetEntityData() *types.CommonEntityData {
    workingRepSent.EntityData.YFilter = workingRepSent.YFilter
    workingRepSent.EntityData.YangName = "working-rep-sent"
    workingRepSent.EntityData.BundleName = "cisco_ios_xr"
    workingRepSent.EntityData.ParentYangName = "packet-statistics"
    workingRepSent.EntityData.SegmentPath = "working-rep-sent"
    workingRepSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    workingRepSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    workingRepSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    workingRepSent.EntityData.Children = make(map[string]types.YChild)
    workingRepSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &workingRepSent.TransmitGood}
    workingRepSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &workingRepSent.TransmitDrop}
    workingRepSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &workingRepSent.TransmitBfdGood}
    workingRepSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &workingRepSent.BfdNoReply}
    workingRepSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(workingRepSent.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "working-rep-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "working-rep-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "working-rep-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_WorkingRepSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "working-rep-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent
// Protect Request Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply
}

func (protectReqSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent) GetEntityData() *types.CommonEntityData {
    protectReqSent.EntityData.YFilter = protectReqSent.YFilter
    protectReqSent.EntityData.YangName = "protect-req-sent"
    protectReqSent.EntityData.BundleName = "cisco_ios_xr"
    protectReqSent.EntityData.ParentYangName = "packet-statistics"
    protectReqSent.EntityData.SegmentPath = "protect-req-sent"
    protectReqSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectReqSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectReqSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectReqSent.EntityData.Children = make(map[string]types.YChild)
    protectReqSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &protectReqSent.TransmitGood}
    protectReqSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &protectReqSent.TransmitDrop}
    protectReqSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &protectReqSent.TransmitBfdGood}
    protectReqSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &protectReqSent.BfdNoReply}
    protectReqSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protectReqSent.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "protect-req-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "protect-req-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "protect-req-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectReqSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "protect-req-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent
// Protect Reply Packet transmit counts
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply
}

func (protectRepSent *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent) GetEntityData() *types.CommonEntityData {
    protectRepSent.EntityData.YFilter = protectRepSent.YFilter
    protectRepSent.EntityData.YangName = "protect-rep-sent"
    protectRepSent.EntityData.BundleName = "cisco_ios_xr"
    protectRepSent.EntityData.ParentYangName = "packet-statistics"
    protectRepSent.EntityData.SegmentPath = "protect-rep-sent"
    protectRepSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectRepSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectRepSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectRepSent.EntityData.Children = make(map[string]types.YChild)
    protectRepSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &protectRepSent.TransmitGood}
    protectRepSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &protectRepSent.TransmitDrop}
    protectRepSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &protectRepSent.TransmitBfdGood}
    protectRepSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &protectRepSent.BfdNoReply}
    protectRepSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protectRepSent.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood
// Transmit good packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "protect-rep-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "protect-rep-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "protect-rep-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Interface_Details_Detail_PacketStatistics_ProtectRepSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "protect-rep-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Packet
// LSPV packet counters operational data
type MplsOam_Packet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet reception counts.
    Received MplsOam_Packet_Received

    // Packet transmit counts.
    Sent MplsOam_Packet_Sent

    // Working Request Packet transmit counts.
    WorkingReqSent MplsOam_Packet_WorkingReqSent

    // Working Reply Packet transmit counts.
    WorkingRepSent MplsOam_Packet_WorkingRepSent

    // Protect Request Packet transmit counts.
    ProtectReqSent MplsOam_Packet_ProtectReqSent

    // Protect Reply Packet transmit counts.
    ProtectRepSent MplsOam_Packet_ProtectRepSent
}

func (packet *MplsOam_Packet) GetEntityData() *types.CommonEntityData {
    packet.EntityData.YFilter = packet.YFilter
    packet.EntityData.YangName = "packet"
    packet.EntityData.BundleName = "cisco_ios_xr"
    packet.EntityData.ParentYangName = "mpls-oam"
    packet.EntityData.SegmentPath = "packet"
    packet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packet.EntityData.Children = make(map[string]types.YChild)
    packet.EntityData.Children["received"] = types.YChild{"Received", &packet.Received}
    packet.EntityData.Children["sent"] = types.YChild{"Sent", &packet.Sent}
    packet.EntityData.Children["working-req-sent"] = types.YChild{"WorkingReqSent", &packet.WorkingReqSent}
    packet.EntityData.Children["working-rep-sent"] = types.YChild{"WorkingRepSent", &packet.WorkingRepSent}
    packet.EntityData.Children["protect-req-sent"] = types.YChild{"ProtectReqSent", &packet.ProtectReqSent}
    packet.EntityData.Children["protect-rep-sent"] = types.YChild{"ProtectRepSent", &packet.ProtectRepSent}
    packet.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(packet.EntityData)
}

// MplsOam_Packet_Received
// Packet reception counts
type MplsOam_Packet_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received good request.
    ReceivedGoodRequest MplsOam_Packet_Received_ReceivedGoodRequest

    // Received good reply.
    ReceivedGoodReply MplsOam_Packet_Received_ReceivedGoodReply

    // Received unknown packets.
    ReceivedUnknown MplsOam_Packet_Received_ReceivedUnknown

    // IP header error.
    ReceivedErrorIpHeader MplsOam_Packet_Received_ReceivedErrorIpHeader

    // UDP header error.
    ReceivedErrorUdpHeader MplsOam_Packet_Received_ReceivedErrorUdpHeader

    // RUNT error.
    ReceivedErrorRunt MplsOam_Packet_Received_ReceivedErrorRunt

    // Dropped queue full.
    ReceivedErrorQueueFull MplsOam_Packet_Received_ReceivedErrorQueueFull

    // General error.
    ReceivedErrorGeneral MplsOam_Packet_Received_ReceivedErrorGeneral

    // Error no Interfaces.
    ReceivedErrorNoInterface MplsOam_Packet_Received_ReceivedErrorNoInterface

    // Error no memory.
    ReceivedErrorNoMemory MplsOam_Packet_Received_ReceivedErrorNoMemory

    // Protect Protocol Received good request.
    ProtectProtocolReceivedGoodRequest MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest

    // Protect Protocol Received good reply.
    ProtectProtocolReceivedGoodReply MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply

    // Received Reqeust with BFD TLV.
    ReceivedGoodBfdRequest MplsOam_Packet_Received_ReceivedGoodBfdRequest

    // Received Reply with BFD TLV.
    ReceivedGoodBfdReply MplsOam_Packet_Received_ReceivedGoodBfdReply
}

func (received *MplsOam_Packet_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xr"
    received.EntityData.ParentYangName = "packet"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Children["received-good-request"] = types.YChild{"ReceivedGoodRequest", &received.ReceivedGoodRequest}
    received.EntityData.Children["received-good-reply"] = types.YChild{"ReceivedGoodReply", &received.ReceivedGoodReply}
    received.EntityData.Children["received-unknown"] = types.YChild{"ReceivedUnknown", &received.ReceivedUnknown}
    received.EntityData.Children["received-error-ip-header"] = types.YChild{"ReceivedErrorIpHeader", &received.ReceivedErrorIpHeader}
    received.EntityData.Children["received-error-udp-header"] = types.YChild{"ReceivedErrorUdpHeader", &received.ReceivedErrorUdpHeader}
    received.EntityData.Children["received-error-runt"] = types.YChild{"ReceivedErrorRunt", &received.ReceivedErrorRunt}
    received.EntityData.Children["received-error-queue-full"] = types.YChild{"ReceivedErrorQueueFull", &received.ReceivedErrorQueueFull}
    received.EntityData.Children["received-error-general"] = types.YChild{"ReceivedErrorGeneral", &received.ReceivedErrorGeneral}
    received.EntityData.Children["received-error-no-interface"] = types.YChild{"ReceivedErrorNoInterface", &received.ReceivedErrorNoInterface}
    received.EntityData.Children["received-error-no-memory"] = types.YChild{"ReceivedErrorNoMemory", &received.ReceivedErrorNoMemory}
    received.EntityData.Children["protect-protocol-received-good-request"] = types.YChild{"ProtectProtocolReceivedGoodRequest", &received.ProtectProtocolReceivedGoodRequest}
    received.EntityData.Children["protect-protocol-received-good-reply"] = types.YChild{"ProtectProtocolReceivedGoodReply", &received.ProtectProtocolReceivedGoodReply}
    received.EntityData.Children["received-good-bfd-request"] = types.YChild{"ReceivedGoodBfdRequest", &received.ReceivedGoodBfdRequest}
    received.EntityData.Children["received-good-bfd-reply"] = types.YChild{"ReceivedGoodBfdReply", &received.ReceivedGoodBfdReply}
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(received.EntityData)
}

// MplsOam_Packet_Received_ReceivedGoodRequest
// Received good request
type MplsOam_Packet_Received_ReceivedGoodRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodRequest *MplsOam_Packet_Received_ReceivedGoodRequest) GetEntityData() *types.CommonEntityData {
    receivedGoodRequest.EntityData.YFilter = receivedGoodRequest.YFilter
    receivedGoodRequest.EntityData.YangName = "received-good-request"
    receivedGoodRequest.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodRequest.EntityData.ParentYangName = "received"
    receivedGoodRequest.EntityData.SegmentPath = "received-good-request"
    receivedGoodRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodRequest.EntityData.Children = make(map[string]types.YChild)
    receivedGoodRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodRequest.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodRequest.Packets}
    receivedGoodRequest.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodRequest.Bytes}
    return &(receivedGoodRequest.EntityData)
}

// MplsOam_Packet_Received_ReceivedGoodReply
// Received good reply
type MplsOam_Packet_Received_ReceivedGoodReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodReply *MplsOam_Packet_Received_ReceivedGoodReply) GetEntityData() *types.CommonEntityData {
    receivedGoodReply.EntityData.YFilter = receivedGoodReply.YFilter
    receivedGoodReply.EntityData.YangName = "received-good-reply"
    receivedGoodReply.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodReply.EntityData.ParentYangName = "received"
    receivedGoodReply.EntityData.SegmentPath = "received-good-reply"
    receivedGoodReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodReply.EntityData.Children = make(map[string]types.YChild)
    receivedGoodReply.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodReply.Packets}
    receivedGoodReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodReply.Bytes}
    return &(receivedGoodReply.EntityData)
}

// MplsOam_Packet_Received_ReceivedUnknown
// Received unknown packets
type MplsOam_Packet_Received_ReceivedUnknown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedUnknown *MplsOam_Packet_Received_ReceivedUnknown) GetEntityData() *types.CommonEntityData {
    receivedUnknown.EntityData.YFilter = receivedUnknown.YFilter
    receivedUnknown.EntityData.YangName = "received-unknown"
    receivedUnknown.EntityData.BundleName = "cisco_ios_xr"
    receivedUnknown.EntityData.ParentYangName = "received"
    receivedUnknown.EntityData.SegmentPath = "received-unknown"
    receivedUnknown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedUnknown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedUnknown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedUnknown.EntityData.Children = make(map[string]types.YChild)
    receivedUnknown.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedUnknown.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedUnknown.Packets}
    receivedUnknown.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedUnknown.Bytes}
    return &(receivedUnknown.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorIpHeader
// IP header error
type MplsOam_Packet_Received_ReceivedErrorIpHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorIpHeader *MplsOam_Packet_Received_ReceivedErrorIpHeader) GetEntityData() *types.CommonEntityData {
    receivedErrorIpHeader.EntityData.YFilter = receivedErrorIpHeader.YFilter
    receivedErrorIpHeader.EntityData.YangName = "received-error-ip-header"
    receivedErrorIpHeader.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorIpHeader.EntityData.ParentYangName = "received"
    receivedErrorIpHeader.EntityData.SegmentPath = "received-error-ip-header"
    receivedErrorIpHeader.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorIpHeader.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorIpHeader.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorIpHeader.EntityData.Children = make(map[string]types.YChild)
    receivedErrorIpHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorIpHeader.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorIpHeader.Packets}
    receivedErrorIpHeader.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorIpHeader.Bytes}
    return &(receivedErrorIpHeader.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorUdpHeader
// UDP header error
type MplsOam_Packet_Received_ReceivedErrorUdpHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorUdpHeader *MplsOam_Packet_Received_ReceivedErrorUdpHeader) GetEntityData() *types.CommonEntityData {
    receivedErrorUdpHeader.EntityData.YFilter = receivedErrorUdpHeader.YFilter
    receivedErrorUdpHeader.EntityData.YangName = "received-error-udp-header"
    receivedErrorUdpHeader.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorUdpHeader.EntityData.ParentYangName = "received"
    receivedErrorUdpHeader.EntityData.SegmentPath = "received-error-udp-header"
    receivedErrorUdpHeader.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorUdpHeader.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorUdpHeader.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorUdpHeader.EntityData.Children = make(map[string]types.YChild)
    receivedErrorUdpHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorUdpHeader.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorUdpHeader.Packets}
    receivedErrorUdpHeader.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorUdpHeader.Bytes}
    return &(receivedErrorUdpHeader.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorRunt
// RUNT error
type MplsOam_Packet_Received_ReceivedErrorRunt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorRunt *MplsOam_Packet_Received_ReceivedErrorRunt) GetEntityData() *types.CommonEntityData {
    receivedErrorRunt.EntityData.YFilter = receivedErrorRunt.YFilter
    receivedErrorRunt.EntityData.YangName = "received-error-runt"
    receivedErrorRunt.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorRunt.EntityData.ParentYangName = "received"
    receivedErrorRunt.EntityData.SegmentPath = "received-error-runt"
    receivedErrorRunt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorRunt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorRunt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorRunt.EntityData.Children = make(map[string]types.YChild)
    receivedErrorRunt.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorRunt.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorRunt.Packets}
    receivedErrorRunt.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorRunt.Bytes}
    return &(receivedErrorRunt.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorQueueFull
// Dropped queue full
type MplsOam_Packet_Received_ReceivedErrorQueueFull struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorQueueFull *MplsOam_Packet_Received_ReceivedErrorQueueFull) GetEntityData() *types.CommonEntityData {
    receivedErrorQueueFull.EntityData.YFilter = receivedErrorQueueFull.YFilter
    receivedErrorQueueFull.EntityData.YangName = "received-error-queue-full"
    receivedErrorQueueFull.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorQueueFull.EntityData.ParentYangName = "received"
    receivedErrorQueueFull.EntityData.SegmentPath = "received-error-queue-full"
    receivedErrorQueueFull.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorQueueFull.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorQueueFull.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorQueueFull.EntityData.Children = make(map[string]types.YChild)
    receivedErrorQueueFull.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorQueueFull.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorQueueFull.Packets}
    receivedErrorQueueFull.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorQueueFull.Bytes}
    return &(receivedErrorQueueFull.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorGeneral
// General error
type MplsOam_Packet_Received_ReceivedErrorGeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorGeneral *MplsOam_Packet_Received_ReceivedErrorGeneral) GetEntityData() *types.CommonEntityData {
    receivedErrorGeneral.EntityData.YFilter = receivedErrorGeneral.YFilter
    receivedErrorGeneral.EntityData.YangName = "received-error-general"
    receivedErrorGeneral.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorGeneral.EntityData.ParentYangName = "received"
    receivedErrorGeneral.EntityData.SegmentPath = "received-error-general"
    receivedErrorGeneral.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorGeneral.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorGeneral.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorGeneral.EntityData.Children = make(map[string]types.YChild)
    receivedErrorGeneral.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorGeneral.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorGeneral.Packets}
    receivedErrorGeneral.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorGeneral.Bytes}
    return &(receivedErrorGeneral.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorNoInterface
// Error no Interfaces
type MplsOam_Packet_Received_ReceivedErrorNoInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoInterface *MplsOam_Packet_Received_ReceivedErrorNoInterface) GetEntityData() *types.CommonEntityData {
    receivedErrorNoInterface.EntityData.YFilter = receivedErrorNoInterface.YFilter
    receivedErrorNoInterface.EntityData.YangName = "received-error-no-interface"
    receivedErrorNoInterface.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorNoInterface.EntityData.ParentYangName = "received"
    receivedErrorNoInterface.EntityData.SegmentPath = "received-error-no-interface"
    receivedErrorNoInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorNoInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorNoInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorNoInterface.EntityData.Children = make(map[string]types.YChild)
    receivedErrorNoInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorNoInterface.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorNoInterface.Packets}
    receivedErrorNoInterface.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorNoInterface.Bytes}
    return &(receivedErrorNoInterface.EntityData)
}

// MplsOam_Packet_Received_ReceivedErrorNoMemory
// Error no memory
type MplsOam_Packet_Received_ReceivedErrorNoMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedErrorNoMemory *MplsOam_Packet_Received_ReceivedErrorNoMemory) GetEntityData() *types.CommonEntityData {
    receivedErrorNoMemory.EntityData.YFilter = receivedErrorNoMemory.YFilter
    receivedErrorNoMemory.EntityData.YangName = "received-error-no-memory"
    receivedErrorNoMemory.EntityData.BundleName = "cisco_ios_xr"
    receivedErrorNoMemory.EntityData.ParentYangName = "received"
    receivedErrorNoMemory.EntityData.SegmentPath = "received-error-no-memory"
    receivedErrorNoMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedErrorNoMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedErrorNoMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedErrorNoMemory.EntityData.Children = make(map[string]types.YChild)
    receivedErrorNoMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedErrorNoMemory.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedErrorNoMemory.Packets}
    receivedErrorNoMemory.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedErrorNoMemory.Bytes}
    return &(receivedErrorNoMemory.EntityData)
}

// MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest
// Protect Protocol Received good request
type MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodRequest *MplsOam_Packet_Received_ProtectProtocolReceivedGoodRequest) GetEntityData() *types.CommonEntityData {
    protectProtocolReceivedGoodRequest.EntityData.YFilter = protectProtocolReceivedGoodRequest.YFilter
    protectProtocolReceivedGoodRequest.EntityData.YangName = "protect-protocol-received-good-request"
    protectProtocolReceivedGoodRequest.EntityData.BundleName = "cisco_ios_xr"
    protectProtocolReceivedGoodRequest.EntityData.ParentYangName = "received"
    protectProtocolReceivedGoodRequest.EntityData.SegmentPath = "protect-protocol-received-good-request"
    protectProtocolReceivedGoodRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectProtocolReceivedGoodRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectProtocolReceivedGoodRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectProtocolReceivedGoodRequest.EntityData.Children = make(map[string]types.YChild)
    protectProtocolReceivedGoodRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    protectProtocolReceivedGoodRequest.EntityData.Leafs["packets"] = types.YLeaf{"Packets", protectProtocolReceivedGoodRequest.Packets}
    protectProtocolReceivedGoodRequest.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", protectProtocolReceivedGoodRequest.Bytes}
    return &(protectProtocolReceivedGoodRequest.EntityData)
}

// MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply
// Protect Protocol Received good reply
type MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (protectProtocolReceivedGoodReply *MplsOam_Packet_Received_ProtectProtocolReceivedGoodReply) GetEntityData() *types.CommonEntityData {
    protectProtocolReceivedGoodReply.EntityData.YFilter = protectProtocolReceivedGoodReply.YFilter
    protectProtocolReceivedGoodReply.EntityData.YangName = "protect-protocol-received-good-reply"
    protectProtocolReceivedGoodReply.EntityData.BundleName = "cisco_ios_xr"
    protectProtocolReceivedGoodReply.EntityData.ParentYangName = "received"
    protectProtocolReceivedGoodReply.EntityData.SegmentPath = "protect-protocol-received-good-reply"
    protectProtocolReceivedGoodReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectProtocolReceivedGoodReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectProtocolReceivedGoodReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectProtocolReceivedGoodReply.EntityData.Children = make(map[string]types.YChild)
    protectProtocolReceivedGoodReply.EntityData.Leafs = make(map[string]types.YLeaf)
    protectProtocolReceivedGoodReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", protectProtocolReceivedGoodReply.Packets}
    protectProtocolReceivedGoodReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", protectProtocolReceivedGoodReply.Bytes}
    return &(protectProtocolReceivedGoodReply.EntityData)
}

// MplsOam_Packet_Received_ReceivedGoodBfdRequest
// Received Reqeust with BFD TLV
type MplsOam_Packet_Received_ReceivedGoodBfdRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdRequest *MplsOam_Packet_Received_ReceivedGoodBfdRequest) GetEntityData() *types.CommonEntityData {
    receivedGoodBfdRequest.EntityData.YFilter = receivedGoodBfdRequest.YFilter
    receivedGoodBfdRequest.EntityData.YangName = "received-good-bfd-request"
    receivedGoodBfdRequest.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodBfdRequest.EntityData.ParentYangName = "received"
    receivedGoodBfdRequest.EntityData.SegmentPath = "received-good-bfd-request"
    receivedGoodBfdRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodBfdRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodBfdRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodBfdRequest.EntityData.Children = make(map[string]types.YChild)
    receivedGoodBfdRequest.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodBfdRequest.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodBfdRequest.Packets}
    receivedGoodBfdRequest.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodBfdRequest.Bytes}
    return &(receivedGoodBfdRequest.EntityData)
}

// MplsOam_Packet_Received_ReceivedGoodBfdReply
// Received Reply with BFD TLV
type MplsOam_Packet_Received_ReceivedGoodBfdReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (receivedGoodBfdReply *MplsOam_Packet_Received_ReceivedGoodBfdReply) GetEntityData() *types.CommonEntityData {
    receivedGoodBfdReply.EntityData.YFilter = receivedGoodBfdReply.YFilter
    receivedGoodBfdReply.EntityData.YangName = "received-good-bfd-reply"
    receivedGoodBfdReply.EntityData.BundleName = "cisco_ios_xr"
    receivedGoodBfdReply.EntityData.ParentYangName = "received"
    receivedGoodBfdReply.EntityData.SegmentPath = "received-good-bfd-reply"
    receivedGoodBfdReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivedGoodBfdReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivedGoodBfdReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivedGoodBfdReply.EntityData.Children = make(map[string]types.YChild)
    receivedGoodBfdReply.EntityData.Leafs = make(map[string]types.YLeaf)
    receivedGoodBfdReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", receivedGoodBfdReply.Packets}
    receivedGoodBfdReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", receivedGoodBfdReply.Bytes}
    return &(receivedGoodBfdReply.EntityData)
}

// MplsOam_Packet_Sent
// Packet transmit counts
type MplsOam_Packet_Sent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_Sent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_Sent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_Sent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_Sent_BfdNoReply
}

func (sent *MplsOam_Packet_Sent) GetEntityData() *types.CommonEntityData {
    sent.EntityData.YFilter = sent.YFilter
    sent.EntityData.YangName = "sent"
    sent.EntityData.BundleName = "cisco_ios_xr"
    sent.EntityData.ParentYangName = "packet"
    sent.EntityData.SegmentPath = "sent"
    sent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sent.EntityData.Children = make(map[string]types.YChild)
    sent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &sent.TransmitGood}
    sent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &sent.TransmitDrop}
    sent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &sent.TransmitBfdGood}
    sent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &sent.BfdNoReply}
    sent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sent.EntityData)
}

// MplsOam_Packet_Sent_TransmitGood
// Transmit good packets
type MplsOam_Packet_Sent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_Sent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Packet_Sent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_Sent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_Sent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Packet_Sent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_Sent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_Sent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Packet_Sent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_Sent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_Sent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Packet_WorkingReqSent
// Working Request Packet transmit counts
type MplsOam_Packet_WorkingReqSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_WorkingReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_WorkingReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_WorkingReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_WorkingReqSent_BfdNoReply
}

func (workingReqSent *MplsOam_Packet_WorkingReqSent) GetEntityData() *types.CommonEntityData {
    workingReqSent.EntityData.YFilter = workingReqSent.YFilter
    workingReqSent.EntityData.YangName = "working-req-sent"
    workingReqSent.EntityData.BundleName = "cisco_ios_xr"
    workingReqSent.EntityData.ParentYangName = "packet"
    workingReqSent.EntityData.SegmentPath = "working-req-sent"
    workingReqSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    workingReqSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    workingReqSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    workingReqSent.EntityData.Children = make(map[string]types.YChild)
    workingReqSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &workingReqSent.TransmitGood}
    workingReqSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &workingReqSent.TransmitDrop}
    workingReqSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &workingReqSent.TransmitBfdGood}
    workingReqSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &workingReqSent.BfdNoReply}
    workingReqSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(workingReqSent.EntityData)
}

// MplsOam_Packet_WorkingReqSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_WorkingReqSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_WorkingReqSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "working-req-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Packet_WorkingReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_WorkingReqSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_WorkingReqSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "working-req-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Packet_WorkingReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_WorkingReqSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_WorkingReqSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "working-req-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Packet_WorkingReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_WorkingReqSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_WorkingReqSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "working-req-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Packet_WorkingRepSent
// Working Reply Packet transmit counts
type MplsOam_Packet_WorkingRepSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_WorkingRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_WorkingRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_WorkingRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_WorkingRepSent_BfdNoReply
}

func (workingRepSent *MplsOam_Packet_WorkingRepSent) GetEntityData() *types.CommonEntityData {
    workingRepSent.EntityData.YFilter = workingRepSent.YFilter
    workingRepSent.EntityData.YangName = "working-rep-sent"
    workingRepSent.EntityData.BundleName = "cisco_ios_xr"
    workingRepSent.EntityData.ParentYangName = "packet"
    workingRepSent.EntityData.SegmentPath = "working-rep-sent"
    workingRepSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    workingRepSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    workingRepSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    workingRepSent.EntityData.Children = make(map[string]types.YChild)
    workingRepSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &workingRepSent.TransmitGood}
    workingRepSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &workingRepSent.TransmitDrop}
    workingRepSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &workingRepSent.TransmitBfdGood}
    workingRepSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &workingRepSent.BfdNoReply}
    workingRepSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(workingRepSent.EntityData)
}

// MplsOam_Packet_WorkingRepSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_WorkingRepSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_WorkingRepSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "working-rep-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Packet_WorkingRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_WorkingRepSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_WorkingRepSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "working-rep-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Packet_WorkingRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_WorkingRepSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_WorkingRepSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "working-rep-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Packet_WorkingRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_WorkingRepSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_WorkingRepSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "working-rep-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Packet_ProtectReqSent
// Protect Request Packet transmit counts
type MplsOam_Packet_ProtectReqSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_ProtectReqSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_ProtectReqSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_ProtectReqSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_ProtectReqSent_BfdNoReply
}

func (protectReqSent *MplsOam_Packet_ProtectReqSent) GetEntityData() *types.CommonEntityData {
    protectReqSent.EntityData.YFilter = protectReqSent.YFilter
    protectReqSent.EntityData.YangName = "protect-req-sent"
    protectReqSent.EntityData.BundleName = "cisco_ios_xr"
    protectReqSent.EntityData.ParentYangName = "packet"
    protectReqSent.EntityData.SegmentPath = "protect-req-sent"
    protectReqSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectReqSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectReqSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectReqSent.EntityData.Children = make(map[string]types.YChild)
    protectReqSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &protectReqSent.TransmitGood}
    protectReqSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &protectReqSent.TransmitDrop}
    protectReqSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &protectReqSent.TransmitBfdGood}
    protectReqSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &protectReqSent.BfdNoReply}
    protectReqSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protectReqSent.EntityData)
}

// MplsOam_Packet_ProtectReqSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_ProtectReqSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_ProtectReqSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "protect-req-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Packet_ProtectReqSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_ProtectReqSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_ProtectReqSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "protect-req-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Packet_ProtectReqSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_ProtectReqSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_ProtectReqSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "protect-req-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Packet_ProtectReqSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_ProtectReqSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_ProtectReqSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "protect-req-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Packet_ProtectRepSent
// Protect Reply Packet transmit counts
type MplsOam_Packet_ProtectRepSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmit good packets.
    TransmitGood MplsOam_Packet_ProtectRepSent_TransmitGood

    // Transmit drop packets.
    TransmitDrop MplsOam_Packet_ProtectRepSent_TransmitDrop

    // Transmit good BFD request packets.
    TransmitBfdGood MplsOam_Packet_ProtectRepSent_TransmitBfdGood

    // No Reply action for echo reqeust of BFD bootstrap.
    BfdNoReply MplsOam_Packet_ProtectRepSent_BfdNoReply
}

func (protectRepSent *MplsOam_Packet_ProtectRepSent) GetEntityData() *types.CommonEntityData {
    protectRepSent.EntityData.YFilter = protectRepSent.YFilter
    protectRepSent.EntityData.YangName = "protect-rep-sent"
    protectRepSent.EntityData.BundleName = "cisco_ios_xr"
    protectRepSent.EntityData.ParentYangName = "packet"
    protectRepSent.EntityData.SegmentPath = "protect-rep-sent"
    protectRepSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectRepSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectRepSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectRepSent.EntityData.Children = make(map[string]types.YChild)
    protectRepSent.EntityData.Children["transmit-good"] = types.YChild{"TransmitGood", &protectRepSent.TransmitGood}
    protectRepSent.EntityData.Children["transmit-drop"] = types.YChild{"TransmitDrop", &protectRepSent.TransmitDrop}
    protectRepSent.EntityData.Children["transmit-bfd-good"] = types.YChild{"TransmitBfdGood", &protectRepSent.TransmitBfdGood}
    protectRepSent.EntityData.Children["bfd-no-reply"] = types.YChild{"BfdNoReply", &protectRepSent.BfdNoReply}
    protectRepSent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protectRepSent.EntityData)
}

// MplsOam_Packet_ProtectRepSent_TransmitGood
// Transmit good packets
type MplsOam_Packet_ProtectRepSent_TransmitGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitGood *MplsOam_Packet_ProtectRepSent_TransmitGood) GetEntityData() *types.CommonEntityData {
    transmitGood.EntityData.YFilter = transmitGood.YFilter
    transmitGood.EntityData.YangName = "transmit-good"
    transmitGood.EntityData.BundleName = "cisco_ios_xr"
    transmitGood.EntityData.ParentYangName = "protect-rep-sent"
    transmitGood.EntityData.SegmentPath = "transmit-good"
    transmitGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitGood.EntityData.Children = make(map[string]types.YChild)
    transmitGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitGood.Packets}
    transmitGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitGood.Bytes}
    return &(transmitGood.EntityData)
}

// MplsOam_Packet_ProtectRepSent_TransmitDrop
// Transmit drop packets
type MplsOam_Packet_ProtectRepSent_TransmitDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitDrop *MplsOam_Packet_ProtectRepSent_TransmitDrop) GetEntityData() *types.CommonEntityData {
    transmitDrop.EntityData.YFilter = transmitDrop.YFilter
    transmitDrop.EntityData.YangName = "transmit-drop"
    transmitDrop.EntityData.BundleName = "cisco_ios_xr"
    transmitDrop.EntityData.ParentYangName = "protect-rep-sent"
    transmitDrop.EntityData.SegmentPath = "transmit-drop"
    transmitDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitDrop.EntityData.Children = make(map[string]types.YChild)
    transmitDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitDrop.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitDrop.Packets}
    transmitDrop.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitDrop.Bytes}
    return &(transmitDrop.EntityData)
}

// MplsOam_Packet_ProtectRepSent_TransmitBfdGood
// Transmit good BFD request packets
type MplsOam_Packet_ProtectRepSent_TransmitBfdGood struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (transmitBfdGood *MplsOam_Packet_ProtectRepSent_TransmitBfdGood) GetEntityData() *types.CommonEntityData {
    transmitBfdGood.EntityData.YFilter = transmitBfdGood.YFilter
    transmitBfdGood.EntityData.YangName = "transmit-bfd-good"
    transmitBfdGood.EntityData.BundleName = "cisco_ios_xr"
    transmitBfdGood.EntityData.ParentYangName = "protect-rep-sent"
    transmitBfdGood.EntityData.SegmentPath = "transmit-bfd-good"
    transmitBfdGood.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitBfdGood.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitBfdGood.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitBfdGood.EntityData.Children = make(map[string]types.YChild)
    transmitBfdGood.EntityData.Leafs = make(map[string]types.YLeaf)
    transmitBfdGood.EntityData.Leafs["packets"] = types.YLeaf{"Packets", transmitBfdGood.Packets}
    transmitBfdGood.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", transmitBfdGood.Bytes}
    return &(transmitBfdGood.EntityData)
}

// MplsOam_Packet_ProtectRepSent_BfdNoReply
// No Reply action for echo reqeust of BFD
// bootstrap
type MplsOam_Packet_ProtectRepSent_BfdNoReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Packets interface{}

    // Byte counter. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    Bytes interface{}
}

func (bfdNoReply *MplsOam_Packet_ProtectRepSent_BfdNoReply) GetEntityData() *types.CommonEntityData {
    bfdNoReply.EntityData.YFilter = bfdNoReply.YFilter
    bfdNoReply.EntityData.YangName = "bfd-no-reply"
    bfdNoReply.EntityData.BundleName = "cisco_ios_xr"
    bfdNoReply.EntityData.ParentYangName = "protect-rep-sent"
    bfdNoReply.EntityData.SegmentPath = "bfd-no-reply"
    bfdNoReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdNoReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdNoReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdNoReply.EntityData.Children = make(map[string]types.YChild)
    bfdNoReply.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNoReply.EntityData.Leafs["packets"] = types.YLeaf{"Packets", bfdNoReply.Packets}
    bfdNoReply.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", bfdNoReply.Bytes}
    return &(bfdNoReply.EntityData)
}

// MplsOam_Global
// LSPV global counters operational data
type MplsOam_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of clients. The type is interface{} with range: 0..4294967295.
    TotalClients interface{}

    // Message statistics.
    MessageStatistics MplsOam_Global_MessageStatistics

    // Collaborator statistics.
    CollaboratorStatistics MplsOam_Global_CollaboratorStatistics
}

func (global *MplsOam_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "mpls-oam"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["message-statistics"] = types.YChild{"MessageStatistics", &global.MessageStatistics}
    global.EntityData.Children["collaborator-statistics"] = types.YChild{"CollaboratorStatistics", &global.CollaboratorStatistics}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["total-clients"] = types.YLeaf{"TotalClients", global.TotalClients}
    return &(global.EntityData)
}

// MplsOam_Global_MessageStatistics
// Message statistics
type MplsOam_Global_MessageStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Message register count. The type is interface{} with range: 0..4294967295.
    RegisterMessages interface{}

    // Message unregister count. The type is interface{} with range:
    // 0..4294967295.
    UnregisterMessages interface{}

    // Message echo submit count. The type is interface{} with range:
    // 0..4294967295.
    EchoSubmitMessages interface{}

    // Message echo cancel count. The type is interface{} with range:
    // 0..4294967295.
    EchoCancelMessages interface{}

    // Message get results count. The type is interface{} with range:
    // 0..4294967295.
    GetResultMessages interface{}

    // Message get configiuration count. The type is interface{} with range:
    // 0..4294967295.
    GetConfigMessages interface{}

    // Message get response count. The type is interface{} with range:
    // 0..4294967295.
    GetResponseMessages interface{}

    // Message property response count. The type is interface{} with range:
    // 0..4294967295.
    PropertyResponseMessages interface{}

    // Message property request count. The type is interface{} with range:
    // 0..4294967295.
    PropertyRequestMessages interface{}

    // Message property block count. The type is interface{} with range:
    // 0..4294967295.
    PropertyBlockMessages interface{}

    // Message thread request count. The type is interface{} with range:
    // 0..4294967295.
    ThreadRequestMessages interface{}
}

func (messageStatistics *MplsOam_Global_MessageStatistics) GetEntityData() *types.CommonEntityData {
    messageStatistics.EntityData.YFilter = messageStatistics.YFilter
    messageStatistics.EntityData.YangName = "message-statistics"
    messageStatistics.EntityData.BundleName = "cisco_ios_xr"
    messageStatistics.EntityData.ParentYangName = "global"
    messageStatistics.EntityData.SegmentPath = "message-statistics"
    messageStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    messageStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    messageStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    messageStatistics.EntityData.Children = make(map[string]types.YChild)
    messageStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    messageStatistics.EntityData.Leafs["register-messages"] = types.YLeaf{"RegisterMessages", messageStatistics.RegisterMessages}
    messageStatistics.EntityData.Leafs["unregister-messages"] = types.YLeaf{"UnregisterMessages", messageStatistics.UnregisterMessages}
    messageStatistics.EntityData.Leafs["echo-submit-messages"] = types.YLeaf{"EchoSubmitMessages", messageStatistics.EchoSubmitMessages}
    messageStatistics.EntityData.Leafs["echo-cancel-messages"] = types.YLeaf{"EchoCancelMessages", messageStatistics.EchoCancelMessages}
    messageStatistics.EntityData.Leafs["get-result-messages"] = types.YLeaf{"GetResultMessages", messageStatistics.GetResultMessages}
    messageStatistics.EntityData.Leafs["get-config-messages"] = types.YLeaf{"GetConfigMessages", messageStatistics.GetConfigMessages}
    messageStatistics.EntityData.Leafs["get-response-messages"] = types.YLeaf{"GetResponseMessages", messageStatistics.GetResponseMessages}
    messageStatistics.EntityData.Leafs["property-response-messages"] = types.YLeaf{"PropertyResponseMessages", messageStatistics.PropertyResponseMessages}
    messageStatistics.EntityData.Leafs["property-request-messages"] = types.YLeaf{"PropertyRequestMessages", messageStatistics.PropertyRequestMessages}
    messageStatistics.EntityData.Leafs["property-block-messages"] = types.YLeaf{"PropertyBlockMessages", messageStatistics.PropertyBlockMessages}
    messageStatistics.EntityData.Leafs["thread-request-messages"] = types.YLeaf{"ThreadRequestMessages", messageStatistics.ThreadRequestMessages}
    return &(messageStatistics.EntityData)
}

// MplsOam_Global_CollaboratorStatistics
// Collaborator statistics
type MplsOam_Global_CollaboratorStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collaborator IPARM counts.
    CollaboratorIParm MplsOam_Global_CollaboratorStatistics_CollaboratorIParm

    // Collaborator IM counts.
    CollaboratorIm MplsOam_Global_CollaboratorStatistics_CollaboratorIm

    // Collaborator NetIO counts.
    CollaboratorNetIo MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo

    // Collaborator RIB counts.
    CollaboratorRib MplsOam_Global_CollaboratorStatistics_CollaboratorRib
}

func (collaboratorStatistics *MplsOam_Global_CollaboratorStatistics) GetEntityData() *types.CommonEntityData {
    collaboratorStatistics.EntityData.YFilter = collaboratorStatistics.YFilter
    collaboratorStatistics.EntityData.YangName = "collaborator-statistics"
    collaboratorStatistics.EntityData.BundleName = "cisco_ios_xr"
    collaboratorStatistics.EntityData.ParentYangName = "global"
    collaboratorStatistics.EntityData.SegmentPath = "collaborator-statistics"
    collaboratorStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collaboratorStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collaboratorStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collaboratorStatistics.EntityData.Children = make(map[string]types.YChild)
    collaboratorStatistics.EntityData.Children["collaborator-i-parm"] = types.YChild{"CollaboratorIParm", &collaboratorStatistics.CollaboratorIParm}
    collaboratorStatistics.EntityData.Children["collaborator-im"] = types.YChild{"CollaboratorIm", &collaboratorStatistics.CollaboratorIm}
    collaboratorStatistics.EntityData.Children["collaborator-net-io"] = types.YChild{"CollaboratorNetIo", &collaboratorStatistics.CollaboratorNetIo}
    collaboratorStatistics.EntityData.Children["collaborator-rib"] = types.YChild{"CollaboratorRib", &collaboratorStatistics.CollaboratorRib}
    collaboratorStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(collaboratorStatistics.EntityData)
}

// MplsOam_Global_CollaboratorStatistics_CollaboratorIParm
// Collaborator IPARM counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorIParm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorIParm *MplsOam_Global_CollaboratorStatistics_CollaboratorIParm) GetEntityData() *types.CommonEntityData {
    collaboratorIParm.EntityData.YFilter = collaboratorIParm.YFilter
    collaboratorIParm.EntityData.YangName = "collaborator-i-parm"
    collaboratorIParm.EntityData.BundleName = "cisco_ios_xr"
    collaboratorIParm.EntityData.ParentYangName = "collaborator-statistics"
    collaboratorIParm.EntityData.SegmentPath = "collaborator-i-parm"
    collaboratorIParm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collaboratorIParm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collaboratorIParm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collaboratorIParm.EntityData.Children = make(map[string]types.YChild)
    collaboratorIParm.EntityData.Leafs = make(map[string]types.YLeaf)
    collaboratorIParm.EntityData.Leafs["ups"] = types.YLeaf{"Ups", collaboratorIParm.Ups}
    collaboratorIParm.EntityData.Leafs["downs"] = types.YLeaf{"Downs", collaboratorIParm.Downs}
    return &(collaboratorIParm.EntityData)
}

// MplsOam_Global_CollaboratorStatistics_CollaboratorIm
// Collaborator IM counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorIm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorIm *MplsOam_Global_CollaboratorStatistics_CollaboratorIm) GetEntityData() *types.CommonEntityData {
    collaboratorIm.EntityData.YFilter = collaboratorIm.YFilter
    collaboratorIm.EntityData.YangName = "collaborator-im"
    collaboratorIm.EntityData.BundleName = "cisco_ios_xr"
    collaboratorIm.EntityData.ParentYangName = "collaborator-statistics"
    collaboratorIm.EntityData.SegmentPath = "collaborator-im"
    collaboratorIm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collaboratorIm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collaboratorIm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collaboratorIm.EntityData.Children = make(map[string]types.YChild)
    collaboratorIm.EntityData.Leafs = make(map[string]types.YLeaf)
    collaboratorIm.EntityData.Leafs["ups"] = types.YLeaf{"Ups", collaboratorIm.Ups}
    collaboratorIm.EntityData.Leafs["downs"] = types.YLeaf{"Downs", collaboratorIm.Downs}
    return &(collaboratorIm.EntityData)
}

// MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo
// Collaborator NetIO counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorNetIo *MplsOam_Global_CollaboratorStatistics_CollaboratorNetIo) GetEntityData() *types.CommonEntityData {
    collaboratorNetIo.EntityData.YFilter = collaboratorNetIo.YFilter
    collaboratorNetIo.EntityData.YangName = "collaborator-net-io"
    collaboratorNetIo.EntityData.BundleName = "cisco_ios_xr"
    collaboratorNetIo.EntityData.ParentYangName = "collaborator-statistics"
    collaboratorNetIo.EntityData.SegmentPath = "collaborator-net-io"
    collaboratorNetIo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collaboratorNetIo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collaboratorNetIo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collaboratorNetIo.EntityData.Children = make(map[string]types.YChild)
    collaboratorNetIo.EntityData.Leafs = make(map[string]types.YLeaf)
    collaboratorNetIo.EntityData.Leafs["ups"] = types.YLeaf{"Ups", collaboratorNetIo.Ups}
    collaboratorNetIo.EntityData.Leafs["downs"] = types.YLeaf{"Downs", collaboratorNetIo.Downs}
    return &(collaboratorNetIo.EntityData)
}

// MplsOam_Global_CollaboratorStatistics_CollaboratorRib
// Collaborator RIB counts
type MplsOam_Global_CollaboratorStatistics_CollaboratorRib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collaborator up counter. The type is interface{} with range: 0..4294967295.
    Ups interface{}

    // Collaborator down counter. The type is interface{} with range:
    // 0..4294967295.
    Downs interface{}
}

func (collaboratorRib *MplsOam_Global_CollaboratorStatistics_CollaboratorRib) GetEntityData() *types.CommonEntityData {
    collaboratorRib.EntityData.YFilter = collaboratorRib.YFilter
    collaboratorRib.EntityData.YangName = "collaborator-rib"
    collaboratorRib.EntityData.BundleName = "cisco_ios_xr"
    collaboratorRib.EntityData.ParentYangName = "collaborator-statistics"
    collaboratorRib.EntityData.SegmentPath = "collaborator-rib"
    collaboratorRib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collaboratorRib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collaboratorRib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collaboratorRib.EntityData.Children = make(map[string]types.YChild)
    collaboratorRib.EntityData.Leafs = make(map[string]types.YLeaf)
    collaboratorRib.EntityData.Leafs["ups"] = types.YLeaf{"Ups", collaboratorRib.Ups}
    collaboratorRib.EntityData.Leafs["downs"] = types.YLeaf{"Downs", collaboratorRib.Downs}
    return &(collaboratorRib.EntityData)
}

