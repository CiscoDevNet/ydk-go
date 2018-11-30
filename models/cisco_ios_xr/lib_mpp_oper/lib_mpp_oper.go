// This module contains a collection of YANG definitions
// for Cisco IOS-XR lib-mpp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   management-plane-protection: Management Plane Protection (MPP)
//     operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package lib_mpp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lib_mpp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lib-mpp-oper management-plane-protection}", reflect.TypeOf(ManagementPlaneProtection{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lib-mpp-oper:management-plane-protection", reflect.TypeOf(ManagementPlaneProtection{}))
}

type Ipv4 struct {
}

func (id Ipv4) String() string {
	return "Cisco-IOS-XR-lib-mpp-oper-sub1:ipv4"
}

type MppAfIdBase struct {
}

func (id MppAfIdBase) String() string {
	return "Cisco-IOS-XR-lib-mpp-oper-sub1:Mpp-af-id-base"
}

type Ipv6 struct {
}

func (id Ipv6) String() string {
	return "Cisco-IOS-XR-lib-mpp-oper-sub1:ipv6"
}

// MppAllow represents MPP protocol types
type MppAllow string

const (
    // SSH protocol
    MppAllow_ssh MppAllow = "ssh"

    // TELNET protocol
    MppAllow_telnet MppAllow = "telnet"

    // SNMP protocol
    MppAllow_snmp MppAllow = "snmp"

    // TFTP protocol
    MppAllow_tftp MppAllow = "tftp"

    // HTTP protocol
    MppAllow_http MppAllow = "http"

    // XML
    MppAllow_xr_xml MppAllow = "xr-xml"

    // NETCONF protocol
    MppAllow_netconf MppAllow = "netconf"

    // All
    MppAllow_all MppAllow = "all"
)

// ManagementPlaneProtection
// Management Plane Protection (MPP) operational
// data
type ManagementPlaneProtection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Management Plane Protection (MPP) outband interface data.
    Outband ManagementPlaneProtection_Outband

    // Management Plane Protection (MPP) inband interface data.
    Inband ManagementPlaneProtection_Inband
}

func (managementPlaneProtection *ManagementPlaneProtection) GetEntityData() *types.CommonEntityData {
    managementPlaneProtection.EntityData.YFilter = managementPlaneProtection.YFilter
    managementPlaneProtection.EntityData.YangName = "management-plane-protection"
    managementPlaneProtection.EntityData.BundleName = "cisco_ios_xr"
    managementPlaneProtection.EntityData.ParentYangName = "Cisco-IOS-XR-lib-mpp-oper"
    managementPlaneProtection.EntityData.SegmentPath = "Cisco-IOS-XR-lib-mpp-oper:management-plane-protection"
    managementPlaneProtection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    managementPlaneProtection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    managementPlaneProtection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    managementPlaneProtection.EntityData.Children = types.NewOrderedMap()
    managementPlaneProtection.EntityData.Children.Append("outband", types.YChild{"Outband", &managementPlaneProtection.Outband})
    managementPlaneProtection.EntityData.Children.Append("inband", types.YChild{"Inband", &managementPlaneProtection.Inband})
    managementPlaneProtection.EntityData.Leafs = types.NewOrderedMap()

    managementPlaneProtection.EntityData.YListKeys = []string {}

    return &(managementPlaneProtection.EntityData)
}

// ManagementPlaneProtection_Outband
// Management Plane Protection (MPP) outband
// interface data
type ManagementPlaneProtection_Outband struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outband VRF information.
    Vrf ManagementPlaneProtection_Outband_Vrf

    // List of inband/outband interfaces.
    Interfaces ManagementPlaneProtection_Outband_Interfaces
}

func (outband *ManagementPlaneProtection_Outband) GetEntityData() *types.CommonEntityData {
    outband.EntityData.YFilter = outband.YFilter
    outband.EntityData.YangName = "outband"
    outband.EntityData.BundleName = "cisco_ios_xr"
    outband.EntityData.ParentYangName = "management-plane-protection"
    outband.EntityData.SegmentPath = "outband"
    outband.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outband.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outband.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outband.EntityData.Children = types.NewOrderedMap()
    outband.EntityData.Children.Append("vrf", types.YChild{"Vrf", &outband.Vrf})
    outband.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &outband.Interfaces})
    outband.EntityData.Leafs = types.NewOrderedMap()

    outband.EntityData.YListKeys = []string {}

    return &(outband.EntityData)
}

// ManagementPlaneProtection_Outband_Vrf
// Outband VRF information
type ManagementPlaneProtection_Outband_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outband VRF name. The type is string.
    VrfName interface{}
}

func (vrf *ManagementPlaneProtection_Outband_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "outband"
    vrf.EntityData.SegmentPath = "vrf"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {}

    return &(vrf.EntityData)
}

// ManagementPlaneProtection_Outband_Interfaces
// List of inband/outband interfaces
type ManagementPlaneProtection_Outband_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPP interface information. The type is slice of
    // ManagementPlaneProtection_Outband_Interfaces_Interface.
    Interface []*ManagementPlaneProtection_Outband_Interfaces_Interface
}

func (interfaces *ManagementPlaneProtection_Outband_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "outband"
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

// ManagementPlaneProtection_Outband_Interfaces_Interface
// MPP interface information
type ManagementPlaneProtection_Outband_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name, specify 'all' for all interfaces.
    // The type is string.
    InterfaceName interface{}

    // MPP Interface protocols. The type is slice of
    // ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol.
    Protocol []*ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol
}

func (self *ManagementPlaneProtection_Outband_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range self.Protocol {
        self.EntityData.Children.Append(types.GetSegmentPath(self.Protocol[i]), types.YChild{"Protocol", self.Protocol[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol
// MPP Interface protocols
type ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPP allow. The type is MppAllow.
    Allow interface{}

    // If TRUE, all peers are allowed. The type is bool.
    IsAllPeersAllowed interface{}

    // List of peer addresses. The type is slice of
    // ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress.
    PeerAddress []*ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress
}

func (protocol *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "interface"
    protocol.EntityData.SegmentPath = "protocol"
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Children.Append("peer-address", types.YChild{"PeerAddress", nil})
    for i := range protocol.PeerAddress {
        protocol.EntityData.Children.Append(types.GetSegmentPath(protocol.PeerAddress[i]), types.YChild{"PeerAddress", protocol.PeerAddress[i]})
    }
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("allow", types.YLeaf{"Allow", protocol.Allow})
    protocol.EntityData.Leafs.Append("is-all-peers-allowed", types.YLeaf{"IsAllPeersAllowed", protocol.IsAllPeersAllowed})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

// ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress
// List of peer addresses
type ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (peerAddress *ManagementPlaneProtection_Outband_Interfaces_Interface_Protocol_PeerAddress) GetEntityData() *types.CommonEntityData {
    peerAddress.EntityData.YFilter = peerAddress.YFilter
    peerAddress.EntityData.YangName = "peer-address"
    peerAddress.EntityData.BundleName = "cisco_ios_xr"
    peerAddress.EntityData.ParentYangName = "protocol"
    peerAddress.EntityData.SegmentPath = "peer-address"
    peerAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddress.EntityData.Children = types.NewOrderedMap()
    peerAddress.EntityData.Leafs = types.NewOrderedMap()
    peerAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddress.AfName})
    peerAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", peerAddress.Ipv4Address})
    peerAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", peerAddress.Ipv6Address})

    peerAddress.EntityData.YListKeys = []string {}

    return &(peerAddress.EntityData)
}

// ManagementPlaneProtection_Inband
// Management Plane Protection (MPP) inband
// interface data
type ManagementPlaneProtection_Inband struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of inband/outband interfaces.
    Interfaces ManagementPlaneProtection_Inband_Interfaces
}

func (inband *ManagementPlaneProtection_Inband) GetEntityData() *types.CommonEntityData {
    inband.EntityData.YFilter = inband.YFilter
    inband.EntityData.YangName = "inband"
    inband.EntityData.BundleName = "cisco_ios_xr"
    inband.EntityData.ParentYangName = "management-plane-protection"
    inband.EntityData.SegmentPath = "inband"
    inband.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inband.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inband.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inband.EntityData.Children = types.NewOrderedMap()
    inband.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &inband.Interfaces})
    inband.EntityData.Leafs = types.NewOrderedMap()

    inband.EntityData.YListKeys = []string {}

    return &(inband.EntityData)
}

// ManagementPlaneProtection_Inband_Interfaces
// List of inband/outband interfaces
type ManagementPlaneProtection_Inband_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPP interface information. The type is slice of
    // ManagementPlaneProtection_Inband_Interfaces_Interface.
    Interface []*ManagementPlaneProtection_Inband_Interfaces_Interface
}

func (interfaces *ManagementPlaneProtection_Inband_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "inband"
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

// ManagementPlaneProtection_Inband_Interfaces_Interface
// MPP interface information
type ManagementPlaneProtection_Inband_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name, specify 'all' for all interfaces.
    // The type is string.
    InterfaceName interface{}

    // MPP Interface protocols. The type is slice of
    // ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol.
    Protocol []*ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol
}

func (self *ManagementPlaneProtection_Inband_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range self.Protocol {
        self.EntityData.Children.Append(types.GetSegmentPath(self.Protocol[i]), types.YChild{"Protocol", self.Protocol[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol
// MPP Interface protocols
type ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPP allow. The type is MppAllow.
    Allow interface{}

    // If TRUE, all peers are allowed. The type is bool.
    IsAllPeersAllowed interface{}

    // List of peer addresses. The type is slice of
    // ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress.
    PeerAddress []*ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress
}

func (protocol *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "interface"
    protocol.EntityData.SegmentPath = "protocol"
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Children.Append("peer-address", types.YChild{"PeerAddress", nil})
    for i := range protocol.PeerAddress {
        protocol.EntityData.Children.Append(types.GetSegmentPath(protocol.PeerAddress[i]), types.YChild{"PeerAddress", protocol.PeerAddress[i]})
    }
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("allow", types.YLeaf{"Allow", protocol.Allow})
    protocol.EntityData.Leafs.Append("is-all-peers-allowed", types.YLeaf{"IsAllPeersAllowed", protocol.IsAllPeersAllowed})

    protocol.EntityData.YListKeys = []string {}

    return &(protocol.EntityData)
}

// ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress
// List of peer addresses
type ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is one of the following: Ipv4Ipv6.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (peerAddress *ManagementPlaneProtection_Inband_Interfaces_Interface_Protocol_PeerAddress) GetEntityData() *types.CommonEntityData {
    peerAddress.EntityData.YFilter = peerAddress.YFilter
    peerAddress.EntityData.YangName = "peer-address"
    peerAddress.EntityData.BundleName = "cisco_ios_xr"
    peerAddress.EntityData.ParentYangName = "protocol"
    peerAddress.EntityData.SegmentPath = "peer-address"
    peerAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddress.EntityData.Children = types.NewOrderedMap()
    peerAddress.EntityData.Leafs = types.NewOrderedMap()
    peerAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddress.AfName})
    peerAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", peerAddress.Ipv4Address})
    peerAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", peerAddress.Ipv6Address})

    peerAddress.EntityData.YListKeys = []string {}

    return &(peerAddress.EntityData)
}

