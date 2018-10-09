// This module contains a collection of YANG definitions
// for PPP and PPPoE operational data.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
// All rights reserved.
package ppp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ppp-oper ppp-data}", reflect.TypeOf(PppData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ppp-oper:ppp-data", reflect.TypeOf(PppData{}))
}

// PppIosAuthType represents Authentication type for the PPP session
type PppIosAuthType string

const (
    // No authentication.
    PppIosAuthType_ppp_ios_auth_none PppIosAuthType = "ppp-ios-auth-none"

    // Challenge Handshake Authentication (CHAP).
    PppIosAuthType_ppp_ios_auth_chap PppIosAuthType = "ppp-ios-auth-chap"

    // Password Authentication Protocol (PAP).
    PppIosAuthType_ppp_ios_auth_pap PppIosAuthType = "ppp-ios-auth-pap"

    // Microsoft CHAP (MS CHAP).
    PppIosAuthType_ppp_ios_auth_mschap PppIosAuthType = "ppp-ios-auth-mschap"

    // Microsoft CHAP, Version 2 (MS CHAP V2).
    PppIosAuthType_ppp_ios_auth_mschap_v2 PppIosAuthType = "ppp-ios-auth-mschap-v2"

    // Extensible Authentication Protocol (EAP).
    PppIosAuthType_ppp_ios_auth_eap PppIosAuthType = "ppp-ios-auth-eap"
)

// PppoeOperationalRole represents The role that this PPPoE session is in
type PppoeOperationalRole string

const (
    // PPPoE device role is a client
    PppoeOperationalRole_pppoe_client PppoeOperationalRole = "pppoe-client"

    // PPPoE device role is a server
    PppoeOperationalRole_pppoe_server PppoeOperationalRole = "pppoe-server"
)

// PppData
// Operational state of PPP
type PppData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of the PPP information. The type is slice of PppData_PppInterface.
    PppInterface []*PppData_PppInterface

    // The PPP statistics.
    PppStatistics PppData_PppStatistics

    // The PPPoE operation information.
    Pppoe PppData_Pppoe
}

func (pppData *PppData) GetEntityData() *types.CommonEntityData {
    pppData.EntityData.YFilter = pppData.YFilter
    pppData.EntityData.YangName = "ppp-data"
    pppData.EntityData.BundleName = "cisco_ios_xe"
    pppData.EntityData.ParentYangName = "Cisco-IOS-XE-ppp-oper"
    pppData.EntityData.SegmentPath = "Cisco-IOS-XE-ppp-oper:ppp-data"
    pppData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppData.EntityData.Children = types.NewOrderedMap()
    pppData.EntityData.Children.Append("ppp-interface", types.YChild{"PppInterface", nil})
    for i := range pppData.PppInterface {
        pppData.EntityData.Children.Append(types.GetSegmentPath(pppData.PppInterface[i]), types.YChild{"PppInterface", pppData.PppInterface[i]})
    }
    pppData.EntityData.Children.Append("ppp-statistics", types.YChild{"PppStatistics", &pppData.PppStatistics})
    pppData.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &pppData.Pppoe})
    pppData.EntityData.Leafs = types.NewOrderedMap()

    pppData.EntityData.YListKeys = []string {}

    return &(pppData.EntityData)
}

// PppData_PppInterface
// A list of the PPP information
type PppData_PppInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ifname of Physical Access (Parent) Interface . The
    // type is string.
    PhyIfname interface{}

    // List of PPPoE sessions on ifname Physical access interface. The type is
    // slice of PppData_PppInterface_PppVa.
    PppVa []*PppData_PppInterface_PppVa
}

func (pppInterface *PppData_PppInterface) GetEntityData() *types.CommonEntityData {
    pppInterface.EntityData.YFilter = pppInterface.YFilter
    pppInterface.EntityData.YangName = "ppp-interface"
    pppInterface.EntityData.BundleName = "cisco_ios_xe"
    pppInterface.EntityData.ParentYangName = "ppp-data"
    pppInterface.EntityData.SegmentPath = "ppp-interface" + types.AddKeyToken(pppInterface.PhyIfname, "phy-ifname")
    pppInterface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppInterface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppInterface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppInterface.EntityData.Children = types.NewOrderedMap()
    pppInterface.EntityData.Children.Append("ppp-va", types.YChild{"PppVa", nil})
    for i := range pppInterface.PppVa {
        pppInterface.EntityData.Children.Append(types.GetSegmentPath(pppInterface.PppVa[i]), types.YChild{"PppVa", pppInterface.PppVa[i]})
    }
    pppInterface.EntityData.Leafs = types.NewOrderedMap()
    pppInterface.EntityData.Leafs.Append("phy-ifname", types.YLeaf{"PhyIfname", pppInterface.PhyIfname})

    pppInterface.EntityData.YListKeys = []string {"PhyIfname"}

    return &(pppInterface.EntityData)
}

// PppData_PppInterface_PppVa
// List of PPPoE sessions on ifname Physical access interface
type PppData_PppInterface_PppVa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPP Virtual Access Interface Name. The type is string.
    VaIfname interface{}

    // VRF of Virtual Access Interface. The type is string.
    VrfName interface{}

    // IP Address assigned/negotiated by PPP. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    InterfaceIp interface{}

    // Gateway IP Address assigned/negotiated by PPP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GatewayIp interface{}

    // Primary DNS IP Address assigned/negotiated by PPP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrimaryDnsIp interface{}

    // Secondary DNS IP Address assigned/negotiated by PPP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SecondaryDnsIp interface{}

    // MTU for PPP interface. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Authentication type for PPP interface. The type is PppIosAuthType.
    AuthType interface{}
}

func (pppVa *PppData_PppInterface_PppVa) GetEntityData() *types.CommonEntityData {
    pppVa.EntityData.YFilter = pppVa.YFilter
    pppVa.EntityData.YangName = "ppp-va"
    pppVa.EntityData.BundleName = "cisco_ios_xe"
    pppVa.EntityData.ParentYangName = "ppp-interface"
    pppVa.EntityData.SegmentPath = "ppp-va"
    pppVa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppVa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppVa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppVa.EntityData.Children = types.NewOrderedMap()
    pppVa.EntityData.Leafs = types.NewOrderedMap()
    pppVa.EntityData.Leafs.Append("va-ifname", types.YLeaf{"VaIfname", pppVa.VaIfname})
    pppVa.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", pppVa.VrfName})
    pppVa.EntityData.Leafs.Append("interface-ip", types.YLeaf{"InterfaceIp", pppVa.InterfaceIp})
    pppVa.EntityData.Leafs.Append("gateway-ip", types.YLeaf{"GatewayIp", pppVa.GatewayIp})
    pppVa.EntityData.Leafs.Append("primary-dns-ip", types.YLeaf{"PrimaryDnsIp", pppVa.PrimaryDnsIp})
    pppVa.EntityData.Leafs.Append("secondary-dns-ip", types.YLeaf{"SecondaryDnsIp", pppVa.SecondaryDnsIp})
    pppVa.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", pppVa.Mtu})
    pppVa.EntityData.Leafs.Append("auth-type", types.YLeaf{"AuthType", pppVa.AuthType})

    pppVa.EntityData.YListKeys = []string {}

    return &(pppVa.EntityData)
}

// PppData_PppStatistics
// The PPP statistics
// This type is a presence type.
type PppData_PppStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // number of PPP LCP pkts. The type is interface{} with range: 0..4294967295.
    PppLcpPkts interface{}

    // number of PPP NCP/IPCP pkts. The type is interface{} with range:
    // 0..4294967295.
    PppIpcpPkts interface{}

    // number of PPP CCP pkts. The type is interface{} with range: 0..4294967295.
    PppCcpPkts interface{}
}

func (pppStatistics *PppData_PppStatistics) GetEntityData() *types.CommonEntityData {
    pppStatistics.EntityData.YFilter = pppStatistics.YFilter
    pppStatistics.EntityData.YangName = "ppp-statistics"
    pppStatistics.EntityData.BundleName = "cisco_ios_xe"
    pppStatistics.EntityData.ParentYangName = "ppp-data"
    pppStatistics.EntityData.SegmentPath = "ppp-statistics"
    pppStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppStatistics.EntityData.Children = types.NewOrderedMap()
    pppStatistics.EntityData.Leafs = types.NewOrderedMap()
    pppStatistics.EntityData.Leafs.Append("ppp-lcp-pkts", types.YLeaf{"PppLcpPkts", pppStatistics.PppLcpPkts})
    pppStatistics.EntityData.Leafs.Append("ppp-ipcp-pkts", types.YLeaf{"PppIpcpPkts", pppStatistics.PppIpcpPkts})
    pppStatistics.EntityData.Leafs.Append("ppp-ccp-pkts", types.YLeaf{"PppCcpPkts", pppStatistics.PppCcpPkts})

    pppStatistics.EntityData.YListKeys = []string {}

    return &(pppStatistics.EntityData)
}

// PppData_Pppoe
// The PPPoE operation information
// This type is a presence type.
type PppData_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The current PPPoE role. The type is PppoeOperationalRole.
    Role interface{}

    // PPPoE session list. The type is slice of PppData_Pppoe_PppoeSessionList.
    PppoeSessionList []*PppData_Pppoe_PppoeSessionList

    // PPPoE statistics.
    PppoeStatistics PppData_Pppoe_PppoeStatistics
}

func (pppoe *PppData_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xe"
    pppoe.EntityData.ParentYangName = "ppp-data"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Children.Append("pppoe-session-list", types.YChild{"PppoeSessionList", nil})
    for i := range pppoe.PppoeSessionList {
        pppoe.EntityData.Children.Append(types.GetSegmentPath(pppoe.PppoeSessionList[i]), types.YChild{"PppoeSessionList", pppoe.PppoeSessionList[i]})
    }
    pppoe.EntityData.Children.Append("pppoe-statistics", types.YChild{"PppoeStatistics", &pppoe.PppoeStatistics})
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("role", types.YLeaf{"Role", pppoe.Role})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// PppData_Pppoe_PppoeSessionList
// PPPoE session list
type PppData_Pppoe_PppoeSessionList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ifname of Physical Access (Parent) Interface. The
    // type is string.
    Ifname interface{}

    // List of PPPoE session on ifname Physical access interface. The type is
    // slice of PppData_Pppoe_PppoeSessionList_Session.
    Session []*PppData_Pppoe_PppoeSessionList_Session
}

func (pppoeSessionList *PppData_Pppoe_PppoeSessionList) GetEntityData() *types.CommonEntityData {
    pppoeSessionList.EntityData.YFilter = pppoeSessionList.YFilter
    pppoeSessionList.EntityData.YangName = "pppoe-session-list"
    pppoeSessionList.EntityData.BundleName = "cisco_ios_xe"
    pppoeSessionList.EntityData.ParentYangName = "pppoe"
    pppoeSessionList.EntityData.SegmentPath = "pppoe-session-list" + types.AddKeyToken(pppoeSessionList.Ifname, "ifname")
    pppoeSessionList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppoeSessionList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppoeSessionList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppoeSessionList.EntityData.Children = types.NewOrderedMap()
    pppoeSessionList.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range pppoeSessionList.Session {
        pppoeSessionList.EntityData.Children.Append(types.GetSegmentPath(pppoeSessionList.Session[i]), types.YChild{"Session", pppoeSessionList.Session[i]})
    }
    pppoeSessionList.EntityData.Leafs = types.NewOrderedMap()
    pppoeSessionList.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", pppoeSessionList.Ifname})

    pppoeSessionList.EntityData.YListKeys = []string {"Ifname"}

    return &(pppoeSessionList.EntityData)
}

// PppData_Pppoe_PppoeSessionList_Session
// List of PPPoE session on ifname Physical access interface
type PppData_Pppoe_PppoeSessionList_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Id of PPPoE sessions. The type is interface{} with range: 0..65535.
    SessionId interface{}

    // Local MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Lmac interface{}

    // Peer MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Rmac interface{}

    // Ifname of Virtual Access Interface. The type is string.
    VaIfname interface{}

    // VRF of Virtual Access Interface . The type is string.
    VrfName interface{}

    // PPPoE session VLAN/QinQ ID  Outer VLAN (QinQ) or only VLAN ID. The type is
    // interface{} with range: 0..65535.
    Dot1qQinqOuterVlan interface{}

    // PPPoE session VLAN/QinQ ID  Inner VLAN ID (QinQ) (if valid). The type is
    // interface{} with range: 0..65535.
    Dot1qQinqInnerVlan interface{}

    // PPPoE service tag name. The type is string.
    ServiceName interface{}

    // Number of packets received over session. The type is interface{} with
    // range: 0..4294967295.
    InPacket interface{}

    // Number of packets sent over session. The type is interface{} with range:
    // 0..4294967295.
    OutPacket interface{}

    // Number of bytes received over session. The type is interface{} with range:
    // 0..18446744073709551615.
    InBytes interface{}

    // Number of bytes sent over session. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBytes interface{}
}

func (session *PppData_Pppoe_PppoeSessionList_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xe"
    session.EntityData.ParentYangName = "pppoe-session-list"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("lmac", types.YLeaf{"Lmac", session.Lmac})
    session.EntityData.Leafs.Append("rmac", types.YLeaf{"Rmac", session.Rmac})
    session.EntityData.Leafs.Append("va-ifname", types.YLeaf{"VaIfname", session.VaIfname})
    session.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", session.VrfName})
    session.EntityData.Leafs.Append("dot1q-qinq-outer-vlan", types.YLeaf{"Dot1qQinqOuterVlan", session.Dot1qQinqOuterVlan})
    session.EntityData.Leafs.Append("dot1q-qinq-inner-vlan", types.YLeaf{"Dot1qQinqInnerVlan", session.Dot1qQinqInnerVlan})
    session.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", session.ServiceName})
    session.EntityData.Leafs.Append("in-packet", types.YLeaf{"InPacket", session.InPacket})
    session.EntityData.Leafs.Append("out-packet", types.YLeaf{"OutPacket", session.OutPacket})
    session.EntityData.Leafs.Append("in-bytes", types.YLeaf{"InBytes", session.InBytes})
    session.EntityData.Leafs.Append("out-bytes", types.YLeaf{"OutBytes", session.OutBytes})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// PppData_Pppoe_PppoeStatistics
// PPPoE statistics
// This type is a presence type.
type PppData_Pppoe_PppoeStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // PPPoE active discovery initiation pkts. The type is interface{} with range:
    // 0..4294967295.
    PppoePadiPkts interface{}

    // PPPoE active discovery offer pkts. The type is interface{} with range:
    // 0..4294967295.
    PppoePadoPkts interface{}

    // PPPoE active discovery response pkts. The type is interface{} with range:
    // 0..4294967295.
    PppoePadrPkts interface{}

    // PPPoE active discovery session pkts. The type is interface{} with range:
    // 0..4294967295.
    PppoePadsPkts interface{}

    // PPPoE active discovery terminate pkts. The type is interface{} with range:
    // 0..4294967295.
    PppoePadtPkts interface{}

    // PPPoE invalid discovery pkts. The type is interface{} with range:
    // 0..4294967295.
    InvalidDiscoveryPkts interface{}
}

func (pppoeStatistics *PppData_Pppoe_PppoeStatistics) GetEntityData() *types.CommonEntityData {
    pppoeStatistics.EntityData.YFilter = pppoeStatistics.YFilter
    pppoeStatistics.EntityData.YangName = "pppoe-statistics"
    pppoeStatistics.EntityData.BundleName = "cisco_ios_xe"
    pppoeStatistics.EntityData.ParentYangName = "pppoe"
    pppoeStatistics.EntityData.SegmentPath = "pppoe-statistics"
    pppoeStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pppoeStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pppoeStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pppoeStatistics.EntityData.Children = types.NewOrderedMap()
    pppoeStatistics.EntityData.Leafs = types.NewOrderedMap()
    pppoeStatistics.EntityData.Leafs.Append("pppoe-padi-pkts", types.YLeaf{"PppoePadiPkts", pppoeStatistics.PppoePadiPkts})
    pppoeStatistics.EntityData.Leafs.Append("pppoe-pado-pkts", types.YLeaf{"PppoePadoPkts", pppoeStatistics.PppoePadoPkts})
    pppoeStatistics.EntityData.Leafs.Append("pppoe-padr-pkts", types.YLeaf{"PppoePadrPkts", pppoeStatistics.PppoePadrPkts})
    pppoeStatistics.EntityData.Leafs.Append("pppoe-pads-pkts", types.YLeaf{"PppoePadsPkts", pppoeStatistics.PppoePadsPkts})
    pppoeStatistics.EntityData.Leafs.Append("pppoe-padt-pkts", types.YLeaf{"PppoePadtPkts", pppoeStatistics.PppoePadtPkts})
    pppoeStatistics.EntityData.Leafs.Append("invalid-discovery-pkts", types.YLeaf{"InvalidDiscoveryPkts", pppoeStatistics.InvalidDiscoveryPkts})

    pppoeStatistics.EntityData.YListKeys = []string {}

    return &(pppoeStatistics.EntityData)
}

